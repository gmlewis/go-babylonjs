// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GUI3DManager represents a babylon.js GUI3DManager.
// Class used to manage 3D user interface
//
// See: http://doc.babylonjs.com/how_to/gui3d
type GUI3DManager struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (g *GUI3DManager) JSObject() js.Value { return g.p }

// GUI3DManager returns a GUI3DManager JavaScript class.
func (gui *GUI) GUI3DManager() *GUI3DManager {
	p := gui.ctx.Get("GUI3DManager")
	return GUI3DManagerFromJSObject(p, gui.ctx)
}

// GUI3DManagerFromJSObject returns a wrapped GUI3DManager JavaScript class.
func GUI3DManagerFromJSObject(p js.Value, ctx js.Value) *GUI3DManager {
	return &GUI3DManager{p: p, ctx: ctx}
}

// GUI3DManagerArrayToJSArray returns a JavaScript Array for the wrapped array.
func GUI3DManagerArrayToJSArray(array []*GUI3DManager) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewGUI3DManagerOpts contains optional parameters for NewGUI3DManager.
type NewGUI3DManagerOpts struct {
	Scene *Scene
}

// NewGUI3DManager returns a new GUI3DManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.gui3dmanager#constructor
func (gui *GUI) NewGUI3DManager(opts *NewGUI3DManagerOpts) *GUI3DManager {
	if opts == nil {
		opts = &NewGUI3DManagerOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	p := gui.ctx.Get("GUI3DManager").New(args...)
	return GUI3DManagerFromJSObject(p, gui.ctx)
}

// AddControl calls the AddControl method on the GUI3DManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.gui3dmanager#addcontrol
func (g *GUI3DManager) AddControl(control *Control3D) *GUI3DManager {

	args := make([]interface{}, 0, 1+0)

	if control == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, control.JSObject())
	}

	retVal := g.p.Call("addControl", args...)
	return GUI3DManagerFromJSObject(retVal, g.ctx)
}

// ContainsControl calls the ContainsControl method on the GUI3DManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.gui3dmanager#containscontrol
func (g *GUI3DManager) ContainsControl(control *Control3D) bool {

	args := make([]interface{}, 0, 1+0)

	if control == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, control.JSObject())
	}

	retVal := g.p.Call("containsControl", args...)
	return retVal.Bool()
}

// Dispose calls the Dispose method on the GUI3DManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.gui3dmanager#dispose
func (g *GUI3DManager) Dispose() {

	g.p.Call("dispose")
}

// RemoveControl calls the RemoveControl method on the GUI3DManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.gui3dmanager#removecontrol
func (g *GUI3DManager) RemoveControl(control *Control3D) *GUI3DManager {

	args := make([]interface{}, 0, 1+0)

	if control == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, control.JSObject())
	}

	retVal := g.p.Call("removeControl", args...)
	return GUI3DManagerFromJSObject(retVal, g.ctx)
}

// OnPickedPointChangedObservable returns the OnPickedPointChangedObservable property of class GUI3DManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.gui3dmanager#onpickedpointchangedobservable
func (g *GUI3DManager) OnPickedPointChangedObservable() *Observable {
	retVal := g.p.Get("onPickedPointChangedObservable")
	return ObservableFromJSObject(retVal, g.ctx)
}

// SetOnPickedPointChangedObservable sets the OnPickedPointChangedObservable property of class GUI3DManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.gui3dmanager#onpickedpointchangedobservable
func (g *GUI3DManager) SetOnPickedPointChangedObservable(onPickedPointChangedObservable *Observable) *GUI3DManager {
	g.p.Set("onPickedPointChangedObservable", onPickedPointChangedObservable.JSObject())
	return g
}

// RootContainer returns the RootContainer property of class GUI3DManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.gui3dmanager#rootcontainer
func (g *GUI3DManager) RootContainer() *Container3D {
	retVal := g.p.Get("rootContainer")
	return Container3DFromJSObject(retVal, g.ctx)
}

// SetRootContainer sets the RootContainer property of class GUI3DManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.gui3dmanager#rootcontainer
func (g *GUI3DManager) SetRootContainer(rootContainer *Container3D) *GUI3DManager {
	g.p.Set("rootContainer", rootContainer.JSObject())
	return g
}

// Scene returns the Scene property of class GUI3DManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.gui3dmanager#scene
func (g *GUI3DManager) Scene() *Scene {
	retVal := g.p.Get("scene")
	return SceneFromJSObject(retVal, g.ctx)
}

// SetScene sets the Scene property of class GUI3DManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.gui3dmanager#scene
func (g *GUI3DManager) SetScene(scene *Scene) *GUI3DManager {
	g.p.Set("scene", scene.JSObject())
	return g
}

// UtilityLayer returns the UtilityLayer property of class GUI3DManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.gui3dmanager#utilitylayer
func (g *GUI3DManager) UtilityLayer() *UtilityLayerRenderer {
	retVal := g.p.Get("utilityLayer")
	return UtilityLayerRendererFromJSObject(retVal, g.ctx)
}

// SetUtilityLayer sets the UtilityLayer property of class GUI3DManager.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.gui3dmanager#utilitylayer
func (g *GUI3DManager) SetUtilityLayer(utilityLayer *UtilityLayerRenderer) *GUI3DManager {
	g.p.Set("utilityLayer", utilityLayer.JSObject())
	return g
}
