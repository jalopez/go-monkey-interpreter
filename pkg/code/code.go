package code

import (
	"encoding/binary"
	"fmt"
)

// Instructions is a byte slice that holds the bytecode instructions.
type Instructions []byte

// Opcode is a single byte that represents a bytecode instruction.
type Opcode byte

// Constants for the opcodes.
const (
	OpConstant Opcode = iota
	OpAdd
	OpSub
	OpMul
	OpDiv
	OpPop
	OpTrue
	OpFalse
	OpEqual
	OpNotEqual
	OpGreaterThan
	OpMinus
	OpBang
	OpJumpNotTruthy
	OpJump
	OpNull
	OpGetGlobal
	OpSetGlobal
	OpArray
	OpIndex
	OpCall
	OpReturnValue
	OpReturn
	OpGetLocal
	OpSetLocal
	OpGetBuiltin
	OpClosure
	OpGetFree
	OpCurrentClosure
)

// Definition is a struct that holds the name and the number of operands for an opcode.
type Definition struct {
	Name          string
	OperandWidths []int
}

// Definitions is a map that holds the definitions for all the opcodes.
var Definitions = map[Opcode]*Definition{
	OpConstant:       {"OpConstant", []int{2}},
	OpAdd:            {"OpAdd", []int{}},
	OpSub:            {"OpSub", []int{}},
	OpMul:            {"OpMul", []int{}},
	OpDiv:            {"OpDiv", []int{}},
	OpPop:            {"OpPop", []int{}},
	OpTrue:           {"OpTrue", []int{}},
	OpFalse:          {"OpFalse", []int{}},
	OpEqual:          {"OpEqual", []int{}},
	OpNotEqual:       {"OpNotEqual", []int{}},
	OpGreaterThan:    {"OpGreaterThan", []int{}},
	OpMinus:          {"OpMinus", []int{}},
	OpBang:           {"OpBang", []int{}},
	OpJumpNotTruthy:  {"OpJumpNotTruthy", []int{2}},
	OpJump:           {"OpJump", []int{2}},
	OpNull:           {"OpNull", []int{}},
	OpGetGlobal:      {"OpGetGlobal", []int{2}},
	OpSetGlobal:      {"OpSetGlobal", []int{2}},
	OpArray:          {"OpArray", []int{2}},
	OpIndex:          {"OpIndex", []int{}},
	OpCall:           {"OpCall", []int{1}},
	OpReturnValue:    {"OpReturnValue", []int{}},
	OpReturn:         {"OpReturn", []int{}},
	OpGetLocal:       {"OpGetLocal", []int{1}},
	OpSetLocal:       {"OpSetLocal", []int{1}},
	OpGetBuiltin:     {"OpGetBuiltin", []int{1}},
	OpClosure:        {"OpClosure", []int{2, 1}},
	OpGetFree:        {"OpGetFree", []int{1}},
	OpCurrentClosure: {"OpCurrentClosure", []int{}},
}

// Lookup returns the definition for the given opcode.
func Lookup(op byte) (*Definition, error) {
	def, ok := Definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}
	return def, nil
}

// Make creates an instruction from an opcode and its operands.
func Make(op Opcode, operands ...int) []byte {
	def, ok := Definitions[op]
	if !ok {
		return []byte{}
	}

	instructionLength := 1
	for _, w := range def.OperandWidths {
		instructionLength += w
	}

	instruction := make([]byte, instructionLength)
	instruction[0] = byte(op)

	offset := 1
	for i, operand := range operands {
		width := def.OperandWidths[i]
		switch width {
		case 0:
			continue
		case 1:
			instruction[offset] = byte(operand)
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(operand))
		}
		offset += width
	}

	return instruction
}

// String returns a string representation of the bytecode instructions.
func (ins Instructions) String() string {
	var out string

	i := 0
	for i < len(ins) {
		def, err := Lookup(ins[i])
		if err != nil {
			return fmt.Sprintf("ERROR: %s\n", err)
		}

		operands, read := readOperands(def, ins[i+1:])

		out += fmt.Sprintf("%04d %s\n", i, formatInstruction(def, operands))

		i += 1 + read
	}

	return out
}

// readOperands reads the operands from the bytecode instructions.
func readOperands(def *Definition, ins Instructions) ([]int, int) {
	operands := make([]int, len(def.OperandWidths))

	offset := 0
	for i, width := range def.OperandWidths {
		switch width {
		case 0:
			continue
		case 1:
			operands[i] = int(ins[offset])
		case 2:
			operands[i] = int(ReadUint16(ins[offset:]))
		}
		offset += width
	}

	return operands, offset
}

// ReadUint16 reads a uint16 from a byte slice.
func ReadUint16(ins Instructions) uint16 {
	return binary.BigEndian.Uint16(ins)
}

// ReadUint8 reads a uint8 from a byte slice.
func ReadUint8(ins Instructions) uint8 {
	return ins[0]
}

// formatInstruction formats the instruction into a string.
func formatInstruction(def *Definition, operands []int) string {
	operandCount := len(def.OperandWidths)

	if len(operands) != operandCount {
		return fmt.Sprintf("ERROR: operand len %d does not match defined %d\n",
			len(operands), operandCount)
	}

	switch operandCount {
	case 0:
		return def.Name
	case 1:
		return fmt.Sprintf("%s %d", def.Name, operands[0])
	case 2:
		return fmt.Sprintf("%s %d %d", def.Name, operands[0], operands[1])
	}

	return fmt.Sprintf("ERROR: unhandled operandCount for %s\n", def.Name)
}
