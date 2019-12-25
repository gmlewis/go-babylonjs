// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GamepadManager represents a babylon.js GamepadManager.
// Manager for handling gamepads
type GamepadManager struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (g *GamepadManager) JSObject() js.Value { return g.p }

// GamepadManager returns a GamepadManager JavaScript class.
func (ba *Babylon) GamepadManager() *GamepadManager {
	p := ba.ctx.Get("GamepadManager")
	return GamepadManagerFromJSObject(p, ba.ctx)
}

// GamepadManagerFromJSObject returns a wrapped GamepadManager JavaScript class.
func GamepadManagerFromJSObject(p js.Value, ctx js.Value) *GamepadManager {
	return &GamepadManager{p: p, ctx: ctx}
}

// GamepadManagerArrayToJSArray returns a JavaScript Array for the wrapped array.
func GamepadManagerArrayToJSArray(array []*GamepadManager) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewGamepadManagerOpts contains optional parameters for NewGamepadManager.
type NewGamepadManagerOpts struct {
	_scene *Scene
}

// NewGamepadManager returns a new GamepadManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepadmanager#constructor
func (ba *Babylon) NewGamepadManager(opts *NewGamepadManagerOpts) *GamepadManager {
	if opts == nil {
		opts = &NewGamepadManagerOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts._scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts._scene.JSObject())
	}

	p := ba.ctx.Get("GamepadManager").New(args...)
	return GamepadManagerFromJSObject(p, ba.ctx)
}

// Dispose calls the Dispose method on the GamepadManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepadmanager#dispose
func (g *GamepadManager) Dispose() {

	g.p.Call("dispose")
}

// GamepadManagerGetGamepadByTypeOpts contains optional parameters for GamepadManager.GetGamepadByType.
type GamepadManagerGetGamepadByTypeOpts struct {
	Type *float64
}

// GetGamepadByType calls the GetGamepadByType method on the GamepadManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepadmanager#getgamepadbytype
func (g *GamepadManager) GetGamepadByType(opts *GamepadManagerGetGamepadByTypeOpts) *Gamepad {
	if opts == nil {
		opts = &GamepadManagerGetGamepadByTypeOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Type == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Type)
	}

	retVal := g.p.Call("getGamepadByType", args...)
	return GamepadFromJSObject(retVal, g.ctx)
}

// Gamepads returns the Gamepads property of class GamepadManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepadmanager#gamepads
func (g *GamepadManager) Gamepads() []*Gamepad {
	retVal := g.p.Get("gamepads")
	result := []*Gamepad{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, GamepadFromJSObject(retVal.Index(ri), g.ctx))
	}
	return result
}

// SetGamepads sets the Gamepads property of class GamepadManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepadmanager#gamepads
func (g *GamepadManager) SetGamepads(gamepads []*Gamepad) *GamepadManager {
	g.p.Set("gamepads", gamepads)
	return g
}

// OnGamepadConnectedObservable returns the OnGamepadConnectedObservable property of class GamepadManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepadmanager#ongamepadconnectedobservable
func (g *GamepadManager) OnGamepadConnectedObservable() *Observable {
	retVal := g.p.Get("onGamepadConnectedObservable")
	return ObservableFromJSObject(retVal, g.ctx)
}

// SetOnGamepadConnectedObservable sets the OnGamepadConnectedObservable property of class GamepadManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepadmanager#ongamepadconnectedobservable
func (g *GamepadManager) SetOnGamepadConnectedObservable(onGamepadConnectedObservable *Observable) *GamepadManager {
	g.p.Set("onGamepadConnectedObservable", onGamepadConnectedObservable.JSObject())
	return g
}

// OnGamepadDisconnectedObservable returns the OnGamepadDisconnectedObservable property of class GamepadManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepadmanager#ongamepaddisconnectedobservable
func (g *GamepadManager) OnGamepadDisconnectedObservable() *Observable {
	retVal := g.p.Get("onGamepadDisconnectedObservable")
	return ObservableFromJSObject(retVal, g.ctx)
}

// SetOnGamepadDisconnectedObservable sets the OnGamepadDisconnectedObservable property of class GamepadManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gamepadmanager#ongamepaddisconnectedobservable
func (g *GamepadManager) SetOnGamepadDisconnectedObservable(onGamepadDisconnectedObservable *Observable) *GamepadManager {
	g.p.Set("onGamepadDisconnectedObservable", onGamepadDisconnectedObservable.JSObject())
	return g
}
