package babylon

import (
	"syscall/js"
)

// ISceneLoaderPluginAsync represents a babylon.js ISceneLoaderPluginAsync.
type ISceneLoaderPluginAsync struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *ISceneLoaderPluginAsync) JSObject() js.Value { return a.p }

// ISceneLoaderPluginAsync returns a ISceneLoaderPluginAsync JavaScript class.
func (ba *Babylon) ISceneLoaderPluginAsync() *ISceneLoaderPluginAsync {
	p := ba.ctx.Get("ISceneLoaderPluginAsync")
	return ISceneLoaderPluginAsyncFromJSObject(p, ba.ctx)
}

// ISceneLoaderPluginAsyncFromJSObject returns a wrapped ISceneLoaderPluginAsync JavaScript class.
func ISceneLoaderPluginAsyncFromJSObject(p js.Value, ctx js.Value) *ISceneLoaderPluginAsync {
	return &ISceneLoaderPluginAsync{p: p, ctx: ctx}
}

// ISceneLoaderPluginAsyncArrayToJSArray returns a JavaScript Array for the wrapped array.
func ISceneLoaderPluginAsyncArrayToJSArray(array []*ISceneLoaderPluginAsync) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}
