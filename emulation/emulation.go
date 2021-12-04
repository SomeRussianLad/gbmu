package emulation

import (
	"gbmu/emulation/controllers"
	"gbmu/emulation/cpu"
	"gbmu/emulation/memory"
	"gbmu/emulation/ppu"
	"sync"
	"time"
)

const CYCLES = 4194304

const (
	MODE_0 = 16 << iota
	MODE_1
	MODE_2
	MODE_3
)

type Emulator struct {
	cpu *cpu.CPU
	ppu *ppu.PPU

	wg *sync.WaitGroup
}

func NewEmulator() *Emulator {
	emulator := &Emulator{}

	memory := memory.NewDMGMemory()

	interrupts := controllers.NewInterrupts(memory)
	divider := controllers.NewDivider(memory)
	timer := controllers.NewTimer(memory, interrupts.Request)
	dma := controllers.NewDMA(memory)
	lcd := controllers.NewLCD(memory)
	// lcd := controllers.NewLCD(memory, interrupts.Request)
	// joypad := controllers.NewJoypad(memory, interrupts.Request)
	// sdt := controllers.NewSDT(memory, interrupts.Request)

	emulator.cpu = cpu.NewCPU(memory, interrupts, divider, timer, dma)
	emulator.ppu = ppu.NewPPU(memory, lcd)

	return emulator
}

func (e *Emulator) Run() {
	quota := CYCLES / MODE_0
	period := time.Second / MODE_0

	ticker := time.NewTicker(period)

	go e.cpu.Run()
	go e.ppu.Run()

	for range ticker.C {
		// TODO(somerussianlad) Both PUs request interrupts. Make sure no race condition occurs!
		for q := quota; q > 0; q -= 4 {
			e.wg.Add(2)
			// g.cpu.nextUpdate <- struct{}
			// g.ppu.nextUpdate <- struct{}
			e.wg.Wait()
		}
		quota += CYCLES
	}
}
