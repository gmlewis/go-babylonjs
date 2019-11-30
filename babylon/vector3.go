package babylon

// Vector3Slice is a slice of Vector3 pointers.
type Vector3Slice []*Vector3

// JSObject returns the underlying JavaScript Array.
func (vs Vector3Slice) JSObject() []interface{} {
	pts := []interface{}{}
	for _, p := range vs {
		pts = append(pts, p.JSObject())
	}
	return pts
}

// SetX sets the x property.
func (v *Vector3) SetX(x float64) *Vector3 {
	v.p.Set("x", x)
	return v
}

// Zero returns a new Vector3 set to (0, 0, 0).
func (v *Vector3) Zero() *Vector3 {
	p := v.ctx.Get("Vector3").New(0, 0, 0)
	return Vector3FromJSObject(p, v.ctx)
}
