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
			skeleton := babylon.SkeletonFromJSObject(skeletons.Index(0), this)

			shadowGenerator.AddShadowCaster(scene.Meshes()[0], &babylon.ShadowGeneratorAddShadowCasterOpts{IncludeDescendants: Bool(true)})
			for index := 0; index < newMeshes.Length(); index++ {
				newMeshes.Index(index).Set("receiveShadows", false)
			}

			helper := scene.CreateDefaultEnvironment(b.NewIEnvironmentHelperOptions().SetEnableGroundShadow(true))
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
			if idleRange != nil {
				scene.BeginAnimation(skeleton, idleRange.From(), idleRange.To(), &babylon.SceneBeginAnimationOpts{Loop: Bool(true)})
			}

			// UI
			gui := b.GUI()
			advancedTexture := gui.AdvancedDynamicTexture().CreateFullscreenUI("UI", nil)
			UIPanel := gui.NewStackPanel(nil)
			UIPanel.SetWidth("220px")
			UIPanel.SetFontSize("14px")
			UIPanel.SetHorizontalAlignment(gui.Control().HORIZONTAL_ALIGNMENT_RIGHT())
			UIPanel.SetVerticalAlignment(gui.Control().VERTICAL_ALIGNMENT_CENTER())
			advancedTexture.AddControl(UIPanel.Control)
			// ..
			button1 := gui.Button().CreateSimpleButton("but1", "Play Idle")
			button1.SetPaddingTop("10px")
			button1.SetWidth("100px")
			button1.SetHeight("50px")
			button1.SetColor("white")
			button1.SetBackground("green")
			button1.OnPointerDownObservable().Add(func(this js.Value, args []js.Value) interface{} {
				if idleRange != nil {
					scene.BeginAnimation(skeleton, idleRange.From(), idleRange.To(), &babylon.SceneBeginAnimationOpts{Loop: Bool(true)})
				}
				return nil
			}, nil)
			UIPanel.AddControl(button1.Control)
			// ..
			button2 := gui.Button().CreateSimpleButton("but2", "Play Walk")
			button2.SetPaddingTop("10px")
			button2.SetWidth("100px")
			button2.SetHeight("50px")
			button2.SetColor("white")
			button2.SetBackground("green")
			button2.OnPointerDownObservable().Add(func(this js.Value, args []js.Value) interface{} {
				if walkRange != nil {
					scene.BeginAnimation(skeleton, walkRange.From(), walkRange.To(), &babylon.SceneBeginAnimationOpts{Loop: Bool(true)})
				}
				return nil
			}, nil)
			UIPanel.AddControl(button2.Control)
			// ..
			button3 := gui.Button().CreateSimpleButton("but3", "Play Run")
			button3.SetPaddingTop("10px")
			button3.SetWidth("100px")
			button3.SetHeight("50px")
			button3.SetColor("white")
			button3.SetBackground("green")
			button3.OnPointerDownObservable().Add(func(this js.Value, args []js.Value) interface{} {
				if runRange != nil {
					scene.BeginAnimation(skeleton, runRange.From(), runRange.To(), &babylon.SceneBeginAnimationOpts{Loop: Bool(true)})
				}
				return nil
			}, nil)
			UIPanel.AddControl(button3.Control)
			// ..
			button4 := gui.Button().CreateSimpleButton("but4", "Play Left")
			button4.SetPaddingTop("10px")
			button4.SetWidth("100px")
			button4.SetHeight("50px")
			button4.SetColor("white")
			button4.SetBackground("green")
			button4.OnPointerDownObservable().Add(func(this js.Value, args []js.Value) interface{} {
				if leftRange != nil {
					scene.BeginAnimation(skeleton, leftRange.From(), leftRange.To(), &babylon.SceneBeginAnimationOpts{Loop: Bool(true)})
				}
				return nil
			}, nil)
			UIPanel.AddControl(button4.Control)
			// ..
			button5 := gui.Button().CreateSimpleButton("but5", "Play Right")
			button5.SetPaddingTop("10px")
			button5.SetWidth("100px")
			button5.SetHeight("50px")
			button5.SetColor("white")
			button5.SetBackground("green")
			button5.OnPointerDownObservable().Add(func(this js.Value, args []js.Value) interface{} {
				if rightRange != nil {
					scene.BeginAnimation(skeleton, rightRange.From(), rightRange.To(), &babylon.SceneBeginAnimationOpts{Loop: Bool(true)})
				}
				return nil
			}, nil)
			UIPanel.AddControl(button5.Control)
			// ..
			button6 := gui.Button().CreateSimpleButton("but6", "Play Blend")
			button6.SetPaddingTop("10px")
			button6.SetWidth("100px")
			button6.SetHeight("50px")
			button6.SetColor("white")
			button6.SetBackground("green")
			button6.OnPointerDownObservable().Add(func(this js.Value, args []js.Value) interface{} {
				if walkRange != nil && leftRange != nil {
					scene.StopAnimation(skeleton, nil)
					walkAnim := scene.BeginWeightedAnimation(skeleton, walkRange.From(), walkRange.To(), 0.5, &babylon.SceneBeginWeightedAnimationOpts{Loop: Bool(true)})
					leftAnim := scene.BeginWeightedAnimation(skeleton, leftRange.From(), leftRange.To(), 0.5, &babylon.SceneBeginWeightedAnimationOpts{Loop: Bool(true)})

					// Note: Sync Speed Ratio With Master Walk Animation
					walkAnim.SyncWith(nil)
					leftAnim.SyncWith(walkAnim)
				}
				return nil
			}, nil)
			UIPanel.AddControl(button6.Control)

			engine.HideLoadingUI()
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

// String returns the pointer to the provided string.
func String(v string) *string {
	return &v
}
