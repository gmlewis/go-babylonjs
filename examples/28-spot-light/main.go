// 28-spot-light is a port to Go of the Babylon example located here:
// https://www.babylonjs-playground.com/#20OAV9#3
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

	engine := b.NewEngine(canvas, &babylon.NewEngineOpts{Antialias: babylon.Bool(true)}) // Generate the BABYLON 3D engine

	/******* Add the create scene function ******/
	createScene := func() *babylon.Scene {
		scene := b.NewScene(engine, nil)

		camera := b.NewArcRotateCamera("Camera", -math.Pi/2, math.Pi/2, 5, b.Vector3().Zero(), scene, nil)
		camera.AttachControl(canvas, &babylon.ArcRotateCameraAttachControlOpts{NoPreventDefault: Bool(true)})

		//Light direction is directly down from a position one unit up, slow decay
		light := b.NewSpotLight("spotLight", b.NewVector3(-1, 1, -1), b.NewVector3(0, -1, 0), math.Pi/2, 10, scene)
		light.SetDiffuse(b.NewColor3(1, 0, 0))
		light.SetSpecular(b.NewColor3(0, 1, 0))

		//Light direction is directly down from a position one unit up, fast decay
		light1 := b.NewSpotLight("spotLight1", b.NewVector3(1, 1, 1), b.NewVector3(0, -1, 0), math.Pi/2, 50, scene)
		light1.SetDiffuse(b.NewColor3(0, 1, 0))
		light1.SetSpecular(b.NewColor3(0, 1, 0))

		b.MeshBuilder().CreateGround("ground", &babylon.GroundOpts{Width: Float64(4), Height: Float64(4)}, scene)

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
