// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DeepCopier represents a babylon.js DeepCopier.
// Class containing a set of static utilities functions for deep copy.
type DeepCopier struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (d *DeepCopier) JSObject() js.Value { return d.p }

// DeepCopier returns a DeepCopier JavaScript class.
func (ba *Babylon) DeepCopier() *DeepCopier {
	p := ba.ctx.Get("DeepCopier")
	return DeepCopierFromJSObject(p, ba.ctx)
}

// DeepCopierFromJSObject returns a wrapped DeepCopier JavaScript class.
func DeepCopierFromJSObject(p js.Value, ctx js.Value) *DeepCopier {
	return &DeepCopier{p: p, ctx: ctx}
}

// DeepCopierArrayToJSArray returns a JavaScript Array for the wrapped array.
func DeepCopierArrayToJSArray(array []*DeepCopier) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// DeepCopierDeepCopyOpts contains optional parameters for DeepCopier.DeepCopy.
type DeepCopierDeepCopyOpts struct {
	DoNotCopyList []string
	MustCopyList  []string
}

// DeepCopy calls the DeepCopy method on the DeepCopier object.
//
// https://doc.babylonjs.com/api/classes/babylon.deepcopier#deepcopy
func (d *DeepCopier) DeepCopy(source JSObject, destination JSObject, opts *DeepCopierDeepCopyOpts) {
	if opts == nil {
		opts = &DeepCopierDeepCopyOpts{}
	}

	args := make([]interface{}, 0, 2+2)

	if source == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, source.JSObject())
	}

	if destination == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, destination.JSObject())
	}

	if opts.DoNotCopyList == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.DoNotCopyList)
	}
	if opts.MustCopyList == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.MustCopyList)
	}

	d.p.Call("DeepCopy", args...)
}
