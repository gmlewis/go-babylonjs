package babylon

// SetIntensity calls the JavaScript method of the same name.
func (l *Light) SetIntensity(intensity float64) {
	l.p.Set("intensity", intensity)
}
