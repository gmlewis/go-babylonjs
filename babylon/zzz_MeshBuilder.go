// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// MeshBuilder represents a babylon.js MeshBuilder.
// Class containing static functions to help procedurally build meshes
type MeshBuilder struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (m *MeshBuilder) JSObject() js.Value { return m.p }

// MeshBuilder returns a MeshBuilder JavaScript class.
func (ba *Babylon) MeshBuilder() *MeshBuilder {
	p := ba.ctx.Get("MeshBuilder")
	return MeshBuilderFromJSObject(p, ba.ctx)
}

// MeshBuilderFromJSObject returns a wrapped MeshBuilder JavaScript class.
func MeshBuilderFromJSObject(p js.Value, ctx js.Value) *MeshBuilder {
	return &MeshBuilder{p: p, ctx: ctx}
}

// MeshBuilderCreateBoxOpts contains optional parameters for MeshBuilder.CreateBox.
type MeshBuilderCreateBoxOpts struct {
	Scene *Scene
}

// CreateBox calls the CreateBox method on the MeshBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.meshbuilder#createbox
func (m *MeshBuilder) CreateBox(name string, options js.Value, opts *MeshBuilderCreateBoxOpts) *Mesh {
	if opts == nil {
		opts = &MeshBuilderCreateBoxOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, options)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	retVal := m.p.Call("CreateBox", args...)
	return MeshFromJSObject(retVal, m.ctx)
}

// MeshBuilderCreateCylinderOpts contains optional parameters for MeshBuilder.CreateCylinder.
type MeshBuilderCreateCylinderOpts struct {
	Scene *Scene
}

// CreateCylinder calls the CreateCylinder method on the MeshBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.meshbuilder#createcylinder
func (m *MeshBuilder) CreateCylinder(name string, options js.Value, opts *MeshBuilderCreateCylinderOpts) *Mesh {
	if opts == nil {
		opts = &MeshBuilderCreateCylinderOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, options)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	retVal := m.p.Call("CreateCylinder", args...)
	return MeshFromJSObject(retVal, m.ctx)
}

// MeshBuilderCreateDashedLinesOpts contains optional parameters for MeshBuilder.CreateDashedLines.
type MeshBuilderCreateDashedLinesOpts struct {
	Scene *Scene
}

// CreateDashedLines calls the CreateDashedLines method on the MeshBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.meshbuilder#createdashedlines
func (m *MeshBuilder) CreateDashedLines(name string, options js.Value, opts *MeshBuilderCreateDashedLinesOpts) *LinesMesh {
	if opts == nil {
		opts = &MeshBuilderCreateDashedLinesOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, options)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	retVal := m.p.Call("CreateDashedLines", args...)
	return LinesMeshFromJSObject(retVal, m.ctx)
}

// CreateDecal calls the CreateDecal method on the MeshBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.meshbuilder#createdecal
func (m *MeshBuilder) CreateDecal(name string, sourceMesh *AbstractMesh, options js.Value) *Mesh {

	args := make([]interface{}, 0, 3+0)

	args = append(args, name)
	args = append(args, sourceMesh.JSObject())
	args = append(args, options)

	retVal := m.p.Call("CreateDecal", args...)
	return MeshFromJSObject(retVal, m.ctx)
}

// MeshBuilderCreateDiscOpts contains optional parameters for MeshBuilder.CreateDisc.
type MeshBuilderCreateDiscOpts struct {
	Scene *Scene
}

// CreateDisc calls the CreateDisc method on the MeshBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.meshbuilder#createdisc
func (m *MeshBuilder) CreateDisc(name string, options js.Value, opts *MeshBuilderCreateDiscOpts) *Mesh {
	if opts == nil {
		opts = &MeshBuilderCreateDiscOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, options)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	retVal := m.p.Call("CreateDisc", args...)
	return MeshFromJSObject(retVal, m.ctx)
}

// MeshBuilderCreateGroundOpts contains optional parameters for MeshBuilder.CreateGround.
type MeshBuilderCreateGroundOpts struct {
	Scene *Scene
}

// CreateGround calls the CreateGround method on the MeshBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.meshbuilder#createground
func (m *MeshBuilder) CreateGround(name string, options js.Value, opts *MeshBuilderCreateGroundOpts) *Mesh {
	if opts == nil {
		opts = &MeshBuilderCreateGroundOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, options)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	retVal := m.p.Call("CreateGround", args...)
	return MeshFromJSObject(retVal, m.ctx)
}

// MeshBuilderCreateGroundFromHeightMapOpts contains optional parameters for MeshBuilder.CreateGroundFromHeightMap.
type MeshBuilderCreateGroundFromHeightMapOpts struct {
	Scene *Scene
}

// CreateGroundFromHeightMap calls the CreateGroundFromHeightMap method on the MeshBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.meshbuilder#creategroundfromheightmap
func (m *MeshBuilder) CreateGroundFromHeightMap(name string, url string, options js.Value, opts *MeshBuilderCreateGroundFromHeightMapOpts) *GroundMesh {
	if opts == nil {
		opts = &MeshBuilderCreateGroundFromHeightMapOpts{}
	}

	args := make([]interface{}, 0, 3+1)

	args = append(args, name)
	args = append(args, url)
	args = append(args, options)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	retVal := m.p.Call("CreateGroundFromHeightMap", args...)
	return GroundMeshFromJSObject(retVal, m.ctx)
}

// MeshBuilderCreateIcoSphereOpts contains optional parameters for MeshBuilder.CreateIcoSphere.
type MeshBuilderCreateIcoSphereOpts struct {
	Scene *Scene
}

// CreateIcoSphere calls the CreateIcoSphere method on the MeshBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.meshbuilder#createicosphere
func (m *MeshBuilder) CreateIcoSphere(name string, options js.Value, opts *MeshBuilderCreateIcoSphereOpts) *Mesh {
	if opts == nil {
		opts = &MeshBuilderCreateIcoSphereOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, options)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	retVal := m.p.Call("CreateIcoSphere", args...)
	return MeshFromJSObject(retVal, m.ctx)
}

// MeshBuilderCreateLatheOpts contains optional parameters for MeshBuilder.CreateLathe.
type MeshBuilderCreateLatheOpts struct {
	Scene *Scene
}

// CreateLathe calls the CreateLathe method on the MeshBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.meshbuilder#createlathe
func (m *MeshBuilder) CreateLathe(name string, options js.Value, opts *MeshBuilderCreateLatheOpts) *Mesh {
	if opts == nil {
		opts = &MeshBuilderCreateLatheOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, options)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	retVal := m.p.Call("CreateLathe", args...)
	return MeshFromJSObject(retVal, m.ctx)
}

// CreateLineSystem calls the CreateLineSystem method on the MeshBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.meshbuilder#createlinesystem
func (m *MeshBuilder) CreateLineSystem(name string, options js.Value, scene *Scene) *LinesMesh {

	args := make([]interface{}, 0, 3+0)

	args = append(args, name)
	args = append(args, options)
	args = append(args, scene.JSObject())

	retVal := m.p.Call("CreateLineSystem", args...)
	return LinesMeshFromJSObject(retVal, m.ctx)
}

// MeshBuilderCreateLinesOpts contains optional parameters for MeshBuilder.CreateLines.
type MeshBuilderCreateLinesOpts struct {
	Scene *Scene
}

// CreateLines calls the CreateLines method on the MeshBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.meshbuilder#createlines
func (m *MeshBuilder) CreateLines(name string, options js.Value, opts *MeshBuilderCreateLinesOpts) *LinesMesh {
	if opts == nil {
		opts = &MeshBuilderCreateLinesOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, options)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	retVal := m.p.Call("CreateLines", args...)
	return LinesMeshFromJSObject(retVal, m.ctx)
}

// MeshBuilderCreatePlaneOpts contains optional parameters for MeshBuilder.CreatePlane.
type MeshBuilderCreatePlaneOpts struct {
	Scene *Scene
}

// CreatePlane calls the CreatePlane method on the MeshBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.meshbuilder#createplane
func (m *MeshBuilder) CreatePlane(name string, options js.Value, opts *MeshBuilderCreatePlaneOpts) *Mesh {
	if opts == nil {
		opts = &MeshBuilderCreatePlaneOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, options)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	retVal := m.p.Call("CreatePlane", args...)
	return MeshFromJSObject(retVal, m.ctx)
}

// MeshBuilderCreatePolygonOpts contains optional parameters for MeshBuilder.CreatePolygon.
type MeshBuilderCreatePolygonOpts struct {
	Scene           *Scene
	EarcutInjection *interface{}
}

// CreatePolygon calls the CreatePolygon method on the MeshBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.meshbuilder#createpolygon
func (m *MeshBuilder) CreatePolygon(name string, options js.Value, opts *MeshBuilderCreatePolygonOpts) *Mesh {
	if opts == nil {
		opts = &MeshBuilderCreatePolygonOpts{}
	}

	args := make([]interface{}, 0, 2+2)

	args = append(args, name)
	args = append(args, options)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}
	if opts.EarcutInjection == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.EarcutInjection)
	}

	retVal := m.p.Call("CreatePolygon", args...)
	return MeshFromJSObject(retVal, m.ctx)
}

// MeshBuilderCreatePolyhedronOpts contains optional parameters for MeshBuilder.CreatePolyhedron.
type MeshBuilderCreatePolyhedronOpts struct {
	Scene *Scene
}

// CreatePolyhedron calls the CreatePolyhedron method on the MeshBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.meshbuilder#createpolyhedron
func (m *MeshBuilder) CreatePolyhedron(name string, options js.Value, opts *MeshBuilderCreatePolyhedronOpts) *Mesh {
	if opts == nil {
		opts = &MeshBuilderCreatePolyhedronOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, options)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	retVal := m.p.Call("CreatePolyhedron", args...)
	return MeshFromJSObject(retVal, m.ctx)
}

// MeshBuilderCreateRibbonOpts contains optional parameters for MeshBuilder.CreateRibbon.
type MeshBuilderCreateRibbonOpts struct {
	Scene *Scene
}

// CreateRibbon calls the CreateRibbon method on the MeshBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.meshbuilder#createribbon
func (m *MeshBuilder) CreateRibbon(name string, options js.Value, opts *MeshBuilderCreateRibbonOpts) *Mesh {
	if opts == nil {
		opts = &MeshBuilderCreateRibbonOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, options)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	retVal := m.p.Call("CreateRibbon", args...)
	return MeshFromJSObject(retVal, m.ctx)
}

// MeshBuilderCreateSphereOpts contains optional parameters for MeshBuilder.CreateSphere.
type MeshBuilderCreateSphereOpts struct {
	Scene *Scene
}

// CreateSphere calls the CreateSphere method on the MeshBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.meshbuilder#createsphere
func (m *MeshBuilder) CreateSphere(name string, options js.Value, opts *MeshBuilderCreateSphereOpts) *Mesh {
	if opts == nil {
		opts = &MeshBuilderCreateSphereOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, options)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	retVal := m.p.Call("CreateSphere", args...)
	return MeshFromJSObject(retVal, m.ctx)
}

// MeshBuilderCreateTiledBoxOpts contains optional parameters for MeshBuilder.CreateTiledBox.
type MeshBuilderCreateTiledBoxOpts struct {
	Scene *Scene
}

// CreateTiledBox calls the CreateTiledBox method on the MeshBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.meshbuilder#createtiledbox
func (m *MeshBuilder) CreateTiledBox(name string, options js.Value, opts *MeshBuilderCreateTiledBoxOpts) *Mesh {
	if opts == nil {
		opts = &MeshBuilderCreateTiledBoxOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, options)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	retVal := m.p.Call("CreateTiledBox", args...)
	return MeshFromJSObject(retVal, m.ctx)
}

// MeshBuilderCreateTiledGroundOpts contains optional parameters for MeshBuilder.CreateTiledGround.
type MeshBuilderCreateTiledGroundOpts struct {
	Scene *Scene
}

// CreateTiledGround calls the CreateTiledGround method on the MeshBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.meshbuilder#createtiledground
func (m *MeshBuilder) CreateTiledGround(name string, options js.Value, opts *MeshBuilderCreateTiledGroundOpts) *Mesh {
	if opts == nil {
		opts = &MeshBuilderCreateTiledGroundOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, options)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	retVal := m.p.Call("CreateTiledGround", args...)
	return MeshFromJSObject(retVal, m.ctx)
}

// MeshBuilderCreateTiledPlaneOpts contains optional parameters for MeshBuilder.CreateTiledPlane.
type MeshBuilderCreateTiledPlaneOpts struct {
	Scene *Scene
}

// CreateTiledPlane calls the CreateTiledPlane method on the MeshBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.meshbuilder#createtiledplane
func (m *MeshBuilder) CreateTiledPlane(name string, options js.Value, opts *MeshBuilderCreateTiledPlaneOpts) *Mesh {
	if opts == nil {
		opts = &MeshBuilderCreateTiledPlaneOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, options)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	retVal := m.p.Call("CreateTiledPlane", args...)
	return MeshFromJSObject(retVal, m.ctx)
}

// MeshBuilderCreateTorusOpts contains optional parameters for MeshBuilder.CreateTorus.
type MeshBuilderCreateTorusOpts struct {
	Scene *Scene
}

// CreateTorus calls the CreateTorus method on the MeshBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.meshbuilder#createtorus
func (m *MeshBuilder) CreateTorus(name string, options js.Value, opts *MeshBuilderCreateTorusOpts) *Mesh {
	if opts == nil {
		opts = &MeshBuilderCreateTorusOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, options)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	retVal := m.p.Call("CreateTorus", args...)
	return MeshFromJSObject(retVal, m.ctx)
}

// MeshBuilderCreateTorusKnotOpts contains optional parameters for MeshBuilder.CreateTorusKnot.
type MeshBuilderCreateTorusKnotOpts struct {
	Scene *Scene
}

// CreateTorusKnot calls the CreateTorusKnot method on the MeshBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.meshbuilder#createtorusknot
func (m *MeshBuilder) CreateTorusKnot(name string, options js.Value, opts *MeshBuilderCreateTorusKnotOpts) *Mesh {
	if opts == nil {
		opts = &MeshBuilderCreateTorusKnotOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, options)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	retVal := m.p.Call("CreateTorusKnot", args...)
	return MeshFromJSObject(retVal, m.ctx)
}

// MeshBuilderCreateTubeOpts contains optional parameters for MeshBuilder.CreateTube.
type MeshBuilderCreateTubeOpts struct {
	Scene *Scene
}

// CreateTube calls the CreateTube method on the MeshBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.meshbuilder#createtube
func (m *MeshBuilder) CreateTube(name string, options js.Value, opts *MeshBuilderCreateTubeOpts) *Mesh {
	if opts == nil {
		opts = &MeshBuilderCreateTubeOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, options)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	retVal := m.p.Call("CreateTube", args...)
	return MeshFromJSObject(retVal, m.ctx)
}

// MeshBuilderExtrudePolygonOpts contains optional parameters for MeshBuilder.ExtrudePolygon.
type MeshBuilderExtrudePolygonOpts struct {
	Scene           *Scene
	EarcutInjection *interface{}
}

// ExtrudePolygon calls the ExtrudePolygon method on the MeshBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.meshbuilder#extrudepolygon
func (m *MeshBuilder) ExtrudePolygon(name string, options js.Value, opts *MeshBuilderExtrudePolygonOpts) *Mesh {
	if opts == nil {
		opts = &MeshBuilderExtrudePolygonOpts{}
	}

	args := make([]interface{}, 0, 2+2)

	args = append(args, name)
	args = append(args, options)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}
	if opts.EarcutInjection == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.EarcutInjection)
	}

	retVal := m.p.Call("ExtrudePolygon", args...)
	return MeshFromJSObject(retVal, m.ctx)
}

// MeshBuilderExtrudeShapeOpts contains optional parameters for MeshBuilder.ExtrudeShape.
type MeshBuilderExtrudeShapeOpts struct {
	Scene *Scene
}

// ExtrudeShape calls the ExtrudeShape method on the MeshBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.meshbuilder#extrudeshape
func (m *MeshBuilder) ExtrudeShape(name string, options js.Value, opts *MeshBuilderExtrudeShapeOpts) *Mesh {
	if opts == nil {
		opts = &MeshBuilderExtrudeShapeOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, options)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	retVal := m.p.Call("ExtrudeShape", args...)
	return MeshFromJSObject(retVal, m.ctx)
}

// MeshBuilderExtrudeShapeCustomOpts contains optional parameters for MeshBuilder.ExtrudeShapeCustom.
type MeshBuilderExtrudeShapeCustomOpts struct {
	Scene *Scene
}

// ExtrudeShapeCustom calls the ExtrudeShapeCustom method on the MeshBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.meshbuilder#extrudeshapecustom
func (m *MeshBuilder) ExtrudeShapeCustom(name string, options js.Value, opts *MeshBuilderExtrudeShapeCustomOpts) *Mesh {
	if opts == nil {
		opts = &MeshBuilderExtrudeShapeCustomOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, options)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	retVal := m.p.Call("ExtrudeShapeCustom", args...)
	return MeshFromJSObject(retVal, m.ctx)
}

/*

 */
