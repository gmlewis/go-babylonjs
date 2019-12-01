// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Style represents a babylon.js Style.
// Define a style used by control to automatically setup properties based on a template.
// Only support font related properties so far
type Style struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *Style) JSObject() js.Value { return s.p }

// Style returns a Style JavaScript class.
func (ba *Babylon) Style() *Style {
	p := ba.ctx.Get("Style")
	return StyleFromJSObject(p, ba.ctx)
}

// StyleFromJSObject returns a wrapped Style JavaScript class.
func StyleFromJSObject(p js.Value, ctx js.Value) *Style {
	return &Style{p: p, ctx: ctx}
}

// StyleArrayToJSArray returns a JavaScript Array for the wrapped array.
func StyleArrayToJSArray(array []*Style) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewStyle returns a new Style object.
//
// https://doc.babylonjs.com/api/classes/babylon.style
func (ba *Babylon) NewStyle(host *AdvancedDynamicTexture) *Style {

	args := make([]interface{}, 0, 1+0)

	args = append(args, host.JSObject())

	p := ba.ctx.Get("Style").New(args...)
	return StyleFromJSObject(p, ba.ctx)
}

// Dispose calls the Dispose method on the Style object.
//
// https://doc.babylonjs.com/api/classes/babylon.style#dispose
func (s *Style) Dispose() {

	s.p.Call("dispose")
}

/*

// FontFamily returns the FontFamily property of class Style.
//
// https://doc.babylonjs.com/api/classes/babylon.style#fontfamily
func (s *Style) FontFamily(fontFamily string) *Style {
	p := ba.ctx.Get("Style").New(fontFamily)
	return StyleFromJSObject(p, ba.ctx)
}

// SetFontFamily sets the FontFamily property of class Style.
//
// https://doc.babylonjs.com/api/classes/babylon.style#fontfamily
func (s *Style) SetFontFamily(fontFamily string) *Style {
	p := ba.ctx.Get("Style").New(fontFamily)
	return StyleFromJSObject(p, ba.ctx)
}

// FontSize returns the FontSize property of class Style.
//
// https://doc.babylonjs.com/api/classes/babylon.style#fontsize
func (s *Style) FontSize(fontSize string) *Style {
	p := ba.ctx.Get("Style").New(fontSize)
	return StyleFromJSObject(p, ba.ctx)
}

// SetFontSize sets the FontSize property of class Style.
//
// https://doc.babylonjs.com/api/classes/babylon.style#fontsize
func (s *Style) SetFontSize(fontSize string) *Style {
	p := ba.ctx.Get("Style").New(fontSize)
	return StyleFromJSObject(p, ba.ctx)
}

// FontStyle returns the FontStyle property of class Style.
//
// https://doc.babylonjs.com/api/classes/babylon.style#fontstyle
func (s *Style) FontStyle(fontStyle string) *Style {
	p := ba.ctx.Get("Style").New(fontStyle)
	return StyleFromJSObject(p, ba.ctx)
}

// SetFontStyle sets the FontStyle property of class Style.
//
// https://doc.babylonjs.com/api/classes/babylon.style#fontstyle
func (s *Style) SetFontStyle(fontStyle string) *Style {
	p := ba.ctx.Get("Style").New(fontStyle)
	return StyleFromJSObject(p, ba.ctx)
}

// FontWeight returns the FontWeight property of class Style.
//
// https://doc.babylonjs.com/api/classes/babylon.style#fontweight
func (s *Style) FontWeight(fontWeight string) *Style {
	p := ba.ctx.Get("Style").New(fontWeight)
	return StyleFromJSObject(p, ba.ctx)
}

// SetFontWeight sets the FontWeight property of class Style.
//
// https://doc.babylonjs.com/api/classes/babylon.style#fontweight
func (s *Style) SetFontWeight(fontWeight string) *Style {
	p := ba.ctx.Get("Style").New(fontWeight)
	return StyleFromJSObject(p, ba.ctx)
}

// OnChangedObservable returns the OnChangedObservable property of class Style.
//
// https://doc.babylonjs.com/api/classes/babylon.style#onchangedobservable
func (s *Style) OnChangedObservable(onChangedObservable *Observable) *Style {
	p := ba.ctx.Get("Style").New(onChangedObservable.JSObject())
	return StyleFromJSObject(p, ba.ctx)
}

// SetOnChangedObservable sets the OnChangedObservable property of class Style.
//
// https://doc.babylonjs.com/api/classes/babylon.style#onchangedobservable
func (s *Style) SetOnChangedObservable(onChangedObservable *Observable) *Style {
	p := ba.ctx.Get("Style").New(onChangedObservable.JSObject())
	return StyleFromJSObject(p, ba.ctx)
}

*/
