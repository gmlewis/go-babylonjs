// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SceneOptions represents a babylon.js SceneOptions.
// Interface defining initialization parameters for Scene class
type SceneOptions struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *SceneOptions) JSObject() js.Value { return s.p }

// SceneOptions returns a SceneOptions JavaScript class.
func (ba *Babylon) SceneOptions() *SceneOptions {
	p := ba.ctx.Get("SceneOptions")
	return SceneOptionsFromJSObject(p, ba.ctx)
}

// SceneOptionsFromJSObject returns a wrapped SceneOptions JavaScript class.
func SceneOptionsFromJSObject(p js.Value, ctx js.Value) *SceneOptions {
	return &SceneOptions{p: p, ctx: ctx}
}

// SceneOptionsArrayToJSArray returns a JavaScript Array for the wrapped array.
func SceneOptionsArrayToJSArray(array []*SceneOptions) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// UseClonedMeshhMap returns the UseClonedMeshhMap property of class SceneOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptions#useclonedmeshhmap
func (s *SceneOptions) UseClonedMeshhMap() bool {
	retVal := s.p.Get("useClonedMeshhMap")
	return retVal.Bool()
}

// SetUseClonedMeshhMap sets the UseClonedMeshhMap property of class SceneOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptions#useclonedmeshhmap
func (s *SceneOptions) SetUseClonedMeshhMap(useClonedMeshhMap bool) *SceneOptions {
	s.p.Set("useClonedMeshhMap", useClonedMeshhMap)
	return s
}

// UseGeometryUniqueIdsMap returns the UseGeometryUniqueIdsMap property of class SceneOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptions#usegeometryuniqueidsmap
func (s *SceneOptions) UseGeometryUniqueIdsMap() bool {
	retVal := s.p.Get("useGeometryUniqueIdsMap")
	return retVal.Bool()
}

// SetUseGeometryUniqueIdsMap sets the UseGeometryUniqueIdsMap property of class SceneOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptions#usegeometryuniqueidsmap
func (s *SceneOptions) SetUseGeometryUniqueIdsMap(useGeometryUniqueIdsMap bool) *SceneOptions {
	s.p.Set("useGeometryUniqueIdsMap", useGeometryUniqueIdsMap)
	return s
}

// UseMaterialMeshMap returns the UseMaterialMeshMap property of class SceneOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptions#usematerialmeshmap
func (s *SceneOptions) UseMaterialMeshMap() bool {
	retVal := s.p.Get("useMaterialMeshMap")
	return retVal.Bool()
}

// SetUseMaterialMeshMap sets the UseMaterialMeshMap property of class SceneOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptions#usematerialmeshmap
func (s *SceneOptions) SetUseMaterialMeshMap(useMaterialMeshMap bool) *SceneOptions {
	s.p.Set("useMaterialMeshMap", useMaterialMeshMap)
	return s
}

// Virtual returns the Virtual property of class SceneOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptions#virtual
func (s *SceneOptions) Virtual() bool {
	retVal := s.p.Get("virtual")
	return retVal.Bool()
}

// SetVirtual sets the Virtual property of class SceneOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptions#virtual
func (s *SceneOptions) SetVirtual(virtual bool) *SceneOptions {
	s.p.Set("virtual", virtual)
	return s
}
