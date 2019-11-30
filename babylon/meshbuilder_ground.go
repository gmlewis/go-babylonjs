package babylon

// GroundOpts represents options passed to CreateGround.
type GroundOpts struct {
	Width        *float64 // size of the width	(default: 1)
	Height       *float64 // size of the height	(default: 1)
	Updatable    *bool    // true if the mesh is updatable	(default: false)
	Subdivisions *float64 // number of square subdivisions	(default: 1)
}

// CreateGround calls the JavaScript method of the same name.
func (b *Babylon) CreateGround(name string, opts *GroundOpts, scene *Scene) *Mesh {
	params := map[string]interface{}{}
	if opts != nil {
		if opts.Width != nil {
			params["width"] = *opts.Width
		}
		if opts.Height != nil {
			params["height"] = *opts.Height
		}
		if opts.Updatable != nil {
			params["updatable"] = *opts.Updatable
		}
		if opts.Subdivisions != nil {
			params["subdivisions"] = *opts.Subdivisions
		}
	}

	mb := b.ctx.Get("MeshBuilder")
	p := mb.Call("CreateGround", name, params, scene.JSObject())
	return MeshFromJSObject(p, b.ctx)
}
