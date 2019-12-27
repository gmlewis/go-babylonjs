// 43-colored-line-system is a port to Go of the Babylon example located here:
// https://www.babylonjs-playground.com/#165IV6#80
// and linked from here:
// https://doc.babylonjs.com/babylon101/parametric_shapes
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

		camera := b.NewArcRotateCamera("Camera", 3*math.Pi/2, 3*math.Pi/8, 30, b.Vector3().Zero(), scene, nil)
		camera.AttachControl(canvas, true, nil)

		b.NewHemisphericLight("hemi", b.NewVector3(0, 1, 0), scene)

		//Array of points to construct lines
		myLines := [][]*babylon.Vector3{
			{
				b.NewVector3(0, 0, 10),
				b.NewVector3(10, 0, 10),
			},
			{
				b.NewVector3(10, 0, 0),
				b.NewVector3(10, 10, 0),
				b.NewVector3(0, 10, 0),
			},
		}

		//Create linesystem with updatable parameter set to true for later changes
		opts := map[string]interface{}{
			"lines":      babylon.Vector3Array2DToJSArray(myLines),
			"updateable": true,
		}
		linesystem := b.MeshBuilder().CreateLineSystem("linesystem", js.ValueOf(opts), scene)
		linesystem.SetColor(b.NewColor3(0, 1, 0))

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
