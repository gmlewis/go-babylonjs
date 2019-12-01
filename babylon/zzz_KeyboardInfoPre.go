// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// KeyboardInfoPre represents a babylon.js KeyboardInfoPre.
// This class is used to store keyboard related info for the onPreKeyboardObservable event.
// Set the skipOnKeyboardObservable property to true if you want the engine to stop any process after this event is triggered, even not calling onKeyboardObservable
type KeyboardInfoPre struct {
	*KeyboardInfo
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (k *KeyboardInfoPre) JSObject() js.Value { return k.p }

// KeyboardInfoPre returns a KeyboardInfoPre JavaScript class.
func (ba *Babylon) KeyboardInfoPre() *KeyboardInfoPre {
	p := ba.ctx.Get("KeyboardInfoPre")
	return KeyboardInfoPreFromJSObject(p, ba.ctx)
}

// KeyboardInfoPreFromJSObject returns a wrapped KeyboardInfoPre JavaScript class.
func KeyboardInfoPreFromJSObject(p js.Value, ctx js.Value) *KeyboardInfoPre {
	return &KeyboardInfoPre{KeyboardInfo: KeyboardInfoFromJSObject(p, ctx), ctx: ctx}
}

// KeyboardInfoPreArrayToJSArray returns a JavaScript Array for the wrapped array.
func KeyboardInfoPreArrayToJSArray(array []*KeyboardInfoPre) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewKeyboardInfoPre returns a new KeyboardInfoPre object.
//
// https://doc.babylonjs.com/api/classes/babylon.keyboardinfopre
func (ba *Babylon) NewKeyboardInfoPre(jsType float64, event js.Value) *KeyboardInfoPre {

	args := make([]interface{}, 0, 2+0)

	args = append(args, jsType)
	args = append(args, event)

	p := ba.ctx.Get("KeyboardInfoPre").New(args...)
	return KeyboardInfoPreFromJSObject(p, ba.ctx)
}

/*

// Event returns the Event property of class KeyboardInfoPre.
//
// https://doc.babylonjs.com/api/classes/babylon.keyboardinfopre#event
func (k *KeyboardInfoPre) Event(event js.Value) *KeyboardInfoPre {
	p := ba.ctx.Get("KeyboardInfoPre").New(event)
	return KeyboardInfoPreFromJSObject(p, ba.ctx)
}

// SetEvent sets the Event property of class KeyboardInfoPre.
//
// https://doc.babylonjs.com/api/classes/babylon.keyboardinfopre#event
func (k *KeyboardInfoPre) SetEvent(event js.Value) *KeyboardInfoPre {
	p := ba.ctx.Get("KeyboardInfoPre").New(event)
	return KeyboardInfoPreFromJSObject(p, ba.ctx)
}

// SkipOnPointerObservable returns the SkipOnPointerObservable property of class KeyboardInfoPre.
//
// https://doc.babylonjs.com/api/classes/babylon.keyboardinfopre#skiponpointerobservable
func (k *KeyboardInfoPre) SkipOnPointerObservable(skipOnPointerObservable bool) *KeyboardInfoPre {
	p := ba.ctx.Get("KeyboardInfoPre").New(skipOnPointerObservable)
	return KeyboardInfoPreFromJSObject(p, ba.ctx)
}

// SetSkipOnPointerObservable sets the SkipOnPointerObservable property of class KeyboardInfoPre.
//
// https://doc.babylonjs.com/api/classes/babylon.keyboardinfopre#skiponpointerobservable
func (k *KeyboardInfoPre) SetSkipOnPointerObservable(skipOnPointerObservable bool) *KeyboardInfoPre {
	p := ba.ctx.Get("KeyboardInfoPre").New(skipOnPointerObservable)
	return KeyboardInfoPreFromJSObject(p, ba.ctx)
}

// Type returns the Type property of class KeyboardInfoPre.
//
// https://doc.babylonjs.com/api/classes/babylon.keyboardinfopre#type
func (k *KeyboardInfoPre) Type(jsType float64) *KeyboardInfoPre {
	p := ba.ctx.Get("KeyboardInfoPre").New(jsType)
	return KeyboardInfoPreFromJSObject(p, ba.ctx)
}

// SetType sets the Type property of class KeyboardInfoPre.
//
// https://doc.babylonjs.com/api/classes/babylon.keyboardinfopre#type
func (k *KeyboardInfoPre) SetType(jsType float64) *KeyboardInfoPre {
	p := ba.ctx.Get("KeyboardInfoPre").New(jsType)
	return KeyboardInfoPreFromJSObject(p, ba.ctx)
}

*/
