// 41-pre-warming-particles is a port to Go of the Babylon example located here:
// https://www.babylonjs-playground.com/#MX2Z99#8
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
		camera := b.NewArcRotateCamera("ArcRotateCamera", 1, 0.8, 5, b.NewVector3(0, 0, 0), scene, nil)
		camera.AttachControl(canvas, true, nil)
		scene.SetClearColor(b.NewColor4(0.0, 0.0, 0.0, 1.0))

		// Create a particle system
		surfaceParticles := b.NewParticleSystem("surfaceParticles", 1600, scene, nil)

		// Texture of each particle
		surfaceParticles.SetParticleTexture(b.NewTexture("https://raw.githubusercontent.com/PatrickRyanMS/BabylonJStextures/master/ParticleSystems/Sun/T_SunSurface.png", scene, nil))

		// Create core sphere
		coreSphere := b.MeshBuilder().CreateSphere("coreSphere", &babylon.SphereOpts{Diameter: Float64(2.01), Segments: Float64(64)}, scene)

		// Create core material
		coreMat := b.NewStandardMaterial("coreMat", scene)
		coreMat.SetEmissiveColor(b.NewColor3(0.3773, 0.0930, 0.0266))

		// Assign core material to sphere
		coreSphere.SetMaterial(coreMat.Material)

		// Pre-warm
		surfaceParticles.SetPreWarmStepOffset(10)
		surfaceParticles.SetPreWarmCycles(100)

		// Initial rotation
		surfaceParticles.SetMinInitialRotation(-2 * math.Pi)
		surfaceParticles.SetMaxInitialRotation(2 * math.Pi)

		// Where the sun particles come from
		sunEmitter := b.NewSphereParticleEmitter(nil)
		sunEmitter.SetRadius(1)
		sunEmitter.SetRadiusRange(0) // emit only from shape surface

		// Assign particles to emitters
		surfaceParticles.SetEmitter(coreSphere.AbstractMesh)                                                                   // the starting object, the emitter
		surfaceParticles.SetParticleEmitterType(babylon.IParticleEmitterTypeFromJSObject(sunEmitter.JSObject(), b.JSObject())) // TODO: Improve this.

		// Color gradient over time
		surfaceParticles.AddColorGradient(0, b.NewColor4(0.8509, 0.4784, 0.1019, 0.0), nil)
		surfaceParticles.AddColorGradient(0.4, b.NewColor4(0.6259, 0.3056, 0.0619, 0.5), nil)
		surfaceParticles.AddColorGradient(0.5, b.NewColor4(0.6039, 0.2887, 0.0579, 0.5), nil)
		surfaceParticles.AddColorGradient(1.0, b.NewColor4(0.3207, 0.0713, 0.0075, 0.0), nil)

		// Size of each particle (random between...
		surfaceParticles.SetMinSize(0.4)
		surfaceParticles.SetMaxSize(0.7)

		// Life time of each particle (random between...
		surfaceParticles.SetMinLifeTime(8.0)
		surfaceParticles.SetMaxLifeTime(8.0)

		// Emission rate
		surfaceParticles.SetEmitRate(200)

		// Blend mode : BLENDMODE_ONEONE, BLENDMODE_STANDARD, or BLENDMODE_ADD
		surfaceParticles.SetBlendMode(b.ParticleSystem().BLENDMODE_ADD())

		// Set the gravity of all particles
		surfaceParticles.SetGravity(b.NewVector3(0, 0, 0))

		// Angular speed, in radians
		surfaceParticles.SetMinAngularSpeed(-0.4)
		surfaceParticles.SetMaxAngularSpeed(0.4)

		// Speed
		surfaceParticles.SetMinEmitPower(0)
		surfaceParticles.SetMaxEmitPower(0)
		surfaceParticles.SetUpdateSpeed(0.005)

		// No billboard
		surfaceParticles.SetIsBillboardBased(false)

		// Start the particle system
		surfaceParticles.Start(nil)

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
