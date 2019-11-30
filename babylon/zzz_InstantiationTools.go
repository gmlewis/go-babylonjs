// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// InstantiationTools represents a babylon.js InstantiationTools.
// Class used to enable instatition of objects by class name
type InstantiationTools struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (i *InstantiationTools) JSObject() js.Value { return i.p }

// InstantiationTools returns a InstantiationTools JavaScript class.
func (b *Babylon) InstantiationTools() *InstantiationTools {
	p := b.ctx.Get("InstantiationTools")
	return InstantiationToolsFromJSObject(p)
}

// InstantiationToolsFromJSObject returns a wrapped InstantiationTools JavaScript class.
func InstantiationToolsFromJSObject(p js.Value) *InstantiationTools {
	return &InstantiationTools{p: p}
}

// TODO: methods
