// 42-sake-arc-reactor is a demonstration of loading an animated
// glTF file and rendering it with BabylonJS using Go and WebAssembly.

// The Sake Ironman: Arc Reactor glTF model is downloaded from:
// https://sketchfab.com/3d-models/ironman-arc-reactor-with-sound-d87547f21f904cfa954f4cf77a1409ac
// It is licensed as: CC Attribution-NonCommercial-NoDerivsCC Attribution-NonCommercial-NoDerivs
// https://creativecommons.org/licenses/by-nc-nd/4.0/
package main

import (
	"log"
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
		// Create a scene.
		scene := b.NewScene(engine, nil)

		// Set up environment
		camera := b.NewArcRotateCamera("ArcRotateCamera", 1, 0.8, 50, b.NewVector3(0, 0, 0), scene, nil)
		camera.AttachControl(canvas, true, nil)

		sceneOpts := &babylon.SceneCreateDefaultCameraOrLightOpts{
			CreateArcRotateCamera: Bool(true),
			Replace:               Bool(true),
			AttachCameraControls:  Bool(true),
		}

		// Append glTF model to scene.
		opts := &babylon.SceneLoaderAppendOpts{
			SceneFilename: String("scene.gltf"),
			Scene:         scene,
			OnSuccess: func(this js.Value, args []js.Value) interface{} {
				scene.CreateDefaultCameraOrLight(sceneOpts)
				scene.SetClearColor(b.NewColor4(0, 0, 0, 1))
				camera = babylon.ArcRotateCameraFromJSObject(scene.ActiveCamera().JSObject(), this)
				camera.SetAlpha(math.Pi)
				log.Printf("camera: r=%v", camera.Radius())
				camera.SetRadius(3000)
				log.Printf("scene loaded")
				return nil
			},
		}
		b.SceneLoader().Append("", opts)

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

// String returns the pointer to the provided string.
func String(v string) *string {
	return &v
}
