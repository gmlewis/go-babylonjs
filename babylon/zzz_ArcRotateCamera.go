// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ArcRotateCamera represents a babylon.js ArcRotateCamera.
// This represents an orbital type of camera.
//
// This camera always points towards a given target position and can be rotated around that target with the target as the centre of rotation. It can be controlled with cursors and mouse, or with touch events.
// Think of this camera as one orbiting its target position, or more imaginatively as a spy satellite orbiting the earth. Its position relative to the target (earth) can be set by three parameters, alpha (radians) the longitudinal rotation, beta (radians) the latitudinal rotation and radius the distance from the target position.
//
// See: http://doc.babylonjs.com/babylon101/cameras#arc-rotate-camera
type ArcRotateCamera struct {
	*TargetCamera
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *ArcRotateCamera) JSObject() js.Value { return a.p }

// ArcRotateCamera returns a ArcRotateCamera JavaScript class.
func (ba *Babylon) ArcRotateCamera() *ArcRotateCamera {
	p := ba.ctx.Get("ArcRotateCamera")
	return ArcRotateCameraFromJSObject(p, ba.ctx)
}

// ArcRotateCameraFromJSObject returns a wrapped ArcRotateCamera JavaScript class.
func ArcRotateCameraFromJSObject(p js.Value, ctx js.Value) *ArcRotateCamera {
	return &ArcRotateCamera{TargetCamera: TargetCameraFromJSObject(p, ctx), ctx: ctx}
}

// ArcRotateCameraArrayToJSArray returns a JavaScript Array for the wrapped array.
func ArcRotateCameraArrayToJSArray(array []*ArcRotateCamera) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewArcRotateCameraOpts contains optional parameters for NewArcRotateCamera.
type NewArcRotateCameraOpts struct {
	SetActiveOnSceneIfNoneActive *bool
}

// NewArcRotateCamera returns a new ArcRotateCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#constructor
func (ba *Babylon) NewArcRotateCamera(name string, alpha float64, beta float64, radius float64, target *Vector3, scene *Scene, opts *NewArcRotateCameraOpts) *ArcRotateCamera {
	if opts == nil {
		opts = &NewArcRotateCameraOpts{}
	}

	args := make([]interface{}, 0, 6+1)

	args = append(args, name)
	args = append(args, alpha)
	args = append(args, beta)
	args = append(args, radius)
	args = append(args, target.JSObject())
	args = append(args, scene.JSObject())

	if opts.SetActiveOnSceneIfNoneActive == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.SetActiveOnSceneIfNoneActive)
	}

	p := ba.ctx.Get("ArcRotateCamera").New(args...)
	return ArcRotateCameraFromJSObject(p, ba.ctx)
}

// ArcRotateCameraAttachControlOpts contains optional parameters for ArcRotateCamera.AttachControl.
type ArcRotateCameraAttachControlOpts struct {
	UseCtrlForPanning  *bool
	PanningMouseButton *float64
}

// AttachControl calls the AttachControl method on the ArcRotateCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#attachcontrol
func (a *ArcRotateCamera) AttachControl(element js.Value, noPreventDefault bool, opts *ArcRotateCameraAttachControlOpts) {
	if opts == nil {
		opts = &ArcRotateCameraAttachControlOpts{}
	}

	args := make([]interface{}, 0, 2+2)

	args = append(args, element)

	args = append(args, noPreventDefault)

	if opts.UseCtrlForPanning == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UseCtrlForPanning)
	}
	if opts.PanningMouseButton == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.PanningMouseButton)
	}

	a.p.Call("attachControl", args...)
}

// CreateRigCamera calls the CreateRigCamera method on the ArcRotateCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#createrigcamera
func (a *ArcRotateCamera) CreateRigCamera(name string, cameraIndex float64) *Camera {

	args := make([]interface{}, 0, 2+0)

	args = append(args, name)

	args = append(args, cameraIndex)

	retVal := a.p.Call("createRigCamera", args...)
	return CameraFromJSObject(retVal, a.ctx)
}

// DetachControl calls the DetachControl method on the ArcRotateCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#detachcontrol
func (a *ArcRotateCamera) DetachControl(element js.Value) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, element)

	a.p.Call("detachControl", args...)
}

// Dispose calls the Dispose method on the ArcRotateCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#dispose
func (a *ArcRotateCamera) Dispose() {

	a.p.Call("dispose")
}

// ArcRotateCameraFocusOnOpts contains optional parameters for ArcRotateCamera.FocusOn.
type ArcRotateCameraFocusOnOpts struct {
	DoNotUpdateMaxZ *bool
}

// FocusOn calls the FocusOn method on the ArcRotateCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#focuson
func (a *ArcRotateCamera) FocusOn(meshesOrMinMaxVectorAndDistance []*AbstractMesh, opts *ArcRotateCameraFocusOnOpts) {
	if opts == nil {
		opts = &ArcRotateCameraFocusOnOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, AbstractMeshArrayToJSArray(meshesOrMinMaxVectorAndDistance))

	if opts.DoNotUpdateMaxZ == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.DoNotUpdateMaxZ)
	}

	a.p.Call("focusOn", args...)
}

// GetClassName calls the GetClassName method on the ArcRotateCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#getclassname
func (a *ArcRotateCamera) GetClassName() string {

	retVal := a.p.Call("getClassName")
	return retVal.String()
}

// RebuildAnglesAndRadius calls the RebuildAnglesAndRadius method on the ArcRotateCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#rebuildanglesandradius
func (a *ArcRotateCamera) RebuildAnglesAndRadius() {

	a.p.Call("rebuildAnglesAndRadius")
}

// SetMatUp calls the SetMatUp method on the ArcRotateCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#setmatup
func (a *ArcRotateCamera) SetMatUp() {

	a.p.Call("setMatUp")
}

// SetPosition calls the SetPosition method on the ArcRotateCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#setposition
func (a *ArcRotateCamera) SetPosition(position *Vector3) {

	args := make([]interface{}, 0, 1+0)

	if position == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, position.JSObject())
	}

	a.p.Call("setPosition", args...)
}

// ArcRotateCameraSetTargetOpts contains optional parameters for ArcRotateCamera.SetTarget.
type ArcRotateCameraSetTargetOpts struct {
	ToBoundingCenter  *bool
	AllowSamePosition *bool
}

// SetTarget calls the SetTarget method on the ArcRotateCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#settarget
func (a *ArcRotateCamera) SetTarget(target *AbstractMesh, opts *ArcRotateCameraSetTargetOpts) {
	if opts == nil {
		opts = &ArcRotateCameraSetTargetOpts{}
	}

	args := make([]interface{}, 0, 1+2)

	if target == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, target.JSObject())
	}

	if opts.ToBoundingCenter == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ToBoundingCenter)
	}
	if opts.AllowSamePosition == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.AllowSamePosition)
	}

	a.p.Call("setTarget", args...)
}

// StoreState calls the StoreState method on the ArcRotateCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#storestate
func (a *ArcRotateCamera) StoreState() *Camera {

	retVal := a.p.Call("storeState")
	return CameraFromJSObject(retVal, a.ctx)
}

// ArcRotateCameraZoomOnOpts contains optional parameters for ArcRotateCamera.ZoomOn.
type ArcRotateCameraZoomOnOpts struct {
	Meshes          []*AbstractMesh
	DoNotUpdateMaxZ *bool
}

// ZoomOn calls the ZoomOn method on the ArcRotateCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#zoomon
func (a *ArcRotateCamera) ZoomOn(opts *ArcRotateCameraZoomOnOpts) {
	if opts == nil {
		opts = &ArcRotateCameraZoomOnOpts{}
	}

	args := make([]interface{}, 0, 0+2)

	if opts.Meshes == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, AbstractMeshArrayToJSArray(opts.Meshes))
	}
	if opts.DoNotUpdateMaxZ == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.DoNotUpdateMaxZ)
	}

	a.p.Call("zoomOn", args...)
}

// AllowUpsideDown returns the AllowUpsideDown property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#allowupsidedown
func (a *ArcRotateCamera) AllowUpsideDown() bool {
	retVal := a.p.Get("allowUpsideDown")
	return retVal.Bool()
}

// SetAllowUpsideDown sets the AllowUpsideDown property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#allowupsidedown
func (a *ArcRotateCamera) SetAllowUpsideDown(allowUpsideDown bool) *ArcRotateCamera {
	a.p.Set("allowUpsideDown", allowUpsideDown)
	return a
}

// Alpha returns the Alpha property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#alpha
func (a *ArcRotateCamera) Alpha() float64 {
	retVal := a.p.Get("alpha")
	return retVal.Float()
}

// SetAlpha sets the Alpha property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#alpha
func (a *ArcRotateCamera) SetAlpha(alpha float64) *ArcRotateCamera {
	a.p.Set("alpha", alpha)
	return a
}

// AngularSensibilityX returns the AngularSensibilityX property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#angularsensibilityx
func (a *ArcRotateCamera) AngularSensibilityX() float64 {
	retVal := a.p.Get("angularSensibilityX")
	return retVal.Float()
}

// SetAngularSensibilityX sets the AngularSensibilityX property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#angularsensibilityx
func (a *ArcRotateCamera) SetAngularSensibilityX(angularSensibilityX float64) *ArcRotateCamera {
	a.p.Set("angularSensibilityX", angularSensibilityX)
	return a
}

// AngularSensibilityY returns the AngularSensibilityY property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#angularsensibilityy
func (a *ArcRotateCamera) AngularSensibilityY() float64 {
	retVal := a.p.Get("angularSensibilityY")
	return retVal.Float()
}

// SetAngularSensibilityY sets the AngularSensibilityY property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#angularsensibilityy
func (a *ArcRotateCamera) SetAngularSensibilityY(angularSensibilityY float64) *ArcRotateCamera {
	a.p.Set("angularSensibilityY", angularSensibilityY)
	return a
}

// AutoRotationBehavior returns the AutoRotationBehavior property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#autorotationbehavior
func (a *ArcRotateCamera) AutoRotationBehavior() *AutoRotationBehavior {
	retVal := a.p.Get("autoRotationBehavior")
	return AutoRotationBehaviorFromJSObject(retVal, a.ctx)
}

// SetAutoRotationBehavior sets the AutoRotationBehavior property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#autorotationbehavior
func (a *ArcRotateCamera) SetAutoRotationBehavior(autoRotationBehavior *AutoRotationBehavior) *ArcRotateCamera {
	a.p.Set("autoRotationBehavior", autoRotationBehavior.JSObject())
	return a
}

// Beta returns the Beta property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#beta
func (a *ArcRotateCamera) Beta() float64 {
	retVal := a.p.Get("beta")
	return retVal.Float()
}

// SetBeta sets the Beta property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#beta
func (a *ArcRotateCamera) SetBeta(beta float64) *ArcRotateCamera {
	a.p.Set("beta", beta)
	return a
}

// BouncingBehavior returns the BouncingBehavior property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#bouncingbehavior
func (a *ArcRotateCamera) BouncingBehavior() *BouncingBehavior {
	retVal := a.p.Get("bouncingBehavior")
	return BouncingBehaviorFromJSObject(retVal, a.ctx)
}

// SetBouncingBehavior sets the BouncingBehavior property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#bouncingbehavior
func (a *ArcRotateCamera) SetBouncingBehavior(bouncingBehavior *BouncingBehavior) *ArcRotateCamera {
	a.p.Set("bouncingBehavior", bouncingBehavior.JSObject())
	return a
}

// CheckCollisions returns the CheckCollisions property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#checkcollisions
func (a *ArcRotateCamera) CheckCollisions() bool {
	retVal := a.p.Get("checkCollisions")
	return retVal.Bool()
}

// SetCheckCollisions sets the CheckCollisions property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#checkcollisions
func (a *ArcRotateCamera) SetCheckCollisions(checkCollisions bool) *ArcRotateCamera {
	a.p.Set("checkCollisions", checkCollisions)
	return a
}

// CollisionRadius returns the CollisionRadius property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#collisionradius
func (a *ArcRotateCamera) CollisionRadius() *Vector3 {
	retVal := a.p.Get("collisionRadius")
	return Vector3FromJSObject(retVal, a.ctx)
}

// SetCollisionRadius sets the CollisionRadius property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#collisionradius
func (a *ArcRotateCamera) SetCollisionRadius(collisionRadius *Vector3) *ArcRotateCamera {
	a.p.Set("collisionRadius", collisionRadius.JSObject())
	return a
}

// FramingBehavior returns the FramingBehavior property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#framingbehavior
func (a *ArcRotateCamera) FramingBehavior() *FramingBehavior {
	retVal := a.p.Get("framingBehavior")
	return FramingBehaviorFromJSObject(retVal, a.ctx)
}

// SetFramingBehavior sets the FramingBehavior property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#framingbehavior
func (a *ArcRotateCamera) SetFramingBehavior(framingBehavior *FramingBehavior) *ArcRotateCamera {
	a.p.Set("framingBehavior", framingBehavior.JSObject())
	return a
}

// InertialAlphaOffset returns the InertialAlphaOffset property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#inertialalphaoffset
func (a *ArcRotateCamera) InertialAlphaOffset() float64 {
	retVal := a.p.Get("inertialAlphaOffset")
	return retVal.Float()
}

// SetInertialAlphaOffset sets the InertialAlphaOffset property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#inertialalphaoffset
func (a *ArcRotateCamera) SetInertialAlphaOffset(inertialAlphaOffset float64) *ArcRotateCamera {
	a.p.Set("inertialAlphaOffset", inertialAlphaOffset)
	return a
}

// InertialBetaOffset returns the InertialBetaOffset property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#inertialbetaoffset
func (a *ArcRotateCamera) InertialBetaOffset() float64 {
	retVal := a.p.Get("inertialBetaOffset")
	return retVal.Float()
}

// SetInertialBetaOffset sets the InertialBetaOffset property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#inertialbetaoffset
func (a *ArcRotateCamera) SetInertialBetaOffset(inertialBetaOffset float64) *ArcRotateCamera {
	a.p.Set("inertialBetaOffset", inertialBetaOffset)
	return a
}

// InertialPanningX returns the InertialPanningX property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#inertialpanningx
func (a *ArcRotateCamera) InertialPanningX() float64 {
	retVal := a.p.Get("inertialPanningX")
	return retVal.Float()
}

// SetInertialPanningX sets the InertialPanningX property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#inertialpanningx
func (a *ArcRotateCamera) SetInertialPanningX(inertialPanningX float64) *ArcRotateCamera {
	a.p.Set("inertialPanningX", inertialPanningX)
	return a
}

// InertialPanningY returns the InertialPanningY property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#inertialpanningy
func (a *ArcRotateCamera) InertialPanningY() float64 {
	retVal := a.p.Get("inertialPanningY")
	return retVal.Float()
}

// SetInertialPanningY sets the InertialPanningY property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#inertialpanningy
func (a *ArcRotateCamera) SetInertialPanningY(inertialPanningY float64) *ArcRotateCamera {
	a.p.Set("inertialPanningY", inertialPanningY)
	return a
}

// InertialRadiusOffset returns the InertialRadiusOffset property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#inertialradiusoffset
func (a *ArcRotateCamera) InertialRadiusOffset() float64 {
	retVal := a.p.Get("inertialRadiusOffset")
	return retVal.Float()
}

// SetInertialRadiusOffset sets the InertialRadiusOffset property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#inertialradiusoffset
func (a *ArcRotateCamera) SetInertialRadiusOffset(inertialRadiusOffset float64) *ArcRotateCamera {
	a.p.Set("inertialRadiusOffset", inertialRadiusOffset)
	return a
}

// Inputs returns the Inputs property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#inputs
func (a *ArcRotateCamera) Inputs() *ArcRotateCameraInputsManager {
	retVal := a.p.Get("inputs")
	return ArcRotateCameraInputsManagerFromJSObject(retVal, a.ctx)
}

// SetInputs sets the Inputs property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#inputs
func (a *ArcRotateCamera) SetInputs(inputs *ArcRotateCameraInputsManager) *ArcRotateCamera {
	a.p.Set("inputs", inputs.JSObject())
	return a
}

// KeysDown returns the KeysDown property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#keysdown
func (a *ArcRotateCamera) KeysDown() []float64 {
	retVal := a.p.Get("keysDown")
	result := []float64{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, retVal.Index(ri).Float())
	}
	return result
}

// SetKeysDown sets the KeysDown property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#keysdown
func (a *ArcRotateCamera) SetKeysDown(keysDown []float64) *ArcRotateCamera {
	a.p.Set("keysDown", keysDown)
	return a
}

// KeysLeft returns the KeysLeft property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#keysleft
func (a *ArcRotateCamera) KeysLeft() []float64 {
	retVal := a.p.Get("keysLeft")
	result := []float64{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, retVal.Index(ri).Float())
	}
	return result
}

// SetKeysLeft sets the KeysLeft property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#keysleft
func (a *ArcRotateCamera) SetKeysLeft(keysLeft []float64) *ArcRotateCamera {
	a.p.Set("keysLeft", keysLeft)
	return a
}

// KeysRight returns the KeysRight property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#keysright
func (a *ArcRotateCamera) KeysRight() []float64 {
	retVal := a.p.Get("keysRight")
	result := []float64{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, retVal.Index(ri).Float())
	}
	return result
}

// SetKeysRight sets the KeysRight property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#keysright
func (a *ArcRotateCamera) SetKeysRight(keysRight []float64) *ArcRotateCamera {
	a.p.Set("keysRight", keysRight)
	return a
}

// KeysUp returns the KeysUp property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#keysup
func (a *ArcRotateCamera) KeysUp() []float64 {
	retVal := a.p.Get("keysUp")
	result := []float64{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, retVal.Index(ri).Float())
	}
	return result
}

// SetKeysUp sets the KeysUp property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#keysup
func (a *ArcRotateCamera) SetKeysUp(keysUp []float64) *ArcRotateCamera {
	a.p.Set("keysUp", keysUp)
	return a
}

// LowerAlphaLimit returns the LowerAlphaLimit property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#loweralphalimit
func (a *ArcRotateCamera) LowerAlphaLimit() float64 {
	retVal := a.p.Get("lowerAlphaLimit")
	return retVal.Float()
}

// SetLowerAlphaLimit sets the LowerAlphaLimit property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#loweralphalimit
func (a *ArcRotateCamera) SetLowerAlphaLimit(lowerAlphaLimit float64) *ArcRotateCamera {
	a.p.Set("lowerAlphaLimit", lowerAlphaLimit)
	return a
}

// LowerBetaLimit returns the LowerBetaLimit property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#lowerbetalimit
func (a *ArcRotateCamera) LowerBetaLimit() float64 {
	retVal := a.p.Get("lowerBetaLimit")
	return retVal.Float()
}

// SetLowerBetaLimit sets the LowerBetaLimit property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#lowerbetalimit
func (a *ArcRotateCamera) SetLowerBetaLimit(lowerBetaLimit float64) *ArcRotateCamera {
	a.p.Set("lowerBetaLimit", lowerBetaLimit)
	return a
}

// LowerRadiusLimit returns the LowerRadiusLimit property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#lowerradiuslimit
func (a *ArcRotateCamera) LowerRadiusLimit() float64 {
	retVal := a.p.Get("lowerRadiusLimit")
	return retVal.Float()
}

// SetLowerRadiusLimit sets the LowerRadiusLimit property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#lowerradiuslimit
func (a *ArcRotateCamera) SetLowerRadiusLimit(lowerRadiusLimit float64) *ArcRotateCamera {
	a.p.Set("lowerRadiusLimit", lowerRadiusLimit)
	return a
}

// OnCollide returns the OnCollide property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#oncollide
func (a *ArcRotateCamera) OnCollide() js.Value {
	retVal := a.p.Get("onCollide")
	return retVal
}

// SetOnCollide sets the OnCollide property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#oncollide
func (a *ArcRotateCamera) SetOnCollide(onCollide JSFunc) *ArcRotateCamera {
	a.p.Set("onCollide", js.FuncOf(onCollide))
	return a
}

// OnMeshTargetChangedObservable returns the OnMeshTargetChangedObservable property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#onmeshtargetchangedobservable
func (a *ArcRotateCamera) OnMeshTargetChangedObservable() *Observable {
	retVal := a.p.Get("onMeshTargetChangedObservable")
	return ObservableFromJSObject(retVal, a.ctx)
}

// SetOnMeshTargetChangedObservable sets the OnMeshTargetChangedObservable property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#onmeshtargetchangedobservable
func (a *ArcRotateCamera) SetOnMeshTargetChangedObservable(onMeshTargetChangedObservable *Observable) *ArcRotateCamera {
	a.p.Set("onMeshTargetChangedObservable", onMeshTargetChangedObservable.JSObject())
	return a
}

// PanningAxis returns the PanningAxis property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#panningaxis
func (a *ArcRotateCamera) PanningAxis() *Vector3 {
	retVal := a.p.Get("panningAxis")
	return Vector3FromJSObject(retVal, a.ctx)
}

// SetPanningAxis sets the PanningAxis property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#panningaxis
func (a *ArcRotateCamera) SetPanningAxis(panningAxis *Vector3) *ArcRotateCamera {
	a.p.Set("panningAxis", panningAxis.JSObject())
	return a
}

// PanningDistanceLimit returns the PanningDistanceLimit property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#panningdistancelimit
func (a *ArcRotateCamera) PanningDistanceLimit() float64 {
	retVal := a.p.Get("panningDistanceLimit")
	return retVal.Float()
}

// SetPanningDistanceLimit sets the PanningDistanceLimit property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#panningdistancelimit
func (a *ArcRotateCamera) SetPanningDistanceLimit(panningDistanceLimit float64) *ArcRotateCamera {
	a.p.Set("panningDistanceLimit", panningDistanceLimit)
	return a
}

// PanningInertia returns the PanningInertia property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#panninginertia
func (a *ArcRotateCamera) PanningInertia() float64 {
	retVal := a.p.Get("panningInertia")
	return retVal.Float()
}

// SetPanningInertia sets the PanningInertia property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#panninginertia
func (a *ArcRotateCamera) SetPanningInertia(panningInertia float64) *ArcRotateCamera {
	a.p.Set("panningInertia", panningInertia)
	return a
}

// PanningOriginTarget returns the PanningOriginTarget property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#panningorigintarget
func (a *ArcRotateCamera) PanningOriginTarget() *Vector3 {
	retVal := a.p.Get("panningOriginTarget")
	return Vector3FromJSObject(retVal, a.ctx)
}

// SetPanningOriginTarget sets the PanningOriginTarget property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#panningorigintarget
func (a *ArcRotateCamera) SetPanningOriginTarget(panningOriginTarget *Vector3) *ArcRotateCamera {
	a.p.Set("panningOriginTarget", panningOriginTarget.JSObject())
	return a
}

// PanningSensibility returns the PanningSensibility property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#panningsensibility
func (a *ArcRotateCamera) PanningSensibility() float64 {
	retVal := a.p.Get("panningSensibility")
	return retVal.Float()
}

// SetPanningSensibility sets the PanningSensibility property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#panningsensibility
func (a *ArcRotateCamera) SetPanningSensibility(panningSensibility float64) *ArcRotateCamera {
	a.p.Set("panningSensibility", panningSensibility)
	return a
}

// PinchDeltaPercentage returns the PinchDeltaPercentage property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#pinchdeltapercentage
func (a *ArcRotateCamera) PinchDeltaPercentage() float64 {
	retVal := a.p.Get("pinchDeltaPercentage")
	return retVal.Float()
}

// SetPinchDeltaPercentage sets the PinchDeltaPercentage property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#pinchdeltapercentage
func (a *ArcRotateCamera) SetPinchDeltaPercentage(pinchDeltaPercentage float64) *ArcRotateCamera {
	a.p.Set("pinchDeltaPercentage", pinchDeltaPercentage)
	return a
}

// PinchPrecision returns the PinchPrecision property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#pinchprecision
func (a *ArcRotateCamera) PinchPrecision() float64 {
	retVal := a.p.Get("pinchPrecision")
	return retVal.Float()
}

// SetPinchPrecision sets the PinchPrecision property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#pinchprecision
func (a *ArcRotateCamera) SetPinchPrecision(pinchPrecision float64) *ArcRotateCamera {
	a.p.Set("pinchPrecision", pinchPrecision)
	return a
}

// PinchToPanMaxDistance returns the PinchToPanMaxDistance property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#pinchtopanmaxdistance
func (a *ArcRotateCamera) PinchToPanMaxDistance() float64 {
	retVal := a.p.Get("pinchToPanMaxDistance")
	return retVal.Float()
}

// SetPinchToPanMaxDistance sets the PinchToPanMaxDistance property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#pinchtopanmaxdistance
func (a *ArcRotateCamera) SetPinchToPanMaxDistance(pinchToPanMaxDistance float64) *ArcRotateCamera {
	a.p.Set("pinchToPanMaxDistance", pinchToPanMaxDistance)
	return a
}

// Position returns the Position property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#position
func (a *ArcRotateCamera) Position() *Vector3 {
	retVal := a.p.Get("position")
	return Vector3FromJSObject(retVal, a.ctx)
}

// Radius returns the Radius property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#radius
func (a *ArcRotateCamera) Radius() float64 {
	retVal := a.p.Get("radius")
	return retVal.Float()
}

// SetRadius sets the Radius property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#radius
func (a *ArcRotateCamera) SetRadius(radius float64) *ArcRotateCamera {
	a.p.Set("radius", radius)
	return a
}

// Target returns the Target property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#target
func (a *ArcRotateCamera) Target() *Vector3 {
	retVal := a.p.Get("target")
	return Vector3FromJSObject(retVal, a.ctx)
}

// TargetScreenOffset returns the TargetScreenOffset property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#targetscreenoffset
func (a *ArcRotateCamera) TargetScreenOffset() *Vector2 {
	retVal := a.p.Get("targetScreenOffset")
	return Vector2FromJSObject(retVal, a.ctx)
}

// SetTargetScreenOffset sets the TargetScreenOffset property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#targetscreenoffset
func (a *ArcRotateCamera) SetTargetScreenOffset(targetScreenOffset *Vector2) *ArcRotateCamera {
	a.p.Set("targetScreenOffset", targetScreenOffset.JSObject())
	return a
}

// UpVector returns the UpVector property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#upvector
func (a *ArcRotateCamera) UpVector() *Vector3 {
	retVal := a.p.Get("upVector")
	return Vector3FromJSObject(retVal, a.ctx)
}

// SetUpVector sets the UpVector property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#upvector
func (a *ArcRotateCamera) SetUpVector(upVector *Vector3) *ArcRotateCamera {
	a.p.Set("upVector", upVector.JSObject())
	return a
}

// UpperAlphaLimit returns the UpperAlphaLimit property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#upperalphalimit
func (a *ArcRotateCamera) UpperAlphaLimit() float64 {
	retVal := a.p.Get("upperAlphaLimit")
	return retVal.Float()
}

// SetUpperAlphaLimit sets the UpperAlphaLimit property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#upperalphalimit
func (a *ArcRotateCamera) SetUpperAlphaLimit(upperAlphaLimit float64) *ArcRotateCamera {
	a.p.Set("upperAlphaLimit", upperAlphaLimit)
	return a
}

// UpperBetaLimit returns the UpperBetaLimit property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#upperbetalimit
func (a *ArcRotateCamera) UpperBetaLimit() float64 {
	retVal := a.p.Get("upperBetaLimit")
	return retVal.Float()
}

// SetUpperBetaLimit sets the UpperBetaLimit property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#upperbetalimit
func (a *ArcRotateCamera) SetUpperBetaLimit(upperBetaLimit float64) *ArcRotateCamera {
	a.p.Set("upperBetaLimit", upperBetaLimit)
	return a
}

// UpperRadiusLimit returns the UpperRadiusLimit property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#upperradiuslimit
func (a *ArcRotateCamera) UpperRadiusLimit() float64 {
	retVal := a.p.Get("upperRadiusLimit")
	return retVal.Float()
}

// SetUpperRadiusLimit sets the UpperRadiusLimit property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#upperradiuslimit
func (a *ArcRotateCamera) SetUpperRadiusLimit(upperRadiusLimit float64) *ArcRotateCamera {
	a.p.Set("upperRadiusLimit", upperRadiusLimit)
	return a
}

// UseAutoRotationBehavior returns the UseAutoRotationBehavior property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#useautorotationbehavior
func (a *ArcRotateCamera) UseAutoRotationBehavior() bool {
	retVal := a.p.Get("useAutoRotationBehavior")
	return retVal.Bool()
}

// SetUseAutoRotationBehavior sets the UseAutoRotationBehavior property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#useautorotationbehavior
func (a *ArcRotateCamera) SetUseAutoRotationBehavior(useAutoRotationBehavior bool) *ArcRotateCamera {
	a.p.Set("useAutoRotationBehavior", useAutoRotationBehavior)
	return a
}

// UseBouncingBehavior returns the UseBouncingBehavior property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#usebouncingbehavior
func (a *ArcRotateCamera) UseBouncingBehavior() bool {
	retVal := a.p.Get("useBouncingBehavior")
	return retVal.Bool()
}

// SetUseBouncingBehavior sets the UseBouncingBehavior property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#usebouncingbehavior
func (a *ArcRotateCamera) SetUseBouncingBehavior(useBouncingBehavior bool) *ArcRotateCamera {
	a.p.Set("useBouncingBehavior", useBouncingBehavior)
	return a
}

// UseFramingBehavior returns the UseFramingBehavior property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#useframingbehavior
func (a *ArcRotateCamera) UseFramingBehavior() bool {
	retVal := a.p.Get("useFramingBehavior")
	return retVal.Bool()
}

// SetUseFramingBehavior sets the UseFramingBehavior property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#useframingbehavior
func (a *ArcRotateCamera) SetUseFramingBehavior(useFramingBehavior bool) *ArcRotateCamera {
	a.p.Set("useFramingBehavior", useFramingBehavior)
	return a
}

// UseInputToRestoreState returns the UseInputToRestoreState property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#useinputtorestorestate
func (a *ArcRotateCamera) UseInputToRestoreState() bool {
	retVal := a.p.Get("useInputToRestoreState")
	return retVal.Bool()
}

// SetUseInputToRestoreState sets the UseInputToRestoreState property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#useinputtorestorestate
func (a *ArcRotateCamera) SetUseInputToRestoreState(useInputToRestoreState bool) *ArcRotateCamera {
	a.p.Set("useInputToRestoreState", useInputToRestoreState)
	return a
}

// WheelDeltaPercentage returns the WheelDeltaPercentage property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#wheeldeltapercentage
func (a *ArcRotateCamera) WheelDeltaPercentage() float64 {
	retVal := a.p.Get("wheelDeltaPercentage")
	return retVal.Float()
}

// SetWheelDeltaPercentage sets the WheelDeltaPercentage property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#wheeldeltapercentage
func (a *ArcRotateCamera) SetWheelDeltaPercentage(wheelDeltaPercentage float64) *ArcRotateCamera {
	a.p.Set("wheelDeltaPercentage", wheelDeltaPercentage)
	return a
}

// WheelPrecision returns the WheelPrecision property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#wheelprecision
func (a *ArcRotateCamera) WheelPrecision() float64 {
	retVal := a.p.Get("wheelPrecision")
	return retVal.Float()
}

// SetWheelPrecision sets the WheelPrecision property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#wheelprecision
func (a *ArcRotateCamera) SetWheelPrecision(wheelPrecision float64) *ArcRotateCamera {
	a.p.Set("wheelPrecision", wheelPrecision)
	return a
}

// ZoomOnFactor returns the ZoomOnFactor property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#zoomonfactor
func (a *ArcRotateCamera) ZoomOnFactor() float64 {
	retVal := a.p.Get("zoomOnFactor")
	return retVal.Float()
}

// SetZoomOnFactor sets the ZoomOnFactor property of class ArcRotateCamera.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecamera#zoomonfactor
func (a *ArcRotateCamera) SetZoomOnFactor(zoomOnFactor float64) *ArcRotateCamera {
	a.p.Set("zoomOnFactor", zoomOnFactor)
	return a
}
