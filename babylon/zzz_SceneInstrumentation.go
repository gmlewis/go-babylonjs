// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SceneInstrumentation represents a babylon.js SceneInstrumentation.
// This class can be used to get instrumentation data from a Babylon engine
//
// See: http://doc.babylonjs.com/how_to/optimizing_your_scene#sceneinstrumentation
type SceneInstrumentation struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *SceneInstrumentation) JSObject() js.Value { return s.p }

// SceneInstrumentation returns a SceneInstrumentation JavaScript class.
func (ba *Babylon) SceneInstrumentation() *SceneInstrumentation {
	p := ba.ctx.Get("SceneInstrumentation")
	return SceneInstrumentationFromJSObject(p, ba.ctx)
}

// SceneInstrumentationFromJSObject returns a wrapped SceneInstrumentation JavaScript class.
func SceneInstrumentationFromJSObject(p js.Value, ctx js.Value) *SceneInstrumentation {
	return &SceneInstrumentation{p: p, ctx: ctx}
}

// SceneInstrumentationArrayToJSArray returns a JavaScript Array for the wrapped array.
func SceneInstrumentationArrayToJSArray(array []*SceneInstrumentation) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewSceneInstrumentation returns a new SceneInstrumentation object.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#constructor
func (ba *Babylon) NewSceneInstrumentation(scene *Scene) *SceneInstrumentation {

	args := make([]interface{}, 0, 1+0)

	args = append(args, scene.JSObject())

	p := ba.ctx.Get("SceneInstrumentation").New(args...)
	return SceneInstrumentationFromJSObject(p, ba.ctx)
}

// Dispose calls the Dispose method on the SceneInstrumentation object.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#dispose
func (s *SceneInstrumentation) Dispose() {

	s.p.Call("dispose")
}

// ActiveMeshesEvaluationTimeCounter returns the ActiveMeshesEvaluationTimeCounter property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#activemeshesevaluationtimecounter
func (s *SceneInstrumentation) ActiveMeshesEvaluationTimeCounter() *PerfCounter {
	retVal := s.p.Get("activeMeshesEvaluationTimeCounter")
	return PerfCounterFromJSObject(retVal, s.ctx)
}

// SetActiveMeshesEvaluationTimeCounter sets the ActiveMeshesEvaluationTimeCounter property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#activemeshesevaluationtimecounter
func (s *SceneInstrumentation) SetActiveMeshesEvaluationTimeCounter(activeMeshesEvaluationTimeCounter *PerfCounter) *SceneInstrumentation {
	s.p.Set("activeMeshesEvaluationTimeCounter", activeMeshesEvaluationTimeCounter.JSObject())
	return s
}

// AnimationsTimeCounter returns the AnimationsTimeCounter property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#animationstimecounter
func (s *SceneInstrumentation) AnimationsTimeCounter() *PerfCounter {
	retVal := s.p.Get("animationsTimeCounter")
	return PerfCounterFromJSObject(retVal, s.ctx)
}

// SetAnimationsTimeCounter sets the AnimationsTimeCounter property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#animationstimecounter
func (s *SceneInstrumentation) SetAnimationsTimeCounter(animationsTimeCounter *PerfCounter) *SceneInstrumentation {
	s.p.Set("animationsTimeCounter", animationsTimeCounter.JSObject())
	return s
}

// CameraRenderTimeCounter returns the CameraRenderTimeCounter property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#camerarendertimecounter
func (s *SceneInstrumentation) CameraRenderTimeCounter() *PerfCounter {
	retVal := s.p.Get("cameraRenderTimeCounter")
	return PerfCounterFromJSObject(retVal, s.ctx)
}

// SetCameraRenderTimeCounter sets the CameraRenderTimeCounter property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#camerarendertimecounter
func (s *SceneInstrumentation) SetCameraRenderTimeCounter(cameraRenderTimeCounter *PerfCounter) *SceneInstrumentation {
	s.p.Set("cameraRenderTimeCounter", cameraRenderTimeCounter.JSObject())
	return s
}

// CaptureActiveMeshesEvaluationTime returns the CaptureActiveMeshesEvaluationTime property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#captureactivemeshesevaluationtime
func (s *SceneInstrumentation) CaptureActiveMeshesEvaluationTime() bool {
	retVal := s.p.Get("captureActiveMeshesEvaluationTime")
	return retVal.Bool()
}

// SetCaptureActiveMeshesEvaluationTime sets the CaptureActiveMeshesEvaluationTime property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#captureactivemeshesevaluationtime
func (s *SceneInstrumentation) SetCaptureActiveMeshesEvaluationTime(captureActiveMeshesEvaluationTime bool) *SceneInstrumentation {
	s.p.Set("captureActiveMeshesEvaluationTime", captureActiveMeshesEvaluationTime)
	return s
}

// CaptureAnimationsTime returns the CaptureAnimationsTime property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#captureanimationstime
func (s *SceneInstrumentation) CaptureAnimationsTime() bool {
	retVal := s.p.Get("captureAnimationsTime")
	return retVal.Bool()
}

// SetCaptureAnimationsTime sets the CaptureAnimationsTime property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#captureanimationstime
func (s *SceneInstrumentation) SetCaptureAnimationsTime(captureAnimationsTime bool) *SceneInstrumentation {
	s.p.Set("captureAnimationsTime", captureAnimationsTime)
	return s
}

// CaptureCameraRenderTime returns the CaptureCameraRenderTime property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#capturecamerarendertime
func (s *SceneInstrumentation) CaptureCameraRenderTime() bool {
	retVal := s.p.Get("captureCameraRenderTime")
	return retVal.Bool()
}

// SetCaptureCameraRenderTime sets the CaptureCameraRenderTime property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#capturecamerarendertime
func (s *SceneInstrumentation) SetCaptureCameraRenderTime(captureCameraRenderTime bool) *SceneInstrumentation {
	s.p.Set("captureCameraRenderTime", captureCameraRenderTime)
	return s
}

// CaptureFrameTime returns the CaptureFrameTime property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#captureframetime
func (s *SceneInstrumentation) CaptureFrameTime() bool {
	retVal := s.p.Get("captureFrameTime")
	return retVal.Bool()
}

// SetCaptureFrameTime sets the CaptureFrameTime property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#captureframetime
func (s *SceneInstrumentation) SetCaptureFrameTime(captureFrameTime bool) *SceneInstrumentation {
	s.p.Set("captureFrameTime", captureFrameTime)
	return s
}

// CaptureInterFrameTime returns the CaptureInterFrameTime property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#captureinterframetime
func (s *SceneInstrumentation) CaptureInterFrameTime() bool {
	retVal := s.p.Get("captureInterFrameTime")
	return retVal.Bool()
}

// SetCaptureInterFrameTime sets the CaptureInterFrameTime property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#captureinterframetime
func (s *SceneInstrumentation) SetCaptureInterFrameTime(captureInterFrameTime bool) *SceneInstrumentation {
	s.p.Set("captureInterFrameTime", captureInterFrameTime)
	return s
}

// CaptureParticlesRenderTime returns the CaptureParticlesRenderTime property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#captureparticlesrendertime
func (s *SceneInstrumentation) CaptureParticlesRenderTime() bool {
	retVal := s.p.Get("captureParticlesRenderTime")
	return retVal.Bool()
}

// SetCaptureParticlesRenderTime sets the CaptureParticlesRenderTime property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#captureparticlesrendertime
func (s *SceneInstrumentation) SetCaptureParticlesRenderTime(captureParticlesRenderTime bool) *SceneInstrumentation {
	s.p.Set("captureParticlesRenderTime", captureParticlesRenderTime)
	return s
}

// CapturePhysicsTime returns the CapturePhysicsTime property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#capturephysicstime
func (s *SceneInstrumentation) CapturePhysicsTime() bool {
	retVal := s.p.Get("capturePhysicsTime")
	return retVal.Bool()
}

// SetCapturePhysicsTime sets the CapturePhysicsTime property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#capturephysicstime
func (s *SceneInstrumentation) SetCapturePhysicsTime(capturePhysicsTime bool) *SceneInstrumentation {
	s.p.Set("capturePhysicsTime", capturePhysicsTime)
	return s
}

// CaptureRenderTargetsRenderTime returns the CaptureRenderTargetsRenderTime property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#capturerendertargetsrendertime
func (s *SceneInstrumentation) CaptureRenderTargetsRenderTime() bool {
	retVal := s.p.Get("captureRenderTargetsRenderTime")
	return retVal.Bool()
}

// SetCaptureRenderTargetsRenderTime sets the CaptureRenderTargetsRenderTime property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#capturerendertargetsrendertime
func (s *SceneInstrumentation) SetCaptureRenderTargetsRenderTime(captureRenderTargetsRenderTime bool) *SceneInstrumentation {
	s.p.Set("captureRenderTargetsRenderTime", captureRenderTargetsRenderTime)
	return s
}

// CaptureRenderTime returns the CaptureRenderTime property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#capturerendertime
func (s *SceneInstrumentation) CaptureRenderTime() bool {
	retVal := s.p.Get("captureRenderTime")
	return retVal.Bool()
}

// SetCaptureRenderTime sets the CaptureRenderTime property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#capturerendertime
func (s *SceneInstrumentation) SetCaptureRenderTime(captureRenderTime bool) *SceneInstrumentation {
	s.p.Set("captureRenderTime", captureRenderTime)
	return s
}

// CaptureSpritesRenderTime returns the CaptureSpritesRenderTime property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#capturespritesrendertime
func (s *SceneInstrumentation) CaptureSpritesRenderTime() bool {
	retVal := s.p.Get("captureSpritesRenderTime")
	return retVal.Bool()
}

// SetCaptureSpritesRenderTime sets the CaptureSpritesRenderTime property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#capturespritesrendertime
func (s *SceneInstrumentation) SetCaptureSpritesRenderTime(captureSpritesRenderTime bool) *SceneInstrumentation {
	s.p.Set("captureSpritesRenderTime", captureSpritesRenderTime)
	return s
}

// DrawCallsCounter returns the DrawCallsCounter property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#drawcallscounter
func (s *SceneInstrumentation) DrawCallsCounter() *PerfCounter {
	retVal := s.p.Get("drawCallsCounter")
	return PerfCounterFromJSObject(retVal, s.ctx)
}

// SetDrawCallsCounter sets the DrawCallsCounter property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#drawcallscounter
func (s *SceneInstrumentation) SetDrawCallsCounter(drawCallsCounter *PerfCounter) *SceneInstrumentation {
	s.p.Set("drawCallsCounter", drawCallsCounter.JSObject())
	return s
}

// FrameTimeCounter returns the FrameTimeCounter property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#frametimecounter
func (s *SceneInstrumentation) FrameTimeCounter() *PerfCounter {
	retVal := s.p.Get("frameTimeCounter")
	return PerfCounterFromJSObject(retVal, s.ctx)
}

// SetFrameTimeCounter sets the FrameTimeCounter property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#frametimecounter
func (s *SceneInstrumentation) SetFrameTimeCounter(frameTimeCounter *PerfCounter) *SceneInstrumentation {
	s.p.Set("frameTimeCounter", frameTimeCounter.JSObject())
	return s
}

// InterFrameTimeCounter returns the InterFrameTimeCounter property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#interframetimecounter
func (s *SceneInstrumentation) InterFrameTimeCounter() *PerfCounter {
	retVal := s.p.Get("interFrameTimeCounter")
	return PerfCounterFromJSObject(retVal, s.ctx)
}

// SetInterFrameTimeCounter sets the InterFrameTimeCounter property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#interframetimecounter
func (s *SceneInstrumentation) SetInterFrameTimeCounter(interFrameTimeCounter *PerfCounter) *SceneInstrumentation {
	s.p.Set("interFrameTimeCounter", interFrameTimeCounter.JSObject())
	return s
}

// ParticlesRenderTimeCounter returns the ParticlesRenderTimeCounter property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#particlesrendertimecounter
func (s *SceneInstrumentation) ParticlesRenderTimeCounter() *PerfCounter {
	retVal := s.p.Get("particlesRenderTimeCounter")
	return PerfCounterFromJSObject(retVal, s.ctx)
}

// SetParticlesRenderTimeCounter sets the ParticlesRenderTimeCounter property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#particlesrendertimecounter
func (s *SceneInstrumentation) SetParticlesRenderTimeCounter(particlesRenderTimeCounter *PerfCounter) *SceneInstrumentation {
	s.p.Set("particlesRenderTimeCounter", particlesRenderTimeCounter.JSObject())
	return s
}

// PhysicsTimeCounter returns the PhysicsTimeCounter property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#physicstimecounter
func (s *SceneInstrumentation) PhysicsTimeCounter() *PerfCounter {
	retVal := s.p.Get("physicsTimeCounter")
	return PerfCounterFromJSObject(retVal, s.ctx)
}

// SetPhysicsTimeCounter sets the PhysicsTimeCounter property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#physicstimecounter
func (s *SceneInstrumentation) SetPhysicsTimeCounter(physicsTimeCounter *PerfCounter) *SceneInstrumentation {
	s.p.Set("physicsTimeCounter", physicsTimeCounter.JSObject())
	return s
}

// RenderTargetsRenderTimeCounter returns the RenderTargetsRenderTimeCounter property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#rendertargetsrendertimecounter
func (s *SceneInstrumentation) RenderTargetsRenderTimeCounter() *PerfCounter {
	retVal := s.p.Get("renderTargetsRenderTimeCounter")
	return PerfCounterFromJSObject(retVal, s.ctx)
}

// SetRenderTargetsRenderTimeCounter sets the RenderTargetsRenderTimeCounter property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#rendertargetsrendertimecounter
func (s *SceneInstrumentation) SetRenderTargetsRenderTimeCounter(renderTargetsRenderTimeCounter *PerfCounter) *SceneInstrumentation {
	s.p.Set("renderTargetsRenderTimeCounter", renderTargetsRenderTimeCounter.JSObject())
	return s
}

// RenderTimeCounter returns the RenderTimeCounter property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#rendertimecounter
func (s *SceneInstrumentation) RenderTimeCounter() *PerfCounter {
	retVal := s.p.Get("renderTimeCounter")
	return PerfCounterFromJSObject(retVal, s.ctx)
}

// SetRenderTimeCounter sets the RenderTimeCounter property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#rendertimecounter
func (s *SceneInstrumentation) SetRenderTimeCounter(renderTimeCounter *PerfCounter) *SceneInstrumentation {
	s.p.Set("renderTimeCounter", renderTimeCounter.JSObject())
	return s
}

// Scene returns the Scene property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#scene
func (s *SceneInstrumentation) Scene() *Scene {
	retVal := s.p.Get("scene")
	return SceneFromJSObject(retVal, s.ctx)
}

// SetScene sets the Scene property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#scene
func (s *SceneInstrumentation) SetScene(scene *Scene) *SceneInstrumentation {
	s.p.Set("scene", scene.JSObject())
	return s
}

// SpritesRenderTimeCounter returns the SpritesRenderTimeCounter property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#spritesrendertimecounter
func (s *SceneInstrumentation) SpritesRenderTimeCounter() *PerfCounter {
	retVal := s.p.Get("spritesRenderTimeCounter")
	return PerfCounterFromJSObject(retVal, s.ctx)
}

// SetSpritesRenderTimeCounter sets the SpritesRenderTimeCounter property of class SceneInstrumentation.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneinstrumentation#spritesrendertimecounter
func (s *SceneInstrumentation) SetSpritesRenderTimeCounter(spritesRenderTimeCounter *PerfCounter) *SceneInstrumentation {
	s.p.Set("spritesRenderTimeCounter", spritesRenderTimeCounter.JSObject())
	return s
}
