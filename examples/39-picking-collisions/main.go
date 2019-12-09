// 39-picking-collisions is a port to Go of the Babylon example located here:
// https://www.babylonjs-playground.com/#NU4F6Y
// and linked from here:
// https://doc.babylonjs.com/babylon101/picking_collisions
package main

import (
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

		// set up environment
		b.NewPointLight("Omni", b.NewVector3(0, 10, 20), scene)
		b.NewFreeCamera("FreeCamera", b.NewVector3(0, 0, -30), scene, nil)

		// Impact impostor
		mb := b.Mesh()
		impact := mb.CreatePlane("impact", 1, scene, nil)
		impactMat := b.NewStandardMaterial("impactMat", scene)
		impact.SetMaterial(impactMat.Material)
		impactMat.SetDiffuseTexture(b.NewTexture("textures/impact.png", scene, nil).BaseTexture)
		impactMat.DiffuseTexture().SetHasAlpha(true)
		impact.SetPosition(b.NewVector3(0, 0, -0.1))

		//Wall
		wall := mb.CreatePlane("wall", 20.0, scene, nil)
		wallMat := b.NewStandardMaterial("wallMat", scene)
		wall.SetMaterial(wallMat.Material)
		wallMat.SetEmissiveColor(b.NewColor3(0.5, 1, 0.5))

		//When pointer down event is raised
		scene.SetOnPointerDown(func(this js.Value, args []js.Value) interface{} {
			pickResult := babylon.PickingInfoFromJSObject(args[1], this)
			// if the click hits the ground object, we change the impact position
			if pickResult.Hit() {
				pp := pickResult.PickedPoint()
				ip := impact.Position()
				ip.SetX(pp.X())
				ip.SetY(pp.Y())
			}
			return nil
		})

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
