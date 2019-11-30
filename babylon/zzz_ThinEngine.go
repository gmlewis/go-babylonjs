// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ThinEngine represents a babylon.js ThinEngine.
// The base engine class (root of all engines)
type ThinEngine struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (t *ThinEngine) JSObject() js.Value { return t.p }

// ThinEngine returns a ThinEngine JavaScript class.
func (ba *Babylon) ThinEngine() *ThinEngine {
	p := ba.ctx.Get("ThinEngine")
	return ThinEngineFromJSObject(p, ba.ctx)
}

// ThinEngineFromJSObject returns a wrapped ThinEngine JavaScript class.
func ThinEngineFromJSObject(p js.Value, ctx js.Value) *ThinEngine {
	return &ThinEngine{p: p, ctx: ctx}
}

// NewThinEngineOpts contains optional parameters for NewThinEngine.
type NewThinEngineOpts struct {
	Antialias *JSBool

	Options *JSValue

	AdaptToDeviceRatio *JSBool
}

// NewThinEngine returns a new ThinEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.thinengine
func (ba *Babylon) NewThinEngine(canvasOrContext js.Value, opts *NewThinEngineOpts) *ThinEngine {
	if opts == nil {
		opts = &NewThinEngineOpts{}
	}

	p := ba.ctx.Get("ThinEngine").New(canvasOrContext, opts.Antialias.JSObject(), opts.Options.JSObject(), opts.AdaptToDeviceRatio.JSObject())
	return ThinEngineFromJSObject(p, ba.ctx)
}

// TODO: methods
