package vm

import (
	"github.com/jalopez/go-monkey-interpreter/pkg/code"
	"github.com/jalopez/go-monkey-interpreter/pkg/object"
)

// Frame holds the frame.
type Frame struct {
	cl          *object.Closure
	ip          int
	basePointer int
}

// NewFrame creates a new frame.
func NewFrame(cl *object.Closure, basePointer int) *Frame {
	return &Frame{
		cl:          cl,
		ip:          -1,
		basePointer: basePointer,
	}
}

// Instructions returns the instructions of the frame.
func (f *Frame) Instructions() code.Instructions {
	return f.cl.Fn.Instructions
}
