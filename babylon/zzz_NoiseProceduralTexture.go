// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// NoiseProceduralTexture represents a babylon.js NoiseProceduralTexture.
// Class used to generate noise procedural textures
type NoiseProceduralTexture struct {
	*ProceduralTexture
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (n *NoiseProceduralTexture) JSObject() js.Value { return n.p }

// NoiseProceduralTexture returns a NoiseProceduralTexture JavaScript class.
func (ba *Babylon) NoiseProceduralTexture() *NoiseProceduralTexture {
	p := ba.ctx.Get("NoiseProceduralTexture")
	return NoiseProceduralTextureFromJSObject(p, ba.ctx)
}

// NoiseProceduralTextureFromJSObject returns a wrapped NoiseProceduralTexture JavaScript class.
func NoiseProceduralTextureFromJSObject(p js.Value, ctx js.Value) *NoiseProceduralTexture {
	return &NoiseProceduralTexture{ProceduralTexture: ProceduralTextureFromJSObject(p, ctx), ctx: ctx}
}

// NoiseProceduralTextureArrayToJSArray returns a JavaScript Array for the wrapped array.
func NoiseProceduralTextureArrayToJSArray(array []*NoiseProceduralTexture) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewNoiseProceduralTextureOpts contains optional parameters for NewNoiseProceduralTexture.
type NewNoiseProceduralTextureOpts struct {
	Size            *float64
	Scene           *Scene
	FallbackTexture *Texture
	GenerateMipMaps *bool
}

// NewNoiseProceduralTexture returns a new NoiseProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.noiseproceduraltexture#constructor
func (ba *Babylon) NewNoiseProceduralTexture(name string, opts *NewNoiseProceduralTextureOpts) *NoiseProceduralTexture {
	if opts == nil {
		opts = &NewNoiseProceduralTextureOpts{}
	}

	args := make([]interface{}, 0, 1+4)

	args = append(args, name)

	if opts.Size == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Size)
	}
	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}
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

	p := ba.ctx.Get("NoiseProceduralTexture").New(args...)
	return NoiseProceduralTextureFromJSObject(p, ba.ctx)
}

// Parse calls the Parse method on the NoiseProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.noiseproceduraltexture#parse
func (n *NoiseProceduralTexture) Parse(parsedTexture JSObject, scene *Scene) *NoiseProceduralTexture {

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

	retVal := n.p.Call("Parse", args...)
	return NoiseProceduralTextureFromJSObject(retVal, n.ctx)
}

// NoiseProceduralTextureRenderOpts contains optional parameters for NoiseProceduralTexture.Render.
type NoiseProceduralTextureRenderOpts struct {
	UseCameraPostProcess *bool
}

// Render calls the Render method on the NoiseProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.noiseproceduraltexture#render
func (n *NoiseProceduralTexture) Render(opts *NoiseProceduralTextureRenderOpts) {
	if opts == nil {
		opts = &NoiseProceduralTextureRenderOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.UseCameraPostProcess == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UseCameraPostProcess)
	}

	n.p.Call("render", args...)
}

// Serialize calls the Serialize method on the NoiseProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.noiseproceduraltexture#serialize
func (n *NoiseProceduralTexture) Serialize() js.Value {

	retVal := n.p.Call("serialize")
	return retVal
}

// AnimationSpeedFactor returns the AnimationSpeedFactor property of class NoiseProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.noiseproceduraltexture#animationspeedfactor
func (n *NoiseProceduralTexture) AnimationSpeedFactor() float64 {
	retVal := n.p.Get("animationSpeedFactor")
	return retVal.Float()
}

// SetAnimationSpeedFactor sets the AnimationSpeedFactor property of class NoiseProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.noiseproceduraltexture#animationspeedfactor
func (n *NoiseProceduralTexture) SetAnimationSpeedFactor(animationSpeedFactor float64) *NoiseProceduralTexture {
	n.p.Set("animationSpeedFactor", animationSpeedFactor)
	return n
}

// Brightness returns the Brightness property of class NoiseProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.noiseproceduraltexture#brightness
func (n *NoiseProceduralTexture) Brightness() float64 {
	retVal := n.p.Get("brightness")
	return retVal.Float()
}

// SetBrightness sets the Brightness property of class NoiseProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.noiseproceduraltexture#brightness
func (n *NoiseProceduralTexture) SetBrightness(brightness float64) *NoiseProceduralTexture {
	n.p.Set("brightness", brightness)
	return n
}

// Octaves returns the Octaves property of class NoiseProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.noiseproceduraltexture#octaves
func (n *NoiseProceduralTexture) Octaves() float64 {
	retVal := n.p.Get("octaves")
	return retVal.Float()
}

// SetOctaves sets the Octaves property of class NoiseProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.noiseproceduraltexture#octaves
func (n *NoiseProceduralTexture) SetOctaves(octaves float64) *NoiseProceduralTexture {
	n.p.Set("octaves", octaves)
	return n
}

// Persistence returns the Persistence property of class NoiseProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.noiseproceduraltexture#persistence
func (n *NoiseProceduralTexture) Persistence() float64 {
	retVal := n.p.Get("persistence")
	return retVal.Float()
}

// SetPersistence sets the Persistence property of class NoiseProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.noiseproceduraltexture#persistence
func (n *NoiseProceduralTexture) SetPersistence(persistence float64) *NoiseProceduralTexture {
	n.p.Set("persistence", persistence)
	return n
}
