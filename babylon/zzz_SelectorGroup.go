// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SelectorGroup represents a babylon.js SelectorGroup.
// Class used to create a RadioGroup
// which contains groups of radio buttons
type SelectorGroup struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *SelectorGroup) JSObject() js.Value { return s.p }

// SelectorGroup returns a SelectorGroup JavaScript class.
func (gui *GUI) SelectorGroup() *SelectorGroup {
	p := gui.ctx.Get("SelectorGroup")
	return SelectorGroupFromJSObject(p, gui.ctx)
}

// SelectorGroupFromJSObject returns a wrapped SelectorGroup JavaScript class.
func SelectorGroupFromJSObject(p js.Value, ctx js.Value) *SelectorGroup {
	return &SelectorGroup{p: p, ctx: ctx}
}

// SelectorGroupArrayToJSArray returns a JavaScript Array for the wrapped array.
func SelectorGroupArrayToJSArray(array []*SelectorGroup) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewSelectorGroup returns a new SelectorGroup object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.selectorgroup#constructor
func (gui *GUI) NewSelectorGroup(name string) *SelectorGroup {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := gui.ctx.Get("SelectorGroup").New(args...)
	return SelectorGroupFromJSObject(p, gui.ctx)
}

// RemoveSelector calls the RemoveSelector method on the SelectorGroup object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.selectorgroup#removeselector
func (s *SelectorGroup) RemoveSelector(selectorNb float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, selectorNb)

	s.p.Call("removeSelector", args...)
}

// GroupPanel returns the GroupPanel property of class SelectorGroup.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.selectorgroup#grouppanel
func (s *SelectorGroup) GroupPanel() *StackPanel {
	retVal := s.p.Get("groupPanel")
	return StackPanelFromJSObject(retVal, s.ctx)
}

// SetGroupPanel sets the GroupPanel property of class SelectorGroup.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.selectorgroup#grouppanel
func (s *SelectorGroup) SetGroupPanel(groupPanel *StackPanel) *SelectorGroup {
	s.p.Set("groupPanel", groupPanel.JSObject())
	return s
}

// Header returns the Header property of class SelectorGroup.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.selectorgroup#header
func (s *SelectorGroup) Header() string {
	retVal := s.p.Get("header")
	return retVal.String()
}

// SetHeader sets the Header property of class SelectorGroup.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.selectorgroup#header
func (s *SelectorGroup) SetHeader(header string) *SelectorGroup {
	s.p.Set("header", header)
	return s
}

// Name returns the Name property of class SelectorGroup.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.selectorgroup#name
func (s *SelectorGroup) Name() string {
	retVal := s.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class SelectorGroup.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.selectorgroup#name
func (s *SelectorGroup) SetName(name string) *SelectorGroup {
	s.p.Set("name", name)
	return s
}

// Selectors returns the Selectors property of class SelectorGroup.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.selectorgroup#selectors
func (s *SelectorGroup) Selectors() []*StackPanel {
	retVal := s.p.Get("selectors")
	result := []*StackPanel{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, StackPanelFromJSObject(retVal.Index(ri), s.ctx))
	}
	return result
}

// SetSelectors sets the Selectors property of class SelectorGroup.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.selectorgroup#selectors
func (s *SelectorGroup) SetSelectors(selectors []*StackPanel) *SelectorGroup {
	s.p.Set("selectors", selectors)
	return s
}
