// 13-position-scaled-rotated-boxes is a port to Go of the Babylon example located here:
// https://www.babylonjs-playground.com/#CURCZC
// and linked from here:
// https://doc.babylonjs.com/babylon101/position
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

		camera := b.NewArcRotateCamera("Camera", math.Pi, math.Pi/8, 150, b.Vector3().Zero(), scene, nil)
		camera.AttachControl(canvas, true, nil)

		b.NewHemisphericLight("hemi", b.NewVector3(0, 1, 0), scene)

		mb := b.MeshBuilder()

		//Creation of 3 boxes and 2 spheres
		box1 := mb.CreateBox("Box1", &babylon.BoxOpts{Size: Float64(6.0)}, scene)
		box2 := mb.CreateBox("Box2", &babylon.BoxOpts{Size: Float64(6.0)}, scene)
		box3 := mb.CreateBox("Box3", &babylon.BoxOpts{Size: Float64(6.0)}, scene)
		box4 := mb.CreateBox("Box4", &babylon.BoxOpts{Size: Float64(6.0)}, scene)
		box5 := mb.CreateBox("Box5", &babylon.BoxOpts{Size: Float64(6.0)}, scene)
		box6 := mb.CreateBox("Box6", &babylon.BoxOpts{Size: Float64(6.0)}, scene)
		box7 := mb.CreateBox("Box7", &babylon.BoxOpts{Size: Float64(6.0)}, scene)

		//Moving boxes on the x axis
		box1.Position().SetX(-20)
		box2.Position().SetX(-10)
		box3.Position().SetX(0)
		box4.Position().SetX(15)
		box5.Position().SetX(30)
		box6.Position().SetX(45)

		//Rotate box around the x axis
		box1.Rotation().SetX(math.Pi / 6)

		//Rotate box around the y axis
		box2.Rotation().SetY(math.Pi / 3)

		//Scaling on the x axis
		box4.Scaling().SetX(2)

		//Scaling on the y axis
		box5.Scaling().SetY(2)

		//Scaling on the z axis
		box6.Scaling().SetZ(2)

		//Moving box7 relatively to box1
		box7.SetParent(box1.Node)
		box7.Position().SetZ(-10)

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
