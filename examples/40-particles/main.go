// 40-particles is a port to Go of the Babylon example located here:
// https://www.babylonjs-playground.com/#WBQ8EM
// and linked from here:
// https://doc.babylonjs.com/babylon101/particles
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

		// Set up environment
		b.NewPointLight("Omni", b.NewVector3(0, 2, 8), scene)
		camera := b.NewArcRotateCamera("ArcRotateCamera", 1, 0.8, 20, b.NewVector3(0, 0, 0), scene, nil)
		camera.AttachControl(canvas, true, nil)

		// Fountain object
		mb := b.Mesh()
		fountain := mb.CreateBox("foutain", 1.0, &babylon.MeshCreateBoxOpts{Scene: scene})

		// Ground
		ground := mb.CreatePlane("ground", 50.0, scene, nil)
		ground.SetPosition(b.NewVector3(0, -10, 0))
		ground.SetRotation(b.NewVector3(math.Pi/2, 0, 0))

		groundMat := b.NewStandardMaterial("groundMat", scene)
		ground.SetMaterial(groundMat.Material)
		groundMat.SetBackFaceCulling(false)
		groundMat.SetDiffuseColor(b.NewColor3(0.3, 0.3, 1))

		// Create a particle system
		particleSystem := b.NewParticleSystem("particles", 2000, scene, nil)

		//Texture of each particle
		particleSystem.SetParticleTexture(b.NewTexture("textures/flare.png", scene, nil))

		// Where the particles come from
		particleSystem.SetEmitter(fountain.AbstractMesh)     // the starting object, the emitter
		particleSystem.SetMinEmitBox(b.NewVector3(-1, 0, 0)) // Starting all from
		particleSystem.SetMaxEmitBox(b.NewVector3(1, 0, 0))  // To...

		// Colors of all particles
		particleSystem.SetColor1(b.NewColor4(0.7, 0.8, 1.0, 1.0))
		particleSystem.SetColor2(b.NewColor4(0.2, 0.5, 1.0, 1.0))
		particleSystem.SetColorDead(b.NewColor4(0, 0, 0.2, 0.0))

		// Size of each particle (random between...
		particleSystem.SetMinSize(0.1)
		particleSystem.SetMaxSize(0.5)

		// Life time of each particle (random between...
		particleSystem.SetMinLifeTime(0.3)
		particleSystem.SetMaxLifeTime(1.5)

		// Emission rate
		particleSystem.SetEmitRate(1500)

		// Blend mode : BLENDMODE_ONEONE, or BLENDMODE_STANDARD
		particleSystem.SetBlendMode(b.ParticleSystem().BLENDMODE_ONEONE())

		// Set the gravity of all particles
		particleSystem.SetGravity(b.NewVector3(0, -9.81, 0))

		// Direction of each particle after it has been emitted
		particleSystem.SetDirection1(b.NewVector3(-7, 8, 3))
		particleSystem.SetDirection2(b.NewVector3(7, 8, -3))

		// Angular speed, in radians
		particleSystem.SetMinAngularSpeed(0)
		particleSystem.SetMaxAngularSpeed(math.Pi)

		// Speed
		particleSystem.SetMinEmitPower(1)
		particleSystem.SetMaxEmitPower(3)
		particleSystem.SetUpdateSpeed(0.005)

		// Start the particle system
		particleSystem.Start(nil)

		// Fountain's animation
		var keys []*babylon.IAnimationKey
		animation := b.NewAnimation("animation", "rotation.x", 30, b.Animation().ANIMATIONTYPE_FLOAT(),
			&babylon.NewAnimationOpts{LoopMode: Float64(b.Animation().ANIMATIONLOOPMODE_CYCLE())})
		// At the animation key 0, the X value of rotation is "0"
		keys = append(keys, b.NewIAnimationKey(0, 0, nil))

		// At the animation key 50, the X value of rotation is Pi
		keys = append(keys, b.NewIAnimationKey(50, math.Pi, nil))

		// At the animation key 100, the X value of rotation is "0"
		keys = append(keys, b.NewIAnimationKey(100, 0, nil))

		// Launch animation
		animation.SetKeys(keys)
		fountain.Animations().Push(animation)
		scene.BeginAnimation(fountain, 0, 100, &babylon.SceneBeginAnimationOpts{Loop: Bool(true)})

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
