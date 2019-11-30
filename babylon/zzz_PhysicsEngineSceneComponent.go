// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PhysicsEngineSceneComponent represents a babylon.js PhysicsEngineSceneComponent.
// Defines the physics engine scene component responsible to manage a physics engine
type PhysicsEngineSceneComponent struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PhysicsEngineSceneComponent) JSObject() js.Value { return p.p }

// PhysicsEngineSceneComponent returns a PhysicsEngineSceneComponent JavaScript class.
func (ba *Babylon) PhysicsEngineSceneComponent() *PhysicsEngineSceneComponent {
	p := ba.ctx.Get("PhysicsEngineSceneComponent")
	return PhysicsEngineSceneComponentFromJSObject(p, ba.ctx)
}

// PhysicsEngineSceneComponentFromJSObject returns a wrapped PhysicsEngineSceneComponent JavaScript class.
func PhysicsEngineSceneComponentFromJSObject(p js.Value, ctx js.Value) *PhysicsEngineSceneComponent {
	return &PhysicsEngineSceneComponent{p: p, ctx: ctx}
}

// NewPhysicsEngineSceneComponent returns a new PhysicsEngineSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsenginescenecomponent
func (ba *Babylon) NewPhysicsEngineSceneComponent(scene *Scene) *PhysicsEngineSceneComponent {
	p := ba.ctx.Get("PhysicsEngineSceneComponent").New(scene.JSObject())
	return PhysicsEngineSceneComponentFromJSObject(p, ba.ctx)
}

// TODO: methods
