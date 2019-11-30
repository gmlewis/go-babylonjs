// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// BasisTools represents a babylon.js BasisTools.
// Used to load .Basis files
// See &lt;a href=&#34;https://github.com/BinomialLLC/basis_universal/tree/master/webgl&#34;&gt;https://github.com/BinomialLLC/basis_universal/tree/master/webgl&lt;/a&gt;
type BasisTools struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (b *BasisTools) JSObject() js.Value { return b.p }

// BasisTools returns a BasisTools JavaScript class.
func (ba *Babylon) BasisTools() *BasisTools {
	p := ba.ctx.Get("BasisTools")
	return BasisToolsFromJSObject(p, ba.ctx)
}

// BasisToolsFromJSObject returns a wrapped BasisTools JavaScript class.
func BasisToolsFromJSObject(p js.Value, ctx js.Value) *BasisTools {
	return &BasisTools{p: p, ctx: ctx}
}

// TODO: methods
