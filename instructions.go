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

func (h *H6502) Fetch() {
	if h.instruction.addrmode != modeIMP {
		h.fetched = h.read(h.operandaddr)
	}
}

func (h *H6502) ADC() {
	h.Fetch()
	var temp uint16 = uint16(h.A) + uint16(h.fetched) + uint16(btou(h.getStat(C)))
	h.setStat(C, (temp&0x100) > 0)
	h.setStat(Z, (temp&0xFF) == 0)
	h.setStat(V, ((uint16(h.A)^temp)&(^uint16(h.fetched)^temp)&0x80) > 0) //overflow flag
	h.setStat(N, (temp&0x80) > 0)
	h.A = uint8(temp & 0xFF)
}

func (h *H6502) ASL() {
	h.Fetch()
	var temp uint16 = uint16(h.fetched)
	temp <<= 1
	h.setStat(C, (temp&0x0100) > 0)
	h.setStat(Z, (temp&0xFF) == 0)
	h.setStat(N, (temp&0x80) > 0)
	if h.instruction.addrmode == modeIMP {
		h.A = uint8(temp & 0xFF)
	} else {
		h.write(h.operandaddr, uint8(temp&0xFF))
	}
}

func (h *H6502) AND() {
	h.Fetch()
	h.A &= h.fetched
	h.setStat(Z, h.A == 0)
	h.setStat(N, (h.A&0x80) != 0)
}

func (h *H6502) BCC() {
	if !h.getStat(C) {
		h.cycles++
		h.operandaddr = h.PC + h.branchoffset
		if (h.PC & 0xF0) != (h.operandaddr & 0xF0) { //if page is diff
			h.cycles++
		}
		h.PC = h.operandaddr
	}
	h.extracyclesinstr = false
}

func (h *H6502) BCS() {
	if h.getStat(C) {
		h.cycles++
		h.operandaddr = h.PC + h.branchoffset
		if (h.PC & 0xF0) != (h.operandaddr & 0xF0) { //if page is diff
			h.cycles++
		}
		h.PC = h.operandaddr
	}
	h.extracyclesinstr = false
}

func (h *H6502) BEQ() {
	if h.getStat(Z) {
		h.cycles++
		h.operandaddr = h.PC + h.branchoffset
		if (h.PC & 0xF0) != (h.operandaddr & 0xF0) { //if page is diff
			h.cycles++
		}
		h.PC = h.operandaddr
	}
	h.extracyclesinstr = false
}

//????
func (h *H6502) BIT() {
}

func (h *H6502) BMI() {
	if h.getStat(N) {
		h.cycles++
		h.operandaddr = h.PC + h.branchoffset
		if (h.PC & 0xF0) != (h.operandaddr & 0xF0) { //if page is diff
			h.cycles++
		}
		h.PC = h.operandaddr
	}
	h.extracyclesinstr = false
}

func (h *H6502) BNE() {
	if !h.getStat(Z) {
		h.cycles++
		h.operandaddr = h.PC + h.branchoffset
		if (h.PC & 0xF0) != (h.operandaddr & 0xF0) { //if page is diff
			h.cycles++
		}
		h.PC = h.operandaddr
	}
	h.extracyclesinstr = false
}

func (h *H6502) BPL() {
	if !h.getStat(N) {
		h.cycles++
		h.operandaddr = h.PC + h.branchoffset
		if (h.PC & 0xF0) != (h.operandaddr & 0xF0) { //if page is diff
			h.cycles++
		}
		h.PC = h.operandaddr
	}
	h.extracyclesinstr = false
}

//????
func (h *H6502) BRK() {

}

func (h *H6502) BVC() {
	if !h.getStat(V) {
		h.cycles++
		h.operandaddr = h.PC + h.branchoffset
		if (h.PC & 0xF0) != (h.operandaddr & 0xF0) { //if page is diff
			h.cycles++
		}
		h.PC = h.operandaddr
	}
	h.extracyclesinstr = false
}

func (h *H6502) BVS() {
	if h.getStat(V) {
		h.cycles++
		h.operandaddr = h.PC + h.branchoffset
		if (h.PC & 0xF0) != (h.operandaddr & 0xF0) { //if page is diff
			h.cycles++
		}
		h.PC = h.operandaddr
	}
	h.extracyclesinstr = false
}

func (h *H6502) CLC() {
	h.setStat(Z, false)
	h.extracyclesinstr = false
}

func (h *H6502) CLD() {
	h.setStat(D, false)
	h.extracyclesinstr = false
}

func (h *H6502) CLI() {
	h.setStat(I, false)
	h.extracyclesinstr = false
}

func (h *H6502) CLV() {
	h.setStat(V, false)
	h.extracyclesinstr = false
}

func (h *H6502) CMP() {
	h.Fetch()
	temp := h.A - h.fetched
	h.setStat(C, h.A > h.fetched)
	h.setStat(Z, (temp&0xFF) == 0)
	h.setStat(N, (temp&0x80) > 0)
	h.extracyclesinstr = true
}

func (h *H6502) CPX() {
	h.Fetch()
	temp := h.X - h.fetched
	h.setStat(C, h.X > h.fetched)
	h.setStat(Z, (temp&0xFF) == 0)
	h.setStat(N, (temp&0x80) > 0)
	h.extracyclesinstr = false
}

func (h *H6502) CPY() {
	h.Fetch()
	temp := h.Y - h.fetched
	h.setStat(C, h.Y > h.fetched)
	h.setStat(Z, temp == 0)
	h.setStat(N, (temp&0x80) > 0)
	h.extracyclesinstr = false
}

func (h *H6502) DEC() {
	h.Fetch()
	temp := h.fetched - 1
	h.setStat(Z, temp == 0)
	h.setStat(N, (temp&0x80) > 0)
	h.write(h.operandaddr, temp)
	h.extracyclesinstr = false
}

func (h *H6502) DEX() {
	h.X--
	h.setStat(Z, h.X == 0)
	h.setStat(N, (h.X&0x80) > 0)
	h.extracyclesinstr = false
}

func (h *H6502) DEY() {
	h.Y--
	h.setStat(Z, h.Y == 0)
	h.setStat(N, (h.Y&0x80) > 0)
	h.extracyclesinstr = false
}

func (h *H6502) EOR() {
	h.Fetch()
	h.A ^= h.fetched
	h.setStat(Z, h.A == 0)
	h.setStat(N, (h.A&0x80) > 0)
	h.extracyclesinstr = true
}

func (h *H6502) INC() {
	h.Fetch()
	temp := h.fetched + 1
	h.setStat(Z, temp == 0)
	h.setStat(N, (temp&0x80) > 0)
	h.write(h.operandaddr, temp)
	h.extracyclesinstr = false
}

func (h *H6502) INX() {
	h.X++
	h.setStat(Z, h.X == 0)
	h.setStat(N, (h.X&0x80) > 0)
	h.extracyclesinstr = false
}

func (h *H6502) INY() {
	h.Y++
	h.setStat(Z, h.X == 0)
	h.setStat(N, (h.X&0x80) > 0)
	h.extracyclesinstr = false
}

func (h *H6502) JMP() {
	h.PC = h.operandaddr
	h.extracyclesinstr = false
}

//STACK DATA:
func (h *H6502) JSR() {

}

func (h *H6502) LDA() {
	h.Fetch()
	h.A = h.fetched
	h.setStat(Z, h.A == 0)
	h.setStat(N, (h.A&0x80) > 0)
	h.extracyclesinstr = true
}
