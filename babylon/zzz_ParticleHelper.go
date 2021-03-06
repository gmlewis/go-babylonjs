// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ParticleHelper represents a babylon.js ParticleHelper.
// This class is made for on one-liner static method to help creating particle system set.
type ParticleHelper struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *ParticleHelper) JSObject() js.Value { return p.p }

// ParticleHelper returns a ParticleHelper JavaScript class.
func (ba *Babylon) ParticleHelper() *ParticleHelper {
	p := ba.ctx.Get("ParticleHelper")
	return ParticleHelperFromJSObject(p, ba.ctx)
}

// ParticleHelperFromJSObject returns a wrapped ParticleHelper JavaScript class.
func ParticleHelperFromJSObject(p js.Value, ctx js.Value) *ParticleHelper {
	return &ParticleHelper{p: p, ctx: ctx}
}

// ParticleHelperArrayToJSArray returns a JavaScript Array for the wrapped array.
func ParticleHelperArrayToJSArray(array []*ParticleHelper) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// ParticleHelperCreateAsyncOpts contains optional parameters for ParticleHelper.CreateAsync.
type ParticleHelperCreateAsyncOpts struct {
	Gpu *bool
}

// CreateAsync calls the CreateAsync method on the ParticleHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlehelper#createasync
func (p *ParticleHelper) CreateAsync(jsType string, scene *Scene, opts *ParticleHelperCreateAsyncOpts) *Promise {
	if opts == nil {
		opts = &ParticleHelperCreateAsyncOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, jsType)

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
	}

	if opts.Gpu == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Gpu)
	}

	retVal := p.p.Call("CreateAsync", args...)
	return PromiseFromJSObject(retVal, p.ctx)
}

// ParticleHelperCreateDefaultOpts contains optional parameters for ParticleHelper.CreateDefault.
type ParticleHelperCreateDefaultOpts struct {
	Capacity *float64
	Scene    *Scene
	UseGPU   *bool
}

// CreateDefault calls the CreateDefault method on the ParticleHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlehelper#createdefault
func (p *ParticleHelper) CreateDefault(emitter *AbstractMesh, opts *ParticleHelperCreateDefaultOpts) *IParticleSystem {
	if opts == nil {
		opts = &ParticleHelperCreateDefaultOpts{}
	}

	args := make([]interface{}, 0, 1+3)

	if emitter == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, emitter.JSObject())
	}

	if opts.Capacity == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Capacity)
	}
	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}
	if opts.UseGPU == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UseGPU)
	}

	retVal := p.p.Call("CreateDefault", args...)
	return IParticleSystemFromJSObject(retVal, p.ctx)
}

// ExportSet calls the ExportSet method on the ParticleHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlehelper#exportset
func (p *ParticleHelper) ExportSet(systems []*IParticleSystem) *ParticleSystemSet {

	args := make([]interface{}, 0, 1+0)

	args = append(args, IParticleSystemArrayToJSArray(systems))

	retVal := p.p.Call("ExportSet", args...)
	return ParticleSystemSetFromJSObject(retVal, p.ctx)
}

// BaseAssetsUrl returns the BaseAssetsUrl property of class ParticleHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.particlehelper#baseassetsurl
func (p *ParticleHelper) BaseAssetsUrl() string {
	retVal := p.p.Get("BaseAssetsUrl")
	return retVal.String()
}

// SetBaseAssetsUrl sets the BaseAssetsUrl property of class ParticleHelper.
//
// https://doc.babylonjs.com/api/classes/babylon.particlehelper#baseassetsurl
func (p *ParticleHelper) SetBaseAssetsUrl(BaseAssetsUrl string) *ParticleHelper {
	p.p.Set("BaseAssetsUrl", BaseAssetsUrl)
	return p
}
