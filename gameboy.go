package main

import (
	"goboy/cpu"
	"goboy/display"
	"goboy/joypad"
	"goboy/memory"
	"goboy/sound"
)

type Gameboy struct {
	CPU     cpu.CPU
	Display display.Display
	Joypad  joypad.Joypad
	Memory  memory.Memory
	Sound   sound.Sound
}

func NewGameboy() Gameboy {
	c := cpu.NewCPU()
	d := display.NewDisplay()
	j := joypad.NewJoypad()
	m := memory.NewMemory()
	s := sound.NewSound()

	c.SetDisplay(d)
	c.SetJoypad(j)
	c.SetMemory(m)
	c.SetSound(s)

	return Gameboy{
		CPU:     c,
		Display: d,
		Joypad:  j,
		Memory:  m,
		Sound:   s,
	}
}

func (g *Gameboy) Launch() {
	// g.CPU.Instructions[0x06].Exec(&g.CPU)
	// g.CPU.Instructions[0x0E].Exec(&g.CPU)
	// g.CPU.Instructions[0x16].Exec(&g.CPU)
	// g.CPU.Instructions[0x1E].Exec(&g.CPU)
	// fmt.Println(g.CPU.Registers)
}

func exitGracefully() {
}
