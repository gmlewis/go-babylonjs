// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// HolographicButton represents a babylon.js HolographicButton.
// Class used to create a holographic button in 3D
type HolographicButton struct {
	*Button3D
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (h *HolographicButton) JSObject() js.Value { return h.p }

// HolographicButton returns a HolographicButton JavaScript class.
func (ba *Babylon) HolographicButton() *HolographicButton {
	p := ba.ctx.Get("HolographicButton")
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// HolographicButtonFromJSObject returns a wrapped HolographicButton JavaScript class.
func HolographicButtonFromJSObject(p js.Value, ctx js.Value) *HolographicButton {
	return &HolographicButton{Button3D: Button3DFromJSObject(p, ctx), ctx: ctx}
}

// NewHolographicButtonOpts contains optional parameters for NewHolographicButton.
type NewHolographicButtonOpts struct {
	Name           *string
	ShareMaterials *bool
}

// NewHolographicButton returns a new HolographicButton object.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton
func (ba *Babylon) NewHolographicButton(opts *NewHolographicButtonOpts) *HolographicButton {
	if opts == nil {
		opts = &NewHolographicButtonOpts{}
	}

	args := make([]interface{}, 0, 0+2)

	if opts.Name == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Name)
	}
	if opts.ShareMaterials == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ShareMaterials)
	}

	p := ba.ctx.Get("HolographicButton").New(args...)
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// AddBehavior calls the AddBehavior method on the HolographicButton object.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#addbehavior
func (h *HolographicButton) AddBehavior(behavior js.Value) *Control3D {

	args := make([]interface{}, 0, 1+0)

	args = append(args, behavior)

	retVal := h.p.Call("addBehavior", args...)
	return Control3DFromJSObject(retVal, h.ctx)
}

// Dispose calls the Dispose method on the HolographicButton object.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#dispose
func (h *HolographicButton) Dispose() {

	args := make([]interface{}, 0, 0+0)

	h.p.Call("dispose", args...)
}

// GetBehaviorByName calls the GetBehaviorByName method on the HolographicButton object.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#getbehaviorbyname
func (h *HolographicButton) GetBehaviorByName(name string) *Control3D {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := h.p.Call("getBehaviorByName", args...)
	return Control3DFromJSObject(retVal, h.ctx)
}

// GetClassName calls the GetClassName method on the HolographicButton object.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#getclassname
func (h *HolographicButton) GetClassName() string {

	args := make([]interface{}, 0, 0+0)

	retVal := h.p.Call("getClassName", args...)
	return retVal.String()
}

// LinkToTransformNode calls the LinkToTransformNode method on the HolographicButton object.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#linktotransformnode
func (h *HolographicButton) LinkToTransformNode(node *TransformNode) *Control3D {

	args := make([]interface{}, 0, 1+0)

	args = append(args, node.JSObject())

	retVal := h.p.Call("linkToTransformNode", args...)
	return Control3DFromJSObject(retVal, h.ctx)
}

// RemoveBehavior calls the RemoveBehavior method on the HolographicButton object.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#removebehavior
func (h *HolographicButton) RemoveBehavior(behavior js.Value) *Control3D {

	args := make([]interface{}, 0, 1+0)

	args = append(args, behavior)

	retVal := h.p.Call("removeBehavior", args...)
	return Control3DFromJSObject(retVal, h.ctx)
}

/*

// BackMaterial returns the BackMaterial property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#backmaterial
func (h *HolographicButton) BackMaterial(backMaterial *FluentMaterial) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(backMaterial.JSObject())
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// SetBackMaterial sets the BackMaterial property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#backmaterial
func (h *HolographicButton) SetBackMaterial(backMaterial *FluentMaterial) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(backMaterial.JSObject())
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// Behaviors returns the Behaviors property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#behaviors
func (h *HolographicButton) Behaviors(behaviors js.Value) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(behaviors)
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// SetBehaviors sets the Behaviors property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#behaviors
func (h *HolographicButton) SetBehaviors(behaviors js.Value) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(behaviors)
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// Content returns the Content property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#content
func (h *HolographicButton) Content(content *Control) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(content.JSObject())
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// SetContent sets the Content property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#content
func (h *HolographicButton) SetContent(content *Control) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(content.JSObject())
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// ContentResolution returns the ContentResolution property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#contentresolution
func (h *HolographicButton) ContentResolution(contentResolution *int) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(contentResolution.JSObject())
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// SetContentResolution sets the ContentResolution property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#contentresolution
func (h *HolographicButton) SetContentResolution(contentResolution *int) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(contentResolution.JSObject())
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// ContentScaleRatio returns the ContentScaleRatio property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#contentscaleratio
func (h *HolographicButton) ContentScaleRatio(contentScaleRatio float64) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(contentScaleRatio)
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// SetContentScaleRatio sets the ContentScaleRatio property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#contentscaleratio
func (h *HolographicButton) SetContentScaleRatio(contentScaleRatio float64) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(contentScaleRatio)
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// FrontMaterial returns the FrontMaterial property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#frontmaterial
func (h *HolographicButton) FrontMaterial(frontMaterial *FluentMaterial) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(frontMaterial.JSObject())
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// SetFrontMaterial sets the FrontMaterial property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#frontmaterial
func (h *HolographicButton) SetFrontMaterial(frontMaterial *FluentMaterial) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(frontMaterial.JSObject())
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// ImageUrl returns the ImageUrl property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#imageurl
func (h *HolographicButton) ImageUrl(imageUrl string) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(imageUrl)
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// SetImageUrl sets the ImageUrl property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#imageurl
func (h *HolographicButton) SetImageUrl(imageUrl string) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(imageUrl)
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// IsVisible returns the IsVisible property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#isvisible
func (h *HolographicButton) IsVisible(isVisible bool) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(isVisible)
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// SetIsVisible sets the IsVisible property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#isvisible
func (h *HolographicButton) SetIsVisible(isVisible bool) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(isVisible)
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// Mesh returns the Mesh property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#mesh
func (h *HolographicButton) Mesh(mesh *AbstractMesh) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(mesh.JSObject())
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// SetMesh sets the Mesh property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#mesh
func (h *HolographicButton) SetMesh(mesh *AbstractMesh) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(mesh.JSObject())
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#name
func (h *HolographicButton) Name(name string) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(name)
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#name
func (h *HolographicButton) SetName(name string) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(name)
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// Node returns the Node property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#node
func (h *HolographicButton) Node(node *TransformNode) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(node.JSObject())
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// SetNode sets the Node property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#node
func (h *HolographicButton) SetNode(node *TransformNode) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(node.JSObject())
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// OnPointerClickObservable returns the OnPointerClickObservable property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#onpointerclickobservable
func (h *HolographicButton) OnPointerClickObservable(onPointerClickObservable *Observable) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(onPointerClickObservable.JSObject())
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// SetOnPointerClickObservable sets the OnPointerClickObservable property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#onpointerclickobservable
func (h *HolographicButton) SetOnPointerClickObservable(onPointerClickObservable *Observable) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(onPointerClickObservable.JSObject())
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// OnPointerDownObservable returns the OnPointerDownObservable property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#onpointerdownobservable
func (h *HolographicButton) OnPointerDownObservable(onPointerDownObservable *Observable) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(onPointerDownObservable.JSObject())
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// SetOnPointerDownObservable sets the OnPointerDownObservable property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#onpointerdownobservable
func (h *HolographicButton) SetOnPointerDownObservable(onPointerDownObservable *Observable) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(onPointerDownObservable.JSObject())
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// OnPointerEnterObservable returns the OnPointerEnterObservable property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#onpointerenterobservable
func (h *HolographicButton) OnPointerEnterObservable(onPointerEnterObservable *Observable) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(onPointerEnterObservable.JSObject())
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// SetOnPointerEnterObservable sets the OnPointerEnterObservable property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#onpointerenterobservable
func (h *HolographicButton) SetOnPointerEnterObservable(onPointerEnterObservable *Observable) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(onPointerEnterObservable.JSObject())
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// OnPointerMoveObservable returns the OnPointerMoveObservable property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#onpointermoveobservable
func (h *HolographicButton) OnPointerMoveObservable(onPointerMoveObservable *Observable) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(onPointerMoveObservable.JSObject())
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// SetOnPointerMoveObservable sets the OnPointerMoveObservable property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#onpointermoveobservable
func (h *HolographicButton) SetOnPointerMoveObservable(onPointerMoveObservable *Observable) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(onPointerMoveObservable.JSObject())
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// OnPointerOutObservable returns the OnPointerOutObservable property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#onpointeroutobservable
func (h *HolographicButton) OnPointerOutObservable(onPointerOutObservable *Observable) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(onPointerOutObservable.JSObject())
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// SetOnPointerOutObservable sets the OnPointerOutObservable property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#onpointeroutobservable
func (h *HolographicButton) SetOnPointerOutObservable(onPointerOutObservable *Observable) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(onPointerOutObservable.JSObject())
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// OnPointerUpObservable returns the OnPointerUpObservable property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#onpointerupobservable
func (h *HolographicButton) OnPointerUpObservable(onPointerUpObservable *Observable) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(onPointerUpObservable.JSObject())
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// SetOnPointerUpObservable sets the OnPointerUpObservable property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#onpointerupobservable
func (h *HolographicButton) SetOnPointerUpObservable(onPointerUpObservable *Observable) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(onPointerUpObservable.JSObject())
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// Parent returns the Parent property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#parent
func (h *HolographicButton) Parent(parent *Container3D) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(parent.JSObject())
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// SetParent sets the Parent property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#parent
func (h *HolographicButton) SetParent(parent *Container3D) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(parent.JSObject())
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// PlateMaterial returns the PlateMaterial property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#platematerial
func (h *HolographicButton) PlateMaterial(plateMaterial *StandardMaterial) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(plateMaterial.JSObject())
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// SetPlateMaterial sets the PlateMaterial property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#platematerial
func (h *HolographicButton) SetPlateMaterial(plateMaterial *StandardMaterial) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(plateMaterial.JSObject())
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// PointerDownAnimation returns the PointerDownAnimation property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#pointerdownanimation
func (h *HolographicButton) PointerDownAnimation(pointerDownAnimation func()) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(pointerDownAnimation)
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// SetPointerDownAnimation sets the PointerDownAnimation property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#pointerdownanimation
func (h *HolographicButton) SetPointerDownAnimation(pointerDownAnimation func()) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(pointerDownAnimation)
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// PointerEnterAnimation returns the PointerEnterAnimation property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#pointerenteranimation
func (h *HolographicButton) PointerEnterAnimation(pointerEnterAnimation func()) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(pointerEnterAnimation)
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// SetPointerEnterAnimation sets the PointerEnterAnimation property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#pointerenteranimation
func (h *HolographicButton) SetPointerEnterAnimation(pointerEnterAnimation func()) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(pointerEnterAnimation)
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// PointerOutAnimation returns the PointerOutAnimation property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#pointeroutanimation
func (h *HolographicButton) PointerOutAnimation(pointerOutAnimation func()) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(pointerOutAnimation)
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// SetPointerOutAnimation sets the PointerOutAnimation property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#pointeroutanimation
func (h *HolographicButton) SetPointerOutAnimation(pointerOutAnimation func()) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(pointerOutAnimation)
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// PointerUpAnimation returns the PointerUpAnimation property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#pointerupanimation
func (h *HolographicButton) PointerUpAnimation(pointerUpAnimation func()) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(pointerUpAnimation)
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// SetPointerUpAnimation sets the PointerUpAnimation property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#pointerupanimation
func (h *HolographicButton) SetPointerUpAnimation(pointerUpAnimation func()) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(pointerUpAnimation)
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// Position returns the Position property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#position
func (h *HolographicButton) Position(position *Vector3) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(position.JSObject())
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// SetPosition sets the Position property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#position
func (h *HolographicButton) SetPosition(position *Vector3) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(position.JSObject())
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// RenderingGroupId returns the RenderingGroupId property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#renderinggroupid
func (h *HolographicButton) RenderingGroupId(renderingGroupId float64) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(renderingGroupId)
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// SetRenderingGroupId sets the RenderingGroupId property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#renderinggroupid
func (h *HolographicButton) SetRenderingGroupId(renderingGroupId float64) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(renderingGroupId)
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// Scaling returns the Scaling property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#scaling
func (h *HolographicButton) Scaling(scaling *Vector3) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(scaling.JSObject())
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// SetScaling sets the Scaling property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#scaling
func (h *HolographicButton) SetScaling(scaling *Vector3) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(scaling.JSObject())
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// ShareMaterials returns the ShareMaterials property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#sharematerials
func (h *HolographicButton) ShareMaterials(shareMaterials bool) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(shareMaterials)
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// SetShareMaterials sets the ShareMaterials property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#sharematerials
func (h *HolographicButton) SetShareMaterials(shareMaterials bool) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(shareMaterials)
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// Text returns the Text property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#text
func (h *HolographicButton) Text(text string) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(text)
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// SetText sets the Text property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#text
func (h *HolographicButton) SetText(text string) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(text)
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// TooltipText returns the TooltipText property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#tooltiptext
func (h *HolographicButton) TooltipText(tooltipText string) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(tooltipText)
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// SetTooltipText sets the TooltipText property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#tooltiptext
func (h *HolographicButton) SetTooltipText(tooltipText string) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(tooltipText)
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// TypeName returns the TypeName property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#typename
func (h *HolographicButton) TypeName(typeName string) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(typeName)
	return HolographicButtonFromJSObject(p, ba.ctx)
}

// SetTypeName sets the TypeName property of class HolographicButton.
//
// https://doc.babylonjs.com/api/classes/babylon.holographicbutton#typename
func (h *HolographicButton) SetTypeName(typeName string) *HolographicButton {
	p := ba.ctx.Get("HolographicButton").New(typeName)
	return HolographicButtonFromJSObject(p, ba.ctx)
}

*/
