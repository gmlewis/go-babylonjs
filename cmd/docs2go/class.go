package main

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	scriptRE = regexp.MustCompile(`<script>.*?</script>`)
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

		logf("\n\nProcessing %v ...", filename)

		buf, err := ioutil.ReadFile(path)
		check("ReadFile: %v", err)

		// Strip problematic <script> sections that the XML parser fails on.
		buf = scriptRE.ReplaceAll(buf, nil)

		d := xml.NewDecoder(bytes.NewReader(buf))
		d.Strict = false
		d.AutoClose = xml.HTMLAutoClose
		d.Entity = xml.HTMLEntity

		html := &ClassHTML{}
		check("xml.Decode: %v", d.Decode(&html))

		for _, div := range html.Div {
			if div.ID != "wrapper" {
				continue
			}
			html.Name = div.H1.InnerXML

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
					for _, indexSection := range section.IndexSections {
						logf("%v: indexSection=%v: %v", t, indexSection.Type(), indexSection.ListItems)
					}
				case "tsd-member-group": // TODO: Parse this for Constructors, Properties, and Methods
				case "tsd-is-not-exported": // ignore
				case "tsd-is-inherited": // ignore
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

		c.m[html.Name] = html

		return nil
	}
}

type ClassHTML struct {
	Title InnerXML `xml:"head>title"`
	Div   []*Div   `xml:"body>div"`

	Name        string   `xml:"-"`
	Summary     string   `xml:"-"`
	Description string   `xml:"-"`
	SeeURL      string   `xml:"-"`
	Parents     []string `xml:"-"`
	Children    []string `xml:"-"`
	Implements  []string `xml:"-"`
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
	IndexSections []*Section `xml:"section>div>section"`
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
