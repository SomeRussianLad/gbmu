package controllers

import "gbmu/emulation/memory"

const (
	ADDR_LCD_CONTROL = uint16(0xFF40)
	ADDR_LCD_STATUS  = uint16(0xFF41)

	ADDR_LCD_SCY = uint16(0xFF42)
	ADDR_LCD_SCX = uint16(0xFF43)
	ADDR_LCD_LY  = uint16(0xFF44)
	ADDR_LCD_LYC = uint16(0xFF45)
	ADDR_LCD_WY  = uint16(0xFF4A)
	ADDR_LCD_WX  = uint16(0xFF4B)

	ADDR_LCD_BGP  = uint16(0xFF47)
	ADDR_LCD_OBP0 = uint16(0xFF48)
	ADDR_LCD_OBP1 = uint16(0xFF49)
)

const (
	LCD_MODE_HBLANK = uint8(iota)
	LCD_MODE_VBLANK
	LCD_MODE_OAM_SEARCH
	LCD_MODE_PIXEL_TRANSFER
)

const (
	LCD_PALETTE_BGP = iota
	LCD_PALETTE_OBP0
	LCD_PALETTE_OBP1
)

type LCD struct {
	control uint8
	status  uint8

	scy uint8
	scx uint8
	ly  uint8
	lyc uint8
	wy  uint8
	wx  uint8

	bgp  uint8
	obp0 uint8
	obp1 uint8

	requestInterrupt func(uint8)
}

func NewLCD(memory memory.Memory, requestInterrupt func(uint8)) *LCD {
	lcd := &LCD{
		// control:          0x91,
		requestInterrupt: requestInterrupt,
	}

	handlers := []struct {
		addr   uint16
		getter func() uint8
		setter func(uint8)
	}{
		{ADDR_LCD_CONTROL, lcd.controlGetter(), lcd.controlSetter()},
		{ADDR_LCD_STATUS, lcd.statusGetter(), lcd.statusSetter()},

		{ADDR_LCD_SCY, lcd.scyGetter(), lcd.scySetter()},
		{ADDR_LCD_SCX, lcd.scxGetter(), lcd.scxSetter()},
		{ADDR_LCD_LY, lcd.lyGetter(), lcd.lySetter()},
		{ADDR_LCD_LYC, lcd.lycGetter(), lcd.lycSetter()},
		{ADDR_LCD_WY, lcd.wyGetter(), lcd.wySetter()},
		{ADDR_LCD_WX, lcd.wxGetter(), lcd.wxSetter()},

		{ADDR_LCD_BGP, lcd.bgpGetter(), lcd.bgpSetter()},
		{ADDR_LCD_OBP0, lcd.obp0Getter(), lcd.obp0Setter()},
		{ADDR_LCD_OBP1, lcd.obp1Getter(), lcd.obp1Setter()},
	}

	for _, h := range handlers {
		memory.AddGetter(h.addr, h.getter)
		memory.AddSetter(h.addr, h.setter)
	}

	return lcd
}

func (l *LCD) IsEnabled() bool {
	return (l.control>>7)&1 == 1
}

func (l *LCD) IsWinEnabled() bool {
	return (l.control>>5)&1 == 1
}

func (l *LCD) IsOBJEnabled() bool {
	return (l.control>>1)&1 == 1
}

// Warning: Some manuals claim that in Non-CGB Mode Bit 0 controls opaqueness of BG and Win simultaneously
func (l *LCD) IsBGEnabled() bool {
	return l.control&1 == 1
}

func (l *LCD) GetWinTileMapAddr() uint16 {
	if (l.control>>6)&1 == 1 {
		return memory.MEM_VRAM_TILE_MAP_1
	}
	return memory.MEM_VRAM_TILE_MAP_0
}

func (l *LCD) GetBGTileMapAddr() uint16 {
	if (l.control>>3)&1 == 1 {
		return memory.MEM_VRAM_TILE_MAP_1
	}
	return memory.MEM_VRAM_TILE_MAP_0
}

func (l *LCD) GetBGAndWinTileDataAddr() (uint16, bool) {
	if (l.control>>4)&1 == 1 {
		return memory.MEM_VRAM_TILE_DATA_0, false
	}
	return memory.MEM_VRAM_TILE_DATA_2, true
}

func (l *LCD) IsOBJDoubleSize() bool {
	return l.control&2 == 1
}

func (l *LCD) IsLYCInterruptSourceEnabled() bool {
	return (l.status>>6)&1 == 1
}

func (l *LCD) IsOAMInterruptSourceEnabled() bool {
	return (l.status>>5)&1 == 1
}

func (l *LCD) IsVBlankInterruptSourceEnabled() bool {
	return (l.status>>4)&1 == 1
}

func (l *LCD) IsHBlankInterruptSourceEnabled() bool {
	return (l.status>>3)&1 == 1
}

func (l *LCD) SetLYCFlag(value bool) {
	if value {
		l.status |= 4
	} else {
		l.status &= 251
	}
}

func (l *LCD) SetMode(value uint8) {
	mode := value & 3

	switch mode {
	case LCD_MODE_HBLANK:
		if l.IsHBlankInterruptSourceEnabled() {
			l.requestInterrupt(LCD_MODE_HBLANK)
		}
	case LCD_MODE_VBLANK:
		if l.IsVBlankInterruptSourceEnabled() {
			l.requestInterrupt(LCD_MODE_HBLANK)
		}
	case LCD_MODE_OAM_SEARCH:
		if l.IsOAMInterruptSourceEnabled() {
			l.requestInterrupt(LCD_MODE_HBLANK)
		}
	}

	l.status = (l.status & 252) | mode
}

func (l *LCD) GetSCY() uint8 {
	return l.scy
}

func (l *LCD) GetSCX() uint8 {
	return l.scx
}

func (l *LCD) GetLY() uint8 {
	return l.ly
}

func (l *LCD) GetWY() uint8 {
	return l.wy
}

func (l *LCD) GetWX() uint8 {
	return l.wx
}

func (l *LCD) IncLY() {
	l.ly++

	if l.IsLYCInterruptSourceEnabled() && l.ly == l.lyc {
		l.requestInterrupt(INT_LCD)
	}

	if l.ly == 154 {
		l.ly = 0
	}
}

// func (l *LCD) GetPalette(paletteType int) int {
// 	return 0
// }

func (l *LCD) IsWindowVisible() bool {
	return l.GetWY() <= 143 && l.GetWX() <= 166
}

func (l *LCD) controlGetter() func() uint8 {
	return func() uint8 {
		return l.control
	}
}

func (l *LCD) statusGetter() func() uint8 {
	return func() uint8 {
		return l.status
	}
}

func (l *LCD) scyGetter() func() uint8 {
	return func() uint8 {
		return l.scy
	}
}

func (l *LCD) scxGetter() func() uint8 {
	return func() uint8 {
		return l.scx
	}
}

func (l *LCD) lyGetter() func() uint8 {
	return func() uint8 {
		return l.ly
	}
}

func (l *LCD) lycGetter() func() uint8 {
	return func() uint8 {
		return l.lyc
	}
}

func (l *LCD) wyGetter() func() uint8 {
	return func() uint8 {
		return l.wy
	}
}

func (l *LCD) wxGetter() func() uint8 {
	return func() uint8 {
		return l.wx
	}
}

func (l *LCD) bgpGetter() func() uint8 {
	return func() uint8 {
		return l.bgp
	}
}

func (l *LCD) obp0Getter() func() uint8 {
	return func() uint8 {
		return l.obp0
	}
}

func (l *LCD) obp1Getter() func() uint8 {
	return func() uint8 {
		return l.obp1
	}
}

func (l *LCD) controlSetter() func(uint8) {
	return func(value uint8) {
		// Warning: possible bug; disabling LCDC, when enabled, resets LY!
		if l.IsEnabled() && (value>>7)&1 == 0 {
			l.ly = 0
		}
		l.control = value
	}
}

func (l *LCD) statusSetter() func(uint8) {
	return func(value uint8) {
		// Warning: possible bug; it is unknown when LYC flag must be reset!
		// Possible cases are:
		// 	when writing to STAT, no matter what:
		// 		l.status = value & 248
		// 	when setting bit 2:
		// 		if (value>>2)&1 == 1 { value &= 251 }
		// 		l.status = value & 252

		if (value>>2)&1 == 1 {
			value &= 251
		}
		l.status = value & 252
	}
}

func (l *LCD) scySetter() func(uint8) {
	return func(value uint8) {
		l.scy = value
	}
}

func (l *LCD) scxSetter() func(uint8) {
	return func(value uint8) {
		l.scx = value
	}
}

func (l *LCD) lySetter() func(uint8) {
	return func(value uint8) {}
}

func (l *LCD) lycSetter() func(uint8) {
	return func(value uint8) {
		l.lyc = value
	}
}

func (l *LCD) wySetter() func(uint8) {
	return func(value uint8) {
		l.wy = value
	}
}

func (l *LCD) wxSetter() func(uint8) {
	// Warning: possible bug; some manuals explicitly tell not to specify values 0-6 for WX
	// If window offset occurs, this setter might be the source of it
	return func(value uint8) {
		l.wx = value
	}
}

func (l *LCD) bgpSetter() func(uint8) {
	return func(value uint8) {
		l.bgp = value
	}
}

func (l *LCD) obp0Setter() func(uint8) {
	return func(value uint8) {
		l.obp0 = value
	}
}

func (l *LCD) obp1Setter() func(uint8) {
	return func(value uint8) {
		l.obp1 = value
	}
}
