// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ProceduralTextureSceneComponent represents a babylon.js ProceduralTextureSceneComponent.
// Defines the Procedural Texture scene component responsible to manage any Procedural Texture
// in a given scene.
type ProceduralTextureSceneComponent struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *ProceduralTextureSceneComponent) JSObject() js.Value { return p.p }

// ProceduralTextureSceneComponent returns a ProceduralTextureSceneComponent JavaScript class.
func (ba *Babylon) ProceduralTextureSceneComponent() *ProceduralTextureSceneComponent {
	p := ba.ctx.Get("ProceduralTextureSceneComponent")
	return ProceduralTextureSceneComponentFromJSObject(p, ba.ctx)
}

// ProceduralTextureSceneComponentFromJSObject returns a wrapped ProceduralTextureSceneComponent JavaScript class.
func ProceduralTextureSceneComponentFromJSObject(p js.Value, ctx js.Value) *ProceduralTextureSceneComponent {
	return &ProceduralTextureSceneComponent{p: p, ctx: ctx}
}

// NewProceduralTextureSceneComponent returns a new ProceduralTextureSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.proceduraltexturescenecomponent
func (ba *Babylon) NewProceduralTextureSceneComponent(scene *Scene) *ProceduralTextureSceneComponent {
	p := ba.ctx.Get("ProceduralTextureSceneComponent").New(scene.JSObject())
	return ProceduralTextureSceneComponentFromJSObject(p, ba.ctx)
}

// TODO: methods
