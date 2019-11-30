package babylon

import "syscall/js"

// JSString represents a string pointer for JavaScript.
type JSString string

// String returns a pointer to a JSString.
func String(s string) *JSString { v := JSString(s); return &v }

// JSBool represents a bool pointer for JavaScript.
type JSBool bool

// Bool returns a pointer to a JSBool.
func Bool(b bool) *JSBool { v := JSBool(b); return &v }

// JSFloat64 represents a float64 pointer for JavaScript.
type JSFloat64 float64

// Float64 returns a pointer to a JSFloat64.
func Float64(f float64) *JSFloat64 { v := JSFloat64(f); return &v }

// JSObject returns the underlying js.Value.
func (j *JSString) JSObject() js.Value {
	if j == nil {
		return js.Null()
	}
	return js.ValueOf(*j)
}

// JSObject returns the underlying js.Value.
func (j *JSBool) JSObject() js.Value {
	if j == nil {
		return js.Null()
	}
	return js.ValueOf(*j)
}

// JSObject returns the underlying js.Value.
func (j *JSFloat64) JSObject() js.Value {
	if j == nil {
		return js.Null()
	}
	return js.ValueOf(*j)
}
