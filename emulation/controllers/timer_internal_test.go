package controllers

import (
	"math/rand"
	"testing"
	"time"

	"gbmu/emulation/memory"
)

const NUM_OF_INSTRUCTIONS = (1 << 22)

func TestTimerUpdate(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	testName := "Checks whether inner cycle counter increments properly, when timer is disabled"
	t.Run(testName, func(t *testing.T) {
		memory := memory.NewDMGMemory()
		interrupts := NewInterrupts(memory)
		timer := NewTimer(memory, interrupts.Request)

		for i := 0; i < NUM_OF_INSTRUCTIONS; i++ {
			cycles := (rand.Intn(6) + 1) * 4

			timer.Update(cycles)

			expectedCycles := 0
			gotCycles := timer.cycles

			if expectedCycles != gotCycles {
				t.Errorf("Inner cycle counter increments, when timer is disabled. Expected %v, got %v", expectedCycles, gotCycles)
				break
			}
		}
	})

	testName = "Checks whether inner cycle counter increments properly, when timer is enabled"
	t.Run(testName, func(t *testing.T) {
		memory := memory.NewDMGMemory()
		interrupts := NewInterrupts(memory)

		for _, c := range []uint8{0, 1, 2, 3} {
			var cyclesSum int

			timer := NewTimer(memory, interrupts.Request)
			memory.Write(ADDR_TIM_CONTROL, 4|c)
			threshold := timer.getThreshold()

			for i := 0; i < NUM_OF_INSTRUCTIONS; i++ {
				cycles := (rand.Intn(6) + 1) * 4
				cyclesSum += cycles

				timer.Update(cycles)

				expectedCycles := cyclesSum % threshold
				gotCycles := timer.cycles

				if expectedCycles != gotCycles {
					t.Errorf("Inner cycle counter stores incorrect value. Expected %v, got %v", expectedCycles, gotCycles)
					break
				}
			}
		}
	})

	testName = "Checks whether inner cycle counter increments and resets properly, when timer is enabled and TAC is changed at random points"
	t.Run(testName, func(t *testing.T) {
		memory := memory.NewDMGMemory()
		interrupts := NewInterrupts(memory)

		for _, c := range []uint8{0, 1, 2, 3} {
			var cyclesSum int

			timer := NewTimer(memory, interrupts.Request)
			memory.Write(ADDR_TIM_CONTROL, 4|c)
			threshold := timer.getThreshold()

			for i := 0; i < NUM_OF_INSTRUCTIONS; i++ {
				if rand.Int()%(1<<10) == 0 {
					cyclesSum = 0
					memory.Write(ADDR_TIM_CONTROL, uint8(rand.Int())|4)
					threshold = timer.getThreshold()
				}

				cycles := (rand.Intn(6) + 1) * 4
				cyclesSum += cycles

				timer.Update(cycles)

				expectedCycles := cyclesSum % threshold
				gotCycles := timer.cycles

				if expectedCycles != gotCycles {
					t.Errorf("Inner cycle counter stores incorrect value. Expected %v, got %v", expectedCycles, gotCycles)
					break
				}
			}
		}
	})

	testName = "Checks whether counter increments properly, when timer is disabled"
	t.Run(testName, func(t *testing.T) {
		memory := memory.NewDMGMemory()
		interrupts := NewInterrupts(memory)
		timer := NewTimer(memory, interrupts.Request)

		for i := 0; i < NUM_OF_INSTRUCTIONS; i++ {
			cycles := (rand.Intn(6) + 1) * 4

			timer.Update(cycles)

			expectedCounter := uint8(0)
			gotCounter := timer.counter

			if expectedCounter != gotCounter {
				t.Errorf("Counter increments, when timer is disabled. Expected %v, got %v", expectedCounter, gotCounter)
				break
			}
		}
	})

	testName = "Checks whether counter increments properly, when timer is enabled"
	t.Run(testName, func(t *testing.T) {
		memory := memory.NewDMGMemory()
		interrupts := NewInterrupts(memory)

		for _, c := range []uint8{0, 1, 2, 3} {
			var cyclesSum int

			timer := NewTimer(memory, interrupts.Request)
			memory.Write(ADDR_TIM_CONTROL, 4|c)
			threshold := timer.getThreshold()

			for i := 0; i < NUM_OF_INSTRUCTIONS; i++ {
				cycles := (rand.Intn(6) + 1) * 4
				cyclesSum += cycles

				timer.Update(cycles)

				expectedCounter := uint8(cyclesSum / threshold)
				gotCounter := timer.counter

				if expectedCounter != gotCounter {
					t.Errorf("Counter stores incorrect value. Expected %v, got %v", expectedCounter, gotCounter)
					break
				}
			}
		}
	})

	testName = "Checks whether counter increments properly, when timer is enabled and TAC is changed at random points"
	t.Run(testName, func(t *testing.T) {
		memory := memory.NewDMGMemory()
		interrupts := NewInterrupts(memory)

		for _, c := range []uint8{0, 1, 2, 3} {
			var cyclesSum int
			var oldCounter uint8

			timer := NewTimer(memory, interrupts.Request)
			memory.Write(ADDR_TIM_CONTROL, 4|c)
			threshold := timer.getThreshold()

			for i := 0; i < NUM_OF_INSTRUCTIONS; i++ {
				if rand.Int()%(1<<10) == 0 {
					cyclesSum = 0
					oldCounter = timer.counter
					memory.Write(ADDR_TIM_CONTROL, uint8(rand.Int())|4)
					threshold = timer.getThreshold()
				}

				cycles := (rand.Intn(6) + 1) * 4
				cyclesSum += cycles

				timer.Update(cycles)

				expectedCounter := uint8(cyclesSum/threshold) + oldCounter
				gotCounter := timer.counter

				if expectedCounter != gotCounter {
					t.Errorf("Counter stores incorrect value. Expected %v, got %v", expectedCounter, gotCounter)
					break
				}
			}
		}
	})

	testName = "Checks whether counter resets properly"
	t.Run(testName, func(t *testing.T) {
		for _, c := range []uint8{0, 1, 2, 3} {
			for m := uint8(0); m < 0xFF; m++ {
				memory := memory.NewDMGMemory()
				interrupts := NewInterrupts(memory)
				timer := NewTimer(memory, interrupts.Request)

				timer.modulo = m
				memory.Write(ADDR_TIM_CONTROL, 4|c)

				for i := 0; i < NUM_OF_INSTRUCTIONS; i++ {
					cycles := (rand.Intn(6) + 1) * 4

					timer.Update(cycles)

					if (interrupts.requested & INT_TIMER) == INT_TIMER {
						expectedCounter := m + func() uint8 {
							if timer.counter == timer.modulo {
								return 0
							}
							if timerClock[timer.control&3] == TIM_CLOCK_1 {
								if timer.cycles == 0 && cycles == 24 {
									return 1
								}
								if timer.cycles == 0 && cycles == 20 {
									return 1
								}
								if timer.cycles == 4 && cycles == 24 {
									return 1
								}
							}
							return 0
						}()
						gotCounter := timer.counter

						if expectedCounter != gotCounter {
							t.Errorf("Counter stores incorrect value with preset TMA. Expected %v, got %v", expectedCounter, gotCounter)
						}
						break
					}
				}
			}
		}
	})
}

func TestRequestInterrupt(t *testing.T) {
	memory := memory.NewDMGMemory()
	interrupts := NewInterrupts(memory)
	timer := NewTimer(memory, interrupts.Request)

	testName := "Checks whether timer correctly requests and sets an interrupt"
	t.Run(testName, func(t *testing.T) {
		for i := uint8(0); i < 0xFF; i++ {
			interrupts.requested = i & 0b00011111

			timer.requestInterrupt(INT_TIMER)

			expectedRequested := interrupts.requested | INT_TIMER
			gotRequested := memory.Read(ADDR_INT_REQUESTED)

			if expectedRequested != gotRequested {
				t.Errorf("Timer requests interrupt incorrectly. Expected 0x%02X, got 0x%02X", expectedRequested, gotRequested)
				t.Logf("Requested: 0b%08b", interrupts.requested)
				break
			}
		}
	})
}

func TestIsStopped(t *testing.T) {
	memory := memory.NewDMGMemory()
	timer := NewTimer(memory, nil)

	testCases := []struct {
		control       uint8
		stopRequested bool
	}{}

	for c := uint8(0); c < 8; c++ {
		stopRequested := (c>>2)%2 == 0
		testCases = append(testCases, struct {
			control       uint8
			stopRequested bool
		}{c, stopRequested})
	}

	testName := "Checks whether timer correctly determines stop request"
	t.Run(testName, func(t *testing.T) {
		for _, tc := range testCases {
			memory.Write(ADDR_TIM_CONTROL, tc.control)

			expected := tc.stopRequested
			got := timer.isStopped()

			if expected != got {
				t.Errorf("Timer determines stop request incorrectly. Expected %v, got %v", expected, got)
				t.Logf("Control: 0b%05b", tc.control)
				break
			}
		}
	})
}

func TestGetClock(t *testing.T) {
	memory := memory.NewDMGMemory()
	timer := NewTimer(memory, nil)

	testCases := []struct {
		control uint8
	}{}

	for c := uint8(0); c < 0xFF; c++ {
		testCases = append(testCases, struct {
			control uint8
		}{c})
	}

	testName := "Checks whether timer correctly determines stop request"
	t.Run(testName, func(t *testing.T) {
		for _, tc := range testCases {
			memory.Write(ADDR_TIM_CONTROL, tc.control)

			expectedClock := func() int {
				switch tc.control & 3 {
				case 0:
					return TIM_CLOCK_0
				case 1:
					return TIM_CLOCK_1
				case 2:
					return TIM_CLOCK_2
				case 3:
					return TIM_CLOCK_3
				}
				return TIM_CLOCK_0
			}()
			gotClock := timer.getClock()

			if expectedClock != gotClock {
				t.Errorf("Timer determines clock incorreclty. Expected %v, got %v", expectedClock, gotClock)
				t.Logf("Control: 0b%05b", tc.control)
				break
			}
		}
	})
}

func TestGetThreshold(t *testing.T) {
	memory := memory.NewDMGMemory()
	timer := NewTimer(memory, nil)

	testCases := []struct {
		control uint8
	}{}

	for c := uint8(0); c < 0xFF; c++ {
		testCases = append(testCases, struct {
			control uint8
		}{c})
	}

	testName := "Checks whether timer correctly determines stop request"
	t.Run(testName, func(t *testing.T) {
		for _, tc := range testCases {
			memory.Write(ADDR_TIM_CONTROL, tc.control)

			expectedThreshold := func() int {
				switch timer.getClock() {
				case TIM_CLOCK_0:
					return 1024
				case TIM_CLOCK_1:
					return 16
				case TIM_CLOCK_2:
					return 64
				case TIM_CLOCK_3:
					return 256
				}
				return 1024
			}()
			gotThreshold := timer.getThreshold()

			if expectedThreshold != gotThreshold {
				t.Errorf("Timer determines clock incorreclty. Expected %v, got %v", expectedThreshold, gotThreshold)
				t.Logf("Control: 0b%05b", tc.control)
				break
			}
		}
	})
}

func TestTimerHandlers(t *testing.T) {
	memory := memory.NewDMGMemory()
	interrupts := NewInterrupts(memory)
	timer := NewTimer(memory, interrupts.Request)

	testCases := []struct {
		value uint8
	}{}

	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 10; i++ {
		testCases = append(testCases, struct{ value uint8 }{uint8(rand.Int())})
	}

	testName := "Checks whether exported getters work properly"
	t.Run(testName, func(t *testing.T) {
		for _, tc := range testCases {
			timer.counter = tc.value
			timer.modulo = tc.value
			timer.control = tc.value

			expectedCounter := tc.value
			expectedModulo := tc.value
			expectedController := tc.value
			gotCounter := memory.Read(ADDR_TIM_COUNTER)
			gotModulo := memory.Read(ADDR_TIM_MODULO)
			gotController := memory.Read(ADDR_TIM_CONTROL)

			if expectedCounter != gotCounter {
				t.Errorf("Counter getter reads incorrect value. Expected 0x%02X, got 0x%02X", expectedCounter, gotCounter)
				break
			}
			if expectedModulo != gotModulo {
				t.Errorf("Modulo getter reads incorrect value. Expected 0x%02X, got 0x%02X", expectedModulo, gotModulo)
				break
			}
			if expectedController != gotController {
				t.Errorf("Control getter reads incorrect value. Expected 0x%02X, got 0x%02X", expectedController, gotController)
				break
			}
		}
	})

	testName = "Checks whether exported setters work properly"
	t.Run(testName, func(t *testing.T) {
		for _, tc := range testCases {
			memory.Write(ADDR_TIM_COUNTER, tc.value)
			memory.Write(ADDR_TIM_MODULO, tc.value)
			memory.Write(ADDR_TIM_CONTROL, tc.value)

			expectedCounter := tc.value
			expectedModulo := tc.value
			expectedController := tc.value
			gotCounter := timer.counter
			gotModulo := timer.modulo
			gotController := timer.control

			if expectedCounter != gotCounter {
				t.Errorf("Counter setter writes incorrect value. Expected 0x%02X, got 0x%02X", expectedCounter, gotCounter)
				break
			}
			if expectedModulo != gotModulo {
				t.Errorf("Modulo setter writes incorrect value. Expected 0x%02X, got 0x%02X", expectedModulo, gotModulo)
				break
			}
			if expectedController != gotController {
				t.Errorf("Control setter writes incorrect value. Expected 0x%02X, got 0x%02X", expectedController, gotController)
				break
			}
		}
	})

	testName = "Checks whether exported getters/setters work properly"
	t.Run(testName, func(t *testing.T) {
		for _, tc := range testCases {
			memory.Write(ADDR_TIM_COUNTER, tc.value)
			memory.Write(ADDR_TIM_MODULO, tc.value)
			memory.Write(ADDR_TIM_CONTROL, tc.value)

			expectedCounter := tc.value
			expectedModulo := tc.value
			expectedController := tc.value
			gotCounter := memory.Read(ADDR_TIM_COUNTER)
			gotModulo := memory.Read(ADDR_TIM_MODULO)
			gotController := memory.Read(ADDR_TIM_CONTROL)

			if expectedCounter != gotCounter {
				t.Errorf("Counter getter/setter reads/writes incorrect value. Expected 0x%02X, got 0x%02X", expectedCounter, gotCounter)
				break
			}
			if expectedModulo != gotModulo {
				t.Errorf("Modulo getter/setter reads/writes incorrect value. Expected 0x%02X, got 0x%02X", expectedModulo, gotModulo)
				break
			}
			if expectedController != gotController {
				t.Errorf("Control getter/setter reads/writes incorrect value. Expected 0x%02X, got 0x%02X", expectedController, gotController)
				break
			}
		}
	})
}
