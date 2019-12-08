// 14-material-color-reaction is a port to Go of the Babylon example located here:
// https://www.babylonjs-playground.com/#20OAV9#325
// and linked from here:
// https://doc.babylonjs.com/babylon101/materials
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
		camera := b.NewArcRotateCamera("Camera", -math.Pi/2, math.Pi/3, 10, b.Vector3().Zero(), scene, nil)
		camera.AttachControl(canvas, true, nil)

		mats := []*babylon.Color3{
			b.NewColor3(1, 1, 0),
			b.NewColor3(1, 0, 1),
			b.NewColor3(0, 1, 1),
			b.NewColor3(1, 1, 1),
		}

		redMat := b.NewStandardMaterial("redMat", scene)
		redMat.SetEmissiveColor(b.NewColor3(1, 0, 0))

		greenMat := b.NewStandardMaterial("greenMat", scene)
		greenMat.SetEmissiveColor(b.NewColor3(0, 1, 0))

		blueMat := b.NewStandardMaterial("blueMat", scene)
		blueMat.SetEmissiveColor(b.NewColor3(0, 0, 1))

		whiteMat := b.NewStandardMaterial("whiteMat", scene)
		whiteMat.SetEmissiveColor(b.NewColor3(1, 1, 1))

		//red light
		lightRed := b.NewSpotLight("spotLight", b.NewVector3(-0.9, 1, -1.8), b.NewVector3(0, -1, 0), math.Pi/2, 1.5, scene)
		lightRed.SetDiffuse(b.NewColor3(1, 0, 0))
		lightRed.SetSpecular(b.NewColor3(0, 0, 0))

		//green light
		lightGreen := b.NewSpotLight("spotLight1", b.NewVector3(0, 1, -0.5), b.NewVector3(0, -1, 0), math.Pi/2, 1.5, scene)
		lightGreen.SetDiffuse(b.NewColor3(0, 1, 0))
		lightGreen.SetSpecular(b.NewColor3(0, 0, 0))

		//blue light
		lightBlue := b.NewSpotLight("spotLight2", b.NewVector3(0.9, 1, -1.8), b.NewVector3(0, -1, 0), math.Pi/2, 1.5, scene)
		lightBlue.SetDiffuse(b.NewColor3(0, 0, 1))
		lightBlue.SetSpecular(b.NewColor3(0, 0, 0))

		//white light
		lightWhite := b.NewSpotLight("spotLight3", b.NewVector3(0, 1, 1), b.NewVector3(0, -1, 0), math.Pi/2, 1.5, scene)
		lightWhite.SetDiffuse(b.NewColor3(1, 1, 1))
		lightWhite.SetSpecular(b.NewColor3(0, 0, 0))

		mb := b.MeshBuilder()

		redSphere := mb.CreateSphere("sphere", &babylon.SphereOpts{Diameter: Float64(0.25)}, scene)
		redSphere.SetMaterial(redMat.Material)
		redSphere.SetPosition(lightRed.Position())

		greenSphere := mb.CreateSphere("sphere", &babylon.SphereOpts{Diameter: Float64(0.25)}, scene)
		greenSphere.SetMaterial(greenMat.Material)
		greenSphere.SetPosition(lightGreen.Position())

		blueSphere := mb.CreateSphere("sphere", &babylon.SphereOpts{Diameter: Float64(0.25)}, scene)
		blueSphere.SetMaterial(blueMat.Material)
		blueSphere.SetPosition(lightBlue.Position())

		whiteSphere := mb.CreateSphere("sphere", &babylon.SphereOpts{Diameter: Float64(0.25)}, scene)
		whiteSphere.SetMaterial(whiteMat.Material)
		whiteSphere.SetPosition(lightWhite.Position())

		groundMat := b.NewStandardMaterial("groundMat", scene)
		groundMat.SetDiffuseColor(mats[0])

		ground := mb.CreateGround("ground", &babylon.GroundOpts{Width: Float64(4), Height: Float64(6)}, scene)
		ground.SetMaterial(groundMat.Material)

		/*******************GUI***********************/
		makeYellow := func(this js.Value, args []js.Value) interface{} {
			groundMat.SetDiffuseColor(mats[0])
			return nil
		}

		makePurple := func(this js.Value, args []js.Value) interface{} {
			groundMat.SetDiffuseColor(mats[1])
			return nil
		}

		makeCyan := func(this js.Value, args []js.Value) interface{} {
			groundMat.SetDiffuseColor(mats[2])
			return nil
		}

		makeWhite := func(this js.Value, args []js.Value) interface{} {
			groundMat.SetDiffuseColor(mats[3])
			return nil
		}

		gui := b.GUI()
		matGroup := gui.NewRadioGroup("Material Color")
		matGroup.AddRadio("Yellow", &babylon.RadioGroupAddRadioOpts{Func: makeYellow, Checked: Bool(true)})
		matGroup.AddRadio("Purple", &babylon.RadioGroupAddRadioOpts{Func: makePurple})
		matGroup.AddRadio("Cyan", &babylon.RadioGroupAddRadioOpts{Func: makeCyan})
		matGroup.AddRadio("White", &babylon.RadioGroupAddRadioOpts{Func: makeWhite})

		advancedTexture := gui.AdvancedDynamicTexture().CreateFullscreenUI("UI", nil)

		selectBox := gui.NewSelectionPanel("sp", &babylon.NewSelectionPanelOpts{Groups: []*babylon.SelectorGroup{matGroup.SelectorGroup}})
		selectBox.SetWidth("0.25")
		selectBox.SetHeight("50%")
		selectBox.SetTop("4px")
		selectBox.SetLeft("4px")
		selectBox.SetBackground("white")
		selectBox.SetHorizontalAlignment(gui.Control().HORIZONTAL_ALIGNMENT_LEFT())
		selectBox.SetVerticalAlignment(gui.Control().VERTICAL_ALIGNMENT_TOP())

		advancedTexture.AddControl(selectBox.Control)

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
