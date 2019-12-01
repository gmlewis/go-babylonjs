package main

import (
	"errors"
	"fmt"
)

func (html *ClassHTML) parseMethods() error {
	for _, div := range html.Div {
		if div.ID != "wrapper" {
			continue
		}

		mainDiv := div.GetDiv(0).GetDiv(0).GetDiv(0)
		if mainDiv == nil {
			return fmt.Errorf("methods: unable to find main div.")
		}

		methodSection := mainDiv.FindSections(func(s *Section) bool {
			return s.HasClass("tsd-member-group") && s.GetH2().GetInnerXML() == "Methods"
		})
		if len(methodSection) == 0 { // Methods are optional
			return nil
		}

		if len(methodSection) != 1 {
			return fmt.Errorf("methods: found %v method sections, want 1", len(methodSection))
		}
		for _, section := range methodSection[0].Sections {
			li := section.GetUL(0).GetLI(0)
			if li == nil || !li.HasClass("tsd-signature") {
				return errors.New("methods: unable to find tsd-signature")
			}
			v := li.GetSignature()
			if ok := v.parseParameters(html.Name, processMethodOverrides); ok {
				html.MethodNames[v.GoName] = v
			}
		}
	}

	return nil
}
