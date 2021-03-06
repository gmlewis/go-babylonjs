// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// NormalMapProceduralTexture represents a babylon.js NormalMapProceduralTexture.
//
type NormalMapProceduralTexture struct {
	*ProceduralTexture
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (n *NormalMapProceduralTexture) JSObject() js.Value { return n.p }

// NormalMapProceduralTexture returns a NormalMapProceduralTexture JavaScript class.
func (ba *Babylon) NormalMapProceduralTexture() *NormalMapProceduralTexture {
	p := ba.ctx.Get("NormalMapProceduralTexture")
	return NormalMapProceduralTextureFromJSObject(p, ba.ctx)
}

// NormalMapProceduralTextureFromJSObject returns a wrapped NormalMapProceduralTexture JavaScript class.
func NormalMapProceduralTextureFromJSObject(p js.Value, ctx js.Value) *NormalMapProceduralTexture {
	return &NormalMapProceduralTexture{ProceduralTexture: ProceduralTextureFromJSObject(p, ctx), ctx: ctx}
}

// NormalMapProceduralTextureArrayToJSArray returns a JavaScript Array for the wrapped array.
func NormalMapProceduralTextureArrayToJSArray(array []*NormalMapProceduralTexture) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewNormalMapProceduralTextureOpts contains optional parameters for NewNormalMapProceduralTexture.
type NewNormalMapProceduralTextureOpts struct {
	FallbackTexture *Texture
	GenerateMipMaps *bool
}

// NewNormalMapProceduralTexture returns a new NormalMapProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalmapproceduraltexture#constructor
func (ba *Babylon) NewNormalMapProceduralTexture(name string, size float64, scene *Scene, opts *NewNormalMapProceduralTextureOpts) *NormalMapProceduralTexture {
	if opts == nil {
		opts = &NewNormalMapProceduralTextureOpts{}
	}

	args := make([]interface{}, 0, 3+2)

	args = append(args, name)
	args = append(args, size)
	args = append(args, scene.JSObject())

	if opts.FallbackTexture == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.FallbackTexture.JSObject())
	}
	if opts.GenerateMipMaps == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.GenerateMipMaps)
	}

	p := ba.ctx.Get("NormalMapProceduralTexture").New(args...)
	return NormalMapProceduralTextureFromJSObject(p, ba.ctx)
}

// Parse calls the Parse method on the NormalMapProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalmapproceduraltexture#parse
func (n *NormalMapProceduralTexture) Parse(parsedTexture JSObject, scene *Scene, rootUrl string) *NormalMapProceduralTexture {

	args := make([]interface{}, 0, 3+0)

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

	args = append(args, rootUrl)

	retVal := n.p.Call("Parse", args...)
	return NormalMapProceduralTextureFromJSObject(retVal, n.ctx)
}

// NormalMapProceduralTextureRenderOpts contains optional parameters for NormalMapProceduralTexture.Render.
type NormalMapProceduralTextureRenderOpts struct {
	UseCameraPostProcess *bool
}

// Render calls the Render method on the NormalMapProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalmapproceduraltexture#render
func (n *NormalMapProceduralTexture) Render(opts *NormalMapProceduralTextureRenderOpts) {
	if opts == nil {
		opts = &NormalMapProceduralTextureRenderOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.UseCameraPostProcess == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UseCameraPostProcess)
	}

	n.p.Call("render", args...)
}

// Resize calls the Resize method on the NormalMapProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalmapproceduraltexture#resize
func (n *NormalMapProceduralTexture) Resize(size JSObject, generateMipMaps JSObject) {

	args := make([]interface{}, 0, 2+0)

	if size == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, size.JSObject())
	}

	if generateMipMaps == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, generateMipMaps.JSObject())
	}

	n.p.Call("resize", args...)
}

// Serialize calls the Serialize method on the NormalMapProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalmapproceduraltexture#serialize
func (n *NormalMapProceduralTexture) Serialize() js.Value {

	retVal := n.p.Call("serialize")
	return retVal
}

// UpdateShaderUniforms calls the UpdateShaderUniforms method on the NormalMapProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalmapproceduraltexture#updateshaderuniforms
func (n *NormalMapProceduralTexture) UpdateShaderUniforms() {

	n.p.Call("updateShaderUniforms")
}

// BaseTexture returns the BaseTexture property of class NormalMapProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.normalmapproceduraltexture#basetexture
func (n *NormalMapProceduralTexture) BaseTexture() *Texture {
	retVal := n.p.Get("baseTexture")
	return TextureFromJSObject(retVal, n.ctx)
}

// SetBaseTexture sets the BaseTexture property of class NormalMapProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.normalmapproceduraltexture#basetexture
func (n *NormalMapProceduralTexture) SetBaseTexture(baseTexture *Texture) *NormalMapProceduralTexture {
	n.p.Set("baseTexture", baseTexture.JSObject())
	return n
}
