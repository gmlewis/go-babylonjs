// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GridMaterial represents a babylon.js GridMaterial.
// The grid materials allows you to wrap any shape with a grid.
// Colors are customizable.
type GridMaterial struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (g *GridMaterial) JSObject() js.Value { return g.p }

// GridMaterial returns a GridMaterial JavaScript class.
func (b *Babylon) GridMaterial() *GridMaterial {
	p := b.ctx.Get("GridMaterial")
	return GridMaterialFromJSObject(p)
}

// GridMaterialFromJSObject returns a wrapped GridMaterial JavaScript class.
func GridMaterialFromJSObject(p js.Value) *GridMaterial {
	return &GridMaterial{p: p}
}

// NewGridMaterial returns a new GridMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.gridmaterial
func (b *Babylon) NewGridMaterial(name string, scene *Scene) *GridMaterial {
	p := b.ctx.Get("GridMaterial").New(name, scene.JSObject())
	return GridMaterialFromJSObject(p)
}

// TODO: methods
