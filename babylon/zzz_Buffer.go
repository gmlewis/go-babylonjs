// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Buffer represents a babylon.js Buffer.
// Class used to store data that will be store in GPU memory
type Buffer struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (b *Buffer) JSObject() js.Value { return b.p }

// Buffer returns a Buffer JavaScript class.
func (ba *Babylon) Buffer() *Buffer {
	p := ba.ctx.Get("Buffer")
	return BufferFromJSObject(p, ba.ctx)
}

// BufferFromJSObject returns a wrapped Buffer JavaScript class.
func BufferFromJSObject(p js.Value, ctx js.Value) *Buffer {
	return &Buffer{p: p, ctx: ctx}
}

// BufferArrayToJSArray returns a JavaScript Array for the wrapped array.
func BufferArrayToJSArray(array []*Buffer) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewBufferOpts contains optional parameters for NewBuffer.
type NewBufferOpts struct {
	Stride                   *float64
	PostponeInternalCreation *bool
	Instanced                *bool
	UseBytes                 *bool
	Divisor                  *float64
}

// NewBuffer returns a new Buffer object.
//
// https://doc.babylonjs.com/api/classes/babylon.buffer#constructor
func (ba *Babylon) NewBuffer(engine JSObject, data []float64, updatable bool, opts *NewBufferOpts) *Buffer {
	if opts == nil {
		opts = &NewBufferOpts{}
	}

	args := make([]interface{}, 0, 3+5)

	args = append(args, engine.JSObject())
	args = append(args, data)
	args = append(args, updatable)

	if opts.Stride == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Stride)
	}
	if opts.PostponeInternalCreation == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.PostponeInternalCreation)
	}
	if opts.Instanced == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Instanced)
	}
	if opts.UseBytes == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UseBytes)
	}
	if opts.Divisor == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Divisor)
	}

	p := ba.ctx.Get("Buffer").New(args...)
	return BufferFromJSObject(p, ba.ctx)
}

// BufferCreateOpts contains optional parameters for Buffer.Create.
type BufferCreateOpts struct {
	Data []float64
}

// Create calls the Create method on the Buffer object.
//
// https://doc.babylonjs.com/api/classes/babylon.buffer#create
func (b *Buffer) Create(opts *BufferCreateOpts) {
	if opts == nil {
		opts = &BufferCreateOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Data == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Data)
	}

	b.p.Call("create", args...)
}

// BufferCreateVertexBufferOpts contains optional parameters for Buffer.CreateVertexBuffer.
type BufferCreateVertexBufferOpts struct {
	Stride    *float64
	Instanced *bool
	UseBytes  *bool
	Divisor   *float64
}

// CreateVertexBuffer calls the CreateVertexBuffer method on the Buffer object.
//
// https://doc.babylonjs.com/api/classes/babylon.buffer#createvertexbuffer
func (b *Buffer) CreateVertexBuffer(kind string, offset float64, size float64, opts *BufferCreateVertexBufferOpts) *VertexBuffer {
	if opts == nil {
		opts = &BufferCreateVertexBufferOpts{}
	}

	args := make([]interface{}, 0, 3+4)

	args = append(args, kind)

	args = append(args, offset)

	args = append(args, size)

	if opts.Stride == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Stride)
	}
	if opts.Instanced == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Instanced)
	}
	if opts.UseBytes == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UseBytes)
	}
	if opts.Divisor == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Divisor)
	}

	retVal := b.p.Call("createVertexBuffer", args...)
	return VertexBufferFromJSObject(retVal, b.ctx)
}

// Dispose calls the Dispose method on the Buffer object.
//
// https://doc.babylonjs.com/api/classes/babylon.buffer#dispose
func (b *Buffer) Dispose() {

	b.p.Call("dispose")
}

// GetBuffer calls the GetBuffer method on the Buffer object.
//
// https://doc.babylonjs.com/api/classes/babylon.buffer#getbuffer
func (b *Buffer) GetBuffer() *DataBuffer {

	retVal := b.p.Call("getBuffer")
	return DataBufferFromJSObject(retVal, b.ctx)
}

// GetData calls the GetData method on the Buffer object.
//
// https://doc.babylonjs.com/api/classes/babylon.buffer#getdata
func (b *Buffer) GetData() []float64 {

	retVal := b.p.Call("getData")
	result := []float64{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, retVal.Index(ri).Float())
	}
	return result
}

// GetStrideSize calls the GetStrideSize method on the Buffer object.
//
// https://doc.babylonjs.com/api/classes/babylon.buffer#getstridesize
func (b *Buffer) GetStrideSize() float64 {

	retVal := b.p.Call("getStrideSize")
	return retVal.Float()
}

// IsUpdatable calls the IsUpdatable method on the Buffer object.
//
// https://doc.babylonjs.com/api/classes/babylon.buffer#isupdatable
func (b *Buffer) IsUpdatable() bool {

	retVal := b.p.Call("isUpdatable")
	return retVal.Bool()
}

// Update calls the Update method on the Buffer object.
//
// https://doc.babylonjs.com/api/classes/babylon.buffer#update
func (b *Buffer) Update(data []float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, data)

	b.p.Call("update", args...)
}

// BufferUpdateDirectlyOpts contains optional parameters for Buffer.UpdateDirectly.
type BufferUpdateDirectlyOpts struct {
	VertexCount *float64
	UseBytes    *bool
}

// UpdateDirectly calls the UpdateDirectly method on the Buffer object.
//
// https://doc.babylonjs.com/api/classes/babylon.buffer#updatedirectly
func (b *Buffer) UpdateDirectly(data []float64, offset float64, opts *BufferUpdateDirectlyOpts) {
	if opts == nil {
		opts = &BufferUpdateDirectlyOpts{}
	}

	args := make([]interface{}, 0, 2+2)

	args = append(args, data)

	args = append(args, offset)

	if opts.VertexCount == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.VertexCount)
	}
	if opts.UseBytes == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UseBytes)
	}

	b.p.Call("updateDirectly", args...)
}

// ByteStride returns the ByteStride property of class Buffer.
//
// https://doc.babylonjs.com/api/classes/babylon.buffer#bytestride
func (b *Buffer) ByteStride() float64 {
	retVal := b.p.Get("byteStride")
	return retVal.Float()
}

// SetByteStride sets the ByteStride property of class Buffer.
//
// https://doc.babylonjs.com/api/classes/babylon.buffer#bytestride
func (b *Buffer) SetByteStride(byteStride float64) *Buffer {
	b.p.Set("byteStride", byteStride)
	return b
}
