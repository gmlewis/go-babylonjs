// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ScatterPanel represents a babylon.js ScatterPanel.
// Class used to create a container panel where items get randomized planar mapping
type ScatterPanel struct {
	*VolumeBasedPanel
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *ScatterPanel) JSObject() js.Value { return s.p }

// ScatterPanel returns a ScatterPanel JavaScript class.
func (ba *Babylon) ScatterPanel() *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel")
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// ScatterPanelFromJSObject returns a wrapped ScatterPanel JavaScript class.
func ScatterPanelFromJSObject(p js.Value, ctx js.Value) *ScatterPanel {
	return &ScatterPanel{VolumeBasedPanel: VolumeBasedPanelFromJSObject(p, ctx), ctx: ctx}
}

// NewScatterPanel returns a new ScatterPanel object.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel
func (ba *Babylon) NewScatterPanel() *ScatterPanel {

	args := make([]interface{}, 0, 0+0)

	p := ba.ctx.Get("ScatterPanel").New(args...)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// AddBehavior calls the AddBehavior method on the ScatterPanel object.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#addbehavior
func (s *ScatterPanel) AddBehavior(behavior js.Value) *Control3D {

	args := make([]interface{}, 0, 1+0)

	args = append(args, behavior)

	retVal := s.p.Call("addBehavior", args...)
	return Control3DFromJSObject(retVal, s.ctx)
}

// AddControl calls the AddControl method on the ScatterPanel object.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#addcontrol
func (s *ScatterPanel) AddControl(control *Control3D) *Container3D {

	args := make([]interface{}, 0, 1+0)

	args = append(args, control.JSObject())

	retVal := s.p.Call("addControl", args...)
	return Container3DFromJSObject(retVal, s.ctx)
}

// ContainsControl calls the ContainsControl method on the ScatterPanel object.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#containscontrol
func (s *ScatterPanel) ContainsControl(control *Control3D) bool {

	args := make([]interface{}, 0, 1+0)

	args = append(args, control.JSObject())

	retVal := s.p.Call("containsControl", args...)
	return retVal.Bool()
}

// Dispose calls the Dispose method on the ScatterPanel object.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#dispose
func (s *ScatterPanel) Dispose() {

	args := make([]interface{}, 0, 0+0)

	s.p.Call("dispose", args...)
}

// GetBehaviorByName calls the GetBehaviorByName method on the ScatterPanel object.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#getbehaviorbyname
func (s *ScatterPanel) GetBehaviorByName(name string) *Control3D {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := s.p.Call("getBehaviorByName", args...)
	return Control3DFromJSObject(retVal, s.ctx)
}

// GetClassName calls the GetClassName method on the ScatterPanel object.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#getclassname
func (s *ScatterPanel) GetClassName() string {

	args := make([]interface{}, 0, 0+0)

	retVal := s.p.Call("getClassName", args...)
	return retVal.String()
}

// LinkToTransformNode calls the LinkToTransformNode method on the ScatterPanel object.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#linktotransformnode
func (s *ScatterPanel) LinkToTransformNode(node *TransformNode) *Control3D {

	args := make([]interface{}, 0, 1+0)

	args = append(args, node.JSObject())

	retVal := s.p.Call("linkToTransformNode", args...)
	return Control3DFromJSObject(retVal, s.ctx)
}

// RemoveBehavior calls the RemoveBehavior method on the ScatterPanel object.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#removebehavior
func (s *ScatterPanel) RemoveBehavior(behavior js.Value) *Control3D {

	args := make([]interface{}, 0, 1+0)

	args = append(args, behavior)

	retVal := s.p.Call("removeBehavior", args...)
	return Control3DFromJSObject(retVal, s.ctx)
}

// RemoveControl calls the RemoveControl method on the ScatterPanel object.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#removecontrol
func (s *ScatterPanel) RemoveControl(control *Control3D) *Container3D {

	args := make([]interface{}, 0, 1+0)

	args = append(args, control.JSObject())

	retVal := s.p.Call("removeControl", args...)
	return Container3DFromJSObject(retVal, s.ctx)
}

// UpdateLayout calls the UpdateLayout method on the ScatterPanel object.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#updatelayout
func (s *ScatterPanel) UpdateLayout() *Container3D {

	args := make([]interface{}, 0, 0+0)

	retVal := s.p.Call("updateLayout", args...)
	return Container3DFromJSObject(retVal, s.ctx)
}

/*

// Behaviors returns the Behaviors property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#behaviors
func (s *ScatterPanel) Behaviors(behaviors js.Value) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(behaviors)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// SetBehaviors sets the Behaviors property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#behaviors
func (s *ScatterPanel) SetBehaviors(behaviors js.Value) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(behaviors)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// BlockLayout returns the BlockLayout property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#blocklayout
func (s *ScatterPanel) BlockLayout(blockLayout bool) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(blockLayout)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// SetBlockLayout sets the BlockLayout property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#blocklayout
func (s *ScatterPanel) SetBlockLayout(blockLayout bool) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(blockLayout)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// Children returns the Children property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#children
func (s *ScatterPanel) Children(children []Control3D) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(children.JSObject())
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// SetChildren sets the Children property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#children
func (s *ScatterPanel) SetChildren(children []Control3D) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(children.JSObject())
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// Columns returns the Columns property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#columns
func (s *ScatterPanel) Columns(columns *int) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(columns.JSObject())
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// SetColumns sets the Columns property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#columns
func (s *ScatterPanel) SetColumns(columns *int) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(columns.JSObject())
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// FACEFORWARDREVERSED_ORIENTATION returns the FACEFORWARDREVERSED_ORIENTATION property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#faceforwardreversed_orientation
func (s *ScatterPanel) FACEFORWARDREVERSED_ORIENTATION(FACEFORWARDREVERSED_ORIENTATION float64) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(FACEFORWARDREVERSED_ORIENTATION)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// SetFACEFORWARDREVERSED_ORIENTATION sets the FACEFORWARDREVERSED_ORIENTATION property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#faceforwardreversed_orientation
func (s *ScatterPanel) SetFACEFORWARDREVERSED_ORIENTATION(FACEFORWARDREVERSED_ORIENTATION float64) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(FACEFORWARDREVERSED_ORIENTATION)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// FACEFORWARD_ORIENTATION returns the FACEFORWARD_ORIENTATION property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#faceforward_orientation
func (s *ScatterPanel) FACEFORWARD_ORIENTATION(FACEFORWARD_ORIENTATION float64) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(FACEFORWARD_ORIENTATION)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// SetFACEFORWARD_ORIENTATION sets the FACEFORWARD_ORIENTATION property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#faceforward_orientation
func (s *ScatterPanel) SetFACEFORWARD_ORIENTATION(FACEFORWARD_ORIENTATION float64) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(FACEFORWARD_ORIENTATION)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// FACEORIGINREVERSED_ORIENTATION returns the FACEORIGINREVERSED_ORIENTATION property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#faceoriginreversed_orientation
func (s *ScatterPanel) FACEORIGINREVERSED_ORIENTATION(FACEORIGINREVERSED_ORIENTATION float64) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(FACEORIGINREVERSED_ORIENTATION)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// SetFACEORIGINREVERSED_ORIENTATION sets the FACEORIGINREVERSED_ORIENTATION property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#faceoriginreversed_orientation
func (s *ScatterPanel) SetFACEORIGINREVERSED_ORIENTATION(FACEORIGINREVERSED_ORIENTATION float64) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(FACEORIGINREVERSED_ORIENTATION)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// FACEORIGIN_ORIENTATION returns the FACEORIGIN_ORIENTATION property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#faceorigin_orientation
func (s *ScatterPanel) FACEORIGIN_ORIENTATION(FACEORIGIN_ORIENTATION float64) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(FACEORIGIN_ORIENTATION)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// SetFACEORIGIN_ORIENTATION sets the FACEORIGIN_ORIENTATION property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#faceorigin_orientation
func (s *ScatterPanel) SetFACEORIGIN_ORIENTATION(FACEORIGIN_ORIENTATION float64) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(FACEORIGIN_ORIENTATION)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// IsVisible returns the IsVisible property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#isvisible
func (s *ScatterPanel) IsVisible(isVisible bool) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(isVisible)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// SetIsVisible sets the IsVisible property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#isvisible
func (s *ScatterPanel) SetIsVisible(isVisible bool) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(isVisible)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// Iteration returns the Iteration property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#iteration
func (s *ScatterPanel) Iteration(iteration *float) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(iteration.JSObject())
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// SetIteration sets the Iteration property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#iteration
func (s *ScatterPanel) SetIteration(iteration *float) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(iteration.JSObject())
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// Margin returns the Margin property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#margin
func (s *ScatterPanel) Margin(margin float64) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(margin)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// SetMargin sets the Margin property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#margin
func (s *ScatterPanel) SetMargin(margin float64) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(margin)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// Mesh returns the Mesh property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#mesh
func (s *ScatterPanel) Mesh(mesh *AbstractMesh) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(mesh.JSObject())
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// SetMesh sets the Mesh property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#mesh
func (s *ScatterPanel) SetMesh(mesh *AbstractMesh) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(mesh.JSObject())
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#name
func (s *ScatterPanel) Name(name string) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(name)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#name
func (s *ScatterPanel) SetName(name string) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(name)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// Node returns the Node property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#node
func (s *ScatterPanel) Node(node *TransformNode) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(node.JSObject())
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// SetNode sets the Node property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#node
func (s *ScatterPanel) SetNode(node *TransformNode) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(node.JSObject())
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// OnPointerClickObservable returns the OnPointerClickObservable property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#onpointerclickobservable
func (s *ScatterPanel) OnPointerClickObservable(onPointerClickObservable *Observable) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(onPointerClickObservable.JSObject())
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// SetOnPointerClickObservable sets the OnPointerClickObservable property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#onpointerclickobservable
func (s *ScatterPanel) SetOnPointerClickObservable(onPointerClickObservable *Observable) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(onPointerClickObservable.JSObject())
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// OnPointerDownObservable returns the OnPointerDownObservable property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#onpointerdownobservable
func (s *ScatterPanel) OnPointerDownObservable(onPointerDownObservable *Observable) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(onPointerDownObservable.JSObject())
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// SetOnPointerDownObservable sets the OnPointerDownObservable property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#onpointerdownobservable
func (s *ScatterPanel) SetOnPointerDownObservable(onPointerDownObservable *Observable) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(onPointerDownObservable.JSObject())
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// OnPointerEnterObservable returns the OnPointerEnterObservable property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#onpointerenterobservable
func (s *ScatterPanel) OnPointerEnterObservable(onPointerEnterObservable *Observable) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(onPointerEnterObservable.JSObject())
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// SetOnPointerEnterObservable sets the OnPointerEnterObservable property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#onpointerenterobservable
func (s *ScatterPanel) SetOnPointerEnterObservable(onPointerEnterObservable *Observable) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(onPointerEnterObservable.JSObject())
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// OnPointerMoveObservable returns the OnPointerMoveObservable property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#onpointermoveobservable
func (s *ScatterPanel) OnPointerMoveObservable(onPointerMoveObservable *Observable) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(onPointerMoveObservable.JSObject())
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// SetOnPointerMoveObservable sets the OnPointerMoveObservable property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#onpointermoveobservable
func (s *ScatterPanel) SetOnPointerMoveObservable(onPointerMoveObservable *Observable) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(onPointerMoveObservable.JSObject())
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// OnPointerOutObservable returns the OnPointerOutObservable property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#onpointeroutobservable
func (s *ScatterPanel) OnPointerOutObservable(onPointerOutObservable *Observable) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(onPointerOutObservable.JSObject())
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// SetOnPointerOutObservable sets the OnPointerOutObservable property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#onpointeroutobservable
func (s *ScatterPanel) SetOnPointerOutObservable(onPointerOutObservable *Observable) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(onPointerOutObservable.JSObject())
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// OnPointerUpObservable returns the OnPointerUpObservable property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#onpointerupobservable
func (s *ScatterPanel) OnPointerUpObservable(onPointerUpObservable *Observable) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(onPointerUpObservable.JSObject())
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// SetOnPointerUpObservable sets the OnPointerUpObservable property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#onpointerupobservable
func (s *ScatterPanel) SetOnPointerUpObservable(onPointerUpObservable *Observable) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(onPointerUpObservable.JSObject())
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// Orientation returns the Orientation property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#orientation
func (s *ScatterPanel) Orientation(orientation float64) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(orientation)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// SetOrientation sets the Orientation property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#orientation
func (s *ScatterPanel) SetOrientation(orientation float64) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(orientation)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// Parent returns the Parent property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#parent
func (s *ScatterPanel) Parent(parent *Container3D) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(parent.JSObject())
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// SetParent sets the Parent property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#parent
func (s *ScatterPanel) SetParent(parent *Container3D) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(parent.JSObject())
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// PointerDownAnimation returns the PointerDownAnimation property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#pointerdownanimation
func (s *ScatterPanel) PointerDownAnimation(pointerDownAnimation func()) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(pointerDownAnimation)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// SetPointerDownAnimation sets the PointerDownAnimation property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#pointerdownanimation
func (s *ScatterPanel) SetPointerDownAnimation(pointerDownAnimation func()) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(pointerDownAnimation)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// PointerEnterAnimation returns the PointerEnterAnimation property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#pointerenteranimation
func (s *ScatterPanel) PointerEnterAnimation(pointerEnterAnimation func()) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(pointerEnterAnimation)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// SetPointerEnterAnimation sets the PointerEnterAnimation property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#pointerenteranimation
func (s *ScatterPanel) SetPointerEnterAnimation(pointerEnterAnimation func()) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(pointerEnterAnimation)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// PointerOutAnimation returns the PointerOutAnimation property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#pointeroutanimation
func (s *ScatterPanel) PointerOutAnimation(pointerOutAnimation func()) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(pointerOutAnimation)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// SetPointerOutAnimation sets the PointerOutAnimation property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#pointeroutanimation
func (s *ScatterPanel) SetPointerOutAnimation(pointerOutAnimation func()) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(pointerOutAnimation)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// PointerUpAnimation returns the PointerUpAnimation property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#pointerupanimation
func (s *ScatterPanel) PointerUpAnimation(pointerUpAnimation func()) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(pointerUpAnimation)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// SetPointerUpAnimation sets the PointerUpAnimation property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#pointerupanimation
func (s *ScatterPanel) SetPointerUpAnimation(pointerUpAnimation func()) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(pointerUpAnimation)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// Position returns the Position property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#position
func (s *ScatterPanel) Position(position *Vector3) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(position.JSObject())
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// SetPosition sets the Position property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#position
func (s *ScatterPanel) SetPosition(position *Vector3) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(position.JSObject())
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// Rows returns the Rows property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#rows
func (s *ScatterPanel) Rows(rows *int) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(rows.JSObject())
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// SetRows sets the Rows property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#rows
func (s *ScatterPanel) SetRows(rows *int) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(rows.JSObject())
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// Scaling returns the Scaling property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#scaling
func (s *ScatterPanel) Scaling(scaling *Vector3) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(scaling.JSObject())
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// SetScaling sets the Scaling property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#scaling
func (s *ScatterPanel) SetScaling(scaling *Vector3) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(scaling.JSObject())
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// TypeName returns the TypeName property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#typename
func (s *ScatterPanel) TypeName(typeName string) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(typeName)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// SetTypeName sets the TypeName property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#typename
func (s *ScatterPanel) SetTypeName(typeName string) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(typeName)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// UNSET_ORIENTATION returns the UNSET_ORIENTATION property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#unset_orientation
func (s *ScatterPanel) UNSET_ORIENTATION(UNSET_ORIENTATION float64) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(UNSET_ORIENTATION)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// SetUNSET_ORIENTATION sets the UNSET_ORIENTATION property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#unset_orientation
func (s *ScatterPanel) SetUNSET_ORIENTATION(UNSET_ORIENTATION float64) *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel").New(UNSET_ORIENTATION)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

*/
