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
func (ba *Babylon) SelectorGroup() *SelectorGroup {
	p := ba.ctx.Get("SelectorGroup")
	return SelectorGroupFromJSObject(p, ba.ctx)
}

// SelectorGroupFromJSObject returns a wrapped SelectorGroup JavaScript class.
func SelectorGroupFromJSObject(p js.Value, ctx js.Value) *SelectorGroup {
	return &SelectorGroup{p: p, ctx: ctx}
}

// NewSelectorGroup returns a new SelectorGroup object.
//
// https://doc.babylonjs.com/api/classes/babylon.selectorgroup
func (ba *Babylon) NewSelectorGroup(name string) *SelectorGroup {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("SelectorGroup").New(args...)
	return SelectorGroupFromJSObject(p, ba.ctx)
}

// RemoveSelector calls the RemoveSelector method on the SelectorGroup object.
//
// https://doc.babylonjs.com/api/classes/babylon.selectorgroup#removeselector
func (s *SelectorGroup) RemoveSelector(selectorNb float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, selectorNb)

	s.p.Call("removeSelector", args...)
}

/*

// GroupPanel returns the GroupPanel property of class SelectorGroup.
//
// https://doc.babylonjs.com/api/classes/babylon.selectorgroup#grouppanel
func (s *SelectorGroup) GroupPanel(groupPanel *StackPanel) *SelectorGroup {
	p := ba.ctx.Get("SelectorGroup").New(groupPanel.JSObject())
	return SelectorGroupFromJSObject(p, ba.ctx)
}

// SetGroupPanel sets the GroupPanel property of class SelectorGroup.
//
// https://doc.babylonjs.com/api/classes/babylon.selectorgroup#grouppanel
func (s *SelectorGroup) SetGroupPanel(groupPanel *StackPanel) *SelectorGroup {
	p := ba.ctx.Get("SelectorGroup").New(groupPanel.JSObject())
	return SelectorGroupFromJSObject(p, ba.ctx)
}

// Header returns the Header property of class SelectorGroup.
//
// https://doc.babylonjs.com/api/classes/babylon.selectorgroup#header
func (s *SelectorGroup) Header(header string) *SelectorGroup {
	p := ba.ctx.Get("SelectorGroup").New(header)
	return SelectorGroupFromJSObject(p, ba.ctx)
}

// SetHeader sets the Header property of class SelectorGroup.
//
// https://doc.babylonjs.com/api/classes/babylon.selectorgroup#header
func (s *SelectorGroup) SetHeader(header string) *SelectorGroup {
	p := ba.ctx.Get("SelectorGroup").New(header)
	return SelectorGroupFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class SelectorGroup.
//
// https://doc.babylonjs.com/api/classes/babylon.selectorgroup#name
func (s *SelectorGroup) Name(name string) *SelectorGroup {
	p := ba.ctx.Get("SelectorGroup").New(name)
	return SelectorGroupFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class SelectorGroup.
//
// https://doc.babylonjs.com/api/classes/babylon.selectorgroup#name
func (s *SelectorGroup) SetName(name string) *SelectorGroup {
	p := ba.ctx.Get("SelectorGroup").New(name)
	return SelectorGroupFromJSObject(p, ba.ctx)
}

// Selectors returns the Selectors property of class SelectorGroup.
//
// https://doc.babylonjs.com/api/classes/babylon.selectorgroup#selectors
func (s *SelectorGroup) Selectors(selectors *StackPanel) *SelectorGroup {
	p := ba.ctx.Get("SelectorGroup").New(selectors.JSObject())
	return SelectorGroupFromJSObject(p, ba.ctx)
}

// SetSelectors sets the Selectors property of class SelectorGroup.
//
// https://doc.babylonjs.com/api/classes/babylon.selectorgroup#selectors
func (s *SelectorGroup) SetSelectors(selectors *StackPanel) *SelectorGroup {
	p := ba.ctx.Get("SelectorGroup").New(selectors.JSObject())
	return SelectorGroupFromJSObject(p, ba.ctx)
}

*/
