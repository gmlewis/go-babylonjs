// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DepthTextureCreationOptions represents a babylon.js DepthTextureCreationOptions.
// Define options used to create a depth texture
type DepthTextureCreationOptions struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (d *DepthTextureCreationOptions) JSObject() js.Value { return d.p }

// DepthTextureCreationOptions returns a DepthTextureCreationOptions JavaScript class.
func (ba *Babylon) DepthTextureCreationOptions() *DepthTextureCreationOptions {
	p := ba.ctx.Get("DepthTextureCreationOptions")
	return DepthTextureCreationOptionsFromJSObject(p, ba.ctx)
}

// DepthTextureCreationOptionsFromJSObject returns a wrapped DepthTextureCreationOptions JavaScript class.
func DepthTextureCreationOptionsFromJSObject(p js.Value, ctx js.Value) *DepthTextureCreationOptions {
	return &DepthTextureCreationOptions{p: p, ctx: ctx}
}

// DepthTextureCreationOptionsArrayToJSArray returns a JavaScript Array for the wrapped array.
func DepthTextureCreationOptionsArrayToJSArray(array []*DepthTextureCreationOptions) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

/*

// BilinearFiltering returns the BilinearFiltering property of class DepthTextureCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.depthtexturecreationoptions#bilinearfiltering
func (d *DepthTextureCreationOptions) BilinearFiltering(bilinearFiltering bool) *DepthTextureCreationOptions {
	p := ba.ctx.Get("DepthTextureCreationOptions").New(bilinearFiltering)
	return DepthTextureCreationOptionsFromJSObject(p, ba.ctx)
}

// SetBilinearFiltering sets the BilinearFiltering property of class DepthTextureCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.depthtexturecreationoptions#bilinearfiltering
func (d *DepthTextureCreationOptions) SetBilinearFiltering(bilinearFiltering bool) *DepthTextureCreationOptions {
	p := ba.ctx.Get("DepthTextureCreationOptions").New(bilinearFiltering)
	return DepthTextureCreationOptionsFromJSObject(p, ba.ctx)
}

// ComparisonFunction returns the ComparisonFunction property of class DepthTextureCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.depthtexturecreationoptions#comparisonfunction
func (d *DepthTextureCreationOptions) ComparisonFunction(comparisonFunction float64) *DepthTextureCreationOptions {
	p := ba.ctx.Get("DepthTextureCreationOptions").New(comparisonFunction)
	return DepthTextureCreationOptionsFromJSObject(p, ba.ctx)
}

// SetComparisonFunction sets the ComparisonFunction property of class DepthTextureCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.depthtexturecreationoptions#comparisonfunction
func (d *DepthTextureCreationOptions) SetComparisonFunction(comparisonFunction float64) *DepthTextureCreationOptions {
	p := ba.ctx.Get("DepthTextureCreationOptions").New(comparisonFunction)
	return DepthTextureCreationOptionsFromJSObject(p, ba.ctx)
}

// GenerateStencil returns the GenerateStencil property of class DepthTextureCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.depthtexturecreationoptions#generatestencil
func (d *DepthTextureCreationOptions) GenerateStencil(generateStencil bool) *DepthTextureCreationOptions {
	p := ba.ctx.Get("DepthTextureCreationOptions").New(generateStencil)
	return DepthTextureCreationOptionsFromJSObject(p, ba.ctx)
}

// SetGenerateStencil sets the GenerateStencil property of class DepthTextureCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.depthtexturecreationoptions#generatestencil
func (d *DepthTextureCreationOptions) SetGenerateStencil(generateStencil bool) *DepthTextureCreationOptions {
	p := ba.ctx.Get("DepthTextureCreationOptions").New(generateStencil)
	return DepthTextureCreationOptionsFromJSObject(p, ba.ctx)
}

// IsCube returns the IsCube property of class DepthTextureCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.depthtexturecreationoptions#iscube
func (d *DepthTextureCreationOptions) IsCube(isCube bool) *DepthTextureCreationOptions {
	p := ba.ctx.Get("DepthTextureCreationOptions").New(isCube)
	return DepthTextureCreationOptionsFromJSObject(p, ba.ctx)
}

// SetIsCube sets the IsCube property of class DepthTextureCreationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.depthtexturecreationoptions#iscube
func (d *DepthTextureCreationOptions) SetIsCube(isCube bool) *DepthTextureCreationOptions {
	p := ba.ctx.Get("DepthTextureCreationOptions").New(isCube)
	return DepthTextureCreationOptionsFromJSObject(p, ba.ctx)
}

*/
