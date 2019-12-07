// 32-six-intersecting-point-lights is a port to Go of the Babylon example located here:
// https://www.babylonjs-playground.com/#IRVAX#0
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

		// Set up camera
		camera := b.NewArcRotateCamera("Camera", 0, 0, 10, b.Vector3().Zero(), scene, nil)
		camera.SetPosition(b.NewVector3(-10, 10, 0))
		camera.AttachControl(canvas, &babylon.ArcRotateCameraAttachControlOpts{NoPreventDefault: Bool(true)})

		// Lights
		light0 := b.NewPointLight("Omni0", b.NewVector3(0, 10, 0), scene)
		light1 := b.NewPointLight("Omni1", b.NewVector3(0, -10, 0), scene)
		light2 := b.NewPointLight("Omni2", b.NewVector3(10, 0, 0), scene)
		light3 := b.NewDirectionalLight("Dir0", b.NewVector3(1, -1, 0), scene)
		light4 := b.NewPointLight("Omni3", b.NewVector3(10, 0, 0), scene)
		light5 := b.NewPointLight("Omni4", b.NewVector3(10, 0, 0), scene)

		material := b.NewStandardMaterial("kosh", scene)
		mb := b.Mesh()
		sphereOpts := &babylon.MeshCreateSphereOpts{Scene: scene}
		sphere := mb.CreateSphere("Sphere", 16, 3, sphereOpts)

		// Creating light sphere
		lightSphere0 := mb.CreateSphere("Sphere0", 16, 0.5, sphereOpts)
		lightSphere1 := mb.CreateSphere("Sphere1", 16, 0.5, sphereOpts)
		lightSphere2 := mb.CreateSphere("Sphere2", 16, 0.5, sphereOpts)
		lightSphere4 := mb.CreateSphere("Sphere4", 16, 0.5, sphereOpts)
		lightSphere5 := mb.CreateSphere("Sphere5", 16, 0.5, sphereOpts)

		redMat := b.NewStandardMaterial("red", scene)
		lightSphere0.SetMaterial(redMat.Material)
		redMat.SetDiffuseColor(b.NewColor3(0, 0, 0))
		redMat.SetSpecularColor(b.NewColor3(0, 0, 0))
		redMat.SetEmissiveColor(b.NewColor3(1, 0, 0))

		greenMat := b.NewStandardMaterial("green", scene)
		lightSphere1.SetMaterial(greenMat.Material)
		greenMat.SetDiffuseColor(b.NewColor3(0, 0, 0))
		greenMat.SetSpecularColor(b.NewColor3(0, 0, 0))
		greenMat.SetEmissiveColor(b.NewColor3(0, 1, 0))

		blueMat1 := b.NewStandardMaterial("blue", scene)
		lightSphere2.SetMaterial(blueMat1.Material)
		blueMat1.SetDiffuseColor(b.NewColor3(0, 0, 0))
		blueMat1.SetSpecularColor(b.NewColor3(0, 0, 0))
		blueMat1.SetEmissiveColor(b.NewColor3(0, 0, 1))

		blueMat2 := b.NewStandardMaterial("blue", scene)
		lightSphere4.SetMaterial(blueMat2.Material)
		blueMat2.SetDiffuseColor(b.NewColor3(0, 0, 0))
		blueMat2.SetSpecularColor(b.NewColor3(0, 0, 0))
		blueMat2.SetEmissiveColor(b.NewColor3(1, 1, 0))

		blueMat3 := b.NewStandardMaterial("blue", scene)
		lightSphere5.SetMaterial(blueMat3.Material)
		blueMat3.SetDiffuseColor(b.NewColor3(0, 0, 0))
		blueMat3.SetSpecularColor(b.NewColor3(0, 0, 0))
		blueMat3.SetEmissiveColor(b.NewColor3(0, 1, 1))

		// Sphere material
		material.SetDiffuseColor(b.NewColor3(1, 1, 1))
		sphere.SetMaterial(material.Material)
		material.SetMaxSimultaneousLights(16)

		// Lights colors
		light0.SetDiffuse(b.NewColor3(1, 0, 0))
		light0.SetSpecular(b.NewColor3(1, 0, 0))

		light1.SetDiffuse(b.NewColor3(0, 1, 0))
		light1.SetSpecular(b.NewColor3(0, 1, 0))

		light2.SetDiffuse(b.NewColor3(0, 0, 1))
		light2.SetSpecular(b.NewColor3(0, 0, 1))

		light3.SetDiffuse(b.NewColor3(1, 1, 1))
		light3.SetSpecular(b.NewColor3(1, 1, 1))

		light4.SetDiffuse(b.NewColor3(1, 1, 0))
		light4.SetSpecular(b.NewColor3(1, 1, 0))

		light5.SetDiffuse(b.NewColor3(0, 1, 1))
		light5.SetSpecular(b.NewColor3(0, 1, 1))

		// Animations
		alpha := 0.0
		scene.SetBeforeRender(func() {
			light0.SetPosition(b.NewVector3(10*math.Sin(alpha), 0, 10*math.Cos(alpha)))
			light1.SetPosition(b.NewVector3(10*math.Sin(alpha), 0, -10*math.Cos(alpha)))
			light2.SetPosition(b.NewVector3(10*math.Cos(alpha), 0, 10*math.Sin(alpha)))
			light4.SetPosition(b.NewVector3(10*math.Cos(alpha), 10*math.Sin(alpha), 0))
			light5.SetPosition(b.NewVector3(10*math.Sin(alpha), -10*math.Cos(alpha), 0))

			lightSphere0.SetPosition(light0.Position())
			lightSphere1.SetPosition(light1.Position())
			lightSphere2.SetPosition(light2.Position())
			lightSphere4.SetPosition(light4.Position())
			lightSphere5.SetPosition(light5.Position())

			alpha += 0.01
		})

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
