// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Button3D represents a babylon.js Button3D.
// Class used to create a button in 3D
type Button3D struct {
	*AbstractButton3D
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (b *Button3D) JSObject() js.Value { return b.p }

// Button3D returns a Button3D JavaScript class.
func (gui *GUI) Button3D() *Button3D {
	p := gui.ctx.Get("Button3D")
	return Button3DFromJSObject(p, gui.ctx)
}

// Button3DFromJSObject returns a wrapped Button3D JavaScript class.
func Button3DFromJSObject(p js.Value, ctx js.Value) *Button3D {
	return &Button3D{AbstractButton3D: AbstractButton3DFromJSObject(p, ctx), ctx: ctx}
}

// Button3DArrayToJSArray returns a JavaScript Array for the wrapped array.
func Button3DArrayToJSArray(array []*Button3D) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewButton3DOpts contains optional parameters for NewButton3D.
type NewButton3DOpts struct {
	Name *string
}

// NewButton3D returns a new Button3D object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.button3d#constructor
func (gui *GUI) NewButton3D(opts *NewButton3DOpts) *Button3D {
	if opts == nil {
		opts = &NewButton3DOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Name == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Name)
	}

	p := gui.ctx.Get("Button3D").New(args...)
	return Button3DFromJSObject(p, gui.ctx)
}

// Dispose calls the Dispose method on the Button3D object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.button3d#dispose
func (b *Button3D) Dispose() {

	b.p.Call("dispose")
}

// Content returns the Content property of class Button3D.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.button3d#content
func (b *Button3D) Content() *Control {
	retVal := b.p.Get("content")
	return ControlFromJSObject(retVal, b.ctx)
}

// SetContent sets the Content property of class Button3D.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.button3d#content
func (b *Button3D) SetContent(content *Control) *Button3D {
	b.p.Set("content", content.JSObject())
	return b
}

// ContentResolution returns the ContentResolution property of class Button3D.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.button3d#contentresolution
func (b *Button3D) ContentResolution() int {
	retVal := b.p.Get("contentResolution")
	return retVal.Int()
}

// SetContentResolution sets the ContentResolution property of class Button3D.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.button3d#contentresolution
func (b *Button3D) SetContentResolution(contentResolution int) *Button3D {
	b.p.Set("contentResolution", contentResolution)
	return b
}

// ContentScaleRatio returns the ContentScaleRatio property of class Button3D.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.button3d#contentscaleratio
func (b *Button3D) ContentScaleRatio() float64 {
	retVal := b.p.Get("contentScaleRatio")
	return retVal.Float()
}

// SetContentScaleRatio sets the ContentScaleRatio property of class Button3D.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.button3d#contentscaleratio
func (b *Button3D) SetContentScaleRatio(contentScaleRatio float64) *Button3D {
	b.p.Set("contentScaleRatio", contentScaleRatio)
	return b
}
