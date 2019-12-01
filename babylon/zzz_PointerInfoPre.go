// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PointerInfoPre represents a babylon.js PointerInfoPre.
// This class is used to store pointer related info for the onPrePointerObservable event.
// Set the skipOnPointerObservable property to true if you want the engine to stop any process after this event is triggered, even not calling onPointerObservable
type PointerInfoPre struct {
	*PointerInfoBase
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PointerInfoPre) JSObject() js.Value { return p.p }

// PointerInfoPre returns a PointerInfoPre JavaScript class.
func (ba *Babylon) PointerInfoPre() *PointerInfoPre {
	p := ba.ctx.Get("PointerInfoPre")
	return PointerInfoPreFromJSObject(p, ba.ctx)
}

// PointerInfoPreFromJSObject returns a wrapped PointerInfoPre JavaScript class.
func PointerInfoPreFromJSObject(p js.Value, ctx js.Value) *PointerInfoPre {
	return &PointerInfoPre{PointerInfoBase: PointerInfoBaseFromJSObject(p, ctx), ctx: ctx}
}

// PointerInfoPreArrayToJSArray returns a JavaScript Array for the wrapped array.
func PointerInfoPreArrayToJSArray(array []*PointerInfoPre) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewPointerInfoPre returns a new PointerInfoPre object.
//
// https://doc.babylonjs.com/api/classes/babylon.pointerinfopre
func (ba *Babylon) NewPointerInfoPre(jsType float64, event js.Value, localX float64, localY float64) *PointerInfoPre {

	args := make([]interface{}, 0, 4+0)

	args = append(args, jsType)
	args = append(args, event)
	args = append(args, localX)
	args = append(args, localY)

	p := ba.ctx.Get("PointerInfoPre").New(args...)
	return PointerInfoPreFromJSObject(p, ba.ctx)
}

/*

// Event returns the Event property of class PointerInfoPre.
//
// https://doc.babylonjs.com/api/classes/babylon.pointerinfopre#event
func (p *PointerInfoPre) Event(event js.Value) *PointerInfoPre {
	p := ba.ctx.Get("PointerInfoPre").New(event)
	return PointerInfoPreFromJSObject(p, ba.ctx)
}

// SetEvent sets the Event property of class PointerInfoPre.
//
// https://doc.babylonjs.com/api/classes/babylon.pointerinfopre#event
func (p *PointerInfoPre) SetEvent(event js.Value) *PointerInfoPre {
	p := ba.ctx.Get("PointerInfoPre").New(event)
	return PointerInfoPreFromJSObject(p, ba.ctx)
}

// LocalPosition returns the LocalPosition property of class PointerInfoPre.
//
// https://doc.babylonjs.com/api/classes/babylon.pointerinfopre#localposition
func (p *PointerInfoPre) LocalPosition(localPosition *Vector2) *PointerInfoPre {
	p := ba.ctx.Get("PointerInfoPre").New(localPosition.JSObject())
	return PointerInfoPreFromJSObject(p, ba.ctx)
}

// SetLocalPosition sets the LocalPosition property of class PointerInfoPre.
//
// https://doc.babylonjs.com/api/classes/babylon.pointerinfopre#localposition
func (p *PointerInfoPre) SetLocalPosition(localPosition *Vector2) *PointerInfoPre {
	p := ba.ctx.Get("PointerInfoPre").New(localPosition.JSObject())
	return PointerInfoPreFromJSObject(p, ba.ctx)
}

// Ray returns the Ray property of class PointerInfoPre.
//
// https://doc.babylonjs.com/api/classes/babylon.pointerinfopre#ray
func (p *PointerInfoPre) Ray(ray *Ray) *PointerInfoPre {
	p := ba.ctx.Get("PointerInfoPre").New(ray.JSObject())
	return PointerInfoPreFromJSObject(p, ba.ctx)
}

// SetRay sets the Ray property of class PointerInfoPre.
//
// https://doc.babylonjs.com/api/classes/babylon.pointerinfopre#ray
func (p *PointerInfoPre) SetRay(ray *Ray) *PointerInfoPre {
	p := ba.ctx.Get("PointerInfoPre").New(ray.JSObject())
	return PointerInfoPreFromJSObject(p, ba.ctx)
}

// SkipOnPointerObservable returns the SkipOnPointerObservable property of class PointerInfoPre.
//
// https://doc.babylonjs.com/api/classes/babylon.pointerinfopre#skiponpointerobservable
func (p *PointerInfoPre) SkipOnPointerObservable(skipOnPointerObservable bool) *PointerInfoPre {
	p := ba.ctx.Get("PointerInfoPre").New(skipOnPointerObservable)
	return PointerInfoPreFromJSObject(p, ba.ctx)
}

// SetSkipOnPointerObservable sets the SkipOnPointerObservable property of class PointerInfoPre.
//
// https://doc.babylonjs.com/api/classes/babylon.pointerinfopre#skiponpointerobservable
func (p *PointerInfoPre) SetSkipOnPointerObservable(skipOnPointerObservable bool) *PointerInfoPre {
	p := ba.ctx.Get("PointerInfoPre").New(skipOnPointerObservable)
	return PointerInfoPreFromJSObject(p, ba.ctx)
}

// Type returns the Type property of class PointerInfoPre.
//
// https://doc.babylonjs.com/api/classes/babylon.pointerinfopre#type
func (p *PointerInfoPre) Type(jsType float64) *PointerInfoPre {
	p := ba.ctx.Get("PointerInfoPre").New(jsType)
	return PointerInfoPreFromJSObject(p, ba.ctx)
}

// SetType sets the Type property of class PointerInfoPre.
//
// https://doc.babylonjs.com/api/classes/babylon.pointerinfopre#type
func (p *PointerInfoPre) SetType(jsType float64) *PointerInfoPre {
	p := ba.ctx.Get("PointerInfoPre").New(jsType)
	return PointerInfoPreFromJSObject(p, ba.ctx)
}

*/
