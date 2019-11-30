// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AssetContainer represents a babylon.js AssetContainer.
// Container with a set of assets that can be added or removed from a scene.
type AssetContainer struct {
	*AbstractScene
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *AssetContainer) JSObject() js.Value { return a.p }

// AssetContainer returns a AssetContainer JavaScript class.
func (ba *Babylon) AssetContainer() *AssetContainer {
	p := ba.ctx.Get("AssetContainer")
	return AssetContainerFromJSObject(p, ba.ctx)
}

// AssetContainerFromJSObject returns a wrapped AssetContainer JavaScript class.
func AssetContainerFromJSObject(p js.Value, ctx js.Value) *AssetContainer {
	return &AssetContainer{AbstractScene: AbstractSceneFromJSObject(p, ctx), ctx: ctx}
}

// NewAssetContainer returns a new AssetContainer object.
//
// https://doc.babylonjs.com/api/classes/babylon.assetcontainer
func (ba *Babylon) NewAssetContainer(scene *Scene) *AssetContainer {
	p := ba.ctx.Get("AssetContainer").New(scene.JSObject())
	return AssetContainerFromJSObject(p, ba.ctx)
}

// TODO: methods
