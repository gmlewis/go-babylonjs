// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AssetsProgressEvent represents a babylon.js AssetsProgressEvent.
// Class used to share progress information about assets loading
type AssetsProgressEvent struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *AssetsProgressEvent) JSObject() js.Value { return a.p }

// AssetsProgressEvent returns a AssetsProgressEvent JavaScript class.
func (ba *Babylon) AssetsProgressEvent() *AssetsProgressEvent {
	p := ba.ctx.Get("AssetsProgressEvent")
	return AssetsProgressEventFromJSObject(p, ba.ctx)
}

// AssetsProgressEventFromJSObject returns a wrapped AssetsProgressEvent JavaScript class.
func AssetsProgressEventFromJSObject(p js.Value, ctx js.Value) *AssetsProgressEvent {
	return &AssetsProgressEvent{p: p, ctx: ctx}
}

// AssetsProgressEventArrayToJSArray returns a JavaScript Array for the wrapped array.
func AssetsProgressEventArrayToJSArray(array []*AssetsProgressEvent) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewAssetsProgressEvent returns a new AssetsProgressEvent object.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsprogressevent#constructor
func (ba *Babylon) NewAssetsProgressEvent(remainingCount float64, totalCount float64, task *AbstractAssetTask) *AssetsProgressEvent {

	args := make([]interface{}, 0, 3+0)

	args = append(args, remainingCount)
	args = append(args, totalCount)
	args = append(args, task.JSObject())

	p := ba.ctx.Get("AssetsProgressEvent").New(args...)
	return AssetsProgressEventFromJSObject(p, ba.ctx)
}

// RemainingCount returns the RemainingCount property of class AssetsProgressEvent.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsprogressevent#remainingcount
func (a *AssetsProgressEvent) RemainingCount() float64 {
	retVal := a.p.Get("remainingCount")
	return retVal.Float()
}

// SetRemainingCount sets the RemainingCount property of class AssetsProgressEvent.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsprogressevent#remainingcount
func (a *AssetsProgressEvent) SetRemainingCount(remainingCount float64) *AssetsProgressEvent {
	a.p.Set("remainingCount", remainingCount)
	return a
}

// Task returns the Task property of class AssetsProgressEvent.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsprogressevent#task
func (a *AssetsProgressEvent) Task() *AbstractAssetTask {
	retVal := a.p.Get("task")
	return AbstractAssetTaskFromJSObject(retVal, a.ctx)
}

// SetTask sets the Task property of class AssetsProgressEvent.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsprogressevent#task
func (a *AssetsProgressEvent) SetTask(task *AbstractAssetTask) *AssetsProgressEvent {
	a.p.Set("task", task.JSObject())
	return a
}

// TotalCount returns the TotalCount property of class AssetsProgressEvent.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsprogressevent#totalcount
func (a *AssetsProgressEvent) TotalCount() float64 {
	retVal := a.p.Get("totalCount")
	return retVal.Float()
}

// SetTotalCount sets the TotalCount property of class AssetsProgressEvent.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsprogressevent#totalcount
func (a *AssetsProgressEvent) SetTotalCount(totalCount float64) *AssetsProgressEvent {
	a.p.Set("totalCount", totalCount)
	return a
}
