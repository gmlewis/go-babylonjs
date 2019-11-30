// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GUI3DManager represents a babylon.js GUI3DManager.
// Class used to manage 3D user interface
//
// See: http://doc.babylonjs.com/how_to/gui3d
type GUI3DManager struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (g *GUI3DManager) JSObject() js.Value { return g.p }

// GUI3DManager returns a GUI3DManager JavaScript class.
func (b *Babylon) GUI3DManager() *GUI3DManager {
	p := b.ctx.Get("GUI3DManager")
	return GUI3DManagerFromJSObject(p)
}

// GUI3DManagerFromJSObject returns a wrapped GUI3DManager JavaScript class.
func GUI3DManagerFromJSObject(p js.Value) *GUI3DManager {
	return &GUI3DManager{p: p}
}

// NewGUI3DManagerOpts contains optional parameters for NewGUI3DManager.
type NewGUI3DManagerOpts struct {
	Scene *Scene
}

// NewGUI3DManager returns a new GUI3DManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui3dmanager
func (b *Babylon) NewGUI3DManager(opts *NewGUI3DManagerOpts) *GUI3DManager {
	if opts == nil {
		opts = &NewGUI3DManagerOpts{}
	}

	p := b.ctx.Get("GUI3DManager").New(opts.Scene.JSObject())
	return GUI3DManagerFromJSObject(p)
}

// TODO: methods
