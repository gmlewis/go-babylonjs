// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Light represents a babylon.js Light.
// Base class of all the lights in Babylon. It groups all the generic information about lights.
// Lights are used, as you would expect, to affect how meshes are seen, in terms of both illumination and colour.
// All meshes allow light to pass through them unless shadow generation is activated. The default number of lights allowed is four but this can be increased.
type Light struct {
	*Node
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (l *Light) JSObject() js.Value { return l.p }

// Light returns a Light JavaScript class.
func (ba *Babylon) Light() *Light {
	p := ba.ctx.Get("Light")
	return LightFromJSObject(p, ba.ctx)
}

// LightFromJSObject returns a wrapped Light JavaScript class.
func LightFromJSObject(p js.Value, ctx js.Value) *Light {
	return &Light{Node: NodeFromJSObject(p, ctx), ctx: ctx}
}

// NewLight returns a new Light object.
//
// https://doc.babylonjs.com/api/classes/babylon.light
func (ba *Babylon) NewLight(name string, scene *Scene) *Light {
	p := ba.ctx.Get("Light").New(name, scene.JSObject())
	return LightFromJSObject(p, ba.ctx)
}

// TODO: methods
