// 25-vr-device-orientation-camera is a port to Go of the Babylon example located here:
// https://www.babylonjs-playground.com/#SRZRWV#4
// and linked from here:
// https://doc.babylonjs.com/babylon101/cameras
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

		/********** DEVICE ORIENTATION CAMERA EXAMPLE **************************/

		// This creates and positions a device orientation camera
		camera := b.NewVRDeviceOrientationFreeCamera("DevOr_camera", b.NewVector3(0, 0, 0), scene, nil)

		// This targets the camera to scene origin
		camera.SetTarget(b.NewVector3(0, 0, 10))

		// This attaches the camera to the canvas
		camera.AttachControl(canvas, true)
		/**************************************************************/

		// This creates a light, aiming 0,1,0 - to the sky (non-mesh)
		b.NewHemisphericLight("light", b.NewVector3(0, 1, 0), scene)

		//Materials
		redMat := b.NewStandardMaterial("red", scene)
		redMat.SetDiffuseColor(b.NewColor3(1, 0, 0))
		redMat.SetEmissiveColor(b.NewColor3(1, 0, 0))
		redMat.SetSpecularColor(b.NewColor3(1, 0, 0))

		greenMat := b.NewStandardMaterial("green", scene)
		greenMat.SetDiffuseColor(b.NewColor3(0, 1, 0))
		greenMat.SetEmissiveColor(b.NewColor3(0, 1, 0))
		greenMat.SetSpecularColor(b.NewColor3(0, 1, 0))

		blueMat := b.NewStandardMaterial("blue", scene)
		blueMat.SetDiffuseColor(b.NewColor3(0, 0, 1))
		blueMat.SetEmissiveColor(b.NewColor3(0, 0, 1))
		blueMat.SetSpecularColor(b.NewColor3(0, 0, 1))

		// Shapes
		mb := b.Mesh()

		planeOpts := &babylon.MeshCreatePlaneOpts{Updatable: Bool(true), SideOrientation: Float64(mb.DOUBLESIDE())}

		plane1 := mb.CreatePlane("plane1", 3, scene, planeOpts)
		plane1.Position().SetX(-3)
		plane1.Position().SetZ(0)
		plane1.SetMaterial(redMat.Material)

		plane2 := mb.CreatePlane("plane2", 3, scene, planeOpts)
		plane2.Position().SetX(3)
		plane2.Position().SetZ(-1.5)
		plane2.SetMaterial(greenMat.Material)

		plane3 := mb.CreatePlane("plane3", 3, scene, planeOpts)
		plane3.Position().SetX(3)
		plane3.Position().SetZ(1.5)
		plane3.SetMaterial(blueMat.Material)

		mb.CreateGround("ground1", 10, 10, 2, &babylon.MeshCreateGroundOpts{Scene: scene})

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
