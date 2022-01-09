package cpu

import (
	"gbmu/emulation/controllers"
	"gbmu/emulation/memory"
)

// In Double Speed Mode the following will operate twice as fast as normal:
//     CPU (8.388608 MHz, 1 cycle = approx. 0.14us)
//     Timer and Divider Registers
//     Serial Port (Link Cable)
//     DMA Transfer to OAM

const CPU_CYCLE_LEN = 4

type CPU struct {
	memory memory.Memory

	interrupts *controllers.Interrupts
	divider    *controllers.Divider
	timer      *controllers.Timer
	dma        *controllers.DMA

	instructions instructions
	registers    *registers

	cycles   int
	isHalted bool
}

func NewCPU(
	memory memory.Memory,
	interrupts *controllers.Interrupts,
	divider *controllers.Divider,
	timer *controllers.Timer,
	dma *controllers.DMA,
) *CPU {
	cpu := &CPU{
		memory: memory,

		interrupts: interrupts,
		divider:    divider,
		timer:      timer,
		dma:        dma,

		instructions: newInstructions(),
		registers:    newRegisters(),
	}

	return cpu
}

func (c *CPU) enableHalt() {
	c.isHalted = true
}

func (c *CPU) disableHalt() {
	c.isHalted = false
}

func (c *CPU) handleInterrupt(interrupt uint8) int {
	vector := c.interrupts.GetVector(interrupt)
	addr := c.registers.getSP()
	msb := uint8(c.registers.getPC() >> 8)
	lsb := uint8(c.registers.getPC() & 0xFF)

	c.memory.Write(addr-1, msb)
	c.memory.Write(addr-2, lsb)

	c.registers.setSP(addr - 2)
	c.registers.setPC(vector)

	c.interrupts.DisableMaster()
	c.interrupts.Acknowledge(interrupt)

	return 20
}

func (c *CPU) readPC() uint8 {
	addr := c.registers.getPC()
	c.registers.incPC()
	return c.memory.Read(addr)
}

func (c *CPU) executeNextInstruction() int {
	opcode := int(c.readPC())
	if opcode == 0xCB {
		opcode = 0xCB00 + int(c.readPC())
	}

	c.instructions[opcode].exec(c)
	return c.instructions[opcode].cycles
}

func (c *CPU) getSpeedMultplier() int {
	return 1
}

func (c *CPU) Update() {
	speedMultiplier := c.getSpeedMultplier()

	c.cycles += speedMultiplier * CPU_CYCLE_LEN

	if c.cycles < 4 {
		return
	}

	if i, ok := c.interrupts.Pending(); ok {
		c.disableHalt()
		if c.interrupts.IsMasterEnabled() {
			c.cycles -= c.handleInterrupt(i)
		}
	}

	if !c.isHalted {
		c.cycles -= c.executeNextInstruction()
	}

	c.divider.Update(speedMultiplier * CPU_CYCLE_LEN)
	c.timer.Update(speedMultiplier * CPU_CYCLE_LEN)
	c.dma.Update(speedMultiplier * CPU_CYCLE_LEN)
}
