package babylon

// SphereOpts represents options passed to CreateSphere.
type SphereOpts struct {
	Diameter *float64
}

// CreateSphere calls the JavaScript method of the same name.
func (b *Babylon) CreateSphere(name string, opts *SphereOpts, scene *Scene) *Mesh {
	mb := b.ctx.Get("MeshBuilder")
	params := map[string]interface{}{}
	if opts != nil {
		if opts.Diameter != nil {
			params["diameter"] = *opts.Diameter
		}
	}
	p := mb.Call("CreateSphere", name, params, scene.JSObject())
	return MeshFromJSObject(p)
}
