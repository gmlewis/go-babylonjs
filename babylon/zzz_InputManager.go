// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// InputManager represents a babylon.js InputManager.
// Class used to manage all inputs for the scene.
type InputManager struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *InputManager) JSObject() js.Value { return i.p }

// InputManager returns a InputManager JavaScript class.
func (ba *Babylon) InputManager() *InputManager {
	p := ba.ctx.Get("InputManager")
	return InputManagerFromJSObject(p, ba.ctx)
}

// InputManagerFromJSObject returns a wrapped InputManager JavaScript class.
func InputManagerFromJSObject(p js.Value, ctx js.Value) *InputManager {
	return &InputManager{p: p, ctx: ctx}
}

// NewInputManager returns a new InputManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputmanager
func (ba *Babylon) NewInputManager(scene *Scene) *InputManager {

	args := make([]interface{}, 0, 1+0)

	args = append(args, scene.JSObject())

	p := ba.ctx.Get("InputManager").New(args...)
	return InputManagerFromJSObject(p, ba.ctx)
}

// InputManagerAttachControlOpts contains optional parameters for InputManager.AttachControl.
type InputManagerAttachControlOpts struct {
	AttachUp   *bool
	AttachDown *bool
	AttachMove *bool
}

// AttachControl calls the AttachControl method on the InputManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputmanager#attachcontrol
func (i *InputManager) AttachControl(opts *InputManagerAttachControlOpts) {
	if opts == nil {
		opts = &InputManagerAttachControlOpts{}
	}

	args := make([]interface{}, 0, 0+3)

	if opts.AttachUp == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.AttachUp)
	}
	if opts.AttachDown == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.AttachDown)
	}
	if opts.AttachMove == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.AttachMove)
	}

	i.p.Call("attachControl", args...)
}

// DetachControl calls the DetachControl method on the InputManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputmanager#detachcontrol
func (i *InputManager) DetachControl() {

	args := make([]interface{}, 0, 0+0)

	i.p.Call("detachControl", args...)
}

// GetPointerOverMesh calls the GetPointerOverMesh method on the InputManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputmanager#getpointerovermesh
func (i *InputManager) GetPointerOverMesh() *AbstractMesh {

	args := make([]interface{}, 0, 0+0)

	retVal := i.p.Call("getPointerOverMesh", args...)
	return AbstractMeshFromJSObject(retVal, i.ctx)
}

// InputManagerIsPointerCapturedOpts contains optional parameters for InputManager.IsPointerCaptured.
type InputManagerIsPointerCapturedOpts struct {
	PointerId *float64
}

// IsPointerCaptured calls the IsPointerCaptured method on the InputManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputmanager#ispointercaptured
func (i *InputManager) IsPointerCaptured(opts *InputManagerIsPointerCapturedOpts) bool {
	if opts == nil {
		opts = &InputManagerIsPointerCapturedOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.PointerId == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.PointerId)
	}

	retVal := i.p.Call("isPointerCaptured", args...)
	return retVal.Bool()
}

// SetPointerOverMesh calls the SetPointerOverMesh method on the InputManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputmanager#setpointerovermesh
func (i *InputManager) SetPointerOverMesh(mesh *AbstractMesh) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, mesh.JSObject())

	i.p.Call("setPointerOverMesh", args...)
}

// InputManagerSimulatePointerDownOpts contains optional parameters for InputManager.SimulatePointerDown.
type InputManagerSimulatePointerDownOpts struct {
	PointerEventInit *PointerEventInit
}

// SimulatePointerDown calls the SimulatePointerDown method on the InputManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputmanager#simulatepointerdown
func (i *InputManager) SimulatePointerDown(pickResult *PickingInfo, opts *InputManagerSimulatePointerDownOpts) {
	if opts == nil {
		opts = &InputManagerSimulatePointerDownOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, pickResult.JSObject())

	if opts.PointerEventInit == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.PointerEventInit.JSObject())
	}

	i.p.Call("simulatePointerDown", args...)
}

// InputManagerSimulatePointerMoveOpts contains optional parameters for InputManager.SimulatePointerMove.
type InputManagerSimulatePointerMoveOpts struct {
	PointerEventInit *PointerEventInit
}

// SimulatePointerMove calls the SimulatePointerMove method on the InputManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputmanager#simulatepointermove
func (i *InputManager) SimulatePointerMove(pickResult *PickingInfo, opts *InputManagerSimulatePointerMoveOpts) {
	if opts == nil {
		opts = &InputManagerSimulatePointerMoveOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, pickResult.JSObject())

	if opts.PointerEventInit == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.PointerEventInit.JSObject())
	}

	i.p.Call("simulatePointerMove", args...)
}

// InputManagerSimulatePointerUpOpts contains optional parameters for InputManager.SimulatePointerUp.
type InputManagerSimulatePointerUpOpts struct {
	PointerEventInit *PointerEventInit
	DoubleTap        *bool
}

// SimulatePointerUp calls the SimulatePointerUp method on the InputManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputmanager#simulatepointerup
func (i *InputManager) SimulatePointerUp(pickResult *PickingInfo, opts *InputManagerSimulatePointerUpOpts) {
	if opts == nil {
		opts = &InputManagerSimulatePointerUpOpts{}
	}

	args := make([]interface{}, 0, 1+2)

	args = append(args, pickResult.JSObject())

	if opts.PointerEventInit == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.PointerEventInit.JSObject())
	}
	if opts.DoubleTap == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.DoubleTap)
	}

	i.p.Call("simulatePointerUp", args...)
}

/*

// DoubleClickDelay returns the DoubleClickDelay property of class InputManager.
//
// https://doc.babylonjs.com/api/classes/babylon.inputmanager#doubleclickdelay
func (i *InputManager) DoubleClickDelay(DoubleClickDelay float64) *InputManager {
	p := ba.ctx.Get("InputManager").New(DoubleClickDelay)
	return InputManagerFromJSObject(p, ba.ctx)
}

// SetDoubleClickDelay sets the DoubleClickDelay property of class InputManager.
//
// https://doc.babylonjs.com/api/classes/babylon.inputmanager#doubleclickdelay
func (i *InputManager) SetDoubleClickDelay(DoubleClickDelay float64) *InputManager {
	p := ba.ctx.Get("InputManager").New(DoubleClickDelay)
	return InputManagerFromJSObject(p, ba.ctx)
}

// DragMovementThreshold returns the DragMovementThreshold property of class InputManager.
//
// https://doc.babylonjs.com/api/classes/babylon.inputmanager#dragmovementthreshold
func (i *InputManager) DragMovementThreshold(DragMovementThreshold float64) *InputManager {
	p := ba.ctx.Get("InputManager").New(DragMovementThreshold)
	return InputManagerFromJSObject(p, ba.ctx)
}

// SetDragMovementThreshold sets the DragMovementThreshold property of class InputManager.
//
// https://doc.babylonjs.com/api/classes/babylon.inputmanager#dragmovementthreshold
func (i *InputManager) SetDragMovementThreshold(DragMovementThreshold float64) *InputManager {
	p := ba.ctx.Get("InputManager").New(DragMovementThreshold)
	return InputManagerFromJSObject(p, ba.ctx)
}

// ExclusiveDoubleClickMode returns the ExclusiveDoubleClickMode property of class InputManager.
//
// https://doc.babylonjs.com/api/classes/babylon.inputmanager#exclusivedoubleclickmode
func (i *InputManager) ExclusiveDoubleClickMode(ExclusiveDoubleClickMode bool) *InputManager {
	p := ba.ctx.Get("InputManager").New(ExclusiveDoubleClickMode)
	return InputManagerFromJSObject(p, ba.ctx)
}

// SetExclusiveDoubleClickMode sets the ExclusiveDoubleClickMode property of class InputManager.
//
// https://doc.babylonjs.com/api/classes/babylon.inputmanager#exclusivedoubleclickmode
func (i *InputManager) SetExclusiveDoubleClickMode(ExclusiveDoubleClickMode bool) *InputManager {
	p := ba.ctx.Get("InputManager").New(ExclusiveDoubleClickMode)
	return InputManagerFromJSObject(p, ba.ctx)
}

// LongPressDelay returns the LongPressDelay property of class InputManager.
//
// https://doc.babylonjs.com/api/classes/babylon.inputmanager#longpressdelay
func (i *InputManager) LongPressDelay(LongPressDelay float64) *InputManager {
	p := ba.ctx.Get("InputManager").New(LongPressDelay)
	return InputManagerFromJSObject(p, ba.ctx)
}

// SetLongPressDelay sets the LongPressDelay property of class InputManager.
//
// https://doc.babylonjs.com/api/classes/babylon.inputmanager#longpressdelay
func (i *InputManager) SetLongPressDelay(LongPressDelay float64) *InputManager {
	p := ba.ctx.Get("InputManager").New(LongPressDelay)
	return InputManagerFromJSObject(p, ba.ctx)
}

// MeshUnderPointer returns the MeshUnderPointer property of class InputManager.
//
// https://doc.babylonjs.com/api/classes/babylon.inputmanager#meshunderpointer
func (i *InputManager) MeshUnderPointer(meshUnderPointer *AbstractMesh) *InputManager {
	p := ba.ctx.Get("InputManager").New(meshUnderPointer.JSObject())
	return InputManagerFromJSObject(p, ba.ctx)
}

// SetMeshUnderPointer sets the MeshUnderPointer property of class InputManager.
//
// https://doc.babylonjs.com/api/classes/babylon.inputmanager#meshunderpointer
func (i *InputManager) SetMeshUnderPointer(meshUnderPointer *AbstractMesh) *InputManager {
	p := ba.ctx.Get("InputManager").New(meshUnderPointer.JSObject())
	return InputManagerFromJSObject(p, ba.ctx)
}

// PointerX returns the PointerX property of class InputManager.
//
// https://doc.babylonjs.com/api/classes/babylon.inputmanager#pointerx
func (i *InputManager) PointerX(pointerX float64) *InputManager {
	p := ba.ctx.Get("InputManager").New(pointerX)
	return InputManagerFromJSObject(p, ba.ctx)
}

// SetPointerX sets the PointerX property of class InputManager.
//
// https://doc.babylonjs.com/api/classes/babylon.inputmanager#pointerx
func (i *InputManager) SetPointerX(pointerX float64) *InputManager {
	p := ba.ctx.Get("InputManager").New(pointerX)
	return InputManagerFromJSObject(p, ba.ctx)
}

// PointerY returns the PointerY property of class InputManager.
//
// https://doc.babylonjs.com/api/classes/babylon.inputmanager#pointery
func (i *InputManager) PointerY(pointerY float64) *InputManager {
	p := ba.ctx.Get("InputManager").New(pointerY)
	return InputManagerFromJSObject(p, ba.ctx)
}

// SetPointerY sets the PointerY property of class InputManager.
//
// https://doc.babylonjs.com/api/classes/babylon.inputmanager#pointery
func (i *InputManager) SetPointerY(pointerY float64) *InputManager {
	p := ba.ctx.Get("InputManager").New(pointerY)
	return InputManagerFromJSObject(p, ba.ctx)
}

// UnTranslatedPointer returns the UnTranslatedPointer property of class InputManager.
//
// https://doc.babylonjs.com/api/classes/babylon.inputmanager#untranslatedpointer
func (i *InputManager) UnTranslatedPointer(unTranslatedPointer *Vector2) *InputManager {
	p := ba.ctx.Get("InputManager").New(unTranslatedPointer.JSObject())
	return InputManagerFromJSObject(p, ba.ctx)
}

// SetUnTranslatedPointer sets the UnTranslatedPointer property of class InputManager.
//
// https://doc.babylonjs.com/api/classes/babylon.inputmanager#untranslatedpointer
func (i *InputManager) SetUnTranslatedPointer(unTranslatedPointer *Vector2) *InputManager {
	p := ba.ctx.Get("InputManager").New(unTranslatedPointer.JSObject())
	return InputManagerFromJSObject(p, ba.ctx)
}

*/
