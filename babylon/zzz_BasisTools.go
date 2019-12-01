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

// GetInternalFormatFromBasisFormat calls the GetInternalFormatFromBasisFormat method on the BasisTools object.
//
// https://doc.babylonjs.com/api/classes/babylon.basistools#getinternalformatfrombasisformat
func (b *BasisTools) GetInternalFormatFromBasisFormat(basisFormat float64) float64 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, basisFormat)

	retVal := b.p.Call("GetInternalFormatFromBasisFormat", args...)
	return retVal.Float()
}

// LoadTextureFromTranscodeResult calls the LoadTextureFromTranscodeResult method on the BasisTools object.
//
// https://doc.babylonjs.com/api/classes/babylon.basistools#loadtexturefromtranscoderesult
func (b *BasisTools) LoadTextureFromTranscodeResult(texture *InternalTexture, transcodeResult *TranscodeResult) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, texture.JSObject())
	args = append(args, transcodeResult.JSObject())

	b.p.Call("LoadTextureFromTranscodeResult", args...)
}

// TranscodeAsync calls the TranscodeAsync method on the BasisTools object.
//
// https://doc.babylonjs.com/api/classes/babylon.basistools#transcodeasync
func (b *BasisTools) TranscodeAsync(imageData js.Value, config *BasisTranscodeConfiguration) *TranscodeResult {

	args := make([]interface{}, 0, 2+0)

	args = append(args, imageData)
	args = append(args, config.JSObject())

	retVal := b.p.Call("TranscodeAsync", args...)
	return TranscodeResultFromJSObject(retVal, b.ctx)
}

/*

// JSModuleURL returns the JSModuleURL property of class BasisTools.
//
// https://doc.babylonjs.com/api/classes/babylon.basistools#jsmoduleurl
func (b *BasisTools) JSModuleURL(JSModuleURL string) *BasisTools {
	p := ba.ctx.Get("BasisTools").New(JSModuleURL)
	return BasisToolsFromJSObject(p, ba.ctx)
}

// SetJSModuleURL sets the JSModuleURL property of class BasisTools.
//
// https://doc.babylonjs.com/api/classes/babylon.basistools#jsmoduleurl
func (b *BasisTools) SetJSModuleURL(JSModuleURL string) *BasisTools {
	p := ba.ctx.Get("BasisTools").New(JSModuleURL)
	return BasisToolsFromJSObject(p, ba.ctx)
}

// WasmModuleURL returns the WasmModuleURL property of class BasisTools.
//
// https://doc.babylonjs.com/api/classes/babylon.basistools#wasmmoduleurl
func (b *BasisTools) WasmModuleURL(WasmModuleURL string) *BasisTools {
	p := ba.ctx.Get("BasisTools").New(WasmModuleURL)
	return BasisToolsFromJSObject(p, ba.ctx)
}

// SetWasmModuleURL sets the WasmModuleURL property of class BasisTools.
//
// https://doc.babylonjs.com/api/classes/babylon.basistools#wasmmoduleurl
func (b *BasisTools) SetWasmModuleURL(WasmModuleURL string) *BasisTools {
	p := ba.ctx.Get("BasisTools").New(WasmModuleURL)
	return BasisToolsFromJSObject(p, ba.ctx)
}

*/
