// 38-mesh-collisions is a port to Go of the Babylon example located here:
// https://www.babylonjs-playground.com/#KQV9SA
// and linked from here:
// https://doc.babylonjs.com/babylon101/intersect_collisions_-_mesh
package main

import (
	"math"
	"syscall/js"

	"github.com/gmlewis/go-babylonjs/babylon"
)

func main() {
	doc := js.Global().Get("document")
	canvas := doc.Call("getElementById", "renderCanvas") // Get the canvas element

	b := babylon.New()

	engine := b.NewEngine(canvas, &babylon.NewEngineOpts{Antialias: Bool(true)}) // Generate the BABYLON 3D engine

	/******* Add the create scene function ******/
	createScene := func() *babylon.Scene {
		scene := b.NewScene(engine, nil)

		camera := b.NewArcRotateCamera("Camera", 1, 0.8, 70, b.NewVector3(5, 0, 0), scene, nil)
		camera.AttachControl(canvas, true, nil)

		// Material
		matPlan := b.NewStandardMaterial("matPlan1", scene)
		matPlan.SetBackFaceCulling(false)
		matPlan.SetEmissiveColor(b.NewColor3(0.2, 1, 0.2))

		matBB := b.NewStandardMaterial("matBB", scene)
		matBB.SetEmissiveColor(b.NewColor3(1, 1, 1))
		matBB.SetWireframe(true)

		mb := b.Mesh()

		// Intersection point
		pointToIntersect := b.NewVector3(-30, 0, 0)
		origin := mb.CreateSphere("origin", 4, 0.3, &babylon.MeshCreateSphereOpts{Scene: scene})
		origin.SetPosition(pointToIntersect)
		origin.SetMaterial(matPlan.Material)

		// Create two planes
		plan1 := mb.CreatePlane("plane1", 20, scene, nil)
		plan1.SetPosition(b.NewVector3(13, 0, 0))
		plan1.Rotation().SetX(-math.Pi / 4)
		plan1.SetMaterial(matPlan.Material)

		plan2 := mb.CreatePlane("plane2", 20, scene, nil)
		plan2.SetPosition(b.NewVector3(-13, 0, 0))
		plan2.Rotation().SetX(-math.Pi / 4)
		plan2.SetMaterial(matPlan.Material)

		// AABB - Axis aligned bounding box
		planAABB := mb.CreateBox("AABB", 20, &babylon.MeshCreateBoxOpts{Scene: scene})
		planAABB.SetMaterial(matBB.Material)
		planAABB.SetPosition(b.NewVector3(13, 0, 0))
		planAABB.SetScaling(b.NewVector3(1, math.Cos(math.Pi/4), math.Cos(math.Pi/4)))

		// OBB - Object bounding box
		planOBB := mb.CreateBox("OBB", 20, &babylon.MeshCreateBoxOpts{Scene: scene})
		planOBB.SetScaling(b.NewVector3(1, 1, 0.05))
		planOBB.Node.SetParent(plan2.Node) // TODO: Without the 1st .Node, this calls .TransformNode, and it doesn't work.
		planOBB.SetMaterial(matBB.Material)

		// Balloons
		balloon1 := mb.CreateSphere("balloon1", 10, 2.0, &babylon.MeshCreateSphereOpts{Scene: scene})
		balloon2 := mb.CreateSphere("balloon2", 10, 2.0, &babylon.MeshCreateSphereOpts{Scene: scene})
		balloon3 := mb.CreateSphere("balloon3", 10, 2.0, &babylon.MeshCreateSphereOpts{Scene: scene})
		mat1 := b.NewStandardMaterial("matBallon", scene)
		mat2 := b.NewStandardMaterial("matBallon", scene)
		mat3 := b.NewStandardMaterial("matBallon", scene)
		balloon1.SetMaterial(mat1.Material)
		balloon2.SetMaterial(mat2.Material)
		balloon3.SetMaterial(mat3.Material)

		balloon1.SetPosition(b.NewVector3(6, 5, 0))
		balloon2.SetPosition(b.NewVector3(-6, 5, 0))
		balloon3.SetPosition(b.NewVector3(-30, 5, 0))

		//Animation
		alpha := math.Pi
		scene.RegisterBeforeRender(func(this js.Value, args []js.Value) interface{} {

			//Balloon 1 intersection -- Precise = false
			if balloon1.IntersectsMesh(plan1.AbstractMesh, &babylon.AbstractMeshIntersectsMeshOpts{Precise: Bool(false)}) {
				mat1.SetEmissiveColor(b.NewColor3(1, 0, 0))
			} else {
				mat1.SetEmissiveColor(b.NewColor3(1, 1, 1))
			}

			//Balloon 2 intersection -- Precise = true
			if balloon2.IntersectsMesh(plan2.AbstractMesh, &babylon.AbstractMeshIntersectsMeshOpts{Precise: Bool(true)}) {
				mat2.SetEmissiveColor(b.NewColor3(1, 0, 0))
			} else {
				mat2.SetEmissiveColor(b.NewColor3(1, 1, 1))
			}

			//balloon 3 intersection on single point
			if balloon3.IntersectsPoint(pointToIntersect) {
				mat3.SetEmissiveColor(b.NewColor3(1, 0, 0))
			} else {
				mat3.SetEmissiveColor(b.NewColor3(1, 1, 1))
			}

			alpha += 0.01
			newY := balloon1.Position().Y() + math.Cos(alpha)/10
			balloon1.Position().SetY(newY)
			balloon2.Position().SetY(newY)
			balloon3.Position().SetY(newY)
			return nil
		})

		return scene
	}
	/******* End of the create scene function ******/

	scene := createScene() //Call the createScene function

	// Register a render loop to repeatedly render the scene
	engine.RunRenderLoop(scene.RenderLoopFunc(nil))

	// Watch for browser/canvas resize events
	window := js.Global().Get("window")
	// Note that engine.ResizeFunc is never released since it is needed for resizing.
	window.Call("addEventListener", "resize", engine.ResizeFunc())

	// prevent program from terminating
	c := make(chan struct{}, 0)
	<-c
}

// Bool returns the pointer to the provided bool.
func Bool(v bool) *bool {
	return &v
}

// Float64 returns the pointer to the provided float64.
func Float64(v float64) *float64 {
	return &v
}
