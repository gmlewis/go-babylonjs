// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// MultiLine represents a babylon.js MultiLine.
// Class used to create multi line control
type MultiLine struct{ *Control }

// JSObject returns the underlying js.Value.
func (m *MultiLine) JSObject() js.Value { return m.p }

// MultiLine returns a MultiLine JavaScript class.
func (b *Babylon) MultiLine() *MultiLine {
	p := b.ctx.Get("MultiLine")
	return MultiLineFromJSObject(p)
}

// MultiLineFromJSObject returns a wrapped MultiLine JavaScript class.
func MultiLineFromJSObject(p js.Value) *MultiLine {
	return &MultiLine{ControlFromJSObject(p)}
}

// NewMultiLineOpts contains optional parameters for NewMultiLine.
type NewMultiLineOpts struct {
	Name *string
}

// NewMultiLine returns a new MultiLine object.
//
// https://doc.babylonjs.com/api/classes/babylon.multiline
func (b *Babylon) NewMultiLine(opts *NewMultiLineOpts) *MultiLine {
	if opts == nil {
		opts = &NewMultiLineOpts{}
	}

	p := b.ctx.Get("MultiLine").New(opts.Name)
	return MultiLineFromJSObject(p)
}

// TODO: methods
