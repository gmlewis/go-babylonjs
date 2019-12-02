// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// HDRCubeTexture represents a babylon.js HDRCubeTexture.
// This represents a texture coming from an HDR input.
//
// The only supported format is currently panorama picture stored in RGBE format.
// Example of such files can be found on HDRLib: <a href="http://hdrlib.com/">http://hdrlib.com/</a>
type HDRCubeTexture struct {
	*BaseTexture
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (h *HDRCubeTexture) JSObject() js.Value { return h.p }

// HDRCubeTexture returns a HDRCubeTexture JavaScript class.
func (ba *Babylon) HDRCubeTexture() *HDRCubeTexture {
	p := ba.ctx.Get("HDRCubeTexture")
	return HDRCubeTextureFromJSObject(p, ba.ctx)
}

// HDRCubeTextureFromJSObject returns a wrapped HDRCubeTexture JavaScript class.
func HDRCubeTextureFromJSObject(p js.Value, ctx js.Value) *HDRCubeTexture {
	return &HDRCubeTexture{BaseTexture: BaseTextureFromJSObject(p, ctx), ctx: ctx}
}

// HDRCubeTextureArrayToJSArray returns a JavaScript Array for the wrapped array.
func HDRCubeTextureArrayToJSArray(array []*HDRCubeTexture) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewHDRCubeTextureOpts contains optional parameters for NewHDRCubeTexture.
type NewHDRCubeTextureOpts struct {
	NoMipmap          *bool
	GenerateHarmonics *bool
	GammaSpace        *bool
	Reserved          *bool
	OnLoad            func()
	OnError           func()
}

// NewHDRCubeTexture returns a new HDRCubeTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetexture
func (ba *Babylon) NewHDRCubeTexture(url string, scene *Scene, size float64, opts *NewHDRCubeTextureOpts) *HDRCubeTexture {
	if opts == nil {
		opts = &NewHDRCubeTextureOpts{}
	}

	args := make([]interface{}, 0, 3+6)

	args = append(args, url)
	args = append(args, scene.JSObject())
	args = append(args, size)

	if opts.NoMipmap == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.NoMipmap)
	}
	if opts.GenerateHarmonics == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.GenerateHarmonics)
	}
	if opts.GammaSpace == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.GammaSpace)
	}
	if opts.Reserved == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Reserved)
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

	p := ba.ctx.Get("HDRCubeTexture").New(args...)
	return HDRCubeTextureFromJSObject(p, ba.ctx)
}

// Clone calls the Clone method on the HDRCubeTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetexture#clone
func (h *HDRCubeTexture) Clone() *HDRCubeTexture {

	retVal := h.p.Call("clone")
	return HDRCubeTextureFromJSObject(retVal, h.ctx)
}

// DelayLoad calls the DelayLoad method on the HDRCubeTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetexture#delayload
func (h *HDRCubeTexture) DelayLoad() {

	h.p.Call("delayLoad")
}

// GetClassName calls the GetClassName method on the HDRCubeTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetexture#getclassname
func (h *HDRCubeTexture) GetClassName() string {

	retVal := h.p.Call("getClassName")
	return retVal.String()
}

// GetReflectionTextureMatrix calls the GetReflectionTextureMatrix method on the HDRCubeTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetexture#getreflectiontexturematrix
func (h *HDRCubeTexture) GetReflectionTextureMatrix() *Matrix {

	retVal := h.p.Call("getReflectionTextureMatrix")
	return MatrixFromJSObject(retVal, h.ctx)
}

// Parse calls the Parse method on the HDRCubeTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetexture#parse
func (h *HDRCubeTexture) Parse(parsedTexture interface{}, scene *Scene, rootUrl string) *HDRCubeTexture {

	args := make([]interface{}, 0, 3+0)

	args = append(args, parsedTexture)
	args = append(args, scene.JSObject())
	args = append(args, rootUrl)

	retVal := h.p.Call("Parse", args...)
	return HDRCubeTextureFromJSObject(retVal, h.ctx)
}

// Serialize calls the Serialize method on the HDRCubeTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetexture#serialize
func (h *HDRCubeTexture) Serialize() interface{} {

	retVal := h.p.Call("serialize")
	return retVal
}

// SetReflectionTextureMatrix calls the SetReflectionTextureMatrix method on the HDRCubeTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetexture#setreflectiontexturematrix
func (h *HDRCubeTexture) SetReflectionTextureMatrix(value *Matrix) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, value.JSObject())

	h.p.Call("setReflectionTextureMatrix", args...)
}

/*

// BoundingBoxPosition returns the BoundingBoxPosition property of class HDRCubeTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetexture#boundingboxposition
func (h *HDRCubeTexture) BoundingBoxPosition(boundingBoxPosition *Vector3) *HDRCubeTexture {
	p := ba.ctx.Get("HDRCubeTexture").New(boundingBoxPosition.JSObject())
	return HDRCubeTextureFromJSObject(p, ba.ctx)
}

// SetBoundingBoxPosition sets the BoundingBoxPosition property of class HDRCubeTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetexture#boundingboxposition
func (h *HDRCubeTexture) SetBoundingBoxPosition(boundingBoxPosition *Vector3) *HDRCubeTexture {
	p := ba.ctx.Get("HDRCubeTexture").New(boundingBoxPosition.JSObject())
	return HDRCubeTextureFromJSObject(p, ba.ctx)
}

// BoundingBoxSize returns the BoundingBoxSize property of class HDRCubeTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetexture#boundingboxsize
func (h *HDRCubeTexture) BoundingBoxSize(boundingBoxSize *Vector3) *HDRCubeTexture {
	p := ba.ctx.Get("HDRCubeTexture").New(boundingBoxSize.JSObject())
	return HDRCubeTextureFromJSObject(p, ba.ctx)
}

// SetBoundingBoxSize sets the BoundingBoxSize property of class HDRCubeTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetexture#boundingboxsize
func (h *HDRCubeTexture) SetBoundingBoxSize(boundingBoxSize *Vector3) *HDRCubeTexture {
	p := ba.ctx.Get("HDRCubeTexture").New(boundingBoxSize.JSObject())
	return HDRCubeTextureFromJSObject(p, ba.ctx)
}

// CoordinatesMode returns the CoordinatesMode property of class HDRCubeTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetexture#coordinatesmode
func (h *HDRCubeTexture) CoordinatesMode(coordinatesMode float64) *HDRCubeTexture {
	p := ba.ctx.Get("HDRCubeTexture").New(coordinatesMode)
	return HDRCubeTextureFromJSObject(p, ba.ctx)
}

// SetCoordinatesMode sets the CoordinatesMode property of class HDRCubeTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetexture#coordinatesmode
func (h *HDRCubeTexture) SetCoordinatesMode(coordinatesMode float64) *HDRCubeTexture {
	p := ba.ctx.Get("HDRCubeTexture").New(coordinatesMode)
	return HDRCubeTextureFromJSObject(p, ba.ctx)
}

// IsBlocking returns the IsBlocking property of class HDRCubeTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetexture#isblocking
func (h *HDRCubeTexture) IsBlocking(isBlocking bool) *HDRCubeTexture {
	p := ba.ctx.Get("HDRCubeTexture").New(isBlocking)
	return HDRCubeTextureFromJSObject(p, ba.ctx)
}

// SetIsBlocking sets the IsBlocking property of class HDRCubeTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetexture#isblocking
func (h *HDRCubeTexture) SetIsBlocking(isBlocking bool) *HDRCubeTexture {
	p := ba.ctx.Get("HDRCubeTexture").New(isBlocking)
	return HDRCubeTextureFromJSObject(p, ba.ctx)
}

// RotationY returns the RotationY property of class HDRCubeTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetexture#rotationy
func (h *HDRCubeTexture) RotationY(rotationY float64) *HDRCubeTexture {
	p := ba.ctx.Get("HDRCubeTexture").New(rotationY)
	return HDRCubeTextureFromJSObject(p, ba.ctx)
}

// SetRotationY sets the RotationY property of class HDRCubeTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetexture#rotationy
func (h *HDRCubeTexture) SetRotationY(rotationY float64) *HDRCubeTexture {
	p := ba.ctx.Get("HDRCubeTexture").New(rotationY)
	return HDRCubeTextureFromJSObject(p, ba.ctx)
}

// Url returns the Url property of class HDRCubeTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetexture#url
func (h *HDRCubeTexture) Url(url string) *HDRCubeTexture {
	p := ba.ctx.Get("HDRCubeTexture").New(url)
	return HDRCubeTextureFromJSObject(p, ba.ctx)
}

// SetUrl sets the Url property of class HDRCubeTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetexture#url
func (h *HDRCubeTexture) SetUrl(url string) *HDRCubeTexture {
	p := ba.ctx.Get("HDRCubeTexture").New(url)
	return HDRCubeTextureFromJSObject(p, ba.ctx)
}

*/
