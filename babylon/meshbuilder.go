package babylon

type SphereOpts struct {
	Diameter *float64
}

func (b *Babylon) CreateSphere(name string, opts *SphereOpts, scene *Scene) *Mesh {
	mb := b.ctx.Get("MeshBuilder")
	params := map[string]interface{}{}
	if opts != nil {
		if opts.Diameter != nil {
			params["diameter"] = *opts.Diameter
		}
	}
	p := mb.Call("CreateSphere", name, params, scene)
	return MeshFromJSObject(p)
}
