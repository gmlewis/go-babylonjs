package babylon

import "syscall/js"

// DOUBLESIDE gets the JavaScript property of the same name.
func (m *Mesh) DOUBLESIDE() *float64 {
	v := m.p.Get("DOUBLESIDE")
	f := v.Float()
	return &f
}

// MergeMeshesOpts contains optional parameters for MergeMeshes.
type MergeMeshesOpts struct {
	DisposeSource          *bool // when true (default), dispose of the vertices from the source meshes
	Allow32BitsIndices     *bool // when the sum of the vertices > 64k, this must be set to true
	MeshSubclass           *Mesh // when set, vertices inserted into this Mesh. Meshes can then be merged into a Mesh sub-class.
	SubdivideWithSubMeshes *bool // when true (false default), subdivide mesh to his subMesh array with meshes source.
	MultiMultiMaterials    *bool // when true (false default), subdivide mesh and accept multiple multi materials, ignores subdivideWithSubMeshes.
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

	args := make([]interface{}, 0, 6)

	args = append(args, ms)

	if opts.DisposeSource == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.DisposeSource)
	}
	if opts.Allow32BitsIndices == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Allow32BitsIndices)
	}
	if opts.MeshSubclass == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.MeshSubclass.JSObject())
	}
	if opts.SubdivideWithSubMeshes == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.SubdivideWithSubMeshes)
	}
	if opts.MultiMultiMaterials == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.MultiMultiMaterials)
	}

	p := m.p.Call("MergeMeshes", args...)
	return MeshFromJSObject(p, m.ctx)
}
