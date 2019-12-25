// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// XmlLoader represents a babylon.js XmlLoader.
// Class used to load GUI via XML.
type XmlLoader struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (x *XmlLoader) JSObject() js.Value { return x.p }

// XmlLoader returns a XmlLoader JavaScript class.
func (gui *GUI) XmlLoader() *XmlLoader {
	p := gui.ctx.Get("XmlLoader")
	return XmlLoaderFromJSObject(p, gui.ctx)
}

// XmlLoaderFromJSObject returns a wrapped XmlLoader JavaScript class.
func XmlLoaderFromJSObject(p js.Value, ctx js.Value) *XmlLoader {
	return &XmlLoader{p: p, ctx: ctx}
}

// XmlLoaderArrayToJSArray returns a JavaScript Array for the wrapped array.
func XmlLoaderArrayToJSArray(array []*XmlLoader) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewXmlLoaderOpts contains optional parameters for NewXmlLoader.
type NewXmlLoaderOpts struct {
	ParentClass js.Value
}

// NewXmlLoader returns a new XmlLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.xmlloader#constructor
func (gui *GUI) NewXmlLoader(opts *NewXmlLoaderOpts) *XmlLoader {
	if opts == nil {
		opts = &NewXmlLoaderOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	args = append(args, opts.ParentClass)

	p := gui.ctx.Get("XmlLoader").New(args...)
	return XmlLoaderFromJSObject(p, gui.ctx)
}

// GetNodeById calls the GetNodeById method on the XmlLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.xmlloader#getnodebyid
func (x *XmlLoader) GetNodeById(id string) js.Value {

	args := make([]interface{}, 0, 1+0)

	args = append(args, id)

	retVal := x.p.Call("getNodeById", args...)
	return retVal
}

// GetNodes calls the GetNodes method on the XmlLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.xmlloader#getnodes
func (x *XmlLoader) GetNodes() js.Value {

	retVal := x.p.Call("getNodes")
	return retVal
}

// IsLoaded calls the IsLoaded method on the XmlLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.xmlloader#isloaded
func (x *XmlLoader) IsLoaded() bool {

	retVal := x.p.Call("isLoaded")
	return retVal.Bool()
}

// LoadLayout calls the LoadLayout method on the XmlLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.xmlloader#loadlayout
func (x *XmlLoader) LoadLayout(xmlFile JSObject, rootNode JSObject, callback JSObject) {

	args := make([]interface{}, 0, 3+0)

	if xmlFile == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, xmlFile.JSObject())
	}

	if rootNode == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, rootNode.JSObject())
	}

	if callback == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, callback.JSObject())
	}

	x.p.Call("loadLayout", args...)
}
