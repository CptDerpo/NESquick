package main

type Bus struct {
	RAM [2 * (1 << 10)]uint8 //internal ram 2KB
	//todo: add cpu 6502
	//todo: add ppu
	//todo: add apu
	//todo: add io
	//todo: ...
}

func (b *Bus) write(addr uint16, data uint8) {
	//RAM space including memory mirroring for CPU RAM
	if addr >= 0x0000 && addr < 0x2000 {
		b.RAM[addr&0x7FF] = data
	}
}

func (b *Bus) read(addr uint16) uint8 {
	if addr < 0x2000 {
		return b.RAM[addr&0x7FF]
	}
	return 0
}
