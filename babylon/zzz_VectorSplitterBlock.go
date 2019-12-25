// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// VectorSplitterBlock represents a babylon.js VectorSplitterBlock.
// Block used to expand a Vector3/4 into 4 outputs (one for each component)
type VectorSplitterBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (v *VectorSplitterBlock) JSObject() js.Value { return v.p }

// VectorSplitterBlock returns a VectorSplitterBlock JavaScript class.
func (ba *Babylon) VectorSplitterBlock() *VectorSplitterBlock {
	p := ba.ctx.Get("VectorSplitterBlock")
	return VectorSplitterBlockFromJSObject(p, ba.ctx)
}

// VectorSplitterBlockFromJSObject returns a wrapped VectorSplitterBlock JavaScript class.
func VectorSplitterBlockFromJSObject(p js.Value, ctx js.Value) *VectorSplitterBlock {
	return &VectorSplitterBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// VectorSplitterBlockArrayToJSArray returns a JavaScript Array for the wrapped array.
func VectorSplitterBlockArrayToJSArray(array []*VectorSplitterBlock) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewVectorSplitterBlock returns a new VectorSplitterBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.vectorsplitterblock#constructor
func (ba *Babylon) NewVectorSplitterBlock(name string) *VectorSplitterBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("VectorSplitterBlock").New(args...)
	return VectorSplitterBlockFromJSObject(p, ba.ctx)
}

// GetClassName calls the GetClassName method on the VectorSplitterBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.vectorsplitterblock#getclassname
func (v *VectorSplitterBlock) GetClassName() string {

	retVal := v.p.Call("getClassName")
	return retVal.String()
}

// W returns the W property of class VectorSplitterBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.vectorsplitterblock#w
func (v *VectorSplitterBlock) W() *NodeMaterialConnectionPoint {
	retVal := v.p.Get("w")
	return NodeMaterialConnectionPointFromJSObject(retVal, v.ctx)
}

// SetW sets the W property of class VectorSplitterBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.vectorsplitterblock#w
func (v *VectorSplitterBlock) SetW(w *NodeMaterialConnectionPoint) *VectorSplitterBlock {
	v.p.Set("w", w.JSObject())
	return v
}

// X returns the X property of class VectorSplitterBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.vectorsplitterblock#x
func (v *VectorSplitterBlock) X() *NodeMaterialConnectionPoint {
	retVal := v.p.Get("x")
	return NodeMaterialConnectionPointFromJSObject(retVal, v.ctx)
}

// SetX sets the X property of class VectorSplitterBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.vectorsplitterblock#x
func (v *VectorSplitterBlock) SetX(x *NodeMaterialConnectionPoint) *VectorSplitterBlock {
	v.p.Set("x", x.JSObject())
	return v
}

// XyIn returns the XyIn property of class VectorSplitterBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.vectorsplitterblock#xyin
func (v *VectorSplitterBlock) XyIn() *NodeMaterialConnectionPoint {
	retVal := v.p.Get("xyIn")
	return NodeMaterialConnectionPointFromJSObject(retVal, v.ctx)
}

// SetXyIn sets the XyIn property of class VectorSplitterBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.vectorsplitterblock#xyin
func (v *VectorSplitterBlock) SetXyIn(xyIn *NodeMaterialConnectionPoint) *VectorSplitterBlock {
	v.p.Set("xyIn", xyIn.JSObject())
	return v
}

// XyOut returns the XyOut property of class VectorSplitterBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.vectorsplitterblock#xyout
func (v *VectorSplitterBlock) XyOut() *NodeMaterialConnectionPoint {
	retVal := v.p.Get("xyOut")
	return NodeMaterialConnectionPointFromJSObject(retVal, v.ctx)
}

// SetXyOut sets the XyOut property of class VectorSplitterBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.vectorsplitterblock#xyout
func (v *VectorSplitterBlock) SetXyOut(xyOut *NodeMaterialConnectionPoint) *VectorSplitterBlock {
	v.p.Set("xyOut", xyOut.JSObject())
	return v
}

// XyzIn returns the XyzIn property of class VectorSplitterBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.vectorsplitterblock#xyzin
func (v *VectorSplitterBlock) XyzIn() *NodeMaterialConnectionPoint {
	retVal := v.p.Get("xyzIn")
	return NodeMaterialConnectionPointFromJSObject(retVal, v.ctx)
}

// SetXyzIn sets the XyzIn property of class VectorSplitterBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.vectorsplitterblock#xyzin
func (v *VectorSplitterBlock) SetXyzIn(xyzIn *NodeMaterialConnectionPoint) *VectorSplitterBlock {
	v.p.Set("xyzIn", xyzIn.JSObject())
	return v
}

// XyzOut returns the XyzOut property of class VectorSplitterBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.vectorsplitterblock#xyzout
func (v *VectorSplitterBlock) XyzOut() *NodeMaterialConnectionPoint {
	retVal := v.p.Get("xyzOut")
	return NodeMaterialConnectionPointFromJSObject(retVal, v.ctx)
}

// SetXyzOut sets the XyzOut property of class VectorSplitterBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.vectorsplitterblock#xyzout
func (v *VectorSplitterBlock) SetXyzOut(xyzOut *NodeMaterialConnectionPoint) *VectorSplitterBlock {
	v.p.Set("xyzOut", xyzOut.JSObject())
	return v
}

// Xyzw returns the Xyzw property of class VectorSplitterBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.vectorsplitterblock#xyzw
func (v *VectorSplitterBlock) Xyzw() *NodeMaterialConnectionPoint {
	retVal := v.p.Get("xyzw")
	return NodeMaterialConnectionPointFromJSObject(retVal, v.ctx)
}

// SetXyzw sets the Xyzw property of class VectorSplitterBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.vectorsplitterblock#xyzw
func (v *VectorSplitterBlock) SetXyzw(xyzw *NodeMaterialConnectionPoint) *VectorSplitterBlock {
	v.p.Set("xyzw", xyzw.JSObject())
	return v
}

// Y returns the Y property of class VectorSplitterBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.vectorsplitterblock#y
func (v *VectorSplitterBlock) Y() *NodeMaterialConnectionPoint {
	retVal := v.p.Get("y")
	return NodeMaterialConnectionPointFromJSObject(retVal, v.ctx)
}

// SetY sets the Y property of class VectorSplitterBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.vectorsplitterblock#y
func (v *VectorSplitterBlock) SetY(y *NodeMaterialConnectionPoint) *VectorSplitterBlock {
	v.p.Set("y", y.JSObject())
	return v
}

// Z returns the Z property of class VectorSplitterBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.vectorsplitterblock#z
func (v *VectorSplitterBlock) Z() *NodeMaterialConnectionPoint {
	retVal := v.p.Get("z")
	return NodeMaterialConnectionPointFromJSObject(retVal, v.ctx)
}

// SetZ sets the Z property of class VectorSplitterBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.vectorsplitterblock#z
func (v *VectorSplitterBlock) SetZ(z *NodeMaterialConnectionPoint) *VectorSplitterBlock {
	v.p.Set("z", z.JSObject())
	return v
}
