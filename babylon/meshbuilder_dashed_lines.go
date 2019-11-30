package babylon

// DashedLinesOpts represents options passed to CreateDashedLines.
type DashedLinesOpts struct {
	Points    []*Vector3 // array of Vector3, the path of the line (REQUIRED)
	DashSize  *float64   // size of the dashes	(default: 3)
	GapSize   *float64   // size of the gaps	(default: 1)
	DashNb    *float64   // intended number of dashes	(default: 200)
	Updatable *bool      // true if the mesh is updatable	(default: false)
	Instance  *LinesMesh // an instance of a line mesh to be updated	(default: null)
}

// CreateDashedLines calls the JavaScript method of the same name.
func (b *Babylon) CreateDashedLines(name string, opts *DashedLinesOpts, scene *Scene) *LinesMesh {
	params := map[string]interface{}{}
	if opts != nil {
		if opts.Points != nil {
			if opts.Points != nil {
				pts := Vector3Slice(opts.Points)
				params["points"] = pts.JSObject()
			}
		}
		if opts.DashSize != nil {
			params["dashSize"] = *opts.DashSize
		}
		if opts.GapSize != nil {
			params["gapSize"] = *opts.GapSize
		}
		if opts.DashNb != nil {
			params["dashNb"] = *opts.DashNb
		}
		if opts.Updatable != nil {
			params["updatable"] = *opts.Updatable
		}
		if opts.Instance != nil {
			params["instance"] = opts.Instance.JSObject()
		}
	}

	mb := b.ctx.Get("MeshBuilder")
	p := mb.Call("CreateDashedLines", name, params, scene.JSObject())
	return LinesMeshFromJSObject(p, b.ctx)
}
