// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// IEffectFallbacks represents a babylon.js IEffectFallbacks.
// Interface used to define common properties for effect fallbacks
type IEffectFallbacks struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *IEffectFallbacks) JSObject() js.Value { return i.p }

// IEffectFallbacks returns a IEffectFallbacks JavaScript class.
func (ba *Babylon) IEffectFallbacks() *IEffectFallbacks {
	p := ba.ctx.Get("IEffectFallbacks")
	return IEffectFallbacksFromJSObject(p, ba.ctx)
}

// IEffectFallbacksFromJSObject returns a wrapped IEffectFallbacks JavaScript class.
func IEffectFallbacksFromJSObject(p js.Value, ctx js.Value) *IEffectFallbacks {
	return &IEffectFallbacks{p: p, ctx: ctx}
}

// IEffectFallbacksArrayToJSArray returns a JavaScript Array for the wrapped array.
func IEffectFallbacksArrayToJSArray(array []*IEffectFallbacks) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// Reduce calls the Reduce method on the IEffectFallbacks object.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectfallbacks#reduce
func (i *IEffectFallbacks) Reduce(currentDefines string, effect *Effect) string {

	args := make([]interface{}, 0, 2+0)

	args = append(args, currentDefines)

	if effect == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, effect.JSObject())
	}

	retVal := i.p.Call("reduce", args...)
	return retVal.String()
}

// UnBindMesh calls the UnBindMesh method on the IEffectFallbacks object.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectfallbacks#unbindmesh
func (i *IEffectFallbacks) UnBindMesh() {

	i.p.Call("unBindMesh")
}

// HasMoreFallbacks returns the HasMoreFallbacks property of class IEffectFallbacks.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectfallbacks#hasmorefallbacks
func (i *IEffectFallbacks) HasMoreFallbacks() bool {
	retVal := i.p.Get("hasMoreFallbacks")
	return retVal.Bool()
}

// SetHasMoreFallbacks sets the HasMoreFallbacks property of class IEffectFallbacks.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectfallbacks#hasmorefallbacks
func (i *IEffectFallbacks) SetHasMoreFallbacks(hasMoreFallbacks bool) *IEffectFallbacks {
	i.p.Set("hasMoreFallbacks", hasMoreFallbacks)
	return i
}
