// 34-lightmaps is a port to Go of the Babylon example located here:
// https://www.babylonjs-playground.com/#ULACCM#2
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
		// Create scene and camera
		scene := b.NewScene(engine, nil)
		camera := b.NewFreeCamera("camera1", b.NewVector3(0, 10, -25), scene, nil)
		camera.SetTarget(b.Vector3().Zero())
		camera.AttachControl(canvas, true)

		// Create light
		light := b.NewPointLight("light", b.NewVector3(3, 3, 2), scene)
		light.SetIntensity(0.7)

		// Create lightmap texture on a material
		lightmap := b.NewTexture("textures/candleopacity.png", scene, nil)
		groundMaterial := b.NewStandardMaterial("groundMaterial", scene)
		groundMaterial.SetLightmapTexture(lightmap.BaseTexture)

		mb := b.Mesh()

		// Apply lightmap to floor mesh
		ground := mb.CreateGround("ground", 20, 20, 4, &babylon.MeshCreateGroundOpts{Scene: scene})
		ground.SetMaterial(groundMaterial.Material)
		ground.SetReceiveShadows(true)

		// Create sphere which will cast shadow on floor mesh
		sphere := mb.CreateSphere("sphere1", 16, 1, &babylon.MeshCreateSphereOpts{Scene: scene})
		sphere.Position().SetZ(-1)
		sphere.Position().SetY(2)
		shadowGenerator := b.NewShadowGenerator(1024, light.IShadowLight(), nil)
		shadowGenerator.AddShadowCaster(sphere.AbstractMesh, nil)

		// Move light in the scene
		curTime := 0.0
		scene.OnBeforeRenderObservable().Add(func(this js.Value, args []js.Value) interface{} {
			curTime += engine.GetDeltaTime()
			light.Position().SetX(math.Sin(curTime/1000.0) * 5.0)
			return nil
		}, nil)

		// Press space to swich lightmap mode
		mode := 0
		js.Global().Get("document").Set("onkeydown", js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			keyCode := args[0].Get("keyCode").Int()
			if keyCode == 0 {
				keyCode = args[0].Get("which").Int()
			}
			if keyCode == 32 {
				mode = (mode + 1) % 3
				if mode == 1 {
					light.SetLightmapMode(b.Light().LIGHTMAP_SPECULAR())
				} else if mode == 2 {
					light.SetLightmapMode(b.Light().LIGHTMAP_SHADOWSONLY())
				} else {
					light.SetLightmapMode(b.Light().LIGHTMAP_DEFAULT())
				}
			}
			return nil
		}))

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
