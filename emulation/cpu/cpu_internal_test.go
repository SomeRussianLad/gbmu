package cpu

import (
	"math/rand"
	"testing"
	"time"

	"gbmu/emulation/controllers"
	"gbmu/emulation/memory"
)

func randomizeRegisters(registers *registers) {
	rand.Seed(time.Now().UTC().UnixNano())
	registers.setA(uint8(rand.Int()))
	registers.setF(uint8(rand.Int()))
	registers.setB(uint8(rand.Int()))
	registers.setC(uint8(rand.Int()))
	registers.setD(uint8(rand.Int()))
	registers.setE(uint8(rand.Int()))
	registers.setH(uint8(rand.Int()))
	registers.setL(uint8(rand.Int()))
	registers.setSP(uint16(rand.Int()))
	registers.setPC(uint16(rand.Int()))
}

func randomizeMemory(memory memory.Memory) {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0x0000; i < 0x10000; i++ {
		memory.Write(uint16(i), uint8(rand.Int()))
	}
}

func read8BitOperand(c *CPU) uint8 {
	addr := c.registers.getPC()
	value := c.memory.Read(addr)
	return value
}

func read16BitOperand(c *CPU) uint16 {
	addr := c.registers.getPC()
	lsb := uint16(c.memory.Read(addr))
	msb := uint16(c.memory.Read(addr + 1))
	value := msb<<8 | lsb
	return value
}

// TODO(somerussianlad): Write tests for cpu.update()

func TestEnableHalt(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	cpu.instructions[0x76].exec(cpu)

	expectedIsHalted := true
	gotIsHalted := cpu.isHalted

	if expectedIsHalted != gotIsHalted {
		t.Errorf("Wrong value in isHalted. Expected %v, got %v", expectedIsHalted, gotIsHalted)
	}
}

func TestDisableHalt(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	cpu.disableHalt()

	expectedIsHalted := false
	gotIsHalted := cpu.isHalted

	if expectedIsHalted != gotIsHalted {
		t.Errorf("Wrong value in isHalted. Expected %v, got %v", expectedIsHalted, gotIsHalted)
	}
}

func TestHandleInterrupt(t *testing.T) {
	memory := memory.NewDMGMemory()
	interrupts := controllers.NewInterrupts(memory)
	cpu := NewCPU(memory, interrupts, nil, nil, nil)

	testCases := []struct {
		interrupt uint8
		sp        uint16
		pc        uint16
	}{}

	rand.Seed(time.Now().UTC().UnixNano())
	for _, i := range []uint8{
		controllers.INT_VBLANK,
		controllers.INT_LCD,
		controllers.INT_TIMER,
		controllers.INT_SERIAL,
		controllers.INT_JOYPAD,
	} {
		testCases = append(testCases, struct {
			interrupt uint8
			sp        uint16
			pc        uint16
		}{i, uint16(rand.Int()), uint16(rand.Int())})
	}

	testName := "Checks whether this method handles interrupt correctly"
	t.Run(testName, func(t *testing.T) {
		for _, tc := range testCases {
			memory.Write(controllers.ADDR_INT_ENABLED, uint8(rand.Int())|tc.interrupt)
			memory.Write(controllers.ADDR_INT_REQUESTED, uint8(rand.Int())|tc.interrupt)
			interrupts.EnableMaster()
			cpu.registers.setSP(tc.sp)
			cpu.registers.setPC(tc.pc)

			expectedMaster := false
			expectedRequested := memory.Read(controllers.ADDR_INT_REQUESTED) & (0xFF - tc.interrupt)
			expectedVector := interrupts.GetVector(tc.interrupt)
			expectedStackedPC := tc.pc

			cpu.handleInterrupt(tc.interrupt)

			gotMaster := interrupts.IsMasterEnabled()
			gotRequested := memory.Read(controllers.ADDR_INT_REQUESTED)
			gotVector := cpu.registers.getPC()
			gotStackedPC := func() uint16 {
				addr := cpu.registers.getSP()
				msb := uint16(memory.Read(addr + 1))
				lsb := uint16(memory.Read(addr))
				value := msb<<8 | lsb
				return value
			}()

			if expectedMaster != gotMaster {
				t.Errorf("Wrong value in master. Expected %v, got %v", expectedMaster, gotMaster)
			}
			if expectedRequested != gotRequested {
				t.Errorf("Wrong value in requested. Expected 0b%05b, got 0b%05b", expectedRequested, gotRequested)
			}
			if expectedVector != gotVector {
				t.Errorf("Wrong vector stored in PC. Expected 0x%04X, got 0x%04X", expectedVector, gotVector)
			}
			if expectedStackedPC != gotStackedPC {
				t.Errorf("Wrong PC stored in stack. Expected 0x%04X, got 0x%04X", expectedStackedPC, gotStackedPC)
				break
			}
		}
	})
}

func TestReadPC(t *testing.T) {
	memory := memory.NewDMGMemory()
	interrupts := controllers.NewInterrupts(memory)
	cpu := NewCPU(memory, interrupts, nil, nil, nil)

	randomizeMemory(memory)

	testName := "Checks whether this method reads memory and auto-increments PC correctly"
	t.Run(testName, func(t *testing.T) {
		for i := 0; i < 10; i++ {
			randomizeRegisters(cpu.registers)

			expectedValue := memory.Read(cpu.registers.getPC())
			expectedPC := cpu.registers.getPC() + 1

			gotValue := cpu.readPC()
			gotPC := cpu.registers.getPC()

			if expectedValue != gotValue {
				t.Errorf("readPC reads incorrect value. Expected 0x%02X, got 0x%02X", expectedValue, gotValue)
			}
			if expectedPC != gotPC {
				t.Errorf("readPC increments PC incorrectly. Expected 0x%04X, got 0x%04X", expectedPC, gotPC)
				break
			}
		}
	})
}

func TestExecuteNextInstructionRegular(t *testing.T) {
	memory := memory.NewDMGMemory()
	interrupts := controllers.NewInterrupts(memory)
	cpu := NewCPU(memory, interrupts, nil, nil, nil)

	irregularInstructions := map[int]struct{}{
		0x20: {}, 0x28: {}, 0x30: {}, 0x38: {},
		0xC0: {}, 0xC2: {}, 0xC4: {}, 0xC8: {}, 0xCA: {}, 0xCC: {},
		0xD0: {}, 0xD2: {}, 0xD4: {}, 0xD8: {}, 0xDA: {}, 0xDC: {},
	}
	testCases := []struct {
		opcode, cycles int
	}{}

	for o, i := range cpu.instructions {
		if _, isIrregular := irregularInstructions[o]; o != 0xCB && o <= 0xFF && !isIrregular {
			testCases = append(testCases, struct {
				opcode int
				cycles int
			}{o, i.cycles})
		}
	}

	testName := "Making sure that instructions with constant amount of cycles return it correctly"
	t.Run(testName, func(t *testing.T) {
		for _, tc := range testCases {
			addr := uint16(rand.Int()) % 0xFF00

			cpu.registers.setPC(addr)
			memory.Write(addr, uint8(tc.opcode))

			expectedCycles := tc.cycles
			gotCycles := cpu.executeNextInstruction()

			if expectedCycles != gotCycles {
				t.Errorf("Wrong amount of returned cycles from regular non-prefixed instruction 0x%02X. Expected %v, got %v", tc.opcode, expectedCycles, gotCycles)
				break
			}
		}
	})
}

func TestExecuteNextInstructionIrregular(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	irregularInstructions := map[int]struct {
		acceptedCycles, rejectedCycles     int
		acceptedCallback, rejectedCallback func(*CPU)
	}{
		0x20: {12, 8, func(c *CPU) { cpu.registers.f.setZ(false) }, func(c *CPU) { cpu.registers.f.setZ(true) }},
		0x28: {12, 8, func(c *CPU) { cpu.registers.f.setZ(true) }, func(c *CPU) { cpu.registers.f.setZ(false) }},
		0x30: {12, 8, func(c *CPU) { cpu.registers.f.setC(false) }, func(c *CPU) { cpu.registers.f.setC(true) }},
		0x38: {12, 8, func(c *CPU) { cpu.registers.f.setC(true) }, func(c *CPU) { cpu.registers.f.setC(false) }},

		0xC0: {20, 8, func(c *CPU) { cpu.registers.f.setZ(false) }, func(c *CPU) { cpu.registers.f.setZ(true) }},
		0xC2: {16, 12, func(c *CPU) { cpu.registers.f.setZ(false) }, func(c *CPU) { cpu.registers.f.setZ(true) }},
		0xC4: {24, 12, func(c *CPU) { cpu.registers.f.setZ(false) }, func(c *CPU) { cpu.registers.f.setZ(true) }},
		0xC8: {20, 8, func(c *CPU) { cpu.registers.f.setZ(true) }, func(c *CPU) { cpu.registers.f.setZ(false) }},
		0xCA: {16, 12, func(c *CPU) { cpu.registers.f.setZ(true) }, func(c *CPU) { cpu.registers.f.setZ(false) }},
		0xCC: {24, 12, func(c *CPU) { cpu.registers.f.setZ(true) }, func(c *CPU) { cpu.registers.f.setZ(false) }},

		0xD0: {20, 8, func(c *CPU) { cpu.registers.f.setC(false) }, func(c *CPU) { cpu.registers.f.setC(true) }},
		0xD2: {16, 12, func(c *CPU) { cpu.registers.f.setC(false) }, func(c *CPU) { cpu.registers.f.setC(true) }},
		0xD4: {24, 12, func(c *CPU) { cpu.registers.f.setC(false) }, func(c *CPU) { cpu.registers.f.setC(true) }},
		0xD8: {20, 8, func(c *CPU) { cpu.registers.f.setC(true) }, func(c *CPU) { cpu.registers.f.setC(false) }},
		0xDA: {16, 12, func(c *CPU) { cpu.registers.f.setC(true) }, func(c *CPU) { cpu.registers.f.setC(false) }},
		0xDC: {24, 12, func(c *CPU) { cpu.registers.f.setC(true) }, func(c *CPU) { cpu.registers.f.setC(false) }},
	}
	testCases := []struct {
		opcode int
	}{}

	for o := range irregularInstructions {
		testCases = append(testCases, struct {
			opcode int
		}{o})
	}

	testName := "Making sure that instructions with variable amount of cycles (and accepted condition) return it correctly"
	t.Run(testName, func(t *testing.T) {
		for _, tc := range testCases {
			addr := uint16(rand.Int()) % 0xFF00

			cpu.registers.setPC(addr)
			memory.Write(addr, uint8(tc.opcode))

			irregularInstructions[tc.opcode].acceptedCallback(cpu)

			expectedCycles := irregularInstructions[tc.opcode].acceptedCycles
			gotCycles := cpu.executeNextInstruction()

			if expectedCycles != gotCycles {
				t.Errorf("Wrong amount of returned cycles from irregular instruction 0x%02X on accepted condition. Expected %v, got %v", tc.opcode, expectedCycles, gotCycles)
			}
		}
	})

	testName = "Making sure that instructions with variable amount of cycles (and rejected condition) return it correctly"
	t.Run(testName, func(t *testing.T) {
		for _, tc := range testCases {
			addr := uint16(rand.Int()) % 0xFF00

			cpu.registers.setPC(addr)
			memory.Write(addr, uint8(tc.opcode))

			irregularInstructions[tc.opcode].rejectedCallback(cpu)

			expectedCycles := irregularInstructions[tc.opcode].rejectedCycles
			gotCycles := cpu.executeNextInstruction()

			if expectedCycles != gotCycles {
				t.Errorf("Wrong amount of returned cycles from irregular instruction 0x%02X on accepted condition. Expected %v, got %v", tc.opcode, expectedCycles, gotCycles)
			}
		}
	})
}

func TestExecuteNextInstructionPrefixed(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	testCases := []struct {
		opcode, cycles int
	}{}

	for o, i := range cpu.instructions {
		if o > 0xFF {
			testCases = append(testCases, struct {
				opcode int
				cycles int
			}{o, i.cycles})
		}
	}

	testName := "Making sure that prefixed instructions return the amount of cycles correctly"
	t.Run(testName, func(t *testing.T) {
		for _, tc := range testCases {
			addr := uint16(rand.Int()) % 0xFF00

			cpu.registers.setPC(addr)
			memory.Write(addr, 0xCB)
			memory.Write(addr+1, uint8(tc.opcode&0xFF))

			expectedCycles := tc.cycles
			gotCycles := cpu.executeNextInstruction()

			if expectedCycles != gotCycles {
				t.Errorf("Wrong amount of returned cycles from prefixed instruction 0x%02X. Expected %v, got %v", tc.opcode, expectedCycles, gotCycles)
			}
		}
	})
}
