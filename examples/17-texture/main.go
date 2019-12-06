// 17-texture is a port to Go of the Babylon example located here:
// https://www.babylonjs-playground.com/#20OAV9#15
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
		camera := b.NewArcRotateCamera("Camera", -math.Pi/2, math.Pi/4, 5, b.Vector3().Zero(), scene, nil)
		camera.AttachControl(canvas, &babylon.ArcRotateCameraAttachControlOpts{NoPreventDefault: Bool(true)})

		//Light direction is up and left
		light := b.NewHemisphericLight("hemiLight", b.NewVector3(-1, 1, 0), scene)
		light.SetDiffuse(b.NewColor3(1, 0, 0))
		light.SetSpecular(b.NewColor3(0, 1, 0))
		light.SetGroundColor(b.NewColor3(0, 1, 0))

		grass0 := b.NewStandardMaterial("grass0", scene)
		grass0.SetDiffuseTexture(b.NewTexture("textures/grass.png", scene, nil).BaseTexture)

		grass1 := b.NewStandardMaterial("grass1", scene)
		grass1.SetEmissiveTexture(b.NewTexture("textures/grass.png", scene, nil).BaseTexture)

		grass2 := b.NewStandardMaterial("grass2", scene)
		grass2.SetAmbientTexture(b.NewTexture("textures/grass.png", scene, nil).BaseTexture)
		grass2.SetDiffuseColor(b.NewColor3(1, 0, 0))

		//diffuse texture
		sphere0 := b.CreateSphere("sphere0", nil, scene)
		sphere0.Position().SetX(-1.5)
		sphere0.JSObject().Set("material", grass0.JSObject())

		//emissive texture
		sphere1 := b.CreateSphere("sphere1", nil, scene)
		sphere1.JSObject().Set("material", grass1.JSObject())

		//ambient texture and diffuse color
		sphere2 := b.CreateSphere("sphere2", nil, scene)
		sphere2.JSObject().Set("material", grass2.JSObject())
		sphere2.Position().SetX(1.5)

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
