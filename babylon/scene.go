package babylon

// Render calls the JavaScript method of the same name.
func (s *Scene) Render() {
	s.p.Call("render")
}
