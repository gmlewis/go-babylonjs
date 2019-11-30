package babylon

import "syscall/js"

// AttachControl calls the JavaScript method of the same name.
func (a *ArcRotateCamera) AttachControl(canvas js.Value, b bool) {
	a.p.Call("attachControl", canvas, b)
}
