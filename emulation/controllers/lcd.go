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
}

func NewLCD(memory memory.Memory) *LCD {
	lcd := &LCD{}

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
		memory.RegisterGetter(h.addr, h.getter)
		memory.RegisterSetter(h.addr, h.setter)
	}

	return lcd
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
		l.control = value
	}
}

func (l *LCD) statusSetter() func(uint8) {
	return func(value uint8) {
		l.status = value
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
	return func(value uint8) {
		l.ly = 0
	}
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
