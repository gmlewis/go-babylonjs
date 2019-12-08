// 23-follow-camera is a port to Go of the Babylon example located here:
// https://www.babylonjs-playground.com/#SRZRWV#6
// and linked from here:
// https://doc.babylonjs.com/babylon101/cameras
package main

import (
	"math"
	"math/rand"
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

		/********** FOLLOW CAMERA EXAMPLE **************************/

		// This creates and initially positions a follow camera
		camera := b.NewFollowCamera("FollowCam", b.NewVector3(0, 10, -10), scene, nil)

		//The goal distance of camera from target
		camera.SetRadius(30)

		// The goal height of camera above local origin (centre) of target
		camera.SetHeightOffset(10)

		// The goal rotation of camera around local origin (centre) of target in x y plane
		camera.SetRotationOffset(0)

		//Acceleration of camera in moving from current to goal position
		camera.SetCameraAcceleration(0.005)

		//The speed at which acceleration is halted
		camera.SetMaxCameraSpeed(10)

		//camera.target is set after the target's creation

		// This attaches the camera to the canvas
		camera.AttachControl(canvas, true)

		/**************************************************************/

		// This creates a light, aiming 0,1,0 - to the sky (non-mesh)
		b.NewHemisphericLight("light", b.NewVector3(0, 1, 0), scene)

		//Material
		mat := b.NewStandardMaterial("mat1", scene)
		mat.SetAlpha(1.0)
		mat.SetDiffuseColor(b.NewColor3(0.5, 0.5, 1.0))
		texture := b.NewTexture("https://i.imgur.com/vxH5bCg.jpg", scene, nil)
		mat.SetDiffuseTexture(texture.BaseTexture)

		//Different face for each side of box to show camera rotation
		hSpriteNb := 3.0 // 3 sprites per row
		vSpriteNb := 2.0 // 2 sprite rows

		faceUV := []*babylon.Vector4{}

		for i := 0.0; i < 6.0; i++ {
			faceUV = append(faceUV, b.NewVector4(i/hSpriteNb, 0, (i+1)/hSpriteNb, 1/vSpriteNb))
		}

		// Shape to follow
		box := b.MeshBuilder().CreateBox("box", &babylon.BoxOpts{Size: Float64(2), FaceUV: faceUV}, scene)
		box.SetPosition(b.NewVector3(20, 0, 10))
		box.SetMaterial(mat.Material)

		//create solid particle system of stationery grey boxes to show movement of box and camera
		boxesSPS := b.NewSolidParticleSystem("boxes", scene, &babylon.NewSolidParticleSystemOpts{Options: map[string]interface{}{"updatable": false}})

		//function to position of grey boxes
		setBoxes := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			particle := babylon.MeshFromJSObject(args[0], this)
			particle.SetPosition(b.NewVector3(-50+rand.Float64()*100, -50+rand.Float64()*100, -50+rand.Float64()*100))
			return nil
		})

		//add 400 boxes
		boxesSPS.AddShape(box, 400, &babylon.SolidParticleSystemAddShapeOpts{Options: map[string]interface{}{"positionFunction": setBoxes}})
		boxesSPS.BuildMesh() // mesh of boxes

		/*****************SET TARGET FOR CAMERA************************/
		camera.SetLockedTarget(box.AbstractMesh)
		/**************************************************************/

		//box movement variables
		alpha := 0.0
		orbitRadius := 20.0

		//Move the box to see that the camera follows it
		scene.RegisterBeforeRender(func() {
			alpha += 0.01
			box.Position().SetX(orbitRadius * math.Cos(alpha))
			box.Position().SetY(orbitRadius * math.Sin(alpha))
			box.Position().SetZ(10 * math.Sin(2*alpha))

			//change the viewing angle of the camera as it follows the box
			camera.SetRotationOffset(math.Mod(18*alpha, 360))
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
