package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	scriptRE     = regexp.MustCompile(`<script>.*?</script>`)
	htmlRE       = regexp.MustCompile(`<([^/].*?)>(.*?)</.*?>`)
	whitespaceRE = regexp.MustCompile(`\s+`)

	unhandledType = map[string]bool{}
)

func init() {
	for _, v := range unhandledTypes {
		unhandledType[v] = true
	}
}

type classes struct {
	m map[string]*ClassHTML
}

func (c *classes) walker() filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		parentDir := filepath.Base(filepath.Dir(path))
		if parentDir != "classes" {
			return nil
		}

		filename := filepath.Base(path)
		if !strings.HasPrefix(filename, "babylon.") {
			logf("Unknown filename %v, skipping", path)
			return nil
		}
		if strings.HasPrefix(filename, "babylon._") {
			logf("Skipping private class %v", path)
			return nil
		}

		if *verbose != "" {
			if !strings.Contains(filename, "."+strings.ToLower(*verbose)+".") {
				return nil
			}
		}

		logf("\n\nProcessing %v ...", filename)

		buf, err := ioutil.ReadFile(path)
		check("ReadFile: %v", err)

		// Strip problematic <script> sections that the XML parser fails on.
		buf = scriptRE.ReplaceAll(buf, nil)

		d := xml.NewDecoder(bytes.NewReader(buf))
		d.Strict = false
		d.AutoClose = xml.HTMLAutoClose
		d.Entity = xml.HTMLEntity

		html := &ClassHTML{
			ConstructorNames: map[string]*Signature{},
			MethodNames:      map[string]*Signature{},
			PropertyNames:    map[string]*Signature{},
		}
		check("xml.Decode: %v", d.Decode(&html))

		for _, div := range html.Div {
			if div.ID != "wrapper" {
				continue
			}
			html.Name = div.H1.InnerXML
			if i := strings.Index(html.Name, "&lt;"); i >= 0 {
				html.Name = html.Name[0:i]
			}

			for _, section := range div.Sections {
				switch t := section.Type(); t {
				case "tsd-comment":
					lines := strings.Split(section.Lead.InnerXML, "\n")
					for i, v := range lines {
						lines[i] = strings.TrimSpace(v)
					}
					html.Summary = strings.Join(lines, "\n// ")

					lines = strings.Split(section.Description.InnerXML, "\n")
					for i, v := range lines {
						lines[i] = strings.TrimSpace(v)
					}
					html.Description = strings.Join(lines, "\n// ")

					html.SeeURL = section.SeeURL.InnerXML
				case "tsd-hierarchy":
					for _, v := range section.ListItems {
						html.Parents = append(html.Parents, v.InnerXML)
					}
					for _, v := range section.Children {
						html.Children = append(html.Children, v.InnerXML)
					}
				case "Implements":
					for _, v := range section.ListItems {
						html.Implements = append(html.Implements, v.InnerXML)
					}
				case "tsd-index-group":
					// for _, indexSection := range section.IndexSections {
					// 	logf("%v: indexSection=%v: %v", t, indexSection.Type(), indexSection.ListItems)
					// }
				case "tsd-member-group", "tsd-is-not-exported", "tsd-is-inherited": // Parse this for Constructors, Properties, and Methods
					for _, v := range section.ConstructorsAndMethods {
						if !strings.Contains(v.Class, "tsd-signature") {
							continue
						}
						switch section.H2.InnerXML {
						case "Constructors":
							if ok := v.parseParameters(html.Name, processConstructorOverrides); ok {
								html.ConstructorNames[v.GoName] = v
							}
						case "Methods":
							if ok := v.parseParameters(html.Name, processMethodOverrides); ok {
								html.MethodNames[v.GoName] = v
							}
						default:
							log.Fatalf("%v: unknown ConstructorsAndMethods: h2=%v", filename, section.H2)
						}
					}
					for _, v := range section.Properties {
						if !strings.Contains(v.Class, "tsd-signature") {
							continue
						}
						if ok := v.parseParameters(html.Name, nil); ok {
							html.PropertyNames[v.GoName] = v
						}
					}
				case "tsd-type-parameters": // ignore
				case "tsd-parent-kind-module": // ignore
				default:
					log.Fatalf("Unknown section type: %v", t)
				}
			}
		}

		logf("html.Name=%v", html.Name)
		logf("html.Summary=%v", html.Summary)
		logf("html.Description=%v", html.Description)
		logf("html.SeeURL=%v", html.SeeURL)
		logf("html.Parents=%v", html.Parents)
		logf("html.Children=%v", html.Children)
		logf("html.Implements=%v", html.Implements)
		logf("html.ConstructorNames=%v", html.ConstructorNames)
		logf("html.MethodNames=%v", html.MethodNames)
		logf("html.PropertyNames=%v", html.PropertyNames)

		c.m[html.Name] = html

		return nil
	}
}

// ClassHTML represents a JavaScript class.
type ClassHTML struct {
	Title InnerXML `xml:"head>title"`
	Div   []*Div   `xml:"body>div"`

	Name             string                `xml:"-"`
	Summary          string                `xml:"-"`
	Description      string                `xml:"-"`
	SeeURL           string                `xml:"-"`
	Parents          []string              `xml:"-"`
	Children         []string              `xml:"-"`
	Implements       []string              `xml:"-"`
	ConstructorNames map[string]*Signature `xml:"-"`
	MethodNames      map[string]*Signature `xml:"-"`
	PropertyNames    map[string]*Signature `xml:"-"`
}

// Div represents an HTML div
type Div struct {
	ID       string     `xml:"id,attr"`
	H1       InnerXML   `xml:"header>div>div>h1"`
	Sections []*Section `xml:"div>div>div>section"`
}

// Section represents an HTML section.
type Section struct {
	Class string   `xml:"class,attr"`
	H2    InnerXML `xml:"h2"`
	H3    InnerXML `xml:"h3"`

	// class: tsd-comment
	Lead        InnerXML `xml:"div>div>p"`
	Description InnerXML `xml:"div>p"`
	SeeURL      InnerXML `xml:"div>dl>dd>p>a"`

	// class: tsd-hierarchy
	// class Implements
	ListItems []InnerXML `xml:"ul>li>a"`

	// class: tsd-hierarchy
	Children []InnerXML `xml:"ul>li>ul>li>ul>li>a"`

	// class: tsd-index-group
	// IndexSections []*Section `xml:"section>div>section"`

	// class: tsd-is-not-exported
	ConstructorsAndMethods []*Signature `xml:"section>ul>li"`
	Properties             []*Signature `xml:"section>div"`
}

// Signature represents the signature of a constructor, method, or property.
type Signature struct {
	Class    string `xml:"class,attr"`
	InnerXML string `xml:",innerxml"`

	GoName string `xml:"-"`
	JSName string `xml:"-"`

	GoReturnType      string `xml:"-"`
	GoReturnStatement string `xml:"-"`

	GoParams   []string `xml:"-"`
	GoOpts     []string `xml:"-"`
	GoOptsName []string `xml:"-"`
	HasOpts    bool     `xml:"-"`

	JSParams []string `xml:"-"`
	JSOpts   []string `xml:"-"`
}

type processOverrider func(className string, s *Signature, names []string, optional []bool, types []string) ([]string, []bool, []string)

func processConstructorOverrides(className string, s *Signature, names []string, optional []bool, types []string) ([]string, []bool, []string) {
	override := true
	switch s.GoName {
	case "NewColor3":
		optional = []bool{false, false, false}
	case "NewVector2":
		optional = []bool{false, false}
	case "NewVector3":
		optional = []bool{false, false, false}
	default:
		override = false
	}
	if override {
		logf("\n\nOVERRIDES: parseParameters[%v]: names(%v)=%v, optional(%v)=%v, types(%v)=%v\n\n", s.GoName, len(names), names, len(optional), optional, len(types), types)
	}
	return names, optional, types
}

func processMethodOverrides(className string, s *Signature, names []string, optional []bool, types []string) ([]string, []bool, []string) {
	override := true
	switch className + "." + s.GoName {
	case "Angle.BetweenTwoPoints":
		names[0] = "av"
	default:
		override = false
	}
	if override {
		logf("\n\nOVERRIDES: parseParameters[%v]: names(%v)=%v, optional(%v)=%v, types(%v)=%v\n\n", s.GoName, len(names), names, len(optional), optional, len(types), types)
	}
	return names, optional, types
}

func (s *Signature) parseParameters(className string, processOverrides processOverrider) bool {
	// First, populate the GoName and JSName fields.
	v := strings.TrimSpace(s.InnerXML)
	if i := strings.Index(v, "<span"); i >= 0 {
		v = v[0:i]
	}
	v = strings.Replace(v, "<wbr>", "", -1)
	v = strings.Replace(v, " ", "", -1)
	if len(v) > 0 {
		s.JSName = v
		s.GoName = strings.ToUpper(v[0:1]) + v[1:]
	} else {
		log.Fatalf("parseParameters: unable to determine name of %v", s.InnerXML)
	}
	if strings.Contains(s.GoName, "&") {
		log.Printf("Skipping %v.%v", className, s.GoName)
		return false
	}

	v = strings.TrimSpace(s.InnerXML)
	logf("ORIGINAL: '%v'\n\n", v)
	v = strings.Replace(v, "<wbr>", "", -1)
	v = strings.Replace(v, "</li>", "", -1)
	v = strings.Replace(v, "</div>", "", -1)
	v = whitespaceRE.ReplaceAllString(v, " ")
	matches := htmlRE.FindAllStringSubmatch(v, -1)

	logf("BEFORE: '%v'\n\n", v)
	var names []string
	var optional []bool
	var types []string
	var ignoreNextType bool
	var lastWasLessThan bool
	for im, m := range matches {
		index := strings.Index(v, m[0])
		if index < 0 {
			log.Fatalf("parseParameters: programming error: im=%v, i=%v, v=%v", im, index, v)
		}
		keyword := v[0:index]
		v = v[index+len(m[0]):]
		logf("SEARCHING for '%v'\nm[1]=`%v`\nm[2]=`%v`\nignoreNextType=%v\nAFTER removing '%v': '%v'\n\n", m[0], m[1], m[2], ignoreNextType, keyword, v)

		addType := func() {
			if ignoreNextType {
				ignoreNextType = false
			} else {
				if i := len(types); i > 0 && types[i-1] == "Array" {
					types[i-1] = "[]" + m[2]
					// } else if i > 0 && types[i-1] == "*Nullable" {
					// 	types[i-1] = "*" + m[2] // Override last type
				} else {
					if lastWasLessThan && len(types) > 0 {
						switch types[len(types)-1] {
						case "Nullable", "DeepImmutable", "Partial": // replace property
							types = types[0 : len(types)-1]
						}
					}
					switch m[2] {
					case "boolean":
						types = append(types, "bool")
					case "number":
						types = append(types, "float64")
					case "TCamera": // special case
						types = append(types, "*Camera")
					case "DataArray":
						types = append(types, "[]float64")
					default:
						types = append(types, m[2])
					}
				}
				lastWasLessThan = false
			}
		}

		switch m[1] {
		case `span class="tsd-signature-symbol"`:
			keyword = strings.TrimPrefix(keyword, ", ")
			if keyword != "" {
				if strings.HasPrefix(m[2], "?") {
					optional = append(optional, true)
					names = append(names, keyword)
				} else if strings.HasPrefix(m[2], ":") {
					optional = append(optional, false)
					names = append(names, keyword)
				}
			} else {
				if strings.TrimSpace(m[2]) == "|" { // keep the first type, ignore the rest.
					ignoreNextType = true
				}
				if strings.Contains(m[2], "&lt;") {
					lastWasLessThan = true
				}
			}
		case `span class="tsd-signature-type"`:
			addType()
		}

		if strings.HasPrefix(m[1], `a href="`) {
			addType()
		}
		logf("names(%v)=%v, optional(%v)=%v, types(%v)=%v", len(names), names, len(optional), optional, len(types), types)
	}

	logf("\n\nparseParameters[%v]: %v\n%#v\nnames(%v)=%v, optional(%v)=%v, types(%v)=%v\n\n", s.GoName, v, matches, len(names), names, len(optional), optional, len(types), types)

	if len(names) != len(optional) || len(names) > len(types) {
		log.Fatalf("badly parsed arguments: \n\nparseParameters[%v]: %v\n%#v\nnames(%v)=%v, optional(%v)=%v, types(%v)=%v\n\n", s.GoName, v, matches, len(names), names, len(optional), optional, len(types), types)
	}

	if processOverrides != nil {
		names, optional, types = processOverrides(className, s, names, optional, types)
	}

	for i, v := range names {
		paramType := "interface{}" // unknown type

		if optional[i] {
			name := strings.Title(v)
			jsName := "opts." + name

			if i < len(types) {
				paramType = types[i]
				if unhandledType[paramType] {
					paramType = "js.Value"
				} else {
					switch paramType {
					case "any":
						paramType = "interface{}"
					case "function":
						paramType = "func()"
					case "[]AbstractMesh", "[]IAnimationKey", "[]Worker":
						paramType = "[]js.Value"
					case "string":
						paramType = "*string"
						jsName = "*" + jsName
					case "float64":
						paramType = "*float64"
						jsName = "*" + jsName
					case "bool":
						paramType = "*bool"
						jsName = "*" + jsName
					default:
						jsName += ".JSObject()"
					}
				}
				if !strings.HasPrefix(paramType, "*") && paramType != "js.Value" && !strings.HasPrefix(paramType, "[]") {
					paramType = "*" + paramType
				}
			}
			s.GoOptsName = append(s.GoOptsName, name)
			s.GoOpts = append(s.GoOpts, fmt.Sprintf("%v %v", name, paramType))
			s.HasOpts = true
			s.JSOpts = append(s.JSOpts, jsName)
		} else {
			name := v
			switch name {
			case "type":
				name = "jsType"
			case "func":
				name = "jsFunc"
			}
			jsName := name
			if i < len(types) {
				paramType = types[i]
				var needsJSObject bool
				paramType, needsJSObject = jsTypeToGoType(paramType)
				if needsJSObject {
					if !strings.HasPrefix(paramType, "*") && paramType != "js.Value" && !strings.HasPrefix(paramType, "[]") {
						paramType = "*" + paramType
					}
					jsName += ".JSObject()"
				}
			}
			s.GoParams = append(s.GoParams, fmt.Sprintf("%v %v", name, paramType))
			s.JSParams = append(s.JSParams, jsName)
		}
	}

	if len(types) > len(names) {
		var needsJSObject bool
		s.GoReturnType, needsJSObject = jsTypeToGoType(types[len(types)-1])

		if s.GoReturnType == "this" {
			s.GoReturnType = className
			needsJSObject = true
		}

		if needsJSObject {
			s.GoReturnStatement = fmt.Sprintf("%vFromJSObject(retVal, %v.ctx)", s.GoReturnType, receiver(className))
			s.GoReturnType = "*" + s.GoReturnType
		} else {
			switch s.GoReturnType {
			case "bool":
				s.GoReturnStatement = "retVal.Bool()"
			case "float64":
				s.GoReturnStatement = "retVal.Float()"
			case "string":
				s.GoReturnStatement = "retVal.String()"
			default:
				s.GoReturnStatement = "retVal"
			}
		}
	}

	logf("parseParameters[%v]: final params: HasOpts=%v, GoParams=%#v, GoOptsName=%#v, GoOpts=%#v, JSParams=%#v, JSOpts=%#v, GoReturnType=%v\n\n", s.GoName, s.HasOpts, s.GoParams, s.GoOptsName, s.GoOpts, s.JSParams, s.JSOpts, s.GoReturnType)
	return true
}

func jsTypeToGoType(paramType string) (goType string, needsJSObject bool) {
	if unhandledType[paramType] {
		return "js.Value", false
	}

	switch paramType {
	case "any":
		paramType = "interface{}"
	case "function":
		paramType = "func()"
	case
		"[]IAnimationKey",
		"[]Mesh",
		"[]PostProcess",
		"[]Worker":
		paramType = "[]js.Value"
	case "string", "float64", "bool", "[]string", "[]float64", "[]bool":
	case "void":
		paramType = ""
	case "Boolean":
		paramType = "bool"
	default:
		needsJSObject = true
	}

	return paramType, needsJSObject
}

// Type returns the JavaScript type of the class in this section.
func (s *Section) Type() string {
	classes := strings.Split(strings.TrimSpace(s.Class), " ")
	class := classes[len(classes)-1]
	if class == "tsd-panel" {
		if s.H2.InnerXML != "" {
			return s.H2.InnerXML
		}
		if s.H3.InnerXML != "" {
			return s.H3.InnerXML
		}
	}
	return class
}

// InnerXML represents the inner text of an XML block.
type InnerXML struct {
	InnerXML string `xml:",innerxml"`
}

var unhandledTypes = []string{
	"*DistanceJointData",
	"*EffectWrapperCreationOptions",
	"*EngineOptions",
	"*IDataBuffer",
	"*IEffectFallbacks",
	"*IEffectRendererOptions",
	"*IEnvironmentHelperOptions",
	"*IGlowLayerOptions",
	"*IHighlightLayerOptions",
	"*IHtmlElementTextureOptions",
	"*IMultiRenderTargetOptions",
	"*INodeMaterialOptions",
	"*IOceanPostProcessOptions",
	"*IPhysicsEnabledObject",
	"*IShaderMaterialOptions",
	"*IShadowLight",
	"*ISoundOptions",
	"*ISoundTrackOptions",
	"*ISpriteManager",
	"*InternalTextureSource",
	"*MeshLoadOptions",
	"*NodeMaterialBlockConnectionPointTypes",
	"*NodeMaterialBlockTargets",
	"*NodeMaterialConnectionPointDirection",
	"*PhysicsImpostorParameters",
	"*PhysicsJointData",
	"*SceneOptions",
	"*SpringJointData",
	"*TonemappingOperator",
	"*VRExperienceHelperOptions",
	"*VideoRecorderOptions",
	"*VideoTextureSettings",
	"*WebVROptions",
	"*WebXRState",
	"*XRInputSource",
	"*XRReferenceSpaceType",
	"*XRSessionMode",
	"ArrayBuffer",
	"ArrayBufferView",
	"ArrayLike",
	"AudioNode",
	"BabylonFileParser",
	"Behavior",
	"CanvasRenderingContext2D",
	"ClipboardEvent",
	"CubeMapInfo",
	"DDSInfo",
	"DeepImmutable",
	"DevicePose",
	"Event",
	"Float32Array",
	"FloatArray",
	"HTMLCanvasElement",
	"HTMLElement",
	"HTMLImageElement",
	"HTMLVideoElement",
	"IAction",
	"IActionEvent",
	"IAnimatable",
	"IAnimationKey",
	"IArrayItem",
	"ICameraInput",
	"ICullable",
	"IDataBuffer",
	"IEasingFunction",
	"IFocusableControl",
	"IInspectorOptions",
	"IMaterialCompilationOptions",
	"IMotorEnabledJoint",
	"IPhysicsEnginePlugin",
	"ISize",
	"IndicesArray",
	"IndividualBabylonFileParser",
	"KeyboardEvent",
	"NodeConstructor",
	"NodeMaterialBlockConnectionPointTypes",
	"NodeMaterialBlockTargets",
	"NodeMaterialDefines",
	"PhysicsImpostorJoint",
	"PointerEvent",
	"Space",
	"StandardMaterialDefines",
	"TrianglePickingPredicate",
	"Uint8Array",
	"WebGLRenderingContext",
	"Worker",
	"null",
	"object",
}
