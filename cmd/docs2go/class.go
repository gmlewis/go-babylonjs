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
		if parentDir != "classes" && parentDir != "interfaces" {
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

		logf("Processing %v/%v ...", parentDir, filename)

		buf, err := ioutil.ReadFile(path)
		check("ReadFile: %v", err)

		// Strip problematic <script> sections that the XML parser fails on.
		buf = scriptRE.ReplaceAll(buf, nil)

		d := xml.NewDecoder(bytes.NewReader(buf))
		d.Strict = false
		d.AutoClose = xml.HTMLAutoClose
		d.Entity = xml.HTMLEntity

		html := &ClassHTML{
			ConstructorNamespaceReceiverName: "ba",       // default
			ConstructorNamespaceReceiverType: "*Babylon", // default

			ConstructorNames: map[string]*Signature{},
			MethodNames:      map[string]*Signature{},
			PropertyNames:    map[string]*Signature{},
		}
		if strings.Contains(filename, "babylon.gui") {
			html.ConstructorNamespaceReceiverName = "gui"
			html.ConstructorNamespaceReceiverType = "*GUI"
		}

		check("xml.Decode: %v", d.Decode(&html))

		for _, div := range html.Div {
			if div.ID != "wrapper" {
				continue
			}

			if s := div.GetHeaderDiv(0).GetDiv(0).GetH1().GetInnerXML(); s != "" {
				html.Name = s
				if i := strings.Index(html.Name, "&lt;"); i >= 0 {
					html.Name = html.Name[0:i]
				}
			} else {
				log.Fatalf("%v: unable to find class name.", filename)
			}
			logf("html.Name=%v", html.Name)
			logf("html.ConstructorNamespaceReceiverName=%v", html.ConstructorNamespaceReceiverName)
			logf("html.ConstructorNamespaceReceiverType=%v", html.ConstructorNamespaceReceiverType)

			mainDiv := div.GetDiv(0).GetDiv(0).GetDiv(0)
			if mainDiv == nil {
				log.Fatalf("%v: unable to find class name.", filename)
			}

			commentSection := mainDiv.FindSections(func(s *Section) bool { return s.HasClass("tsd-comment") })
			if len(commentSection) > 0 { // A summary and description are optional.
				if len(commentSection) > 1 {
					log.Fatalf("%v: found %v tsd-comment sections, want 1", filename, len(commentSection))
				}
				if s := commentSection[0].GetDiv(0).GetDiv(0).GetP(0).GetInnerXML(); s != "" {
					lines := strings.Split(s, "\n")
					for i, v := range lines {
						lines[i] = strings.TrimSpace(v)
					}
					html.Summary = strings.Join(lines, "\n// ")
				} else {
					log.Fatalf("%v: unable to find html.Summary.", filename)
				}
				if s := commentSection[0].GetDiv(0).GetP(0).GetInnerXML(); s != "" { // optional
					lines := strings.Split(s, "\n")
					for i, v := range lines {
						lines[i] = strings.TrimSpace(v)
					}
					html.Description = strings.Join(lines, "\n// ")
				}
				var seeLines []string
				for _, dd := range commentSection[0].GetDiv(0).GetDL(0).GetDDs() {
					if s := dd.GetP(0).GetA(0).GetInnerXML(); s != "" { // optional
						seeLines = append(seeLines, s)
					}
				}
				html.SeeURL = strings.Join(seeLines, "\n// See: ")
			}

			hierarchySection := mainDiv.FindSections(func(s *Section) bool { return s.HasClass("tsd-hierarchy") })
			if len(hierarchySection) != 1 {
				log.Fatalf("%v: found %v tsd-hierarchy sections, want 1", filename, len(hierarchySection))
			}
			for _, ul := range hierarchySection[0].ULs {
				if s := ul.GetLI(0).GetA(0).GetInnerXML(); s != "" && s != html.Name { // optional
					html.Parents = append(html.Parents, s)
				}
				if s := ul.GetLI(0).GetUL(0).GetLI(0).GetUL(0).GetLI(0).GetA(0).GetInnerXML(); s != "" { // optional
					html.Children = append(html.Children, s)
				}
			}

			// Special cases...
			switch html.Name {
			case "StandardMaterial":
				html.Parents = append(html.Parents, "Material")
			}

			implementsSection := mainDiv.FindSections(func(s *Section) bool { return s.H3.GetInnerXML() == "Implements" })
			if len(implementsSection) > 0 { // optional
				for _, ul := range implementsSection[0].ULs {
					if s := ul.GetLI(0).GetA(0).GetInnerXML(); s != "" {
						html.Implements = append(html.Implements, s)
					} else if s := ul.GetLI(0).GetSpan(0).GetInnerXML(); s != "" {
						html.Implements = append(html.Implements, s)
					} else {
						log.Fatalf("%v: unable to find Implements.", filename)
					}
				}
			}
		}

		if err := html.parseConstructors(); err != nil {
			log.Fatalf("%v: %v", filename, err)
		}
		if err := html.parseMethods(); err != nil {
			log.Fatalf("%v: %v", filename, err)
		}
		if err := html.parseProperties(); err != nil {
			log.Fatalf("%v: %v", filename, err)
		}

		logf("html.Name=%v", html.Name)
		logf("html.ConstructorNamespaceReceiverName=%v", html.ConstructorNamespaceReceiverName)
		logf("html.ConstructorNamespaceReceiverType=%v", html.ConstructorNamespaceReceiverType)
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

	ConstructorNamespaceReceiverName string `xml:"-"`
	ConstructorNamespaceReceiverType string `xml:"-"`

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
	Class    string          `xml:"class,attr"`
	Classes  map[string]bool `xml:"-"`
	InnerXML string          `xml:",innerxml"`

	HeaderDivs []*Div `xml:"header>div"`

	ID   string  `xml:"id,attr"`
	Name string  `xml:"name,attr"`
	H1   *Header `xml:"h1"`
	H2   *Header `xml:"h2"`
	H3   *Header `xml:"h3"`
	H4   *Header `xml:"h4"`
	H5   *Header `xml:"h5"`

	Divs []*Div `xml:"div"`

	Ps []*P `xml:"p"`

	DLs []*DL `xml:"dl"`
	ULs []*UL `xml:"ul"`

	Sections []*Section `xml:"section"`
}

func (d *Div) GetSignature() *Signature {
	return &Signature{
		ConstructorNamespaceReceiverName: "ba",       // default
		ConstructorNamespaceReceiverType: "*Babylon", // default

		Class:    d.Class,
		Classes:  d.Classes,
		InnerXML: d.InnerXML,
	}
}

type FindSectionFunc func(*Section) bool

func (d *Div) FindSections(f FindSectionFunc) []*Section {
	if d != nil {
		var result []*Section
		for _, section := range d.Sections {
			if f(section) {
				result = append(result, section)
			}
		}
		return result
	}
	return nil
}

func (d *Div) GetHeaderDiv(i int) *Div {
	if d != nil && i < len(d.HeaderDivs) {
		return d.HeaderDivs[i]
	}
	return nil
}

func (d *Div) GetDiv(i int) *Div {
	if d != nil && i < len(d.Divs) {
		return d.Divs[i]
	}
	return nil
}

func (d *Div) GetSection(i int) *Section {
	if d != nil && i < len(d.Sections) {
		return d.Sections[i]
	}
	return nil
}

func (d *Div) GetP(i int) *P {
	if d != nil && i < len(d.Ps) {
		return d.Ps[i]
	}
	return nil
}

func (d *Div) GetDL(i int) *DL {
	if d != nil && i < len(d.DLs) {
		return d.DLs[i]
	}
	return nil
}

func (d *Div) GetH1() *Header {
	if d != nil && d.H1 != nil {
		return d.H1
	}
	return nil
}

func (d *Div) GetInnerXML() string {
	if d != nil {
		return d.InnerXML
	}
	return ""
}

// Section represents an HTML section.
type Section struct {
	Class   string          `xml:"class,attr"`
	Classes map[string]bool `xml:"-"`

	InnerXML string  `xml:",innerxml"`
	H1       *Header `xml:"h1"`
	H2       *Header `xml:"h2"`
	H3       *Header `xml:"h3"`
	H4       *Header `xml:"h4"`
	H5       *Header `xml:"h5"`

	// // class: tsd-comment
	// Lead        InnerXML `xml:"div>div>p"`
	// Description InnerXML `xml:"div>p"`
	// SeeURL      InnerXML `xml:"div>dl>dd>p>a"`
	//
	// // class: tsd-hierarchy
	// // class Implements
	// ListItems []InnerXML `xml:"ul>li>a"`
	//
	// // class: tsd-hierarchy
	// Children []InnerXML `xml:"ul>li>ul>li>ul>li>a"`
	//
	// // class: tsd-index-group
	// // IndexSections []*Section `xml:"section>div>section"`
	//
	// // class: tsd-is-not-exported
	// ConstructorsAndMethods []*Signature `xml:"section>ul>li"`
	// Properties             []*Signature `xml:"section>div"`

	Divs []*Div `xml:"div"`

	ULs []*UL `xml:"ul"`

	Sections []*Section `xml:"section"`
}

func (s *Section) GetSection(i int) *Section {
	if s != nil && i < len(s.Sections) {
		return s.Sections[i]
	}
	return nil
}

func (s *Section) GetH2() *Header {
	if s != nil && s.H2 != nil {
		return s.H2
	}
	return nil
}

func (s *Section) GetDiv(i int) *Div {
	if s != nil && i < len(s.Divs) {
		return s.Divs[i]
	}
	return nil
}

func (s *Section) GetDivs() []*Div {
	if s != nil {
		return s.Divs
	}
	return nil
}

func (s *Section) GetUL(i int) *UL {
	if s != nil && i < len(s.ULs) {
		return s.ULs[i]
	}
	return nil
}

func (s *Section) GetULs() []*UL {
	if s != nil {
		return s.ULs
	}
	return nil
}

// DL represents an HTML dl.
type DL struct {
	Class   string          `xml:"class,attr"`
	Classes map[string]bool `xml:"-"`

	InnerXML string `xml:",innerxml"`

	DTs []*DT `xml:"dt"`
	DDs []*DD `xml:"dd"`
}

func (d *DL) GetDDs() []*DD {
	if d != nil {
		return d.DDs
	}
	return nil
}

// DD represents an HTML dd.
type DD struct {
	Class   string          `xml:"class,attr"`
	Classes map[string]bool `xml:"-"`

	InnerXML string `xml:",innerxml"`

	Ps []*P `xml:"p"`
}

func (d *DD) GetP(i int) *P {
	if d != nil && i < len(d.Ps) {
		return d.Ps[i]
	}
	return nil
}

// DT represents an HTML dt.
type DT struct {
	Class   string          `xml:"class,attr"`
	Classes map[string]bool `xml:"-"`

	InnerXML string `xml:",innerxml"`
}

// UL represents an HTML unsigned list.
type UL struct {
	Class   string          `xml:"class,attr"`
	Classes map[string]bool `xml:"-"`

	InnerXML string `xml:",innerxml"`

	LIs []*LI `xml:"li"`
}

func (u *UL) GetLI(i int) *LI {
	if u != nil && i < len(u.LIs) {
		return u.LIs[i]
	}
	return nil
}

func (u *UL) GetLIs() []*LI {
	if u != nil {
		return u.LIs
	}
	return nil
}

// Span represents an HTML span.
type Span struct {
	Class   string          `xml:"class,attr"`
	Classes map[string]bool `xml:"-"`

	InnerXML string `xml:",innerxml"`
}

func (s *Span) GetInnerXML() string {
	if s != nil {
		return s.InnerXML
	}
	return ""
}

// LI represents an HTML list item.
type LI struct {
	Class   string          `xml:"class,attr"`
	Classes map[string]bool `xml:"-"`

	InnerXML string  `xml:",innerxml"`
	H2       *Header `xml:"h2"`
	H3       *Header `xml:"h3"`
	H4       *Header `xml:"h4"`
	H5       *Header `xml:"h5"`

	As []*Anchor `xml:"a"`

	Divs []*Div `xml:"div"`

	Spans []*Span `xml:"span"`

	ULs []*UL `xml:"ul"`
}

func (l *LI) GetSignature() *Signature {
	return &Signature{
		ConstructorNamespaceReceiverName: "ba",       // default
		ConstructorNamespaceReceiverType: "*Babylon", // default

		Class:    l.Class,
		Classes:  l.Classes,
		InnerXML: l.InnerXML,
	}
}

func (l *LI) GetA(i int) *Anchor {
	if l != nil && i < len(l.As) {
		return l.As[i]
	}
	return nil
}

func (l *LI) GetSpan(i int) *Span {
	if l != nil && i < len(l.Spans) {
		return l.Spans[i]
	}
	return nil
}

func (l *LI) GetUL(i int) *UL {
	if l != nil && i < len(l.ULs) {
		return l.ULs[i]
	}
	return nil
}

// P represents an HTML paragraph.
type P struct {
	Class   string          `xml:"class,attr"`
	Classes map[string]bool `xml:"-"`

	InnerXML string  `xml:",innerxml"`
	H2       *Header `xml:"h2"`
	H3       *Header `xml:"h3"`
	H4       *Header `xml:"h4"`
	H5       *Header `xml:"h5"`

	As []*Anchor `xml:"a"`

	Divs []*Div `xml:"div"`
}

func (p *P) GetA(i int) *Anchor {
	if p != nil && i < len(p.As) {
		return p.As[i]
	}
	return nil
}

func (p *P) GetInnerXML() string {
	if p != nil {
		return p.InnerXML
	}
	return ""
}

// Header represents an HTML header.
type Header struct {
	Class   string          `xml:"class,attr"`
	Classes map[string]bool `xml:"-"`

	InnerXML string `xml:",innerxml"`
}

func (h *Header) GetInnerXML() string {
	if h != nil {
		return h.InnerXML
	}
	return ""
}

// Anchor represents an HTML anchor.
type Anchor struct {
	Class   string          `xml:"class,attr"`
	Classes map[string]bool `xml:"-"`

	HRef     string `xml:"href,attr"`
	Name     string `xml:"name,attr"`
	InnerXML string `xml:",innerxml"`
}

func (a *Anchor) GetInnerXML() string {
	if a != nil {
		return a.InnerXML
	}
	return ""
}

// Signature represents the signature of a constructor, method, or property.
// It can be copied from an LI or DIV.
type Signature struct {
	Class   string
	Classes map[string]bool

	ConstructorNamespaceReceiverName string
	ConstructorNamespaceReceiverType string

	InnerXML string

	GoName string
	JSName string

	GoReturnType      string
	GoReturnStatement string

	HasOpts     bool
	WriteSetter bool

	GoParamsName     []string
	GoParamsType     []string
	NeedsArrayHelper []string
	JSParams         []string

	GoOptsName           []string
	GoOptsType           []string
	OptsNeedsArrayHelper []string
	JSOpts               []string
}

type processOverrider func(className string, s *Signature, names []string, optional []bool, types []string) (bool, []string, []bool, []string)

func processConstructorOverrides(className string, s *Signature, names []string, optional []bool, types []string) (bool, []string, []bool, []string) {
	var avoidUsingOptions bool
	override := true
	switch s.GoName {
	case "NewColor3":
		optional = []bool{false, false, false}
	case "NewColor4":
		optional = []bool{false, false, false, false}
	case "NewDynamicTexture":
		types[1] = "interface{}"
	case "NewObservable":
		avoidUsingOptions = true
	case "NewVector2":
		optional = []bool{false, false}
	case "NewVector3":
		optional = []bool{false, false, false}
	default:
		override = false
	}
	if override {
		logf("\n\nconstructor OVERRIDES: parseParameters[%v]: names(%v)=%v, optional(%v)=%v, types(%v)=%v\n\n", s.GoName, len(names), names, len(optional), optional, len(types), types)
	}
	return avoidUsingOptions, names, optional, types
}

func processMethodOverrides(className string, s *Signature, names []string, optional []bool, types []string) (bool, []string, []bool, []string) {
	var avoidUsingOptions bool
	override := true
	switch className + "." + s.GoName {
	case "Angle.BetweenTwoPoints":
		names[0] = "av"
	case "ArcRotateCamera.AttachControl":
		optional[1] = false
	case "Camera.AttachControl", "FollowCamera.AttachControl", "FreeCamera.AttachControl",
		"AbstractMesh.ToString", "Scene.CreateDefaultEnvironment":
		avoidUsingOptions = true
	case "InstancedMesh.GetFacetNormal", "InstancedMesh.GetFacetNormalToRef", "InstancedMesh.GetFacetPosition", "InstancedMesh.GetFacetPositionToRef":
		names[0] = "index"
	case "PointsCloudSystem.AddSurfacePoints", "PointsCloudSystem.AddVolumePoints":
		names[4] = "numRange"
	case "SceneLoader.ImportMesh":
		types[0] = "interface{}"
	case "Quaternion.Inverse", "Quaternion.InverseToRef":
		names[0] = "v"
	case "Vector3.CheckExtends", "Vector3WithInfo.CheckExtends":
		names[0] = "vec"
	case "Vector3.SetAll", "Vector3WithInfo.SetAll", "Vector4.SetAll":
		names[0] = "f"
	default:
		override = false
	}
	if override {
		logf("\n\nmethod OVERRIDES: parseParameters[%v]: names(%v)=%v, optional(%v)=%v, types(%v)=%v\n\n", s.GoName, len(names), names, len(optional), optional, len(types), types)
	}
	return avoidUsingOptions, names, optional, types
}

func processPropertiesOverrides(className string, s *Signature, names []string, optional []bool, types []string) (bool, []string, []bool, []string) {
	var avoidUsingOptions bool
	s.WriteSetter = true
	override := true
	switch className + "." + s.GoName {
	case "ArcRotateCamera.Position",
		"ArcRotateCamera.Target",
		"BaseSubMesh.Effect",
		"Bone.Position",
		"Bone.Rotation",
		"Bone.RotationQuaternion",
		"IPhysicsEngine.Gravity",
		"LensRenderingPipeline.ChromaticAberration",
		"LensRenderingPipeline.DarkenOutOfFocus",
		"LensRenderingPipeline.EdgeBlur",
		"LensRenderingPipeline.EdgeDistortion",
		"LensRenderingPipeline.GrainAmount",
		"LensRenderingPipeline.HighlightsGain",
		"LensRenderingPipeline.HighlightsThreshold",
		"Light.Intensity",
		"PhysicsEngine.Gravity",
		"PhysicsImpostor.Mass",
		"PhysicsRaycastResult.HitDistance",
		"PointsCloudSystem.Particles",
		"Scene.ClearColor",
		"ShadowGenerator.Darkness",
		"ShadowGenerator.TransparencyShadow",
		"SolidParticleSystem.Particles",
		"Tools.CorsBehavior",
		"TransformNode.AbsolutePosition",
		"VolumetricLightScatteringPostProcess.CustomMeshPosition":
		s.WriteSetter = false
	case "Light.Range":
		names[0] = "r"
	case "Matrix.M", "Matrix2D.M":
		names[0] = "mm"
	case "ReflectionTextureBlock.R":
		names[0] = "rr"
	default:
		override = false
	}
	if override {
		logf("\n\nproperties OVERRIDES: parseParameters[%v]: names(%v)=%v, optional(%v)=%v, types(%v)=%v\n\n", s.GoName, len(names), names, len(optional), optional, len(types), types)
	}
	return avoidUsingOptions, names, optional, types
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
		logf("SEARCHING for '%v'\nm[1]=`%v`\nm[2]=`%v`\nignoreNextType=%v\nlastWasLessThan=%v\nAFTER removing keyword='%v': '%v'\n\n", m[0], m[1], m[2], ignoreNextType, lastWasLessThan, keyword, v)

		addType := func() {
			if ignoreNextType {
				ignoreNextType = false
			} else {
				var arrayPrefix string
				if i := len(types); i > 0 && types[i-1] == "Array" {
					types = types[0 : len(types)-1]
					arrayPrefix = "[]"
					lastWasLessThan = false
				} else if i := len(types); i > 0 && types[i-1] == "[]Array" {
					types = types[0 : len(types)-1]
					arrayPrefix = "[][]"
					lastWasLessThan = false
				} else if lastWasLessThan && len(types) > 0 {
					switch types[len(types)-1] {
					case "Nullable", "DeepImmutable", "Partial": // replace property
						types = types[0 : len(types)-1]
						lastWasLessThan = false
					case "[]Nullable", "[]DeepImmutable", "[]Partial": // replace property
						types = types[0 : len(types)-1]
						arrayPrefix = "[]"
						lastWasLessThan = false
					}
				}

				if !lastWasLessThan {
					switch m[2] {
					case "boolean":
						types = append(types, "bool")
					case "number", "float":
						types = append(types, "float64")
					case "TCamera": // special case
						types = append(types, "Camera")
					case "DataArray":
						types = append(types, "[]float64")
					default:
						types = append(types, m[2])
					}

					if arrayPrefix != "" {
						types[len(types)-1] = arrayPrefix + types[len(types)-1]
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
				if m[2] == "[]" &&
					len(types) > 0 &&
					!unhandledType[types[len(types)-1]] {
					logf("ADDING ARRAY...")
					switch t := types[len(types)-1]; t {
					case "any", "object", "js.Value", "[]js.Value", "function": // ignore
					case "String":
						types[len(types)-1] = "[]string"
					case "string", "float64", "int", "bool",
						"[]string", "[]float64", "[]int", "[]bool":
						types[len(types)-1] = "[]" + t
					default:
						if strings.HasPrefix(t, "[]") {
							types[len(types)-1] = "[]" + t
						} else {
							types[len(types)-1] = "[]*" + t
						}
					}
				}
			}
		case `span class="tsd-signature-type"`:
			addType()
		}

		if strings.HasPrefix(m[1], `a href="babylon.gui.`) {
			addType()
			s.ConstructorNamespaceReceiverName = "gui"
			s.ConstructorNamespaceReceiverType = "*GUI"
		} else if strings.HasPrefix(m[1], `a href="`) {
			addType()
		}
		logf("names(%v)=%v, optional(%v)=%v, types(%v)=%v", len(names), names, len(optional), optional, len(types), types)
	}

	logf("\n\nparseParameters[%v]: %v\n%#v\nnames(%v)=%v, optional(%v)=%v, types(%v)=%v, CNRName=%v, CNRT=%v\n\n", s.GoName, v, matches, len(names), names, len(optional), optional, len(types), types, s.ConstructorNamespaceReceiverName, s.ConstructorNamespaceReceiverType)

	if len(names) != len(optional) || len(names) > len(types) {
		log.Fatalf("badly parsed arguments: \n\nparseParameters[%v]: %v\n%#v\nnames(%v)=%v, optional(%v)=%v, types(%v)=%v\n\n", s.GoName, v, matches, len(names), names, len(optional), optional, len(types), types)
	}

	var avoidUsingOptions bool
	if processOverrides != nil {
		avoidUsingOptions, names, optional, types = processOverrides(className, s, names, optional, types)
	}

	for i, v := range names {
		paramType := "interface{}" // unknown type

		if !avoidUsingOptions && optional[i] {
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
					case "object":
						paramType = "map[string]interface{}"
					case "function", "Function":
						paramType = "JSFunc"
						jsName = fmt.Sprintf(`js.FuncOf(opts.%v) /* never freed! */`, name)
					case
						"[]AbstractMesh",
						"[]IInternalTextureLoader",
						"[]KeyPropertySet",
						"[]PickingInfo",
						"[]Worker":
						paramType = "[]js.Value"
					case "string":
						paramType = "*string"
						jsName = "*" + jsName
					case "int":
						paramType = "*int"
						jsName = "*" + jsName
					case "float64":
						paramType = "*float64"
						jsName = "*" + jsName
					case "bool":
						paramType = "*bool"
						jsName = "*" + jsName
					case "[]float64", "[]string": // leave unmodified, no .JSObject()
					default:
						jsName += ".JSObject()"
					}
				}
				if !isReferenceType(paramType) {
					paramType = "*" + paramType
				}
			}

			if strings.HasPrefix(paramType, "[]*") {
				s.OptsNeedsArrayHelper = append(s.OptsNeedsArrayHelper, fmt.Sprintf("%vArrayToJSArray(opts.%v)", paramType[3:], name))
				logf("paramType=%v, OptsNeedsArrayHelper=%v", paramType, s.OptsNeedsArrayHelper[len(s.OptsNeedsArrayHelper)-1])
			} else if strings.HasPrefix(paramType, "[][]*") {
				s.OptsNeedsArrayHelper = append(s.OptsNeedsArrayHelper, fmt.Sprintf("%vArray2DToJSArray(opts.%v)", paramType[5:], name))
				logf("paramType=%v, OptsNeedsArrayHelper=%v", paramType, s.OptsNeedsArrayHelper[len(s.OptsNeedsArrayHelper)-1])
			} else {
				s.OptsNeedsArrayHelper = append(s.OptsNeedsArrayHelper, "")
			}

			s.GoOptsName = append(s.GoOptsName, name)
			s.GoOptsType = append(s.GoOptsType, paramType)
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
			var needsArrayHelper bool
			if i < len(types) {
				paramType = types[i]
				var needsJSObject bool
				paramType, needsJSObject, needsArrayHelper = jsTypeToGoType(paramType)

				switch paramType {
				case "object":
					paramType = "js.Value"
					needsJSObject = false
				case "func()", "JSFunc": // special case - function callback
					paramType = "JSFunc"
					jsName = fmt.Sprintf(`js.FuncOf(%v)`, name)
					needsJSObject = false
				case "[]*Worker", "[]*object", "[]object":
					paramType = "[]js.Value"
					needsJSObject = false
				}

				if needsJSObject {
					if !isReferenceType(paramType) {
						paramType = "*" + paramType
					}
					jsName += ".JSObject()"
				}
			}
			if needsArrayHelper {
				s.NeedsArrayHelper = append(s.NeedsArrayHelper, fmt.Sprintf("%vArrayToJSArray(%v)", paramType[3:], name))
			} else {
				s.NeedsArrayHelper = append(s.NeedsArrayHelper, "")
			}
			s.GoParamsName = append(s.GoParamsName, name)
			s.GoParamsType = append(s.GoParamsType, paramType)
			s.JSParams = append(s.JSParams, jsName)
		}
	}

	if len(types) > 0 {
		var needsJSObject bool
		s.GoReturnType, needsJSObject, _ = jsTypeToGoType(types[len(types)-1])

		switch s.GoReturnType {
		case "this":
			s.GoReturnType = className
			needsJSObject = true
		case "[]*any", "[]*object", "[]object":
			s.GoReturnType = "js.Value" // "[]interface{}"
			needsJSObject = false
		case "object":
			// s.GoReturnType = "map[string]interface{}"
			s.GoReturnType = "js.Value"
			needsJSObject = false
		case "func()", "JSFunc", "JSObject":
			s.GoReturnType = "js.Value"
			needsJSObject = false
		}

		if strings.HasPrefix(s.GoReturnType, "[]float64") {
			lines := []string{fmt.Sprintf("result := %v{}", s.GoReturnType)}
			lines = append(lines, "for ri := 0; ri < retVal.Length(); ri++ {")
			lines = append(lines, "  result = append(result, retVal.Index(ri).Float())")
			lines = append(lines, "}")
			lines = append(lines, "return result")
			s.GoReturnStatement = strings.Join(lines, "\n")
		} else if strings.HasPrefix(s.GoReturnType, "[]string") {
			lines := []string{fmt.Sprintf("result := %v{}", s.GoReturnType)}
			lines = append(lines, "for ri := 0; ri < retVal.Length(); ri++ {")
			lines = append(lines, "  result = append(result, retVal.Index(ri).String())")
			lines = append(lines, "}")
			lines = append(lines, "return result")
			s.GoReturnStatement = strings.Join(lines, "\n")
		} else if strings.HasPrefix(s.GoReturnType, "[]*") {
			lines := []string{fmt.Sprintf("result := %v{}", s.GoReturnType)}
			lines = append(lines, "for ri := 0; ri < retVal.Length(); ri++ {")
			lines = append(lines, fmt.Sprintf("  result = append(result, %vFromJSObject(retVal.Index(ri), %v.ctx))", s.GoReturnType[3:], receiver(className)))
			lines = append(lines, "}")
			lines = append(lines, "return result")
			s.GoReturnStatement = strings.Join(lines, "\n")
		} else if needsJSObject {
			s.GoReturnStatement = fmt.Sprintf("return %vFromJSObject(retVal, %v.ctx)", s.GoReturnType, receiver(className))
			s.GoReturnType = "*" + s.GoReturnType
		} else {
			if strings.HasPrefix(s.GoReturnType, "[][]") {
				s.GoReturnType = fmt.Sprintf("js.Value /* %v */", s.GoReturnType) // Not yet supported
			}

			switch s.GoReturnType {
			case "bool":
				s.GoReturnStatement = "return retVal.Bool()"
			case "int":
				s.GoReturnStatement = "return retVal.Int()"
			case "float64":
				s.GoReturnStatement = "return retVal.Float()"
			case "string":
				s.GoReturnStatement = "return retVal.String()"
			default:
				s.GoReturnStatement = "return retVal"
			}
		}
	}

	logf("parseParameters[%v]: final params: JSName=%v, HasOpts=%v, GoParamsName=%#v, GoParamsType=%#v, GoOptsName=%#v, GoOptsType=%#v, JSParams=%#v, JSOpts=%#v, GoReturnType=%v, GoReturnStatement=%v\n\n", s.GoName, s.JSName, s.HasOpts, s.GoParamsName, s.GoParamsType, s.GoOptsName, s.GoOptsType, s.JSParams, s.JSOpts, s.GoReturnType, s.GoReturnStatement)
	return true
}

func isReferenceType(s string) bool {
	return strings.HasPrefix(s, "*") || s == "js.Value" || strings.HasPrefix(s, "[]") || strings.HasPrefix(s, "map[") || strings.HasPrefix(s, "func(") || s == "JSFunc" || s == "JSObject" || s == "interface{}"
}

func jsTypeToGoType(paramType string) (goType string, needsJSObject, needsArrayHelper bool) {
	if unhandledType[paramType] {
		return "js.Value", false, false
	}

	switch paramType {
	case "any":
		paramType = "JSObject"
		needsJSObject = true
	case "interface{}": // needsJSObject=false
	case "function", "Function":
		paramType = "JSFunc"
	case "string", "float64", "bool", "[]string", "[]float64", "[]bool":
	case "[]*string":
		paramType = "[]string"
	case "[]*float64":
		paramType = "[]float64"
	case "[]*bool":
		paramType = "[]bool"
	case "void":
		paramType = ""
	case "String":
		paramType = "string"
	case "Boolean":
		paramType = "bool"
	case "*int", "int":
		paramType = "int"
	default:
		needsJSObject = true
	}

	if strings.HasPrefix(paramType, "[]*") {
		return paramType, false, true
	}

	if strings.HasPrefix(paramType, "[]") && !strings.HasPrefix(paramType, "[][]") && paramType != "[]string" && paramType != "[]float64" && paramType != "[]bool" {
		return "[]*" + paramType[2:], false, true
	}

	if strings.HasPrefix(paramType, "[]") {
		needsJSObject = false
	}

	return paramType, needsJSObject, false
}

// HasClass checks if this HTML element has the named class.
func (el *Div) HasClass(className string) bool {
	if el.Classes != nil {
		return el.Classes[className]
	}

	el.Classes = map[string]bool{}
	classes := strings.Split(strings.TrimSpace(el.Class), " ")
	for _, class := range classes {
		el.Classes[strings.ToLower(class)] = true
	}
	return el.Classes[className]
}

// HasClass checks if this HTML element has the named class.
func (el *Section) HasClass(className string) bool {
	if el == nil {
		log.Printf("HasClass operated on null Section!")
		return false
	}

	if el.Classes != nil {
		return el.Classes[className]
	}

	el.Classes = map[string]bool{}
	classes := strings.Split(strings.TrimSpace(el.Class), " ")
	for _, class := range classes {
		el.Classes[strings.ToLower(class)] = true
	}
	return el.Classes[className]
}

// HasClass checks if this HTML element has the named class.
func (el *UL) HasClass(className string) bool {
	if el.Classes != nil {
		return el.Classes[className]
	}

	el.Classes = map[string]bool{}
	classes := strings.Split(strings.TrimSpace(el.Class), " ")
	for _, class := range classes {
		el.Classes[strings.ToLower(class)] = true
	}
	return el.Classes[className]
}

// HasClass checks if this HTML element has the named class.
func (el *LI) HasClass(className string) bool {
	if el.Classes != nil {
		return el.Classes[className]
	}

	el.Classes = map[string]bool{}
	classes := strings.Split(strings.TrimSpace(el.Class), " ")
	for _, class := range classes {
		el.Classes[strings.ToLower(class)] = true
	}
	return el.Classes[className]
}

// HasClass checks if this HTML element has the named class.
func (el *Header) HasClass(className string) bool {
	if el.Classes != nil {
		return el.Classes[className]
	}

	el.Classes = map[string]bool{}
	classes := strings.Split(strings.TrimSpace(el.Class), " ")
	for _, class := range classes {
		el.Classes[strings.ToLower(class)] = true
	}
	return el.Classes[className]
}

// HasClass checks if this HTML element has the named class.
func (el *Anchor) HasClass(className string) bool {
	if el.Classes != nil {
		return el.Classes[className]
	}

	el.Classes = map[string]bool{}
	classes := strings.Split(strings.TrimSpace(el.Class), " ")
	for _, class := range classes {
		el.Classes[strings.ToLower(class)] = true
	}
	return el.Classes[className]
}

// HasClass checks if this HTML element has the named class.
func (el *Signature) HasClass(className string) bool {
	if el.Classes != nil {
		return el.Classes[className]
	}

	el.Classes = map[string]bool{}
	classes := strings.Split(strings.TrimSpace(el.Class), " ")
	for _, class := range classes {
		el.Classes[strings.ToLower(class)] = true
	}
	return el.Classes[className]
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
	"AccessorComponentType",
	"AlphaState",
	"AnimatedInputBlockTypes",
	"AnimationChannelTargetPath",
	"AnimationKeyInterpolation",
	"ArrayBuffer",
	"ArrayBufferView",
	"ArrayLike",
	"AssetTaskState",
	"AudioBuffer",
	"AudioContext",
	"AudioNode",
	"BabylonFileParser",
	"Behavior",
	"Blob",
	"CanvasRenderingContext2D",
	"ClientRect",
	"ClipboardEvent",
	"Collider",
	"CubeMapInfo",
	"DDSInfo",
	"DeepImmutable",
	"DepthCullingState",
	"DepthSortedParticle",
	"DevicePose",
	"DistanceJointData",
	"Document",
	"EXT_disjoint_timer_query",
	"EXT_texture_filter_anisotropic",
	"EffectWrapperCreationOptions",
	"EngineCapabilities",
	"EngineOptions",
	"EngineOptions",
	"EnvironmentTextureInfo",
	"Event",
	"File",
	"Float32Array",
	"FloatArray",
	"GLTFLoaderAnimationStartMode",
	"GLTFLoaderCoordinateSystemMode",
	"GLTFLoaderState",
	"GainNode",
	"GetDOMTextContent",
	"HDRInfo",
	"HTMLButtonElement",
	"HTMLCanvasElement",
	"HTMLElement",
	"HTMLImageElement",
	"HTMLVideoElement",
	"ICollisionCoordinator",
	"ICollisionCoordinatorFromJSObject",
	"IColor3Like",
	"IColor4Like",
	"IImageProcessingConfigurationDefines",
	"IMaterialAnisotropicDefines",
	"IMaterialBRDFDefines",
	"IMaterialClearCoatDefines",
	"IMaterialCompilationOptions",
	"IMaterialSheenDefines",
	"IMaterialSubSurfaceDefines",
	"IMatrixLike",
	"IPhysicsEnginePlugin",
	"IPlaneLike",
	"ISceneLike",
	"ISmartArrayLike",
	"IVector2Like",
	"IVector3Like",
	"IVector4Like",
	"IViewportLike",
	"IViewportOwnerLike",
	"IndicesArray",
	"IndividualBabylonFileParser",
	"InspectableType",
	"Int32Array",
	"InternalTextureSource",
	"IntersectionInfo",
	"IsWindowObjectExist",
	"JoystickAxis",
	"KeyboardEvent",
	"MediaStream",
	"MediaStreamTrack",
	"MediaTrackConstraints",
	"MeshLoadOptions",
	"MeshLoadOptions",
	"NodeConstructor",
	"NodeMaterialBlockConnectionPointTypes",
	"NodeMaterialBlockTargets",
	"NodeMaterialConnectionPointCompatibilityStates",
	"NodeMaterialConnectionPointDirection",
	"NodeMaterialDefines",
	"NodeMaterialSystemValues",
	"Object",
	"Orientation",
	"PBRMaterialDefines",
	"PhysicsGravitationalFieldEventData",
	"PhysicsImpostorJoint",
	"PhysicsImpostorParameters",
	"PhysicsJointData",
	"PhysicsRadialImpulseFalloff",
	"PhysicsUpdraftMode",
	"PointerEvent",
	"PointerEventInit",
	"PoseEnabledControllerType",
	"ProgressEvent",
	"RegExp",
	"SceneOptions",
	"SimplificationType",
	"Space",
	"SpringJointData",
	"StandardMaterialDefines",
	"StencilState",
	"SubEmitterType",
	"TEX",
	"TextWrapping",
	"TonemappingOperator",
	"TrianglePickingPredicate",
	"TrigonometryBlockOperations",
	"Uint8Array",
	"VRExperienceHelperOptions",
	"VideoRecorderOptions",
	"VideoTextureSettings",
	"WEBGL_compressed_texture_s3tc",
	"WaveBlockKind",
	"WebGLBuffer",
	"WebGLProgram",
	"WebGLQuery",
	"WebGLRenderingContext",
	"WebGLTransformFeedback",
	"WebGLUniformLocation",
	"WebGLVertexArrayObject",
	"WebVROptions",
	"WebXRState",
	"Window",
	"Worker",
	"XMLHttpRequestResponseType",
	"XREye",
	"XRFrame",
	"XRInputSource",
	"XRReferenceSpace",
	"XRReferenceSpaceType",
	"XRRenderState",
	"XRSession",
	"XRSessionMode",
	"XRWebGLLayer",
	"XRWebGLLayerOptions",
	"_InstancesBatch",
	"_TimeToken",
	"null",
	"unknown",
	// "IAction",
	// "IActionEvent",
	// "IAnimatable",
	// "IAnimation",
	// "IArrayItem",
	// "IBufferView",
	// "ICamera",
	// "ICameraInput",
	// "ICollisionCoordinator",
	// "ICullable",
	// "IDataBuffer",
	// "IDisplayChangedEventArgs",
	// "IEasingFunction",
	// "IEffectFallbacks",
	// "IEffectRendererOptions",
	// "IEnvironmentHelperOptions",
	// "IExportOptions",
	// "IFocusableControl",
	// "IGlowLayerOptions",
	// "IGlowLayerOptions",
	// "IHighlightLayerOptions",
	// "IHighlightLayerOptions",
	// "IHtmlElementTextureOptions",
	// "IImage",
	// "IInspectorOptions",
	// "IInternalTextureLoader",
	// "ILoadingScreen",
	// "IMaterial",
	// "IMotorEnabledJoint",
	// "IMultiRenderTargetOptions",
	// "INode",
	// "INodeMaterialEditorOptions",
	// "INodeMaterialOptions",
	// "IOceanPostProcessOptions",
	// "IParticleSystem",
	// "IPhysicsEnabledObject",
	// "IPipelineContext",
	// "IProperty",
	// "IScene",
	// "IShaderMaterialOptions",
	// "IShadowGenerator",
	// "IShadowLight",
	// "ISize",
	// "ISoundOptions",
	// "ISoundTrackOptions",
	// "ISpriteManager",
	// "ITextureInfo",
	// "IValueGradient",
}
