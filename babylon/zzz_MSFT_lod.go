// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// MSFT_lod represents a babylon.js MSFT_lod.
// &lt;a href=&#34;https://github.com/KhronosGroup/glTF/tree/master/extensions/2.0/Vendor/MSFT_lod&#34;&gt;Specification&lt;/a&gt;
type MSFT_lod struct{}

// JSObject returns the underlying js.Value.
func (m *MSFT_lod) JSObject() js.Value { return m.p }

// MSFT_lod returns a MSFT_lod JavaScript class.
func (b *Babylon) MSFT_lod() *MSFT_lod {
	p := b.ctx.Get("MSFT_lod")
	return MSFT_lodFromJSObject(p)
}

// MSFT_lodFromJSObject returns a wrapped MSFT_lod JavaScript class.
func MSFT_lodFromJSObject(p js.Value) *MSFT_lod {
	return &MSFT_lod{p: p}
}

// NewMSFT_lod returns a new MSFT_lod object.
//
// https://doc.babylonjs.com/api/classes/babylon.msft_lod
func (b *Babylon) NewMSFT_lod(todo parameters) *MSFT_lod {
	p := b.ctx.Get("MSFT_lod").New(todo)
	return MSFT_lodFromJSObject(p)
}

// TODO: methods
