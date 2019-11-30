package babylon

// Render calls the JavaScript method of the same name.
func (s *Scene) Render() {
	s.p.Call("render")
}

// SetClearColor calls the JavaScript method of the same name.
func (s *Scene) SetClearColor(color *Color3) {
	s.p.Set("clearColor", color.JSObject())
}
