package emulation

import (
	"gbmu/emulation/controllers"
	"gbmu/emulation/cpu"
	"gbmu/emulation/display"
	"gbmu/emulation/memory"
	"gbmu/emulation/ppu"
	"time"
)

const (
	EMULATION_CYCLES_PER_SECOND = 1024 * 64
	EMULATION_DIVIDER           = 16
)

type Emulator struct {
	display display.Display
	cpu     *cpu.CPU
	ppu     *ppu.PPU
}

func NewEmulator() *Emulator {
	memory := memory.NewDMGMemory()

	interrupts := controllers.NewInterrupts(memory)
	divider := controllers.NewDivider(memory)
	timer := controllers.NewTimer(memory, interrupts.Request)
	dma := controllers.NewDMA(memory)
	lcd := controllers.NewLCD(memory, interrupts.Request)
	joypad := controllers.NewJoypad(memory, interrupts.Request)
	sdt := controllers.NewSDT(memory, interrupts.Request)

	display := display.NewFyneGUI()
	cpu := cpu.NewCPU(memory, interrupts, divider, timer, dma)
	ppu := ppu.NewPPU(display, memory, lcd)

	emulator := &Emulator{
		display: display,
		cpu:     cpu,
		ppu:     ppu,
	}

	return emulator
}

func (e *Emulator) Run() {
	cyclesQuota := EMULATION_CYCLES_PER_SECOND / EMULATION_DIVIDER
	period := time.Second / EMULATION_DIVIDER

	go func() {
		ticker := time.NewTicker(period)
		for range ticker.C {
			for cq := cyclesQuota; cq > 0; cq -= 4 {
				e.cpu.Update()
				e.ppu.Update()
			}
		}
	}()

	e.display.Run()
}
