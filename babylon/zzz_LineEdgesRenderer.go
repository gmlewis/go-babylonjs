// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// LineEdgesRenderer represents a babylon.js LineEdgesRenderer.
// LineEdgesRenderer for LineMeshes to remove unnecessary triangulation
type LineEdgesRenderer struct{ *EdgesRenderer }

// JSObject returns the underlying js.Value.
func (l *LineEdgesRenderer) JSObject() js.Value { return l.p }

// LineEdgesRenderer returns a LineEdgesRenderer JavaScript class.
func (b *Babylon) LineEdgesRenderer() *LineEdgesRenderer {
	p := b.ctx.Get("LineEdgesRenderer")
	return LineEdgesRendererFromJSObject(p)
}

// LineEdgesRendererFromJSObject returns a wrapped LineEdgesRenderer JavaScript class.
func LineEdgesRendererFromJSObject(p js.Value) *LineEdgesRenderer {
	return &LineEdgesRenderer{EdgesRendererFromJSObject(p)}
}

// NewLineEdgesRendererOpts contains optional parameters for NewLineEdgesRenderer.
type NewLineEdgesRendererOpts struct {
	Epsilon *float64

	CheckVerticesInsteadOfIndices *bool
}

// NewLineEdgesRenderer returns a new LineEdgesRenderer object.
//
// https://doc.babylonjs.com/api/classes/babylon.lineedgesrenderer
func (b *Babylon) NewLineEdgesRenderer(source *AbstractMesh, opts *NewLineEdgesRendererOpts) *LineEdgesRenderer {
	if opts == nil {
		opts = &NewLineEdgesRendererOpts{}
	}

	p := b.ctx.Get("LineEdgesRenderer").New(source.JSObject(), opts.Epsilon, opts.CheckVerticesInsteadOfIndices.JSObject())
	return LineEdgesRendererFromJSObject(p)
}

// TODO: methods
