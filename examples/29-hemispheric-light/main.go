// 29-hemispheric-light is a port to Go of the Babylon example located here:
// https://www.babylonjs-playground.com/#20OAV9#5
// and linked from here:
// https://doc.babylonjs.com/babylon101/lights
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

		camera := b.NewArcRotateCamera("Camera", -math.Pi/2, math.Pi/2, 5, b.Vector3().Zero(), scene, nil)
		camera.AttachControl(canvas, true, nil)

		light := b.NewHemisphericLight("hemiLight", b.NewVector3(-1, 1, 0), scene)
		light.SetDiffuse(b.NewColor3(1, 0, 0))
		light.SetSpecular(b.NewColor3(0, 1, 0))
		light.SetGroundColor(b.NewColor3(0, 1, 0))

		sphere := b.MeshBuilder().CreateSphere("sphere", nil, scene)
		sphere.Position().SetZ(1)

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
