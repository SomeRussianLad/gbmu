package controllers

import (
	"math/rand"
	"testing"
	"time"

	"gbmu/emulation/memory"
)

func TestIsMasterEnabled(t *testing.T) {
	memory := memory.NewDMGMemory()
	interrupts := NewInterrupts(memory)

	testCases := []struct {
		master bool
	}{
		{true},
		{false},
	}

	for _, tc := range testCases {
		interrupts.master = tc.master

		expectedMaster := tc.master
		gotMaster := interrupts.IsMasterEnabled()

		if expectedMaster != gotMaster {
			t.Errorf("Wrong master value. Expected %v, got %v", expectedMaster, gotMaster)
		}
	}
}

func TestIsDelayed(t *testing.T) {
	memory := memory.NewDMGMemory()
	interrupts := NewInterrupts(memory)

	testCases := []struct {
		delay bool
	}{
		{true},
		{false},
	}

	for _, tc := range testCases {
		interrupts.delay = tc.delay

		expectedDelay := tc.delay
		gotDelay := interrupts.IsDelayed()

		if expectedDelay != gotDelay {
			t.Errorf("Wrong delay value. Expected %v, got %v", expectedDelay, gotDelay)
		}
	}
}

func TestMasterSetters(t *testing.T) {
	memory := memory.NewDMGMemory()
	interrupts := NewInterrupts(memory)

	testCases := []struct {
		f      func()
		master bool
	}{
		{interrupts.EnableMaster, true},
		{interrupts.DisableMaster, false},
	}

	for _, tc := range testCases {
		tc.f()

		expectedMaster := tc.master
		gotMaster := interrupts.master

		if expectedMaster != gotMaster {
			t.Errorf("Wrong master value. Expected %v, got %v", expectedMaster, gotMaster)
		}
	}
}

func TestRequest(t *testing.T) {
	memory := memory.NewDMGMemory()
	interrupts := NewInterrupts(memory)

	testCases := []struct {
		interrupt uint8
	}{}

	for i := uint8(0); i < 0xFF; i++ {
		testCases = append(testCases, struct{ interrupt uint8 }{i})
	}

	rand.Seed(time.Now().UTC().UnixNano())
	testName := "Checks whether it sets interrupt request bits correctly"
	t.Run(testName, func(t *testing.T) {
		for _, tc := range testCases {
			interrupts.requested = uint8(rand.Int())

			interrupts.Request(tc.interrupt)

			expectedRequested := interrupts.requested | tc.interrupt
			gotRequested := interrupts.requested

			if expectedRequested != gotRequested {
				t.Errorf("Wrong requested value. Expected 0b%05b, got 0b%05b", expectedRequested, gotRequested)
				t.Logf("Passed interrupt value: 0b%05b", tc.interrupt)
				break
			}
		}
	})
}

func TestPending(t *testing.T) {
	memory := memory.NewDMGMemory()
	interrupts := NewInterrupts(memory)

	testCases := []struct {
		delay                       bool
		enabled, requested, pending uint8
		ok                          bool
	}{}

	for _, delay := range []bool{false, true} {
		for enabled := uint8(0); enabled < 32; enabled++ {
			for requested := uint8(0); requested < 32; requested++ {
				pending, ok := func() (uint8, bool) {
					if !delay {
						pending := enabled & requested
						if pending%2 == 1 {
							return INT_VBLANK, true
						}
						if (pending>>1)%2 == 1 {
							return INT_LCD, true
						}
						if (pending>>2)%2 == 1 {
							return INT_TIMER, true
						}
						if (pending>>3)%2 == 1 {
							return INT_SERIAL, true
						}
						if (pending>>4)%2 == 1 {
							return INT_JOYPAD, true
						}
					}
					return 0, false
				}()
				testCases = append(testCases, struct {
					delay                       bool
					enabled, requested, pending uint8
					ok                          bool
				}{delay, enabled, requested, pending, ok})
			}
		}
	}

	testName := "Fetches top-priority interrupt if enabled, 0 otherwise"
	t.Run(testName, func(t *testing.T) {
		for _, tc := range testCases {
			interrupts.delay = tc.delay
			memory.Write(ADDR_INT_ENABLED, tc.enabled)
			memory.Write(ADDR_INT_REQUESTED, tc.requested)

			expectedInterrupt, expectedOk := tc.pending, tc.ok
			gotInterrupt, gotOk := interrupts.Pending()

			if expectedInterrupt != gotInterrupt {
				t.Errorf("Wrong interrupt value. Expected 0b%05b, got 0b%05b", expectedInterrupt, gotInterrupt)
				break
			}
			if expectedOk != gotOk {
				t.Errorf("Wrong ok value. Expected %v, got %v", expectedOk, gotOk)
				break
			}
		}
	})
}

func TestGetVector(t *testing.T) {
	memory := memory.NewDMGMemory()
	interrupts := NewInterrupts(memory)

	testCases := []struct {
		interrupt uint8
		vector    uint16
	}{
		{INT_VBLANK, VEC_VBLANK},
		{INT_LCD, VEC_LCD},
		{INT_TIMER, VEC_TIMER},
		{INT_SERIAL, VEC_SERIAL},
		{INT_JOYPAD, VEC_JOYPAD},
	}

	testName := "Checks whether it fetches correct interrupt vector"
	t.Run(testName, func(t *testing.T) {
		for _, tc := range testCases {
			expectedVector := tc.vector
			gotVector := interrupts.GetVector(tc.interrupt)

			if expectedVector != gotVector {
				t.Errorf("Wrong vector. Expected 0x%04X, got 0x%04X", expectedVector, gotVector)
				t.Logf("INT: %05b", tc.interrupt)
			}
		}
	})
}

func TestAcknowledge(t *testing.T) {
	memory := memory.NewDMGMemory()
	interrupts := NewInterrupts(memory)

	testCases := []struct {
		interrupt, requested uint8
	}{}

	for _, interrupt := range []uint8{
		INT_VBLANK,
		INT_LCD,
		INT_TIMER,
		INT_SERIAL,
		INT_JOYPAD,
	} {
		for requested := uint8(0); requested < 32; requested++ {
			testCases = append(testCases, struct {
				interrupt, requested uint8
			}{interrupt, requested})
		}
	}

	testName := "Checks whether it clears handled interrupt bit correctly"
	t.Run(testName, func(t *testing.T) {
		for _, tc := range testCases {
			memory.Write(ADDR_INT_REQUESTED, tc.requested)

			interrupts.Acknowledge(tc.interrupt)

			expectedRequested := tc.requested & (0xFF - func() uint8 {
				if tc.interrupt%2 > 0 {
					return INT_VBLANK
				}
				if (tc.interrupt>>1)%2 > 0 {
					return INT_LCD
				}
				if (tc.interrupt>>2)%2 > 0 {
					return INT_TIMER
				}
				if (tc.interrupt>>3)%2 > 0 {
					return INT_SERIAL
				}
				if (tc.interrupt>>4)%2 > 0 {
					return INT_JOYPAD
				}
				return 0
			}())
			gotRequested := memory.Read(ADDR_INT_REQUESTED)

			if expectedRequested != gotRequested {
				t.Errorf("Wrong requested value after acknowledge. Expected 0b%05b, got 0b%05b", expectedRequested, gotRequested)
				break
			}
		}
	})
}

func TestInterruptsHandlers(t *testing.T) {
	memory := memory.NewDMGMemory()
	interrupts := NewInterrupts(memory)

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
			interrupts.enabled = tc.value
			interrupts.requested = tc.value

			expectedEnabled := tc.value
			expectedRequested := tc.value
			gotEnabled := memory.Read(ADDR_INT_ENABLED)
			gotRequested := memory.Read(ADDR_INT_REQUESTED)

			if expectedEnabled != gotEnabled {
				t.Errorf("Enabled getter reads incorrect value. Expected 0x%02X, got 0x%02X", expectedEnabled, gotEnabled)
				break
			}
			if expectedRequested != gotRequested {
				t.Errorf("Requested getter reads incorrect value. Expected 0x%02X, got 0x%02X", expectedRequested, gotRequested)
				break
			}
		}
	})

	testName = "Checks whether exported setters work properly"
	t.Run(testName, func(t *testing.T) {
		for _, tc := range testCases {
			memory.Write(ADDR_INT_ENABLED, tc.value)
			memory.Write(ADDR_INT_REQUESTED, tc.value)

			expectedEnabled := tc.value
			expectedRequested := tc.value
			gotEnabled := interrupts.enabled
			gotRequested := interrupts.requested

			if expectedEnabled != gotEnabled {
				t.Errorf("Enabled setter writes incorrect value. Expected 0x%02X, got 0x%02X", expectedEnabled, gotEnabled)
				break
			}
			if expectedRequested != gotRequested {
				t.Errorf("Requested setter writes incorrect value. Expected 0x%02X, got 0x%02X", expectedRequested, gotRequested)
				break
			}
		}
	})

	testName = "Checks whether exported getters/setters work properly"
	t.Run(testName, func(t *testing.T) {
		for _, tc := range testCases {
			memory.Write(ADDR_INT_ENABLED, tc.value)
			memory.Write(ADDR_INT_REQUESTED, tc.value)

			expectedEnabled := tc.value
			expectedRequested := tc.value
			gotEnabled := memory.Read(ADDR_INT_ENABLED)
			gotRequested := memory.Read(ADDR_INT_REQUESTED)

			if expectedEnabled != gotEnabled {
				t.Errorf("Enabled getter/sette reads/writes incorrect value. Expected 0x%02X, got 0x%02X", expectedEnabled, gotEnabled)
				break
			}
			if expectedRequested != gotRequested {
				t.Errorf("Requested getter/sette reads/writes incorrect value. Expected 0x%02X, got 0x%02X", expectedRequested, gotRequested)
				break
			}
		}
	})
}
