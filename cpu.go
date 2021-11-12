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
		h.read(h.PC)
		h.PC++
		h.cycles = h.instruction.cycles

		h.instruction.addrfunc(h)
		h.instruction.instfunc(h)
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
}

func (h *H6502) Step() {
	h.ReadInstruction(h.read(h.PC))
	h.PC++
	h.instruction.addrfunc(h)
	h.instruction.instfunc(h)
	h.Print()
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
