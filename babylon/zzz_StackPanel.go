// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// StackPanel represents a babylon.js StackPanel.
// Class used to create a 2D stack panel container
type StackPanel struct {
	*Container
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *StackPanel) JSObject() js.Value { return s.p }

// StackPanel returns a StackPanel JavaScript class.
func (gui *GUI) StackPanel() *StackPanel {
	p := gui.ctx.Get("StackPanel")
	return StackPanelFromJSObject(p, gui.ctx)
}

// StackPanelFromJSObject returns a wrapped StackPanel JavaScript class.
func StackPanelFromJSObject(p js.Value, ctx js.Value) *StackPanel {
	return &StackPanel{Container: ContainerFromJSObject(p, ctx), ctx: ctx}
}

// StackPanelArrayToJSArray returns a JavaScript Array for the wrapped array.
func StackPanelArrayToJSArray(array []*StackPanel) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewStackPanelOpts contains optional parameters for NewStackPanel.
type NewStackPanelOpts struct {
	Name *string
}

// NewStackPanel returns a new StackPanel object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.stackpanel#constructor
func (gui *GUI) NewStackPanel(opts *NewStackPanelOpts) *StackPanel {
	if opts == nil {
		opts = &NewStackPanelOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Name == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Name)
	}

	p := gui.ctx.Get("StackPanel").New(args...)
	return StackPanelFromJSObject(p, gui.ctx)
}

// Height returns the Height property of class StackPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.stackpanel#height
func (s *StackPanel) Height() string {
	retVal := s.p.Get("height")
	return retVal.String()
}

// SetHeight sets the Height property of class StackPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.stackpanel#height
func (s *StackPanel) SetHeight(height string) *StackPanel {
	s.p.Set("height", height)
	return s
}

// IgnoreLayoutWarnings returns the IgnoreLayoutWarnings property of class StackPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.stackpanel#ignorelayoutwarnings
func (s *StackPanel) IgnoreLayoutWarnings() bool {
	retVal := s.p.Get("ignoreLayoutWarnings")
	return retVal.Bool()
}

// SetIgnoreLayoutWarnings sets the IgnoreLayoutWarnings property of class StackPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.stackpanel#ignorelayoutwarnings
func (s *StackPanel) SetIgnoreLayoutWarnings(ignoreLayoutWarnings bool) *StackPanel {
	s.p.Set("ignoreLayoutWarnings", ignoreLayoutWarnings)
	return s
}

// IsVertical returns the IsVertical property of class StackPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.stackpanel#isvertical
func (s *StackPanel) IsVertical() bool {
	retVal := s.p.Get("isVertical")
	return retVal.Bool()
}

// SetIsVertical sets the IsVertical property of class StackPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.stackpanel#isvertical
func (s *StackPanel) SetIsVertical(isVertical bool) *StackPanel {
	s.p.Set("isVertical", isVertical)
	return s
}

// Name returns the Name property of class StackPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.stackpanel#name
func (s *StackPanel) Name() string {
	retVal := s.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class StackPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.stackpanel#name
func (s *StackPanel) SetName(name string) *StackPanel {
	s.p.Set("name", name)
	return s
}

// Width returns the Width property of class StackPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.stackpanel#width
func (s *StackPanel) Width() string {
	retVal := s.p.Get("width")
	return retVal.String()
}

// SetWidth sets the Width property of class StackPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.stackpanel#width
func (s *StackPanel) SetWidth(width string) *StackPanel {
	s.p.Set("width", width)
	return s
}
