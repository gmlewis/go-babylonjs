package babylon

/*
// Scene represents a babylon.js Scene.
type Scene struct{ *AbstractScene }

// JSObject returns the underlying js.Value.
func (s *Scene) JSObject() js.Value { return s.p }

// Scene returns a Scene JavaScript class.
func (b *Babylon) Scene() *Scene {
	p := b.ctx.Get("Scene")
	return SceneFromJSObject(p)
}

// SceneFromJSObject returns a wrapped Scene JavaScript class.
func SceneFromJSObject(p js.Value) *Scene {
	return &Scene{AbstractSceneFromJSObject(p)}
}

// NewScene returns a new Scene object.
//
// https://doc.babylonjs.com/api/classes/babylon.scene
func (b *Babylon) NewScene(engine *Engine) *Scene {
	p := b.ctx.Get("Scene").New(engine.JSObject())
	return SceneFromJSObject(p)
}
*/

// Render calls the JavaScript method of the same name.
func (s *Scene) Render() {
	s.p.Call("render")
}
