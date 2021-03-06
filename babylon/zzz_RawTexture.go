// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// RawTexture represents a babylon.js RawTexture.
// Raw texture can help creating a texture directly from an array of data.
// This can be super useful if you either get the data from an uncompressed source or
// if you wish to create your texture pixel by pixel.
type RawTexture struct {
	*Texture
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (r *RawTexture) JSObject() js.Value { return r.p }

// RawTexture returns a RawTexture JavaScript class.
func (ba *Babylon) RawTexture() *RawTexture {
	p := ba.ctx.Get("RawTexture")
	return RawTextureFromJSObject(p, ba.ctx)
}

// RawTextureFromJSObject returns a wrapped RawTexture JavaScript class.
func RawTextureFromJSObject(p js.Value, ctx js.Value) *RawTexture {
	return &RawTexture{Texture: TextureFromJSObject(p, ctx), ctx: ctx}
}

// RawTextureArrayToJSArray returns a JavaScript Array for the wrapped array.
func RawTextureArrayToJSArray(array []*RawTexture) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewRawTextureOpts contains optional parameters for NewRawTexture.
type NewRawTextureOpts struct {
	GenerateMipMaps *bool
	InvertY         *bool
	SamplingMode    *float64
	Type            *float64
}

// NewRawTexture returns a new RawTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.rawtexture#constructor
func (ba *Babylon) NewRawTexture(data js.Value, width float64, height float64, format float64, scene *Scene, opts *NewRawTextureOpts) *RawTexture {
	if opts == nil {
		opts = &NewRawTextureOpts{}
	}

	args := make([]interface{}, 0, 5+4)

	args = append(args, data)
	args = append(args, width)
	args = append(args, height)
	args = append(args, format)
	args = append(args, scene.JSObject())

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
	if opts.Type == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Type)
	}

	p := ba.ctx.Get("RawTexture").New(args...)
	return RawTextureFromJSObject(p, ba.ctx)
}

// RawTextureCreateAlphaTextureOpts contains optional parameters for RawTexture.CreateAlphaTexture.
type RawTextureCreateAlphaTextureOpts struct {
	GenerateMipMaps *bool
	InvertY         *bool
	SamplingMode    *float64
}

// CreateAlphaTexture calls the CreateAlphaTexture method on the RawTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.rawtexture#createalphatexture
func (r *RawTexture) CreateAlphaTexture(data js.Value, width float64, height float64, scene *Scene, opts *RawTextureCreateAlphaTextureOpts) *RawTexture {
	if opts == nil {
		opts = &RawTextureCreateAlphaTextureOpts{}
	}

	args := make([]interface{}, 0, 4+3)

	args = append(args, data)

	args = append(args, width)

	args = append(args, height)

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
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

	retVal := r.p.Call("CreateAlphaTexture", args...)
	return RawTextureFromJSObject(retVal, r.ctx)
}

// RawTextureCreateLuminanceAlphaTextureOpts contains optional parameters for RawTexture.CreateLuminanceAlphaTexture.
type RawTextureCreateLuminanceAlphaTextureOpts struct {
	GenerateMipMaps *bool
	InvertY         *bool
	SamplingMode    *float64
}

// CreateLuminanceAlphaTexture calls the CreateLuminanceAlphaTexture method on the RawTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.rawtexture#createluminancealphatexture
func (r *RawTexture) CreateLuminanceAlphaTexture(data js.Value, width float64, height float64, scene *Scene, opts *RawTextureCreateLuminanceAlphaTextureOpts) *RawTexture {
	if opts == nil {
		opts = &RawTextureCreateLuminanceAlphaTextureOpts{}
	}

	args := make([]interface{}, 0, 4+3)

	args = append(args, data)

	args = append(args, width)

	args = append(args, height)

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
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

	retVal := r.p.Call("CreateLuminanceAlphaTexture", args...)
	return RawTextureFromJSObject(retVal, r.ctx)
}

// RawTextureCreateLuminanceTextureOpts contains optional parameters for RawTexture.CreateLuminanceTexture.
type RawTextureCreateLuminanceTextureOpts struct {
	GenerateMipMaps *bool
	InvertY         *bool
	SamplingMode    *float64
}

// CreateLuminanceTexture calls the CreateLuminanceTexture method on the RawTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.rawtexture#createluminancetexture
func (r *RawTexture) CreateLuminanceTexture(data js.Value, width float64, height float64, scene *Scene, opts *RawTextureCreateLuminanceTextureOpts) *RawTexture {
	if opts == nil {
		opts = &RawTextureCreateLuminanceTextureOpts{}
	}

	args := make([]interface{}, 0, 4+3)

	args = append(args, data)

	args = append(args, width)

	args = append(args, height)

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
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

	retVal := r.p.Call("CreateLuminanceTexture", args...)
	return RawTextureFromJSObject(retVal, r.ctx)
}

// RawTextureCreateRGBATextureOpts contains optional parameters for RawTexture.CreateRGBATexture.
type RawTextureCreateRGBATextureOpts struct {
	GenerateMipMaps *bool
	InvertY         *bool
	SamplingMode    *float64
	Type            *float64
}

// CreateRGBATexture calls the CreateRGBATexture method on the RawTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.rawtexture#creatergbatexture
func (r *RawTexture) CreateRGBATexture(data js.Value, width float64, height float64, scene *Scene, opts *RawTextureCreateRGBATextureOpts) *RawTexture {
	if opts == nil {
		opts = &RawTextureCreateRGBATextureOpts{}
	}

	args := make([]interface{}, 0, 4+4)

	args = append(args, data)

	args = append(args, width)

	args = append(args, height)

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
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
	if opts.Type == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Type)
	}

	retVal := r.p.Call("CreateRGBATexture", args...)
	return RawTextureFromJSObject(retVal, r.ctx)
}

// RawTextureCreateRGBTextureOpts contains optional parameters for RawTexture.CreateRGBTexture.
type RawTextureCreateRGBTextureOpts struct {
	GenerateMipMaps *bool
	InvertY         *bool
	SamplingMode    *float64
	Type            *float64
}

// CreateRGBTexture calls the CreateRGBTexture method on the RawTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.rawtexture#creatergbtexture
func (r *RawTexture) CreateRGBTexture(data js.Value, width float64, height float64, scene *Scene, opts *RawTextureCreateRGBTextureOpts) *RawTexture {
	if opts == nil {
		opts = &RawTextureCreateRGBTextureOpts{}
	}

	args := make([]interface{}, 0, 4+4)

	args = append(args, data)

	args = append(args, width)

	args = append(args, height)

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
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
	if opts.Type == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Type)
	}

	retVal := r.p.Call("CreateRGBTexture", args...)
	return RawTextureFromJSObject(retVal, r.ctx)
}

// RawTextureCreateRTextureOpts contains optional parameters for RawTexture.CreateRTexture.
type RawTextureCreateRTextureOpts struct {
	GenerateMipMaps *bool
	InvertY         *bool
	SamplingMode    *float64
	Type            *float64
}

// CreateRTexture calls the CreateRTexture method on the RawTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.rawtexture#creatertexture
func (r *RawTexture) CreateRTexture(data js.Value, width float64, height float64, scene *Scene, opts *RawTextureCreateRTextureOpts) *RawTexture {
	if opts == nil {
		opts = &RawTextureCreateRTextureOpts{}
	}

	args := make([]interface{}, 0, 4+4)

	args = append(args, data)

	args = append(args, width)

	args = append(args, height)

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
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
	if opts.Type == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Type)
	}

	retVal := r.p.Call("CreateRTexture", args...)
	return RawTextureFromJSObject(retVal, r.ctx)
}

// Update calls the Update method on the RawTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.rawtexture#update
func (r *RawTexture) Update(data js.Value) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, data)

	r.p.Call("update", args...)
}

// Format returns the Format property of class RawTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.rawtexture#format
func (r *RawTexture) Format() float64 {
	retVal := r.p.Get("format")
	return retVal.Float()
}

// SetFormat sets the Format property of class RawTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.rawtexture#format
func (r *RawTexture) SetFormat(format float64) *RawTexture {
	r.p.Set("format", format)
	return r
}
