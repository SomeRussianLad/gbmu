package ppu

import (
	"gbmu/emulation/controllers"
	"gbmu/emulation/memory"
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"
)

type testPusher struct {
	data        []uint8
	pushCounter int
}

func (tp *testPusher) push(data []uint8) bool {
	if len(tp.data) < 8 {
		tp.data = append(tp.data, data...)
		tp.pushCounter++
		return true
	}
	return false
}

func TestFetcher(t *testing.T) {
	tp := &testPusher{}

	memory := memory.NewDMGMemory()
	interrupts := controllers.NewInterrupts(memory)
	lcd := controllers.NewLCD(memory, interrupts.Request)
	fetcher := newFetcher(memory, lcd, tp)

	min, max := image.Point{0, 0}, image.Point{160, 16}
	rectangle := image.Rectangle{min, max}
	palette := color.Palette{color0, color1, color2, color3}
	image := image.NewPaletted(rectangle, palette)

	var scanline []uint8

	memory.Write(controllers.ADDR_LCD_CONTROL, (1 << 4))
	memory.Write(controllers.ADDR_LCD_SCX, 0)
	memory.Write(controllers.ADDR_LCD_SCY, 0)

	for i, v := range vramDump {
		memory.Write(0x8000+uint16(i), v)
	}
	for i := 0; i < 0xA000-0x9800; i++ {
		memory.Write(0x9800+uint16(i), uint8(i))
	}

	for y := 0; y < 16; y++ {
		fetcher.init(0, uint8(y), TILE_TYPE_BG)

		for tp.pushCounter < 20 {
			fetcher.update()
			// t.Log(fetcher.tileMapX, fetcher.tileMapY)

			if len(tp.data) == 8 {
				scanline = append(scanline, tp.data...)
				tp.data = make([]uint8, 0, 8)
			}
		}

		for x, v := range scanline {
			image.SetColorIndex(x, y, v)
		}

		scanline = make([]uint8, 0, 160)
		tp.data = make([]uint8, 0, 8)
		tp.pushCounter = 0

		lcd.IncLY()
	}

	f, err := os.Create("image.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	png.Encode(f, image)
}
