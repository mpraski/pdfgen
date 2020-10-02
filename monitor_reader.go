package main

import (
	"io"
)

// MonitorWriter monitors the number of bytes written
type MonitorWriter struct {
	w writer
	t int64
}

type writer interface {
	io.Writer
	io.ReaderFrom
}

// newMonitorWriter contructs a new MonitorWriter
func newMonitorWriter(w writer) *MonitorWriter {
	return &MonitorWriter{w: w}
}

// Write {inherit}
func (w *MonitorWriter) Write(p []byte) (n int, err error) {
	n, err = w.w.Write(p)
	w.t += int64(n)
	return
}

// ReadFrom {inherit}
func (w *MonitorWriter) ReadFrom(r io.Reader) (n int64, err error) {
	n, err = w.w.ReadFrom(r)
	w.t += int64(n)
	return
}

// Total returns total bytes written
func (w *MonitorWriter) Total() int64 {
	return w.t
}
