package main

type Instruction struct {
	opcode   uint8 //current inst opcode
	addrmode uint8
	size     uint8        //size of instruction
	cycles   uint8        //cycles needed for inst
	addrfunc func(*H6502) //function for addressmode
	instfunc func(*H6502) //function for instruction
}

var InstructionTable = map[uint8]Instruction{
	//ADC
	0x69: {0x69, modeIMM, 2, 2, (*H6502).IMM, (*H6502).ADC},
	0x65: {0x69, modeZP0, 2, 3, (*H6502).ZP0, (*H6502).ADC},
	0x75: {0x75, modeZPX, 2, 4, (*H6502).ZPX, (*H6502).ADC},
	0x6D: {0x6D, modeABS, 3, 4, (*H6502).ABS, (*H6502).ADC},
	0x7D: {0x6D, modeABX, 3, 4, (*H6502).ABX, (*H6502).ADC},
	0x79: {0x79, modeABY, 3, 4, (*H6502).ABY, (*H6502).ADC},
	0x61: {0x61, modeIZX, 2, 6, (*H6502).IZX, (*H6502).ADC},
	0x71: {0x71, modeIZY, 2, 5, (*H6502).IZY, (*H6502).ADC},

	//AND
	0x29: {0x29, modeIMM, 2, 2, (*H6502).IMM, (*H6502).AND},
	0x25: {0x25, modeZP0, 2, 3, (*H6502).ZP0, (*H6502).AND},
	0x35: {0x35, modeZPX, 2, 4, (*H6502).ZPX, (*H6502).AND},
	0x2D: {0x2D, modeABS, 3, 4, (*H6502).ABS, (*H6502).AND},
	0x3D: {0x3D, modeABX, 3, 4, (*H6502).ABX, (*H6502).AND},
	0x39: {0x39, modeABY, 3, 4, (*H6502).ABY, (*H6502).AND},
	0x21: {0x21, modeIZX, 2, 6, (*H6502).IZX, (*H6502).AND},
	0x31: {0x31, modeIZY, 2, 5, (*H6502).IZY, (*H6502).AND},

	//ASL
	0x0A: {0x0A, modeIMP, 1, 2, (*H6502).IMP, (*H6502).ASL},
	0x06: {0x06, modeZP0, 2, 5, (*H6502).ZP0, (*H6502).ASL},
	0x16: {0x16, modeZPX, 2, 6, (*H6502).ZPX, (*H6502).ASL},
	0x0E: {0x0E, modeABS, 3, 6, (*H6502).ABS, (*H6502).ASL},
	0x1E: {0x6D, modeABX, 3, 7, (*H6502).ABX, (*H6502).ASL},

	//BCC
	0x90: {0x90, modeREL, 2, 2, (*H6502).REL, (*H6502).BCC},

	//BCS
	0xB0: {0xB0, modeREL, 2, 2, (*H6502).REL, (*H6502).BCS},

	//BEQ
	0xF0: {0xF0, modeREL, 2, 2, (*H6502).REL, (*H6502).BEQ},

	//BIT
	0x24: {0x24, modeZP0, 2, 3, (*H6502).ZP0, (*H6502).BIT},
	0x2C: {0x2C, modeABS, 3, 4, (*H6502).ABS, (*H6502).BIT},

	//BMI
	0x30: {0x30, modeREL, 2, 2, (*H6502).REL, (*H6502).BMI},

	//BNE
	0xD0: {0xD0, modeREL, 2, 2, (*H6502).REL, (*H6502).BNE},

	//BPL
	0x10: {0x10, modeREL, 2, 2, (*H6502).REL, (*H6502).BPL},

	//BRK
	0x00: {0x00, modeIMP, 1, 7, (*H6502).IMP, (*H6502).BRK},

	//BVC
	0x50: {0x50, modeREL, 2, 2, (*H6502).REL, (*H6502).BVC},

	//BVS
	0x70: {0x70, modeREL, 2, 2, (*H6502).REL, (*H6502).BVS},

	//CLC
	0x18: {0x18, modeIMP, 1, 2, (*H6502).IMP, (*H6502).CLC},

	//CLD
	0xD8: {0xD8, modeIMP, 1, 2, (*H6502).IMP, (*H6502).CLD},

	//CLI
	0x58: {0x58, modeIMP, 1, 2, (*H6502).IMP, (*H6502).CLI},

	//CLV
	0xB8: {0xB8, modeIMP, 1, 2, (*H6502).IMP, (*H6502).CLV},

	//CMP
	0xC9: {0xC9, modeIMM, 2, 2, (*H6502).IMM, (*H6502).CMP},
	0xC5: {0xC5, modeZP0, 2, 3, (*H6502).ZP0, (*H6502).CMP},
	0xD5: {0xD5, modeZPX, 2, 4, (*H6502).ZPX, (*H6502).CMP},
	0xCD: {0xCD, modeABS, 3, 4, (*H6502).ABS, (*H6502).CMP},
	0xDD: {0xDD, modeABX, 3, 4, (*H6502).ABX, (*H6502).CMP},
	0xD9: {0xD9, modeABY, 3, 4, (*H6502).ABY, (*H6502).CMP},
	0xC1: {0xC1, modeIZX, 2, 6, (*H6502).IZX, (*H6502).CMP},
	0xD1: {0xD1, modeIZY, 2, 5, (*H6502).IZY, (*H6502).CMP},

	//CPX
	0xE0: {0xE0, modeIMM, 2, 2, (*H6502).IMM, (*H6502).CPX},
	0xE4: {0xE4, modeZP0, 2, 3, (*H6502).ZP0, (*H6502).CPX},
	0xEC: {0xEC, modeABS, 3, 4, (*H6502).ABS, (*H6502).CPX},

	//CPY
	0xC0: {0xC0, modeIMM, 2, 2, (*H6502).IMM, (*H6502).CPY},
	0xC4: {0xC4, modeZP0, 2, 3, (*H6502).ZP0, (*H6502).CPY},
	0xCC: {0xCC, modeABS, 3, 4, (*H6502).ABS, (*H6502).CPY},

	//DEC
	0xC6: {0xC6, modeZP0, 2, 5, (*H6502).ZP0, (*H6502).DEC},
	0xD6: {0xD6, modeZPX, 2, 6, (*H6502).ZPX, (*H6502).DEC},
	0xCE: {0xCE, modeABS, 3, 6, (*H6502).ABS, (*H6502).DEC},
	0xDE: {0xCE, modeABX, 3, 7, (*H6502).ABX, (*H6502).DEC},

	//DEX
	0xCA: {0xCA, modeIMP, 1, 2, (*H6502).IMP, (*H6502).DEX},

	//DEY
	0x88: {0x88, modeIMP, 1, 2, (*H6502).IMP, (*H6502).DEY},

	//EOR
	0x49: {0xC9, modeIMM, 2, 2, (*H6502).IMM, (*H6502).EOR},
	0x45: {0xC5, modeZP0, 2, 3, (*H6502).ZP0, (*H6502).EOR},
	0x55: {0xD5, modeZPX, 2, 4, (*H6502).ZPX, (*H6502).EOR},
	0x4D: {0xCD, modeABS, 3, 4, (*H6502).ABS, (*H6502).EOR},
	0x5D: {0xDD, modeABX, 3, 4, (*H6502).ABX, (*H6502).EOR},
	0x59: {0xD9, modeABY, 3, 4, (*H6502).ABY, (*H6502).EOR},
	0x41: {0xC1, modeIZX, 2, 6, (*H6502).IZX, (*H6502).EOR},
	0x51: {0xD1, modeIZY, 2, 5, (*H6502).IZY, (*H6502).EOR},

	//INC
	0xE6: {0xC6, modeZP0, 2, 5, (*H6502).ZP0, (*H6502).INC},
	0xF6: {0xD6, modeZPX, 2, 6, (*H6502).ZPX, (*H6502).INC},
	0xEE: {0xCE, modeABS, 3, 6, (*H6502).ABS, (*H6502).INC},
	0xFE: {0xCE, modeABX, 3, 7, (*H6502).ABX, (*H6502).INC},

	//INX
	0xE8: {0xE8, modeIMP, 1, 2, (*H6502).IMP, (*H6502).INX},

	//INY
	0xC8: {0xE8, modeIMP, 1, 2, (*H6502).IMP, (*H6502).INY},

	//JMP
	0x4C: {0x4C, modeABS, 3, 3, (*H6502).ABS, (*H6502).JMP},
	0x6C: {0x6C, modeIND, 3, 5, (*H6502).IND, (*H6502).JMP},

	//JSR
	0x20: {0x20, modeABS, 1, 2, (*H6502).ABS, (*H6502).JSR},

	//LDA
	0xA9: {0xA9, modeIMM, 2, 2, (*H6502).IMM, (*H6502).LDA},
	0xA5: {0xA5, modeZP0, 2, 3, (*H6502).ZP0, (*H6502).LDA},
	0xB5: {0xB5, modeZPX, 2, 4, (*H6502).ZPX, (*H6502).LDA},
	0xAD: {0xAD, modeABS, 3, 4, (*H6502).ABS, (*H6502).LDA},
	0xBD: {0xBD, modeABX, 3, 4, (*H6502).ABX, (*H6502).LDA},
	0xB9: {0xB9, modeABY, 3, 4, (*H6502).ABY, (*H6502).LDA},
	0xA1: {0xA1, modeIZX, 2, 6, (*H6502).IZX, (*H6502).LDA},
	0xB1: {0xB1, modeIZY, 2, 5, (*H6502).IZY, (*H6502).LDA},

	//LDX
	0xA2: {0xA2, modeIMM, 2, 2, (*H6502).IMM, (*H6502).LDX},
	0xA6: {0xA6, modeZP0, 2, 3, (*H6502).ZP0, (*H6502).LDX},
	0xB6: {0xA6, modeZPY, 2, 4, (*H6502).ZPY, (*H6502).LDX},
	0xAE: {0xAE, modeABS, 3, 4, (*H6502).ABS, (*H6502).LDX},
	0xBE: {0xBE, modeABY, 3, 4, (*H6502).ABY, (*H6502).LDX},

	//LDY
	0xA0: {0xA0, modeIMM, 2, 2, (*H6502).IMM, (*H6502).LDY},
	0xA4: {0xA4, modeZP0, 2, 3, (*H6502).ZP0, (*H6502).LDY},
	0xB4: {0xB4, modeZPY, 2, 4, (*H6502).ZPY, (*H6502).LDY},
	0xAC: {0xAC, modeABS, 3, 4, (*H6502).ABS, (*H6502).LDY},
	0xBC: {0xBC, modeABY, 3, 4, (*H6502).ABY, (*H6502).LDY},
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

func (h *H6502) BIT() {
	h.Fetch()
	h.setStat(Z, (h.A&h.fetched) == 0)
	h.setStat(N, (h.fetched&0x80) > 0)
	h.setStat(V, (h.fetched&0x40) > 0)
	h.extracyclesinstr = false
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

func (h *H6502) BRK() {
	h.PC++
	h.setStat(I, true)
	h.write(0x0100+uint16(h.SP), uint8(h.PC>>8))
	h.SP--
	h.write(0x0100+uint16(h.SP), uint8(h.PC&0xFF))
	h.setStat(B, true)
	h.write(0x0100+uint16(h.SP), h.STATREG)
	h.SP--
	h.setStat(B, false)
	h.PC = (uint16(h.read(0xFFFF)) << 8) | uint16(h.read(0xFFFE))
	h.extracyclesinstr = false
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

func (h *H6502) JSR() {
	h.write((0x0100 + uint16(h.SP)), uint8((h.PC>>8)&0x00FF))
	h.SP--
	h.write((0x0100 + uint16(h.SP)), uint8(h.PC&0x00FF))
	h.SP--
	h.PC = h.operandaddr
	h.extracyclesinstr = false
}

func (h *H6502) LDA() {
	h.Fetch()
	h.A = h.fetched
	h.setStat(Z, h.A == 0)
	h.setStat(N, (h.A&0x80) > 0)
	h.extracyclesinstr = true
}

func (h *H6502) LDX() {
	h.Fetch()
	h.X = h.fetched
	h.setStat(Z, h.X == 0)
	h.setStat(N, (h.X&0x80) > 0)
	h.extracyclesinstr = true
}

func (h *H6502) LDY() {
	h.Fetch()
	h.Y = h.fetched
	h.setStat(Z, h.Y == 0)
	h.setStat(N, (h.Y&0x80) > 0)
	h.extracyclesinstr = true
}

func (h *H6502) LSR() {
	h.Fetch()
	temp := (h.fetched >> 1)
	h.setStat(C, (h.fetched&0x01) > 0)
	h.setStat(Z, temp == 0)
	h.setStat(N, (temp&0x80) > 0)

	if h.instruction.addrmode == modeIMP {
		h.A = temp
	} else {
		h.write(h.operandaddr, temp)
	}
	h.extracyclesinstr = false
}

func (h *H6502) NOP() {
	h.extracyclesinstr = false
}

func (h *H6502) ORA() {
	h.Fetch()
	h.A |= h.fetched
	h.setStat(Z, h.A == 0)
	h.setStat(N, (h.A&0x80) > 0)
	h.extracyclesinstr = true
}

func (h *H6502) PHA() {
	h.write(0x0100+uint16(h.SP), h.A)
	h.SP--
	h.extracyclesinstr = false
}

func (h *H6502) PPA() {
	h.write(0x0100+uint16(h.SP), h.STATREG)
	h.SP--
	h.extracyclesinstr = false
}

func (h *H6502) PLA() {
	h.A = h.read(uint16(h.SP))
	h.SP++
	h.setStat(Z, h.A == 0)
	h.setStat(N, (h.A&0x80) > 0)
	h.extracyclesinstr = false
}

func (h *H6502) PLP() {
	h.STATREG = h.read(uint16(h.SP))
	h.SP++
	h.extracyclesinstr = false
}

func (h *H6502) ROL() {
	h.Fetch()
	temp := (h.fetched << 1) | btou(h.getStat(C))
	h.setStat(C, (h.fetched&0x80) > 0)
	h.setStat(Z, temp == 0)
	h.setStat(N, temp&0x80 > 0)

	if h.instruction.addrmode == modeIMP {
		h.A = temp
	} else {
		h.write(h.operandaddr, temp)
	}
	h.extracyclesinstr = false
}

func (h *H6502) ROR() {
	h.Fetch()
	temp := (h.fetched >> 1) | (btou(h.getStat(C)) << 7)
	h.setStat(C, (h.fetched&0x01) > 0)
	h.setStat(Z, temp == 0)
	h.setStat(N, (temp&0x80) > 0)

	if h.instruction.addrmode == modeIMP {
		h.A = temp
	} else {
		h.write(h.operandaddr, temp)
	}
	h.extracyclesinstr = false
}

func (h *H6502) RTI() {
	h.SP++
	h.STATREG = h.read(uint16(h.SP))
	h.SP++
	lo := h.read(uint16(h.SP))
	h.SP++
	hi := h.read(uint16(h.SP))
	h.PC = (uint16(hi) << 8) | uint16(lo)
	h.extracyclesinstr = false
}

func (h *H6502) RTS() {
	h.SP++
	lo := h.read(0x0100 + uint16(h.SP))
	h.SP++
	hi := h.read(0x0100 + uint16(h.SP))
	h.PC = ((uint16(hi) << 8) | uint16(lo)) - 1
	h.extracyclesinstr = false
}

func (h *H6502) SBC() {
	h.Fetch()
	var temp uint16 = uint16(h.A) + (uint16(h.fetched) ^ 0x00FF) + uint16(btou(h.getStat(C)))
	h.setStat(C, (temp&0x100) > 0)
	h.setStat(Z, (temp&0xFF) == 0)
	h.setStat(V, ((uint16(h.A)^temp)&(^uint16(h.fetched)^temp)&0x80) > 0) //overflow flag
	h.setStat(N, (temp&0x80) > 0)
	h.A = uint8(temp & 0xFF)
	h.extracyclesinstr = true
}

func (h *H6502) SEC() {
	h.setStat(C, true)
	h.extracyclesinstr = false
}

func (h *H6502) SED() {
	h.setStat(D, true)
	h.extracyclesinstr = false
}

func (h *H6502) SEI() {
	h.setStat(I, true)
	h.extracyclesinstr = false
}

func (h *H6502) STA() {
	h.write(h.operandaddr, h.A)
	h.extracyclesinstr = false
}

func (h *H6502) STX() {
	h.write(h.operandaddr, h.X)
	h.extracyclesinstr = false
}

func (h *H6502) STY() {
	h.write(h.operandaddr, h.Y)
	h.extracyclesinstr = false
}

func (h *H6502) TAX() {
	h.X = h.A
	h.setStat(Z, h.X == 0)
	h.setStat(N, (h.X&0x80) > 0)
	h.extracyclesinstr = false
}

func (h *H6502) TAY() {
	h.Y = h.A
	h.setStat(Z, h.Y == 0)
	h.setStat(N, (h.Y&0x80) > 0)
	h.extracyclesinstr = false
}

func (h *H6502) TSX() {
	h.X = h.SP
	h.setStat(Z, h.X == 0)
	h.setStat(N, (h.X&0x80) > 0)
	h.extracyclesinstr = false
}

func (h *H6502) TSA() {
	h.A = h.X
	h.setStat(Z, h.A == 0)
	h.setStat(N, (h.A&0x80) > 0)
	h.extracyclesinstr = false
}

func (h *H6502) TXS() {
	h.SP = h.X
	h.extracyclesinstr = false
}

func (h *H6502) TYA() {
	h.A = h.Y
	h.setStat(Z, h.A == 0)
	h.setStat(N, (h.A&0x80) > 0)
	h.extracyclesinstr = false
}
