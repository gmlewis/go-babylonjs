// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DevicePose represents a babylon.js DevicePose.
// This is a copy of VRPose. See <a href="https://developer.mozilla.org/en-US/docs/Web/API/VRPose">https://developer.mozilla.org/en-US/docs/Web/API/VRPose</a>
// IMPORTANT!! The data is right-hand data.
type DevicePose struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (d *DevicePose) JSObject() js.Value { return d.p }

// DevicePose returns a DevicePose JavaScript class.
func (ba *Babylon) DevicePose() *DevicePose {
	p := ba.ctx.Get("DevicePose")
	return DevicePoseFromJSObject(p, ba.ctx)
}

// DevicePoseFromJSObject returns a wrapped DevicePose JavaScript class.
func DevicePoseFromJSObject(p js.Value, ctx js.Value) *DevicePose {
	return &DevicePose{p: p, ctx: ctx}
}

// DevicePoseArrayToJSArray returns a JavaScript Array for the wrapped array.
func DevicePoseArrayToJSArray(array []*DevicePose) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// AngularAcceleration returns the AngularAcceleration property of class DevicePose.
//
// https://doc.babylonjs.com/api/classes/babylon.devicepose#angularacceleration
func (d *DevicePose) AngularAcceleration() js.Value {
	retVal := d.p.Get("angularAcceleration")
	return retVal
}

// SetAngularAcceleration sets the AngularAcceleration property of class DevicePose.
//
// https://doc.babylonjs.com/api/classes/babylon.devicepose#angularacceleration
func (d *DevicePose) SetAngularAcceleration(angularAcceleration js.Value) *DevicePose {
	d.p.Set("angularAcceleration", angularAcceleration)
	return d
}

// AngularVelocity returns the AngularVelocity property of class DevicePose.
//
// https://doc.babylonjs.com/api/classes/babylon.devicepose#angularvelocity
func (d *DevicePose) AngularVelocity() js.Value {
	retVal := d.p.Get("angularVelocity")
	return retVal
}

// SetAngularVelocity sets the AngularVelocity property of class DevicePose.
//
// https://doc.babylonjs.com/api/classes/babylon.devicepose#angularvelocity
func (d *DevicePose) SetAngularVelocity(angularVelocity js.Value) *DevicePose {
	d.p.Set("angularVelocity", angularVelocity)
	return d
}

// LinearAcceleration returns the LinearAcceleration property of class DevicePose.
//
// https://doc.babylonjs.com/api/classes/babylon.devicepose#linearacceleration
func (d *DevicePose) LinearAcceleration() js.Value {
	retVal := d.p.Get("linearAcceleration")
	return retVal
}

// SetLinearAcceleration sets the LinearAcceleration property of class DevicePose.
//
// https://doc.babylonjs.com/api/classes/babylon.devicepose#linearacceleration
func (d *DevicePose) SetLinearAcceleration(linearAcceleration js.Value) *DevicePose {
	d.p.Set("linearAcceleration", linearAcceleration)
	return d
}

// LinearVelocity returns the LinearVelocity property of class DevicePose.
//
// https://doc.babylonjs.com/api/classes/babylon.devicepose#linearvelocity
func (d *DevicePose) LinearVelocity() js.Value {
	retVal := d.p.Get("linearVelocity")
	return retVal
}

// SetLinearVelocity sets the LinearVelocity property of class DevicePose.
//
// https://doc.babylonjs.com/api/classes/babylon.devicepose#linearvelocity
func (d *DevicePose) SetLinearVelocity(linearVelocity js.Value) *DevicePose {
	d.p.Set("linearVelocity", linearVelocity)
	return d
}

// Orientation returns the Orientation property of class DevicePose.
//
// https://doc.babylonjs.com/api/classes/babylon.devicepose#orientation
func (d *DevicePose) Orientation() js.Value {
	retVal := d.p.Get("orientation")
	return retVal
}

// SetOrientation sets the Orientation property of class DevicePose.
//
// https://doc.babylonjs.com/api/classes/babylon.devicepose#orientation
func (d *DevicePose) SetOrientation(orientation js.Value) *DevicePose {
	d.p.Set("orientation", orientation)
	return d
}

// Position returns the Position property of class DevicePose.
//
// https://doc.babylonjs.com/api/classes/babylon.devicepose#position
func (d *DevicePose) Position() js.Value {
	retVal := d.p.Get("position")
	return retVal
}

// SetPosition sets the Position property of class DevicePose.
//
// https://doc.babylonjs.com/api/classes/babylon.devicepose#position
func (d *DevicePose) SetPosition(position js.Value) *DevicePose {
	d.p.Set("position", position)
	return d
}
