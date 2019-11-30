// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// VRCameraMetrics represents a babylon.js VRCameraMetrics.
// This represents all the required metrics to create a VR camera.
//
// See: http://doc.babylonjs.com/babylon101/cameras#device-orientation-camera
type VRCameraMetrics struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (v *VRCameraMetrics) JSObject() js.Value { return v.p }

// VRCameraMetrics returns a VRCameraMetrics JavaScript class.
func (b *Babylon) VRCameraMetrics() *VRCameraMetrics {
	p := b.ctx.Get("VRCameraMetrics")
	return VRCameraMetricsFromJSObject(p)
}

// VRCameraMetricsFromJSObject returns a wrapped VRCameraMetrics JavaScript class.
func VRCameraMetricsFromJSObject(p js.Value) *VRCameraMetrics {
	return &VRCameraMetrics{p: p}
}

// TODO: methods
