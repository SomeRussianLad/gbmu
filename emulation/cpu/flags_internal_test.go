package cpu

import (
	"math/rand"
	"testing"
	"time"

	"gbmu/emulation/memory"
)

func TestGetValue(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	testCases := []struct {
		value uint8
	}{}

	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 10; i++ {
		testCases = append(testCases, struct{ value uint8 }{uint8(rand.Int())})
	}

	testName := "Checks that getValue reads value properly"
	t.Run(testName, func(t *testing.T) {
		for _, tc := range testCases {
			cpu.registers.f.value = tc.value

			expectedValue := tc.value & 0xF0 // four rightmost bits of F are always cleared
			gotValue := cpu.registers.f.getValue()

			if expectedValue != gotValue {
				t.Errorf("Wrong value. Expected 0x%02X, got 0x%02X", expectedValue, gotValue)
				break
			}
		}
	})
}

func TestSetValue(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	testCases := []struct {
		value uint8
	}{}

	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 10; i++ {
		testCases = append(testCases, struct{ value uint8 }{uint8(rand.Int())})
	}

	testName := "Checks that setValue sets value properly"
	t.Run(testName, func(t *testing.T) {
		for _, tc := range testCases {
			cpu.registers.f.setValue(tc.value)

			expectedValue := tc.value & 0xF0 // four rightmost bits of F are always cleared
			gotValue := cpu.registers.f.value

			if expectedValue != gotValue {
				t.Errorf("Wrong value. Expected 0x%02X, got 0x%02X", expectedValue, gotValue)
				break
			}
		}
	})
}

func TestGetCarry(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	testCases := []struct {
		value uint8
	}{}

	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 10; i++ {
		testCases = append(testCases, struct{ value uint8 }{uint8(rand.Int())})
	}

	testName := "Checks that getCarry reads Carry flag properly"
	t.Run(testName, func(t *testing.T) {
		for _, tc := range testCases {
			cpu.registers.f.setValue(tc.value)

			expectedValue := (tc.value >> 4) & 1
			gotValue := cpu.registers.f.getCarry()

			if expectedValue != gotValue {
				t.Errorf("Wrong value. Expected 0x%02X, got 0x%02X", expectedValue, gotValue)
				break
			}
		}
	})
}

func TestFlagsGetters(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	testCases := []struct {
		value uint8
	}{}

	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 10; i++ {
		testCases = append(testCases, struct{ value uint8 }{uint8(rand.Int())})
	}

	testName := "Checks that flag getters read flag values properly"
	t.Run(testName, func(t *testing.T) {
		for _, tc := range testCases {
			cpu.registers.f.setValue(tc.value)

			expectedZ := (tc.value>>7)&1 == 1
			expectedN := (tc.value>>6)&1 == 1
			expectedH := (tc.value>>5)&1 == 1
			expectedC := (tc.value>>4)&1 == 1
			gotZ := cpu.registers.f.getZ()
			gotN := cpu.registers.f.getN()
			gotH := cpu.registers.f.getH()
			gotC := cpu.registers.f.getC()

			if expectedZ != gotZ {
				t.Errorf("Wrong value of Z. Expected %v, got %v", expectedZ, gotZ)
				break
			}
			if expectedN != gotN {
				t.Errorf("Wrong value of N. Expected %v, got %v", expectedN, gotN)
				break
			}
			if expectedH != gotH {
				t.Errorf("Wrong value of H. Expected %v, got %v", expectedH, gotH)
				break
			}
			if expectedC != gotC {
				t.Errorf("Wrong value of C. Expected %v, got %v", expectedC, gotC)
				break
			}
		}
	})
}

func TestFlagsSetters(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	testCases := []struct {
		z, n, h, c bool
	}{}

	for i := 0; i < 10; i++ {
		testCases = append(testCases, struct{ z, n, h, c bool }{
			rand.Int()%2 > 0,
			rand.Int()%2 > 0,
			rand.Int()%2 > 0,
			rand.Int()%2 > 0,
		})
	}

	testName := "Checks that flag setters set flag values properly"
	t.Run(testName, func(t *testing.T) {
		for _, tc := range testCases {
			cpu.registers.f.setZ(tc.z)
			cpu.registers.f.setN(tc.n)
			cpu.registers.f.setH(tc.h)
			cpu.registers.f.setC(tc.c)

			expectedZ := tc.z
			expectedN := tc.n
			expectedH := tc.h
			expectedC := tc.c

			gotZ := (cpu.registers.f.value>>7)&1 == 1
			gotN := (cpu.registers.f.value>>6)&1 == 1
			gotH := (cpu.registers.f.value>>5)&1 == 1
			gotC := (cpu.registers.f.value>>4)&1 == 1

			if expectedZ != gotZ {
				t.Errorf("Wrong value of Z. Expected %v, got %v", expectedZ, gotZ)
				break
			}
			if expectedN != gotN {
				t.Errorf("Wrong value of N. Expected %v, got %v", expectedN, gotN)
				break
			}
			if expectedH != gotH {
				t.Errorf("Wrong value of H. Expected %v, got %v", expectedH, gotH)
				break
			}
			if expectedC != gotC {
				t.Errorf("Wrong value of C. Expected %v, got %v", expectedC, gotC)
				break
			}
		}
	})
}
