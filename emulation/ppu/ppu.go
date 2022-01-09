package ppu

import (
	"gbmu/emulation/controllers"
	"gbmu/emulation/display"
	"gbmu/emulation/memory"
	"image"
	"image/color"
)

// freq:                4194304 Hz
// pixels:              70224
// pixels per scanline: 456
// fps:                 ~59.7 (freq / pixels)

// OAM entry:
// Byte 0: Position X
// Byte 1: Position Y
// Byte 2: Tile Index
// Byte 3: Flags{
// 	0: Priority: 0/1 (0: drawn over BG, 1: drawn below BG),
// 	1: FlipX: 0/1 (0: default, 1: mirror tile horizontally),
// 	2: FlipY: 0/1 (0: default, 1: mirror tile vertically),
// 	3: Palette: 0/1 (0: use palette_0, 1: use palette_1),
// }

// Tile data: 3 Maps * 128 tiles * 16 Bytes = 6144 (0x1800) Bytes
// Tile map: 2 Maps * 1024 pointers * 1 Bytes = 2048 (0x800) Bytes
// Tile map entry: 1 Byte pointer to the tile in the currently selected tile data map

const (
	PPU_STATE_OAM_SEARCH = iota
	PPU_STATE_PIXEL_TRANSFER
	PPU_STATE_HBLANK
	PPU_STATE_VBLANK
)

const (
	PPU_DOT_LEN   = 4                   // Length of a PPU dot, in cycles
	PPU_SCANL_LEN = 456 * PPU_DOT_LEN   // Length of a scanline, in cycles
	PPU_FRAME_LEN = 154 * PPU_SCANL_LEN // Length of a frame, in cycles

	PPU_OAM_SEARCH_END = 10 * PPU_DOT_LEN
)

type pusher interface {
	push([]uint8) bool
}

type PPU struct {
	display.Drawer

	memory memory.Memory

	lcd      *controllers.LCD
	fifo     *fifo
	objQueue *objQueue

	cycles int
	state  uint8

	frame *image.Paletted
}

func NewPPU(d display.Drawer, memory memory.Memory, lcd *controllers.LCD) *PPU {
	color0 := color.RGBA{231, 255, 214, 255}
	color1 := color.RGBA{136, 192, 112, 255}
	color2 := color.RGBA{52, 104, 86, 255}
	color3 := color.RGBA{8, 24, 32, 255}

	min, max := image.Point{0, 0}, image.Point{160, 144}
	rectangle := image.Rectangle{min, max}
	palette := color.Palette{color0, color1, color2, color3}
	image := image.NewPaletted(rectangle, palette)

	ppu := &PPU{
		Drawer: d,
		memory: memory,
		lcd:    lcd,
		fifo:   newFIFO(memory, lcd),
		frame:  image,
	}

	return ppu
}

func (p *PPU) addScanline(scanline []uint8, y int) {
	for x, colorID := range scanline {
		p.frame.SetColorIndex(x, y, colorID)
	}
}

func (p *PPU) performOAMSearch() {
	p.fifo.objQueue.update()

	if p.cycles == PPU_OAM_SEARCH_END {
		p.fifo.init()
		p.lcd.SetMode(controllers.LCD_MODE_PIXEL_TRANSFER)
		p.state = PPU_STATE_PIXEL_TRANSFER
	}
}

func (p *PPU) performPixelTransfer() {
	p.cycles += p.fifo.update()

	if len(p.fifo.scanline) == 160 {
		line := p.lcd.GetLY()
		p.addScanline(p.fifo.scanline, int(line))

		p.lcd.IncLY()
		p.lcd.SetMode(controllers.LCD_MODE_HBLANK)
		p.state = PPU_STATE_HBLANK
	}
}

func (p *PPU) performHBlank() {
	if p.cycles == PPU_SCANL_LEN {
		p.cycles = 0
		p.lcd.IncLY()
		if p.lcd.GetLY() == 144 {
			p.lcd.SetMode(controllers.LCD_MODE_VBLANK)
			p.state = PPU_STATE_VBLANK
		} else {
			p.Draw(p.frame)
			p.fifo.objQueue.init()
			p.lcd.SetMode(controllers.LCD_MODE_OAM_SEARCH)
			p.state = PPU_STATE_OAM_SEARCH
		}
	}
}

func (p *PPU) performVBlank() {
	if p.cycles == PPU_SCANL_LEN {
		p.cycles = 0
		p.lcd.IncLY()
		if p.lcd.GetLY() == 0 {
			p.Draw(p.frame)
			p.fifo.objQueue.init()
			p.lcd.SetMode(controllers.LCD_MODE_OAM_SEARCH)
			p.state = PPU_STATE_OAM_SEARCH
		}
	}
}

func (p *PPU) Update() {
	p.cycles += 4

	if p.lcd.IsEnabled() {
		switch p.state {
		case PPU_STATE_OAM_SEARCH:
			p.performOAMSearch()

		case PPU_STATE_PIXEL_TRANSFER:
			p.performPixelTransfer()

		case PPU_STATE_HBLANK:
			p.performHBlank()

		case PPU_STATE_VBLANK:
			p.performVBlank()
		}
	}
}
