// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// XmlLoader represents a babylon.js XmlLoader.
// Class used to load GUI via XML.
type XmlLoader struct{}

// JSObject returns the underlying js.Value.
func (x *XmlLoader) JSObject() js.Value { return x.p }

// XmlLoader returns a XmlLoader JavaScript class.
func (b *Babylon) XmlLoader() *XmlLoader {
	p := b.ctx.Get("XmlLoader")
	return XmlLoaderFromJSObject(p)
}

// XmlLoaderFromJSObject returns a wrapped XmlLoader JavaScript class.
func XmlLoaderFromJSObject(p js.Value) *XmlLoader {
	return &XmlLoader{p: p}
}

// NewXmlLoader returns a new XmlLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.xmlloader
func (b *Babylon) NewXmlLoader(todo parameters) *XmlLoader {
	p := b.ctx.Get("XmlLoader").New(todo)
	return XmlLoaderFromJSObject(p)
}

// TODO: methods
