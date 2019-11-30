// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// InputText represents a babylon.js InputText.
// Class used to create input text control
type InputText struct{ *Control }

// JSObject returns the underlying js.Value.
func (i *InputText) JSObject() js.Value { return i.p }

// InputText returns a InputText JavaScript class.
func (b *Babylon) InputText() *InputText {
	p := b.ctx.Get("InputText")
	return InputTextFromJSObject(p)
}

// InputTextFromJSObject returns a wrapped InputText JavaScript class.
func InputTextFromJSObject(p js.Value) *InputText {
	return &InputText{ControlFromJSObject(p)}
}

// NewInputTextOpts contains optional parameters for NewInputText.
type NewInputTextOpts struct {
	Name *string

	Text *string
}

// NewInputText returns a new InputText object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputtext
func (b *Babylon) NewInputText(opts *NewInputTextOpts) *InputText {
	if opts == nil {
		opts = &NewInputTextOpts{}
	}

	p := b.ctx.Get("InputText").New(opts.Name, opts.Text)
	return InputTextFromJSObject(p)
}

// TODO: methods
