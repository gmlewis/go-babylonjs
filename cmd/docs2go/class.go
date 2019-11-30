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
)

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
							v.parseParameters(processConstructorOverrides)
							html.ConstructorNames[v.Name()] = v
						case "Methods":
							v.parseParameters(nil)
							html.MethodNames[v.Name()] = v
						default:
							log.Fatalf("%v: unknown ConstructorsAndMethods: h2=%v", filename, section.H2)
						}
					}
					for _, v := range section.Properties {
						if !strings.Contains(v.Class, "tsd-signature") {
							continue
						}
						v.parseParameters(nil)
						html.PropertyNames[v.Name()] = v
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

type Div struct {
	ID       string     `xml:"id,attr"`
	H1       InnerXML   `xml:"header>div>div>h1"`
	Sections []*Section `xml:"div>div>div>section"`
}

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

type Signature struct {
	Class    string `xml:"class,attr"`
	InnerXML string `xml:",innerxml"`

	GoParams []string `xml:"-"`
	GoOpts   []string `xml:"-"`
	HasOpts  bool     `xml:"-"`

	JSParams []string `xml:"-"`
}

func (s *Signature) Name() string {
	v := strings.TrimSpace(s.InnerXML)
	if i := strings.Index(v, "<span"); i >= 0 {
		v = v[0:i]
	}
	v = strings.Replace(v, "<wbr>", "", -1)
	v = strings.Replace(v, " ", "", -1)
	if len(v) > 0 {
		v = strings.ToUpper(v[0:1]) + v[1:]
	}
	return v
}

type processOverrider func(s *Signature, names []string, optional []bool, types []string) ([]string, []bool, []string)

func processConstructorOverrides(s *Signature, names []string, optional []bool, types []string) ([]string, []bool, []string) {
	override := true
	switch s.Name() {
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
		logf("\n\nOVERRIDES: parseParameters[%v]: names(%v)=%v, optional(%v)=%v, types(%v)=%v\n\n", s.Name(), len(names), names, len(optional), optional, len(types), types)
	}
	return names, optional, types
}

func (s *Signature) parseParameters(processOverrides processOverrider) {
	v := strings.TrimSpace(s.InnerXML)
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
	for im, m := range matches {
		index := strings.Index(v, m[0])
		if index < 0 {
			log.Fatalf("parseParameters: programming error: im=%v, i=%v, v=%v", im, index, v)
		}
		keyword := v[0:index]
		v = v[index+len(m[0]):]
		logf("SEARCHING for '%v'\nm[1]=`%v`\nm[2]=`%v`\nignoreNextType=%v\nAFTER removing '%v': '%v'\n\n", m[0], m[1], m[2], ignoreNextType, keyword, v)

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
			}
		case `span class="tsd-signature-type"`:
			if ignoreNextType {
				ignoreNextType = false
			} else {
				if i := len(types); i > 0 && types[i-1] == "Array" {
					types[i-1] = "[]" + m[2]
					// } else if i > 0 && types[i-1] == "*Nullable" {
					// 	types[i-1] = "*" + m[2] // Override last type
				} else {
					switch m[2] {
					case "Nullable", "DeepImmutable", "Partial": // ignore and pick up next class
					case "boolean":
						types = append(types, "bool")
					case "number":
						types = append(types, "float64")
					case "TCamera": // special case
						types = append(types, "*Camera")
					default:
						types = append(types, m[2])
					}
				}
			}
		}

		if strings.HasPrefix(m[1], `a href="`) {
			if ignoreNextType {
				ignoreNextType = false
			} else {
				if i := len(types); i > 0 && types[i-1] == "Array" {
					types[i-1] = "[]" + m[2]
					// } else if i > 0 && types[i-1] == "*Nullable" {
					// 	types[i-1] = "*" + m[2] // Override last type
				} else {
					switch m[2] {
					case "Nullable", "DeepImmutable", "Partial": // ignore and pick up next class
					case "DataArray":
						types = append(types, "[]float64")
					default:
						types = append(types, "*"+m[2])
					}
				}
			}
		}
		logf("names(%v)=%v, optional(%v)=%v, types(%v)=%v", len(names), names, len(optional), optional, len(types), types)
	}

	logf("\n\nparseParameters[%v]: %v\n%#v\nnames(%v)=%v, optional(%v)=%v, types(%v)=%v\n\n", s.Name(), v, matches, len(names), names, len(optional), optional, len(types), types)

	if len(names) != len(optional) || len(names) > len(types) {
		log.Fatalf("badly parsed arguments: \n\nparseParameters[%v]: %v\n%#v\nnames(%v)=%v, optional(%v)=%v, types(%v)=%v\n\n", s.Name(), v, matches, len(names), names, len(optional), optional, len(types), types)
	}

	if processOverrides != nil {
		names, optional, types = processOverrides(s, names, optional, types)
	}

	for i, v := range names {
		paramType := "interface{}" // unknown type

		if optional[i] {
			name := strings.Title(v)
			jsName := name

			if i < len(types) {
				paramType = types[i]
				switch paramType {
				case "any":
					paramType = "interface{}"
				case "function":
					paramType = "func()"
				case "[]AbstractMesh", "[]Worker":
					paramType = "[]js.Value"
				case "string":
					paramType = "JSString"
					jsName += ".JSObject()"
				case "float64":
					paramType = "JSFloat64"
					jsName += ".JSObject()"
				case "bool":
					paramType = "JSBool"
					jsName += ".JSObject()"
				case "*DistanceJointData",
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
					"*InternalTextureSource",
					"*IOceanPostProcessOptions",
					"*IPhysicsEnabledObject",
					"*IShaderMaterialOptions",
					"*IShadowLight",
					"*ISoundOptions",
					"*ISoundTrackOptions",
					"*ISpriteManager",
					"*MeshLoadOptions",
					"*NodeMaterialBlockConnectionPointTypes",
					"*NodeMaterialBlockTargets",
					"*NodeMaterialConnectionPointDirection",
					"*PhysicsImpostorParameters",
					"*PhysicsJointData",
					"*SceneOptions",
					"*SpringJointData",
					"*TonemappingOperator",
					"*VideoRecorderOptions",
					"*VideoTextureSettings",
					"*VRExperienceHelperOptions",
					"*WebVROptions",
					"*WebXRState",
					"*XRInputSource",
					"*XRReferenceSpaceType",
					"*XRSessionMode",
					"ArrayBufferView",
					"ClipboardEvent",
					"HTMLCanvasElement",
					"HTMLElement",
					"HTMLVideoElement",
					"IPhysicsEnginePlugin",
					"KeyboardEvent",
					"PointerEvent",
					"WebGLRenderingContext",
					"Worker",
					"null",
					"object":
					paramType = "JSValue"
					jsName += ".JSObject()"
				default:
					jsName += ".JSObject()"
				}
				if !strings.HasPrefix(paramType, "*") && paramType != "js.Value" && !strings.HasPrefix(paramType, "[]") {
					paramType = "*" + paramType
				}
			}
			s.GoOpts = append(s.GoOpts, fmt.Sprintf("%v %v", name, paramType))
			s.HasOpts = true
			s.JSParams = append(s.JSParams, "opts."+jsName)
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
				switch paramType {
				case "any":
					paramType = "interface{}"
				case "function":
					paramType = "func()"
				case "[]PostProcess",
					"[]Mesh",
					"[]Worker":
					paramType = "[]js.Value"
				case "string", "float64", "bool", "[]string", "[]float64", "[]bool":
				case
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
					"*InternalTextureSource",
					"*IOceanPostProcessOptions",
					"*IPhysicsEnabledObject",
					"*IShaderMaterialOptions",
					"*IShadowLight",
					"*ISoundOptions",
					"*ISoundTrackOptions",
					"*ISpriteManager",
					"*MeshLoadOptions",
					"*NodeMaterialBlockConnectionPointTypes",
					"*NodeMaterialBlockTargets",
					"*NodeMaterialConnectionPointDirection",
					"*PhysicsImpostorParameters",
					"*PhysicsJointData",
					"*SceneOptions",
					"*SpringJointData",
					"*TonemappingOperator",
					"*VideoRecorderOptions",
					"*VideoTextureSettings",
					"*VRExperienceHelperOptions",
					"*WebVROptions",
					"*WebXRState",
					"*XRInputSource",
					"*XRReferenceSpaceType",
					"*XRSessionMode",
					"ArrayBufferView",
					"ClipboardEvent",
					"HTMLCanvasElement",
					"HTMLElement",
					"HTMLVideoElement",
					"IPhysicsEnginePlugin",
					"KeyboardEvent",
					"PointerEvent",
					"WebGLRenderingContext",
					"Worker",
					"null",
					"object":
					paramType = "js.Value"
				default:
					jsName += ".JSObject()"
				}
			}
			s.GoParams = append(s.GoParams, fmt.Sprintf("%v %v", name, paramType))
			s.JSParams = append(s.JSParams, jsName)
		}
	}

	logf("parseParameters[%v]: final params: HasOpts=%v, GoParams=%#v, GoOpts=%#v, JSParams=%#v\n\n", s.Name(), s.HasOpts, s.GoParams, s.GoOpts, s.JSParams)
}

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

type InnerXML struct {
	InnerXML string `xml:",innerxml"`
}
