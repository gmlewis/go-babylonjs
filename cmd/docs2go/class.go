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
			log.Printf("Unknown filename %v, skipping", path)
			return nil
		}
		if strings.HasPrefix(filename, "babylon._") {
			log.Printf("Skipping private class %v", path)
			return nil
		}

		log.Printf("Processing %v ...", filename)

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

		log.Printf("html.Title=%v", html.Title.InnerXML)

		for _, div := range html.Div {
			if div.ID != "wrapper" {
				continue
			}
			log.Printf("html.Div:\nid=%v\nh1=%v", div.ID, div.H1.InnerXML)

			for _, section := range div.Sections {
				switch t := section.Type(); t {
				case "tsd-comment":
					log.Printf("%v: lead=%v\ndescription=%v\nseeURL=%v", t, section.Lead.InnerXML, section.Description.InnerXML, section.SeeURL.InnerXML)
				case "tsd-hierarchy":
					log.Printf("%v: parent=%v", t, section.ListItems)
				case "Implements":
					log.Printf("%v: implements=%v", t, section.ListItems)
				case "tsd-index-group":
					for _, indexSection := range section.IndexSections {
						log.Printf("%v: indexSection=%v: %v", t, indexSection.Type(), indexSection.ListItems)
					}
				case "tsd-member-group": // ignore
				case "tsd-is-not-exported": // ignore
				case "tsd-is-inherited": // ignore
				case "tsd-type-parameters": // ignore
				case "tsd-parent-kind-module": // ignore
				default:
					log.Fatalf("Unknown section type: %v", t)
				}
			}
		}

		return nil
	}
}

type ClassHTML struct {
	Title InnerXML `xml:"head>title"`
	Div   []*Div   `xml:"body>div"`
}

type Div struct {
	ID string   `xml:"id,attr"`
	H1 InnerXML `xml:"header>div>div>h1"`
	// Lead        InnerXML `xml:"div>div>div>section>div>div>p"`
	// Description InnerXML `xml:"div>div>div>section>div>p"`
	// SeeURL      InnerXML `xml:"div>div>div>section>div>dl>dd>p>a"`

	Sections []*Section `xml:"div>div>div>section"`
	// Parent InnerXML `xml:"div>div>div>section>ul>li>a"`
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
