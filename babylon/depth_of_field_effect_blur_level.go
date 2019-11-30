package babylon

import "syscall/js"

// DepthOfFieldEffectBlurLevel represents a Babylon enum.
type DepthOfFieldEffectBlurLevel int

var (
	// DepthOfFieldEffectBlurLevelLow represents a Babylon enum.
	DepthOfFieldEffectBlurLevelLow DepthOfFieldEffectBlurLevel = 0
	// DepthOfFieldEffectBlurLevelMedium represents a Babylon enum.
	DepthOfFieldEffectBlurLevelMedium DepthOfFieldEffectBlurLevel = 1
	// DepthOfFieldEffectBlurLevelHigh represents a Babylon enum.
	DepthOfFieldEffectBlurLevelHigh DepthOfFieldEffectBlurLevel = 2
)

// JSObject returns the underlying js.Value.
func (d *DepthOfFieldEffectBlurLevel) JSObject() js.Value {
	if d == nil {
		return js.Null()
	}
	return js.ValueOf(*d)
}
