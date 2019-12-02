// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DynamicTexture represents a babylon.js DynamicTexture.
// A class extending Texture allowing drawing on a texture
//
// See: http://doc.babylonjs.com/how_to/dynamictexture
type DynamicTexture struct {
	*Texture
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (d *DynamicTexture) JSObject() js.Value { return d.p }

// DynamicTexture returns a DynamicTexture JavaScript class.
func (ba *Babylon) DynamicTexture() *DynamicTexture {
	p := ba.ctx.Get("DynamicTexture")
	return DynamicTextureFromJSObject(p, ba.ctx)
}

// DynamicTextureFromJSObject returns a wrapped DynamicTexture JavaScript class.
func DynamicTextureFromJSObject(p js.Value, ctx js.Value) *DynamicTexture {
	return &DynamicTexture{Texture: TextureFromJSObject(p, ctx), ctx: ctx}
}

// DynamicTextureArrayToJSArray returns a JavaScript Array for the wrapped array.
func DynamicTextureArrayToJSArray(array []*DynamicTexture) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewDynamicTextureOpts contains optional parameters for NewDynamicTexture.
type NewDynamicTextureOpts struct {
	SamplingMode *float64
	Format       *float64
}

// NewDynamicTexture returns a new DynamicTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.dynamictexture
func (ba *Babylon) NewDynamicTexture(name string, options interface{}, scene *Scene, generateMipMaps bool, opts *NewDynamicTextureOpts) *DynamicTexture {
	if opts == nil {
		opts = &NewDynamicTextureOpts{}
	}

	args := make([]interface{}, 0, 4+2)

	args = append(args, name)
	args = append(args, options)
	args = append(args, scene.JSObject())
	args = append(args, generateMipMaps)

	if opts.SamplingMode == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.SamplingMode)
	}
	if opts.Format == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Format)
	}

	p := ba.ctx.Get("DynamicTexture").New(args...)
	return DynamicTextureFromJSObject(p, ba.ctx)
}

// Clear calls the Clear method on the DynamicTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.dynamictexture#clear
func (d *DynamicTexture) Clear() {

	d.p.Call("clear")
}

// Clone calls the Clone method on the DynamicTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.dynamictexture#clone
func (d *DynamicTexture) Clone() *DynamicTexture {

	retVal := d.p.Call("clone")
	return DynamicTextureFromJSObject(retVal, d.ctx)
}

// DynamicTextureDrawTextOpts contains optional parameters for DynamicTexture.DrawText.
type DynamicTextureDrawTextOpts struct {
	InvertY *bool
	Update  *bool
}

// DrawText calls the DrawText method on the DynamicTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.dynamictexture#drawtext
func (d *DynamicTexture) DrawText(text string, x float64, y float64, font string, color string, clearColor string, opts *DynamicTextureDrawTextOpts) {
	if opts == nil {
		opts = &DynamicTextureDrawTextOpts{}
	}

	args := make([]interface{}, 0, 6+2)

	args = append(args, text)
	args = append(args, x)
	args = append(args, y)
	args = append(args, font)
	args = append(args, color)
	args = append(args, clearColor)

	if opts.InvertY == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.InvertY)
	}
	if opts.Update == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Update)
	}

	d.p.Call("drawText", args...)
}

// GetClassName calls the GetClassName method on the DynamicTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.dynamictexture#getclassname
func (d *DynamicTexture) GetClassName() string {

	retVal := d.p.Call("getClassName")
	return retVal.String()
}

// GetContext calls the GetContext method on the DynamicTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.dynamictexture#getcontext
func (d *DynamicTexture) GetContext() js.Value {

	retVal := d.p.Call("getContext")
	return retVal
}

// Scale calls the Scale method on the DynamicTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.dynamictexture#scale
func (d *DynamicTexture) Scale(ratio float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, ratio)

	d.p.Call("scale", args...)
}

// ScaleTo calls the ScaleTo method on the DynamicTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.dynamictexture#scaleto
func (d *DynamicTexture) ScaleTo(width float64, height float64) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, width)
	args = append(args, height)

	d.p.Call("scaleTo", args...)
}

// Serialize calls the Serialize method on the DynamicTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.dynamictexture#serialize
func (d *DynamicTexture) Serialize() interface{} {

	retVal := d.p.Call("serialize")
	return retVal
}

// DynamicTextureUpdateOpts contains optional parameters for DynamicTexture.Update.
type DynamicTextureUpdateOpts struct {
	InvertY     *bool
	PremulAlpha *bool
}

// Update calls the Update method on the DynamicTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.dynamictexture#update
func (d *DynamicTexture) Update(opts *DynamicTextureUpdateOpts) {
	if opts == nil {
		opts = &DynamicTextureUpdateOpts{}
	}

	args := make([]interface{}, 0, 0+2)

	if opts.InvertY == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.InvertY)
	}
	if opts.PremulAlpha == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.PremulAlpha)
	}

	d.p.Call("update", args...)
}

/*

// CanRescale returns the CanRescale property of class DynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.dynamictexture#canrescale
func (d *DynamicTexture) CanRescale(canRescale bool) *DynamicTexture {
	p := ba.ctx.Get("DynamicTexture").New(canRescale)
	return DynamicTextureFromJSObject(p, ba.ctx)
}

// SetCanRescale sets the CanRescale property of class DynamicTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.dynamictexture#canrescale
func (d *DynamicTexture) SetCanRescale(canRescale bool) *DynamicTexture {
	p := ba.ctx.Get("DynamicTexture").New(canRescale)
	return DynamicTextureFromJSObject(p, ba.ctx)
}

*/
