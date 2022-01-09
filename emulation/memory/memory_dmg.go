package memory

type DMGMemory struct {
	memory [0x10000]uint8

	bootROM *bootROM
	mbc     mbc

	getterHandlers map[uint16]func() uint8
	setterHandlers map[uint16]func(uint8)
}

func NewDMGMemory() *DMGMemory {
	bootROM := newBootROM()

	memory := &DMGMemory{
		// memory: func() [0x10000]uint8 {
		// 	var memory [0x10000]uint8

		// 	f, _ := os.Open("./emulation/memory/Tetris.gb")

		// 	f.Read(memory[:])
		// 	f.Close()

		// 	return memory
		// }(),
		bootROM: bootROM,

		getterHandlers: make(map[uint16]func() uint8),
		setterHandlers: make(map[uint16]func(uint8)),
	}

	return memory
}

func (m *DMGMemory) Read(addr uint16) uint8 {
	switch {
	case addr < 0x100 && m.bootROM.isEnabled():
		if m.memory[MEM_BOOT_ROM_ENABLE] != 0x00 {
			m.bootROM.disable()
			break
		}
		return m.bootROM.read(addr)
	}

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

func (m *DMGMemory) AddGetter(addr uint16, handler func() uint8) {
	m.getterHandlers[addr] = handler
}

func (m *DMGMemory) AddSetter(addr uint16, handler func(uint8)) {
	m.setterHandlers[addr] = handler
}
