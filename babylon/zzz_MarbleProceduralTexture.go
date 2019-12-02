// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// MarbleProceduralTexture represents a babylon.js MarbleProceduralTexture.
//
type MarbleProceduralTexture struct {
	*ProceduralTexture
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (m *MarbleProceduralTexture) JSObject() js.Value { return m.p }

// MarbleProceduralTexture returns a MarbleProceduralTexture JavaScript class.
func (ba *Babylon) MarbleProceduralTexture() *MarbleProceduralTexture {
	p := ba.ctx.Get("MarbleProceduralTexture")
	return MarbleProceduralTextureFromJSObject(p, ba.ctx)
}

// MarbleProceduralTextureFromJSObject returns a wrapped MarbleProceduralTexture JavaScript class.
func MarbleProceduralTextureFromJSObject(p js.Value, ctx js.Value) *MarbleProceduralTexture {
	return &MarbleProceduralTexture{ProceduralTexture: ProceduralTextureFromJSObject(p, ctx), ctx: ctx}
}

// MarbleProceduralTextureArrayToJSArray returns a JavaScript Array for the wrapped array.
func MarbleProceduralTextureArrayToJSArray(array []*MarbleProceduralTexture) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewMarbleProceduralTextureOpts contains optional parameters for NewMarbleProceduralTexture.
type NewMarbleProceduralTextureOpts struct {
	FallbackTexture *Texture
	GenerateMipMaps *bool
}

// NewMarbleProceduralTexture returns a new MarbleProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.marbleproceduraltexture
func (ba *Babylon) NewMarbleProceduralTexture(name string, size float64, scene *Scene, opts *NewMarbleProceduralTextureOpts) *MarbleProceduralTexture {
	if opts == nil {
		opts = &NewMarbleProceduralTextureOpts{}
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

	p := ba.ctx.Get("MarbleProceduralTexture").New(args...)
	return MarbleProceduralTextureFromJSObject(p, ba.ctx)
}

// Parse calls the Parse method on the MarbleProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.marbleproceduraltexture#parse
func (m *MarbleProceduralTexture) Parse(parsedTexture interface{}, scene *Scene, rootUrl string) *MarbleProceduralTexture {

	args := make([]interface{}, 0, 3+0)

	args = append(args, parsedTexture)
	args = append(args, scene.JSObject())
	args = append(args, rootUrl)

	retVal := m.p.Call("Parse", args...)
	return MarbleProceduralTextureFromJSObject(retVal, m.ctx)
}

// Serialize calls the Serialize method on the MarbleProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.marbleproceduraltexture#serialize
func (m *MarbleProceduralTexture) Serialize() interface{} {

	retVal := m.p.Call("serialize")
	return retVal
}

// UpdateShaderUniforms calls the UpdateShaderUniforms method on the MarbleProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.marbleproceduraltexture#updateshaderuniforms
func (m *MarbleProceduralTexture) UpdateShaderUniforms() {

	m.p.Call("updateShaderUniforms")
}

/*

// Amplitude returns the Amplitude property of class MarbleProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.marbleproceduraltexture#amplitude
func (m *MarbleProceduralTexture) Amplitude(amplitude float64) *MarbleProceduralTexture {
	p := ba.ctx.Get("MarbleProceduralTexture").New(amplitude)
	return MarbleProceduralTextureFromJSObject(p, ba.ctx)
}

// SetAmplitude sets the Amplitude property of class MarbleProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.marbleproceduraltexture#amplitude
func (m *MarbleProceduralTexture) SetAmplitude(amplitude float64) *MarbleProceduralTexture {
	p := ba.ctx.Get("MarbleProceduralTexture").New(amplitude)
	return MarbleProceduralTextureFromJSObject(p, ba.ctx)
}

// JointColor returns the JointColor property of class MarbleProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.marbleproceduraltexture#jointcolor
func (m *MarbleProceduralTexture) JointColor(jointColor *Color3) *MarbleProceduralTexture {
	p := ba.ctx.Get("MarbleProceduralTexture").New(jointColor.JSObject())
	return MarbleProceduralTextureFromJSObject(p, ba.ctx)
}

// SetJointColor sets the JointColor property of class MarbleProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.marbleproceduraltexture#jointcolor
func (m *MarbleProceduralTexture) SetJointColor(jointColor *Color3) *MarbleProceduralTexture {
	p := ba.ctx.Get("MarbleProceduralTexture").New(jointColor.JSObject())
	return MarbleProceduralTextureFromJSObject(p, ba.ctx)
}

// NumberOfTilesHeight returns the NumberOfTilesHeight property of class MarbleProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.marbleproceduraltexture#numberoftilesheight
func (m *MarbleProceduralTexture) NumberOfTilesHeight(numberOfTilesHeight float64) *MarbleProceduralTexture {
	p := ba.ctx.Get("MarbleProceduralTexture").New(numberOfTilesHeight)
	return MarbleProceduralTextureFromJSObject(p, ba.ctx)
}

// SetNumberOfTilesHeight sets the NumberOfTilesHeight property of class MarbleProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.marbleproceduraltexture#numberoftilesheight
func (m *MarbleProceduralTexture) SetNumberOfTilesHeight(numberOfTilesHeight float64) *MarbleProceduralTexture {
	p := ba.ctx.Get("MarbleProceduralTexture").New(numberOfTilesHeight)
	return MarbleProceduralTextureFromJSObject(p, ba.ctx)
}

// NumberOfTilesWidth returns the NumberOfTilesWidth property of class MarbleProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.marbleproceduraltexture#numberoftileswidth
func (m *MarbleProceduralTexture) NumberOfTilesWidth(numberOfTilesWidth float64) *MarbleProceduralTexture {
	p := ba.ctx.Get("MarbleProceduralTexture").New(numberOfTilesWidth)
	return MarbleProceduralTextureFromJSObject(p, ba.ctx)
}

// SetNumberOfTilesWidth sets the NumberOfTilesWidth property of class MarbleProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.marbleproceduraltexture#numberoftileswidth
func (m *MarbleProceduralTexture) SetNumberOfTilesWidth(numberOfTilesWidth float64) *MarbleProceduralTexture {
	p := ba.ctx.Get("MarbleProceduralTexture").New(numberOfTilesWidth)
	return MarbleProceduralTextureFromJSObject(p, ba.ctx)
}

*/
