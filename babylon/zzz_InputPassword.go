// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// InputPassword represents a babylon.js InputPassword.
// Class used to create a password control
type InputPassword struct{ *InputText }

// JSObject returns the underlying js.Value.
func (i *InputPassword) JSObject() js.Value { return i.p }

// InputPassword returns a InputPassword JavaScript class.
func (b *Babylon) InputPassword() *InputPassword {
	p := b.ctx.Get("InputPassword")
	return InputPasswordFromJSObject(p)
}

// InputPasswordFromJSObject returns a wrapped InputPassword JavaScript class.
func InputPasswordFromJSObject(p js.Value) *InputPassword {
	return &InputPassword{InputTextFromJSObject(p)}
}

// NewInputPasswordOpts contains optional parameters for NewInputPassword.
type NewInputPasswordOpts struct {
	Name *string

	Text *string
}

// NewInputPassword returns a new InputPassword object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputpassword
func (b *Babylon) NewInputPassword(opts *NewInputPasswordOpts) *InputPassword {
	if opts == nil {
		opts = &NewInputPasswordOpts{}
	}

	p := b.ctx.Get("InputPassword").New(opts.Name, opts.Text)
	return InputPasswordFromJSObject(p)
}

// TODO: methods
