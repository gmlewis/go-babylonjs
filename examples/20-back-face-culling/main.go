// 20-back-face-culling is a port to Go of the Babylon example located here:
// https://www.babylonjs-playground.com/#YDO1F#20
// and linked from here:
// https://doc.babylonjs.com/babylon101/materials
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
		camera := b.NewArcRotateCamera("Camera", 13*math.Pi/8, math.Pi/4, 5, b.Vector3().Zero(), scene, nil)
		camera.AttachControl(canvas, &babylon.ArcRotateCameraAttachControlOpts{NoPreventDefault: Bool(false)})

		light := b.NewHemisphericLight("light1", b.NewVector3(0, 1, 0), scene)
		light.SetIntensity(0.7)

		pl := b.NewPointLight("pl", b.Vector3().Zero(), scene)
		pl.SetDiffuse(b.NewColor3(1, 1, 1))
		pl.SetSpecular(b.NewColor3(1, 1, 1))
		pl.SetIntensity(0.8)

		mat := b.NewStandardMaterial("dog", scene)
		mat.SetDiffuseTexture(b.NewTexture("https://upload.wikimedia.org/wikipedia/commons/8/87/Alaskan_Malamute%2BBlank.png", scene, nil).BaseTexture)
		mat.DiffuseTexture().SetHasAlpha(true)
		mat.JSObject().Set("backFaceCulling", true)
		box := b.CreateBox("box", nil, scene)
		box.JSObject().Set("material", mat.JSObject())

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
