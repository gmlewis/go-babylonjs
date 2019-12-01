// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Logger represents a babylon.js Logger.
// Logger used througouht the application to allow configuration of
// the log level required for the messages.
type Logger struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (l *Logger) JSObject() js.Value { return l.p }

// Logger returns a Logger JavaScript class.
func (ba *Babylon) Logger() *Logger {
	p := ba.ctx.Get("Logger")
	return LoggerFromJSObject(p, ba.ctx)
}

// LoggerFromJSObject returns a wrapped Logger JavaScript class.
func LoggerFromJSObject(p js.Value, ctx js.Value) *Logger {
	return &Logger{p: p, ctx: ctx}
}

// LoggerArrayToJSArray returns a JavaScript Array for the wrapped array.
func LoggerArrayToJSArray(array []*Logger) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// ClearLogCache calls the ClearLogCache method on the Logger object.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#clearlogcache
func (l *Logger) ClearLogCache() {

	l.p.Call("ClearLogCache")
}

/*

// AllLogLevel returns the AllLogLevel property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#allloglevel
func (l *Logger) AllLogLevel(AllLogLevel float64) *Logger {
	p := ba.ctx.Get("Logger").New(AllLogLevel)
	return LoggerFromJSObject(p, ba.ctx)
}

// SetAllLogLevel sets the AllLogLevel property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#allloglevel
func (l *Logger) SetAllLogLevel(AllLogLevel float64) *Logger {
	p := ba.ctx.Get("Logger").New(AllLogLevel)
	return LoggerFromJSObject(p, ba.ctx)
}

// Error returns the Error property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#error
func (l *Logger) Error(Error func()) *Logger {
	p := ba.ctx.Get("Logger").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {Error(); return nil}))
	return LoggerFromJSObject(p, ba.ctx)
}

// SetError sets the Error property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#error
func (l *Logger) SetError(Error func()) *Logger {
	p := ba.ctx.Get("Logger").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {Error(); return nil}))
	return LoggerFromJSObject(p, ba.ctx)
}

// ErrorLogLevel returns the ErrorLogLevel property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#errorloglevel
func (l *Logger) ErrorLogLevel(ErrorLogLevel float64) *Logger {
	p := ba.ctx.Get("Logger").New(ErrorLogLevel)
	return LoggerFromJSObject(p, ba.ctx)
}

// SetErrorLogLevel sets the ErrorLogLevel property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#errorloglevel
func (l *Logger) SetErrorLogLevel(ErrorLogLevel float64) *Logger {
	p := ba.ctx.Get("Logger").New(ErrorLogLevel)
	return LoggerFromJSObject(p, ba.ctx)
}

// ErrorsCount returns the ErrorsCount property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#errorscount
func (l *Logger) ErrorsCount(errorsCount float64) *Logger {
	p := ba.ctx.Get("Logger").New(errorsCount)
	return LoggerFromJSObject(p, ba.ctx)
}

// SetErrorsCount sets the ErrorsCount property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#errorscount
func (l *Logger) SetErrorsCount(errorsCount float64) *Logger {
	p := ba.ctx.Get("Logger").New(errorsCount)
	return LoggerFromJSObject(p, ba.ctx)
}

// Log returns the Log property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#log
func (l *Logger) Log(Log func()) *Logger {
	p := ba.ctx.Get("Logger").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {Log(); return nil}))
	return LoggerFromJSObject(p, ba.ctx)
}

// SetLog sets the Log property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#log
func (l *Logger) SetLog(Log func()) *Logger {
	p := ba.ctx.Get("Logger").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {Log(); return nil}))
	return LoggerFromJSObject(p, ba.ctx)
}

// LogCache returns the LogCache property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#logcache
func (l *Logger) LogCache(LogCache string) *Logger {
	p := ba.ctx.Get("Logger").New(LogCache)
	return LoggerFromJSObject(p, ba.ctx)
}

// SetLogCache sets the LogCache property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#logcache
func (l *Logger) SetLogCache(LogCache string) *Logger {
	p := ba.ctx.Get("Logger").New(LogCache)
	return LoggerFromJSObject(p, ba.ctx)
}

// LogLevels returns the LogLevels property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#loglevels
func (l *Logger) LogLevels(LogLevels float64) *Logger {
	p := ba.ctx.Get("Logger").New(LogLevels)
	return LoggerFromJSObject(p, ba.ctx)
}

// SetLogLevels sets the LogLevels property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#loglevels
func (l *Logger) SetLogLevels(LogLevels float64) *Logger {
	p := ba.ctx.Get("Logger").New(LogLevels)
	return LoggerFromJSObject(p, ba.ctx)
}

// MessageLogLevel returns the MessageLogLevel property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#messageloglevel
func (l *Logger) MessageLogLevel(MessageLogLevel float64) *Logger {
	p := ba.ctx.Get("Logger").New(MessageLogLevel)
	return LoggerFromJSObject(p, ba.ctx)
}

// SetMessageLogLevel sets the MessageLogLevel property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#messageloglevel
func (l *Logger) SetMessageLogLevel(MessageLogLevel float64) *Logger {
	p := ba.ctx.Get("Logger").New(MessageLogLevel)
	return LoggerFromJSObject(p, ba.ctx)
}

// NoneLogLevel returns the NoneLogLevel property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#noneloglevel
func (l *Logger) NoneLogLevel(NoneLogLevel float64) *Logger {
	p := ba.ctx.Get("Logger").New(NoneLogLevel)
	return LoggerFromJSObject(p, ba.ctx)
}

// SetNoneLogLevel sets the NoneLogLevel property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#noneloglevel
func (l *Logger) SetNoneLogLevel(NoneLogLevel float64) *Logger {
	p := ba.ctx.Get("Logger").New(NoneLogLevel)
	return LoggerFromJSObject(p, ba.ctx)
}

// OnNewCacheEntry returns the OnNewCacheEntry property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#onnewcacheentry
func (l *Logger) OnNewCacheEntry(OnNewCacheEntry func()) *Logger {
	p := ba.ctx.Get("Logger").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {OnNewCacheEntry(); return nil}))
	return LoggerFromJSObject(p, ba.ctx)
}

// SetOnNewCacheEntry sets the OnNewCacheEntry property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#onnewcacheentry
func (l *Logger) SetOnNewCacheEntry(OnNewCacheEntry func()) *Logger {
	p := ba.ctx.Get("Logger").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {OnNewCacheEntry(); return nil}))
	return LoggerFromJSObject(p, ba.ctx)
}

// Warn returns the Warn property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#warn
func (l *Logger) Warn(Warn func()) *Logger {
	p := ba.ctx.Get("Logger").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {Warn(); return nil}))
	return LoggerFromJSObject(p, ba.ctx)
}

// SetWarn sets the Warn property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#warn
func (l *Logger) SetWarn(Warn func()) *Logger {
	p := ba.ctx.Get("Logger").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {Warn(); return nil}))
	return LoggerFromJSObject(p, ba.ctx)
}

// WarningLogLevel returns the WarningLogLevel property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#warningloglevel
func (l *Logger) WarningLogLevel(WarningLogLevel float64) *Logger {
	p := ba.ctx.Get("Logger").New(WarningLogLevel)
	return LoggerFromJSObject(p, ba.ctx)
}

// SetWarningLogLevel sets the WarningLogLevel property of class Logger.
//
// https://doc.babylonjs.com/api/classes/babylon.logger#warningloglevel
func (l *Logger) SetWarningLogLevel(WarningLogLevel float64) *Logger {
	p := ba.ctx.Get("Logger").New(WarningLogLevel)
	return LoggerFromJSObject(p, ba.ctx)
}

*/
