package main

import (
	"log"
	"math"
	"syscall/js"

	"github.com/gmlewis/go-babylonjs/babylon"
)

func main() {
	window := js.Global().Get("window")
	doc := js.Global().Get("document")
	canvas := doc.Call("getElementById", "renderCanvas") // Get the canvas element
	log.Printf("window=%T=%v, doc=%T=%v, canvas=%T=%v", window, window, doc, doc, canvas, canvas)

	b := babylon.New()

	engine := b.NewEngine(canvas, true) // Generate the BABYLON 3D engine

	/******* Add the create scene function ******/
	createScene := func() *babylon.Scene {

		// Create the scene space
		scene := b.NewScene(engine)

		// Add a camera to the scene and attach it to the canvas
		camera := b.NewArcRotateCamera("Camera", math.Pi/2, math.Pi/2, 2, b.NewVector3(0, 0, 5), scene)
		camera.AttachControl(canvas, true)

		// Add lights to the scene
		b.NewHemisphericLight("light1", b.NewVector3(1, 1, 0), scene)
		b.NewPointLight("light2", b.NewVector3(0, 1, -1), scene)

		// Add and manipulate meshes in the scene
		b.CreateSphere("sphere", &babylon.SphereOpts{Diameter: Float64(2)}, scene)

		return scene
	}
	/******* End of the create scene function ******/

	scene := createScene() //Call the createScene function

	// Register a render loop to repeatedly render the scene
	engine.RunRenderLoop(func() {
		scene.Render()
	})

	// Watch for browser/canvas resize events
	window.Call("addEventListener", "resize", func() { // Get the canvas element
		engine.Resize()
	})
}

func Float64(v float64) *float64 {
	return &v
}
