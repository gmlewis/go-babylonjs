// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// StringDictionary represents a babylon.js StringDictionary.
// This class implement a typical dictionary using a string as key and the generic type T as value.
// The underlying implementation relies on an associative array to ensure the best performances.
// The value can be anything including &amp;#39;null&amp;#39; but except &amp;#39;undefined&amp;#39;
type StringDictionary struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *StringDictionary) JSObject() js.Value { return s.p }

// StringDictionary returns a StringDictionary JavaScript class.
func (ba *Babylon) StringDictionary() *StringDictionary {
	p := ba.ctx.Get("StringDictionary")
	return StringDictionaryFromJSObject(p, ba.ctx)
}

// StringDictionaryFromJSObject returns a wrapped StringDictionary JavaScript class.
func StringDictionaryFromJSObject(p js.Value, ctx js.Value) *StringDictionary {
	return &StringDictionary{p: p, ctx: ctx}
}

// Add calls the Add method on the StringDictionary object.
//
// https://doc.babylonjs.com/api/classes/babylon.stringdictionary#add
func (s *StringDictionary) Add(key string, value *T) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, key)
	args = append(args, value.JSObject())

	retVal := s.p.Call("add", args...)
	return retVal.Bool()
}

// Clear calls the Clear method on the StringDictionary object.
//
// https://doc.babylonjs.com/api/classes/babylon.stringdictionary#clear
func (s *StringDictionary) Clear() {

	args := make([]interface{}, 0, 0+0)

	s.p.Call("clear", args...)
}

// Contains calls the Contains method on the StringDictionary object.
//
// https://doc.babylonjs.com/api/classes/babylon.stringdictionary#contains
func (s *StringDictionary) Contains(key string) bool {

	args := make([]interface{}, 0, 1+0)

	args = append(args, key)

	retVal := s.p.Call("contains", args...)
	return retVal.Bool()
}

// CopyFrom calls the CopyFrom method on the StringDictionary object.
//
// https://doc.babylonjs.com/api/classes/babylon.stringdictionary#copyfrom
func (s *StringDictionary) CopyFrom(source *StringDictionary) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, source.JSObject())

	s.p.Call("copyFrom", args...)
}

// ForEach calls the ForEach method on the StringDictionary object.
//
// https://doc.babylonjs.com/api/classes/babylon.stringdictionary#foreach
func (s *StringDictionary) ForEach(callback func()) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, callback)

	s.p.Call("forEach", args...)
}

// Get calls the Get method on the StringDictionary object.
//
// https://doc.babylonjs.com/api/classes/babylon.stringdictionary#get
func (s *StringDictionary) Get(key string) *T {

	args := make([]interface{}, 0, 1+0)

	args = append(args, key)

	retVal := s.p.Call("get", args...)
	return TFromJSObject(retVal, s.ctx)
}

// GetAndRemove calls the GetAndRemove method on the StringDictionary object.
//
// https://doc.babylonjs.com/api/classes/babylon.stringdictionary#getandremove
func (s *StringDictionary) GetAndRemove(key string) *T {

	args := make([]interface{}, 0, 1+0)

	args = append(args, key)

	retVal := s.p.Call("getAndRemove", args...)
	return TFromJSObject(retVal, s.ctx)
}

// GetOrAdd calls the GetOrAdd method on the StringDictionary object.
//
// https://doc.babylonjs.com/api/classes/babylon.stringdictionary#getoradd
func (s *StringDictionary) GetOrAdd(key string, val *T) *T {

	args := make([]interface{}, 0, 2+0)

	args = append(args, key)
	args = append(args, val.JSObject())

	retVal := s.p.Call("getOrAdd", args...)
	return TFromJSObject(retVal, s.ctx)
}

// GetOrAddWithFactory calls the GetOrAddWithFactory method on the StringDictionary object.
//
// https://doc.babylonjs.com/api/classes/babylon.stringdictionary#getoraddwithfactory
func (s *StringDictionary) GetOrAddWithFactory(key string, factory func()) *T {

	args := make([]interface{}, 0, 2+0)

	args = append(args, key)
	args = append(args, factory)

	retVal := s.p.Call("getOrAddWithFactory", args...)
	return TFromJSObject(retVal, s.ctx)
}

// Remove calls the Remove method on the StringDictionary object.
//
// https://doc.babylonjs.com/api/classes/babylon.stringdictionary#remove
func (s *StringDictionary) Remove(key string) bool {

	args := make([]interface{}, 0, 1+0)

	args = append(args, key)

	retVal := s.p.Call("remove", args...)
	return retVal.Bool()
}

// Set calls the Set method on the StringDictionary object.
//
// https://doc.babylonjs.com/api/classes/babylon.stringdictionary#set
func (s *StringDictionary) Set(key string, value *T) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, key)
	args = append(args, value.JSObject())

	retVal := s.p.Call("set", args...)
	return retVal.Bool()
}

/*

// Count returns the Count property of class StringDictionary.
//
// https://doc.babylonjs.com/api/classes/babylon.stringdictionary#count
func (s *StringDictionary) Count(count float64) *StringDictionary {
	p := ba.ctx.Get("StringDictionary").New(count)
	return StringDictionaryFromJSObject(p, ba.ctx)
}

// SetCount sets the Count property of class StringDictionary.
//
// https://doc.babylonjs.com/api/classes/babylon.stringdictionary#count
func (s *StringDictionary) SetCount(count float64) *StringDictionary {
	p := ba.ctx.Get("StringDictionary").New(count)
	return StringDictionaryFromJSObject(p, ba.ctx)
}

*/
