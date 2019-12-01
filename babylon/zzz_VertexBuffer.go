// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// VertexBuffer represents a babylon.js VertexBuffer.
// Specialized buffer used to store vertex data
type VertexBuffer struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (v *VertexBuffer) JSObject() js.Value { return v.p }

// VertexBuffer returns a VertexBuffer JavaScript class.
func (ba *Babylon) VertexBuffer() *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer")
	return VertexBufferFromJSObject(p, ba.ctx)
}

// VertexBufferFromJSObject returns a wrapped VertexBuffer JavaScript class.
func VertexBufferFromJSObject(p js.Value, ctx js.Value) *VertexBuffer {
	return &VertexBuffer{p: p, ctx: ctx}
}

// VertexBufferArrayToJSArray returns a JavaScript Array for the wrapped array.
func VertexBufferArrayToJSArray(array []*VertexBuffer) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewVertexBufferOpts contains optional parameters for NewVertexBuffer.
type NewVertexBufferOpts struct {
	PostponeInternalCreation *bool
	Stride                   *float64
	Instanced                *bool
	Offset                   *float64
	Size                     *float64
	Type                     *float64
	Normalized               *bool
	UseBytes                 *bool
	Divisor                  *float64
}

// NewVertexBuffer returns a new VertexBuffer object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer
func (ba *Babylon) NewVertexBuffer(engine interface{}, data []*float64, kind string, updatable bool, opts *NewVertexBufferOpts) *VertexBuffer {
	if opts == nil {
		opts = &NewVertexBufferOpts{}
	}

	args := make([]interface{}, 0, 4+9)

	args = append(args, engine)
	args = append(args, data)
	args = append(args, kind)
	args = append(args, updatable)

	if opts.PostponeInternalCreation == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.PostponeInternalCreation)
	}
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
	if opts.Offset == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Offset)
	}
	if opts.Size == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Size)
	}
	if opts.Type == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Type)
	}
	if opts.Normalized == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Normalized)
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

	p := ba.ctx.Get("VertexBuffer").New(args...)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// VertexBufferCreateOpts contains optional parameters for VertexBuffer.Create.
type VertexBufferCreateOpts struct {
	Data []float64
}

// Create calls the Create method on the VertexBuffer object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#create
func (v *VertexBuffer) Create(opts *VertexBufferCreateOpts) {
	if opts == nil {
		opts = &VertexBufferCreateOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Data == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Data.JSObject())
	}

	v.p.Call("create", args...)
}

// DeduceStride calls the DeduceStride method on the VertexBuffer object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#deducestride
func (v *VertexBuffer) DeduceStride(kind string) float64 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, kind)

	retVal := v.p.Call("DeduceStride", args...)
	return retVal.Float()
}

// Dispose calls the Dispose method on the VertexBuffer object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#dispose
func (v *VertexBuffer) Dispose() {

	v.p.Call("dispose")
}

// ForEach calls the ForEach method on the VertexBuffer object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#foreach
func (v *VertexBuffer) ForEach(data []*float64, byteOffset float64, byteStride float64, componentCount float64, componentType float64, count float64, normalized bool, callback func()) {

	args := make([]interface{}, 0, 8+0)

	args = append(args, float64ArrayToJSArray(data))
	args = append(args, byteOffset)
	args = append(args, byteStride)
	args = append(args, componentCount)
	args = append(args, componentType)
	args = append(args, count)
	args = append(args, normalized)
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { callback(); return nil }))

	v.p.Call("ForEach", args...)
}

// GetBuffer calls the GetBuffer method on the VertexBuffer object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#getbuffer
func (v *VertexBuffer) GetBuffer() *DataBuffer {

	retVal := v.p.Call("getBuffer")
	return DataBufferFromJSObject(retVal, v.ctx)
}

// GetData calls the GetData method on the VertexBuffer object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#getdata
func (v *VertexBuffer) GetData() []*float64 {

	retVal := v.p.Call("getData")
	return retVal
}

// GetInstanceDivisor calls the GetInstanceDivisor method on the VertexBuffer object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#getinstancedivisor
func (v *VertexBuffer) GetInstanceDivisor() float64 {

	retVal := v.p.Call("getInstanceDivisor")
	return retVal.Float()
}

// GetIsInstanced calls the GetIsInstanced method on the VertexBuffer object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#getisinstanced
func (v *VertexBuffer) GetIsInstanced() bool {

	retVal := v.p.Call("getIsInstanced")
	return retVal.Bool()
}

// GetKind calls the GetKind method on the VertexBuffer object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#getkind
func (v *VertexBuffer) GetKind() string {

	retVal := v.p.Call("getKind")
	return retVal.String()
}

// GetOffset calls the GetOffset method on the VertexBuffer object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#getoffset
func (v *VertexBuffer) GetOffset() float64 {

	retVal := v.p.Call("getOffset")
	return retVal.Float()
}

// GetSize calls the GetSize method on the VertexBuffer object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#getsize
func (v *VertexBuffer) GetSize() float64 {

	retVal := v.p.Call("getSize")
	return retVal.Float()
}

// GetStrideSize calls the GetStrideSize method on the VertexBuffer object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#getstridesize
func (v *VertexBuffer) GetStrideSize() float64 {

	retVal := v.p.Call("getStrideSize")
	return retVal.Float()
}

// GetTypeByteLength calls the GetTypeByteLength method on the VertexBuffer object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#gettypebytelength
func (v *VertexBuffer) GetTypeByteLength(jsType float64) float64 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, jsType)

	retVal := v.p.Call("GetTypeByteLength", args...)
	return retVal.Float()
}

// IsUpdatable calls the IsUpdatable method on the VertexBuffer object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#isupdatable
func (v *VertexBuffer) IsUpdatable() bool {

	retVal := v.p.Call("isUpdatable")
	return retVal.Bool()
}

// Update calls the Update method on the VertexBuffer object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#update
func (v *VertexBuffer) Update(data []*float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, float64ArrayToJSArray(data))

	v.p.Call("update", args...)
}

// VertexBufferUpdateDirectlyOpts contains optional parameters for VertexBuffer.UpdateDirectly.
type VertexBufferUpdateDirectlyOpts struct {
	UseBytes *bool
}

// UpdateDirectly calls the UpdateDirectly method on the VertexBuffer object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#updatedirectly
func (v *VertexBuffer) UpdateDirectly(data []*float64, offset float64, opts *VertexBufferUpdateDirectlyOpts) {
	if opts == nil {
		opts = &VertexBufferUpdateDirectlyOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, float64ArrayToJSArray(data))
	args = append(args, offset)

	if opts.UseBytes == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UseBytes)
	}

	v.p.Call("updateDirectly", args...)
}

/*

// BYTE returns the BYTE property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#byte
func (v *VertexBuffer) BYTE(BYTE float64) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(BYTE)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// SetBYTE sets the BYTE property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#byte
func (v *VertexBuffer) SetBYTE(BYTE float64) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(BYTE)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// ByteOffset returns the ByteOffset property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#byteoffset
func (v *VertexBuffer) ByteOffset(byteOffset float64) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(byteOffset)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// SetByteOffset sets the ByteOffset property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#byteoffset
func (v *VertexBuffer) SetByteOffset(byteOffset float64) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(byteOffset)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// ByteStride returns the ByteStride property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#bytestride
func (v *VertexBuffer) ByteStride(byteStride float64) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(byteStride)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// SetByteStride sets the ByteStride property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#bytestride
func (v *VertexBuffer) SetByteStride(byteStride float64) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(byteStride)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// ColorKind returns the ColorKind property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#colorkind
func (v *VertexBuffer) ColorKind(ColorKind string) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(ColorKind)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// SetColorKind sets the ColorKind property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#colorkind
func (v *VertexBuffer) SetColorKind(ColorKind string) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(ColorKind)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// FLOAT returns the FLOAT property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#float
func (v *VertexBuffer) FLOAT(FLOAT float64) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(FLOAT)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// SetFLOAT sets the FLOAT property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#float
func (v *VertexBuffer) SetFLOAT(FLOAT float64) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(FLOAT)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// INT returns the INT property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#int
func (v *VertexBuffer) INT(INT float64) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(INT)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// SetINT sets the INT property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#int
func (v *VertexBuffer) SetINT(INT float64) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(INT)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// InstanceDivisor returns the InstanceDivisor property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#instancedivisor
func (v *VertexBuffer) InstanceDivisor(instanceDivisor float64) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(instanceDivisor)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// SetInstanceDivisor sets the InstanceDivisor property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#instancedivisor
func (v *VertexBuffer) SetInstanceDivisor(instanceDivisor float64) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(instanceDivisor)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// MatricesIndicesExtraKind returns the MatricesIndicesExtraKind property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#matricesindicesextrakind
func (v *VertexBuffer) MatricesIndicesExtraKind(MatricesIndicesExtraKind string) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(MatricesIndicesExtraKind)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// SetMatricesIndicesExtraKind sets the MatricesIndicesExtraKind property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#matricesindicesextrakind
func (v *VertexBuffer) SetMatricesIndicesExtraKind(MatricesIndicesExtraKind string) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(MatricesIndicesExtraKind)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// MatricesIndicesKind returns the MatricesIndicesKind property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#matricesindiceskind
func (v *VertexBuffer) MatricesIndicesKind(MatricesIndicesKind string) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(MatricesIndicesKind)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// SetMatricesIndicesKind sets the MatricesIndicesKind property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#matricesindiceskind
func (v *VertexBuffer) SetMatricesIndicesKind(MatricesIndicesKind string) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(MatricesIndicesKind)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// MatricesWeightsExtraKind returns the MatricesWeightsExtraKind property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#matricesweightsextrakind
func (v *VertexBuffer) MatricesWeightsExtraKind(MatricesWeightsExtraKind string) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(MatricesWeightsExtraKind)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// SetMatricesWeightsExtraKind sets the MatricesWeightsExtraKind property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#matricesweightsextrakind
func (v *VertexBuffer) SetMatricesWeightsExtraKind(MatricesWeightsExtraKind string) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(MatricesWeightsExtraKind)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// MatricesWeightsKind returns the MatricesWeightsKind property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#matricesweightskind
func (v *VertexBuffer) MatricesWeightsKind(MatricesWeightsKind string) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(MatricesWeightsKind)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// SetMatricesWeightsKind sets the MatricesWeightsKind property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#matricesweightskind
func (v *VertexBuffer) SetMatricesWeightsKind(MatricesWeightsKind string) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(MatricesWeightsKind)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// NormalKind returns the NormalKind property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#normalkind
func (v *VertexBuffer) NormalKind(NormalKind string) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(NormalKind)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// SetNormalKind sets the NormalKind property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#normalkind
func (v *VertexBuffer) SetNormalKind(NormalKind string) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(NormalKind)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// Normalized returns the Normalized property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#normalized
func (v *VertexBuffer) Normalized(normalized bool) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(normalized)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// SetNormalized sets the Normalized property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#normalized
func (v *VertexBuffer) SetNormalized(normalized bool) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(normalized)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// PositionKind returns the PositionKind property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#positionkind
func (v *VertexBuffer) PositionKind(PositionKind string) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(PositionKind)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// SetPositionKind sets the PositionKind property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#positionkind
func (v *VertexBuffer) SetPositionKind(PositionKind string) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(PositionKind)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// SHORT returns the SHORT property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#short
func (v *VertexBuffer) SHORT(SHORT float64) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(SHORT)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// SetSHORT sets the SHORT property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#short
func (v *VertexBuffer) SetSHORT(SHORT float64) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(SHORT)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// TangentKind returns the TangentKind property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#tangentkind
func (v *VertexBuffer) TangentKind(TangentKind string) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(TangentKind)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// SetTangentKind sets the TangentKind property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#tangentkind
func (v *VertexBuffer) SetTangentKind(TangentKind string) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(TangentKind)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// Type returns the Type property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#type
func (v *VertexBuffer) Type(jsType float64) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(jsType)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// SetType sets the Type property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#type
func (v *VertexBuffer) SetType(jsType float64) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(jsType)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// UNSIGNED_BYTE returns the UNSIGNED_BYTE property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#unsigned_byte
func (v *VertexBuffer) UNSIGNED_BYTE(UNSIGNED_BYTE float64) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(UNSIGNED_BYTE)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// SetUNSIGNED_BYTE sets the UNSIGNED_BYTE property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#unsigned_byte
func (v *VertexBuffer) SetUNSIGNED_BYTE(UNSIGNED_BYTE float64) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(UNSIGNED_BYTE)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// UNSIGNED_INT returns the UNSIGNED_INT property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#unsigned_int
func (v *VertexBuffer) UNSIGNED_INT(UNSIGNED_INT float64) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(UNSIGNED_INT)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// SetUNSIGNED_INT sets the UNSIGNED_INT property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#unsigned_int
func (v *VertexBuffer) SetUNSIGNED_INT(UNSIGNED_INT float64) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(UNSIGNED_INT)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// UNSIGNED_SHORT returns the UNSIGNED_SHORT property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#unsigned_short
func (v *VertexBuffer) UNSIGNED_SHORT(UNSIGNED_SHORT float64) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(UNSIGNED_SHORT)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// SetUNSIGNED_SHORT sets the UNSIGNED_SHORT property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#unsigned_short
func (v *VertexBuffer) SetUNSIGNED_SHORT(UNSIGNED_SHORT float64) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(UNSIGNED_SHORT)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// UV2Kind returns the UV2Kind property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#uv2kind
func (v *VertexBuffer) UV2Kind(UV2Kind string) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(UV2Kind)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// SetUV2Kind sets the UV2Kind property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#uv2kind
func (v *VertexBuffer) SetUV2Kind(UV2Kind string) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(UV2Kind)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// UV3Kind returns the UV3Kind property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#uv3kind
func (v *VertexBuffer) UV3Kind(UV3Kind string) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(UV3Kind)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// SetUV3Kind sets the UV3Kind property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#uv3kind
func (v *VertexBuffer) SetUV3Kind(UV3Kind string) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(UV3Kind)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// UV4Kind returns the UV4Kind property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#uv4kind
func (v *VertexBuffer) UV4Kind(UV4Kind string) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(UV4Kind)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// SetUV4Kind sets the UV4Kind property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#uv4kind
func (v *VertexBuffer) SetUV4Kind(UV4Kind string) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(UV4Kind)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// UV5Kind returns the UV5Kind property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#uv5kind
func (v *VertexBuffer) UV5Kind(UV5Kind string) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(UV5Kind)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// SetUV5Kind sets the UV5Kind property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#uv5kind
func (v *VertexBuffer) SetUV5Kind(UV5Kind string) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(UV5Kind)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// UV6Kind returns the UV6Kind property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#uv6kind
func (v *VertexBuffer) UV6Kind(UV6Kind string) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(UV6Kind)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// SetUV6Kind sets the UV6Kind property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#uv6kind
func (v *VertexBuffer) SetUV6Kind(UV6Kind string) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(UV6Kind)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// UVKind returns the UVKind property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#uvkind
func (v *VertexBuffer) UVKind(UVKind string) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(UVKind)
	return VertexBufferFromJSObject(p, ba.ctx)
}

// SetUVKind sets the UVKind property of class VertexBuffer.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer#uvkind
func (v *VertexBuffer) SetUVKind(UVKind string) *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer").New(UVKind)
	return VertexBufferFromJSObject(p, ba.ctx)
}

*/
