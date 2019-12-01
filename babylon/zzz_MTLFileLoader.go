// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// MTLFileLoader represents a babylon.js MTLFileLoader.
// Class reading and parsing the MTL file bundled with the obj file.
type MTLFileLoader struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (m *MTLFileLoader) JSObject() js.Value { return m.p }

// MTLFileLoader returns a MTLFileLoader JavaScript class.
func (ba *Babylon) MTLFileLoader() *MTLFileLoader {
	p := ba.ctx.Get("MTLFileLoader")
	return MTLFileLoaderFromJSObject(p, ba.ctx)
}

// MTLFileLoaderFromJSObject returns a wrapped MTLFileLoader JavaScript class.
func MTLFileLoaderFromJSObject(p js.Value, ctx js.Value) *MTLFileLoader {
	return &MTLFileLoader{p: p, ctx: ctx}
}

// ParseMTL calls the ParseMTL method on the MTLFileLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.mtlfileloader#parsemtl
func (m *MTLFileLoader) ParseMTL(scene *Scene, data string, rootUrl string) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, scene.JSObject())
	args = append(args, data)
	args = append(args, rootUrl)

	m.p.Call("parseMTL", args...)
}

/*

// Materials returns the Materials property of class MTLFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.mtlfileloader#materials
func (m *MTLFileLoader) Materials(materials *StandardMaterial) *MTLFileLoader {
	p := ba.ctx.Get("MTLFileLoader").New(materials.JSObject())
	return MTLFileLoaderFromJSObject(p, ba.ctx)
}

// SetMaterials sets the Materials property of class MTLFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.mtlfileloader#materials
func (m *MTLFileLoader) SetMaterials(materials *StandardMaterial) *MTLFileLoader {
	p := ba.ctx.Get("MTLFileLoader").New(materials.JSObject())
	return MTLFileLoaderFromJSObject(p, ba.ctx)
}

*/
