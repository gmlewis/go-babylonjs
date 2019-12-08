package babylon

// IShadowLight is a helper function to reduce typing.
func (pl *PointLight) IShadowLight() *IShadowLight {
	return IShadowLightFromJSObject(pl.p, pl.ctx)
}
