// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Bone represents a babylon.js Bone.
// Class used to store bone information
//
// See: http://doc.babylonjs.com/how_to/how_to_use_bones_and_skeletons
type Bone struct{ *Node }

// JSObject returns the underlying js.Value.
func (b *Bone) JSObject() js.Value { return b.p }

// Bone returns a Bone JavaScript class.
func (b *Babylon) Bone() *Bone {
	p := b.ctx.Get("Bone")
	return BoneFromJSObject(p)
}

// BoneFromJSObject returns a wrapped Bone JavaScript class.
func BoneFromJSObject(p js.Value) *Bone {
	return &Bone{NodeFromJSObject(p)}
}

// NewBoneOpts contains optional parameters for NewBone.
type NewBoneOpts struct {
	ParentBone *Bone

	LocalMatrix *Matrix

	RestPose *Matrix

	BaseMatrix *Matrix

	Index *float64
}

// NewBone returns a new Bone object.
//
// https://doc.babylonjs.com/api/classes/babylon.bone
func (b *Babylon) NewBone(name string, skeleton *Skeleton, opts *NewBoneOpts) *Bone {
	if opts == nil {
		opts = &NewBoneOpts{}
	}

	p := b.ctx.Get("Bone").New(name, skeleton.JSObject(), opts.ParentBone.JSObject(), opts.LocalMatrix.JSObject(), opts.RestPose.JSObject(), opts.BaseMatrix.JSObject(), opts.Index)
	return BoneFromJSObject(p)
}

// TODO: methods
