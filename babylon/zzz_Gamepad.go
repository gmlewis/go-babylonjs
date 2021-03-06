// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Gamepad represents a babylon.js Gamepad.
// Represents a gamepad
type Gamepad struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (g *Gamepad) JSObject() js.Value { return g.p }

// Gamepad returns a Gamepad JavaScript class.
func (ba *Babylon) Gamepad() *Gamepad {
	p := ba.ctx.Get("Gamepad")
	return GamepadFromJSObject(p, ba.ctx)
}

// GamepadFromJSObject returns a wrapped Gamepad JavaScript class.
func GamepadFromJSObject(p js.Value, ctx js.Value) *Gamepad {
	return &Gamepad{p: p, ctx: ctx}
}

// GamepadArrayToJSArray returns a JavaScript Array for the wrapped array.
func GamepadArrayToJSArray(array []*Gamepad) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewGamepadOpts contains optional parameters for NewGamepad.
type NewGamepadOpts struct {
	LeftStickX  *float64
	LeftStickY  *float64
	RightStickX *float64
	RightStickY *float64
}

// NewGamepad returns a new Gamepad object.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepad#constructor
func (ba *Babylon) NewGamepad(id string, index float64, browserGamepad JSObject, opts *NewGamepadOpts) *Gamepad {
	if opts == nil {
		opts = &NewGamepadOpts{}
	}

	args := make([]interface{}, 0, 3+4)

	args = append(args, id)
	args = append(args, index)
	args = append(args, browserGamepad.JSObject())

	if opts.LeftStickX == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.LeftStickX)
	}
	if opts.LeftStickY == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.LeftStickY)
	}
	if opts.RightStickX == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.RightStickX)
	}
	if opts.RightStickY == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.RightStickY)
	}

	p := ba.ctx.Get("Gamepad").New(args...)
	return GamepadFromJSObject(p, ba.ctx)
}

// Dispose calls the Dispose method on the Gamepad object.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepad#dispose
func (g *Gamepad) Dispose() {

	g.p.Call("dispose")
}

// Onleftstickchanged calls the Onleftstickchanged method on the Gamepad object.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepad#onleftstickchanged
func (g *Gamepad) Onleftstickchanged(callback JSFunc) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, js.FuncOf(callback))

	g.p.Call("onleftstickchanged", args...)
}

// Onrightstickchanged calls the Onrightstickchanged method on the Gamepad object.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepad#onrightstickchanged
func (g *Gamepad) Onrightstickchanged(callback JSFunc) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, js.FuncOf(callback))

	g.p.Call("onrightstickchanged", args...)
}

// Update calls the Update method on the Gamepad object.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepad#update
func (g *Gamepad) Update() {

	g.p.Call("update")
}

// BrowserGamepad returns the BrowserGamepad property of class Gamepad.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepad#browsergamepad
func (g *Gamepad) BrowserGamepad() js.Value {
	retVal := g.p.Get("browserGamepad")
	return retVal
}

// SetBrowserGamepad sets the BrowserGamepad property of class Gamepad.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepad#browsergamepad
func (g *Gamepad) SetBrowserGamepad(browserGamepad JSObject) *Gamepad {
	g.p.Set("browserGamepad", browserGamepad.JSObject())
	return g
}

// DUALSHOCK returns the DUALSHOCK property of class Gamepad.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepad#dualshock
func (g *Gamepad) DUALSHOCK() float64 {
	retVal := g.p.Get("DUALSHOCK")
	return retVal.Float()
}

// SetDUALSHOCK sets the DUALSHOCK property of class Gamepad.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepad#dualshock
func (g *Gamepad) SetDUALSHOCK(DUALSHOCK float64) *Gamepad {
	g.p.Set("DUALSHOCK", DUALSHOCK)
	return g
}

// GAMEPAD returns the GAMEPAD property of class Gamepad.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepad#gamepad
func (g *Gamepad) GAMEPAD() float64 {
	retVal := g.p.Get("GAMEPAD")
	return retVal.Float()
}

// SetGAMEPAD sets the GAMEPAD property of class Gamepad.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepad#gamepad
func (g *Gamepad) SetGAMEPAD(GAMEPAD float64) *Gamepad {
	g.p.Set("GAMEPAD", GAMEPAD)
	return g
}

// GENERIC returns the GENERIC property of class Gamepad.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepad#generic
func (g *Gamepad) GENERIC() float64 {
	retVal := g.p.Get("GENERIC")
	return retVal.Float()
}

// SetGENERIC sets the GENERIC property of class Gamepad.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepad#generic
func (g *Gamepad) SetGENERIC(GENERIC float64) *Gamepad {
	g.p.Set("GENERIC", GENERIC)
	return g
}

// Id returns the Id property of class Gamepad.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepad#id
func (g *Gamepad) Id() string {
	retVal := g.p.Get("id")
	return retVal.String()
}

// SetId sets the Id property of class Gamepad.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepad#id
func (g *Gamepad) SetId(id string) *Gamepad {
	g.p.Set("id", id)
	return g
}

// Index returns the Index property of class Gamepad.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepad#index
func (g *Gamepad) Index() float64 {
	retVal := g.p.Get("index")
	return retVal.Float()
}

// SetIndex sets the Index property of class Gamepad.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepad#index
func (g *Gamepad) SetIndex(index float64) *Gamepad {
	g.p.Set("index", index)
	return g
}

// IsConnected returns the IsConnected property of class Gamepad.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepad#isconnected
func (g *Gamepad) IsConnected() bool {
	retVal := g.p.Get("isConnected")
	return retVal.Bool()
}

// SetIsConnected sets the IsConnected property of class Gamepad.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepad#isconnected
func (g *Gamepad) SetIsConnected(isConnected bool) *Gamepad {
	g.p.Set("isConnected", isConnected)
	return g
}

// LeftStick returns the LeftStick property of class Gamepad.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepad#leftstick
func (g *Gamepad) LeftStick() *StickValues {
	retVal := g.p.Get("leftStick")
	return StickValuesFromJSObject(retVal, g.ctx)
}

// SetLeftStick sets the LeftStick property of class Gamepad.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepad#leftstick
func (g *Gamepad) SetLeftStick(leftStick *StickValues) *Gamepad {
	g.p.Set("leftStick", leftStick.JSObject())
	return g
}

// POSE_ENABLED returns the POSE_ENABLED property of class Gamepad.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepad#pose_enabled
func (g *Gamepad) POSE_ENABLED() float64 {
	retVal := g.p.Get("POSE_ENABLED")
	return retVal.Float()
}

// SetPOSE_ENABLED sets the POSE_ENABLED property of class Gamepad.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepad#pose_enabled
func (g *Gamepad) SetPOSE_ENABLED(POSE_ENABLED float64) *Gamepad {
	g.p.Set("POSE_ENABLED", POSE_ENABLED)
	return g
}

// RightStick returns the RightStick property of class Gamepad.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepad#rightstick
func (g *Gamepad) RightStick() *StickValues {
	retVal := g.p.Get("rightStick")
	return StickValuesFromJSObject(retVal, g.ctx)
}

// SetRightStick sets the RightStick property of class Gamepad.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepad#rightstick
func (g *Gamepad) SetRightStick(rightStick *StickValues) *Gamepad {
	g.p.Set("rightStick", rightStick.JSObject())
	return g
}

// Type returns the Type property of class Gamepad.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepad#type
func (g *Gamepad) Type() float64 {
	retVal := g.p.Get("type")
	return retVal.Float()
}

// SetType sets the Type property of class Gamepad.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepad#type
func (g *Gamepad) SetType(jsType float64) *Gamepad {
	g.p.Set("type", jsType)
	return g
}

// XBOX returns the XBOX property of class Gamepad.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepad#xbox
func (g *Gamepad) XBOX() float64 {
	retVal := g.p.Get("XBOX")
	return retVal.Float()
}

// SetXBOX sets the XBOX property of class Gamepad.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepad#xbox
func (g *Gamepad) SetXBOX(XBOX float64) *Gamepad {
	g.p.Set("XBOX", XBOX)
	return g
}
