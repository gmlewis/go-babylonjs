// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PostProcessRenderPipeline represents a babylon.js PostProcessRenderPipeline.
// PostProcessRenderPipeline
//
// See: https://doc.babylonjs.com/how_to/how_to_use_postprocessrenderpipeline
type PostProcessRenderPipeline struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PostProcessRenderPipeline) JSObject() js.Value { return p.p }

// PostProcessRenderPipeline returns a PostProcessRenderPipeline JavaScript class.
func (ba *Babylon) PostProcessRenderPipeline() *PostProcessRenderPipeline {
	p := ba.ctx.Get("PostProcessRenderPipeline")
	return PostProcessRenderPipelineFromJSObject(p, ba.ctx)
}

// PostProcessRenderPipelineFromJSObject returns a wrapped PostProcessRenderPipeline JavaScript class.
func PostProcessRenderPipelineFromJSObject(p js.Value, ctx js.Value) *PostProcessRenderPipeline {
	return &PostProcessRenderPipeline{p: p, ctx: ctx}
}

// PostProcessRenderPipelineArrayToJSArray returns a JavaScript Array for the wrapped array.
func PostProcessRenderPipelineArrayToJSArray(array []*PostProcessRenderPipeline) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewPostProcessRenderPipeline returns a new PostProcessRenderPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocessrenderpipeline#constructor
func (ba *Babylon) NewPostProcessRenderPipeline(engine *Engine, name string) *PostProcessRenderPipeline {

	args := make([]interface{}, 0, 2+0)

	args = append(args, engine.JSObject())
	args = append(args, name)

	p := ba.ctx.Get("PostProcessRenderPipeline").New(args...)
	return PostProcessRenderPipelineFromJSObject(p, ba.ctx)
}

// AddEffect calls the AddEffect method on the PostProcessRenderPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocessrenderpipeline#addeffect
func (p *PostProcessRenderPipeline) AddEffect(renderEffect *PostProcessRenderEffect) {

	args := make([]interface{}, 0, 1+0)

	if renderEffect == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, renderEffect.JSObject())
	}

	p.p.Call("addEffect", args...)
}

// Dispose calls the Dispose method on the PostProcessRenderPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocessrenderpipeline#dispose
func (p *PostProcessRenderPipeline) Dispose() {

	p.p.Call("dispose")
}

// GetClassName calls the GetClassName method on the PostProcessRenderPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocessrenderpipeline#getclassname
func (p *PostProcessRenderPipeline) GetClassName() string {

	retVal := p.p.Call("getClassName")
	return retVal.String()
}

// Cameras returns the Cameras property of class PostProcessRenderPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocessrenderpipeline#cameras
func (p *PostProcessRenderPipeline) Cameras() []*Camera {
	retVal := p.p.Get("cameras")
	result := []*Camera{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, CameraFromJSObject(retVal.Index(ri), p.ctx))
	}
	return result
}

// SetCameras sets the Cameras property of class PostProcessRenderPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocessrenderpipeline#cameras
func (p *PostProcessRenderPipeline) SetCameras(cameras []*Camera) *PostProcessRenderPipeline {
	p.p.Set("cameras", cameras)
	return p
}

// InspectableCustomProperties returns the InspectableCustomProperties property of class PostProcessRenderPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocessrenderpipeline#inspectablecustomproperties
func (p *PostProcessRenderPipeline) InspectableCustomProperties() []*IInspectable {
	retVal := p.p.Get("inspectableCustomProperties")
	result := []*IInspectable{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, IInspectableFromJSObject(retVal.Index(ri), p.ctx))
	}
	return result
}

// SetInspectableCustomProperties sets the InspectableCustomProperties property of class PostProcessRenderPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocessrenderpipeline#inspectablecustomproperties
func (p *PostProcessRenderPipeline) SetInspectableCustomProperties(inspectableCustomProperties []*IInspectable) *PostProcessRenderPipeline {
	p.p.Set("inspectableCustomProperties", inspectableCustomProperties)
	return p
}

// IsSupported returns the IsSupported property of class PostProcessRenderPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocessrenderpipeline#issupported
func (p *PostProcessRenderPipeline) IsSupported() bool {
	retVal := p.p.Get("isSupported")
	return retVal.Bool()
}

// SetIsSupported sets the IsSupported property of class PostProcessRenderPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocessrenderpipeline#issupported
func (p *PostProcessRenderPipeline) SetIsSupported(isSupported bool) *PostProcessRenderPipeline {
	p.p.Set("isSupported", isSupported)
	return p
}

// Name returns the Name property of class PostProcessRenderPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocessrenderpipeline#name
func (p *PostProcessRenderPipeline) Name() string {
	retVal := p.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class PostProcessRenderPipeline.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocessrenderpipeline#name
func (p *PostProcessRenderPipeline) SetName(name string) *PostProcessRenderPipeline {
	p.p.Set("name", name)
	return p
}
