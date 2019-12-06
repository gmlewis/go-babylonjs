// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// InstantiatedEntries represents a babylon.js InstantiatedEntries.
// Class used to store the output of the AssetContainer.instantiateAllMeshesToScene function
type InstantiatedEntries struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *InstantiatedEntries) JSObject() js.Value { return i.p }

// InstantiatedEntries returns a InstantiatedEntries JavaScript class.
func (ba *Babylon) InstantiatedEntries() *InstantiatedEntries {
	p := ba.ctx.Get("InstantiatedEntries")
	return InstantiatedEntriesFromJSObject(p, ba.ctx)
}

// InstantiatedEntriesFromJSObject returns a wrapped InstantiatedEntries JavaScript class.
func InstantiatedEntriesFromJSObject(p js.Value, ctx js.Value) *InstantiatedEntries {
	return &InstantiatedEntries{p: p, ctx: ctx}
}

// InstantiatedEntriesArrayToJSArray returns a JavaScript Array for the wrapped array.
func InstantiatedEntriesArrayToJSArray(array []*InstantiatedEntries) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// AnimationGroups returns the AnimationGroups property of class InstantiatedEntries.
//
// https://doc.babylonjs.com/api/classes/babylon.instantiatedentries#animationgroups
func (i *InstantiatedEntries) AnimationGroups() []*AnimationGroup {
	retVal := i.p.Get("animationGroups")
	result := []*AnimationGroup{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, AnimationGroupFromJSObject(retVal.Index(ri), i.ctx))
	}
	return result
}

// SetAnimationGroups sets the AnimationGroups property of class InstantiatedEntries.
//
// https://doc.babylonjs.com/api/classes/babylon.instantiatedentries#animationgroups
func (i *InstantiatedEntries) SetAnimationGroups(animationGroups []*AnimationGroup) *InstantiatedEntries {
	i.p.Set("animationGroups", animationGroups)
	return i
}

// RootNodes returns the RootNodes property of class InstantiatedEntries.
//
// https://doc.babylonjs.com/api/classes/babylon.instantiatedentries#rootnodes
func (i *InstantiatedEntries) RootNodes() []*TransformNode {
	retVal := i.p.Get("rootNodes")
	result := []*TransformNode{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, TransformNodeFromJSObject(retVal.Index(ri), i.ctx))
	}
	return result
}

// SetRootNodes sets the RootNodes property of class InstantiatedEntries.
//
// https://doc.babylonjs.com/api/classes/babylon.instantiatedentries#rootnodes
func (i *InstantiatedEntries) SetRootNodes(rootNodes []*TransformNode) *InstantiatedEntries {
	i.p.Set("rootNodes", rootNodes)
	return i
}

// Skeletons returns the Skeletons property of class InstantiatedEntries.
//
// https://doc.babylonjs.com/api/classes/babylon.instantiatedentries#skeletons
func (i *InstantiatedEntries) Skeletons() []*Skeleton {
	retVal := i.p.Get("skeletons")
	result := []*Skeleton{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, SkeletonFromJSObject(retVal.Index(ri), i.ctx))
	}
	return result
}

// SetSkeletons sets the Skeletons property of class InstantiatedEntries.
//
// https://doc.babylonjs.com/api/classes/babylon.instantiatedentries#skeletons
func (i *InstantiatedEntries) SetSkeletons(skeletons []*Skeleton) *InstantiatedEntries {
	i.p.Set("skeletons", skeletons)
	return i
}
