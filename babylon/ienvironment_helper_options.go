package babylon

import (
	"syscall/js"
)

// NewIEnvironmentHelperOptions returns a new IEnvironmentHelperOptions object.
// For some reason, the docs don't mention a constructor, but 'new' is called
// in a demo: https://www.babylonjs-playground.com/#DMLMIP#1
//
//
// https://doc.babylonjs.com/api/classes/babylon.IEnvironmentHelperOptions
func (ba *Babylon) NewIEnvironmentHelperOptions() *IEnvironmentHelperOptions {
	obj := map[string]interface{}{}
	p := js.ValueOf(obj)
	return IEnvironmentHelperOptionsFromJSObject(p, ba.ctx)
}
