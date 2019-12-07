// 08-parametric-shapes-spiral-from-lines is a port to Go of the Babylon example located here:
// https://www.babylonjs-playground.com/#165IV6#61
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

	engine := b.NewEngine(canvas, &babylon.NewEngineOpts{Antialias: babylon.Bool(true)}) // Generate the BABYLON 3D engine

	/******* Add the create scene function ******/
	createScene := func() *babylon.Scene {
		scene := b.NewScene(engine, nil)

		camera := b.NewArcRotateCamera("Camera", 0, 0, 0, b.Vector3().Zero(), scene, nil)
		camera.SetPosition(b.NewVector3(5, 5, -5))
		camera.AttachControl(canvas, &babylon.ArcRotateCameraAttachControlOpts{NoPreventDefault: babylon.Bool(true)})

		b.NewHemisphericLight("hemi", b.NewVector3(0, 1, 0), scene)

		// Array of points to construct a spiral with lines
		var myPoints []*babylon.Vector3

		deltaTheta := 0.1
		deltaY := 0.005

		radius := 1.0
		theta := 0.0
		Y := 0.0
		for i := 0; i < 400; i++ {
			myPoints = append(myPoints, b.NewVector3(radius*math.Cos(theta), Y, radius*math.Sin(theta)))
			theta += deltaTheta
			Y += deltaY
		}

		// Create lines
		b.MeshBuilder().CreateLines("lines", &babylon.LinesOpts{Points: myPoints}, scene)

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
