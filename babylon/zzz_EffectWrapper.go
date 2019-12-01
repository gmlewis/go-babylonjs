// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// EffectWrapper represents a babylon.js EffectWrapper.
// Wraps an effect to be used for rendering
type EffectWrapper struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (e *EffectWrapper) JSObject() js.Value { return e.p }

// EffectWrapper returns a EffectWrapper JavaScript class.
func (ba *Babylon) EffectWrapper() *EffectWrapper {
	p := ba.ctx.Get("EffectWrapper")
	return EffectWrapperFromJSObject(p, ba.ctx)
}

// EffectWrapperFromJSObject returns a wrapped EffectWrapper JavaScript class.
func EffectWrapperFromJSObject(p js.Value, ctx js.Value) *EffectWrapper {
	return &EffectWrapper{p: p, ctx: ctx}
}

// NewEffectWrapper returns a new EffectWrapper object.
//
// https://doc.babylonjs.com/api/classes/babylon.effectwrapper
func (ba *Babylon) NewEffectWrapper(creationOptions *EffectWrapperCreationOptions) *EffectWrapper {

	args := make([]interface{}, 0, 1+0)

	args = append(args, creationOptions.JSObject())

	p := ba.ctx.Get("EffectWrapper").New(args...)
	return EffectWrapperFromJSObject(p, ba.ctx)
}

// Dispose calls the Dispose method on the EffectWrapper object.
//
// https://doc.babylonjs.com/api/classes/babylon.effectwrapper#dispose
func (e *EffectWrapper) Dispose() {

	args := make([]interface{}, 0, 0+0)

	e.p.Call("dispose", args...)
}

/*

// Effect returns the Effect property of class EffectWrapper.
//
// https://doc.babylonjs.com/api/classes/babylon.effectwrapper#effect
func (e *EffectWrapper) Effect(effect *Effect) *EffectWrapper {
	p := ba.ctx.Get("EffectWrapper").New(effect.JSObject())
	return EffectWrapperFromJSObject(p, ba.ctx)
}

// SetEffect sets the Effect property of class EffectWrapper.
//
// https://doc.babylonjs.com/api/classes/babylon.effectwrapper#effect
func (e *EffectWrapper) SetEffect(effect *Effect) *EffectWrapper {
	p := ba.ctx.Get("EffectWrapper").New(effect.JSObject())
	return EffectWrapperFromJSObject(p, ba.ctx)
}

// OnApplyObservable returns the OnApplyObservable property of class EffectWrapper.
//
// https://doc.babylonjs.com/api/classes/babylon.effectwrapper#onapplyobservable
func (e *EffectWrapper) OnApplyObservable(onApplyObservable *Observable) *EffectWrapper {
	p := ba.ctx.Get("EffectWrapper").New(onApplyObservable.JSObject())
	return EffectWrapperFromJSObject(p, ba.ctx)
}

// SetOnApplyObservable sets the OnApplyObservable property of class EffectWrapper.
//
// https://doc.babylonjs.com/api/classes/babylon.effectwrapper#onapplyobservable
func (e *EffectWrapper) SetOnApplyObservable(onApplyObservable *Observable) *EffectWrapper {
	p := ba.ctx.Get("EffectWrapper").New(onApplyObservable.JSObject())
	return EffectWrapperFromJSObject(p, ba.ctx)
}

*/
