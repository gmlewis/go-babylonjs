// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ScrollViewer represents a babylon.js ScrollViewer.
// Class used to hold a viewer window and sliders in a grid
type ScrollViewer struct {
	*Rectangle
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *ScrollViewer) JSObject() js.Value { return s.p }

// ScrollViewer returns a ScrollViewer JavaScript class.
func (gui *GUI) ScrollViewer() *ScrollViewer {
	p := gui.ctx.Get("ScrollViewer")
	return ScrollViewerFromJSObject(p, gui.ctx)
}

// ScrollViewerFromJSObject returns a wrapped ScrollViewer JavaScript class.
func ScrollViewerFromJSObject(p js.Value, ctx js.Value) *ScrollViewer {
	return &ScrollViewer{Rectangle: RectangleFromJSObject(p, ctx), ctx: ctx}
}

// ScrollViewerArrayToJSArray returns a JavaScript Array for the wrapped array.
func ScrollViewerArrayToJSArray(array []*ScrollViewer) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewScrollViewerOpts contains optional parameters for NewScrollViewer.
type NewScrollViewerOpts struct {
	Name         *string
	IsImageBased *bool
}

// NewScrollViewer returns a new ScrollViewer object.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer
func (gui *GUI) NewScrollViewer(opts *NewScrollViewerOpts) *ScrollViewer {
	if opts == nil {
		opts = &NewScrollViewerOpts{}
	}

	args := make([]interface{}, 0, 0+2)

	if opts.Name == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Name)
	}
	if opts.IsImageBased == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.IsImageBased)
	}

	p := gui.ctx.Get("ScrollViewer").New(args...)
	return ScrollViewerFromJSObject(p, gui.ctx)
}

// AddControl calls the AddControl method on the ScrollViewer object.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer#addcontrol
func (s *ScrollViewer) AddControl(control *Control) *Container {

	args := make([]interface{}, 0, 1+0)

	args = append(args, control.JSObject())

	retVal := s.p.Call("addControl", args...)
	return ContainerFromJSObject(retVal, s.ctx)
}

// Dispose calls the Dispose method on the ScrollViewer object.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer#dispose
func (s *ScrollViewer) Dispose() {

	s.p.Call("dispose")
}

// RemoveControl calls the RemoveControl method on the ScrollViewer object.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer#removecontrol
func (s *ScrollViewer) RemoveControl(control *Control) *Container {

	args := make([]interface{}, 0, 1+0)

	args = append(args, control.JSObject())

	retVal := s.p.Call("removeControl", args...)
	return ContainerFromJSObject(retVal, s.ctx)
}

// ResetWindow calls the ResetWindow method on the ScrollViewer object.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer#resetwindow
func (s *ScrollViewer) ResetWindow() {

	s.p.Call("resetWindow")
}

// _flagDescendantsAsMatrixDirty calls the _flagDescendantsAsMatrixDirty method on the ScrollViewer object.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer#_flagdescendantsasmatrixdirty
func (s *ScrollViewer) _flagDescendantsAsMatrixDirty() {

	s.p.Call("_flagDescendantsAsMatrixDirty")
}

// _link calls the _link method on the ScrollViewer object.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer#_link
func (s *ScrollViewer) _link(host *AdvancedDynamicTexture) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, host.JSObject())

	s.p.Call("_link", args...)
}

// _renderHighlightSpecific calls the _renderHighlightSpecific method on the ScrollViewer object.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer#_renderhighlightspecific
func (s *ScrollViewer) _renderHighlightSpecific(context js.Value) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, context)

	s.p.Call("_renderHighlightSpecific", args...)
}

// BarBackground returns the BarBackground property of class ScrollViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer#barbackground
func (s *ScrollViewer) BarBackground() string {
	retVal := s.p.Get("barBackground")
	return retVal.String()
}

// SetBarBackground sets the BarBackground property of class ScrollViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer#barbackground
func (s *ScrollViewer) SetBarBackground(barBackground string) *ScrollViewer {
	s.p.Set("barBackground", barBackground)
	return s
}

// BarColor returns the BarColor property of class ScrollViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer#barcolor
func (s *ScrollViewer) BarColor() string {
	retVal := s.p.Get("barColor")
	return retVal.String()
}

// SetBarColor sets the BarColor property of class ScrollViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer#barcolor
func (s *ScrollViewer) SetBarColor(barColor string) *ScrollViewer {
	s.p.Set("barColor", barColor)
	return s
}

// BarImage returns the BarImage property of class ScrollViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer#barimage
func (s *ScrollViewer) BarImage() *Image {
	retVal := s.p.Get("barImage")
	return ImageFromJSObject(retVal, s.ctx)
}

// SetBarImage sets the BarImage property of class ScrollViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer#barimage
func (s *ScrollViewer) SetBarImage(barImage *Image) *ScrollViewer {
	s.p.Set("barImage", barImage.JSObject())
	return s
}

// BarImageHeight returns the BarImageHeight property of class ScrollViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer#barimageheight
func (s *ScrollViewer) BarImageHeight() float64 {
	retVal := s.p.Get("barImageHeight")
	return retVal.Float()
}

// SetBarImageHeight sets the BarImageHeight property of class ScrollViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer#barimageheight
func (s *ScrollViewer) SetBarImageHeight(barImageHeight float64) *ScrollViewer {
	s.p.Set("barImageHeight", barImageHeight)
	return s
}

// BarSize returns the BarSize property of class ScrollViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer#barsize
func (s *ScrollViewer) BarSize() float64 {
	retVal := s.p.Get("barSize")
	return retVal.Float()
}

// SetBarSize sets the BarSize property of class ScrollViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer#barsize
func (s *ScrollViewer) SetBarSize(barSize float64) *ScrollViewer {
	s.p.Set("barSize", barSize)
	return s
}

// Children returns the Children property of class ScrollViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer#children
func (s *ScrollViewer) Children() []*Control {
	retVal := s.p.Get("children")
	result := []*Control{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, ControlFromJSObject(retVal.Index(ri), s.ctx))
	}
	return result
}

// SetChildren sets the Children property of class ScrollViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer#children
func (s *ScrollViewer) SetChildren(children []*Control) *ScrollViewer {
	s.p.Set("children", children)
	return s
}

// HorizontalBar returns the HorizontalBar property of class ScrollViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer#horizontalbar
func (s *ScrollViewer) HorizontalBar() *ScrollBar {
	retVal := s.p.Get("horizontalBar")
	return ScrollBarFromJSObject(retVal, s.ctx)
}

// SetHorizontalBar sets the HorizontalBar property of class ScrollViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer#horizontalbar
func (s *ScrollViewer) SetHorizontalBar(horizontalBar *ScrollBar) *ScrollViewer {
	s.p.Set("horizontalBar", horizontalBar.JSObject())
	return s
}

// ScrollBackground returns the ScrollBackground property of class ScrollViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer#scrollbackground
func (s *ScrollViewer) ScrollBackground() string {
	retVal := s.p.Get("scrollBackground")
	return retVal.String()
}

// SetScrollBackground sets the ScrollBackground property of class ScrollViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer#scrollbackground
func (s *ScrollViewer) SetScrollBackground(scrollBackground string) *ScrollViewer {
	s.p.Set("scrollBackground", scrollBackground)
	return s
}

// ThumbHeight returns the ThumbHeight property of class ScrollViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer#thumbheight
func (s *ScrollViewer) ThumbHeight() float64 {
	retVal := s.p.Get("thumbHeight")
	return retVal.Float()
}

// SetThumbHeight sets the ThumbHeight property of class ScrollViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer#thumbheight
func (s *ScrollViewer) SetThumbHeight(thumbHeight float64) *ScrollViewer {
	s.p.Set("thumbHeight", thumbHeight)
	return s
}

// ThumbImage returns the ThumbImage property of class ScrollViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer#thumbimage
func (s *ScrollViewer) ThumbImage() *Image {
	retVal := s.p.Get("thumbImage")
	return ImageFromJSObject(retVal, s.ctx)
}

// SetThumbImage sets the ThumbImage property of class ScrollViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer#thumbimage
func (s *ScrollViewer) SetThumbImage(thumbImage *Image) *ScrollViewer {
	s.p.Set("thumbImage", thumbImage.JSObject())
	return s
}

// ThumbLength returns the ThumbLength property of class ScrollViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer#thumblength
func (s *ScrollViewer) ThumbLength() float64 {
	retVal := s.p.Get("thumbLength")
	return retVal.Float()
}

// SetThumbLength sets the ThumbLength property of class ScrollViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer#thumblength
func (s *ScrollViewer) SetThumbLength(thumbLength float64) *ScrollViewer {
	s.p.Set("thumbLength", thumbLength)
	return s
}

// VerticalBar returns the VerticalBar property of class ScrollViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer#verticalbar
func (s *ScrollViewer) VerticalBar() *ScrollBar {
	retVal := s.p.Get("verticalBar")
	return ScrollBarFromJSObject(retVal, s.ctx)
}

// SetVerticalBar sets the VerticalBar property of class ScrollViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer#verticalbar
func (s *ScrollViewer) SetVerticalBar(verticalBar *ScrollBar) *ScrollViewer {
	s.p.Set("verticalBar", verticalBar.JSObject())
	return s
}

// WheelPrecision returns the WheelPrecision property of class ScrollViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer#wheelprecision
func (s *ScrollViewer) WheelPrecision() float64 {
	retVal := s.p.Get("wheelPrecision")
	return retVal.Float()
}

// SetWheelPrecision sets the WheelPrecision property of class ScrollViewer.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollviewer#wheelprecision
func (s *ScrollViewer) SetWheelPrecision(wheelPrecision float64) *ScrollViewer {
	s.p.Set("wheelPrecision", wheelPrecision)
	return s
}
