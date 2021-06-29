package memory

type DMGMemory struct {
	memory [0x10000]uint8
}

func NewDMGMemory() *DMGMemory {
	return &DMGMemory{}
}

func (m *DMGMemory) Read(addr uint16) uint8 {
	return m.memory[addr]
}

func (m *DMGMemory) Write(addr uint16, value uint8) {
	m.memory[addr] = value
}
