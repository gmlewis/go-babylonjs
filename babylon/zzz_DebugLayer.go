// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DebugLayer represents a babylon.js DebugLayer.
// The debug layer (aka Inspector) is the go to tool in order to better understand
// what is happening in your scene
//
// See: http://doc.babylonjs.com/features/playground_debuglayer
type DebugLayer struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (d *DebugLayer) JSObject() js.Value { return d.p }

// DebugLayer returns a DebugLayer JavaScript class.
func (ba *Babylon) DebugLayer() *DebugLayer {
	p := ba.ctx.Get("DebugLayer")
	return DebugLayerFromJSObject(p, ba.ctx)
}

// DebugLayerFromJSObject returns a wrapped DebugLayer JavaScript class.
func DebugLayerFromJSObject(p js.Value, ctx js.Value) *DebugLayer {
	return &DebugLayer{p: p, ctx: ctx}
}

// DebugLayerArrayToJSArray returns a JavaScript Array for the wrapped array.
func DebugLayerArrayToJSArray(array []*DebugLayer) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewDebugLayer returns a new DebugLayer object.
//
// https://doc.babylonjs.com/api/classes/babylon.debuglayer
func (ba *Babylon) NewDebugLayer(scene *Scene) *DebugLayer {

	args := make([]interface{}, 0, 1+0)

	args = append(args, scene.JSObject())

	p := ba.ctx.Get("DebugLayer").New(args...)
	return DebugLayerFromJSObject(p, ba.ctx)
}

// Hide calls the Hide method on the DebugLayer object.
//
// https://doc.babylonjs.com/api/classes/babylon.debuglayer#hide
func (d *DebugLayer) Hide() {

	d.p.Call("hide")
}

// IsVisible calls the IsVisible method on the DebugLayer object.
//
// https://doc.babylonjs.com/api/classes/babylon.debuglayer#isvisible
func (d *DebugLayer) IsVisible() bool {

	retVal := d.p.Call("isVisible")
	return retVal.Bool()
}

// DebugLayerSelectOpts contains optional parameters for DebugLayer.Select.
type DebugLayerSelectOpts struct {
	LineContainerTitle *string
}

// Select calls the Select method on the DebugLayer object.
//
// https://doc.babylonjs.com/api/classes/babylon.debuglayer#select
func (d *DebugLayer) Select(entity interface{}, opts *DebugLayerSelectOpts) {
	if opts == nil {
		opts = &DebugLayerSelectOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, entity)

	if opts.LineContainerTitle == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.LineContainerTitle)
	}

	d.p.Call("select", args...)
}

// DebugLayerShowOpts contains optional parameters for DebugLayer.Show.
type DebugLayerShowOpts struct {
	Config js.Value
}

// Show calls the Show method on the DebugLayer object.
//
// https://doc.babylonjs.com/api/classes/babylon.debuglayer#show
func (d *DebugLayer) Show(opts *DebugLayerShowOpts) *Promise {
	if opts == nil {
		opts = &DebugLayerShowOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Config == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Config)
	}

	retVal := d.p.Call("show", args...)
	return PromiseFromJSObject(retVal, d.ctx)
}

/*

// InspectorURL returns the InspectorURL property of class DebugLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.debuglayer#inspectorurl
func (d *DebugLayer) InspectorURL(InspectorURL string) *DebugLayer {
	p := ba.ctx.Get("DebugLayer").New(InspectorURL)
	return DebugLayerFromJSObject(p, ba.ctx)
}

// SetInspectorURL sets the InspectorURL property of class DebugLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.debuglayer#inspectorurl
func (d *DebugLayer) SetInspectorURL(InspectorURL string) *DebugLayer {
	p := ba.ctx.Get("DebugLayer").New(InspectorURL)
	return DebugLayerFromJSObject(p, ba.ctx)
}

// OnPropertyChangedObservable returns the OnPropertyChangedObservable property of class DebugLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.debuglayer#onpropertychangedobservable
func (d *DebugLayer) OnPropertyChangedObservable(onPropertyChangedObservable interface{}) *DebugLayer {
	p := ba.ctx.Get("DebugLayer").New(onPropertyChangedObservable)
	return DebugLayerFromJSObject(p, ba.ctx)
}

// SetOnPropertyChangedObservable sets the OnPropertyChangedObservable property of class DebugLayer.
//
// https://doc.babylonjs.com/api/classes/babylon.debuglayer#onpropertychangedobservable
func (d *DebugLayer) SetOnPropertyChangedObservable(onPropertyChangedObservable interface{}) *DebugLayer {
	p := ba.ctx.Get("DebugLayer").New(onPropertyChangedObservable)
	return DebugLayerFromJSObject(p, ba.ctx)
}

*/
