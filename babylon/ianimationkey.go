package babylon

import (
	"syscall/js"
)

// IAnimationKey represents a babylon.js IAnimationKey.
type IAnimationKey struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *IAnimationKey) JSObject() js.Value { return a.p }

// IAnimationKey returns a IAnimationKey JavaScript class.
func (ba *Babylon) IAnimationKey() *IAnimationKey {
	p := ba.ctx.Get("IAnimationKey")
	return IAnimationKeyFromJSObject(p, ba.ctx)
}

// IAnimationKeyFromJSObject returns a wrapped IAnimationKey JavaScript class.
func IAnimationKeyFromJSObject(p js.Value, ctx js.Value) *IAnimationKey {
	return &IAnimationKey{p: p, ctx: ctx}
}

// IAnimationKeyArrayToJSArray returns a JavaScript Array for the wrapped array.
func IAnimationKeyArrayToJSArray(array []*IAnimationKey) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}
