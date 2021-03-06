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
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#constructor
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

// A returns the A property of class ColorMergerBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#a
func (c *ColorMergerBlock) A() *NodeMaterialConnectionPoint {
	retVal := c.p.Get("a")
	return NodeMaterialConnectionPointFromJSObject(retVal, c.ctx)
}

// SetA sets the A property of class ColorMergerBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#a
func (c *ColorMergerBlock) SetA(a *NodeMaterialConnectionPoint) *ColorMergerBlock {
	c.p.Set("a", a.JSObject())
	return c
}

// B returns the B property of class ColorMergerBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#b
func (c *ColorMergerBlock) B() *NodeMaterialConnectionPoint {
	retVal := c.p.Get("b")
	return NodeMaterialConnectionPointFromJSObject(retVal, c.ctx)
}

// SetB sets the B property of class ColorMergerBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#b
func (c *ColorMergerBlock) SetB(b *NodeMaterialConnectionPoint) *ColorMergerBlock {
	c.p.Set("b", b.JSObject())
	return c
}

// G returns the G property of class ColorMergerBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#g
func (c *ColorMergerBlock) G() *NodeMaterialConnectionPoint {
	retVal := c.p.Get("g")
	return NodeMaterialConnectionPointFromJSObject(retVal, c.ctx)
}

// SetG sets the G property of class ColorMergerBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#g
func (c *ColorMergerBlock) SetG(g *NodeMaterialConnectionPoint) *ColorMergerBlock {
	c.p.Set("g", g.JSObject())
	return c
}

// R returns the R property of class ColorMergerBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#r
func (c *ColorMergerBlock) R() *NodeMaterialConnectionPoint {
	retVal := c.p.Get("r")
	return NodeMaterialConnectionPointFromJSObject(retVal, c.ctx)
}

// SetR sets the R property of class ColorMergerBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#r
func (c *ColorMergerBlock) SetR(r *NodeMaterialConnectionPoint) *ColorMergerBlock {
	c.p.Set("r", r.JSObject())
	return c
}

// Rgb returns the Rgb property of class ColorMergerBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#rgb
func (c *ColorMergerBlock) Rgb() *NodeMaterialConnectionPoint {
	retVal := c.p.Get("rgb")
	return NodeMaterialConnectionPointFromJSObject(retVal, c.ctx)
}

// SetRgb sets the Rgb property of class ColorMergerBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#rgb
func (c *ColorMergerBlock) SetRgb(rgb *NodeMaterialConnectionPoint) *ColorMergerBlock {
	c.p.Set("rgb", rgb.JSObject())
	return c
}

// RgbIn returns the RgbIn property of class ColorMergerBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#rgbin
func (c *ColorMergerBlock) RgbIn() *NodeMaterialConnectionPoint {
	retVal := c.p.Get("rgbIn")
	return NodeMaterialConnectionPointFromJSObject(retVal, c.ctx)
}

// SetRgbIn sets the RgbIn property of class ColorMergerBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#rgbin
func (c *ColorMergerBlock) SetRgbIn(rgbIn *NodeMaterialConnectionPoint) *ColorMergerBlock {
	c.p.Set("rgbIn", rgbIn.JSObject())
	return c
}

// RgbOut returns the RgbOut property of class ColorMergerBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#rgbout
func (c *ColorMergerBlock) RgbOut() *NodeMaterialConnectionPoint {
	retVal := c.p.Get("rgbOut")
	return NodeMaterialConnectionPointFromJSObject(retVal, c.ctx)
}

// SetRgbOut sets the RgbOut property of class ColorMergerBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#rgbout
func (c *ColorMergerBlock) SetRgbOut(rgbOut *NodeMaterialConnectionPoint) *ColorMergerBlock {
	c.p.Set("rgbOut", rgbOut.JSObject())
	return c
}

// Rgba returns the Rgba property of class ColorMergerBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#rgba
func (c *ColorMergerBlock) Rgba() *NodeMaterialConnectionPoint {
	retVal := c.p.Get("rgba")
	return NodeMaterialConnectionPointFromJSObject(retVal, c.ctx)
}

// SetRgba sets the Rgba property of class ColorMergerBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.colormergerblock#rgba
func (c *ColorMergerBlock) SetRgba(rgba *NodeMaterialConnectionPoint) *ColorMergerBlock {
	c.p.Set("rgba", rgba.JSObject())
	return c
}
