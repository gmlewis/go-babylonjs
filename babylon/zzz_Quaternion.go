// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Quaternion represents a babylon.js Quaternion.
// Class used to store quaternion data
//
// See: http://doc.babylonjs.com/features/position,_rotation,_scaling
type Quaternion struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (q *Quaternion) JSObject() js.Value { return q.p }

// Quaternion returns a Quaternion JavaScript class.
func (ba *Babylon) Quaternion() *Quaternion {
	p := ba.ctx.Get("Quaternion")
	return QuaternionFromJSObject(p, ba.ctx)
}

// QuaternionFromJSObject returns a wrapped Quaternion JavaScript class.
func QuaternionFromJSObject(p js.Value, ctx js.Value) *Quaternion {
	return &Quaternion{p: p, ctx: ctx}
}

// NewQuaternionOpts contains optional parameters for NewQuaternion.
type NewQuaternionOpts struct {
	X *float64
	Y *float64
	Z *float64
	W *float64
}

// NewQuaternion returns a new Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion
func (ba *Babylon) NewQuaternion(opts *NewQuaternionOpts) *Quaternion {
	if opts == nil {
		opts = &NewQuaternionOpts{}
	}

	args := make([]interface{}, 0, 0+4)

	if opts.X == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.X)
	}
	if opts.Y == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Y)
	}
	if opts.Z == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Z)
	}
	if opts.W == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.W)
	}

	p := ba.ctx.Get("Quaternion").New(args...)
	return QuaternionFromJSObject(p, ba.ctx)
}

// Add calls the Add method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#add
func (q *Quaternion) Add(other *Quaternion) *Quaternion {

	args := make([]interface{}, 0, 1+0)

	args = append(args, other.JSObject())

	retVal := q.p.Call("add", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// AddInPlace calls the AddInPlace method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#addinplace
func (q *Quaternion) AddInPlace(other *Quaternion) *Quaternion {

	args := make([]interface{}, 0, 1+0)

	args = append(args, other.JSObject())

	retVal := q.p.Call("addInPlace", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// AreClose calls the AreClose method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#areclose
func (q *Quaternion) AreClose(quat0 *Quaternion, quat1 *Quaternion) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, quat0.JSObject())
	args = append(args, quat1.JSObject())

	retVal := q.p.Call("AreClose", args...)
	return retVal.Bool()
}

// AsArray calls the AsArray method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#asarray
func (q *Quaternion) AsArray() float64 {

	args := make([]interface{}, 0, 0+0)

	retVal := q.p.Call("asArray", args...)
	return retVal.Float()
}

// Clone calls the Clone method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#clone
func (q *Quaternion) Clone() *Quaternion {

	args := make([]interface{}, 0, 0+0)

	retVal := q.p.Call("clone", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// Conjugate calls the Conjugate method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#conjugate
func (q *Quaternion) Conjugate() *Quaternion {

	args := make([]interface{}, 0, 0+0)

	retVal := q.p.Call("conjugate", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// ConjugateInPlace calls the ConjugateInPlace method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#conjugateinplace
func (q *Quaternion) ConjugateInPlace() *Quaternion {

	args := make([]interface{}, 0, 0+0)

	retVal := q.p.Call("conjugateInPlace", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// ConjugateToRef calls the ConjugateToRef method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#conjugatetoref
func (q *Quaternion) ConjugateToRef(ref *Quaternion) *Quaternion {

	args := make([]interface{}, 0, 1+0)

	args = append(args, ref.JSObject())

	retVal := q.p.Call("conjugateToRef", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// CopyFrom calls the CopyFrom method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#copyfrom
func (q *Quaternion) CopyFrom(other *Quaternion) *Quaternion {

	args := make([]interface{}, 0, 1+0)

	args = append(args, other.JSObject())

	retVal := q.p.Call("copyFrom", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// CopyFromFloats calls the CopyFromFloats method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#copyfromfloats
func (q *Quaternion) CopyFromFloats(x float64, y float64, z float64, w float64) *Quaternion {

	args := make([]interface{}, 0, 4+0)

	args = append(args, x)
	args = append(args, y)
	args = append(args, z)
	args = append(args, w)

	retVal := q.p.Call("copyFromFloats", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// Dot calls the Dot method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#dot
func (q *Quaternion) Dot(left *Quaternion, right *Quaternion) float64 {

	args := make([]interface{}, 0, 2+0)

	args = append(args, left.JSObject())
	args = append(args, right.JSObject())

	retVal := q.p.Call("Dot", args...)
	return retVal.Float()
}

// Equals calls the Equals method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#equals
func (q *Quaternion) Equals(otherQuaternion *Quaternion) bool {

	args := make([]interface{}, 0, 1+0)

	args = append(args, otherQuaternion.JSObject())

	retVal := q.p.Call("equals", args...)
	return retVal.Bool()
}

// QuaternionFromArrayOpts contains optional parameters for Quaternion.FromArray.
type QuaternionFromArrayOpts struct {
	Offset *float64
}

// FromArray calls the FromArray method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#fromarray
func (q *Quaternion) FromArray(array js.Value, opts *QuaternionFromArrayOpts) *Quaternion {
	if opts == nil {
		opts = &QuaternionFromArrayOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, array)

	if opts.Offset == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Offset)
	}

	retVal := q.p.Call("FromArray", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// FromEulerAngles calls the FromEulerAngles method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#fromeulerangles
func (q *Quaternion) FromEulerAngles(x float64, y float64, z float64) *Quaternion {

	args := make([]interface{}, 0, 3+0)

	args = append(args, x)
	args = append(args, y)
	args = append(args, z)

	retVal := q.p.Call("FromEulerAngles", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// FromEulerAnglesToRef calls the FromEulerAnglesToRef method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#fromeuleranglestoref
func (q *Quaternion) FromEulerAnglesToRef(x float64, y float64, z float64, result *Quaternion) *Quaternion {

	args := make([]interface{}, 0, 4+0)

	args = append(args, x)
	args = append(args, y)
	args = append(args, z)
	args = append(args, result.JSObject())

	retVal := q.p.Call("FromEulerAnglesToRef", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// FromEulerVector calls the FromEulerVector method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#fromeulervector
func (q *Quaternion) FromEulerVector(vec *Vector3) *Quaternion {

	args := make([]interface{}, 0, 1+0)

	args = append(args, vec.JSObject())

	retVal := q.p.Call("FromEulerVector", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// FromEulerVectorToRef calls the FromEulerVectorToRef method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#fromeulervectortoref
func (q *Quaternion) FromEulerVectorToRef(vec *Vector3, result *Quaternion) *Quaternion {

	args := make([]interface{}, 0, 2+0)

	args = append(args, vec.JSObject())
	args = append(args, result.JSObject())

	retVal := q.p.Call("FromEulerVectorToRef", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// FromRotationMatrix calls the FromRotationMatrix method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#fromrotationmatrix
func (q *Quaternion) FromRotationMatrix(matrix *Matrix) *Quaternion {

	args := make([]interface{}, 0, 1+0)

	args = append(args, matrix.JSObject())

	retVal := q.p.Call("FromRotationMatrix", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// FromRotationMatrixToRef calls the FromRotationMatrixToRef method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#fromrotationmatrixtoref
func (q *Quaternion) FromRotationMatrixToRef(matrix *Matrix, result *Quaternion) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, matrix.JSObject())
	args = append(args, result.JSObject())

	q.p.Call("FromRotationMatrixToRef", args...)
}

// GetClassName calls the GetClassName method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#getclassname
func (q *Quaternion) GetClassName() string {

	args := make([]interface{}, 0, 0+0)

	retVal := q.p.Call("getClassName", args...)
	return retVal.String()
}

// GetHashCode calls the GetHashCode method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#gethashcode
func (q *Quaternion) GetHashCode() float64 {

	args := make([]interface{}, 0, 0+0)

	retVal := q.p.Call("getHashCode", args...)
	return retVal.Float()
}

// Hermite calls the Hermite method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#hermite
func (q *Quaternion) Hermite(value1 *Quaternion, tangent1 *Quaternion, value2 *Quaternion, tangent2 *Quaternion, amount float64) *Quaternion {

	args := make([]interface{}, 0, 5+0)

	args = append(args, value1.JSObject())
	args = append(args, tangent1.JSObject())
	args = append(args, value2.JSObject())
	args = append(args, tangent2.JSObject())
	args = append(args, amount)

	retVal := q.p.Call("Hermite", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// Identity calls the Identity method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#identity
func (q *Quaternion) Identity() *Quaternion {

	args := make([]interface{}, 0, 0+0)

	retVal := q.p.Call("Identity", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// Inverse calls the Inverse method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#inverse
func (q *Quaternion) Inverse(q *Quaternion) *Quaternion {

	args := make([]interface{}, 0, 1+0)

	args = append(args, q.JSObject())

	retVal := q.p.Call("Inverse", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// InverseToRef calls the InverseToRef method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#inversetoref
func (q *Quaternion) InverseToRef(q *Quaternion, result *Quaternion) *Quaternion {

	args := make([]interface{}, 0, 2+0)

	args = append(args, q.JSObject())
	args = append(args, result.JSObject())

	retVal := q.p.Call("InverseToRef", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// IsIdentity calls the IsIdentity method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#isidentity
func (q *Quaternion) IsIdentity(quaternion *Quaternion) bool {

	args := make([]interface{}, 0, 1+0)

	args = append(args, quaternion.JSObject())

	retVal := q.p.Call("IsIdentity", args...)
	return retVal.Bool()
}

// Length calls the Length method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#length
func (q *Quaternion) Length() float64 {

	args := make([]interface{}, 0, 0+0)

	retVal := q.p.Call("length", args...)
	return retVal.Float()
}

// Multiply calls the Multiply method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#multiply
func (q *Quaternion) Multiply(q1 *Quaternion) *Quaternion {

	args := make([]interface{}, 0, 1+0)

	args = append(args, q1.JSObject())

	retVal := q.p.Call("multiply", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// MultiplyInPlace calls the MultiplyInPlace method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#multiplyinplace
func (q *Quaternion) MultiplyInPlace(q1 *Quaternion) *Quaternion {

	args := make([]interface{}, 0, 1+0)

	args = append(args, q1.JSObject())

	retVal := q.p.Call("multiplyInPlace", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// MultiplyToRef calls the MultiplyToRef method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#multiplytoref
func (q *Quaternion) MultiplyToRef(q1 *Quaternion, result *Quaternion) *Quaternion {

	args := make([]interface{}, 0, 2+0)

	args = append(args, q1.JSObject())
	args = append(args, result.JSObject())

	retVal := q.p.Call("multiplyToRef", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// Normalize calls the Normalize method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#normalize
func (q *Quaternion) Normalize() *Quaternion {

	args := make([]interface{}, 0, 0+0)

	retVal := q.p.Call("normalize", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// RotationAlphaBetaGamma calls the RotationAlphaBetaGamma method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#rotationalphabetagamma
func (q *Quaternion) RotationAlphaBetaGamma(alpha float64, beta float64, gamma float64) *Quaternion {

	args := make([]interface{}, 0, 3+0)

	args = append(args, alpha)
	args = append(args, beta)
	args = append(args, gamma)

	retVal := q.p.Call("RotationAlphaBetaGamma", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// RotationAlphaBetaGammaToRef calls the RotationAlphaBetaGammaToRef method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#rotationalphabetagammatoref
func (q *Quaternion) RotationAlphaBetaGammaToRef(alpha float64, beta float64, gamma float64, result *Quaternion) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, alpha)
	args = append(args, beta)
	args = append(args, gamma)
	args = append(args, result.JSObject())

	q.p.Call("RotationAlphaBetaGammaToRef", args...)
}

// RotationAxis calls the RotationAxis method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#rotationaxis
func (q *Quaternion) RotationAxis(axis *Vector3, angle float64) *Quaternion {

	args := make([]interface{}, 0, 2+0)

	args = append(args, axis.JSObject())
	args = append(args, angle)

	retVal := q.p.Call("RotationAxis", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// RotationAxisToRef calls the RotationAxisToRef method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#rotationaxistoref
func (q *Quaternion) RotationAxisToRef(axis *Vector3, angle float64, result *Quaternion) *Quaternion {

	args := make([]interface{}, 0, 3+0)

	args = append(args, axis.JSObject())
	args = append(args, angle)
	args = append(args, result.JSObject())

	retVal := q.p.Call("RotationAxisToRef", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// RotationQuaternionFromAxis calls the RotationQuaternionFromAxis method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#rotationquaternionfromaxis
func (q *Quaternion) RotationQuaternionFromAxis(axis1 *Vector3, axis2 *Vector3, axis3 *Vector3) *Quaternion {

	args := make([]interface{}, 0, 3+0)

	args = append(args, axis1.JSObject())
	args = append(args, axis2.JSObject())
	args = append(args, axis3.JSObject())

	retVal := q.p.Call("RotationQuaternionFromAxis", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// RotationQuaternionFromAxisToRef calls the RotationQuaternionFromAxisToRef method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#rotationquaternionfromaxistoref
func (q *Quaternion) RotationQuaternionFromAxisToRef(axis1 *Vector3, axis2 *Vector3, axis3 *Vector3, ref *Quaternion) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, axis1.JSObject())
	args = append(args, axis2.JSObject())
	args = append(args, axis3.JSObject())
	args = append(args, ref.JSObject())

	q.p.Call("RotationQuaternionFromAxisToRef", args...)
}

// RotationYawPitchRoll calls the RotationYawPitchRoll method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#rotationyawpitchroll
func (q *Quaternion) RotationYawPitchRoll(yaw float64, pitch float64, roll float64) *Quaternion {

	args := make([]interface{}, 0, 3+0)

	args = append(args, yaw)
	args = append(args, pitch)
	args = append(args, roll)

	retVal := q.p.Call("RotationYawPitchRoll", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// RotationYawPitchRollToRef calls the RotationYawPitchRollToRef method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#rotationyawpitchrolltoref
func (q *Quaternion) RotationYawPitchRollToRef(yaw float64, pitch float64, roll float64, result *Quaternion) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, yaw)
	args = append(args, pitch)
	args = append(args, roll)
	args = append(args, result.JSObject())

	q.p.Call("RotationYawPitchRollToRef", args...)
}

// Scale calls the Scale method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#scale
func (q *Quaternion) Scale(value float64) *Quaternion {

	args := make([]interface{}, 0, 1+0)

	args = append(args, value)

	retVal := q.p.Call("scale", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// ScaleAndAddToRef calls the ScaleAndAddToRef method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#scaleandaddtoref
func (q *Quaternion) ScaleAndAddToRef(scale float64, result *Quaternion) *Quaternion {

	args := make([]interface{}, 0, 2+0)

	args = append(args, scale)
	args = append(args, result.JSObject())

	retVal := q.p.Call("scaleAndAddToRef", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// ScaleInPlace calls the ScaleInPlace method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#scaleinplace
func (q *Quaternion) ScaleInPlace(value float64) *Quaternion {

	args := make([]interface{}, 0, 1+0)

	args = append(args, value)

	retVal := q.p.Call("scaleInPlace", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// ScaleToRef calls the ScaleToRef method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#scaletoref
func (q *Quaternion) ScaleToRef(scale float64, result *Quaternion) *Quaternion {

	args := make([]interface{}, 0, 2+0)

	args = append(args, scale)
	args = append(args, result.JSObject())

	retVal := q.p.Call("scaleToRef", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// Set calls the Set method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#set
func (q *Quaternion) Set(x float64, y float64, z float64, w float64) *Quaternion {

	args := make([]interface{}, 0, 4+0)

	args = append(args, x)
	args = append(args, y)
	args = append(args, z)
	args = append(args, w)

	retVal := q.p.Call("set", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// Slerp calls the Slerp method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#slerp
func (q *Quaternion) Slerp(left *Quaternion, right *Quaternion, amount float64) *Quaternion {

	args := make([]interface{}, 0, 3+0)

	args = append(args, left.JSObject())
	args = append(args, right.JSObject())
	args = append(args, amount)

	retVal := q.p.Call("Slerp", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// SlerpToRef calls the SlerpToRef method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#slerptoref
func (q *Quaternion) SlerpToRef(left *Quaternion, right *Quaternion, amount float64, result *Quaternion) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, left.JSObject())
	args = append(args, right.JSObject())
	args = append(args, amount)
	args = append(args, result.JSObject())

	q.p.Call("SlerpToRef", args...)
}

// Subtract calls the Subtract method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#subtract
func (q *Quaternion) Subtract(other *Quaternion) *Quaternion {

	args := make([]interface{}, 0, 1+0)

	args = append(args, other.JSObject())

	retVal := q.p.Call("subtract", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// QuaternionToEulerAnglesOpts contains optional parameters for Quaternion.ToEulerAngles.
type QuaternionToEulerAnglesOpts struct {
	Order *string
}

// ToEulerAngles calls the ToEulerAngles method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#toeulerangles
func (q *Quaternion) ToEulerAngles(opts *QuaternionToEulerAnglesOpts) *Vector3 {
	if opts == nil {
		opts = &QuaternionToEulerAnglesOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Order == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Order)
	}

	retVal := q.p.Call("toEulerAngles", args...)
	return Vector3FromJSObject(retVal, q.ctx)
}

// ToEulerAnglesToRef calls the ToEulerAnglesToRef method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#toeuleranglestoref
func (q *Quaternion) ToEulerAnglesToRef(result *Vector3) *Quaternion {

	args := make([]interface{}, 0, 1+0)

	args = append(args, result.JSObject())

	retVal := q.p.Call("toEulerAnglesToRef", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// ToRotationMatrix calls the ToRotationMatrix method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#torotationmatrix
func (q *Quaternion) ToRotationMatrix(result *Matrix) *Quaternion {

	args := make([]interface{}, 0, 1+0)

	args = append(args, result.JSObject())

	retVal := q.p.Call("toRotationMatrix", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

// ToString calls the ToString method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#tostring
func (q *Quaternion) ToString() string {

	args := make([]interface{}, 0, 0+0)

	retVal := q.p.Call("toString", args...)
	return retVal.String()
}

// Zero calls the Zero method on the Quaternion object.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#zero
func (q *Quaternion) Zero() *Quaternion {

	args := make([]interface{}, 0, 0+0)

	retVal := q.p.Call("Zero", args...)
	return QuaternionFromJSObject(retVal, q.ctx)
}

/*

// W returns the W property of class Quaternion.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#w
func (q *Quaternion) W(w float64) *Quaternion {
	p := ba.ctx.Get("Quaternion").New(w)
	return QuaternionFromJSObject(p, ba.ctx)
}

// SetW sets the W property of class Quaternion.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#w
func (q *Quaternion) SetW(w float64) *Quaternion {
	p := ba.ctx.Get("Quaternion").New(w)
	return QuaternionFromJSObject(p, ba.ctx)
}

// X returns the X property of class Quaternion.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#x
func (q *Quaternion) X(x float64) *Quaternion {
	p := ba.ctx.Get("Quaternion").New(x)
	return QuaternionFromJSObject(p, ba.ctx)
}

// SetX sets the X property of class Quaternion.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#x
func (q *Quaternion) SetX(x float64) *Quaternion {
	p := ba.ctx.Get("Quaternion").New(x)
	return QuaternionFromJSObject(p, ba.ctx)
}

// Y returns the Y property of class Quaternion.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#y
func (q *Quaternion) Y(y float64) *Quaternion {
	p := ba.ctx.Get("Quaternion").New(y)
	return QuaternionFromJSObject(p, ba.ctx)
}

// SetY sets the Y property of class Quaternion.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#y
func (q *Quaternion) SetY(y float64) *Quaternion {
	p := ba.ctx.Get("Quaternion").New(y)
	return QuaternionFromJSObject(p, ba.ctx)
}

// Z returns the Z property of class Quaternion.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#z
func (q *Quaternion) Z(z float64) *Quaternion {
	p := ba.ctx.Get("Quaternion").New(z)
	return QuaternionFromJSObject(p, ba.ctx)
}

// SetZ sets the Z property of class Quaternion.
//
// https://doc.babylonjs.com/api/classes/babylon.quaternion#z
func (q *Quaternion) SetZ(z float64) *Quaternion {
	p := ba.ctx.Get("Quaternion").New(z)
	return QuaternionFromJSObject(p, ba.ctx)
}

*/
