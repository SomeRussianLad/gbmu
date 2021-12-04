package controllers

import (
	"math/rand"
	"testing"
	"time"

	"gbmu/emulation/memory"
)

func TestDividerUpdate(t *testing.T) {
	memory := memory.NewDMGMemory()
	divider := NewDivider(memory)

	testCases := []struct {
		cycles int
	}{}

	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < (1 << 18); i++ {
		testCases = append(testCases, struct{ cycles int }{(rand.Intn(5) + 1) * 4})
	}

	var cyclesSum int
	testName := "Checks whether update works properly"
	t.Run(testName, func(t *testing.T) {
		for _, tc := range testCases {
			divider.Update(tc.cycles)
			cyclesSum += tc.cycles

			expectedCounter := uint8(cyclesSum / 256)
			gotCounter := divider.counter

			if expectedCounter != gotCounter {
				t.Errorf("Divider increments incorrectly. Expected 0x%04X, got 0x%04X", expectedCounter, gotCounter)
				t.Logf("Machine cycles (total): %v", cyclesSum)
				break
			}
		}
	})
}

func TestDividerHandlers(t *testing.T) {
	memory := memory.NewDMGMemory()
	divider := NewDivider(memory)

	testCases := []struct {
		value uint8
	}{}

	for i := 0; i < 10; i++ {
		testCases = append(testCases, struct{ value uint8 }{uint8(rand.Int())})
	}

	testName := "Checks whether exported getters work properly"
	t.Run(testName, func(t *testing.T) {
		for _, tc := range testCases {
			divider.counter = tc.value

			expectedCounter := tc.value
			gotCounter := memory.Read(ADDR_DIV_COUNTER)

			if expectedCounter != gotCounter {
				t.Errorf("Counter getter reads incorrect value. Expected 0x%02X, got 0x%02X", expectedCounter, gotCounter)
				break
			}
		}
	})

	testName = "Checks whether exported setters work properly"
	t.Run(testName, func(t *testing.T) {
		for _, tc := range testCases {
			memory.Write(ADDR_DIV_COUNTER, tc.value)

			expectedCounter := uint8(0) // Writing to DIV resets its counter
			gotCounter := divider.counter

			if expectedCounter != gotCounter {
				t.Errorf("Counter setter writes incorrect value. Expected 0x%02X, got 0x%02X", expectedCounter, gotCounter)
				break
			}
		}
	})

	testName = "Checks whether exported getters/setters work properly"
	t.Run(testName, func(t *testing.T) {
		for _, tc := range testCases {
			memory.Write(ADDR_DIV_COUNTER, tc.value)

			expectedCounter := uint8(0) // Writing to DIV resets its counter
			gotCounter := memory.Read(ADDR_DIV_COUNTER)

			if expectedCounter != gotCounter {
				t.Errorf("Counter getter/setter reads/writes incorrect value. Expected 0x%02X, got 0x%02X", expectedCounter, gotCounter)
				break
			}
		}
	})
}
