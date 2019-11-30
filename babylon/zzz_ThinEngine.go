// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ThinEngine represents a babylon.js ThinEngine.
// The base engine class (root of all engines)
type ThinEngine struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (t *ThinEngine) JSObject() js.Value { return t.p }

// ThinEngine returns a ThinEngine JavaScript class.
func (b *Babylon) ThinEngine() *ThinEngine {
	p := b.ctx.Get("ThinEngine")
	return ThinEngineFromJSObject(p)
}

// ThinEngineFromJSObject returns a wrapped ThinEngine JavaScript class.
func ThinEngineFromJSObject(p js.Value) *ThinEngine {
	return &ThinEngine{p: p}
}

// NewThinEngineOpts contains optional parameters for NewThinEngine.
type NewThinEngineOpts struct {
	Antialias *bool

	Options js.Value

	AdaptToDeviceRatio *bool
}

// NewThinEngine returns a new ThinEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.thinengine
func (b *Babylon) NewThinEngine(canvasOrContext js.Value, opts *NewThinEngineOpts) *ThinEngine {
	if opts == nil {
		opts = &NewThinEngineOpts{}
	}

	p := b.ctx.Get("ThinEngine").New(canvasOrContext, opts.Antialias.JSObject(), opts.Options, opts.AdaptToDeviceRatio.JSObject())
	return ThinEngineFromJSObject(p)
}

// TODO: methods
