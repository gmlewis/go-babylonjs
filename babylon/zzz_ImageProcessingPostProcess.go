// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ImageProcessingPostProcess represents a babylon.js ImageProcessingPostProcess.
// ImageProcessingPostProcess
//
// See: https://doc.babylonjs.com/how_to/how_to_use_postprocesses#imageprocessing
type ImageProcessingPostProcess struct {
	*PostProcess
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *ImageProcessingPostProcess) JSObject() js.Value { return i.p }

// ImageProcessingPostProcess returns a ImageProcessingPostProcess JavaScript class.
func (ba *Babylon) ImageProcessingPostProcess() *ImageProcessingPostProcess {
	p := ba.ctx.Get("ImageProcessingPostProcess")
	return ImageProcessingPostProcessFromJSObject(p, ba.ctx)
}

// ImageProcessingPostProcessFromJSObject returns a wrapped ImageProcessingPostProcess JavaScript class.
func ImageProcessingPostProcessFromJSObject(p js.Value, ctx js.Value) *ImageProcessingPostProcess {
	return &ImageProcessingPostProcess{PostProcess: PostProcessFromJSObject(p, ctx), ctx: ctx}
}

// ImageProcessingPostProcessArrayToJSArray returns a JavaScript Array for the wrapped array.
func ImageProcessingPostProcessArrayToJSArray(array []*ImageProcessingPostProcess) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewImageProcessingPostProcessOpts contains optional parameters for NewImageProcessingPostProcess.
type NewImageProcessingPostProcessOpts struct {
	Camera                       *Camera
	SamplingMode                 *float64
	Engine                       *Engine
	Reusable                     *bool
	TextureType                  *float64
	ImageProcessingConfiguration *ImageProcessingConfiguration
}

// NewImageProcessingPostProcess returns a new ImageProcessingPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#constructor
func (ba *Babylon) NewImageProcessingPostProcess(name string, options float64, opts *NewImageProcessingPostProcessOpts) *ImageProcessingPostProcess {
	if opts == nil {
		opts = &NewImageProcessingPostProcessOpts{}
	}

	args := make([]interface{}, 0, 2+6)

	args = append(args, name)
	args = append(args, options)

	if opts.Camera == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Camera.JSObject())
	}
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
	if opts.ImageProcessingConfiguration == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ImageProcessingConfiguration.JSObject())
	}

	p := ba.ctx.Get("ImageProcessingPostProcess").New(args...)
	return ImageProcessingPostProcessFromJSObject(p, ba.ctx)
}

// ImageProcessingPostProcessDisposeOpts contains optional parameters for ImageProcessingPostProcess.Dispose.
type ImageProcessingPostProcessDisposeOpts struct {
	Camera *Camera
}

// Dispose calls the Dispose method on the ImageProcessingPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#dispose
func (i *ImageProcessingPostProcess) Dispose(opts *ImageProcessingPostProcessDisposeOpts) {
	if opts == nil {
		opts = &ImageProcessingPostProcessDisposeOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Camera == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Camera.JSObject())
	}

	i.p.Call("dispose", args...)
}

// GetClassName calls the GetClassName method on the ImageProcessingPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#getclassname
func (i *ImageProcessingPostProcess) GetClassName() string {

	retVal := i.p.Call("getClassName")
	return retVal.String()
}

// ColorCurves returns the ColorCurves property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#colorcurves
func (i *ImageProcessingPostProcess) ColorCurves() *ColorCurves {
	retVal := i.p.Get("colorCurves")
	return ColorCurvesFromJSObject(retVal, i.ctx)
}

// SetColorCurves sets the ColorCurves property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#colorcurves
func (i *ImageProcessingPostProcess) SetColorCurves(colorCurves *ColorCurves) *ImageProcessingPostProcess {
	i.p.Set("colorCurves", colorCurves.JSObject())
	return i
}

// ColorCurvesEnabled returns the ColorCurvesEnabled property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#colorcurvesenabled
func (i *ImageProcessingPostProcess) ColorCurvesEnabled() bool {
	retVal := i.p.Get("colorCurvesEnabled")
	return retVal.Bool()
}

// SetColorCurvesEnabled sets the ColorCurvesEnabled property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#colorcurvesenabled
func (i *ImageProcessingPostProcess) SetColorCurvesEnabled(colorCurvesEnabled bool) *ImageProcessingPostProcess {
	i.p.Set("colorCurvesEnabled", colorCurvesEnabled)
	return i
}

// ColorGradingEnabled returns the ColorGradingEnabled property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#colorgradingenabled
func (i *ImageProcessingPostProcess) ColorGradingEnabled() bool {
	retVal := i.p.Get("colorGradingEnabled")
	return retVal.Bool()
}

// SetColorGradingEnabled sets the ColorGradingEnabled property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#colorgradingenabled
func (i *ImageProcessingPostProcess) SetColorGradingEnabled(colorGradingEnabled bool) *ImageProcessingPostProcess {
	i.p.Set("colorGradingEnabled", colorGradingEnabled)
	return i
}

// ColorGradingTexture returns the ColorGradingTexture property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#colorgradingtexture
func (i *ImageProcessingPostProcess) ColorGradingTexture() *BaseTexture {
	retVal := i.p.Get("colorGradingTexture")
	return BaseTextureFromJSObject(retVal, i.ctx)
}

// SetColorGradingTexture sets the ColorGradingTexture property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#colorgradingtexture
func (i *ImageProcessingPostProcess) SetColorGradingTexture(colorGradingTexture *BaseTexture) *ImageProcessingPostProcess {
	i.p.Set("colorGradingTexture", colorGradingTexture.JSObject())
	return i
}

// Contrast returns the Contrast property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#contrast
func (i *ImageProcessingPostProcess) Contrast() float64 {
	retVal := i.p.Get("contrast")
	return retVal.Float()
}

// SetContrast sets the Contrast property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#contrast
func (i *ImageProcessingPostProcess) SetContrast(contrast float64) *ImageProcessingPostProcess {
	i.p.Set("contrast", contrast)
	return i
}

// Exposure returns the Exposure property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#exposure
func (i *ImageProcessingPostProcess) Exposure() float64 {
	retVal := i.p.Get("exposure")
	return retVal.Float()
}

// SetExposure sets the Exposure property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#exposure
func (i *ImageProcessingPostProcess) SetExposure(exposure float64) *ImageProcessingPostProcess {
	i.p.Set("exposure", exposure)
	return i
}

// FromLinearSpace returns the FromLinearSpace property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#fromlinearspace
func (i *ImageProcessingPostProcess) FromLinearSpace() bool {
	retVal := i.p.Get("fromLinearSpace")
	return retVal.Bool()
}

// SetFromLinearSpace sets the FromLinearSpace property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#fromlinearspace
func (i *ImageProcessingPostProcess) SetFromLinearSpace(fromLinearSpace bool) *ImageProcessingPostProcess {
	i.p.Set("fromLinearSpace", fromLinearSpace)
	return i
}

// ImageProcessingConfiguration returns the ImageProcessingConfiguration property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#imageprocessingconfiguration
func (i *ImageProcessingPostProcess) ImageProcessingConfiguration() *ImageProcessingConfiguration {
	retVal := i.p.Get("imageProcessingConfiguration")
	return ImageProcessingConfigurationFromJSObject(retVal, i.ctx)
}

// SetImageProcessingConfiguration sets the ImageProcessingConfiguration property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#imageprocessingconfiguration
func (i *ImageProcessingPostProcess) SetImageProcessingConfiguration(imageProcessingConfiguration *ImageProcessingConfiguration) *ImageProcessingPostProcess {
	i.p.Set("imageProcessingConfiguration", imageProcessingConfiguration.JSObject())
	return i
}

// ToneMappingEnabled returns the ToneMappingEnabled property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#tonemappingenabled
func (i *ImageProcessingPostProcess) ToneMappingEnabled() bool {
	retVal := i.p.Get("toneMappingEnabled")
	return retVal.Bool()
}

// SetToneMappingEnabled sets the ToneMappingEnabled property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#tonemappingenabled
func (i *ImageProcessingPostProcess) SetToneMappingEnabled(toneMappingEnabled bool) *ImageProcessingPostProcess {
	i.p.Set("toneMappingEnabled", toneMappingEnabled)
	return i
}

// ToneMappingType returns the ToneMappingType property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#tonemappingtype
func (i *ImageProcessingPostProcess) ToneMappingType() float64 {
	retVal := i.p.Get("toneMappingType")
	return retVal.Float()
}

// SetToneMappingType sets the ToneMappingType property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#tonemappingtype
func (i *ImageProcessingPostProcess) SetToneMappingType(toneMappingType float64) *ImageProcessingPostProcess {
	i.p.Set("toneMappingType", toneMappingType)
	return i
}

// VignetteBlendMode returns the VignetteBlendMode property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#vignetteblendmode
func (i *ImageProcessingPostProcess) VignetteBlendMode() float64 {
	retVal := i.p.Get("vignetteBlendMode")
	return retVal.Float()
}

// SetVignetteBlendMode sets the VignetteBlendMode property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#vignetteblendmode
func (i *ImageProcessingPostProcess) SetVignetteBlendMode(vignetteBlendMode float64) *ImageProcessingPostProcess {
	i.p.Set("vignetteBlendMode", vignetteBlendMode)
	return i
}

// VignetteCameraFov returns the VignetteCameraFov property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#vignettecamerafov
func (i *ImageProcessingPostProcess) VignetteCameraFov() float64 {
	retVal := i.p.Get("vignetteCameraFov")
	return retVal.Float()
}

// SetVignetteCameraFov sets the VignetteCameraFov property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#vignettecamerafov
func (i *ImageProcessingPostProcess) SetVignetteCameraFov(vignetteCameraFov float64) *ImageProcessingPostProcess {
	i.p.Set("vignetteCameraFov", vignetteCameraFov)
	return i
}

// VignetteCentreX returns the VignetteCentreX property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#vignettecentrex
func (i *ImageProcessingPostProcess) VignetteCentreX() float64 {
	retVal := i.p.Get("vignetteCentreX")
	return retVal.Float()
}

// SetVignetteCentreX sets the VignetteCentreX property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#vignettecentrex
func (i *ImageProcessingPostProcess) SetVignetteCentreX(vignetteCentreX float64) *ImageProcessingPostProcess {
	i.p.Set("vignetteCentreX", vignetteCentreX)
	return i
}

// VignetteCentreY returns the VignetteCentreY property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#vignettecentrey
func (i *ImageProcessingPostProcess) VignetteCentreY() float64 {
	retVal := i.p.Get("vignetteCentreY")
	return retVal.Float()
}

// SetVignetteCentreY sets the VignetteCentreY property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#vignettecentrey
func (i *ImageProcessingPostProcess) SetVignetteCentreY(vignetteCentreY float64) *ImageProcessingPostProcess {
	i.p.Set("vignetteCentreY", vignetteCentreY)
	return i
}

// VignetteColor returns the VignetteColor property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#vignettecolor
func (i *ImageProcessingPostProcess) VignetteColor() *Color4 {
	retVal := i.p.Get("vignetteColor")
	return Color4FromJSObject(retVal, i.ctx)
}

// SetVignetteColor sets the VignetteColor property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#vignettecolor
func (i *ImageProcessingPostProcess) SetVignetteColor(vignetteColor *Color4) *ImageProcessingPostProcess {
	i.p.Set("vignetteColor", vignetteColor.JSObject())
	return i
}

// VignetteEnabled returns the VignetteEnabled property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#vignetteenabled
func (i *ImageProcessingPostProcess) VignetteEnabled() bool {
	retVal := i.p.Get("vignetteEnabled")
	return retVal.Bool()
}

// SetVignetteEnabled sets the VignetteEnabled property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#vignetteenabled
func (i *ImageProcessingPostProcess) SetVignetteEnabled(vignetteEnabled bool) *ImageProcessingPostProcess {
	i.p.Set("vignetteEnabled", vignetteEnabled)
	return i
}

// VignetteStretch returns the VignetteStretch property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#vignettestretch
func (i *ImageProcessingPostProcess) VignetteStretch() float64 {
	retVal := i.p.Get("vignetteStretch")
	return retVal.Float()
}

// SetVignetteStretch sets the VignetteStretch property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#vignettestretch
func (i *ImageProcessingPostProcess) SetVignetteStretch(vignetteStretch float64) *ImageProcessingPostProcess {
	i.p.Set("vignetteStretch", vignetteStretch)
	return i
}

// VignetteWeight returns the VignetteWeight property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#vignetteweight
func (i *ImageProcessingPostProcess) VignetteWeight() float64 {
	retVal := i.p.Get("vignetteWeight")
	return retVal.Float()
}

// SetVignetteWeight sets the VignetteWeight property of class ImageProcessingPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.imageprocessingpostprocess#vignetteweight
func (i *ImageProcessingPostProcess) SetVignetteWeight(vignetteWeight float64) *ImageProcessingPostProcess {
	i.p.Set("vignetteWeight", vignetteWeight)
	return i
}
