// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AsyncLoop represents a babylon.js AsyncLoop.
// An implementation of a loop for asynchronous functions.
type AsyncLoop struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *AsyncLoop) JSObject() js.Value { return a.p }

// AsyncLoop returns a AsyncLoop JavaScript class.
func (ba *Babylon) AsyncLoop() *AsyncLoop {
	p := ba.ctx.Get("AsyncLoop")
	return AsyncLoopFromJSObject(p, ba.ctx)
}

// AsyncLoopFromJSObject returns a wrapped AsyncLoop JavaScript class.
func AsyncLoopFromJSObject(p js.Value, ctx js.Value) *AsyncLoop {
	return &AsyncLoop{p: p, ctx: ctx}
}

// AsyncLoopArrayToJSArray returns a JavaScript Array for the wrapped array.
func AsyncLoopArrayToJSArray(array []*AsyncLoop) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewAsyncLoopOpts contains optional parameters for NewAsyncLoop.
type NewAsyncLoopOpts struct {
	Offset *float64
}

// NewAsyncLoop returns a new AsyncLoop object.
//
// https://doc.babylonjs.com/api/classes/babylon.asyncloop
func (ba *Babylon) NewAsyncLoop(iterations float64, jsFunc func(), successCallback func(), opts *NewAsyncLoopOpts) *AsyncLoop {
	if opts == nil {
		opts = &NewAsyncLoopOpts{}
	}

	args := make([]interface{}, 0, 3+1)

	args = append(args, iterations)
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { jsFunc(); return nil }))
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { successCallback(); return nil }))

	if opts.Offset == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Offset)
	}

	p := ba.ctx.Get("AsyncLoop").New(args...)
	return AsyncLoopFromJSObject(p, ba.ctx)
}

// BreakLoop calls the BreakLoop method on the AsyncLoop object.
//
// https://doc.babylonjs.com/api/classes/babylon.asyncloop#breakloop
func (a *AsyncLoop) BreakLoop() {

	a.p.Call("breakLoop")
}

// ExecuteNext calls the ExecuteNext method on the AsyncLoop object.
//
// https://doc.babylonjs.com/api/classes/babylon.asyncloop#executenext
func (a *AsyncLoop) ExecuteNext() {

	a.p.Call("executeNext")
}

// AsyncLoopRunOpts contains optional parameters for AsyncLoop.Run.
type AsyncLoopRunOpts struct {
	Offset *float64
}

// Run calls the Run method on the AsyncLoop object.
//
// https://doc.babylonjs.com/api/classes/babylon.asyncloop#run
func (a *AsyncLoop) Run(iterations float64, fn func(), successCallback func(), opts *AsyncLoopRunOpts) *AsyncLoop {
	if opts == nil {
		opts = &AsyncLoopRunOpts{}
	}

	args := make([]interface{}, 0, 3+1)

	args = append(args, iterations)
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { fn(); return nil }))
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { successCallback(); return nil }))

	if opts.Offset == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Offset)
	}

	retVal := a.p.Call("Run", args...)
	return AsyncLoopFromJSObject(retVal, a.ctx)
}

// AsyncLoopSyncAsyncForLoopOpts contains optional parameters for AsyncLoop.SyncAsyncForLoop.
type AsyncLoopSyncAsyncForLoopOpts struct {
	BreakFunction func()
	Timeout       *float64
}

// SyncAsyncForLoop calls the SyncAsyncForLoop method on the AsyncLoop object.
//
// https://doc.babylonjs.com/api/classes/babylon.asyncloop#syncasyncforloop
func (a *AsyncLoop) SyncAsyncForLoop(iterations float64, syncedIterations float64, fn func(), callback func(), opts *AsyncLoopSyncAsyncForLoopOpts) *AsyncLoop {
	if opts == nil {
		opts = &AsyncLoopSyncAsyncForLoopOpts{}
	}

	args := make([]interface{}, 0, 4+2)

	args = append(args, iterations)
	args = append(args, syncedIterations)
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { fn(); return nil }))
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { callback(); return nil }))

	if opts.BreakFunction == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { opts.BreakFunction(); return nil }) /* never freed! */)
	}
	if opts.Timeout == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Timeout)
	}

	retVal := a.p.Call("SyncAsyncForLoop", args...)
	return AsyncLoopFromJSObject(retVal, a.ctx)
}

// Index returns the Index property of class AsyncLoop.
//
// https://doc.babylonjs.com/api/classes/babylon.asyncloop#index
func (a *AsyncLoop) Index() float64 {
	retVal := a.p.Get("index")
	return retVal.Float()
}

// SetIndex sets the Index property of class AsyncLoop.
//
// https://doc.babylonjs.com/api/classes/babylon.asyncloop#index
func (a *AsyncLoop) SetIndex(index float64) *AsyncLoop {
	a.p.Set("index", index)
	return a
}

// Iterations returns the Iterations property of class AsyncLoop.
//
// https://doc.babylonjs.com/api/classes/babylon.asyncloop#iterations
func (a *AsyncLoop) Iterations() float64 {
	retVal := a.p.Get("iterations")
	return retVal.Float()
}

// SetIterations sets the Iterations property of class AsyncLoop.
//
// https://doc.babylonjs.com/api/classes/babylon.asyncloop#iterations
func (a *AsyncLoop) SetIterations(iterations float64) *AsyncLoop {
	a.p.Set("iterations", iterations)
	return a
}
