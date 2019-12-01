package babylon

import (
	"syscall/js"
)

// ISimplificationSettings represents a babylon.js ISimplificationSettings.
type ISimplificationSettings struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *ISimplificationSettings) JSObject() js.Value { return a.p }

// ISimplificationSettings returns a ISimplificationSettings JavaScript class.
func (ba *Babylon) ISimplificationSettings() *ISimplificationSettings {
	p := ba.ctx.Get("ISimplificationSettings")
	return ISimplificationSettingsFromJSObject(p, ba.ctx)
}

// ISimplificationSettingsFromJSObject returns a wrapped ISimplificationSettings JavaScript class.
func ISimplificationSettingsFromJSObject(p js.Value, ctx js.Value) *ISimplificationSettings {
	return &ISimplificationSettings{p: p, ctx: ctx}
}

// ISimplificationSettingsArrayToJSArray returns a JavaScript Array for the wrapped array.
func ISimplificationSettingsArrayToJSArray(array []*ISimplificationSettings) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}
