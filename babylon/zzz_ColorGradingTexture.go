// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ColorGradingTexture represents a babylon.js ColorGradingTexture.
// This represents a color grading texture. This acts as a lookup table LUT, useful during post process
// It can help converting any input color in a desired output one. This can then be used to create effects
// from sepia, black and white to sixties or futuristic rendering...
//
// The only supported format is currently 3dl.
// More information on LUT: <a href="https://en.wikipedia.org/wiki/3D_lookup_table">https://en.wikipedia.org/wiki/3D_lookup_table</a>
type ColorGradingTexture struct {
	*BaseTexture
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (c *ColorGradingTexture) JSObject() js.Value { return c.p }

// ColorGradingTexture returns a ColorGradingTexture JavaScript class.
func (ba *Babylon) ColorGradingTexture() *ColorGradingTexture {
	p := ba.ctx.Get("ColorGradingTexture")
	return ColorGradingTextureFromJSObject(p, ba.ctx)
}

// ColorGradingTextureFromJSObject returns a wrapped ColorGradingTexture JavaScript class.
func ColorGradingTextureFromJSObject(p js.Value, ctx js.Value) *ColorGradingTexture {
	return &ColorGradingTexture{BaseTexture: BaseTextureFromJSObject(p, ctx), ctx: ctx}
}

// ColorGradingTextureArrayToJSArray returns a JavaScript Array for the wrapped array.
func ColorGradingTextureArrayToJSArray(array []*ColorGradingTexture) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewColorGradingTexture returns a new ColorGradingTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.colorgradingtexture#constructor
func (ba *Babylon) NewColorGradingTexture(url string, scene *Scene) *ColorGradingTexture {

	args := make([]interface{}, 0, 2+0)

	args = append(args, url)
	args = append(args, scene.JSObject())

	p := ba.ctx.Get("ColorGradingTexture").New(args...)
	return ColorGradingTextureFromJSObject(p, ba.ctx)
}

// Clone calls the Clone method on the ColorGradingTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.colorgradingtexture#clone
func (c *ColorGradingTexture) Clone() *ColorGradingTexture {

	retVal := c.p.Call("clone")
	return ColorGradingTextureFromJSObject(retVal, c.ctx)
}

// DelayLoad calls the DelayLoad method on the ColorGradingTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.colorgradingtexture#delayload
func (c *ColorGradingTexture) DelayLoad() {

	c.p.Call("delayLoad")
}

// GetTextureMatrix calls the GetTextureMatrix method on the ColorGradingTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.colorgradingtexture#gettexturematrix
func (c *ColorGradingTexture) GetTextureMatrix() *Matrix {

	retVal := c.p.Call("getTextureMatrix")
	return MatrixFromJSObject(retVal, c.ctx)
}

// Parse calls the Parse method on the ColorGradingTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.colorgradingtexture#parse
func (c *ColorGradingTexture) Parse(parsedTexture JSObject, scene *Scene) *ColorGradingTexture {

	args := make([]interface{}, 0, 2+0)

	if parsedTexture == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, parsedTexture.JSObject())
	}

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
	}

	retVal := c.p.Call("Parse", args...)
	return ColorGradingTextureFromJSObject(retVal, c.ctx)
}

// Serialize calls the Serialize method on the ColorGradingTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.colorgradingtexture#serialize
func (c *ColorGradingTexture) Serialize() js.Value {

	retVal := c.p.Call("serialize")
	return retVal
}

// Url returns the Url property of class ColorGradingTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.colorgradingtexture#url
func (c *ColorGradingTexture) Url() string {
	retVal := c.p.Get("url")
	return retVal.String()
}

// SetUrl sets the Url property of class ColorGradingTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.colorgradingtexture#url
func (c *ColorGradingTexture) SetUrl(url string) *ColorGradingTexture {
	c.p.Set("url", url)
	return c
}
