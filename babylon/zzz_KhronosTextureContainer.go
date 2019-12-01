// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// KhronosTextureContainer represents a babylon.js KhronosTextureContainer.
// for description see &lt;a href=&#34;https://www.khronos.org/opengles/sdk/tools/KTX/&#34;&gt;https://www.khronos.org/opengles/sdk/tools/KTX/&lt;/a&gt;
// for file layout see &lt;a href=&#34;https://www.khronos.org/opengles/sdk/tools/KTX/file_format_spec/&#34;&gt;https://www.khronos.org/opengles/sdk/tools/KTX/file_format_spec/&lt;/a&gt;
type KhronosTextureContainer struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (k *KhronosTextureContainer) JSObject() js.Value { return k.p }

// KhronosTextureContainer returns a KhronosTextureContainer JavaScript class.
func (ba *Babylon) KhronosTextureContainer() *KhronosTextureContainer {
	p := ba.ctx.Get("KhronosTextureContainer")
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

// KhronosTextureContainerFromJSObject returns a wrapped KhronosTextureContainer JavaScript class.
func KhronosTextureContainerFromJSObject(p js.Value, ctx js.Value) *KhronosTextureContainer {
	return &KhronosTextureContainer{p: p, ctx: ctx}
}

// KhronosTextureContainerArrayToJSArray returns a JavaScript Array for the wrapped array.
func KhronosTextureContainerArrayToJSArray(array []*KhronosTextureContainer) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewKhronosTextureContainerOpts contains optional parameters for NewKhronosTextureContainer.
type NewKhronosTextureContainerOpts struct {
	ThreeDExpected       *bool
	TextureArrayExpected *bool
}

// NewKhronosTextureContainer returns a new KhronosTextureContainer object.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer
func (ba *Babylon) NewKhronosTextureContainer(arrayBuffer interface{}, facesExpected float64, opts *NewKhronosTextureContainerOpts) *KhronosTextureContainer {
	if opts == nil {
		opts = &NewKhronosTextureContainerOpts{}
	}

	args := make([]interface{}, 0, 2+2)

	args = append(args, arrayBuffer)
	args = append(args, facesExpected)

	if opts.ThreeDExpected == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ThreeDExpected)
	}
	if opts.TextureArrayExpected == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.TextureArrayExpected)
	}

	p := ba.ctx.Get("KhronosTextureContainer").New(args...)
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

/*

// ArrayBuffer returns the ArrayBuffer property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#arraybuffer
func (k *KhronosTextureContainer) ArrayBuffer(arrayBuffer interface{}) *KhronosTextureContainer {
	p := ba.ctx.Get("KhronosTextureContainer").New(arrayBuffer)
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

// SetArrayBuffer sets the ArrayBuffer property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#arraybuffer
func (k *KhronosTextureContainer) SetArrayBuffer(arrayBuffer interface{}) *KhronosTextureContainer {
	p := ba.ctx.Get("KhronosTextureContainer").New(arrayBuffer)
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

// BytesOfKeyValueData returns the BytesOfKeyValueData property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#bytesofkeyvaluedata
func (k *KhronosTextureContainer) BytesOfKeyValueData(bytesOfKeyValueData float64) *KhronosTextureContainer {
	p := ba.ctx.Get("KhronosTextureContainer").New(bytesOfKeyValueData)
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

// SetBytesOfKeyValueData sets the BytesOfKeyValueData property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#bytesofkeyvaluedata
func (k *KhronosTextureContainer) SetBytesOfKeyValueData(bytesOfKeyValueData float64) *KhronosTextureContainer {
	p := ba.ctx.Get("KhronosTextureContainer").New(bytesOfKeyValueData)
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

// GlBaseInternalFormat returns the GlBaseInternalFormat property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#glbaseinternalformat
func (k *KhronosTextureContainer) GlBaseInternalFormat(glBaseInternalFormat float64) *KhronosTextureContainer {
	p := ba.ctx.Get("KhronosTextureContainer").New(glBaseInternalFormat)
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

// SetGlBaseInternalFormat sets the GlBaseInternalFormat property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#glbaseinternalformat
func (k *KhronosTextureContainer) SetGlBaseInternalFormat(glBaseInternalFormat float64) *KhronosTextureContainer {
	p := ba.ctx.Get("KhronosTextureContainer").New(glBaseInternalFormat)
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

// GlFormat returns the GlFormat property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#glformat
func (k *KhronosTextureContainer) GlFormat(glFormat float64) *KhronosTextureContainer {
	p := ba.ctx.Get("KhronosTextureContainer").New(glFormat)
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

// SetGlFormat sets the GlFormat property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#glformat
func (k *KhronosTextureContainer) SetGlFormat(glFormat float64) *KhronosTextureContainer {
	p := ba.ctx.Get("KhronosTextureContainer").New(glFormat)
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

// GlInternalFormat returns the GlInternalFormat property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#glinternalformat
func (k *KhronosTextureContainer) GlInternalFormat(glInternalFormat float64) *KhronosTextureContainer {
	p := ba.ctx.Get("KhronosTextureContainer").New(glInternalFormat)
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

// SetGlInternalFormat sets the GlInternalFormat property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#glinternalformat
func (k *KhronosTextureContainer) SetGlInternalFormat(glInternalFormat float64) *KhronosTextureContainer {
	p := ba.ctx.Get("KhronosTextureContainer").New(glInternalFormat)
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

// GlType returns the GlType property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#gltype
func (k *KhronosTextureContainer) GlType(glType float64) *KhronosTextureContainer {
	p := ba.ctx.Get("KhronosTextureContainer").New(glType)
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

// SetGlType sets the GlType property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#gltype
func (k *KhronosTextureContainer) SetGlType(glType float64) *KhronosTextureContainer {
	p := ba.ctx.Get("KhronosTextureContainer").New(glType)
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

// GlTypeSize returns the GlTypeSize property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#gltypesize
func (k *KhronosTextureContainer) GlTypeSize(glTypeSize float64) *KhronosTextureContainer {
	p := ba.ctx.Get("KhronosTextureContainer").New(glTypeSize)
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

// SetGlTypeSize sets the GlTypeSize property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#gltypesize
func (k *KhronosTextureContainer) SetGlTypeSize(glTypeSize float64) *KhronosTextureContainer {
	p := ba.ctx.Get("KhronosTextureContainer").New(glTypeSize)
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

// IsInvalid returns the IsInvalid property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#isinvalid
func (k *KhronosTextureContainer) IsInvalid(isInvalid bool) *KhronosTextureContainer {
	p := ba.ctx.Get("KhronosTextureContainer").New(isInvalid)
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

// SetIsInvalid sets the IsInvalid property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#isinvalid
func (k *KhronosTextureContainer) SetIsInvalid(isInvalid bool) *KhronosTextureContainer {
	p := ba.ctx.Get("KhronosTextureContainer").New(isInvalid)
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

// LoadType returns the LoadType property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#loadtype
func (k *KhronosTextureContainer) LoadType(loadType float64) *KhronosTextureContainer {
	p := ba.ctx.Get("KhronosTextureContainer").New(loadType)
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

// SetLoadType sets the LoadType property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#loadtype
func (k *KhronosTextureContainer) SetLoadType(loadType float64) *KhronosTextureContainer {
	p := ba.ctx.Get("KhronosTextureContainer").New(loadType)
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

// NumberOfArrayElements returns the NumberOfArrayElements property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#numberofarrayelements
func (k *KhronosTextureContainer) NumberOfArrayElements(numberOfArrayElements float64) *KhronosTextureContainer {
	p := ba.ctx.Get("KhronosTextureContainer").New(numberOfArrayElements)
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

// SetNumberOfArrayElements sets the NumberOfArrayElements property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#numberofarrayelements
func (k *KhronosTextureContainer) SetNumberOfArrayElements(numberOfArrayElements float64) *KhronosTextureContainer {
	p := ba.ctx.Get("KhronosTextureContainer").New(numberOfArrayElements)
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

// NumberOfFaces returns the NumberOfFaces property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#numberoffaces
func (k *KhronosTextureContainer) NumberOfFaces(numberOfFaces float64) *KhronosTextureContainer {
	p := ba.ctx.Get("KhronosTextureContainer").New(numberOfFaces)
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

// SetNumberOfFaces sets the NumberOfFaces property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#numberoffaces
func (k *KhronosTextureContainer) SetNumberOfFaces(numberOfFaces float64) *KhronosTextureContainer {
	p := ba.ctx.Get("KhronosTextureContainer").New(numberOfFaces)
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

// NumberOfMipmapLevels returns the NumberOfMipmapLevels property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#numberofmipmaplevels
func (k *KhronosTextureContainer) NumberOfMipmapLevels(numberOfMipmapLevels float64) *KhronosTextureContainer {
	p := ba.ctx.Get("KhronosTextureContainer").New(numberOfMipmapLevels)
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

// SetNumberOfMipmapLevels sets the NumberOfMipmapLevels property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#numberofmipmaplevels
func (k *KhronosTextureContainer) SetNumberOfMipmapLevels(numberOfMipmapLevels float64) *KhronosTextureContainer {
	p := ba.ctx.Get("KhronosTextureContainer").New(numberOfMipmapLevels)
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

// PixelDepth returns the PixelDepth property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#pixeldepth
func (k *KhronosTextureContainer) PixelDepth(pixelDepth float64) *KhronosTextureContainer {
	p := ba.ctx.Get("KhronosTextureContainer").New(pixelDepth)
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

// SetPixelDepth sets the PixelDepth property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#pixeldepth
func (k *KhronosTextureContainer) SetPixelDepth(pixelDepth float64) *KhronosTextureContainer {
	p := ba.ctx.Get("KhronosTextureContainer").New(pixelDepth)
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

// PixelHeight returns the PixelHeight property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#pixelheight
func (k *KhronosTextureContainer) PixelHeight(pixelHeight float64) *KhronosTextureContainer {
	p := ba.ctx.Get("KhronosTextureContainer").New(pixelHeight)
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

// SetPixelHeight sets the PixelHeight property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#pixelheight
func (k *KhronosTextureContainer) SetPixelHeight(pixelHeight float64) *KhronosTextureContainer {
	p := ba.ctx.Get("KhronosTextureContainer").New(pixelHeight)
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

// PixelWidth returns the PixelWidth property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#pixelwidth
func (k *KhronosTextureContainer) PixelWidth(pixelWidth float64) *KhronosTextureContainer {
	p := ba.ctx.Get("KhronosTextureContainer").New(pixelWidth)
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

// SetPixelWidth sets the PixelWidth property of class KhronosTextureContainer.
//
// https://doc.babylonjs.com/api/classes/babylon.khronostexturecontainer#pixelwidth
func (k *KhronosTextureContainer) SetPixelWidth(pixelWidth float64) *KhronosTextureContainer {
	p := ba.ctx.Get("KhronosTextureContainer").New(pixelWidth)
	return KhronosTextureContainerFromJSObject(p, ba.ctx)
}

*/
