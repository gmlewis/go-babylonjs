// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AdvancedDynamicTextureInstrumentation represents a babylon.js AdvancedDynamicTextureInstrumentation.
// This class can be used to get instrumentation data from a AdvancedDynamicTexture object
type AdvancedDynamicTextureInstrumentation struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *AdvancedDynamicTextureInstrumentation) JSObject() js.Value { return a.p }

// AdvancedDynamicTextureInstrumentation returns a AdvancedDynamicTextureInstrumentation JavaScript class.
func (ba *Babylon) AdvancedDynamicTextureInstrumentation() *AdvancedDynamicTextureInstrumentation {
	p := ba.ctx.Get("AdvancedDynamicTextureInstrumentation")
	return AdvancedDynamicTextureInstrumentationFromJSObject(p, ba.ctx)
}

// AdvancedDynamicTextureInstrumentationFromJSObject returns a wrapped AdvancedDynamicTextureInstrumentation JavaScript class.
func AdvancedDynamicTextureInstrumentationFromJSObject(p js.Value, ctx js.Value) *AdvancedDynamicTextureInstrumentation {
	return &AdvancedDynamicTextureInstrumentation{p: p, ctx: ctx}
}

// AdvancedDynamicTextureInstrumentationArrayToJSArray returns a JavaScript Array for the wrapped array.
func AdvancedDynamicTextureInstrumentationArrayToJSArray(array []*AdvancedDynamicTextureInstrumentation) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewAdvancedDynamicTextureInstrumentation returns a new AdvancedDynamicTextureInstrumentation object.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictextureinstrumentation
func (ba *Babylon) NewAdvancedDynamicTextureInstrumentation(texture *AdvancedDynamicTexture) *AdvancedDynamicTextureInstrumentation {

	args := make([]interface{}, 0, 1+0)

	args = append(args, texture.JSObject())

	p := ba.ctx.Get("AdvancedDynamicTextureInstrumentation").New(args...)
	return AdvancedDynamicTextureInstrumentationFromJSObject(p, ba.ctx)
}

// Dispose calls the Dispose method on the AdvancedDynamicTextureInstrumentation object.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictextureinstrumentation#dispose
func (a *AdvancedDynamicTextureInstrumentation) Dispose() {

	a.p.Call("dispose")
}

/*

// CaptureLayoutTime returns the CaptureLayoutTime property of class AdvancedDynamicTextureInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictextureinstrumentation#capturelayouttime
func (a *AdvancedDynamicTextureInstrumentation) CaptureLayoutTime(captureLayoutTime bool) *AdvancedDynamicTextureInstrumentation {
	p := ba.ctx.Get("AdvancedDynamicTextureInstrumentation").New(captureLayoutTime)
	return AdvancedDynamicTextureInstrumentationFromJSObject(p, ba.ctx)
}

// SetCaptureLayoutTime sets the CaptureLayoutTime property of class AdvancedDynamicTextureInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictextureinstrumentation#capturelayouttime
func (a *AdvancedDynamicTextureInstrumentation) SetCaptureLayoutTime(captureLayoutTime bool) *AdvancedDynamicTextureInstrumentation {
	p := ba.ctx.Get("AdvancedDynamicTextureInstrumentation").New(captureLayoutTime)
	return AdvancedDynamicTextureInstrumentationFromJSObject(p, ba.ctx)
}

// CaptureRenderTime returns the CaptureRenderTime property of class AdvancedDynamicTextureInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictextureinstrumentation#capturerendertime
func (a *AdvancedDynamicTextureInstrumentation) CaptureRenderTime(captureRenderTime bool) *AdvancedDynamicTextureInstrumentation {
	p := ba.ctx.Get("AdvancedDynamicTextureInstrumentation").New(captureRenderTime)
	return AdvancedDynamicTextureInstrumentationFromJSObject(p, ba.ctx)
}

// SetCaptureRenderTime sets the CaptureRenderTime property of class AdvancedDynamicTextureInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictextureinstrumentation#capturerendertime
func (a *AdvancedDynamicTextureInstrumentation) SetCaptureRenderTime(captureRenderTime bool) *AdvancedDynamicTextureInstrumentation {
	p := ba.ctx.Get("AdvancedDynamicTextureInstrumentation").New(captureRenderTime)
	return AdvancedDynamicTextureInstrumentationFromJSObject(p, ba.ctx)
}

// LayoutTimeCounter returns the LayoutTimeCounter property of class AdvancedDynamicTextureInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictextureinstrumentation#layouttimecounter
func (a *AdvancedDynamicTextureInstrumentation) LayoutTimeCounter(layoutTimeCounter *PerfCounter) *AdvancedDynamicTextureInstrumentation {
	p := ba.ctx.Get("AdvancedDynamicTextureInstrumentation").New(layoutTimeCounter.JSObject())
	return AdvancedDynamicTextureInstrumentationFromJSObject(p, ba.ctx)
}

// SetLayoutTimeCounter sets the LayoutTimeCounter property of class AdvancedDynamicTextureInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictextureinstrumentation#layouttimecounter
func (a *AdvancedDynamicTextureInstrumentation) SetLayoutTimeCounter(layoutTimeCounter *PerfCounter) *AdvancedDynamicTextureInstrumentation {
	p := ba.ctx.Get("AdvancedDynamicTextureInstrumentation").New(layoutTimeCounter.JSObject())
	return AdvancedDynamicTextureInstrumentationFromJSObject(p, ba.ctx)
}

// RenderTimeCounter returns the RenderTimeCounter property of class AdvancedDynamicTextureInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictextureinstrumentation#rendertimecounter
func (a *AdvancedDynamicTextureInstrumentation) RenderTimeCounter(renderTimeCounter *PerfCounter) *AdvancedDynamicTextureInstrumentation {
	p := ba.ctx.Get("AdvancedDynamicTextureInstrumentation").New(renderTimeCounter.JSObject())
	return AdvancedDynamicTextureInstrumentationFromJSObject(p, ba.ctx)
}

// SetRenderTimeCounter sets the RenderTimeCounter property of class AdvancedDynamicTextureInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictextureinstrumentation#rendertimecounter
func (a *AdvancedDynamicTextureInstrumentation) SetRenderTimeCounter(renderTimeCounter *PerfCounter) *AdvancedDynamicTextureInstrumentation {
	p := ba.ctx.Get("AdvancedDynamicTextureInstrumentation").New(renderTimeCounter.JSObject())
	return AdvancedDynamicTextureInstrumentationFromJSObject(p, ba.ctx)
}

// Texture returns the Texture property of class AdvancedDynamicTextureInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictextureinstrumentation#texture
func (a *AdvancedDynamicTextureInstrumentation) Texture(texture *AdvancedDynamicTexture) *AdvancedDynamicTextureInstrumentation {
	p := ba.ctx.Get("AdvancedDynamicTextureInstrumentation").New(texture.JSObject())
	return AdvancedDynamicTextureInstrumentationFromJSObject(p, ba.ctx)
}

// SetTexture sets the Texture property of class AdvancedDynamicTextureInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.advanceddynamictextureinstrumentation#texture
func (a *AdvancedDynamicTextureInstrumentation) SetTexture(texture *AdvancedDynamicTexture) *AdvancedDynamicTextureInstrumentation {
	p := ba.ctx.Get("AdvancedDynamicTextureInstrumentation").New(texture.JSObject())
	return AdvancedDynamicTextureInstrumentationFromJSObject(p, ba.ctx)
}

*/
