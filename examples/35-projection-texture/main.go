// 35-projection-texture is a port to Go of the Babylon example located here:
// https://www.babylonjs-playground.com/#CQNGRK
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

		// Setup environment
		camera := b.NewArcRotateCamera("Camera", 0, 0.8, 90, b.Vector3().Zero(), scene, nil)
		camera.SetLowerBetaLimit(0.1)
		camera.SetUpperBetaLimit((math.Pi / 2) * 0.9)
		camera.SetLowerRadiusLimit(30)
		camera.SetUpperRadiusLimit(150)
		camera.AttachControl(canvas, true, nil)

		light := b.NewHemisphericLight("dir01", b.NewVector3(0, 1, 0), scene)
		light.SetIntensity(0.1)

		// spot
		spotLight := b.NewSpotLight("spot02", b.NewVector3(30, 40, 30),
			b.NewVector3(-1, -2, -1), 1.1, 16, scene)
		spotLight.SetProjectionTexture(b.NewTexture("textures/co.png", scene, nil).BaseTexture)
		spotLight.SetDirectionToTarget(b.Vector3().Zero())
		spotLight.SetIntensity(1.5)

		// Ground
		ground := b.Mesh().CreateGroundFromHeightMap("ground", "textures/heightMap.png", 100, 100, 100, 0, 10, scene, &babylon.MeshCreateGroundFromHeightMapOpts{Updatable: Bool(false)})
		groundMaterial := b.NewStandardMaterial("ground", scene)
		diffuseTexture := b.NewTexture("textures/ground.jpg", scene, nil)
		groundMaterial.SetDiffuseTexture(diffuseTexture.BaseTexture)
		diffuseTexture.SetUScale(6)
		diffuseTexture.SetVScale(6)
		groundMaterial.SetSpecularColor(b.NewColor3(0, 0, 0))
		ground.Position().SetY(-2.05)
		ground.SetMaterial(groundMaterial.Material)

		// Animations
		alpha := 0.0
		scene.RegisterBeforeRender(func(this js.Value, args []js.Value) interface{} {
			spotLight.SetPosition(b.NewVector3(math.Cos(alpha)*60, 40, math.Sin(alpha)*60))
			spotLight.SetDirectionToTarget(b.Vector3().Zero())
			alpha += 0.01
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
