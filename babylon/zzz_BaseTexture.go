// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// BaseTexture represents a babylon.js BaseTexture.
// Base class of all the textures in babylon.
// It groups all the common properties the materials, post process, lights... might need
// in order to make a correct use of the texture.
type BaseTexture struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (b *BaseTexture) JSObject() js.Value { return b.p }

// BaseTexture returns a BaseTexture JavaScript class.
func (ba *Babylon) BaseTexture() *BaseTexture {
	p := ba.ctx.Get("BaseTexture")
	return BaseTextureFromJSObject(p, ba.ctx)
}

// BaseTextureFromJSObject returns a wrapped BaseTexture JavaScript class.
func BaseTextureFromJSObject(p js.Value, ctx js.Value) *BaseTexture {
	return &BaseTexture{p: p, ctx: ctx}
}

// BaseTextureArrayToJSArray returns a JavaScript Array for the wrapped array.
func BaseTextureArrayToJSArray(array []*BaseTexture) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewBaseTexture returns a new BaseTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture
func (ba *Babylon) NewBaseTexture(scene *Scene) *BaseTexture {

	args := make([]interface{}, 0, 1+0)

	args = append(args, scene.JSObject())

	p := ba.ctx.Get("BaseTexture").New(args...)
	return BaseTextureFromJSObject(p, ba.ctx)
}

// Clone calls the Clone method on the BaseTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#clone
func (b *BaseTexture) Clone() *BaseTexture {

	retVal := b.p.Call("clone")
	return BaseTextureFromJSObject(retVal, b.ctx)
}

// DelayLoad calls the DelayLoad method on the BaseTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#delayload
func (b *BaseTexture) DelayLoad() {

	b.p.Call("delayLoad")
}

// Dispose calls the Dispose method on the BaseTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#dispose
func (b *BaseTexture) Dispose() {

	b.p.Call("dispose")
}

// GetBaseSize calls the GetBaseSize method on the BaseTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#getbasesize
func (b *BaseTexture) GetBaseSize() *ISize {

	retVal := b.p.Call("getBaseSize")
	return ISizeFromJSObject(retVal, b.ctx)
}

// GetClassName calls the GetClassName method on the BaseTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#getclassname
func (b *BaseTexture) GetClassName() string {

	retVal := b.p.Call("getClassName")
	return retVal.String()
}

// GetInternalTexture calls the GetInternalTexture method on the BaseTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#getinternaltexture
func (b *BaseTexture) GetInternalTexture() *InternalTexture {

	retVal := b.p.Call("getInternalTexture")
	return InternalTextureFromJSObject(retVal, b.ctx)
}

// GetReflectionTextureMatrix calls the GetReflectionTextureMatrix method on the BaseTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#getreflectiontexturematrix
func (b *BaseTexture) GetReflectionTextureMatrix() *Matrix {

	retVal := b.p.Call("getReflectionTextureMatrix")
	return MatrixFromJSObject(retVal, b.ctx)
}

// GetScene calls the GetScene method on the BaseTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#getscene
func (b *BaseTexture) GetScene() *Scene {

	retVal := b.p.Call("getScene")
	return SceneFromJSObject(retVal, b.ctx)
}

// GetSize calls the GetSize method on the BaseTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#getsize
func (b *BaseTexture) GetSize() *ISize {

	retVal := b.p.Call("getSize")
	return ISizeFromJSObject(retVal, b.ctx)
}

// GetTextureMatrix calls the GetTextureMatrix method on the BaseTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#gettexturematrix
func (b *BaseTexture) GetTextureMatrix() *Matrix {

	retVal := b.p.Call("getTextureMatrix")
	return MatrixFromJSObject(retVal, b.ctx)
}

// IsReady calls the IsReady method on the BaseTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#isready
func (b *BaseTexture) IsReady() bool {

	retVal := b.p.Call("isReady")
	return retVal.Bool()
}

// IsReadyOrNotBlocking calls the IsReadyOrNotBlocking method on the BaseTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#isreadyornotblocking
func (b *BaseTexture) IsReadyOrNotBlocking() bool {

	retVal := b.p.Call("isReadyOrNotBlocking")
	return retVal.Bool()
}

// BaseTextureReadPixelsOpts contains optional parameters for BaseTexture.ReadPixels.
type BaseTextureReadPixelsOpts struct {
	FaceIndex *float64
	Level     *float64
	Buffer    js.Value
}

// ReadPixels calls the ReadPixels method on the BaseTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#readpixels
func (b *BaseTexture) ReadPixels(opts *BaseTextureReadPixelsOpts) js.Value {
	if opts == nil {
		opts = &BaseTextureReadPixelsOpts{}
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
	args = append(args, opts.Buffer)

	retVal := b.p.Call("readPixels", args...)
	return retVal
}

// ReleaseInternalTexture calls the ReleaseInternalTexture method on the BaseTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#releaseinternaltexture
func (b *BaseTexture) ReleaseInternalTexture() {

	b.p.Call("releaseInternalTexture")
}

// Scale calls the Scale method on the BaseTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#scale
func (b *BaseTexture) Scale(ratio float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, ratio)

	b.p.Call("scale", args...)
}

// Serialize calls the Serialize method on the BaseTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#serialize
func (b *BaseTexture) Serialize() interface{} {

	retVal := b.p.Call("serialize")
	return retVal
}

// ToString calls the ToString method on the BaseTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#tostring
func (b *BaseTexture) ToString() string {

	retVal := b.p.Call("toString")
	return retVal.String()
}

// UpdateSamplingMode calls the UpdateSamplingMode method on the BaseTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#updatesamplingmode
func (b *BaseTexture) UpdateSamplingMode(samplingMode float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, samplingMode)

	b.p.Call("updateSamplingMode", args...)
}

// WhenAllReady calls the WhenAllReady method on the BaseTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#whenallready
func (b *BaseTexture) WhenAllReady(textures []*BaseTexture, callback func()) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, BaseTextureArrayToJSArray(textures))
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { callback(); return nil }))

	b.p.Call("WhenAllReady", args...)
}

// Animations returns the Animations property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#animations
func (b *BaseTexture) Animations() []*Animation {
	retVal := b.p.Get("animations")
	result := []*Animation{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, AnimationFromJSObject(retVal.Index(ri), b.ctx))
	}
	return result
}

// SetAnimations sets the Animations property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#animations
func (b *BaseTexture) SetAnimations(animations []*Animation) *BaseTexture {
	b.p.Set("animations", animations)
	return b
}

// AnisotropicFilteringLevel returns the AnisotropicFilteringLevel property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#anisotropicfilteringlevel
func (b *BaseTexture) AnisotropicFilteringLevel() float64 {
	retVal := b.p.Get("anisotropicFilteringLevel")
	return retVal.Float()
}

// SetAnisotropicFilteringLevel sets the AnisotropicFilteringLevel property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#anisotropicfilteringlevel
func (b *BaseTexture) SetAnisotropicFilteringLevel(anisotropicFilteringLevel float64) *BaseTexture {
	b.p.Set("anisotropicFilteringLevel", anisotropicFilteringLevel)
	return b
}

// CanRescale returns the CanRescale property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#canrescale
func (b *BaseTexture) CanRescale() bool {
	retVal := b.p.Get("canRescale")
	return retVal.Bool()
}

// SetCanRescale sets the CanRescale property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#canrescale
func (b *BaseTexture) SetCanRescale(canRescale bool) *BaseTexture {
	b.p.Set("canRescale", canRescale)
	return b
}

// CoordinatesIndex returns the CoordinatesIndex property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#coordinatesindex
func (b *BaseTexture) CoordinatesIndex() float64 {
	retVal := b.p.Get("coordinatesIndex")
	return retVal.Float()
}

// SetCoordinatesIndex sets the CoordinatesIndex property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#coordinatesindex
func (b *BaseTexture) SetCoordinatesIndex(coordinatesIndex float64) *BaseTexture {
	b.p.Set("coordinatesIndex", coordinatesIndex)
	return b
}

// CoordinatesMode returns the CoordinatesMode property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#coordinatesmode
func (b *BaseTexture) CoordinatesMode() float64 {
	retVal := b.p.Get("coordinatesMode")
	return retVal.Float()
}

// SetCoordinatesMode sets the CoordinatesMode property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#coordinatesmode
func (b *BaseTexture) SetCoordinatesMode(coordinatesMode float64) *BaseTexture {
	b.p.Set("coordinatesMode", coordinatesMode)
	return b
}

// DEFAULT_ANISOTROPIC_FILTERING_LEVEL returns the DEFAULT_ANISOTROPIC_FILTERING_LEVEL property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#default_anisotropic_filtering_level
func (b *BaseTexture) DEFAULT_ANISOTROPIC_FILTERING_LEVEL() float64 {
	retVal := b.p.Get("DEFAULT_ANISOTROPIC_FILTERING_LEVEL")
	return retVal.Float()
}

// SetDEFAULT_ANISOTROPIC_FILTERING_LEVEL sets the DEFAULT_ANISOTROPIC_FILTERING_LEVEL property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#default_anisotropic_filtering_level
func (b *BaseTexture) SetDEFAULT_ANISOTROPIC_FILTERING_LEVEL(DEFAULT_ANISOTROPIC_FILTERING_LEVEL float64) *BaseTexture {
	b.p.Set("DEFAULT_ANISOTROPIC_FILTERING_LEVEL", DEFAULT_ANISOTROPIC_FILTERING_LEVEL)
	return b
}

// DelayLoadState returns the DelayLoadState property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#delayloadstate
func (b *BaseTexture) DelayLoadState() float64 {
	retVal := b.p.Get("delayLoadState")
	return retVal.Float()
}

// SetDelayLoadState sets the DelayLoadState property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#delayloadstate
func (b *BaseTexture) SetDelayLoadState(delayLoadState float64) *BaseTexture {
	b.p.Set("delayLoadState", delayLoadState)
	return b
}

// GammaSpace returns the GammaSpace property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#gammaspace
func (b *BaseTexture) GammaSpace() bool {
	retVal := b.p.Get("gammaSpace")
	return retVal.Bool()
}

// SetGammaSpace sets the GammaSpace property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#gammaspace
func (b *BaseTexture) SetGammaSpace(gammaSpace bool) *BaseTexture {
	b.p.Set("gammaSpace", gammaSpace)
	return b
}

// GetAlphaFromRGB returns the GetAlphaFromRGB property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#getalphafromrgb
func (b *BaseTexture) GetAlphaFromRGB() bool {
	retVal := b.p.Get("getAlphaFromRGB")
	return retVal.Bool()
}

// SetGetAlphaFromRGB sets the GetAlphaFromRGB property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#getalphafromrgb
func (b *BaseTexture) SetGetAlphaFromRGB(getAlphaFromRGB bool) *BaseTexture {
	b.p.Set("getAlphaFromRGB", getAlphaFromRGB)
	return b
}

// HasAlpha returns the HasAlpha property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#hasalpha
func (b *BaseTexture) HasAlpha() bool {
	retVal := b.p.Get("hasAlpha")
	return retVal.Bool()
}

// SetHasAlpha sets the HasAlpha property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#hasalpha
func (b *BaseTexture) SetHasAlpha(hasAlpha bool) *BaseTexture {
	b.p.Set("hasAlpha", hasAlpha)
	return b
}

// InvertZ returns the InvertZ property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#invertz
func (b *BaseTexture) InvertZ() bool {
	retVal := b.p.Get("invertZ")
	return retVal.Bool()
}

// SetInvertZ sets the InvertZ property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#invertz
func (b *BaseTexture) SetInvertZ(invertZ bool) *BaseTexture {
	b.p.Set("invertZ", invertZ)
	return b
}

// IrradianceTexture returns the IrradianceTexture property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#irradiancetexture
func (b *BaseTexture) IrradianceTexture() *BaseTexture {
	retVal := b.p.Get("irradianceTexture")
	return BaseTextureFromJSObject(retVal, b.ctx)
}

// SetIrradianceTexture sets the IrradianceTexture property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#irradiancetexture
func (b *BaseTexture) SetIrradianceTexture(irradianceTexture *BaseTexture) *BaseTexture {
	b.p.Set("irradianceTexture", irradianceTexture.JSObject())
	return b
}

// Is2DArray returns the Is2DArray property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#is2darray
func (b *BaseTexture) Is2DArray() bool {
	retVal := b.p.Get("is2DArray")
	return retVal.Bool()
}

// SetIs2DArray sets the Is2DArray property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#is2darray
func (b *BaseTexture) SetIs2DArray(is2DArray bool) *BaseTexture {
	b.p.Set("is2DArray", is2DArray)
	return b
}

// Is3D returns the Is3D property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#is3d
func (b *BaseTexture) Is3D() bool {
	retVal := b.p.Get("is3D")
	return retVal.Bool()
}

// SetIs3D sets the Is3D property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#is3d
func (b *BaseTexture) SetIs3D(is3D bool) *BaseTexture {
	b.p.Set("is3D", is3D)
	return b
}

// IsBlocking returns the IsBlocking property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#isblocking
func (b *BaseTexture) IsBlocking() bool {
	retVal := b.p.Get("isBlocking")
	return retVal.Bool()
}

// SetIsBlocking sets the IsBlocking property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#isblocking
func (b *BaseTexture) SetIsBlocking(isBlocking bool) *BaseTexture {
	b.p.Set("isBlocking", isBlocking)
	return b
}

// IsCube returns the IsCube property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#iscube
func (b *BaseTexture) IsCube() bool {
	retVal := b.p.Get("isCube")
	return retVal.Bool()
}

// SetIsCube sets the IsCube property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#iscube
func (b *BaseTexture) SetIsCube(isCube bool) *BaseTexture {
	b.p.Set("isCube", isCube)
	return b
}

// IsRGBD returns the IsRGBD property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#isrgbd
func (b *BaseTexture) IsRGBD() bool {
	retVal := b.p.Get("isRGBD")
	return retVal.Bool()
}

// SetIsRGBD sets the IsRGBD property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#isrgbd
func (b *BaseTexture) SetIsRGBD(isRGBD bool) *BaseTexture {
	b.p.Set("isRGBD", isRGBD)
	return b
}

// IsRenderTarget returns the IsRenderTarget property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#isrendertarget
func (b *BaseTexture) IsRenderTarget() bool {
	retVal := b.p.Get("isRenderTarget")
	return retVal.Bool()
}

// SetIsRenderTarget sets the IsRenderTarget property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#isrendertarget
func (b *BaseTexture) SetIsRenderTarget(isRenderTarget bool) *BaseTexture {
	b.p.Set("isRenderTarget", isRenderTarget)
	return b
}

// Level returns the Level property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#level
func (b *BaseTexture) Level() float64 {
	retVal := b.p.Get("level")
	return retVal.Float()
}

// SetLevel sets the Level property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#level
func (b *BaseTexture) SetLevel(level float64) *BaseTexture {
	b.p.Set("level", level)
	return b
}

// LinearSpecularLOD returns the LinearSpecularLOD property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#linearspecularlod
func (b *BaseTexture) LinearSpecularLOD() bool {
	retVal := b.p.Get("linearSpecularLOD")
	return retVal.Bool()
}

// SetLinearSpecularLOD sets the LinearSpecularLOD property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#linearspecularlod
func (b *BaseTexture) SetLinearSpecularLOD(linearSpecularLOD bool) *BaseTexture {
	b.p.Set("linearSpecularLOD", linearSpecularLOD)
	return b
}

// LodGenerationOffset returns the LodGenerationOffset property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#lodgenerationoffset
func (b *BaseTexture) LodGenerationOffset() float64 {
	retVal := b.p.Get("lodGenerationOffset")
	return retVal.Float()
}

// SetLodGenerationOffset sets the LodGenerationOffset property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#lodgenerationoffset
func (b *BaseTexture) SetLodGenerationOffset(lodGenerationOffset float64) *BaseTexture {
	b.p.Set("lodGenerationOffset", lodGenerationOffset)
	return b
}

// LodGenerationScale returns the LodGenerationScale property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#lodgenerationscale
func (b *BaseTexture) LodGenerationScale() float64 {
	retVal := b.p.Get("lodGenerationScale")
	return retVal.Float()
}

// SetLodGenerationScale sets the LodGenerationScale property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#lodgenerationscale
func (b *BaseTexture) SetLodGenerationScale(lodGenerationScale float64) *BaseTexture {
	b.p.Set("lodGenerationScale", lodGenerationScale)
	return b
}

// Metadata returns the Metadata property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#metadata
func (b *BaseTexture) Metadata() interface{} {
	retVal := b.p.Get("metadata")
	return retVal
}

// SetMetadata sets the Metadata property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#metadata
func (b *BaseTexture) SetMetadata(metadata interface{}) *BaseTexture {
	b.p.Set("metadata", metadata)
	return b
}

// Name returns the Name property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#name
func (b *BaseTexture) Name() string {
	retVal := b.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#name
func (b *BaseTexture) SetName(name string) *BaseTexture {
	b.p.Set("name", name)
	return b
}

// NoMipmap returns the NoMipmap property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#nomipmap
func (b *BaseTexture) NoMipmap() bool {
	retVal := b.p.Get("noMipmap")
	return retVal.Bool()
}

// SetNoMipmap sets the NoMipmap property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#nomipmap
func (b *BaseTexture) SetNoMipmap(noMipmap bool) *BaseTexture {
	b.p.Set("noMipmap", noMipmap)
	return b
}

// OnDispose returns the OnDispose property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#ondispose
func (b *BaseTexture) OnDispose() js.Value {
	retVal := b.p.Get("onDispose")
	return retVal
}

// SetOnDispose sets the OnDispose property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#ondispose
func (b *BaseTexture) SetOnDispose(onDispose func()) *BaseTexture {
	b.p.Set("onDispose", js.FuncOf(func(this js.Value, args []js.Value) interface{} { onDispose(); return nil }))
	return b
}

// OnDisposeObservable returns the OnDisposeObservable property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#ondisposeobservable
func (b *BaseTexture) OnDisposeObservable() *Observable {
	retVal := b.p.Get("onDisposeObservable")
	return ObservableFromJSObject(retVal, b.ctx)
}

// SetOnDisposeObservable sets the OnDisposeObservable property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#ondisposeobservable
func (b *BaseTexture) SetOnDisposeObservable(onDisposeObservable *Observable) *BaseTexture {
	b.p.Set("onDisposeObservable", onDisposeObservable.JSObject())
	return b
}

// ReservedDataStore returns the ReservedDataStore property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#reserveddatastore
func (b *BaseTexture) ReservedDataStore() interface{} {
	retVal := b.p.Get("reservedDataStore")
	return retVal
}

// SetReservedDataStore sets the ReservedDataStore property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#reserveddatastore
func (b *BaseTexture) SetReservedDataStore(reservedDataStore interface{}) *BaseTexture {
	b.p.Set("reservedDataStore", reservedDataStore)
	return b
}

// SphericalPolynomial returns the SphericalPolynomial property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#sphericalpolynomial
func (b *BaseTexture) SphericalPolynomial() *SphericalPolynomial {
	retVal := b.p.Get("sphericalPolynomial")
	return SphericalPolynomialFromJSObject(retVal, b.ctx)
}

// SetSphericalPolynomial sets the SphericalPolynomial property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#sphericalpolynomial
func (b *BaseTexture) SetSphericalPolynomial(sphericalPolynomial *SphericalPolynomial) *BaseTexture {
	b.p.Set("sphericalPolynomial", sphericalPolynomial.JSObject())
	return b
}

// TextureFormat returns the TextureFormat property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#textureformat
func (b *BaseTexture) TextureFormat() float64 {
	retVal := b.p.Get("textureFormat")
	return retVal.Float()
}

// SetTextureFormat sets the TextureFormat property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#textureformat
func (b *BaseTexture) SetTextureFormat(textureFormat float64) *BaseTexture {
	b.p.Set("textureFormat", textureFormat)
	return b
}

// TextureType returns the TextureType property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#texturetype
func (b *BaseTexture) TextureType() float64 {
	retVal := b.p.Get("textureType")
	return retVal.Float()
}

// SetTextureType sets the TextureType property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#texturetype
func (b *BaseTexture) SetTextureType(textureType float64) *BaseTexture {
	b.p.Set("textureType", textureType)
	return b
}

// Uid returns the Uid property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#uid
func (b *BaseTexture) Uid() string {
	retVal := b.p.Get("uid")
	return retVal.String()
}

// SetUid sets the Uid property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#uid
func (b *BaseTexture) SetUid(uid string) *BaseTexture {
	b.p.Set("uid", uid)
	return b
}

// UniqueId returns the UniqueId property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#uniqueid
func (b *BaseTexture) UniqueId() float64 {
	retVal := b.p.Get("uniqueId")
	return retVal.Float()
}

// SetUniqueId sets the UniqueId property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#uniqueid
func (b *BaseTexture) SetUniqueId(uniqueId float64) *BaseTexture {
	b.p.Set("uniqueId", uniqueId)
	return b
}

// WrapR returns the WrapR property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#wrapr
func (b *BaseTexture) WrapR() float64 {
	retVal := b.p.Get("wrapR")
	return retVal.Float()
}

// SetWrapR sets the WrapR property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#wrapr
func (b *BaseTexture) SetWrapR(wrapR float64) *BaseTexture {
	b.p.Set("wrapR", wrapR)
	return b
}

// WrapU returns the WrapU property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#wrapu
func (b *BaseTexture) WrapU() float64 {
	retVal := b.p.Get("wrapU")
	return retVal.Float()
}

// SetWrapU sets the WrapU property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#wrapu
func (b *BaseTexture) SetWrapU(wrapU float64) *BaseTexture {
	b.p.Set("wrapU", wrapU)
	return b
}

// WrapV returns the WrapV property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#wrapv
func (b *BaseTexture) WrapV() float64 {
	retVal := b.p.Get("wrapV")
	return retVal.Float()
}

// SetWrapV sets the WrapV property of class BaseTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.basetexture#wrapv
func (b *BaseTexture) SetWrapV(wrapV float64) *BaseTexture {
	b.p.Set("wrapV", wrapV)
	return b
}
