// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SoundTrack represents a babylon.js SoundTrack.
// It could be useful to isolate your music &amp; sounds on several tracks to better manage volume on a grouped instance of sounds.
// It will be also used in a future release to apply effects on a specific track.
//
// See: http://doc.babylonjs.com/how_to/playing_sounds_and_music#using-sound-tracks
type SoundTrack struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *SoundTrack) JSObject() js.Value { return s.p }

// SoundTrack returns a SoundTrack JavaScript class.
func (ba *Babylon) SoundTrack() *SoundTrack {
	p := ba.ctx.Get("SoundTrack")
	return SoundTrackFromJSObject(p, ba.ctx)
}

// SoundTrackFromJSObject returns a wrapped SoundTrack JavaScript class.
func SoundTrackFromJSObject(p js.Value, ctx js.Value) *SoundTrack {
	return &SoundTrack{p: p, ctx: ctx}
}

// SoundTrackArrayToJSArray returns a JavaScript Array for the wrapped array.
func SoundTrackArrayToJSArray(array []*SoundTrack) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewSoundTrackOpts contains optional parameters for NewSoundTrack.
type NewSoundTrackOpts struct {
	Options *ISoundTrackOptions
}

// NewSoundTrack returns a new SoundTrack object.
//
// https://doc.babylonjs.com/api/classes/babylon.soundtrack#constructor
func (ba *Babylon) NewSoundTrack(scene *Scene, opts *NewSoundTrackOpts) *SoundTrack {
	if opts == nil {
		opts = &NewSoundTrackOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, scene.JSObject())

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options.JSObject())
	}

	p := ba.ctx.Get("SoundTrack").New(args...)
	return SoundTrackFromJSObject(p, ba.ctx)
}

// AddSound calls the AddSound method on the SoundTrack object.
//
// https://doc.babylonjs.com/api/classes/babylon.soundtrack#addsound
func (s *SoundTrack) AddSound(sound *Sound) {

	args := make([]interface{}, 0, 1+0)

	if sound == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, sound.JSObject())
	}

	s.p.Call("AddSound", args...)
}

// ConnectToAnalyser calls the ConnectToAnalyser method on the SoundTrack object.
//
// https://doc.babylonjs.com/api/classes/babylon.soundtrack#connecttoanalyser
func (s *SoundTrack) ConnectToAnalyser(analyser *Analyser) {

	args := make([]interface{}, 0, 1+0)

	if analyser == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, analyser.JSObject())
	}

	s.p.Call("connectToAnalyser", args...)
}

// Dispose calls the Dispose method on the SoundTrack object.
//
// https://doc.babylonjs.com/api/classes/babylon.soundtrack#dispose
func (s *SoundTrack) Dispose() {

	s.p.Call("dispose")
}

// RemoveSound calls the RemoveSound method on the SoundTrack object.
//
// https://doc.babylonjs.com/api/classes/babylon.soundtrack#removesound
func (s *SoundTrack) RemoveSound(sound *Sound) {

	args := make([]interface{}, 0, 1+0)

	if sound == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, sound.JSObject())
	}

	s.p.Call("RemoveSound", args...)
}

// SetVolume calls the SetVolume method on the SoundTrack object.
//
// https://doc.babylonjs.com/api/classes/babylon.soundtrack#setvolume
func (s *SoundTrack) SetVolume(newVolume float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, newVolume)

	s.p.Call("setVolume", args...)
}

// SwitchPanningModelToEqualPower calls the SwitchPanningModelToEqualPower method on the SoundTrack object.
//
// https://doc.babylonjs.com/api/classes/babylon.soundtrack#switchpanningmodeltoequalpower
func (s *SoundTrack) SwitchPanningModelToEqualPower() {

	s.p.Call("switchPanningModelToEqualPower")
}

// SwitchPanningModelToHRTF calls the SwitchPanningModelToHRTF method on the SoundTrack object.
//
// https://doc.babylonjs.com/api/classes/babylon.soundtrack#switchpanningmodeltohrtf
func (s *SoundTrack) SwitchPanningModelToHRTF() {

	s.p.Call("switchPanningModelToHRTF")
}

// Id returns the Id property of class SoundTrack.
//
// https://doc.babylonjs.com/api/classes/babylon.soundtrack#id
func (s *SoundTrack) Id() float64 {
	retVal := s.p.Get("id")
	return retVal.Float()
}

// SetId sets the Id property of class SoundTrack.
//
// https://doc.babylonjs.com/api/classes/babylon.soundtrack#id
func (s *SoundTrack) SetId(id float64) *SoundTrack {
	s.p.Set("id", id)
	return s
}

// SoundCollection returns the SoundCollection property of class SoundTrack.
//
// https://doc.babylonjs.com/api/classes/babylon.soundtrack#soundcollection
func (s *SoundTrack) SoundCollection() []*Sound {
	retVal := s.p.Get("soundCollection")
	result := []*Sound{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, SoundFromJSObject(retVal.Index(ri), s.ctx))
	}
	return result
}

// SetSoundCollection sets the SoundCollection property of class SoundTrack.
//
// https://doc.babylonjs.com/api/classes/babylon.soundtrack#soundcollection
func (s *SoundTrack) SetSoundCollection(soundCollection []*Sound) *SoundTrack {
	s.p.Set("soundCollection", soundCollection)
	return s
}
