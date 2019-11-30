// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ShadowGeneratorSceneComponent represents a babylon.js ShadowGeneratorSceneComponent.
// Defines the shadow generator component responsible to manage any shadow generators
// in a given scene.
type ShadowGeneratorSceneComponent struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (s *ShadowGeneratorSceneComponent) JSObject() js.Value { return s.p }

// ShadowGeneratorSceneComponent returns a ShadowGeneratorSceneComponent JavaScript class.
func (b *Babylon) ShadowGeneratorSceneComponent() *ShadowGeneratorSceneComponent {
	p := b.ctx.Get("ShadowGeneratorSceneComponent")
	return ShadowGeneratorSceneComponentFromJSObject(p)
}

// ShadowGeneratorSceneComponentFromJSObject returns a wrapped ShadowGeneratorSceneComponent JavaScript class.
func ShadowGeneratorSceneComponentFromJSObject(p js.Value) *ShadowGeneratorSceneComponent {
	return &ShadowGeneratorSceneComponent{p: p}
}

// NewShadowGeneratorSceneComponent returns a new ShadowGeneratorSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowgeneratorscenecomponent
func (b *Babylon) NewShadowGeneratorSceneComponent(scene *Scene) *ShadowGeneratorSceneComponent {
	p := b.ctx.Get("ShadowGeneratorSceneComponent").New(scene.JSObject())
	return ShadowGeneratorSceneComponentFromJSObject(p)
}

// TODO: methods
