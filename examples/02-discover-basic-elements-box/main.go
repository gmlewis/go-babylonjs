// 02-discover-basic-elements-box is a port to Go of the Babylon example located here:
// https://www.babylonjs-playground.com/#3QW4J1#1
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

	engine := b.NewEngine(canvas, &babylon.NewEngineOpts{Antialias: babylon.Bool(true)}) // Generate the BABYLON 3D engine

	/******* Add the create scene function ******/
	createScene := func() *babylon.Scene {

		// Create the scene space
		scene := b.NewScene(engine, nil)

		// Add a camera to the scene and attach it to the canvas
		camera := b.NewArcRotateCamera("Camera", 3*math.Pi/4, math.Pi/4, 4, b.Vector3().Zero(), scene, nil)
		camera.AttachControl(canvas, &babylon.ArcRotateCameraAttachControlOpts{NoPreventDefault: babylon.Bool(true)})

		// Add lights to the scene
		b.NewHemisphericLight("light1", b.NewVector3(1, 1, 0), scene)
		b.NewPointLight("light2", b.NewVector3(0, 1, -1), scene)

		// Add and manipulate meshes in the scene
		b.CreateBox("box", &babylon.BoxOpts{Height: Float64(1), Width: Float64(0.75), Depth: Float64(0.25)}, scene)

		return scene
	}
	/******* End of the create scene function ******/

	scene := createScene() //Call the createScene function

	// Register a render loop to repeatedly render the scene
	engine.RunRenderLoop(func() {
		scene.Render(nil)
	})

	// Watch for browser/canvas resize events
	window := js.Global().Get("window")
	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		engine.Resize()
		return nil
	})
	// Note that "cb" is never released since it is needed for resizing.
	window.Call("addEventListener", "resize", cb)

	// prevent program from terminating
	c := make(chan struct{}, 0)
	<-c
}

// Float64 returns the pointer to the provided float64.
func Float64(v float64) *float64 {
	return &v
}
