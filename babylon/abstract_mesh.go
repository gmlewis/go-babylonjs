package babylon

import "syscall/js"

// AbstractMeshArray makes it easier to manipulate the JavaScript arrays.
type AbstractMeshArray struct {
	p  js.Value
	am []*AbstractMesh
}

// Push pushes meshes to an AbstractMesh.
func (ama *AbstractMeshArray) Push(meshes ...*AbstractMesh) {
	args := []interface{}{}
	for _, m := range meshes {
		args = append(args, m.JSObject())
	}
	ama.p.Call("push", args...)
}
