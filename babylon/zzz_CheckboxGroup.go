// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// CheckboxGroup represents a babylon.js CheckboxGroup.
// Class used to create a CheckboxGroup
// which contains groups of checkbox buttons
type CheckboxGroup struct {
	*SelectorGroup
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (c *CheckboxGroup) JSObject() js.Value { return c.p }

// CheckboxGroup returns a CheckboxGroup JavaScript class.
func (ba *Babylon) CheckboxGroup() *CheckboxGroup {
	p := ba.ctx.Get("CheckboxGroup")
	return CheckboxGroupFromJSObject(p, ba.ctx)
}

// CheckboxGroupFromJSObject returns a wrapped CheckboxGroup JavaScript class.
func CheckboxGroupFromJSObject(p js.Value, ctx js.Value) *CheckboxGroup {
	return &CheckboxGroup{SelectorGroup: SelectorGroupFromJSObject(p, ctx), ctx: ctx}
}

// CheckboxGroupArrayToJSArray returns a JavaScript Array for the wrapped array.
func CheckboxGroupArrayToJSArray(array []*CheckboxGroup) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewCheckboxGroup returns a new CheckboxGroup object.
//
// https://doc.babylonjs.com/api/classes/babylon.checkboxgroup
func (gui *GUI) NewCheckboxGroup(name string) *CheckboxGroup {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := gui.ctx.Get("CheckboxGroup").New(args...)
	return CheckboxGroupFromJSObject(p, gui.ctx)
}

// CheckboxGroupAddCheckboxOpts contains optional parameters for CheckboxGroup.AddCheckbox.
type CheckboxGroupAddCheckboxOpts struct {
	Func    func()
	Checked *bool
}

// AddCheckbox calls the AddCheckbox method on the CheckboxGroup object.
//
// https://doc.babylonjs.com/api/classes/babylon.checkboxgroup#addcheckbox
func (c *CheckboxGroup) AddCheckbox(text string, opts *CheckboxGroupAddCheckboxOpts) {
	if opts == nil {
		opts = &CheckboxGroupAddCheckboxOpts{}
	}

	args := make([]interface{}, 0, 1+2)

	args = append(args, text)

	if opts.Func == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { opts.Func(); return nil }) /* never freed! */)
	}
	if opts.Checked == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Checked)
	}

	c.p.Call("addCheckbox", args...)
}
