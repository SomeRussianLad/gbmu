package ppu

import (
	"gbmu/emulation/controllers"
	"gbmu/emulation/memory"
)

type fifo struct {
	lcd *controllers.LCD

	fetcher *fetcher
	// objQueue *objQueue

	pixelX     uint8
	pixelY     uint8
	pixelShift uint8
	pixels     []uint8

	scanline []uint8
}

func newFIFO(memory memory.Memory, lcd *controllers.LCD) *fifo {
	fifo := &fifo{
		lcd:      lcd,
		pixels:   make([]uint8, 0, 16),
		scanline: make([]uint8, 0, 160),
	}

	fifo.fetcher = newFetcher(memory, lcd, fifo)

	return fifo
}

func (f *fifo) isWindowXYConditionMet() bool {
	return f.pixelX+7 >= f.lcd.GetWX() && f.lcd.GetLY() >= f.lcd.GetWY()
}

func (f *fifo) isWindowConditionMet() bool {
	return f.lcd.IsWinEnabled() &&
		f.lcd.IsWindowVisible() &&
		f.isWindowXYConditionMet()
}

func (f *fifo) push(pixels []uint8) bool {
	if len(f.pixels) < 8 {
		f.pixels = append(f.pixels, pixels...)
		return true
	}
	return false
}

func (f *fifo) init() {
	// f.fetcher.init(f.pixelX, f.pixelY, TILE_TYPE_BG)

	f.pixelX = 0
	f.pixelY = f.lcd.GetLY()
	f.pixelShift = f.lcd.GetSCX() % 8
	f.pixels = f.pixels[:0]

	f.scanline = f.scanline[:0]

	f.fetcher.init(f.pixelX, f.pixelY, TILE_TYPE_BG)
}

func (f *fifo) shift() uint8 {
	pixel := f.pixels[0]

	f.pixelX++
	if f.pixelShift > 0 {
		f.pixelShift--
	}

	f.pixels = f.pixels[1:]

	return pixel
}

func (f *fifo) update() int {
	f.fetcher.update()

	if f.fetcher.tileType != TILE_TYPE_WIN && f.isWindowConditionMet() {
		f.fetcher.init(f.pixelX, f.pixelY, TILE_TYPE_WIN)

		f.pixelShift = f.pixelX - f.lcd.GetWX() + 7
		f.pixels = f.pixels[:0]
	}

	if len(f.pixels) < 8 {
		return PPU_DOT_LEN
	}

	if f.pixelShift > 0 {
		f.shift()
		return PPU_DOT_LEN
	}

	f.scanline = append(f.scanline, f.shift())

	return PPU_DOT_LEN
}
