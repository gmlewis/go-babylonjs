// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// MorphTargetManager represents a babylon.js MorphTargetManager.
// This class is used to deform meshes using morphing between different targets
//
// See: http://doc.babylonjs.com/how_to/how_to_use_morphtargets
type MorphTargetManager struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (m *MorphTargetManager) JSObject() js.Value { return m.p }

// MorphTargetManager returns a MorphTargetManager JavaScript class.
func (b *Babylon) MorphTargetManager() *MorphTargetManager {
	p := b.ctx.Get("MorphTargetManager")
	return MorphTargetManagerFromJSObject(p)
}

// MorphTargetManagerFromJSObject returns a wrapped MorphTargetManager JavaScript class.
func MorphTargetManagerFromJSObject(p js.Value) *MorphTargetManager {
	return &MorphTargetManager{p: p}
}

// NewMorphTargetManagerOpts contains optional parameters for NewMorphTargetManager.
type NewMorphTargetManagerOpts struct {
	Scene *Scene
}

// NewMorphTargetManager returns a new MorphTargetManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.morphtargetmanager
func (b *Babylon) NewMorphTargetManager(opts *NewMorphTargetManagerOpts) *MorphTargetManager {
	if opts == nil {
		opts = &NewMorphTargetManagerOpts{}
	}

	p := b.ctx.Get("MorphTargetManager").New(opts.Scene.JSObject())
	return MorphTargetManagerFromJSObject(p)
}

// TODO: methods
