package vm

import (
	"github.com/jalopez/go-monkey-interpreter/pkg/code"
	"github.com/jalopez/go-monkey-interpreter/pkg/object"
)

// Frame holds the frame.
type Frame struct {
	fn          *object.CompiledFunction
	ip          int
	basePointer int
}

// NewFrame creates a new frame.
func NewFrame(fn *object.CompiledFunction, basePointer int) *Frame {
	return &Frame{fn: fn,
		ip:          -1,
		basePointer: basePointer,
	}
}

// Instructions returns the instructions of the frame.
func (f *Frame) Instructions() code.Instructions {
	return f.fn.Instructions
}
