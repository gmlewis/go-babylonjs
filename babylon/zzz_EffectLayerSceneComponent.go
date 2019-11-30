// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// EffectLayerSceneComponent represents a babylon.js EffectLayerSceneComponent.
// Defines the layer scene component responsible to manage any effect layers
// in a given scene.
type EffectLayerSceneComponent struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (e *EffectLayerSceneComponent) JSObject() js.Value { return e.p }

// EffectLayerSceneComponent returns a EffectLayerSceneComponent JavaScript class.
func (ba *Babylon) EffectLayerSceneComponent() *EffectLayerSceneComponent {
	p := ba.ctx.Get("EffectLayerSceneComponent")
	return EffectLayerSceneComponentFromJSObject(p, ba.ctx)
}

// EffectLayerSceneComponentFromJSObject returns a wrapped EffectLayerSceneComponent JavaScript class.
func EffectLayerSceneComponentFromJSObject(p js.Value, ctx js.Value) *EffectLayerSceneComponent {
	return &EffectLayerSceneComponent{p: p, ctx: ctx}
}

// NewEffectLayerSceneComponent returns a new EffectLayerSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.effectlayerscenecomponent
func (ba *Babylon) NewEffectLayerSceneComponent(scene *Scene) *EffectLayerSceneComponent {
	p := ba.ctx.Get("EffectLayerSceneComponent").New(scene.JSObject())
	return EffectLayerSceneComponentFromJSObject(p, ba.ctx)
}

// TODO: methods
