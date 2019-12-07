package babylon

import "syscall/js"

// AttachControl calls the AttachControl method on the FreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.freecamera#attachcontrol
func (c *FreeCamera) AttachControl(element js.Value, noPreventDefault *bool) {
	if noPreventDefault != nil {
		c.p.Call("attachControl", element, *noPreventDefault)
	} else {
		c.p.Call("attachControl", element)
	}
}
