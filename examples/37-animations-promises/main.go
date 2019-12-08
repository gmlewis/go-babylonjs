// 37-animations-promises is a port to Go of the Babylon example located here:
// https://www.babylonjs-playground.com/#HZBCXR
// and linked from here:
// https://doc.babylonjs.com/babylon101/animations
package main

import (
	"log"
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

		b.NewPointLight("Omni", b.NewVector3(0, 100, 100), scene)
		camera := b.NewArcRotateCamera("Camera", 0, 0.8, 100, b.Vector3().Zero(), scene, nil)
		camera.AttachControl(canvas, true, nil)

		box1 := b.Mesh().CreateBox("Box1", 10.0, &babylon.MeshCreateBoxOpts{Scene: scene})

		materialBox := b.NewStandardMaterial("mat", scene)
		materialBox.SetDiffuseColor(b.NewColor3(0, 1, 0))

		box1.SetMaterial(materialBox.Material)

		//Create a scaling animation at 30 FPS
		animationBox := b.NewAnimation("tutoAnimation", "scaling.x", 30, b.Animation().ANIMATIONTYPE_FLOAT(),
			&babylon.NewAnimationOpts{LoopMode: Float64(b.Animation().ANIMATIONLOOPMODE_CYCLE())})
		// Animation keys
		var keys []*babylon.IAnimationKey
		//At the animation key 0, the value of scaling is "1"
		keys = append(keys, b.NewIAnimationKey(0, 1, nil))

		//At the animation key 20, the value of scaling is "0.2"
		keys = append(keys, b.NewIAnimationKey(20, 0.2, nil))

		//At the animation key 100, the value of scaling is "1"
		keys = append(keys, b.NewIAnimationKey(100, 1, nil))

		animationBox.SetKeys(keys)

		box1.Animations().Push(animationBox)

		js.Global().Call("setTimeout", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			anim := scene.BeginAnimation(box1, 0, 100, &babylon.SceneBeginAnimationOpts{Loop: Bool(false)})

			log.Printf("before")
			anim.WaitAsync()
			log.Printf("after")
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
