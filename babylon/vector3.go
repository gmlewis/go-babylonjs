package babylon

// Zero returns a new Vector3 set to (0.0, 0.0, 0.0).
func (v *Vector3) Zero() *Vector3 {
	p := v.ctx.Get("Vector3").New(0, 0, 0)
	return Vector3FromJSObject(p, v.ctx)
}
