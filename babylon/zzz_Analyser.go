// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Analyser represents a babylon.js Analyser.
// Class used to work with sound analyzer using fast fourier transform (FFT)
//
// See: http://doc.babylonjs.com/how_to/playing_sounds_and_music
type Analyser struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *Analyser) JSObject() js.Value { return a.p }

// Analyser returns a Analyser JavaScript class.
func (ba *Babylon) Analyser() *Analyser {
	p := ba.ctx.Get("Analyser")
	return AnalyserFromJSObject(p, ba.ctx)
}

// AnalyserFromJSObject returns a wrapped Analyser JavaScript class.
func AnalyserFromJSObject(p js.Value, ctx js.Value) *Analyser {
	return &Analyser{p: p, ctx: ctx}
}

// AnalyserArrayToJSArray returns a JavaScript Array for the wrapped array.
func AnalyserArrayToJSArray(array []*Analyser) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewAnalyser returns a new Analyser object.
//
// https://doc.babylonjs.com/api/classes/babylon.analyser#constructor
func (ba *Babylon) NewAnalyser(scene *Scene) *Analyser {

	args := make([]interface{}, 0, 1+0)

	args = append(args, scene.JSObject())

	p := ba.ctx.Get("Analyser").New(args...)
	return AnalyserFromJSObject(p, ba.ctx)
}

// ConnectAudioNodes calls the ConnectAudioNodes method on the Analyser object.
//
// https://doc.babylonjs.com/api/classes/babylon.analyser#connectaudionodes
func (a *Analyser) ConnectAudioNodes(inputAudioNode js.Value, outputAudioNode js.Value) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, inputAudioNode)

	args = append(args, outputAudioNode)

	a.p.Call("connectAudioNodes", args...)
}

// Dispose calls the Dispose method on the Analyser object.
//
// https://doc.babylonjs.com/api/classes/babylon.analyser#dispose
func (a *Analyser) Dispose() {

	a.p.Call("dispose")
}

// DrawDebugCanvas calls the DrawDebugCanvas method on the Analyser object.
//
// https://doc.babylonjs.com/api/classes/babylon.analyser#drawdebugcanvas
func (a *Analyser) DrawDebugCanvas() {

	a.p.Call("drawDebugCanvas")
}

// GetByteFrequencyData calls the GetByteFrequencyData method on the Analyser object.
//
// https://doc.babylonjs.com/api/classes/babylon.analyser#getbytefrequencydata
func (a *Analyser) GetByteFrequencyData() js.Value {

	retVal := a.p.Call("getByteFrequencyData")
	return retVal
}

// GetByteTimeDomainData calls the GetByteTimeDomainData method on the Analyser object.
//
// https://doc.babylonjs.com/api/classes/babylon.analyser#getbytetimedomaindata
func (a *Analyser) GetByteTimeDomainData() js.Value {

	retVal := a.p.Call("getByteTimeDomainData")
	return retVal
}

// GetFloatFrequencyData calls the GetFloatFrequencyData method on the Analyser object.
//
// https://doc.babylonjs.com/api/classes/babylon.analyser#getfloatfrequencydata
func (a *Analyser) GetFloatFrequencyData() js.Value {

	retVal := a.p.Call("getFloatFrequencyData")
	return retVal
}

// GetFrequencyBinCount calls the GetFrequencyBinCount method on the Analyser object.
//
// https://doc.babylonjs.com/api/classes/babylon.analyser#getfrequencybincount
func (a *Analyser) GetFrequencyBinCount() float64 {

	retVal := a.p.Call("getFrequencyBinCount")
	return retVal.Float()
}

// StopDebugCanvas calls the StopDebugCanvas method on the Analyser object.
//
// https://doc.babylonjs.com/api/classes/babylon.analyser#stopdebugcanvas
func (a *Analyser) StopDebugCanvas() {

	a.p.Call("stopDebugCanvas")
}

// BARGRAPHAMPLITUDE returns the BARGRAPHAMPLITUDE property of class Analyser.
//
// https://doc.babylonjs.com/api/classes/babylon.analyser#bargraphamplitude
func (a *Analyser) BARGRAPHAMPLITUDE() float64 {
	retVal := a.p.Get("BARGRAPHAMPLITUDE")
	return retVal.Float()
}

// SetBARGRAPHAMPLITUDE sets the BARGRAPHAMPLITUDE property of class Analyser.
//
// https://doc.babylonjs.com/api/classes/babylon.analyser#bargraphamplitude
func (a *Analyser) SetBARGRAPHAMPLITUDE(BARGRAPHAMPLITUDE float64) *Analyser {
	a.p.Set("BARGRAPHAMPLITUDE", BARGRAPHAMPLITUDE)
	return a
}

// DEBUGCANVASPOS returns the DEBUGCANVASPOS property of class Analyser.
//
// https://doc.babylonjs.com/api/classes/babylon.analyser#debugcanvaspos
func (a *Analyser) DEBUGCANVASPOS() js.Value {
	retVal := a.p.Get("DEBUGCANVASPOS")
	return retVal
}

// SetDEBUGCANVASPOS sets the DEBUGCANVASPOS property of class Analyser.
//
// https://doc.babylonjs.com/api/classes/babylon.analyser#debugcanvaspos
func (a *Analyser) SetDEBUGCANVASPOS(DEBUGCANVASPOS js.Value) *Analyser {
	a.p.Set("DEBUGCANVASPOS", DEBUGCANVASPOS)
	return a
}

// DEBUGCANVASSIZE returns the DEBUGCANVASSIZE property of class Analyser.
//
// https://doc.babylonjs.com/api/classes/babylon.analyser#debugcanvassize
func (a *Analyser) DEBUGCANVASSIZE() js.Value {
	retVal := a.p.Get("DEBUGCANVASSIZE")
	return retVal
}

// SetDEBUGCANVASSIZE sets the DEBUGCANVASSIZE property of class Analyser.
//
// https://doc.babylonjs.com/api/classes/babylon.analyser#debugcanvassize
func (a *Analyser) SetDEBUGCANVASSIZE(DEBUGCANVASSIZE js.Value) *Analyser {
	a.p.Set("DEBUGCANVASSIZE", DEBUGCANVASSIZE)
	return a
}

// FFT_SIZE returns the FFT_SIZE property of class Analyser.
//
// https://doc.babylonjs.com/api/classes/babylon.analyser#fft_size
func (a *Analyser) FFT_SIZE() float64 {
	retVal := a.p.Get("FFT_SIZE")
	return retVal.Float()
}

// SetFFT_SIZE sets the FFT_SIZE property of class Analyser.
//
// https://doc.babylonjs.com/api/classes/babylon.analyser#fft_size
func (a *Analyser) SetFFT_SIZE(FFT_SIZE float64) *Analyser {
	a.p.Set("FFT_SIZE", FFT_SIZE)
	return a
}

// SMOOTHING returns the SMOOTHING property of class Analyser.
//
// https://doc.babylonjs.com/api/classes/babylon.analyser#smoothing
func (a *Analyser) SMOOTHING() float64 {
	retVal := a.p.Get("SMOOTHING")
	return retVal.Float()
}

// SetSMOOTHING sets the SMOOTHING property of class Analyser.
//
// https://doc.babylonjs.com/api/classes/babylon.analyser#smoothing
func (a *Analyser) SetSMOOTHING(SMOOTHING float64) *Analyser {
	a.p.Set("SMOOTHING", SMOOTHING)
	return a
}
