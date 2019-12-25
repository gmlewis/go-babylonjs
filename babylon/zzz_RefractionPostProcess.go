// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// RefractionPostProcess represents a babylon.js RefractionPostProcess.
// Post process which applies a refractin texture
//
// See: https://doc.babylonjs.com/how_to/how_to_use_postprocesses#refraction
type RefractionPostProcess struct {
	*PostProcess
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (r *RefractionPostProcess) JSObject() js.Value { return r.p }

// RefractionPostProcess returns a RefractionPostProcess JavaScript class.
func (ba *Babylon) RefractionPostProcess() *RefractionPostProcess {
	p := ba.ctx.Get("RefractionPostProcess")
	return RefractionPostProcessFromJSObject(p, ba.ctx)
}

// RefractionPostProcessFromJSObject returns a wrapped RefractionPostProcess JavaScript class.
func RefractionPostProcessFromJSObject(p js.Value, ctx js.Value) *RefractionPostProcess {
	return &RefractionPostProcess{PostProcess: PostProcessFromJSObject(p, ctx), ctx: ctx}
}

// RefractionPostProcessArrayToJSArray returns a JavaScript Array for the wrapped array.
func RefractionPostProcessArrayToJSArray(array []*RefractionPostProcess) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewRefractionPostProcessOpts contains optional parameters for NewRefractionPostProcess.
type NewRefractionPostProcessOpts struct {
	SamplingMode *float64
	Engine       *Engine
	Reusable     *bool
}

// NewRefractionPostProcess returns a new RefractionPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.refractionpostprocess#constructor
func (ba *Babylon) NewRefractionPostProcess(name string, refractionTextureUrl string, color *Color3, depth float64, colorLevel float64, options float64, camera *Camera, opts *NewRefractionPostProcessOpts) *RefractionPostProcess {
	if opts == nil {
		opts = &NewRefractionPostProcessOpts{}
	}

	args := make([]interface{}, 0, 7+3)

	args = append(args, name)
	args = append(args, refractionTextureUrl)
	args = append(args, color.JSObject())
	args = append(args, depth)
	args = append(args, colorLevel)
	args = append(args, options)
	args = append(args, camera.JSObject())

	if opts.SamplingMode == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.SamplingMode)
	}
	if opts.Engine == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Engine.JSObject())
	}
	if opts.Reusable == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Reusable)
	}

	p := ba.ctx.Get("RefractionPostProcess").New(args...)
	return RefractionPostProcessFromJSObject(p, ba.ctx)
}

// Dispose calls the Dispose method on the RefractionPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.refractionpostprocess#dispose
func (r *RefractionPostProcess) Dispose(camera *Camera) {

	args := make([]interface{}, 0, 1+0)

	if camera == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, camera.JSObject())
	}

	r.p.Call("dispose", args...)
}

// Color returns the Color property of class RefractionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.refractionpostprocess#color
func (r *RefractionPostProcess) Color() *Color3 {
	retVal := r.p.Get("color")
	return Color3FromJSObject(retVal, r.ctx)
}

// SetColor sets the Color property of class RefractionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.refractionpostprocess#color
func (r *RefractionPostProcess) SetColor(color *Color3) *RefractionPostProcess {
	r.p.Set("color", color.JSObject())
	return r
}

// ColorLevel returns the ColorLevel property of class RefractionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.refractionpostprocess#colorlevel
func (r *RefractionPostProcess) ColorLevel() float64 {
	retVal := r.p.Get("colorLevel")
	return retVal.Float()
}

// SetColorLevel sets the ColorLevel property of class RefractionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.refractionpostprocess#colorlevel
func (r *RefractionPostProcess) SetColorLevel(colorLevel float64) *RefractionPostProcess {
	r.p.Set("colorLevel", colorLevel)
	return r
}

// Depth returns the Depth property of class RefractionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.refractionpostprocess#depth
func (r *RefractionPostProcess) Depth() float64 {
	retVal := r.p.Get("depth")
	return retVal.Float()
}

// SetDepth sets the Depth property of class RefractionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.refractionpostprocess#depth
func (r *RefractionPostProcess) SetDepth(depth float64) *RefractionPostProcess {
	r.p.Set("depth", depth)
	return r
}

// RefractionTexture returns the RefractionTexture property of class RefractionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.refractionpostprocess#refractiontexture
func (r *RefractionPostProcess) RefractionTexture() *Texture {
	retVal := r.p.Get("refractionTexture")
	return TextureFromJSObject(retVal, r.ctx)
}

// SetRefractionTexture sets the RefractionTexture property of class RefractionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.refractionpostprocess#refractiontexture
func (r *RefractionPostProcess) SetRefractionTexture(refractionTexture *Texture) *RefractionPostProcess {
	r.p.Set("refractionTexture", refractionTexture.JSObject())
	return r
}
