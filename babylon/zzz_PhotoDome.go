// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PhotoDome represents a babylon.js PhotoDome.
// Display a 360 degree photo on an approximately spherical surface, useful for VR applications or skyboxes.
// As a subclass of TransformNode, this allow parenting to the camera with different locations in the scene.
// This class achieves its effect with a Texture and a correctly configured BackgroundMaterial on an inverted sphere.
// Potential additions to this helper include zoom and and non-infinite distance rendering effects.
type PhotoDome struct {
	*TransformNode
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PhotoDome) JSObject() js.Value { return p.p }

// PhotoDome returns a PhotoDome JavaScript class.
func (ba *Babylon) PhotoDome() *PhotoDome {
	p := ba.ctx.Get("PhotoDome")
	return PhotoDomeFromJSObject(p, ba.ctx)
}

// PhotoDomeFromJSObject returns a wrapped PhotoDome JavaScript class.
func PhotoDomeFromJSObject(p js.Value, ctx js.Value) *PhotoDome {
	return &PhotoDome{TransformNode: TransformNodeFromJSObject(p, ctx), ctx: ctx}
}

// PhotoDomeArrayToJSArray returns a JavaScript Array for the wrapped array.
func PhotoDomeArrayToJSArray(array []*PhotoDome) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewPhotoDomeOpts contains optional parameters for NewPhotoDome.
type NewPhotoDomeOpts struct {
	OnError JSFunc
}

// NewPhotoDome returns a new PhotoDome object.
//
// https://doc.babylonjs.com/api/classes/babylon.photodome#constructor
func (ba *Babylon) NewPhotoDome(name string, urlOfPhoto string, options js.Value, scene *Scene, opts *NewPhotoDomeOpts) *PhotoDome {
	if opts == nil {
		opts = &NewPhotoDomeOpts{}
	}

	args := make([]interface{}, 0, 4+1)

	args = append(args, name)
	args = append(args, urlOfPhoto)
	args = append(args, options)
	args = append(args, scene.JSObject())

	if opts.OnError == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(opts.OnError) /* never freed! */)
	}

	p := ba.ctx.Get("PhotoDome").New(args...)
	return PhotoDomeFromJSObject(p, ba.ctx)
}

// PhotoDomeDisposeOpts contains optional parameters for PhotoDome.Dispose.
type PhotoDomeDisposeOpts struct {
	DoNotRecurse               *bool
	DisposeMaterialAndTextures *bool
}

// Dispose calls the Dispose method on the PhotoDome object.
//
// https://doc.babylonjs.com/api/classes/babylon.photodome#dispose
func (p *PhotoDome) Dispose(opts *PhotoDomeDisposeOpts) {
	if opts == nil {
		opts = &PhotoDomeDisposeOpts{}
	}

	args := make([]interface{}, 0, 0+2)

	if opts.DoNotRecurse == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.DoNotRecurse)
	}
	if opts.DisposeMaterialAndTextures == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.DisposeMaterialAndTextures)
	}

	p.p.Call("dispose", args...)
}

// FovMultiplier returns the FovMultiplier property of class PhotoDome.
//
// https://doc.babylonjs.com/api/classes/babylon.photodome#fovmultiplier
func (p *PhotoDome) FovMultiplier() float64 {
	retVal := p.p.Get("fovMultiplier")
	return retVal.Float()
}

// SetFovMultiplier sets the FovMultiplier property of class PhotoDome.
//
// https://doc.babylonjs.com/api/classes/babylon.photodome#fovmultiplier
func (p *PhotoDome) SetFovMultiplier(fovMultiplier float64) *PhotoDome {
	p.p.Set("fovMultiplier", fovMultiplier)
	return p
}

// ImageMode returns the ImageMode property of class PhotoDome.
//
// https://doc.babylonjs.com/api/classes/babylon.photodome#imagemode
func (p *PhotoDome) ImageMode() float64 {
	retVal := p.p.Get("imageMode")
	return retVal.Float()
}

// SetImageMode sets the ImageMode property of class PhotoDome.
//
// https://doc.babylonjs.com/api/classes/babylon.photodome#imagemode
func (p *PhotoDome) SetImageMode(imageMode float64) *PhotoDome {
	p.p.Set("imageMode", imageMode)
	return p
}

// MODE_MONOSCOPIC returns the MODE_MONOSCOPIC property of class PhotoDome.
//
// https://doc.babylonjs.com/api/classes/babylon.photodome#mode_monoscopic
func (p *PhotoDome) MODE_MONOSCOPIC() float64 {
	retVal := p.p.Get("MODE_MONOSCOPIC")
	return retVal.Float()
}

// SetMODE_MONOSCOPIC sets the MODE_MONOSCOPIC property of class PhotoDome.
//
// https://doc.babylonjs.com/api/classes/babylon.photodome#mode_monoscopic
func (p *PhotoDome) SetMODE_MONOSCOPIC(MODE_MONOSCOPIC float64) *PhotoDome {
	p.p.Set("MODE_MONOSCOPIC", MODE_MONOSCOPIC)
	return p
}

// MODE_SIDEBYSIDE returns the MODE_SIDEBYSIDE property of class PhotoDome.
//
// https://doc.babylonjs.com/api/classes/babylon.photodome#mode_sidebyside
func (p *PhotoDome) MODE_SIDEBYSIDE() float64 {
	retVal := p.p.Get("MODE_SIDEBYSIDE")
	return retVal.Float()
}

// SetMODE_SIDEBYSIDE sets the MODE_SIDEBYSIDE property of class PhotoDome.
//
// https://doc.babylonjs.com/api/classes/babylon.photodome#mode_sidebyside
func (p *PhotoDome) SetMODE_SIDEBYSIDE(MODE_SIDEBYSIDE float64) *PhotoDome {
	p.p.Set("MODE_SIDEBYSIDE", MODE_SIDEBYSIDE)
	return p
}

// MODE_TOPBOTTOM returns the MODE_TOPBOTTOM property of class PhotoDome.
//
// https://doc.babylonjs.com/api/classes/babylon.photodome#mode_topbottom
func (p *PhotoDome) MODE_TOPBOTTOM() float64 {
	retVal := p.p.Get("MODE_TOPBOTTOM")
	return retVal.Float()
}

// SetMODE_TOPBOTTOM sets the MODE_TOPBOTTOM property of class PhotoDome.
//
// https://doc.babylonjs.com/api/classes/babylon.photodome#mode_topbottom
func (p *PhotoDome) SetMODE_TOPBOTTOM(MODE_TOPBOTTOM float64) *PhotoDome {
	p.p.Set("MODE_TOPBOTTOM", MODE_TOPBOTTOM)
	return p
}

// Mesh returns the Mesh property of class PhotoDome.
//
// https://doc.babylonjs.com/api/classes/babylon.photodome#mesh
func (p *PhotoDome) Mesh() *Mesh {
	retVal := p.p.Get("mesh")
	return MeshFromJSObject(retVal, p.ctx)
}

// SetMesh sets the Mesh property of class PhotoDome.
//
// https://doc.babylonjs.com/api/classes/babylon.photodome#mesh
func (p *PhotoDome) SetMesh(mesh *Mesh) *PhotoDome {
	p.p.Set("mesh", mesh.JSObject())
	return p
}

// OnLoadErrorObservable returns the OnLoadErrorObservable property of class PhotoDome.
//
// https://doc.babylonjs.com/api/classes/babylon.photodome#onloaderrorobservable
func (p *PhotoDome) OnLoadErrorObservable() *Observable {
	retVal := p.p.Get("onLoadErrorObservable")
	return ObservableFromJSObject(retVal, p.ctx)
}

// SetOnLoadErrorObservable sets the OnLoadErrorObservable property of class PhotoDome.
//
// https://doc.babylonjs.com/api/classes/babylon.photodome#onloaderrorobservable
func (p *PhotoDome) SetOnLoadErrorObservable(onLoadErrorObservable *Observable) *PhotoDome {
	p.p.Set("onLoadErrorObservable", onLoadErrorObservable.JSObject())
	return p
}

// PhotoTexture returns the PhotoTexture property of class PhotoDome.
//
// https://doc.babylonjs.com/api/classes/babylon.photodome#phototexture
func (p *PhotoDome) PhotoTexture() *Texture {
	retVal := p.p.Get("photoTexture")
	return TextureFromJSObject(retVal, p.ctx)
}

// SetPhotoTexture sets the PhotoTexture property of class PhotoDome.
//
// https://doc.babylonjs.com/api/classes/babylon.photodome#phototexture
func (p *PhotoDome) SetPhotoTexture(photoTexture *Texture) *PhotoDome {
	p.p.Set("photoTexture", photoTexture.JSObject())
	return p
}
