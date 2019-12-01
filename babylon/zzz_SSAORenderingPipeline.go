// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SSAORenderingPipeline represents a babylon.js SSAORenderingPipeline.
// Render pipeline to produce ssao effect
type SSAORenderingPipeline struct {
	*PostProcessRenderPipeline
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *SSAORenderingPipeline) JSObject() js.Value { return s.p }

// SSAORenderingPipeline returns a SSAORenderingPipeline JavaScript class.
func (ba *Babylon) SSAORenderingPipeline() *SSAORenderingPipeline {
	p := ba.ctx.Get("SSAORenderingPipeline")
	return SSAORenderingPipelineFromJSObject(p, ba.ctx)
}

// SSAORenderingPipelineFromJSObject returns a wrapped SSAORenderingPipeline JavaScript class.
func SSAORenderingPipelineFromJSObject(p js.Value, ctx js.Value) *SSAORenderingPipeline {
	return &SSAORenderingPipeline{PostProcessRenderPipeline: PostProcessRenderPipelineFromJSObject(p, ctx), ctx: ctx}
}

// SSAORenderingPipelineArrayToJSArray returns a JavaScript Array for the wrapped array.
func SSAORenderingPipelineArrayToJSArray(array []*SSAORenderingPipeline) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewSSAORenderingPipelineOpts contains optional parameters for NewSSAORenderingPipeline.
type NewSSAORenderingPipelineOpts struct {
	Cameras *Camera
}

// NewSSAORenderingPipeline returns a new SSAORenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline
func (ba *Babylon) NewSSAORenderingPipeline(name string, scene *Scene, ratio interface{}, opts *NewSSAORenderingPipelineOpts) *SSAORenderingPipeline {
	if opts == nil {
		opts = &NewSSAORenderingPipelineOpts{}
	}

	args := make([]interface{}, 0, 3+1)

	args = append(args, name)
	args = append(args, scene.JSObject())
	args = append(args, ratio)

	if opts.Cameras == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Cameras.JSObject())
	}

	p := ba.ctx.Get("SSAORenderingPipeline").New(args...)
	return SSAORenderingPipelineFromJSObject(p, ba.ctx)
}

// AddEffect calls the AddEffect method on the SSAORenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#addeffect
func (s *SSAORenderingPipeline) AddEffect(renderEffect *PostProcessRenderEffect) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, renderEffect.JSObject())

	s.p.Call("addEffect", args...)
}

// SSAORenderingPipelineDisposeOpts contains optional parameters for SSAORenderingPipeline.Dispose.
type SSAORenderingPipelineDisposeOpts struct {
	DisableDepthRender *bool
}

// Dispose calls the Dispose method on the SSAORenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#dispose
func (s *SSAORenderingPipeline) Dispose(opts *SSAORenderingPipelineDisposeOpts) {
	if opts == nil {
		opts = &SSAORenderingPipelineDisposeOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.DisableDepthRender == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.DisableDepthRender)
	}

	s.p.Call("dispose", args...)
}

// GetClassName calls the GetClassName method on the SSAORenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#getclassname
func (s *SSAORenderingPipeline) GetClassName() string {

	retVal := s.p.Call("getClassName")
	return retVal.String()
}

/*

// Area returns the Area property of class SSAORenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#area
func (s *SSAORenderingPipeline) Area(area float64) *SSAORenderingPipeline {
	p := ba.ctx.Get("SSAORenderingPipeline").New(area)
	return SSAORenderingPipelineFromJSObject(p, ba.ctx)
}

// SetArea sets the Area property of class SSAORenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#area
func (s *SSAORenderingPipeline) SetArea(area float64) *SSAORenderingPipeline {
	p := ba.ctx.Get("SSAORenderingPipeline").New(area)
	return SSAORenderingPipelineFromJSObject(p, ba.ctx)
}

// Base returns the Base property of class SSAORenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#base
func (s *SSAORenderingPipeline) Base(base float64) *SSAORenderingPipeline {
	p := ba.ctx.Get("SSAORenderingPipeline").New(base)
	return SSAORenderingPipelineFromJSObject(p, ba.ctx)
}

// SetBase sets the Base property of class SSAORenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#base
func (s *SSAORenderingPipeline) SetBase(base float64) *SSAORenderingPipeline {
	p := ba.ctx.Get("SSAORenderingPipeline").New(base)
	return SSAORenderingPipelineFromJSObject(p, ba.ctx)
}

// Cameras returns the Cameras property of class SSAORenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#cameras
func (s *SSAORenderingPipeline) Cameras(cameras *Camera) *SSAORenderingPipeline {
	p := ba.ctx.Get("SSAORenderingPipeline").New(cameras.JSObject())
	return SSAORenderingPipelineFromJSObject(p, ba.ctx)
}

// SetCameras sets the Cameras property of class SSAORenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#cameras
func (s *SSAORenderingPipeline) SetCameras(cameras *Camera) *SSAORenderingPipeline {
	p := ba.ctx.Get("SSAORenderingPipeline").New(cameras.JSObject())
	return SSAORenderingPipelineFromJSObject(p, ba.ctx)
}

// FallOff returns the FallOff property of class SSAORenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#falloff
func (s *SSAORenderingPipeline) FallOff(fallOff float64) *SSAORenderingPipeline {
	p := ba.ctx.Get("SSAORenderingPipeline").New(fallOff)
	return SSAORenderingPipelineFromJSObject(p, ba.ctx)
}

// SetFallOff sets the FallOff property of class SSAORenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#falloff
func (s *SSAORenderingPipeline) SetFallOff(fallOff float64) *SSAORenderingPipeline {
	p := ba.ctx.Get("SSAORenderingPipeline").New(fallOff)
	return SSAORenderingPipelineFromJSObject(p, ba.ctx)
}

// InspectableCustomProperties returns the InspectableCustomProperties property of class SSAORenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#inspectablecustomproperties
func (s *SSAORenderingPipeline) InspectableCustomProperties(inspectableCustomProperties *IInspectable) *SSAORenderingPipeline {
	p := ba.ctx.Get("SSAORenderingPipeline").New(inspectableCustomProperties.JSObject())
	return SSAORenderingPipelineFromJSObject(p, ba.ctx)
}

// SetInspectableCustomProperties sets the InspectableCustomProperties property of class SSAORenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#inspectablecustomproperties
func (s *SSAORenderingPipeline) SetInspectableCustomProperties(inspectableCustomProperties *IInspectable) *SSAORenderingPipeline {
	p := ba.ctx.Get("SSAORenderingPipeline").New(inspectableCustomProperties.JSObject())
	return SSAORenderingPipelineFromJSObject(p, ba.ctx)
}

// IsSupported returns the IsSupported property of class SSAORenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#issupported
func (s *SSAORenderingPipeline) IsSupported(isSupported bool) *SSAORenderingPipeline {
	p := ba.ctx.Get("SSAORenderingPipeline").New(isSupported)
	return SSAORenderingPipelineFromJSObject(p, ba.ctx)
}

// SetIsSupported sets the IsSupported property of class SSAORenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#issupported
func (s *SSAORenderingPipeline) SetIsSupported(isSupported bool) *SSAORenderingPipeline {
	p := ba.ctx.Get("SSAORenderingPipeline").New(isSupported)
	return SSAORenderingPipelineFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class SSAORenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#name
func (s *SSAORenderingPipeline) Name(name string) *SSAORenderingPipeline {
	p := ba.ctx.Get("SSAORenderingPipeline").New(name)
	return SSAORenderingPipelineFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class SSAORenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#name
func (s *SSAORenderingPipeline) SetName(name string) *SSAORenderingPipeline {
	p := ba.ctx.Get("SSAORenderingPipeline").New(name)
	return SSAORenderingPipelineFromJSObject(p, ba.ctx)
}

// Radius returns the Radius property of class SSAORenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#radius
func (s *SSAORenderingPipeline) Radius(radius float64) *SSAORenderingPipeline {
	p := ba.ctx.Get("SSAORenderingPipeline").New(radius)
	return SSAORenderingPipelineFromJSObject(p, ba.ctx)
}

// SetRadius sets the Radius property of class SSAORenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#radius
func (s *SSAORenderingPipeline) SetRadius(radius float64) *SSAORenderingPipeline {
	p := ba.ctx.Get("SSAORenderingPipeline").New(radius)
	return SSAORenderingPipelineFromJSObject(p, ba.ctx)
}

// Scene returns the Scene property of class SSAORenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#scene
func (s *SSAORenderingPipeline) Scene(scene *Scene) *SSAORenderingPipeline {
	p := ba.ctx.Get("SSAORenderingPipeline").New(scene.JSObject())
	return SSAORenderingPipelineFromJSObject(p, ba.ctx)
}

// SetScene sets the Scene property of class SSAORenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#scene
func (s *SSAORenderingPipeline) SetScene(scene *Scene) *SSAORenderingPipeline {
	p := ba.ctx.Get("SSAORenderingPipeline").New(scene.JSObject())
	return SSAORenderingPipelineFromJSObject(p, ba.ctx)
}

// TotalStrength returns the TotalStrength property of class SSAORenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#totalstrength
func (s *SSAORenderingPipeline) TotalStrength(totalStrength float64) *SSAORenderingPipeline {
	p := ba.ctx.Get("SSAORenderingPipeline").New(totalStrength)
	return SSAORenderingPipelineFromJSObject(p, ba.ctx)
}

// SetTotalStrength sets the TotalStrength property of class SSAORenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#totalstrength
func (s *SSAORenderingPipeline) SetTotalStrength(totalStrength float64) *SSAORenderingPipeline {
	p := ba.ctx.Get("SSAORenderingPipeline").New(totalStrength)
	return SSAORenderingPipelineFromJSObject(p, ba.ctx)
}

*/
