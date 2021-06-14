package cpu

import (
	"goboy/display"
	"goboy/joypad"
	"goboy/memory"
	"goboy/sound"
)

type CPU struct {
	Flags        Flags
	Instructions Instructions
	Registers    Registers
	Timer        Timer

	Display display.Display
	Joypad  joypad.Joypad
	Memory  memory.Memory
	Sound   sound.Sound
}

func NewCPU() CPU {
	f := NewFlags()
	i := NewInstructions()
	r := NewRegisters()
	t := NewTimer()

	return CPU{
		Flags:        f,
		Instructions: i,
		Registers:    r,
		Timer:        t,
	}
}

func (c *CPU) SetDisplay(d display.Display) {
	c.Display = d
}

func (c *CPU) SetJoypad(j joypad.Joypad) {
	c.Joypad = j
}

func (c *CPU) SetMemory(m memory.Memory) {
	c.Memory = m
}

func (c *CPU) SetSound(s sound.Sound) {
	c.Sound = s
}

func (c *CPU) ReadPC() uint8 {
	addr := c.Registers.GetPC()
	c.Registers.IncPC()
	return c.ReadMemory(addr)
}

func (c *CPU) ReadMemory(addr uint16) uint8 {
	return c.Memory.ReadMemory(addr)
}

func (c *CPU) WriteMemory(addr uint16, value uint8) {
	c.Memory.WriteMemory(addr, value)
}
