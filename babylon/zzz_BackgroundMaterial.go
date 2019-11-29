// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// BackgroundMaterial represents a babylon.js BackgroundMaterial.
// Background material used to create an efficient environement around your scene.
type BackgroundMaterial struct{}

// JSObject returns the underlying js.Value.
func (b *BackgroundMaterial) JSObject() js.Value { return b.p }

// BackgroundMaterial returns a BackgroundMaterial JavaScript class.
func (b *Babylon) BackgroundMaterial() *BackgroundMaterial {
	p := b.ctx.Get("BackgroundMaterial")
	return BackgroundMaterialFromJSObject(p)
}

// BackgroundMaterialFromJSObject returns a wrapped BackgroundMaterial JavaScript class.
func BackgroundMaterialFromJSObject(p js.Value) *BackgroundMaterial {
	return &BackgroundMaterial{p: p}
}

// NewBackgroundMaterial returns a new BackgroundMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.backgroundmaterial
func (b *Babylon) NewBackgroundMaterial(todo parameters) *BackgroundMaterial {
	p := b.ctx.Get("BackgroundMaterial").New(todo)
	return BackgroundMaterialFromJSObject(p)
}

// TODO: methods