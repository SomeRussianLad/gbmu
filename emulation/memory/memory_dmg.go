package memory

type DMGMemory struct {
	memory         [0x10000]uint8
	getterHandlers map[uint16]func() uint8
	setterHandlers map[uint16]func(uint8)
}

func NewDMGMemory() *DMGMemory {
	return &DMGMemory{
		getterHandlers: make(map[uint16]func() uint8),
		setterHandlers: make(map[uint16]func(uint8)),
	}
}

func (m *DMGMemory) Read(addr uint16) uint8 {
	if h, ok := m.getterHandlers[addr]; ok {
		m.memory[addr] = h()
	}
	return m.memory[addr]
}

func (m *DMGMemory) Write(addr uint16, value uint8) {
	if h, ok := m.setterHandlers[addr]; ok {
		h(value)
	}
	m.memory[addr] = value
}

func (m *DMGMemory) RegisterGetter(addr uint16, handler func() uint8) {
	m.getterHandlers[addr] = handler
}

func (m *DMGMemory) RegisterSetter(addr uint16, handler func(uint8)) {
	m.setterHandlers[addr] = handler
}
