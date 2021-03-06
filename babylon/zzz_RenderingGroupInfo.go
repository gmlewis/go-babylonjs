// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// RenderingGroupInfo represents a babylon.js RenderingGroupInfo.
// This class is used by the onRenderingGroupObservable
type RenderingGroupInfo struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (r *RenderingGroupInfo) JSObject() js.Value { return r.p }

// RenderingGroupInfo returns a RenderingGroupInfo JavaScript class.
func (ba *Babylon) RenderingGroupInfo() *RenderingGroupInfo {
	p := ba.ctx.Get("RenderingGroupInfo")
	return RenderingGroupInfoFromJSObject(p, ba.ctx)
}

// RenderingGroupInfoFromJSObject returns a wrapped RenderingGroupInfo JavaScript class.
func RenderingGroupInfoFromJSObject(p js.Value, ctx js.Value) *RenderingGroupInfo {
	return &RenderingGroupInfo{p: p, ctx: ctx}
}

// RenderingGroupInfoArrayToJSArray returns a JavaScript Array for the wrapped array.
func RenderingGroupInfoArrayToJSArray(array []*RenderingGroupInfo) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// Camera returns the Camera property of class RenderingGroupInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.renderinggroupinfo#camera
func (r *RenderingGroupInfo) Camera() *Camera {
	retVal := r.p.Get("camera")
	return CameraFromJSObject(retVal, r.ctx)
}

// SetCamera sets the Camera property of class RenderingGroupInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.renderinggroupinfo#camera
func (r *RenderingGroupInfo) SetCamera(camera *Camera) *RenderingGroupInfo {
	r.p.Set("camera", camera.JSObject())
	return r
}

// RenderingGroupId returns the RenderingGroupId property of class RenderingGroupInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.renderinggroupinfo#renderinggroupid
func (r *RenderingGroupInfo) RenderingGroupId() float64 {
	retVal := r.p.Get("renderingGroupId")
	return retVal.Float()
}

// SetRenderingGroupId sets the RenderingGroupId property of class RenderingGroupInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.renderinggroupinfo#renderinggroupid
func (r *RenderingGroupInfo) SetRenderingGroupId(renderingGroupId float64) *RenderingGroupInfo {
	r.p.Set("renderingGroupId", renderingGroupId)
	return r
}

// Scene returns the Scene property of class RenderingGroupInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.renderinggroupinfo#scene
func (r *RenderingGroupInfo) Scene() *Scene {
	retVal := r.p.Get("scene")
	return SceneFromJSObject(retVal, r.ctx)
}

// SetScene sets the Scene property of class RenderingGroupInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.renderinggroupinfo#scene
func (r *RenderingGroupInfo) SetScene(scene *Scene) *RenderingGroupInfo {
	r.p.Set("scene", scene.JSObject())
	return r
}
