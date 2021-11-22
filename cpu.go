package main

import "fmt"

const (
	LDA = iota + 1
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

type H6502 struct {
	bus *Bus //16 bit address bus

	A uint8 //Accumulator
	Y uint8 //Index Register Y
	X uint8 //Index Register X

	PC uint16 //Program Counter
	SP uint8  //Stack Pointer

	instruction      *Instruction //stores the instruction
	cycles           uint8        //cycles left for instruction
	fetched          uint8        //fetched data from instruction
	extracyclesaddr  bool
	extracyclesinstr bool

	operandaddr  uint16 //address of operand data location
	branchoffset uint16

	STATREG uint8 //statusregister
}

func (h *H6502) Clock() {
	if h.cycles == 0 {
		h.ReadInstruction(h.read(h.PC))
		h.PC++
		h.cycles = h.instruction.cycles

		h.instruction.addrfunc(h)
		h.instruction.instfunc(h)

		h.cycles += btou(h.extracyclesaddr) & btou(h.extracyclesinstr)
	}
	h.cycles--
}

func (h *H6502) Print() {
	fmt.Printf("A: %X,  Y: %X, X: %X\nPC: %X, SP: %X\nfetched: %X, STATREG: %X\n\n", h.A, h.Y, h.X, h.PC, h.SP, h.fetched, h.STATREG)
}

func (h *H6502) setStat(register uint8, status bool) {
	if status {
		h.STATREG |= 1 << register
	} else {
		h.STATREG &^= 1 << register
	}
}

func (h *H6502) getStat(stat uint8) bool {
	return ((h.STATREG >> stat) & 1) == 1
}

func (h *H6502) Reset() {
	h.PC = 0x0000
	h.SP = 0x00FF //do startup routine

	h.operandaddr = 0xFFFA
	lo := uint16(h.read(h.operandaddr))
	hi := uint16(h.read(h.operandaddr + 1))
	h.PC = (hi << 8) | lo

	h.A = 0
	h.X = 0
	h.Y = 0
	h.SP = 0xFD
	h.STATREG = 0

	h.operandaddr = 0
	h.branchoffset = 0
	h.fetched = 0

	h.cycles = 8
}

func (h *H6502) IRQ() {
	if h.getStat(I) {
		//push pc to stack
		h.write(0x0100+uint16(h.SP), uint8(h.PC>>8))
		h.SP--
		h.write(0x0100+uint16(h.SP), uint8(h.PC&0xFF))
		h.SP--

		//push status register to stack
		h.setStat(B, false)
		h.setStat(I, true)
		h.write(0x0100+uint16(h.SP), h.STATREG)
		h.SP--

		//read new program counter
		h.operandaddr = 0xFFFE
		lo := uint16(h.read(h.operandaddr))
		hi := uint16(h.read(h.operandaddr + 1))
		h.PC = (hi << 8) | lo

		h.cycles = 7
	}
}

func (h *H6502) NMI() {
	//same as IRQ without able to be ignored
	h.write(0x0100+uint16(h.SP), uint8(h.PC>>8))
	h.SP--
	h.write(0x0100+uint16(h.SP), uint8(h.PC&0xFF))
	h.SP--

	//push status register to stack
	h.setStat(B, false)
	h.setStat(I, true)
	h.write(0x0100+uint16(h.SP), h.STATREG)
	h.SP--

	//read new program counter
	h.operandaddr = 0xFFFA
	lo := uint16(h.read(h.operandaddr))
	hi := uint16(h.read(h.operandaddr + 1))
	h.PC = (hi << 8) | lo

	h.cycles = 7
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
