// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// TextFileAssetTask represents a babylon.js TextFileAssetTask.
// Define a task used by AssetsManager to load text content
type TextFileAssetTask struct {
	*AbstractAssetTask
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (t *TextFileAssetTask) JSObject() js.Value { return t.p }

// TextFileAssetTask returns a TextFileAssetTask JavaScript class.
func (ba *Babylon) TextFileAssetTask() *TextFileAssetTask {
	p := ba.ctx.Get("TextFileAssetTask")
	return TextFileAssetTaskFromJSObject(p, ba.ctx)
}

// TextFileAssetTaskFromJSObject returns a wrapped TextFileAssetTask JavaScript class.
func TextFileAssetTaskFromJSObject(p js.Value, ctx js.Value) *TextFileAssetTask {
	return &TextFileAssetTask{AbstractAssetTask: AbstractAssetTaskFromJSObject(p, ctx), ctx: ctx}
}

// TextFileAssetTaskArrayToJSArray returns a JavaScript Array for the wrapped array.
func TextFileAssetTaskArrayToJSArray(array []*TextFileAssetTask) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewTextFileAssetTask returns a new TextFileAssetTask object.
//
// https://doc.babylonjs.com/api/classes/babylon.textfileassettask
func (ba *Babylon) NewTextFileAssetTask(name string, url string) *TextFileAssetTask {

	args := make([]interface{}, 0, 2+0)

	args = append(args, name)
	args = append(args, url)

	p := ba.ctx.Get("TextFileAssetTask").New(args...)
	return TextFileAssetTaskFromJSObject(p, ba.ctx)
}

// RunTask calls the RunTask method on the TextFileAssetTask object.
//
// https://doc.babylonjs.com/api/classes/babylon.textfileassettask#runtask
func (t *TextFileAssetTask) RunTask(scene *Scene, onSuccess func(), onError func()) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, scene.JSObject())
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { onSuccess(); return nil }))
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { onError(); return nil }))

	t.p.Call("runTask", args...)
}

/*

// Name returns the Name property of class TextFileAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.textfileassettask#name
func (t *TextFileAssetTask) Name(name string) *TextFileAssetTask {
	p := ba.ctx.Get("TextFileAssetTask").New(name)
	return TextFileAssetTaskFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class TextFileAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.textfileassettask#name
func (t *TextFileAssetTask) SetName(name string) *TextFileAssetTask {
	p := ba.ctx.Get("TextFileAssetTask").New(name)
	return TextFileAssetTaskFromJSObject(p, ba.ctx)
}

// OnError returns the OnError property of class TextFileAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.textfileassettask#onerror
func (t *TextFileAssetTask) OnError(onError func()) *TextFileAssetTask {
	p := ba.ctx.Get("TextFileAssetTask").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onError(); return nil}))
	return TextFileAssetTaskFromJSObject(p, ba.ctx)
}

// SetOnError sets the OnError property of class TextFileAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.textfileassettask#onerror
func (t *TextFileAssetTask) SetOnError(onError func()) *TextFileAssetTask {
	p := ba.ctx.Get("TextFileAssetTask").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onError(); return nil}))
	return TextFileAssetTaskFromJSObject(p, ba.ctx)
}

// OnSuccess returns the OnSuccess property of class TextFileAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.textfileassettask#onsuccess
func (t *TextFileAssetTask) OnSuccess(onSuccess func()) *TextFileAssetTask {
	p := ba.ctx.Get("TextFileAssetTask").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onSuccess(); return nil}))
	return TextFileAssetTaskFromJSObject(p, ba.ctx)
}

// SetOnSuccess sets the OnSuccess property of class TextFileAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.textfileassettask#onsuccess
func (t *TextFileAssetTask) SetOnSuccess(onSuccess func()) *TextFileAssetTask {
	p := ba.ctx.Get("TextFileAssetTask").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onSuccess(); return nil}))
	return TextFileAssetTaskFromJSObject(p, ba.ctx)
}

// Text returns the Text property of class TextFileAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.textfileassettask#text
func (t *TextFileAssetTask) Text(text string) *TextFileAssetTask {
	p := ba.ctx.Get("TextFileAssetTask").New(text)
	return TextFileAssetTaskFromJSObject(p, ba.ctx)
}

// SetText sets the Text property of class TextFileAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.textfileassettask#text
func (t *TextFileAssetTask) SetText(text string) *TextFileAssetTask {
	p := ba.ctx.Get("TextFileAssetTask").New(text)
	return TextFileAssetTaskFromJSObject(p, ba.ctx)
}

// Url returns the Url property of class TextFileAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.textfileassettask#url
func (t *TextFileAssetTask) Url(url string) *TextFileAssetTask {
	p := ba.ctx.Get("TextFileAssetTask").New(url)
	return TextFileAssetTaskFromJSObject(p, ba.ctx)
}

// SetUrl sets the Url property of class TextFileAssetTask.
//
// https://doc.babylonjs.com/api/classes/babylon.textfileassettask#url
func (t *TextFileAssetTask) SetUrl(url string) *TextFileAssetTask {
	p := ba.ctx.Get("TextFileAssetTask").New(url)
	return TextFileAssetTaskFromJSObject(p, ba.ctx)
}

*/
