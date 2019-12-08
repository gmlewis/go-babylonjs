package babylon

import "syscall/js"

// SetClearColor calls the JavaScript method of the same name.
func (s *Scene) SetClearColor(color *Color4) {
	s.p.Set("clearColor", color.JSObject())
}

// RenderLoopFunc is a convenience function to reduce typing.
func (s *Scene) RenderLoopFunc(opts *SceneRenderOpts) JSFunc {
	return func(this js.Value, args []js.Value) interface{} {
		s.Render(opts)
		return nil
	}
}
