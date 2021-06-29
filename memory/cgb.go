package memory

type CGBMemory struct {
	memory [0x10000]uint8
}

func NewCGBMemory() *CGBMemory {
	return &CGBMemory{}
}

func (m *CGBMemory) Read(addr uint16) uint8 {
	return m.memory[addr]
}

func (m *CGBMemory) Write(addr uint16, value uint8) {
	m.memory[addr] = value
}
