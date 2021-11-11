package main

import "fmt"

const (
	modeImmediate = iota + 1
	modeZeroPage
	modeZeroPageX
	modeAbsolute
	modeAbsoluteX
	modeAbsoluteY
)

// Status flag ids
const (
	C = iota     //Carry flag
	Z            //Zero flag
	I            //Interrupt disable flag
	D            //Decimal Mode flag
	B            //Break Command flag
	V = iota + 1 //Overflow flag
	N            //Negative flag (if bit 7 set to 1)
)

// Instruction ids
const (
	LDA = iota + 1
)

type Instruction struct {
	id       uint8 //id of inst
	opcode   uint8 //current inst opcode
	cycles   uint8 //cycles remaining for inst
	size     uint8 //size of instruction
	addrmode uint8 //addressing mode of inst
}

var InstructionTable = map[uint8]Instruction{
	0xA9: {LDA, 0xA9, 2, 2, modeImmediate},
}

type H6502 struct {
	bus *Bus //16 bit address bus

	A uint8 //Accumulator
	Y uint8 //Index Register Y
	X uint8 //Index Register X

	PC uint16 //Program Counter
	SP uint8  //Stack Pointer

	instruction *Instruction //stores the instruction
	op1         uint8        //operand for 2 byte instructions
	op2         uint16       //operand for 3 byte instructions

	STATREG uint8 //statusregister
}

func (h *H6502) Print() {
	fmt.Printf("A: %X,  Y: %X, X: %X\nPC: %X, SP: %X\nop1: %X, op2: %X, STATREG: %X\n\n", h.A, h.Y, h.X, h.PC, h.SP, h.op1, h.op2, h.STATREG)
}

func (h *H6502) setStat(register uint8, status bool) {
	if status {
		h.STATREG |= 1 << register
	} else {
		h.STATREG &^= 1 << register
	}
}

func (h *H6502) Reset() {
	h.PC = 0x0000
	h.SP = 0x00FF //do startup routine
}

func (h *H6502) Step() {
	h.ReadInstruction(h.read(h.PC))
	h.Execute()
	h.Print()
}

func (h *H6502) ReadInstruction(opcode uint8) {
	ins := InstructionTable[opcode]
	h.instruction = &ins
	if ins.size == 2 {
		h.op1 = h.read(h.PC + 1)
	} else if ins.size == 3 {
		h.op2 = h.read16(h.PC + 1)
	} else {
		h.op1, h.op2 = 0, 0
	}
}

func (h *H6502) Execute() {
	switch h.instruction.id {
	case LDA:
		h.LDA()
	}
}

func (h *H6502) connectBus(b *Bus) {
	h.bus = b
}

func (h *H6502) write(addr uint16, data uint8) {
	h.bus.write(addr, data)
}

func (h *H6502) read(addr uint16) uint8 {
	return h.bus.read(addr)
}

func (h *H6502) read16(addr uint16) uint16 {
	return h.bus.read16(addr)
}

func (h *H6502) LDA() {
	switch h.instruction.addrmode {
	case modeImmediate:
		h.A = h.op1
		h.setStat(N, true)
		h.setStat(Z, true)
	}
}
