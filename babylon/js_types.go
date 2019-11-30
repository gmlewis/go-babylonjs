package babylon

import "syscall/js"

// JSString represents a string pointer for JavaScript.
type JSString string

// JSBool represents a bool pointer for JavaScript.
type JSBool bool

// JSFloat64 represents a float64 pointer for JavaScript.
type JSFloat64 float64

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
