// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// IDisplayChangedEventArgs represents a babylon.js IDisplayChangedEventArgs.
// Defines the interface used by display changed events
type IDisplayChangedEventArgs struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *IDisplayChangedEventArgs) JSObject() js.Value { return i.p }

// IDisplayChangedEventArgs returns a IDisplayChangedEventArgs JavaScript class.
func (ba *Babylon) IDisplayChangedEventArgs() *IDisplayChangedEventArgs {
	p := ba.ctx.Get("IDisplayChangedEventArgs")
	return IDisplayChangedEventArgsFromJSObject(p, ba.ctx)
}

// IDisplayChangedEventArgsFromJSObject returns a wrapped IDisplayChangedEventArgs JavaScript class.
func IDisplayChangedEventArgsFromJSObject(p js.Value, ctx js.Value) *IDisplayChangedEventArgs {
	return &IDisplayChangedEventArgs{p: p, ctx: ctx}
}

// IDisplayChangedEventArgsArrayToJSArray returns a JavaScript Array for the wrapped array.
func IDisplayChangedEventArgsArrayToJSArray(array []*IDisplayChangedEventArgs) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// VrDisplay returns the VrDisplay property of class IDisplayChangedEventArgs.
//
// https://doc.babylonjs.com/api/classes/babylon.idisplaychangedeventargs#vrdisplay
func (i *IDisplayChangedEventArgs) VrDisplay() js.Value {
	retVal := i.p.Get("vrDisplay")
	return retVal
}

// SetVrDisplay sets the VrDisplay property of class IDisplayChangedEventArgs.
//
// https://doc.babylonjs.com/api/classes/babylon.idisplaychangedeventargs#vrdisplay
func (i *IDisplayChangedEventArgs) SetVrDisplay(vrDisplay JSObject) *IDisplayChangedEventArgs {
	i.p.Set("vrDisplay", vrDisplay.JSObject())
	return i
}

// VrSupported returns the VrSupported property of class IDisplayChangedEventArgs.
//
// https://doc.babylonjs.com/api/classes/babylon.idisplaychangedeventargs#vrsupported
func (i *IDisplayChangedEventArgs) VrSupported() bool {
	retVal := i.p.Get("vrSupported")
	return retVal.Bool()
}

// SetVrSupported sets the VrSupported property of class IDisplayChangedEventArgs.
//
// https://doc.babylonjs.com/api/classes/babylon.idisplaychangedeventargs#vrsupported
func (i *IDisplayChangedEventArgs) SetVrSupported(vrSupported bool) *IDisplayChangedEventArgs {
	i.p.Set("vrSupported", vrSupported)
	return i
}
