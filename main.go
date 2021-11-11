package main

func main() {
	var bus Bus
	var cpu H6502
	cpu.Print()
	cpu.connectBus(&bus)
	bus.RAM[0x0000] = 0xA9 //LDA immediate
	bus.RAM[0x0001] = 0xFF //LDA immediate value
	cpu.Step()
}
