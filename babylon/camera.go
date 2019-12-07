package babylon

import "syscall/js"

// AttachControl calls the AttachControl method on the Camera object.
//
// https://doc.babylonjs.com/api/classes/babylon.camera#attachcontrol
func (c *Camera) AttachControl(element js.Value, noPreventDefault *bool) {
	if noPreventDefault != nil {
		c.p.Call("attachControl", element, *noPreventDefault)
	} else {
		c.p.Call("attachControl", element)
	}
}
