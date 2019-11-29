// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PostProcess represents a babylon.js PostProcess.
// PostProcess can be used to apply a shader to a texture after it has been rendered
// See &lt;a href=&#34;https://doc.babylonjs.com/how_to/how_to_use_postprocesses&#34;&gt;https://doc.babylonjs.com/how_to/how_to_use_postprocesses&lt;/a&gt;
type PostProcess struct{}

// JSObject returns the underlying js.Value.
func (p *PostProcess) JSObject() js.Value { return p.p }

// PostProcess returns a PostProcess JavaScript class.
func (b *Babylon) PostProcess() *PostProcess {
	p := b.ctx.Get("PostProcess")
	return PostProcessFromJSObject(p)
}

// PostProcessFromJSObject returns a wrapped PostProcess JavaScript class.
func PostProcessFromJSObject(p js.Value) *PostProcess {
	return &PostProcess{p: p}
}

// NewPostProcess returns a new PostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocess
func (b *Babylon) NewPostProcess(todo parameters) *PostProcess {
	p := b.ctx.Get("PostProcess").New(todo)
	return PostProcessFromJSObject(p)
}

// TODO: methods