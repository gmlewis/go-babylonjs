// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// WorkerPool represents a babylon.js WorkerPool.
// Helper class to push actions to a pool of workers.
type WorkerPool struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (w *WorkerPool) JSObject() js.Value { return w.p }

// WorkerPool returns a WorkerPool JavaScript class.
func (ba *Babylon) WorkerPool() *WorkerPool {
	p := ba.ctx.Get("WorkerPool")
	return WorkerPoolFromJSObject(p, ba.ctx)
}

// WorkerPoolFromJSObject returns a wrapped WorkerPool JavaScript class.
func WorkerPoolFromJSObject(p js.Value, ctx js.Value) *WorkerPool {
	return &WorkerPool{p: p, ctx: ctx}
}

// WorkerPoolArrayToJSArray returns a JavaScript Array for the wrapped array.
func WorkerPoolArrayToJSArray(array []*WorkerPool) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewWorkerPool returns a new WorkerPool object.
//
// https://doc.babylonjs.com/api/classes/babylon.workerpool
func (ba *Babylon) NewWorkerPool(workers []*Worker) *WorkerPool {

	args := make([]interface{}, 0, 1+0)

	args = append(args, workers)

	p := ba.ctx.Get("WorkerPool").New(args...)
	return WorkerPoolFromJSObject(p, ba.ctx)
}

// Dispose calls the Dispose method on the WorkerPool object.
//
// https://doc.babylonjs.com/api/classes/babylon.workerpool#dispose
func (w *WorkerPool) Dispose() {

	w.p.Call("dispose")
}

// Push calls the Push method on the WorkerPool object.
//
// https://doc.babylonjs.com/api/classes/babylon.workerpool#push
func (w *WorkerPool) Push(action func()) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { action(); return nil }))

	w.p.Call("push", args...)
}

/*

 */
