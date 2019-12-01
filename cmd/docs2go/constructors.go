package main

import (
	"errors"
	"fmt"
)

func (html *ClassHTML) parseConstructors() error {
	for _, div := range html.Div {
		if div.ID != "wrapper" {
			continue
		}

		mainDiv := div.GetDiv(0).GetDiv(0).GetDiv(0)
		if mainDiv == nil {
			return fmt.Errorf("constructors: unable to find main div.")
		}

		constructorSection := mainDiv.FindSections(func(s *Section) bool {
			return s.HasClass("tsd-member-group") && s.GetH2().GetInnerXML() == "Constructors"
		})
		if len(constructorSection) == 0 { // No constructor for abstract classes.
			return nil
		}

		if len(constructorSection) != 1 {
			return fmt.Errorf("constructors: found %v constructor sections, want 1", len(constructorSection))
		}
		li := constructorSection[0].GetSection(0).GetUL(0).GetLI(0)
		if !li.HasClass("tsd-signature") {
			return errors.New("constructors: unable to find tsd-signature")
		}
		v := li.GetSignature()
		if ok := v.parseParameters(html.Name, processConstructorOverrides); ok {
			html.ConstructorNames[v.GoName] = v
		}
	}

	return nil
}
