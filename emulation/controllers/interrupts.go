package controllers

import "gbmu/emulation/memory"

const (
	ADDR_INT_ENABLED   = uint16(0xFFFF)
	ADDR_INT_REQUESTED = uint16(0xFF0F)
)

const (
	VEC_VBLANK = uint16(0x0040)
	VEC_LCD    = uint16(0x0048)
	VEC_TIMER  = uint16(0x0050)
	VEC_SERIAL = uint16(0x0058)
	VEC_JOYPAD = uint16(0x0060)

	VEC_NIL = uint16(0x0000)
)

const (
	INT_VBLANK = uint8(1 << iota)
	INT_LCD
	INT_TIMER
	INT_SERIAL
	INT_JOYPAD

	INT_NIL = uint8(0)
)

var vectors map[uint8]uint16 = map[uint8]uint16{
	INT_VBLANK: VEC_VBLANK,
	INT_LCD:    VEC_LCD,
	INT_TIMER:  VEC_TIMER,
	INT_SERIAL: VEC_SERIAL,
	INT_JOYPAD: VEC_JOYPAD,

	INT_NIL: VEC_NIL,
}

type Interrupts struct {
	master    bool
	enabled   uint8 // 0xFFFF
	requested uint8 // 0xFF0F

	delay bool
}

func NewInterrupts(memory memory.Memory) *Interrupts {
	interrupts := &Interrupts{}

	handlers := []struct {
		addr   uint16
		getter func() uint8
		setter func(uint8)
	}{
		{ADDR_INT_ENABLED, interrupts.enabledGetter(), interrupts.enabledSetter()},
		{ADDR_INT_REQUESTED, interrupts.requestedGetter(), interrupts.requestedSetter()},
	}

	for _, h := range handlers {
		memory.AddGetter(h.addr, h.getter)
		memory.AddSetter(h.addr, h.setter)
	}

	return interrupts
}

func (i *Interrupts) IsMasterEnabled() bool {
	return i.master
}

func (i *Interrupts) IsDelayed() bool {
	return i.delay
}

func (i *Interrupts) EnableMaster() {
	i.master = true
}

func (i *Interrupts) DisableMaster() {
	i.master = false
}

func (i *Interrupts) EnableDelay() {
	i.delay = true
}

func (i *Interrupts) Request(interrupt uint8) {
	i.requested = i.requested | interrupt
}

func (i *Interrupts) Pending() (uint8, bool) {
	if i.delay {
		i.delay = false
		return 0, false
	}

	pending := i.enabled & i.requested
	switch {
	case pending&INT_VBLANK == INT_VBLANK:
		return INT_VBLANK, true

	case pending&INT_LCD == INT_LCD:
		return INT_LCD, true

	case pending&INT_TIMER == INT_TIMER:
		return INT_TIMER, true

	case pending&INT_SERIAL == INT_SERIAL:
		return INT_SERIAL, true

	case pending&INT_JOYPAD == INT_JOYPAD:
		return INT_JOYPAD, true
	}

	return 0, false
}

func (i *Interrupts) GetVector(interrupt uint8) uint16 {
	return vectors[interrupt]
}

func (i *Interrupts) Acknowledge(interrupt uint8) {
	i.requested &= 0xFF - interrupt
}

func (i *Interrupts) enabledGetter() func() uint8 {
	return func() uint8 {
		return i.enabled
	}
}

func (i *Interrupts) requestedGetter() func() uint8 {
	return func() uint8 {
		return i.requested
	}
}

func (i *Interrupts) enabledSetter() func(uint8) {
	return func(value uint8) {
		i.enabled = value
	}
}

func (i *Interrupts) requestedSetter() func(uint8) {
	return func(value uint8) {
		i.requested = value
	}
}
