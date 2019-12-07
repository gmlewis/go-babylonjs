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
