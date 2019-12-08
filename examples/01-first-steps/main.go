// 01-first-steps is a port to Go of the Babylon example located here:
// https://www.babylonjs-playground.com/#WG9OY#1
// and linked from here:
// https://doc.babylonjs.com/babylon101/first
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
		camera := b.NewArcRotateCamera("Camera", math.Pi/2, math.Pi/2, 2, b.NewVector3(0, 0, 5), scene, nil)
		camera.AttachControl(canvas, true, nil)

		// Add lights to the scene
		b.NewHemisphericLight("light1", b.NewVector3(1, 1, 0), scene)
		b.NewPointLight("light2", b.NewVector3(0, 1, -1), scene)

		// Add and manipulate meshes in the scene
		b.MeshBuilder().CreateSphere("sphere", &babylon.SphereOpts{Diameter: Float64(2)}, scene)

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

// Bool returns the pointer to the provided bool.
func Bool(v bool) *bool {
	return &v
}

// Float64 returns the pointer to the provided float64.
func Float64(v float64) *float64 {
	return &v
}
