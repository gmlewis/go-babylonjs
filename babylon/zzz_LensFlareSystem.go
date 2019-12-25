// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// LensFlareSystem represents a babylon.js LensFlareSystem.
// This represents a Lens Flare System or the shiny effect created by the light reflection on the  camera lenses.
// It is usually composed of several <code>lensFlare</code>.
//
// See: http://doc.babylonjs.com/how_to/how_to_use_lens_flares
type LensFlareSystem struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (l *LensFlareSystem) JSObject() js.Value { return l.p }

// LensFlareSystem returns a LensFlareSystem JavaScript class.
func (ba *Babylon) LensFlareSystem() *LensFlareSystem {
	p := ba.ctx.Get("LensFlareSystem")
	return LensFlareSystemFromJSObject(p, ba.ctx)
}

// LensFlareSystemFromJSObject returns a wrapped LensFlareSystem JavaScript class.
func LensFlareSystemFromJSObject(p js.Value, ctx js.Value) *LensFlareSystem {
	return &LensFlareSystem{p: p, ctx: ctx}
}

// LensFlareSystemArrayToJSArray returns a JavaScript Array for the wrapped array.
func LensFlareSystemArrayToJSArray(array []*LensFlareSystem) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewLensFlareSystem returns a new LensFlareSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystem#constructor
func (ba *Babylon) NewLensFlareSystem(name string, emitter JSObject, scene *Scene) *LensFlareSystem {

	args := make([]interface{}, 0, 3+0)

	args = append(args, name)
	args = append(args, emitter.JSObject())
	args = append(args, scene.JSObject())

	p := ba.ctx.Get("LensFlareSystem").New(args...)
	return LensFlareSystemFromJSObject(p, ba.ctx)
}

// Dispose calls the Dispose method on the LensFlareSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystem#dispose
func (l *LensFlareSystem) Dispose() {

	l.p.Call("dispose")
}

// GetEmitter calls the GetEmitter method on the LensFlareSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystem#getemitter
func (l *LensFlareSystem) GetEmitter() js.Value {

	retVal := l.p.Call("getEmitter")
	return retVal
}

// GetEmitterPosition calls the GetEmitterPosition method on the LensFlareSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystem#getemitterposition
func (l *LensFlareSystem) GetEmitterPosition() *Vector3 {

	retVal := l.p.Call("getEmitterPosition")
	return Vector3FromJSObject(retVal, l.ctx)
}

// GetScene calls the GetScene method on the LensFlareSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystem#getscene
func (l *LensFlareSystem) GetScene() *Scene {

	retVal := l.p.Call("getScene")
	return SceneFromJSObject(retVal, l.ctx)
}

// Parse calls the Parse method on the LensFlareSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystem#parse
func (l *LensFlareSystem) Parse(parsedLensFlareSystem JSObject, scene *Scene, rootUrl string) *LensFlareSystem {

	args := make([]interface{}, 0, 3+0)

	if parsedLensFlareSystem == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, parsedLensFlareSystem.JSObject())
	}

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
	}

	args = append(args, rootUrl)

	retVal := l.p.Call("Parse", args...)
	return LensFlareSystemFromJSObject(retVal, l.ctx)
}

// Serialize calls the Serialize method on the LensFlareSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystem#serialize
func (l *LensFlareSystem) Serialize() js.Value {

	retVal := l.p.Call("serialize")
	return retVal
}

// SetEmitter calls the SetEmitter method on the LensFlareSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystem#setemitter
func (l *LensFlareSystem) SetEmitter(newEmitter JSObject) {

	args := make([]interface{}, 0, 1+0)

	if newEmitter == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, newEmitter.JSObject())
	}

	l.p.Call("setEmitter", args...)
}

// BorderLimit returns the BorderLimit property of class LensFlareSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystem#borderlimit
func (l *LensFlareSystem) BorderLimit() float64 {
	retVal := l.p.Get("borderLimit")
	return retVal.Float()
}

// SetBorderLimit sets the BorderLimit property of class LensFlareSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystem#borderlimit
func (l *LensFlareSystem) SetBorderLimit(borderLimit float64) *LensFlareSystem {
	l.p.Set("borderLimit", borderLimit)
	return l
}

// Id returns the Id property of class LensFlareSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystem#id
func (l *LensFlareSystem) Id() string {
	retVal := l.p.Get("id")
	return retVal.String()
}

// SetId sets the Id property of class LensFlareSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystem#id
func (l *LensFlareSystem) SetId(id string) *LensFlareSystem {
	l.p.Set("id", id)
	return l
}

// IsEnabled returns the IsEnabled property of class LensFlareSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystem#isenabled
func (l *LensFlareSystem) IsEnabled() bool {
	retVal := l.p.Get("isEnabled")
	return retVal.Bool()
}

// SetIsEnabled sets the IsEnabled property of class LensFlareSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystem#isenabled
func (l *LensFlareSystem) SetIsEnabled(isEnabled bool) *LensFlareSystem {
	l.p.Set("isEnabled", isEnabled)
	return l
}

// LayerMask returns the LayerMask property of class LensFlareSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystem#layermask
func (l *LensFlareSystem) LayerMask() float64 {
	retVal := l.p.Get("layerMask")
	return retVal.Float()
}

// SetLayerMask sets the LayerMask property of class LensFlareSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystem#layermask
func (l *LensFlareSystem) SetLayerMask(layerMask float64) *LensFlareSystem {
	l.p.Set("layerMask", layerMask)
	return l
}

// LensFlares returns the LensFlares property of class LensFlareSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystem#lensflares
func (l *LensFlareSystem) LensFlares() []*LensFlare {
	retVal := l.p.Get("lensFlares")
	result := []*LensFlare{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, LensFlareFromJSObject(retVal.Index(ri), l.ctx))
	}
	return result
}

// SetLensFlares sets the LensFlares property of class LensFlareSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystem#lensflares
func (l *LensFlareSystem) SetLensFlares(lensFlares []*LensFlare) *LensFlareSystem {
	l.p.Set("lensFlares", lensFlares)
	return l
}

// MeshesSelectionPredicate returns the MeshesSelectionPredicate property of class LensFlareSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystem#meshesselectionpredicate
func (l *LensFlareSystem) MeshesSelectionPredicate() js.Value {
	retVal := l.p.Get("meshesSelectionPredicate")
	return retVal
}

// SetMeshesSelectionPredicate sets the MeshesSelectionPredicate property of class LensFlareSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystem#meshesselectionpredicate
func (l *LensFlareSystem) SetMeshesSelectionPredicate(meshesSelectionPredicate JSFunc) *LensFlareSystem {
	l.p.Set("meshesSelectionPredicate", js.FuncOf(meshesSelectionPredicate))
	return l
}

// Name returns the Name property of class LensFlareSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystem#name
func (l *LensFlareSystem) Name() string {
	retVal := l.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class LensFlareSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystem#name
func (l *LensFlareSystem) SetName(name string) *LensFlareSystem {
	l.p.Set("name", name)
	return l
}

// ViewportBorder returns the ViewportBorder property of class LensFlareSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystem#viewportborder
func (l *LensFlareSystem) ViewportBorder() float64 {
	retVal := l.p.Get("viewportBorder")
	return retVal.Float()
}

// SetViewportBorder sets the ViewportBorder property of class LensFlareSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.lensflaresystem#viewportborder
func (l *LensFlareSystem) SetViewportBorder(viewportBorder float64) *LensFlareSystem {
	l.p.Set("viewportBorder", viewportBorder)
	return l
}
