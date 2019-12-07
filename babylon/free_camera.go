package babylon

import "syscall/js"

// AttachControl calls the AttachControl method on the FreeCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.freecamera#attachcontrol
func (f *FreeCamera) AttachControl(element js.Value, noPreventDefault *bool) {
	if noPreventDefault != nil {
		f.p.Call("attachControl", element, *noPreventDefault)
	} else {
		f.p.Call("attachControl", element)
	}
}
