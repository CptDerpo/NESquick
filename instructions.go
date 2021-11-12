package main

type Instruction struct {
	opcode   uint8 //current inst opcode
	addrmode uint8
	cycles   uint8        //cycles needed for inst
	size     uint8        //size of instruction
	addrfunc func(*H6502) //function for addressmode
	instfunc func(*H6502) //function for instruction
}

var InstructionTable = map[uint8]Instruction{
	0xA9: {0xA9, modeIMM, 2, 2, (*H6502).IMM, (*H6502).LDA},
	0xDD: {0xDD, modeIMM, 2, 2, (*H6502).IMM, (*H6502).AND},
}

func (h *H6502) ReadInstruction(opcode uint8) {
	ins := InstructionTable[opcode]
	h.instruction = &ins
}

//fetch instruction data
func (h *H6502) Fetch() {
	if h.instruction.addrmode != modeIMP {
		h.fetched = h.read(h.operandaddr)
	}
}

func (h *H6502) LDA() {
	h.Fetch()
	h.A = h.fetched
	if h.A == 0 {
		h.setStat(Z, true)
	}
	if (h.A & 0x80) == 1 {
		h.setStat(N, true)
	}
}

func (h *H6502) AND() {
	h.Fetch()
	h.A &= h.fetched
	if h.A == 0 {
		h.setStat(Z, true)
	}
	if (h.A & 0x80) == 1 {
		h.setStat(N, true)
	}
}
