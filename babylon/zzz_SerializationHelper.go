// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SerializationHelper represents a babylon.js SerializationHelper.
// Class used to help serialization objects
type SerializationHelper struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (s *SerializationHelper) JSObject() js.Value { return s.p }

// SerializationHelper returns a SerializationHelper JavaScript class.
func (b *Babylon) SerializationHelper() *SerializationHelper {
	p := b.ctx.Get("SerializationHelper")
	return SerializationHelperFromJSObject(p)
}

// SerializationHelperFromJSObject returns a wrapped SerializationHelper JavaScript class.
func SerializationHelperFromJSObject(p js.Value) *SerializationHelper {
	return &SerializationHelper{p: p}
}

// TODO: methods
