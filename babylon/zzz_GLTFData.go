// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GLTFData represents a babylon.js GLTFData.
// Class for holding and downloading glTF file data
type GLTFData struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (g *GLTFData) JSObject() js.Value { return g.p }

// GLTFData returns a GLTFData JavaScript class.
func (ba *Babylon) GLTFData() *GLTFData {
	p := ba.ctx.Get("GLTFData")
	return GLTFDataFromJSObject(p, ba.ctx)
}

// GLTFDataFromJSObject returns a wrapped GLTFData JavaScript class.
func GLTFDataFromJSObject(p js.Value, ctx js.Value) *GLTFData {
	return &GLTFData{p: p, ctx: ctx}
}

// GLTFDataArrayToJSArray returns a JavaScript Array for the wrapped array.
func GLTFDataArrayToJSArray(array []*GLTFData) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewGLTFData returns a new GLTFData object.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfdata#constructor
func (ba *Babylon) NewGLTFData() *GLTFData {

	args := make([]interface{}, 0, 0+0)

	p := ba.ctx.Get("GLTFData").New(args...)
	return GLTFDataFromJSObject(p, ba.ctx)
}

// DownloadFiles calls the DownloadFiles method on the GLTFData object.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfdata#downloadfiles
func (g *GLTFData) DownloadFiles() {

	g.p.Call("downloadFiles")
}

// GlTFFiles returns the GlTFFiles property of class GLTFData.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfdata#gltffiles
func (g *GLTFData) GlTFFiles() js.Value {
	retVal := g.p.Get("glTFFiles")
	return retVal
}

// SetGlTFFiles sets the GlTFFiles property of class GLTFData.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfdata#gltffiles
func (g *GLTFData) SetGlTFFiles(glTFFiles js.Value) *GLTFData {
	g.p.Set("glTFFiles", glTFFiles)
	return g
}
