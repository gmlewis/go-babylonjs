package babylon

// Color4Slice is a slice of Color4 pointers.
type Color4Slice []*Color4

// JSObject returns the underlying JavaScript Array.
func (vs Color4Slice) JSObject() []interface{} {
	pts := []interface{}{}
	for _, p := range vs {
		pts = append(pts, p.JSObject())
	}
	return pts
}
