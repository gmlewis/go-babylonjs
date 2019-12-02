// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// MeshButton3D represents a babylon.js MeshButton3D.
// Class used to create an interactable object. It&#39;s a 3D button using a mesh coming from the current scene
type MeshButton3D struct {
	*Button3D
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (m *MeshButton3D) JSObject() js.Value { return m.p }

// MeshButton3D returns a MeshButton3D JavaScript class.
func (ba *Babylon) MeshButton3D() *MeshButton3D {
	p := ba.ctx.Get("MeshButton3D")
	return MeshButton3DFromJSObject(p, ba.ctx)
}

// MeshButton3DFromJSObject returns a wrapped MeshButton3D JavaScript class.
func MeshButton3DFromJSObject(p js.Value, ctx js.Value) *MeshButton3D {
	return &MeshButton3D{Button3D: Button3DFromJSObject(p, ctx), ctx: ctx}
}

// MeshButton3DArrayToJSArray returns a JavaScript Array for the wrapped array.
func MeshButton3DArrayToJSArray(array []*MeshButton3D) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewMeshButton3DOpts contains optional parameters for NewMeshButton3D.
type NewMeshButton3DOpts struct {
	Name *string
}

// NewMeshButton3D returns a new MeshButton3D object.
//
// https://doc.babylonjs.com/api/classes/babylon.meshbutton3d
func (ba *Babylon) NewMeshButton3D(mesh *Mesh, opts *NewMeshButton3DOpts) *MeshButton3D {
	if opts == nil {
		opts = &NewMeshButton3DOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, mesh.JSObject())

	if opts.Name == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Name)
	}

	p := ba.ctx.Get("MeshButton3D").New(args...)
	return MeshButton3DFromJSObject(p, ba.ctx)
}

/*

 */
