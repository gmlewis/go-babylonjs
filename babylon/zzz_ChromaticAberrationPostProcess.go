// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ChromaticAberrationPostProcess represents a babylon.js ChromaticAberrationPostProcess.
// The ChromaticAberrationPostProcess separates the rgb channels in an image to produce chromatic distortion around the edges of the screen
type ChromaticAberrationPostProcess struct {
	*PostProcess
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (c *ChromaticAberrationPostProcess) JSObject() js.Value { return c.p }

// ChromaticAberrationPostProcess returns a ChromaticAberrationPostProcess JavaScript class.
func (ba *Babylon) ChromaticAberrationPostProcess() *ChromaticAberrationPostProcess {
	p := ba.ctx.Get("ChromaticAberrationPostProcess")
	return ChromaticAberrationPostProcessFromJSObject(p, ba.ctx)
}

// ChromaticAberrationPostProcessFromJSObject returns a wrapped ChromaticAberrationPostProcess JavaScript class.
func ChromaticAberrationPostProcessFromJSObject(p js.Value, ctx js.Value) *ChromaticAberrationPostProcess {
	return &ChromaticAberrationPostProcess{PostProcess: PostProcessFromJSObject(p, ctx), ctx: ctx}
}

// ChromaticAberrationPostProcessArrayToJSArray returns a JavaScript Array for the wrapped array.
func ChromaticAberrationPostProcessArrayToJSArray(array []*ChromaticAberrationPostProcess) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewChromaticAberrationPostProcessOpts contains optional parameters for NewChromaticAberrationPostProcess.
type NewChromaticAberrationPostProcessOpts struct {
	SamplingMode     *float64
	Engine           *Engine
	Reusable         *bool
	TextureType      *float64
	BlockCompilation *bool
}

// NewChromaticAberrationPostProcess returns a new ChromaticAberrationPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.chromaticaberrationpostprocess
func (ba *Babylon) NewChromaticAberrationPostProcess(name string, screenWidth float64, screenHeight float64, options float64, camera *Camera, opts *NewChromaticAberrationPostProcessOpts) *ChromaticAberrationPostProcess {
	if opts == nil {
		opts = &NewChromaticAberrationPostProcessOpts{}
	}

	args := make([]interface{}, 0, 5+5)

	args = append(args, name)
	args = append(args, screenWidth)
	args = append(args, screenHeight)
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
	if opts.TextureType == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.TextureType)
	}
	if opts.BlockCompilation == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.BlockCompilation)
	}

	p := ba.ctx.Get("ChromaticAberrationPostProcess").New(args...)
	return ChromaticAberrationPostProcessFromJSObject(p, ba.ctx)
}

/*

// AberrationAmount returns the AberrationAmount property of class ChromaticAberrationPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.chromaticaberrationpostprocess#aberrationamount
func (c *ChromaticAberrationPostProcess) AberrationAmount(aberrationAmount float64) *ChromaticAberrationPostProcess {
	p := ba.ctx.Get("ChromaticAberrationPostProcess").New(aberrationAmount)
	return ChromaticAberrationPostProcessFromJSObject(p, ba.ctx)
}

// SetAberrationAmount sets the AberrationAmount property of class ChromaticAberrationPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.chromaticaberrationpostprocess#aberrationamount
func (c *ChromaticAberrationPostProcess) SetAberrationAmount(aberrationAmount float64) *ChromaticAberrationPostProcess {
	p := ba.ctx.Get("ChromaticAberrationPostProcess").New(aberrationAmount)
	return ChromaticAberrationPostProcessFromJSObject(p, ba.ctx)
}

// CenterPosition returns the CenterPosition property of class ChromaticAberrationPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.chromaticaberrationpostprocess#centerposition
func (c *ChromaticAberrationPostProcess) CenterPosition(centerPosition *Vector2) *ChromaticAberrationPostProcess {
	p := ba.ctx.Get("ChromaticAberrationPostProcess").New(centerPosition.JSObject())
	return ChromaticAberrationPostProcessFromJSObject(p, ba.ctx)
}

// SetCenterPosition sets the CenterPosition property of class ChromaticAberrationPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.chromaticaberrationpostprocess#centerposition
func (c *ChromaticAberrationPostProcess) SetCenterPosition(centerPosition *Vector2) *ChromaticAberrationPostProcess {
	p := ba.ctx.Get("ChromaticAberrationPostProcess").New(centerPosition.JSObject())
	return ChromaticAberrationPostProcessFromJSObject(p, ba.ctx)
}

// Direction returns the Direction property of class ChromaticAberrationPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.chromaticaberrationpostprocess#direction
func (c *ChromaticAberrationPostProcess) Direction(direction *Vector2) *ChromaticAberrationPostProcess {
	p := ba.ctx.Get("ChromaticAberrationPostProcess").New(direction.JSObject())
	return ChromaticAberrationPostProcessFromJSObject(p, ba.ctx)
}

// SetDirection sets the Direction property of class ChromaticAberrationPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.chromaticaberrationpostprocess#direction
func (c *ChromaticAberrationPostProcess) SetDirection(direction *Vector2) *ChromaticAberrationPostProcess {
	p := ba.ctx.Get("ChromaticAberrationPostProcess").New(direction.JSObject())
	return ChromaticAberrationPostProcessFromJSObject(p, ba.ctx)
}

// RadialIntensity returns the RadialIntensity property of class ChromaticAberrationPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.chromaticaberrationpostprocess#radialintensity
func (c *ChromaticAberrationPostProcess) RadialIntensity(radialIntensity float64) *ChromaticAberrationPostProcess {
	p := ba.ctx.Get("ChromaticAberrationPostProcess").New(radialIntensity)
	return ChromaticAberrationPostProcessFromJSObject(p, ba.ctx)
}

// SetRadialIntensity sets the RadialIntensity property of class ChromaticAberrationPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.chromaticaberrationpostprocess#radialintensity
func (c *ChromaticAberrationPostProcess) SetRadialIntensity(radialIntensity float64) *ChromaticAberrationPostProcess {
	p := ba.ctx.Get("ChromaticAberrationPostProcess").New(radialIntensity)
	return ChromaticAberrationPostProcessFromJSObject(p, ba.ctx)
}

*/
