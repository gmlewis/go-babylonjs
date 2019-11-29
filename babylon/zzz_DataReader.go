// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DataReader represents a babylon.js DataReader.
// Utility class for reading from a data buffer
type DataReader struct{}

// JSObject returns the underlying js.Value.
func (d *DataReader) JSObject() js.Value { return d.p }

// DataReader returns a DataReader JavaScript class.
func (b *Babylon) DataReader() *DataReader {
	p := b.ctx.Get("DataReader")
	return DataReaderFromJSObject(p)
}

// DataReaderFromJSObject returns a wrapped DataReader JavaScript class.
func DataReaderFromJSObject(p js.Value) *DataReader {
	return &DataReader{p: p}
}

// NewDataReader returns a new DataReader object.
//
// https://doc.babylonjs.com/api/classes/babylon.datareader
func (b *Babylon) NewDataReader(todo parameters) *DataReader {
	p := b.ctx.Get("DataReader").New(todo)
	return DataReaderFromJSObject(p)
}

// TODO: methods
