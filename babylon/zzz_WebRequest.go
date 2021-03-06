// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// WebRequest represents a babylon.js WebRequest.
// Extended version of XMLHttpRequest with support for customizations (headers, ...)
type WebRequest struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (w *WebRequest) JSObject() js.Value { return w.p }

// WebRequest returns a WebRequest JavaScript class.
func (ba *Babylon) WebRequest() *WebRequest {
	p := ba.ctx.Get("WebRequest")
	return WebRequestFromJSObject(p, ba.ctx)
}

// WebRequestFromJSObject returns a wrapped WebRequest JavaScript class.
func WebRequestFromJSObject(p js.Value, ctx js.Value) *WebRequest {
	return &WebRequest{p: p, ctx: ctx}
}

// WebRequestArrayToJSArray returns a JavaScript Array for the wrapped array.
func WebRequestArrayToJSArray(array []*WebRequest) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// Abort calls the Abort method on the WebRequest object.
//
// https://doc.babylonjs.com/api/classes/babylon.webrequest#abort
func (w *WebRequest) Abort() {

	w.p.Call("abort")
}

// GetResponseHeader calls the GetResponseHeader method on the WebRequest object.
//
// https://doc.babylonjs.com/api/classes/babylon.webrequest#getresponseheader
func (w *WebRequest) GetResponseHeader(name string) string {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := w.p.Call("getResponseHeader", args...)
	return retVal.String()
}

// Open calls the Open method on the WebRequest object.
//
// https://doc.babylonjs.com/api/classes/babylon.webrequest#open
func (w *WebRequest) Open(method string, url string) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, method)

	args = append(args, url)

	w.p.Call("open", args...)
}

// WebRequestSendOpts contains optional parameters for WebRequest.Send.
type WebRequestSendOpts struct {
	Body js.Value
}

// Send calls the Send method on the WebRequest object.
//
// https://doc.babylonjs.com/api/classes/babylon.webrequest#send
func (w *WebRequest) Send(opts *WebRequestSendOpts) {
	if opts == nil {
		opts = &WebRequestSendOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	args = append(args, opts.Body)

	w.p.Call("send", args...)
}

// SetRequestHeader calls the SetRequestHeader method on the WebRequest object.
//
// https://doc.babylonjs.com/api/classes/babylon.webrequest#setrequestheader
func (w *WebRequest) SetRequestHeader(name string, value string) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, name)

	args = append(args, value)

	w.p.Call("setRequestHeader", args...)
}

// CustomRequestHeaders returns the CustomRequestHeaders property of class WebRequest.
//
// https://doc.babylonjs.com/api/classes/babylon.webrequest#customrequestheaders
func (w *WebRequest) CustomRequestHeaders() js.Value {
	retVal := w.p.Get("CustomRequestHeaders")
	return retVal
}

// SetCustomRequestHeaders sets the CustomRequestHeaders property of class WebRequest.
//
// https://doc.babylonjs.com/api/classes/babylon.webrequest#customrequestheaders
func (w *WebRequest) SetCustomRequestHeaders(CustomRequestHeaders js.Value) *WebRequest {
	w.p.Set("CustomRequestHeaders", CustomRequestHeaders)
	return w
}

// CustomRequestModifiers returns the CustomRequestModifiers property of class WebRequest.
//
// https://doc.babylonjs.com/api/classes/babylon.webrequest#customrequestmodifiers
func (w *WebRequest) CustomRequestModifiers() js.Value {
	retVal := w.p.Get("CustomRequestModifiers")
	return retVal
}

// SetCustomRequestModifiers sets the CustomRequestModifiers property of class WebRequest.
//
// https://doc.babylonjs.com/api/classes/babylon.webrequest#customrequestmodifiers
func (w *WebRequest) SetCustomRequestModifiers(CustomRequestModifiers JSFunc) *WebRequest {
	w.p.Set("CustomRequestModifiers", js.FuncOf(CustomRequestModifiers))
	return w
}

// Onprogress returns the Onprogress property of class WebRequest.
//
// https://doc.babylonjs.com/api/classes/babylon.webrequest#onprogress
func (w *WebRequest) Onprogress() js.Value {
	retVal := w.p.Get("onprogress")
	return retVal
}

// SetOnprogress sets the Onprogress property of class WebRequest.
//
// https://doc.babylonjs.com/api/classes/babylon.webrequest#onprogress
func (w *WebRequest) SetOnprogress(onprogress JSFunc) *WebRequest {
	w.p.Set("onprogress", js.FuncOf(onprogress))
	return w
}

// ReadyState returns the ReadyState property of class WebRequest.
//
// https://doc.babylonjs.com/api/classes/babylon.webrequest#readystate
func (w *WebRequest) ReadyState() float64 {
	retVal := w.p.Get("readyState")
	return retVal.Float()
}

// SetReadyState sets the ReadyState property of class WebRequest.
//
// https://doc.babylonjs.com/api/classes/babylon.webrequest#readystate
func (w *WebRequest) SetReadyState(readyState float64) *WebRequest {
	w.p.Set("readyState", readyState)
	return w
}

// Response returns the Response property of class WebRequest.
//
// https://doc.babylonjs.com/api/classes/babylon.webrequest#response
func (w *WebRequest) Response() js.Value {
	retVal := w.p.Get("response")
	return retVal
}

// SetResponse sets the Response property of class WebRequest.
//
// https://doc.babylonjs.com/api/classes/babylon.webrequest#response
func (w *WebRequest) SetResponse(response JSObject) *WebRequest {
	w.p.Set("response", response.JSObject())
	return w
}

// ResponseText returns the ResponseText property of class WebRequest.
//
// https://doc.babylonjs.com/api/classes/babylon.webrequest#responsetext
func (w *WebRequest) ResponseText() string {
	retVal := w.p.Get("responseText")
	return retVal.String()
}

// SetResponseText sets the ResponseText property of class WebRequest.
//
// https://doc.babylonjs.com/api/classes/babylon.webrequest#responsetext
func (w *WebRequest) SetResponseText(responseText string) *WebRequest {
	w.p.Set("responseText", responseText)
	return w
}

// ResponseType returns the ResponseType property of class WebRequest.
//
// https://doc.babylonjs.com/api/classes/babylon.webrequest#responsetype
func (w *WebRequest) ResponseType() js.Value {
	retVal := w.p.Get("responseType")
	return retVal
}

// SetResponseType sets the ResponseType property of class WebRequest.
//
// https://doc.babylonjs.com/api/classes/babylon.webrequest#responsetype
func (w *WebRequest) SetResponseType(responseType js.Value) *WebRequest {
	w.p.Set("responseType", responseType)
	return w
}

// ResponseURL returns the ResponseURL property of class WebRequest.
//
// https://doc.babylonjs.com/api/classes/babylon.webrequest#responseurl
func (w *WebRequest) ResponseURL() string {
	retVal := w.p.Get("responseURL")
	return retVal.String()
}

// SetResponseURL sets the ResponseURL property of class WebRequest.
//
// https://doc.babylonjs.com/api/classes/babylon.webrequest#responseurl
func (w *WebRequest) SetResponseURL(responseURL string) *WebRequest {
	w.p.Set("responseURL", responseURL)
	return w
}

// Status returns the Status property of class WebRequest.
//
// https://doc.babylonjs.com/api/classes/babylon.webrequest#status
func (w *WebRequest) Status() float64 {
	retVal := w.p.Get("status")
	return retVal.Float()
}

// SetStatus sets the Status property of class WebRequest.
//
// https://doc.babylonjs.com/api/classes/babylon.webrequest#status
func (w *WebRequest) SetStatus(status float64) *WebRequest {
	w.p.Set("status", status)
	return w
}

// StatusText returns the StatusText property of class WebRequest.
//
// https://doc.babylonjs.com/api/classes/babylon.webrequest#statustext
func (w *WebRequest) StatusText() string {
	retVal := w.p.Get("statusText")
	return retVal.String()
}

// SetStatusText sets the StatusText property of class WebRequest.
//
// https://doc.babylonjs.com/api/classes/babylon.webrequest#statustext
func (w *WebRequest) SetStatusText(statusText string) *WebRequest {
	w.p.Set("statusText", statusText)
	return w
}
