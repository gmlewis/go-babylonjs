package babylon

import "syscall/js"

// JSString represents a string pointer for JavaScript.
type JSString string

// String returns a pointer to a JSString.
func String(s string) *JSString { v := JSString(s); return &v }

// JSObject returns the underlying js.Value.
func (j *JSString) JSObject() js.Value {
	if j == nil {
		return js.Null()
	}
	return js.ValueOf(string(*j))
}

// JSBool represents a bool pointer for JavaScript.
type JSBool bool

// Bool returns a pointer to a JSBool.
func Bool(b bool) *JSBool { v := JSBool(b); return &v }

// JSObject returns the underlying js.Value.
func (j *JSBool) JSObject() js.Value {
	if j == nil {
		return js.Null()
	}
	return js.ValueOf(bool(*j))
}

// JSFloat64 represents a float64 pointer for JavaScript.
type JSFloat64 float64

// Float64 returns a pointer to a JSFloat64.
func Float64(f float64) *JSFloat64 { v := JSFloat64(f); return &v }

// JSObject returns the underlying js.Value.
func (j *JSFloat64) JSObject() js.Value {
	if j == nil {
		return js.Null()
	}
	return js.ValueOf(float64(*j))
}

// JSValue represents a js.Value pointer for JavaScript.
type JSValue js.Value

// Value returns a pointer to a JSValue.
func Value(v js.Value) *JSValue { result := JSValue(v); return &result }

// JSObject returns the underlying js.Value.
func (j *JSValue) JSObject() js.Value {
	if j == nil {
		return js.Null()
	}
	return js.ValueOf(js.Value(*j))
}
