package babylon

// IShadowLight is a helper function to reduce typing.
func (dl *DirectionalLight) IShadowLight() *IShadowLight {
	return IShadowLightFromJSObject(dl.p, dl.ctx)
}
