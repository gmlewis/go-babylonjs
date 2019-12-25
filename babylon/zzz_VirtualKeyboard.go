// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// VirtualKeyboard represents a babylon.js VirtualKeyboard.
// Class used to create virtual keyboard
type VirtualKeyboard struct {
	*StackPanel
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (v *VirtualKeyboard) JSObject() js.Value { return v.p }

// VirtualKeyboard returns a VirtualKeyboard JavaScript class.
func (gui *GUI) VirtualKeyboard() *VirtualKeyboard {
	p := gui.ctx.Get("VirtualKeyboard")
	return VirtualKeyboardFromJSObject(p, gui.ctx)
}

// VirtualKeyboardFromJSObject returns a wrapped VirtualKeyboard JavaScript class.
func VirtualKeyboardFromJSObject(p js.Value, ctx js.Value) *VirtualKeyboard {
	return &VirtualKeyboard{StackPanel: StackPanelFromJSObject(p, ctx), ctx: ctx}
}

// VirtualKeyboardArrayToJSArray returns a JavaScript Array for the wrapped array.
func VirtualKeyboardArrayToJSArray(array []*VirtualKeyboard) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewVirtualKeyboardOpts contains optional parameters for NewVirtualKeyboard.
type NewVirtualKeyboardOpts struct {
	Name *string
}

// NewVirtualKeyboard returns a new VirtualKeyboard object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.virtualkeyboard#constructor
func (gui *GUI) NewVirtualKeyboard(opts *NewVirtualKeyboardOpts) *VirtualKeyboard {
	if opts == nil {
		opts = &NewVirtualKeyboardOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Name == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Name)
	}

	p := gui.ctx.Get("VirtualKeyboard").New(args...)
	return VirtualKeyboardFromJSObject(p, gui.ctx)
}

// VirtualKeyboardAddKeysRowOpts contains optional parameters for VirtualKeyboard.AddKeysRow.
type VirtualKeyboardAddKeysRowOpts struct {
	PropertySets []js.Value
}

// AddKeysRow calls the AddKeysRow method on the VirtualKeyboard object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.virtualkeyboard#addkeysrow
func (v *VirtualKeyboard) AddKeysRow(keys []string, opts *VirtualKeyboardAddKeysRowOpts) {
	if opts == nil {
		opts = &VirtualKeyboardAddKeysRowOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, keys)

	if opts.PropertySets == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.PropertySets)
	}

	v.p.Call("addKeysRow", args...)
}

// ApplyShiftState calls the ApplyShiftState method on the VirtualKeyboard object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.virtualkeyboard#applyshiftstate
func (v *VirtualKeyboard) ApplyShiftState(shiftState float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, shiftState)

	v.p.Call("applyShiftState", args...)
}

// Connect calls the Connect method on the VirtualKeyboard object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.virtualkeyboard#connect
func (v *VirtualKeyboard) Connect(input *InputText) {

	args := make([]interface{}, 0, 1+0)

	if input == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, input.JSObject())
	}

	v.p.Call("connect", args...)
}

// VirtualKeyboardCreateDefaultLayoutOpts contains optional parameters for VirtualKeyboard.CreateDefaultLayout.
type VirtualKeyboardCreateDefaultLayoutOpts struct {
	Name *string
}

// CreateDefaultLayout calls the CreateDefaultLayout method on the VirtualKeyboard object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.virtualkeyboard#createdefaultlayout
func (v *VirtualKeyboard) CreateDefaultLayout(opts *VirtualKeyboardCreateDefaultLayoutOpts) *VirtualKeyboard {
	if opts == nil {
		opts = &VirtualKeyboardCreateDefaultLayoutOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Name == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Name)
	}

	retVal := v.p.Call("CreateDefaultLayout", args...)
	return VirtualKeyboardFromJSObject(retVal, v.ctx)
}

// VirtualKeyboardDisconnectOpts contains optional parameters for VirtualKeyboard.Disconnect.
type VirtualKeyboardDisconnectOpts struct {
	Input *InputText
}

// Disconnect calls the Disconnect method on the VirtualKeyboard object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.virtualkeyboard#disconnect
func (v *VirtualKeyboard) Disconnect(opts *VirtualKeyboardDisconnectOpts) {
	if opts == nil {
		opts = &VirtualKeyboardDisconnectOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Input == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Input.JSObject())
	}

	v.p.Call("disconnect", args...)
}

// Dispose calls the Dispose method on the VirtualKeyboard object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.virtualkeyboard#dispose
func (v *VirtualKeyboard) Dispose() {

	v.p.Call("dispose")
}

// ConnectedInputText returns the ConnectedInputText property of class VirtualKeyboard.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.virtualkeyboard#connectedinputtext
func (v *VirtualKeyboard) ConnectedInputText() *InputText {
	retVal := v.p.Get("connectedInputText")
	return InputTextFromJSObject(retVal, v.ctx)
}

// SetConnectedInputText sets the ConnectedInputText property of class VirtualKeyboard.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.virtualkeyboard#connectedinputtext
func (v *VirtualKeyboard) SetConnectedInputText(connectedInputText *InputText) *VirtualKeyboard {
	v.p.Set("connectedInputText", connectedInputText.JSObject())
	return v
}

// DefaultButtonBackground returns the DefaultButtonBackground property of class VirtualKeyboard.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.virtualkeyboard#defaultbuttonbackground
func (v *VirtualKeyboard) DefaultButtonBackground() string {
	retVal := v.p.Get("defaultButtonBackground")
	return retVal.String()
}

// SetDefaultButtonBackground sets the DefaultButtonBackground property of class VirtualKeyboard.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.virtualkeyboard#defaultbuttonbackground
func (v *VirtualKeyboard) SetDefaultButtonBackground(defaultButtonBackground string) *VirtualKeyboard {
	v.p.Set("defaultButtonBackground", defaultButtonBackground)
	return v
}

// DefaultButtonColor returns the DefaultButtonColor property of class VirtualKeyboard.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.virtualkeyboard#defaultbuttoncolor
func (v *VirtualKeyboard) DefaultButtonColor() string {
	retVal := v.p.Get("defaultButtonColor")
	return retVal.String()
}

// SetDefaultButtonColor sets the DefaultButtonColor property of class VirtualKeyboard.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.virtualkeyboard#defaultbuttoncolor
func (v *VirtualKeyboard) SetDefaultButtonColor(defaultButtonColor string) *VirtualKeyboard {
	v.p.Set("defaultButtonColor", defaultButtonColor)
	return v
}

// DefaultButtonHeight returns the DefaultButtonHeight property of class VirtualKeyboard.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.virtualkeyboard#defaultbuttonheight
func (v *VirtualKeyboard) DefaultButtonHeight() string {
	retVal := v.p.Get("defaultButtonHeight")
	return retVal.String()
}

// SetDefaultButtonHeight sets the DefaultButtonHeight property of class VirtualKeyboard.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.virtualkeyboard#defaultbuttonheight
func (v *VirtualKeyboard) SetDefaultButtonHeight(defaultButtonHeight string) *VirtualKeyboard {
	v.p.Set("defaultButtonHeight", defaultButtonHeight)
	return v
}

// DefaultButtonPaddingBottom returns the DefaultButtonPaddingBottom property of class VirtualKeyboard.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.virtualkeyboard#defaultbuttonpaddingbottom
func (v *VirtualKeyboard) DefaultButtonPaddingBottom() string {
	retVal := v.p.Get("defaultButtonPaddingBottom")
	return retVal.String()
}

// SetDefaultButtonPaddingBottom sets the DefaultButtonPaddingBottom property of class VirtualKeyboard.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.virtualkeyboard#defaultbuttonpaddingbottom
func (v *VirtualKeyboard) SetDefaultButtonPaddingBottom(defaultButtonPaddingBottom string) *VirtualKeyboard {
	v.p.Set("defaultButtonPaddingBottom", defaultButtonPaddingBottom)
	return v
}

// DefaultButtonPaddingLeft returns the DefaultButtonPaddingLeft property of class VirtualKeyboard.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.virtualkeyboard#defaultbuttonpaddingleft
func (v *VirtualKeyboard) DefaultButtonPaddingLeft() string {
	retVal := v.p.Get("defaultButtonPaddingLeft")
	return retVal.String()
}

// SetDefaultButtonPaddingLeft sets the DefaultButtonPaddingLeft property of class VirtualKeyboard.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.virtualkeyboard#defaultbuttonpaddingleft
func (v *VirtualKeyboard) SetDefaultButtonPaddingLeft(defaultButtonPaddingLeft string) *VirtualKeyboard {
	v.p.Set("defaultButtonPaddingLeft", defaultButtonPaddingLeft)
	return v
}

// DefaultButtonPaddingRight returns the DefaultButtonPaddingRight property of class VirtualKeyboard.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.virtualkeyboard#defaultbuttonpaddingright
func (v *VirtualKeyboard) DefaultButtonPaddingRight() string {
	retVal := v.p.Get("defaultButtonPaddingRight")
	return retVal.String()
}

// SetDefaultButtonPaddingRight sets the DefaultButtonPaddingRight property of class VirtualKeyboard.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.virtualkeyboard#defaultbuttonpaddingright
func (v *VirtualKeyboard) SetDefaultButtonPaddingRight(defaultButtonPaddingRight string) *VirtualKeyboard {
	v.p.Set("defaultButtonPaddingRight", defaultButtonPaddingRight)
	return v
}

// DefaultButtonPaddingTop returns the DefaultButtonPaddingTop property of class VirtualKeyboard.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.virtualkeyboard#defaultbuttonpaddingtop
func (v *VirtualKeyboard) DefaultButtonPaddingTop() string {
	retVal := v.p.Get("defaultButtonPaddingTop")
	return retVal.String()
}

// SetDefaultButtonPaddingTop sets the DefaultButtonPaddingTop property of class VirtualKeyboard.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.virtualkeyboard#defaultbuttonpaddingtop
func (v *VirtualKeyboard) SetDefaultButtonPaddingTop(defaultButtonPaddingTop string) *VirtualKeyboard {
	v.p.Set("defaultButtonPaddingTop", defaultButtonPaddingTop)
	return v
}

// DefaultButtonWidth returns the DefaultButtonWidth property of class VirtualKeyboard.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.virtualkeyboard#defaultbuttonwidth
func (v *VirtualKeyboard) DefaultButtonWidth() string {
	retVal := v.p.Get("defaultButtonWidth")
	return retVal.String()
}

// SetDefaultButtonWidth sets the DefaultButtonWidth property of class VirtualKeyboard.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.virtualkeyboard#defaultbuttonwidth
func (v *VirtualKeyboard) SetDefaultButtonWidth(defaultButtonWidth string) *VirtualKeyboard {
	v.p.Set("defaultButtonWidth", defaultButtonWidth)
	return v
}

// OnKeyPressObservable returns the OnKeyPressObservable property of class VirtualKeyboard.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.virtualkeyboard#onkeypressobservable
func (v *VirtualKeyboard) OnKeyPressObservable() *Observable {
	retVal := v.p.Get("onKeyPressObservable")
	return ObservableFromJSObject(retVal, v.ctx)
}

// SetOnKeyPressObservable sets the OnKeyPressObservable property of class VirtualKeyboard.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.virtualkeyboard#onkeypressobservable
func (v *VirtualKeyboard) SetOnKeyPressObservable(onKeyPressObservable *Observable) *VirtualKeyboard {
	v.p.Set("onKeyPressObservable", onKeyPressObservable.JSObject())
	return v
}

// SelectedShiftThickness returns the SelectedShiftThickness property of class VirtualKeyboard.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.virtualkeyboard#selectedshiftthickness
func (v *VirtualKeyboard) SelectedShiftThickness() float64 {
	retVal := v.p.Get("selectedShiftThickness")
	return retVal.Float()
}

// SetSelectedShiftThickness sets the SelectedShiftThickness property of class VirtualKeyboard.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.virtualkeyboard#selectedshiftthickness
func (v *VirtualKeyboard) SetSelectedShiftThickness(selectedShiftThickness float64) *VirtualKeyboard {
	v.p.Set("selectedShiftThickness", selectedShiftThickness)
	return v
}

// ShiftButtonColor returns the ShiftButtonColor property of class VirtualKeyboard.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.virtualkeyboard#shiftbuttoncolor
func (v *VirtualKeyboard) ShiftButtonColor() string {
	retVal := v.p.Get("shiftButtonColor")
	return retVal.String()
}

// SetShiftButtonColor sets the ShiftButtonColor property of class VirtualKeyboard.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.virtualkeyboard#shiftbuttoncolor
func (v *VirtualKeyboard) SetShiftButtonColor(shiftButtonColor string) *VirtualKeyboard {
	v.p.Set("shiftButtonColor", shiftButtonColor)
	return v
}

// ShiftState returns the ShiftState property of class VirtualKeyboard.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.virtualkeyboard#shiftstate
func (v *VirtualKeyboard) ShiftState() float64 {
	retVal := v.p.Get("shiftState")
	return retVal.Float()
}

// SetShiftState sets the ShiftState property of class VirtualKeyboard.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.virtualkeyboard#shiftstate
func (v *VirtualKeyboard) SetShiftState(shiftState float64) *VirtualKeyboard {
	v.p.Set("shiftState", shiftState)
	return v
}
