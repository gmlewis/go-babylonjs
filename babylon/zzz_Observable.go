// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Observable represents a babylon.js Observable.
// The Observable class is a simple implementation of the Observable pattern.
//
// There&#39;s one slight particularity though: a given Observable can notify its observer using a particular mask value, only the Observers registered with this mask value will be notified.
// This enable a more fine grained execution without having to rely on multiple different Observable objects.
// For instance you may have a given Observable that have four different types of notifications: Move (mask = 0x01), Stop (mask = 0x02), Turn Right (mask = 0X04), Turn Left (mask = 0X08).
// A given observer can register itself with only Move and Stop (mask = 0x03), then it will only be notified when one of these two occurs and will never be for Turn Left/Right.
type Observable struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (o *Observable) JSObject() js.Value { return o.p }

// Observable returns a Observable JavaScript class.
func (ba *Babylon) Observable() *Observable {
	p := ba.ctx.Get("Observable")
	return ObservableFromJSObject(p, ba.ctx)
}

// ObservableFromJSObject returns a wrapped Observable JavaScript class.
func ObservableFromJSObject(p js.Value, ctx js.Value) *Observable {
	return &Observable{p: p, ctx: ctx}
}

// ObservableArrayToJSArray returns a JavaScript Array for the wrapped array.
func ObservableArrayToJSArray(array []*Observable) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewObservableOpts contains optional parameters for NewObservable.
type NewObservableOpts struct {
	OnObserverAdded func()
}

// NewObservable returns a new Observable object.
//
// https://doc.babylonjs.com/api/classes/babylon.observable
func (ba *Babylon) NewObservable(opts *NewObservableOpts) *Observable {
	if opts == nil {
		opts = &NewObservableOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.OnObserverAdded == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { opts.OnObserverAdded(); return nil }) /* never freed! */)
	}

	p := ba.ctx.Get("Observable").New(args...)
	return ObservableFromJSObject(p, ba.ctx)
}

// ObservableAddOpts contains optional parameters for Observable.Add.
type ObservableAddOpts struct {
	Mask                  *float64
	InsertFirst           *bool
	Scope                 *interface{}
	UnregisterOnFirstCall *bool
}

// Add calls the Add method on the Observable object.
//
// https://doc.babylonjs.com/api/classes/babylon.observable#add
func (o *Observable) Add(callback func(), opts *ObservableAddOpts) *Observer {
	if opts == nil {
		opts = &ObservableAddOpts{}
	}

	args := make([]interface{}, 0, 1+4)

	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { callback(); return nil }))

	if opts.Mask == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Mask)
	}
	if opts.InsertFirst == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.InsertFirst)
	}
	if opts.Scope == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scope)
	}
	if opts.UnregisterOnFirstCall == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UnregisterOnFirstCall)
	}

	retVal := o.p.Call("add", args...)
	return ObserverFromJSObject(retVal, o.ctx)
}

// AddOnce calls the AddOnce method on the Observable object.
//
// https://doc.babylonjs.com/api/classes/babylon.observable#addonce
func (o *Observable) AddOnce(callback func()) *Observer {

	args := make([]interface{}, 0, 1+0)

	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { callback(); return nil }))

	retVal := o.p.Call("addOnce", args...)
	return ObserverFromJSObject(retVal, o.ctx)
}

// Clear calls the Clear method on the Observable object.
//
// https://doc.babylonjs.com/api/classes/babylon.observable#clear
func (o *Observable) Clear() {

	o.p.Call("clear")
}

// Clone calls the Clone method on the Observable object.
//
// https://doc.babylonjs.com/api/classes/babylon.observable#clone
func (o *Observable) Clone() *Observable {

	retVal := o.p.Call("clone")
	return ObservableFromJSObject(retVal, o.ctx)
}

// HasObservers calls the HasObservers method on the Observable object.
//
// https://doc.babylonjs.com/api/classes/babylon.observable#hasobservers
func (o *Observable) HasObservers() bool {

	retVal := o.p.Call("hasObservers")
	return retVal.Bool()
}

// ObservableHasSpecificMaskOpts contains optional parameters for Observable.HasSpecificMask.
type ObservableHasSpecificMaskOpts struct {
	Mask *float64
}

// HasSpecificMask calls the HasSpecificMask method on the Observable object.
//
// https://doc.babylonjs.com/api/classes/babylon.observable#hasspecificmask
func (o *Observable) HasSpecificMask(opts *ObservableHasSpecificMaskOpts) bool {
	if opts == nil {
		opts = &ObservableHasSpecificMaskOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Mask == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Mask)
	}

	retVal := o.p.Call("hasSpecificMask", args...)
	return retVal.Bool()
}

// MakeObserverBottomPriority calls the MakeObserverBottomPriority method on the Observable object.
//
// https://doc.babylonjs.com/api/classes/babylon.observable#makeobserverbottompriority
func (o *Observable) MakeObserverBottomPriority(observer *Observer) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, observer.JSObject())

	o.p.Call("makeObserverBottomPriority", args...)
}

// MakeObserverTopPriority calls the MakeObserverTopPriority method on the Observable object.
//
// https://doc.babylonjs.com/api/classes/babylon.observable#makeobservertoppriority
func (o *Observable) MakeObserverTopPriority(observer *Observer) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, observer.JSObject())

	o.p.Call("makeObserverTopPriority", args...)
}

// ObservableNotifyObserverOpts contains optional parameters for Observable.NotifyObserver.
type ObservableNotifyObserverOpts struct {
	Mask *float64
}

// NotifyObserver calls the NotifyObserver method on the Observable object.
//
// https://doc.babylonjs.com/api/classes/babylon.observable#notifyobserver
func (o *Observable) NotifyObserver(observer *Observer, eventData *T, opts *ObservableNotifyObserverOpts) {
	if opts == nil {
		opts = &ObservableNotifyObserverOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, observer.JSObject())
	args = append(args, eventData.JSObject())

	if opts.Mask == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Mask)
	}

	o.p.Call("notifyObserver", args...)
}

// ObservableNotifyObserversOpts contains optional parameters for Observable.NotifyObservers.
type ObservableNotifyObserversOpts struct {
	Mask          *float64
	Target        *interface{}
	CurrentTarget *interface{}
}

// NotifyObservers calls the NotifyObservers method on the Observable object.
//
// https://doc.babylonjs.com/api/classes/babylon.observable#notifyobservers
func (o *Observable) NotifyObservers(eventData *T, opts *ObservableNotifyObserversOpts) bool {
	if opts == nil {
		opts = &ObservableNotifyObserversOpts{}
	}

	args := make([]interface{}, 0, 1+3)

	args = append(args, eventData.JSObject())

	if opts.Mask == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Mask)
	}
	if opts.Target == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Target)
	}
	if opts.CurrentTarget == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.CurrentTarget)
	}

	retVal := o.p.Call("notifyObservers", args...)
	return retVal.Bool()
}

// ObservableNotifyObserversWithPromiseOpts contains optional parameters for Observable.NotifyObserversWithPromise.
type ObservableNotifyObserversWithPromiseOpts struct {
	Mask          *float64
	Target        *interface{}
	CurrentTarget *interface{}
}

// NotifyObserversWithPromise calls the NotifyObserversWithPromise method on the Observable object.
//
// https://doc.babylonjs.com/api/classes/babylon.observable#notifyobserverswithpromise
func (o *Observable) NotifyObserversWithPromise(eventData *T, opts *ObservableNotifyObserversWithPromiseOpts) *Promise {
	if opts == nil {
		opts = &ObservableNotifyObserversWithPromiseOpts{}
	}

	args := make([]interface{}, 0, 1+3)

	args = append(args, eventData.JSObject())

	if opts.Mask == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Mask)
	}
	if opts.Target == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Target)
	}
	if opts.CurrentTarget == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.CurrentTarget)
	}

	retVal := o.p.Call("notifyObserversWithPromise", args...)
	return PromiseFromJSObject(retVal, o.ctx)
}

// Remove calls the Remove method on the Observable object.
//
// https://doc.babylonjs.com/api/classes/babylon.observable#remove
func (o *Observable) Remove(observer *Observer) bool {

	args := make([]interface{}, 0, 1+0)

	args = append(args, observer.JSObject())

	retVal := o.p.Call("remove", args...)
	return retVal.Bool()
}

// ObservableRemoveCallbackOpts contains optional parameters for Observable.RemoveCallback.
type ObservableRemoveCallbackOpts struct {
	Scope *interface{}
}

// RemoveCallback calls the RemoveCallback method on the Observable object.
//
// https://doc.babylonjs.com/api/classes/babylon.observable#removecallback
func (o *Observable) RemoveCallback(callback func(), opts *ObservableRemoveCallbackOpts) bool {
	if opts == nil {
		opts = &ObservableRemoveCallbackOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { callback(); return nil }))

	if opts.Scope == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scope)
	}

	retVal := o.p.Call("removeCallback", args...)
	return retVal.Bool()
}

// Observers returns the Observers property of class Observable.
//
// https://doc.babylonjs.com/api/classes/babylon.observable#observers
func (o *Observable) Observers() []*Observer {
	retVal := o.p.Get("observers")
	result := []*Observer{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, ObserverFromJSObject(retVal.Index(ri), o.ctx))
	}
	return result
}

// SetObservers sets the Observers property of class Observable.
//
// https://doc.babylonjs.com/api/classes/babylon.observable#observers
func (o *Observable) SetObservers(observers []*Observer) *Observable {
	o.p.Set("observers", observers)
	return o
}
