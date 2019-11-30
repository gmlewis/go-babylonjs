// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// VRExperienceHelper represents a babylon.js VRExperienceHelper.
// Helps to quickly add VR support to an existing scene.
// See &lt;a href=&#34;http://doc.babylonjs.com/how_to/webvr_helper&#34;&gt;http://doc.babylonjs.com/how_to/webvr_helper&lt;/a&gt;
type VRExperienceHelper struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (v *VRExperienceHelper) JSObject() js.Value { return v.p }

// VRExperienceHelper returns a VRExperienceHelper JavaScript class.
func (b *Babylon) VRExperienceHelper() *VRExperienceHelper {
	p := b.ctx.Get("VRExperienceHelper")
	return VRExperienceHelperFromJSObject(p)
}

// VRExperienceHelperFromJSObject returns a wrapped VRExperienceHelper JavaScript class.
func VRExperienceHelperFromJSObject(p js.Value) *VRExperienceHelper {
	return &VRExperienceHelper{p: p}
}

// NewVRExperienceHelperOpts contains optional parameters for NewVRExperienceHelper.
type NewVRExperienceHelperOpts struct {
	WebVROptions *VRExperienceHelperOptions
}

// NewVRExperienceHelper returns a new VRExperienceHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrexperiencehelper
func (b *Babylon) NewVRExperienceHelper(scene *Scene, opts *NewVRExperienceHelperOpts) *VRExperienceHelper {
	if opts == nil {
		opts = &NewVRExperienceHelperOpts{}
	}

	p := b.ctx.Get("VRExperienceHelper").New(scene.JSObject(), opts.WebVROptions.JSObject())
	return VRExperienceHelperFromJSObject(p)
}

// TODO: methods
