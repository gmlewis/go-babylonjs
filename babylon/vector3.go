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

// Vector3Array2DToJSArray returns a JavaScript 2D Array for the wrapped array.
func Vector3Array2DToJSArray(array [][]*Vector3) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, Vector3ArrayToJSArray(v))
	}
	return result
}
