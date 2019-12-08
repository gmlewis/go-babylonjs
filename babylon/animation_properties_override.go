package babylon

// NewAnimationPropertiesOverride returns a new AnimationPropertiesOverride object.
// For some reason, the docs don't mention a constructor, but 'new' is called
// in a demo: https://www.babylonjs-playground.com/#DMLMIP#1
//
//
// https://doc.babylonjs.com/api/classes/babylon.animationpropertiesoverride
func (ba *Babylon) NewAnimationPropertiesOverride() *AnimationPropertiesOverride {
	p := ba.ctx.Get("AnimationPropertiesOverride").New()
	return AnimationPropertiesOverrideFromJSObject(p, ba.ctx)
}
