package cpu

import (
	"math/rand"
	"testing"

	"gbmu/emulation/memory"
)

func TestRegistersGetters(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	testName := "Checks that registers' getters read registers' values properly"
	t.Run(testName, func(t *testing.T) {
		for i := 0; i < 10; i++ {
			randomizeRegisters(cpu.registers)

			expectedA := cpu.registers.a
			expectedF := cpu.registers.f.value
			expectedB := cpu.registers.b
			expectedC := cpu.registers.c
			expectedD := cpu.registers.d
			expectedE := cpu.registers.e
			expectedH := cpu.registers.h
			expectedL := cpu.registers.l
			expectedBC := uint16(cpu.registers.b)<<8 | uint16(cpu.registers.c)
			expectedDE := uint16(cpu.registers.d)<<8 | uint16(cpu.registers.e)
			expectedHL := uint16(cpu.registers.h)<<8 | uint16(cpu.registers.l)
			expectedSP := cpu.registers.sp
			expectedPC := cpu.registers.pc
			gotA := cpu.registers.getA()
			gotF := cpu.registers.getF()
			gotB := cpu.registers.getB()
			gotC := cpu.registers.getC()
			gotD := cpu.registers.getD()
			gotE := cpu.registers.getE()
			gotH := cpu.registers.getH()
			gotL := cpu.registers.getL()
			gotBC := cpu.registers.getBC()
			gotDE := cpu.registers.getDE()
			gotHL := cpu.registers.getHL()
			gotSP := cpu.registers.getSP()
			gotPC := cpu.registers.getPC()

			if expectedA != gotA {
				t.Errorf("Wrong value of A. Expected 0x%02X, got 0x%02X", expectedA, gotA)
				break
			}
			if expectedF != gotF {
				t.Errorf("Wrong value of F. Expected 0x%02X, got 0x%02X", expectedF, gotF)
				break
			}
			if expectedB != gotB {
				t.Errorf("Wrong value of B. Expected 0x%02X, got 0x%02X", expectedB, gotB)
				break
			}
			if expectedC != gotC {
				t.Errorf("Wrong value of C. Expected 0x%02X, got 0x%02X", expectedC, gotC)
				break
			}
			if expectedD != gotD {
				t.Errorf("Wrong value of D. Expected 0x%02X, got 0x%02X", expectedD, gotD)
				break
			}
			if expectedE != gotE {
				t.Errorf("Wrong value of E. Expected 0x%02X, got 0x%02X", expectedE, gotE)
				break
			}
			if expectedH != gotH {
				t.Errorf("Wrong value of H. Expected 0x%02X, got 0x%02X", expectedH, gotH)
				break
			}
			if expectedL != gotL {
				t.Errorf("Wrong value of L. Expected 0x%02X, got 0x%02X", expectedL, gotL)
				break
			}
			if expectedBC != gotBC {
				t.Errorf("Wrong value of BC. Expected 0x%04X, got 0x%04X", expectedBC, gotBC)
				break
			}
			if expectedDE != gotDE {
				t.Errorf("Wrong value of DE. Expected 0x%04X, got 0x%04X", expectedDE, gotDE)
				break
			}
			if expectedHL != gotHL {
				t.Errorf("Wrong value of HL. Expected 0x%04X, got 0x%04X", expectedHL, gotHL)
				break
			}
			if expectedSP != gotSP {
				t.Errorf("Wrong value of SP. Expected 0x%04X, got 0x%04X", expectedSP, gotSP)
				break
			}
			if expectedPC != gotPC {
				t.Errorf("Wrong value of PC. Expected 0x%04X, got 0x%04X", expectedPC, gotPC)
				break
			}
		}
	})
}

func TestRegistersSetters(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	testName := "Checks that registers' getters set registers' values properly"
	t.Run(testName, func(t *testing.T) {
		for i := 0; i < 10; i++ {
			expectedA := uint8(rand.Int())
			expectedF := uint8(rand.Int()) & 0xF0 // four rightmost bits of F are always cleared
			expectedB := uint8(rand.Int())
			expectedC := uint8(rand.Int())
			expectedD := uint8(rand.Int())
			expectedE := uint8(rand.Int())
			expectedH := uint8(rand.Int())
			expectedL := uint8(rand.Int())

			cpu.registers.setA(expectedA)
			cpu.registers.setF(expectedF)
			cpu.registers.setB(expectedB)
			cpu.registers.setC(expectedC)
			cpu.registers.setD(expectedD)
			cpu.registers.setE(expectedE)
			cpu.registers.setH(expectedH)
			cpu.registers.setL(expectedL)

			gotA := cpu.registers.a
			gotF := cpu.registers.f.value
			gotB := cpu.registers.b
			gotC := cpu.registers.c
			gotD := cpu.registers.d
			gotE := cpu.registers.e
			gotH := cpu.registers.h
			gotL := cpu.registers.l

			if expectedA != gotA {
				t.Errorf("Wrong value of A. Expected 0x%02X, got 0x%02X", expectedA, gotA)
				break
			}
			if expectedF != gotF {
				t.Errorf("Wrong value of F. Expected 0x%02X, got 0x%02X", expectedF, gotF)
				break
			}
			if expectedB != gotB {
				t.Errorf("Wrong value of B. Expected 0x%02X, got 0x%02X", expectedB, gotB)
				break
			}
			if expectedC != gotC {
				t.Errorf("Wrong value of C. Expected 0x%02X, got 0x%02X", expectedC, gotC)
				break
			}
			if expectedD != gotD {
				t.Errorf("Wrong value of D. Expected 0x%02X, got 0x%02X", expectedD, gotD)
				break
			}
			if expectedE != gotE {
				t.Errorf("Wrong value of E. Expected 0x%02X, got 0x%02X", expectedE, gotE)
				break
			}
			if expectedH != gotH {
				t.Errorf("Wrong value of H. Expected 0x%02X, got 0x%02X", expectedH, gotH)
				break
			}
			if expectedL != gotL {
				t.Errorf("Wrong value of L. Expected 0x%02X, got 0x%02X", expectedL, gotL)
				break
			}

			expectedAF := uint16(rand.Int()) & 0xFFF0 // four rightmost bits of F are always cleared
			expectedBC := uint16(rand.Int())
			expectedDE := uint16(rand.Int())
			expectedHL := uint16(rand.Int())
			expectedSP := uint16(rand.Int())
			expectedPC := uint16(rand.Int())

			cpu.registers.setAF(expectedAF)
			cpu.registers.setBC(expectedBC)
			cpu.registers.setDE(expectedDE)
			cpu.registers.setHL(expectedHL)
			cpu.registers.setSP(expectedSP)
			cpu.registers.setPC(expectedPC)

			gotAF := uint16(cpu.registers.a)<<8 | uint16(cpu.registers.f.value)
			gotBC := uint16(cpu.registers.b)<<8 | uint16(cpu.registers.c)
			gotDE := uint16(cpu.registers.d)<<8 | uint16(cpu.registers.e)
			gotHL := uint16(cpu.registers.h)<<8 | uint16(cpu.registers.l)
			gotSP := cpu.registers.sp
			gotPC := cpu.registers.pc

			if expectedAF != gotAF {
				t.Errorf("Wrong value of AF. Expected 0x%04X, got 0x%04X", expectedAF, gotAF)
				break
			}
			if expectedBC != gotBC {
				t.Errorf("Wrong value of BC. Expected 0x%04X, got 0x%04X", expectedBC, gotBC)
				break
			}
			if expectedDE != gotDE {
				t.Errorf("Wrong value of DE. Expected 0x%04X, got 0x%04X", expectedDE, gotDE)
				break
			}
			if expectedHL != gotHL {
				t.Errorf("Wrong value of HL. Expected 0x%04X, got 0x%04X", expectedHL, gotHL)
				break
			}
			if expectedSP != gotSP {
				t.Errorf("Wrong value of SP. Expected 0x%04X, got 0x%04X", expectedSP, gotSP)
				break
			}
			if expectedPC != gotPC {
				t.Errorf("Wrong value of PC. Expected 0x%04X, got 0x%04X", expectedPC, gotPC)
				break
			}
		}
	})
}

func TestRegistersIncrementers(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	testName := "Checks that registers' incrementers increment registers' values properly"
	t.Run(testName, func(t *testing.T) {
		for i := 0; i < 10; i++ {
			randomizeRegisters(cpu.registers)

			expectedA := cpu.registers.a + 1
			expectedB := cpu.registers.b + 1
			expectedC := cpu.registers.c + 1
			expectedD := cpu.registers.d + 1
			expectedE := cpu.registers.e + 1
			expectedH := cpu.registers.h + 1
			expectedL := cpu.registers.l + 1

			cpu.registers.incA()
			cpu.registers.incB()
			cpu.registers.incC()
			cpu.registers.incD()
			cpu.registers.incE()
			cpu.registers.incH()
			cpu.registers.incL()

			gotA := cpu.registers.getA()
			gotB := cpu.registers.getB()
			gotC := cpu.registers.getC()
			gotD := cpu.registers.getD()
			gotE := cpu.registers.getE()
			gotH := cpu.registers.getH()
			gotL := cpu.registers.getL()

			if expectedA != gotA {
				t.Errorf("Wrong value of A. Expected 0x%02X, got 0x%02X", expectedA, gotA)
				break
			}
			if expectedB != gotB {
				t.Errorf("Wrong value of B. Expected 0x%02X, got 0x%02X", expectedB, gotB)
				break
			}
			if expectedC != gotC {
				t.Errorf("Wrong value of C. Expected 0x%02X, got 0x%02X", expectedC, gotC)
				break
			}
			if expectedD != gotD {
				t.Errorf("Wrong value of D. Expected 0x%02X, got 0x%02X", expectedD, gotD)
				break
			}
			if expectedE != gotE {
				t.Errorf("Wrong value of E. Expected 0x%02X, got 0x%02X", expectedE, gotE)
				break
			}
			if expectedH != gotH {
				t.Errorf("Wrong value of H. Expected 0x%02X, got 0x%02X", expectedH, gotH)
				break
			}
			if expectedL != gotL {
				t.Errorf("Wrong value of L. Expected 0x%02X, got 0x%02X", expectedL, gotL)
				break
			}

			randomizeRegisters(cpu.registers)

			expectedBC := uint16(cpu.registers.b)<<8 | uint16(cpu.registers.c) + 1
			expectedDE := uint16(cpu.registers.d)<<8 | uint16(cpu.registers.e) + 1
			expectedHL := uint16(cpu.registers.h)<<8 | uint16(cpu.registers.l) + 1
			expectedSP := cpu.registers.sp + 1
			expectedPC := cpu.registers.pc + 1

			cpu.registers.incBC()
			cpu.registers.incDE()
			cpu.registers.incHL()
			cpu.registers.incSP()
			cpu.registers.incPC()

			gotBC := cpu.registers.getBC()
			gotDE := cpu.registers.getDE()
			gotHL := cpu.registers.getHL()
			gotSP := cpu.registers.getSP()
			gotPC := cpu.registers.getPC()

			if expectedBC != gotBC {
				t.Errorf("Wrong value of BC. Expected 0x%04X, got 0x%04X", expectedBC, gotBC)
				break
			}
			if expectedDE != gotDE {
				t.Errorf("Wrong value of DE. Expected 0x%04X, got 0x%04X", expectedDE, gotDE)
				break
			}
			if expectedHL != gotHL {
				t.Errorf("Wrong value of HL. Expected 0x%04X, got 0x%04X", expectedHL, gotHL)
				break
			}
			if expectedSP != gotSP {
				t.Errorf("Wrong value of SP. Expected 0x%04X, got 0x%04X", expectedSP, gotSP)
				break
			}
			if expectedPC != gotPC {
				t.Errorf("Wrong value of PC. Expected 0x%04X, got 0x%04X", expectedPC, gotPC)
				break
			}
		}
	})
}

func TestRegistersDecrementers(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	testName := "Checks that registers' decrementers decrement registers' values properly"
	t.Run(testName, func(t *testing.T) {
		for i := 0; i < 100; i++ {
			randomizeRegisters(cpu.registers)

			expectedA := cpu.registers.a - 1
			expectedB := cpu.registers.b - 1
			expectedC := cpu.registers.c - 1
			expectedD := cpu.registers.d - 1
			expectedE := cpu.registers.e - 1
			expectedH := cpu.registers.h - 1
			expectedL := cpu.registers.l - 1

			cpu.registers.decA()
			cpu.registers.decB()
			cpu.registers.decC()
			cpu.registers.decD()
			cpu.registers.decE()
			cpu.registers.decH()
			cpu.registers.decL()

			gotA := cpu.registers.getA()
			gotB := cpu.registers.getB()
			gotC := cpu.registers.getC()
			gotD := cpu.registers.getD()
			gotE := cpu.registers.getE()
			gotH := cpu.registers.getH()
			gotL := cpu.registers.getL()

			if expectedA != gotA {
				t.Errorf("Wrong value of A. Expected 0x%02X, got 0x%02X", expectedA, gotA)
				break
			}
			if expectedB != gotB {
				t.Errorf("Wrong value of B. Expected 0x%02X, got 0x%02X", expectedB, gotB)
				break
			}
			if expectedC != gotC {
				t.Errorf("Wrong value of C. Expected 0x%02X, got 0x%02X", expectedC, gotC)
				break
			}
			if expectedD != gotD {
				t.Errorf("Wrong value of D. Expected 0x%02X, got 0x%02X", expectedD, gotD)
				break
			}
			if expectedE != gotE {
				t.Errorf("Wrong value of E. Expected 0x%02X, got 0x%02X", expectedE, gotE)
				break
			}
			if expectedH != gotH {
				t.Errorf("Wrong value of H. Expected 0x%02X, got 0x%02X", expectedH, gotH)
				break
			}
			if expectedL != gotL {
				t.Errorf("Wrong value of L. Expected 0x%02X, got 0x%02X", expectedL, gotL)
				break
			}

			randomizeRegisters(cpu.registers)

			expectedBC := uint16(cpu.registers.b)<<8 | uint16(cpu.registers.c) - 1
			expectedDE := uint16(cpu.registers.d)<<8 | uint16(cpu.registers.e) - 1
			expectedHL := uint16(cpu.registers.h)<<8 | uint16(cpu.registers.l) - 1
			expectedSP := cpu.registers.sp - 1

			cpu.registers.decBC()
			cpu.registers.decDE()
			cpu.registers.decHL()
			cpu.registers.decSP()

			gotBC := cpu.registers.getBC()
			gotDE := cpu.registers.getDE()
			gotHL := cpu.registers.getHL()
			gotSP := cpu.registers.getSP()

			if expectedBC != gotBC {
				t.Errorf("Wrong value of BC. Expected 0x%04X, got 0x%04X", expectedBC, gotBC)
				break
			}
			if expectedDE != gotDE {
				t.Errorf("Wrong value of DE. Expected 0x%04X, got 0x%04X", expectedDE, gotDE)
				break
			}
			if expectedHL != gotHL {
				t.Errorf("Wrong value of HL. Expected 0x%04X, got 0x%04X", expectedHL, gotHL)
				break
			}
			if expectedSP != gotSP {
				t.Errorf("Wrong value of SP. Expected 0x%04X, got 0x%04X", expectedSP, gotSP)
				break
			}
		}
	})
}
