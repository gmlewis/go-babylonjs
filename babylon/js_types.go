package babylon

import "syscall/js"

type JSString string

type JSBool bool

type JSFloat64 float64

func (j *JSString) JSObject() js.Value {
	if j == nil {
		return js.Null()
	}
	return js.ValueOf(*j)
}

func (j *JSBool) JSObject() js.Value {
	if j == nil {
		return js.Null()
	}
	return js.ValueOf(*j)
}

func (j *JSFloat64) JSObject() js.Value {
	if j == nil {
		return js.Null()
	}
	return js.ValueOf(*j)
}
