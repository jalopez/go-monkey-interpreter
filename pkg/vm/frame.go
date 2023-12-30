package vm

import (
	"github.com/jalopez/go-monkey-interpreter/pkg/code"
	"github.com/jalopez/go-monkey-interpreter/pkg/object"
)

// Frame holds the frame.
type Frame struct {
	fn *object.CompiledFunction
	ip int
}

// NewFrame creates a new frame.
func NewFrame(fn *object.CompiledFunction) *Frame {
	return &Frame{fn: fn, ip: -1}
}

// Instructions returns the instructions of the frame.
func (f *Frame) Instructions() code.Instructions {
	return f.fn.Instructions
}
