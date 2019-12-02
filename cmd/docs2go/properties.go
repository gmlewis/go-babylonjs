package main

import (
	"errors"
	"fmt"
)

func (html *ClassHTML) parseProperties() error {
	for _, div := range html.Div {
		if div.ID != "wrapper" {
			continue
		}

		mainDiv := div.GetDiv(0).GetDiv(0).GetDiv(0)
		if mainDiv == nil {
			return fmt.Errorf("properties: unable to find main div.")
		}

		propertiesSection := mainDiv.FindSections(func(s *Section) bool {
			return s.HasClass("tsd-member-group") && s.GetH2().GetInnerXML() == "Properties"
		})
		if len(propertiesSection) == 0 { // Properties are optional
			return nil
		}

		if len(propertiesSection) != 1 {
			return fmt.Errorf("properties: found %v properties sections, want 1", len(propertiesSection))
		}
		for _, section := range propertiesSection[0].Sections {
			if section.HasClass("tsd-is-inherited") {
				continue
			}
			div := section.GetDiv(0)
			if div == nil || !div.HasClass("tsd-signature") {
				return errors.New("properties: unable to find tsd-signature")
			}
			v := div.GetSignature()
			if ok := v.parseParameters(html.Name, processPropertiesOverrides); ok {
				html.PropertyNames[v.GoName] = v
			}
		}
	}

	return nil
}
