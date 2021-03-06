// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ISceneSerializableComponent represents a babylon.js ISceneSerializableComponent.
// This represents a SERIALIZABLE scene component.
//
// This extends Scene Component to add Serialization methods on top.
type ISceneSerializableComponent struct {
	*ISceneComponent
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *ISceneSerializableComponent) JSObject() js.Value { return i.p }

// ISceneSerializableComponent returns a ISceneSerializableComponent JavaScript class.
func (ba *Babylon) ISceneSerializableComponent() *ISceneSerializableComponent {
	p := ba.ctx.Get("ISceneSerializableComponent")
	return ISceneSerializableComponentFromJSObject(p, ba.ctx)
}

// ISceneSerializableComponentFromJSObject returns a wrapped ISceneSerializableComponent JavaScript class.
func ISceneSerializableComponentFromJSObject(p js.Value, ctx js.Value) *ISceneSerializableComponent {
	return &ISceneSerializableComponent{ISceneComponent: ISceneComponentFromJSObject(p, ctx), ctx: ctx}
}

// ISceneSerializableComponentArrayToJSArray returns a JavaScript Array for the wrapped array.
func ISceneSerializableComponentArrayToJSArray(array []*ISceneSerializableComponent) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// AddFromContainer calls the AddFromContainer method on the ISceneSerializableComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneserializablecomponent#addfromcontainer
func (i *ISceneSerializableComponent) AddFromContainer(container *AbstractScene) {

	args := make([]interface{}, 0, 1+0)

	if container == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, container.JSObject())
	}

	i.p.Call("addFromContainer", args...)
}

// ISceneSerializableComponentRemoveFromContainerOpts contains optional parameters for ISceneSerializableComponent.RemoveFromContainer.
type ISceneSerializableComponentRemoveFromContainerOpts struct {
	Dispose *bool
}

// RemoveFromContainer calls the RemoveFromContainer method on the ISceneSerializableComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneserializablecomponent#removefromcontainer
func (i *ISceneSerializableComponent) RemoveFromContainer(container *AbstractScene, opts *ISceneSerializableComponentRemoveFromContainerOpts) {
	if opts == nil {
		opts = &ISceneSerializableComponentRemoveFromContainerOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	if container == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, container.JSObject())
	}

	if opts.Dispose == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Dispose)
	}

	i.p.Call("removeFromContainer", args...)
}

// Serialize calls the Serialize method on the ISceneSerializableComponent object.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneserializablecomponent#serialize
func (i *ISceneSerializableComponent) Serialize(serializationObject JSObject) {

	args := make([]interface{}, 0, 1+0)

	if serializationObject == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, serializationObject.JSObject())
	}

	i.p.Call("serialize", args...)
}
