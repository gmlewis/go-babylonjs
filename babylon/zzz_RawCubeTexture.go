// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// RawCubeTexture represents a babylon.js RawCubeTexture.
// Raw cube texture where the raw buffers are passed in
type RawCubeTexture struct {
	*CubeTexture
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (r *RawCubeTexture) JSObject() js.Value { return r.p }

// RawCubeTexture returns a RawCubeTexture JavaScript class.
func (ba *Babylon) RawCubeTexture() *RawCubeTexture {
	p := ba.ctx.Get("RawCubeTexture")
	return RawCubeTextureFromJSObject(p, ba.ctx)
}

// RawCubeTextureFromJSObject returns a wrapped RawCubeTexture JavaScript class.
func RawCubeTextureFromJSObject(p js.Value, ctx js.Value) *RawCubeTexture {
	return &RawCubeTexture{CubeTexture: CubeTextureFromJSObject(p, ctx), ctx: ctx}
}

// RawCubeTextureArrayToJSArray returns a JavaScript Array for the wrapped array.
func RawCubeTextureArrayToJSArray(array []*RawCubeTexture) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewRawCubeTextureOpts contains optional parameters for NewRawCubeTexture.
type NewRawCubeTextureOpts struct {
	Format          *float64
	Type            *float64
	GenerateMipMaps *bool
	InvertY         *bool
	SamplingMode    *float64
	Compression     *string
}

// NewRawCubeTexture returns a new RawCubeTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.rawcubetexture#constructor
func (ba *Babylon) NewRawCubeTexture(scene *Scene, data js.Value, size float64, opts *NewRawCubeTextureOpts) *RawCubeTexture {
	if opts == nil {
		opts = &NewRawCubeTextureOpts{}
	}

	args := make([]interface{}, 0, 3+6)

	args = append(args, scene.JSObject())
	args = append(args, data)
	args = append(args, size)

	if opts.Format == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Format)
	}
	if opts.Type == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Type)
	}
	if opts.GenerateMipMaps == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.GenerateMipMaps)
	}
	if opts.InvertY == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.InvertY)
	}
	if opts.SamplingMode == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.SamplingMode)
	}
	if opts.Compression == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Compression)
	}

	p := ba.ctx.Get("RawCubeTexture").New(args...)
	return RawCubeTextureFromJSObject(p, ba.ctx)
}

// Clone calls the Clone method on the RawCubeTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.rawcubetexture#clone
func (r *RawCubeTexture) Clone() *CubeTexture {

	retVal := r.p.Call("clone")
	return CubeTextureFromJSObject(retVal, r.ctx)
}

// RawCubeTextureUpdateOpts contains optional parameters for RawCubeTexture.Update.
type RawCubeTextureUpdateOpts struct {
	Compression *string
}

// Update calls the Update method on the RawCubeTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.rawcubetexture#update
func (r *RawCubeTexture) Update(data js.Value, format float64, jsType float64, invertY bool, opts *RawCubeTextureUpdateOpts) {
	if opts == nil {
		opts = &RawCubeTextureUpdateOpts{}
	}

	args := make([]interface{}, 0, 4+1)

	args = append(args, data)

	args = append(args, format)

	args = append(args, jsType)

	args = append(args, invertY)

	if opts.Compression == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Compression)
	}

	r.p.Call("update", args...)
}

// RawCubeTextureUpdateRGBDAsyncOpts contains optional parameters for RawCubeTexture.UpdateRGBDAsync.
type RawCubeTextureUpdateRGBDAsyncOpts struct {
	SphericalPolynomial *SphericalPolynomial
	LodScale            *float64
	LodOffset           *float64
}

// UpdateRGBDAsync calls the UpdateRGBDAsync method on the RawCubeTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.rawcubetexture#updatergbdasync
func (r *RawCubeTexture) UpdateRGBDAsync(data js.Value, opts *RawCubeTextureUpdateRGBDAsyncOpts) *Promise {
	if opts == nil {
		opts = &RawCubeTextureUpdateRGBDAsyncOpts{}
	}

	args := make([]interface{}, 0, 1+3)

	args = append(args, data)

	if opts.SphericalPolynomial == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.SphericalPolynomial.JSObject())
	}
	if opts.LodScale == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.LodScale)
	}
	if opts.LodOffset == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.LodOffset)
	}

	retVal := r.p.Call("updateRGBDAsync", args...)
	return PromiseFromJSObject(retVal, r.ctx)
}
