// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DracoCompression represents a babylon.js DracoCompression.
// Draco compression (&lt;a href=&#34;https://google.github.io/draco/&#34;&gt;https://google.github.io/draco/&lt;/a&gt;)
//
// To decode Draco compressed data, get the default DracoCompression object and call decodeMeshAsync:
//
// See: https://www.babylonjs-playground.com/#N3EK4B#0
type DracoCompression struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (d *DracoCompression) JSObject() js.Value { return d.p }

// DracoCompression returns a DracoCompression JavaScript class.
func (b *Babylon) DracoCompression() *DracoCompression {
	p := b.ctx.Get("DracoCompression")
	return DracoCompressionFromJSObject(p)
}

// DracoCompressionFromJSObject returns a wrapped DracoCompression JavaScript class.
func DracoCompressionFromJSObject(p js.Value) *DracoCompression {
	return &DracoCompression{p: p}
}

// NewDracoCompressionOpts contains optional parameters for NewDracoCompression.
type NewDracoCompressionOpts struct {
	NumWorkers *float64
}

// NewDracoCompression returns a new DracoCompression object.
//
// https://doc.babylonjs.com/api/classes/babylon.dracocompression
func (b *Babylon) NewDracoCompression(opts *NewDracoCompressionOpts) *DracoCompression {
	if opts == nil {
		opts = &NewDracoCompressionOpts{}
	}

	p := b.ctx.Get("DracoCompression").New(opts.NumWorkers)
	return DracoCompressionFromJSObject(p)
}

// TODO: methods
