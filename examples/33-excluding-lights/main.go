// 33-excluding-lights is a port to Go of the Babylon example located here:
// https://www.babylonjs-playground.com/#20OAV9#8
// and linked from here:
// https://doc.babylonjs.com/babylon101/lights
package main

import (
	"fmt"
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

		// Set up camera
		camera := b.NewArcRotateCamera("Camera", -math.Pi/2, math.Pi/4, 8, b.Vector3().Zero(), scene, nil)
		camera.AttachControl(canvas, true, nil)

		//Light direction is up and left
		light0 := b.NewHemisphericLight("hemiLight", b.NewVector3(-1, 1, 0), scene)
		light0.SetDiffuse(b.NewColor3(1, 0, 0))
		light0.SetSpecular(b.NewColor3(0, 1, 0))
		light0.SetGroundColor(b.NewColor3(0, 1, 0))

		light1 := b.NewHemisphericLight("hemiLight", b.NewVector3(-1, 1, 0), scene)
		light1.SetDiffuse(b.NewColor3(1, 1, 1))
		light1.SetSpecular(b.NewColor3(1, 1, 1))
		light1.SetGroundColor(b.NewColor3(0, 0, 0))

		sphere := b.MeshBuilder().CreateSphere("sphere", &babylon.SphereOpts{Diameter: Float64(0.5)}, scene)

		spheres := []*babylon.AbstractMesh{}
		for i := 0; i < 25; i++ {
			name := fmt.Sprintf("sphere%v", i)
			s := sphere.Clone(&babylon.MeshCloneOpts{Name: &name})
			s.Position().SetX(-2.0 + float64(i%5))
			s.Position().SetY(2.0 - math.Floor(float64(i)/5.0))
			spheres = append(spheres, s)
		}

		// Try commenting out one or both of these lines to see the effect.
		light0.ExcludedMeshes().Push(spheres[7], spheres[18])
		light1.IncludedOnlyMeshes().Push(spheres[7], spheres[18])

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
