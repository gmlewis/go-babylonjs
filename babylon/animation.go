package babylon

import "syscall/js"

// AnimationArray makes it easier to manipulate the JavaScript arrays.
type AnimationArray struct {
	p   js.Value
	ctx js.Value
	a   []*Animation
}

// JSObject returns the underlying *js.JavaScript class.
func (aa *AnimationArray) JSObject() js.Value { return aa.p }

// AnimationArrayFromJSObject returns a wrapped AnimationArray JavaScript class.
func AnimationArrayFromJSObject(p js.Value, ctx js.Value) *AnimationArray {
	result := []*Animation{}
	for ri := 0; ri < p.Length(); ri++ {
		result = append(result, AnimationFromJSObject(p.Index(ri), ctx))
	}
	return &AnimationArray{p: p, a: result, ctx: ctx}
}

// Push pushes meshes to an Animation.
func (aa *AnimationArray) Push(meshes ...*Animation) {
	args := []interface{}{}
	for _, m := range meshes {
		aa.a = append(aa.a, m)
		args = append(args, m.JSObject())
	}
	aa.p.Call("push", args...)
}
