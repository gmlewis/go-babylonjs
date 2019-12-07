package babylon

// SphereOpts represents options passed to CreateSphere.
type SphereOpts struct {
	Segments        *float64 // number of horizontal segments	(default: 32)
	Diameter        *float64 // diameter of the sphere	(default: 1)
	DiameterX       *float64 // diameter on X axis, overwrites diameter property	(default: diameter)
	DiameterY       *float64 // diameter on Y axis, overwrites diameter property	(default: diameter)
	DiameterZ       *float64 // diameter on Z axis, overwrites diameter property	(default: diameter)
	Arc             *float64 // ratio of the circumference (latitude) between 0 and 1	(default: 1)
	Slice           *float64 // ratio of the height (longitude) between 0 and 1	(default: 1)
	Updatable       *bool    // true if the mesh is updatable	(default: false)
	SideOrientation *float64 // side orientation	(default: DEFAULTSIDE)
}

// CreateSphere calls the JavaScript method of the same name.
func (mb *MeshBuilder) CreateSphere(name string, opts *SphereOpts, scene *Scene) *Mesh {
	params := map[string]interface{}{}
	if opts != nil {
		if opts.Segments != nil {
			params["segments"] = *opts.Segments
		}
		if opts.Diameter != nil {
			params["diameter"] = *opts.Diameter
		}
		if opts.DiameterX != nil {
			params["diameterX"] = *opts.DiameterX
		}
		if opts.DiameterY != nil {
			params["diameterY"] = *opts.DiameterY
		}
		if opts.DiameterZ != nil {
			params["diameterZ"] = *opts.DiameterZ
		}
		if opts.Arc != nil {
			params["arc"] = *opts.Arc
		}
		if opts.Slice != nil {
			params["slice"] = *opts.Slice
		}
		if opts.Updatable != nil {
			params["updatable"] = *opts.Updatable
		}
		if opts.SideOrientation != nil {
			params["sideOrientation"] = *opts.SideOrientation
		}
	}

	p := mb.p.Call("CreateSphere", name, params, scene.JSObject())
	return MeshFromJSObject(p, mb.ctx)
}
