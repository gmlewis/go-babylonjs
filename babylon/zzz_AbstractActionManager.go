// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AbstractActionManager represents a babylon.js AbstractActionManager.
// Abstract class used to decouple action Manager from scene and meshes.
// Do not instantiate.
//
// See: http://doc.babylonjs.com/how_to/how_to_use_actions
type AbstractActionManager struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (a *AbstractActionManager) JSObject() js.Value { return a.p }

// AbstractActionManager returns a AbstractActionManager JavaScript class.
func (b *Babylon) AbstractActionManager() *AbstractActionManager {
	p := b.ctx.Get("AbstractActionManager")
	return AbstractActionManagerFromJSObject(p)
}

// AbstractActionManagerFromJSObject returns a wrapped AbstractActionManager JavaScript class.
func AbstractActionManagerFromJSObject(p js.Value) *AbstractActionManager {
	return &AbstractActionManager{p: p}
}

// TODO: methods
