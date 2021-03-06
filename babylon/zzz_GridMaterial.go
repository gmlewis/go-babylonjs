// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GridMaterial represents a babylon.js GridMaterial.
// The grid materials allows you to wrap any shape with a grid.
// Colors are customizable.
type GridMaterial struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (g *GridMaterial) JSObject() js.Value { return g.p }

// GridMaterial returns a GridMaterial JavaScript class.
func (ba *Babylon) GridMaterial() *GridMaterial {
	p := ba.ctx.Get("GridMaterial")
	return GridMaterialFromJSObject(p, ba.ctx)
}

// GridMaterialFromJSObject returns a wrapped GridMaterial JavaScript class.
func GridMaterialFromJSObject(p js.Value, ctx js.Value) *GridMaterial {
	return &GridMaterial{p: p, ctx: ctx}
}

// GridMaterialArrayToJSArray returns a JavaScript Array for the wrapped array.
func GridMaterialArrayToJSArray(array []*GridMaterial) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewGridMaterial returns a new GridMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.gridmaterial#constructor
func (ba *Babylon) NewGridMaterial(name string, scene *Scene) *GridMaterial {

	args := make([]interface{}, 0, 2+0)

	args = append(args, name)
	args = append(args, scene.JSObject())

	p := ba.ctx.Get("GridMaterial").New(args...)
	return GridMaterialFromJSObject(p, ba.ctx)
}

// BindForSubMesh calls the BindForSubMesh method on the GridMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.gridmaterial#bindforsubmesh
func (g *GridMaterial) BindForSubMesh(world *Matrix, mesh *Mesh, subMesh *SubMesh) {

	args := make([]interface{}, 0, 3+0)

	if world == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, world.JSObject())
	}

	if mesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, mesh.JSObject())
	}

	if subMesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, subMesh.JSObject())
	}

	g.p.Call("bindForSubMesh", args...)
}

// Clone calls the Clone method on the GridMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.gridmaterial#clone
func (g *GridMaterial) Clone(name string) *GridMaterial {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := g.p.Call("clone", args...)
	return GridMaterialFromJSObject(retVal, g.ctx)
}

// GridMaterialDisposeOpts contains optional parameters for GridMaterial.Dispose.
type GridMaterialDisposeOpts struct {
	ForceDisposeEffect *bool
}

// Dispose calls the Dispose method on the GridMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.gridmaterial#dispose
func (g *GridMaterial) Dispose(opts *GridMaterialDisposeOpts) {
	if opts == nil {
		opts = &GridMaterialDisposeOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForceDisposeEffect == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ForceDisposeEffect)
	}

	g.p.Call("dispose", args...)
}

// GetClassName calls the GetClassName method on the GridMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.gridmaterial#getclassname
func (g *GridMaterial) GetClassName() string {

	retVal := g.p.Call("getClassName")
	return retVal.String()
}

// GridMaterialIsReadyForSubMeshOpts contains optional parameters for GridMaterial.IsReadyForSubMesh.
type GridMaterialIsReadyForSubMeshOpts struct {
	UseInstances *bool
}

// IsReadyForSubMesh calls the IsReadyForSubMesh method on the GridMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.gridmaterial#isreadyforsubmesh
func (g *GridMaterial) IsReadyForSubMesh(mesh *AbstractMesh, subMesh *SubMesh, opts *GridMaterialIsReadyForSubMeshOpts) bool {
	if opts == nil {
		opts = &GridMaterialIsReadyForSubMeshOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	if mesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, mesh.JSObject())
	}

	if subMesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, subMesh.JSObject())
	}

	if opts.UseInstances == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UseInstances)
	}

	retVal := g.p.Call("isReadyForSubMesh", args...)
	return retVal.Bool()
}

// NeedAlphaBlending calls the NeedAlphaBlending method on the GridMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.gridmaterial#needalphablending
func (g *GridMaterial) NeedAlphaBlending() bool {

	retVal := g.p.Call("needAlphaBlending")
	return retVal.Bool()
}

// NeedAlphaBlendingForMesh calls the NeedAlphaBlendingForMesh method on the GridMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.gridmaterial#needalphablendingformesh
func (g *GridMaterial) NeedAlphaBlendingForMesh(mesh *AbstractMesh) bool {

	args := make([]interface{}, 0, 1+0)

	if mesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, mesh.JSObject())
	}

	retVal := g.p.Call("needAlphaBlendingForMesh", args...)
	return retVal.Bool()
}

// Parse calls the Parse method on the GridMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.gridmaterial#parse
func (g *GridMaterial) Parse(source JSObject, scene *Scene, rootUrl string) *GridMaterial {

	args := make([]interface{}, 0, 3+0)

	if source == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, source.JSObject())
	}

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
	}

	args = append(args, rootUrl)

	retVal := g.p.Call("Parse", args...)
	return GridMaterialFromJSObject(retVal, g.ctx)
}

// Serialize calls the Serialize method on the GridMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.gridmaterial#serialize
func (g *GridMaterial) Serialize() js.Value {

	retVal := g.p.Call("serialize")
	return retVal
}

// GridOffset returns the GridOffset property of class GridMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.gridmaterial#gridoffset
func (g *GridMaterial) GridOffset() *Vector3 {
	retVal := g.p.Get("gridOffset")
	return Vector3FromJSObject(retVal, g.ctx)
}

// SetGridOffset sets the GridOffset property of class GridMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.gridmaterial#gridoffset
func (g *GridMaterial) SetGridOffset(gridOffset *Vector3) *GridMaterial {
	g.p.Set("gridOffset", gridOffset.JSObject())
	return g
}

// GridRatio returns the GridRatio property of class GridMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.gridmaterial#gridratio
func (g *GridMaterial) GridRatio() float64 {
	retVal := g.p.Get("gridRatio")
	return retVal.Float()
}

// SetGridRatio sets the GridRatio property of class GridMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.gridmaterial#gridratio
func (g *GridMaterial) SetGridRatio(gridRatio float64) *GridMaterial {
	g.p.Set("gridRatio", gridRatio)
	return g
}

// LineColor returns the LineColor property of class GridMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.gridmaterial#linecolor
func (g *GridMaterial) LineColor() *Color3 {
	retVal := g.p.Get("lineColor")
	return Color3FromJSObject(retVal, g.ctx)
}

// SetLineColor sets the LineColor property of class GridMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.gridmaterial#linecolor
func (g *GridMaterial) SetLineColor(lineColor *Color3) *GridMaterial {
	g.p.Set("lineColor", lineColor.JSObject())
	return g
}

// MainColor returns the MainColor property of class GridMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.gridmaterial#maincolor
func (g *GridMaterial) MainColor() *Color3 {
	retVal := g.p.Get("mainColor")
	return Color3FromJSObject(retVal, g.ctx)
}

// SetMainColor sets the MainColor property of class GridMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.gridmaterial#maincolor
func (g *GridMaterial) SetMainColor(mainColor *Color3) *GridMaterial {
	g.p.Set("mainColor", mainColor.JSObject())
	return g
}

// MajorUnitFrequency returns the MajorUnitFrequency property of class GridMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.gridmaterial#majorunitfrequency
func (g *GridMaterial) MajorUnitFrequency() float64 {
	retVal := g.p.Get("majorUnitFrequency")
	return retVal.Float()
}

// SetMajorUnitFrequency sets the MajorUnitFrequency property of class GridMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.gridmaterial#majorunitfrequency
func (g *GridMaterial) SetMajorUnitFrequency(majorUnitFrequency float64) *GridMaterial {
	g.p.Set("majorUnitFrequency", majorUnitFrequency)
	return g
}

// MinorUnitVisibility returns the MinorUnitVisibility property of class GridMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.gridmaterial#minorunitvisibility
func (g *GridMaterial) MinorUnitVisibility() float64 {
	retVal := g.p.Get("minorUnitVisibility")
	return retVal.Float()
}

// SetMinorUnitVisibility sets the MinorUnitVisibility property of class GridMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.gridmaterial#minorunitvisibility
func (g *GridMaterial) SetMinorUnitVisibility(minorUnitVisibility float64) *GridMaterial {
	g.p.Set("minorUnitVisibility", minorUnitVisibility)
	return g
}

// Opacity returns the Opacity property of class GridMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.gridmaterial#opacity
func (g *GridMaterial) Opacity() float64 {
	retVal := g.p.Get("opacity")
	return retVal.Float()
}

// SetOpacity sets the Opacity property of class GridMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.gridmaterial#opacity
func (g *GridMaterial) SetOpacity(opacity float64) *GridMaterial {
	g.p.Set("opacity", opacity)
	return g
}

// OpacityTexture returns the OpacityTexture property of class GridMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.gridmaterial#opacitytexture
func (g *GridMaterial) OpacityTexture() *BaseTexture {
	retVal := g.p.Get("opacityTexture")
	return BaseTextureFromJSObject(retVal, g.ctx)
}

// SetOpacityTexture sets the OpacityTexture property of class GridMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.gridmaterial#opacitytexture
func (g *GridMaterial) SetOpacityTexture(opacityTexture *BaseTexture) *GridMaterial {
	g.p.Set("opacityTexture", opacityTexture.JSObject())
	return g
}

// PreMultiplyAlpha returns the PreMultiplyAlpha property of class GridMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.gridmaterial#premultiplyalpha
func (g *GridMaterial) PreMultiplyAlpha() bool {
	retVal := g.p.Get("preMultiplyAlpha")
	return retVal.Bool()
}

// SetPreMultiplyAlpha sets the PreMultiplyAlpha property of class GridMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.gridmaterial#premultiplyalpha
func (g *GridMaterial) SetPreMultiplyAlpha(preMultiplyAlpha bool) *GridMaterial {
	g.p.Set("preMultiplyAlpha", preMultiplyAlpha)
	return g
}
