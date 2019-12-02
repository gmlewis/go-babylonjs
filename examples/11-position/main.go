// 11-position is a port to Go of the Babylon example located here:
// https://www.babylonjs-playground.com/#UBWFJT#2
// and linked from here:
// https://doc.babylonjs.com/babylon101/position
package main

import (
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
		scene.SetClearColor(b.NewColor4(0.5, 0.5, 0.5, 1))

		// camera
		camera := b.NewArcRotateCamera("camera1", 0, 0, 0, b.NewVector3(2, 3, 4), scene, nil)
		camera.SetPosition(b.NewVector3(10, 3, -10))
		camera.AttachControl(canvas, &babylon.ArcRotateCameraAttachControlOpts{NoPreventDefault: babylon.Bool(true)})
		// lights
		light := b.NewHemisphericLight("light1", b.NewVector3(1, 0.5, 0), scene)
		light.SetIntensity(0.8)

		/*******************************Local Axes****************************/
		localAxes := func(size float64) *babylon.Mesh {
			xOpts := &babylon.LinesOpts{
				Points: []*babylon.Vector3{
					b.Vector3().Zero(),
					b.NewVector3(size, 0, 0),
					b.NewVector3(size*0.95, 0.05*size, 0),
					b.NewVector3(size, 0, 0),
					b.NewVector3(size*0.95, -0.05*size, 0),
				},
			}
			pilotLocalAxisX := b.CreateLines("pilotLocalAxisX", xOpts, scene)
			pilotLocalAxisX.SetColor(b.NewColor3(1, 0, 0))

			yOpts := &babylon.LinesOpts{
				Points: []*babylon.Vector3{
					b.Vector3().Zero(),
					b.NewVector3(0, size, 0),
					b.NewVector3(-0.05*size, size*0.95, 0),
					b.NewVector3(0, size, 0),
					b.NewVector3(0.05*size, size*0.95, 0),
				},
			}
			pilotLocalAxisY := b.CreateLines("pilotLocalAxisY", yOpts, scene)
			pilotLocalAxisY.SetColor(b.NewColor3(0, 1, 0))

			zOpts := &babylon.LinesOpts{
				Points: []*babylon.Vector3{
					b.Vector3().Zero(),
					b.NewVector3(0, 0, size),
					b.NewVector3(0, -0.05*size, size*0.95),
					b.NewVector3(0, 0, size),
					b.NewVector3(0, 0.05*size, size*0.95),
				},
			}
			pilotLocalAxisZ := b.CreateLines("pilotLocalAxisZ", zOpts, scene)
			pilotLocalAxisZ.SetColor(b.NewColor3(0, 0, 1))

			localOrigin := b.CreateBox("local_origin", &babylon.BoxOpts{Size: Float64(1)}, scene)
			localOrigin.SetIsVisible(false)

			pilotLocalAxisX.SetParent(localOrigin)
			pilotLocalAxisY.SetParent(localOrigin)
			pilotLocalAxisZ.SetParent(localOrigin)

			return localOrigin
		}
		/*******************************End Local Axes****************************/

		/************Start Pilot*********************************/
		body := b.CreateCylinder("body", &babylon.CylinderOpts{Height: Float64(0.75), DiameterTop: Float64(0.2), DiameterBottom: Float64(0.5), Tessellation: Float64(6), Subdivisions: Float64(1)}, scene)
		arm := b.CreateBox("arm", &babylon.BoxOpts{Height: Float64(0.75), Width: Float64(0.3), Depth: Float64(0.1875)}, scene)
		arm.Position().SetX(0.125)
		pilot := b.Mesh().MergeMeshes([]*babylon.Mesh{body, arm}, &babylon.MergeMeshesOpts{DisposeSource: Bool(true)})

		localOrigin := localAxes(2)
		localOrigin.parent = pilot
		/*************End Pilot****************************************/

		//#####################BABYLON 101 DEMO CODE POSITION###################

		pilot.SetPosition(b.NewVector3(2, 3, 4))

		//#############################################################

		/*********************************Start World Axes********************/
		showAxis := func(size float64) {
			makeTextPlane := func(text string, color *babylon.Color4, size float64) *babylon.Plane {
				dynamicTexture := b.NewDynamicTexture("DynamicTexture", 50, scene, true)
				dynamicTexture.SetHasAlpha(true)
				dynamicTexture.DrawText(text, 5, 40, "bold 36px Arial", color, "transparent", true)
				plane := b.CreatePlane("TextPlane", size, scene, true)
				plane.SetMaterial(b.NewStandardMaterial("TextPlaneMaterial", scene))
				plane.Material.SetBackFaceCulling(false)
				plane.Material.SetSpecularColor(b.NewColor3(0, 0, 0))
				plane.Material.SetDiffuseTexture(dynamicTexture)
				return plane
			}

			axisX := b.CreateLines("axisX", []*babylon.Vector3{
				b.Vector3().Zero(), b.NewVector3(size, 0, 0), b.NewVector3(size*0.95, 0.05*size, 0),
				b.NewVector3(size, 0, 0), b.NewVector3(size*0.95, -0.05*size, 0),
			}, scene)
			axisX.SetColor(b.NewColor3(1, 0, 0))
			xChar := makeTextPlane("X", "red", size/10)
			xChar.SetPosition(b.NewVector3(0.9*size, -0.05*size, 0))
			axisY := b.CreateLines("axisY", []*babylon.Vector3{
				b.Vector3().Zero(), b.NewVector3(0, size, 0), b.NewVector3(-0.05*size, size*0.95, 0),
				b.NewVector3(0, size, 0), b.NewVector3(0.05*size, size*0.95, 0),
			}, scene)
			axisY.SetColor(b.NewColor3(0, 1, 0))
			yChar := makeTextPlane("Y", "green", size/10)
			yChar.SetPosition(b.NewVector3(0, 0.9*size, -0.05*size))
			axisZ := b.CreateLines("axisZ", []*babylon.Vector3{
				b.Vector3().Zero(), b.NewVector3(0, 0, size), b.NewVector3(0, -0.05*size, size*0.95),
				b.NewVector3(0, 0, size), b.NewVector3(0, 0.05*size, size*0.95),
			}, scene)
			axisZ.SetColor(b.NewColor3(0, 0, 1))
			zChar := makeTextPlane("Z", "blue", size/10)
			zChar.SetPosition(b.NewVector3(0, 0.05*size, 0.9*size))
		}
		/***************************End World Axes***************************/

		showAxis(8)

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
