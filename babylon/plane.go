package babylon

// Normalize calls the JavaScript method of the same name.
func (p *Plane) Normalize() {
	p.p.Call("normalize")
}
