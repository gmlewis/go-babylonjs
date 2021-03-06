// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ISimplificationSettings represents a babylon.js ISimplificationSettings.
// Expected simplification settings.
// Quality should be between 0 and 1 (1 being 100%, 0 being 0%)
//
// See: http://doc.babylonjs.com/how_to/in-browser_mesh_simplification
type ISimplificationSettings struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *ISimplificationSettings) JSObject() js.Value { return i.p }

// ISimplificationSettings returns a ISimplificationSettings JavaScript class.
func (ba *Babylon) ISimplificationSettings() *ISimplificationSettings {
	p := ba.ctx.Get("ISimplificationSettings")
	return ISimplificationSettingsFromJSObject(p, ba.ctx)
}

// ISimplificationSettingsFromJSObject returns a wrapped ISimplificationSettings JavaScript class.
func ISimplificationSettingsFromJSObject(p js.Value, ctx js.Value) *ISimplificationSettings {
	return &ISimplificationSettings{p: p, ctx: ctx}
}

// ISimplificationSettingsArrayToJSArray returns a JavaScript Array for the wrapped array.
func ISimplificationSettingsArrayToJSArray(array []*ISimplificationSettings) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// Distance returns the Distance property of class ISimplificationSettings.
//
// https://doc.babylonjs.com/api/classes/babylon.isimplificationsettings#distance
func (i *ISimplificationSettings) Distance() float64 {
	retVal := i.p.Get("distance")
	return retVal.Float()
}

// SetDistance sets the Distance property of class ISimplificationSettings.
//
// https://doc.babylonjs.com/api/classes/babylon.isimplificationsettings#distance
func (i *ISimplificationSettings) SetDistance(distance float64) *ISimplificationSettings {
	i.p.Set("distance", distance)
	return i
}

// OptimizeMesh returns the OptimizeMesh property of class ISimplificationSettings.
//
// https://doc.babylonjs.com/api/classes/babylon.isimplificationsettings#optimizemesh
func (i *ISimplificationSettings) OptimizeMesh() bool {
	retVal := i.p.Get("optimizeMesh")
	return retVal.Bool()
}

// SetOptimizeMesh sets the OptimizeMesh property of class ISimplificationSettings.
//
// https://doc.babylonjs.com/api/classes/babylon.isimplificationsettings#optimizemesh
func (i *ISimplificationSettings) SetOptimizeMesh(optimizeMesh bool) *ISimplificationSettings {
	i.p.Set("optimizeMesh", optimizeMesh)
	return i
}

// Quality returns the Quality property of class ISimplificationSettings.
//
// https://doc.babylonjs.com/api/classes/babylon.isimplificationsettings#quality
func (i *ISimplificationSettings) Quality() float64 {
	retVal := i.p.Get("quality")
	return retVal.Float()
}

// SetQuality sets the Quality property of class ISimplificationSettings.
//
// https://doc.babylonjs.com/api/classes/babylon.isimplificationsettings#quality
func (i *ISimplificationSettings) SetQuality(quality float64) *ISimplificationSettings {
	i.p.Set("quality", quality)
	return i
}
