// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GrassProceduralTexture represents a babylon.js GrassProceduralTexture.
//
type GrassProceduralTexture struct {
	*ProceduralTexture
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (g *GrassProceduralTexture) JSObject() js.Value { return g.p }

// GrassProceduralTexture returns a GrassProceduralTexture JavaScript class.
func (ba *Babylon) GrassProceduralTexture() *GrassProceduralTexture {
	p := ba.ctx.Get("GrassProceduralTexture")
	return GrassProceduralTextureFromJSObject(p, ba.ctx)
}

// GrassProceduralTextureFromJSObject returns a wrapped GrassProceduralTexture JavaScript class.
func GrassProceduralTextureFromJSObject(p js.Value, ctx js.Value) *GrassProceduralTexture {
	return &GrassProceduralTexture{ProceduralTexture: ProceduralTextureFromJSObject(p, ctx), ctx: ctx}
}

// GrassProceduralTextureArrayToJSArray returns a JavaScript Array for the wrapped array.
func GrassProceduralTextureArrayToJSArray(array []*GrassProceduralTexture) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewGrassProceduralTextureOpts contains optional parameters for NewGrassProceduralTexture.
type NewGrassProceduralTextureOpts struct {
	FallbackTexture *Texture
	GenerateMipMaps *bool
}

// NewGrassProceduralTexture returns a new GrassProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.grassproceduraltexture#constructor
func (ba *Babylon) NewGrassProceduralTexture(name string, size float64, scene *Scene, opts *NewGrassProceduralTextureOpts) *GrassProceduralTexture {
	if opts == nil {
		opts = &NewGrassProceduralTextureOpts{}
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

	p := ba.ctx.Get("GrassProceduralTexture").New(args...)
	return GrassProceduralTextureFromJSObject(p, ba.ctx)
}

// Parse calls the Parse method on the GrassProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.grassproceduraltexture#parse
func (g *GrassProceduralTexture) Parse(parsedTexture JSObject, scene *Scene, rootUrl string) *GrassProceduralTexture {

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

	retVal := g.p.Call("Parse", args...)
	return GrassProceduralTextureFromJSObject(retVal, g.ctx)
}

// Serialize calls the Serialize method on the GrassProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.grassproceduraltexture#serialize
func (g *GrassProceduralTexture) Serialize() js.Value {

	retVal := g.p.Call("serialize")
	return retVal
}

// UpdateShaderUniforms calls the UpdateShaderUniforms method on the GrassProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.grassproceduraltexture#updateshaderuniforms
func (g *GrassProceduralTexture) UpdateShaderUniforms() {

	g.p.Call("updateShaderUniforms")
}

// GrassColors returns the GrassColors property of class GrassProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.grassproceduraltexture#grasscolors
func (g *GrassProceduralTexture) GrassColors() []*Color3 {
	retVal := g.p.Get("grassColors")
	result := []*Color3{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, Color3FromJSObject(retVal.Index(ri), g.ctx))
	}
	return result
}

// SetGrassColors sets the GrassColors property of class GrassProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.grassproceduraltexture#grasscolors
func (g *GrassProceduralTexture) SetGrassColors(grassColors []*Color3) *GrassProceduralTexture {
	g.p.Set("grassColors", grassColors)
	return g
}

// GroundColor returns the GroundColor property of class GrassProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.grassproceduraltexture#groundcolor
func (g *GrassProceduralTexture) GroundColor() *Color3 {
	retVal := g.p.Get("groundColor")
	return Color3FromJSObject(retVal, g.ctx)
}

// SetGroundColor sets the GroundColor property of class GrassProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.grassproceduraltexture#groundcolor
func (g *GrassProceduralTexture) SetGroundColor(groundColor *Color3) *GrassProceduralTexture {
	g.p.Set("groundColor", groundColor.JSObject())
	return g
}
