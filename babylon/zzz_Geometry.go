// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Geometry represents a babylon.js Geometry.
// Class used to store geometry data (vertex buffers + index buffer)
type Geometry struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (g *Geometry) JSObject() js.Value { return g.p }

// Geometry returns a Geometry JavaScript class.
func (ba *Babylon) Geometry() *Geometry {
	p := ba.ctx.Get("Geometry")
	return GeometryFromJSObject(p, ba.ctx)
}

// GeometryFromJSObject returns a wrapped Geometry JavaScript class.
func GeometryFromJSObject(p js.Value, ctx js.Value) *Geometry {
	return &Geometry{p: p, ctx: ctx}
}

// GeometryArrayToJSArray returns a JavaScript Array for the wrapped array.
func GeometryArrayToJSArray(array []*Geometry) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewGeometryOpts contains optional parameters for NewGeometry.
type NewGeometryOpts struct {
	VertexData *VertexData
	Updatable  *bool
	Mesh       *Mesh
}

// NewGeometry returns a new Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry
func (ba *Babylon) NewGeometry(id string, scene *Scene, opts *NewGeometryOpts) *Geometry {
	if opts == nil {
		opts = &NewGeometryOpts{}
	}

	args := make([]interface{}, 0, 2+3)

	args = append(args, id)
	args = append(args, scene.JSObject())

	if opts.VertexData == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.VertexData.JSObject())
	}
	if opts.Updatable == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Updatable)
	}
	if opts.Mesh == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Mesh.JSObject())
	}

	p := ba.ctx.Get("Geometry").New(args...)
	return GeometryFromJSObject(p, ba.ctx)
}

// ApplyToMesh calls the ApplyToMesh method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#applytomesh
func (g *Geometry) ApplyToMesh(mesh *Mesh) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, mesh.JSObject())

	g.p.Call("applyToMesh", args...)
}

// Copy calls the Copy method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#copy
func (g *Geometry) Copy(id string) *Geometry {

	args := make([]interface{}, 0, 1+0)

	args = append(args, id)

	retVal := g.p.Call("copy", args...)
	return GeometryFromJSObject(retVal, g.ctx)
}

// CreateGeometryForMesh calls the CreateGeometryForMesh method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#creategeometryformesh
func (g *Geometry) CreateGeometryForMesh(mesh *Mesh) *Geometry {

	args := make([]interface{}, 0, 1+0)

	args = append(args, mesh.JSObject())

	retVal := g.p.Call("CreateGeometryForMesh", args...)
	return GeometryFromJSObject(retVal, g.ctx)
}

// Dispose calls the Dispose method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#dispose
func (g *Geometry) Dispose() {

	g.p.Call("dispose")
}

// ExtractFromMesh calls the ExtractFromMesh method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#extractfrommesh
func (g *Geometry) ExtractFromMesh(mesh *Mesh, id string) *Geometry {

	args := make([]interface{}, 0, 2+0)

	args = append(args, mesh.JSObject())
	args = append(args, id)

	retVal := g.p.Call("ExtractFromMesh", args...)
	return GeometryFromJSObject(retVal, g.ctx)
}

// GetEngine calls the GetEngine method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#getengine
func (g *Geometry) GetEngine() *Engine {

	retVal := g.p.Call("getEngine")
	return EngineFromJSObject(retVal, g.ctx)
}

// GetIndexBuffer calls the GetIndexBuffer method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#getindexbuffer
func (g *Geometry) GetIndexBuffer() *DataBuffer {

	retVal := g.p.Call("getIndexBuffer")
	return DataBufferFromJSObject(retVal, g.ctx)
}

// GeometryGetIndicesOpts contains optional parameters for Geometry.GetIndices.
type GeometryGetIndicesOpts struct {
	CopyWhenShared *bool
	ForceCopy      *bool
}

// GetIndices calls the GetIndices method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#getindices
func (g *Geometry) GetIndices(opts *GeometryGetIndicesOpts) js.Value {
	if opts == nil {
		opts = &GeometryGetIndicesOpts{}
	}

	args := make([]interface{}, 0, 0+2)

	if opts.CopyWhenShared == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.CopyWhenShared)
	}
	if opts.ForceCopy == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ForceCopy)
	}

	retVal := g.p.Call("getIndices", args...)
	return retVal
}

// GetScene calls the GetScene method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#getscene
func (g *Geometry) GetScene() *Scene {

	retVal := g.p.Call("getScene")
	return SceneFromJSObject(retVal, g.ctx)
}

// GetTotalIndices calls the GetTotalIndices method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#gettotalindices
func (g *Geometry) GetTotalIndices() float64 {

	retVal := g.p.Call("getTotalIndices")
	return retVal.Float()
}

// GetTotalVertices calls the GetTotalVertices method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#gettotalvertices
func (g *Geometry) GetTotalVertices() float64 {

	retVal := g.p.Call("getTotalVertices")
	return retVal.Float()
}

// GetVertexBuffer calls the GetVertexBuffer method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#getvertexbuffer
func (g *Geometry) GetVertexBuffer(kind string) *VertexBuffer {

	args := make([]interface{}, 0, 1+0)

	args = append(args, kind)

	retVal := g.p.Call("getVertexBuffer", args...)
	return VertexBufferFromJSObject(retVal, g.ctx)
}

// GetVertexBuffers calls the GetVertexBuffers method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#getvertexbuffers
func (g *Geometry) GetVertexBuffers() js.Value {

	retVal := g.p.Call("getVertexBuffers")
	return retVal
}

// GeometryGetVerticesDataOpts contains optional parameters for Geometry.GetVerticesData.
type GeometryGetVerticesDataOpts struct {
	CopyWhenShared *bool
	ForceCopy      *bool
}

// GetVerticesData calls the GetVerticesData method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#getverticesdata
func (g *Geometry) GetVerticesData(kind string, opts *GeometryGetVerticesDataOpts) js.Value {
	if opts == nil {
		opts = &GeometryGetVerticesDataOpts{}
	}

	args := make([]interface{}, 0, 1+2)

	args = append(args, kind)

	if opts.CopyWhenShared == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.CopyWhenShared)
	}
	if opts.ForceCopy == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ForceCopy)
	}

	retVal := g.p.Call("getVerticesData", args...)
	return retVal
}

// GetVerticesDataKinds calls the GetVerticesDataKinds method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#getverticesdatakinds
func (g *Geometry) GetVerticesDataKinds() string {

	retVal := g.p.Call("getVerticesDataKinds")
	return retVal.String()
}

// IsDisposed calls the IsDisposed method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#isdisposed
func (g *Geometry) IsDisposed() bool {

	retVal := g.p.Call("isDisposed")
	return retVal.Bool()
}

// IsReady calls the IsReady method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#isready
func (g *Geometry) IsReady() bool {

	retVal := g.p.Call("isReady")
	return retVal.Bool()
}

// IsVertexBufferUpdatable calls the IsVertexBufferUpdatable method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#isvertexbufferupdatable
func (g *Geometry) IsVertexBufferUpdatable(kind string) bool {

	args := make([]interface{}, 0, 1+0)

	args = append(args, kind)

	retVal := g.p.Call("isVertexBufferUpdatable", args...)
	return retVal.Bool()
}

// IsVerticesDataPresent calls the IsVerticesDataPresent method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#isverticesdatapresent
func (g *Geometry) IsVerticesDataPresent(kind string) bool {

	args := make([]interface{}, 0, 1+0)

	args = append(args, kind)

	retVal := g.p.Call("isVerticesDataPresent", args...)
	return retVal.Bool()
}

// GeometryLoadOpts contains optional parameters for Geometry.Load.
type GeometryLoadOpts struct {
	OnLoaded func()
}

// Load calls the Load method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#load
func (g *Geometry) Load(scene *Scene, opts *GeometryLoadOpts) {
	if opts == nil {
		opts = &GeometryLoadOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, scene.JSObject())

	if opts.OnLoaded == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { opts.OnLoaded(); return nil }) /* never freed! */)
	}

	g.p.Call("load", args...)
}

// Parse calls the Parse method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#parse
func (g *Geometry) Parse(parsedVertexData interface{}, scene *Scene, rootUrl string) *Geometry {

	args := make([]interface{}, 0, 3+0)

	args = append(args, parsedVertexData)
	args = append(args, scene.JSObject())
	args = append(args, rootUrl)

	retVal := g.p.Call("Parse", args...)
	return GeometryFromJSObject(retVal, g.ctx)
}

// RandomId calls the RandomId method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#randomid
func (g *Geometry) RandomId() string {

	retVal := g.p.Call("RandomId")
	return retVal.String()
}

// GeometryReleaseForMeshOpts contains optional parameters for Geometry.ReleaseForMesh.
type GeometryReleaseForMeshOpts struct {
	ShouldDispose *bool
}

// ReleaseForMesh calls the ReleaseForMesh method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#releaseformesh
func (g *Geometry) ReleaseForMesh(mesh *Mesh, opts *GeometryReleaseForMeshOpts) {
	if opts == nil {
		opts = &GeometryReleaseForMeshOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, mesh.JSObject())

	if opts.ShouldDispose == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ShouldDispose)
	}

	g.p.Call("releaseForMesh", args...)
}

// RemoveVerticesData calls the RemoveVerticesData method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#removeverticesdata
func (g *Geometry) RemoveVerticesData(kind string) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, kind)

	g.p.Call("removeVerticesData", args...)
}

// Serialize calls the Serialize method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#serialize
func (g *Geometry) Serialize() interface{} {

	retVal := g.p.Call("serialize")
	return retVal
}

// SerializeVerticeData calls the SerializeVerticeData method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#serializeverticedata
func (g *Geometry) SerializeVerticeData() interface{} {

	retVal := g.p.Call("serializeVerticeData")
	return retVal
}

// GeometrySetAllVerticesDataOpts contains optional parameters for Geometry.SetAllVerticesData.
type GeometrySetAllVerticesDataOpts struct {
	Updatable *bool
}

// SetAllVerticesData calls the SetAllVerticesData method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#setallverticesdata
func (g *Geometry) SetAllVerticesData(vertexData *VertexData, opts *GeometrySetAllVerticesDataOpts) {
	if opts == nil {
		opts = &GeometrySetAllVerticesDataOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, vertexData.JSObject())

	if opts.Updatable == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Updatable)
	}

	g.p.Call("setAllVerticesData", args...)
}

// GeometrySetIndicesOpts contains optional parameters for Geometry.SetIndices.
type GeometrySetIndicesOpts struct {
	TotalVertices *float64
	Updatable     *bool
}

// SetIndices calls the SetIndices method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#setindices
func (g *Geometry) SetIndices(indices js.Value, opts *GeometrySetIndicesOpts) {
	if opts == nil {
		opts = &GeometrySetIndicesOpts{}
	}

	args := make([]interface{}, 0, 1+2)

	args = append(args, indices)

	if opts.TotalVertices == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.TotalVertices)
	}
	if opts.Updatable == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Updatable)
	}

	g.p.Call("setIndices", args...)
}

// GeometrySetVerticesBufferOpts contains optional parameters for Geometry.SetVerticesBuffer.
type GeometrySetVerticesBufferOpts struct {
	TotalVertices *float64
}

// SetVerticesBuffer calls the SetVerticesBuffer method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#setverticesbuffer
func (g *Geometry) SetVerticesBuffer(buffer *VertexBuffer, opts *GeometrySetVerticesBufferOpts) {
	if opts == nil {
		opts = &GeometrySetVerticesBufferOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, buffer.JSObject())

	if opts.TotalVertices == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.TotalVertices)
	}

	g.p.Call("setVerticesBuffer", args...)
}

// GeometrySetVerticesDataOpts contains optional parameters for Geometry.SetVerticesData.
type GeometrySetVerticesDataOpts struct {
	Updatable *bool
	Stride    *float64
}

// SetVerticesData calls the SetVerticesData method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#setverticesdata
func (g *Geometry) SetVerticesData(kind string, data js.Value, opts *GeometrySetVerticesDataOpts) {
	if opts == nil {
		opts = &GeometrySetVerticesDataOpts{}
	}

	args := make([]interface{}, 0, 2+2)

	args = append(args, kind)
	args = append(args, data)

	if opts.Updatable == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Updatable)
	}
	if opts.Stride == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Stride)
	}

	g.p.Call("setVerticesData", args...)
}

// ToLeftHanded calls the ToLeftHanded method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#tolefthanded
func (g *Geometry) ToLeftHanded() {

	g.p.Call("toLeftHanded")
}

// GeometryUpdateIndicesOpts contains optional parameters for Geometry.UpdateIndices.
type GeometryUpdateIndicesOpts struct {
	Offset        *float64
	GpuMemoryOnly *bool
}

// UpdateIndices calls the UpdateIndices method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#updateindices
func (g *Geometry) UpdateIndices(indices js.Value, opts *GeometryUpdateIndicesOpts) {
	if opts == nil {
		opts = &GeometryUpdateIndicesOpts{}
	}

	args := make([]interface{}, 0, 1+2)

	args = append(args, indices)

	if opts.Offset == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Offset)
	}
	if opts.GpuMemoryOnly == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.GpuMemoryOnly)
	}

	g.p.Call("updateIndices", args...)
}

// GeometryUpdateVerticesDataOpts contains optional parameters for Geometry.UpdateVerticesData.
type GeometryUpdateVerticesDataOpts struct {
	UpdateExtends *bool
}

// UpdateVerticesData calls the UpdateVerticesData method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#updateverticesdata
func (g *Geometry) UpdateVerticesData(kind string, data js.Value, opts *GeometryUpdateVerticesDataOpts) {
	if opts == nil {
		opts = &GeometryUpdateVerticesDataOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, kind)
	args = append(args, data)

	if opts.UpdateExtends == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UpdateExtends)
	}

	g.p.Call("updateVerticesData", args...)
}

// GeometryUpdateVerticesDataDirectlyOpts contains optional parameters for Geometry.UpdateVerticesDataDirectly.
type GeometryUpdateVerticesDataDirectlyOpts struct {
	UseBytes *bool
}

// UpdateVerticesDataDirectly calls the UpdateVerticesDataDirectly method on the Geometry object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#updateverticesdatadirectly
func (g *Geometry) UpdateVerticesDataDirectly(kind string, data []float64, offset float64, opts *GeometryUpdateVerticesDataDirectlyOpts) {
	if opts == nil {
		opts = &GeometryUpdateVerticesDataDirectlyOpts{}
	}

	args := make([]interface{}, 0, 3+1)

	args = append(args, kind)
	args = append(args, data)
	args = append(args, offset)

	if opts.UseBytes == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UseBytes)
	}

	g.p.Call("updateVerticesDataDirectly", args...)
}

// BoundingBias returns the BoundingBias property of class Geometry.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#boundingbias
func (g *Geometry) BoundingBias() *Vector2 {
	retVal := g.p.Get("boundingBias")
	return Vector2FromJSObject(retVal, g.ctx)
}

// SetBoundingBias sets the BoundingBias property of class Geometry.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#boundingbias
func (g *Geometry) SetBoundingBias(boundingBias *Vector2) *Geometry {
	g.p.Set("boundingBias", boundingBias.JSObject())
	return g
}

// DelayLoadState returns the DelayLoadState property of class Geometry.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#delayloadstate
func (g *Geometry) DelayLoadState() float64 {
	retVal := g.p.Get("delayLoadState")
	return retVal.Float()
}

// SetDelayLoadState sets the DelayLoadState property of class Geometry.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#delayloadstate
func (g *Geometry) SetDelayLoadState(delayLoadState float64) *Geometry {
	g.p.Set("delayLoadState", delayLoadState)
	return g
}

// DelayLoadingFile returns the DelayLoadingFile property of class Geometry.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#delayloadingfile
func (g *Geometry) DelayLoadingFile() string {
	retVal := g.p.Get("delayLoadingFile")
	return retVal.String()
}

// SetDelayLoadingFile sets the DelayLoadingFile property of class Geometry.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#delayloadingfile
func (g *Geometry) SetDelayLoadingFile(delayLoadingFile string) *Geometry {
	g.p.Set("delayLoadingFile", delayLoadingFile)
	return g
}

// DoNotSerialize returns the DoNotSerialize property of class Geometry.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#donotserialize
func (g *Geometry) DoNotSerialize() bool {
	retVal := g.p.Get("doNotSerialize")
	return retVal.Bool()
}

// SetDoNotSerialize sets the DoNotSerialize property of class Geometry.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#donotserialize
func (g *Geometry) SetDoNotSerialize(doNotSerialize bool) *Geometry {
	g.p.Set("doNotSerialize", doNotSerialize)
	return g
}

// Extend returns the Extend property of class Geometry.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#extend
func (g *Geometry) Extend() js.Value {
	retVal := g.p.Get("extend")
	return retVal
}

// SetExtend sets the Extend property of class Geometry.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#extend
func (g *Geometry) SetExtend(extend js.Value) *Geometry {
	g.p.Set("extend", extend)
	return g
}

// Id returns the Id property of class Geometry.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#id
func (g *Geometry) Id() string {
	retVal := g.p.Get("id")
	return retVal.String()
}

// SetId sets the Id property of class Geometry.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#id
func (g *Geometry) SetId(id string) *Geometry {
	g.p.Set("id", id)
	return g
}

// OnGeometryUpdated returns the OnGeometryUpdated property of class Geometry.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#ongeometryupdated
func (g *Geometry) OnGeometryUpdated() js.Value {
	retVal := g.p.Get("onGeometryUpdated")
	return retVal
}

// SetOnGeometryUpdated sets the OnGeometryUpdated property of class Geometry.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#ongeometryupdated
func (g *Geometry) SetOnGeometryUpdated(onGeometryUpdated func()) *Geometry {
	g.p.Set("onGeometryUpdated", js.FuncOf(func(this js.Value, args []js.Value) interface{} { onGeometryUpdated(); return nil }))
	return g
}

// UniqueId returns the UniqueId property of class Geometry.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#uniqueid
func (g *Geometry) UniqueId() float64 {
	retVal := g.p.Get("uniqueId")
	return retVal.Float()
}

// SetUniqueId sets the UniqueId property of class Geometry.
//
// https://doc.babylonjs.com/api/classes/babylon.geometry#uniqueid
func (g *Geometry) SetUniqueId(uniqueId float64) *Geometry {
	g.p.Set("uniqueId", uniqueId)
	return g
}
