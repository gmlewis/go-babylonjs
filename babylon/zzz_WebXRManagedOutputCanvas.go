// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// WebXRManagedOutputCanvas represents a babylon.js WebXRManagedOutputCanvas.
// Creates a canvas that is added/removed from the webpage when entering/exiting XR
type WebXRManagedOutputCanvas struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (w *WebXRManagedOutputCanvas) JSObject() js.Value { return w.p }

// WebXRManagedOutputCanvas returns a WebXRManagedOutputCanvas JavaScript class.
func (ba *Babylon) WebXRManagedOutputCanvas() *WebXRManagedOutputCanvas {
	p := ba.ctx.Get("WebXRManagedOutputCanvas")
	return WebXRManagedOutputCanvasFromJSObject(p, ba.ctx)
}

// WebXRManagedOutputCanvasFromJSObject returns a wrapped WebXRManagedOutputCanvas JavaScript class.
func WebXRManagedOutputCanvasFromJSObject(p js.Value, ctx js.Value) *WebXRManagedOutputCanvas {
	return &WebXRManagedOutputCanvas{p: p, ctx: ctx}
}

// NewWebXRManagedOutputCanvasOpts contains optional parameters for NewWebXRManagedOutputCanvas.
type NewWebXRManagedOutputCanvasOpts struct {
	Canvas                   js.Value
	OnStateChangedObservable *Observable
	Configuration            *WebXRState
}

// NewWebXRManagedOutputCanvas returns a new WebXRManagedOutputCanvas object.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrmanagedoutputcanvas
func (ba *Babylon) NewWebXRManagedOutputCanvas(engine *ThinEngine, opts *NewWebXRManagedOutputCanvasOpts) *WebXRManagedOutputCanvas {
	if opts == nil {
		opts = &NewWebXRManagedOutputCanvasOpts{}
	}

	args := make([]interface{}, 0, 1+3)

	args = append(args, engine.JSObject())

	if opts.Canvas == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Canvas)
	}
	if opts.OnStateChangedObservable == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.OnStateChangedObservable.JSObject())
	}
	if opts.Configuration == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Configuration.JSObject())
	}

	p := ba.ctx.Get("WebXRManagedOutputCanvas").New(args...)
	return WebXRManagedOutputCanvasFromJSObject(p, ba.ctx)
}

// Dispose calls the Dispose method on the WebXRManagedOutputCanvas object.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrmanagedoutputcanvas#dispose
func (w *WebXRManagedOutputCanvas) Dispose() {

	args := make([]interface{}, 0, 0+0)

	w.p.Call("dispose", args...)
}

// InitializeXRLayerAsync calls the InitializeXRLayerAsync method on the WebXRManagedOutputCanvas object.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrmanagedoutputcanvas#initializexrlayerasync
func (w *WebXRManagedOutputCanvas) InitializeXRLayerAsync(xrSession interface{}) interface{} {

	args := make([]interface{}, 0, 1+0)

	args = append(args, xrSession)

	retVal := w.p.Call("initializeXRLayerAsync", args...)
	return retVal
}

/*

// CanvasContext returns the CanvasContext property of class WebXRManagedOutputCanvas.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrmanagedoutputcanvas#canvascontext
func (w *WebXRManagedOutputCanvas) CanvasContext(canvasContext js.Value) *WebXRManagedOutputCanvas {
	p := ba.ctx.Get("WebXRManagedOutputCanvas").New(canvasContext)
	return WebXRManagedOutputCanvasFromJSObject(p, ba.ctx)
}

// SetCanvasContext sets the CanvasContext property of class WebXRManagedOutputCanvas.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrmanagedoutputcanvas#canvascontext
func (w *WebXRManagedOutputCanvas) SetCanvasContext(canvasContext js.Value) *WebXRManagedOutputCanvas {
	p := ba.ctx.Get("WebXRManagedOutputCanvas").New(canvasContext)
	return WebXRManagedOutputCanvasFromJSObject(p, ba.ctx)
}

// XrLayer returns the XrLayer property of class WebXRManagedOutputCanvas.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrmanagedoutputcanvas#xrlayer
func (w *WebXRManagedOutputCanvas) XrLayer(xrLayer *XRWebGLLayer) *WebXRManagedOutputCanvas {
	p := ba.ctx.Get("WebXRManagedOutputCanvas").New(xrLayer.JSObject())
	return WebXRManagedOutputCanvasFromJSObject(p, ba.ctx)
}

// SetXrLayer sets the XrLayer property of class WebXRManagedOutputCanvas.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrmanagedoutputcanvas#xrlayer
func (w *WebXRManagedOutputCanvas) SetXrLayer(xrLayer *XRWebGLLayer) *WebXRManagedOutputCanvas {
	p := ba.ctx.Get("WebXRManagedOutputCanvas").New(xrLayer.JSObject())
	return WebXRManagedOutputCanvasFromJSObject(p, ba.ctx)
}

*/
