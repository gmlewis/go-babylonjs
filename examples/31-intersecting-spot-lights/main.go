// 31-intersecting-spot-lights is a port to Go of the Babylon example located here:
// https://www.babylonjs-playground.com/#20OAV9#9
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

		camera := b.NewArcRotateCamera("Camera", -math.Pi/2, math.Pi/4, 5, b.Vector3().Zero(), scene, nil)
		camera.AttachControl(canvas, true, nil)

		//red light
		light := b.NewSpotLight("spotLight", b.NewVector3(-math.Cos(math.Pi/6), 1, -math.Sin(math.Pi/6)), b.NewVector3(0, -1, 0), math.Pi/2, 1.5, scene)
		light.SetDiffuse(b.NewColor3(1, 0, 0))

		//green light
		light1 := b.NewSpotLight("spotLight1", b.NewVector3(0, 1, 1-math.Sin(math.Pi/6)), b.NewVector3(0, -1, 0), math.Pi/2, 1.5, scene)
		light1.SetDiffuse(b.NewColor3(0, 1, 0))

		//blue light
		light2 := b.NewSpotLight("spotLight2", b.NewVector3(math.Cos(math.Pi/6), 1, -math.Sin(math.Pi/6)), b.NewVector3(0, -1, 0), math.Pi/2, 1.5, scene)
		light2.SetDiffuse(b.NewColor3(0, 0, 1))

		b.MeshBuilder().CreateGround("ground", &babylon.GroundOpts{Width: Float64(4), Height: Float64(4)}, scene)

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
