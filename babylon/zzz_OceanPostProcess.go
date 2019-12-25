// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// OceanPostProcess represents a babylon.js OceanPostProcess.
// OceanPostProcess helps rendering an infinite ocean surface that can reflect and refract environment.
//
// Simmply add it to your scene and let the nerd that lives in you have fun.
// Example usage:
// var pp = new OceanPostProcess(&quot;myOcean&quot;, camera);
// pp.reflectionEnabled = true;
// pp.refractionEnabled = true;
type OceanPostProcess struct {
	*PostProcess
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (o *OceanPostProcess) JSObject() js.Value { return o.p }

// OceanPostProcess returns a OceanPostProcess JavaScript class.
func (ba *Babylon) OceanPostProcess() *OceanPostProcess {
	p := ba.ctx.Get("OceanPostProcess")
	return OceanPostProcessFromJSObject(p, ba.ctx)
}

// OceanPostProcessFromJSObject returns a wrapped OceanPostProcess JavaScript class.
func OceanPostProcessFromJSObject(p js.Value, ctx js.Value) *OceanPostProcess {
	return &OceanPostProcess{PostProcess: PostProcessFromJSObject(p, ctx), ctx: ctx}
}

// OceanPostProcessArrayToJSArray returns a JavaScript Array for the wrapped array.
func OceanPostProcessArrayToJSArray(array []*OceanPostProcess) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewOceanPostProcessOpts contains optional parameters for NewOceanPostProcess.
type NewOceanPostProcessOpts struct {
	Options *IOceanPostProcessOptions
}

// NewOceanPostProcess returns a new OceanPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.oceanpostprocess#constructor
func (ba *Babylon) NewOceanPostProcess(name string, camera *TargetCamera, opts *NewOceanPostProcessOpts) *OceanPostProcess {
	if opts == nil {
		opts = &NewOceanPostProcessOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, camera.JSObject())

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options.JSObject())
	}

	p := ba.ctx.Get("OceanPostProcess").New(args...)
	return OceanPostProcessFromJSObject(p, ba.ctx)
}

// IsSupported returns the IsSupported property of class OceanPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.oceanpostprocess#issupported
func (o *OceanPostProcess) IsSupported() bool {
	retVal := o.p.Get("isSupported")
	return retVal.Bool()
}

// SetIsSupported sets the IsSupported property of class OceanPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.oceanpostprocess#issupported
func (o *OceanPostProcess) SetIsSupported(isSupported bool) *OceanPostProcess {
	o.p.Set("isSupported", isSupported)
	return o
}

// ReflectionEnabled returns the ReflectionEnabled property of class OceanPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.oceanpostprocess#reflectionenabled
func (o *OceanPostProcess) ReflectionEnabled() bool {
	retVal := o.p.Get("reflectionEnabled")
	return retVal.Bool()
}

// SetReflectionEnabled sets the ReflectionEnabled property of class OceanPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.oceanpostprocess#reflectionenabled
func (o *OceanPostProcess) SetReflectionEnabled(reflectionEnabled bool) *OceanPostProcess {
	o.p.Set("reflectionEnabled", reflectionEnabled)
	return o
}

// ReflectionTexture returns the ReflectionTexture property of class OceanPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.oceanpostprocess#reflectiontexture
func (o *OceanPostProcess) ReflectionTexture() *MirrorTexture {
	retVal := o.p.Get("reflectionTexture")
	return MirrorTextureFromJSObject(retVal, o.ctx)
}

// SetReflectionTexture sets the ReflectionTexture property of class OceanPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.oceanpostprocess#reflectiontexture
func (o *OceanPostProcess) SetReflectionTexture(reflectionTexture *MirrorTexture) *OceanPostProcess {
	o.p.Set("reflectionTexture", reflectionTexture.JSObject())
	return o
}

// RefractionEnabled returns the RefractionEnabled property of class OceanPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.oceanpostprocess#refractionenabled
func (o *OceanPostProcess) RefractionEnabled() bool {
	retVal := o.p.Get("refractionEnabled")
	return retVal.Bool()
}

// SetRefractionEnabled sets the RefractionEnabled property of class OceanPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.oceanpostprocess#refractionenabled
func (o *OceanPostProcess) SetRefractionEnabled(refractionEnabled bool) *OceanPostProcess {
	o.p.Set("refractionEnabled", refractionEnabled)
	return o
}

// RefractionTexture returns the RefractionTexture property of class OceanPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.oceanpostprocess#refractiontexture
func (o *OceanPostProcess) RefractionTexture() *RenderTargetTexture {
	retVal := o.p.Get("refractionTexture")
	return RenderTargetTextureFromJSObject(retVal, o.ctx)
}

// SetRefractionTexture sets the RefractionTexture property of class OceanPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.oceanpostprocess#refractiontexture
func (o *OceanPostProcess) SetRefractionTexture(refractionTexture *RenderTargetTexture) *OceanPostProcess {
	o.p.Set("refractionTexture", refractionTexture.JSObject())
	return o
}
