// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SelectionPanel represents a babylon.js SelectionPanel.
// Class used to hold the controls for the checkboxes, radio buttons and sliders
//
// See: http://doc.babylonjs.com/how_to/selector
type SelectionPanel struct {
	*Rectangle
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *SelectionPanel) JSObject() js.Value { return s.p }

// SelectionPanel returns a SelectionPanel JavaScript class.
func (ba *Babylon) SelectionPanel() *SelectionPanel {
	p := ba.ctx.Get("SelectionPanel")
	return SelectionPanelFromJSObject(p, ba.ctx)
}

// SelectionPanelFromJSObject returns a wrapped SelectionPanel JavaScript class.
func SelectionPanelFromJSObject(p js.Value, ctx js.Value) *SelectionPanel {
	return &SelectionPanel{Rectangle: RectangleFromJSObject(p, ctx), ctx: ctx}
}

// SelectionPanelArrayToJSArray returns a JavaScript Array for the wrapped array.
func SelectionPanelArrayToJSArray(array []*SelectionPanel) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewSelectionPanelOpts contains optional parameters for NewSelectionPanel.
type NewSelectionPanelOpts struct {
	Groups *SelectorGroup
}

// NewSelectionPanel returns a new SelectionPanel object.
//
// https://doc.babylonjs.com/api/classes/babylon.selectionpanel
func (ba *Babylon) NewSelectionPanel(name string, opts *NewSelectionPanelOpts) *SelectionPanel {
	if opts == nil {
		opts = &NewSelectionPanelOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, name)

	if opts.Groups == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Groups.JSObject())
	}

	p := ba.ctx.Get("GUI").Get("SelectionPanel").New(args...)
	return SelectionPanelFromJSObject(p, ba.ctx)
}

// AddGroup calls the AddGroup method on the SelectionPanel object.
//
// https://doc.babylonjs.com/api/classes/babylon.selectionpanel#addgroup
func (s *SelectionPanel) AddGroup(group *SelectorGroup) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, group.JSObject())

	s.p.Call("addGroup", args...)
}

// SelectionPanelAddToGroupCheckboxOpts contains optional parameters for SelectionPanel.AddToGroupCheckbox.
type SelectionPanelAddToGroupCheckboxOpts struct {
	Func    func()
	Checked *bool
}

// AddToGroupCheckbox calls the AddToGroupCheckbox method on the SelectionPanel object.
//
// https://doc.babylonjs.com/api/classes/babylon.selectionpanel#addtogroupcheckbox
func (s *SelectionPanel) AddToGroupCheckbox(groupNb float64, label string, opts *SelectionPanelAddToGroupCheckboxOpts) {
	if opts == nil {
		opts = &SelectionPanelAddToGroupCheckboxOpts{}
	}

	args := make([]interface{}, 0, 2+2)

	args = append(args, groupNb)
	args = append(args, label)

	if opts.Func == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { opts.Func(); return nil }) /* never freed! */)
	}
	if opts.Checked == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Checked)
	}

	s.p.Call("addToGroupCheckbox", args...)
}

// SelectionPanelAddToGroupRadioOpts contains optional parameters for SelectionPanel.AddToGroupRadio.
type SelectionPanelAddToGroupRadioOpts struct {
	Func    func()
	Checked *bool
}

// AddToGroupRadio calls the AddToGroupRadio method on the SelectionPanel object.
//
// https://doc.babylonjs.com/api/classes/babylon.selectionpanel#addtogroupradio
func (s *SelectionPanel) AddToGroupRadio(groupNb float64, label string, opts *SelectionPanelAddToGroupRadioOpts) {
	if opts == nil {
		opts = &SelectionPanelAddToGroupRadioOpts{}
	}

	args := make([]interface{}, 0, 2+2)

	args = append(args, groupNb)
	args = append(args, label)

	if opts.Func == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { opts.Func(); return nil }) /* never freed! */)
	}
	if opts.Checked == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Checked)
	}

	s.p.Call("addToGroupRadio", args...)
}

// SelectionPanelAddToGroupSliderOpts contains optional parameters for SelectionPanel.AddToGroupSlider.
type SelectionPanelAddToGroupSliderOpts struct {
	Func  func()
	Unit  *string
	Min   *float64
	Max   *float64
	Value *float64
	OnVal func()
}

// AddToGroupSlider calls the AddToGroupSlider method on the SelectionPanel object.
//
// https://doc.babylonjs.com/api/classes/babylon.selectionpanel#addtogroupslider
func (s *SelectionPanel) AddToGroupSlider(groupNb float64, label string, opts *SelectionPanelAddToGroupSliderOpts) {
	if opts == nil {
		opts = &SelectionPanelAddToGroupSliderOpts{}
	}

	args := make([]interface{}, 0, 2+6)

	args = append(args, groupNb)
	args = append(args, label)

	if opts.Func == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { opts.Func(); return nil }) /* never freed! */)
	}
	if opts.Unit == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Unit)
	}
	if opts.Min == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Min)
	}
	if opts.Max == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Max)
	}
	if opts.Value == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Value)
	}
	if opts.OnVal == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { opts.OnVal(); return nil }) /* never freed! */)
	}

	s.p.Call("addToGroupSlider", args...)
}

// Relabel calls the Relabel method on the SelectionPanel object.
//
// https://doc.babylonjs.com/api/classes/babylon.selectionpanel#relabel
func (s *SelectionPanel) Relabel(label string, groupNb float64, selectorNb float64) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, label)
	args = append(args, groupNb)
	args = append(args, selectorNb)

	s.p.Call("relabel", args...)
}

// RemoveFromGroupSelector calls the RemoveFromGroupSelector method on the SelectionPanel object.
//
// https://doc.babylonjs.com/api/classes/babylon.selectionpanel#removefromgroupselector
func (s *SelectionPanel) RemoveFromGroupSelector(groupNb float64, selectorNb float64) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, groupNb)
	args = append(args, selectorNb)

	s.p.Call("removeFromGroupSelector", args...)
}

// RemoveGroup calls the RemoveGroup method on the SelectionPanel object.
//
// https://doc.babylonjs.com/api/classes/babylon.selectionpanel#removegroup
func (s *SelectionPanel) RemoveGroup(groupNb float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, groupNb)

	s.p.Call("removeGroup", args...)
}

// SetHeaderName calls the SetHeaderName method on the SelectionPanel object.
//
// https://doc.babylonjs.com/api/classes/babylon.selectionpanel#setheadername
func (s *SelectionPanel) SetHeaderName(label string, groupNb float64) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, label)
	args = append(args, groupNb)

	s.p.Call("setHeaderName", args...)
}

// BarColor returns the BarColor property of class SelectionPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.selectionpanel#barcolor
func (s *SelectionPanel) BarColor() string {
	retVal := s.p.Get("barColor")
	return retVal.String()
}

// SetBarColor sets the BarColor property of class SelectionPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.selectionpanel#barcolor
func (s *SelectionPanel) SetBarColor(barColor string) *SelectionPanel {
	s.p.Set("barColor", barColor)
	return s
}

// BarHeight returns the BarHeight property of class SelectionPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.selectionpanel#barheight
func (s *SelectionPanel) BarHeight() string {
	retVal := s.p.Get("barHeight")
	return retVal.String()
}

// SetBarHeight sets the BarHeight property of class SelectionPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.selectionpanel#barheight
func (s *SelectionPanel) SetBarHeight(barHeight string) *SelectionPanel {
	s.p.Set("barHeight", barHeight)
	return s
}

// ButtonBackground returns the ButtonBackground property of class SelectionPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.selectionpanel#buttonbackground
func (s *SelectionPanel) ButtonBackground() string {
	retVal := s.p.Get("buttonBackground")
	return retVal.String()
}

// SetButtonBackground sets the ButtonBackground property of class SelectionPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.selectionpanel#buttonbackground
func (s *SelectionPanel) SetButtonBackground(buttonBackground string) *SelectionPanel {
	s.p.Set("buttonBackground", buttonBackground)
	return s
}

// ButtonColor returns the ButtonColor property of class SelectionPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.selectionpanel#buttoncolor
func (s *SelectionPanel) ButtonColor() string {
	retVal := s.p.Get("buttonColor")
	return retVal.String()
}

// SetButtonColor sets the ButtonColor property of class SelectionPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.selectionpanel#buttoncolor
func (s *SelectionPanel) SetButtonColor(buttonColor string) *SelectionPanel {
	s.p.Set("buttonColor", buttonColor)
	return s
}

// Groups returns the Groups property of class SelectionPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.selectionpanel#groups
func (s *SelectionPanel) Groups() *SelectorGroup {
	retVal := s.p.Get("groups")
	return SelectorGroupFromJSObject(retVal, s.ctx)
}

// SetGroups sets the Groups property of class SelectionPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.selectionpanel#groups
func (s *SelectionPanel) SetGroups(groups *SelectorGroup) *SelectionPanel {
	s.p.Set("groups", groups.JSObject())
	return s
}

// HeaderColor returns the HeaderColor property of class SelectionPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.selectionpanel#headercolor
func (s *SelectionPanel) HeaderColor() string {
	retVal := s.p.Get("headerColor")
	return retVal.String()
}

// SetHeaderColor sets the HeaderColor property of class SelectionPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.selectionpanel#headercolor
func (s *SelectionPanel) SetHeaderColor(headerColor string) *SelectionPanel {
	s.p.Set("headerColor", headerColor)
	return s
}

// LabelColor returns the LabelColor property of class SelectionPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.selectionpanel#labelcolor
func (s *SelectionPanel) LabelColor() string {
	retVal := s.p.Get("labelColor")
	return retVal.String()
}

// SetLabelColor sets the LabelColor property of class SelectionPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.selectionpanel#labelcolor
func (s *SelectionPanel) SetLabelColor(labelColor string) *SelectionPanel {
	s.p.Set("labelColor", labelColor)
	return s
}

// Name returns the Name property of class SelectionPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.selectionpanel#name
func (s *SelectionPanel) Name() string {
	retVal := s.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class SelectionPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.selectionpanel#name
func (s *SelectionPanel) SetName(name string) *SelectionPanel {
	s.p.Set("name", name)
	return s
}

// SpacerHeight returns the SpacerHeight property of class SelectionPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.selectionpanel#spacerheight
func (s *SelectionPanel) SpacerHeight() string {
	retVal := s.p.Get("spacerHeight")
	return retVal.String()
}

// SetSpacerHeight sets the SpacerHeight property of class SelectionPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.selectionpanel#spacerheight
func (s *SelectionPanel) SetSpacerHeight(spacerHeight string) *SelectionPanel {
	s.p.Set("spacerHeight", spacerHeight)
	return s
}
