package controllers

import "gbmu/emulation/memory"

const (
	ADDR_TIM_COUNTER = uint16(0xFF05)
	ADDR_TIM_MODULO  = uint16(0xFF06)
	ADDR_TIM_CONTROL = uint16(0xFF07)
)

const (
	TIM_CLOCK_0 = 4096
	TIM_CLOCK_1 = 262144
	TIM_CLOCK_2 = 65536
	TIM_CLOCK_3 = 16384
)

var (
	timerClock = map[uint8]int{
		0: TIM_CLOCK_0,
		1: TIM_CLOCK_1,
		2: TIM_CLOCK_2,
		3: TIM_CLOCK_3,
	}

	timerThreshold = map[int]int{
		TIM_CLOCK_0: 1024,
		TIM_CLOCK_1: 16,
		TIM_CLOCK_2: 64,
		TIM_CLOCK_3: 256,
	}
)

type Timer struct {
	counter uint8 // 0xFF05
	modulo  uint8 // 0xFF06
	control uint8 // 0xFF07

	cycles int

	requestInterrupt func(uint8)
}

func NewTimer(memory memory.Memory, requestInterrupt func(uint8)) *Timer {
	timer := &Timer{
		requestInterrupt: requestInterrupt,
	}

	handlers := []struct {
		addr   uint16
		getter func() uint8
		setter func(uint8)
	}{
		{ADDR_TIM_COUNTER, timer.counterGetter(), timer.counterSetter()},
		{ADDR_TIM_MODULO, timer.moduloGetter(), timer.moduloSetter()},
		{ADDR_TIM_CONTROL, timer.controllerGetter(), timer.controlSetter()},
	}

	for _, h := range handlers {
		memory.RegisterGetter(h.addr, h.getter)
		memory.RegisterSetter(h.addr, h.setter)
	}

	return timer
}

func (t *Timer) Update(cycles int) {
	if !t.isStopped() {
		t.cycles += cycles

		for threshold := t.getThreshold(); t.cycles >= threshold; t.cycles -= threshold {
			t.counter++

			if t.counter == 0x00 {
				t.counter = t.modulo
				t.requestInterrupt(INT_TIMER)
			}
		}
	}
}

// isStopped checks if timer is stopped based on TAC.
// Returns true if timer is stopped, false otherwise.
func (t *Timer) isStopped() bool {
	return t.control&4 == 0
}

// getClock returns current timer clock based on TAC Input Clock Select.
func (t *Timer) getClock() int {
	return timerClock[t.control&3]
}

// getThreshold returns the amount of machine cycles, after reaching which the timer counter must be increased
func (t *Timer) getThreshold() int {
	return timerThreshold[t.getClock()]
}

func (t *Timer) counterGetter() func() uint8 {
	return func() uint8 {
		return t.counter
	}
}

func (t *Timer) moduloGetter() func() uint8 {
	return func() uint8 {
		return t.modulo
	}
}

func (t *Timer) controllerGetter() func() uint8 {
	return func() uint8 {
		return t.control
	}
}

func (t *Timer) counterSetter() func(uint8) {
	return func(value uint8) {
		t.counter = 0
	}
}

func (t *Timer) moduloSetter() func(uint8) {
	return func(value uint8) {
		t.modulo = value
	}
}

func (t *Timer) controlSetter() func(uint8) {
	return func(value uint8) {
		t.control = value
		// TODO(somerussianlad) Test this with the working emulator
		// Apparently the inner cycles counter should be reset only when switching to the different frequency in TAC
		t.cycles = 0
	}
}
