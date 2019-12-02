// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Container represents a babylon.js Container.
// Root class for 2D containers
//
// See: http://doc.babylonjs.com/how_to/gui#containers
type Container struct {
	*Control
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (c *Container) JSObject() js.Value { return c.p }

// Container returns a Container JavaScript class.
func (ba *Babylon) Container() *Container {
	p := ba.ctx.Get("Container")
	return ContainerFromJSObject(p, ba.ctx)
}

// ContainerFromJSObject returns a wrapped Container JavaScript class.
func ContainerFromJSObject(p js.Value, ctx js.Value) *Container {
	return &Container{Control: ControlFromJSObject(p, ctx), ctx: ctx}
}

// ContainerArrayToJSArray returns a JavaScript Array for the wrapped array.
func ContainerArrayToJSArray(array []*Container) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewContainerOpts contains optional parameters for NewContainer.
type NewContainerOpts struct {
	Name *string
}

// NewContainer returns a new Container object.
//
// https://doc.babylonjs.com/api/classes/babylon.container
func (ba *Babylon) NewContainer(opts *NewContainerOpts) *Container {
	if opts == nil {
		opts = &NewContainerOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Name == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Name)
	}

	p := ba.ctx.Get("Container").New(args...)
	return ContainerFromJSObject(p, ba.ctx)
}

// AddControl calls the AddControl method on the Container object.
//
// https://doc.babylonjs.com/api/classes/babylon.container#addcontrol
func (c *Container) AddControl(control *Control) *Container {

	args := make([]interface{}, 0, 1+0)

	args = append(args, control.JSObject())

	retVal := c.p.Call("addControl", args...)
	return ContainerFromJSObject(retVal, c.ctx)
}

// ClearControls calls the ClearControls method on the Container object.
//
// https://doc.babylonjs.com/api/classes/babylon.container#clearcontrols
func (c *Container) ClearControls() *Container {

	retVal := c.p.Call("clearControls")
	return ContainerFromJSObject(retVal, c.ctx)
}

// ContainsControl calls the ContainsControl method on the Container object.
//
// https://doc.babylonjs.com/api/classes/babylon.container#containscontrol
func (c *Container) ContainsControl(control *Control) bool {

	args := make([]interface{}, 0, 1+0)

	args = append(args, control.JSObject())

	retVal := c.p.Call("containsControl", args...)
	return retVal.Bool()
}

// Dispose calls the Dispose method on the Container object.
//
// https://doc.babylonjs.com/api/classes/babylon.container#dispose
func (c *Container) Dispose() {

	c.p.Call("dispose")
}

// GetChildByName calls the GetChildByName method on the Container object.
//
// https://doc.babylonjs.com/api/classes/babylon.container#getchildbyname
func (c *Container) GetChildByName(name string) *Control {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := c.p.Call("getChildByName", args...)
	return ControlFromJSObject(retVal, c.ctx)
}

// GetChildByType calls the GetChildByType method on the Container object.
//
// https://doc.babylonjs.com/api/classes/babylon.container#getchildbytype
func (c *Container) GetChildByType(name string, jsType string) *Control {

	args := make([]interface{}, 0, 2+0)

	args = append(args, name)
	args = append(args, jsType)

	retVal := c.p.Call("getChildByType", args...)
	return ControlFromJSObject(retVal, c.ctx)
}

// ContainerGetDescendantsToRefOpts contains optional parameters for Container.GetDescendantsToRef.
type ContainerGetDescendantsToRefOpts struct {
	DirectDescendantsOnly *bool
	Predicate             func()
}

// GetDescendantsToRef calls the GetDescendantsToRef method on the Container object.
//
// https://doc.babylonjs.com/api/classes/babylon.container#getdescendantstoref
func (c *Container) GetDescendantsToRef(results *Control, opts *ContainerGetDescendantsToRefOpts) {
	if opts == nil {
		opts = &ContainerGetDescendantsToRefOpts{}
	}

	args := make([]interface{}, 0, 1+2)

	args = append(args, results.JSObject())

	if opts.DirectDescendantsOnly == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.DirectDescendantsOnly)
	}
	if opts.Predicate == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Predicate)
	}

	c.p.Call("getDescendantsToRef", args...)
}

// RemoveControl calls the RemoveControl method on the Container object.
//
// https://doc.babylonjs.com/api/classes/babylon.container#removecontrol
func (c *Container) RemoveControl(control *Control) *Container {

	args := make([]interface{}, 0, 1+0)

	args = append(args, control.JSObject())

	retVal := c.p.Call("removeControl", args...)
	return ContainerFromJSObject(retVal, c.ctx)
}

// _flagDescendantsAsMatrixDirty calls the _flagDescendantsAsMatrixDirty method on the Container object.
//
// https://doc.babylonjs.com/api/classes/babylon.container#_flagdescendantsasmatrixdirty
func (c *Container) _flagDescendantsAsMatrixDirty() {

	c.p.Call("_flagDescendantsAsMatrixDirty")
}

/*

// AdaptHeightToChildren returns the AdaptHeightToChildren property of class Container.
//
// https://doc.babylonjs.com/api/classes/babylon.container#adaptheighttochildren
func (c *Container) AdaptHeightToChildren(adaptHeightToChildren bool) *Container {
	p := ba.ctx.Get("Container").New(adaptHeightToChildren)
	return ContainerFromJSObject(p, ba.ctx)
}

// SetAdaptHeightToChildren sets the AdaptHeightToChildren property of class Container.
//
// https://doc.babylonjs.com/api/classes/babylon.container#adaptheighttochildren
func (c *Container) SetAdaptHeightToChildren(adaptHeightToChildren bool) *Container {
	p := ba.ctx.Get("Container").New(adaptHeightToChildren)
	return ContainerFromJSObject(p, ba.ctx)
}

// AdaptWidthToChildren returns the AdaptWidthToChildren property of class Container.
//
// https://doc.babylonjs.com/api/classes/babylon.container#adaptwidthtochildren
func (c *Container) AdaptWidthToChildren(adaptWidthToChildren bool) *Container {
	p := ba.ctx.Get("Container").New(adaptWidthToChildren)
	return ContainerFromJSObject(p, ba.ctx)
}

// SetAdaptWidthToChildren sets the AdaptWidthToChildren property of class Container.
//
// https://doc.babylonjs.com/api/classes/babylon.container#adaptwidthtochildren
func (c *Container) SetAdaptWidthToChildren(adaptWidthToChildren bool) *Container {
	p := ba.ctx.Get("Container").New(adaptWidthToChildren)
	return ContainerFromJSObject(p, ba.ctx)
}

// Background returns the Background property of class Container.
//
// https://doc.babylonjs.com/api/classes/babylon.container#background
func (c *Container) Background(background string) *Container {
	p := ba.ctx.Get("Container").New(background)
	return ContainerFromJSObject(p, ba.ctx)
}

// SetBackground sets the Background property of class Container.
//
// https://doc.babylonjs.com/api/classes/babylon.container#background
func (c *Container) SetBackground(background string) *Container {
	p := ba.ctx.Get("Container").New(background)
	return ContainerFromJSObject(p, ba.ctx)
}

// Children returns the Children property of class Container.
//
// https://doc.babylonjs.com/api/classes/babylon.container#children
func (c *Container) Children(children *Control) *Container {
	p := ba.ctx.Get("Container").New(children.JSObject())
	return ContainerFromJSObject(p, ba.ctx)
}

// SetChildren sets the Children property of class Container.
//
// https://doc.babylonjs.com/api/classes/babylon.container#children
func (c *Container) SetChildren(children *Control) *Container {
	p := ba.ctx.Get("Container").New(children.JSObject())
	return ContainerFromJSObject(p, ba.ctx)
}

// LogLayoutCycleErrors returns the LogLayoutCycleErrors property of class Container.
//
// https://doc.babylonjs.com/api/classes/babylon.container#loglayoutcycleerrors
func (c *Container) LogLayoutCycleErrors(logLayoutCycleErrors bool) *Container {
	p := ba.ctx.Get("Container").New(logLayoutCycleErrors)
	return ContainerFromJSObject(p, ba.ctx)
}

// SetLogLayoutCycleErrors sets the LogLayoutCycleErrors property of class Container.
//
// https://doc.babylonjs.com/api/classes/babylon.container#loglayoutcycleerrors
func (c *Container) SetLogLayoutCycleErrors(logLayoutCycleErrors bool) *Container {
	p := ba.ctx.Get("Container").New(logLayoutCycleErrors)
	return ContainerFromJSObject(p, ba.ctx)
}

// MaxLayoutCycle returns the MaxLayoutCycle property of class Container.
//
// https://doc.babylonjs.com/api/classes/babylon.container#maxlayoutcycle
func (c *Container) MaxLayoutCycle(maxLayoutCycle float64) *Container {
	p := ba.ctx.Get("Container").New(maxLayoutCycle)
	return ContainerFromJSObject(p, ba.ctx)
}

// SetMaxLayoutCycle sets the MaxLayoutCycle property of class Container.
//
// https://doc.babylonjs.com/api/classes/babylon.container#maxlayoutcycle
func (c *Container) SetMaxLayoutCycle(maxLayoutCycle float64) *Container {
	p := ba.ctx.Get("Container").New(maxLayoutCycle)
	return ContainerFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class Container.
//
// https://doc.babylonjs.com/api/classes/babylon.container#name
func (c *Container) Name(name string) *Container {
	p := ba.ctx.Get("Container").New(name)
	return ContainerFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class Container.
//
// https://doc.babylonjs.com/api/classes/babylon.container#name
func (c *Container) SetName(name string) *Container {
	p := ba.ctx.Get("Container").New(name)
	return ContainerFromJSObject(p, ba.ctx)
}

*/
