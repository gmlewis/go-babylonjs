// 05-discover-basic-elements-source-plane is a port to Go of the Babylon example located here:
// https://www.babylonjs-playground.com/#LXZPJK#2
// and linked from here:
// https://doc.babylonjs.com/babylon101/discover_basic_elements
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

		// Create the scene space
		scene := b.NewScene(engine, nil)

		// Add a camera to the scene and attach it to the canvas
		camera := b.NewArcRotateCamera("Camera", -math.Pi/2, math.Pi/2, 4, b.Vector3().Zero(), scene, nil)
		camera.AttachControl(canvas, true, nil)

		// Add lights to the scene
		b.NewHemisphericLight("light1", b.NewVector3(1, 1, 0), scene)
		b.NewPointLight("light2", b.NewVector3(0, 1, -1), scene)

		sourcePlane := b.NewPlane(0, -1, 1, 0)
		sourcePlane.Normalize()

		// Add and manipulate meshes in the scene
		b.MeshBuilder().CreatePlane("plane", &babylon.PlaneOpts{Height: Float64(2), Width: Float64(1), SourcePlane: sourcePlane, SideOrientation: Float64(b.Mesh().DOUBLESIDE())}, scene)

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
