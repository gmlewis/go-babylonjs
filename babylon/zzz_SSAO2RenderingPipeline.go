// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SSAO2RenderingPipeline represents a babylon.js SSAO2RenderingPipeline.
// Render pipeline to produce ssao effect
type SSAO2RenderingPipeline struct {
	*PostProcessRenderPipeline
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *SSAO2RenderingPipeline) JSObject() js.Value { return s.p }

// SSAO2RenderingPipeline returns a SSAO2RenderingPipeline JavaScript class.
func (ba *Babylon) SSAO2RenderingPipeline() *SSAO2RenderingPipeline {
	p := ba.ctx.Get("SSAO2RenderingPipeline")
	return SSAO2RenderingPipelineFromJSObject(p, ba.ctx)
}

// SSAO2RenderingPipelineFromJSObject returns a wrapped SSAO2RenderingPipeline JavaScript class.
func SSAO2RenderingPipelineFromJSObject(p js.Value, ctx js.Value) *SSAO2RenderingPipeline {
	return &SSAO2RenderingPipeline{PostProcessRenderPipeline: PostProcessRenderPipelineFromJSObject(p, ctx), ctx: ctx}
}

// SSAO2RenderingPipelineArrayToJSArray returns a JavaScript Array for the wrapped array.
func SSAO2RenderingPipelineArrayToJSArray(array []*SSAO2RenderingPipeline) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewSSAO2RenderingPipelineOpts contains optional parameters for NewSSAO2RenderingPipeline.
type NewSSAO2RenderingPipelineOpts struct {
	Cameras []*Camera
}

// NewSSAO2RenderingPipeline returns a new SSAO2RenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.ssao2renderingpipeline
func (ba *Babylon) NewSSAO2RenderingPipeline(name string, scene *Scene, ratio interface{}, opts *NewSSAO2RenderingPipelineOpts) *SSAO2RenderingPipeline {
	if opts == nil {
		opts = &NewSSAO2RenderingPipelineOpts{}
	}

	args := make([]interface{}, 0, 3+1)

	args = append(args, name)
	args = append(args, scene.JSObject())
	args = append(args, ratio)

	if opts.Cameras == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, CameraArrayToJSArray(opts.Cameras))
	}

	p := ba.ctx.Get("SSAO2RenderingPipeline").New(args...)
	return SSAO2RenderingPipelineFromJSObject(p, ba.ctx)
}

// SSAO2RenderingPipelineDisposeOpts contains optional parameters for SSAO2RenderingPipeline.Dispose.
type SSAO2RenderingPipelineDisposeOpts struct {
	DisableGeometryBufferRenderer *bool
}

// Dispose calls the Dispose method on the SSAO2RenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.ssao2renderingpipeline#dispose
func (s *SSAO2RenderingPipeline) Dispose(opts *SSAO2RenderingPipelineDisposeOpts) {
	if opts == nil {
		opts = &SSAO2RenderingPipelineDisposeOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.DisableGeometryBufferRenderer == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.DisableGeometryBufferRenderer)
	}

	s.p.Call("dispose", args...)
}

// GetClassName calls the GetClassName method on the SSAO2RenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.ssao2renderingpipeline#getclassname
func (s *SSAO2RenderingPipeline) GetClassName() string {

	retVal := s.p.Call("getClassName")
	return retVal.String()
}

// Parse calls the Parse method on the SSAO2RenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.ssao2renderingpipeline#parse
func (s *SSAO2RenderingPipeline) Parse(source interface{}, scene *Scene, rootUrl string) *SSAO2RenderingPipeline {

	args := make([]interface{}, 0, 3+0)

	args = append(args, source)
	args = append(args, scene.JSObject())
	args = append(args, rootUrl)

	retVal := s.p.Call("Parse", args...)
	return SSAO2RenderingPipelineFromJSObject(retVal, s.ctx)
}

// Serialize calls the Serialize method on the SSAO2RenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.ssao2renderingpipeline#serialize
func (s *SSAO2RenderingPipeline) Serialize() interface{} {

	retVal := s.p.Call("serialize")
	return retVal
}

// Base returns the Base property of class SSAO2RenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssao2renderingpipeline#base
func (s *SSAO2RenderingPipeline) Base() float64 {
	retVal := s.p.Get("base")
	return retVal.Float()
}

// SetBase sets the Base property of class SSAO2RenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssao2renderingpipeline#base
func (s *SSAO2RenderingPipeline) SetBase(base float64) *SSAO2RenderingPipeline {
	s.p.Set("base", base)
	return s
}

// ExpensiveBlur returns the ExpensiveBlur property of class SSAO2RenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssao2renderingpipeline#expensiveblur
func (s *SSAO2RenderingPipeline) ExpensiveBlur() bool {
	retVal := s.p.Get("expensiveBlur")
	return retVal.Bool()
}

// SetExpensiveBlur sets the ExpensiveBlur property of class SSAO2RenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssao2renderingpipeline#expensiveblur
func (s *SSAO2RenderingPipeline) SetExpensiveBlur(expensiveBlur bool) *SSAO2RenderingPipeline {
	s.p.Set("expensiveBlur", expensiveBlur)
	return s
}

// IsSupported returns the IsSupported property of class SSAO2RenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssao2renderingpipeline#issupported
func (s *SSAO2RenderingPipeline) IsSupported() bool {
	retVal := s.p.Get("IsSupported")
	return retVal.Bool()
}

// SetIsSupported sets the IsSupported property of class SSAO2RenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssao2renderingpipeline#issupported
func (s *SSAO2RenderingPipeline) SetIsSupported(IsSupported bool) *SSAO2RenderingPipeline {
	s.p.Set("IsSupported", IsSupported)
	return s
}

// MaxZ returns the MaxZ property of class SSAO2RenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssao2renderingpipeline#maxz
func (s *SSAO2RenderingPipeline) MaxZ() float64 {
	retVal := s.p.Get("maxZ")
	return retVal.Float()
}

// SetMaxZ sets the MaxZ property of class SSAO2RenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssao2renderingpipeline#maxz
func (s *SSAO2RenderingPipeline) SetMaxZ(maxZ float64) *SSAO2RenderingPipeline {
	s.p.Set("maxZ", maxZ)
	return s
}

// MinZAspect returns the MinZAspect property of class SSAO2RenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssao2renderingpipeline#minzaspect
func (s *SSAO2RenderingPipeline) MinZAspect() float64 {
	retVal := s.p.Get("minZAspect")
	return retVal.Float()
}

// SetMinZAspect sets the MinZAspect property of class SSAO2RenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssao2renderingpipeline#minzaspect
func (s *SSAO2RenderingPipeline) SetMinZAspect(minZAspect float64) *SSAO2RenderingPipeline {
	s.p.Set("minZAspect", minZAspect)
	return s
}

// Radius returns the Radius property of class SSAO2RenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssao2renderingpipeline#radius
func (s *SSAO2RenderingPipeline) Radius() float64 {
	retVal := s.p.Get("radius")
	return retVal.Float()
}

// SetRadius sets the Radius property of class SSAO2RenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssao2renderingpipeline#radius
func (s *SSAO2RenderingPipeline) SetRadius(radius float64) *SSAO2RenderingPipeline {
	s.p.Set("radius", radius)
	return s
}

// Samples returns the Samples property of class SSAO2RenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssao2renderingpipeline#samples
func (s *SSAO2RenderingPipeline) Samples() float64 {
	retVal := s.p.Get("samples")
	return retVal.Float()
}

// SetSamples sets the Samples property of class SSAO2RenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssao2renderingpipeline#samples
func (s *SSAO2RenderingPipeline) SetSamples(samples float64) *SSAO2RenderingPipeline {
	s.p.Set("samples", samples)
	return s
}

// Scene returns the Scene property of class SSAO2RenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssao2renderingpipeline#scene
func (s *SSAO2RenderingPipeline) Scene() *Scene {
	retVal := s.p.Get("scene")
	return SceneFromJSObject(retVal, s.ctx)
}

// SetScene sets the Scene property of class SSAO2RenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssao2renderingpipeline#scene
func (s *SSAO2RenderingPipeline) SetScene(scene *Scene) *SSAO2RenderingPipeline {
	s.p.Set("scene", scene.JSObject())
	return s
}

// TextureSamples returns the TextureSamples property of class SSAO2RenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssao2renderingpipeline#texturesamples
func (s *SSAO2RenderingPipeline) TextureSamples() float64 {
	retVal := s.p.Get("textureSamples")
	return retVal.Float()
}

// SetTextureSamples sets the TextureSamples property of class SSAO2RenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssao2renderingpipeline#texturesamples
func (s *SSAO2RenderingPipeline) SetTextureSamples(textureSamples float64) *SSAO2RenderingPipeline {
	s.p.Set("textureSamples", textureSamples)
	return s
}

// TotalStrength returns the TotalStrength property of class SSAO2RenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssao2renderingpipeline#totalstrength
func (s *SSAO2RenderingPipeline) TotalStrength() float64 {
	retVal := s.p.Get("totalStrength")
	return retVal.Float()
}

// SetTotalStrength sets the TotalStrength property of class SSAO2RenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.ssao2renderingpipeline#totalstrength
func (s *SSAO2RenderingPipeline) SetTotalStrength(totalStrength float64) *SSAO2RenderingPipeline {
	s.p.Set("totalStrength", totalStrength)
	return s
}
