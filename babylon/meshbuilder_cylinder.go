package babylon

// CylinderOpts represents options passed to CreateCylinder.
type CylinderOpts struct {
	Arc             *float64
	BackUVs         *Vector4
	Cap             *float64
	Diameter        *float64
	DiameterBottom  *float64
	DiameterTop     *float64
	Enclose         *bool
	FaceColors      []*Color4
	FaceUV          []*Vector4
	FrontUVs        *Vector4
	HasRings        *bool
	Height          *float64
	SideOrientation *float64
	Subdivisions    *float64
	Tessellation    *float64
	Updatable       *bool
}

// CreateCylinder calls the JavaScript method of the same name.
func (mb *MeshBuilder) CreateCylinder(name string, opts *CylinderOpts, scene *Scene) *Mesh {
	params := map[string]interface{}{}
	if opts != nil {
		if opts.Arc != nil {
			params["arc"] = *opts.Arc
		}
		if opts.BackUVs != nil {
			params["backUVs"] = *opts.BackUVs
		}
		if opts.Cap != nil {
			params["cap"] = *opts.Cap
		}
		if opts.Diameter != nil {
			params["diameter"] = *opts.Diameter
		}
		if opts.DiameterBottom != nil {
			params["diameterBottom"] = *opts.DiameterBottom
		}
		if opts.DiameterTop != nil {
			params["diameterTop"] = *opts.DiameterTop
		}
		if opts.Enclose != nil {
			params["enclose"] = *opts.Enclose
		}
		if opts.FaceColors != nil {
			colors := Color4Slice(opts.FaceColors)
			params["faceColors"] = colors.JSObject()
		}
		if opts.FaceUV != nil {
			pts := Vector4Slice(opts.FaceUV)
			params["faceUV"] = pts.JSObject()
		}
		if opts.FrontUVs != nil {
			params["frontUVs"] = *opts.FrontUVs
		}
		if opts.HasRings != nil {
			params["hasRings"] = *opts.HasRings
		}
		if opts.Height != nil {
			params["height"] = *opts.Height
		}
		if opts.SideOrientation != nil {
			params["sideOrientation"] = *opts.SideOrientation
		}
		if opts.Subdivisions != nil {
			params["subdivisions"] = *opts.Subdivisions
		}
		if opts.Tessellation != nil {
			params["tessellation"] = *opts.Tessellation
		}
		if opts.Updatable != nil {
			params["updatable"] = *opts.Updatable
		}
	}

	p := mb.p.Call("CreateCylinder", name, params, scene.JSObject())
	return MeshFromJSObject(p, mb.ctx)
}
