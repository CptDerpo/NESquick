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

func (h *H6502) ADC() {
	h.Fetch()
	var temp uint16 = uint16(h.A) + uint16(h.fetched) + uint16(btou(h.getStat(C)))
	h.setStat(C, (temp&0x100) != 0)
	h.setStat(Z, (temp&0xFF) == 0)
	h.setStat(V) //overflow flag
	h.setStat(N, (temp&0x80) != 0)
	h.A = uint8(temp & 0xFF)
}

func (h *H6502) AND() {
	h.Fetch()
	h.A &= h.fetched
	h.setStat(Z, h.A == 0)
	h.setStat(N, (h.A&0x80) != 0)
}

func (h *H6502) LDA() {
	h.Fetch()
	h.A = h.fetched
	if h.A == 0 {
		h.setStat(Z, true)
	}
	if (h.A & 0x80) != 0 {
		h.setStat(N, true)
	}
}
