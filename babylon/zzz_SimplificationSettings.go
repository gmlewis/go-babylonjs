// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SimplificationSettings represents a babylon.js SimplificationSettings.
// Class used to specify simplification options
//
// See: http://doc.babylonjs.com/how_to/in-browser_mesh_simplification
type SimplificationSettings struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *SimplificationSettings) JSObject() js.Value { return s.p }

// SimplificationSettings returns a SimplificationSettings JavaScript class.
func (ba *Babylon) SimplificationSettings() *SimplificationSettings {
	p := ba.ctx.Get("SimplificationSettings")
	return SimplificationSettingsFromJSObject(p, ba.ctx)
}

// SimplificationSettingsFromJSObject returns a wrapped SimplificationSettings JavaScript class.
func SimplificationSettingsFromJSObject(p js.Value, ctx js.Value) *SimplificationSettings {
	return &SimplificationSettings{p: p, ctx: ctx}
}

// SimplificationSettingsArrayToJSArray returns a JavaScript Array for the wrapped array.
func SimplificationSettingsArrayToJSArray(array []*SimplificationSettings) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewSimplificationSettingsOpts contains optional parameters for NewSimplificationSettings.
type NewSimplificationSettingsOpts struct {
	OptimizeMesh *bool
}

// NewSimplificationSettings returns a new SimplificationSettings object.
//
// https://doc.babylonjs.com/api/classes/babylon.simplificationsettings
func (ba *Babylon) NewSimplificationSettings(quality float64, distance float64, opts *NewSimplificationSettingsOpts) *SimplificationSettings {
	if opts == nil {
		opts = &NewSimplificationSettingsOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, quality)
	args = append(args, distance)

	if opts.OptimizeMesh == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.OptimizeMesh)
	}

	p := ba.ctx.Get("SimplificationSettings").New(args...)
	return SimplificationSettingsFromJSObject(p, ba.ctx)
}

/*

// Distance returns the Distance property of class SimplificationSettings.
//
// https://doc.babylonjs.com/api/classes/babylon.simplificationsettings#distance
func (s *SimplificationSettings) Distance(distance float64) *SimplificationSettings {
	p := ba.ctx.Get("SimplificationSettings").New(distance)
	return SimplificationSettingsFromJSObject(p, ba.ctx)
}

// SetDistance sets the Distance property of class SimplificationSettings.
//
// https://doc.babylonjs.com/api/classes/babylon.simplificationsettings#distance
func (s *SimplificationSettings) SetDistance(distance float64) *SimplificationSettings {
	p := ba.ctx.Get("SimplificationSettings").New(distance)
	return SimplificationSettingsFromJSObject(p, ba.ctx)
}

// OptimizeMesh returns the OptimizeMesh property of class SimplificationSettings.
//
// https://doc.babylonjs.com/api/classes/babylon.simplificationsettings#optimizemesh
func (s *SimplificationSettings) OptimizeMesh(optimizeMesh bool) *SimplificationSettings {
	p := ba.ctx.Get("SimplificationSettings").New(optimizeMesh)
	return SimplificationSettingsFromJSObject(p, ba.ctx)
}

// SetOptimizeMesh sets the OptimizeMesh property of class SimplificationSettings.
//
// https://doc.babylonjs.com/api/classes/babylon.simplificationsettings#optimizemesh
func (s *SimplificationSettings) SetOptimizeMesh(optimizeMesh bool) *SimplificationSettings {
	p := ba.ctx.Get("SimplificationSettings").New(optimizeMesh)
	return SimplificationSettingsFromJSObject(p, ba.ctx)
}

// Quality returns the Quality property of class SimplificationSettings.
//
// https://doc.babylonjs.com/api/classes/babylon.simplificationsettings#quality
func (s *SimplificationSettings) Quality(quality float64) *SimplificationSettings {
	p := ba.ctx.Get("SimplificationSettings").New(quality)
	return SimplificationSettingsFromJSObject(p, ba.ctx)
}

// SetQuality sets the Quality property of class SimplificationSettings.
//
// https://doc.babylonjs.com/api/classes/babylon.simplificationsettings#quality
func (s *SimplificationSettings) SetQuality(quality float64) *SimplificationSettings {
	p := ba.ctx.Get("SimplificationSettings").New(quality)
	return SimplificationSettingsFromJSObject(p, ba.ctx)
}

*/
