// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// EquiRectangularCubeTexture represents a babylon.js EquiRectangularCubeTexture.
// This represents a texture coming from an equirectangular image supported by the web browser canvas.
type EquiRectangularCubeTexture struct {
	*BaseTexture
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (e *EquiRectangularCubeTexture) JSObject() js.Value { return e.p }

// EquiRectangularCubeTexture returns a EquiRectangularCubeTexture JavaScript class.
func (ba *Babylon) EquiRectangularCubeTexture() *EquiRectangularCubeTexture {
	p := ba.ctx.Get("EquiRectangularCubeTexture")
	return EquiRectangularCubeTextureFromJSObject(p, ba.ctx)
}

// EquiRectangularCubeTextureFromJSObject returns a wrapped EquiRectangularCubeTexture JavaScript class.
func EquiRectangularCubeTextureFromJSObject(p js.Value, ctx js.Value) *EquiRectangularCubeTexture {
	return &EquiRectangularCubeTexture{BaseTexture: BaseTextureFromJSObject(p, ctx), ctx: ctx}
}

// EquiRectangularCubeTextureArrayToJSArray returns a JavaScript Array for the wrapped array.
func EquiRectangularCubeTextureArrayToJSArray(array []*EquiRectangularCubeTexture) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewEquiRectangularCubeTextureOpts contains optional parameters for NewEquiRectangularCubeTexture.
type NewEquiRectangularCubeTextureOpts struct {
	NoMipmap   *bool
	GammaSpace *bool
	OnLoad     func()
	OnError    func()
}

// NewEquiRectangularCubeTexture returns a new EquiRectangularCubeTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.equirectangularcubetexture
func (ba *Babylon) NewEquiRectangularCubeTexture(url string, scene *Scene, size float64, opts *NewEquiRectangularCubeTextureOpts) *EquiRectangularCubeTexture {
	if opts == nil {
		opts = &NewEquiRectangularCubeTextureOpts{}
	}

	args := make([]interface{}, 0, 3+4)

	args = append(args, url)
	args = append(args, scene.JSObject())
	args = append(args, size)

	if opts.NoMipmap == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.NoMipmap)
	}
	if opts.GammaSpace == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.GammaSpace)
	}
	if opts.OnLoad == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.OnLoad)
	}
	if opts.OnError == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.OnError)
	}

	p := ba.ctx.Get("EquiRectangularCubeTexture").New(args...)
	return EquiRectangularCubeTextureFromJSObject(p, ba.ctx)
}

// Clone calls the Clone method on the EquiRectangularCubeTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.equirectangularcubetexture#clone
func (e *EquiRectangularCubeTexture) Clone() *EquiRectangularCubeTexture {

	retVal := e.p.Call("clone")
	return EquiRectangularCubeTextureFromJSObject(retVal, e.ctx)
}

// GetClassName calls the GetClassName method on the EquiRectangularCubeTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.equirectangularcubetexture#getclassname
func (e *EquiRectangularCubeTexture) GetClassName() string {

	retVal := e.p.Call("getClassName")
	return retVal.String()
}

/*

// CoordinatesMode returns the CoordinatesMode property of class EquiRectangularCubeTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.equirectangularcubetexture#coordinatesmode
func (e *EquiRectangularCubeTexture) CoordinatesMode(coordinatesMode float64) *EquiRectangularCubeTexture {
	p := ba.ctx.Get("EquiRectangularCubeTexture").New(coordinatesMode)
	return EquiRectangularCubeTextureFromJSObject(p, ba.ctx)
}

// SetCoordinatesMode sets the CoordinatesMode property of class EquiRectangularCubeTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.equirectangularcubetexture#coordinatesmode
func (e *EquiRectangularCubeTexture) SetCoordinatesMode(coordinatesMode float64) *EquiRectangularCubeTexture {
	p := ba.ctx.Get("EquiRectangularCubeTexture").New(coordinatesMode)
	return EquiRectangularCubeTextureFromJSObject(p, ba.ctx)
}

// Url returns the Url property of class EquiRectangularCubeTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.equirectangularcubetexture#url
func (e *EquiRectangularCubeTexture) Url(url string) *EquiRectangularCubeTexture {
	p := ba.ctx.Get("EquiRectangularCubeTexture").New(url)
	return EquiRectangularCubeTextureFromJSObject(p, ba.ctx)
}

// SetUrl sets the Url property of class EquiRectangularCubeTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.equirectangularcubetexture#url
func (e *EquiRectangularCubeTexture) SetUrl(url string) *EquiRectangularCubeTexture {
	p := ba.ctx.Get("EquiRectangularCubeTexture").New(url)
	return EquiRectangularCubeTextureFromJSObject(p, ba.ctx)
}

*/
