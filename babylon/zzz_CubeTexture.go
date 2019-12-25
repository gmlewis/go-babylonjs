// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// CubeTexture represents a babylon.js CubeTexture.
// Class for creating a cube texture
type CubeTexture struct {
	*BaseTexture
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (c *CubeTexture) JSObject() js.Value { return c.p }

// CubeTexture returns a CubeTexture JavaScript class.
func (ba *Babylon) CubeTexture() *CubeTexture {
	p := ba.ctx.Get("CubeTexture")
	return CubeTextureFromJSObject(p, ba.ctx)
}

// CubeTextureFromJSObject returns a wrapped CubeTexture JavaScript class.
func CubeTextureFromJSObject(p js.Value, ctx js.Value) *CubeTexture {
	return &CubeTexture{BaseTexture: BaseTextureFromJSObject(p, ctx), ctx: ctx}
}

// CubeTextureArrayToJSArray returns a JavaScript Array for the wrapped array.
func CubeTextureArrayToJSArray(array []*CubeTexture) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewCubeTextureOpts contains optional parameters for NewCubeTexture.
type NewCubeTextureOpts struct {
	Extensions        []string
	NoMipmap          *bool
	Files             []string
	OnLoad            JSFunc
	OnError           JSFunc
	Format            *float64
	Prefiltered       *bool
	ForcedExtension   interface{}
	CreatePolynomials *bool
	LodScale          *float64
	LodOffset         *float64
}

// NewCubeTexture returns a new CubeTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.cubetexture#constructor
func (ba *Babylon) NewCubeTexture(rootUrl string, scene *Scene, opts *NewCubeTextureOpts) *CubeTexture {
	if opts == nil {
		opts = &NewCubeTextureOpts{}
	}

	args := make([]interface{}, 0, 2+11)

	args = append(args, rootUrl)
	args = append(args, scene.JSObject())

	if opts.Extensions == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Extensions)
	}
	if opts.NoMipmap == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.NoMipmap)
	}
	if opts.Files == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Files)
	}
	if opts.OnLoad == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(opts.OnLoad) /* never freed! */)
	}
	if opts.OnError == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(opts.OnError) /* never freed! */)
	}
	if opts.Format == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Format)
	}
	if opts.Prefiltered == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Prefiltered)
	}
	if opts.ForcedExtension == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForcedExtension)
	}
	if opts.CreatePolynomials == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.CreatePolynomials)
	}
	if opts.LodScale == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.LodScale)
	}
	if opts.LodOffset == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.LodOffset)
	}

	p := ba.ctx.Get("CubeTexture").New(args...)
	return CubeTextureFromJSObject(p, ba.ctx)
}

// Clone calls the Clone method on the CubeTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.cubetexture#clone
func (c *CubeTexture) Clone() *CubeTexture {

	retVal := c.p.Call("clone")
	return CubeTextureFromJSObject(retVal, c.ctx)
}

// CubeTextureCreateFromImagesOpts contains optional parameters for CubeTexture.CreateFromImages.
type CubeTextureCreateFromImagesOpts struct {
	NoMipmap *bool
}

// CreateFromImages calls the CreateFromImages method on the CubeTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.cubetexture#createfromimages
func (c *CubeTexture) CreateFromImages(files []string, scene *Scene, opts *CubeTextureCreateFromImagesOpts) *CubeTexture {
	if opts == nil {
		opts = &CubeTextureCreateFromImagesOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, files)

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
	}

	if opts.NoMipmap == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.NoMipmap)
	}

	retVal := c.p.Call("CreateFromImages", args...)
	return CubeTextureFromJSObject(retVal, c.ctx)
}

// CubeTextureCreateFromPrefilteredDataOpts contains optional parameters for CubeTexture.CreateFromPrefilteredData.
type CubeTextureCreateFromPrefilteredDataOpts struct {
	ForcedExtension   interface{}
	CreatePolynomials *bool
}

// CreateFromPrefilteredData calls the CreateFromPrefilteredData method on the CubeTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.cubetexture#createfromprefiltereddata
func (c *CubeTexture) CreateFromPrefilteredData(url string, scene *Scene, opts *CubeTextureCreateFromPrefilteredDataOpts) *CubeTexture {
	if opts == nil {
		opts = &CubeTextureCreateFromPrefilteredDataOpts{}
	}

	args := make([]interface{}, 0, 2+2)

	args = append(args, url)

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
	}

	if opts.ForcedExtension == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForcedExtension)
	}
	if opts.CreatePolynomials == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.CreatePolynomials)
	}

	retVal := c.p.Call("CreateFromPrefilteredData", args...)
	return CubeTextureFromJSObject(retVal, c.ctx)
}

// CubeTextureDelayLoadOpts contains optional parameters for CubeTexture.DelayLoad.
type CubeTextureDelayLoadOpts struct {
	ForcedExtension *string
}

// DelayLoad calls the DelayLoad method on the CubeTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.cubetexture#delayload
func (c *CubeTexture) DelayLoad(opts *CubeTextureDelayLoadOpts) {
	if opts == nil {
		opts = &CubeTextureDelayLoadOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForcedExtension == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ForcedExtension)
	}

	c.p.Call("delayLoad", args...)
}

// GetClassName calls the GetClassName method on the CubeTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.cubetexture#getclassname
func (c *CubeTexture) GetClassName() string {

	retVal := c.p.Call("getClassName")
	return retVal.String()
}

// GetReflectionTextureMatrix calls the GetReflectionTextureMatrix method on the CubeTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.cubetexture#getreflectiontexturematrix
func (c *CubeTexture) GetReflectionTextureMatrix() *Matrix {

	retVal := c.p.Call("getReflectionTextureMatrix")
	return MatrixFromJSObject(retVal, c.ctx)
}

// Parse calls the Parse method on the CubeTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.cubetexture#parse
func (c *CubeTexture) Parse(parsedTexture JSObject, scene *Scene, rootUrl string) *CubeTexture {

	args := make([]interface{}, 0, 3+0)

	if parsedTexture == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, parsedTexture.JSObject())
	}

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
	}

	args = append(args, rootUrl)

	retVal := c.p.Call("Parse", args...)
	return CubeTextureFromJSObject(retVal, c.ctx)
}

// SetReflectionTextureMatrix calls the SetReflectionTextureMatrix method on the CubeTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.cubetexture#setreflectiontexturematrix
func (c *CubeTexture) SetReflectionTextureMatrix(value *Matrix) {

	args := make([]interface{}, 0, 1+0)

	if value == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, value.JSObject())
	}

	c.p.Call("setReflectionTextureMatrix", args...)
}

// CubeTextureUpdateURLOpts contains optional parameters for CubeTexture.UpdateURL.
type CubeTextureUpdateURLOpts struct {
	ForcedExtension *string
	OnLoad          JSFunc
}

// UpdateURL calls the UpdateURL method on the CubeTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.cubetexture#updateurl
func (c *CubeTexture) UpdateURL(url string, opts *CubeTextureUpdateURLOpts) {
	if opts == nil {
		opts = &CubeTextureUpdateURLOpts{}
	}

	args := make([]interface{}, 0, 1+2)

	args = append(args, url)

	if opts.ForcedExtension == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ForcedExtension)
	}
	if opts.OnLoad == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(opts.OnLoad) /* never freed! */)
	}

	c.p.Call("updateURL", args...)
}

// BoundingBoxPosition returns the BoundingBoxPosition property of class CubeTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.cubetexture#boundingboxposition
func (c *CubeTexture) BoundingBoxPosition() *Vector3 {
	retVal := c.p.Get("boundingBoxPosition")
	return Vector3FromJSObject(retVal, c.ctx)
}

// SetBoundingBoxPosition sets the BoundingBoxPosition property of class CubeTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.cubetexture#boundingboxposition
func (c *CubeTexture) SetBoundingBoxPosition(boundingBoxPosition *Vector3) *CubeTexture {
	c.p.Set("boundingBoxPosition", boundingBoxPosition.JSObject())
	return c
}

// BoundingBoxSize returns the BoundingBoxSize property of class CubeTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.cubetexture#boundingboxsize
func (c *CubeTexture) BoundingBoxSize() *Vector3 {
	retVal := c.p.Get("boundingBoxSize")
	return Vector3FromJSObject(retVal, c.ctx)
}

// SetBoundingBoxSize sets the BoundingBoxSize property of class CubeTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.cubetexture#boundingboxsize
func (c *CubeTexture) SetBoundingBoxSize(boundingBoxSize *Vector3) *CubeTexture {
	c.p.Set("boundingBoxSize", boundingBoxSize.JSObject())
	return c
}

// IsPrefiltered returns the IsPrefiltered property of class CubeTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.cubetexture#isprefiltered
func (c *CubeTexture) IsPrefiltered() bool {
	retVal := c.p.Get("isPrefiltered")
	return retVal.Bool()
}

// SetIsPrefiltered sets the IsPrefiltered property of class CubeTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.cubetexture#isprefiltered
func (c *CubeTexture) SetIsPrefiltered(isPrefiltered bool) *CubeTexture {
	c.p.Set("isPrefiltered", isPrefiltered)
	return c
}

// NoMipmap returns the NoMipmap property of class CubeTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.cubetexture#nomipmap
func (c *CubeTexture) NoMipmap() bool {
	retVal := c.p.Get("noMipmap")
	return retVal.Bool()
}

// SetNoMipmap sets the NoMipmap property of class CubeTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.cubetexture#nomipmap
func (c *CubeTexture) SetNoMipmap(noMipmap bool) *CubeTexture {
	c.p.Set("noMipmap", noMipmap)
	return c
}

// RotationY returns the RotationY property of class CubeTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.cubetexture#rotationy
func (c *CubeTexture) RotationY() float64 {
	retVal := c.p.Get("rotationY")
	return retVal.Float()
}

// SetRotationY sets the RotationY property of class CubeTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.cubetexture#rotationy
func (c *CubeTexture) SetRotationY(rotationY float64) *CubeTexture {
	c.p.Set("rotationY", rotationY)
	return c
}

// Url returns the Url property of class CubeTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.cubetexture#url
func (c *CubeTexture) Url() string {
	retVal := c.p.Get("url")
	return retVal.String()
}

// SetUrl sets the Url property of class CubeTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.cubetexture#url
func (c *CubeTexture) SetUrl(url string) *CubeTexture {
	c.p.Set("url", url)
	return c
}
