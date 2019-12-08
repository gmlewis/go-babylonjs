// 12-position is a port to Go of the Babylon example located here:
// https://www.babylonjs-playground.com/#1ZMJQV#2
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
		scene.SetClearColor(b.NewColor4(.5, .5, .5, 1))

		// camera
		camera := b.NewArcRotateCamera("camera1", 0, 0, 0, b.NewVector3(5, 3, 0), scene, nil)
		camera.SetPosition(b.NewVector3(10.253, 5.82251, -9.45717))
		camera.AttachControl(canvas, true, nil)

		// lights
		light := b.NewHemisphericLight("light1", b.NewVector3(1, 0.5, 0), scene)
		light.SetIntensity(0.8)

		// material
		mat := b.NewStandardMaterial("mat1", scene)
		mat.SetAlpha(1.0)
		mat.SetDiffuseColor(b.NewColor3(0.5, 0.5, 1.0))
		mat.SetBackFaceCulling(false)

		mb := b.MeshBuilder()

		//Local Axes
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
			pilotLocalAxisX := mb.CreateLines("pilotLocalAxisX", xOpts, scene)
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
			pilotLocalAxisY := mb.CreateLines("pilotLocalAxisY", yOpts, scene)
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
			pilotLocalAxisZ := mb.CreateLines("pilotLocalAxisZ", zOpts, scene)
			pilotLocalAxisZ.SetColor(b.NewColor3(0, 0, 1))

			localOrigin := mb.CreateBox("localOrigin", &babylon.BoxOpts{Size: Float64(1)}, scene)
			localOrigin.SetIsVisible(false)

			pilotLocalAxisX.SetParent(localOrigin.Node)
			pilotLocalAxisY.SetParent(localOrigin.Node)
			pilotLocalAxisZ.SetParent(localOrigin.Node)

			return localOrigin
		}

		/************Creations of Pilot*********************************/
		pilot := mb.CreateCylinder("pilot", &babylon.CylinderOpts{
			Height:         Float64(0.75),
			DiameterTop:    Float64(0.2),
			DiameterBottom: Float64(0.5),
			Tessellation:   Float64(6),
			Subdivisions:   Float64(1),
		}, scene)
		greyMat := b.NewStandardMaterial("grey", scene)
		greyMat.SetEmissiveColor(b.NewColor3(0.2, 0.2, 0.2))
		pilot.SetMaterial(greyMat.Material)

		arm := mb.CreateBox("arm", &babylon.BoxOpts{Height: Float64(0.75), Width: Float64(0.3), Depth: Float64(0.1875)}, scene)
		arm.SetMaterial(greyMat.Material)
		arm.Position().SetX(0.125)
		arm.SetParent(pilot.Node)
		/****************************************************************/

		localOrigin := localAxes(2)

		/********************Placement of Pilot and Local Axes***********/
		pilot.SetPosition(b.NewVector3(2, 3, 4))
		localOrigin.SetPosition(pilot.Position())

		// rotation order
		pilot.Rotation().SetZ(math.Pi / 2)
		pilot.Rotation().SetX(math.Pi / 2)
		pilot.Rotation().SetY(math.Pi / 2)

		/****************************************************************/

		// show axis
		showAxis := func(size float64) {
			makeTextPlane := func(text, color string, size float64) *babylon.Mesh {
				dynamicTexture := b.NewDynamicTexture("DynamicTexture", 50, scene, true, nil)
				dynamicTexture.SetHasAlpha(true)
				dynamicTexture.DrawText(text, 5, 40, "bold 36px Arial", color, "transparent", &babylon.DynamicTextureDrawTextOpts{InvertY: Bool(true)})
				plane := mb.CreatePlane("TextPlane", &babylon.PlaneOpts{Size: Float64(size)}, scene)
				plane.SetMaterial(b.NewStandardMaterial("TextPlaneMaterial", scene).Material)
				plane.Material().SetBackFaceCulling(false)
				plane.Material().SetSpecularColor(b.NewColor3(0, 0, 0))
				plane.Material().SetDiffuseTexture(dynamicTexture.BaseTexture)
				return plane
			}

			xOpts := &babylon.LinesOpts{
				Points: []*babylon.Vector3{
					b.Vector3().Zero(),
					b.NewVector3(size, 0, 0),
					b.NewVector3(size*0.95, 0.05*size, 0),
					b.NewVector3(size, 0, 0),
					b.NewVector3(size*0.95, -0.05*size, 0),
				},
			}
			axisX := mb.CreateLines("axisX", xOpts, scene)
			axisX.SetColor(b.NewColor3(1, 0, 0))
			xChar := makeTextPlane("X", "red", size/10)
			xChar.SetPosition(b.NewVector3(0.9*size, -0.05*size, 0))

			yOpts := &babylon.LinesOpts{
				Points: []*babylon.Vector3{
					b.Vector3().Zero(),
					b.NewVector3(0, size, 0),
					b.NewVector3(-0.05*size, size*0.95, 0),
					b.NewVector3(0, size, 0),
					b.NewVector3(0.05*size, size*0.95, 0),
				},
			}
			axisY := mb.CreateLines("axisY", yOpts, scene)
			axisY.SetColor(b.NewColor3(0, 1, 0))
			yChar := makeTextPlane("Y", "green", size/10)
			yChar.SetPosition(b.NewVector3(0, 0.9*size, -0.05*size))

			zOpts := &babylon.LinesOpts{
				Points: []*babylon.Vector3{
					b.Vector3().Zero(),
					b.NewVector3(0, 0, size),
					b.NewVector3(0, -0.05*size, size*0.95),
					b.NewVector3(0, 0, size),
					b.NewVector3(0, 0.05*size, size*0.95),
				},
			}
			axisZ := mb.CreateLines("axisZ", zOpts, scene)
			axisZ.SetColor(b.NewColor3(0, 0, 1))
			zChar := makeTextPlane("Z", "blue", size/10)
			zChar.SetPosition(b.NewVector3(0, 0.05*size, 0.9*size))
		}

		showAxis(8)

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
