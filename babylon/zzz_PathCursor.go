// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PathCursor represents a babylon.js PathCursor.
// A cursor which tracks a point on a path
type PathCursor struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (p *PathCursor) JSObject() js.Value { return p.p }

// PathCursor returns a PathCursor JavaScript class.
func (b *Babylon) PathCursor() *PathCursor {
	p := b.ctx.Get("PathCursor")
	return PathCursorFromJSObject(p)
}

// PathCursorFromJSObject returns a wrapped PathCursor JavaScript class.
func PathCursorFromJSObject(p js.Value) *PathCursor {
	return &PathCursor{p: p}
}

// NewPathCursor returns a new PathCursor object.
//
// https://doc.babylonjs.com/api/classes/babylon.pathcursor
func (b *Babylon) NewPathCursor(path *Path2) *PathCursor {
	p := b.ctx.Get("PathCursor").New(path.JSObject())
	return PathCursorFromJSObject(p)
}

// TODO: methods
