package babylon

// PlaneOpts represents options passed to CreatePlane.
type PlaneOpts struct {
	Size            *float64   // side size of the plane	(default: 1)
	Width           *float64   // size of the width	(default: size)
	Height          *float64   // size of the height	(default: size)
	Updatable       *bool      // true if the mesh is updatable	(default: false)
	SideOrientation *float64   // side orientation	(default: DEFAULTSIDE)
	FrontUVs        []*Vector4 // array of Vector4, ONLY WHEN sideOrientation:BABYLON.Mesh.DOUBLESIDE set	(default: Vector4(0,0, 1,1))
	BackUVs         []*Vector4 // array of Vector4, ONLY WHEN sideOrientation:BABYLON.Mesh.DOUBLESIDE set	(default: Vector4(0,0, 1,1))
	SourcePlane     *Plane     // source plane (maths) the mesh will be transformed to	(default: null)
}

// CreatePlane calls the JavaScript method of the same name.
func (b *Babylon) CreatePlane(name string, opts *PlaneOpts, scene *Scene) *Mesh {
	params := map[string]interface{}{}
	if opts != nil {
		if opts.Size != nil {
			params["size"] = *opts.Size
		}
		if opts.Width != nil {
			params["width"] = *opts.Width
		}
		if opts.Height != nil {
			params["height"] = *opts.Height
		}
		if opts.Updatable != nil {
			params["updatable"] = *opts.Updatable
		}
		if opts.SideOrientation != nil {
			params["sideOrientation"] = *opts.SideOrientation
		}
		if opts.FrontUVs != nil {
			params["frontUVs"] = opts.FrontUVs
		}
		if opts.BackUVs != nil {
			params["backUVs"] = opts.BackUVs
		}
		if opts.SourcePlane != nil {
			params["sourcePlane"] = *opts.SourcePlane
		}
	}

	mb := b.ctx.Get("MeshBuilder")
	p := mb.Call("CreatePlane", name, params, scene.JSObject())
	return MeshFromJSObject(p, b.ctx)
}
