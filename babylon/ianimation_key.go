package babylon

import (
	"syscall/js"
)

// NewIAnimationKeyOpts contains optional parameters for NewIAnimationKey.
type NewIAnimationKeyOpts struct {
	InTangent     JSObject
	Interpolation *float64
	OutTangent    JSObject
}

// NewIAnimationKey returns a new IAnimationKey object.
//
// https://doc.babylonjs.com/api/classes/babylon.IAnimationKey
func (ba *Babylon) NewIAnimationKey(frame, value float64, opts *NewIAnimationKeyOpts) *IAnimationKey {
	obj := map[string]interface{}{
		"frame": frame,
		"value": value,
	}

	if opts != nil {
		if opts.InTangent != nil {
			obj["inTangent"] = opts.InTangent.JSObject()
		}
		if opts.Interpolation != nil {
			obj["interpolation"] = *opts.Interpolation
		}
		if opts.OutTangent != nil {
			obj["outTangent"] = opts.OutTangent.JSObject()
		}
	}

	p := js.ValueOf(obj)
	return IAnimationKeyFromJSObject(p, ba.ctx)
}
