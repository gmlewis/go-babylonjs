package babylon

// Vector4Slice is a slice of Vector4 pointers.
type Vector4Slice []*Vector4

// JSObject returns the underlying JavaScript Array.
func (vs Vector4Slice) JSObject() []interface{} {
	pts := []interface{}{}
	for _, p := range vs {
		pts = append(pts, p.JSObject())
	}
	return pts
}

// Zero returns a new Vector4 set to (0, 0, 0, 0).
func (v *Vector4) Zero() *Vector4 {
	p := v.ctx.Get("Vector4").New(0, 0, 0, 0)
	return Vector4FromJSObject(p, v.ctx)
}
