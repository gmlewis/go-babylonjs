// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AsciiArtFontTexture represents a babylon.js AsciiArtFontTexture.
// AsciiArtFontTexture is the helper class used to easily create your ascii art font texture.
//
// It basically takes care rendering the font front the given font size to a texture.
// This is used later on in the postprocess.
type AsciiArtFontTexture struct {
	*BaseTexture
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *AsciiArtFontTexture) JSObject() js.Value { return a.p }

// AsciiArtFontTexture returns a AsciiArtFontTexture JavaScript class.
func (ba *Babylon) AsciiArtFontTexture() *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture")
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// AsciiArtFontTextureFromJSObject returns a wrapped AsciiArtFontTexture JavaScript class.
func AsciiArtFontTextureFromJSObject(p js.Value, ctx js.Value) *AsciiArtFontTexture {
	return &AsciiArtFontTexture{BaseTexture: BaseTextureFromJSObject(p, ctx), ctx: ctx}
}

// AsciiArtFontTextureArrayToJSArray returns a JavaScript Array for the wrapped array.
func AsciiArtFontTextureArrayToJSArray(array []*AsciiArtFontTexture) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewAsciiArtFontTextureOpts contains optional parameters for NewAsciiArtFontTexture.
type NewAsciiArtFontTextureOpts struct {
	Scene *Scene
}

// NewAsciiArtFontTexture returns a new AsciiArtFontTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture
func (ba *Babylon) NewAsciiArtFontTexture(name string, font string, text string, opts *NewAsciiArtFontTextureOpts) *AsciiArtFontTexture {
	if opts == nil {
		opts = &NewAsciiArtFontTextureOpts{}
	}

	args := make([]interface{}, 0, 3+1)

	args = append(args, name)
	args = append(args, font)
	args = append(args, text)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	p := ba.ctx.Get("AsciiArtFontTexture").New(args...)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// Clone calls the Clone method on the AsciiArtFontTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#clone
func (a *AsciiArtFontTexture) Clone() *AsciiArtFontTexture {

	retVal := a.p.Call("clone")
	return AsciiArtFontTextureFromJSObject(retVal, a.ctx)
}

// DelayLoad calls the DelayLoad method on the AsciiArtFontTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#delayload
func (a *AsciiArtFontTexture) DelayLoad() {

	a.p.Call("delayLoad")
}

// Dispose calls the Dispose method on the AsciiArtFontTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#dispose
func (a *AsciiArtFontTexture) Dispose() {

	a.p.Call("dispose")
}

// GetBaseSize calls the GetBaseSize method on the AsciiArtFontTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#getbasesize
func (a *AsciiArtFontTexture) GetBaseSize() js.Value {

	retVal := a.p.Call("getBaseSize")
	return retVal
}

// GetClassName calls the GetClassName method on the AsciiArtFontTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#getclassname
func (a *AsciiArtFontTexture) GetClassName() string {

	retVal := a.p.Call("getClassName")
	return retVal.String()
}

// GetInternalTexture calls the GetInternalTexture method on the AsciiArtFontTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#getinternaltexture
func (a *AsciiArtFontTexture) GetInternalTexture() *InternalTexture {

	retVal := a.p.Call("getInternalTexture")
	return InternalTextureFromJSObject(retVal, a.ctx)
}

// GetReflectionTextureMatrix calls the GetReflectionTextureMatrix method on the AsciiArtFontTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#getreflectiontexturematrix
func (a *AsciiArtFontTexture) GetReflectionTextureMatrix() *Matrix {

	retVal := a.p.Call("getReflectionTextureMatrix")
	return MatrixFromJSObject(retVal, a.ctx)
}

// GetScene calls the GetScene method on the AsciiArtFontTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#getscene
func (a *AsciiArtFontTexture) GetScene() *Scene {

	retVal := a.p.Call("getScene")
	return SceneFromJSObject(retVal, a.ctx)
}

// GetSize calls the GetSize method on the AsciiArtFontTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#getsize
func (a *AsciiArtFontTexture) GetSize() js.Value {

	retVal := a.p.Call("getSize")
	return retVal
}

// GetTextureMatrix calls the GetTextureMatrix method on the AsciiArtFontTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#gettexturematrix
func (a *AsciiArtFontTexture) GetTextureMatrix() *Matrix {

	retVal := a.p.Call("getTextureMatrix")
	return MatrixFromJSObject(retVal, a.ctx)
}

// IsReady calls the IsReady method on the AsciiArtFontTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#isready
func (a *AsciiArtFontTexture) IsReady() bool {

	retVal := a.p.Call("isReady")
	return retVal.Bool()
}

// IsReadyOrNotBlocking calls the IsReadyOrNotBlocking method on the AsciiArtFontTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#isreadyornotblocking
func (a *AsciiArtFontTexture) IsReadyOrNotBlocking() bool {

	retVal := a.p.Call("isReadyOrNotBlocking")
	return retVal.Bool()
}

// Parse calls the Parse method on the AsciiArtFontTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#parse
func (a *AsciiArtFontTexture) Parse(source interface{}, scene *Scene) *AsciiArtFontTexture {

	args := make([]interface{}, 0, 2+0)

	args = append(args, source)
	args = append(args, scene.JSObject())

	retVal := a.p.Call("Parse", args...)
	return AsciiArtFontTextureFromJSObject(retVal, a.ctx)
}

// AsciiArtFontTextureReadPixelsOpts contains optional parameters for AsciiArtFontTexture.ReadPixels.
type AsciiArtFontTextureReadPixelsOpts struct {
	FaceIndex *float64
	Level     *float64
	Buffer    js.Value
}

// ReadPixels calls the ReadPixels method on the AsciiArtFontTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#readpixels
func (a *AsciiArtFontTexture) ReadPixels(opts *AsciiArtFontTextureReadPixelsOpts) js.Value {
	if opts == nil {
		opts = &AsciiArtFontTextureReadPixelsOpts{}
	}

	args := make([]interface{}, 0, 0+3)

	if opts.FaceIndex == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.FaceIndex)
	}
	if opts.Level == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Level)
	}
	if opts.Buffer == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Buffer)
	}

	retVal := a.p.Call("readPixels", args...)
	return retVal
}

// ReleaseInternalTexture calls the ReleaseInternalTexture method on the AsciiArtFontTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#releaseinternaltexture
func (a *AsciiArtFontTexture) ReleaseInternalTexture() {

	a.p.Call("releaseInternalTexture")
}

// Scale calls the Scale method on the AsciiArtFontTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#scale
func (a *AsciiArtFontTexture) Scale(ratio float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, ratio)

	a.p.Call("scale", args...)
}

// Serialize calls the Serialize method on the AsciiArtFontTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#serialize
func (a *AsciiArtFontTexture) Serialize() interface{} {

	retVal := a.p.Call("serialize")
	return retVal
}

// ToString calls the ToString method on the AsciiArtFontTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#tostring
func (a *AsciiArtFontTexture) ToString() string {

	retVal := a.p.Call("toString")
	return retVal.String()
}

// UpdateSamplingMode calls the UpdateSamplingMode method on the AsciiArtFontTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#updatesamplingmode
func (a *AsciiArtFontTexture) UpdateSamplingMode(samplingMode float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, samplingMode)

	a.p.Call("updateSamplingMode", args...)
}

// WhenAllReady calls the WhenAllReady method on the AsciiArtFontTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#whenallready
func (a *AsciiArtFontTexture) WhenAllReady(textures *BaseTexture, callback func()) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, textures.JSObject())
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { callback(); return nil }))

	a.p.Call("WhenAllReady", args...)
}

/*

// Animations returns the Animations property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#animations
func (a *AsciiArtFontTexture) Animations(animations *Animation) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(animations.JSObject())
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetAnimations sets the Animations property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#animations
func (a *AsciiArtFontTexture) SetAnimations(animations *Animation) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(animations.JSObject())
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// AnisotropicFilteringLevel returns the AnisotropicFilteringLevel property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#anisotropicfilteringlevel
func (a *AsciiArtFontTexture) AnisotropicFilteringLevel(anisotropicFilteringLevel float64) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(anisotropicFilteringLevel)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetAnisotropicFilteringLevel sets the AnisotropicFilteringLevel property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#anisotropicfilteringlevel
func (a *AsciiArtFontTexture) SetAnisotropicFilteringLevel(anisotropicFilteringLevel float64) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(anisotropicFilteringLevel)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// CanRescale returns the CanRescale property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#canrescale
func (a *AsciiArtFontTexture) CanRescale(canRescale bool) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(canRescale)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetCanRescale sets the CanRescale property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#canrescale
func (a *AsciiArtFontTexture) SetCanRescale(canRescale bool) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(canRescale)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// CharSize returns the CharSize property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#charsize
func (a *AsciiArtFontTexture) CharSize(charSize float64) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(charSize)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetCharSize sets the CharSize property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#charsize
func (a *AsciiArtFontTexture) SetCharSize(charSize float64) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(charSize)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// CoordinatesIndex returns the CoordinatesIndex property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#coordinatesindex
func (a *AsciiArtFontTexture) CoordinatesIndex(coordinatesIndex float64) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(coordinatesIndex)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetCoordinatesIndex sets the CoordinatesIndex property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#coordinatesindex
func (a *AsciiArtFontTexture) SetCoordinatesIndex(coordinatesIndex float64) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(coordinatesIndex)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// CoordinatesMode returns the CoordinatesMode property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#coordinatesmode
func (a *AsciiArtFontTexture) CoordinatesMode(coordinatesMode float64) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(coordinatesMode)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetCoordinatesMode sets the CoordinatesMode property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#coordinatesmode
func (a *AsciiArtFontTexture) SetCoordinatesMode(coordinatesMode float64) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(coordinatesMode)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// DEFAULT_ANISOTROPIC_FILTERING_LEVEL returns the DEFAULT_ANISOTROPIC_FILTERING_LEVEL property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#default_anisotropic_filtering_level
func (a *AsciiArtFontTexture) DEFAULT_ANISOTROPIC_FILTERING_LEVEL(DEFAULT_ANISOTROPIC_FILTERING_LEVEL float64) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(DEFAULT_ANISOTROPIC_FILTERING_LEVEL)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetDEFAULT_ANISOTROPIC_FILTERING_LEVEL sets the DEFAULT_ANISOTROPIC_FILTERING_LEVEL property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#default_anisotropic_filtering_level
func (a *AsciiArtFontTexture) SetDEFAULT_ANISOTROPIC_FILTERING_LEVEL(DEFAULT_ANISOTROPIC_FILTERING_LEVEL float64) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(DEFAULT_ANISOTROPIC_FILTERING_LEVEL)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// DelayLoadState returns the DelayLoadState property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#delayloadstate
func (a *AsciiArtFontTexture) DelayLoadState(delayLoadState float64) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(delayLoadState)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetDelayLoadState sets the DelayLoadState property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#delayloadstate
func (a *AsciiArtFontTexture) SetDelayLoadState(delayLoadState float64) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(delayLoadState)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// GammaSpace returns the GammaSpace property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#gammaspace
func (a *AsciiArtFontTexture) GammaSpace(gammaSpace bool) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(gammaSpace)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetGammaSpace sets the GammaSpace property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#gammaspace
func (a *AsciiArtFontTexture) SetGammaSpace(gammaSpace bool) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(gammaSpace)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// GetAlphaFromRGB returns the GetAlphaFromRGB property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#getalphafromrgb
func (a *AsciiArtFontTexture) GetAlphaFromRGB(getAlphaFromRGB bool) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(getAlphaFromRGB)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetGetAlphaFromRGB sets the GetAlphaFromRGB property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#getalphafromrgb
func (a *AsciiArtFontTexture) SetGetAlphaFromRGB(getAlphaFromRGB bool) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(getAlphaFromRGB)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// HasAlpha returns the HasAlpha property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#hasalpha
func (a *AsciiArtFontTexture) HasAlpha(hasAlpha bool) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(hasAlpha)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetHasAlpha sets the HasAlpha property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#hasalpha
func (a *AsciiArtFontTexture) SetHasAlpha(hasAlpha bool) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(hasAlpha)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// InvertZ returns the InvertZ property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#invertz
func (a *AsciiArtFontTexture) InvertZ(invertZ bool) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(invertZ)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetInvertZ sets the InvertZ property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#invertz
func (a *AsciiArtFontTexture) SetInvertZ(invertZ bool) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(invertZ)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// IrradianceTexture returns the IrradianceTexture property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#irradiancetexture
func (a *AsciiArtFontTexture) IrradianceTexture(irradianceTexture *BaseTexture) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(irradianceTexture.JSObject())
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetIrradianceTexture sets the IrradianceTexture property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#irradiancetexture
func (a *AsciiArtFontTexture) SetIrradianceTexture(irradianceTexture *BaseTexture) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(irradianceTexture.JSObject())
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// Is2DArray returns the Is2DArray property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#is2darray
func (a *AsciiArtFontTexture) Is2DArray(is2DArray bool) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(is2DArray)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetIs2DArray sets the Is2DArray property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#is2darray
func (a *AsciiArtFontTexture) SetIs2DArray(is2DArray bool) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(is2DArray)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// Is3D returns the Is3D property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#is3d
func (a *AsciiArtFontTexture) Is3D(is3D bool) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(is3D)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetIs3D sets the Is3D property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#is3d
func (a *AsciiArtFontTexture) SetIs3D(is3D bool) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(is3D)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// IsBlocking returns the IsBlocking property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#isblocking
func (a *AsciiArtFontTexture) IsBlocking(isBlocking bool) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(isBlocking)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetIsBlocking sets the IsBlocking property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#isblocking
func (a *AsciiArtFontTexture) SetIsBlocking(isBlocking bool) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(isBlocking)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// IsCube returns the IsCube property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#iscube
func (a *AsciiArtFontTexture) IsCube(isCube bool) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(isCube)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetIsCube sets the IsCube property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#iscube
func (a *AsciiArtFontTexture) SetIsCube(isCube bool) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(isCube)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// IsRGBD returns the IsRGBD property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#isrgbd
func (a *AsciiArtFontTexture) IsRGBD(isRGBD bool) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(isRGBD)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetIsRGBD sets the IsRGBD property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#isrgbd
func (a *AsciiArtFontTexture) SetIsRGBD(isRGBD bool) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(isRGBD)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// IsRenderTarget returns the IsRenderTarget property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#isrendertarget
func (a *AsciiArtFontTexture) IsRenderTarget(isRenderTarget bool) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(isRenderTarget)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetIsRenderTarget sets the IsRenderTarget property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#isrendertarget
func (a *AsciiArtFontTexture) SetIsRenderTarget(isRenderTarget bool) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(isRenderTarget)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// Level returns the Level property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#level
func (a *AsciiArtFontTexture) Level(level float64) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(level)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetLevel sets the Level property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#level
func (a *AsciiArtFontTexture) SetLevel(level float64) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(level)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// LinearSpecularLOD returns the LinearSpecularLOD property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#linearspecularlod
func (a *AsciiArtFontTexture) LinearSpecularLOD(linearSpecularLOD bool) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(linearSpecularLOD)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetLinearSpecularLOD sets the LinearSpecularLOD property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#linearspecularlod
func (a *AsciiArtFontTexture) SetLinearSpecularLOD(linearSpecularLOD bool) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(linearSpecularLOD)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// LodGenerationOffset returns the LodGenerationOffset property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#lodgenerationoffset
func (a *AsciiArtFontTexture) LodGenerationOffset(lodGenerationOffset float64) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(lodGenerationOffset)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetLodGenerationOffset sets the LodGenerationOffset property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#lodgenerationoffset
func (a *AsciiArtFontTexture) SetLodGenerationOffset(lodGenerationOffset float64) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(lodGenerationOffset)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// LodGenerationScale returns the LodGenerationScale property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#lodgenerationscale
func (a *AsciiArtFontTexture) LodGenerationScale(lodGenerationScale float64) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(lodGenerationScale)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetLodGenerationScale sets the LodGenerationScale property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#lodgenerationscale
func (a *AsciiArtFontTexture) SetLodGenerationScale(lodGenerationScale float64) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(lodGenerationScale)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// Metadata returns the Metadata property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#metadata
func (a *AsciiArtFontTexture) Metadata(metadata interface{}) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(metadata)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetMetadata sets the Metadata property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#metadata
func (a *AsciiArtFontTexture) SetMetadata(metadata interface{}) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(metadata)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#name
func (a *AsciiArtFontTexture) Name(name string) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(name)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#name
func (a *AsciiArtFontTexture) SetName(name string) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(name)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// NoMipmap returns the NoMipmap property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#nomipmap
func (a *AsciiArtFontTexture) NoMipmap(noMipmap bool) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(noMipmap)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetNoMipmap sets the NoMipmap property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#nomipmap
func (a *AsciiArtFontTexture) SetNoMipmap(noMipmap bool) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(noMipmap)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// OnDispose returns the OnDispose property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#ondispose
func (a *AsciiArtFontTexture) OnDispose(onDispose func()) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onDispose(); return nil}))
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetOnDispose sets the OnDispose property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#ondispose
func (a *AsciiArtFontTexture) SetOnDispose(onDispose func()) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onDispose(); return nil}))
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// OnDisposeObservable returns the OnDisposeObservable property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#ondisposeobservable
func (a *AsciiArtFontTexture) OnDisposeObservable(onDisposeObservable *Observable) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(onDisposeObservable.JSObject())
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetOnDisposeObservable sets the OnDisposeObservable property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#ondisposeobservable
func (a *AsciiArtFontTexture) SetOnDisposeObservable(onDisposeObservable *Observable) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(onDisposeObservable.JSObject())
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// ReservedDataStore returns the ReservedDataStore property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#reserveddatastore
func (a *AsciiArtFontTexture) ReservedDataStore(reservedDataStore interface{}) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(reservedDataStore)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetReservedDataStore sets the ReservedDataStore property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#reserveddatastore
func (a *AsciiArtFontTexture) SetReservedDataStore(reservedDataStore interface{}) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(reservedDataStore)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SphericalPolynomial returns the SphericalPolynomial property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#sphericalpolynomial
func (a *AsciiArtFontTexture) SphericalPolynomial(sphericalPolynomial *SphericalPolynomial) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(sphericalPolynomial.JSObject())
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetSphericalPolynomial sets the SphericalPolynomial property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#sphericalpolynomial
func (a *AsciiArtFontTexture) SetSphericalPolynomial(sphericalPolynomial *SphericalPolynomial) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(sphericalPolynomial.JSObject())
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// TextureFormat returns the TextureFormat property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#textureformat
func (a *AsciiArtFontTexture) TextureFormat(textureFormat float64) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(textureFormat)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetTextureFormat sets the TextureFormat property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#textureformat
func (a *AsciiArtFontTexture) SetTextureFormat(textureFormat float64) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(textureFormat)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// TextureType returns the TextureType property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#texturetype
func (a *AsciiArtFontTexture) TextureType(textureType float64) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(textureType)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetTextureType sets the TextureType property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#texturetype
func (a *AsciiArtFontTexture) SetTextureType(textureType float64) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(textureType)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// Uid returns the Uid property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#uid
func (a *AsciiArtFontTexture) Uid(uid string) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(uid)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetUid sets the Uid property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#uid
func (a *AsciiArtFontTexture) SetUid(uid string) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(uid)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// UniqueId returns the UniqueId property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#uniqueid
func (a *AsciiArtFontTexture) UniqueId(uniqueId float64) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(uniqueId)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetUniqueId sets the UniqueId property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#uniqueid
func (a *AsciiArtFontTexture) SetUniqueId(uniqueId float64) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(uniqueId)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// WrapR returns the WrapR property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#wrapr
func (a *AsciiArtFontTexture) WrapR(wrapR float64) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(wrapR)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetWrapR sets the WrapR property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#wrapr
func (a *AsciiArtFontTexture) SetWrapR(wrapR float64) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(wrapR)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// WrapU returns the WrapU property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#wrapu
func (a *AsciiArtFontTexture) WrapU(wrapU float64) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(wrapU)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetWrapU sets the WrapU property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#wrapu
func (a *AsciiArtFontTexture) SetWrapU(wrapU float64) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(wrapU)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// WrapV returns the WrapV property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#wrapv
func (a *AsciiArtFontTexture) WrapV(wrapV float64) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(wrapV)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// SetWrapV sets the WrapV property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#wrapv
func (a *AsciiArtFontTexture) SetWrapV(wrapV float64) *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture").New(wrapV)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

*/
