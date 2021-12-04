package controllers

import "gbmu/emulation/memory"

const ADDR_DIV_COUNTER = uint16(0xFF04)

type Divider struct {
	counter uint8 // 0xFF04

	cycles int
}

func NewDivider(memory memory.Memory) *Divider {
	divider := &Divider{}

	handlers := []struct {
		addr   uint16
		getter func() uint8
		setter func(uint8)
	}{
		{ADDR_DIV_COUNTER, divider.counterGetter(), divider.counterSetter()},
	}

	for _, h := range handlers {
		memory.RegisterGetter(h.addr, h.getter)
		memory.RegisterSetter(h.addr, h.setter)
	}

	return divider
}

func (d *Divider) Update(cycles int) {
	d.cycles += cycles

	if threshold := 256; d.cycles >= threshold { // CPU_CLOCK / DIV_CLOCK is always 256, no matter the CPU clock
		d.counter++
		d.cycles -= threshold
	}
}

func (d *Divider) counterGetter() func() uint8 {
	return func() uint8 {
		return d.counter
	}
}

func (d *Divider) counterSetter() func(uint8) {
	return func(v uint8) {
		d.counter = 0
	}
}
