// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// LensFlareSystemSceneComponent represents a babylon.js LensFlareSystemSceneComponent.
// Defines the lens flare scene component responsible to manage any lens flares
// in a given scene.
type LensFlareSystemSceneComponent struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (l *LensFlareSystemSceneComponent) JSObject() js.Value { return l.p }

// LensFlareSystemSceneComponent returns a LensFlareSystemSceneComponent JavaScript class.
func (b *Babylon) LensFlareSystemSceneComponent() *LensFlareSystemSceneComponent {
	p := b.ctx.Get("LensFlareSystemSceneComponent")
	return LensFlareSystemSceneComponentFromJSObject(p)
}

// LensFlareSystemSceneComponentFromJSObject returns a wrapped LensFlareSystemSceneComponent JavaScript class.
func LensFlareSystemSceneComponentFromJSObject(p js.Value) *LensFlareSystemSceneComponent {
	return &LensFlareSystemSceneComponent{p: p}
}

// NewLensFlareSystemSceneComponent returns a new LensFlareSystemSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystemscenecomponent
func (b *Babylon) NewLensFlareSystemSceneComponent(scene *Scene) *LensFlareSystemSceneComponent {
	p := b.ctx.Get("LensFlareSystemSceneComponent").New(scene.JSObject())
	return LensFlareSystemSceneComponentFromJSObject(p)
}

// TODO: methods
