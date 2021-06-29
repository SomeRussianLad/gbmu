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

func (c *CPU) readPC() uint8 {
	addr := c.Registers.GetPC()
	c.Registers.IncPC()
	return c.Memory.Read(addr)
}

func (c *CPU) read8BitOperand() uint8 {
	value := c.readPC()
	return value
}

func (c *CPU) read16BitOperand() uint16 {
	msb := uint16(c.readPC())
	lsb := uint16(c.readPC())
	value := msb<<8 | lsb
	return value
}
