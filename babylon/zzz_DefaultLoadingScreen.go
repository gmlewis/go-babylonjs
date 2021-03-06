// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DefaultLoadingScreen represents a babylon.js DefaultLoadingScreen.
// Class used for the default loading screen
//
// See: http://doc.babylonjs.com/how_to/creating_a_custom_loading_screen
type DefaultLoadingScreen struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (d *DefaultLoadingScreen) JSObject() js.Value { return d.p }

// DefaultLoadingScreen returns a DefaultLoadingScreen JavaScript class.
func (ba *Babylon) DefaultLoadingScreen() *DefaultLoadingScreen {
	p := ba.ctx.Get("DefaultLoadingScreen")
	return DefaultLoadingScreenFromJSObject(p, ba.ctx)
}

// DefaultLoadingScreenFromJSObject returns a wrapped DefaultLoadingScreen JavaScript class.
func DefaultLoadingScreenFromJSObject(p js.Value, ctx js.Value) *DefaultLoadingScreen {
	return &DefaultLoadingScreen{p: p, ctx: ctx}
}

// DefaultLoadingScreenArrayToJSArray returns a JavaScript Array for the wrapped array.
func DefaultLoadingScreenArrayToJSArray(array []*DefaultLoadingScreen) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewDefaultLoadingScreenOpts contains optional parameters for NewDefaultLoadingScreen.
type NewDefaultLoadingScreenOpts struct {
	_loadingText               *string
	_loadingDivBackgroundColor *string
}

// NewDefaultLoadingScreen returns a new DefaultLoadingScreen object.
//
// https://doc.babylonjs.com/api/classes/babylon.defaultloadingscreen#constructor
func (ba *Babylon) NewDefaultLoadingScreen(_renderingCanvas js.Value, opts *NewDefaultLoadingScreenOpts) *DefaultLoadingScreen {
	if opts == nil {
		opts = &NewDefaultLoadingScreenOpts{}
	}

	args := make([]interface{}, 0, 1+2)

	args = append(args, _renderingCanvas)

	if opts._loadingText == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts._loadingText)
	}
	if opts._loadingDivBackgroundColor == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts._loadingDivBackgroundColor)
	}

	p := ba.ctx.Get("DefaultLoadingScreen").New(args...)
	return DefaultLoadingScreenFromJSObject(p, ba.ctx)
}

// DisplayLoadingUI calls the DisplayLoadingUI method on the DefaultLoadingScreen object.
//
// https://doc.babylonjs.com/api/classes/babylon.defaultloadingscreen#displayloadingui
func (d *DefaultLoadingScreen) DisplayLoadingUI() {

	d.p.Call("displayLoadingUI")
}

// HideLoadingUI calls the HideLoadingUI method on the DefaultLoadingScreen object.
//
// https://doc.babylonjs.com/api/classes/babylon.defaultloadingscreen#hideloadingui
func (d *DefaultLoadingScreen) HideLoadingUI() {

	d.p.Call("hideLoadingUI")
}

// DefaultLogoUrl returns the DefaultLogoUrl property of class DefaultLoadingScreen.
//
// https://doc.babylonjs.com/api/classes/babylon.defaultloadingscreen#defaultlogourl
func (d *DefaultLoadingScreen) DefaultLogoUrl() string {
	retVal := d.p.Get("DefaultLogoUrl")
	return retVal.String()
}

// SetDefaultLogoUrl sets the DefaultLogoUrl property of class DefaultLoadingScreen.
//
// https://doc.babylonjs.com/api/classes/babylon.defaultloadingscreen#defaultlogourl
func (d *DefaultLoadingScreen) SetDefaultLogoUrl(DefaultLogoUrl string) *DefaultLoadingScreen {
	d.p.Set("DefaultLogoUrl", DefaultLogoUrl)
	return d
}

// DefaultSpinnerUrl returns the DefaultSpinnerUrl property of class DefaultLoadingScreen.
//
// https://doc.babylonjs.com/api/classes/babylon.defaultloadingscreen#defaultspinnerurl
func (d *DefaultLoadingScreen) DefaultSpinnerUrl() string {
	retVal := d.p.Get("DefaultSpinnerUrl")
	return retVal.String()
}

// SetDefaultSpinnerUrl sets the DefaultSpinnerUrl property of class DefaultLoadingScreen.
//
// https://doc.babylonjs.com/api/classes/babylon.defaultloadingscreen#defaultspinnerurl
func (d *DefaultLoadingScreen) SetDefaultSpinnerUrl(DefaultSpinnerUrl string) *DefaultLoadingScreen {
	d.p.Set("DefaultSpinnerUrl", DefaultSpinnerUrl)
	return d
}

// LoadingUIBackgroundColor returns the LoadingUIBackgroundColor property of class DefaultLoadingScreen.
//
// https://doc.babylonjs.com/api/classes/babylon.defaultloadingscreen#loadinguibackgroundcolor
func (d *DefaultLoadingScreen) LoadingUIBackgroundColor() string {
	retVal := d.p.Get("loadingUIBackgroundColor")
	return retVal.String()
}

// SetLoadingUIBackgroundColor sets the LoadingUIBackgroundColor property of class DefaultLoadingScreen.
//
// https://doc.babylonjs.com/api/classes/babylon.defaultloadingscreen#loadinguibackgroundcolor
func (d *DefaultLoadingScreen) SetLoadingUIBackgroundColor(loadingUIBackgroundColor string) *DefaultLoadingScreen {
	d.p.Set("loadingUIBackgroundColor", loadingUIBackgroundColor)
	return d
}

// LoadingUIText returns the LoadingUIText property of class DefaultLoadingScreen.
//
// https://doc.babylonjs.com/api/classes/babylon.defaultloadingscreen#loadinguitext
func (d *DefaultLoadingScreen) LoadingUIText() string {
	retVal := d.p.Get("loadingUIText")
	return retVal.String()
}

// SetLoadingUIText sets the LoadingUIText property of class DefaultLoadingScreen.
//
// https://doc.babylonjs.com/api/classes/babylon.defaultloadingscreen#loadinguitext
func (d *DefaultLoadingScreen) SetLoadingUIText(loadingUIText string) *DefaultLoadingScreen {
	d.p.Set("loadingUIText", loadingUIText)
	return d
}
