// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PhysicsVortexEventOptions represents a babylon.js PhysicsVortexEventOptions.
// Options fot the vortex event
//
// See: https://doc.babylonjs.com/how_to/using_the_physics_engine#further-functionality-of-the-impostor-class
type PhysicsVortexEventOptions struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PhysicsVortexEventOptions) JSObject() js.Value { return p.p }

// PhysicsVortexEventOptions returns a PhysicsVortexEventOptions JavaScript class.
func (ba *Babylon) PhysicsVortexEventOptions() *PhysicsVortexEventOptions {
	p := ba.ctx.Get("PhysicsVortexEventOptions")
	return PhysicsVortexEventOptionsFromJSObject(p, ba.ctx)
}

// PhysicsVortexEventOptionsFromJSObject returns a wrapped PhysicsVortexEventOptions JavaScript class.
func PhysicsVortexEventOptionsFromJSObject(p js.Value, ctx js.Value) *PhysicsVortexEventOptions {
	return &PhysicsVortexEventOptions{p: p, ctx: ctx}
}

/*

// CentrifugalForceMultiplier returns the CentrifugalForceMultiplier property of class PhysicsVortexEventOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsvortexeventoptions#centrifugalforcemultiplier
func (p *PhysicsVortexEventOptions) CentrifugalForceMultiplier(centrifugalForceMultiplier float64) *PhysicsVortexEventOptions {
	p := ba.ctx.Get("PhysicsVortexEventOptions").New(centrifugalForceMultiplier)
	return PhysicsVortexEventOptionsFromJSObject(p, ba.ctx)
}

// SetCentrifugalForceMultiplier sets the CentrifugalForceMultiplier property of class PhysicsVortexEventOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsvortexeventoptions#centrifugalforcemultiplier
func (p *PhysicsVortexEventOptions) SetCentrifugalForceMultiplier(centrifugalForceMultiplier float64) *PhysicsVortexEventOptions {
	p := ba.ctx.Get("PhysicsVortexEventOptions").New(centrifugalForceMultiplier)
	return PhysicsVortexEventOptionsFromJSObject(p, ba.ctx)
}

// CentripetalForceMultiplier returns the CentripetalForceMultiplier property of class PhysicsVortexEventOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsvortexeventoptions#centripetalforcemultiplier
func (p *PhysicsVortexEventOptions) CentripetalForceMultiplier(centripetalForceMultiplier float64) *PhysicsVortexEventOptions {
	p := ba.ctx.Get("PhysicsVortexEventOptions").New(centripetalForceMultiplier)
	return PhysicsVortexEventOptionsFromJSObject(p, ba.ctx)
}

// SetCentripetalForceMultiplier sets the CentripetalForceMultiplier property of class PhysicsVortexEventOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsvortexeventoptions#centripetalforcemultiplier
func (p *PhysicsVortexEventOptions) SetCentripetalForceMultiplier(centripetalForceMultiplier float64) *PhysicsVortexEventOptions {
	p := ba.ctx.Get("PhysicsVortexEventOptions").New(centripetalForceMultiplier)
	return PhysicsVortexEventOptionsFromJSObject(p, ba.ctx)
}

// CentripetalForceThreshold returns the CentripetalForceThreshold property of class PhysicsVortexEventOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsvortexeventoptions#centripetalforcethreshold
func (p *PhysicsVortexEventOptions) CentripetalForceThreshold(centripetalForceThreshold float64) *PhysicsVortexEventOptions {
	p := ba.ctx.Get("PhysicsVortexEventOptions").New(centripetalForceThreshold)
	return PhysicsVortexEventOptionsFromJSObject(p, ba.ctx)
}

// SetCentripetalForceThreshold sets the CentripetalForceThreshold property of class PhysicsVortexEventOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsvortexeventoptions#centripetalforcethreshold
func (p *PhysicsVortexEventOptions) SetCentripetalForceThreshold(centripetalForceThreshold float64) *PhysicsVortexEventOptions {
	p := ba.ctx.Get("PhysicsVortexEventOptions").New(centripetalForceThreshold)
	return PhysicsVortexEventOptionsFromJSObject(p, ba.ctx)
}

// Height returns the Height property of class PhysicsVortexEventOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsvortexeventoptions#height
func (p *PhysicsVortexEventOptions) Height(height float64) *PhysicsVortexEventOptions {
	p := ba.ctx.Get("PhysicsVortexEventOptions").New(height)
	return PhysicsVortexEventOptionsFromJSObject(p, ba.ctx)
}

// SetHeight sets the Height property of class PhysicsVortexEventOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsvortexeventoptions#height
func (p *PhysicsVortexEventOptions) SetHeight(height float64) *PhysicsVortexEventOptions {
	p := ba.ctx.Get("PhysicsVortexEventOptions").New(height)
	return PhysicsVortexEventOptionsFromJSObject(p, ba.ctx)
}

// Radius returns the Radius property of class PhysicsVortexEventOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsvortexeventoptions#radius
func (p *PhysicsVortexEventOptions) Radius(radius float64) *PhysicsVortexEventOptions {
	p := ba.ctx.Get("PhysicsVortexEventOptions").New(radius)
	return PhysicsVortexEventOptionsFromJSObject(p, ba.ctx)
}

// SetRadius sets the Radius property of class PhysicsVortexEventOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsvortexeventoptions#radius
func (p *PhysicsVortexEventOptions) SetRadius(radius float64) *PhysicsVortexEventOptions {
	p := ba.ctx.Get("PhysicsVortexEventOptions").New(radius)
	return PhysicsVortexEventOptionsFromJSObject(p, ba.ctx)
}

// Strength returns the Strength property of class PhysicsVortexEventOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsvortexeventoptions#strength
func (p *PhysicsVortexEventOptions) Strength(strength float64) *PhysicsVortexEventOptions {
	p := ba.ctx.Get("PhysicsVortexEventOptions").New(strength)
	return PhysicsVortexEventOptionsFromJSObject(p, ba.ctx)
}

// SetStrength sets the Strength property of class PhysicsVortexEventOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsvortexeventoptions#strength
func (p *PhysicsVortexEventOptions) SetStrength(strength float64) *PhysicsVortexEventOptions {
	p := ba.ctx.Get("PhysicsVortexEventOptions").New(strength)
	return PhysicsVortexEventOptionsFromJSObject(p, ba.ctx)
}

// UpdraftForceMultiplier returns the UpdraftForceMultiplier property of class PhysicsVortexEventOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsvortexeventoptions#updraftforcemultiplier
func (p *PhysicsVortexEventOptions) UpdraftForceMultiplier(updraftForceMultiplier float64) *PhysicsVortexEventOptions {
	p := ba.ctx.Get("PhysicsVortexEventOptions").New(updraftForceMultiplier)
	return PhysicsVortexEventOptionsFromJSObject(p, ba.ctx)
}

// SetUpdraftForceMultiplier sets the UpdraftForceMultiplier property of class PhysicsVortexEventOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsvortexeventoptions#updraftforcemultiplier
func (p *PhysicsVortexEventOptions) SetUpdraftForceMultiplier(updraftForceMultiplier float64) *PhysicsVortexEventOptions {
	p := ba.ctx.Get("PhysicsVortexEventOptions").New(updraftForceMultiplier)
	return PhysicsVortexEventOptionsFromJSObject(p, ba.ctx)
}

*/
