// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Matrix2D represents a babylon.js Matrix2D.
// Class used to provide 2D matrix features
type Matrix2D struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (m *Matrix2D) JSObject() js.Value { return m.p }

// Matrix2D returns a Matrix2D JavaScript class.
func (ba *Babylon) Matrix2D() *Matrix2D {
	p := ba.ctx.Get("Matrix2D")
	return Matrix2DFromJSObject(p, ba.ctx)
}

// Matrix2DFromJSObject returns a wrapped Matrix2D JavaScript class.
func Matrix2DFromJSObject(p js.Value, ctx js.Value) *Matrix2D {
	return &Matrix2D{p: p, ctx: ctx}
}

// Matrix2DArrayToJSArray returns a JavaScript Array for the wrapped array.
func Matrix2DArrayToJSArray(array []*Matrix2D) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewMatrix2D returns a new Matrix2D object.
//
// https://doc.babylonjs.com/api/classes/babylon.matrix2d
func (ba *Babylon) NewMatrix2D(m00 float64, m01 float64, m10 float64, m11 float64, m20 float64, m21 float64) *Matrix2D {

	args := make([]interface{}, 0, 6+0)

	args = append(args, m00)
	args = append(args, m01)
	args = append(args, m10)
	args = append(args, m11)
	args = append(args, m20)
	args = append(args, m21)

	p := ba.ctx.Get("Matrix2D").New(args...)
	return Matrix2DFromJSObject(p, ba.ctx)
}

// ComposeToRef calls the ComposeToRef method on the Matrix2D object.
//
// https://doc.babylonjs.com/api/classes/babylon.matrix2d#composetoref
func (m *Matrix2D) ComposeToRef(tx float64, ty float64, angle float64, scaleX float64, scaleY float64, parentMatrix *Matrix2D, result *Matrix2D) {

	args := make([]interface{}, 0, 7+0)

	args = append(args, tx)
	args = append(args, ty)
	args = append(args, angle)
	args = append(args, scaleX)
	args = append(args, scaleY)
	args = append(args, parentMatrix.JSObject())
	args = append(args, result.JSObject())

	m.p.Call("ComposeToRef", args...)
}

// Determinant calls the Determinant method on the Matrix2D object.
//
// https://doc.babylonjs.com/api/classes/babylon.matrix2d#determinant
func (m *Matrix2D) Determinant() float64 {

	retVal := m.p.Call("determinant")
	return retVal.Float()
}

// FromValues calls the FromValues method on the Matrix2D object.
//
// https://doc.babylonjs.com/api/classes/babylon.matrix2d#fromvalues
func (m *Matrix2D) FromValues(m00 float64, m01 float64, m10 float64, m11 float64, m20 float64, m21 float64) *Matrix2D {

	args := make([]interface{}, 0, 6+0)

	args = append(args, m00)
	args = append(args, m01)
	args = append(args, m10)
	args = append(args, m11)
	args = append(args, m20)
	args = append(args, m21)

	retVal := m.p.Call("fromValues", args...)
	return Matrix2DFromJSObject(retVal, m.ctx)
}

// Identity calls the Identity method on the Matrix2D object.
//
// https://doc.babylonjs.com/api/classes/babylon.matrix2d#identity
func (m *Matrix2D) Identity() *Matrix2D {

	retVal := m.p.Call("Identity")
	return Matrix2DFromJSObject(retVal, m.ctx)
}

// InvertToRef calls the InvertToRef method on the Matrix2D object.
//
// https://doc.babylonjs.com/api/classes/babylon.matrix2d#inverttoref
func (m *Matrix2D) InvertToRef(result *Matrix2D) *Matrix2D {

	args := make([]interface{}, 0, 1+0)

	args = append(args, result.JSObject())

	retVal := m.p.Call("invertToRef", args...)
	return Matrix2DFromJSObject(retVal, m.ctx)
}

// MultiplyToRef calls the MultiplyToRef method on the Matrix2D object.
//
// https://doc.babylonjs.com/api/classes/babylon.matrix2d#multiplytoref
func (m *Matrix2D) MultiplyToRef(other *Matrix2D, result *Matrix2D) *Matrix2D {

	args := make([]interface{}, 0, 2+0)

	args = append(args, other.JSObject())
	args = append(args, result.JSObject())

	retVal := m.p.Call("multiplyToRef", args...)
	return Matrix2DFromJSObject(retVal, m.ctx)
}

// RotationToRef calls the RotationToRef method on the Matrix2D object.
//
// https://doc.babylonjs.com/api/classes/babylon.matrix2d#rotationtoref
func (m *Matrix2D) RotationToRef(angle float64, result *Matrix2D) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, angle)
	args = append(args, result.JSObject())

	m.p.Call("RotationToRef", args...)
}

// ScalingToRef calls the ScalingToRef method on the Matrix2D object.
//
// https://doc.babylonjs.com/api/classes/babylon.matrix2d#scalingtoref
func (m *Matrix2D) ScalingToRef(x float64, y float64, result *Matrix2D) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, x)
	args = append(args, y)
	args = append(args, result.JSObject())

	m.p.Call("ScalingToRef", args...)
}

// TransformCoordinates calls the TransformCoordinates method on the Matrix2D object.
//
// https://doc.babylonjs.com/api/classes/babylon.matrix2d#transformcoordinates
func (m *Matrix2D) TransformCoordinates(x float64, y float64, result *Vector2) *Matrix2D {

	args := make([]interface{}, 0, 3+0)

	args = append(args, x)
	args = append(args, y)
	args = append(args, result.JSObject())

	retVal := m.p.Call("transformCoordinates", args...)
	return Matrix2DFromJSObject(retVal, m.ctx)
}

// TranslationToRef calls the TranslationToRef method on the Matrix2D object.
//
// https://doc.babylonjs.com/api/classes/babylon.matrix2d#translationtoref
func (m *Matrix2D) TranslationToRef(x float64, y float64, result *Matrix2D) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, x)
	args = append(args, y)
	args = append(args, result.JSObject())

	m.p.Call("TranslationToRef", args...)
}

/*

// M returns the M property of class Matrix2D.
//
// https://doc.babylonjs.com/api/classes/babylon.matrix2d#m
func (m *Matrix2D) M(m js.Value) *Matrix2D {
	p := ba.ctx.Get("Matrix2D").New(m)
	return Matrix2DFromJSObject(p, ba.ctx)
}

// SetM sets the M property of class Matrix2D.
//
// https://doc.babylonjs.com/api/classes/babylon.matrix2d#m
func (m *Matrix2D) SetM(m js.Value) *Matrix2D {
	p := ba.ctx.Get("Matrix2D").New(m)
	return Matrix2DFromJSObject(p, ba.ctx)
}

*/
