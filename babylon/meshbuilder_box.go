package babylon

// BoxOpts represents options passed to CreateBox.
type BoxOpts struct {
	Size            *float64   // size of each box side	(default: 1)
	Height          *float64   // height size, overwrites size property	(default: size)
	Width           *float64   // width size, overwrites size property	(default: size)
	Depth           *float64   // depth size, overwrites size property	(default: size)
	FaceColors      []*Color4  // array of 6 Color4, one per box face	  (default: Color4(1, 1, 1, 1) for each side)
	FaceUV          []*Vector4 // array of 6 Vector4, one per box face	(default: UVs(0, 0, 1, 1) for each side)
	Updatable       *bool      // true if the mesh is updatable	(default: false)
	SideOrientation *float64   // side orientation	(default: DEFAULTSIDE)
}

// CreateBox calls the JavaScript method of the same name.
func (b *Babylon) CreateBox(name string, opts *BoxOpts, scene *Scene) *Mesh {
	params := map[string]interface{}{}
	if opts != nil {
		if opts.Size != nil {
			params["size"] = *opts.Size
		}
		if opts.Height != nil {
			params["height"] = *opts.Height
		}
		if opts.Width != nil {
			params["width"] = *opts.Width
		}
		if opts.Depth != nil {
			params["depth"] = *opts.Depth
		}
		if opts.FaceColors != nil {
			params["faceColors"] = opts.FaceColors
		}
		if opts.FaceUV != nil {
			params["faceUV"] = opts.FaceUV
		}
		if opts.Updatable != nil {
			params["updatable"] = *opts.Updatable
		}
		if opts.SideOrientation != nil {
			params["sideOrientation"] = *opts.SideOrientation
		}
	}

	mb := b.ctx.Get("MeshBuilder")
	p := mb.Call("CreateBox", name, params, scene.JSObject())
	return MeshFromJSObject(p, b.ctx)
}
