// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// CloudProceduralTexture represents a babylon.js CloudProceduralTexture.
//
type CloudProceduralTexture struct {
	*ProceduralTexture
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (c *CloudProceduralTexture) JSObject() js.Value { return c.p }

// CloudProceduralTexture returns a CloudProceduralTexture JavaScript class.
func (ba *Babylon) CloudProceduralTexture() *CloudProceduralTexture {
	p := ba.ctx.Get("CloudProceduralTexture")
	return CloudProceduralTextureFromJSObject(p, ba.ctx)
}

// CloudProceduralTextureFromJSObject returns a wrapped CloudProceduralTexture JavaScript class.
func CloudProceduralTextureFromJSObject(p js.Value, ctx js.Value) *CloudProceduralTexture {
	return &CloudProceduralTexture{ProceduralTexture: ProceduralTextureFromJSObject(p, ctx), ctx: ctx}
}

// CloudProceduralTextureArrayToJSArray returns a JavaScript Array for the wrapped array.
func CloudProceduralTextureArrayToJSArray(array []*CloudProceduralTexture) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewCloudProceduralTextureOpts contains optional parameters for NewCloudProceduralTexture.
type NewCloudProceduralTextureOpts struct {
	FallbackTexture *Texture
	GenerateMipMaps *bool
}

// NewCloudProceduralTexture returns a new CloudProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.cloudproceduraltexture#constructor
func (ba *Babylon) NewCloudProceduralTexture(name string, size float64, scene *Scene, opts *NewCloudProceduralTextureOpts) *CloudProceduralTexture {
	if opts == nil {
		opts = &NewCloudProceduralTextureOpts{}
	}

	args := make([]interface{}, 0, 3+2)

	args = append(args, name)
	args = append(args, size)
	args = append(args, scene.JSObject())

	if opts.FallbackTexture == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.FallbackTexture.JSObject())
	}
	if opts.GenerateMipMaps == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.GenerateMipMaps)
	}

	p := ba.ctx.Get("CloudProceduralTexture").New(args...)
	return CloudProceduralTextureFromJSObject(p, ba.ctx)
}

// Parse calls the Parse method on the CloudProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.cloudproceduraltexture#parse
func (c *CloudProceduralTexture) Parse(parsedTexture JSObject, scene *Scene, rootUrl string) *CloudProceduralTexture {

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
	return CloudProceduralTextureFromJSObject(retVal, c.ctx)
}

// Serialize calls the Serialize method on the CloudProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.cloudproceduraltexture#serialize
func (c *CloudProceduralTexture) Serialize() js.Value {

	retVal := c.p.Call("serialize")
	return retVal
}

// UpdateShaderUniforms calls the UpdateShaderUniforms method on the CloudProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.cloudproceduraltexture#updateshaderuniforms
func (c *CloudProceduralTexture) UpdateShaderUniforms() {

	c.p.Call("updateShaderUniforms")
}

// CloudColor returns the CloudColor property of class CloudProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.cloudproceduraltexture#cloudcolor
func (c *CloudProceduralTexture) CloudColor() *Color4 {
	retVal := c.p.Get("cloudColor")
	return Color4FromJSObject(retVal, c.ctx)
}

// SetCloudColor sets the CloudColor property of class CloudProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.cloudproceduraltexture#cloudcolor
func (c *CloudProceduralTexture) SetCloudColor(cloudColor *Color4) *CloudProceduralTexture {
	c.p.Set("cloudColor", cloudColor.JSObject())
	return c
}

// SkyColor returns the SkyColor property of class CloudProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.cloudproceduraltexture#skycolor
func (c *CloudProceduralTexture) SkyColor() *Color4 {
	retVal := c.p.Get("skyColor")
	return Color4FromJSObject(retVal, c.ctx)
}

// SetSkyColor sets the SkyColor property of class CloudProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.cloudproceduraltexture#skycolor
func (c *CloudProceduralTexture) SetSkyColor(skyColor *Color4) *CloudProceduralTexture {
	c.p.Set("skyColor", skyColor.JSObject())
	return c
}
