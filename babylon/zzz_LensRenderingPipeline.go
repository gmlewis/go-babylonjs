// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// LensRenderingPipeline represents a babylon.js LensRenderingPipeline.
// BABYLON.JS Chromatic Aberration GLSL Shader
// Author: Olivier Guyot
// Separates very slightly R, G and B colors on the edges of the screen
// Inspired by Francois Tarlier &amp;amp; Martins Upitis
type LensRenderingPipeline struct {
	*PostProcessRenderPipeline
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (l *LensRenderingPipeline) JSObject() js.Value { return l.p }

// LensRenderingPipeline returns a LensRenderingPipeline JavaScript class.
func (ba *Babylon) LensRenderingPipeline() *LensRenderingPipeline {
	p := ba.ctx.Get("LensRenderingPipeline")
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

// LensRenderingPipelineFromJSObject returns a wrapped LensRenderingPipeline JavaScript class.
func LensRenderingPipelineFromJSObject(p js.Value, ctx js.Value) *LensRenderingPipeline {
	return &LensRenderingPipeline{PostProcessRenderPipeline: PostProcessRenderPipelineFromJSObject(p, ctx), ctx: ctx}
}

// LensRenderingPipelineArrayToJSArray returns a JavaScript Array for the wrapped array.
func LensRenderingPipelineArrayToJSArray(array []*LensRenderingPipeline) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewLensRenderingPipelineOpts contains optional parameters for NewLensRenderingPipeline.
type NewLensRenderingPipelineOpts struct {
	Ratio   *float64
	Cameras *Camera
}

// NewLensRenderingPipeline returns a new LensRenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline
func (ba *Babylon) NewLensRenderingPipeline(name string, parameters interface{}, scene *Scene, opts *NewLensRenderingPipelineOpts) *LensRenderingPipeline {
	if opts == nil {
		opts = &NewLensRenderingPipelineOpts{}
	}

	args := make([]interface{}, 0, 3+2)

	args = append(args, name)
	args = append(args, parameters)
	args = append(args, scene.JSObject())

	if opts.Ratio == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Ratio)
	}
	if opts.Cameras == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Cameras.JSObject())
	}

	p := ba.ctx.Get("LensRenderingPipeline").New(args...)
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

// AddEffect calls the AddEffect method on the LensRenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#addeffect
func (l *LensRenderingPipeline) AddEffect(renderEffect *PostProcessRenderEffect) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, renderEffect.JSObject())

	l.p.Call("addEffect", args...)
}

// DisableChromaticAberration calls the DisableChromaticAberration method on the LensRenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#disablechromaticaberration
func (l *LensRenderingPipeline) DisableChromaticAberration() {

	l.p.Call("disableChromaticAberration")
}

// DisableDepthOfField calls the DisableDepthOfField method on the LensRenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#disabledepthoffield
func (l *LensRenderingPipeline) DisableDepthOfField() {

	l.p.Call("disableDepthOfField")
}

// DisableEdgeBlur calls the DisableEdgeBlur method on the LensRenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#disableedgeblur
func (l *LensRenderingPipeline) DisableEdgeBlur() {

	l.p.Call("disableEdgeBlur")
}

// DisableEdgeDistortion calls the DisableEdgeDistortion method on the LensRenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#disableedgedistortion
func (l *LensRenderingPipeline) DisableEdgeDistortion() {

	l.p.Call("disableEdgeDistortion")
}

// DisableGrain calls the DisableGrain method on the LensRenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#disablegrain
func (l *LensRenderingPipeline) DisableGrain() {

	l.p.Call("disableGrain")
}

// DisableHighlights calls the DisableHighlights method on the LensRenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#disablehighlights
func (l *LensRenderingPipeline) DisableHighlights() {

	l.p.Call("disableHighlights")
}

// DisableNoiseBlur calls the DisableNoiseBlur method on the LensRenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#disablenoiseblur
func (l *LensRenderingPipeline) DisableNoiseBlur() {

	l.p.Call("disableNoiseBlur")
}

// DisablePentagonBokeh calls the DisablePentagonBokeh method on the LensRenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#disablepentagonbokeh
func (l *LensRenderingPipeline) DisablePentagonBokeh() {

	l.p.Call("disablePentagonBokeh")
}

// LensRenderingPipelineDisposeOpts contains optional parameters for LensRenderingPipeline.Dispose.
type LensRenderingPipelineDisposeOpts struct {
	DisableDepthRender *bool
}

// Dispose calls the Dispose method on the LensRenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#dispose
func (l *LensRenderingPipeline) Dispose(opts *LensRenderingPipelineDisposeOpts) {
	if opts == nil {
		opts = &LensRenderingPipelineDisposeOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.DisableDepthRender == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.DisableDepthRender)
	}

	l.p.Call("dispose", args...)
}

// EnableNoiseBlur calls the EnableNoiseBlur method on the LensRenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#enablenoiseblur
func (l *LensRenderingPipeline) EnableNoiseBlur() {

	l.p.Call("enableNoiseBlur")
}

// EnablePentagonBokeh calls the EnablePentagonBokeh method on the LensRenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#enablepentagonbokeh
func (l *LensRenderingPipeline) EnablePentagonBokeh() {

	l.p.Call("enablePentagonBokeh")
}

// GetClassName calls the GetClassName method on the LensRenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#getclassname
func (l *LensRenderingPipeline) GetClassName() string {

	retVal := l.p.Call("getClassName")
	return retVal.String()
}

// SetAperture calls the SetAperture method on the LensRenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#setaperture
func (l *LensRenderingPipeline) SetAperture(amount float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, amount)

	l.p.Call("setAperture", args...)
}

// SetChromaticAberration calls the SetChromaticAberration method on the LensRenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#setchromaticaberration
func (l *LensRenderingPipeline) SetChromaticAberration(amount float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, amount)

	l.p.Call("setChromaticAberration", args...)
}

// SetDarkenOutOfFocus calls the SetDarkenOutOfFocus method on the LensRenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#setdarkenoutoffocus
func (l *LensRenderingPipeline) SetDarkenOutOfFocus(amount float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, amount)

	l.p.Call("setDarkenOutOfFocus", args...)
}

// SetEdgeBlur calls the SetEdgeBlur method on the LensRenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#setedgeblur
func (l *LensRenderingPipeline) SetEdgeBlur(amount float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, amount)

	l.p.Call("setEdgeBlur", args...)
}

// SetEdgeDistortion calls the SetEdgeDistortion method on the LensRenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#setedgedistortion
func (l *LensRenderingPipeline) SetEdgeDistortion(amount float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, amount)

	l.p.Call("setEdgeDistortion", args...)
}

// SetFocusDistance calls the SetFocusDistance method on the LensRenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#setfocusdistance
func (l *LensRenderingPipeline) SetFocusDistance(amount float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, amount)

	l.p.Call("setFocusDistance", args...)
}

// SetGrainAmount calls the SetGrainAmount method on the LensRenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#setgrainamount
func (l *LensRenderingPipeline) SetGrainAmount(amount float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, amount)

	l.p.Call("setGrainAmount", args...)
}

// SetHighlightsGain calls the SetHighlightsGain method on the LensRenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#sethighlightsgain
func (l *LensRenderingPipeline) SetHighlightsGain(amount float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, amount)

	l.p.Call("setHighlightsGain", args...)
}

// SetHighlightsThreshold calls the SetHighlightsThreshold method on the LensRenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#sethighlightsthreshold
func (l *LensRenderingPipeline) SetHighlightsThreshold(amount float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, amount)

	l.p.Call("setHighlightsThreshold", args...)
}

/*

// BlurNoise returns the BlurNoise property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#blurnoise
func (l *LensRenderingPipeline) BlurNoise(blurNoise bool) *LensRenderingPipeline {
	p := ba.ctx.Get("LensRenderingPipeline").New(blurNoise)
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

// SetBlurNoise sets the BlurNoise property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#blurnoise
func (l *LensRenderingPipeline) SetBlurNoise(blurNoise bool) *LensRenderingPipeline {
	p := ba.ctx.Get("LensRenderingPipeline").New(blurNoise)
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

// Cameras returns the Cameras property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#cameras
func (l *LensRenderingPipeline) Cameras(cameras *Camera) *LensRenderingPipeline {
	p := ba.ctx.Get("LensRenderingPipeline").New(cameras.JSObject())
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

// SetCameras sets the Cameras property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#cameras
func (l *LensRenderingPipeline) SetCameras(cameras *Camera) *LensRenderingPipeline {
	p := ba.ctx.Get("LensRenderingPipeline").New(cameras.JSObject())
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

// ChromaticAberration returns the ChromaticAberration property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#chromaticaberration
func (l *LensRenderingPipeline) ChromaticAberration(chromaticAberration float64) *LensRenderingPipeline {
	p := ba.ctx.Get("LensRenderingPipeline").New(chromaticAberration)
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

// SetChromaticAberration sets the ChromaticAberration property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#chromaticaberration
func (l *LensRenderingPipeline) SetChromaticAberration(chromaticAberration float64) *LensRenderingPipeline {
	p := ba.ctx.Get("LensRenderingPipeline").New(chromaticAberration)
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

// DarkenOutOfFocus returns the DarkenOutOfFocus property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#darkenoutoffocus
func (l *LensRenderingPipeline) DarkenOutOfFocus(darkenOutOfFocus float64) *LensRenderingPipeline {
	p := ba.ctx.Get("LensRenderingPipeline").New(darkenOutOfFocus)
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

// SetDarkenOutOfFocus sets the DarkenOutOfFocus property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#darkenoutoffocus
func (l *LensRenderingPipeline) SetDarkenOutOfFocus(darkenOutOfFocus float64) *LensRenderingPipeline {
	p := ba.ctx.Get("LensRenderingPipeline").New(darkenOutOfFocus)
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

// DofAperture returns the DofAperture property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#dofaperture
func (l *LensRenderingPipeline) DofAperture(dofAperture float64) *LensRenderingPipeline {
	p := ba.ctx.Get("LensRenderingPipeline").New(dofAperture)
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

// SetDofAperture sets the DofAperture property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#dofaperture
func (l *LensRenderingPipeline) SetDofAperture(dofAperture float64) *LensRenderingPipeline {
	p := ba.ctx.Get("LensRenderingPipeline").New(dofAperture)
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

// DofDistortion returns the DofDistortion property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#dofdistortion
func (l *LensRenderingPipeline) DofDistortion(dofDistortion float64) *LensRenderingPipeline {
	p := ba.ctx.Get("LensRenderingPipeline").New(dofDistortion)
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

// SetDofDistortion sets the DofDistortion property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#dofdistortion
func (l *LensRenderingPipeline) SetDofDistortion(dofDistortion float64) *LensRenderingPipeline {
	p := ba.ctx.Get("LensRenderingPipeline").New(dofDistortion)
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

// EdgeBlur returns the EdgeBlur property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#edgeblur
func (l *LensRenderingPipeline) EdgeBlur(edgeBlur float64) *LensRenderingPipeline {
	p := ba.ctx.Get("LensRenderingPipeline").New(edgeBlur)
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

// SetEdgeBlur sets the EdgeBlur property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#edgeblur
func (l *LensRenderingPipeline) SetEdgeBlur(edgeBlur float64) *LensRenderingPipeline {
	p := ba.ctx.Get("LensRenderingPipeline").New(edgeBlur)
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

// EdgeDistortion returns the EdgeDistortion property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#edgedistortion
func (l *LensRenderingPipeline) EdgeDistortion(edgeDistortion float64) *LensRenderingPipeline {
	p := ba.ctx.Get("LensRenderingPipeline").New(edgeDistortion)
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

// SetEdgeDistortion sets the EdgeDistortion property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#edgedistortion
func (l *LensRenderingPipeline) SetEdgeDistortion(edgeDistortion float64) *LensRenderingPipeline {
	p := ba.ctx.Get("LensRenderingPipeline").New(edgeDistortion)
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

// GrainAmount returns the GrainAmount property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#grainamount
func (l *LensRenderingPipeline) GrainAmount(grainAmount float64) *LensRenderingPipeline {
	p := ba.ctx.Get("LensRenderingPipeline").New(grainAmount)
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

// SetGrainAmount sets the GrainAmount property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#grainamount
func (l *LensRenderingPipeline) SetGrainAmount(grainAmount float64) *LensRenderingPipeline {
	p := ba.ctx.Get("LensRenderingPipeline").New(grainAmount)
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

// HighlightsGain returns the HighlightsGain property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#highlightsgain
func (l *LensRenderingPipeline) HighlightsGain(highlightsGain float64) *LensRenderingPipeline {
	p := ba.ctx.Get("LensRenderingPipeline").New(highlightsGain)
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

// SetHighlightsGain sets the HighlightsGain property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#highlightsgain
func (l *LensRenderingPipeline) SetHighlightsGain(highlightsGain float64) *LensRenderingPipeline {
	p := ba.ctx.Get("LensRenderingPipeline").New(highlightsGain)
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

// HighlightsThreshold returns the HighlightsThreshold property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#highlightsthreshold
func (l *LensRenderingPipeline) HighlightsThreshold(highlightsThreshold float64) *LensRenderingPipeline {
	p := ba.ctx.Get("LensRenderingPipeline").New(highlightsThreshold)
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

// SetHighlightsThreshold sets the HighlightsThreshold property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#highlightsthreshold
func (l *LensRenderingPipeline) SetHighlightsThreshold(highlightsThreshold float64) *LensRenderingPipeline {
	p := ba.ctx.Get("LensRenderingPipeline").New(highlightsThreshold)
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

// InspectableCustomProperties returns the InspectableCustomProperties property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#inspectablecustomproperties
func (l *LensRenderingPipeline) InspectableCustomProperties(inspectableCustomProperties *IInspectable) *LensRenderingPipeline {
	p := ba.ctx.Get("LensRenderingPipeline").New(inspectableCustomProperties.JSObject())
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

// SetInspectableCustomProperties sets the InspectableCustomProperties property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#inspectablecustomproperties
func (l *LensRenderingPipeline) SetInspectableCustomProperties(inspectableCustomProperties *IInspectable) *LensRenderingPipeline {
	p := ba.ctx.Get("LensRenderingPipeline").New(inspectableCustomProperties.JSObject())
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

// IsSupported returns the IsSupported property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#issupported
func (l *LensRenderingPipeline) IsSupported(isSupported bool) *LensRenderingPipeline {
	p := ba.ctx.Get("LensRenderingPipeline").New(isSupported)
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

// SetIsSupported sets the IsSupported property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#issupported
func (l *LensRenderingPipeline) SetIsSupported(isSupported bool) *LensRenderingPipeline {
	p := ba.ctx.Get("LensRenderingPipeline").New(isSupported)
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#name
func (l *LensRenderingPipeline) Name(name string) *LensRenderingPipeline {
	p := ba.ctx.Get("LensRenderingPipeline").New(name)
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#name
func (l *LensRenderingPipeline) SetName(name string) *LensRenderingPipeline {
	p := ba.ctx.Get("LensRenderingPipeline").New(name)
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

// PentagonBokeh returns the PentagonBokeh property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#pentagonbokeh
func (l *LensRenderingPipeline) PentagonBokeh(pentagonBokeh bool) *LensRenderingPipeline {
	p := ba.ctx.Get("LensRenderingPipeline").New(pentagonBokeh)
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

// SetPentagonBokeh sets the PentagonBokeh property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#pentagonbokeh
func (l *LensRenderingPipeline) SetPentagonBokeh(pentagonBokeh bool) *LensRenderingPipeline {
	p := ba.ctx.Get("LensRenderingPipeline").New(pentagonBokeh)
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

// Scene returns the Scene property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#scene
func (l *LensRenderingPipeline) Scene(scene *Scene) *LensRenderingPipeline {
	p := ba.ctx.Get("LensRenderingPipeline").New(scene.JSObject())
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

// SetScene sets the Scene property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#scene
func (l *LensRenderingPipeline) SetScene(scene *Scene) *LensRenderingPipeline {
	p := ba.ctx.Get("LensRenderingPipeline").New(scene.JSObject())
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
}

*/
