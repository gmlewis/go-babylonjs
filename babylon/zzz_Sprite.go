// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Sprite represents a babylon.js Sprite.
// Class used to represent a sprite
//
// See: http://doc.babylonjs.com/babylon101/sprites
type Sprite struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *Sprite) JSObject() js.Value { return s.p }

// Sprite returns a Sprite JavaScript class.
func (ba *Babylon) Sprite() *Sprite {
	p := ba.ctx.Get("Sprite")
	return SpriteFromJSObject(p, ba.ctx)
}

// SpriteFromJSObject returns a wrapped Sprite JavaScript class.
func SpriteFromJSObject(p js.Value, ctx js.Value) *Sprite {
	return &Sprite{p: p, ctx: ctx}
}

// SpriteArrayToJSArray returns a JavaScript Array for the wrapped array.
func SpriteArrayToJSArray(array []*Sprite) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewSprite returns a new Sprite object.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite
func (ba *Babylon) NewSprite(name string, manager js.Value) *Sprite {

	args := make([]interface{}, 0, 2+0)

	args = append(args, name)
	args = append(args, manager)

	p := ba.ctx.Get("Sprite").New(args...)
	return SpriteFromJSObject(p, ba.ctx)
}

// Dispose calls the Dispose method on the Sprite object.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#dispose
func (s *Sprite) Dispose() {

	s.p.Call("dispose")
}

// PlayAnimation calls the PlayAnimation method on the Sprite object.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#playanimation
func (s *Sprite) PlayAnimation(from float64, to float64, loop bool, delay float64, onAnimationEnd func()) {

	args := make([]interface{}, 0, 5+0)

	args = append(args, from)
	args = append(args, to)
	args = append(args, loop)
	args = append(args, delay)
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { onAnimationEnd(); return nil }))

	s.p.Call("playAnimation", args...)
}

// StopAnimation calls the StopAnimation method on the Sprite object.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#stopanimation
func (s *Sprite) StopAnimation() {

	s.p.Call("stopAnimation")
}

/*

// ActionManager returns the ActionManager property of class Sprite.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#actionmanager
func (s *Sprite) ActionManager(actionManager *ActionManager) *Sprite {
	p := ba.ctx.Get("Sprite").New(actionManager.JSObject())
	return SpriteFromJSObject(p, ba.ctx)
}

// SetActionManager sets the ActionManager property of class Sprite.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#actionmanager
func (s *Sprite) SetActionManager(actionManager *ActionManager) *Sprite {
	p := ba.ctx.Get("Sprite").New(actionManager.JSObject())
	return SpriteFromJSObject(p, ba.ctx)
}

// Angle returns the Angle property of class Sprite.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#angle
func (s *Sprite) Angle(angle float64) *Sprite {
	p := ba.ctx.Get("Sprite").New(angle)
	return SpriteFromJSObject(p, ba.ctx)
}

// SetAngle sets the Angle property of class Sprite.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#angle
func (s *Sprite) SetAngle(angle float64) *Sprite {
	p := ba.ctx.Get("Sprite").New(angle)
	return SpriteFromJSObject(p, ba.ctx)
}

// Animations returns the Animations property of class Sprite.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#animations
func (s *Sprite) Animations(animations *Animation) *Sprite {
	p := ba.ctx.Get("Sprite").New(animations.JSObject())
	return SpriteFromJSObject(p, ba.ctx)
}

// SetAnimations sets the Animations property of class Sprite.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#animations
func (s *Sprite) SetAnimations(animations *Animation) *Sprite {
	p := ba.ctx.Get("Sprite").New(animations.JSObject())
	return SpriteFromJSObject(p, ba.ctx)
}

// CellIndex returns the CellIndex property of class Sprite.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#cellindex
func (s *Sprite) CellIndex(cellIndex float64) *Sprite {
	p := ba.ctx.Get("Sprite").New(cellIndex)
	return SpriteFromJSObject(p, ba.ctx)
}

// SetCellIndex sets the CellIndex property of class Sprite.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#cellindex
func (s *Sprite) SetCellIndex(cellIndex float64) *Sprite {
	p := ba.ctx.Get("Sprite").New(cellIndex)
	return SpriteFromJSObject(p, ba.ctx)
}

// CellRef returns the CellRef property of class Sprite.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#cellref
func (s *Sprite) CellRef(cellRef string) *Sprite {
	p := ba.ctx.Get("Sprite").New(cellRef)
	return SpriteFromJSObject(p, ba.ctx)
}

// SetCellRef sets the CellRef property of class Sprite.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#cellref
func (s *Sprite) SetCellRef(cellRef string) *Sprite {
	p := ba.ctx.Get("Sprite").New(cellRef)
	return SpriteFromJSObject(p, ba.ctx)
}

// Color returns the Color property of class Sprite.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#color
func (s *Sprite) Color(color *Color4) *Sprite {
	p := ba.ctx.Get("Sprite").New(color.JSObject())
	return SpriteFromJSObject(p, ba.ctx)
}

// SetColor sets the Color property of class Sprite.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#color
func (s *Sprite) SetColor(color *Color4) *Sprite {
	p := ba.ctx.Get("Sprite").New(color.JSObject())
	return SpriteFromJSObject(p, ba.ctx)
}

// DisposeWhenFinishedAnimating returns the DisposeWhenFinishedAnimating property of class Sprite.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#disposewhenfinishedanimating
func (s *Sprite) DisposeWhenFinishedAnimating(disposeWhenFinishedAnimating bool) *Sprite {
	p := ba.ctx.Get("Sprite").New(disposeWhenFinishedAnimating)
	return SpriteFromJSObject(p, ba.ctx)
}

// SetDisposeWhenFinishedAnimating sets the DisposeWhenFinishedAnimating property of class Sprite.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#disposewhenfinishedanimating
func (s *Sprite) SetDisposeWhenFinishedAnimating(disposeWhenFinishedAnimating bool) *Sprite {
	p := ba.ctx.Get("Sprite").New(disposeWhenFinishedAnimating)
	return SpriteFromJSObject(p, ba.ctx)
}

// Height returns the Height property of class Sprite.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#height
func (s *Sprite) Height(height float64) *Sprite {
	p := ba.ctx.Get("Sprite").New(height)
	return SpriteFromJSObject(p, ba.ctx)
}

// SetHeight sets the Height property of class Sprite.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#height
func (s *Sprite) SetHeight(height float64) *Sprite {
	p := ba.ctx.Get("Sprite").New(height)
	return SpriteFromJSObject(p, ba.ctx)
}

// InvertU returns the InvertU property of class Sprite.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#invertu
func (s *Sprite) InvertU(invertU float64) *Sprite {
	p := ba.ctx.Get("Sprite").New(invertU)
	return SpriteFromJSObject(p, ba.ctx)
}

// SetInvertU sets the InvertU property of class Sprite.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#invertu
func (s *Sprite) SetInvertU(invertU float64) *Sprite {
	p := ba.ctx.Get("Sprite").New(invertU)
	return SpriteFromJSObject(p, ba.ctx)
}

// InvertV returns the InvertV property of class Sprite.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#invertv
func (s *Sprite) InvertV(invertV float64) *Sprite {
	p := ba.ctx.Get("Sprite").New(invertV)
	return SpriteFromJSObject(p, ba.ctx)
}

// SetInvertV sets the InvertV property of class Sprite.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#invertv
func (s *Sprite) SetInvertV(invertV float64) *Sprite {
	p := ba.ctx.Get("Sprite").New(invertV)
	return SpriteFromJSObject(p, ba.ctx)
}

// IsPickable returns the IsPickable property of class Sprite.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#ispickable
func (s *Sprite) IsPickable(isPickable bool) *Sprite {
	p := ba.ctx.Get("Sprite").New(isPickable)
	return SpriteFromJSObject(p, ba.ctx)
}

// SetIsPickable sets the IsPickable property of class Sprite.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#ispickable
func (s *Sprite) SetIsPickable(isPickable bool) *Sprite {
	p := ba.ctx.Get("Sprite").New(isPickable)
	return SpriteFromJSObject(p, ba.ctx)
}

// IsVisible returns the IsVisible property of class Sprite.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#isvisible
func (s *Sprite) IsVisible(isVisible bool) *Sprite {
	p := ba.ctx.Get("Sprite").New(isVisible)
	return SpriteFromJSObject(p, ba.ctx)
}

// SetIsVisible sets the IsVisible property of class Sprite.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#isvisible
func (s *Sprite) SetIsVisible(isVisible bool) *Sprite {
	p := ba.ctx.Get("Sprite").New(isVisible)
	return SpriteFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class Sprite.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#name
func (s *Sprite) Name(name string) *Sprite {
	p := ba.ctx.Get("Sprite").New(name)
	return SpriteFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class Sprite.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#name
func (s *Sprite) SetName(name string) *Sprite {
	p := ba.ctx.Get("Sprite").New(name)
	return SpriteFromJSObject(p, ba.ctx)
}

// Position returns the Position property of class Sprite.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#position
func (s *Sprite) Position(position *Vector3) *Sprite {
	p := ba.ctx.Get("Sprite").New(position.JSObject())
	return SpriteFromJSObject(p, ba.ctx)
}

// SetPosition sets the Position property of class Sprite.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#position
func (s *Sprite) SetPosition(position *Vector3) *Sprite {
	p := ba.ctx.Get("Sprite").New(position.JSObject())
	return SpriteFromJSObject(p, ba.ctx)
}

// Size returns the Size property of class Sprite.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#size
func (s *Sprite) Size(size float64) *Sprite {
	p := ba.ctx.Get("Sprite").New(size)
	return SpriteFromJSObject(p, ba.ctx)
}

// SetSize sets the Size property of class Sprite.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#size
func (s *Sprite) SetSize(size float64) *Sprite {
	p := ba.ctx.Get("Sprite").New(size)
	return SpriteFromJSObject(p, ba.ctx)
}

// Width returns the Width property of class Sprite.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#width
func (s *Sprite) Width(width float64) *Sprite {
	p := ba.ctx.Get("Sprite").New(width)
	return SpriteFromJSObject(p, ba.ctx)
}

// SetWidth sets the Width property of class Sprite.
//
// https://doc.babylonjs.com/api/classes/babylon.sprite#width
func (s *Sprite) SetWidth(width float64) *Sprite {
	p := ba.ctx.Get("Sprite").New(width)
	return SpriteFromJSObject(p, ba.ctx)
}

*/
