// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Grid represents a babylon.js Grid.
// Class used to create a 2D grid container
type Grid struct {
	*Container
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (g *Grid) JSObject() js.Value { return g.p }

// Grid returns a Grid JavaScript class.
func (ba *Babylon) Grid() *Grid {
	p := ba.ctx.Get("Grid")
	return GridFromJSObject(p, ba.ctx)
}

// GridFromJSObject returns a wrapped Grid JavaScript class.
func GridFromJSObject(p js.Value, ctx js.Value) *Grid {
	return &Grid{Container: ContainerFromJSObject(p, ctx), ctx: ctx}
}

// NewGridOpts contains optional parameters for NewGrid.
type NewGridOpts struct {
	Name *JSString
}

// NewGrid returns a new Grid object.
//
// https://doc.babylonjs.com/api/classes/babylon.grid
func (ba *Babylon) NewGrid(opts *NewGridOpts) *Grid {
	if opts == nil {
		opts = &NewGridOpts{}
	}

	p := ba.ctx.Get("Grid").New(opts.Name.JSObject())
	return GridFromJSObject(p, ba.ctx)
}

// TODO: methods
