// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AudioSceneComponent represents a babylon.js AudioSceneComponent.
// Defines the sound scene component responsible to manage any sounds
// in a given scene.
type AudioSceneComponent struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *AudioSceneComponent) JSObject() js.Value { return a.p }

// AudioSceneComponent returns a AudioSceneComponent JavaScript class.
func (ba *Babylon) AudioSceneComponent() *AudioSceneComponent {
	p := ba.ctx.Get("AudioSceneComponent")
	return AudioSceneComponentFromJSObject(p, ba.ctx)
}

// AudioSceneComponentFromJSObject returns a wrapped AudioSceneComponent JavaScript class.
func AudioSceneComponentFromJSObject(p js.Value, ctx js.Value) *AudioSceneComponent {
	return &AudioSceneComponent{p: p, ctx: ctx}
}

// AudioSceneComponentArrayToJSArray returns a JavaScript Array for the wrapped array.
func AudioSceneComponentArrayToJSArray(array []*AudioSceneComponent) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewAudioSceneComponent returns a new AudioSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.audioscenecomponent#constructor
func (ba *Babylon) NewAudioSceneComponent(scene *Scene) *AudioSceneComponent {

	args := make([]interface{}, 0, 1+0)

	args = append(args, scene.JSObject())

	p := ba.ctx.Get("AudioSceneComponent").New(args...)
	return AudioSceneComponentFromJSObject(p, ba.ctx)
}

// AddFromContainer calls the AddFromContainer method on the AudioSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.audioscenecomponent#addfromcontainer
func (a *AudioSceneComponent) AddFromContainer(container *AbstractScene) {

	args := make([]interface{}, 0, 1+0)

	if container == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, container.JSObject())
	}

	a.p.Call("addFromContainer", args...)
}

// DisableAudio calls the DisableAudio method on the AudioSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.audioscenecomponent#disableaudio
func (a *AudioSceneComponent) DisableAudio() {

	a.p.Call("disableAudio")
}

// Dispose calls the Dispose method on the AudioSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.audioscenecomponent#dispose
func (a *AudioSceneComponent) Dispose() {

	a.p.Call("dispose")
}

// EnableAudio calls the EnableAudio method on the AudioSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.audioscenecomponent#enableaudio
func (a *AudioSceneComponent) EnableAudio() {

	a.p.Call("enableAudio")
}

// Rebuild calls the Rebuild method on the AudioSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.audioscenecomponent#rebuild
func (a *AudioSceneComponent) Rebuild() {

	a.p.Call("rebuild")
}

// Register calls the Register method on the AudioSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.audioscenecomponent#register
func (a *AudioSceneComponent) Register() {

	a.p.Call("register")
}

// AudioSceneComponentRemoveFromContainerOpts contains optional parameters for AudioSceneComponent.RemoveFromContainer.
type AudioSceneComponentRemoveFromContainerOpts struct {
	Dispose *bool
}

// RemoveFromContainer calls the RemoveFromContainer method on the AudioSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.audioscenecomponent#removefromcontainer
func (a *AudioSceneComponent) RemoveFromContainer(container *AbstractScene, opts *AudioSceneComponentRemoveFromContainerOpts) {
	if opts == nil {
		opts = &AudioSceneComponentRemoveFromContainerOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	if container == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, container.JSObject())
	}

	if opts.Dispose == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Dispose)
	}

	a.p.Call("removeFromContainer", args...)
}

// Serialize calls the Serialize method on the AudioSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.audioscenecomponent#serialize
func (a *AudioSceneComponent) Serialize(serializationObject JSObject) {

	args := make([]interface{}, 0, 1+0)

	if serializationObject == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, serializationObject.JSObject())
	}

	a.p.Call("serialize", args...)
}

// SwitchAudioModeForHeadphones calls the SwitchAudioModeForHeadphones method on the AudioSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.audioscenecomponent#switchaudiomodeforheadphones
func (a *AudioSceneComponent) SwitchAudioModeForHeadphones() {

	a.p.Call("switchAudioModeForHeadphones")
}

// SwitchAudioModeForNormalSpeakers calls the SwitchAudioModeForNormalSpeakers method on the AudioSceneComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.audioscenecomponent#switchaudiomodefornormalspeakers
func (a *AudioSceneComponent) SwitchAudioModeForNormalSpeakers() {

	a.p.Call("switchAudioModeForNormalSpeakers")
}

// AudioEnabled returns the AudioEnabled property of class AudioSceneComponent.
//
// https://doc.babylonjs.com/api/classes/babylon.audioscenecomponent#audioenabled
func (a *AudioSceneComponent) AudioEnabled() bool {
	retVal := a.p.Get("audioEnabled")
	return retVal.Bool()
}

// SetAudioEnabled sets the AudioEnabled property of class AudioSceneComponent.
//
// https://doc.babylonjs.com/api/classes/babylon.audioscenecomponent#audioenabled
func (a *AudioSceneComponent) SetAudioEnabled(audioEnabled bool) *AudioSceneComponent {
	a.p.Set("audioEnabled", audioEnabled)
	return a
}

// AudioListenerPositionProvider returns the AudioListenerPositionProvider property of class AudioSceneComponent.
//
// https://doc.babylonjs.com/api/classes/babylon.audioscenecomponent#audiolistenerpositionprovider
func (a *AudioSceneComponent) AudioListenerPositionProvider() js.Value {
	retVal := a.p.Get("audioListenerPositionProvider")
	return retVal
}

// SetAudioListenerPositionProvider sets the AudioListenerPositionProvider property of class AudioSceneComponent.
//
// https://doc.babylonjs.com/api/classes/babylon.audioscenecomponent#audiolistenerpositionprovider
func (a *AudioSceneComponent) SetAudioListenerPositionProvider(audioListenerPositionProvider JSFunc) *AudioSceneComponent {
	a.p.Set("audioListenerPositionProvider", js.FuncOf(audioListenerPositionProvider))
	return a
}

// AudioPositioningRefreshRate returns the AudioPositioningRefreshRate property of class AudioSceneComponent.
//
// https://doc.babylonjs.com/api/classes/babylon.audioscenecomponent#audiopositioningrefreshrate
func (a *AudioSceneComponent) AudioPositioningRefreshRate() float64 {
	retVal := a.p.Get("audioPositioningRefreshRate")
	return retVal.Float()
}

// SetAudioPositioningRefreshRate sets the AudioPositioningRefreshRate property of class AudioSceneComponent.
//
// https://doc.babylonjs.com/api/classes/babylon.audioscenecomponent#audiopositioningrefreshrate
func (a *AudioSceneComponent) SetAudioPositioningRefreshRate(audioPositioningRefreshRate float64) *AudioSceneComponent {
	a.p.Set("audioPositioningRefreshRate", audioPositioningRefreshRate)
	return a
}

// Headphone returns the Headphone property of class AudioSceneComponent.
//
// https://doc.babylonjs.com/api/classes/babylon.audioscenecomponent#headphone
func (a *AudioSceneComponent) Headphone() bool {
	retVal := a.p.Get("headphone")
	return retVal.Bool()
}

// SetHeadphone sets the Headphone property of class AudioSceneComponent.
//
// https://doc.babylonjs.com/api/classes/babylon.audioscenecomponent#headphone
func (a *AudioSceneComponent) SetHeadphone(headphone bool) *AudioSceneComponent {
	a.p.Set("headphone", headphone)
	return a
}

// Name returns the Name property of class AudioSceneComponent.
//
// https://doc.babylonjs.com/api/classes/babylon.audioscenecomponent#name
func (a *AudioSceneComponent) Name() string {
	retVal := a.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class AudioSceneComponent.
//
// https://doc.babylonjs.com/api/classes/babylon.audioscenecomponent#name
func (a *AudioSceneComponent) SetName(name string) *AudioSceneComponent {
	a.p.Set("name", name)
	return a
}

// Scene returns the Scene property of class AudioSceneComponent.
//
// https://doc.babylonjs.com/api/classes/babylon.audioscenecomponent#scene
func (a *AudioSceneComponent) Scene() *Scene {
	retVal := a.p.Get("scene")
	return SceneFromJSObject(retVal, a.ctx)
}

// SetScene sets the Scene property of class AudioSceneComponent.
//
// https://doc.babylonjs.com/api/classes/babylon.audioscenecomponent#scene
func (a *AudioSceneComponent) SetScene(scene *Scene) *AudioSceneComponent {
	a.p.Set("scene", scene.JSObject())
	return a
}
