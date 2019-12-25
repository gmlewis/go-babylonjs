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
	Cameras []*Camera
}

// NewSSAORenderingPipeline returns a new SSAORenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#constructor
func (ba *Babylon) NewSSAORenderingPipeline(name string, scene *Scene, ratio JSObject, opts *NewSSAORenderingPipelineOpts) *SSAORenderingPipeline {
	if opts == nil {
		opts = &NewSSAORenderingPipelineOpts{}
	}

	args := make([]interface{}, 0, 3+1)

	args = append(args, name)
	args = append(args, scene.JSObject())
	args = append(args, ratio.JSObject())

	if opts.Cameras == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, CameraArrayToJSArray(opts.Cameras))
	}

	p := ba.ctx.Get("SSAORenderingPipeline").New(args...)
	return SSAORenderingPipelineFromJSObject(p, ba.ctx)
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

// Area returns the Area property of class SSAORenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#area
func (s *SSAORenderingPipeline) Area() float64 {
	retVal := s.p.Get("area")
	return retVal.Float()
}

// SetArea sets the Area property of class SSAORenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#area
func (s *SSAORenderingPipeline) SetArea(area float64) *SSAORenderingPipeline {
	s.p.Set("area", area)
	return s
}

// Base returns the Base property of class SSAORenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#base
func (s *SSAORenderingPipeline) Base() float64 {
	retVal := s.p.Get("base")
	return retVal.Float()
}

// SetBase sets the Base property of class SSAORenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#base
func (s *SSAORenderingPipeline) SetBase(base float64) *SSAORenderingPipeline {
	s.p.Set("base", base)
	return s
}

// FallOff returns the FallOff property of class SSAORenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#falloff
func (s *SSAORenderingPipeline) FallOff() float64 {
	retVal := s.p.Get("fallOff")
	return retVal.Float()
}

// SetFallOff sets the FallOff property of class SSAORenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#falloff
func (s *SSAORenderingPipeline) SetFallOff(fallOff float64) *SSAORenderingPipeline {
	s.p.Set("fallOff", fallOff)
	return s
}

// Radius returns the Radius property of class SSAORenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#radius
func (s *SSAORenderingPipeline) Radius() float64 {
	retVal := s.p.Get("radius")
	return retVal.Float()
}

// SetRadius sets the Radius property of class SSAORenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#radius
func (s *SSAORenderingPipeline) SetRadius(radius float64) *SSAORenderingPipeline {
	s.p.Set("radius", radius)
	return s
}

// Scene returns the Scene property of class SSAORenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#scene
func (s *SSAORenderingPipeline) Scene() *Scene {
	retVal := s.p.Get("scene")
	return SceneFromJSObject(retVal, s.ctx)
}

// SetScene sets the Scene property of class SSAORenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#scene
func (s *SSAORenderingPipeline) SetScene(scene *Scene) *SSAORenderingPipeline {
	s.p.Set("scene", scene.JSObject())
	return s
}

// TotalStrength returns the TotalStrength property of class SSAORenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#totalstrength
func (s *SSAORenderingPipeline) TotalStrength() float64 {
	retVal := s.p.Get("totalStrength")
	return retVal.Float()
}

// SetTotalStrength sets the TotalStrength property of class SSAORenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssaorenderingpipeline#totalstrength
func (s *SSAORenderingPipeline) SetTotalStrength(totalStrength float64) *SSAORenderingPipeline {
	s.p.Set("totalStrength", totalStrength)
	return s
}
