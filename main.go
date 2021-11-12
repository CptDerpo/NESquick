package main

func main() {
	var bus Bus
	var cpu H6502
	cpu.Print()
	cpu.connectBus(&bus)
	bus.RAM[0x0000] = 0xA9 //LDA immediate
	bus.RAM[0x0001] = 0xFF //LDA immediate value
	bus.RAM[0x0002] = 0xA9 //AND immediate
	bus.RAM[0x0003] = 0xC3 //AND immediate value
	bus.RAM[0x0004] = 0xA9
	bus.RAM[0x0005] = 0x00

	cpu.Step()
	cpu.Step()
	cpu.Step()
}
