package babylon

// Position returns the position property.
func (t *TransformNode) Position() *Vector3 {
	p := t.p.Get("position")
	return Vector3FromJSObject(p, t.ctx)
}
