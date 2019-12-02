// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ColorMergerBlock represents a babylon.js ColorMergerBlock.
// Block used to create a Color3/4 out of individual inputs (one for each component)
type ColorMergerBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (c *ColorMergerBlock) JSObject() js.Value { return c.p }

// ColorMergerBlock returns a ColorMergerBlock JavaScript class.
func (ba *Babylon) ColorMergerBlock() *ColorMergerBlock {
	p := ba.ctx.Get("ColorMergerBlock")
	return ColorMergerBlockFromJSObject(p, ba.ctx)
}

// ColorMergerBlockFromJSObject returns a wrapped ColorMergerBlock JavaScript class.
func ColorMergerBlockFromJSObject(p js.Value, ctx js.Value) *ColorMergerBlock {
	return &ColorMergerBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// ColorMergerBlockArrayToJSArray returns a JavaScript Array for the wrapped array.
func ColorMergerBlockArrayToJSArray(array []*ColorMergerBlock) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewColorMergerBlock returns a new ColorMergerBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock
func (ba *Babylon) NewColorMergerBlock(name string) *ColorMergerBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("ColorMergerBlock").New(args...)
	return ColorMergerBlockFromJSObject(p, ba.ctx)
}

// GetClassName calls the GetClassName method on the ColorMergerBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#getclassname
func (c *ColorMergerBlock) GetClassName() string {

	retVal := c.p.Call("getClassName")
	return retVal.String()
}

/*

// A returns the A property of class ColorMergerBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#a
func (c *ColorMergerBlock) A(a *NodeMaterialConnectionPoint) *ColorMergerBlock {
	p := ba.ctx.Get("ColorMergerBlock").New(a.JSObject())
	return ColorMergerBlockFromJSObject(p, ba.ctx)
}

// SetA sets the A property of class ColorMergerBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#a
func (c *ColorMergerBlock) SetA(a *NodeMaterialConnectionPoint) *ColorMergerBlock {
	p := ba.ctx.Get("ColorMergerBlock").New(a.JSObject())
	return ColorMergerBlockFromJSObject(p, ba.ctx)
}

// B returns the B property of class ColorMergerBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#b
func (c *ColorMergerBlock) B(b *NodeMaterialConnectionPoint) *ColorMergerBlock {
	p := ba.ctx.Get("ColorMergerBlock").New(b.JSObject())
	return ColorMergerBlockFromJSObject(p, ba.ctx)
}

// SetB sets the B property of class ColorMergerBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#b
func (c *ColorMergerBlock) SetB(b *NodeMaterialConnectionPoint) *ColorMergerBlock {
	p := ba.ctx.Get("ColorMergerBlock").New(b.JSObject())
	return ColorMergerBlockFromJSObject(p, ba.ctx)
}

// G returns the G property of class ColorMergerBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#g
func (c *ColorMergerBlock) G(g *NodeMaterialConnectionPoint) *ColorMergerBlock {
	p := ba.ctx.Get("ColorMergerBlock").New(g.JSObject())
	return ColorMergerBlockFromJSObject(p, ba.ctx)
}

// SetG sets the G property of class ColorMergerBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#g
func (c *ColorMergerBlock) SetG(g *NodeMaterialConnectionPoint) *ColorMergerBlock {
	p := ba.ctx.Get("ColorMergerBlock").New(g.JSObject())
	return ColorMergerBlockFromJSObject(p, ba.ctx)
}

// R returns the R property of class ColorMergerBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#r
func (c *ColorMergerBlock) R(r *NodeMaterialConnectionPoint) *ColorMergerBlock {
	p := ba.ctx.Get("ColorMergerBlock").New(r.JSObject())
	return ColorMergerBlockFromJSObject(p, ba.ctx)
}

// SetR sets the R property of class ColorMergerBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#r
func (c *ColorMergerBlock) SetR(r *NodeMaterialConnectionPoint) *ColorMergerBlock {
	p := ba.ctx.Get("ColorMergerBlock").New(r.JSObject())
	return ColorMergerBlockFromJSObject(p, ba.ctx)
}

// Rgb returns the Rgb property of class ColorMergerBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#rgb
func (c *ColorMergerBlock) Rgb(rgb *NodeMaterialConnectionPoint) *ColorMergerBlock {
	p := ba.ctx.Get("ColorMergerBlock").New(rgb.JSObject())
	return ColorMergerBlockFromJSObject(p, ba.ctx)
}

// SetRgb sets the Rgb property of class ColorMergerBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#rgb
func (c *ColorMergerBlock) SetRgb(rgb *NodeMaterialConnectionPoint) *ColorMergerBlock {
	p := ba.ctx.Get("ColorMergerBlock").New(rgb.JSObject())
	return ColorMergerBlockFromJSObject(p, ba.ctx)
}

// RgbIn returns the RgbIn property of class ColorMergerBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#rgbin
func (c *ColorMergerBlock) RgbIn(rgbIn *NodeMaterialConnectionPoint) *ColorMergerBlock {
	p := ba.ctx.Get("ColorMergerBlock").New(rgbIn.JSObject())
	return ColorMergerBlockFromJSObject(p, ba.ctx)
}

// SetRgbIn sets the RgbIn property of class ColorMergerBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#rgbin
func (c *ColorMergerBlock) SetRgbIn(rgbIn *NodeMaterialConnectionPoint) *ColorMergerBlock {
	p := ba.ctx.Get("ColorMergerBlock").New(rgbIn.JSObject())
	return ColorMergerBlockFromJSObject(p, ba.ctx)
}

// RgbOut returns the RgbOut property of class ColorMergerBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#rgbout
func (c *ColorMergerBlock) RgbOut(rgbOut *NodeMaterialConnectionPoint) *ColorMergerBlock {
	p := ba.ctx.Get("ColorMergerBlock").New(rgbOut.JSObject())
	return ColorMergerBlockFromJSObject(p, ba.ctx)
}

// SetRgbOut sets the RgbOut property of class ColorMergerBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#rgbout
func (c *ColorMergerBlock) SetRgbOut(rgbOut *NodeMaterialConnectionPoint) *ColorMergerBlock {
	p := ba.ctx.Get("ColorMergerBlock").New(rgbOut.JSObject())
	return ColorMergerBlockFromJSObject(p, ba.ctx)
}

// Rgba returns the Rgba property of class ColorMergerBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#rgba
func (c *ColorMergerBlock) Rgba(rgba *NodeMaterialConnectionPoint) *ColorMergerBlock {
	p := ba.ctx.Get("ColorMergerBlock").New(rgba.JSObject())
	return ColorMergerBlockFromJSObject(p, ba.ctx)
}

// SetRgba sets the Rgba property of class ColorMergerBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#rgba
func (c *ColorMergerBlock) SetRgba(rgba *NodeMaterialConnectionPoint) *ColorMergerBlock {
	p := ba.ctx.Get("ColorMergerBlock").New(rgba.JSObject())
	return ColorMergerBlockFromJSObject(p, ba.ctx)
}

*/
