// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// LensRenderingPipeline represents a babylon.js LensRenderingPipeline.
// BABYLON.JS Chromatic Aberration GLSL Shader
// Author: Olivier Guyot
// Separates very slightly R, G and B colors on the edges of the screen
// Inspired by Francois Tarlier &amp; Martins Upitis
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
	Cameras []*Camera
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
		args = append(args, CameraArrayToJSArray(opts.Cameras))
	}

	p := ba.ctx.Get("LensRenderingPipeline").New(args...)
	return LensRenderingPipelineFromJSObject(p, ba.ctx)
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

// BlurNoise returns the BlurNoise property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#blurnoise
func (l *LensRenderingPipeline) BlurNoise() bool {
	retVal := l.p.Get("blurNoise")
	return retVal.Bool()
}

// SetBlurNoise sets the BlurNoise property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#blurnoise
func (l *LensRenderingPipeline) SetBlurNoise(blurNoise bool) *LensRenderingPipeline {
	l.p.Set("blurNoise", blurNoise)
	return l
}

// ChromaticAberration returns the ChromaticAberration property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#chromaticaberration
func (l *LensRenderingPipeline) ChromaticAberration() float64 {
	retVal := l.p.Get("chromaticAberration")
	return retVal.Float()
}

// DarkenOutOfFocus returns the DarkenOutOfFocus property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#darkenoutoffocus
func (l *LensRenderingPipeline) DarkenOutOfFocus() float64 {
	retVal := l.p.Get("darkenOutOfFocus")
	return retVal.Float()
}

// DofAperture returns the DofAperture property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#dofaperture
func (l *LensRenderingPipeline) DofAperture() float64 {
	retVal := l.p.Get("dofAperture")
	return retVal.Float()
}

// SetDofAperture sets the DofAperture property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#dofaperture
func (l *LensRenderingPipeline) SetDofAperture(dofAperture float64) *LensRenderingPipeline {
	l.p.Set("dofAperture", dofAperture)
	return l
}

// DofDistortion returns the DofDistortion property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#dofdistortion
func (l *LensRenderingPipeline) DofDistortion() float64 {
	retVal := l.p.Get("dofDistortion")
	return retVal.Float()
}

// SetDofDistortion sets the DofDistortion property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#dofdistortion
func (l *LensRenderingPipeline) SetDofDistortion(dofDistortion float64) *LensRenderingPipeline {
	l.p.Set("dofDistortion", dofDistortion)
	return l
}

// EdgeBlur returns the EdgeBlur property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#edgeblur
func (l *LensRenderingPipeline) EdgeBlur() float64 {
	retVal := l.p.Get("edgeBlur")
	return retVal.Float()
}

// EdgeDistortion returns the EdgeDistortion property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#edgedistortion
func (l *LensRenderingPipeline) EdgeDistortion() float64 {
	retVal := l.p.Get("edgeDistortion")
	return retVal.Float()
}

// GrainAmount returns the GrainAmount property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#grainamount
func (l *LensRenderingPipeline) GrainAmount() float64 {
	retVal := l.p.Get("grainAmount")
	return retVal.Float()
}

// HighlightsGain returns the HighlightsGain property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#highlightsgain
func (l *LensRenderingPipeline) HighlightsGain() float64 {
	retVal := l.p.Get("highlightsGain")
	return retVal.Float()
}

// HighlightsThreshold returns the HighlightsThreshold property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#highlightsthreshold
func (l *LensRenderingPipeline) HighlightsThreshold() float64 {
	retVal := l.p.Get("highlightsThreshold")
	return retVal.Float()
}

// PentagonBokeh returns the PentagonBokeh property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#pentagonbokeh
func (l *LensRenderingPipeline) PentagonBokeh() bool {
	retVal := l.p.Get("pentagonBokeh")
	return retVal.Bool()
}

// SetPentagonBokeh sets the PentagonBokeh property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#pentagonbokeh
func (l *LensRenderingPipeline) SetPentagonBokeh(pentagonBokeh bool) *LensRenderingPipeline {
	l.p.Set("pentagonBokeh", pentagonBokeh)
	return l
}

// Scene returns the Scene property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#scene
func (l *LensRenderingPipeline) Scene() *Scene {
	retVal := l.p.Get("scene")
	return SceneFromJSObject(retVal, l.ctx)
}

// SetScene sets the Scene property of class LensRenderingPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.lensrenderingpipeline#scene
func (l *LensRenderingPipeline) SetScene(scene *Scene) *LensRenderingPipeline {
	l.p.Set("scene", scene.JSObject())
	return l
}
