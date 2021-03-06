// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Color3Gradient represents a babylon.js Color3Gradient.
// Class used to store color 3 gradient
type Color3Gradient struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (c *Color3Gradient) JSObject() js.Value { return c.p }

// Color3Gradient returns a Color3Gradient JavaScript class.
func (ba *Babylon) Color3Gradient() *Color3Gradient {
	p := ba.ctx.Get("Color3Gradient")
	return Color3GradientFromJSObject(p, ba.ctx)
}

// Color3GradientFromJSObject returns a wrapped Color3Gradient JavaScript class.
func Color3GradientFromJSObject(p js.Value, ctx js.Value) *Color3Gradient {
	return &Color3Gradient{p: p, ctx: ctx}
}

// Color3GradientArrayToJSArray returns a JavaScript Array for the wrapped array.
func Color3GradientArrayToJSArray(array []*Color3Gradient) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// Color returns the Color property of class Color3Gradient.
//
// https://doc.babylonjs.com/api/classes/babylon.color3gradient#color
func (c *Color3Gradient) Color() *Color3 {
	retVal := c.p.Get("color")
	return Color3FromJSObject(retVal, c.ctx)
}

// SetColor sets the Color property of class Color3Gradient.
//
// https://doc.babylonjs.com/api/classes/babylon.color3gradient#color
func (c *Color3Gradient) SetColor(color *Color3) *Color3Gradient {
	c.p.Set("color", color.JSObject())
	return c
}

// Gradient returns the Gradient property of class Color3Gradient.
//
// https://doc.babylonjs.com/api/classes/babylon.color3gradient#gradient
func (c *Color3Gradient) Gradient() float64 {
	retVal := c.p.Get("gradient")
	return retVal.Float()
}

// SetGradient sets the Gradient property of class Color3Gradient.
//
// https://doc.babylonjs.com/api/classes/babylon.color3gradient#gradient
func (c *Color3Gradient) SetGradient(gradient float64) *Color3Gradient {
	c.p.Set("gradient", gradient)
	return c
}
