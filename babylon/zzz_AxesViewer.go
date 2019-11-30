// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AxesViewer represents a babylon.js AxesViewer.
// The Axes viewer will show 3 axes in a specific point in space
type AxesViewer struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (a *AxesViewer) JSObject() js.Value { return a.p }

// AxesViewer returns a AxesViewer JavaScript class.
func (b *Babylon) AxesViewer() *AxesViewer {
	p := b.ctx.Get("AxesViewer")
	return AxesViewerFromJSObject(p)
}

// AxesViewerFromJSObject returns a wrapped AxesViewer JavaScript class.
func AxesViewerFromJSObject(p js.Value) *AxesViewer {
	return &AxesViewer{p: p}
}

// NewAxesViewerOpts contains optional parameters for NewAxesViewer.
type NewAxesViewerOpts struct {
	ScaleLines *float64

	RenderingGroupId *float64

	XAxis *TransformNode

	YAxis *TransformNode

	ZAxis *TransformNode
}

// NewAxesViewer returns a new AxesViewer object.
//
// https://doc.babylonjs.com/api/classes/babylon.axesviewer
func (b *Babylon) NewAxesViewer(scene *Scene, opts *NewAxesViewerOpts) *AxesViewer {
	if opts == nil {
		opts = &NewAxesViewerOpts{}
	}

	p := b.ctx.Get("AxesViewer").New(scene.JSObject(), opts.ScaleLines, opts.RenderingGroupId, opts.XAxis.JSObject(), opts.YAxis.JSObject(), opts.ZAxis.JSObject())
	return AxesViewerFromJSObject(p)
}

// TODO: methods
