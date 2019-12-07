package babylon

import (
	"syscall/js"
)

// AttachControl calls the AttachControl method on the FollowCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamera#attachcontrol
func (f *FollowCamera) AttachControl(element js.Value, noPreventDefault *bool) {
	if noPreventDefault != nil {
		f.p.Call("attachControl", element, *noPreventDefault)
	} else {
		f.p.Call("attachControl", element)
	}
}
