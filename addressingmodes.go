package main

const (
	modeIMP = iota + 1
	modeIMM
	modeZP0
	modeZPX
	modeZPY
	modeABS
	modeABX
	modeABY
	modeIND
	modeREL
)

//IMP and ACC
func (h *H6502) IMP() {
	h.fetched = h.A
	h.extracyclesaddr = false
}

func (h *H6502) IMM() {
	h.operandaddr = h.PC
	h.PC++
	h.extracyclesaddr = false
}

func (h *H6502) ZP0() {
	h.operandaddr = uint16(h.read(h.PC))
	h.PC++
	h.operandaddr &= 0x00FF
	h.extracyclesaddr = false
}

func (h *H6502) ZPX() {
	h.operandaddr = uint16(h.read(h.PC) + h.X)
	h.PC++
	h.operandaddr &= 0x00FF
	h.extracyclesaddr = false
}

func (h *H6502) ZPY() {
	h.operandaddr = uint16(h.read(h.PC) + h.Y)
	h.PC++
	h.operandaddr &= 0x00FF
	h.extracyclesaddr = false
}

func (h *H6502) ABS() {
	h.operandaddr = uint16(h.read(h.PC)) //low bytes of address
	h.PC++
	h.operandaddr |= (uint16(h.read(h.PC)) << 8) //high bytes of address
	h.PC++
	h.extracyclesaddr = false
}

//May need more cycles
func (h *H6502) ABX() {
	h.operandaddr = uint16(h.read(h.PC)) //low bytes of address
	h.PC++
	h.operandaddr |= (uint16(h.read(h.PC)) << 8) //high bytes of address
	hi := h.operandaddr & 0xFF00
	h.PC++
	h.operandaddr += uint16(h.X)

	if (h.operandaddr & 0xFF00) != (hi << 8) { //check if high bytes (page) changed or not
		h.extracyclesaddr = true
	} else {
		h.extracyclesaddr = false
	}
}

//May need more cycles
func (h *H6502) ABY() {
	h.operandaddr = uint16(h.read(h.PC)) //low bytes of address
	h.PC++
	h.operandaddr |= (uint16(h.read(h.PC)) << 8) //high bytes of address
	hi := h.operandaddr & 0xFF00
	h.PC++
	h.operandaddr += uint16(h.Y)

	if (h.operandaddr & 0xFF00) != (hi << 8) { //check if high bytes (page) changed or not
		h.extracyclesaddr = true
	} else {
		h.extracyclesaddr = false
	}
}

func (h *H6502) IND() {
	h.operandaddr = uint16(h.bus.read(h.PC))
	h.PC++
	h.operandaddr |= (uint16(h.read(h.PC)) << 8) //memory location of where memory location is

	//check if indirect jmp (XXFF)
	h.operandaddr = (uint16(h.read(h.operandaddr)) | (uint16(h.read(h.operandaddr+1)) << 8))
	h.extracyclesaddr = false
}

func (h *H6502) REL() {
	h.branchoffset = uint16(h.read(h.PC))
	h.PC++
	//careful signed, can jump forward and backwards (7th bit is sign bit)
	if (h.branchoffset & 0x80) > 0 { //if signed byte
		h.branchoffset |= 0xFF00
	}
	h.extracyclesaddr = false
}
