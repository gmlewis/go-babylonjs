package babylon

// LinesOpts represents options passed to CreateLines.
type LinesOpts struct {
	Points         []*Vector3 // array of Vector3, the path of the line (REQUIRED)
	Updatable      *bool      // true if the mesh is updatable	(default: false)
	Instance       *LinesMesh // an instance of a line mesh to be updated	(default: null)
	Colors         []*Color4  // array of Color4, each point color	(default: null)
	UseVertexAlpha *bool      // false if the alpha blending is not required (faster)	(default: true)
}

// CreateLines calls the JavaScript method of the same name.
func (b *Babylon) CreateLines(name string, opts *LinesOpts, scene *Scene) *LinesMesh {
	params := map[string]interface{}{}
	if opts != nil {
		if opts.Points != nil {
			pts := Vector3Slice(opts.Points)
			params["points"] = pts.JSObject()
		}
		if opts.Updatable != nil {
			params["updatable"] = *opts.Updatable
		}
		if opts.Instance != nil {
			params["instance"] = opts.Instance.JSObject()
		}
		if opts.Colors != nil {
			colors := Color4Slice(opts.Colors)
			params["colors"] = colors.JSObject()
		}
		if opts.UseVertexAlpha != nil {
			params["useVertexAlpha"] = *opts.UseVertexAlpha
		}
	}

	mb := b.ctx.Get("MeshBuilder")
	p := mb.Call("CreateLines", name, params, scene.JSObject())
	return LinesMeshFromJSObject(p, b.ctx)
}
