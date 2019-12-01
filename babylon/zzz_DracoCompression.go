// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DracoCompression represents a babylon.js DracoCompression.
// Draco compression (&lt;a href=&#34;https://google.github.io/draco/&#34;&gt;https://google.github.io/draco/&lt;/a&gt;)
//
// This class wraps the Draco module.
//
// See: https://www.babylonjs-playground.com/#N3EK4B#0
type DracoCompression struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (d *DracoCompression) JSObject() js.Value { return d.p }

// DracoCompression returns a DracoCompression JavaScript class.
func (ba *Babylon) DracoCompression() *DracoCompression {
	p := ba.ctx.Get("DracoCompression")
	return DracoCompressionFromJSObject(p, ba.ctx)
}

// DracoCompressionFromJSObject returns a wrapped DracoCompression JavaScript class.
func DracoCompressionFromJSObject(p js.Value, ctx js.Value) *DracoCompression {
	return &DracoCompression{p: p, ctx: ctx}
}

// DracoCompressionArrayToJSArray returns a JavaScript Array for the wrapped array.
func DracoCompressionArrayToJSArray(array []*DracoCompression) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewDracoCompressionOpts contains optional parameters for NewDracoCompression.
type NewDracoCompressionOpts struct {
	NumWorkers *float64
}

// NewDracoCompression returns a new DracoCompression object.
//
// https://doc.babylonjs.com/api/classes/babylon.dracocompression
func (ba *Babylon) NewDracoCompression(opts *NewDracoCompressionOpts) *DracoCompression {
	if opts == nil {
		opts = &NewDracoCompressionOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.NumWorkers == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.NumWorkers)
	}

	p := ba.ctx.Get("DracoCompression").New(args...)
	return DracoCompressionFromJSObject(p, ba.ctx)
}

// DracoCompressionDecodeMeshAsyncOpts contains optional parameters for DracoCompression.DecodeMeshAsync.
type DracoCompressionDecodeMeshAsyncOpts struct {
	Attributes js.Value
}

// DecodeMeshAsync calls the DecodeMeshAsync method on the DracoCompression object.
//
// https://doc.babylonjs.com/api/classes/babylon.dracocompression#decodemeshasync
func (d *DracoCompression) DecodeMeshAsync(data js.Value, opts *DracoCompressionDecodeMeshAsyncOpts) *Promise {
	if opts == nil {
		opts = &DracoCompressionDecodeMeshAsyncOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, data)

	if opts.Attributes == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Attributes)
	}

	retVal := d.p.Call("decodeMeshAsync", args...)
	return PromiseFromJSObject(retVal, d.ctx)
}

// Dispose calls the Dispose method on the DracoCompression object.
//
// https://doc.babylonjs.com/api/classes/babylon.dracocompression#dispose
func (d *DracoCompression) Dispose() {

	d.p.Call("dispose")
}

// WhenReadyAsync calls the WhenReadyAsync method on the DracoCompression object.
//
// https://doc.babylonjs.com/api/classes/babylon.dracocompression#whenreadyasync
func (d *DracoCompression) WhenReadyAsync() *Promise {

	retVal := d.p.Call("whenReadyAsync")
	return PromiseFromJSObject(retVal, d.ctx)
}

/*

// Configuration returns the Configuration property of class DracoCompression.
//
// https://doc.babylonjs.com/api/classes/babylon.dracocompression#configuration
func (d *DracoCompression) Configuration(Configuration *IDracoCompressionConfiguration) *DracoCompression {
	p := ba.ctx.Get("DracoCompression").New(Configuration.JSObject())
	return DracoCompressionFromJSObject(p, ba.ctx)
}

// SetConfiguration sets the Configuration property of class DracoCompression.
//
// https://doc.babylonjs.com/api/classes/babylon.dracocompression#configuration
func (d *DracoCompression) SetConfiguration(Configuration *IDracoCompressionConfiguration) *DracoCompression {
	p := ba.ctx.Get("DracoCompression").New(Configuration.JSObject())
	return DracoCompressionFromJSObject(p, ba.ctx)
}

// DecoderAvailable returns the DecoderAvailable property of class DracoCompression.
//
// https://doc.babylonjs.com/api/classes/babylon.dracocompression#decoderavailable
func (d *DracoCompression) DecoderAvailable(DecoderAvailable bool) *DracoCompression {
	p := ba.ctx.Get("DracoCompression").New(DecoderAvailable)
	return DracoCompressionFromJSObject(p, ba.ctx)
}

// SetDecoderAvailable sets the DecoderAvailable property of class DracoCompression.
//
// https://doc.babylonjs.com/api/classes/babylon.dracocompression#decoderavailable
func (d *DracoCompression) SetDecoderAvailable(DecoderAvailable bool) *DracoCompression {
	p := ba.ctx.Get("DracoCompression").New(DecoderAvailable)
	return DracoCompressionFromJSObject(p, ba.ctx)
}

// Default returns the Default property of class DracoCompression.
//
// https://doc.babylonjs.com/api/classes/babylon.dracocompression#default
func (d *DracoCompression) Default(Default *DracoCompression) *DracoCompression {
	p := ba.ctx.Get("DracoCompression").New(Default.JSObject())
	return DracoCompressionFromJSObject(p, ba.ctx)
}

// SetDefault sets the Default property of class DracoCompression.
//
// https://doc.babylonjs.com/api/classes/babylon.dracocompression#default
func (d *DracoCompression) SetDefault(Default *DracoCompression) *DracoCompression {
	p := ba.ctx.Get("DracoCompression").New(Default.JSObject())
	return DracoCompressionFromJSObject(p, ba.ctx)
}

// DefaultNumWorkers returns the DefaultNumWorkers property of class DracoCompression.
//
// https://doc.babylonjs.com/api/classes/babylon.dracocompression#defaultnumworkers
func (d *DracoCompression) DefaultNumWorkers(DefaultNumWorkers float64) *DracoCompression {
	p := ba.ctx.Get("DracoCompression").New(DefaultNumWorkers)
	return DracoCompressionFromJSObject(p, ba.ctx)
}

// SetDefaultNumWorkers sets the DefaultNumWorkers property of class DracoCompression.
//
// https://doc.babylonjs.com/api/classes/babylon.dracocompression#defaultnumworkers
func (d *DracoCompression) SetDefaultNumWorkers(DefaultNumWorkers float64) *DracoCompression {
	p := ba.ctx.Get("DracoCompression").New(DefaultNumWorkers)
	return DracoCompressionFromJSObject(p, ba.ctx)
}

*/
