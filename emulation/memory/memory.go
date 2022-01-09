package memory

const (
	MEM_VRAM_TILE_DATA_0 = uint16(0x8000) // 0x8000-0x87FF
	MEM_VRAM_TILE_DATA_1 = uint16(0x8800) // 0x8800-0x8FFF
	MEM_VRAM_TILE_DATA_2 = uint16(0x9000) // 0x9000-0x97FF
	MEM_VRAM_TILE_MAP_0  = uint16(0x9800) // 0x9800-0x9BFF
	MEM_VRAM_TILE_MAP_1  = uint16(0x9C00) // 0x9C00-0x9FFF

	MEM_EXTERNAL_RAM = uint16(0xA000) // 0xA000-0xBFFF, must be locked until unlocked by MBC
)

const (
	MEM_BOOT_ROM_ENABLE = uint16(0xFF50)
)

// const (
// 	MEM_ROM_LEN = MEM_RAM - MEM_ROM
// )

// General Memory Map
// 0000-3FFF   16KB ROM Bank 00     (in cartridge, fixed at bank 00)
// 4000-7FFF   16KB ROM Bank 01..NN (in cartridge, switchable bank number)
// 8000-9FFF   8KB Video RAM (VRAM) (switchable bank 0-1 in CGB Mode)
// A000-BFFF   8KB External RAM     (in cartridge, switchable bank, if any)
// C000-CFFF   4KB Work RAM Bank 0 (WRAM)
// D000-DFFF   4KB Work RAM Bank 1 (WRAM)  (switchable bank 1-7 in CGB Mode)
// E000-FDFF   Same as C000-DDFF (ECHO)    (typically not used)
// FE00-FE9F   Sprite Attribute Table (OAM)
// FEA0-FEFF   Not Usable
// FF00-FF7F   I/O Ports
// FF80-FFFE   High RAM (HRAM)
// FFFF        Interrupt Enable Register

// TODO(somerussianlad): Add switchable Banks 0-1 for VRAM

// Locks:
// FEA0-FEFF range must be locked at all times

type Memory interface {
	Read(uint16) uint8
	Write(uint16, uint8)
	AddGetter(uint16, func() uint8)
	AddSetter(uint16, func(uint8))

	// RLock(uint16)
	// WLock(uint16)
	// RWLock(uint16)

	// RUnlock(uint16)
	// WUnlock(uint16)
	// RWUnlock(uint16)
}

// type HandlerAdder interface {
// 	AddGetter(uint16, func() uint8)
// 	AddSetter(uint16, func(uint8))
// }
