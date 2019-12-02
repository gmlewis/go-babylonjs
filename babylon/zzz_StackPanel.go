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
func (ba *Babylon) StackPanel() *StackPanel {
	p := ba.ctx.Get("StackPanel")
	return StackPanelFromJSObject(p, ba.ctx)
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
// https://doc.babylonjs.com/api/classes/babylon.stackpanel
func (ba *Babylon) NewStackPanel(opts *NewStackPanelOpts) *StackPanel {
	if opts == nil {
		opts = &NewStackPanelOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Name == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Name)
	}

	p := ba.ctx.Get("StackPanel").New(args...)
	return StackPanelFromJSObject(p, ba.ctx)
}

/*

// Height returns the Height property of class StackPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.stackpanel#height
func (s *StackPanel) Height(height string) *StackPanel {
	p := ba.ctx.Get("StackPanel").New(height)
	return StackPanelFromJSObject(p, ba.ctx)
}

// SetHeight sets the Height property of class StackPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.stackpanel#height
func (s *StackPanel) SetHeight(height string) *StackPanel {
	p := ba.ctx.Get("StackPanel").New(height)
	return StackPanelFromJSObject(p, ba.ctx)
}

// IgnoreLayoutWarnings returns the IgnoreLayoutWarnings property of class StackPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.stackpanel#ignorelayoutwarnings
func (s *StackPanel) IgnoreLayoutWarnings(ignoreLayoutWarnings bool) *StackPanel {
	p := ba.ctx.Get("StackPanel").New(ignoreLayoutWarnings)
	return StackPanelFromJSObject(p, ba.ctx)
}

// SetIgnoreLayoutWarnings sets the IgnoreLayoutWarnings property of class StackPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.stackpanel#ignorelayoutwarnings
func (s *StackPanel) SetIgnoreLayoutWarnings(ignoreLayoutWarnings bool) *StackPanel {
	p := ba.ctx.Get("StackPanel").New(ignoreLayoutWarnings)
	return StackPanelFromJSObject(p, ba.ctx)
}

// IsVertical returns the IsVertical property of class StackPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.stackpanel#isvertical
func (s *StackPanel) IsVertical(isVertical bool) *StackPanel {
	p := ba.ctx.Get("StackPanel").New(isVertical)
	return StackPanelFromJSObject(p, ba.ctx)
}

// SetIsVertical sets the IsVertical property of class StackPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.stackpanel#isvertical
func (s *StackPanel) SetIsVertical(isVertical bool) *StackPanel {
	p := ba.ctx.Get("StackPanel").New(isVertical)
	return StackPanelFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class StackPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.stackpanel#name
func (s *StackPanel) Name(name string) *StackPanel {
	p := ba.ctx.Get("StackPanel").New(name)
	return StackPanelFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class StackPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.stackpanel#name
func (s *StackPanel) SetName(name string) *StackPanel {
	p := ba.ctx.Get("StackPanel").New(name)
	return StackPanelFromJSObject(p, ba.ctx)
}

// Width returns the Width property of class StackPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.stackpanel#width
func (s *StackPanel) Width(width string) *StackPanel {
	p := ba.ctx.Get("StackPanel").New(width)
	return StackPanelFromJSObject(p, ba.ctx)
}

// SetWidth sets the Width property of class StackPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.stackpanel#width
func (s *StackPanel) SetWidth(width string) *StackPanel {
	p := ba.ctx.Get("StackPanel").New(width)
	return StackPanelFromJSObject(p, ba.ctx)
}

*/
