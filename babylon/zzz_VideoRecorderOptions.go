// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// VideoRecorderOptions represents a babylon.js VideoRecorderOptions.
// This represents the different options available for the video capture.
type VideoRecorderOptions struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (v *VideoRecorderOptions) JSObject() js.Value { return v.p }

// VideoRecorderOptions returns a VideoRecorderOptions JavaScript class.
func (ba *Babylon) VideoRecorderOptions() *VideoRecorderOptions {
	p := ba.ctx.Get("VideoRecorderOptions")
	return VideoRecorderOptionsFromJSObject(p, ba.ctx)
}

// VideoRecorderOptionsFromJSObject returns a wrapped VideoRecorderOptions JavaScript class.
func VideoRecorderOptionsFromJSObject(p js.Value, ctx js.Value) *VideoRecorderOptions {
	return &VideoRecorderOptions{p: p, ctx: ctx}
}

// VideoRecorderOptionsArrayToJSArray returns a JavaScript Array for the wrapped array.
func VideoRecorderOptionsArrayToJSArray(array []*VideoRecorderOptions) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// AudioTracks returns the AudioTracks property of class VideoRecorderOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.videorecorderoptions#audiotracks
func (v *VideoRecorderOptions) AudioTracks() js.Value {
	retVal := v.p.Get("audioTracks")
	return retVal
}

// SetAudioTracks sets the AudioTracks property of class VideoRecorderOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.videorecorderoptions#audiotracks
func (v *VideoRecorderOptions) SetAudioTracks(audioTracks js.Value) *VideoRecorderOptions {
	v.p.Set("audioTracks", audioTracks)
	return v
}

// Fps returns the Fps property of class VideoRecorderOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.videorecorderoptions#fps
func (v *VideoRecorderOptions) Fps() float64 {
	retVal := v.p.Get("fps")
	return retVal.Float()
}

// SetFps sets the Fps property of class VideoRecorderOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.videorecorderoptions#fps
func (v *VideoRecorderOptions) SetFps(fps float64) *VideoRecorderOptions {
	v.p.Set("fps", fps)
	return v
}

// MimeType returns the MimeType property of class VideoRecorderOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.videorecorderoptions#mimetype
func (v *VideoRecorderOptions) MimeType() string {
	retVal := v.p.Get("mimeType")
	return retVal.String()
}

// SetMimeType sets the MimeType property of class VideoRecorderOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.videorecorderoptions#mimetype
func (v *VideoRecorderOptions) SetMimeType(mimeType string) *VideoRecorderOptions {
	v.p.Set("mimeType", mimeType)
	return v
}

// RecordChunckSize returns the RecordChunckSize property of class VideoRecorderOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.videorecorderoptions#recordchuncksize
func (v *VideoRecorderOptions) RecordChunckSize() float64 {
	retVal := v.p.Get("recordChunckSize")
	return retVal.Float()
}

// SetRecordChunckSize sets the RecordChunckSize property of class VideoRecorderOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.videorecorderoptions#recordchuncksize
func (v *VideoRecorderOptions) SetRecordChunckSize(recordChunckSize float64) *VideoRecorderOptions {
	v.p.Set("recordChunckSize", recordChunckSize)
	return v
}
