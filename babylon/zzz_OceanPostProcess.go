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
// var pp = new OceanPostProcess(&amp;quot;myOcean&amp;quot;, camera);
// pp.reflectionEnabled = true;
// pp.refractionEnabled = true;
type OceanPostProcess struct{ *PostProcess }

// JSObject returns the underlying js.Value.
func (o *OceanPostProcess) JSObject() js.Value { return o.p }

// OceanPostProcess returns a OceanPostProcess JavaScript class.
func (b *Babylon) OceanPostProcess() *OceanPostProcess {
	p := b.ctx.Get("OceanPostProcess")
	return OceanPostProcessFromJSObject(p)
}

// OceanPostProcessFromJSObject returns a wrapped OceanPostProcess JavaScript class.
func OceanPostProcessFromJSObject(p js.Value) *OceanPostProcess {
	return &OceanPostProcess{PostProcessFromJSObject(p)}
}

// NewOceanPostProcessOpts contains optional parameters for NewOceanPostProcess.
type NewOceanPostProcessOpts struct {
	Options *IOceanPostProcessOptions
}

// NewOceanPostProcess returns a new OceanPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.oceanpostprocess
func (b *Babylon) NewOceanPostProcess(name string, camera *TargetCamera, opts *NewOceanPostProcessOpts) *OceanPostProcess {
	if opts == nil {
		opts = &NewOceanPostProcessOpts{}
	}

	p := b.ctx.Get("OceanPostProcess").New(name, camera.JSObject(), opts.Options.JSObject())
	return OceanPostProcessFromJSObject(p)
}

// TODO: methods
