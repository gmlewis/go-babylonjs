package babylon

// DOUBLESIDE calls the JavaScript method of the same name.
func (m *Mesh) DOUBLESIDE() *float64 {
	v := m.p.Get("DOUBLESIDE")
	f := v.Float()
	return &f
}

// MergeMeshesOpts contains optional parameters for MergeMeshes.
type MergeMeshesOpts struct {
	DisposeSource          *JSBool // when true (default), dispose of the vertices from the source meshes
	Allow32BitsIndices     *JSBool // when the sum of the vertices > 64k, this must be set to true
	MeshSubclass           *Mesh   // when set, vertices inserted into this Mesh. Meshes can then be merged into a Mesh sub-class.
	SubdivideWithSubMeshes *JSBool // when true (false default), subdivide mesh to his subMesh array with meshes source.
	MultiMultiMaterials    *JSBool // when true (false default), subdivide mesh and accept multiple multi materials, ignores subdivideWithSubMeshes.
}

// MergeMeshes merges the array of meshes into a single mesh for performance reasons.
func (m *Mesh) MergeMeshes(meshes []*Mesh, opts *MergeMeshesOpts) *Mesh {
	if opts == nil {
		opts = &MergeMeshesOpts{}
	}

	var ms []interface{}
	for _, mesh := range meshes {
		ms = append(ms, mesh.JSObject())
	}
	p := m.p.Call("MergeMeshes", ms,
		opts.DisposeSource.JSObject(),
		opts.Allow32BitsIndices.JSObject(),
		opts.MeshSubclass.JSObject(),
		opts.SubdivideWithSubMeshes.JSObject(),
		opts.MultiMultiMaterials.JSObject(),
	)
	return MeshFromJSObject(p, m.ctx)
}
