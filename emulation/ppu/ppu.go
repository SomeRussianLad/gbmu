package ppu

import (
	"gbmu/emulation/controllers"
	"gbmu/emulation/memory"
)

// freq: 4194304 Hz
// pixels: 70224
// pixels per scanline: 456
// fps: ~59.7 (freq / pixels)

type PPU struct {
	lcd *controllers.LCD

	memory memory.Memory
}

func NewPPU(memory memory.Memory, lcd *controllers.LCD) *PPU {
	ppu := &PPU{
		memory: memory,
		lcd:    lcd,
	}

	return ppu
}

func (p *PPU) update() {}

func (p *PPU) Run() {
	for {
		p.update()
	}
}
