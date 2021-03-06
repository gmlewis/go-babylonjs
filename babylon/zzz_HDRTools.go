// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// HDRTools represents a babylon.js HDRTools.
// This groups tools to convert HDR texture to native colors array.
type HDRTools struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (h *HDRTools) JSObject() js.Value { return h.p }

// HDRTools returns a HDRTools JavaScript class.
func (ba *Babylon) HDRTools() *HDRTools {
	p := ba.ctx.Get("HDRTools")
	return HDRToolsFromJSObject(p, ba.ctx)
}

// HDRToolsFromJSObject returns a wrapped HDRTools JavaScript class.
func HDRToolsFromJSObject(p js.Value, ctx js.Value) *HDRTools {
	return &HDRTools{p: p, ctx: ctx}
}

// HDRToolsArrayToJSArray returns a JavaScript Array for the wrapped array.
func HDRToolsArrayToJSArray(array []*HDRTools) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// GetCubeMapTextureData calls the GetCubeMapTextureData method on the HDRTools object.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrtools#getcubemaptexturedata
func (h *HDRTools) GetCubeMapTextureData(buffer js.Value, size float64) js.Value {

	args := make([]interface{}, 0, 2+0)

	args = append(args, buffer)

	args = append(args, size)

	retVal := h.p.Call("GetCubeMapTextureData", args...)
	return retVal
}

// RGBE_ReadHeader calls the RGBE_ReadHeader method on the HDRTools object.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrtools#rgbe_readheader
func (h *HDRTools) RGBE_ReadHeader(uint8array js.Value) js.Value {

	args := make([]interface{}, 0, 1+0)

	args = append(args, uint8array)

	retVal := h.p.Call("RGBE_ReadHeader", args...)
	return retVal
}

// RGBE_ReadPixels calls the RGBE_ReadPixels method on the HDRTools object.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrtools#rgbe_readpixels
func (h *HDRTools) RGBE_ReadPixels(uint8array js.Value, hdrInfo js.Value) js.Value {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uint8array)

	args = append(args, hdrInfo)

	retVal := h.p.Call("RGBE_ReadPixels", args...)
	return retVal
}
