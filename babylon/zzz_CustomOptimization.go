// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// CustomOptimization represents a babylon.js CustomOptimization.
// Defines an optimization based on user defined callback.
//
// See: http://doc.babylonjs.com/how_to/how_to_use_sceneoptimizer
type CustomOptimization struct {
	*SceneOptimization
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (c *CustomOptimization) JSObject() js.Value { return c.p }

// CustomOptimization returns a CustomOptimization JavaScript class.
func (ba *Babylon) CustomOptimization() *CustomOptimization {
	p := ba.ctx.Get("CustomOptimization")
	return CustomOptimizationFromJSObject(p, ba.ctx)
}

// CustomOptimizationFromJSObject returns a wrapped CustomOptimization JavaScript class.
func CustomOptimizationFromJSObject(p js.Value, ctx js.Value) *CustomOptimization {
	return &CustomOptimization{SceneOptimization: SceneOptimizationFromJSObject(p, ctx), ctx: ctx}
}

// NewCustomOptimizationOpts contains optional parameters for NewCustomOptimization.
type NewCustomOptimizationOpts struct {
	Priority *float64
}

// NewCustomOptimization returns a new CustomOptimization object.
//
// https://doc.babylonjs.com/api/classes/babylon.customoptimization
func (ba *Babylon) NewCustomOptimization(opts *NewCustomOptimizationOpts) *CustomOptimization {
	if opts == nil {
		opts = &NewCustomOptimizationOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Priority == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Priority)
	}

	p := ba.ctx.Get("CustomOptimization").New(args...)
	return CustomOptimizationFromJSObject(p, ba.ctx)
}

// Apply calls the Apply method on the CustomOptimization object.
//
// https://doc.babylonjs.com/api/classes/babylon.customoptimization#apply
func (c *CustomOptimization) Apply(scene *Scene, optimizer *SceneOptimizer) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, scene.JSObject())
	args = append(args, optimizer.JSObject())

	retVal := c.p.Call("apply", args...)
	return retVal.Bool()
}

// GetDescription calls the GetDescription method on the CustomOptimization object.
//
// https://doc.babylonjs.com/api/classes/babylon.customoptimization#getdescription
func (c *CustomOptimization) GetDescription() string {

	args := make([]interface{}, 0, 0+0)

	retVal := c.p.Call("getDescription", args...)
	return retVal.String()
}

/*

// OnApply returns the OnApply property of class CustomOptimization.
//
// https://doc.babylonjs.com/api/classes/babylon.customoptimization#onapply
func (c *CustomOptimization) OnApply(onApply func()) *CustomOptimization {
	p := ba.ctx.Get("CustomOptimization").New(onApply)
	return CustomOptimizationFromJSObject(p, ba.ctx)
}

// SetOnApply sets the OnApply property of class CustomOptimization.
//
// https://doc.babylonjs.com/api/classes/babylon.customoptimization#onapply
func (c *CustomOptimization) SetOnApply(onApply func()) *CustomOptimization {
	p := ba.ctx.Get("CustomOptimization").New(onApply)
	return CustomOptimizationFromJSObject(p, ba.ctx)
}

// OnGetDescription returns the OnGetDescription property of class CustomOptimization.
//
// https://doc.babylonjs.com/api/classes/babylon.customoptimization#ongetdescription
func (c *CustomOptimization) OnGetDescription(onGetDescription func()) *CustomOptimization {
	p := ba.ctx.Get("CustomOptimization").New(onGetDescription)
	return CustomOptimizationFromJSObject(p, ba.ctx)
}

// SetOnGetDescription sets the OnGetDescription property of class CustomOptimization.
//
// https://doc.babylonjs.com/api/classes/babylon.customoptimization#ongetdescription
func (c *CustomOptimization) SetOnGetDescription(onGetDescription func()) *CustomOptimization {
	p := ba.ctx.Get("CustomOptimization").New(onGetDescription)
	return CustomOptimizationFromJSObject(p, ba.ctx)
}

// Priority returns the Priority property of class CustomOptimization.
//
// https://doc.babylonjs.com/api/classes/babylon.customoptimization#priority
func (c *CustomOptimization) Priority(priority float64) *CustomOptimization {
	p := ba.ctx.Get("CustomOptimization").New(priority)
	return CustomOptimizationFromJSObject(p, ba.ctx)
}

// SetPriority sets the Priority property of class CustomOptimization.
//
// https://doc.babylonjs.com/api/classes/babylon.customoptimization#priority
func (c *CustomOptimization) SetPriority(priority float64) *CustomOptimization {
	p := ba.ctx.Get("CustomOptimization").New(priority)
	return CustomOptimizationFromJSObject(p, ba.ctx)
}

*/
