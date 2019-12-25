// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SkeletonViewer represents a babylon.js SkeletonViewer.
// Class used to render a debug view of a given skeleton
//
// See: http://www.babylonjs-playground.com/#1BZJVJ#8
type SkeletonViewer struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *SkeletonViewer) JSObject() js.Value { return s.p }

// SkeletonViewer returns a SkeletonViewer JavaScript class.
func (ba *Babylon) SkeletonViewer() *SkeletonViewer {
	p := ba.ctx.Get("SkeletonViewer")
	return SkeletonViewerFromJSObject(p, ba.ctx)
}

// SkeletonViewerFromJSObject returns a wrapped SkeletonViewer JavaScript class.
func SkeletonViewerFromJSObject(p js.Value, ctx js.Value) *SkeletonViewer {
	return &SkeletonViewer{p: p, ctx: ctx}
}

// SkeletonViewerArrayToJSArray returns a JavaScript Array for the wrapped array.
func SkeletonViewerArrayToJSArray(array []*SkeletonViewer) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewSkeletonViewerOpts contains optional parameters for NewSkeletonViewer.
type NewSkeletonViewerOpts struct {
	AutoUpdateBonesMatrices *bool
	RenderingGroupId        *float64
}

// NewSkeletonViewer returns a new SkeletonViewer object.
//
// https://doc.babylonjs.com/api/classes/babylon.skeletonviewer#constructor
func (ba *Babylon) NewSkeletonViewer(skeleton *Skeleton, mesh *AbstractMesh, scene *Scene, opts *NewSkeletonViewerOpts) *SkeletonViewer {
	if opts == nil {
		opts = &NewSkeletonViewerOpts{}
	}

	args := make([]interface{}, 0, 3+2)

	args = append(args, skeleton.JSObject())
	args = append(args, mesh.JSObject())
	args = append(args, scene.JSObject())

	if opts.AutoUpdateBonesMatrices == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.AutoUpdateBonesMatrices)
	}
	if opts.RenderingGroupId == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.RenderingGroupId)
	}

	p := ba.ctx.Get("SkeletonViewer").New(args...)
	return SkeletonViewerFromJSObject(p, ba.ctx)
}

// Dispose calls the Dispose method on the SkeletonViewer object.
//
// https://doc.babylonjs.com/api/classes/babylon.skeletonviewer#dispose
func (s *SkeletonViewer) Dispose() {

	s.p.Call("dispose")
}

// Update calls the Update method on the SkeletonViewer object.
//
// https://doc.babylonjs.com/api/classes/babylon.skeletonviewer#update
func (s *SkeletonViewer) Update() {

	s.p.Call("update")
}

// AutoUpdateBonesMatrices returns the AutoUpdateBonesMatrices property of class SkeletonViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.skeletonviewer#autoupdatebonesmatrices
func (s *SkeletonViewer) AutoUpdateBonesMatrices() bool {
	retVal := s.p.Get("autoUpdateBonesMatrices")
	return retVal.Bool()
}

// SetAutoUpdateBonesMatrices sets the AutoUpdateBonesMatrices property of class SkeletonViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.skeletonviewer#autoupdatebonesmatrices
func (s *SkeletonViewer) SetAutoUpdateBonesMatrices(autoUpdateBonesMatrices bool) *SkeletonViewer {
	s.p.Set("autoUpdateBonesMatrices", autoUpdateBonesMatrices)
	return s
}

// Color returns the Color property of class SkeletonViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.skeletonviewer#color
func (s *SkeletonViewer) Color() *Color3 {
	retVal := s.p.Get("color")
	return Color3FromJSObject(retVal, s.ctx)
}

// SetColor sets the Color property of class SkeletonViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.skeletonviewer#color
func (s *SkeletonViewer) SetColor(color *Color3) *SkeletonViewer {
	s.p.Set("color", color.JSObject())
	return s
}

// DebugMesh returns the DebugMesh property of class SkeletonViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.skeletonviewer#debugmesh
func (s *SkeletonViewer) DebugMesh() *LinesMesh {
	retVal := s.p.Get("debugMesh")
	return LinesMeshFromJSObject(retVal, s.ctx)
}

// SetDebugMesh sets the DebugMesh property of class SkeletonViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.skeletonviewer#debugmesh
func (s *SkeletonViewer) SetDebugMesh(debugMesh *LinesMesh) *SkeletonViewer {
	s.p.Set("debugMesh", debugMesh.JSObject())
	return s
}

// IsEnabled returns the IsEnabled property of class SkeletonViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.skeletonviewer#isenabled
func (s *SkeletonViewer) IsEnabled() bool {
	retVal := s.p.Get("isEnabled")
	return retVal.Bool()
}

// SetIsEnabled sets the IsEnabled property of class SkeletonViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.skeletonviewer#isenabled
func (s *SkeletonViewer) SetIsEnabled(isEnabled bool) *SkeletonViewer {
	s.p.Set("isEnabled", isEnabled)
	return s
}

// Mesh returns the Mesh property of class SkeletonViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.skeletonviewer#mesh
func (s *SkeletonViewer) Mesh() *AbstractMesh {
	retVal := s.p.Get("mesh")
	return AbstractMeshFromJSObject(retVal, s.ctx)
}

// SetMesh sets the Mesh property of class SkeletonViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.skeletonviewer#mesh
func (s *SkeletonViewer) SetMesh(mesh *AbstractMesh) *SkeletonViewer {
	s.p.Set("mesh", mesh.JSObject())
	return s
}

// RenderingGroupId returns the RenderingGroupId property of class SkeletonViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.skeletonviewer#renderinggroupid
func (s *SkeletonViewer) RenderingGroupId() float64 {
	retVal := s.p.Get("renderingGroupId")
	return retVal.Float()
}

// SetRenderingGroupId sets the RenderingGroupId property of class SkeletonViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.skeletonviewer#renderinggroupid
func (s *SkeletonViewer) SetRenderingGroupId(renderingGroupId float64) *SkeletonViewer {
	s.p.Set("renderingGroupId", renderingGroupId)
	return s
}

// Skeleton returns the Skeleton property of class SkeletonViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.skeletonviewer#skeleton
func (s *SkeletonViewer) Skeleton() *Skeleton {
	retVal := s.p.Get("skeleton")
	return SkeletonFromJSObject(retVal, s.ctx)
}

// SetSkeleton sets the Skeleton property of class SkeletonViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.skeletonviewer#skeleton
func (s *SkeletonViewer) SetSkeleton(skeleton *Skeleton) *SkeletonViewer {
	s.p.Set("skeleton", skeleton.JSObject())
	return s
}
