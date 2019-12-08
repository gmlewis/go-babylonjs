// 36-basic-animation is a port to Go of the Babylon example located here:
// https://www.babylonjs-playground.com/#DMLMIP#1
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
		engine.SetEnableOfflineSupport(false)

		// This is really important to tell Babylon.js to use decomposeLerp and matrix interpolation
		b.Animation().SetAllowMatricesInterpolation(true)

		scene := b.NewScene(engine, nil)

		camera := b.NewArcRotateCamera("camera1", math.Pi/2, math.Pi/4, 3, b.NewVector3(0, 1, 0), scene, nil)
		camera.AttachControl(canvas, true, nil)

		camera.SetLowerRadiusLimit(2)
		camera.SetUpperRadiusLimit(10)
		camera.SetWheelDeltaPercentage(0.01)

		light := b.NewHemisphericLight("light1", b.NewVector3(0, 1, 0), scene)
		light.SetIntensity(0.6)
		light.SetSpecular(b.Color3().Black())

		light2 := b.NewDirectionalLight("dir01", b.NewVector3(0, -0.5, -1.0), scene)
		light2.SetPosition(b.NewVector3(0, 5, 5))

		// Shadows
		shadowGenerator := b.NewShadowGenerator(1024, light2.IShadowLight(), nil)
		shadowGenerator.SetUseBlurExponentialShadowMap(true)
		shadowGenerator.SetBlurKernel(32)

		engine.DisplayLoadingUI()

		onSuccess := func(this js.Value, args []js.Value) interface{} {
			newMeshes, skeletons := args[0], args[2]
			skeleton := SkeletonFromJSObject(skeletons.Index(0), this)

			shadowGenerator.AddShadowCaster(scene.Meshes()[0], true)
			for index := 0; index < newMeshes.Length(); index++ {
				newMeshes.Index(index).Set("receiveShadows", false)
			}

			helper := scene.CreateDefaultEnvironment(&babylon.IEnvironmentHelperOptions{
				EnableGroundShadow: Bool(true),
			})
			helper.SetMainColor(b.Color3().Gray())
			groundPos := helper.Ground().Position()
			groundPos.SetY(groundPos.Y() + 0.01)

			// ROBOT
			skeleton.SetAnimationPropertiesOverride(b.NewAnimationPropertiesOverride())
			skeleton.AnimationPropertiesOverride().SetEnableBlending(true)
			skeleton.AnimationPropertiesOverride().SetBlendingSpeed(0.05)
			skeleton.AnimationPropertiesOverride().SetLoopMode(1)

			idleRange := skeleton.GetAnimationRange("YBot_Idle")
			walkRange := skeleton.GetAnimationRange("YBot_Walk")
			runRange := skeleton.GetAnimationRange("YBot_Run")
			leftRange := skeleton.GetAnimationRange("YBot_LeftStrafeWalk")
			rightRange := skeleton.GetAnimationRange("YBot_RightStrafeWalk")

			// IDLE
			if idleRange {
				scene.BeginAnimation(skeleton, idleRange.from, idleRange.to, true)
			}

			// UI
			advancedTexture := b.GUI().AdvancedDynamicTexture.CreateFullscreenUI("UI")
			UiPanel := b.NewGUI.StackPanel()
			UiPanel.width = "220px"
			UiPanel.fontSize = "14px"
			UiPanel.horizontalAlignment = b.GUI().Control.HORIZONTAL_ALIGNMENT_RIGHT
			UiPanel.verticalAlignment = b.GUI().Control.VERTICAL_ALIGNMENT_CENTER
			advancedTexture.addControl(UiPanel)
			// ..
			button := b.GUI().Button.CreateSimpleButton("but1", "Play Idle")
			button.paddingTop = "10px"
			button.width = "100px"
			button.height = "50px"
			button.color = "white"
			button.background = "green"
			button.OnPointerDownObservable().Add(func(this js.Value, args []js.Value) interface{} {
				if idleRange {
					scene.beginAnimation(skeleton, idleRange.from, idleRange.to, true)
				}
			})
			UiPanel.addControl(button)
			// ..
			button1 := b.GUI().Button.CreateSimpleButton("but2", "Play Walk")
			button1.paddingTop = "10px"
			button1.width = "100px"
			button1.height = "50px"
			button1.color = "white"
			button1.background = "green"
			button1.onPointerDownObservable.add(func(this js.Value, args []js.Value) interface{} {
				if walkRange {
					scene.beginAnimation(skeleton, walkRange.from, walkRange.to, true)
				}
			})
			UiPanel.addControl(button1)
			// ..
			button1 := b.GUI().Button.CreateSimpleButton("but3", "Play Run")
			button1.paddingTop = "10px"
			button1.width = "100px"
			button1.height = "50px"
			button1.color = "white"
			button1.background = "green"
			button1.onPointerDownObservable.add(func(this js.Value, args []js.Value) interface{} {
				if runRange {
					scene.beginAnimation(skeleton, runRange.from, runRange.to, true)
				}
			})
			UiPanel.addControl(button1)
			// ..
			button1 := b.GUI().Button.CreateSimpleButton("but4", "Play Left")
			button1.paddingTop = "10px"
			button1.width = "100px"
			button1.height = "50px"
			button1.color = "white"
			button1.background = "green"
			button1.onPointerDownObservable.add(func(this js.Value, args []js.Value) interface{} {
				if leftRange {
					scene.beginAnimation(skeleton, leftRange.from, leftRange.to, true)
				}
			})
			UiPanel.addControl(button1)
			// ..
			button1 := b.GUI().Button.CreateSimpleButton("but5", "Play Right")
			button1.paddingTop = "10px"
			button1.width = "100px"
			button1.height = "50px"
			button1.color = "white"
			button1.background = "green"
			button1.onPointerDownObservable.add(func(this js.Value, args []js.Value) interface{} {
				if rightRange {
					scene.beginAnimation(skeleton, rightRange.from, rightRange.to, true)
				}
			})
			UiPanel.addControl(button1)
			// ..
			button1 := b.GUI().Button.CreateSimpleButton("but6", "Play Blend")
			button1.paddingTop = "10px"
			button1.width = "100px"
			button1.height = "50px"
			button1.color = "white"
			button1.background = "green"
			button1.onPointerDownObservable.add(func(this js.Value, args []js.Value) interface{} {
				if walkRange && leftRange {
					scene.stopAnimation(skeleton)
					walkAnim := scene.beginWeightedAnimation(skeleton, walkRange.from, walkRange.to, 0.5, true)
					leftAnim := scene.beginWeightedAnimation(skeleton, leftRange.from, leftRange.to, 0.5, true)

					// Note: Sync Speed Ratio With Master Walk Animation
					walkAnim.syncWith(null)
					leftAnim.syncWith(walkAnim)
				}
			})
			UiPanel.addControl(button1)

			engine.hideLoadingUI()
			return nil
		}
		opts := &babylon.SceneLoaderImportMeshOpts{
			SceneFilename: String("dummy3.babylon"),
			Scene:         scene,
			OnSuccess:     onSuccess,
		}
		b.SceneLoader().ImportMesh("", "https://www.babylonjs-playground.com/scenes/", opts)

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
