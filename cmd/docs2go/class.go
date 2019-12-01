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

			if s := div.GetHeaderDiv(0).GetDiv(0).GetH1().GetInnerXML(); s != "" {
				html.Name = s
				if i := strings.Index(html.Name, "&lt;"); i >= 0 {
					html.Name = html.Name[0:i]
				}
			} else {
				log.Fatalf("%v: unable to find class name.", filename)
			}
			logf("html.Name=%v", html.Name)

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
				if s := ul.GetLI(0).GetA(0).GetInnerXML(); s != "" { // optional
					html.Parents = append(html.Parents, s)
				}
				if s := ul.GetLI(0).GetUL(0).GetLI(0).GetUL(0).GetLI(0).GetA(0).GetInnerXML(); s != "" { // optional
					html.Children = append(html.Children, s)
				}
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

		/*
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
		*/

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

	InnerXML string

	GoName string
	JSName string

	GoReturnType      string
	GoReturnStatement string

	HasOpts bool

	GoParams         []string
	GoOpts           []string
	GoOptsName       []string
	NeedsArrayHelper []string
	JSParams         []string
	JSOpts           []string
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
	case "InstancedMesh.GetFacetNormal", "InstancedMesh.GetFacetNormalToRef", "InstancedMesh.GetFacetPosition", "InstancedMesh.GetFacetPositionToRef":
		names[0] = "index"
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
					case "number":
						types = append(types, "float64")
					case "TCamera": // special case
						types = append(types, "*Camera")
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
					case "function", "Function":
						paramType = "func()"
					case
						"[]AbstractMesh",
						"[]IInternalTextureLoader",
						"[]Worker":
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
			var needsArrayHelper bool
			if i < len(types) {
				paramType = types[i]
				var needsJSObject bool
				paramType, needsJSObject, needsArrayHelper = jsTypeToGoType(paramType)
				if needsJSObject {
					if !strings.HasPrefix(paramType, "*") && paramType != "js.Value" && !strings.HasPrefix(paramType, "[]") {
						paramType = "*" + paramType
					}
					jsName += ".JSObject()"
				}

				// special case - function callback
				if paramType == "func()" {
					jsName = fmt.Sprintf(`js.FuncOf(func(this js.Value, args []js.Value) interface{} {%v(); return nil})`, name)
				}
			}
			if needsArrayHelper {
				s.NeedsArrayHelper = append(s.NeedsArrayHelper, fmt.Sprintf("%vArrayToJSArray(%v)", paramType[3:], name))
			} else {
				s.NeedsArrayHelper = append(s.NeedsArrayHelper, "")
			}
			s.GoParams = append(s.GoParams, fmt.Sprintf("%v %v", name, paramType))
			s.JSParams = append(s.JSParams, jsName)
		}
	}

	if len(types) > len(names) {
		var needsJSObject bool
		s.GoReturnType, needsJSObject, _ = jsTypeToGoType(types[len(types)-1])

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

func jsTypeToGoType(paramType string) (goType string, needsJSObject, needsArrayHelper bool) {
	if unhandledType[paramType] {
		return "js.Value", false, false
	}

	if strings.HasPrefix(paramType, "[]*") {
		return paramType, false, true
	}

	if strings.HasPrefix(paramType, "[]") {
		return "[]*" + paramType[2:], false, true
	}

	switch paramType {
	case "any":
		paramType = "interface{}"
	case "function", "Function":
		paramType = "func()"
	case "string", "float64", "bool", "[]string", "[]float64", "[]bool":
	case "void":
		paramType = ""
	case "Boolean":
		paramType = "bool"
	default:
		needsJSObject = true
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
	"ArrayBuffer",
	"ArrayBufferView",
	"ArrayLike",
	"AudioNode",
	"BabylonFileParser",
	"Behavior",
	"CanvasRenderingContext2D",
	"ClientRect",
	"ClipboardEvent",
	"Collider",
	"CubeMapInfo",
	"DDSInfo",
	"DeepImmutable",
	"DevicePose",
	"DistanceJointData",
	"DistanceJointData",
	"Document",
	"EffectWrapperCreationOptions",
	"EffectWrapperCreationOptions",
	"EngineCapabilities",
	"EngineOptions",
	"EngineOptions",
	"EnvironmentTextureInfo",
	"Event",
	"Float32Array",
	"FloatArray",
	"HDRInfo",
	"HTMLCanvasElement",
	"HTMLElement",
	"HTMLImageElement",
	"HTMLVideoElement",
	"IAction",
	"IActionEvent",
	"IAnimatable",
	"IAnimation",
	"IArrayItem",
	"IBufferView",
	"ICamera",
	"ICameraInput",
	"IColor3Like",
	"IColor3Like",
	"IColor4Like",
	"ICullable",
	"IDataBuffer",
	"IDataBuffer",
	"IDisplayChangedEventArgs",
	"IEasingFunction",
	"IEffectFallbacks",
	"IEffectFallbacks",
	"IEffectRendererOptions",
	"IEffectRendererOptions",
	"IEnvironmentHelperOptions",
	"IEnvironmentHelperOptions",
	"IExportOptions",
	"IFocusableControl",
	"IGlowLayerOptions",
	"IGlowLayerOptions",
	"IHighlightLayerOptions",
	"IHighlightLayerOptions",
	"IHtmlElementTextureOptions",
	"IHtmlElementTextureOptions",
	"IImage",
	"IImageProcessingConfigurationDefines",
	"IInspectorOptions",
	"IInternalTextureLoader",
	"ILoadingScreen",
	"IMaterial",
	"IMaterialAnisotropicDefines",
	"IMaterialBRDFDefines",
	"IMaterialClearCoatDefines",
	"IMaterialCompilationOptions",
	"IMaterialSheenDefines",
	"IMaterialSubSurfaceDefines",
	"IMatrixLike",
	"IMotorEnabledJoint",
	"IMultiRenderTargetOptions",
	"IMultiRenderTargetOptions",
	"INode",
	"INodeMaterialEditorOptions",
	"INodeMaterialOptions",
	"IOceanPostProcessOptions",
	"IOceanPostProcessOptions",
	"IParticleSystem",
	"IPhysicsEnabledObject",
	"IPhysicsEnginePlugin",
	"IPipelineContext",
	"IPlaneLike",
	"IProperty",
	"IScene",
	"ISceneLike",
	"IShaderMaterialOptions",
	"IShadowGenerator",
	"IShadowLight",
	"IShadowLight",
	"ISize",
	"ISmartArrayLike",
	"ISoundOptions",
	"ISoundTrackOptions",
	"ISpriteManager",
	"ITextureInfo",
	"IValueGradient",
	"IVector2Like",
	"IVector3Like",
	"IVector4Like",
	"IViewportLike",
	"IViewportOwnerLike",
	"IndicesArray",
	"IndividualBabylonFileParser",
	"Int32Array",
	"InternalTextureSource",
	"InternalTextureSource",
	"KeyboardEvent",
	"MeshLoadOptions",
	"MeshLoadOptions",
	"NodeConstructor",
	"NodeMaterialBlockConnectionPointTypes",
	"NodeMaterialBlockConnectionPointTypes",
	"NodeMaterialBlockTargets",
	"NodeMaterialBlockTargets",
	"NodeMaterialConnectionPointCompatibilityStates",
	"NodeMaterialConnectionPointDirection",
	"NodeMaterialConnectionPointDirection",
	"NodeMaterialDefines",
	"NodeMaterialSystemValues",
	"PBRMaterialDefines",
	"PhysicsGravitationalFieldEventData",
	"PhysicsImpostorJoint",
	"PhysicsImpostorParameters",
	"PhysicsJointData",
	"PhysicsJointData",
	"PhysicsRadialImpulseFalloff",
	"PhysicsUpdraftMode",
	"PointerEvent",
	"PointerEventInit",
	"SceneOptions",
	"SimplificationType",
	"Space",
	"SpringJointData",
	"StandardMaterialDefines",
	"TonemappingOperator",
	"TrianglePickingPredicate",
	"Uint8Array",
	"VRExperienceHelperOptions",
	"VideoRecorderOptions",
	"VideoTextureSettings",
	"WebGLBuffer",
	"WebGLProgram",
	"WebGLQuery",
	"WebGLRenderingContext",
	"WebGLTransformFeedback",
	"WebGLUniformLocation",
	"WebGLVertexArrayObject",
	"WebVROptions",
	"WebVROptions",
	"WebXRState",
	"Window",
	"Worker",
	"XRInputSource",
	"XRReferenceSpaceType",
	"XRSessionMode",
	"_InstancesBatch",
	"_TimeToken",
	"null",
	"object",
	"unknown",
}
