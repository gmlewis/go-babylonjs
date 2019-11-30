package babylon

// DOUBLESIDE calls the JavaScript method of the same name.
func (m *Mesh) DOUBLESIDE() *float64 {
	v := m.p.Get("DOUBLESIDE")
	f := v.Float()
	return &f
}
