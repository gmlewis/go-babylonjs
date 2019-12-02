// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Grid represents a babylon.js Grid.
// Class used to create a 2D grid container
type Grid struct {
	*Container
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (g *Grid) JSObject() js.Value { return g.p }

// Grid returns a Grid JavaScript class.
func (ba *Babylon) Grid() *Grid {
	p := ba.ctx.Get("Grid")
	return GridFromJSObject(p, ba.ctx)
}

// GridFromJSObject returns a wrapped Grid JavaScript class.
func GridFromJSObject(p js.Value, ctx js.Value) *Grid {
	return &Grid{Container: ContainerFromJSObject(p, ctx), ctx: ctx}
}

// GridArrayToJSArray returns a JavaScript Array for the wrapped array.
func GridArrayToJSArray(array []*Grid) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewGridOpts contains optional parameters for NewGrid.
type NewGridOpts struct {
	Name *string
}

// NewGrid returns a new Grid object.
//
// https://doc.babylonjs.com/api/classes/babylon.grid
func (ba *Babylon) NewGrid(opts *NewGridOpts) *Grid {
	if opts == nil {
		opts = &NewGridOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Name == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Name)
	}

	p := ba.ctx.Get("Grid").New(args...)
	return GridFromJSObject(p, ba.ctx)
}

// GridAddColumnDefinitionOpts contains optional parameters for Grid.AddColumnDefinition.
type GridAddColumnDefinitionOpts struct {
	IsPixel *bool
}

// AddColumnDefinition calls the AddColumnDefinition method on the Grid object.
//
// https://doc.babylonjs.com/api/classes/babylon.grid#addcolumndefinition
func (g *Grid) AddColumnDefinition(width float64, opts *GridAddColumnDefinitionOpts) *Grid {
	if opts == nil {
		opts = &GridAddColumnDefinitionOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, width)

	if opts.IsPixel == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.IsPixel)
	}

	retVal := g.p.Call("addColumnDefinition", args...)
	return GridFromJSObject(retVal, g.ctx)
}

// GridAddControlOpts contains optional parameters for Grid.AddControl.
type GridAddControlOpts struct {
	Row    *float64
	Column *float64
}

// AddControl calls the AddControl method on the Grid object.
//
// https://doc.babylonjs.com/api/classes/babylon.grid#addcontrol
func (g *Grid) AddControl(control *Control, opts *GridAddControlOpts) *Grid {
	if opts == nil {
		opts = &GridAddControlOpts{}
	}

	args := make([]interface{}, 0, 1+2)

	args = append(args, control.JSObject())

	if opts.Row == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Row)
	}
	if opts.Column == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Column)
	}

	retVal := g.p.Call("addControl", args...)
	return GridFromJSObject(retVal, g.ctx)
}

// GridAddRowDefinitionOpts contains optional parameters for Grid.AddRowDefinition.
type GridAddRowDefinitionOpts struct {
	IsPixel *bool
}

// AddRowDefinition calls the AddRowDefinition method on the Grid object.
//
// https://doc.babylonjs.com/api/classes/babylon.grid#addrowdefinition
func (g *Grid) AddRowDefinition(height float64, opts *GridAddRowDefinitionOpts) *Grid {
	if opts == nil {
		opts = &GridAddRowDefinitionOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, height)

	if opts.IsPixel == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.IsPixel)
	}

	retVal := g.p.Call("addRowDefinition", args...)
	return GridFromJSObject(retVal, g.ctx)
}

// Dispose calls the Dispose method on the Grid object.
//
// https://doc.babylonjs.com/api/classes/babylon.grid#dispose
func (g *Grid) Dispose() {

	g.p.Call("dispose")
}

// GetChildCellInfo calls the GetChildCellInfo method on the Grid object.
//
// https://doc.babylonjs.com/api/classes/babylon.grid#getchildcellinfo
func (g *Grid) GetChildCellInfo(child *Control) string {

	args := make([]interface{}, 0, 1+0)

	args = append(args, child.JSObject())

	retVal := g.p.Call("getChildCellInfo", args...)
	return retVal.String()
}

// GetChildrenAt calls the GetChildrenAt method on the Grid object.
//
// https://doc.babylonjs.com/api/classes/babylon.grid#getchildrenat
func (g *Grid) GetChildrenAt(row float64, column float64) []*Control {

	args := make([]interface{}, 0, 2+0)

	args = append(args, row)
	args = append(args, column)

	retVal := g.p.Call("getChildrenAt", args...)
	result := []*Control{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, ControlFromJSObject(retVal.Index(ri), g.ctx))
	}
	return result
}

// GetColumnDefinition calls the GetColumnDefinition method on the Grid object.
//
// https://doc.babylonjs.com/api/classes/babylon.grid#getcolumndefinition
func (g *Grid) GetColumnDefinition(index float64) *ValueAndUnit {

	args := make([]interface{}, 0, 1+0)

	args = append(args, index)

	retVal := g.p.Call("getColumnDefinition", args...)
	return ValueAndUnitFromJSObject(retVal, g.ctx)
}

// GetRowDefinition calls the GetRowDefinition method on the Grid object.
//
// https://doc.babylonjs.com/api/classes/babylon.grid#getrowdefinition
func (g *Grid) GetRowDefinition(index float64) *ValueAndUnit {

	args := make([]interface{}, 0, 1+0)

	args = append(args, index)

	retVal := g.p.Call("getRowDefinition", args...)
	return ValueAndUnitFromJSObject(retVal, g.ctx)
}

// RemoveColumnDefinition calls the RemoveColumnDefinition method on the Grid object.
//
// https://doc.babylonjs.com/api/classes/babylon.grid#removecolumndefinition
func (g *Grid) RemoveColumnDefinition(index float64) *Grid {

	args := make([]interface{}, 0, 1+0)

	args = append(args, index)

	retVal := g.p.Call("removeColumnDefinition", args...)
	return GridFromJSObject(retVal, g.ctx)
}

// RemoveControl calls the RemoveControl method on the Grid object.
//
// https://doc.babylonjs.com/api/classes/babylon.grid#removecontrol
func (g *Grid) RemoveControl(control *Control) *Container {

	args := make([]interface{}, 0, 1+0)

	args = append(args, control.JSObject())

	retVal := g.p.Call("removeControl", args...)
	return ContainerFromJSObject(retVal, g.ctx)
}

// RemoveRowDefinition calls the RemoveRowDefinition method on the Grid object.
//
// https://doc.babylonjs.com/api/classes/babylon.grid#removerowdefinition
func (g *Grid) RemoveRowDefinition(index float64) *Grid {

	args := make([]interface{}, 0, 1+0)

	args = append(args, index)

	retVal := g.p.Call("removeRowDefinition", args...)
	return GridFromJSObject(retVal, g.ctx)
}

// GridSetColumnDefinitionOpts contains optional parameters for Grid.SetColumnDefinition.
type GridSetColumnDefinitionOpts struct {
	IsPixel *bool
}

// SetColumnDefinition calls the SetColumnDefinition method on the Grid object.
//
// https://doc.babylonjs.com/api/classes/babylon.grid#setcolumndefinition
func (g *Grid) SetColumnDefinition(index float64, width float64, opts *GridSetColumnDefinitionOpts) *Grid {
	if opts == nil {
		opts = &GridSetColumnDefinitionOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, index)
	args = append(args, width)

	if opts.IsPixel == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.IsPixel)
	}

	retVal := g.p.Call("setColumnDefinition", args...)
	return GridFromJSObject(retVal, g.ctx)
}

// GridSetRowDefinitionOpts contains optional parameters for Grid.SetRowDefinition.
type GridSetRowDefinitionOpts struct {
	IsPixel *bool
}

// SetRowDefinition calls the SetRowDefinition method on the Grid object.
//
// https://doc.babylonjs.com/api/classes/babylon.grid#setrowdefinition
func (g *Grid) SetRowDefinition(index float64, height float64, opts *GridSetRowDefinitionOpts) *Grid {
	if opts == nil {
		opts = &GridSetRowDefinitionOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, index)
	args = append(args, height)

	if opts.IsPixel == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.IsPixel)
	}

	retVal := g.p.Call("setRowDefinition", args...)
	return GridFromJSObject(retVal, g.ctx)
}

// _flagDescendantsAsMatrixDirty calls the _flagDescendantsAsMatrixDirty method on the Grid object.
//
// https://doc.babylonjs.com/api/classes/babylon.grid#_flagdescendantsasmatrixdirty
func (g *Grid) _flagDescendantsAsMatrixDirty() {

	g.p.Call("_flagDescendantsAsMatrixDirty")
}

// _renderHighlightSpecific calls the _renderHighlightSpecific method on the Grid object.
//
// https://doc.babylonjs.com/api/classes/babylon.grid#_renderhighlightspecific
func (g *Grid) _renderHighlightSpecific(context js.Value) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, context)

	g.p.Call("_renderHighlightSpecific", args...)
}

/*

// Cells returns the Cells property of class Grid.
//
// https://doc.babylonjs.com/api/classes/babylon.grid#cells
func (g *Grid) Cells(cells js.Value) *Grid {
	p := ba.ctx.Get("Grid").New(cells)
	return GridFromJSObject(p, ba.ctx)
}

// SetCells sets the Cells property of class Grid.
//
// https://doc.babylonjs.com/api/classes/babylon.grid#cells
func (g *Grid) SetCells(cells js.Value) *Grid {
	p := ba.ctx.Get("Grid").New(cells)
	return GridFromJSObject(p, ba.ctx)
}

// Children returns the Children property of class Grid.
//
// https://doc.babylonjs.com/api/classes/babylon.grid#children
func (g *Grid) Children(children *Control) *Grid {
	p := ba.ctx.Get("Grid").New(children.JSObject())
	return GridFromJSObject(p, ba.ctx)
}

// SetChildren sets the Children property of class Grid.
//
// https://doc.babylonjs.com/api/classes/babylon.grid#children
func (g *Grid) SetChildren(children *Control) *Grid {
	p := ba.ctx.Get("Grid").New(children.JSObject())
	return GridFromJSObject(p, ba.ctx)
}

// ColumnCount returns the ColumnCount property of class Grid.
//
// https://doc.babylonjs.com/api/classes/babylon.grid#columncount
func (g *Grid) ColumnCount(columnCount float64) *Grid {
	p := ba.ctx.Get("Grid").New(columnCount)
	return GridFromJSObject(p, ba.ctx)
}

// SetColumnCount sets the ColumnCount property of class Grid.
//
// https://doc.babylonjs.com/api/classes/babylon.grid#columncount
func (g *Grid) SetColumnCount(columnCount float64) *Grid {
	p := ba.ctx.Get("Grid").New(columnCount)
	return GridFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class Grid.
//
// https://doc.babylonjs.com/api/classes/babylon.grid#name
func (g *Grid) Name(name string) *Grid {
	p := ba.ctx.Get("Grid").New(name)
	return GridFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class Grid.
//
// https://doc.babylonjs.com/api/classes/babylon.grid#name
func (g *Grid) SetName(name string) *Grid {
	p := ba.ctx.Get("Grid").New(name)
	return GridFromJSObject(p, ba.ctx)
}

// RowCount returns the RowCount property of class Grid.
//
// https://doc.babylonjs.com/api/classes/babylon.grid#rowcount
func (g *Grid) RowCount(rowCount float64) *Grid {
	p := ba.ctx.Get("Grid").New(rowCount)
	return GridFromJSObject(p, ba.ctx)
}

// SetRowCount sets the RowCount property of class Grid.
//
// https://doc.babylonjs.com/api/classes/babylon.grid#rowcount
func (g *Grid) SetRowCount(rowCount float64) *Grid {
	p := ba.ctx.Get("Grid").New(rowCount)
	return GridFromJSObject(p, ba.ctx)
}

*/
