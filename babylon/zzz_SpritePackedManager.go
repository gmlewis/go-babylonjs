// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SpritePackedManager represents a babylon.js SpritePackedManager.
// Class used to manage multiple sprites of different sizes on the same spritesheet
//
// See: http://doc.babylonjs.com/babylon101/sprites
type SpritePackedManager struct {
	*SpriteManager
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *SpritePackedManager) JSObject() js.Value { return s.p }

// SpritePackedManager returns a SpritePackedManager JavaScript class.
func (ba *Babylon) SpritePackedManager() *SpritePackedManager {
	p := ba.ctx.Get("SpritePackedManager")
	return SpritePackedManagerFromJSObject(p, ba.ctx)
}

// SpritePackedManagerFromJSObject returns a wrapped SpritePackedManager JavaScript class.
func SpritePackedManagerFromJSObject(p js.Value, ctx js.Value) *SpritePackedManager {
	return &SpritePackedManager{SpriteManager: SpriteManagerFromJSObject(p, ctx), ctx: ctx}
}

// NewSpritePackedManagerOpts contains optional parameters for NewSpritePackedManager.
type NewSpritePackedManagerOpts struct {
	SpriteJSON *JSString

	Epsilon *JSFloat64

	SamplingMode *JSFloat64
}

// NewSpritePackedManager returns a new SpritePackedManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.spritepackedmanager
func (ba *Babylon) NewSpritePackedManager(name string, imgUrl string, capacity float64, scene *Scene, opts *NewSpritePackedManagerOpts) *SpritePackedManager {
	if opts == nil {
		opts = &NewSpritePackedManagerOpts{}
	}

	p := ba.ctx.Get("SpritePackedManager").New(name, imgUrl, capacity, scene.JSObject(), opts.SpriteJSON.JSObject(), opts.Epsilon.JSObject(), opts.SamplingMode.JSObject())
	return SpritePackedManagerFromJSObject(p, ba.ctx)
}

// TODO: methods
