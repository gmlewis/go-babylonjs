package babylon

// SetClearColor calls the JavaScript method of the same name.
func (s *Scene) SetClearColor(color *Color4) {
	s.p.Set("clearColor", color.JSObject())
}
