// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// LensFlareSystem represents a babylon.js LensFlareSystem.
// This represents a Lens Flare System or the shiny effect created by the light reflection on the  camera lenses.
// It is usually composed of several &lt;code&gt;lensFlare&lt;/code&gt;.

//
// See: http://doc.babylonjs.com/how_to/how_to_use_lens_flares
type LensFlareSystem struct{}

// JSObject returns the underlying js.Value.
func (l *LensFlareSystem) JSObject() js.Value { return l.p }

// LensFlareSystem returns a LensFlareSystem JavaScript class.
func (b *Babylon) LensFlareSystem() *LensFlareSystem {
	p := b.ctx.Get("LensFlareSystem")
	return LensFlareSystemFromJSObject(p)
}

// LensFlareSystemFromJSObject returns a wrapped LensFlareSystem JavaScript class.
func LensFlareSystemFromJSObject(p js.Value) *LensFlareSystem {
	return &LensFlareSystem{p: p}
}

// NewLensFlareSystem returns a new LensFlareSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystem
func (b *Babylon) NewLensFlareSystem(todo parameters) *LensFlareSystem {
	p := b.ctx.Get("LensFlareSystem").New(todo)
	return LensFlareSystemFromJSObject(p)
}

// TODO: methods
