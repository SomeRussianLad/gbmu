package cpu

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"testing"

	"gbmu/emulation/controllers"
	"gbmu/emulation/memory"
)

// Checks for proper amount of PC increments after execution
func TestPC(t *testing.T) {
	memory := memory.NewDMGMemory()
	interrupts := controllers.NewInterrupts(memory)
	cpu := NewCPU(memory, interrupts, nil, nil, nil)

	reSkip, _ := regexp.Compile(`JP|JR|CALL|RET|RST`)

	for o, i := range cpu.instructions {
		// Any instructions that JUMP must be skipped as they overwrite PC
		// completely and are not relevant in this test
		if reSkip.Match([]byte(i.mnemonic)) {
			continue
		}

		testName := fmt.Sprintf("Executes %s", i.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)

			pc0 := cpu.registers.getPC()

			// Fake instruction read, that auto-increments PC
			cpu.registers.incPC()

			// Because PREFIX instructions are 2 bytes long, it is
			// required to increment PC one more time
			if o > 0xFF {
				cpu.registers.incPC()
			}

			i.exec(cpu)
			pc1 := cpu.registers.getPC()

			if int(pc1-pc0) != i.length {
				t.Errorf("PC incremented incorrectly. Expected %v steps, got %v", int(pc1-pc0), i.length)
			}
		})
	}
}

func TestFlagsConformity(t *testing.T) {
	memory := memory.NewDMGMemory()
	interrupts := controllers.NewInterrupts(memory)
	cpu := NewCPU(memory, interrupts, nil, nil, nil)

	randomizeMemory(memory)

	for _, i := range cpu.instructions {
		testName := fmt.Sprintf("Executes %s", i.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			i.exec(cpu)
			value := cpu.registers.getF()

			if cpu.registers.f.getZ() != ((value & 128) == 128) {
				t.Errorf("Conformity broken. Flag Z: %t, RegF[7]: %t", cpu.registers.f.getZ(), (value&128) == 128)
			}
			if cpu.registers.f.getN() != ((value & 64) == 64) {
				t.Errorf("Conformity broken. Flag N: %t, RegF[6]: %t", cpu.registers.f.getN(), (value&64) == 64)
			}
			if cpu.registers.f.getH() != ((value & 32) == 32) {
				t.Errorf("Conformity broken. Flag H: %t, RegF[5]: %t", cpu.registers.f.getH(), (value&32) == 32)
			}
			if cpu.registers.f.getC() != ((value & 16) == 16) {
				t.Errorf("Conformity broken. Flag C: %t, RegF[4]: %t", cpu.registers.f.getC(), (value&16) == 16)
			}
		})
	}
}

// LD r,r
func TestLDInstructions1(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	instructions := []struct {
		instruction instruction
		to, from    func() uint8
	}{
		{cpu.instructions[0x40], cpu.registers.getB, cpu.registers.getB},
		{cpu.instructions[0x41], cpu.registers.getB, cpu.registers.getC},
		{cpu.instructions[0x42], cpu.registers.getB, cpu.registers.getD},
		{cpu.instructions[0x43], cpu.registers.getB, cpu.registers.getE},
		{cpu.instructions[0x44], cpu.registers.getB, cpu.registers.getH},
		{cpu.instructions[0x45], cpu.registers.getB, cpu.registers.getL},
		{cpu.instructions[0x47], cpu.registers.getB, cpu.registers.getA},
		{cpu.instructions[0x48], cpu.registers.getC, cpu.registers.getB},
		{cpu.instructions[0x49], cpu.registers.getC, cpu.registers.getC},
		{cpu.instructions[0x4A], cpu.registers.getC, cpu.registers.getD},
		{cpu.instructions[0x4B], cpu.registers.getC, cpu.registers.getE},
		{cpu.instructions[0x4C], cpu.registers.getC, cpu.registers.getH},
		{cpu.instructions[0x4D], cpu.registers.getC, cpu.registers.getL},
		{cpu.instructions[0x4F], cpu.registers.getC, cpu.registers.getA},
		{cpu.instructions[0x50], cpu.registers.getD, cpu.registers.getB},
		{cpu.instructions[0x51], cpu.registers.getD, cpu.registers.getC},
		{cpu.instructions[0x52], cpu.registers.getD, cpu.registers.getD},
		{cpu.instructions[0x53], cpu.registers.getD, cpu.registers.getE},
		{cpu.instructions[0x54], cpu.registers.getD, cpu.registers.getH},
		{cpu.instructions[0x55], cpu.registers.getD, cpu.registers.getL},
		{cpu.instructions[0x57], cpu.registers.getD, cpu.registers.getA},
		{cpu.instructions[0x58], cpu.registers.getE, cpu.registers.getB},
		{cpu.instructions[0x59], cpu.registers.getE, cpu.registers.getC},
		{cpu.instructions[0x5A], cpu.registers.getE, cpu.registers.getD},
		{cpu.instructions[0x5B], cpu.registers.getE, cpu.registers.getE},
		{cpu.instructions[0x5C], cpu.registers.getE, cpu.registers.getH},
		{cpu.instructions[0x5D], cpu.registers.getE, cpu.registers.getL},
		{cpu.instructions[0x5F], cpu.registers.getE, cpu.registers.getA},
		{cpu.instructions[0x60], cpu.registers.getH, cpu.registers.getB},
		{cpu.instructions[0x61], cpu.registers.getH, cpu.registers.getC},
		{cpu.instructions[0x62], cpu.registers.getH, cpu.registers.getD},
		{cpu.instructions[0x63], cpu.registers.getH, cpu.registers.getE},
		{cpu.instructions[0x64], cpu.registers.getH, cpu.registers.getH},
		{cpu.instructions[0x65], cpu.registers.getH, cpu.registers.getL},
		{cpu.instructions[0x67], cpu.registers.getH, cpu.registers.getA},
		{cpu.instructions[0x68], cpu.registers.getL, cpu.registers.getB},
		{cpu.instructions[0x69], cpu.registers.getL, cpu.registers.getC},
		{cpu.instructions[0x6A], cpu.registers.getL, cpu.registers.getD},
		{cpu.instructions[0x6B], cpu.registers.getL, cpu.registers.getE},
		{cpu.instructions[0x6C], cpu.registers.getL, cpu.registers.getH},
		{cpu.instructions[0x6D], cpu.registers.getL, cpu.registers.getL},
		{cpu.instructions[0x6F], cpu.registers.getL, cpu.registers.getA},
		{cpu.instructions[0x78], cpu.registers.getA, cpu.registers.getB},
		{cpu.instructions[0x79], cpu.registers.getA, cpu.registers.getC},
		{cpu.instructions[0x7A], cpu.registers.getA, cpu.registers.getD},
		{cpu.instructions[0x7B], cpu.registers.getA, cpu.registers.getE},
		{cpu.instructions[0x7C], cpu.registers.getA, cpu.registers.getH},
		{cpu.instructions[0x7D], cpu.registers.getA, cpu.registers.getL},
		{cpu.instructions[0x7F], cpu.registers.getA, cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			tc.instruction.exec(cpu)
			value1 := tc.from()
			value2 := tc.to()
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
		})
	}
}

// LD r,n
func TestLDInstructions2(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		to          func() uint8
	}{
		{cpu.instructions[0x06], cpu.registers.getB},
		{cpu.instructions[0x0E], cpu.registers.getC},
		{cpu.instructions[0x16], cpu.registers.getD},
		{cpu.instructions[0x1E], cpu.registers.getE},
		{cpu.instructions[0x26], cpu.registers.getH},
		{cpu.instructions[0x2E], cpu.registers.getL},
		{cpu.instructions[0x3E], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			value1 := read8BitOperand(cpu)
			tc.instruction.exec(cpu)
			value2 := tc.to()
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
		})
	}
}

// LD r,(HL)
func TestLDInstructions3(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		to          func() uint8
	}{
		{cpu.instructions[0x46], cpu.registers.getB},
		{cpu.instructions[0x4E], cpu.registers.getC},
		{cpu.instructions[0x56], cpu.registers.getD},
		{cpu.instructions[0x5E], cpu.registers.getE},
		{cpu.instructions[0x66], cpu.registers.getH},
		{cpu.instructions[0x6E], cpu.registers.getL},
		{cpu.instructions[0x7E], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := cpu.registers.getHL()
			value1 := cpu.memory.Read(addr)
			tc.instruction.exec(cpu)
			value2 := tc.to()
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
		})
	}
}

// LD (HL),r
func TestLDInstructions4(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		from        func() uint8
	}{
		{cpu.instructions[0x70], cpu.registers.getB},
		{cpu.instructions[0x71], cpu.registers.getC},
		{cpu.instructions[0x72], cpu.registers.getD},
		{cpu.instructions[0x73], cpu.registers.getE},
		{cpu.instructions[0x74], cpu.registers.getH},
		{cpu.instructions[0x75], cpu.registers.getL},
		{cpu.instructions[0x77], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			tc.instruction.exec(cpu)
			addr := cpu.registers.getHL()
			if addr == controllers.ADDR_DIV_COUNTER { // writing to 0xFF04 does nothing; skip
				return
			}
			value1 := tc.from()
			value2 := cpu.memory.Read(addr)
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
		})
	}
}

// LD (HL),n
func TestLDInstructions5(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
	}{
		{cpu.instructions[0x36]},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := cpu.registers.getHL()
			value1 := read8BitOperand(cpu)
			tc.instruction.exec(cpu)
			value2 := cpu.memory.Read(addr)
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
		})
	}
}

// LD A,(BC)
func TestLDInstructions6(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		to          func() uint8
	}{
		{cpu.instructions[0x0A], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			tc.instruction.exec(cpu)
			addr := cpu.registers.getBC()
			value1 := cpu.memory.Read(addr)
			value2 := tc.to()
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
		})
	}
}

// LD A,(DE)
func TestLDInstructions7(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		to          func() uint8
	}{
		{cpu.instructions[0x1A], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			tc.instruction.exec(cpu)
			addr := cpu.registers.getDE()
			value1 := cpu.memory.Read(addr)
			value2 := tc.to()
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
		})
	}
}

// LD A,(nn)
func TestLDInstructions8(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		to          func() uint8
	}{
		{cpu.instructions[0xFA], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := read16BitOperand(cpu)
			tc.instruction.exec(cpu)
			value1 := cpu.memory.Read(addr)
			value2 := tc.to()
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
		})
	}
}

// LD (BC),A
func TestLDInstructions9(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		from        func() uint8
	}{
		{cpu.instructions[0x02], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			tc.instruction.exec(cpu)
			addr := cpu.registers.getBC()
			value1 := tc.from()
			value2 := cpu.memory.Read(addr)
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
		})
	}
}

// LD (DE),A
func TestLDInstructions10(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		from        func() uint8
	}{
		{cpu.instructions[0x12], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			tc.instruction.exec(cpu)
			addr := cpu.registers.getDE()
			value1 := tc.from()
			value2 := cpu.memory.Read(addr)
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
		})
	}
}

// LD (nn),A
func TestLDInstructions11(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		from        func() uint8
	}{
		{cpu.instructions[0xEA], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := read16BitOperand(cpu)
			tc.instruction.exec(cpu)
			value1 := tc.from()
			value2 := cpu.memory.Read(addr)
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
		})
	}
}

// LDH A,(n)
func TestLDInstructions12(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		to          func() uint8
	}{
		{cpu.instructions[0xF0], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			msb := uint16(0xFF)
			lsb := uint16(read8BitOperand(cpu))
			addr := msb<<8 | lsb
			tc.instruction.exec(cpu)
			value1 := cpu.memory.Read(addr)
			value2 := tc.to()
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
		})
	}
}

// LDH (n),A
func TestLDInstructions13(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		from        func() uint8
	}{
		{cpu.instructions[0xE0], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			msb := uint16(0xFF)
			lsb := uint16(read8BitOperand(cpu))
			addr := msb<<8 | lsb
			for addr == controllers.ADDR_DIV_COUNTER { // writing to 0xFF04 does nothing; randomize once more
				cpu.registers.setPC(uint16(rand.Int()))
				lsb = uint16(read8BitOperand(cpu))
				addr = msb<<8 | lsb
			}
			tc.instruction.exec(cpu)
			value1 := tc.from()
			value2 := cpu.memory.Read(addr)
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
			if value1 != value2 {
				log.Printf("adr: 0x%04X\n", addr)
			}
		})
	}
}

// LD A,(C)
func TestLDInstructions14(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		to, from    func() uint8
	}{
		{cpu.instructions[0xF2], cpu.registers.getA, cpu.registers.getC},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			tc.instruction.exec(cpu)
			msb := uint16(0xFF)
			lsb := uint16(tc.from())
			addr := msb<<8 | lsb
			value1 := cpu.memory.Read(addr)
			value2 := tc.to()
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
		})
	}
}

// LD (C),A
func TestLDInstructions15(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		to, from    func() uint8
	}{
		{cpu.instructions[0xE2], cpu.registers.getC, cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			tc.instruction.exec(cpu)
			msb := uint16(0xFF)
			lsb := uint16(tc.to())
			addr := msb<<8 | lsb
			if addr == controllers.ADDR_DIV_COUNTER { // writing to 0xFF04 does nothing; skip
				return
			}
			value1 := tc.from()
			value2 := cpu.memory.Read(addr)
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
		})
	}
}

// LDI (HL),A
func TestLDInstructions16(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		from        func() uint8
	}{
		{cpu.instructions[0x22], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := cpu.registers.getHL()
			tc.instruction.exec(cpu)
			value1 := tc.from()
			value2 := cpu.memory.Read(addr)
			prevAddr := addr
			currAddr := cpu.registers.getHL()
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
			if prevAddr+1 != currAddr {
				t.Errorf("Expected 0x%04X, got 0x%04X", prevAddr+1, currAddr)
			}
		})
	}
}

// LDI A,(HL)
func TestLDInstructions17(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		to          func() uint8
	}{
		{cpu.instructions[0x2A], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := cpu.registers.getHL()
			tc.instruction.exec(cpu)
			value1 := cpu.memory.Read(addr)
			value2 := tc.to()
			prevAddr := addr
			currAddr := cpu.registers.getHL()
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
			if prevAddr+1 != currAddr {
				t.Errorf("Expected 0x%04X, got 0x%04X", prevAddr+1, currAddr)
			}
		})
	}
}

// LDD (HL),A
func TestLDInstructions18(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		from        func() uint8
	}{
		{cpu.instructions[0x32], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := cpu.registers.getHL()
			if addr == controllers.ADDR_DIV_COUNTER { // writing to 0xFF04 does nothing; skip
				return
			}
			tc.instruction.exec(cpu)
			value1 := tc.from()
			value2 := cpu.memory.Read(addr)
			prevAddr := addr
			currAddr := cpu.registers.getHL()
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
			if prevAddr-1 != currAddr {
				t.Errorf("Expected 0x%04X, got 0x%04X", prevAddr-1, currAddr)
			}
		})
	}
}

// LDD A,(HL)
func TestLDInstructions19(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		to          func() uint8
	}{
		{cpu.instructions[0x3A], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := cpu.registers.getHL()
			tc.instruction.exec(cpu)
			value1 := cpu.memory.Read(addr)
			value2 := tc.to()
			prevAddr := addr
			currAddr := cpu.registers.getHL()
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
			if prevAddr-1 != currAddr {
				t.Errorf("Expected 0x%04X, got 0x%04X", prevAddr-1, currAddr)
			}
		})
	}
}

// LD rr,nn
func TestLDInstructions20(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		to          func() uint16
	}{
		{cpu.instructions[0x01], cpu.registers.getBC},
		{cpu.instructions[0x11], cpu.registers.getDE},
		{cpu.instructions[0x21], cpu.registers.getHL},
		{cpu.instructions[0x31], cpu.registers.getSP},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			lsb := uint16(cpu.memory.Read(cpu.registers.getPC()))
			msb := uint16(cpu.memory.Read(cpu.registers.getPC() + 1))
			value1 := msb<<8 | lsb
			tc.instruction.exec(cpu)
			value2 := tc.to()
			if value1 != value2 {
				t.Errorf("Expected 0x%04X, got 0x%04X", value1, value2)
			}
		})
	}
}

// LD (nn),SP
func TestLDInstructions21(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
	}{
		{cpu.instructions[0x08]},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := read16BitOperand(cpu)
			tc.instruction.exec(cpu)
			value1 := cpu.memory.Read(addr)
			value2 := cpu.memory.Read(addr + 1)
			if value1 != uint8(cpu.registers.getSP()&0xFF) {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, uint8(cpu.registers.getSP()&0xFF))
			}
			if value2 != uint8(cpu.registers.getSP()>>8) {
				t.Errorf("Expected 0x%02X, got 0x%02X", value2, uint8(cpu.registers.getSP()>>8))
			}
		})
	}
}

// LDHL SP+e
func TestLDInstructions22(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		to          func() uint16
	}{
		{cpu.instructions[0xF8], cpu.registers.getHL},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			sp := int(cpu.registers.getSP())
			e := int(int8(read8BitOperand(cpu)))
			tc.instruction.exec(cpu)
			value1 := uint16(sp + e)
			value2 := tc.to()
			flagH := false
			flagC := false
			if e < 0 {
				flagH = (sp & 0xF) < (e & 0xF)
				flagC = (sp & 0xFF) < (e & 0xFF)
			} else {
				flagH = (sp&0xF)+(e&0xF) > 0xF
				flagC = (sp&0xFF)+(e&0xFF) > 0xFF
			}
			if value1 != value2 {
				t.Errorf("Expected 0x%04X, got 0x%04X", value1, value2)
			}
			if false != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", false, cpu.registers.f.getZ())
			}
			if false != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", false, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// LD SP,HL
func TestLDInstructions23(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		to, from    func() uint16
	}{
		{cpu.instructions[0xF9], cpu.registers.getSP, cpu.registers.getHL},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			tc.instruction.exec(cpu)
			if tc.from() != tc.to() {
				t.Errorf("Expected 0x%04X, got 0x%04X", tc.from(), tc.to())
			}
		})
	}
}

// PUSH rr
func TestPUSHInstructions(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction  instruction
		From1, From2 func() uint8
	}{
		{cpu.instructions[0xC5], cpu.registers.getB, cpu.registers.getC},
		{cpu.instructions[0xD5], cpu.registers.getD, cpu.registers.getE},
		{cpu.instructions[0xE5], cpu.registers.getH, cpu.registers.getL},
		{cpu.instructions[0xF5], cpu.registers.getA, cpu.registers.getF},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := cpu.registers.getSP()
			if addr == controllers.ADDR_DIV_COUNTER+1 || addr == controllers.ADDR_DIV_COUNTER+2 { // writing to 0xFF04 does nothing; skip
				return
			}
			tc.instruction.exec(cpu)
			value1 := tc.From1()
			value2 := tc.From2()
			value3 := cpu.memory.Read(addr - 1)
			value4 := cpu.memory.Read(addr - 2)
			prevAddr := addr
			currAddr := cpu.registers.getSP()
			if value1 != value3 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value3)
			}
			if value2 != value4 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value2, value4)
			}
			if prevAddr-2 != currAddr {
				t.Errorf("Expected 0x%04X, got 0x%04X", prevAddr-2, currAddr)
			}
		})
	}
}

// POP rr
func TestPOPInstructions1(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		To1, To2    func() uint8
	}{
		{cpu.instructions[0xC1], cpu.registers.getC, cpu.registers.getB},
		{cpu.instructions[0xD1], cpu.registers.getE, cpu.registers.getD},
		{cpu.instructions[0xE1], cpu.registers.getL, cpu.registers.getH},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := cpu.registers.getSP()
			tc.instruction.exec(cpu)
			value1 := cpu.memory.Read(addr)
			value2 := cpu.memory.Read(addr + 1)
			value3 := tc.To1()
			value4 := tc.To2()
			prevAddr := addr
			currAddr := cpu.registers.getSP()
			if value1 != value3 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value3)
			}
			if value2 != value4 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value2, value4)
			}
			if prevAddr+2 != currAddr {
				t.Errorf("Expected 0x%04X, got 0x%04X", prevAddr+2, currAddr)
			}
		})
	}
}

// POP AF
func TestPOPInstructions2(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		To1, To2    func() uint8
	}{
		{cpu.instructions[0xF1], cpu.registers.getF, cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := cpu.registers.getSP()
			tc.instruction.exec(cpu)
			value1 := cpu.memory.Read(addr)
			value2 := cpu.memory.Read(addr + 1)
			value3 := tc.To1()
			value4 := tc.To2()
			prevAddr := addr
			currAddr := cpu.registers.getSP()
			if value1&0xF0 != value3 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1&0xF0, value3)
			}
			if value2 != value4 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value2, value4)
			}
			if prevAddr+2 != currAddr {
				t.Errorf("Expected 0x%04X, got 0x%04X", prevAddr+2, currAddr)
			}
		})
	}
}

// ADD A,r
func TestADDInstructions1(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	instructions := []struct {
		instruction instruction
		to, from    func() uint8
	}{
		{cpu.instructions[0x80], cpu.registers.getA, cpu.registers.getB},
		{cpu.instructions[0x81], cpu.registers.getA, cpu.registers.getC},
		{cpu.instructions[0x82], cpu.registers.getA, cpu.registers.getD},
		{cpu.instructions[0x83], cpu.registers.getA, cpu.registers.getE},
		{cpu.instructions[0x84], cpu.registers.getA, cpu.registers.getH},
		{cpu.instructions[0x85], cpu.registers.getA, cpu.registers.getL},
		{cpu.instructions[0x87], cpu.registers.getA, cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			value1 := tc.to()
			value2 := tc.from()
			flagZ := value1+value2 == 0
			flagN := false
			flagH := (value1&0xF)+(value2&0xF) > 0xF
			flagC := uint16(value1)+uint16(value2) > 0xFF
			tc.instruction.exec(cpu)
			if value1+value2 != tc.to() {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1+value2, tc.to())
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if flagN != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// ADD A,n
func TestADDInstructions2(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		to          func() uint8
	}{
		{cpu.instructions[0xC6], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			value1 := tc.to()
			value2 := read8BitOperand(cpu)
			flagZ := value1+value2 == 0
			flagN := false
			flagH := (value1&0xF)+(value2&0xF) > 0xF
			flagC := uint16(value1)+uint16(value2) > 0xFF
			tc.instruction.exec(cpu)
			if value1+value2 != tc.to() {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1+value2, tc.to())
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if flagN != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// ADD A,(HL)
func TestADDInstructions3(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		to          func() uint8
	}{
		{cpu.instructions[0x86], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := cpu.registers.getHL()
			value1 := tc.to()
			value2 := cpu.memory.Read(addr)
			flagZ := value1+value2 == 0
			flagN := false
			flagH := (value1&0xF)+(value2&0xF) > 0xF
			flagC := uint16(value1)+uint16(value2) > 0xFF
			tc.instruction.exec(cpu)
			if value1+value2 != tc.to() {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1+value2, tc.to())
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if flagN != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// ADD HL,rr
func TestADDInstructions4(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	instructions := []struct {
		instruction instruction
		to, from    func() uint16
	}{
		{cpu.instructions[0x09], cpu.registers.getHL, cpu.registers.getBC},
		{cpu.instructions[0x19], cpu.registers.getHL, cpu.registers.getDE},
		{cpu.instructions[0x29], cpu.registers.getHL, cpu.registers.getHL},
		{cpu.instructions[0x39], cpu.registers.getHL, cpu.registers.getSP},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			value1 := tc.to()
			value2 := tc.from()
			flagZ := cpu.registers.f.getZ()
			flagN := false
			flagH := (value1&0xFFF)+(value2&0xFFF) > 0xFFF
			flagC := uint32(value1)+uint32(value2) > 0xFFFF
			tc.instruction.exec(cpu)
			if value1+value2 != tc.to() {
				t.Errorf("Expected 0x%04X, got 0x%04X", value1+value2, tc.to())
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if flagN != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// ADD SP,e
func TestADDInstructions5(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		to          func() uint16
	}{
		{cpu.instructions[0xE8], cpu.registers.getSP},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			for i := 0; i < 1000; i++ {
				randomizeRegisters(cpu.registers)
				value1 := int(cpu.registers.getSP())
				value2 := int(int8(read8BitOperand(cpu)))
				flagZ := false
				flagN := false
				flagH := false
				flagC := false
				if value2 < 0 {
					flagH = (value1 & 0xF) < (value2 & 0xF)
					flagC = (value1 & 0xFF) < (value2 & 0xFF)
				} else {
					flagH = (value1&0xF)+(value2&0xF) > 0xF
					flagC = (value1&0xFF)+(value2&0xFF) > 0xFF
				}
				tc.instruction.exec(cpu)
				if uint16(value1+value2) != tc.to() {
					t.Errorf("Expected 0x%04X, got 0x%04X", uint16(value1+value2), tc.to())
				}
				if flagZ != cpu.registers.f.getZ() {
					t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
				}
				if flagN != cpu.registers.f.getN() {
					t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
				}
				if flagH != cpu.registers.f.getH() {
					t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
				}
				if flagC != cpu.registers.f.getC() {
					t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
				}
			}
		})
	}
}

// ADC A,r
func TestADCInstructions1(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	instructions := []struct {
		instruction instruction
		to, from    func() uint8
	}{
		{cpu.instructions[0x88], cpu.registers.getA, cpu.registers.getB},
		{cpu.instructions[0x89], cpu.registers.getA, cpu.registers.getC},
		{cpu.instructions[0x8A], cpu.registers.getA, cpu.registers.getD},
		{cpu.instructions[0x8B], cpu.registers.getA, cpu.registers.getE},
		{cpu.instructions[0x8C], cpu.registers.getA, cpu.registers.getH},
		{cpu.instructions[0x8D], cpu.registers.getA, cpu.registers.getL},
		{cpu.instructions[0x8F], cpu.registers.getA, cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			carry := cpu.registers.f.getCarry()
			value1 := tc.to()
			value2 := tc.from()
			flagZ := value1+value2+carry == 0
			flagN := false
			flagH := (value1&0xF)+(value2&0xF)+carry > 0xF
			flagC := uint16(value1)+uint16(value2)+uint16(carry) > 0xFF
			tc.instruction.exec(cpu)
			if value1+value2+carry != tc.to() {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1+value2+carry, tc.to())
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if flagN != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// ADC A,n
func TestADCInstructions2(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		to          func() uint8
	}{
		{cpu.instructions[0xCE], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			carry := cpu.registers.f.getCarry()
			value1 := tc.to()
			value2 := read8BitOperand(cpu)
			flagZ := value1+value2+carry == 0
			flagN := false
			flagH := (value1&0xF)+(value2&0xF)+carry > 0xF
			flagC := uint16(value1)+uint16(value2)+uint16(carry) > 0xFF
			tc.instruction.exec(cpu)
			if value1+value2+carry != tc.to() {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1+value2+carry, tc.to())
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if flagN != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// ADC A,(HL)
func TestADCInstructions3(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		to          func() uint8
	}{
		{cpu.instructions[0x8E], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := cpu.registers.getHL()
			carry := cpu.registers.f.getCarry()
			value1 := tc.to()
			value2 := cpu.memory.Read(addr)
			flagZ := value1+value2+carry == 0
			flagN := false
			flagH := (value1&0xF)+(value2&0xF)+carry > 0xF
			flagC := uint16(value1)+uint16(value2)+uint16(carry) > 0xFF
			tc.instruction.exec(cpu)
			if value1+value2+carry != tc.to() {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1+value2+carry, tc.to())
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if flagN != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// SUB A,r
func TestSUBInstructions1(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	instructions := []struct {
		instruction instruction
		to, from    func() uint8
	}{
		{cpu.instructions[0x90], cpu.registers.getA, cpu.registers.getB},
		{cpu.instructions[0x91], cpu.registers.getA, cpu.registers.getC},
		{cpu.instructions[0x92], cpu.registers.getA, cpu.registers.getD},
		{cpu.instructions[0x93], cpu.registers.getA, cpu.registers.getE},
		{cpu.instructions[0x94], cpu.registers.getA, cpu.registers.getH},
		{cpu.instructions[0x95], cpu.registers.getA, cpu.registers.getL},
		{cpu.instructions[0x97], cpu.registers.getA, cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			value1 := tc.to()
			value2 := tc.from()
			flagZ := value1-value2 == 0
			flagN := true
			flagH := (value1 & 0xF) < (value2 & 0xF)
			flagC := uint16(value1) < uint16(value2)
			tc.instruction.exec(cpu)
			if value1-value2 != tc.to() {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1-value2, tc.to())
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if flagN != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// SUB A,n
func TestSUBInstructions2(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		to          func() uint8
	}{
		{cpu.instructions[0xD6], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			value1 := tc.to()
			value2 := read8BitOperand(cpu)
			flagZ := value1-value2 == 0
			flagN := true
			flagH := (value1 & 0xF) < (value2 & 0xF)
			flagC := uint16(value1) < uint16(value2)
			tc.instruction.exec(cpu)
			if value1-value2 != tc.to() {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1-value2, tc.to())
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if flagN != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// SUB A,(HL)
func TestSUBInstructions3(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		to          func() uint8
	}{
		{cpu.instructions[0x96], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := cpu.registers.getHL()
			value1 := tc.to()
			value2 := cpu.memory.Read(addr)
			flagZ := value1-value2 == 0
			flagN := true
			flagH := (value1 & 0xF) < (value2 & 0xF)
			flagC := uint16(value1) < uint16(value2)
			tc.instruction.exec(cpu)
			if value1-value2 != tc.to() {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1-value2, tc.to())
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if flagN != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// SBC A,r
func TestSBCInstructions1(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	instructions := []struct {
		instruction instruction
		to, from    func() uint8
	}{
		{cpu.instructions[0x98], cpu.registers.getA, cpu.registers.getB},
		{cpu.instructions[0x99], cpu.registers.getA, cpu.registers.getC},
		{cpu.instructions[0x9A], cpu.registers.getA, cpu.registers.getD},
		{cpu.instructions[0x9B], cpu.registers.getA, cpu.registers.getE},
		{cpu.instructions[0x9C], cpu.registers.getA, cpu.registers.getH},
		{cpu.instructions[0x9D], cpu.registers.getA, cpu.registers.getL},
		{cpu.instructions[0x9F], cpu.registers.getA, cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			carry := cpu.registers.f.getCarry()
			value1 := tc.to()
			value2 := tc.from()
			flagZ := value1-value2-carry == 0
			flagN := true
			flagH := (value1 & 0xF) < ((value2 & 0xF) + carry)
			flagC := uint16(value1) < (uint16(value2) + uint16(carry))
			tc.instruction.exec(cpu)
			if value1-value2-carry != tc.to() {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1-value2-carry, tc.to())
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if flagN != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// SBC A,n
func TestSBCInstructions2(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		to          func() uint8
	}{
		{cpu.instructions[0xDE], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			carry := cpu.registers.f.getCarry()
			value1 := tc.to()
			value2 := read8BitOperand(cpu)
			flagZ := value1-value2-carry == 0
			flagN := true
			flagH := (value1 & 0xF) < ((value2 & 0xF) + carry)
			flagC := uint16(value1) < (uint16(value2) + uint16(carry))
			tc.instruction.exec(cpu)
			if value1-value2-carry != tc.to() {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1-value2-carry, tc.to())
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if flagN != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// SBC A,(HL)
func TestSBCInstructions3(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		to          func() uint8
	}{
		{cpu.instructions[0x9E], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := cpu.registers.getHL()
			carry := cpu.registers.f.getCarry()
			value1 := tc.to()
			value2 := cpu.memory.Read(addr)
			flagZ := value1-value2-carry == 0
			flagN := true
			flagH := (value1 & 0xF) < ((value2 & 0xF) + carry)
			flagC := uint16(value1) < (uint16(value2) + uint16(carry))
			tc.instruction.exec(cpu)
			if value1-value2-carry != tc.to() {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1-value2-carry, tc.to())
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if flagN != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// AND A,r
func TestANDInstructions1(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	instructions := []struct {
		instruction instruction
		to, from    func() uint8
	}{
		{cpu.instructions[0xA0], cpu.registers.getA, cpu.registers.getB},
		{cpu.instructions[0xA1], cpu.registers.getA, cpu.registers.getC},
		{cpu.instructions[0xA2], cpu.registers.getA, cpu.registers.getD},
		{cpu.instructions[0xA3], cpu.registers.getA, cpu.registers.getE},
		{cpu.instructions[0xA4], cpu.registers.getA, cpu.registers.getH},
		{cpu.instructions[0xA5], cpu.registers.getA, cpu.registers.getL},
		{cpu.instructions[0xA7], cpu.registers.getA, cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			value1 := tc.to()
			value2 := tc.from()
			flagZ := value1&value2 == 0
			flagN := false
			flagH := true
			flagC := false
			tc.instruction.exec(cpu)
			if value1&value2 != tc.to() {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1&value2, tc.to())
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if flagN != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// AND A,n
func TestANDInstructions2(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		to          func() uint8
	}{
		{cpu.instructions[0xE6], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			value1 := tc.to()
			value2 := read8BitOperand(cpu)
			flagZ := value1&value2 == 0
			flagN := false
			flagH := true
			flagC := false
			tc.instruction.exec(cpu)
			if value1&value2 != tc.to() {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1&value2, tc.to())
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if flagN != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// AND A,(HL)
func TestANDInstructions3(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		to          func() uint8
	}{
		{cpu.instructions[0xA6], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := cpu.registers.getHL()
			value1 := tc.to()
			value2 := cpu.memory.Read(addr)
			flagZ := value1&value2 == 0
			flagN := false
			flagH := true
			flagC := false
			tc.instruction.exec(cpu)
			if value1&value2 != tc.to() {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1&value2, tc.to())
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if flagN != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// XOR A,r
func TestXORInstructions1(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	instructions := []struct {
		instruction instruction
		to, from    func() uint8
	}{
		{cpu.instructions[0xA8], cpu.registers.getA, cpu.registers.getB},
		{cpu.instructions[0xA9], cpu.registers.getA, cpu.registers.getC},
		{cpu.instructions[0xAA], cpu.registers.getA, cpu.registers.getD},
		{cpu.instructions[0xAB], cpu.registers.getA, cpu.registers.getE},
		{cpu.instructions[0xAC], cpu.registers.getA, cpu.registers.getH},
		{cpu.instructions[0xAD], cpu.registers.getA, cpu.registers.getL},
		{cpu.instructions[0xAF], cpu.registers.getA, cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			value1 := tc.to()
			value2 := tc.from()
			flagZ := value1^value2 == 0
			flagN := false
			flagH := false
			flagC := false
			tc.instruction.exec(cpu)
			if value1^value2 != tc.to() {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1^value2, tc.to())
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if flagN != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// XOR A,n
func TestXORInstructions2(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		to          func() uint8
	}{
		{cpu.instructions[0xEE], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			value1 := tc.to()
			value2 := read8BitOperand(cpu)
			flagZ := value1^value2 == 0
			flagN := false
			flagH := false
			flagC := false
			tc.instruction.exec(cpu)
			if value1^value2 != tc.to() {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1^value2, tc.to())
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if flagN != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// XOR A,(HL)
func TestXORInstructions3(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		to          func() uint8
	}{
		{cpu.instructions[0xAE], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := cpu.registers.getHL()
			value1 := cpu.registers.getA()
			value2 := cpu.memory.Read(addr)
			flagZ := value1^value2 == 0
			flagN := false
			flagH := false
			flagC := false
			tc.instruction.exec(cpu)
			if value1^value2 != tc.to() {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1^value2, tc.to())
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if flagN != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// OR A,r
func TestORInstructions1(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	instructions := []struct {
		instruction instruction
		to, from    func() uint8
	}{
		{cpu.instructions[0xB0], cpu.registers.getA, cpu.registers.getB},
		{cpu.instructions[0xB1], cpu.registers.getA, cpu.registers.getC},
		{cpu.instructions[0xB2], cpu.registers.getA, cpu.registers.getD},
		{cpu.instructions[0xB3], cpu.registers.getA, cpu.registers.getE},
		{cpu.instructions[0xB4], cpu.registers.getA, cpu.registers.getH},
		{cpu.instructions[0xB5], cpu.registers.getA, cpu.registers.getL},
		{cpu.instructions[0xB7], cpu.registers.getA, cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			value1 := tc.to()
			value2 := tc.from()
			flagZ := value1|value2 == 0
			flagN := false
			flagH := false
			flagC := false
			tc.instruction.exec(cpu)
			if value1|value2 != tc.to() {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1|value2, tc.to())
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if flagN != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// OR A,n
func TestORInstructions2(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		to          func() uint8
	}{
		{cpu.instructions[0xF6], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			value1 := tc.to()
			value2 := read8BitOperand(cpu)
			flagZ := value1|value2 == 0
			flagN := false
			flagH := false
			flagC := false
			tc.instruction.exec(cpu)
			if value1|value2 != tc.to() {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1|value2, tc.to())
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if flagN != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// OR A,(HL)
func TestORInstructions3(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		to          func() uint8
	}{
		{cpu.instructions[0xB6], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := cpu.registers.getHL()
			value1 := tc.to()
			value2 := cpu.memory.Read(addr)
			flagZ := value1|value2 == 0
			flagN := false
			flagH := false
			flagC := false
			tc.instruction.exec(cpu)
			if value1|value2 != tc.to() {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1|value2, tc.to())
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if flagN != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// CP A,r
func TestCPInstructions1(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	instructions := []struct {
		instruction instruction
		to, from    func() uint8
	}{
		{cpu.instructions[0xB8], cpu.registers.getA, cpu.registers.getB},
		{cpu.instructions[0xB9], cpu.registers.getA, cpu.registers.getC},
		{cpu.instructions[0xBA], cpu.registers.getA, cpu.registers.getD},
		{cpu.instructions[0xBB], cpu.registers.getA, cpu.registers.getE},
		{cpu.instructions[0xBC], cpu.registers.getA, cpu.registers.getH},
		{cpu.instructions[0xBD], cpu.registers.getA, cpu.registers.getL},
		{cpu.instructions[0xBF], cpu.registers.getA, cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			value1 := tc.to()
			value2 := tc.from()
			flagZ := value1 == value2
			flagN := true
			flagH := (value1 & 0xF) < (value2 & 0xF)
			flagC := value1 < value2
			tc.instruction.exec(cpu)
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if flagN != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// CP A,n
func TestCPInstructions2(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		to          func() uint8
	}{
		{cpu.instructions[0xD6], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			value1 := tc.to()
			value2 := read8BitOperand(cpu)
			flagZ := value1 == value2
			flagN := true
			flagH := (value1 & 0xF) < (value2 & 0xF)
			flagC := value1 < value2
			tc.instruction.exec(cpu)
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if flagN != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// CP A,(HL)
func TestCPInstructions3(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		to          func() uint8
	}{
		{cpu.instructions[0xBE], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := cpu.registers.getHL()
			value1 := tc.to()
			value2 := cpu.memory.Read(addr)
			flagZ := value1 == value2
			flagN := true
			flagH := (value1 & 0xF) < (value2 & 0xF)
			flagC := value1 < value2
			tc.instruction.exec(cpu)
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if flagN != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// INC r
func TestINCInstructions1(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	instructions := []struct {
		instruction instruction
		from        func() uint8
	}{
		{cpu.instructions[0x04], cpu.registers.getB},
		{cpu.instructions[0x0C], cpu.registers.getC},
		{cpu.instructions[0x14], cpu.registers.getD},
		{cpu.instructions[0x1C], cpu.registers.getE},
		{cpu.instructions[0x24], cpu.registers.getH},
		{cpu.instructions[0x2C], cpu.registers.getL},
		{cpu.instructions[0x3C], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			value1 := tc.from()
			value2 := value1 + 1
			flagZ := value2 == 0
			flagN := false
			flagH := value2&0xF == 0
			tc.instruction.exec(cpu)
			if value2 != tc.from() {
				t.Errorf("Expected 0x%02X, got 0x%02X", value2, tc.from())
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if flagN != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
		})
	}
}

// INC (HL)
func TestINCInstructions2(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
	}{
		{cpu.instructions[0x34]},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := cpu.registers.getHL()
			value1 := cpu.memory.Read(addr)
			value2 := value1 + 1
			flagZ := value2 == 0
			flagN := false
			flagH := value2&0xF == 0
			tc.instruction.exec(cpu)
			if value2 != cpu.memory.Read(addr) {
				t.Errorf("Expected 0x%02X, got 0x%02X", value2, cpu.memory.Read(addr))
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if flagN != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
		})
	}
}

// INC rr
func TestINCInstructions3(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		from        func() uint16
	}{
		{cpu.instructions[0x03], cpu.registers.getBC},
		{cpu.instructions[0x13], cpu.registers.getDE},
		{cpu.instructions[0x23], cpu.registers.getHL},
		{cpu.instructions[0x33], cpu.registers.getSP},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			value1 := tc.from()
			value2 := value1 + 1
			tc.instruction.exec(cpu)
			if value2 != tc.from() {
				t.Errorf("Expected 0x%04X, got 0x%04X", value2, tc.from())
			}
		})
	}
}

// DEC r
func TestDECInstructions1(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	instructions := []struct {
		instruction instruction
		from        func() uint8
	}{
		{cpu.instructions[0x05], cpu.registers.getB},
		{cpu.instructions[0x0D], cpu.registers.getC},
		{cpu.instructions[0x15], cpu.registers.getD},
		{cpu.instructions[0x1D], cpu.registers.getE},
		{cpu.instructions[0x25], cpu.registers.getH},
		{cpu.instructions[0x2D], cpu.registers.getL},
		{cpu.instructions[0x3D], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			value1 := tc.from()
			value2 := value1 - 1
			flagZ := value2 == 0
			flagN := true
			flagH := value1&0xF == 0
			tc.instruction.exec(cpu)
			if value2 != tc.from() {
				t.Errorf("Expected 0x%02X, got 0x%02X", value2, tc.from())
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if flagN != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
		})
	}
}

// DEC (HL)
func TestDECInstructions2(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
	}{
		{cpu.instructions[0x35]},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := cpu.registers.getHL()
			value1 := cpu.memory.Read(addr)
			value2 := value1 - 1
			flagZ := value2 == 0
			flagN := true
			flagH := value1&0xF == 0
			tc.instruction.exec(cpu)
			if value2 != cpu.memory.Read(addr) {
				t.Errorf("Expected 0x%02X, got 0x%02X", value2, cpu.memory.Read(addr))
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if flagN != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
		})
	}
}

// DEC rr
func TestDECInstructions3(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		from        func() uint16
	}{
		{cpu.instructions[0x0B], cpu.registers.getBC},
		{cpu.instructions[0x1B], cpu.registers.getDE},
		{cpu.instructions[0x2B], cpu.registers.getHL},
		{cpu.instructions[0x3B], cpu.registers.getSP},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			value1 := tc.from()
			value2 := value1 - 1
			tc.instruction.exec(cpu)
			if value2 != tc.from() {
				t.Errorf("Expected 0x%04X, got 0x%04X", value2, tc.from())
			}
		})
	}
}

// DAA
func TestDAAInstructions(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	instructions := []struct {
		instruction instruction
	}{
		{cpu.instructions[0x27]},
	}
	operations := []struct {
		PrevA                      uint8
		PrevZ, PrevN, PrevH, PrevC bool
		CurrA                      uint8
		CurrZ, CurrN, CurrH, CurrC bool
	}{
		{0x00, false, false, false, false, 0x00, true, false, false, false},
		{0x00, false, false, false, true, 0x60, false, false, false, true},
		{0x00, false, false, true, false, 0x06, false, false, false, false},
		{0x00, false, false, true, true, 0x66, false, false, false, true},
		{0x00, false, true, false, false, 0x00, true, true, false, false},
		{0x00, false, true, false, true, 0xa0, false, true, false, true},
		{0x00, false, true, true, false, 0xfa, false, true, false, false},
		{0x00, false, true, true, true, 0x9a, false, true, false, true},
		{0x00, true, false, false, false, 0x00, true, false, false, false},
		{0x00, true, false, false, true, 0x60, false, false, false, true},
		{0x00, true, false, true, false, 0x06, false, false, false, false},
		{0x00, true, false, true, true, 0x66, false, false, false, true},
		{0x00, true, true, false, false, 0x00, true, true, false, false},
		{0x00, true, true, false, true, 0xa0, false, true, false, true},
		{0x00, true, true, true, false, 0xfa, false, true, false, false},
		{0x00, true, true, true, true, 0x9a, false, true, false, true},
		{0x01, false, false, false, false, 0x01, false, false, false, false},
		{0x01, false, false, false, true, 0x61, false, false, false, true},
		{0x01, false, false, true, false, 0x07, false, false, false, false},
		{0x01, false, false, true, true, 0x67, false, false, false, true},
		{0x01, false, true, false, false, 0x01, false, true, false, false},
		{0x01, false, true, false, true, 0xa1, false, true, false, true},
		{0x01, false, true, true, false, 0xfb, false, true, false, false},
		{0x01, false, true, true, true, 0x9b, false, true, false, true},
		{0x01, true, false, false, false, 0x01, false, false, false, false},
		{0x01, true, false, false, true, 0x61, false, false, false, true},
		{0x01, true, false, true, false, 0x07, false, false, false, false},
		{0x01, true, false, true, true, 0x67, false, false, false, true},
		{0x01, true, true, false, false, 0x01, false, true, false, false},
		{0x01, true, true, false, true, 0xa1, false, true, false, true},
		{0x01, true, true, true, false, 0xfb, false, true, false, false},
		{0x01, true, true, true, true, 0x9b, false, true, false, true},
		{0x02, false, false, false, false, 0x02, false, false, false, false},
		{0x02, false, false, false, true, 0x62, false, false, false, true},
		{0x02, false, false, true, false, 0x08, false, false, false, false},
		{0x02, false, false, true, true, 0x68, false, false, false, true},
		{0x02, false, true, false, false, 0x02, false, true, false, false},
		{0x02, false, true, false, true, 0xa2, false, true, false, true},
		{0x02, false, true, true, false, 0xfc, false, true, false, false},
		{0x02, false, true, true, true, 0x9c, false, true, false, true},
		{0x02, true, false, false, false, 0x02, false, false, false, false},
		{0x02, true, false, false, true, 0x62, false, false, false, true},
		{0x02, true, false, true, false, 0x08, false, false, false, false},
		{0x02, true, false, true, true, 0x68, false, false, false, true},
		{0x02, true, true, false, false, 0x02, false, true, false, false},
		{0x02, true, true, false, true, 0xa2, false, true, false, true},
		{0x02, true, true, true, false, 0xfc, false, true, false, false},
		{0x02, true, true, true, true, 0x9c, false, true, false, true},
		{0x03, false, false, false, false, 0x03, false, false, false, false},
		{0x03, false, false, false, true, 0x63, false, false, false, true},
		{0x03, false, false, true, false, 0x09, false, false, false, false},
		{0x03, false, false, true, true, 0x69, false, false, false, true},
		{0x03, false, true, false, false, 0x03, false, true, false, false},
		{0x03, false, true, false, true, 0xa3, false, true, false, true},
		{0x03, false, true, true, false, 0xfd, false, true, false, false},
		{0x03, false, true, true, true, 0x9d, false, true, false, true},
		{0x03, true, false, false, false, 0x03, false, false, false, false},
		{0x03, true, false, false, true, 0x63, false, false, false, true},
		{0x03, true, false, true, false, 0x09, false, false, false, false},
		{0x03, true, false, true, true, 0x69, false, false, false, true},
		{0x03, true, true, false, false, 0x03, false, true, false, false},
		{0x03, true, true, false, true, 0xa3, false, true, false, true},
		{0x03, true, true, true, false, 0xfd, false, true, false, false},
		{0x03, true, true, true, true, 0x9d, false, true, false, true},
		{0x04, false, false, false, false, 0x04, false, false, false, false},
		{0x04, false, false, false, true, 0x64, false, false, false, true},
		{0x04, false, false, true, false, 0x0a, false, false, false, false},
		{0x04, false, false, true, true, 0x6a, false, false, false, true},
		{0x04, false, true, false, false, 0x04, false, true, false, false},
		{0x04, false, true, false, true, 0xa4, false, true, false, true},
		{0x04, false, true, true, false, 0xfe, false, true, false, false},
		{0x04, false, true, true, true, 0x9e, false, true, false, true},
		{0x04, true, false, false, false, 0x04, false, false, false, false},
		{0x04, true, false, false, true, 0x64, false, false, false, true},
		{0x04, true, false, true, false, 0x0a, false, false, false, false},
		{0x04, true, false, true, true, 0x6a, false, false, false, true},
		{0x04, true, true, false, false, 0x04, false, true, false, false},
		{0x04, true, true, false, true, 0xa4, false, true, false, true},
		{0x04, true, true, true, false, 0xfe, false, true, false, false},
		{0x04, true, true, true, true, 0x9e, false, true, false, true},
		{0x05, false, false, false, false, 0x05, false, false, false, false},
		{0x05, false, false, false, true, 0x65, false, false, false, true},
		{0x05, false, false, true, false, 0x0b, false, false, false, false},
		{0x05, false, false, true, true, 0x6b, false, false, false, true},
		{0x05, false, true, false, false, 0x05, false, true, false, false},
		{0x05, false, true, false, true, 0xa5, false, true, false, true},
		{0x05, false, true, true, false, 0xff, false, true, false, false},
		{0x05, false, true, true, true, 0x9f, false, true, false, true},
		{0x05, true, false, false, false, 0x05, false, false, false, false},
		{0x05, true, false, false, true, 0x65, false, false, false, true},
		{0x05, true, false, true, false, 0x0b, false, false, false, false},
		{0x05, true, false, true, true, 0x6b, false, false, false, true},
		{0x05, true, true, false, false, 0x05, false, true, false, false},
		{0x05, true, true, false, true, 0xa5, false, true, false, true},
		{0x05, true, true, true, false, 0xff, false, true, false, false},
		{0x05, true, true, true, true, 0x9f, false, true, false, true},
		{0x06, false, false, false, false, 0x06, false, false, false, false},
		{0x06, false, false, false, true, 0x66, false, false, false, true},
		{0x06, false, false, true, false, 0x0c, false, false, false, false},
		{0x06, false, false, true, true, 0x6c, false, false, false, true},
		{0x06, false, true, false, false, 0x06, false, true, false, false},
		{0x06, false, true, false, true, 0xa6, false, true, false, true},
		{0x06, false, true, true, false, 0x00, true, true, false, false},
		{0x06, false, true, true, true, 0xa0, false, true, false, true},
		{0x06, true, false, false, false, 0x06, false, false, false, false},
		{0x06, true, false, false, true, 0x66, false, false, false, true},
		{0x06, true, false, true, false, 0x0c, false, false, false, false},
		{0x06, true, false, true, true, 0x6c, false, false, false, true},
		{0x06, true, true, false, false, 0x06, false, true, false, false},
		{0x06, true, true, false, true, 0xa6, false, true, false, true},
		{0x06, true, true, true, false, 0x00, true, true, false, false},
		{0x06, true, true, true, true, 0xa0, false, true, false, true},
		{0x07, false, false, false, false, 0x07, false, false, false, false},
		{0x07, false, false, false, true, 0x67, false, false, false, true},
		{0x07, false, false, true, false, 0x0d, false, false, false, false},
		{0x07, false, false, true, true, 0x6d, false, false, false, true},
		{0x07, false, true, false, false, 0x07, false, true, false, false},
		{0x07, false, true, false, true, 0xa7, false, true, false, true},
		{0x07, false, true, true, false, 0x01, false, true, false, false},
		{0x07, false, true, true, true, 0xa1, false, true, false, true},
		{0x07, true, false, false, false, 0x07, false, false, false, false},
		{0x07, true, false, false, true, 0x67, false, false, false, true},
		{0x07, true, false, true, false, 0x0d, false, false, false, false},
		{0x07, true, false, true, true, 0x6d, false, false, false, true},
		{0x07, true, true, false, false, 0x07, false, true, false, false},
		{0x07, true, true, false, true, 0xa7, false, true, false, true},
		{0x07, true, true, true, false, 0x01, false, true, false, false},
		{0x07, true, true, true, true, 0xa1, false, true, false, true},
		{0x08, false, false, false, false, 0x08, false, false, false, false},
		{0x08, false, false, false, true, 0x68, false, false, false, true},
		{0x08, false, false, true, false, 0x0e, false, false, false, false},
		{0x08, false, false, true, true, 0x6e, false, false, false, true},
		{0x08, false, true, false, false, 0x08, false, true, false, false},
		{0x08, false, true, false, true, 0xa8, false, true, false, true},
		{0x08, false, true, true, false, 0x02, false, true, false, false},
		{0x08, false, true, true, true, 0xa2, false, true, false, true},
		{0x08, true, false, false, false, 0x08, false, false, false, false},
		{0x08, true, false, false, true, 0x68, false, false, false, true},
		{0x08, true, false, true, false, 0x0e, false, false, false, false},
		{0x08, true, false, true, true, 0x6e, false, false, false, true},
		{0x08, true, true, false, false, 0x08, false, true, false, false},
		{0x08, true, true, false, true, 0xa8, false, true, false, true},
		{0x08, true, true, true, false, 0x02, false, true, false, false},
		{0x08, true, true, true, true, 0xa2, false, true, false, true},
		{0x09, false, false, false, false, 0x09, false, false, false, false},
		{0x09, false, false, false, true, 0x69, false, false, false, true},
		{0x09, false, false, true, false, 0x0f, false, false, false, false},
		{0x09, false, false, true, true, 0x6f, false, false, false, true},
		{0x09, false, true, false, false, 0x09, false, true, false, false},
		{0x09, false, true, false, true, 0xa9, false, true, false, true},
		{0x09, false, true, true, false, 0x03, false, true, false, false},
		{0x09, false, true, true, true, 0xa3, false, true, false, true},
		{0x09, true, false, false, false, 0x09, false, false, false, false},
		{0x09, true, false, false, true, 0x69, false, false, false, true},
		{0x09, true, false, true, false, 0x0f, false, false, false, false},
		{0x09, true, false, true, true, 0x6f, false, false, false, true},
		{0x09, true, true, false, false, 0x09, false, true, false, false},
		{0x09, true, true, false, true, 0xa9, false, true, false, true},
		{0x09, true, true, true, false, 0x03, false, true, false, false},
		{0x09, true, true, true, true, 0xa3, false, true, false, true},
		{0x0a, false, false, false, false, 0x10, false, false, false, false},
		{0x0a, false, false, false, true, 0x70, false, false, false, true},
		{0x0a, false, false, true, false, 0x10, false, false, false, false},
		{0x0a, false, false, true, true, 0x70, false, false, false, true},
		{0x0a, false, true, false, false, 0x0a, false, true, false, false},
		{0x0a, false, true, false, true, 0xaa, false, true, false, true},
		{0x0a, false, true, true, false, 0x04, false, true, false, false},
		{0x0a, false, true, true, true, 0xa4, false, true, false, true},
		{0x0a, true, false, false, false, 0x10, false, false, false, false},
		{0x0a, true, false, false, true, 0x70, false, false, false, true},
		{0x0a, true, false, true, false, 0x10, false, false, false, false},
		{0x0a, true, false, true, true, 0x70, false, false, false, true},
		{0x0a, true, true, false, false, 0x0a, false, true, false, false},
		{0x0a, true, true, false, true, 0xaa, false, true, false, true},
		{0x0a, true, true, true, false, 0x04, false, true, false, false},
		{0x0a, true, true, true, true, 0xa4, false, true, false, true},
		{0x0b, false, false, false, false, 0x11, false, false, false, false},
		{0x0b, false, false, false, true, 0x71, false, false, false, true},
		{0x0b, false, false, true, false, 0x11, false, false, false, false},
		{0x0b, false, false, true, true, 0x71, false, false, false, true},
		{0x0b, false, true, false, false, 0x0b, false, true, false, false},
		{0x0b, false, true, false, true, 0xab, false, true, false, true},
		{0x0b, false, true, true, false, 0x05, false, true, false, false},
		{0x0b, false, true, true, true, 0xa5, false, true, false, true},
		{0x0b, true, false, false, false, 0x11, false, false, false, false},
		{0x0b, true, false, false, true, 0x71, false, false, false, true},
		{0x0b, true, false, true, false, 0x11, false, false, false, false},
		{0x0b, true, false, true, true, 0x71, false, false, false, true},
		{0x0b, true, true, false, false, 0x0b, false, true, false, false},
		{0x0b, true, true, false, true, 0xab, false, true, false, true},
		{0x0b, true, true, true, false, 0x05, false, true, false, false},
		{0x0b, true, true, true, true, 0xa5, false, true, false, true},
		{0x0c, false, false, false, false, 0x12, false, false, false, false},
		{0x0c, false, false, false, true, 0x72, false, false, false, true},
		{0x0c, false, false, true, false, 0x12, false, false, false, false},
		{0x0c, false, false, true, true, 0x72, false, false, false, true},
		{0x0c, false, true, false, false, 0x0c, false, true, false, false},
		{0x0c, false, true, false, true, 0xac, false, true, false, true},
		{0x0c, false, true, true, false, 0x06, false, true, false, false},
		{0x0c, false, true, true, true, 0xa6, false, true, false, true},
		{0x0c, true, false, false, false, 0x12, false, false, false, false},
		{0x0c, true, false, false, true, 0x72, false, false, false, true},
		{0x0c, true, false, true, false, 0x12, false, false, false, false},
		{0x0c, true, false, true, true, 0x72, false, false, false, true},
		{0x0c, true, true, false, false, 0x0c, false, true, false, false},
		{0x0c, true, true, false, true, 0xac, false, true, false, true},
		{0x0c, true, true, true, false, 0x06, false, true, false, false},
		{0x0c, true, true, true, true, 0xa6, false, true, false, true},
		{0x0d, false, false, false, false, 0x13, false, false, false, false},
		{0x0d, false, false, false, true, 0x73, false, false, false, true},
		{0x0d, false, false, true, false, 0x13, false, false, false, false},
		{0x0d, false, false, true, true, 0x73, false, false, false, true},
		{0x0d, false, true, false, false, 0x0d, false, true, false, false},
		{0x0d, false, true, false, true, 0xad, false, true, false, true},
		{0x0d, false, true, true, false, 0x07, false, true, false, false},
		{0x0d, false, true, true, true, 0xa7, false, true, false, true},
		{0x0d, true, false, false, false, 0x13, false, false, false, false},
		{0x0d, true, false, false, true, 0x73, false, false, false, true},
		{0x0d, true, false, true, false, 0x13, false, false, false, false},
		{0x0d, true, false, true, true, 0x73, false, false, false, true},
		{0x0d, true, true, false, false, 0x0d, false, true, false, false},
		{0x0d, true, true, false, true, 0xad, false, true, false, true},
		{0x0d, true, true, true, false, 0x07, false, true, false, false},
		{0x0d, true, true, true, true, 0xa7, false, true, false, true},
		{0x0e, false, false, false, false, 0x14, false, false, false, false},
		{0x0e, false, false, false, true, 0x74, false, false, false, true},
		{0x0e, false, false, true, false, 0x14, false, false, false, false},
		{0x0e, false, false, true, true, 0x74, false, false, false, true},
		{0x0e, false, true, false, false, 0x0e, false, true, false, false},
		{0x0e, false, true, false, true, 0xae, false, true, false, true},
		{0x0e, false, true, true, false, 0x08, false, true, false, false},
		{0x0e, false, true, true, true, 0xa8, false, true, false, true},
		{0x0e, true, false, false, false, 0x14, false, false, false, false},
		{0x0e, true, false, false, true, 0x74, false, false, false, true},
		{0x0e, true, false, true, false, 0x14, false, false, false, false},
		{0x0e, true, false, true, true, 0x74, false, false, false, true},
		{0x0e, true, true, false, false, 0x0e, false, true, false, false},
		{0x0e, true, true, false, true, 0xae, false, true, false, true},
		{0x0e, true, true, true, false, 0x08, false, true, false, false},
		{0x0e, true, true, true, true, 0xa8, false, true, false, true},
		{0x0f, false, false, false, false, 0x15, false, false, false, false},
		{0x0f, false, false, false, true, 0x75, false, false, false, true},
		{0x0f, false, false, true, false, 0x15, false, false, false, false},
		{0x0f, false, false, true, true, 0x75, false, false, false, true},
		{0x0f, false, true, false, false, 0x0f, false, true, false, false},
		{0x0f, false, true, false, true, 0xaf, false, true, false, true},
		{0x0f, false, true, true, false, 0x09, false, true, false, false},
		{0x0f, false, true, true, true, 0xa9, false, true, false, true},
		{0x0f, true, false, false, false, 0x15, false, false, false, false},
		{0x0f, true, false, false, true, 0x75, false, false, false, true},
		{0x0f, true, false, true, false, 0x15, false, false, false, false},
		{0x0f, true, false, true, true, 0x75, false, false, false, true},
		{0x0f, true, true, false, false, 0x0f, false, true, false, false},
		{0x0f, true, true, false, true, 0xaf, false, true, false, true},
		{0x0f, true, true, true, false, 0x09, false, true, false, false},
		{0x0f, true, true, true, true, 0xa9, false, true, false, true},
		{0x10, false, false, false, false, 0x10, false, false, false, false},
		{0x10, false, false, false, true, 0x70, false, false, false, true},
		{0x10, false, false, true, false, 0x16, false, false, false, false},
		{0x10, false, false, true, true, 0x76, false, false, false, true},
		{0x10, false, true, false, false, 0x10, false, true, false, false},
		{0x10, false, true, false, true, 0xb0, false, true, false, true},
		{0x10, false, true, true, false, 0x0a, false, true, false, false},
		{0x10, false, true, true, true, 0xaa, false, true, false, true},
		{0x10, true, false, false, false, 0x10, false, false, false, false},
		{0x10, true, false, false, true, 0x70, false, false, false, true},
		{0x10, true, false, true, false, 0x16, false, false, false, false},
		{0x10, true, false, true, true, 0x76, false, false, false, true},
		{0x10, true, true, false, false, 0x10, false, true, false, false},
		{0x10, true, true, false, true, 0xb0, false, true, false, true},
		{0x10, true, true, true, false, 0x0a, false, true, false, false},
		{0x10, true, true, true, true, 0xaa, false, true, false, true},
		{0x11, false, false, false, false, 0x11, false, false, false, false},
		{0x11, false, false, false, true, 0x71, false, false, false, true},
		{0x11, false, false, true, false, 0x17, false, false, false, false},
		{0x11, false, false, true, true, 0x77, false, false, false, true},
		{0x11, false, true, false, false, 0x11, false, true, false, false},
		{0x11, false, true, false, true, 0xb1, false, true, false, true},
		{0x11, false, true, true, false, 0x0b, false, true, false, false},
		{0x11, false, true, true, true, 0xab, false, true, false, true},
		{0x11, true, false, false, false, 0x11, false, false, false, false},
		{0x11, true, false, false, true, 0x71, false, false, false, true},
		{0x11, true, false, true, false, 0x17, false, false, false, false},
		{0x11, true, false, true, true, 0x77, false, false, false, true},
		{0x11, true, true, false, false, 0x11, false, true, false, false},
		{0x11, true, true, false, true, 0xb1, false, true, false, true},
		{0x11, true, true, true, false, 0x0b, false, true, false, false},
		{0x11, true, true, true, true, 0xab, false, true, false, true},
		{0x12, false, false, false, false, 0x12, false, false, false, false},
		{0x12, false, false, false, true, 0x72, false, false, false, true},
		{0x12, false, false, true, false, 0x18, false, false, false, false},
		{0x12, false, false, true, true, 0x78, false, false, false, true},
		{0x12, false, true, false, false, 0x12, false, true, false, false},
		{0x12, false, true, false, true, 0xb2, false, true, false, true},
		{0x12, false, true, true, false, 0x0c, false, true, false, false},
		{0x12, false, true, true, true, 0xac, false, true, false, true},
		{0x12, true, false, false, false, 0x12, false, false, false, false},
		{0x12, true, false, false, true, 0x72, false, false, false, true},
		{0x12, true, false, true, false, 0x18, false, false, false, false},
		{0x12, true, false, true, true, 0x78, false, false, false, true},
		{0x12, true, true, false, false, 0x12, false, true, false, false},
		{0x12, true, true, false, true, 0xb2, false, true, false, true},
		{0x12, true, true, true, false, 0x0c, false, true, false, false},
		{0x12, true, true, true, true, 0xac, false, true, false, true},
		{0x13, false, false, false, false, 0x13, false, false, false, false},
		{0x13, false, false, false, true, 0x73, false, false, false, true},
		{0x13, false, false, true, false, 0x19, false, false, false, false},
		{0x13, false, false, true, true, 0x79, false, false, false, true},
		{0x13, false, true, false, false, 0x13, false, true, false, false},
		{0x13, false, true, false, true, 0xb3, false, true, false, true},
		{0x13, false, true, true, false, 0x0d, false, true, false, false},
		{0x13, false, true, true, true, 0xad, false, true, false, true},
		{0x13, true, false, false, false, 0x13, false, false, false, false},
		{0x13, true, false, false, true, 0x73, false, false, false, true},
		{0x13, true, false, true, false, 0x19, false, false, false, false},
		{0x13, true, false, true, true, 0x79, false, false, false, true},
		{0x13, true, true, false, false, 0x13, false, true, false, false},
		{0x13, true, true, false, true, 0xb3, false, true, false, true},
		{0x13, true, true, true, false, 0x0d, false, true, false, false},
		{0x13, true, true, true, true, 0xad, false, true, false, true},
		{0x14, false, false, false, false, 0x14, false, false, false, false},
		{0x14, false, false, false, true, 0x74, false, false, false, true},
		{0x14, false, false, true, false, 0x1a, false, false, false, false},
		{0x14, false, false, true, true, 0x7a, false, false, false, true},
		{0x14, false, true, false, false, 0x14, false, true, false, false},
		{0x14, false, true, false, true, 0xb4, false, true, false, true},
		{0x14, false, true, true, false, 0x0e, false, true, false, false},
		{0x14, false, true, true, true, 0xae, false, true, false, true},
		{0x14, true, false, false, false, 0x14, false, false, false, false},
		{0x14, true, false, false, true, 0x74, false, false, false, true},
		{0x14, true, false, true, false, 0x1a, false, false, false, false},
		{0x14, true, false, true, true, 0x7a, false, false, false, true},
		{0x14, true, true, false, false, 0x14, false, true, false, false},
		{0x14, true, true, false, true, 0xb4, false, true, false, true},
		{0x14, true, true, true, false, 0x0e, false, true, false, false},
		{0x14, true, true, true, true, 0xae, false, true, false, true},
		{0x15, false, false, false, false, 0x15, false, false, false, false},
		{0x15, false, false, false, true, 0x75, false, false, false, true},
		{0x15, false, false, true, false, 0x1b, false, false, false, false},
		{0x15, false, false, true, true, 0x7b, false, false, false, true},
		{0x15, false, true, false, false, 0x15, false, true, false, false},
		{0x15, false, true, false, true, 0xb5, false, true, false, true},
		{0x15, false, true, true, false, 0x0f, false, true, false, false},
		{0x15, false, true, true, true, 0xaf, false, true, false, true},
		{0x15, true, false, false, false, 0x15, false, false, false, false},
		{0x15, true, false, false, true, 0x75, false, false, false, true},
		{0x15, true, false, true, false, 0x1b, false, false, false, false},
		{0x15, true, false, true, true, 0x7b, false, false, false, true},
		{0x15, true, true, false, false, 0x15, false, true, false, false},
		{0x15, true, true, false, true, 0xb5, false, true, false, true},
		{0x15, true, true, true, false, 0x0f, false, true, false, false},
		{0x15, true, true, true, true, 0xaf, false, true, false, true},
		{0x16, false, false, false, false, 0x16, false, false, false, false},
		{0x16, false, false, false, true, 0x76, false, false, false, true},
		{0x16, false, false, true, false, 0x1c, false, false, false, false},
		{0x16, false, false, true, true, 0x7c, false, false, false, true},
		{0x16, false, true, false, false, 0x16, false, true, false, false},
		{0x16, false, true, false, true, 0xb6, false, true, false, true},
		{0x16, false, true, true, false, 0x10, false, true, false, false},
		{0x16, false, true, true, true, 0xb0, false, true, false, true},
		{0x16, true, false, false, false, 0x16, false, false, false, false},
		{0x16, true, false, false, true, 0x76, false, false, false, true},
		{0x16, true, false, true, false, 0x1c, false, false, false, false},
		{0x16, true, false, true, true, 0x7c, false, false, false, true},
		{0x16, true, true, false, false, 0x16, false, true, false, false},
		{0x16, true, true, false, true, 0xb6, false, true, false, true},
		{0x16, true, true, true, false, 0x10, false, true, false, false},
		{0x16, true, true, true, true, 0xb0, false, true, false, true},
		{0x17, false, false, false, false, 0x17, false, false, false, false},
		{0x17, false, false, false, true, 0x77, false, false, false, true},
		{0x17, false, false, true, false, 0x1d, false, false, false, false},
		{0x17, false, false, true, true, 0x7d, false, false, false, true},
		{0x17, false, true, false, false, 0x17, false, true, false, false},
		{0x17, false, true, false, true, 0xb7, false, true, false, true},
		{0x17, false, true, true, false, 0x11, false, true, false, false},
		{0x17, false, true, true, true, 0xb1, false, true, false, true},
		{0x17, true, false, false, false, 0x17, false, false, false, false},
		{0x17, true, false, false, true, 0x77, false, false, false, true},
		{0x17, true, false, true, false, 0x1d, false, false, false, false},
		{0x17, true, false, true, true, 0x7d, false, false, false, true},
		{0x17, true, true, false, false, 0x17, false, true, false, false},
		{0x17, true, true, false, true, 0xb7, false, true, false, true},
		{0x17, true, true, true, false, 0x11, false, true, false, false},
		{0x17, true, true, true, true, 0xb1, false, true, false, true},
		{0x18, false, false, false, false, 0x18, false, false, false, false},
		{0x18, false, false, false, true, 0x78, false, false, false, true},
		{0x18, false, false, true, false, 0x1e, false, false, false, false},
		{0x18, false, false, true, true, 0x7e, false, false, false, true},
		{0x18, false, true, false, false, 0x18, false, true, false, false},
		{0x18, false, true, false, true, 0xb8, false, true, false, true},
		{0x18, false, true, true, false, 0x12, false, true, false, false},
		{0x18, false, true, true, true, 0xb2, false, true, false, true},
		{0x18, true, false, false, false, 0x18, false, false, false, false},
		{0x18, true, false, false, true, 0x78, false, false, false, true},
		{0x18, true, false, true, false, 0x1e, false, false, false, false},
		{0x18, true, false, true, true, 0x7e, false, false, false, true},
		{0x18, true, true, false, false, 0x18, false, true, false, false},
		{0x18, true, true, false, true, 0xb8, false, true, false, true},
		{0x18, true, true, true, false, 0x12, false, true, false, false},
		{0x18, true, true, true, true, 0xb2, false, true, false, true},
		{0x19, false, false, false, false, 0x19, false, false, false, false},
		{0x19, false, false, false, true, 0x79, false, false, false, true},
		{0x19, false, false, true, false, 0x1f, false, false, false, false},
		{0x19, false, false, true, true, 0x7f, false, false, false, true},
		{0x19, false, true, false, false, 0x19, false, true, false, false},
		{0x19, false, true, false, true, 0xb9, false, true, false, true},
		{0x19, false, true, true, false, 0x13, false, true, false, false},
		{0x19, false, true, true, true, 0xb3, false, true, false, true},
		{0x19, true, false, false, false, 0x19, false, false, false, false},
		{0x19, true, false, false, true, 0x79, false, false, false, true},
		{0x19, true, false, true, false, 0x1f, false, false, false, false},
		{0x19, true, false, true, true, 0x7f, false, false, false, true},
		{0x19, true, true, false, false, 0x19, false, true, false, false},
		{0x19, true, true, false, true, 0xb9, false, true, false, true},
		{0x19, true, true, true, false, 0x13, false, true, false, false},
		{0x19, true, true, true, true, 0xb3, false, true, false, true},
		{0x1a, false, false, false, false, 0x20, false, false, false, false},
		{0x1a, false, false, false, true, 0x80, false, false, false, true},
		{0x1a, false, false, true, false, 0x20, false, false, false, false},
		{0x1a, false, false, true, true, 0x80, false, false, false, true},
		{0x1a, false, true, false, false, 0x1a, false, true, false, false},
		{0x1a, false, true, false, true, 0xba, false, true, false, true},
		{0x1a, false, true, true, false, 0x14, false, true, false, false},
		{0x1a, false, true, true, true, 0xb4, false, true, false, true},
		{0x1a, true, false, false, false, 0x20, false, false, false, false},
		{0x1a, true, false, false, true, 0x80, false, false, false, true},
		{0x1a, true, false, true, false, 0x20, false, false, false, false},
		{0x1a, true, false, true, true, 0x80, false, false, false, true},
		{0x1a, true, true, false, false, 0x1a, false, true, false, false},
		{0x1a, true, true, false, true, 0xba, false, true, false, true},
		{0x1a, true, true, true, false, 0x14, false, true, false, false},
		{0x1a, true, true, true, true, 0xb4, false, true, false, true},
		{0x1b, false, false, false, false, 0x21, false, false, false, false},
		{0x1b, false, false, false, true, 0x81, false, false, false, true},
		{0x1b, false, false, true, false, 0x21, false, false, false, false},
		{0x1b, false, false, true, true, 0x81, false, false, false, true},
		{0x1b, false, true, false, false, 0x1b, false, true, false, false},
		{0x1b, false, true, false, true, 0xbb, false, true, false, true},
		{0x1b, false, true, true, false, 0x15, false, true, false, false},
		{0x1b, false, true, true, true, 0xb5, false, true, false, true},
		{0x1b, true, false, false, false, 0x21, false, false, false, false},
		{0x1b, true, false, false, true, 0x81, false, false, false, true},
		{0x1b, true, false, true, false, 0x21, false, false, false, false},
		{0x1b, true, false, true, true, 0x81, false, false, false, true},
		{0x1b, true, true, false, false, 0x1b, false, true, false, false},
		{0x1b, true, true, false, true, 0xbb, false, true, false, true},
		{0x1b, true, true, true, false, 0x15, false, true, false, false},
		{0x1b, true, true, true, true, 0xb5, false, true, false, true},
		{0x1c, false, false, false, false, 0x22, false, false, false, false},
		{0x1c, false, false, false, true, 0x82, false, false, false, true},
		{0x1c, false, false, true, false, 0x22, false, false, false, false},
		{0x1c, false, false, true, true, 0x82, false, false, false, true},
		{0x1c, false, true, false, false, 0x1c, false, true, false, false},
		{0x1c, false, true, false, true, 0xbc, false, true, false, true},
		{0x1c, false, true, true, false, 0x16, false, true, false, false},
		{0x1c, false, true, true, true, 0xb6, false, true, false, true},
		{0x1c, true, false, false, false, 0x22, false, false, false, false},
		{0x1c, true, false, false, true, 0x82, false, false, false, true},
		{0x1c, true, false, true, false, 0x22, false, false, false, false},
		{0x1c, true, false, true, true, 0x82, false, false, false, true},
		{0x1c, true, true, false, false, 0x1c, false, true, false, false},
		{0x1c, true, true, false, true, 0xbc, false, true, false, true},
		{0x1c, true, true, true, false, 0x16, false, true, false, false},
		{0x1c, true, true, true, true, 0xb6, false, true, false, true},
		{0x1d, false, false, false, false, 0x23, false, false, false, false},
		{0x1d, false, false, false, true, 0x83, false, false, false, true},
		{0x1d, false, false, true, false, 0x23, false, false, false, false},
		{0x1d, false, false, true, true, 0x83, false, false, false, true},
		{0x1d, false, true, false, false, 0x1d, false, true, false, false},
		{0x1d, false, true, false, true, 0xbd, false, true, false, true},
		{0x1d, false, true, true, false, 0x17, false, true, false, false},
		{0x1d, false, true, true, true, 0xb7, false, true, false, true},
		{0x1d, true, false, false, false, 0x23, false, false, false, false},
		{0x1d, true, false, false, true, 0x83, false, false, false, true},
		{0x1d, true, false, true, false, 0x23, false, false, false, false},
		{0x1d, true, false, true, true, 0x83, false, false, false, true},
		{0x1d, true, true, false, false, 0x1d, false, true, false, false},
		{0x1d, true, true, false, true, 0xbd, false, true, false, true},
		{0x1d, true, true, true, false, 0x17, false, true, false, false},
		{0x1d, true, true, true, true, 0xb7, false, true, false, true},
		{0x1e, false, false, false, false, 0x24, false, false, false, false},
		{0x1e, false, false, false, true, 0x84, false, false, false, true},
		{0x1e, false, false, true, false, 0x24, false, false, false, false},
		{0x1e, false, false, true, true, 0x84, false, false, false, true},
		{0x1e, false, true, false, false, 0x1e, false, true, false, false},
		{0x1e, false, true, false, true, 0xbe, false, true, false, true},
		{0x1e, false, true, true, false, 0x18, false, true, false, false},
		{0x1e, false, true, true, true, 0xb8, false, true, false, true},
		{0x1e, true, false, false, false, 0x24, false, false, false, false},
		{0x1e, true, false, false, true, 0x84, false, false, false, true},
		{0x1e, true, false, true, false, 0x24, false, false, false, false},
		{0x1e, true, false, true, true, 0x84, false, false, false, true},
		{0x1e, true, true, false, false, 0x1e, false, true, false, false},
		{0x1e, true, true, false, true, 0xbe, false, true, false, true},
		{0x1e, true, true, true, false, 0x18, false, true, false, false},
		{0x1e, true, true, true, true, 0xb8, false, true, false, true},
		{0x1f, false, false, false, false, 0x25, false, false, false, false},
		{0x1f, false, false, false, true, 0x85, false, false, false, true},
		{0x1f, false, false, true, false, 0x25, false, false, false, false},
		{0x1f, false, false, true, true, 0x85, false, false, false, true},
		{0x1f, false, true, false, false, 0x1f, false, true, false, false},
		{0x1f, false, true, false, true, 0xbf, false, true, false, true},
		{0x1f, false, true, true, false, 0x19, false, true, false, false},
		{0x1f, false, true, true, true, 0xb9, false, true, false, true},
		{0x1f, true, false, false, false, 0x25, false, false, false, false},
		{0x1f, true, false, false, true, 0x85, false, false, false, true},
		{0x1f, true, false, true, false, 0x25, false, false, false, false},
		{0x1f, true, false, true, true, 0x85, false, false, false, true},
		{0x1f, true, true, false, false, 0x1f, false, true, false, false},
		{0x1f, true, true, false, true, 0xbf, false, true, false, true},
		{0x1f, true, true, true, false, 0x19, false, true, false, false},
		{0x1f, true, true, true, true, 0xb9, false, true, false, true},
		{0x20, false, false, false, false, 0x20, false, false, false, false},
		{0x20, false, false, false, true, 0x80, false, false, false, true},
		{0x20, false, false, true, false, 0x26, false, false, false, false},
		{0x20, false, false, true, true, 0x86, false, false, false, true},
		{0x20, false, true, false, false, 0x20, false, true, false, false},
		{0x20, false, true, false, true, 0xc0, false, true, false, true},
		{0x20, false, true, true, false, 0x1a, false, true, false, false},
		{0x20, false, true, true, true, 0xba, false, true, false, true},
		{0x20, true, false, false, false, 0x20, false, false, false, false},
		{0x20, true, false, false, true, 0x80, false, false, false, true},
		{0x20, true, false, true, false, 0x26, false, false, false, false},
		{0x20, true, false, true, true, 0x86, false, false, false, true},
		{0x20, true, true, false, false, 0x20, false, true, false, false},
		{0x20, true, true, false, true, 0xc0, false, true, false, true},
		{0x20, true, true, true, false, 0x1a, false, true, false, false},
		{0x20, true, true, true, true, 0xba, false, true, false, true},
		{0x21, false, false, false, false, 0x21, false, false, false, false},
		{0x21, false, false, false, true, 0x81, false, false, false, true},
		{0x21, false, false, true, false, 0x27, false, false, false, false},
		{0x21, false, false, true, true, 0x87, false, false, false, true},
		{0x21, false, true, false, false, 0x21, false, true, false, false},
		{0x21, false, true, false, true, 0xc1, false, true, false, true},
		{0x21, false, true, true, false, 0x1b, false, true, false, false},
		{0x21, false, true, true, true, 0xbb, false, true, false, true},
		{0x21, true, false, false, false, 0x21, false, false, false, false},
		{0x21, true, false, false, true, 0x81, false, false, false, true},
		{0x21, true, false, true, false, 0x27, false, false, false, false},
		{0x21, true, false, true, true, 0x87, false, false, false, true},
		{0x21, true, true, false, false, 0x21, false, true, false, false},
		{0x21, true, true, false, true, 0xc1, false, true, false, true},
		{0x21, true, true, true, false, 0x1b, false, true, false, false},
		{0x21, true, true, true, true, 0xbb, false, true, false, true},
		{0x22, false, false, false, false, 0x22, false, false, false, false},
		{0x22, false, false, false, true, 0x82, false, false, false, true},
		{0x22, false, false, true, false, 0x28, false, false, false, false},
		{0x22, false, false, true, true, 0x88, false, false, false, true},
		{0x22, false, true, false, false, 0x22, false, true, false, false},
		{0x22, false, true, false, true, 0xc2, false, true, false, true},
		{0x22, false, true, true, false, 0x1c, false, true, false, false},
		{0x22, false, true, true, true, 0xbc, false, true, false, true},
		{0x22, true, false, false, false, 0x22, false, false, false, false},
		{0x22, true, false, false, true, 0x82, false, false, false, true},
		{0x22, true, false, true, false, 0x28, false, false, false, false},
		{0x22, true, false, true, true, 0x88, false, false, false, true},
		{0x22, true, true, false, false, 0x22, false, true, false, false},
		{0x22, true, true, false, true, 0xc2, false, true, false, true},
		{0x22, true, true, true, false, 0x1c, false, true, false, false},
		{0x22, true, true, true, true, 0xbc, false, true, false, true},
		{0x23, false, false, false, false, 0x23, false, false, false, false},
		{0x23, false, false, false, true, 0x83, false, false, false, true},
		{0x23, false, false, true, false, 0x29, false, false, false, false},
		{0x23, false, false, true, true, 0x89, false, false, false, true},
		{0x23, false, true, false, false, 0x23, false, true, false, false},
		{0x23, false, true, false, true, 0xc3, false, true, false, true},
		{0x23, false, true, true, false, 0x1d, false, true, false, false},
		{0x23, false, true, true, true, 0xbd, false, true, false, true},
		{0x23, true, false, false, false, 0x23, false, false, false, false},
		{0x23, true, false, false, true, 0x83, false, false, false, true},
		{0x23, true, false, true, false, 0x29, false, false, false, false},
		{0x23, true, false, true, true, 0x89, false, false, false, true},
		{0x23, true, true, false, false, 0x23, false, true, false, false},
		{0x23, true, true, false, true, 0xc3, false, true, false, true},
		{0x23, true, true, true, false, 0x1d, false, true, false, false},
		{0x23, true, true, true, true, 0xbd, false, true, false, true},
		{0x24, false, false, false, false, 0x24, false, false, false, false},
		{0x24, false, false, false, true, 0x84, false, false, false, true},
		{0x24, false, false, true, false, 0x2a, false, false, false, false},
		{0x24, false, false, true, true, 0x8a, false, false, false, true},
		{0x24, false, true, false, false, 0x24, false, true, false, false},
		{0x24, false, true, false, true, 0xc4, false, true, false, true},
		{0x24, false, true, true, false, 0x1e, false, true, false, false},
		{0x24, false, true, true, true, 0xbe, false, true, false, true},
		{0x24, true, false, false, false, 0x24, false, false, false, false},
		{0x24, true, false, false, true, 0x84, false, false, false, true},
		{0x24, true, false, true, false, 0x2a, false, false, false, false},
		{0x24, true, false, true, true, 0x8a, false, false, false, true},
		{0x24, true, true, false, false, 0x24, false, true, false, false},
		{0x24, true, true, false, true, 0xc4, false, true, false, true},
		{0x24, true, true, true, false, 0x1e, false, true, false, false},
		{0x24, true, true, true, true, 0xbe, false, true, false, true},
		{0x25, false, false, false, false, 0x25, false, false, false, false},
		{0x25, false, false, false, true, 0x85, false, false, false, true},
		{0x25, false, false, true, false, 0x2b, false, false, false, false},
		{0x25, false, false, true, true, 0x8b, false, false, false, true},
		{0x25, false, true, false, false, 0x25, false, true, false, false},
		{0x25, false, true, false, true, 0xc5, false, true, false, true},
		{0x25, false, true, true, false, 0x1f, false, true, false, false},
		{0x25, false, true, true, true, 0xbf, false, true, false, true},
		{0x25, true, false, false, false, 0x25, false, false, false, false},
		{0x25, true, false, false, true, 0x85, false, false, false, true},
		{0x25, true, false, true, false, 0x2b, false, false, false, false},
		{0x25, true, false, true, true, 0x8b, false, false, false, true},
		{0x25, true, true, false, false, 0x25, false, true, false, false},
		{0x25, true, true, false, true, 0xc5, false, true, false, true},
		{0x25, true, true, true, false, 0x1f, false, true, false, false},
		{0x25, true, true, true, true, 0xbf, false, true, false, true},
		{0x26, false, false, false, false, 0x26, false, false, false, false},
		{0x26, false, false, false, true, 0x86, false, false, false, true},
		{0x26, false, false, true, false, 0x2c, false, false, false, false},
		{0x26, false, false, true, true, 0x8c, false, false, false, true},
		{0x26, false, true, false, false, 0x26, false, true, false, false},
		{0x26, false, true, false, true, 0xc6, false, true, false, true},
		{0x26, false, true, true, false, 0x20, false, true, false, false},
		{0x26, false, true, true, true, 0xc0, false, true, false, true},
		{0x26, true, false, false, false, 0x26, false, false, false, false},
		{0x26, true, false, false, true, 0x86, false, false, false, true},
		{0x26, true, false, true, false, 0x2c, false, false, false, false},
		{0x26, true, false, true, true, 0x8c, false, false, false, true},
		{0x26, true, true, false, false, 0x26, false, true, false, false},
		{0x26, true, true, false, true, 0xc6, false, true, false, true},
		{0x26, true, true, true, false, 0x20, false, true, false, false},
		{0x26, true, true, true, true, 0xc0, false, true, false, true},
		{0x27, false, false, false, false, 0x27, false, false, false, false},
		{0x27, false, false, false, true, 0x87, false, false, false, true},
		{0x27, false, false, true, false, 0x2d, false, false, false, false},
		{0x27, false, false, true, true, 0x8d, false, false, false, true},
		{0x27, false, true, false, false, 0x27, false, true, false, false},
		{0x27, false, true, false, true, 0xc7, false, true, false, true},
		{0x27, false, true, true, false, 0x21, false, true, false, false},
		{0x27, false, true, true, true, 0xc1, false, true, false, true},
		{0x27, true, false, false, false, 0x27, false, false, false, false},
		{0x27, true, false, false, true, 0x87, false, false, false, true},
		{0x27, true, false, true, false, 0x2d, false, false, false, false},
		{0x27, true, false, true, true, 0x8d, false, false, false, true},
		{0x27, true, true, false, false, 0x27, false, true, false, false},
		{0x27, true, true, false, true, 0xc7, false, true, false, true},
		{0x27, true, true, true, false, 0x21, false, true, false, false},
		{0x27, true, true, true, true, 0xc1, false, true, false, true},
		{0x28, false, false, false, false, 0x28, false, false, false, false},
		{0x28, false, false, false, true, 0x88, false, false, false, true},
		{0x28, false, false, true, false, 0x2e, false, false, false, false},
		{0x28, false, false, true, true, 0x8e, false, false, false, true},
		{0x28, false, true, false, false, 0x28, false, true, false, false},
		{0x28, false, true, false, true, 0xc8, false, true, false, true},
		{0x28, false, true, true, false, 0x22, false, true, false, false},
		{0x28, false, true, true, true, 0xc2, false, true, false, true},
		{0x28, true, false, false, false, 0x28, false, false, false, false},
		{0x28, true, false, false, true, 0x88, false, false, false, true},
		{0x28, true, false, true, false, 0x2e, false, false, false, false},
		{0x28, true, false, true, true, 0x8e, false, false, false, true},
		{0x28, true, true, false, false, 0x28, false, true, false, false},
		{0x28, true, true, false, true, 0xc8, false, true, false, true},
		{0x28, true, true, true, false, 0x22, false, true, false, false},
		{0x28, true, true, true, true, 0xc2, false, true, false, true},
		{0x29, false, false, false, false, 0x29, false, false, false, false},
		{0x29, false, false, false, true, 0x89, false, false, false, true},
		{0x29, false, false, true, false, 0x2f, false, false, false, false},
		{0x29, false, false, true, true, 0x8f, false, false, false, true},
		{0x29, false, true, false, false, 0x29, false, true, false, false},
		{0x29, false, true, false, true, 0xc9, false, true, false, true},
		{0x29, false, true, true, false, 0x23, false, true, false, false},
		{0x29, false, true, true, true, 0xc3, false, true, false, true},
		{0x29, true, false, false, false, 0x29, false, false, false, false},
		{0x29, true, false, false, true, 0x89, false, false, false, true},
		{0x29, true, false, true, false, 0x2f, false, false, false, false},
		{0x29, true, false, true, true, 0x8f, false, false, false, true},
		{0x29, true, true, false, false, 0x29, false, true, false, false},
		{0x29, true, true, false, true, 0xc9, false, true, false, true},
		{0x29, true, true, true, false, 0x23, false, true, false, false},
		{0x29, true, true, true, true, 0xc3, false, true, false, true},
		{0x2a, false, false, false, false, 0x30, false, false, false, false},
		{0x2a, false, false, false, true, 0x90, false, false, false, true},
		{0x2a, false, false, true, false, 0x30, false, false, false, false},
		{0x2a, false, false, true, true, 0x90, false, false, false, true},
		{0x2a, false, true, false, false, 0x2a, false, true, false, false},
		{0x2a, false, true, false, true, 0xca, false, true, false, true},
		{0x2a, false, true, true, false, 0x24, false, true, false, false},
		{0x2a, false, true, true, true, 0xc4, false, true, false, true},
		{0x2a, true, false, false, false, 0x30, false, false, false, false},
		{0x2a, true, false, false, true, 0x90, false, false, false, true},
		{0x2a, true, false, true, false, 0x30, false, false, false, false},
		{0x2a, true, false, true, true, 0x90, false, false, false, true},
		{0x2a, true, true, false, false, 0x2a, false, true, false, false},
		{0x2a, true, true, false, true, 0xca, false, true, false, true},
		{0x2a, true, true, true, false, 0x24, false, true, false, false},
		{0x2a, true, true, true, true, 0xc4, false, true, false, true},
		{0x2b, false, false, false, false, 0x31, false, false, false, false},
		{0x2b, false, false, false, true, 0x91, false, false, false, true},
		{0x2b, false, false, true, false, 0x31, false, false, false, false},
		{0x2b, false, false, true, true, 0x91, false, false, false, true},
		{0x2b, false, true, false, false, 0x2b, false, true, false, false},
		{0x2b, false, true, false, true, 0xcb, false, true, false, true},
		{0x2b, false, true, true, false, 0x25, false, true, false, false},
		{0x2b, false, true, true, true, 0xc5, false, true, false, true},
		{0x2b, true, false, false, false, 0x31, false, false, false, false},
		{0x2b, true, false, false, true, 0x91, false, false, false, true},
		{0x2b, true, false, true, false, 0x31, false, false, false, false},
		{0x2b, true, false, true, true, 0x91, false, false, false, true},
		{0x2b, true, true, false, false, 0x2b, false, true, false, false},
		{0x2b, true, true, false, true, 0xcb, false, true, false, true},
		{0x2b, true, true, true, false, 0x25, false, true, false, false},
		{0x2b, true, true, true, true, 0xc5, false, true, false, true},
		{0x2c, false, false, false, false, 0x32, false, false, false, false},
		{0x2c, false, false, false, true, 0x92, false, false, false, true},
		{0x2c, false, false, true, false, 0x32, false, false, false, false},
		{0x2c, false, false, true, true, 0x92, false, false, false, true},
		{0x2c, false, true, false, false, 0x2c, false, true, false, false},
		{0x2c, false, true, false, true, 0xcc, false, true, false, true},
		{0x2c, false, true, true, false, 0x26, false, true, false, false},
		{0x2c, false, true, true, true, 0xc6, false, true, false, true},
		{0x2c, true, false, false, false, 0x32, false, false, false, false},
		{0x2c, true, false, false, true, 0x92, false, false, false, true},
		{0x2c, true, false, true, false, 0x32, false, false, false, false},
		{0x2c, true, false, true, true, 0x92, false, false, false, true},
		{0x2c, true, true, false, false, 0x2c, false, true, false, false},
		{0x2c, true, true, false, true, 0xcc, false, true, false, true},
		{0x2c, true, true, true, false, 0x26, false, true, false, false},
		{0x2c, true, true, true, true, 0xc6, false, true, false, true},
		{0x2d, false, false, false, false, 0x33, false, false, false, false},
		{0x2d, false, false, false, true, 0x93, false, false, false, true},
		{0x2d, false, false, true, false, 0x33, false, false, false, false},
		{0x2d, false, false, true, true, 0x93, false, false, false, true},
		{0x2d, false, true, false, false, 0x2d, false, true, false, false},
		{0x2d, false, true, false, true, 0xcd, false, true, false, true},
		{0x2d, false, true, true, false, 0x27, false, true, false, false},
		{0x2d, false, true, true, true, 0xc7, false, true, false, true},
		{0x2d, true, false, false, false, 0x33, false, false, false, false},
		{0x2d, true, false, false, true, 0x93, false, false, false, true},
		{0x2d, true, false, true, false, 0x33, false, false, false, false},
		{0x2d, true, false, true, true, 0x93, false, false, false, true},
		{0x2d, true, true, false, false, 0x2d, false, true, false, false},
		{0x2d, true, true, false, true, 0xcd, false, true, false, true},
		{0x2d, true, true, true, false, 0x27, false, true, false, false},
		{0x2d, true, true, true, true, 0xc7, false, true, false, true},
		{0x2e, false, false, false, false, 0x34, false, false, false, false},
		{0x2e, false, false, false, true, 0x94, false, false, false, true},
		{0x2e, false, false, true, false, 0x34, false, false, false, false},
		{0x2e, false, false, true, true, 0x94, false, false, false, true},
		{0x2e, false, true, false, false, 0x2e, false, true, false, false},
		{0x2e, false, true, false, true, 0xce, false, true, false, true},
		{0x2e, false, true, true, false, 0x28, false, true, false, false},
		{0x2e, false, true, true, true, 0xc8, false, true, false, true},
		{0x2e, true, false, false, false, 0x34, false, false, false, false},
		{0x2e, true, false, false, true, 0x94, false, false, false, true},
		{0x2e, true, false, true, false, 0x34, false, false, false, false},
		{0x2e, true, false, true, true, 0x94, false, false, false, true},
		{0x2e, true, true, false, false, 0x2e, false, true, false, false},
		{0x2e, true, true, false, true, 0xce, false, true, false, true},
		{0x2e, true, true, true, false, 0x28, false, true, false, false},
		{0x2e, true, true, true, true, 0xc8, false, true, false, true},
		{0x2f, false, false, false, false, 0x35, false, false, false, false},
		{0x2f, false, false, false, true, 0x95, false, false, false, true},
		{0x2f, false, false, true, false, 0x35, false, false, false, false},
		{0x2f, false, false, true, true, 0x95, false, false, false, true},
		{0x2f, false, true, false, false, 0x2f, false, true, false, false},
		{0x2f, false, true, false, true, 0xcf, false, true, false, true},
		{0x2f, false, true, true, false, 0x29, false, true, false, false},
		{0x2f, false, true, true, true, 0xc9, false, true, false, true},
		{0x2f, true, false, false, false, 0x35, false, false, false, false},
		{0x2f, true, false, false, true, 0x95, false, false, false, true},
		{0x2f, true, false, true, false, 0x35, false, false, false, false},
		{0x2f, true, false, true, true, 0x95, false, false, false, true},
		{0x2f, true, true, false, false, 0x2f, false, true, false, false},
		{0x2f, true, true, false, true, 0xcf, false, true, false, true},
		{0x2f, true, true, true, false, 0x29, false, true, false, false},
		{0x2f, true, true, true, true, 0xc9, false, true, false, true},
		{0x30, false, false, false, false, 0x30, false, false, false, false},
		{0x30, false, false, false, true, 0x90, false, false, false, true},
		{0x30, false, false, true, false, 0x36, false, false, false, false},
		{0x30, false, false, true, true, 0x96, false, false, false, true},
		{0x30, false, true, false, false, 0x30, false, true, false, false},
		{0x30, false, true, false, true, 0xd0, false, true, false, true},
		{0x30, false, true, true, false, 0x2a, false, true, false, false},
		{0x30, false, true, true, true, 0xca, false, true, false, true},
		{0x30, true, false, false, false, 0x30, false, false, false, false},
		{0x30, true, false, false, true, 0x90, false, false, false, true},
		{0x30, true, false, true, false, 0x36, false, false, false, false},
		{0x30, true, false, true, true, 0x96, false, false, false, true},
		{0x30, true, true, false, false, 0x30, false, true, false, false},
		{0x30, true, true, false, true, 0xd0, false, true, false, true},
		{0x30, true, true, true, false, 0x2a, false, true, false, false},
		{0x30, true, true, true, true, 0xca, false, true, false, true},
		{0x31, false, false, false, false, 0x31, false, false, false, false},
		{0x31, false, false, false, true, 0x91, false, false, false, true},
		{0x31, false, false, true, false, 0x37, false, false, false, false},
		{0x31, false, false, true, true, 0x97, false, false, false, true},
		{0x31, false, true, false, false, 0x31, false, true, false, false},
		{0x31, false, true, false, true, 0xd1, false, true, false, true},
		{0x31, false, true, true, false, 0x2b, false, true, false, false},
		{0x31, false, true, true, true, 0xcb, false, true, false, true},
		{0x31, true, false, false, false, 0x31, false, false, false, false},
		{0x31, true, false, false, true, 0x91, false, false, false, true},
		{0x31, true, false, true, false, 0x37, false, false, false, false},
		{0x31, true, false, true, true, 0x97, false, false, false, true},
		{0x31, true, true, false, false, 0x31, false, true, false, false},
		{0x31, true, true, false, true, 0xd1, false, true, false, true},
		{0x31, true, true, true, false, 0x2b, false, true, false, false},
		{0x31, true, true, true, true, 0xcb, false, true, false, true},
		{0x32, false, false, false, false, 0x32, false, false, false, false},
		{0x32, false, false, false, true, 0x92, false, false, false, true},
		{0x32, false, false, true, false, 0x38, false, false, false, false},
		{0x32, false, false, true, true, 0x98, false, false, false, true},
		{0x32, false, true, false, false, 0x32, false, true, false, false},
		{0x32, false, true, false, true, 0xd2, false, true, false, true},
		{0x32, false, true, true, false, 0x2c, false, true, false, false},
		{0x32, false, true, true, true, 0xcc, false, true, false, true},
		{0x32, true, false, false, false, 0x32, false, false, false, false},
		{0x32, true, false, false, true, 0x92, false, false, false, true},
		{0x32, true, false, true, false, 0x38, false, false, false, false},
		{0x32, true, false, true, true, 0x98, false, false, false, true},
		{0x32, true, true, false, false, 0x32, false, true, false, false},
		{0x32, true, true, false, true, 0xd2, false, true, false, true},
		{0x32, true, true, true, false, 0x2c, false, true, false, false},
		{0x32, true, true, true, true, 0xcc, false, true, false, true},
		{0x33, false, false, false, false, 0x33, false, false, false, false},
		{0x33, false, false, false, true, 0x93, false, false, false, true},
		{0x33, false, false, true, false, 0x39, false, false, false, false},
		{0x33, false, false, true, true, 0x99, false, false, false, true},
		{0x33, false, true, false, false, 0x33, false, true, false, false},
		{0x33, false, true, false, true, 0xd3, false, true, false, true},
		{0x33, false, true, true, false, 0x2d, false, true, false, false},
		{0x33, false, true, true, true, 0xcd, false, true, false, true},
		{0x33, true, false, false, false, 0x33, false, false, false, false},
		{0x33, true, false, false, true, 0x93, false, false, false, true},
		{0x33, true, false, true, false, 0x39, false, false, false, false},
		{0x33, true, false, true, true, 0x99, false, false, false, true},
		{0x33, true, true, false, false, 0x33, false, true, false, false},
		{0x33, true, true, false, true, 0xd3, false, true, false, true},
		{0x33, true, true, true, false, 0x2d, false, true, false, false},
		{0x33, true, true, true, true, 0xcd, false, true, false, true},
		{0x34, false, false, false, false, 0x34, false, false, false, false},
		{0x34, false, false, false, true, 0x94, false, false, false, true},
		{0x34, false, false, true, false, 0x3a, false, false, false, false},
		{0x34, false, false, true, true, 0x9a, false, false, false, true},
		{0x34, false, true, false, false, 0x34, false, true, false, false},
		{0x34, false, true, false, true, 0xd4, false, true, false, true},
		{0x34, false, true, true, false, 0x2e, false, true, false, false},
		{0x34, false, true, true, true, 0xce, false, true, false, true},
		{0x34, true, false, false, false, 0x34, false, false, false, false},
		{0x34, true, false, false, true, 0x94, false, false, false, true},
		{0x34, true, false, true, false, 0x3a, false, false, false, false},
		{0x34, true, false, true, true, 0x9a, false, false, false, true},
		{0x34, true, true, false, false, 0x34, false, true, false, false},
		{0x34, true, true, false, true, 0xd4, false, true, false, true},
		{0x34, true, true, true, false, 0x2e, false, true, false, false},
		{0x34, true, true, true, true, 0xce, false, true, false, true},
		{0x35, false, false, false, false, 0x35, false, false, false, false},
		{0x35, false, false, false, true, 0x95, false, false, false, true},
		{0x35, false, false, true, false, 0x3b, false, false, false, false},
		{0x35, false, false, true, true, 0x9b, false, false, false, true},
		{0x35, false, true, false, false, 0x35, false, true, false, false},
		{0x35, false, true, false, true, 0xd5, false, true, false, true},
		{0x35, false, true, true, false, 0x2f, false, true, false, false},
		{0x35, false, true, true, true, 0xcf, false, true, false, true},
		{0x35, true, false, false, false, 0x35, false, false, false, false},
		{0x35, true, false, false, true, 0x95, false, false, false, true},
		{0x35, true, false, true, false, 0x3b, false, false, false, false},
		{0x35, true, false, true, true, 0x9b, false, false, false, true},
		{0x35, true, true, false, false, 0x35, false, true, false, false},
		{0x35, true, true, false, true, 0xd5, false, true, false, true},
		{0x35, true, true, true, false, 0x2f, false, true, false, false},
		{0x35, true, true, true, true, 0xcf, false, true, false, true},
		{0x36, false, false, false, false, 0x36, false, false, false, false},
		{0x36, false, false, false, true, 0x96, false, false, false, true},
		{0x36, false, false, true, false, 0x3c, false, false, false, false},
		{0x36, false, false, true, true, 0x9c, false, false, false, true},
		{0x36, false, true, false, false, 0x36, false, true, false, false},
		{0x36, false, true, false, true, 0xd6, false, true, false, true},
		{0x36, false, true, true, false, 0x30, false, true, false, false},
		{0x36, false, true, true, true, 0xd0, false, true, false, true},
		{0x36, true, false, false, false, 0x36, false, false, false, false},
		{0x36, true, false, false, true, 0x96, false, false, false, true},
		{0x36, true, false, true, false, 0x3c, false, false, false, false},
		{0x36, true, false, true, true, 0x9c, false, false, false, true},
		{0x36, true, true, false, false, 0x36, false, true, false, false},
		{0x36, true, true, false, true, 0xd6, false, true, false, true},
		{0x36, true, true, true, false, 0x30, false, true, false, false},
		{0x36, true, true, true, true, 0xd0, false, true, false, true},
		{0x37, false, false, false, false, 0x37, false, false, false, false},
		{0x37, false, false, false, true, 0x97, false, false, false, true},
		{0x37, false, false, true, false, 0x3d, false, false, false, false},
		{0x37, false, false, true, true, 0x9d, false, false, false, true},
		{0x37, false, true, false, false, 0x37, false, true, false, false},
		{0x37, false, true, false, true, 0xd7, false, true, false, true},
		{0x37, false, true, true, false, 0x31, false, true, false, false},
		{0x37, false, true, true, true, 0xd1, false, true, false, true},
		{0x37, true, false, false, false, 0x37, false, false, false, false},
		{0x37, true, false, false, true, 0x97, false, false, false, true},
		{0x37, true, false, true, false, 0x3d, false, false, false, false},
		{0x37, true, false, true, true, 0x9d, false, false, false, true},
		{0x37, true, true, false, false, 0x37, false, true, false, false},
		{0x37, true, true, false, true, 0xd7, false, true, false, true},
		{0x37, true, true, true, false, 0x31, false, true, false, false},
		{0x37, true, true, true, true, 0xd1, false, true, false, true},
		{0x38, false, false, false, false, 0x38, false, false, false, false},
		{0x38, false, false, false, true, 0x98, false, false, false, true},
		{0x38, false, false, true, false, 0x3e, false, false, false, false},
		{0x38, false, false, true, true, 0x9e, false, false, false, true},
		{0x38, false, true, false, false, 0x38, false, true, false, false},
		{0x38, false, true, false, true, 0xd8, false, true, false, true},
		{0x38, false, true, true, false, 0x32, false, true, false, false},
		{0x38, false, true, true, true, 0xd2, false, true, false, true},
		{0x38, true, false, false, false, 0x38, false, false, false, false},
		{0x38, true, false, false, true, 0x98, false, false, false, true},
		{0x38, true, false, true, false, 0x3e, false, false, false, false},
		{0x38, true, false, true, true, 0x9e, false, false, false, true},
		{0x38, true, true, false, false, 0x38, false, true, false, false},
		{0x38, true, true, false, true, 0xd8, false, true, false, true},
		{0x38, true, true, true, false, 0x32, false, true, false, false},
		{0x38, true, true, true, true, 0xd2, false, true, false, true},
		{0x39, false, false, false, false, 0x39, false, false, false, false},
		{0x39, false, false, false, true, 0x99, false, false, false, true},
		{0x39, false, false, true, false, 0x3f, false, false, false, false},
		{0x39, false, false, true, true, 0x9f, false, false, false, true},
		{0x39, false, true, false, false, 0x39, false, true, false, false},
		{0x39, false, true, false, true, 0xd9, false, true, false, true},
		{0x39, false, true, true, false, 0x33, false, true, false, false},
		{0x39, false, true, true, true, 0xd3, false, true, false, true},
		{0x39, true, false, false, false, 0x39, false, false, false, false},
		{0x39, true, false, false, true, 0x99, false, false, false, true},
		{0x39, true, false, true, false, 0x3f, false, false, false, false},
		{0x39, true, false, true, true, 0x9f, false, false, false, true},
		{0x39, true, true, false, false, 0x39, false, true, false, false},
		{0x39, true, true, false, true, 0xd9, false, true, false, true},
		{0x39, true, true, true, false, 0x33, false, true, false, false},
		{0x39, true, true, true, true, 0xd3, false, true, false, true},
		{0x3a, false, false, false, false, 0x40, false, false, false, false},
		{0x3a, false, false, false, true, 0xa0, false, false, false, true},
		{0x3a, false, false, true, false, 0x40, false, false, false, false},
		{0x3a, false, false, true, true, 0xa0, false, false, false, true},
		{0x3a, false, true, false, false, 0x3a, false, true, false, false},
		{0x3a, false, true, false, true, 0xda, false, true, false, true},
		{0x3a, false, true, true, false, 0x34, false, true, false, false},
		{0x3a, false, true, true, true, 0xd4, false, true, false, true},
		{0x3a, true, false, false, false, 0x40, false, false, false, false},
		{0x3a, true, false, false, true, 0xa0, false, false, false, true},
		{0x3a, true, false, true, false, 0x40, false, false, false, false},
		{0x3a, true, false, true, true, 0xa0, false, false, false, true},
		{0x3a, true, true, false, false, 0x3a, false, true, false, false},
		{0x3a, true, true, false, true, 0xda, false, true, false, true},
		{0x3a, true, true, true, false, 0x34, false, true, false, false},
		{0x3a, true, true, true, true, 0xd4, false, true, false, true},
		{0x3b, false, false, false, false, 0x41, false, false, false, false},
		{0x3b, false, false, false, true, 0xa1, false, false, false, true},
		{0x3b, false, false, true, false, 0x41, false, false, false, false},
		{0x3b, false, false, true, true, 0xa1, false, false, false, true},
		{0x3b, false, true, false, false, 0x3b, false, true, false, false},
		{0x3b, false, true, false, true, 0xdb, false, true, false, true},
		{0x3b, false, true, true, false, 0x35, false, true, false, false},
		{0x3b, false, true, true, true, 0xd5, false, true, false, true},
		{0x3b, true, false, false, false, 0x41, false, false, false, false},
		{0x3b, true, false, false, true, 0xa1, false, false, false, true},
		{0x3b, true, false, true, false, 0x41, false, false, false, false},
		{0x3b, true, false, true, true, 0xa1, false, false, false, true},
		{0x3b, true, true, false, false, 0x3b, false, true, false, false},
		{0x3b, true, true, false, true, 0xdb, false, true, false, true},
		{0x3b, true, true, true, false, 0x35, false, true, false, false},
		{0x3b, true, true, true, true, 0xd5, false, true, false, true},
		{0x3c, false, false, false, false, 0x42, false, false, false, false},
		{0x3c, false, false, false, true, 0xa2, false, false, false, true},
		{0x3c, false, false, true, false, 0x42, false, false, false, false},
		{0x3c, false, false, true, true, 0xa2, false, false, false, true},
		{0x3c, false, true, false, false, 0x3c, false, true, false, false},
		{0x3c, false, true, false, true, 0xdc, false, true, false, true},
		{0x3c, false, true, true, false, 0x36, false, true, false, false},
		{0x3c, false, true, true, true, 0xd6, false, true, false, true},
		{0x3c, true, false, false, false, 0x42, false, false, false, false},
		{0x3c, true, false, false, true, 0xa2, false, false, false, true},
		{0x3c, true, false, true, false, 0x42, false, false, false, false},
		{0x3c, true, false, true, true, 0xa2, false, false, false, true},
		{0x3c, true, true, false, false, 0x3c, false, true, false, false},
		{0x3c, true, true, false, true, 0xdc, false, true, false, true},
		{0x3c, true, true, true, false, 0x36, false, true, false, false},
		{0x3c, true, true, true, true, 0xd6, false, true, false, true},
		{0x3d, false, false, false, false, 0x43, false, false, false, false},
		{0x3d, false, false, false, true, 0xa3, false, false, false, true},
		{0x3d, false, false, true, false, 0x43, false, false, false, false},
		{0x3d, false, false, true, true, 0xa3, false, false, false, true},
		{0x3d, false, true, false, false, 0x3d, false, true, false, false},
		{0x3d, false, true, false, true, 0xdd, false, true, false, true},
		{0x3d, false, true, true, false, 0x37, false, true, false, false},
		{0x3d, false, true, true, true, 0xd7, false, true, false, true},
		{0x3d, true, false, false, false, 0x43, false, false, false, false},
		{0x3d, true, false, false, true, 0xa3, false, false, false, true},
		{0x3d, true, false, true, false, 0x43, false, false, false, false},
		{0x3d, true, false, true, true, 0xa3, false, false, false, true},
		{0x3d, true, true, false, false, 0x3d, false, true, false, false},
		{0x3d, true, true, false, true, 0xdd, false, true, false, true},
		{0x3d, true, true, true, false, 0x37, false, true, false, false},
		{0x3d, true, true, true, true, 0xd7, false, true, false, true},
		{0x3e, false, false, false, false, 0x44, false, false, false, false},
		{0x3e, false, false, false, true, 0xa4, false, false, false, true},
		{0x3e, false, false, true, false, 0x44, false, false, false, false},
		{0x3e, false, false, true, true, 0xa4, false, false, false, true},
		{0x3e, false, true, false, false, 0x3e, false, true, false, false},
		{0x3e, false, true, false, true, 0xde, false, true, false, true},
		{0x3e, false, true, true, false, 0x38, false, true, false, false},
		{0x3e, false, true, true, true, 0xd8, false, true, false, true},
		{0x3e, true, false, false, false, 0x44, false, false, false, false},
		{0x3e, true, false, false, true, 0xa4, false, false, false, true},
		{0x3e, true, false, true, false, 0x44, false, false, false, false},
		{0x3e, true, false, true, true, 0xa4, false, false, false, true},
		{0x3e, true, true, false, false, 0x3e, false, true, false, false},
		{0x3e, true, true, false, true, 0xde, false, true, false, true},
		{0x3e, true, true, true, false, 0x38, false, true, false, false},
		{0x3e, true, true, true, true, 0xd8, false, true, false, true},
		{0x3f, false, false, false, false, 0x45, false, false, false, false},
		{0x3f, false, false, false, true, 0xa5, false, false, false, true},
		{0x3f, false, false, true, false, 0x45, false, false, false, false},
		{0x3f, false, false, true, true, 0xa5, false, false, false, true},
		{0x3f, false, true, false, false, 0x3f, false, true, false, false},
		{0x3f, false, true, false, true, 0xdf, false, true, false, true},
		{0x3f, false, true, true, false, 0x39, false, true, false, false},
		{0x3f, false, true, true, true, 0xd9, false, true, false, true},
		{0x3f, true, false, false, false, 0x45, false, false, false, false},
		{0x3f, true, false, false, true, 0xa5, false, false, false, true},
		{0x3f, true, false, true, false, 0x45, false, false, false, false},
		{0x3f, true, false, true, true, 0xa5, false, false, false, true},
		{0x3f, true, true, false, false, 0x3f, false, true, false, false},
		{0x3f, true, true, false, true, 0xdf, false, true, false, true},
		{0x3f, true, true, true, false, 0x39, false, true, false, false},
		{0x3f, true, true, true, true, 0xd9, false, true, false, true},
		{0x40, false, false, false, false, 0x40, false, false, false, false},
		{0x40, false, false, false, true, 0xa0, false, false, false, true},
		{0x40, false, false, true, false, 0x46, false, false, false, false},
		{0x40, false, false, true, true, 0xa6, false, false, false, true},
		{0x40, false, true, false, false, 0x40, false, true, false, false},
		{0x40, false, true, false, true, 0xe0, false, true, false, true},
		{0x40, false, true, true, false, 0x3a, false, true, false, false},
		{0x40, false, true, true, true, 0xda, false, true, false, true},
		{0x40, true, false, false, false, 0x40, false, false, false, false},
		{0x40, true, false, false, true, 0xa0, false, false, false, true},
		{0x40, true, false, true, false, 0x46, false, false, false, false},
		{0x40, true, false, true, true, 0xa6, false, false, false, true},
		{0x40, true, true, false, false, 0x40, false, true, false, false},
		{0x40, true, true, false, true, 0xe0, false, true, false, true},
		{0x40, true, true, true, false, 0x3a, false, true, false, false},
		{0x40, true, true, true, true, 0xda, false, true, false, true},
		{0x41, false, false, false, false, 0x41, false, false, false, false},
		{0x41, false, false, false, true, 0xa1, false, false, false, true},
		{0x41, false, false, true, false, 0x47, false, false, false, false},
		{0x41, false, false, true, true, 0xa7, false, false, false, true},
		{0x41, false, true, false, false, 0x41, false, true, false, false},
		{0x41, false, true, false, true, 0xe1, false, true, false, true},
		{0x41, false, true, true, false, 0x3b, false, true, false, false},
		{0x41, false, true, true, true, 0xdb, false, true, false, true},
		{0x41, true, false, false, false, 0x41, false, false, false, false},
		{0x41, true, false, false, true, 0xa1, false, false, false, true},
		{0x41, true, false, true, false, 0x47, false, false, false, false},
		{0x41, true, false, true, true, 0xa7, false, false, false, true},
		{0x41, true, true, false, false, 0x41, false, true, false, false},
		{0x41, true, true, false, true, 0xe1, false, true, false, true},
		{0x41, true, true, true, false, 0x3b, false, true, false, false},
		{0x41, true, true, true, true, 0xdb, false, true, false, true},
		{0x42, false, false, false, false, 0x42, false, false, false, false},
		{0x42, false, false, false, true, 0xa2, false, false, false, true},
		{0x42, false, false, true, false, 0x48, false, false, false, false},
		{0x42, false, false, true, true, 0xa8, false, false, false, true},
		{0x42, false, true, false, false, 0x42, false, true, false, false},
		{0x42, false, true, false, true, 0xe2, false, true, false, true},
		{0x42, false, true, true, false, 0x3c, false, true, false, false},
		{0x42, false, true, true, true, 0xdc, false, true, false, true},
		{0x42, true, false, false, false, 0x42, false, false, false, false},
		{0x42, true, false, false, true, 0xa2, false, false, false, true},
		{0x42, true, false, true, false, 0x48, false, false, false, false},
		{0x42, true, false, true, true, 0xa8, false, false, false, true},
		{0x42, true, true, false, false, 0x42, false, true, false, false},
		{0x42, true, true, false, true, 0xe2, false, true, false, true},
		{0x42, true, true, true, false, 0x3c, false, true, false, false},
		{0x42, true, true, true, true, 0xdc, false, true, false, true},
		{0x43, false, false, false, false, 0x43, false, false, false, false},
		{0x43, false, false, false, true, 0xa3, false, false, false, true},
		{0x43, false, false, true, false, 0x49, false, false, false, false},
		{0x43, false, false, true, true, 0xa9, false, false, false, true},
		{0x43, false, true, false, false, 0x43, false, true, false, false},
		{0x43, false, true, false, true, 0xe3, false, true, false, true},
		{0x43, false, true, true, false, 0x3d, false, true, false, false},
		{0x43, false, true, true, true, 0xdd, false, true, false, true},
		{0x43, true, false, false, false, 0x43, false, false, false, false},
		{0x43, true, false, false, true, 0xa3, false, false, false, true},
		{0x43, true, false, true, false, 0x49, false, false, false, false},
		{0x43, true, false, true, true, 0xa9, false, false, false, true},
		{0x43, true, true, false, false, 0x43, false, true, false, false},
		{0x43, true, true, false, true, 0xe3, false, true, false, true},
		{0x43, true, true, true, false, 0x3d, false, true, false, false},
		{0x43, true, true, true, true, 0xdd, false, true, false, true},
		{0x44, false, false, false, false, 0x44, false, false, false, false},
		{0x44, false, false, false, true, 0xa4, false, false, false, true},
		{0x44, false, false, true, false, 0x4a, false, false, false, false},
		{0x44, false, false, true, true, 0xaa, false, false, false, true},
		{0x44, false, true, false, false, 0x44, false, true, false, false},
		{0x44, false, true, false, true, 0xe4, false, true, false, true},
		{0x44, false, true, true, false, 0x3e, false, true, false, false},
		{0x44, false, true, true, true, 0xde, false, true, false, true},
		{0x44, true, false, false, false, 0x44, false, false, false, false},
		{0x44, true, false, false, true, 0xa4, false, false, false, true},
		{0x44, true, false, true, false, 0x4a, false, false, false, false},
		{0x44, true, false, true, true, 0xaa, false, false, false, true},
		{0x44, true, true, false, false, 0x44, false, true, false, false},
		{0x44, true, true, false, true, 0xe4, false, true, false, true},
		{0x44, true, true, true, false, 0x3e, false, true, false, false},
		{0x44, true, true, true, true, 0xde, false, true, false, true},
		{0x45, false, false, false, false, 0x45, false, false, false, false},
		{0x45, false, false, false, true, 0xa5, false, false, false, true},
		{0x45, false, false, true, false, 0x4b, false, false, false, false},
		{0x45, false, false, true, true, 0xab, false, false, false, true},
		{0x45, false, true, false, false, 0x45, false, true, false, false},
		{0x45, false, true, false, true, 0xe5, false, true, false, true},
		{0x45, false, true, true, false, 0x3f, false, true, false, false},
		{0x45, false, true, true, true, 0xdf, false, true, false, true},
		{0x45, true, false, false, false, 0x45, false, false, false, false},
		{0x45, true, false, false, true, 0xa5, false, false, false, true},
		{0x45, true, false, true, false, 0x4b, false, false, false, false},
		{0x45, true, false, true, true, 0xab, false, false, false, true},
		{0x45, true, true, false, false, 0x45, false, true, false, false},
		{0x45, true, true, false, true, 0xe5, false, true, false, true},
		{0x45, true, true, true, false, 0x3f, false, true, false, false},
		{0x45, true, true, true, true, 0xdf, false, true, false, true},
		{0x46, false, false, false, false, 0x46, false, false, false, false},
		{0x46, false, false, false, true, 0xa6, false, false, false, true},
		{0x46, false, false, true, false, 0x4c, false, false, false, false},
		{0x46, false, false, true, true, 0xac, false, false, false, true},
		{0x46, false, true, false, false, 0x46, false, true, false, false},
		{0x46, false, true, false, true, 0xe6, false, true, false, true},
		{0x46, false, true, true, false, 0x40, false, true, false, false},
		{0x46, false, true, true, true, 0xe0, false, true, false, true},
		{0x46, true, false, false, false, 0x46, false, false, false, false},
		{0x46, true, false, false, true, 0xa6, false, false, false, true},
		{0x46, true, false, true, false, 0x4c, false, false, false, false},
		{0x46, true, false, true, true, 0xac, false, false, false, true},
		{0x46, true, true, false, false, 0x46, false, true, false, false},
		{0x46, true, true, false, true, 0xe6, false, true, false, true},
		{0x46, true, true, true, false, 0x40, false, true, false, false},
		{0x46, true, true, true, true, 0xe0, false, true, false, true},
		{0x47, false, false, false, false, 0x47, false, false, false, false},
		{0x47, false, false, false, true, 0xa7, false, false, false, true},
		{0x47, false, false, true, false, 0x4d, false, false, false, false},
		{0x47, false, false, true, true, 0xad, false, false, false, true},
		{0x47, false, true, false, false, 0x47, false, true, false, false},
		{0x47, false, true, false, true, 0xe7, false, true, false, true},
		{0x47, false, true, true, false, 0x41, false, true, false, false},
		{0x47, false, true, true, true, 0xe1, false, true, false, true},
		{0x47, true, false, false, false, 0x47, false, false, false, false},
		{0x47, true, false, false, true, 0xa7, false, false, false, true},
		{0x47, true, false, true, false, 0x4d, false, false, false, false},
		{0x47, true, false, true, true, 0xad, false, false, false, true},
		{0x47, true, true, false, false, 0x47, false, true, false, false},
		{0x47, true, true, false, true, 0xe7, false, true, false, true},
		{0x47, true, true, true, false, 0x41, false, true, false, false},
		{0x47, true, true, true, true, 0xe1, false, true, false, true},
		{0x48, false, false, false, false, 0x48, false, false, false, false},
		{0x48, false, false, false, true, 0xa8, false, false, false, true},
		{0x48, false, false, true, false, 0x4e, false, false, false, false},
		{0x48, false, false, true, true, 0xae, false, false, false, true},
		{0x48, false, true, false, false, 0x48, false, true, false, false},
		{0x48, false, true, false, true, 0xe8, false, true, false, true},
		{0x48, false, true, true, false, 0x42, false, true, false, false},
		{0x48, false, true, true, true, 0xe2, false, true, false, true},
		{0x48, true, false, false, false, 0x48, false, false, false, false},
		{0x48, true, false, false, true, 0xa8, false, false, false, true},
		{0x48, true, false, true, false, 0x4e, false, false, false, false},
		{0x48, true, false, true, true, 0xae, false, false, false, true},
		{0x48, true, true, false, false, 0x48, false, true, false, false},
		{0x48, true, true, false, true, 0xe8, false, true, false, true},
		{0x48, true, true, true, false, 0x42, false, true, false, false},
		{0x48, true, true, true, true, 0xe2, false, true, false, true},
		{0x49, false, false, false, false, 0x49, false, false, false, false},
		{0x49, false, false, false, true, 0xa9, false, false, false, true},
		{0x49, false, false, true, false, 0x4f, false, false, false, false},
		{0x49, false, false, true, true, 0xaf, false, false, false, true},
		{0x49, false, true, false, false, 0x49, false, true, false, false},
		{0x49, false, true, false, true, 0xe9, false, true, false, true},
		{0x49, false, true, true, false, 0x43, false, true, false, false},
		{0x49, false, true, true, true, 0xe3, false, true, false, true},
		{0x49, true, false, false, false, 0x49, false, false, false, false},
		{0x49, true, false, false, true, 0xa9, false, false, false, true},
		{0x49, true, false, true, false, 0x4f, false, false, false, false},
		{0x49, true, false, true, true, 0xaf, false, false, false, true},
		{0x49, true, true, false, false, 0x49, false, true, false, false},
		{0x49, true, true, false, true, 0xe9, false, true, false, true},
		{0x49, true, true, true, false, 0x43, false, true, false, false},
		{0x49, true, true, true, true, 0xe3, false, true, false, true},
		{0x4a, false, false, false, false, 0x50, false, false, false, false},
		{0x4a, false, false, false, true, 0xb0, false, false, false, true},
		{0x4a, false, false, true, false, 0x50, false, false, false, false},
		{0x4a, false, false, true, true, 0xb0, false, false, false, true},
		{0x4a, false, true, false, false, 0x4a, false, true, false, false},
		{0x4a, false, true, false, true, 0xea, false, true, false, true},
		{0x4a, false, true, true, false, 0x44, false, true, false, false},
		{0x4a, false, true, true, true, 0xe4, false, true, false, true},
		{0x4a, true, false, false, false, 0x50, false, false, false, false},
		{0x4a, true, false, false, true, 0xb0, false, false, false, true},
		{0x4a, true, false, true, false, 0x50, false, false, false, false},
		{0x4a, true, false, true, true, 0xb0, false, false, false, true},
		{0x4a, true, true, false, false, 0x4a, false, true, false, false},
		{0x4a, true, true, false, true, 0xea, false, true, false, true},
		{0x4a, true, true, true, false, 0x44, false, true, false, false},
		{0x4a, true, true, true, true, 0xe4, false, true, false, true},
		{0x4b, false, false, false, false, 0x51, false, false, false, false},
		{0x4b, false, false, false, true, 0xb1, false, false, false, true},
		{0x4b, false, false, true, false, 0x51, false, false, false, false},
		{0x4b, false, false, true, true, 0xb1, false, false, false, true},
		{0x4b, false, true, false, false, 0x4b, false, true, false, false},
		{0x4b, false, true, false, true, 0xeb, false, true, false, true},
		{0x4b, false, true, true, false, 0x45, false, true, false, false},
		{0x4b, false, true, true, true, 0xe5, false, true, false, true},
		{0x4b, true, false, false, false, 0x51, false, false, false, false},
		{0x4b, true, false, false, true, 0xb1, false, false, false, true},
		{0x4b, true, false, true, false, 0x51, false, false, false, false},
		{0x4b, true, false, true, true, 0xb1, false, false, false, true},
		{0x4b, true, true, false, false, 0x4b, false, true, false, false},
		{0x4b, true, true, false, true, 0xeb, false, true, false, true},
		{0x4b, true, true, true, false, 0x45, false, true, false, false},
		{0x4b, true, true, true, true, 0xe5, false, true, false, true},
		{0x4c, false, false, false, false, 0x52, false, false, false, false},
		{0x4c, false, false, false, true, 0xb2, false, false, false, true},
		{0x4c, false, false, true, false, 0x52, false, false, false, false},
		{0x4c, false, false, true, true, 0xb2, false, false, false, true},
		{0x4c, false, true, false, false, 0x4c, false, true, false, false},
		{0x4c, false, true, false, true, 0xec, false, true, false, true},
		{0x4c, false, true, true, false, 0x46, false, true, false, false},
		{0x4c, false, true, true, true, 0xe6, false, true, false, true},
		{0x4c, true, false, false, false, 0x52, false, false, false, false},
		{0x4c, true, false, false, true, 0xb2, false, false, false, true},
		{0x4c, true, false, true, false, 0x52, false, false, false, false},
		{0x4c, true, false, true, true, 0xb2, false, false, false, true},
		{0x4c, true, true, false, false, 0x4c, false, true, false, false},
		{0x4c, true, true, false, true, 0xec, false, true, false, true},
		{0x4c, true, true, true, false, 0x46, false, true, false, false},
		{0x4c, true, true, true, true, 0xe6, false, true, false, true},
		{0x4d, false, false, false, false, 0x53, false, false, false, false},
		{0x4d, false, false, false, true, 0xb3, false, false, false, true},
		{0x4d, false, false, true, false, 0x53, false, false, false, false},
		{0x4d, false, false, true, true, 0xb3, false, false, false, true},
		{0x4d, false, true, false, false, 0x4d, false, true, false, false},
		{0x4d, false, true, false, true, 0xed, false, true, false, true},
		{0x4d, false, true, true, false, 0x47, false, true, false, false},
		{0x4d, false, true, true, true, 0xe7, false, true, false, true},
		{0x4d, true, false, false, false, 0x53, false, false, false, false},
		{0x4d, true, false, false, true, 0xb3, false, false, false, true},
		{0x4d, true, false, true, false, 0x53, false, false, false, false},
		{0x4d, true, false, true, true, 0xb3, false, false, false, true},
		{0x4d, true, true, false, false, 0x4d, false, true, false, false},
		{0x4d, true, true, false, true, 0xed, false, true, false, true},
		{0x4d, true, true, true, false, 0x47, false, true, false, false},
		{0x4d, true, true, true, true, 0xe7, false, true, false, true},
		{0x4e, false, false, false, false, 0x54, false, false, false, false},
		{0x4e, false, false, false, true, 0xb4, false, false, false, true},
		{0x4e, false, false, true, false, 0x54, false, false, false, false},
		{0x4e, false, false, true, true, 0xb4, false, false, false, true},
		{0x4e, false, true, false, false, 0x4e, false, true, false, false},
		{0x4e, false, true, false, true, 0xee, false, true, false, true},
		{0x4e, false, true, true, false, 0x48, false, true, false, false},
		{0x4e, false, true, true, true, 0xe8, false, true, false, true},
		{0x4e, true, false, false, false, 0x54, false, false, false, false},
		{0x4e, true, false, false, true, 0xb4, false, false, false, true},
		{0x4e, true, false, true, false, 0x54, false, false, false, false},
		{0x4e, true, false, true, true, 0xb4, false, false, false, true},
		{0x4e, true, true, false, false, 0x4e, false, true, false, false},
		{0x4e, true, true, false, true, 0xee, false, true, false, true},
		{0x4e, true, true, true, false, 0x48, false, true, false, false},
		{0x4e, true, true, true, true, 0xe8, false, true, false, true},
		{0x4f, false, false, false, false, 0x55, false, false, false, false},
		{0x4f, false, false, false, true, 0xb5, false, false, false, true},
		{0x4f, false, false, true, false, 0x55, false, false, false, false},
		{0x4f, false, false, true, true, 0xb5, false, false, false, true},
		{0x4f, false, true, false, false, 0x4f, false, true, false, false},
		{0x4f, false, true, false, true, 0xef, false, true, false, true},
		{0x4f, false, true, true, false, 0x49, false, true, false, false},
		{0x4f, false, true, true, true, 0xe9, false, true, false, true},
		{0x4f, true, false, false, false, 0x55, false, false, false, false},
		{0x4f, true, false, false, true, 0xb5, false, false, false, true},
		{0x4f, true, false, true, false, 0x55, false, false, false, false},
		{0x4f, true, false, true, true, 0xb5, false, false, false, true},
		{0x4f, true, true, false, false, 0x4f, false, true, false, false},
		{0x4f, true, true, false, true, 0xef, false, true, false, true},
		{0x4f, true, true, true, false, 0x49, false, true, false, false},
		{0x4f, true, true, true, true, 0xe9, false, true, false, true},
		{0x50, false, false, false, false, 0x50, false, false, false, false},
		{0x50, false, false, false, true, 0xb0, false, false, false, true},
		{0x50, false, false, true, false, 0x56, false, false, false, false},
		{0x50, false, false, true, true, 0xb6, false, false, false, true},
		{0x50, false, true, false, false, 0x50, false, true, false, false},
		{0x50, false, true, false, true, 0xf0, false, true, false, true},
		{0x50, false, true, true, false, 0x4a, false, true, false, false},
		{0x50, false, true, true, true, 0xea, false, true, false, true},
		{0x50, true, false, false, false, 0x50, false, false, false, false},
		{0x50, true, false, false, true, 0xb0, false, false, false, true},
		{0x50, true, false, true, false, 0x56, false, false, false, false},
		{0x50, true, false, true, true, 0xb6, false, false, false, true},
		{0x50, true, true, false, false, 0x50, false, true, false, false},
		{0x50, true, true, false, true, 0xf0, false, true, false, true},
		{0x50, true, true, true, false, 0x4a, false, true, false, false},
		{0x50, true, true, true, true, 0xea, false, true, false, true},
		{0x51, false, false, false, false, 0x51, false, false, false, false},
		{0x51, false, false, false, true, 0xb1, false, false, false, true},
		{0x51, false, false, true, false, 0x57, false, false, false, false},
		{0x51, false, false, true, true, 0xb7, false, false, false, true},
		{0x51, false, true, false, false, 0x51, false, true, false, false},
		{0x51, false, true, false, true, 0xf1, false, true, false, true},
		{0x51, false, true, true, false, 0x4b, false, true, false, false},
		{0x51, false, true, true, true, 0xeb, false, true, false, true},
		{0x51, true, false, false, false, 0x51, false, false, false, false},
		{0x51, true, false, false, true, 0xb1, false, false, false, true},
		{0x51, true, false, true, false, 0x57, false, false, false, false},
		{0x51, true, false, true, true, 0xb7, false, false, false, true},
		{0x51, true, true, false, false, 0x51, false, true, false, false},
		{0x51, true, true, false, true, 0xf1, false, true, false, true},
		{0x51, true, true, true, false, 0x4b, false, true, false, false},
		{0x51, true, true, true, true, 0xeb, false, true, false, true},
		{0x52, false, false, false, false, 0x52, false, false, false, false},
		{0x52, false, false, false, true, 0xb2, false, false, false, true},
		{0x52, false, false, true, false, 0x58, false, false, false, false},
		{0x52, false, false, true, true, 0xb8, false, false, false, true},
		{0x52, false, true, false, false, 0x52, false, true, false, false},
		{0x52, false, true, false, true, 0xf2, false, true, false, true},
		{0x52, false, true, true, false, 0x4c, false, true, false, false},
		{0x52, false, true, true, true, 0xec, false, true, false, true},
		{0x52, true, false, false, false, 0x52, false, false, false, false},
		{0x52, true, false, false, true, 0xb2, false, false, false, true},
		{0x52, true, false, true, false, 0x58, false, false, false, false},
		{0x52, true, false, true, true, 0xb8, false, false, false, true},
		{0x52, true, true, false, false, 0x52, false, true, false, false},
		{0x52, true, true, false, true, 0xf2, false, true, false, true},
		{0x52, true, true, true, false, 0x4c, false, true, false, false},
		{0x52, true, true, true, true, 0xec, false, true, false, true},
		{0x53, false, false, false, false, 0x53, false, false, false, false},
		{0x53, false, false, false, true, 0xb3, false, false, false, true},
		{0x53, false, false, true, false, 0x59, false, false, false, false},
		{0x53, false, false, true, true, 0xb9, false, false, false, true},
		{0x53, false, true, false, false, 0x53, false, true, false, false},
		{0x53, false, true, false, true, 0xf3, false, true, false, true},
		{0x53, false, true, true, false, 0x4d, false, true, false, false},
		{0x53, false, true, true, true, 0xed, false, true, false, true},
		{0x53, true, false, false, false, 0x53, false, false, false, false},
		{0x53, true, false, false, true, 0xb3, false, false, false, true},
		{0x53, true, false, true, false, 0x59, false, false, false, false},
		{0x53, true, false, true, true, 0xb9, false, false, false, true},
		{0x53, true, true, false, false, 0x53, false, true, false, false},
		{0x53, true, true, false, true, 0xf3, false, true, false, true},
		{0x53, true, true, true, false, 0x4d, false, true, false, false},
		{0x53, true, true, true, true, 0xed, false, true, false, true},
		{0x54, false, false, false, false, 0x54, false, false, false, false},
		{0x54, false, false, false, true, 0xb4, false, false, false, true},
		{0x54, false, false, true, false, 0x5a, false, false, false, false},
		{0x54, false, false, true, true, 0xba, false, false, false, true},
		{0x54, false, true, false, false, 0x54, false, true, false, false},
		{0x54, false, true, false, true, 0xf4, false, true, false, true},
		{0x54, false, true, true, false, 0x4e, false, true, false, false},
		{0x54, false, true, true, true, 0xee, false, true, false, true},
		{0x54, true, false, false, false, 0x54, false, false, false, false},
		{0x54, true, false, false, true, 0xb4, false, false, false, true},
		{0x54, true, false, true, false, 0x5a, false, false, false, false},
		{0x54, true, false, true, true, 0xba, false, false, false, true},
		{0x54, true, true, false, false, 0x54, false, true, false, false},
		{0x54, true, true, false, true, 0xf4, false, true, false, true},
		{0x54, true, true, true, false, 0x4e, false, true, false, false},
		{0x54, true, true, true, true, 0xee, false, true, false, true},
		{0x55, false, false, false, false, 0x55, false, false, false, false},
		{0x55, false, false, false, true, 0xb5, false, false, false, true},
		{0x55, false, false, true, false, 0x5b, false, false, false, false},
		{0x55, false, false, true, true, 0xbb, false, false, false, true},
		{0x55, false, true, false, false, 0x55, false, true, false, false},
		{0x55, false, true, false, true, 0xf5, false, true, false, true},
		{0x55, false, true, true, false, 0x4f, false, true, false, false},
		{0x55, false, true, true, true, 0xef, false, true, false, true},
		{0x55, true, false, false, false, 0x55, false, false, false, false},
		{0x55, true, false, false, true, 0xb5, false, false, false, true},
		{0x55, true, false, true, false, 0x5b, false, false, false, false},
		{0x55, true, false, true, true, 0xbb, false, false, false, true},
		{0x55, true, true, false, false, 0x55, false, true, false, false},
		{0x55, true, true, false, true, 0xf5, false, true, false, true},
		{0x55, true, true, true, false, 0x4f, false, true, false, false},
		{0x55, true, true, true, true, 0xef, false, true, false, true},
		{0x56, false, false, false, false, 0x56, false, false, false, false},
		{0x56, false, false, false, true, 0xb6, false, false, false, true},
		{0x56, false, false, true, false, 0x5c, false, false, false, false},
		{0x56, false, false, true, true, 0xbc, false, false, false, true},
		{0x56, false, true, false, false, 0x56, false, true, false, false},
		{0x56, false, true, false, true, 0xf6, false, true, false, true},
		{0x56, false, true, true, false, 0x50, false, true, false, false},
		{0x56, false, true, true, true, 0xf0, false, true, false, true},
		{0x56, true, false, false, false, 0x56, false, false, false, false},
		{0x56, true, false, false, true, 0xb6, false, false, false, true},
		{0x56, true, false, true, false, 0x5c, false, false, false, false},
		{0x56, true, false, true, true, 0xbc, false, false, false, true},
		{0x56, true, true, false, false, 0x56, false, true, false, false},
		{0x56, true, true, false, true, 0xf6, false, true, false, true},
		{0x56, true, true, true, false, 0x50, false, true, false, false},
		{0x56, true, true, true, true, 0xf0, false, true, false, true},
		{0x57, false, false, false, false, 0x57, false, false, false, false},
		{0x57, false, false, false, true, 0xb7, false, false, false, true},
		{0x57, false, false, true, false, 0x5d, false, false, false, false},
		{0x57, false, false, true, true, 0xbd, false, false, false, true},
		{0x57, false, true, false, false, 0x57, false, true, false, false},
		{0x57, false, true, false, true, 0xf7, false, true, false, true},
		{0x57, false, true, true, false, 0x51, false, true, false, false},
		{0x57, false, true, true, true, 0xf1, false, true, false, true},
		{0x57, true, false, false, false, 0x57, false, false, false, false},
		{0x57, true, false, false, true, 0xb7, false, false, false, true},
		{0x57, true, false, true, false, 0x5d, false, false, false, false},
		{0x57, true, false, true, true, 0xbd, false, false, false, true},
		{0x57, true, true, false, false, 0x57, false, true, false, false},
		{0x57, true, true, false, true, 0xf7, false, true, false, true},
		{0x57, true, true, true, false, 0x51, false, true, false, false},
		{0x57, true, true, true, true, 0xf1, false, true, false, true},
		{0x58, false, false, false, false, 0x58, false, false, false, false},
		{0x58, false, false, false, true, 0xb8, false, false, false, true},
		{0x58, false, false, true, false, 0x5e, false, false, false, false},
		{0x58, false, false, true, true, 0xbe, false, false, false, true},
		{0x58, false, true, false, false, 0x58, false, true, false, false},
		{0x58, false, true, false, true, 0xf8, false, true, false, true},
		{0x58, false, true, true, false, 0x52, false, true, false, false},
		{0x58, false, true, true, true, 0xf2, false, true, false, true},
		{0x58, true, false, false, false, 0x58, false, false, false, false},
		{0x58, true, false, false, true, 0xb8, false, false, false, true},
		{0x58, true, false, true, false, 0x5e, false, false, false, false},
		{0x58, true, false, true, true, 0xbe, false, false, false, true},
		{0x58, true, true, false, false, 0x58, false, true, false, false},
		{0x58, true, true, false, true, 0xf8, false, true, false, true},
		{0x58, true, true, true, false, 0x52, false, true, false, false},
		{0x58, true, true, true, true, 0xf2, false, true, false, true},
		{0x59, false, false, false, false, 0x59, false, false, false, false},
		{0x59, false, false, false, true, 0xb9, false, false, false, true},
		{0x59, false, false, true, false, 0x5f, false, false, false, false},
		{0x59, false, false, true, true, 0xbf, false, false, false, true},
		{0x59, false, true, false, false, 0x59, false, true, false, false},
		{0x59, false, true, false, true, 0xf9, false, true, false, true},
		{0x59, false, true, true, false, 0x53, false, true, false, false},
		{0x59, false, true, true, true, 0xf3, false, true, false, true},
		{0x59, true, false, false, false, 0x59, false, false, false, false},
		{0x59, true, false, false, true, 0xb9, false, false, false, true},
		{0x59, true, false, true, false, 0x5f, false, false, false, false},
		{0x59, true, false, true, true, 0xbf, false, false, false, true},
		{0x59, true, true, false, false, 0x59, false, true, false, false},
		{0x59, true, true, false, true, 0xf9, false, true, false, true},
		{0x59, true, true, true, false, 0x53, false, true, false, false},
		{0x59, true, true, true, true, 0xf3, false, true, false, true},
		{0x5a, false, false, false, false, 0x60, false, false, false, false},
		{0x5a, false, false, false, true, 0xc0, false, false, false, true},
		{0x5a, false, false, true, false, 0x60, false, false, false, false},
		{0x5a, false, false, true, true, 0xc0, false, false, false, true},
		{0x5a, false, true, false, false, 0x5a, false, true, false, false},
		{0x5a, false, true, false, true, 0xfa, false, true, false, true},
		{0x5a, false, true, true, false, 0x54, false, true, false, false},
		{0x5a, false, true, true, true, 0xf4, false, true, false, true},
		{0x5a, true, false, false, false, 0x60, false, false, false, false},
		{0x5a, true, false, false, true, 0xc0, false, false, false, true},
		{0x5a, true, false, true, false, 0x60, false, false, false, false},
		{0x5a, true, false, true, true, 0xc0, false, false, false, true},
		{0x5a, true, true, false, false, 0x5a, false, true, false, false},
		{0x5a, true, true, false, true, 0xfa, false, true, false, true},
		{0x5a, true, true, true, false, 0x54, false, true, false, false},
		{0x5a, true, true, true, true, 0xf4, false, true, false, true},
		{0x5b, false, false, false, false, 0x61, false, false, false, false},
		{0x5b, false, false, false, true, 0xc1, false, false, false, true},
		{0x5b, false, false, true, false, 0x61, false, false, false, false},
		{0x5b, false, false, true, true, 0xc1, false, false, false, true},
		{0x5b, false, true, false, false, 0x5b, false, true, false, false},
		{0x5b, false, true, false, true, 0xfb, false, true, false, true},
		{0x5b, false, true, true, false, 0x55, false, true, false, false},
		{0x5b, false, true, true, true, 0xf5, false, true, false, true},
		{0x5b, true, false, false, false, 0x61, false, false, false, false},
		{0x5b, true, false, false, true, 0xc1, false, false, false, true},
		{0x5b, true, false, true, false, 0x61, false, false, false, false},
		{0x5b, true, false, true, true, 0xc1, false, false, false, true},
		{0x5b, true, true, false, false, 0x5b, false, true, false, false},
		{0x5b, true, true, false, true, 0xfb, false, true, false, true},
		{0x5b, true, true, true, false, 0x55, false, true, false, false},
		{0x5b, true, true, true, true, 0xf5, false, true, false, true},
		{0x5c, false, false, false, false, 0x62, false, false, false, false},
		{0x5c, false, false, false, true, 0xc2, false, false, false, true},
		{0x5c, false, false, true, false, 0x62, false, false, false, false},
		{0x5c, false, false, true, true, 0xc2, false, false, false, true},
		{0x5c, false, true, false, false, 0x5c, false, true, false, false},
		{0x5c, false, true, false, true, 0xfc, false, true, false, true},
		{0x5c, false, true, true, false, 0x56, false, true, false, false},
		{0x5c, false, true, true, true, 0xf6, false, true, false, true},
		{0x5c, true, false, false, false, 0x62, false, false, false, false},
		{0x5c, true, false, false, true, 0xc2, false, false, false, true},
		{0x5c, true, false, true, false, 0x62, false, false, false, false},
		{0x5c, true, false, true, true, 0xc2, false, false, false, true},
		{0x5c, true, true, false, false, 0x5c, false, true, false, false},
		{0x5c, true, true, false, true, 0xfc, false, true, false, true},
		{0x5c, true, true, true, false, 0x56, false, true, false, false},
		{0x5c, true, true, true, true, 0xf6, false, true, false, true},
		{0x5d, false, false, false, false, 0x63, false, false, false, false},
		{0x5d, false, false, false, true, 0xc3, false, false, false, true},
		{0x5d, false, false, true, false, 0x63, false, false, false, false},
		{0x5d, false, false, true, true, 0xc3, false, false, false, true},
		{0x5d, false, true, false, false, 0x5d, false, true, false, false},
		{0x5d, false, true, false, true, 0xfd, false, true, false, true},
		{0x5d, false, true, true, false, 0x57, false, true, false, false},
		{0x5d, false, true, true, true, 0xf7, false, true, false, true},
		{0x5d, true, false, false, false, 0x63, false, false, false, false},
		{0x5d, true, false, false, true, 0xc3, false, false, false, true},
		{0x5d, true, false, true, false, 0x63, false, false, false, false},
		{0x5d, true, false, true, true, 0xc3, false, false, false, true},
		{0x5d, true, true, false, false, 0x5d, false, true, false, false},
		{0x5d, true, true, false, true, 0xfd, false, true, false, true},
		{0x5d, true, true, true, false, 0x57, false, true, false, false},
		{0x5d, true, true, true, true, 0xf7, false, true, false, true},
		{0x5e, false, false, false, false, 0x64, false, false, false, false},
		{0x5e, false, false, false, true, 0xc4, false, false, false, true},
		{0x5e, false, false, true, false, 0x64, false, false, false, false},
		{0x5e, false, false, true, true, 0xc4, false, false, false, true},
		{0x5e, false, true, false, false, 0x5e, false, true, false, false},
		{0x5e, false, true, false, true, 0xfe, false, true, false, true},
		{0x5e, false, true, true, false, 0x58, false, true, false, false},
		{0x5e, false, true, true, true, 0xf8, false, true, false, true},
		{0x5e, true, false, false, false, 0x64, false, false, false, false},
		{0x5e, true, false, false, true, 0xc4, false, false, false, true},
		{0x5e, true, false, true, false, 0x64, false, false, false, false},
		{0x5e, true, false, true, true, 0xc4, false, false, false, true},
		{0x5e, true, true, false, false, 0x5e, false, true, false, false},
		{0x5e, true, true, false, true, 0xfe, false, true, false, true},
		{0x5e, true, true, true, false, 0x58, false, true, false, false},
		{0x5e, true, true, true, true, 0xf8, false, true, false, true},
		{0x5f, false, false, false, false, 0x65, false, false, false, false},
		{0x5f, false, false, false, true, 0xc5, false, false, false, true},
		{0x5f, false, false, true, false, 0x65, false, false, false, false},
		{0x5f, false, false, true, true, 0xc5, false, false, false, true},
		{0x5f, false, true, false, false, 0x5f, false, true, false, false},
		{0x5f, false, true, false, true, 0xff, false, true, false, true},
		{0x5f, false, true, true, false, 0x59, false, true, false, false},
		{0x5f, false, true, true, true, 0xf9, false, true, false, true},
		{0x5f, true, false, false, false, 0x65, false, false, false, false},
		{0x5f, true, false, false, true, 0xc5, false, false, false, true},
		{0x5f, true, false, true, false, 0x65, false, false, false, false},
		{0x5f, true, false, true, true, 0xc5, false, false, false, true},
		{0x5f, true, true, false, false, 0x5f, false, true, false, false},
		{0x5f, true, true, false, true, 0xff, false, true, false, true},
		{0x5f, true, true, true, false, 0x59, false, true, false, false},
		{0x5f, true, true, true, true, 0xf9, false, true, false, true},
		{0x60, false, false, false, false, 0x60, false, false, false, false},
		{0x60, false, false, false, true, 0xc0, false, false, false, true},
		{0x60, false, false, true, false, 0x66, false, false, false, false},
		{0x60, false, false, true, true, 0xc6, false, false, false, true},
		{0x60, false, true, false, false, 0x60, false, true, false, false},
		{0x60, false, true, false, true, 0x00, true, true, false, true},
		{0x60, false, true, true, false, 0x5a, false, true, false, false},
		{0x60, false, true, true, true, 0xfa, false, true, false, true},
		{0x60, true, false, false, false, 0x60, false, false, false, false},
		{0x60, true, false, false, true, 0xc0, false, false, false, true},
		{0x60, true, false, true, false, 0x66, false, false, false, false},
		{0x60, true, false, true, true, 0xc6, false, false, false, true},
		{0x60, true, true, false, false, 0x60, false, true, false, false},
		{0x60, true, true, false, true, 0x00, true, true, false, true},
		{0x60, true, true, true, false, 0x5a, false, true, false, false},
		{0x60, true, true, true, true, 0xfa, false, true, false, true},
		{0x61, false, false, false, false, 0x61, false, false, false, false},
		{0x61, false, false, false, true, 0xc1, false, false, false, true},
		{0x61, false, false, true, false, 0x67, false, false, false, false},
		{0x61, false, false, true, true, 0xc7, false, false, false, true},
		{0x61, false, true, false, false, 0x61, false, true, false, false},
		{0x61, false, true, false, true, 0x01, false, true, false, true},
		{0x61, false, true, true, false, 0x5b, false, true, false, false},
		{0x61, false, true, true, true, 0xfb, false, true, false, true},
		{0x61, true, false, false, false, 0x61, false, false, false, false},
		{0x61, true, false, false, true, 0xc1, false, false, false, true},
		{0x61, true, false, true, false, 0x67, false, false, false, false},
		{0x61, true, false, true, true, 0xc7, false, false, false, true},
		{0x61, true, true, false, false, 0x61, false, true, false, false},
		{0x61, true, true, false, true, 0x01, false, true, false, true},
		{0x61, true, true, true, false, 0x5b, false, true, false, false},
		{0x61, true, true, true, true, 0xfb, false, true, false, true},
		{0x62, false, false, false, false, 0x62, false, false, false, false},
		{0x62, false, false, false, true, 0xc2, false, false, false, true},
		{0x62, false, false, true, false, 0x68, false, false, false, false},
		{0x62, false, false, true, true, 0xc8, false, false, false, true},
		{0x62, false, true, false, false, 0x62, false, true, false, false},
		{0x62, false, true, false, true, 0x02, false, true, false, true},
		{0x62, false, true, true, false, 0x5c, false, true, false, false},
		{0x62, false, true, true, true, 0xfc, false, true, false, true},
		{0x62, true, false, false, false, 0x62, false, false, false, false},
		{0x62, true, false, false, true, 0xc2, false, false, false, true},
		{0x62, true, false, true, false, 0x68, false, false, false, false},
		{0x62, true, false, true, true, 0xc8, false, false, false, true},
		{0x62, true, true, false, false, 0x62, false, true, false, false},
		{0x62, true, true, false, true, 0x02, false, true, false, true},
		{0x62, true, true, true, false, 0x5c, false, true, false, false},
		{0x62, true, true, true, true, 0xfc, false, true, false, true},
		{0x63, false, false, false, false, 0x63, false, false, false, false},
		{0x63, false, false, false, true, 0xc3, false, false, false, true},
		{0x63, false, false, true, false, 0x69, false, false, false, false},
		{0x63, false, false, true, true, 0xc9, false, false, false, true},
		{0x63, false, true, false, false, 0x63, false, true, false, false},
		{0x63, false, true, false, true, 0x03, false, true, false, true},
		{0x63, false, true, true, false, 0x5d, false, true, false, false},
		{0x63, false, true, true, true, 0xfd, false, true, false, true},
		{0x63, true, false, false, false, 0x63, false, false, false, false},
		{0x63, true, false, false, true, 0xc3, false, false, false, true},
		{0x63, true, false, true, false, 0x69, false, false, false, false},
		{0x63, true, false, true, true, 0xc9, false, false, false, true},
		{0x63, true, true, false, false, 0x63, false, true, false, false},
		{0x63, true, true, false, true, 0x03, false, true, false, true},
		{0x63, true, true, true, false, 0x5d, false, true, false, false},
		{0x63, true, true, true, true, 0xfd, false, true, false, true},
		{0x64, false, false, false, false, 0x64, false, false, false, false},
		{0x64, false, false, false, true, 0xc4, false, false, false, true},
		{0x64, false, false, true, false, 0x6a, false, false, false, false},
		{0x64, false, false, true, true, 0xca, false, false, false, true},
		{0x64, false, true, false, false, 0x64, false, true, false, false},
		{0x64, false, true, false, true, 0x04, false, true, false, true},
		{0x64, false, true, true, false, 0x5e, false, true, false, false},
		{0x64, false, true, true, true, 0xfe, false, true, false, true},
		{0x64, true, false, false, false, 0x64, false, false, false, false},
		{0x64, true, false, false, true, 0xc4, false, false, false, true},
		{0x64, true, false, true, false, 0x6a, false, false, false, false},
		{0x64, true, false, true, true, 0xca, false, false, false, true},
		{0x64, true, true, false, false, 0x64, false, true, false, false},
		{0x64, true, true, false, true, 0x04, false, true, false, true},
		{0x64, true, true, true, false, 0x5e, false, true, false, false},
		{0x64, true, true, true, true, 0xfe, false, true, false, true},
		{0x65, false, false, false, false, 0x65, false, false, false, false},
		{0x65, false, false, false, true, 0xc5, false, false, false, true},
		{0x65, false, false, true, false, 0x6b, false, false, false, false},
		{0x65, false, false, true, true, 0xcb, false, false, false, true},
		{0x65, false, true, false, false, 0x65, false, true, false, false},
		{0x65, false, true, false, true, 0x05, false, true, false, true},
		{0x65, false, true, true, false, 0x5f, false, true, false, false},
		{0x65, false, true, true, true, 0xff, false, true, false, true},
		{0x65, true, false, false, false, 0x65, false, false, false, false},
		{0x65, true, false, false, true, 0xc5, false, false, false, true},
		{0x65, true, false, true, false, 0x6b, false, false, false, false},
		{0x65, true, false, true, true, 0xcb, false, false, false, true},
		{0x65, true, true, false, false, 0x65, false, true, false, false},
		{0x65, true, true, false, true, 0x05, false, true, false, true},
		{0x65, true, true, true, false, 0x5f, false, true, false, false},
		{0x65, true, true, true, true, 0xff, false, true, false, true},
		{0x66, false, false, false, false, 0x66, false, false, false, false},
		{0x66, false, false, false, true, 0xc6, false, false, false, true},
		{0x66, false, false, true, false, 0x6c, false, false, false, false},
		{0x66, false, false, true, true, 0xcc, false, false, false, true},
		{0x66, false, true, false, false, 0x66, false, true, false, false},
		{0x66, false, true, false, true, 0x06, false, true, false, true},
		{0x66, false, true, true, false, 0x60, false, true, false, false},
		{0x66, false, true, true, true, 0x00, true, true, false, true},
		{0x66, true, false, false, false, 0x66, false, false, false, false},
		{0x66, true, false, false, true, 0xc6, false, false, false, true},
		{0x66, true, false, true, false, 0x6c, false, false, false, false},
		{0x66, true, false, true, true, 0xcc, false, false, false, true},
		{0x66, true, true, false, false, 0x66, false, true, false, false},
		{0x66, true, true, false, true, 0x06, false, true, false, true},
		{0x66, true, true, true, false, 0x60, false, true, false, false},
		{0x66, true, true, true, true, 0x00, true, true, false, true},
		{0x67, false, false, false, false, 0x67, false, false, false, false},
		{0x67, false, false, false, true, 0xc7, false, false, false, true},
		{0x67, false, false, true, false, 0x6d, false, false, false, false},
		{0x67, false, false, true, true, 0xcd, false, false, false, true},
		{0x67, false, true, false, false, 0x67, false, true, false, false},
		{0x67, false, true, false, true, 0x07, false, true, false, true},
		{0x67, false, true, true, false, 0x61, false, true, false, false},
		{0x67, false, true, true, true, 0x01, false, true, false, true},
		{0x67, true, false, false, false, 0x67, false, false, false, false},
		{0x67, true, false, false, true, 0xc7, false, false, false, true},
		{0x67, true, false, true, false, 0x6d, false, false, false, false},
		{0x67, true, false, true, true, 0xcd, false, false, false, true},
		{0x67, true, true, false, false, 0x67, false, true, false, false},
		{0x67, true, true, false, true, 0x07, false, true, false, true},
		{0x67, true, true, true, false, 0x61, false, true, false, false},
		{0x67, true, true, true, true, 0x01, false, true, false, true},
		{0x68, false, false, false, false, 0x68, false, false, false, false},
		{0x68, false, false, false, true, 0xc8, false, false, false, true},
		{0x68, false, false, true, false, 0x6e, false, false, false, false},
		{0x68, false, false, true, true, 0xce, false, false, false, true},
		{0x68, false, true, false, false, 0x68, false, true, false, false},
		{0x68, false, true, false, true, 0x08, false, true, false, true},
		{0x68, false, true, true, false, 0x62, false, true, false, false},
		{0x68, false, true, true, true, 0x02, false, true, false, true},
		{0x68, true, false, false, false, 0x68, false, false, false, false},
		{0x68, true, false, false, true, 0xc8, false, false, false, true},
		{0x68, true, false, true, false, 0x6e, false, false, false, false},
		{0x68, true, false, true, true, 0xce, false, false, false, true},
		{0x68, true, true, false, false, 0x68, false, true, false, false},
		{0x68, true, true, false, true, 0x08, false, true, false, true},
		{0x68, true, true, true, false, 0x62, false, true, false, false},
		{0x68, true, true, true, true, 0x02, false, true, false, true},
		{0x69, false, false, false, false, 0x69, false, false, false, false},
		{0x69, false, false, false, true, 0xc9, false, false, false, true},
		{0x69, false, false, true, false, 0x6f, false, false, false, false},
		{0x69, false, false, true, true, 0xcf, false, false, false, true},
		{0x69, false, true, false, false, 0x69, false, true, false, false},
		{0x69, false, true, false, true, 0x09, false, true, false, true},
		{0x69, false, true, true, false, 0x63, false, true, false, false},
		{0x69, false, true, true, true, 0x03, false, true, false, true},
		{0x69, true, false, false, false, 0x69, false, false, false, false},
		{0x69, true, false, false, true, 0xc9, false, false, false, true},
		{0x69, true, false, true, false, 0x6f, false, false, false, false},
		{0x69, true, false, true, true, 0xcf, false, false, false, true},
		{0x69, true, true, false, false, 0x69, false, true, false, false},
		{0x69, true, true, false, true, 0x09, false, true, false, true},
		{0x69, true, true, true, false, 0x63, false, true, false, false},
		{0x69, true, true, true, true, 0x03, false, true, false, true},
		{0x6a, false, false, false, false, 0x70, false, false, false, false},
		{0x6a, false, false, false, true, 0xd0, false, false, false, true},
		{0x6a, false, false, true, false, 0x70, false, false, false, false},
		{0x6a, false, false, true, true, 0xd0, false, false, false, true},
		{0x6a, false, true, false, false, 0x6a, false, true, false, false},
		{0x6a, false, true, false, true, 0x0a, false, true, false, true},
		{0x6a, false, true, true, false, 0x64, false, true, false, false},
		{0x6a, false, true, true, true, 0x04, false, true, false, true},
		{0x6a, true, false, false, false, 0x70, false, false, false, false},
		{0x6a, true, false, false, true, 0xd0, false, false, false, true},
		{0x6a, true, false, true, false, 0x70, false, false, false, false},
		{0x6a, true, false, true, true, 0xd0, false, false, false, true},
		{0x6a, true, true, false, false, 0x6a, false, true, false, false},
		{0x6a, true, true, false, true, 0x0a, false, true, false, true},
		{0x6a, true, true, true, false, 0x64, false, true, false, false},
		{0x6a, true, true, true, true, 0x04, false, true, false, true},
		{0x6b, false, false, false, false, 0x71, false, false, false, false},
		{0x6b, false, false, false, true, 0xd1, false, false, false, true},
		{0x6b, false, false, true, false, 0x71, false, false, false, false},
		{0x6b, false, false, true, true, 0xd1, false, false, false, true},
		{0x6b, false, true, false, false, 0x6b, false, true, false, false},
		{0x6b, false, true, false, true, 0x0b, false, true, false, true},
		{0x6b, false, true, true, false, 0x65, false, true, false, false},
		{0x6b, false, true, true, true, 0x05, false, true, false, true},
		{0x6b, true, false, false, false, 0x71, false, false, false, false},
		{0x6b, true, false, false, true, 0xd1, false, false, false, true},
		{0x6b, true, false, true, false, 0x71, false, false, false, false},
		{0x6b, true, false, true, true, 0xd1, false, false, false, true},
		{0x6b, true, true, false, false, 0x6b, false, true, false, false},
		{0x6b, true, true, false, true, 0x0b, false, true, false, true},
		{0x6b, true, true, true, false, 0x65, false, true, false, false},
		{0x6b, true, true, true, true, 0x05, false, true, false, true},
		{0x6c, false, false, false, false, 0x72, false, false, false, false},
		{0x6c, false, false, false, true, 0xd2, false, false, false, true},
		{0x6c, false, false, true, false, 0x72, false, false, false, false},
		{0x6c, false, false, true, true, 0xd2, false, false, false, true},
		{0x6c, false, true, false, false, 0x6c, false, true, false, false},
		{0x6c, false, true, false, true, 0x0c, false, true, false, true},
		{0x6c, false, true, true, false, 0x66, false, true, false, false},
		{0x6c, false, true, true, true, 0x06, false, true, false, true},
		{0x6c, true, false, false, false, 0x72, false, false, false, false},
		{0x6c, true, false, false, true, 0xd2, false, false, false, true},
		{0x6c, true, false, true, false, 0x72, false, false, false, false},
		{0x6c, true, false, true, true, 0xd2, false, false, false, true},
		{0x6c, true, true, false, false, 0x6c, false, true, false, false},
		{0x6c, true, true, false, true, 0x0c, false, true, false, true},
		{0x6c, true, true, true, false, 0x66, false, true, false, false},
		{0x6c, true, true, true, true, 0x06, false, true, false, true},
		{0x6d, false, false, false, false, 0x73, false, false, false, false},
		{0x6d, false, false, false, true, 0xd3, false, false, false, true},
		{0x6d, false, false, true, false, 0x73, false, false, false, false},
		{0x6d, false, false, true, true, 0xd3, false, false, false, true},
		{0x6d, false, true, false, false, 0x6d, false, true, false, false},
		{0x6d, false, true, false, true, 0x0d, false, true, false, true},
		{0x6d, false, true, true, false, 0x67, false, true, false, false},
		{0x6d, false, true, true, true, 0x07, false, true, false, true},
		{0x6d, true, false, false, false, 0x73, false, false, false, false},
		{0x6d, true, false, false, true, 0xd3, false, false, false, true},
		{0x6d, true, false, true, false, 0x73, false, false, false, false},
		{0x6d, true, false, true, true, 0xd3, false, false, false, true},
		{0x6d, true, true, false, false, 0x6d, false, true, false, false},
		{0x6d, true, true, false, true, 0x0d, false, true, false, true},
		{0x6d, true, true, true, false, 0x67, false, true, false, false},
		{0x6d, true, true, true, true, 0x07, false, true, false, true},
		{0x6e, false, false, false, false, 0x74, false, false, false, false},
		{0x6e, false, false, false, true, 0xd4, false, false, false, true},
		{0x6e, false, false, true, false, 0x74, false, false, false, false},
		{0x6e, false, false, true, true, 0xd4, false, false, false, true},
		{0x6e, false, true, false, false, 0x6e, false, true, false, false},
		{0x6e, false, true, false, true, 0x0e, false, true, false, true},
		{0x6e, false, true, true, false, 0x68, false, true, false, false},
		{0x6e, false, true, true, true, 0x08, false, true, false, true},
		{0x6e, true, false, false, false, 0x74, false, false, false, false},
		{0x6e, true, false, false, true, 0xd4, false, false, false, true},
		{0x6e, true, false, true, false, 0x74, false, false, false, false},
		{0x6e, true, false, true, true, 0xd4, false, false, false, true},
		{0x6e, true, true, false, false, 0x6e, false, true, false, false},
		{0x6e, true, true, false, true, 0x0e, false, true, false, true},
		{0x6e, true, true, true, false, 0x68, false, true, false, false},
		{0x6e, true, true, true, true, 0x08, false, true, false, true},
		{0x6f, false, false, false, false, 0x75, false, false, false, false},
		{0x6f, false, false, false, true, 0xd5, false, false, false, true},
		{0x6f, false, false, true, false, 0x75, false, false, false, false},
		{0x6f, false, false, true, true, 0xd5, false, false, false, true},
		{0x6f, false, true, false, false, 0x6f, false, true, false, false},
		{0x6f, false, true, false, true, 0x0f, false, true, false, true},
		{0x6f, false, true, true, false, 0x69, false, true, false, false},
		{0x6f, false, true, true, true, 0x09, false, true, false, true},
		{0x6f, true, false, false, false, 0x75, false, false, false, false},
		{0x6f, true, false, false, true, 0xd5, false, false, false, true},
		{0x6f, true, false, true, false, 0x75, false, false, false, false},
		{0x6f, true, false, true, true, 0xd5, false, false, false, true},
		{0x6f, true, true, false, false, 0x6f, false, true, false, false},
		{0x6f, true, true, false, true, 0x0f, false, true, false, true},
		{0x6f, true, true, true, false, 0x69, false, true, false, false},
		{0x6f, true, true, true, true, 0x09, false, true, false, true},
		{0x70, false, false, false, false, 0x70, false, false, false, false},
		{0x70, false, false, false, true, 0xd0, false, false, false, true},
		{0x70, false, false, true, false, 0x76, false, false, false, false},
		{0x70, false, false, true, true, 0xd6, false, false, false, true},
		{0x70, false, true, false, false, 0x70, false, true, false, false},
		{0x70, false, true, false, true, 0x10, false, true, false, true},
		{0x70, false, true, true, false, 0x6a, false, true, false, false},
		{0x70, false, true, true, true, 0x0a, false, true, false, true},
		{0x70, true, false, false, false, 0x70, false, false, false, false},
		{0x70, true, false, false, true, 0xd0, false, false, false, true},
		{0x70, true, false, true, false, 0x76, false, false, false, false},
		{0x70, true, false, true, true, 0xd6, false, false, false, true},
		{0x70, true, true, false, false, 0x70, false, true, false, false},
		{0x70, true, true, false, true, 0x10, false, true, false, true},
		{0x70, true, true, true, false, 0x6a, false, true, false, false},
		{0x70, true, true, true, true, 0x0a, false, true, false, true},
		{0x71, false, false, false, false, 0x71, false, false, false, false},
		{0x71, false, false, false, true, 0xd1, false, false, false, true},
		{0x71, false, false, true, false, 0x77, false, false, false, false},
		{0x71, false, false, true, true, 0xd7, false, false, false, true},
		{0x71, false, true, false, false, 0x71, false, true, false, false},
		{0x71, false, true, false, true, 0x11, false, true, false, true},
		{0x71, false, true, true, false, 0x6b, false, true, false, false},
		{0x71, false, true, true, true, 0x0b, false, true, false, true},
		{0x71, true, false, false, false, 0x71, false, false, false, false},
		{0x71, true, false, false, true, 0xd1, false, false, false, true},
		{0x71, true, false, true, false, 0x77, false, false, false, false},
		{0x71, true, false, true, true, 0xd7, false, false, false, true},
		{0x71, true, true, false, false, 0x71, false, true, false, false},
		{0x71, true, true, false, true, 0x11, false, true, false, true},
		{0x71, true, true, true, false, 0x6b, false, true, false, false},
		{0x71, true, true, true, true, 0x0b, false, true, false, true},
		{0x72, false, false, false, false, 0x72, false, false, false, false},
		{0x72, false, false, false, true, 0xd2, false, false, false, true},
		{0x72, false, false, true, false, 0x78, false, false, false, false},
		{0x72, false, false, true, true, 0xd8, false, false, false, true},
		{0x72, false, true, false, false, 0x72, false, true, false, false},
		{0x72, false, true, false, true, 0x12, false, true, false, true},
		{0x72, false, true, true, false, 0x6c, false, true, false, false},
		{0x72, false, true, true, true, 0x0c, false, true, false, true},
		{0x72, true, false, false, false, 0x72, false, false, false, false},
		{0x72, true, false, false, true, 0xd2, false, false, false, true},
		{0x72, true, false, true, false, 0x78, false, false, false, false},
		{0x72, true, false, true, true, 0xd8, false, false, false, true},
		{0x72, true, true, false, false, 0x72, false, true, false, false},
		{0x72, true, true, false, true, 0x12, false, true, false, true},
		{0x72, true, true, true, false, 0x6c, false, true, false, false},
		{0x72, true, true, true, true, 0x0c, false, true, false, true},
		{0x73, false, false, false, false, 0x73, false, false, false, false},
		{0x73, false, false, false, true, 0xd3, false, false, false, true},
		{0x73, false, false, true, false, 0x79, false, false, false, false},
		{0x73, false, false, true, true, 0xd9, false, false, false, true},
		{0x73, false, true, false, false, 0x73, false, true, false, false},
		{0x73, false, true, false, true, 0x13, false, true, false, true},
		{0x73, false, true, true, false, 0x6d, false, true, false, false},
		{0x73, false, true, true, true, 0x0d, false, true, false, true},
		{0x73, true, false, false, false, 0x73, false, false, false, false},
		{0x73, true, false, false, true, 0xd3, false, false, false, true},
		{0x73, true, false, true, false, 0x79, false, false, false, false},
		{0x73, true, false, true, true, 0xd9, false, false, false, true},
		{0x73, true, true, false, false, 0x73, false, true, false, false},
		{0x73, true, true, false, true, 0x13, false, true, false, true},
		{0x73, true, true, true, false, 0x6d, false, true, false, false},
		{0x73, true, true, true, true, 0x0d, false, true, false, true},
		{0x74, false, false, false, false, 0x74, false, false, false, false},
		{0x74, false, false, false, true, 0xd4, false, false, false, true},
		{0x74, false, false, true, false, 0x7a, false, false, false, false},
		{0x74, false, false, true, true, 0xda, false, false, false, true},
		{0x74, false, true, false, false, 0x74, false, true, false, false},
		{0x74, false, true, false, true, 0x14, false, true, false, true},
		{0x74, false, true, true, false, 0x6e, false, true, false, false},
		{0x74, false, true, true, true, 0x0e, false, true, false, true},
		{0x74, true, false, false, false, 0x74, false, false, false, false},
		{0x74, true, false, false, true, 0xd4, false, false, false, true},
		{0x74, true, false, true, false, 0x7a, false, false, false, false},
		{0x74, true, false, true, true, 0xda, false, false, false, true},
		{0x74, true, true, false, false, 0x74, false, true, false, false},
		{0x74, true, true, false, true, 0x14, false, true, false, true},
		{0x74, true, true, true, false, 0x6e, false, true, false, false},
		{0x74, true, true, true, true, 0x0e, false, true, false, true},
		{0x75, false, false, false, false, 0x75, false, false, false, false},
		{0x75, false, false, false, true, 0xd5, false, false, false, true},
		{0x75, false, false, true, false, 0x7b, false, false, false, false},
		{0x75, false, false, true, true, 0xdb, false, false, false, true},
		{0x75, false, true, false, false, 0x75, false, true, false, false},
		{0x75, false, true, false, true, 0x15, false, true, false, true},
		{0x75, false, true, true, false, 0x6f, false, true, false, false},
		{0x75, false, true, true, true, 0x0f, false, true, false, true},
		{0x75, true, false, false, false, 0x75, false, false, false, false},
		{0x75, true, false, false, true, 0xd5, false, false, false, true},
		{0x75, true, false, true, false, 0x7b, false, false, false, false},
		{0x75, true, false, true, true, 0xdb, false, false, false, true},
		{0x75, true, true, false, false, 0x75, false, true, false, false},
		{0x75, true, true, false, true, 0x15, false, true, false, true},
		{0x75, true, true, true, false, 0x6f, false, true, false, false},
		{0x75, true, true, true, true, 0x0f, false, true, false, true},
		{0x76, false, false, false, false, 0x76, false, false, false, false},
		{0x76, false, false, false, true, 0xd6, false, false, false, true},
		{0x76, false, false, true, false, 0x7c, false, false, false, false},
		{0x76, false, false, true, true, 0xdc, false, false, false, true},
		{0x76, false, true, false, false, 0x76, false, true, false, false},
		{0x76, false, true, false, true, 0x16, false, true, false, true},
		{0x76, false, true, true, false, 0x70, false, true, false, false},
		{0x76, false, true, true, true, 0x10, false, true, false, true},
		{0x76, true, false, false, false, 0x76, false, false, false, false},
		{0x76, true, false, false, true, 0xd6, false, false, false, true},
		{0x76, true, false, true, false, 0x7c, false, false, false, false},
		{0x76, true, false, true, true, 0xdc, false, false, false, true},
		{0x76, true, true, false, false, 0x76, false, true, false, false},
		{0x76, true, true, false, true, 0x16, false, true, false, true},
		{0x76, true, true, true, false, 0x70, false, true, false, false},
		{0x76, true, true, true, true, 0x10, false, true, false, true},
		{0x77, false, false, false, false, 0x77, false, false, false, false},
		{0x77, false, false, false, true, 0xd7, false, false, false, true},
		{0x77, false, false, true, false, 0x7d, false, false, false, false},
		{0x77, false, false, true, true, 0xdd, false, false, false, true},
		{0x77, false, true, false, false, 0x77, false, true, false, false},
		{0x77, false, true, false, true, 0x17, false, true, false, true},
		{0x77, false, true, true, false, 0x71, false, true, false, false},
		{0x77, false, true, true, true, 0x11, false, true, false, true},
		{0x77, true, false, false, false, 0x77, false, false, false, false},
		{0x77, true, false, false, true, 0xd7, false, false, false, true},
		{0x77, true, false, true, false, 0x7d, false, false, false, false},
		{0x77, true, false, true, true, 0xdd, false, false, false, true},
		{0x77, true, true, false, false, 0x77, false, true, false, false},
		{0x77, true, true, false, true, 0x17, false, true, false, true},
		{0x77, true, true, true, false, 0x71, false, true, false, false},
		{0x77, true, true, true, true, 0x11, false, true, false, true},
		{0x78, false, false, false, false, 0x78, false, false, false, false},
		{0x78, false, false, false, true, 0xd8, false, false, false, true},
		{0x78, false, false, true, false, 0x7e, false, false, false, false},
		{0x78, false, false, true, true, 0xde, false, false, false, true},
		{0x78, false, true, false, false, 0x78, false, true, false, false},
		{0x78, false, true, false, true, 0x18, false, true, false, true},
		{0x78, false, true, true, false, 0x72, false, true, false, false},
		{0x78, false, true, true, true, 0x12, false, true, false, true},
		{0x78, true, false, false, false, 0x78, false, false, false, false},
		{0x78, true, false, false, true, 0xd8, false, false, false, true},
		{0x78, true, false, true, false, 0x7e, false, false, false, false},
		{0x78, true, false, true, true, 0xde, false, false, false, true},
		{0x78, true, true, false, false, 0x78, false, true, false, false},
		{0x78, true, true, false, true, 0x18, false, true, false, true},
		{0x78, true, true, true, false, 0x72, false, true, false, false},
		{0x78, true, true, true, true, 0x12, false, true, false, true},
		{0x79, false, false, false, false, 0x79, false, false, false, false},
		{0x79, false, false, false, true, 0xd9, false, false, false, true},
		{0x79, false, false, true, false, 0x7f, false, false, false, false},
		{0x79, false, false, true, true, 0xdf, false, false, false, true},
		{0x79, false, true, false, false, 0x79, false, true, false, false},
		{0x79, false, true, false, true, 0x19, false, true, false, true},
		{0x79, false, true, true, false, 0x73, false, true, false, false},
		{0x79, false, true, true, true, 0x13, false, true, false, true},
		{0x79, true, false, false, false, 0x79, false, false, false, false},
		{0x79, true, false, false, true, 0xd9, false, false, false, true},
		{0x79, true, false, true, false, 0x7f, false, false, false, false},
		{0x79, true, false, true, true, 0xdf, false, false, false, true},
		{0x79, true, true, false, false, 0x79, false, true, false, false},
		{0x79, true, true, false, true, 0x19, false, true, false, true},
		{0x79, true, true, true, false, 0x73, false, true, false, false},
		{0x79, true, true, true, true, 0x13, false, true, false, true},
		{0x7a, false, false, false, false, 0x80, false, false, false, false},
		{0x7a, false, false, false, true, 0xe0, false, false, false, true},
		{0x7a, false, false, true, false, 0x80, false, false, false, false},
		{0x7a, false, false, true, true, 0xe0, false, false, false, true},
		{0x7a, false, true, false, false, 0x7a, false, true, false, false},
		{0x7a, false, true, false, true, 0x1a, false, true, false, true},
		{0x7a, false, true, true, false, 0x74, false, true, false, false},
		{0x7a, false, true, true, true, 0x14, false, true, false, true},
		{0x7a, true, false, false, false, 0x80, false, false, false, false},
		{0x7a, true, false, false, true, 0xe0, false, false, false, true},
		{0x7a, true, false, true, false, 0x80, false, false, false, false},
		{0x7a, true, false, true, true, 0xe0, false, false, false, true},
		{0x7a, true, true, false, false, 0x7a, false, true, false, false},
		{0x7a, true, true, false, true, 0x1a, false, true, false, true},
		{0x7a, true, true, true, false, 0x74, false, true, false, false},
		{0x7a, true, true, true, true, 0x14, false, true, false, true},
		{0x7b, false, false, false, false, 0x81, false, false, false, false},
		{0x7b, false, false, false, true, 0xe1, false, false, false, true},
		{0x7b, false, false, true, false, 0x81, false, false, false, false},
		{0x7b, false, false, true, true, 0xe1, false, false, false, true},
		{0x7b, false, true, false, false, 0x7b, false, true, false, false},
		{0x7b, false, true, false, true, 0x1b, false, true, false, true},
		{0x7b, false, true, true, false, 0x75, false, true, false, false},
		{0x7b, false, true, true, true, 0x15, false, true, false, true},
		{0x7b, true, false, false, false, 0x81, false, false, false, false},
		{0x7b, true, false, false, true, 0xe1, false, false, false, true},
		{0x7b, true, false, true, false, 0x81, false, false, false, false},
		{0x7b, true, false, true, true, 0xe1, false, false, false, true},
		{0x7b, true, true, false, false, 0x7b, false, true, false, false},
		{0x7b, true, true, false, true, 0x1b, false, true, false, true},
		{0x7b, true, true, true, false, 0x75, false, true, false, false},
		{0x7b, true, true, true, true, 0x15, false, true, false, true},
		{0x7c, false, false, false, false, 0x82, false, false, false, false},
		{0x7c, false, false, false, true, 0xe2, false, false, false, true},
		{0x7c, false, false, true, false, 0x82, false, false, false, false},
		{0x7c, false, false, true, true, 0xe2, false, false, false, true},
		{0x7c, false, true, false, false, 0x7c, false, true, false, false},
		{0x7c, false, true, false, true, 0x1c, false, true, false, true},
		{0x7c, false, true, true, false, 0x76, false, true, false, false},
		{0x7c, false, true, true, true, 0x16, false, true, false, true},
		{0x7c, true, false, false, false, 0x82, false, false, false, false},
		{0x7c, true, false, false, true, 0xe2, false, false, false, true},
		{0x7c, true, false, true, false, 0x82, false, false, false, false},
		{0x7c, true, false, true, true, 0xe2, false, false, false, true},
		{0x7c, true, true, false, false, 0x7c, false, true, false, false},
		{0x7c, true, true, false, true, 0x1c, false, true, false, true},
		{0x7c, true, true, true, false, 0x76, false, true, false, false},
		{0x7c, true, true, true, true, 0x16, false, true, false, true},
		{0x7d, false, false, false, false, 0x83, false, false, false, false},
		{0x7d, false, false, false, true, 0xe3, false, false, false, true},
		{0x7d, false, false, true, false, 0x83, false, false, false, false},
		{0x7d, false, false, true, true, 0xe3, false, false, false, true},
		{0x7d, false, true, false, false, 0x7d, false, true, false, false},
		{0x7d, false, true, false, true, 0x1d, false, true, false, true},
		{0x7d, false, true, true, false, 0x77, false, true, false, false},
		{0x7d, false, true, true, true, 0x17, false, true, false, true},
		{0x7d, true, false, false, false, 0x83, false, false, false, false},
		{0x7d, true, false, false, true, 0xe3, false, false, false, true},
		{0x7d, true, false, true, false, 0x83, false, false, false, false},
		{0x7d, true, false, true, true, 0xe3, false, false, false, true},
		{0x7d, true, true, false, false, 0x7d, false, true, false, false},
		{0x7d, true, true, false, true, 0x1d, false, true, false, true},
		{0x7d, true, true, true, false, 0x77, false, true, false, false},
		{0x7d, true, true, true, true, 0x17, false, true, false, true},
		{0x7e, false, false, false, false, 0x84, false, false, false, false},
		{0x7e, false, false, false, true, 0xe4, false, false, false, true},
		{0x7e, false, false, true, false, 0x84, false, false, false, false},
		{0x7e, false, false, true, true, 0xe4, false, false, false, true},
		{0x7e, false, true, false, false, 0x7e, false, true, false, false},
		{0x7e, false, true, false, true, 0x1e, false, true, false, true},
		{0x7e, false, true, true, false, 0x78, false, true, false, false},
		{0x7e, false, true, true, true, 0x18, false, true, false, true},
		{0x7e, true, false, false, false, 0x84, false, false, false, false},
		{0x7e, true, false, false, true, 0xe4, false, false, false, true},
		{0x7e, true, false, true, false, 0x84, false, false, false, false},
		{0x7e, true, false, true, true, 0xe4, false, false, false, true},
		{0x7e, true, true, false, false, 0x7e, false, true, false, false},
		{0x7e, true, true, false, true, 0x1e, false, true, false, true},
		{0x7e, true, true, true, false, 0x78, false, true, false, false},
		{0x7e, true, true, true, true, 0x18, false, true, false, true},
		{0x7f, false, false, false, false, 0x85, false, false, false, false},
		{0x7f, false, false, false, true, 0xe5, false, false, false, true},
		{0x7f, false, false, true, false, 0x85, false, false, false, false},
		{0x7f, false, false, true, true, 0xe5, false, false, false, true},
		{0x7f, false, true, false, false, 0x7f, false, true, false, false},
		{0x7f, false, true, false, true, 0x1f, false, true, false, true},
		{0x7f, false, true, true, false, 0x79, false, true, false, false},
		{0x7f, false, true, true, true, 0x19, false, true, false, true},
		{0x7f, true, false, false, false, 0x85, false, false, false, false},
		{0x7f, true, false, false, true, 0xe5, false, false, false, true},
		{0x7f, true, false, true, false, 0x85, false, false, false, false},
		{0x7f, true, false, true, true, 0xe5, false, false, false, true},
		{0x7f, true, true, false, false, 0x7f, false, true, false, false},
		{0x7f, true, true, false, true, 0x1f, false, true, false, true},
		{0x7f, true, true, true, false, 0x79, false, true, false, false},
		{0x7f, true, true, true, true, 0x19, false, true, false, true},
		{0x80, false, false, false, false, 0x80, false, false, false, false},
		{0x80, false, false, false, true, 0xe0, false, false, false, true},
		{0x80, false, false, true, false, 0x86, false, false, false, false},
		{0x80, false, false, true, true, 0xe6, false, false, false, true},
		{0x80, false, true, false, false, 0x80, false, true, false, false},
		{0x80, false, true, false, true, 0x20, false, true, false, true},
		{0x80, false, true, true, false, 0x7a, false, true, false, false},
		{0x80, false, true, true, true, 0x1a, false, true, false, true},
		{0x80, true, false, false, false, 0x80, false, false, false, false},
		{0x80, true, false, false, true, 0xe0, false, false, false, true},
		{0x80, true, false, true, false, 0x86, false, false, false, false},
		{0x80, true, false, true, true, 0xe6, false, false, false, true},
		{0x80, true, true, false, false, 0x80, false, true, false, false},
		{0x80, true, true, false, true, 0x20, false, true, false, true},
		{0x80, true, true, true, false, 0x7a, false, true, false, false},
		{0x80, true, true, true, true, 0x1a, false, true, false, true},
		{0x81, false, false, false, false, 0x81, false, false, false, false},
		{0x81, false, false, false, true, 0xe1, false, false, false, true},
		{0x81, false, false, true, false, 0x87, false, false, false, false},
		{0x81, false, false, true, true, 0xe7, false, false, false, true},
		{0x81, false, true, false, false, 0x81, false, true, false, false},
		{0x81, false, true, false, true, 0x21, false, true, false, true},
		{0x81, false, true, true, false, 0x7b, false, true, false, false},
		{0x81, false, true, true, true, 0x1b, false, true, false, true},
		{0x81, true, false, false, false, 0x81, false, false, false, false},
		{0x81, true, false, false, true, 0xe1, false, false, false, true},
		{0x81, true, false, true, false, 0x87, false, false, false, false},
		{0x81, true, false, true, true, 0xe7, false, false, false, true},
		{0x81, true, true, false, false, 0x81, false, true, false, false},
		{0x81, true, true, false, true, 0x21, false, true, false, true},
		{0x81, true, true, true, false, 0x7b, false, true, false, false},
		{0x81, true, true, true, true, 0x1b, false, true, false, true},
		{0x82, false, false, false, false, 0x82, false, false, false, false},
		{0x82, false, false, false, true, 0xe2, false, false, false, true},
		{0x82, false, false, true, false, 0x88, false, false, false, false},
		{0x82, false, false, true, true, 0xe8, false, false, false, true},
		{0x82, false, true, false, false, 0x82, false, true, false, false},
		{0x82, false, true, false, true, 0x22, false, true, false, true},
		{0x82, false, true, true, false, 0x7c, false, true, false, false},
		{0x82, false, true, true, true, 0x1c, false, true, false, true},
		{0x82, true, false, false, false, 0x82, false, false, false, false},
		{0x82, true, false, false, true, 0xe2, false, false, false, true},
		{0x82, true, false, true, false, 0x88, false, false, false, false},
		{0x82, true, false, true, true, 0xe8, false, false, false, true},
		{0x82, true, true, false, false, 0x82, false, true, false, false},
		{0x82, true, true, false, true, 0x22, false, true, false, true},
		{0x82, true, true, true, false, 0x7c, false, true, false, false},
		{0x82, true, true, true, true, 0x1c, false, true, false, true},
		{0x83, false, false, false, false, 0x83, false, false, false, false},
		{0x83, false, false, false, true, 0xe3, false, false, false, true},
		{0x83, false, false, true, false, 0x89, false, false, false, false},
		{0x83, false, false, true, true, 0xe9, false, false, false, true},
		{0x83, false, true, false, false, 0x83, false, true, false, false},
		{0x83, false, true, false, true, 0x23, false, true, false, true},
		{0x83, false, true, true, false, 0x7d, false, true, false, false},
		{0x83, false, true, true, true, 0x1d, false, true, false, true},
		{0x83, true, false, false, false, 0x83, false, false, false, false},
		{0x83, true, false, false, true, 0xe3, false, false, false, true},
		{0x83, true, false, true, false, 0x89, false, false, false, false},
		{0x83, true, false, true, true, 0xe9, false, false, false, true},
		{0x83, true, true, false, false, 0x83, false, true, false, false},
		{0x83, true, true, false, true, 0x23, false, true, false, true},
		{0x83, true, true, true, false, 0x7d, false, true, false, false},
		{0x83, true, true, true, true, 0x1d, false, true, false, true},
		{0x84, false, false, false, false, 0x84, false, false, false, false},
		{0x84, false, false, false, true, 0xe4, false, false, false, true},
		{0x84, false, false, true, false, 0x8a, false, false, false, false},
		{0x84, false, false, true, true, 0xea, false, false, false, true},
		{0x84, false, true, false, false, 0x84, false, true, false, false},
		{0x84, false, true, false, true, 0x24, false, true, false, true},
		{0x84, false, true, true, false, 0x7e, false, true, false, false},
		{0x84, false, true, true, true, 0x1e, false, true, false, true},
		{0x84, true, false, false, false, 0x84, false, false, false, false},
		{0x84, true, false, false, true, 0xe4, false, false, false, true},
		{0x84, true, false, true, false, 0x8a, false, false, false, false},
		{0x84, true, false, true, true, 0xea, false, false, false, true},
		{0x84, true, true, false, false, 0x84, false, true, false, false},
		{0x84, true, true, false, true, 0x24, false, true, false, true},
		{0x84, true, true, true, false, 0x7e, false, true, false, false},
		{0x84, true, true, true, true, 0x1e, false, true, false, true},
		{0x85, false, false, false, false, 0x85, false, false, false, false},
		{0x85, false, false, false, true, 0xe5, false, false, false, true},
		{0x85, false, false, true, false, 0x8b, false, false, false, false},
		{0x85, false, false, true, true, 0xeb, false, false, false, true},
		{0x85, false, true, false, false, 0x85, false, true, false, false},
		{0x85, false, true, false, true, 0x25, false, true, false, true},
		{0x85, false, true, true, false, 0x7f, false, true, false, false},
		{0x85, false, true, true, true, 0x1f, false, true, false, true},
		{0x85, true, false, false, false, 0x85, false, false, false, false},
		{0x85, true, false, false, true, 0xe5, false, false, false, true},
		{0x85, true, false, true, false, 0x8b, false, false, false, false},
		{0x85, true, false, true, true, 0xeb, false, false, false, true},
		{0x85, true, true, false, false, 0x85, false, true, false, false},
		{0x85, true, true, false, true, 0x25, false, true, false, true},
		{0x85, true, true, true, false, 0x7f, false, true, false, false},
		{0x85, true, true, true, true, 0x1f, false, true, false, true},
		{0x86, false, false, false, false, 0x86, false, false, false, false},
		{0x86, false, false, false, true, 0xe6, false, false, false, true},
		{0x86, false, false, true, false, 0x8c, false, false, false, false},
		{0x86, false, false, true, true, 0xec, false, false, false, true},
		{0x86, false, true, false, false, 0x86, false, true, false, false},
		{0x86, false, true, false, true, 0x26, false, true, false, true},
		{0x86, false, true, true, false, 0x80, false, true, false, false},
		{0x86, false, true, true, true, 0x20, false, true, false, true},
		{0x86, true, false, false, false, 0x86, false, false, false, false},
		{0x86, true, false, false, true, 0xe6, false, false, false, true},
		{0x86, true, false, true, false, 0x8c, false, false, false, false},
		{0x86, true, false, true, true, 0xec, false, false, false, true},
		{0x86, true, true, false, false, 0x86, false, true, false, false},
		{0x86, true, true, false, true, 0x26, false, true, false, true},
		{0x86, true, true, true, false, 0x80, false, true, false, false},
		{0x86, true, true, true, true, 0x20, false, true, false, true},
		{0x87, false, false, false, false, 0x87, false, false, false, false},
		{0x87, false, false, false, true, 0xe7, false, false, false, true},
		{0x87, false, false, true, false, 0x8d, false, false, false, false},
		{0x87, false, false, true, true, 0xed, false, false, false, true},
		{0x87, false, true, false, false, 0x87, false, true, false, false},
		{0x87, false, true, false, true, 0x27, false, true, false, true},
		{0x87, false, true, true, false, 0x81, false, true, false, false},
		{0x87, false, true, true, true, 0x21, false, true, false, true},
		{0x87, true, false, false, false, 0x87, false, false, false, false},
		{0x87, true, false, false, true, 0xe7, false, false, false, true},
		{0x87, true, false, true, false, 0x8d, false, false, false, false},
		{0x87, true, false, true, true, 0xed, false, false, false, true},
		{0x87, true, true, false, false, 0x87, false, true, false, false},
		{0x87, true, true, false, true, 0x27, false, true, false, true},
		{0x87, true, true, true, false, 0x81, false, true, false, false},
		{0x87, true, true, true, true, 0x21, false, true, false, true},
		{0x88, false, false, false, false, 0x88, false, false, false, false},
		{0x88, false, false, false, true, 0xe8, false, false, false, true},
		{0x88, false, false, true, false, 0x8e, false, false, false, false},
		{0x88, false, false, true, true, 0xee, false, false, false, true},
		{0x88, false, true, false, false, 0x88, false, true, false, false},
		{0x88, false, true, false, true, 0x28, false, true, false, true},
		{0x88, false, true, true, false, 0x82, false, true, false, false},
		{0x88, false, true, true, true, 0x22, false, true, false, true},
		{0x88, true, false, false, false, 0x88, false, false, false, false},
		{0x88, true, false, false, true, 0xe8, false, false, false, true},
		{0x88, true, false, true, false, 0x8e, false, false, false, false},
		{0x88, true, false, true, true, 0xee, false, false, false, true},
		{0x88, true, true, false, false, 0x88, false, true, false, false},
		{0x88, true, true, false, true, 0x28, false, true, false, true},
		{0x88, true, true, true, false, 0x82, false, true, false, false},
		{0x88, true, true, true, true, 0x22, false, true, false, true},
		{0x89, false, false, false, false, 0x89, false, false, false, false},
		{0x89, false, false, false, true, 0xe9, false, false, false, true},
		{0x89, false, false, true, false, 0x8f, false, false, false, false},
		{0x89, false, false, true, true, 0xef, false, false, false, true},
		{0x89, false, true, false, false, 0x89, false, true, false, false},
		{0x89, false, true, false, true, 0x29, false, true, false, true},
		{0x89, false, true, true, false, 0x83, false, true, false, false},
		{0x89, false, true, true, true, 0x23, false, true, false, true},
		{0x89, true, false, false, false, 0x89, false, false, false, false},
		{0x89, true, false, false, true, 0xe9, false, false, false, true},
		{0x89, true, false, true, false, 0x8f, false, false, false, false},
		{0x89, true, false, true, true, 0xef, false, false, false, true},
		{0x89, true, true, false, false, 0x89, false, true, false, false},
		{0x89, true, true, false, true, 0x29, false, true, false, true},
		{0x89, true, true, true, false, 0x83, false, true, false, false},
		{0x89, true, true, true, true, 0x23, false, true, false, true},
		{0x8a, false, false, false, false, 0x90, false, false, false, false},
		{0x8a, false, false, false, true, 0xf0, false, false, false, true},
		{0x8a, false, false, true, false, 0x90, false, false, false, false},
		{0x8a, false, false, true, true, 0xf0, false, false, false, true},
		{0x8a, false, true, false, false, 0x8a, false, true, false, false},
		{0x8a, false, true, false, true, 0x2a, false, true, false, true},
		{0x8a, false, true, true, false, 0x84, false, true, false, false},
		{0x8a, false, true, true, true, 0x24, false, true, false, true},
		{0x8a, true, false, false, false, 0x90, false, false, false, false},
		{0x8a, true, false, false, true, 0xf0, false, false, false, true},
		{0x8a, true, false, true, false, 0x90, false, false, false, false},
		{0x8a, true, false, true, true, 0xf0, false, false, false, true},
		{0x8a, true, true, false, false, 0x8a, false, true, false, false},
		{0x8a, true, true, false, true, 0x2a, false, true, false, true},
		{0x8a, true, true, true, false, 0x84, false, true, false, false},
		{0x8a, true, true, true, true, 0x24, false, true, false, true},
		{0x8b, false, false, false, false, 0x91, false, false, false, false},
		{0x8b, false, false, false, true, 0xf1, false, false, false, true},
		{0x8b, false, false, true, false, 0x91, false, false, false, false},
		{0x8b, false, false, true, true, 0xf1, false, false, false, true},
		{0x8b, false, true, false, false, 0x8b, false, true, false, false},
		{0x8b, false, true, false, true, 0x2b, false, true, false, true},
		{0x8b, false, true, true, false, 0x85, false, true, false, false},
		{0x8b, false, true, true, true, 0x25, false, true, false, true},
		{0x8b, true, false, false, false, 0x91, false, false, false, false},
		{0x8b, true, false, false, true, 0xf1, false, false, false, true},
		{0x8b, true, false, true, false, 0x91, false, false, false, false},
		{0x8b, true, false, true, true, 0xf1, false, false, false, true},
		{0x8b, true, true, false, false, 0x8b, false, true, false, false},
		{0x8b, true, true, false, true, 0x2b, false, true, false, true},
		{0x8b, true, true, true, false, 0x85, false, true, false, false},
		{0x8b, true, true, true, true, 0x25, false, true, false, true},
		{0x8c, false, false, false, false, 0x92, false, false, false, false},
		{0x8c, false, false, false, true, 0xf2, false, false, false, true},
		{0x8c, false, false, true, false, 0x92, false, false, false, false},
		{0x8c, false, false, true, true, 0xf2, false, false, false, true},
		{0x8c, false, true, false, false, 0x8c, false, true, false, false},
		{0x8c, false, true, false, true, 0x2c, false, true, false, true},
		{0x8c, false, true, true, false, 0x86, false, true, false, false},
		{0x8c, false, true, true, true, 0x26, false, true, false, true},
		{0x8c, true, false, false, false, 0x92, false, false, false, false},
		{0x8c, true, false, false, true, 0xf2, false, false, false, true},
		{0x8c, true, false, true, false, 0x92, false, false, false, false},
		{0x8c, true, false, true, true, 0xf2, false, false, false, true},
		{0x8c, true, true, false, false, 0x8c, false, true, false, false},
		{0x8c, true, true, false, true, 0x2c, false, true, false, true},
		{0x8c, true, true, true, false, 0x86, false, true, false, false},
		{0x8c, true, true, true, true, 0x26, false, true, false, true},
		{0x8d, false, false, false, false, 0x93, false, false, false, false},
		{0x8d, false, false, false, true, 0xf3, false, false, false, true},
		{0x8d, false, false, true, false, 0x93, false, false, false, false},
		{0x8d, false, false, true, true, 0xf3, false, false, false, true},
		{0x8d, false, true, false, false, 0x8d, false, true, false, false},
		{0x8d, false, true, false, true, 0x2d, false, true, false, true},
		{0x8d, false, true, true, false, 0x87, false, true, false, false},
		{0x8d, false, true, true, true, 0x27, false, true, false, true},
		{0x8d, true, false, false, false, 0x93, false, false, false, false},
		{0x8d, true, false, false, true, 0xf3, false, false, false, true},
		{0x8d, true, false, true, false, 0x93, false, false, false, false},
		{0x8d, true, false, true, true, 0xf3, false, false, false, true},
		{0x8d, true, true, false, false, 0x8d, false, true, false, false},
		{0x8d, true, true, false, true, 0x2d, false, true, false, true},
		{0x8d, true, true, true, false, 0x87, false, true, false, false},
		{0x8d, true, true, true, true, 0x27, false, true, false, true},
		{0x8e, false, false, false, false, 0x94, false, false, false, false},
		{0x8e, false, false, false, true, 0xf4, false, false, false, true},
		{0x8e, false, false, true, false, 0x94, false, false, false, false},
		{0x8e, false, false, true, true, 0xf4, false, false, false, true},
		{0x8e, false, true, false, false, 0x8e, false, true, false, false},
		{0x8e, false, true, false, true, 0x2e, false, true, false, true},
		{0x8e, false, true, true, false, 0x88, false, true, false, false},
		{0x8e, false, true, true, true, 0x28, false, true, false, true},
		{0x8e, true, false, false, false, 0x94, false, false, false, false},
		{0x8e, true, false, false, true, 0xf4, false, false, false, true},
		{0x8e, true, false, true, false, 0x94, false, false, false, false},
		{0x8e, true, false, true, true, 0xf4, false, false, false, true},
		{0x8e, true, true, false, false, 0x8e, false, true, false, false},
		{0x8e, true, true, false, true, 0x2e, false, true, false, true},
		{0x8e, true, true, true, false, 0x88, false, true, false, false},
		{0x8e, true, true, true, true, 0x28, false, true, false, true},
		{0x8f, false, false, false, false, 0x95, false, false, false, false},
		{0x8f, false, false, false, true, 0xf5, false, false, false, true},
		{0x8f, false, false, true, false, 0x95, false, false, false, false},
		{0x8f, false, false, true, true, 0xf5, false, false, false, true},
		{0x8f, false, true, false, false, 0x8f, false, true, false, false},
		{0x8f, false, true, false, true, 0x2f, false, true, false, true},
		{0x8f, false, true, true, false, 0x89, false, true, false, false},
		{0x8f, false, true, true, true, 0x29, false, true, false, true},
		{0x8f, true, false, false, false, 0x95, false, false, false, false},
		{0x8f, true, false, false, true, 0xf5, false, false, false, true},
		{0x8f, true, false, true, false, 0x95, false, false, false, false},
		{0x8f, true, false, true, true, 0xf5, false, false, false, true},
		{0x8f, true, true, false, false, 0x8f, false, true, false, false},
		{0x8f, true, true, false, true, 0x2f, false, true, false, true},
		{0x8f, true, true, true, false, 0x89, false, true, false, false},
		{0x8f, true, true, true, true, 0x29, false, true, false, true},
		{0x90, false, false, false, false, 0x90, false, false, false, false},
		{0x90, false, false, false, true, 0xf0, false, false, false, true},
		{0x90, false, false, true, false, 0x96, false, false, false, false},
		{0x90, false, false, true, true, 0xf6, false, false, false, true},
		{0x90, false, true, false, false, 0x90, false, true, false, false},
		{0x90, false, true, false, true, 0x30, false, true, false, true},
		{0x90, false, true, true, false, 0x8a, false, true, false, false},
		{0x90, false, true, true, true, 0x2a, false, true, false, true},
		{0x90, true, false, false, false, 0x90, false, false, false, false},
		{0x90, true, false, false, true, 0xf0, false, false, false, true},
		{0x90, true, false, true, false, 0x96, false, false, false, false},
		{0x90, true, false, true, true, 0xf6, false, false, false, true},
		{0x90, true, true, false, false, 0x90, false, true, false, false},
		{0x90, true, true, false, true, 0x30, false, true, false, true},
		{0x90, true, true, true, false, 0x8a, false, true, false, false},
		{0x90, true, true, true, true, 0x2a, false, true, false, true},
		{0x91, false, false, false, false, 0x91, false, false, false, false},
		{0x91, false, false, false, true, 0xf1, false, false, false, true},
		{0x91, false, false, true, false, 0x97, false, false, false, false},
		{0x91, false, false, true, true, 0xf7, false, false, false, true},
		{0x91, false, true, false, false, 0x91, false, true, false, false},
		{0x91, false, true, false, true, 0x31, false, true, false, true},
		{0x91, false, true, true, false, 0x8b, false, true, false, false},
		{0x91, false, true, true, true, 0x2b, false, true, false, true},
		{0x91, true, false, false, false, 0x91, false, false, false, false},
		{0x91, true, false, false, true, 0xf1, false, false, false, true},
		{0x91, true, false, true, false, 0x97, false, false, false, false},
		{0x91, true, false, true, true, 0xf7, false, false, false, true},
		{0x91, true, true, false, false, 0x91, false, true, false, false},
		{0x91, true, true, false, true, 0x31, false, true, false, true},
		{0x91, true, true, true, false, 0x8b, false, true, false, false},
		{0x91, true, true, true, true, 0x2b, false, true, false, true},
		{0x92, false, false, false, false, 0x92, false, false, false, false},
		{0x92, false, false, false, true, 0xf2, false, false, false, true},
		{0x92, false, false, true, false, 0x98, false, false, false, false},
		{0x92, false, false, true, true, 0xf8, false, false, false, true},
		{0x92, false, true, false, false, 0x92, false, true, false, false},
		{0x92, false, true, false, true, 0x32, false, true, false, true},
		{0x92, false, true, true, false, 0x8c, false, true, false, false},
		{0x92, false, true, true, true, 0x2c, false, true, false, true},
		{0x92, true, false, false, false, 0x92, false, false, false, false},
		{0x92, true, false, false, true, 0xf2, false, false, false, true},
		{0x92, true, false, true, false, 0x98, false, false, false, false},
		{0x92, true, false, true, true, 0xf8, false, false, false, true},
		{0x92, true, true, false, false, 0x92, false, true, false, false},
		{0x92, true, true, false, true, 0x32, false, true, false, true},
		{0x92, true, true, true, false, 0x8c, false, true, false, false},
		{0x92, true, true, true, true, 0x2c, false, true, false, true},
		{0x93, false, false, false, false, 0x93, false, false, false, false},
		{0x93, false, false, false, true, 0xf3, false, false, false, true},
		{0x93, false, false, true, false, 0x99, false, false, false, false},
		{0x93, false, false, true, true, 0xf9, false, false, false, true},
		{0x93, false, true, false, false, 0x93, false, true, false, false},
		{0x93, false, true, false, true, 0x33, false, true, false, true},
		{0x93, false, true, true, false, 0x8d, false, true, false, false},
		{0x93, false, true, true, true, 0x2d, false, true, false, true},
		{0x93, true, false, false, false, 0x93, false, false, false, false},
		{0x93, true, false, false, true, 0xf3, false, false, false, true},
		{0x93, true, false, true, false, 0x99, false, false, false, false},
		{0x93, true, false, true, true, 0xf9, false, false, false, true},
		{0x93, true, true, false, false, 0x93, false, true, false, false},
		{0x93, true, true, false, true, 0x33, false, true, false, true},
		{0x93, true, true, true, false, 0x8d, false, true, false, false},
		{0x93, true, true, true, true, 0x2d, false, true, false, true},
		{0x94, false, false, false, false, 0x94, false, false, false, false},
		{0x94, false, false, false, true, 0xf4, false, false, false, true},
		{0x94, false, false, true, false, 0x9a, false, false, false, false},
		{0x94, false, false, true, true, 0xfa, false, false, false, true},
		{0x94, false, true, false, false, 0x94, false, true, false, false},
		{0x94, false, true, false, true, 0x34, false, true, false, true},
		{0x94, false, true, true, false, 0x8e, false, true, false, false},
		{0x94, false, true, true, true, 0x2e, false, true, false, true},
		{0x94, true, false, false, false, 0x94, false, false, false, false},
		{0x94, true, false, false, true, 0xf4, false, false, false, true},
		{0x94, true, false, true, false, 0x9a, false, false, false, false},
		{0x94, true, false, true, true, 0xfa, false, false, false, true},
		{0x94, true, true, false, false, 0x94, false, true, false, false},
		{0x94, true, true, false, true, 0x34, false, true, false, true},
		{0x94, true, true, true, false, 0x8e, false, true, false, false},
		{0x94, true, true, true, true, 0x2e, false, true, false, true},
		{0x95, false, false, false, false, 0x95, false, false, false, false},
		{0x95, false, false, false, true, 0xf5, false, false, false, true},
		{0x95, false, false, true, false, 0x9b, false, false, false, false},
		{0x95, false, false, true, true, 0xfb, false, false, false, true},
		{0x95, false, true, false, false, 0x95, false, true, false, false},
		{0x95, false, true, false, true, 0x35, false, true, false, true},
		{0x95, false, true, true, false, 0x8f, false, true, false, false},
		{0x95, false, true, true, true, 0x2f, false, true, false, true},
		{0x95, true, false, false, false, 0x95, false, false, false, false},
		{0x95, true, false, false, true, 0xf5, false, false, false, true},
		{0x95, true, false, true, false, 0x9b, false, false, false, false},
		{0x95, true, false, true, true, 0xfb, false, false, false, true},
		{0x95, true, true, false, false, 0x95, false, true, false, false},
		{0x95, true, true, false, true, 0x35, false, true, false, true},
		{0x95, true, true, true, false, 0x8f, false, true, false, false},
		{0x95, true, true, true, true, 0x2f, false, true, false, true},
		{0x96, false, false, false, false, 0x96, false, false, false, false},
		{0x96, false, false, false, true, 0xf6, false, false, false, true},
		{0x96, false, false, true, false, 0x9c, false, false, false, false},
		{0x96, false, false, true, true, 0xfc, false, false, false, true},
		{0x96, false, true, false, false, 0x96, false, true, false, false},
		{0x96, false, true, false, true, 0x36, false, true, false, true},
		{0x96, false, true, true, false, 0x90, false, true, false, false},
		{0x96, false, true, true, true, 0x30, false, true, false, true},
		{0x96, true, false, false, false, 0x96, false, false, false, false},
		{0x96, true, false, false, true, 0xf6, false, false, false, true},
		{0x96, true, false, true, false, 0x9c, false, false, false, false},
		{0x96, true, false, true, true, 0xfc, false, false, false, true},
		{0x96, true, true, false, false, 0x96, false, true, false, false},
		{0x96, true, true, false, true, 0x36, false, true, false, true},
		{0x96, true, true, true, false, 0x90, false, true, false, false},
		{0x96, true, true, true, true, 0x30, false, true, false, true},
		{0x97, false, false, false, false, 0x97, false, false, false, false},
		{0x97, false, false, false, true, 0xf7, false, false, false, true},
		{0x97, false, false, true, false, 0x9d, false, false, false, false},
		{0x97, false, false, true, true, 0xfd, false, false, false, true},
		{0x97, false, true, false, false, 0x97, false, true, false, false},
		{0x97, false, true, false, true, 0x37, false, true, false, true},
		{0x97, false, true, true, false, 0x91, false, true, false, false},
		{0x97, false, true, true, true, 0x31, false, true, false, true},
		{0x97, true, false, false, false, 0x97, false, false, false, false},
		{0x97, true, false, false, true, 0xf7, false, false, false, true},
		{0x97, true, false, true, false, 0x9d, false, false, false, false},
		{0x97, true, false, true, true, 0xfd, false, false, false, true},
		{0x97, true, true, false, false, 0x97, false, true, false, false},
		{0x97, true, true, false, true, 0x37, false, true, false, true},
		{0x97, true, true, true, false, 0x91, false, true, false, false},
		{0x97, true, true, true, true, 0x31, false, true, false, true},
		{0x98, false, false, false, false, 0x98, false, false, false, false},
		{0x98, false, false, false, true, 0xf8, false, false, false, true},
		{0x98, false, false, true, false, 0x9e, false, false, false, false},
		{0x98, false, false, true, true, 0xfe, false, false, false, true},
		{0x98, false, true, false, false, 0x98, false, true, false, false},
		{0x98, false, true, false, true, 0x38, false, true, false, true},
		{0x98, false, true, true, false, 0x92, false, true, false, false},
		{0x98, false, true, true, true, 0x32, false, true, false, true},
		{0x98, true, false, false, false, 0x98, false, false, false, false},
		{0x98, true, false, false, true, 0xf8, false, false, false, true},
		{0x98, true, false, true, false, 0x9e, false, false, false, false},
		{0x98, true, false, true, true, 0xfe, false, false, false, true},
		{0x98, true, true, false, false, 0x98, false, true, false, false},
		{0x98, true, true, false, true, 0x38, false, true, false, true},
		{0x98, true, true, true, false, 0x92, false, true, false, false},
		{0x98, true, true, true, true, 0x32, false, true, false, true},
		{0x99, false, false, false, false, 0x99, false, false, false, false},
		{0x99, false, false, false, true, 0xf9, false, false, false, true},
		{0x99, false, false, true, false, 0x9f, false, false, false, false},
		{0x99, false, false, true, true, 0xff, false, false, false, true},
		{0x99, false, true, false, false, 0x99, false, true, false, false},
		{0x99, false, true, false, true, 0x39, false, true, false, true},
		{0x99, false, true, true, false, 0x93, false, true, false, false},
		{0x99, false, true, true, true, 0x33, false, true, false, true},
		{0x99, true, false, false, false, 0x99, false, false, false, false},
		{0x99, true, false, false, true, 0xf9, false, false, false, true},
		{0x99, true, false, true, false, 0x9f, false, false, false, false},
		{0x99, true, false, true, true, 0xff, false, false, false, true},
		{0x99, true, true, false, false, 0x99, false, true, false, false},
		{0x99, true, true, false, true, 0x39, false, true, false, true},
		{0x99, true, true, true, false, 0x93, false, true, false, false},
		{0x99, true, true, true, true, 0x33, false, true, false, true},
		{0x9a, false, false, false, false, 0x00, true, false, false, true},
		{0x9a, false, false, false, true, 0x00, true, false, false, true},
		{0x9a, false, false, true, false, 0x00, true, false, false, true},
		{0x9a, false, false, true, true, 0x00, true, false, false, true},
		{0x9a, false, true, false, false, 0x9a, false, true, false, false},
		{0x9a, false, true, false, true, 0x3a, false, true, false, true},
		{0x9a, false, true, true, false, 0x94, false, true, false, false},
		{0x9a, false, true, true, true, 0x34, false, true, false, true},
		{0x9a, true, false, false, false, 0x00, true, false, false, true},
		{0x9a, true, false, false, true, 0x00, true, false, false, true},
		{0x9a, true, false, true, false, 0x00, true, false, false, true},
		{0x9a, true, false, true, true, 0x00, true, false, false, true},
		{0x9a, true, true, false, false, 0x9a, false, true, false, false},
		{0x9a, true, true, false, true, 0x3a, false, true, false, true},
		{0x9a, true, true, true, false, 0x94, false, true, false, false},
		{0x9a, true, true, true, true, 0x34, false, true, false, true},
		{0x9b, false, false, false, false, 0x01, false, false, false, true},
		{0x9b, false, false, false, true, 0x01, false, false, false, true},
		{0x9b, false, false, true, false, 0x01, false, false, false, true},
		{0x9b, false, false, true, true, 0x01, false, false, false, true},
		{0x9b, false, true, false, false, 0x9b, false, true, false, false},
		{0x9b, false, true, false, true, 0x3b, false, true, false, true},
		{0x9b, false, true, true, false, 0x95, false, true, false, false},
		{0x9b, false, true, true, true, 0x35, false, true, false, true},
		{0x9b, true, false, false, false, 0x01, false, false, false, true},
		{0x9b, true, false, false, true, 0x01, false, false, false, true},
		{0x9b, true, false, true, false, 0x01, false, false, false, true},
		{0x9b, true, false, true, true, 0x01, false, false, false, true},
		{0x9b, true, true, false, false, 0x9b, false, true, false, false},
		{0x9b, true, true, false, true, 0x3b, false, true, false, true},
		{0x9b, true, true, true, false, 0x95, false, true, false, false},
		{0x9b, true, true, true, true, 0x35, false, true, false, true},
		{0x9c, false, false, false, false, 0x02, false, false, false, true},
		{0x9c, false, false, false, true, 0x02, false, false, false, true},
		{0x9c, false, false, true, false, 0x02, false, false, false, true},
		{0x9c, false, false, true, true, 0x02, false, false, false, true},
		{0x9c, false, true, false, false, 0x9c, false, true, false, false},
		{0x9c, false, true, false, true, 0x3c, false, true, false, true},
		{0x9c, false, true, true, false, 0x96, false, true, false, false},
		{0x9c, false, true, true, true, 0x36, false, true, false, true},
		{0x9c, true, false, false, false, 0x02, false, false, false, true},
		{0x9c, true, false, false, true, 0x02, false, false, false, true},
		{0x9c, true, false, true, false, 0x02, false, false, false, true},
		{0x9c, true, false, true, true, 0x02, false, false, false, true},
		{0x9c, true, true, false, false, 0x9c, false, true, false, false},
		{0x9c, true, true, false, true, 0x3c, false, true, false, true},
		{0x9c, true, true, true, false, 0x96, false, true, false, false},
		{0x9c, true, true, true, true, 0x36, false, true, false, true},
		{0x9d, false, false, false, false, 0x03, false, false, false, true},
		{0x9d, false, false, false, true, 0x03, false, false, false, true},
		{0x9d, false, false, true, false, 0x03, false, false, false, true},
		{0x9d, false, false, true, true, 0x03, false, false, false, true},
		{0x9d, false, true, false, false, 0x9d, false, true, false, false},
		{0x9d, false, true, false, true, 0x3d, false, true, false, true},
		{0x9d, false, true, true, false, 0x97, false, true, false, false},
		{0x9d, false, true, true, true, 0x37, false, true, false, true},
		{0x9d, true, false, false, false, 0x03, false, false, false, true},
		{0x9d, true, false, false, true, 0x03, false, false, false, true},
		{0x9d, true, false, true, false, 0x03, false, false, false, true},
		{0x9d, true, false, true, true, 0x03, false, false, false, true},
		{0x9d, true, true, false, false, 0x9d, false, true, false, false},
		{0x9d, true, true, false, true, 0x3d, false, true, false, true},
		{0x9d, true, true, true, false, 0x97, false, true, false, false},
		{0x9d, true, true, true, true, 0x37, false, true, false, true},
		{0x9e, false, false, false, false, 0x04, false, false, false, true},
		{0x9e, false, false, false, true, 0x04, false, false, false, true},
		{0x9e, false, false, true, false, 0x04, false, false, false, true},
		{0x9e, false, false, true, true, 0x04, false, false, false, true},
		{0x9e, false, true, false, false, 0x9e, false, true, false, false},
		{0x9e, false, true, false, true, 0x3e, false, true, false, true},
		{0x9e, false, true, true, false, 0x98, false, true, false, false},
		{0x9e, false, true, true, true, 0x38, false, true, false, true},
		{0x9e, true, false, false, false, 0x04, false, false, false, true},
		{0x9e, true, false, false, true, 0x04, false, false, false, true},
		{0x9e, true, false, true, false, 0x04, false, false, false, true},
		{0x9e, true, false, true, true, 0x04, false, false, false, true},
		{0x9e, true, true, false, false, 0x9e, false, true, false, false},
		{0x9e, true, true, false, true, 0x3e, false, true, false, true},
		{0x9e, true, true, true, false, 0x98, false, true, false, false},
		{0x9e, true, true, true, true, 0x38, false, true, false, true},
		{0x9f, false, false, false, false, 0x05, false, false, false, true},
		{0x9f, false, false, false, true, 0x05, false, false, false, true},
		{0x9f, false, false, true, false, 0x05, false, false, false, true},
		{0x9f, false, false, true, true, 0x05, false, false, false, true},
		{0x9f, false, true, false, false, 0x9f, false, true, false, false},
		{0x9f, false, true, false, true, 0x3f, false, true, false, true},
		{0x9f, false, true, true, false, 0x99, false, true, false, false},
		{0x9f, false, true, true, true, 0x39, false, true, false, true},
		{0x9f, true, false, false, false, 0x05, false, false, false, true},
		{0x9f, true, false, false, true, 0x05, false, false, false, true},
		{0x9f, true, false, true, false, 0x05, false, false, false, true},
		{0x9f, true, false, true, true, 0x05, false, false, false, true},
		{0x9f, true, true, false, false, 0x9f, false, true, false, false},
		{0x9f, true, true, false, true, 0x3f, false, true, false, true},
		{0x9f, true, true, true, false, 0x99, false, true, false, false},
		{0x9f, true, true, true, true, 0x39, false, true, false, true},
		{0xa0, false, false, false, false, 0x00, true, false, false, true},
		{0xa0, false, false, false, true, 0x00, true, false, false, true},
		{0xa0, false, false, true, false, 0x06, false, false, false, true},
		{0xa0, false, false, true, true, 0x06, false, false, false, true},
		{0xa0, false, true, false, false, 0xa0, false, true, false, false},
		{0xa0, false, true, false, true, 0x40, false, true, false, true},
		{0xa0, false, true, true, false, 0x9a, false, true, false, false},
		{0xa0, false, true, true, true, 0x3a, false, true, false, true},
		{0xa0, true, false, false, false, 0x00, true, false, false, true},
		{0xa0, true, false, false, true, 0x00, true, false, false, true},
		{0xa0, true, false, true, false, 0x06, false, false, false, true},
		{0xa0, true, false, true, true, 0x06, false, false, false, true},
		{0xa0, true, true, false, false, 0xa0, false, true, false, false},
		{0xa0, true, true, false, true, 0x40, false, true, false, true},
		{0xa0, true, true, true, false, 0x9a, false, true, false, false},
		{0xa0, true, true, true, true, 0x3a, false, true, false, true},
		{0xa1, false, false, false, false, 0x01, false, false, false, true},
		{0xa1, false, false, false, true, 0x01, false, false, false, true},
		{0xa1, false, false, true, false, 0x07, false, false, false, true},
		{0xa1, false, false, true, true, 0x07, false, false, false, true},
		{0xa1, false, true, false, false, 0xa1, false, true, false, false},
		{0xa1, false, true, false, true, 0x41, false, true, false, true},
		{0xa1, false, true, true, false, 0x9b, false, true, false, false},
		{0xa1, false, true, true, true, 0x3b, false, true, false, true},
		{0xa1, true, false, false, false, 0x01, false, false, false, true},
		{0xa1, true, false, false, true, 0x01, false, false, false, true},
		{0xa1, true, false, true, false, 0x07, false, false, false, true},
		{0xa1, true, false, true, true, 0x07, false, false, false, true},
		{0xa1, true, true, false, false, 0xa1, false, true, false, false},
		{0xa1, true, true, false, true, 0x41, false, true, false, true},
		{0xa1, true, true, true, false, 0x9b, false, true, false, false},
		{0xa1, true, true, true, true, 0x3b, false, true, false, true},
		{0xa2, false, false, false, false, 0x02, false, false, false, true},
		{0xa2, false, false, false, true, 0x02, false, false, false, true},
		{0xa2, false, false, true, false, 0x08, false, false, false, true},
		{0xa2, false, false, true, true, 0x08, false, false, false, true},
		{0xa2, false, true, false, false, 0xa2, false, true, false, false},
		{0xa2, false, true, false, true, 0x42, false, true, false, true},
		{0xa2, false, true, true, false, 0x9c, false, true, false, false},
		{0xa2, false, true, true, true, 0x3c, false, true, false, true},
		{0xa2, true, false, false, false, 0x02, false, false, false, true},
		{0xa2, true, false, false, true, 0x02, false, false, false, true},
		{0xa2, true, false, true, false, 0x08, false, false, false, true},
		{0xa2, true, false, true, true, 0x08, false, false, false, true},
		{0xa2, true, true, false, false, 0xa2, false, true, false, false},
		{0xa2, true, true, false, true, 0x42, false, true, false, true},
		{0xa2, true, true, true, false, 0x9c, false, true, false, false},
		{0xa2, true, true, true, true, 0x3c, false, true, false, true},
		{0xa3, false, false, false, false, 0x03, false, false, false, true},
		{0xa3, false, false, false, true, 0x03, false, false, false, true},
		{0xa3, false, false, true, false, 0x09, false, false, false, true},
		{0xa3, false, false, true, true, 0x09, false, false, false, true},
		{0xa3, false, true, false, false, 0xa3, false, true, false, false},
		{0xa3, false, true, false, true, 0x43, false, true, false, true},
		{0xa3, false, true, true, false, 0x9d, false, true, false, false},
		{0xa3, false, true, true, true, 0x3d, false, true, false, true},
		{0xa3, true, false, false, false, 0x03, false, false, false, true},
		{0xa3, true, false, false, true, 0x03, false, false, false, true},
		{0xa3, true, false, true, false, 0x09, false, false, false, true},
		{0xa3, true, false, true, true, 0x09, false, false, false, true},
		{0xa3, true, true, false, false, 0xa3, false, true, false, false},
		{0xa3, true, true, false, true, 0x43, false, true, false, true},
		{0xa3, true, true, true, false, 0x9d, false, true, false, false},
		{0xa3, true, true, true, true, 0x3d, false, true, false, true},
		{0xa4, false, false, false, false, 0x04, false, false, false, true},
		{0xa4, false, false, false, true, 0x04, false, false, false, true},
		{0xa4, false, false, true, false, 0x0a, false, false, false, true},
		{0xa4, false, false, true, true, 0x0a, false, false, false, true},
		{0xa4, false, true, false, false, 0xa4, false, true, false, false},
		{0xa4, false, true, false, true, 0x44, false, true, false, true},
		{0xa4, false, true, true, false, 0x9e, false, true, false, false},
		{0xa4, false, true, true, true, 0x3e, false, true, false, true},
		{0xa4, true, false, false, false, 0x04, false, false, false, true},
		{0xa4, true, false, false, true, 0x04, false, false, false, true},
		{0xa4, true, false, true, false, 0x0a, false, false, false, true},
		{0xa4, true, false, true, true, 0x0a, false, false, false, true},
		{0xa4, true, true, false, false, 0xa4, false, true, false, false},
		{0xa4, true, true, false, true, 0x44, false, true, false, true},
		{0xa4, true, true, true, false, 0x9e, false, true, false, false},
		{0xa4, true, true, true, true, 0x3e, false, true, false, true},
		{0xa5, false, false, false, false, 0x05, false, false, false, true},
		{0xa5, false, false, false, true, 0x05, false, false, false, true},
		{0xa5, false, false, true, false, 0x0b, false, false, false, true},
		{0xa5, false, false, true, true, 0x0b, false, false, false, true},
		{0xa5, false, true, false, false, 0xa5, false, true, false, false},
		{0xa5, false, true, false, true, 0x45, false, true, false, true},
		{0xa5, false, true, true, false, 0x9f, false, true, false, false},
		{0xa5, false, true, true, true, 0x3f, false, true, false, true},
		{0xa5, true, false, false, false, 0x05, false, false, false, true},
		{0xa5, true, false, false, true, 0x05, false, false, false, true},
		{0xa5, true, false, true, false, 0x0b, false, false, false, true},
		{0xa5, true, false, true, true, 0x0b, false, false, false, true},
		{0xa5, true, true, false, false, 0xa5, false, true, false, false},
		{0xa5, true, true, false, true, 0x45, false, true, false, true},
		{0xa5, true, true, true, false, 0x9f, false, true, false, false},
		{0xa5, true, true, true, true, 0x3f, false, true, false, true},
		{0xa6, false, false, false, false, 0x06, false, false, false, true},
		{0xa6, false, false, false, true, 0x06, false, false, false, true},
		{0xa6, false, false, true, false, 0x0c, false, false, false, true},
		{0xa6, false, false, true, true, 0x0c, false, false, false, true},
		{0xa6, false, true, false, false, 0xa6, false, true, false, false},
		{0xa6, false, true, false, true, 0x46, false, true, false, true},
		{0xa6, false, true, true, false, 0xa0, false, true, false, false},
		{0xa6, false, true, true, true, 0x40, false, true, false, true},
		{0xa6, true, false, false, false, 0x06, false, false, false, true},
		{0xa6, true, false, false, true, 0x06, false, false, false, true},
		{0xa6, true, false, true, false, 0x0c, false, false, false, true},
		{0xa6, true, false, true, true, 0x0c, false, false, false, true},
		{0xa6, true, true, false, false, 0xa6, false, true, false, false},
		{0xa6, true, true, false, true, 0x46, false, true, false, true},
		{0xa6, true, true, true, false, 0xa0, false, true, false, false},
		{0xa6, true, true, true, true, 0x40, false, true, false, true},
		{0xa7, false, false, false, false, 0x07, false, false, false, true},
		{0xa7, false, false, false, true, 0x07, false, false, false, true},
		{0xa7, false, false, true, false, 0x0d, false, false, false, true},
		{0xa7, false, false, true, true, 0x0d, false, false, false, true},
		{0xa7, false, true, false, false, 0xa7, false, true, false, false},
		{0xa7, false, true, false, true, 0x47, false, true, false, true},
		{0xa7, false, true, true, false, 0xa1, false, true, false, false},
		{0xa7, false, true, true, true, 0x41, false, true, false, true},
		{0xa7, true, false, false, false, 0x07, false, false, false, true},
		{0xa7, true, false, false, true, 0x07, false, false, false, true},
		{0xa7, true, false, true, false, 0x0d, false, false, false, true},
		{0xa7, true, false, true, true, 0x0d, false, false, false, true},
		{0xa7, true, true, false, false, 0xa7, false, true, false, false},
		{0xa7, true, true, false, true, 0x47, false, true, false, true},
		{0xa7, true, true, true, false, 0xa1, false, true, false, false},
		{0xa7, true, true, true, true, 0x41, false, true, false, true},
		{0xa8, false, false, false, false, 0x08, false, false, false, true},
		{0xa8, false, false, false, true, 0x08, false, false, false, true},
		{0xa8, false, false, true, false, 0x0e, false, false, false, true},
		{0xa8, false, false, true, true, 0x0e, false, false, false, true},
		{0xa8, false, true, false, false, 0xa8, false, true, false, false},
		{0xa8, false, true, false, true, 0x48, false, true, false, true},
		{0xa8, false, true, true, false, 0xa2, false, true, false, false},
		{0xa8, false, true, true, true, 0x42, false, true, false, true},
		{0xa8, true, false, false, false, 0x08, false, false, false, true},
		{0xa8, true, false, false, true, 0x08, false, false, false, true},
		{0xa8, true, false, true, false, 0x0e, false, false, false, true},
		{0xa8, true, false, true, true, 0x0e, false, false, false, true},
		{0xa8, true, true, false, false, 0xa8, false, true, false, false},
		{0xa8, true, true, false, true, 0x48, false, true, false, true},
		{0xa8, true, true, true, false, 0xa2, false, true, false, false},
		{0xa8, true, true, true, true, 0x42, false, true, false, true},
		{0xa9, false, false, false, false, 0x09, false, false, false, true},
		{0xa9, false, false, false, true, 0x09, false, false, false, true},
		{0xa9, false, false, true, false, 0x0f, false, false, false, true},
		{0xa9, false, false, true, true, 0x0f, false, false, false, true},
		{0xa9, false, true, false, false, 0xa9, false, true, false, false},
		{0xa9, false, true, false, true, 0x49, false, true, false, true},
		{0xa9, false, true, true, false, 0xa3, false, true, false, false},
		{0xa9, false, true, true, true, 0x43, false, true, false, true},
		{0xa9, true, false, false, false, 0x09, false, false, false, true},
		{0xa9, true, false, false, true, 0x09, false, false, false, true},
		{0xa9, true, false, true, false, 0x0f, false, false, false, true},
		{0xa9, true, false, true, true, 0x0f, false, false, false, true},
		{0xa9, true, true, false, false, 0xa9, false, true, false, false},
		{0xa9, true, true, false, true, 0x49, false, true, false, true},
		{0xa9, true, true, true, false, 0xa3, false, true, false, false},
		{0xa9, true, true, true, true, 0x43, false, true, false, true},
		{0xaa, false, false, false, false, 0x10, false, false, false, true},
		{0xaa, false, false, false, true, 0x10, false, false, false, true},
		{0xaa, false, false, true, false, 0x10, false, false, false, true},
		{0xaa, false, false, true, true, 0x10, false, false, false, true},
		{0xaa, false, true, false, false, 0xaa, false, true, false, false},
		{0xaa, false, true, false, true, 0x4a, false, true, false, true},
		{0xaa, false, true, true, false, 0xa4, false, true, false, false},
		{0xaa, false, true, true, true, 0x44, false, true, false, true},
		{0xaa, true, false, false, false, 0x10, false, false, false, true},
		{0xaa, true, false, false, true, 0x10, false, false, false, true},
		{0xaa, true, false, true, false, 0x10, false, false, false, true},
		{0xaa, true, false, true, true, 0x10, false, false, false, true},
		{0xaa, true, true, false, false, 0xaa, false, true, false, false},
		{0xaa, true, true, false, true, 0x4a, false, true, false, true},
		{0xaa, true, true, true, false, 0xa4, false, true, false, false},
		{0xaa, true, true, true, true, 0x44, false, true, false, true},
		{0xab, false, false, false, false, 0x11, false, false, false, true},
		{0xab, false, false, false, true, 0x11, false, false, false, true},
		{0xab, false, false, true, false, 0x11, false, false, false, true},
		{0xab, false, false, true, true, 0x11, false, false, false, true},
		{0xab, false, true, false, false, 0xab, false, true, false, false},
		{0xab, false, true, false, true, 0x4b, false, true, false, true},
		{0xab, false, true, true, false, 0xa5, false, true, false, false},
		{0xab, false, true, true, true, 0x45, false, true, false, true},
		{0xab, true, false, false, false, 0x11, false, false, false, true},
		{0xab, true, false, false, true, 0x11, false, false, false, true},
		{0xab, true, false, true, false, 0x11, false, false, false, true},
		{0xab, true, false, true, true, 0x11, false, false, false, true},
		{0xab, true, true, false, false, 0xab, false, true, false, false},
		{0xab, true, true, false, true, 0x4b, false, true, false, true},
		{0xab, true, true, true, false, 0xa5, false, true, false, false},
		{0xab, true, true, true, true, 0x45, false, true, false, true},
		{0xac, false, false, false, false, 0x12, false, false, false, true},
		{0xac, false, false, false, true, 0x12, false, false, false, true},
		{0xac, false, false, true, false, 0x12, false, false, false, true},
		{0xac, false, false, true, true, 0x12, false, false, false, true},
		{0xac, false, true, false, false, 0xac, false, true, false, false},
		{0xac, false, true, false, true, 0x4c, false, true, false, true},
		{0xac, false, true, true, false, 0xa6, false, true, false, false},
		{0xac, false, true, true, true, 0x46, false, true, false, true},
		{0xac, true, false, false, false, 0x12, false, false, false, true},
		{0xac, true, false, false, true, 0x12, false, false, false, true},
		{0xac, true, false, true, false, 0x12, false, false, false, true},
		{0xac, true, false, true, true, 0x12, false, false, false, true},
		{0xac, true, true, false, false, 0xac, false, true, false, false},
		{0xac, true, true, false, true, 0x4c, false, true, false, true},
		{0xac, true, true, true, false, 0xa6, false, true, false, false},
		{0xac, true, true, true, true, 0x46, false, true, false, true},
		{0xad, false, false, false, false, 0x13, false, false, false, true},
		{0xad, false, false, false, true, 0x13, false, false, false, true},
		{0xad, false, false, true, false, 0x13, false, false, false, true},
		{0xad, false, false, true, true, 0x13, false, false, false, true},
		{0xad, false, true, false, false, 0xad, false, true, false, false},
		{0xad, false, true, false, true, 0x4d, false, true, false, true},
		{0xad, false, true, true, false, 0xa7, false, true, false, false},
		{0xad, false, true, true, true, 0x47, false, true, false, true},
		{0xad, true, false, false, false, 0x13, false, false, false, true},
		{0xad, true, false, false, true, 0x13, false, false, false, true},
		{0xad, true, false, true, false, 0x13, false, false, false, true},
		{0xad, true, false, true, true, 0x13, false, false, false, true},
		{0xad, true, true, false, false, 0xad, false, true, false, false},
		{0xad, true, true, false, true, 0x4d, false, true, false, true},
		{0xad, true, true, true, false, 0xa7, false, true, false, false},
		{0xad, true, true, true, true, 0x47, false, true, false, true},
		{0xae, false, false, false, false, 0x14, false, false, false, true},
		{0xae, false, false, false, true, 0x14, false, false, false, true},
		{0xae, false, false, true, false, 0x14, false, false, false, true},
		{0xae, false, false, true, true, 0x14, false, false, false, true},
		{0xae, false, true, false, false, 0xae, false, true, false, false},
		{0xae, false, true, false, true, 0x4e, false, true, false, true},
		{0xae, false, true, true, false, 0xa8, false, true, false, false},
		{0xae, false, true, true, true, 0x48, false, true, false, true},
		{0xae, true, false, false, false, 0x14, false, false, false, true},
		{0xae, true, false, false, true, 0x14, false, false, false, true},
		{0xae, true, false, true, false, 0x14, false, false, false, true},
		{0xae, true, false, true, true, 0x14, false, false, false, true},
		{0xae, true, true, false, false, 0xae, false, true, false, false},
		{0xae, true, true, false, true, 0x4e, false, true, false, true},
		{0xae, true, true, true, false, 0xa8, false, true, false, false},
		{0xae, true, true, true, true, 0x48, false, true, false, true},
		{0xaf, false, false, false, false, 0x15, false, false, false, true},
		{0xaf, false, false, false, true, 0x15, false, false, false, true},
		{0xaf, false, false, true, false, 0x15, false, false, false, true},
		{0xaf, false, false, true, true, 0x15, false, false, false, true},
		{0xaf, false, true, false, false, 0xaf, false, true, false, false},
		{0xaf, false, true, false, true, 0x4f, false, true, false, true},
		{0xaf, false, true, true, false, 0xa9, false, true, false, false},
		{0xaf, false, true, true, true, 0x49, false, true, false, true},
		{0xaf, true, false, false, false, 0x15, false, false, false, true},
		{0xaf, true, false, false, true, 0x15, false, false, false, true},
		{0xaf, true, false, true, false, 0x15, false, false, false, true},
		{0xaf, true, false, true, true, 0x15, false, false, false, true},
		{0xaf, true, true, false, false, 0xaf, false, true, false, false},
		{0xaf, true, true, false, true, 0x4f, false, true, false, true},
		{0xaf, true, true, true, false, 0xa9, false, true, false, false},
		{0xaf, true, true, true, true, 0x49, false, true, false, true},
		{0xb0, false, false, false, false, 0x10, false, false, false, true},
		{0xb0, false, false, false, true, 0x10, false, false, false, true},
		{0xb0, false, false, true, false, 0x16, false, false, false, true},
		{0xb0, false, false, true, true, 0x16, false, false, false, true},
		{0xb0, false, true, false, false, 0xb0, false, true, false, false},
		{0xb0, false, true, false, true, 0x50, false, true, false, true},
		{0xb0, false, true, true, false, 0xaa, false, true, false, false},
		{0xb0, false, true, true, true, 0x4a, false, true, false, true},
		{0xb0, true, false, false, false, 0x10, false, false, false, true},
		{0xb0, true, false, false, true, 0x10, false, false, false, true},
		{0xb0, true, false, true, false, 0x16, false, false, false, true},
		{0xb0, true, false, true, true, 0x16, false, false, false, true},
		{0xb0, true, true, false, false, 0xb0, false, true, false, false},
		{0xb0, true, true, false, true, 0x50, false, true, false, true},
		{0xb0, true, true, true, false, 0xaa, false, true, false, false},
		{0xb0, true, true, true, true, 0x4a, false, true, false, true},
		{0xb1, false, false, false, false, 0x11, false, false, false, true},
		{0xb1, false, false, false, true, 0x11, false, false, false, true},
		{0xb1, false, false, true, false, 0x17, false, false, false, true},
		{0xb1, false, false, true, true, 0x17, false, false, false, true},
		{0xb1, false, true, false, false, 0xb1, false, true, false, false},
		{0xb1, false, true, false, true, 0x51, false, true, false, true},
		{0xb1, false, true, true, false, 0xab, false, true, false, false},
		{0xb1, false, true, true, true, 0x4b, false, true, false, true},
		{0xb1, true, false, false, false, 0x11, false, false, false, true},
		{0xb1, true, false, false, true, 0x11, false, false, false, true},
		{0xb1, true, false, true, false, 0x17, false, false, false, true},
		{0xb1, true, false, true, true, 0x17, false, false, false, true},
		{0xb1, true, true, false, false, 0xb1, false, true, false, false},
		{0xb1, true, true, false, true, 0x51, false, true, false, true},
		{0xb1, true, true, true, false, 0xab, false, true, false, false},
		{0xb1, true, true, true, true, 0x4b, false, true, false, true},
		{0xb2, false, false, false, false, 0x12, false, false, false, true},
		{0xb2, false, false, false, true, 0x12, false, false, false, true},
		{0xb2, false, false, true, false, 0x18, false, false, false, true},
		{0xb2, false, false, true, true, 0x18, false, false, false, true},
		{0xb2, false, true, false, false, 0xb2, false, true, false, false},
		{0xb2, false, true, false, true, 0x52, false, true, false, true},
		{0xb2, false, true, true, false, 0xac, false, true, false, false},
		{0xb2, false, true, true, true, 0x4c, false, true, false, true},
		{0xb2, true, false, false, false, 0x12, false, false, false, true},
		{0xb2, true, false, false, true, 0x12, false, false, false, true},
		{0xb2, true, false, true, false, 0x18, false, false, false, true},
		{0xb2, true, false, true, true, 0x18, false, false, false, true},
		{0xb2, true, true, false, false, 0xb2, false, true, false, false},
		{0xb2, true, true, false, true, 0x52, false, true, false, true},
		{0xb2, true, true, true, false, 0xac, false, true, false, false},
		{0xb2, true, true, true, true, 0x4c, false, true, false, true},
		{0xb3, false, false, false, false, 0x13, false, false, false, true},
		{0xb3, false, false, false, true, 0x13, false, false, false, true},
		{0xb3, false, false, true, false, 0x19, false, false, false, true},
		{0xb3, false, false, true, true, 0x19, false, false, false, true},
		{0xb3, false, true, false, false, 0xb3, false, true, false, false},
		{0xb3, false, true, false, true, 0x53, false, true, false, true},
		{0xb3, false, true, true, false, 0xad, false, true, false, false},
		{0xb3, false, true, true, true, 0x4d, false, true, false, true},
		{0xb3, true, false, false, false, 0x13, false, false, false, true},
		{0xb3, true, false, false, true, 0x13, false, false, false, true},
		{0xb3, true, false, true, false, 0x19, false, false, false, true},
		{0xb3, true, false, true, true, 0x19, false, false, false, true},
		{0xb3, true, true, false, false, 0xb3, false, true, false, false},
		{0xb3, true, true, false, true, 0x53, false, true, false, true},
		{0xb3, true, true, true, false, 0xad, false, true, false, false},
		{0xb3, true, true, true, true, 0x4d, false, true, false, true},
		{0xb4, false, false, false, false, 0x14, false, false, false, true},
		{0xb4, false, false, false, true, 0x14, false, false, false, true},
		{0xb4, false, false, true, false, 0x1a, false, false, false, true},
		{0xb4, false, false, true, true, 0x1a, false, false, false, true},
		{0xb4, false, true, false, false, 0xb4, false, true, false, false},
		{0xb4, false, true, false, true, 0x54, false, true, false, true},
		{0xb4, false, true, true, false, 0xae, false, true, false, false},
		{0xb4, false, true, true, true, 0x4e, false, true, false, true},
		{0xb4, true, false, false, false, 0x14, false, false, false, true},
		{0xb4, true, false, false, true, 0x14, false, false, false, true},
		{0xb4, true, false, true, false, 0x1a, false, false, false, true},
		{0xb4, true, false, true, true, 0x1a, false, false, false, true},
		{0xb4, true, true, false, false, 0xb4, false, true, false, false},
		{0xb4, true, true, false, true, 0x54, false, true, false, true},
		{0xb4, true, true, true, false, 0xae, false, true, false, false},
		{0xb4, true, true, true, true, 0x4e, false, true, false, true},
		{0xb5, false, false, false, false, 0x15, false, false, false, true},
		{0xb5, false, false, false, true, 0x15, false, false, false, true},
		{0xb5, false, false, true, false, 0x1b, false, false, false, true},
		{0xb5, false, false, true, true, 0x1b, false, false, false, true},
		{0xb5, false, true, false, false, 0xb5, false, true, false, false},
		{0xb5, false, true, false, true, 0x55, false, true, false, true},
		{0xb5, false, true, true, false, 0xaf, false, true, false, false},
		{0xb5, false, true, true, true, 0x4f, false, true, false, true},
		{0xb5, true, false, false, false, 0x15, false, false, false, true},
		{0xb5, true, false, false, true, 0x15, false, false, false, true},
		{0xb5, true, false, true, false, 0x1b, false, false, false, true},
		{0xb5, true, false, true, true, 0x1b, false, false, false, true},
		{0xb5, true, true, false, false, 0xb5, false, true, false, false},
		{0xb5, true, true, false, true, 0x55, false, true, false, true},
		{0xb5, true, true, true, false, 0xaf, false, true, false, false},
		{0xb5, true, true, true, true, 0x4f, false, true, false, true},
		{0xb6, false, false, false, false, 0x16, false, false, false, true},
		{0xb6, false, false, false, true, 0x16, false, false, false, true},
		{0xb6, false, false, true, false, 0x1c, false, false, false, true},
		{0xb6, false, false, true, true, 0x1c, false, false, false, true},
		{0xb6, false, true, false, false, 0xb6, false, true, false, false},
		{0xb6, false, true, false, true, 0x56, false, true, false, true},
		{0xb6, false, true, true, false, 0xb0, false, true, false, false},
		{0xb6, false, true, true, true, 0x50, false, true, false, true},
		{0xb6, true, false, false, false, 0x16, false, false, false, true},
		{0xb6, true, false, false, true, 0x16, false, false, false, true},
		{0xb6, true, false, true, false, 0x1c, false, false, false, true},
		{0xb6, true, false, true, true, 0x1c, false, false, false, true},
		{0xb6, true, true, false, false, 0xb6, false, true, false, false},
		{0xb6, true, true, false, true, 0x56, false, true, false, true},
		{0xb6, true, true, true, false, 0xb0, false, true, false, false},
		{0xb6, true, true, true, true, 0x50, false, true, false, true},
		{0xb7, false, false, false, false, 0x17, false, false, false, true},
		{0xb7, false, false, false, true, 0x17, false, false, false, true},
		{0xb7, false, false, true, false, 0x1d, false, false, false, true},
		{0xb7, false, false, true, true, 0x1d, false, false, false, true},
		{0xb7, false, true, false, false, 0xb7, false, true, false, false},
		{0xb7, false, true, false, true, 0x57, false, true, false, true},
		{0xb7, false, true, true, false, 0xb1, false, true, false, false},
		{0xb7, false, true, true, true, 0x51, false, true, false, true},
		{0xb7, true, false, false, false, 0x17, false, false, false, true},
		{0xb7, true, false, false, true, 0x17, false, false, false, true},
		{0xb7, true, false, true, false, 0x1d, false, false, false, true},
		{0xb7, true, false, true, true, 0x1d, false, false, false, true},
		{0xb7, true, true, false, false, 0xb7, false, true, false, false},
		{0xb7, true, true, false, true, 0x57, false, true, false, true},
		{0xb7, true, true, true, false, 0xb1, false, true, false, false},
		{0xb7, true, true, true, true, 0x51, false, true, false, true},
		{0xb8, false, false, false, false, 0x18, false, false, false, true},
		{0xb8, false, false, false, true, 0x18, false, false, false, true},
		{0xb8, false, false, true, false, 0x1e, false, false, false, true},
		{0xb8, false, false, true, true, 0x1e, false, false, false, true},
		{0xb8, false, true, false, false, 0xb8, false, true, false, false},
		{0xb8, false, true, false, true, 0x58, false, true, false, true},
		{0xb8, false, true, true, false, 0xb2, false, true, false, false},
		{0xb8, false, true, true, true, 0x52, false, true, false, true},
		{0xb8, true, false, false, false, 0x18, false, false, false, true},
		{0xb8, true, false, false, true, 0x18, false, false, false, true},
		{0xb8, true, false, true, false, 0x1e, false, false, false, true},
		{0xb8, true, false, true, true, 0x1e, false, false, false, true},
		{0xb8, true, true, false, false, 0xb8, false, true, false, false},
		{0xb8, true, true, false, true, 0x58, false, true, false, true},
		{0xb8, true, true, true, false, 0xb2, false, true, false, false},
		{0xb8, true, true, true, true, 0x52, false, true, false, true},
		{0xb9, false, false, false, false, 0x19, false, false, false, true},
		{0xb9, false, false, false, true, 0x19, false, false, false, true},
		{0xb9, false, false, true, false, 0x1f, false, false, false, true},
		{0xb9, false, false, true, true, 0x1f, false, false, false, true},
		{0xb9, false, true, false, false, 0xb9, false, true, false, false},
		{0xb9, false, true, false, true, 0x59, false, true, false, true},
		{0xb9, false, true, true, false, 0xb3, false, true, false, false},
		{0xb9, false, true, true, true, 0x53, false, true, false, true},
		{0xb9, true, false, false, false, 0x19, false, false, false, true},
		{0xb9, true, false, false, true, 0x19, false, false, false, true},
		{0xb9, true, false, true, false, 0x1f, false, false, false, true},
		{0xb9, true, false, true, true, 0x1f, false, false, false, true},
		{0xb9, true, true, false, false, 0xb9, false, true, false, false},
		{0xb9, true, true, false, true, 0x59, false, true, false, true},
		{0xb9, true, true, true, false, 0xb3, false, true, false, false},
		{0xb9, true, true, true, true, 0x53, false, true, false, true},
		{0xba, false, false, false, false, 0x20, false, false, false, true},
		{0xba, false, false, false, true, 0x20, false, false, false, true},
		{0xba, false, false, true, false, 0x20, false, false, false, true},
		{0xba, false, false, true, true, 0x20, false, false, false, true},
		{0xba, false, true, false, false, 0xba, false, true, false, false},
		{0xba, false, true, false, true, 0x5a, false, true, false, true},
		{0xba, false, true, true, false, 0xb4, false, true, false, false},
		{0xba, false, true, true, true, 0x54, false, true, false, true},
		{0xba, true, false, false, false, 0x20, false, false, false, true},
		{0xba, true, false, false, true, 0x20, false, false, false, true},
		{0xba, true, false, true, false, 0x20, false, false, false, true},
		{0xba, true, false, true, true, 0x20, false, false, false, true},
		{0xba, true, true, false, false, 0xba, false, true, false, false},
		{0xba, true, true, false, true, 0x5a, false, true, false, true},
		{0xba, true, true, true, false, 0xb4, false, true, false, false},
		{0xba, true, true, true, true, 0x54, false, true, false, true},
		{0xbb, false, false, false, false, 0x21, false, false, false, true},
		{0xbb, false, false, false, true, 0x21, false, false, false, true},
		{0xbb, false, false, true, false, 0x21, false, false, false, true},
		{0xbb, false, false, true, true, 0x21, false, false, false, true},
		{0xbb, false, true, false, false, 0xbb, false, true, false, false},
		{0xbb, false, true, false, true, 0x5b, false, true, false, true},
		{0xbb, false, true, true, false, 0xb5, false, true, false, false},
		{0xbb, false, true, true, true, 0x55, false, true, false, true},
		{0xbb, true, false, false, false, 0x21, false, false, false, true},
		{0xbb, true, false, false, true, 0x21, false, false, false, true},
		{0xbb, true, false, true, false, 0x21, false, false, false, true},
		{0xbb, true, false, true, true, 0x21, false, false, false, true},
		{0xbb, true, true, false, false, 0xbb, false, true, false, false},
		{0xbb, true, true, false, true, 0x5b, false, true, false, true},
		{0xbb, true, true, true, false, 0xb5, false, true, false, false},
		{0xbb, true, true, true, true, 0x55, false, true, false, true},
		{0xbc, false, false, false, false, 0x22, false, false, false, true},
		{0xbc, false, false, false, true, 0x22, false, false, false, true},
		{0xbc, false, false, true, false, 0x22, false, false, false, true},
		{0xbc, false, false, true, true, 0x22, false, false, false, true},
		{0xbc, false, true, false, false, 0xbc, false, true, false, false},
		{0xbc, false, true, false, true, 0x5c, false, true, false, true},
		{0xbc, false, true, true, false, 0xb6, false, true, false, false},
		{0xbc, false, true, true, true, 0x56, false, true, false, true},
		{0xbc, true, false, false, false, 0x22, false, false, false, true},
		{0xbc, true, false, false, true, 0x22, false, false, false, true},
		{0xbc, true, false, true, false, 0x22, false, false, false, true},
		{0xbc, true, false, true, true, 0x22, false, false, false, true},
		{0xbc, true, true, false, false, 0xbc, false, true, false, false},
		{0xbc, true, true, false, true, 0x5c, false, true, false, true},
		{0xbc, true, true, true, false, 0xb6, false, true, false, false},
		{0xbc, true, true, true, true, 0x56, false, true, false, true},
		{0xbd, false, false, false, false, 0x23, false, false, false, true},
		{0xbd, false, false, false, true, 0x23, false, false, false, true},
		{0xbd, false, false, true, false, 0x23, false, false, false, true},
		{0xbd, false, false, true, true, 0x23, false, false, false, true},
		{0xbd, false, true, false, false, 0xbd, false, true, false, false},
		{0xbd, false, true, false, true, 0x5d, false, true, false, true},
		{0xbd, false, true, true, false, 0xb7, false, true, false, false},
		{0xbd, false, true, true, true, 0x57, false, true, false, true},
		{0xbd, true, false, false, false, 0x23, false, false, false, true},
		{0xbd, true, false, false, true, 0x23, false, false, false, true},
		{0xbd, true, false, true, false, 0x23, false, false, false, true},
		{0xbd, true, false, true, true, 0x23, false, false, false, true},
		{0xbd, true, true, false, false, 0xbd, false, true, false, false},
		{0xbd, true, true, false, true, 0x5d, false, true, false, true},
		{0xbd, true, true, true, false, 0xb7, false, true, false, false},
		{0xbd, true, true, true, true, 0x57, false, true, false, true},
		{0xbe, false, false, false, false, 0x24, false, false, false, true},
		{0xbe, false, false, false, true, 0x24, false, false, false, true},
		{0xbe, false, false, true, false, 0x24, false, false, false, true},
		{0xbe, false, false, true, true, 0x24, false, false, false, true},
		{0xbe, false, true, false, false, 0xbe, false, true, false, false},
		{0xbe, false, true, false, true, 0x5e, false, true, false, true},
		{0xbe, false, true, true, false, 0xb8, false, true, false, false},
		{0xbe, false, true, true, true, 0x58, false, true, false, true},
		{0xbe, true, false, false, false, 0x24, false, false, false, true},
		{0xbe, true, false, false, true, 0x24, false, false, false, true},
		{0xbe, true, false, true, false, 0x24, false, false, false, true},
		{0xbe, true, false, true, true, 0x24, false, false, false, true},
		{0xbe, true, true, false, false, 0xbe, false, true, false, false},
		{0xbe, true, true, false, true, 0x5e, false, true, false, true},
		{0xbe, true, true, true, false, 0xb8, false, true, false, false},
		{0xbe, true, true, true, true, 0x58, false, true, false, true},
		{0xbf, false, false, false, false, 0x25, false, false, false, true},
		{0xbf, false, false, false, true, 0x25, false, false, false, true},
		{0xbf, false, false, true, false, 0x25, false, false, false, true},
		{0xbf, false, false, true, true, 0x25, false, false, false, true},
		{0xbf, false, true, false, false, 0xbf, false, true, false, false},
		{0xbf, false, true, false, true, 0x5f, false, true, false, true},
		{0xbf, false, true, true, false, 0xb9, false, true, false, false},
		{0xbf, false, true, true, true, 0x59, false, true, false, true},
		{0xbf, true, false, false, false, 0x25, false, false, false, true},
		{0xbf, true, false, false, true, 0x25, false, false, false, true},
		{0xbf, true, false, true, false, 0x25, false, false, false, true},
		{0xbf, true, false, true, true, 0x25, false, false, false, true},
		{0xbf, true, true, false, false, 0xbf, false, true, false, false},
		{0xbf, true, true, false, true, 0x5f, false, true, false, true},
		{0xbf, true, true, true, false, 0xb9, false, true, false, false},
		{0xbf, true, true, true, true, 0x59, false, true, false, true},
		{0xc0, false, false, false, false, 0x20, false, false, false, true},
		{0xc0, false, false, false, true, 0x20, false, false, false, true},
		{0xc0, false, false, true, false, 0x26, false, false, false, true},
		{0xc0, false, false, true, true, 0x26, false, false, false, true},
		{0xc0, false, true, false, false, 0xc0, false, true, false, false},
		{0xc0, false, true, false, true, 0x60, false, true, false, true},
		{0xc0, false, true, true, false, 0xba, false, true, false, false},
		{0xc0, false, true, true, true, 0x5a, false, true, false, true},
		{0xc0, true, false, false, false, 0x20, false, false, false, true},
		{0xc0, true, false, false, true, 0x20, false, false, false, true},
		{0xc0, true, false, true, false, 0x26, false, false, false, true},
		{0xc0, true, false, true, true, 0x26, false, false, false, true},
		{0xc0, true, true, false, false, 0xc0, false, true, false, false},
		{0xc0, true, true, false, true, 0x60, false, true, false, true},
		{0xc0, true, true, true, false, 0xba, false, true, false, false},
		{0xc0, true, true, true, true, 0x5a, false, true, false, true},
		{0xc1, false, false, false, false, 0x21, false, false, false, true},
		{0xc1, false, false, false, true, 0x21, false, false, false, true},
		{0xc1, false, false, true, false, 0x27, false, false, false, true},
		{0xc1, false, false, true, true, 0x27, false, false, false, true},
		{0xc1, false, true, false, false, 0xc1, false, true, false, false},
		{0xc1, false, true, false, true, 0x61, false, true, false, true},
		{0xc1, false, true, true, false, 0xbb, false, true, false, false},
		{0xc1, false, true, true, true, 0x5b, false, true, false, true},
		{0xc1, true, false, false, false, 0x21, false, false, false, true},
		{0xc1, true, false, false, true, 0x21, false, false, false, true},
		{0xc1, true, false, true, false, 0x27, false, false, false, true},
		{0xc1, true, false, true, true, 0x27, false, false, false, true},
		{0xc1, true, true, false, false, 0xc1, false, true, false, false},
		{0xc1, true, true, false, true, 0x61, false, true, false, true},
		{0xc1, true, true, true, false, 0xbb, false, true, false, false},
		{0xc1, true, true, true, true, 0x5b, false, true, false, true},
		{0xc2, false, false, false, false, 0x22, false, false, false, true},
		{0xc2, false, false, false, true, 0x22, false, false, false, true},
		{0xc2, false, false, true, false, 0x28, false, false, false, true},
		{0xc2, false, false, true, true, 0x28, false, false, false, true},
		{0xc2, false, true, false, false, 0xc2, false, true, false, false},
		{0xc2, false, true, false, true, 0x62, false, true, false, true},
		{0xc2, false, true, true, false, 0xbc, false, true, false, false},
		{0xc2, false, true, true, true, 0x5c, false, true, false, true},
		{0xc2, true, false, false, false, 0x22, false, false, false, true},
		{0xc2, true, false, false, true, 0x22, false, false, false, true},
		{0xc2, true, false, true, false, 0x28, false, false, false, true},
		{0xc2, true, false, true, true, 0x28, false, false, false, true},
		{0xc2, true, true, false, false, 0xc2, false, true, false, false},
		{0xc2, true, true, false, true, 0x62, false, true, false, true},
		{0xc2, true, true, true, false, 0xbc, false, true, false, false},
		{0xc2, true, true, true, true, 0x5c, false, true, false, true},
		{0xc3, false, false, false, false, 0x23, false, false, false, true},
		{0xc3, false, false, false, true, 0x23, false, false, false, true},
		{0xc3, false, false, true, false, 0x29, false, false, false, true},
		{0xc3, false, false, true, true, 0x29, false, false, false, true},
		{0xc3, false, true, false, false, 0xc3, false, true, false, false},
		{0xc3, false, true, false, true, 0x63, false, true, false, true},
		{0xc3, false, true, true, false, 0xbd, false, true, false, false},
		{0xc3, false, true, true, true, 0x5d, false, true, false, true},
		{0xc3, true, false, false, false, 0x23, false, false, false, true},
		{0xc3, true, false, false, true, 0x23, false, false, false, true},
		{0xc3, true, false, true, false, 0x29, false, false, false, true},
		{0xc3, true, false, true, true, 0x29, false, false, false, true},
		{0xc3, true, true, false, false, 0xc3, false, true, false, false},
		{0xc3, true, true, false, true, 0x63, false, true, false, true},
		{0xc3, true, true, true, false, 0xbd, false, true, false, false},
		{0xc3, true, true, true, true, 0x5d, false, true, false, true},
		{0xc4, false, false, false, false, 0x24, false, false, false, true},
		{0xc4, false, false, false, true, 0x24, false, false, false, true},
		{0xc4, false, false, true, false, 0x2a, false, false, false, true},
		{0xc4, false, false, true, true, 0x2a, false, false, false, true},
		{0xc4, false, true, false, false, 0xc4, false, true, false, false},
		{0xc4, false, true, false, true, 0x64, false, true, false, true},
		{0xc4, false, true, true, false, 0xbe, false, true, false, false},
		{0xc4, false, true, true, true, 0x5e, false, true, false, true},
		{0xc4, true, false, false, false, 0x24, false, false, false, true},
		{0xc4, true, false, false, true, 0x24, false, false, false, true},
		{0xc4, true, false, true, false, 0x2a, false, false, false, true},
		{0xc4, true, false, true, true, 0x2a, false, false, false, true},
		{0xc4, true, true, false, false, 0xc4, false, true, false, false},
		{0xc4, true, true, false, true, 0x64, false, true, false, true},
		{0xc4, true, true, true, false, 0xbe, false, true, false, false},
		{0xc4, true, true, true, true, 0x5e, false, true, false, true},
		{0xc5, false, false, false, false, 0x25, false, false, false, true},
		{0xc5, false, false, false, true, 0x25, false, false, false, true},
		{0xc5, false, false, true, false, 0x2b, false, false, false, true},
		{0xc5, false, false, true, true, 0x2b, false, false, false, true},
		{0xc5, false, true, false, false, 0xc5, false, true, false, false},
		{0xc5, false, true, false, true, 0x65, false, true, false, true},
		{0xc5, false, true, true, false, 0xbf, false, true, false, false},
		{0xc5, false, true, true, true, 0x5f, false, true, false, true},
		{0xc5, true, false, false, false, 0x25, false, false, false, true},
		{0xc5, true, false, false, true, 0x25, false, false, false, true},
		{0xc5, true, false, true, false, 0x2b, false, false, false, true},
		{0xc5, true, false, true, true, 0x2b, false, false, false, true},
		{0xc5, true, true, false, false, 0xc5, false, true, false, false},
		{0xc5, true, true, false, true, 0x65, false, true, false, true},
		{0xc5, true, true, true, false, 0xbf, false, true, false, false},
		{0xc5, true, true, true, true, 0x5f, false, true, false, true},
		{0xc6, false, false, false, false, 0x26, false, false, false, true},
		{0xc6, false, false, false, true, 0x26, false, false, false, true},
		{0xc6, false, false, true, false, 0x2c, false, false, false, true},
		{0xc6, false, false, true, true, 0x2c, false, false, false, true},
		{0xc6, false, true, false, false, 0xc6, false, true, false, false},
		{0xc6, false, true, false, true, 0x66, false, true, false, true},
		{0xc6, false, true, true, false, 0xc0, false, true, false, false},
		{0xc6, false, true, true, true, 0x60, false, true, false, true},
		{0xc6, true, false, false, false, 0x26, false, false, false, true},
		{0xc6, true, false, false, true, 0x26, false, false, false, true},
		{0xc6, true, false, true, false, 0x2c, false, false, false, true},
		{0xc6, true, false, true, true, 0x2c, false, false, false, true},
		{0xc6, true, true, false, false, 0xc6, false, true, false, false},
		{0xc6, true, true, false, true, 0x66, false, true, false, true},
		{0xc6, true, true, true, false, 0xc0, false, true, false, false},
		{0xc6, true, true, true, true, 0x60, false, true, false, true},
		{0xc7, false, false, false, false, 0x27, false, false, false, true},
		{0xc7, false, false, false, true, 0x27, false, false, false, true},
		{0xc7, false, false, true, false, 0x2d, false, false, false, true},
		{0xc7, false, false, true, true, 0x2d, false, false, false, true},
		{0xc7, false, true, false, false, 0xc7, false, true, false, false},
		{0xc7, false, true, false, true, 0x67, false, true, false, true},
		{0xc7, false, true, true, false, 0xc1, false, true, false, false},
		{0xc7, false, true, true, true, 0x61, false, true, false, true},
		{0xc7, true, false, false, false, 0x27, false, false, false, true},
		{0xc7, true, false, false, true, 0x27, false, false, false, true},
		{0xc7, true, false, true, false, 0x2d, false, false, false, true},
		{0xc7, true, false, true, true, 0x2d, false, false, false, true},
		{0xc7, true, true, false, false, 0xc7, false, true, false, false},
		{0xc7, true, true, false, true, 0x67, false, true, false, true},
		{0xc7, true, true, true, false, 0xc1, false, true, false, false},
		{0xc7, true, true, true, true, 0x61, false, true, false, true},
		{0xc8, false, false, false, false, 0x28, false, false, false, true},
		{0xc8, false, false, false, true, 0x28, false, false, false, true},
		{0xc8, false, false, true, false, 0x2e, false, false, false, true},
		{0xc8, false, false, true, true, 0x2e, false, false, false, true},
		{0xc8, false, true, false, false, 0xc8, false, true, false, false},
		{0xc8, false, true, false, true, 0x68, false, true, false, true},
		{0xc8, false, true, true, false, 0xc2, false, true, false, false},
		{0xc8, false, true, true, true, 0x62, false, true, false, true},
		{0xc8, true, false, false, false, 0x28, false, false, false, true},
		{0xc8, true, false, false, true, 0x28, false, false, false, true},
		{0xc8, true, false, true, false, 0x2e, false, false, false, true},
		{0xc8, true, false, true, true, 0x2e, false, false, false, true},
		{0xc8, true, true, false, false, 0xc8, false, true, false, false},
		{0xc8, true, true, false, true, 0x68, false, true, false, true},
		{0xc8, true, true, true, false, 0xc2, false, true, false, false},
		{0xc8, true, true, true, true, 0x62, false, true, false, true},
		{0xc9, false, false, false, false, 0x29, false, false, false, true},
		{0xc9, false, false, false, true, 0x29, false, false, false, true},
		{0xc9, false, false, true, false, 0x2f, false, false, false, true},
		{0xc9, false, false, true, true, 0x2f, false, false, false, true},
		{0xc9, false, true, false, false, 0xc9, false, true, false, false},
		{0xc9, false, true, false, true, 0x69, false, true, false, true},
		{0xc9, false, true, true, false, 0xc3, false, true, false, false},
		{0xc9, false, true, true, true, 0x63, false, true, false, true},
		{0xc9, true, false, false, false, 0x29, false, false, false, true},
		{0xc9, true, false, false, true, 0x29, false, false, false, true},
		{0xc9, true, false, true, false, 0x2f, false, false, false, true},
		{0xc9, true, false, true, true, 0x2f, false, false, false, true},
		{0xc9, true, true, false, false, 0xc9, false, true, false, false},
		{0xc9, true, true, false, true, 0x69, false, true, false, true},
		{0xc9, true, true, true, false, 0xc3, false, true, false, false},
		{0xc9, true, true, true, true, 0x63, false, true, false, true},
		{0xca, false, false, false, false, 0x30, false, false, false, true},
		{0xca, false, false, false, true, 0x30, false, false, false, true},
		{0xca, false, false, true, false, 0x30, false, false, false, true},
		{0xca, false, false, true, true, 0x30, false, false, false, true},
		{0xca, false, true, false, false, 0xca, false, true, false, false},
		{0xca, false, true, false, true, 0x6a, false, true, false, true},
		{0xca, false, true, true, false, 0xc4, false, true, false, false},
		{0xca, false, true, true, true, 0x64, false, true, false, true},
		{0xca, true, false, false, false, 0x30, false, false, false, true},
		{0xca, true, false, false, true, 0x30, false, false, false, true},
		{0xca, true, false, true, false, 0x30, false, false, false, true},
		{0xca, true, false, true, true, 0x30, false, false, false, true},
		{0xca, true, true, false, false, 0xca, false, true, false, false},
		{0xca, true, true, false, true, 0x6a, false, true, false, true},
		{0xca, true, true, true, false, 0xc4, false, true, false, false},
		{0xca, true, true, true, true, 0x64, false, true, false, true},
		{0xcb, false, false, false, false, 0x31, false, false, false, true},
		{0xcb, false, false, false, true, 0x31, false, false, false, true},
		{0xcb, false, false, true, false, 0x31, false, false, false, true},
		{0xcb, false, false, true, true, 0x31, false, false, false, true},
		{0xcb, false, true, false, false, 0xcb, false, true, false, false},
		{0xcb, false, true, false, true, 0x6b, false, true, false, true},
		{0xcb, false, true, true, false, 0xc5, false, true, false, false},
		{0xcb, false, true, true, true, 0x65, false, true, false, true},
		{0xcb, true, false, false, false, 0x31, false, false, false, true},
		{0xcb, true, false, false, true, 0x31, false, false, false, true},
		{0xcb, true, false, true, false, 0x31, false, false, false, true},
		{0xcb, true, false, true, true, 0x31, false, false, false, true},
		{0xcb, true, true, false, false, 0xcb, false, true, false, false},
		{0xcb, true, true, false, true, 0x6b, false, true, false, true},
		{0xcb, true, true, true, false, 0xc5, false, true, false, false},
		{0xcb, true, true, true, true, 0x65, false, true, false, true},
		{0xcc, false, false, false, false, 0x32, false, false, false, true},
		{0xcc, false, false, false, true, 0x32, false, false, false, true},
		{0xcc, false, false, true, false, 0x32, false, false, false, true},
		{0xcc, false, false, true, true, 0x32, false, false, false, true},
		{0xcc, false, true, false, false, 0xcc, false, true, false, false},
		{0xcc, false, true, false, true, 0x6c, false, true, false, true},
		{0xcc, false, true, true, false, 0xc6, false, true, false, false},
		{0xcc, false, true, true, true, 0x66, false, true, false, true},
		{0xcc, true, false, false, false, 0x32, false, false, false, true},
		{0xcc, true, false, false, true, 0x32, false, false, false, true},
		{0xcc, true, false, true, false, 0x32, false, false, false, true},
		{0xcc, true, false, true, true, 0x32, false, false, false, true},
		{0xcc, true, true, false, false, 0xcc, false, true, false, false},
		{0xcc, true, true, false, true, 0x6c, false, true, false, true},
		{0xcc, true, true, true, false, 0xc6, false, true, false, false},
		{0xcc, true, true, true, true, 0x66, false, true, false, true},
		{0xcd, false, false, false, false, 0x33, false, false, false, true},
		{0xcd, false, false, false, true, 0x33, false, false, false, true},
		{0xcd, false, false, true, false, 0x33, false, false, false, true},
		{0xcd, false, false, true, true, 0x33, false, false, false, true},
		{0xcd, false, true, false, false, 0xcd, false, true, false, false},
		{0xcd, false, true, false, true, 0x6d, false, true, false, true},
		{0xcd, false, true, true, false, 0xc7, false, true, false, false},
		{0xcd, false, true, true, true, 0x67, false, true, false, true},
		{0xcd, true, false, false, false, 0x33, false, false, false, true},
		{0xcd, true, false, false, true, 0x33, false, false, false, true},
		{0xcd, true, false, true, false, 0x33, false, false, false, true},
		{0xcd, true, false, true, true, 0x33, false, false, false, true},
		{0xcd, true, true, false, false, 0xcd, false, true, false, false},
		{0xcd, true, true, false, true, 0x6d, false, true, false, true},
		{0xcd, true, true, true, false, 0xc7, false, true, false, false},
		{0xcd, true, true, true, true, 0x67, false, true, false, true},
		{0xce, false, false, false, false, 0x34, false, false, false, true},
		{0xce, false, false, false, true, 0x34, false, false, false, true},
		{0xce, false, false, true, false, 0x34, false, false, false, true},
		{0xce, false, false, true, true, 0x34, false, false, false, true},
		{0xce, false, true, false, false, 0xce, false, true, false, false},
		{0xce, false, true, false, true, 0x6e, false, true, false, true},
		{0xce, false, true, true, false, 0xc8, false, true, false, false},
		{0xce, false, true, true, true, 0x68, false, true, false, true},
		{0xce, true, false, false, false, 0x34, false, false, false, true},
		{0xce, true, false, false, true, 0x34, false, false, false, true},
		{0xce, true, false, true, false, 0x34, false, false, false, true},
		{0xce, true, false, true, true, 0x34, false, false, false, true},
		{0xce, true, true, false, false, 0xce, false, true, false, false},
		{0xce, true, true, false, true, 0x6e, false, true, false, true},
		{0xce, true, true, true, false, 0xc8, false, true, false, false},
		{0xce, true, true, true, true, 0x68, false, true, false, true},
		{0xcf, false, false, false, false, 0x35, false, false, false, true},
		{0xcf, false, false, false, true, 0x35, false, false, false, true},
		{0xcf, false, false, true, false, 0x35, false, false, false, true},
		{0xcf, false, false, true, true, 0x35, false, false, false, true},
		{0xcf, false, true, false, false, 0xcf, false, true, false, false},
		{0xcf, false, true, false, true, 0x6f, false, true, false, true},
		{0xcf, false, true, true, false, 0xc9, false, true, false, false},
		{0xcf, false, true, true, true, 0x69, false, true, false, true},
		{0xcf, true, false, false, false, 0x35, false, false, false, true},
		{0xcf, true, false, false, true, 0x35, false, false, false, true},
		{0xcf, true, false, true, false, 0x35, false, false, false, true},
		{0xcf, true, false, true, true, 0x35, false, false, false, true},
		{0xcf, true, true, false, false, 0xcf, false, true, false, false},
		{0xcf, true, true, false, true, 0x6f, false, true, false, true},
		{0xcf, true, true, true, false, 0xc9, false, true, false, false},
		{0xcf, true, true, true, true, 0x69, false, true, false, true},
		{0xd0, false, false, false, false, 0x30, false, false, false, true},
		{0xd0, false, false, false, true, 0x30, false, false, false, true},
		{0xd0, false, false, true, false, 0x36, false, false, false, true},
		{0xd0, false, false, true, true, 0x36, false, false, false, true},
		{0xd0, false, true, false, false, 0xd0, false, true, false, false},
		{0xd0, false, true, false, true, 0x70, false, true, false, true},
		{0xd0, false, true, true, false, 0xca, false, true, false, false},
		{0xd0, false, true, true, true, 0x6a, false, true, false, true},
		{0xd0, true, false, false, false, 0x30, false, false, false, true},
		{0xd0, true, false, false, true, 0x30, false, false, false, true},
		{0xd0, true, false, true, false, 0x36, false, false, false, true},
		{0xd0, true, false, true, true, 0x36, false, false, false, true},
		{0xd0, true, true, false, false, 0xd0, false, true, false, false},
		{0xd0, true, true, false, true, 0x70, false, true, false, true},
		{0xd0, true, true, true, false, 0xca, false, true, false, false},
		{0xd0, true, true, true, true, 0x6a, false, true, false, true},
		{0xd1, false, false, false, false, 0x31, false, false, false, true},
		{0xd1, false, false, false, true, 0x31, false, false, false, true},
		{0xd1, false, false, true, false, 0x37, false, false, false, true},
		{0xd1, false, false, true, true, 0x37, false, false, false, true},
		{0xd1, false, true, false, false, 0xd1, false, true, false, false},
		{0xd1, false, true, false, true, 0x71, false, true, false, true},
		{0xd1, false, true, true, false, 0xcb, false, true, false, false},
		{0xd1, false, true, true, true, 0x6b, false, true, false, true},
		{0xd1, true, false, false, false, 0x31, false, false, false, true},
		{0xd1, true, false, false, true, 0x31, false, false, false, true},
		{0xd1, true, false, true, false, 0x37, false, false, false, true},
		{0xd1, true, false, true, true, 0x37, false, false, false, true},
		{0xd1, true, true, false, false, 0xd1, false, true, false, false},
		{0xd1, true, true, false, true, 0x71, false, true, false, true},
		{0xd1, true, true, true, false, 0xcb, false, true, false, false},
		{0xd1, true, true, true, true, 0x6b, false, true, false, true},
		{0xd2, false, false, false, false, 0x32, false, false, false, true},
		{0xd2, false, false, false, true, 0x32, false, false, false, true},
		{0xd2, false, false, true, false, 0x38, false, false, false, true},
		{0xd2, false, false, true, true, 0x38, false, false, false, true},
		{0xd2, false, true, false, false, 0xd2, false, true, false, false},
		{0xd2, false, true, false, true, 0x72, false, true, false, true},
		{0xd2, false, true, true, false, 0xcc, false, true, false, false},
		{0xd2, false, true, true, true, 0x6c, false, true, false, true},
		{0xd2, true, false, false, false, 0x32, false, false, false, true},
		{0xd2, true, false, false, true, 0x32, false, false, false, true},
		{0xd2, true, false, true, false, 0x38, false, false, false, true},
		{0xd2, true, false, true, true, 0x38, false, false, false, true},
		{0xd2, true, true, false, false, 0xd2, false, true, false, false},
		{0xd2, true, true, false, true, 0x72, false, true, false, true},
		{0xd2, true, true, true, false, 0xcc, false, true, false, false},
		{0xd2, true, true, true, true, 0x6c, false, true, false, true},
		{0xd3, false, false, false, false, 0x33, false, false, false, true},
		{0xd3, false, false, false, true, 0x33, false, false, false, true},
		{0xd3, false, false, true, false, 0x39, false, false, false, true},
		{0xd3, false, false, true, true, 0x39, false, false, false, true},
		{0xd3, false, true, false, false, 0xd3, false, true, false, false},
		{0xd3, false, true, false, true, 0x73, false, true, false, true},
		{0xd3, false, true, true, false, 0xcd, false, true, false, false},
		{0xd3, false, true, true, true, 0x6d, false, true, false, true},
		{0xd3, true, false, false, false, 0x33, false, false, false, true},
		{0xd3, true, false, false, true, 0x33, false, false, false, true},
		{0xd3, true, false, true, false, 0x39, false, false, false, true},
		{0xd3, true, false, true, true, 0x39, false, false, false, true},
		{0xd3, true, true, false, false, 0xd3, false, true, false, false},
		{0xd3, true, true, false, true, 0x73, false, true, false, true},
		{0xd3, true, true, true, false, 0xcd, false, true, false, false},
		{0xd3, true, true, true, true, 0x6d, false, true, false, true},
		{0xd4, false, false, false, false, 0x34, false, false, false, true},
		{0xd4, false, false, false, true, 0x34, false, false, false, true},
		{0xd4, false, false, true, false, 0x3a, false, false, false, true},
		{0xd4, false, false, true, true, 0x3a, false, false, false, true},
		{0xd4, false, true, false, false, 0xd4, false, true, false, false},
		{0xd4, false, true, false, true, 0x74, false, true, false, true},
		{0xd4, false, true, true, false, 0xce, false, true, false, false},
		{0xd4, false, true, true, true, 0x6e, false, true, false, true},
		{0xd4, true, false, false, false, 0x34, false, false, false, true},
		{0xd4, true, false, false, true, 0x34, false, false, false, true},
		{0xd4, true, false, true, false, 0x3a, false, false, false, true},
		{0xd4, true, false, true, true, 0x3a, false, false, false, true},
		{0xd4, true, true, false, false, 0xd4, false, true, false, false},
		{0xd4, true, true, false, true, 0x74, false, true, false, true},
		{0xd4, true, true, true, false, 0xce, false, true, false, false},
		{0xd4, true, true, true, true, 0x6e, false, true, false, true},
		{0xd5, false, false, false, false, 0x35, false, false, false, true},
		{0xd5, false, false, false, true, 0x35, false, false, false, true},
		{0xd5, false, false, true, false, 0x3b, false, false, false, true},
		{0xd5, false, false, true, true, 0x3b, false, false, false, true},
		{0xd5, false, true, false, false, 0xd5, false, true, false, false},
		{0xd5, false, true, false, true, 0x75, false, true, false, true},
		{0xd5, false, true, true, false, 0xcf, false, true, false, false},
		{0xd5, false, true, true, true, 0x6f, false, true, false, true},
		{0xd5, true, false, false, false, 0x35, false, false, false, true},
		{0xd5, true, false, false, true, 0x35, false, false, false, true},
		{0xd5, true, false, true, false, 0x3b, false, false, false, true},
		{0xd5, true, false, true, true, 0x3b, false, false, false, true},
		{0xd5, true, true, false, false, 0xd5, false, true, false, false},
		{0xd5, true, true, false, true, 0x75, false, true, false, true},
		{0xd5, true, true, true, false, 0xcf, false, true, false, false},
		{0xd5, true, true, true, true, 0x6f, false, true, false, true},
		{0xd6, false, false, false, false, 0x36, false, false, false, true},
		{0xd6, false, false, false, true, 0x36, false, false, false, true},
		{0xd6, false, false, true, false, 0x3c, false, false, false, true},
		{0xd6, false, false, true, true, 0x3c, false, false, false, true},
		{0xd6, false, true, false, false, 0xd6, false, true, false, false},
		{0xd6, false, true, false, true, 0x76, false, true, false, true},
		{0xd6, false, true, true, false, 0xd0, false, true, false, false},
		{0xd6, false, true, true, true, 0x70, false, true, false, true},
		{0xd6, true, false, false, false, 0x36, false, false, false, true},
		{0xd6, true, false, false, true, 0x36, false, false, false, true},
		{0xd6, true, false, true, false, 0x3c, false, false, false, true},
		{0xd6, true, false, true, true, 0x3c, false, false, false, true},
		{0xd6, true, true, false, false, 0xd6, false, true, false, false},
		{0xd6, true, true, false, true, 0x76, false, true, false, true},
		{0xd6, true, true, true, false, 0xd0, false, true, false, false},
		{0xd6, true, true, true, true, 0x70, false, true, false, true},
		{0xd7, false, false, false, false, 0x37, false, false, false, true},
		{0xd7, false, false, false, true, 0x37, false, false, false, true},
		{0xd7, false, false, true, false, 0x3d, false, false, false, true},
		{0xd7, false, false, true, true, 0x3d, false, false, false, true},
		{0xd7, false, true, false, false, 0xd7, false, true, false, false},
		{0xd7, false, true, false, true, 0x77, false, true, false, true},
		{0xd7, false, true, true, false, 0xd1, false, true, false, false},
		{0xd7, false, true, true, true, 0x71, false, true, false, true},
		{0xd7, true, false, false, false, 0x37, false, false, false, true},
		{0xd7, true, false, false, true, 0x37, false, false, false, true},
		{0xd7, true, false, true, false, 0x3d, false, false, false, true},
		{0xd7, true, false, true, true, 0x3d, false, false, false, true},
		{0xd7, true, true, false, false, 0xd7, false, true, false, false},
		{0xd7, true, true, false, true, 0x77, false, true, false, true},
		{0xd7, true, true, true, false, 0xd1, false, true, false, false},
		{0xd7, true, true, true, true, 0x71, false, true, false, true},
		{0xd8, false, false, false, false, 0x38, false, false, false, true},
		{0xd8, false, false, false, true, 0x38, false, false, false, true},
		{0xd8, false, false, true, false, 0x3e, false, false, false, true},
		{0xd8, false, false, true, true, 0x3e, false, false, false, true},
		{0xd8, false, true, false, false, 0xd8, false, true, false, false},
		{0xd8, false, true, false, true, 0x78, false, true, false, true},
		{0xd8, false, true, true, false, 0xd2, false, true, false, false},
		{0xd8, false, true, true, true, 0x72, false, true, false, true},
		{0xd8, true, false, false, false, 0x38, false, false, false, true},
		{0xd8, true, false, false, true, 0x38, false, false, false, true},
		{0xd8, true, false, true, false, 0x3e, false, false, false, true},
		{0xd8, true, false, true, true, 0x3e, false, false, false, true},
		{0xd8, true, true, false, false, 0xd8, false, true, false, false},
		{0xd8, true, true, false, true, 0x78, false, true, false, true},
		{0xd8, true, true, true, false, 0xd2, false, true, false, false},
		{0xd8, true, true, true, true, 0x72, false, true, false, true},
		{0xd9, false, false, false, false, 0x39, false, false, false, true},
		{0xd9, false, false, false, true, 0x39, false, false, false, true},
		{0xd9, false, false, true, false, 0x3f, false, false, false, true},
		{0xd9, false, false, true, true, 0x3f, false, false, false, true},
		{0xd9, false, true, false, false, 0xd9, false, true, false, false},
		{0xd9, false, true, false, true, 0x79, false, true, false, true},
		{0xd9, false, true, true, false, 0xd3, false, true, false, false},
		{0xd9, false, true, true, true, 0x73, false, true, false, true},
		{0xd9, true, false, false, false, 0x39, false, false, false, true},
		{0xd9, true, false, false, true, 0x39, false, false, false, true},
		{0xd9, true, false, true, false, 0x3f, false, false, false, true},
		{0xd9, true, false, true, true, 0x3f, false, false, false, true},
		{0xd9, true, true, false, false, 0xd9, false, true, false, false},
		{0xd9, true, true, false, true, 0x79, false, true, false, true},
		{0xd9, true, true, true, false, 0xd3, false, true, false, false},
		{0xd9, true, true, true, true, 0x73, false, true, false, true},
		{0xda, false, false, false, false, 0x40, false, false, false, true},
		{0xda, false, false, false, true, 0x40, false, false, false, true},
		{0xda, false, false, true, false, 0x40, false, false, false, true},
		{0xda, false, false, true, true, 0x40, false, false, false, true},
		{0xda, false, true, false, false, 0xda, false, true, false, false},
		{0xda, false, true, false, true, 0x7a, false, true, false, true},
		{0xda, false, true, true, false, 0xd4, false, true, false, false},
		{0xda, false, true, true, true, 0x74, false, true, false, true},
		{0xda, true, false, false, false, 0x40, false, false, false, true},
		{0xda, true, false, false, true, 0x40, false, false, false, true},
		{0xda, true, false, true, false, 0x40, false, false, false, true},
		{0xda, true, false, true, true, 0x40, false, false, false, true},
		{0xda, true, true, false, false, 0xda, false, true, false, false},
		{0xda, true, true, false, true, 0x7a, false, true, false, true},
		{0xda, true, true, true, false, 0xd4, false, true, false, false},
		{0xda, true, true, true, true, 0x74, false, true, false, true},
		{0xdb, false, false, false, false, 0x41, false, false, false, true},
		{0xdb, false, false, false, true, 0x41, false, false, false, true},
		{0xdb, false, false, true, false, 0x41, false, false, false, true},
		{0xdb, false, false, true, true, 0x41, false, false, false, true},
		{0xdb, false, true, false, false, 0xdb, false, true, false, false},
		{0xdb, false, true, false, true, 0x7b, false, true, false, true},
		{0xdb, false, true, true, false, 0xd5, false, true, false, false},
		{0xdb, false, true, true, true, 0x75, false, true, false, true},
		{0xdb, true, false, false, false, 0x41, false, false, false, true},
		{0xdb, true, false, false, true, 0x41, false, false, false, true},
		{0xdb, true, false, true, false, 0x41, false, false, false, true},
		{0xdb, true, false, true, true, 0x41, false, false, false, true},
		{0xdb, true, true, false, false, 0xdb, false, true, false, false},
		{0xdb, true, true, false, true, 0x7b, false, true, false, true},
		{0xdb, true, true, true, false, 0xd5, false, true, false, false},
		{0xdb, true, true, true, true, 0x75, false, true, false, true},
		{0xdc, false, false, false, false, 0x42, false, false, false, true},
		{0xdc, false, false, false, true, 0x42, false, false, false, true},
		{0xdc, false, false, true, false, 0x42, false, false, false, true},
		{0xdc, false, false, true, true, 0x42, false, false, false, true},
		{0xdc, false, true, false, false, 0xdc, false, true, false, false},
		{0xdc, false, true, false, true, 0x7c, false, true, false, true},
		{0xdc, false, true, true, false, 0xd6, false, true, false, false},
		{0xdc, false, true, true, true, 0x76, false, true, false, true},
		{0xdc, true, false, false, false, 0x42, false, false, false, true},
		{0xdc, true, false, false, true, 0x42, false, false, false, true},
		{0xdc, true, false, true, false, 0x42, false, false, false, true},
		{0xdc, true, false, true, true, 0x42, false, false, false, true},
		{0xdc, true, true, false, false, 0xdc, false, true, false, false},
		{0xdc, true, true, false, true, 0x7c, false, true, false, true},
		{0xdc, true, true, true, false, 0xd6, false, true, false, false},
		{0xdc, true, true, true, true, 0x76, false, true, false, true},
		{0xdd, false, false, false, false, 0x43, false, false, false, true},
		{0xdd, false, false, false, true, 0x43, false, false, false, true},
		{0xdd, false, false, true, false, 0x43, false, false, false, true},
		{0xdd, false, false, true, true, 0x43, false, false, false, true},
		{0xdd, false, true, false, false, 0xdd, false, true, false, false},
		{0xdd, false, true, false, true, 0x7d, false, true, false, true},
		{0xdd, false, true, true, false, 0xd7, false, true, false, false},
		{0xdd, false, true, true, true, 0x77, false, true, false, true},
		{0xdd, true, false, false, false, 0x43, false, false, false, true},
		{0xdd, true, false, false, true, 0x43, false, false, false, true},
		{0xdd, true, false, true, false, 0x43, false, false, false, true},
		{0xdd, true, false, true, true, 0x43, false, false, false, true},
		{0xdd, true, true, false, false, 0xdd, false, true, false, false},
		{0xdd, true, true, false, true, 0x7d, false, true, false, true},
		{0xdd, true, true, true, false, 0xd7, false, true, false, false},
		{0xdd, true, true, true, true, 0x77, false, true, false, true},
		{0xde, false, false, false, false, 0x44, false, false, false, true},
		{0xde, false, false, false, true, 0x44, false, false, false, true},
		{0xde, false, false, true, false, 0x44, false, false, false, true},
		{0xde, false, false, true, true, 0x44, false, false, false, true},
		{0xde, false, true, false, false, 0xde, false, true, false, false},
		{0xde, false, true, false, true, 0x7e, false, true, false, true},
		{0xde, false, true, true, false, 0xd8, false, true, false, false},
		{0xde, false, true, true, true, 0x78, false, true, false, true},
		{0xde, true, false, false, false, 0x44, false, false, false, true},
		{0xde, true, false, false, true, 0x44, false, false, false, true},
		{0xde, true, false, true, false, 0x44, false, false, false, true},
		{0xde, true, false, true, true, 0x44, false, false, false, true},
		{0xde, true, true, false, false, 0xde, false, true, false, false},
		{0xde, true, true, false, true, 0x7e, false, true, false, true},
		{0xde, true, true, true, false, 0xd8, false, true, false, false},
		{0xde, true, true, true, true, 0x78, false, true, false, true},
		{0xdf, false, false, false, false, 0x45, false, false, false, true},
		{0xdf, false, false, false, true, 0x45, false, false, false, true},
		{0xdf, false, false, true, false, 0x45, false, false, false, true},
		{0xdf, false, false, true, true, 0x45, false, false, false, true},
		{0xdf, false, true, false, false, 0xdf, false, true, false, false},
		{0xdf, false, true, false, true, 0x7f, false, true, false, true},
		{0xdf, false, true, true, false, 0xd9, false, true, false, false},
		{0xdf, false, true, true, true, 0x79, false, true, false, true},
		{0xdf, true, false, false, false, 0x45, false, false, false, true},
		{0xdf, true, false, false, true, 0x45, false, false, false, true},
		{0xdf, true, false, true, false, 0x45, false, false, false, true},
		{0xdf, true, false, true, true, 0x45, false, false, false, true},
		{0xdf, true, true, false, false, 0xdf, false, true, false, false},
		{0xdf, true, true, false, true, 0x7f, false, true, false, true},
		{0xdf, true, true, true, false, 0xd9, false, true, false, false},
		{0xdf, true, true, true, true, 0x79, false, true, false, true},
		{0xe0, false, false, false, false, 0x40, false, false, false, true},
		{0xe0, false, false, false, true, 0x40, false, false, false, true},
		{0xe0, false, false, true, false, 0x46, false, false, false, true},
		{0xe0, false, false, true, true, 0x46, false, false, false, true},
		{0xe0, false, true, false, false, 0xe0, false, true, false, false},
		{0xe0, false, true, false, true, 0x80, false, true, false, true},
		{0xe0, false, true, true, false, 0xda, false, true, false, false},
		{0xe0, false, true, true, true, 0x7a, false, true, false, true},
		{0xe0, true, false, false, false, 0x40, false, false, false, true},
		{0xe0, true, false, false, true, 0x40, false, false, false, true},
		{0xe0, true, false, true, false, 0x46, false, false, false, true},
		{0xe0, true, false, true, true, 0x46, false, false, false, true},
		{0xe0, true, true, false, false, 0xe0, false, true, false, false},
		{0xe0, true, true, false, true, 0x80, false, true, false, true},
		{0xe0, true, true, true, false, 0xda, false, true, false, false},
		{0xe0, true, true, true, true, 0x7a, false, true, false, true},
		{0xe1, false, false, false, false, 0x41, false, false, false, true},
		{0xe1, false, false, false, true, 0x41, false, false, false, true},
		{0xe1, false, false, true, false, 0x47, false, false, false, true},
		{0xe1, false, false, true, true, 0x47, false, false, false, true},
		{0xe1, false, true, false, false, 0xe1, false, true, false, false},
		{0xe1, false, true, false, true, 0x81, false, true, false, true},
		{0xe1, false, true, true, false, 0xdb, false, true, false, false},
		{0xe1, false, true, true, true, 0x7b, false, true, false, true},
		{0xe1, true, false, false, false, 0x41, false, false, false, true},
		{0xe1, true, false, false, true, 0x41, false, false, false, true},
		{0xe1, true, false, true, false, 0x47, false, false, false, true},
		{0xe1, true, false, true, true, 0x47, false, false, false, true},
		{0xe1, true, true, false, false, 0xe1, false, true, false, false},
		{0xe1, true, true, false, true, 0x81, false, true, false, true},
		{0xe1, true, true, true, false, 0xdb, false, true, false, false},
		{0xe1, true, true, true, true, 0x7b, false, true, false, true},
		{0xe2, false, false, false, false, 0x42, false, false, false, true},
		{0xe2, false, false, false, true, 0x42, false, false, false, true},
		{0xe2, false, false, true, false, 0x48, false, false, false, true},
		{0xe2, false, false, true, true, 0x48, false, false, false, true},
		{0xe2, false, true, false, false, 0xe2, false, true, false, false},
		{0xe2, false, true, false, true, 0x82, false, true, false, true},
		{0xe2, false, true, true, false, 0xdc, false, true, false, false},
		{0xe2, false, true, true, true, 0x7c, false, true, false, true},
		{0xe2, true, false, false, false, 0x42, false, false, false, true},
		{0xe2, true, false, false, true, 0x42, false, false, false, true},
		{0xe2, true, false, true, false, 0x48, false, false, false, true},
		{0xe2, true, false, true, true, 0x48, false, false, false, true},
		{0xe2, true, true, false, false, 0xe2, false, true, false, false},
		{0xe2, true, true, false, true, 0x82, false, true, false, true},
		{0xe2, true, true, true, false, 0xdc, false, true, false, false},
		{0xe2, true, true, true, true, 0x7c, false, true, false, true},
		{0xe3, false, false, false, false, 0x43, false, false, false, true},
		{0xe3, false, false, false, true, 0x43, false, false, false, true},
		{0xe3, false, false, true, false, 0x49, false, false, false, true},
		{0xe3, false, false, true, true, 0x49, false, false, false, true},
		{0xe3, false, true, false, false, 0xe3, false, true, false, false},
		{0xe3, false, true, false, true, 0x83, false, true, false, true},
		{0xe3, false, true, true, false, 0xdd, false, true, false, false},
		{0xe3, false, true, true, true, 0x7d, false, true, false, true},
		{0xe3, true, false, false, false, 0x43, false, false, false, true},
		{0xe3, true, false, false, true, 0x43, false, false, false, true},
		{0xe3, true, false, true, false, 0x49, false, false, false, true},
		{0xe3, true, false, true, true, 0x49, false, false, false, true},
		{0xe3, true, true, false, false, 0xe3, false, true, false, false},
		{0xe3, true, true, false, true, 0x83, false, true, false, true},
		{0xe3, true, true, true, false, 0xdd, false, true, false, false},
		{0xe3, true, true, true, true, 0x7d, false, true, false, true},
		{0xe4, false, false, false, false, 0x44, false, false, false, true},
		{0xe4, false, false, false, true, 0x44, false, false, false, true},
		{0xe4, false, false, true, false, 0x4a, false, false, false, true},
		{0xe4, false, false, true, true, 0x4a, false, false, false, true},
		{0xe4, false, true, false, false, 0xe4, false, true, false, false},
		{0xe4, false, true, false, true, 0x84, false, true, false, true},
		{0xe4, false, true, true, false, 0xde, false, true, false, false},
		{0xe4, false, true, true, true, 0x7e, false, true, false, true},
		{0xe4, true, false, false, false, 0x44, false, false, false, true},
		{0xe4, true, false, false, true, 0x44, false, false, false, true},
		{0xe4, true, false, true, false, 0x4a, false, false, false, true},
		{0xe4, true, false, true, true, 0x4a, false, false, false, true},
		{0xe4, true, true, false, false, 0xe4, false, true, false, false},
		{0xe4, true, true, false, true, 0x84, false, true, false, true},
		{0xe4, true, true, true, false, 0xde, false, true, false, false},
		{0xe4, true, true, true, true, 0x7e, false, true, false, true},
		{0xe5, false, false, false, false, 0x45, false, false, false, true},
		{0xe5, false, false, false, true, 0x45, false, false, false, true},
		{0xe5, false, false, true, false, 0x4b, false, false, false, true},
		{0xe5, false, false, true, true, 0x4b, false, false, false, true},
		{0xe5, false, true, false, false, 0xe5, false, true, false, false},
		{0xe5, false, true, false, true, 0x85, false, true, false, true},
		{0xe5, false, true, true, false, 0xdf, false, true, false, false},
		{0xe5, false, true, true, true, 0x7f, false, true, false, true},
		{0xe5, true, false, false, false, 0x45, false, false, false, true},
		{0xe5, true, false, false, true, 0x45, false, false, false, true},
		{0xe5, true, false, true, false, 0x4b, false, false, false, true},
		{0xe5, true, false, true, true, 0x4b, false, false, false, true},
		{0xe5, true, true, false, false, 0xe5, false, true, false, false},
		{0xe5, true, true, false, true, 0x85, false, true, false, true},
		{0xe5, true, true, true, false, 0xdf, false, true, false, false},
		{0xe5, true, true, true, true, 0x7f, false, true, false, true},
		{0xe6, false, false, false, false, 0x46, false, false, false, true},
		{0xe6, false, false, false, true, 0x46, false, false, false, true},
		{0xe6, false, false, true, false, 0x4c, false, false, false, true},
		{0xe6, false, false, true, true, 0x4c, false, false, false, true},
		{0xe6, false, true, false, false, 0xe6, false, true, false, false},
		{0xe6, false, true, false, true, 0x86, false, true, false, true},
		{0xe6, false, true, true, false, 0xe0, false, true, false, false},
		{0xe6, false, true, true, true, 0x80, false, true, false, true},
		{0xe6, true, false, false, false, 0x46, false, false, false, true},
		{0xe6, true, false, false, true, 0x46, false, false, false, true},
		{0xe6, true, false, true, false, 0x4c, false, false, false, true},
		{0xe6, true, false, true, true, 0x4c, false, false, false, true},
		{0xe6, true, true, false, false, 0xe6, false, true, false, false},
		{0xe6, true, true, false, true, 0x86, false, true, false, true},
		{0xe6, true, true, true, false, 0xe0, false, true, false, false},
		{0xe6, true, true, true, true, 0x80, false, true, false, true},
		{0xe7, false, false, false, false, 0x47, false, false, false, true},
		{0xe7, false, false, false, true, 0x47, false, false, false, true},
		{0xe7, false, false, true, false, 0x4d, false, false, false, true},
		{0xe7, false, false, true, true, 0x4d, false, false, false, true},
		{0xe7, false, true, false, false, 0xe7, false, true, false, false},
		{0xe7, false, true, false, true, 0x87, false, true, false, true},
		{0xe7, false, true, true, false, 0xe1, false, true, false, false},
		{0xe7, false, true, true, true, 0x81, false, true, false, true},
		{0xe7, true, false, false, false, 0x47, false, false, false, true},
		{0xe7, true, false, false, true, 0x47, false, false, false, true},
		{0xe7, true, false, true, false, 0x4d, false, false, false, true},
		{0xe7, true, false, true, true, 0x4d, false, false, false, true},
		{0xe7, true, true, false, false, 0xe7, false, true, false, false},
		{0xe7, true, true, false, true, 0x87, false, true, false, true},
		{0xe7, true, true, true, false, 0xe1, false, true, false, false},
		{0xe7, true, true, true, true, 0x81, false, true, false, true},
		{0xe8, false, false, false, false, 0x48, false, false, false, true},
		{0xe8, false, false, false, true, 0x48, false, false, false, true},
		{0xe8, false, false, true, false, 0x4e, false, false, false, true},
		{0xe8, false, false, true, true, 0x4e, false, false, false, true},
		{0xe8, false, true, false, false, 0xe8, false, true, false, false},
		{0xe8, false, true, false, true, 0x88, false, true, false, true},
		{0xe8, false, true, true, false, 0xe2, false, true, false, false},
		{0xe8, false, true, true, true, 0x82, false, true, false, true},
		{0xe8, true, false, false, false, 0x48, false, false, false, true},
		{0xe8, true, false, false, true, 0x48, false, false, false, true},
		{0xe8, true, false, true, false, 0x4e, false, false, false, true},
		{0xe8, true, false, true, true, 0x4e, false, false, false, true},
		{0xe8, true, true, false, false, 0xe8, false, true, false, false},
		{0xe8, true, true, false, true, 0x88, false, true, false, true},
		{0xe8, true, true, true, false, 0xe2, false, true, false, false},
		{0xe8, true, true, true, true, 0x82, false, true, false, true},
		{0xe9, false, false, false, false, 0x49, false, false, false, true},
		{0xe9, false, false, false, true, 0x49, false, false, false, true},
		{0xe9, false, false, true, false, 0x4f, false, false, false, true},
		{0xe9, false, false, true, true, 0x4f, false, false, false, true},
		{0xe9, false, true, false, false, 0xe9, false, true, false, false},
		{0xe9, false, true, false, true, 0x89, false, true, false, true},
		{0xe9, false, true, true, false, 0xe3, false, true, false, false},
		{0xe9, false, true, true, true, 0x83, false, true, false, true},
		{0xe9, true, false, false, false, 0x49, false, false, false, true},
		{0xe9, true, false, false, true, 0x49, false, false, false, true},
		{0xe9, true, false, true, false, 0x4f, false, false, false, true},
		{0xe9, true, false, true, true, 0x4f, false, false, false, true},
		{0xe9, true, true, false, false, 0xe9, false, true, false, false},
		{0xe9, true, true, false, true, 0x89, false, true, false, true},
		{0xe9, true, true, true, false, 0xe3, false, true, false, false},
		{0xe9, true, true, true, true, 0x83, false, true, false, true},
		{0xea, false, false, false, false, 0x50, false, false, false, true},
		{0xea, false, false, false, true, 0x50, false, false, false, true},
		{0xea, false, false, true, false, 0x50, false, false, false, true},
		{0xea, false, false, true, true, 0x50, false, false, false, true},
		{0xea, false, true, false, false, 0xea, false, true, false, false},
		{0xea, false, true, false, true, 0x8a, false, true, false, true},
		{0xea, false, true, true, false, 0xe4, false, true, false, false},
		{0xea, false, true, true, true, 0x84, false, true, false, true},
		{0xea, true, false, false, false, 0x50, false, false, false, true},
		{0xea, true, false, false, true, 0x50, false, false, false, true},
		{0xea, true, false, true, false, 0x50, false, false, false, true},
		{0xea, true, false, true, true, 0x50, false, false, false, true},
		{0xea, true, true, false, false, 0xea, false, true, false, false},
		{0xea, true, true, false, true, 0x8a, false, true, false, true},
		{0xea, true, true, true, false, 0xe4, false, true, false, false},
		{0xea, true, true, true, true, 0x84, false, true, false, true},
		{0xeb, false, false, false, false, 0x51, false, false, false, true},
		{0xeb, false, false, false, true, 0x51, false, false, false, true},
		{0xeb, false, false, true, false, 0x51, false, false, false, true},
		{0xeb, false, false, true, true, 0x51, false, false, false, true},
		{0xeb, false, true, false, false, 0xeb, false, true, false, false},
		{0xeb, false, true, false, true, 0x8b, false, true, false, true},
		{0xeb, false, true, true, false, 0xe5, false, true, false, false},
		{0xeb, false, true, true, true, 0x85, false, true, false, true},
		{0xeb, true, false, false, false, 0x51, false, false, false, true},
		{0xeb, true, false, false, true, 0x51, false, false, false, true},
		{0xeb, true, false, true, false, 0x51, false, false, false, true},
		{0xeb, true, false, true, true, 0x51, false, false, false, true},
		{0xeb, true, true, false, false, 0xeb, false, true, false, false},
		{0xeb, true, true, false, true, 0x8b, false, true, false, true},
		{0xeb, true, true, true, false, 0xe5, false, true, false, false},
		{0xeb, true, true, true, true, 0x85, false, true, false, true},
		{0xec, false, false, false, false, 0x52, false, false, false, true},
		{0xec, false, false, false, true, 0x52, false, false, false, true},
		{0xec, false, false, true, false, 0x52, false, false, false, true},
		{0xec, false, false, true, true, 0x52, false, false, false, true},
		{0xec, false, true, false, false, 0xec, false, true, false, false},
		{0xec, false, true, false, true, 0x8c, false, true, false, true},
		{0xec, false, true, true, false, 0xe6, false, true, false, false},
		{0xec, false, true, true, true, 0x86, false, true, false, true},
		{0xec, true, false, false, false, 0x52, false, false, false, true},
		{0xec, true, false, false, true, 0x52, false, false, false, true},
		{0xec, true, false, true, false, 0x52, false, false, false, true},
		{0xec, true, false, true, true, 0x52, false, false, false, true},
		{0xec, true, true, false, false, 0xec, false, true, false, false},
		{0xec, true, true, false, true, 0x8c, false, true, false, true},
		{0xec, true, true, true, false, 0xe6, false, true, false, false},
		{0xec, true, true, true, true, 0x86, false, true, false, true},
		{0xed, false, false, false, false, 0x53, false, false, false, true},
		{0xed, false, false, false, true, 0x53, false, false, false, true},
		{0xed, false, false, true, false, 0x53, false, false, false, true},
		{0xed, false, false, true, true, 0x53, false, false, false, true},
		{0xed, false, true, false, false, 0xed, false, true, false, false},
		{0xed, false, true, false, true, 0x8d, false, true, false, true},
		{0xed, false, true, true, false, 0xe7, false, true, false, false},
		{0xed, false, true, true, true, 0x87, false, true, false, true},
		{0xed, true, false, false, false, 0x53, false, false, false, true},
		{0xed, true, false, false, true, 0x53, false, false, false, true},
		{0xed, true, false, true, false, 0x53, false, false, false, true},
		{0xed, true, false, true, true, 0x53, false, false, false, true},
		{0xed, true, true, false, false, 0xed, false, true, false, false},
		{0xed, true, true, false, true, 0x8d, false, true, false, true},
		{0xed, true, true, true, false, 0xe7, false, true, false, false},
		{0xed, true, true, true, true, 0x87, false, true, false, true},
		{0xee, false, false, false, false, 0x54, false, false, false, true},
		{0xee, false, false, false, true, 0x54, false, false, false, true},
		{0xee, false, false, true, false, 0x54, false, false, false, true},
		{0xee, false, false, true, true, 0x54, false, false, false, true},
		{0xee, false, true, false, false, 0xee, false, true, false, false},
		{0xee, false, true, false, true, 0x8e, false, true, false, true},
		{0xee, false, true, true, false, 0xe8, false, true, false, false},
		{0xee, false, true, true, true, 0x88, false, true, false, true},
		{0xee, true, false, false, false, 0x54, false, false, false, true},
		{0xee, true, false, false, true, 0x54, false, false, false, true},
		{0xee, true, false, true, false, 0x54, false, false, false, true},
		{0xee, true, false, true, true, 0x54, false, false, false, true},
		{0xee, true, true, false, false, 0xee, false, true, false, false},
		{0xee, true, true, false, true, 0x8e, false, true, false, true},
		{0xee, true, true, true, false, 0xe8, false, true, false, false},
		{0xee, true, true, true, true, 0x88, false, true, false, true},
		{0xef, false, false, false, false, 0x55, false, false, false, true},
		{0xef, false, false, false, true, 0x55, false, false, false, true},
		{0xef, false, false, true, false, 0x55, false, false, false, true},
		{0xef, false, false, true, true, 0x55, false, false, false, true},
		{0xef, false, true, false, false, 0xef, false, true, false, false},
		{0xef, false, true, false, true, 0x8f, false, true, false, true},
		{0xef, false, true, true, false, 0xe9, false, true, false, false},
		{0xef, false, true, true, true, 0x89, false, true, false, true},
		{0xef, true, false, false, false, 0x55, false, false, false, true},
		{0xef, true, false, false, true, 0x55, false, false, false, true},
		{0xef, true, false, true, false, 0x55, false, false, false, true},
		{0xef, true, false, true, true, 0x55, false, false, false, true},
		{0xef, true, true, false, false, 0xef, false, true, false, false},
		{0xef, true, true, false, true, 0x8f, false, true, false, true},
		{0xef, true, true, true, false, 0xe9, false, true, false, false},
		{0xef, true, true, true, true, 0x89, false, true, false, true},
		{0xf0, false, false, false, false, 0x50, false, false, false, true},
		{0xf0, false, false, false, true, 0x50, false, false, false, true},
		{0xf0, false, false, true, false, 0x56, false, false, false, true},
		{0xf0, false, false, true, true, 0x56, false, false, false, true},
		{0xf0, false, true, false, false, 0xf0, false, true, false, false},
		{0xf0, false, true, false, true, 0x90, false, true, false, true},
		{0xf0, false, true, true, false, 0xea, false, true, false, false},
		{0xf0, false, true, true, true, 0x8a, false, true, false, true},
		{0xf0, true, false, false, false, 0x50, false, false, false, true},
		{0xf0, true, false, false, true, 0x50, false, false, false, true},
		{0xf0, true, false, true, false, 0x56, false, false, false, true},
		{0xf0, true, false, true, true, 0x56, false, false, false, true},
		{0xf0, true, true, false, false, 0xf0, false, true, false, false},
		{0xf0, true, true, false, true, 0x90, false, true, false, true},
		{0xf0, true, true, true, false, 0xea, false, true, false, false},
		{0xf0, true, true, true, true, 0x8a, false, true, false, true},
		{0xf1, false, false, false, false, 0x51, false, false, false, true},
		{0xf1, false, false, false, true, 0x51, false, false, false, true},
		{0xf1, false, false, true, false, 0x57, false, false, false, true},
		{0xf1, false, false, true, true, 0x57, false, false, false, true},
		{0xf1, false, true, false, false, 0xf1, false, true, false, false},
		{0xf1, false, true, false, true, 0x91, false, true, false, true},
		{0xf1, false, true, true, false, 0xeb, false, true, false, false},
		{0xf1, false, true, true, true, 0x8b, false, true, false, true},
		{0xf1, true, false, false, false, 0x51, false, false, false, true},
		{0xf1, true, false, false, true, 0x51, false, false, false, true},
		{0xf1, true, false, true, false, 0x57, false, false, false, true},
		{0xf1, true, false, true, true, 0x57, false, false, false, true},
		{0xf1, true, true, false, false, 0xf1, false, true, false, false},
		{0xf1, true, true, false, true, 0x91, false, true, false, true},
		{0xf1, true, true, true, false, 0xeb, false, true, false, false},
		{0xf1, true, true, true, true, 0x8b, false, true, false, true},
		{0xf2, false, false, false, false, 0x52, false, false, false, true},
		{0xf2, false, false, false, true, 0x52, false, false, false, true},
		{0xf2, false, false, true, false, 0x58, false, false, false, true},
		{0xf2, false, false, true, true, 0x58, false, false, false, true},
		{0xf2, false, true, false, false, 0xf2, false, true, false, false},
		{0xf2, false, true, false, true, 0x92, false, true, false, true},
		{0xf2, false, true, true, false, 0xec, false, true, false, false},
		{0xf2, false, true, true, true, 0x8c, false, true, false, true},
		{0xf2, true, false, false, false, 0x52, false, false, false, true},
		{0xf2, true, false, false, true, 0x52, false, false, false, true},
		{0xf2, true, false, true, false, 0x58, false, false, false, true},
		{0xf2, true, false, true, true, 0x58, false, false, false, true},
		{0xf2, true, true, false, false, 0xf2, false, true, false, false},
		{0xf2, true, true, false, true, 0x92, false, true, false, true},
		{0xf2, true, true, true, false, 0xec, false, true, false, false},
		{0xf2, true, true, true, true, 0x8c, false, true, false, true},
		{0xf3, false, false, false, false, 0x53, false, false, false, true},
		{0xf3, false, false, false, true, 0x53, false, false, false, true},
		{0xf3, false, false, true, false, 0x59, false, false, false, true},
		{0xf3, false, false, true, true, 0x59, false, false, false, true},
		{0xf3, false, true, false, false, 0xf3, false, true, false, false},
		{0xf3, false, true, false, true, 0x93, false, true, false, true},
		{0xf3, false, true, true, false, 0xed, false, true, false, false},
		{0xf3, false, true, true, true, 0x8d, false, true, false, true},
		{0xf3, true, false, false, false, 0x53, false, false, false, true},
		{0xf3, true, false, false, true, 0x53, false, false, false, true},
		{0xf3, true, false, true, false, 0x59, false, false, false, true},
		{0xf3, true, false, true, true, 0x59, false, false, false, true},
		{0xf3, true, true, false, false, 0xf3, false, true, false, false},
		{0xf3, true, true, false, true, 0x93, false, true, false, true},
		{0xf3, true, true, true, false, 0xed, false, true, false, false},
		{0xf3, true, true, true, true, 0x8d, false, true, false, true},
		{0xf4, false, false, false, false, 0x54, false, false, false, true},
		{0xf4, false, false, false, true, 0x54, false, false, false, true},
		{0xf4, false, false, true, false, 0x5a, false, false, false, true},
		{0xf4, false, false, true, true, 0x5a, false, false, false, true},
		{0xf4, false, true, false, false, 0xf4, false, true, false, false},
		{0xf4, false, true, false, true, 0x94, false, true, false, true},
		{0xf4, false, true, true, false, 0xee, false, true, false, false},
		{0xf4, false, true, true, true, 0x8e, false, true, false, true},
		{0xf4, true, false, false, false, 0x54, false, false, false, true},
		{0xf4, true, false, false, true, 0x54, false, false, false, true},
		{0xf4, true, false, true, false, 0x5a, false, false, false, true},
		{0xf4, true, false, true, true, 0x5a, false, false, false, true},
		{0xf4, true, true, false, false, 0xf4, false, true, false, false},
		{0xf4, true, true, false, true, 0x94, false, true, false, true},
		{0xf4, true, true, true, false, 0xee, false, true, false, false},
		{0xf4, true, true, true, true, 0x8e, false, true, false, true},
		{0xf5, false, false, false, false, 0x55, false, false, false, true},
		{0xf5, false, false, false, true, 0x55, false, false, false, true},
		{0xf5, false, false, true, false, 0x5b, false, false, false, true},
		{0xf5, false, false, true, true, 0x5b, false, false, false, true},
		{0xf5, false, true, false, false, 0xf5, false, true, false, false},
		{0xf5, false, true, false, true, 0x95, false, true, false, true},
		{0xf5, false, true, true, false, 0xef, false, true, false, false},
		{0xf5, false, true, true, true, 0x8f, false, true, false, true},
		{0xf5, true, false, false, false, 0x55, false, false, false, true},
		{0xf5, true, false, false, true, 0x55, false, false, false, true},
		{0xf5, true, false, true, false, 0x5b, false, false, false, true},
		{0xf5, true, false, true, true, 0x5b, false, false, false, true},
		{0xf5, true, true, false, false, 0xf5, false, true, false, false},
		{0xf5, true, true, false, true, 0x95, false, true, false, true},
		{0xf5, true, true, true, false, 0xef, false, true, false, false},
		{0xf5, true, true, true, true, 0x8f, false, true, false, true},
		{0xf6, false, false, false, false, 0x56, false, false, false, true},
		{0xf6, false, false, false, true, 0x56, false, false, false, true},
		{0xf6, false, false, true, false, 0x5c, false, false, false, true},
		{0xf6, false, false, true, true, 0x5c, false, false, false, true},
		{0xf6, false, true, false, false, 0xf6, false, true, false, false},
		{0xf6, false, true, false, true, 0x96, false, true, false, true},
		{0xf6, false, true, true, false, 0xf0, false, true, false, false},
		{0xf6, false, true, true, true, 0x90, false, true, false, true},
		{0xf6, true, false, false, false, 0x56, false, false, false, true},
		{0xf6, true, false, false, true, 0x56, false, false, false, true},
		{0xf6, true, false, true, false, 0x5c, false, false, false, true},
		{0xf6, true, false, true, true, 0x5c, false, false, false, true},
		{0xf6, true, true, false, false, 0xf6, false, true, false, false},
		{0xf6, true, true, false, true, 0x96, false, true, false, true},
		{0xf6, true, true, true, false, 0xf0, false, true, false, false},
		{0xf6, true, true, true, true, 0x90, false, true, false, true},
		{0xf7, false, false, false, false, 0x57, false, false, false, true},
		{0xf7, false, false, false, true, 0x57, false, false, false, true},
		{0xf7, false, false, true, false, 0x5d, false, false, false, true},
		{0xf7, false, false, true, true, 0x5d, false, false, false, true},
		{0xf7, false, true, false, false, 0xf7, false, true, false, false},
		{0xf7, false, true, false, true, 0x97, false, true, false, true},
		{0xf7, false, true, true, false, 0xf1, false, true, false, false},
		{0xf7, false, true, true, true, 0x91, false, true, false, true},
		{0xf7, true, false, false, false, 0x57, false, false, false, true},
		{0xf7, true, false, false, true, 0x57, false, false, false, true},
		{0xf7, true, false, true, false, 0x5d, false, false, false, true},
		{0xf7, true, false, true, true, 0x5d, false, false, false, true},
		{0xf7, true, true, false, false, 0xf7, false, true, false, false},
		{0xf7, true, true, false, true, 0x97, false, true, false, true},
		{0xf7, true, true, true, false, 0xf1, false, true, false, false},
		{0xf7, true, true, true, true, 0x91, false, true, false, true},
		{0xf8, false, false, false, false, 0x58, false, false, false, true},
		{0xf8, false, false, false, true, 0x58, false, false, false, true},
		{0xf8, false, false, true, false, 0x5e, false, false, false, true},
		{0xf8, false, false, true, true, 0x5e, false, false, false, true},
		{0xf8, false, true, false, false, 0xf8, false, true, false, false},
		{0xf8, false, true, false, true, 0x98, false, true, false, true},
		{0xf8, false, true, true, false, 0xf2, false, true, false, false},
		{0xf8, false, true, true, true, 0x92, false, true, false, true},
		{0xf8, true, false, false, false, 0x58, false, false, false, true},
		{0xf8, true, false, false, true, 0x58, false, false, false, true},
		{0xf8, true, false, true, false, 0x5e, false, false, false, true},
		{0xf8, true, false, true, true, 0x5e, false, false, false, true},
		{0xf8, true, true, false, false, 0xf8, false, true, false, false},
		{0xf8, true, true, false, true, 0x98, false, true, false, true},
		{0xf8, true, true, true, false, 0xf2, false, true, false, false},
		{0xf8, true, true, true, true, 0x92, false, true, false, true},
		{0xf9, false, false, false, false, 0x59, false, false, false, true},
		{0xf9, false, false, false, true, 0x59, false, false, false, true},
		{0xf9, false, false, true, false, 0x5f, false, false, false, true},
		{0xf9, false, false, true, true, 0x5f, false, false, false, true},
		{0xf9, false, true, false, false, 0xf9, false, true, false, false},
		{0xf9, false, true, false, true, 0x99, false, true, false, true},
		{0xf9, false, true, true, false, 0xf3, false, true, false, false},
		{0xf9, false, true, true, true, 0x93, false, true, false, true},
		{0xf9, true, false, false, false, 0x59, false, false, false, true},
		{0xf9, true, false, false, true, 0x59, false, false, false, true},
		{0xf9, true, false, true, false, 0x5f, false, false, false, true},
		{0xf9, true, false, true, true, 0x5f, false, false, false, true},
		{0xf9, true, true, false, false, 0xf9, false, true, false, false},
		{0xf9, true, true, false, true, 0x99, false, true, false, true},
		{0xf9, true, true, true, false, 0xf3, false, true, false, false},
		{0xf9, true, true, true, true, 0x93, false, true, false, true},
		{0xfa, false, false, false, false, 0x60, false, false, false, true},
		{0xfa, false, false, false, true, 0x60, false, false, false, true},
		{0xfa, false, false, true, false, 0x60, false, false, false, true},
		{0xfa, false, false, true, true, 0x60, false, false, false, true},
		{0xfa, false, true, false, false, 0xfa, false, true, false, false},
		{0xfa, false, true, false, true, 0x9a, false, true, false, true},
		{0xfa, false, true, true, false, 0xf4, false, true, false, false},
		{0xfa, false, true, true, true, 0x94, false, true, false, true},
		{0xfa, true, false, false, false, 0x60, false, false, false, true},
		{0xfa, true, false, false, true, 0x60, false, false, false, true},
		{0xfa, true, false, true, false, 0x60, false, false, false, true},
		{0xfa, true, false, true, true, 0x60, false, false, false, true},
		{0xfa, true, true, false, false, 0xfa, false, true, false, false},
		{0xfa, true, true, false, true, 0x9a, false, true, false, true},
		{0xfa, true, true, true, false, 0xf4, false, true, false, false},
		{0xfa, true, true, true, true, 0x94, false, true, false, true},
		{0xfb, false, false, false, false, 0x61, false, false, false, true},
		{0xfb, false, false, false, true, 0x61, false, false, false, true},
		{0xfb, false, false, true, false, 0x61, false, false, false, true},
		{0xfb, false, false, true, true, 0x61, false, false, false, true},
		{0xfb, false, true, false, false, 0xfb, false, true, false, false},
		{0xfb, false, true, false, true, 0x9b, false, true, false, true},
		{0xfb, false, true, true, false, 0xf5, false, true, false, false},
		{0xfb, false, true, true, true, 0x95, false, true, false, true},
		{0xfb, true, false, false, false, 0x61, false, false, false, true},
		{0xfb, true, false, false, true, 0x61, false, false, false, true},
		{0xfb, true, false, true, false, 0x61, false, false, false, true},
		{0xfb, true, false, true, true, 0x61, false, false, false, true},
		{0xfb, true, true, false, false, 0xfb, false, true, false, false},
		{0xfb, true, true, false, true, 0x9b, false, true, false, true},
		{0xfb, true, true, true, false, 0xf5, false, true, false, false},
		{0xfb, true, true, true, true, 0x95, false, true, false, true},
		{0xfc, false, false, false, false, 0x62, false, false, false, true},
		{0xfc, false, false, false, true, 0x62, false, false, false, true},
		{0xfc, false, false, true, false, 0x62, false, false, false, true},
		{0xfc, false, false, true, true, 0x62, false, false, false, true},
		{0xfc, false, true, false, false, 0xfc, false, true, false, false},
		{0xfc, false, true, false, true, 0x9c, false, true, false, true},
		{0xfc, false, true, true, false, 0xf6, false, true, false, false},
		{0xfc, false, true, true, true, 0x96, false, true, false, true},
		{0xfc, true, false, false, false, 0x62, false, false, false, true},
		{0xfc, true, false, false, true, 0x62, false, false, false, true},
		{0xfc, true, false, true, false, 0x62, false, false, false, true},
		{0xfc, true, false, true, true, 0x62, false, false, false, true},
		{0xfc, true, true, false, false, 0xfc, false, true, false, false},
		{0xfc, true, true, false, true, 0x9c, false, true, false, true},
		{0xfc, true, true, true, false, 0xf6, false, true, false, false},
		{0xfc, true, true, true, true, 0x96, false, true, false, true},
		{0xfd, false, false, false, false, 0x63, false, false, false, true},
		{0xfd, false, false, false, true, 0x63, false, false, false, true},
		{0xfd, false, false, true, false, 0x63, false, false, false, true},
		{0xfd, false, false, true, true, 0x63, false, false, false, true},
		{0xfd, false, true, false, false, 0xfd, false, true, false, false},
		{0xfd, false, true, false, true, 0x9d, false, true, false, true},
		{0xfd, false, true, true, false, 0xf7, false, true, false, false},
		{0xfd, false, true, true, true, 0x97, false, true, false, true},
		{0xfd, true, false, false, false, 0x63, false, false, false, true},
		{0xfd, true, false, false, true, 0x63, false, false, false, true},
		{0xfd, true, false, true, false, 0x63, false, false, false, true},
		{0xfd, true, false, true, true, 0x63, false, false, false, true},
		{0xfd, true, true, false, false, 0xfd, false, true, false, false},
		{0xfd, true, true, false, true, 0x9d, false, true, false, true},
		{0xfd, true, true, true, false, 0xf7, false, true, false, false},
		{0xfd, true, true, true, true, 0x97, false, true, false, true},
		{0xfe, false, false, false, false, 0x64, false, false, false, true},
		{0xfe, false, false, false, true, 0x64, false, false, false, true},
		{0xfe, false, false, true, false, 0x64, false, false, false, true},
		{0xfe, false, false, true, true, 0x64, false, false, false, true},
		{0xfe, false, true, false, false, 0xfe, false, true, false, false},
		{0xfe, false, true, false, true, 0x9e, false, true, false, true},
		{0xfe, false, true, true, false, 0xf8, false, true, false, false},
		{0xfe, false, true, true, true, 0x98, false, true, false, true},
		{0xfe, true, false, false, false, 0x64, false, false, false, true},
		{0xfe, true, false, false, true, 0x64, false, false, false, true},
		{0xfe, true, false, true, false, 0x64, false, false, false, true},
		{0xfe, true, false, true, true, 0x64, false, false, false, true},
		{0xfe, true, true, false, false, 0xfe, false, true, false, false},
		{0xfe, true, true, false, true, 0x9e, false, true, false, true},
		{0xfe, true, true, true, false, 0xf8, false, true, false, false},
		{0xfe, true, true, true, true, 0x98, false, true, false, true},
		{0xff, false, false, false, false, 0x65, false, false, false, true},
		{0xff, false, false, false, true, 0x65, false, false, false, true},
		{0xff, false, false, true, false, 0x65, false, false, false, true},
		{0xff, false, false, true, true, 0x65, false, false, false, true},
		{0xff, false, true, false, false, 0xff, false, true, false, false},
		{0xff, false, true, false, true, 0x9f, false, true, false, true},
		{0xff, false, true, true, false, 0xf9, false, true, false, false},
		{0xff, false, true, true, true, 0x99, false, true, false, true},
		{0xff, true, false, false, false, 0x65, false, false, false, true},
		{0xff, true, false, false, true, 0x65, false, false, false, true},
		{0xff, true, false, true, false, 0x65, false, false, false, true},
		{0xff, true, false, true, true, 0x65, false, false, false, true},
		{0xff, true, true, false, false, 0xff, false, true, false, false},
		{0xff, true, true, false, true, 0x9f, false, true, false, true},
		{0xff, true, true, true, false, 0xf9, false, true, false, false},
		{0xff, true, true, true, true, 0x99, false, true, false, true},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			for _, operation := range operations {
				cpu.registers.setA(operation.PrevA)
				cpu.registers.f.setZ(operation.PrevZ)
				cpu.registers.f.setN(operation.PrevN)
				cpu.registers.f.setH(operation.PrevH)
				cpu.registers.f.setC(operation.PrevC)
				tc.instruction.exec(cpu)
				if operation.CurrA != cpu.registers.getA() {
					t.Errorf("Expected 0x%02X, got 0x%02X", operation.CurrA, cpu.registers.getA())
				}
				if operation.CurrZ != cpu.registers.f.getZ() {
					t.Errorf("Flag Z. Expected %t, got %t", operation.CurrZ, cpu.registers.f.getZ())
				}
				if operation.CurrN != cpu.registers.f.getN() {
					t.Errorf("Flag N. Expected %t, got %t", operation.CurrN, cpu.registers.f.getN())
				}
				if operation.CurrH != cpu.registers.f.getH() {
					t.Errorf("Flag H. Expected %t, got %t", operation.CurrH, cpu.registers.f.getH())
				}
				if operation.CurrC != cpu.registers.f.getC() {
					t.Errorf("Flag C. Expected %t, got %t", operation.CurrC, cpu.registers.f.getC())
				}
			}
		})
	}
}

// CPL
func TestCPLInstructions(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	instructions := []struct {
		instruction instruction
	}{
		{cpu.instructions[0x2F]},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			value1 := cpu.registers.getA() ^ 0xFF
			tc.instruction.exec(cpu)
			value2 := cpu.registers.getA()
			flagN := true
			flagH := true
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
			if flagN != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
		})
	}
}

// RLCA
func TestRLCA(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	instructions := []struct {
		instruction instruction
	}{
		{cpu.instructions[0x07]},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			bit7 := cpu.registers.getA() >> 7
			value1 := (cpu.registers.getA() << 1) + bit7
			tc.instruction.exec(cpu)
			value2 := cpu.registers.getA()
			flagC := bit7 == 1
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
			if false != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", false, cpu.registers.f.getZ())
			}
			if false != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", false, cpu.registers.f.getN())
			}
			if false != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", false, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// RLA
func TestRLA(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	instructions := []struct {
		instruction instruction
	}{
		{cpu.instructions[0x17]},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			carry := cpu.registers.f.getCarry()
			bit7 := cpu.registers.getA() >> 7
			value1 := (cpu.registers.getA() << 1) + carry
			tc.instruction.exec(cpu)
			value2 := cpu.registers.getA()
			flagC := bit7 == 1
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
			if false != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", false, cpu.registers.f.getZ())
			}
			if false != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", false, cpu.registers.f.getN())
			}
			if false != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", false, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// RRCA
func TestRRCA(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	instructions := []struct {
		instruction instruction
	}{
		{cpu.instructions[0x0F]},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			bit0 := cpu.registers.getA() << 7
			value1 := (cpu.registers.getA() >> 1) + bit0
			tc.instruction.exec(cpu)
			value2 := cpu.registers.getA()
			flagC := bit0 == 128
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
			if false != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", false, cpu.registers.f.getZ())
			}
			if false != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", false, cpu.registers.f.getN())
			}
			if false != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", false, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// RRA
func TestRRA(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	instructions := []struct {
		instruction instruction
	}{
		{cpu.instructions[0x1F]},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			carry := cpu.registers.f.getCarry() << 7
			bit0 := cpu.registers.getA() << 7
			value1 := (cpu.registers.getA() >> 1) + carry
			tc.instruction.exec(cpu)
			value2 := cpu.registers.getA()
			flagC := bit0 == 128
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
			if false != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", false, cpu.registers.f.getZ())
			}
			if false != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", false, cpu.registers.f.getN())
			}
			if false != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", false, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// RLC r
func TestRLCInstructions1(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	instructions := []struct {
		instruction instruction
		to          func() uint8
	}{
		{cpu.instructions[0xCB00], cpu.registers.getB},
		{cpu.instructions[0xCB01], cpu.registers.getC},
		{cpu.instructions[0xCB02], cpu.registers.getD},
		{cpu.instructions[0xCB03], cpu.registers.getE},
		{cpu.instructions[0xCB04], cpu.registers.getH},
		{cpu.instructions[0xCB05], cpu.registers.getL},
		{cpu.instructions[0xCB07], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			bit7 := tc.to() >> 7
			value1 := (tc.to() << 1) + bit7
			tc.instruction.exec(cpu)
			value2 := tc.to()
			flagZ := value2 == 0
			flagC := bit7 == 1
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if false != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", false, cpu.registers.f.getN())
			}
			if false != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", false, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// RLC (HL)
func TestRLCInstructions2(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
	}{
		{cpu.instructions[0xCB06]},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := cpu.registers.getHL()
			bit7 := cpu.memory.Read(addr) >> 7
			value1 := (cpu.memory.Read(addr) << 1) + bit7
			tc.instruction.exec(cpu)
			value2 := cpu.memory.Read(addr)
			flagZ := value2 == 0
			flagC := bit7 == 1
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if false != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", false, cpu.registers.f.getN())
			}
			if false != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", false, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// RRC r
func TestRRCInstructions1(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	instructions := []struct {
		instruction instruction
		to          func() uint8
	}{
		{cpu.instructions[0xCB08], cpu.registers.getB},
		{cpu.instructions[0xCB09], cpu.registers.getC},
		{cpu.instructions[0xCB0A], cpu.registers.getD},
		{cpu.instructions[0xCB0B], cpu.registers.getE},
		{cpu.instructions[0xCB0C], cpu.registers.getH},
		{cpu.instructions[0xCB0D], cpu.registers.getL},
		{cpu.instructions[0xCB0F], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			bit0 := tc.to() << 7
			value1 := (tc.to() >> 1) + bit0
			tc.instruction.exec(cpu)
			value2 := tc.to()
			flagZ := value2 == 0
			flagC := bit0 == 128
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if false != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", false, cpu.registers.f.getN())
			}
			if false != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", false, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// RRC (HL)
func TestRRCInstructions2(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
	}{
		{cpu.instructions[0xCB0E]},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := cpu.registers.getHL()
			bit0 := cpu.memory.Read(addr) << 7
			value1 := (cpu.memory.Read(addr) >> 1) + bit0
			tc.instruction.exec(cpu)
			value2 := cpu.memory.Read(addr)
			flagZ := value2 == 0
			flagC := bit0 == 128
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if false != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", false, cpu.registers.f.getN())
			}
			if false != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", false, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// RL r
func TestRLInstructions1(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	instructions := []struct {
		instruction instruction
		to          func() uint8
	}{
		{cpu.instructions[0xCB10], cpu.registers.getB},
		{cpu.instructions[0xCB11], cpu.registers.getC},
		{cpu.instructions[0xCB12], cpu.registers.getD},
		{cpu.instructions[0xCB13], cpu.registers.getE},
		{cpu.instructions[0xCB14], cpu.registers.getH},
		{cpu.instructions[0xCB15], cpu.registers.getL},
		{cpu.instructions[0xCB17], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			carry := cpu.registers.f.getCarry()
			bit7 := tc.to() >> 7
			value1 := (tc.to() << 1) + carry
			tc.instruction.exec(cpu)
			value2 := tc.to()
			flagZ := value2 == 0
			flagC := bit7 == 1
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if false != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", false, cpu.registers.f.getN())
			}
			if false != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", false, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// RL (HL)
func TestRLInstructions2(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
	}{
		{cpu.instructions[0xCB16]},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := cpu.registers.getHL()
			carry := cpu.registers.f.getCarry()
			bit7 := cpu.memory.Read(addr) >> 7
			value1 := (cpu.memory.Read(addr) << 1) + carry
			tc.instruction.exec(cpu)
			value2 := cpu.memory.Read(addr)
			flagZ := value2 == 0
			flagC := bit7 == 1
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if false != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", false, cpu.registers.f.getN())
			}
			if false != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", false, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// RR r
func TestRRInstructions1(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	instructions := []struct {
		instruction instruction
		to          func() uint8
	}{
		{cpu.instructions[0xCB18], cpu.registers.getB},
		{cpu.instructions[0xCB19], cpu.registers.getC},
		{cpu.instructions[0xCB1A], cpu.registers.getD},
		{cpu.instructions[0xCB1B], cpu.registers.getE},
		{cpu.instructions[0xCB1C], cpu.registers.getH},
		{cpu.instructions[0xCB1D], cpu.registers.getL},
		{cpu.instructions[0xCB1F], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			carry := cpu.registers.f.getCarry() << 7
			bit0 := tc.to() << 7
			value1 := (tc.to() >> 1) + carry
			tc.instruction.exec(cpu)
			value2 := tc.to()
			flagZ := value2 == 0
			flagC := bit0 == 128
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if false != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", false, cpu.registers.f.getN())
			}
			if false != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", false, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// RR (HL)
func TestRRInstructions2(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
	}{
		{cpu.instructions[0xCB1E]},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := cpu.registers.getHL()
			carry := cpu.registers.f.getCarry() << 7
			bit0 := cpu.memory.Read(addr) << 7
			value1 := (cpu.memory.Read(addr) >> 1) + carry
			tc.instruction.exec(cpu)
			value2 := cpu.memory.Read(addr)
			flagZ := value2 == 0
			flagC := bit0 == 128
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if false != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", false, cpu.registers.f.getN())
			}
			if false != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", false, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// SLA r
func TestSLAInstructions1(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	instructions := []struct {
		instruction instruction
		to          func() uint8
	}{
		{cpu.instructions[0xCB20], cpu.registers.getB},
		{cpu.instructions[0xCB21], cpu.registers.getC},
		{cpu.instructions[0xCB22], cpu.registers.getD},
		{cpu.instructions[0xCB23], cpu.registers.getE},
		{cpu.instructions[0xCB24], cpu.registers.getH},
		{cpu.instructions[0xCB25], cpu.registers.getL},
		{cpu.instructions[0xCB27], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			bit7 := tc.to() >> 7
			value1 := tc.to() << 1
			tc.instruction.exec(cpu)
			value2 := tc.to()
			flagZ := value2 == 0
			flagC := bit7 == 1
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if false != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", false, cpu.registers.f.getN())
			}
			if false != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", false, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// SLA (HL)
func TestSLAInstructions2(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
	}{
		{cpu.instructions[0xCB26]},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := cpu.registers.getHL()
			bit7 := cpu.memory.Read(addr) >> 7
			value1 := cpu.memory.Read(addr) << 1
			tc.instruction.exec(cpu)
			value2 := cpu.memory.Read(addr)
			flagZ := value2 == 0
			flagC := bit7 == 1
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if false != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", false, cpu.registers.f.getN())
			}
			if false != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", false, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// SRA r
func TestSRAInstructions1(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	instructions := []struct {
		instruction instruction
		to          func() uint8
	}{
		{cpu.instructions[0xCB28], cpu.registers.getB},
		{cpu.instructions[0xCB29], cpu.registers.getC},
		{cpu.instructions[0xCB2A], cpu.registers.getD},
		{cpu.instructions[0xCB2B], cpu.registers.getE},
		{cpu.instructions[0xCB2C], cpu.registers.getH},
		{cpu.instructions[0xCB2D], cpu.registers.getL},
		{cpu.instructions[0xCB2F], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			bit0 := tc.to() << 7
			bit7 := tc.to() & 128
			value1 := (tc.to() >> 1) + bit7
			tc.instruction.exec(cpu)
			value2 := tc.to()
			flagZ := value2 == 0
			flagC := bit0 == 128
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if false != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", false, cpu.registers.f.getN())
			}
			if false != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", false, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// SRA (HL)
func TestSRAInstructions2(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
	}{
		{cpu.instructions[0xCB2E]},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := cpu.registers.getHL()
			bit0 := cpu.memory.Read(addr) << 7
			bit7 := cpu.memory.Read(addr) & 128
			value1 := (cpu.memory.Read(addr) >> 1) + bit7
			tc.instruction.exec(cpu)
			value2 := cpu.memory.Read(addr)
			flagZ := value2 == 0
			flagC := bit0 == 128
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if false != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", false, cpu.registers.f.getN())
			}
			if false != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", false, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// SWAP r
func TestSWAPInstructions1(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	instructions := []struct {
		instruction instruction
		to          func() uint8
	}{
		{cpu.instructions[0xCB30], cpu.registers.getB},
		{cpu.instructions[0xCB31], cpu.registers.getC},
		{cpu.instructions[0xCB32], cpu.registers.getD},
		{cpu.instructions[0xCB33], cpu.registers.getE},
		{cpu.instructions[0xCB34], cpu.registers.getH},
		{cpu.instructions[0xCB35], cpu.registers.getL},
		{cpu.instructions[0xCB37], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			lo := tc.to() % 16
			hi := tc.to() >> 4
			value1 := lo<<4 | hi
			tc.instruction.exec(cpu)
			value2 := tc.to()
			flagZ := value2 == 0
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if false != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", false, cpu.registers.f.getN())
			}
			if false != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", false, cpu.registers.f.getH())
			}
			if false != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", false, cpu.registers.f.getC())
			}
		})
	}
}

// SWAP (HL)
func TestSWAPInstructions2(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
	}{
		{cpu.instructions[0xCB36]},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := cpu.registers.getHL()
			lo := cpu.memory.Read(addr) % 16
			hi := cpu.memory.Read(addr) >> 4
			value1 := lo<<4 | hi
			tc.instruction.exec(cpu)
			value2 := cpu.memory.Read(addr)
			flagZ := value2 == 0
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if false != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", false, cpu.registers.f.getN())
			}
			if false != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", false, cpu.registers.f.getH())
			}
			if false != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", false, cpu.registers.f.getC())
			}
		})
	}
}

// SRL r
func TestSRLInstructions1(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	instructions := []struct {
		instruction instruction
		to          func() uint8
	}{
		{cpu.instructions[0xCB38], cpu.registers.getB},
		{cpu.instructions[0xCB39], cpu.registers.getC},
		{cpu.instructions[0xCB3A], cpu.registers.getD},
		{cpu.instructions[0xCB3B], cpu.registers.getE},
		{cpu.instructions[0xCB3C], cpu.registers.getH},
		{cpu.instructions[0xCB3D], cpu.registers.getL},
		{cpu.instructions[0xCB3F], cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			bit0 := tc.to() << 7
			value1 := tc.to() >> 1
			tc.instruction.exec(cpu)
			value2 := tc.to()
			flagZ := value2 == 0
			flagC := bit0 == 128
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if false != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", false, cpu.registers.f.getN())
			}
			if false != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", false, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// SRL (HL)
func TestSRLInstructions2(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
	}{
		{cpu.instructions[0xCB3E]},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := cpu.registers.getHL()
			bit0 := cpu.memory.Read(addr) << 7
			value1 := cpu.memory.Read(addr) >> 1
			tc.instruction.exec(cpu)
			value2 := cpu.memory.Read(addr)
			flagZ := value2 == 0
			flagC := bit0 == 128
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if false != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", false, cpu.registers.f.getN())
			}
			if false != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", false, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// BIT n,r
func TestBITInstructions1(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	instructions := []struct {
		instruction instruction
		bitIndex    int
		to          func() uint8
	}{
		{cpu.instructions[0xCB40], 0, cpu.registers.getB},
		{cpu.instructions[0xCB41], 0, cpu.registers.getC},
		{cpu.instructions[0xCB42], 0, cpu.registers.getD},
		{cpu.instructions[0xCB43], 0, cpu.registers.getE},
		{cpu.instructions[0xCB44], 0, cpu.registers.getH},
		{cpu.instructions[0xCB45], 0, cpu.registers.getL},
		{cpu.instructions[0xCB47], 0, cpu.registers.getA},
		{cpu.instructions[0xCB48], 1, cpu.registers.getB},
		{cpu.instructions[0xCB49], 1, cpu.registers.getC},
		{cpu.instructions[0xCB4A], 1, cpu.registers.getD},
		{cpu.instructions[0xCB4B], 1, cpu.registers.getE},
		{cpu.instructions[0xCB4C], 1, cpu.registers.getH},
		{cpu.instructions[0xCB4D], 1, cpu.registers.getL},
		{cpu.instructions[0xCB4F], 1, cpu.registers.getA},
		{cpu.instructions[0xCB50], 2, cpu.registers.getB},
		{cpu.instructions[0xCB51], 2, cpu.registers.getC},
		{cpu.instructions[0xCB52], 2, cpu.registers.getD},
		{cpu.instructions[0xCB53], 2, cpu.registers.getE},
		{cpu.instructions[0xCB54], 2, cpu.registers.getH},
		{cpu.instructions[0xCB55], 2, cpu.registers.getL},
		{cpu.instructions[0xCB57], 2, cpu.registers.getA},
		{cpu.instructions[0xCB58], 3, cpu.registers.getB},
		{cpu.instructions[0xCB59], 3, cpu.registers.getC},
		{cpu.instructions[0xCB5A], 3, cpu.registers.getD},
		{cpu.instructions[0xCB5B], 3, cpu.registers.getE},
		{cpu.instructions[0xCB5C], 3, cpu.registers.getH},
		{cpu.instructions[0xCB5D], 3, cpu.registers.getL},
		{cpu.instructions[0xCB5F], 3, cpu.registers.getA},
		{cpu.instructions[0xCB60], 4, cpu.registers.getB},
		{cpu.instructions[0xCB61], 4, cpu.registers.getC},
		{cpu.instructions[0xCB62], 4, cpu.registers.getD},
		{cpu.instructions[0xCB63], 4, cpu.registers.getE},
		{cpu.instructions[0xCB64], 4, cpu.registers.getH},
		{cpu.instructions[0xCB65], 4, cpu.registers.getL},
		{cpu.instructions[0xCB67], 4, cpu.registers.getA},
		{cpu.instructions[0xCB68], 5, cpu.registers.getB},
		{cpu.instructions[0xCB69], 5, cpu.registers.getC},
		{cpu.instructions[0xCB6A], 5, cpu.registers.getD},
		{cpu.instructions[0xCB6B], 5, cpu.registers.getE},
		{cpu.instructions[0xCB6C], 5, cpu.registers.getH},
		{cpu.instructions[0xCB6D], 5, cpu.registers.getL},
		{cpu.instructions[0xCB6F], 5, cpu.registers.getA},
		{cpu.instructions[0xCB70], 6, cpu.registers.getB},
		{cpu.instructions[0xCB71], 6, cpu.registers.getC},
		{cpu.instructions[0xCB72], 6, cpu.registers.getD},
		{cpu.instructions[0xCB73], 6, cpu.registers.getE},
		{cpu.instructions[0xCB74], 6, cpu.registers.getH},
		{cpu.instructions[0xCB75], 6, cpu.registers.getL},
		{cpu.instructions[0xCB77], 6, cpu.registers.getA},
		{cpu.instructions[0xCB78], 7, cpu.registers.getB},
		{cpu.instructions[0xCB79], 7, cpu.registers.getC},
		{cpu.instructions[0xCB7A], 7, cpu.registers.getD},
		{cpu.instructions[0xCB7B], 7, cpu.registers.getE},
		{cpu.instructions[0xCB7C], 7, cpu.registers.getH},
		{cpu.instructions[0xCB7D], 7, cpu.registers.getL},
		{cpu.instructions[0xCB7F], 7, cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			value1 := (tc.to() & (1 << tc.bitIndex)) == (1 << tc.bitIndex)
			tc.instruction.exec(cpu)
			value2 := !cpu.registers.f.getZ()
			flagZ := !value1
			if value1 != value2 {
				t.Errorf("Expected %v, got %v", value1, value2)
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if false != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", false, cpu.registers.f.getN())
			}
			if true != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", true, cpu.registers.f.getH())
			}
		})
	}
}

// BIT n,(HL)
func TestBITInstructions2(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		bitIndex    int
	}{
		{cpu.instructions[0xCB46], 0},
		{cpu.instructions[0xCB4E], 1},
		{cpu.instructions[0xCB56], 2},
		{cpu.instructions[0xCB5E], 3},
		{cpu.instructions[0xCB66], 4},
		{cpu.instructions[0xCB6E], 5},
		{cpu.instructions[0xCB76], 6},
		{cpu.instructions[0xCB7E], 7},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := cpu.registers.getHL()
			value1 := (cpu.memory.Read(addr) & (1 << tc.bitIndex)) == (1 << tc.bitIndex)
			tc.instruction.exec(cpu)
			value2 := !cpu.registers.f.getZ()
			flagZ := !value1
			if value1 != value2 {
				t.Errorf("Expected %v, got %v", value1, value2)
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if false != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", false, cpu.registers.f.getN())
			}
			if true != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", true, cpu.registers.f.getH())
			}
		})
	}
}

// RES n,r
func TestRESInstructions1(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	instructions := []struct {
		instruction instruction
		bitIndex    int
		to          func() uint8
	}{
		{cpu.instructions[0xCB80], 0, cpu.registers.getB},
		{cpu.instructions[0xCB81], 0, cpu.registers.getC},
		{cpu.instructions[0xCB82], 0, cpu.registers.getD},
		{cpu.instructions[0xCB83], 0, cpu.registers.getE},
		{cpu.instructions[0xCB84], 0, cpu.registers.getH},
		{cpu.instructions[0xCB85], 0, cpu.registers.getL},
		{cpu.instructions[0xCB87], 0, cpu.registers.getA},
		{cpu.instructions[0xCB88], 1, cpu.registers.getB},
		{cpu.instructions[0xCB89], 1, cpu.registers.getC},
		{cpu.instructions[0xCB8A], 1, cpu.registers.getD},
		{cpu.instructions[0xCB8B], 1, cpu.registers.getE},
		{cpu.instructions[0xCB8C], 1, cpu.registers.getH},
		{cpu.instructions[0xCB8D], 1, cpu.registers.getL},
		{cpu.instructions[0xCB8F], 1, cpu.registers.getA},
		{cpu.instructions[0xCB90], 2, cpu.registers.getB},
		{cpu.instructions[0xCB91], 2, cpu.registers.getC},
		{cpu.instructions[0xCB92], 2, cpu.registers.getD},
		{cpu.instructions[0xCB93], 2, cpu.registers.getE},
		{cpu.instructions[0xCB94], 2, cpu.registers.getH},
		{cpu.instructions[0xCB95], 2, cpu.registers.getL},
		{cpu.instructions[0xCB97], 2, cpu.registers.getA},
		{cpu.instructions[0xCB98], 3, cpu.registers.getB},
		{cpu.instructions[0xCB99], 3, cpu.registers.getC},
		{cpu.instructions[0xCB9A], 3, cpu.registers.getD},
		{cpu.instructions[0xCB9B], 3, cpu.registers.getE},
		{cpu.instructions[0xCB9C], 3, cpu.registers.getH},
		{cpu.instructions[0xCB9D], 3, cpu.registers.getL},
		{cpu.instructions[0xCB9F], 3, cpu.registers.getA},
		{cpu.instructions[0xCBA0], 4, cpu.registers.getB},
		{cpu.instructions[0xCBA1], 4, cpu.registers.getC},
		{cpu.instructions[0xCBA2], 4, cpu.registers.getD},
		{cpu.instructions[0xCBA3], 4, cpu.registers.getE},
		{cpu.instructions[0xCBA4], 4, cpu.registers.getH},
		{cpu.instructions[0xCBA5], 4, cpu.registers.getL},
		{cpu.instructions[0xCBA7], 4, cpu.registers.getA},
		{cpu.instructions[0xCBA8], 5, cpu.registers.getB},
		{cpu.instructions[0xCBA9], 5, cpu.registers.getC},
		{cpu.instructions[0xCBAA], 5, cpu.registers.getD},
		{cpu.instructions[0xCBAB], 5, cpu.registers.getE},
		{cpu.instructions[0xCBAC], 5, cpu.registers.getH},
		{cpu.instructions[0xCBAD], 5, cpu.registers.getL},
		{cpu.instructions[0xCBAF], 5, cpu.registers.getA},
		{cpu.instructions[0xCBB0], 6, cpu.registers.getB},
		{cpu.instructions[0xCBB1], 6, cpu.registers.getC},
		{cpu.instructions[0xCBB2], 6, cpu.registers.getD},
		{cpu.instructions[0xCBB3], 6, cpu.registers.getE},
		{cpu.instructions[0xCBB4], 6, cpu.registers.getH},
		{cpu.instructions[0xCBB5], 6, cpu.registers.getL},
		{cpu.instructions[0xCBB7], 6, cpu.registers.getA},
		{cpu.instructions[0xCBB8], 7, cpu.registers.getB},
		{cpu.instructions[0xCBB9], 7, cpu.registers.getC},
		{cpu.instructions[0xCBBA], 7, cpu.registers.getD},
		{cpu.instructions[0xCBBB], 7, cpu.registers.getE},
		{cpu.instructions[0xCBBC], 7, cpu.registers.getH},
		{cpu.instructions[0xCBBD], 7, cpu.registers.getL},
		{cpu.instructions[0xCBBF], 7, cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			value1 := tc.to() & (0xFF - (1 << tc.bitIndex))
			tc.instruction.exec(cpu)
			value2 := tc.to()
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
		})
	}
}

// RES n,(HL)
func TestRESInstructions2(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		bitIndex    int
	}{
		{cpu.instructions[0xCB86], 0},
		{cpu.instructions[0xCB8E], 1},
		{cpu.instructions[0xCB96], 2},
		{cpu.instructions[0xCB9E], 3},
		{cpu.instructions[0xCBA6], 4},
		{cpu.instructions[0xCBAE], 5},
		{cpu.instructions[0xCBB6], 6},
		{cpu.instructions[0xCBBE], 7},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := cpu.registers.getHL()
			value1 := cpu.memory.Read(addr) & (0xFF - (1 << tc.bitIndex))
			tc.instruction.exec(cpu)
			value2 := cpu.memory.Read(addr)
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
		})
	}
}

// SET n,r
func TestSETInstructions1(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	instructions := []struct {
		instruction instruction
		bitIndex    int
		to          func() uint8
	}{
		{cpu.instructions[0xCBC0], 0, cpu.registers.getB},
		{cpu.instructions[0xCBC1], 0, cpu.registers.getC},
		{cpu.instructions[0xCBC2], 0, cpu.registers.getD},
		{cpu.instructions[0xCBC3], 0, cpu.registers.getE},
		{cpu.instructions[0xCBC4], 0, cpu.registers.getH},
		{cpu.instructions[0xCBC5], 0, cpu.registers.getL},
		{cpu.instructions[0xCBC7], 0, cpu.registers.getA},
		{cpu.instructions[0xCBC8], 1, cpu.registers.getB},
		{cpu.instructions[0xCBC9], 1, cpu.registers.getC},
		{cpu.instructions[0xCBCA], 1, cpu.registers.getD},
		{cpu.instructions[0xCBCB], 1, cpu.registers.getE},
		{cpu.instructions[0xCBCC], 1, cpu.registers.getH},
		{cpu.instructions[0xCBCD], 1, cpu.registers.getL},
		{cpu.instructions[0xCBCF], 1, cpu.registers.getA},
		{cpu.instructions[0xCBD0], 2, cpu.registers.getB},
		{cpu.instructions[0xCBD1], 2, cpu.registers.getC},
		{cpu.instructions[0xCBD2], 2, cpu.registers.getD},
		{cpu.instructions[0xCBD3], 2, cpu.registers.getE},
		{cpu.instructions[0xCBD4], 2, cpu.registers.getH},
		{cpu.instructions[0xCBD5], 2, cpu.registers.getL},
		{cpu.instructions[0xCBD7], 2, cpu.registers.getA},
		{cpu.instructions[0xCBD8], 3, cpu.registers.getB},
		{cpu.instructions[0xCBD9], 3, cpu.registers.getC},
		{cpu.instructions[0xCBDA], 3, cpu.registers.getD},
		{cpu.instructions[0xCBDB], 3, cpu.registers.getE},
		{cpu.instructions[0xCBDC], 3, cpu.registers.getH},
		{cpu.instructions[0xCBDD], 3, cpu.registers.getL},
		{cpu.instructions[0xCBDF], 3, cpu.registers.getA},
		{cpu.instructions[0xCBE0], 4, cpu.registers.getB},
		{cpu.instructions[0xCBE1], 4, cpu.registers.getC},
		{cpu.instructions[0xCBE2], 4, cpu.registers.getD},
		{cpu.instructions[0xCBE3], 4, cpu.registers.getE},
		{cpu.instructions[0xCBE4], 4, cpu.registers.getH},
		{cpu.instructions[0xCBE5], 4, cpu.registers.getL},
		{cpu.instructions[0xCBE7], 4, cpu.registers.getA},
		{cpu.instructions[0xCBE8], 5, cpu.registers.getB},
		{cpu.instructions[0xCBE9], 5, cpu.registers.getC},
		{cpu.instructions[0xCBEA], 5, cpu.registers.getD},
		{cpu.instructions[0xCBEB], 5, cpu.registers.getE},
		{cpu.instructions[0xCBEC], 5, cpu.registers.getH},
		{cpu.instructions[0xCBED], 5, cpu.registers.getL},
		{cpu.instructions[0xCBEF], 5, cpu.registers.getA},
		{cpu.instructions[0xCBF0], 6, cpu.registers.getB},
		{cpu.instructions[0xCBF1], 6, cpu.registers.getC},
		{cpu.instructions[0xCBF2], 6, cpu.registers.getD},
		{cpu.instructions[0xCBF3], 6, cpu.registers.getE},
		{cpu.instructions[0xCBF4], 6, cpu.registers.getH},
		{cpu.instructions[0xCBF5], 6, cpu.registers.getL},
		{cpu.instructions[0xCBF7], 6, cpu.registers.getA},
		{cpu.instructions[0xCBF8], 7, cpu.registers.getB},
		{cpu.instructions[0xCBF9], 7, cpu.registers.getC},
		{cpu.instructions[0xCBFA], 7, cpu.registers.getD},
		{cpu.instructions[0xCBFB], 7, cpu.registers.getE},
		{cpu.instructions[0xCBFC], 7, cpu.registers.getH},
		{cpu.instructions[0xCBFD], 7, cpu.registers.getL},
		{cpu.instructions[0xCBFF], 7, cpu.registers.getA},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			value1 := tc.to() | (1 << tc.bitIndex)
			tc.instruction.exec(cpu)
			value2 := tc.to()
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
		})
	}
}

// SET n,(HL)
func TestSETInstructions2(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		bitIndex    int
	}{
		{cpu.instructions[0xCBC6], 0},
		{cpu.instructions[0xCBCE], 1},
		{cpu.instructions[0xCBD6], 2},
		{cpu.instructions[0xCBDE], 3},
		{cpu.instructions[0xCBE6], 4},
		{cpu.instructions[0xCBEE], 5},
		{cpu.instructions[0xCBF6], 6},
		{cpu.instructions[0xCBFE], 7},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := cpu.registers.getHL()
			value1 := cpu.memory.Read(addr) | (1 << tc.bitIndex)
			tc.instruction.exec(cpu)
			value2 := cpu.memory.Read(addr)
			if value1 != value2 {
				t.Errorf("Expected 0x%02X, got 0x%02X", value1, value2)
			}
		})
	}
}

// SCF
func TestSCFInstructions(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	instructions := []struct {
		instruction instruction
	}{
		{cpu.instructions[0x37]},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			flagN := false
			flagH := false
			flagC := true
			tc.instruction.exec(cpu)
			if flagN != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// CCF
func TestCCFInstructions(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	instructions := []struct {
		instruction instruction
	}{
		{cpu.instructions[0x3F]},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			flagN := false
			flagH := false
			flagC := !cpu.registers.f.getC()
			tc.instruction.exec(cpu)
			if flagN != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// NOP
func TestNOPInstructions(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	instructions := []struct {
		instruction instruction
	}{
		{cpu.instructions[0x00]},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			regA := cpu.registers.getA()
			regF := cpu.registers.getF()
			regB := cpu.registers.getB()
			regC := cpu.registers.getC()
			regD := cpu.registers.getD()
			regE := cpu.registers.getE()
			regH := cpu.registers.getH()
			regL := cpu.registers.getL()
			regSP := cpu.registers.getSP()
			regPC := cpu.registers.getPC()
			flagZ := cpu.registers.f.getZ()
			flagN := cpu.registers.f.getN()
			flagH := cpu.registers.f.getH()
			flagC := cpu.registers.f.getC()
			tc.instruction.exec(cpu)
			if regA != cpu.registers.getA() {
				t.Errorf("Reg A. Expected 0x%02X, got 0x%02X", regA, cpu.registers.getA())
			}
			if regF != cpu.registers.getF() {
				t.Errorf("Reg F. Expected 0x%02X, got 0x%02X", regF, cpu.registers.getF())
			}
			if regB != cpu.registers.getB() {
				t.Errorf("Reg B. Expected 0x%02X, got 0x%02X", regB, cpu.registers.getB())
			}
			if regC != cpu.registers.getC() {
				t.Errorf("Reg C. Expected 0x%02X, got 0x%02X", regC, cpu.registers.getC())
			}
			if regD != cpu.registers.getD() {
				t.Errorf("Reg D. Expected 0x%02X, got 0x%02X", regD, cpu.registers.getD())
			}
			if regE != cpu.registers.getE() {
				t.Errorf("Reg E. Expected 0x%02X, got 0x%02X", regE, cpu.registers.getE())
			}
			if regH != cpu.registers.getH() {
				t.Errorf("Reg H. Expected 0x%02X, got 0x%02X", regH, cpu.registers.getH())
			}
			if regL != cpu.registers.getL() {
				t.Errorf("Reg L. Expected 0x%02X, got 0x%02X", regL, cpu.registers.getL())
			}
			if regSP != cpu.registers.getSP() {
				t.Errorf("Reg SP. Expected 0x%04X, got 0x%04X", regSP, cpu.registers.getSP())
			}
			if regPC != cpu.registers.getPC() {
				t.Errorf("Reg PC. Expected 0x%04X, got 0x%04X", regPC, cpu.registers.getPC())
			}
			if flagZ != cpu.registers.f.getZ() {
				t.Errorf("Flag Z. Expected %t, got %t", flagZ, cpu.registers.f.getZ())
			}
			if flagN != cpu.registers.f.getN() {
				t.Errorf("Flag N. Expected %t, got %t", flagN, cpu.registers.f.getN())
			}
			if flagH != cpu.registers.f.getH() {
				t.Errorf("Flag H. Expected %t, got %t", flagH, cpu.registers.f.getH())
			}
			if flagC != cpu.registers.f.getC() {
				t.Errorf("Flag C. Expected %t, got %t", flagC, cpu.registers.f.getC())
			}
		})
	}
}

// JP nn
func TestJPInstructions1(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
	}{
		{cpu.instructions[0xC3]},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := cpu.registers.getPC()
			msb := uint16(cpu.memory.Read(addr + 1))
			lsb := uint16(cpu.memory.Read(addr))
			value1 := msb<<8 | lsb
			tc.instruction.exec(cpu)
			value2 := cpu.registers.getPC()
			if value1 != value2 {
				t.Errorf("Expected 0x%04X, got 0x%04X", value1, value2)
			}
		})
	}
}

// JP HL
func TestJPInstructions2(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	instructions := []struct {
		instruction instruction
	}{
		{cpu.instructions[0xE9]},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			value1 := cpu.registers.getHL()
			tc.instruction.exec(cpu)
			value2 := cpu.registers.getPC()
			if value1 != value2 {
				t.Errorf("Expected 0x%04X, got 0x%04X", value1, value2)
			}
		})
	}
}

// JP !c,nn
func TestJPInstructions3(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		flag        func() bool
	}{
		{cpu.instructions[0xC2], cpu.registers.f.getZ},
		{cpu.instructions[0xD2], cpu.registers.f.getC},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			var value1 uint16
			addr := cpu.registers.getPC()
			msb := uint16(cpu.memory.Read(addr + 1))
			lsb := uint16(cpu.memory.Read(addr))
			if !tc.flag() {
				value1 = msb<<8 | lsb
			} else {
				value1 = cpu.registers.getPC() + 2
			}
			tc.instruction.exec(cpu)
			value2 := cpu.registers.getPC()
			if value1 != value2 {
				t.Errorf("Expected 0x%04X, got 0x%04X", value1, value2)
			}
		})
	}
}

// JP f,nn
func TestJPInstructions4(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		flag        func() bool
	}{
		{cpu.instructions[0xCA], cpu.registers.f.getZ},
		{cpu.instructions[0xDA], cpu.registers.f.getC},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			var value1 uint16
			addr := cpu.registers.getPC()
			msb := uint16(cpu.memory.Read(addr + 1))
			lsb := uint16(cpu.memory.Read(addr))
			if tc.flag() {
				value1 = msb<<8 | lsb
			} else {
				value1 = cpu.registers.getPC() + 2
			}
			tc.instruction.exec(cpu)
			value2 := cpu.registers.getPC()
			if value1 != value2 {
				t.Errorf("Expected 0x%04X, got 0x%04X", value1, value2)
			}
		})
	}
}

// JR e
func TestJRInstructions1(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
	}{
		{cpu.instructions[0x18]},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			pc := int(cpu.registers.getPC())
			e := int(int8(read8BitOperand(cpu)))
			value1 := uint16(pc + e)
			tc.instruction.exec(cpu)
			value2 := cpu.registers.getPC()
			if value1 != value2 {
				t.Errorf("Expected 0x%04X, got 0x%04X", value1, value2)
			}
		})
	}
}

// JR !c,nn
func TestJRInstructions2(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		flag        func() bool
	}{
		{cpu.instructions[0x20], cpu.registers.f.getZ},
		{cpu.instructions[0x30], cpu.registers.f.getC},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			var value1 uint16
			randomizeRegisters(cpu.registers)
			pc := int(cpu.registers.getPC())
			e := int(int8(read8BitOperand(cpu)))
			if !tc.flag() {
				value1 = uint16(pc + e)
			} else {
				value1 = cpu.registers.getPC() + 1
			}
			tc.instruction.exec(cpu)
			value2 := cpu.registers.getPC()
			if value1 != value2 {
				t.Errorf("Expected 0x%04X, got 0x%04X", value1, value2)
			}
		})
	}
}

// JR f,e
func TestJRInstructions3(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		flag        func() bool
	}{
		{cpu.instructions[0x28], cpu.registers.f.getZ},
		{cpu.instructions[0x38], cpu.registers.f.getC},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			var value1 uint16
			randomizeRegisters(cpu.registers)
			pc := int(cpu.registers.getPC())
			e := int(int8(read8BitOperand(cpu)))
			if tc.flag() {
				value1 = uint16(pc + e)
			} else {
				value1 = cpu.registers.getPC() + 1
			}
			tc.instruction.exec(cpu)
			value2 := cpu.registers.getPC()
			if value1 != value2 {
				t.Errorf("Expected 0x%04X, got 0x%04X", value1, value2)
			}
		})
	}
}

// CALL nn
func TestCALLInstructions1(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
	}{
		{cpu.instructions[0xCD]},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			value1 := read16BitOperand(cpu)
			addr := cpu.registers.getSP()
			if addr == controllers.ADDR_DIV_COUNTER+1 || addr == controllers.ADDR_DIV_COUNTER+2 { // writing to 0xFF04 does nothing; skip
				return
			}
			msb := uint8((cpu.registers.getPC() + 2) >> 8)
			lsb := uint8((cpu.registers.getPC() + 2) & 0xFF)
			tc.instruction.exec(cpu)
			value2 := cpu.registers.getPC()
			if msb != cpu.memory.Read(addr-1) {
				t.Errorf("SP-1. Expected 0x%02X, got 0x%02X", msb, cpu.memory.Read(addr-1))
			}
			if lsb != cpu.memory.Read(addr-2) {
				t.Errorf("SP-2. Expected 0x%02X, got 0x%02X", lsb, cpu.memory.Read(addr-2))
			}
			if addr-2 != cpu.registers.getSP() {
				t.Errorf("SP. Expected 0x%04X, got 0x%04X", addr-2, cpu.registers.getSP())
			}
			if value1 != value2 {
				t.Errorf("PC. Expected 0x%04X, got 0x%04X", value1, value2)
			}
		})
	}
}

// CALL !f, nn
func TestCALLInstructions2(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		flag        func() bool
	}{
		{cpu.instructions[0xC4], cpu.registers.f.getZ},
		{cpu.instructions[0xD4], cpu.registers.f.getC},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			var value1 uint16
			var msb, lsb uint8
			addr := cpu.registers.getSP()
			if addr == controllers.ADDR_DIV_COUNTER+1 || addr == controllers.ADDR_DIV_COUNTER+2 { // writing to 0xFF04 does nothing; skip
				return
			}
			if !tc.flag() {
				value1 = read16BitOperand(cpu)
				msb = uint8((cpu.registers.getPC() + 2) >> 8)
				lsb = uint8((cpu.registers.getPC() + 2) & 0xFF)
			} else {
				value1 = cpu.registers.getPC() + 2
				msb = cpu.memory.Read(addr - 1)
				lsb = cpu.memory.Read(addr - 2)
			}
			tc.instruction.exec(cpu)
			value2 := cpu.registers.getPC()
			if !tc.flag() {
				if addr-2 != cpu.registers.getSP() {
					t.Errorf("SP. Expected 0x%04X, got 0x%04X", addr-2, cpu.registers.getSP())
				}
			} else {
				if addr != cpu.registers.getSP() {
					t.Errorf("SP. Expected 0x%04X, got 0x%04X", addr, cpu.registers.getSP())
				}
			}
			if msb != cpu.memory.Read(addr-1) {
				t.Errorf("SP-1. Expected 0x%02X, got 0x%02X", msb, cpu.memory.Read(addr-1))
			}
			if lsb != cpu.memory.Read(addr-2) {
				t.Errorf("SP-2. Expected 0x%02X, got 0x%02X", lsb, cpu.memory.Read(addr-2))
			}
			if value1 != value2 {
				t.Errorf("PC. Expected 0x%04X, got 0x%04X", value1, value2)
			}
		})
	}
}

// CALL f, nn
func TestCALLInstructions3(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		flag        func() bool
	}{
		{cpu.instructions[0xCC], cpu.registers.f.getZ},
		{cpu.instructions[0xDC], cpu.registers.f.getC},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			var value1 uint16
			var msb, lsb uint8
			addr := cpu.registers.getSP()
			if addr == controllers.ADDR_DIV_COUNTER+1 || addr == controllers.ADDR_DIV_COUNTER+2 { // writing to 0xFF04 does nothing; skip
				return
			}
			if tc.flag() {
				value1 = read16BitOperand(cpu)
				msb = uint8((cpu.registers.getPC() + 2) >> 8)
				lsb = uint8((cpu.registers.getPC() + 2) & 0xFF)
			} else {
				value1 = cpu.registers.getPC() + 2
				msb = cpu.memory.Read(addr - 1)
				lsb = cpu.memory.Read(addr - 2)
			}
			tc.instruction.exec(cpu)
			value2 := cpu.registers.getPC()
			if tc.flag() {
				if addr-2 != cpu.registers.getSP() {
					t.Errorf("SP. Expected 0x%04X, got 0x%04X", addr-2, cpu.registers.getSP())
				}
			} else {
				if addr != cpu.registers.getSP() {
					t.Errorf("SP. Expected 0x%04X, got 0x%04X", addr, cpu.registers.getSP())
				}
			}
			if msb != cpu.memory.Read(addr-1) {
				t.Errorf("SP-1. Expected 0x%02X, got 0x%02X", msb, cpu.memory.Read(addr-1))
			}
			if lsb != cpu.memory.Read(addr-2) {
				t.Errorf("SP-2. Expected 0x%02X, got 0x%02X", lsb, cpu.memory.Read(addr-2))
			}
			if value1 != value2 {
				t.Errorf("PC. Expected 0x%04X, got 0x%04X", value1, value2)
			}
		})
	}
}

// RET
func TestRETInstructions1(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
	}{
		{cpu.instructions[0xC9]},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := cpu.registers.getSP()
			msb := uint16(cpu.memory.Read(addr + 1))
			lsb := uint16(cpu.memory.Read(addr))
			value1 := msb<<8 | lsb
			tc.instruction.exec(cpu)
			value2 := cpu.registers.getPC()
			if uint8(msb) != cpu.memory.Read(addr+1) {
				t.Errorf("SP-1. Expected 0x%02X, got 0x%02X", uint8(msb), cpu.memory.Read(addr+1))
			}
			if uint8(lsb) != cpu.memory.Read(addr) {
				t.Errorf("SP-2. Expected 0x%02X, got 0x%02X", uint8(lsb), cpu.memory.Read(addr))
			}
			if addr+2 != cpu.registers.getSP() {
				t.Errorf("SP. Expected 0x%04X, got 0x%04X", addr+2, cpu.registers.getSP())
			}
			if value1 != value2 {
				t.Errorf("PC. Expected 0x%04X, got 0x%04X", value1, value2)
			}
		})
	}
}

// RET !f
func TestRETInstructions2(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		flag        func() bool
	}{
		{cpu.instructions[0xC0], cpu.registers.f.getZ},
		{cpu.instructions[0xD0], cpu.registers.f.getC},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			var value1 uint16
			addr := cpu.registers.getSP()
			if !tc.flag() {
				msb := uint16(cpu.memory.Read(addr + 1))
				lsb := uint16(cpu.memory.Read(addr))
				value1 = msb<<8 | lsb
			} else {
				value1 = cpu.registers.getPC()
			}
			tc.instruction.exec(cpu)
			value2 := cpu.registers.getPC()
			if !tc.flag() {
				if addr+2 != cpu.registers.getSP() {
					t.Errorf("SP. Expected 0x%04X, got 0x%04X", addr+2, cpu.registers.getSP())
				}
			} else {
				if addr != cpu.registers.getSP() {
					t.Errorf("SP. Expected 0x%04X, got 0x%04X", addr, cpu.registers.getSP())
				}
			}
			if value1 != value2 {
				t.Errorf("PC. Expected 0x%04X, got 0x%04X", value1, value2)
			}
		})
	}
}

// RET f
func TestRETInstructions3(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		flag        func() bool
	}{
		{cpu.instructions[0xC8], cpu.registers.f.getZ},
		{cpu.instructions[0xD8], cpu.registers.f.getC},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			var value1 uint16
			addr := cpu.registers.getSP()
			if tc.flag() {
				msb := uint16(cpu.memory.Read(addr + 1))
				lsb := uint16(cpu.memory.Read(addr))
				value1 = msb<<8 | lsb
			} else {
				value1 = cpu.registers.getPC()
			}
			tc.instruction.exec(cpu)
			value2 := cpu.registers.getPC()
			if tc.flag() {
				if addr+2 != cpu.registers.getSP() {
					t.Errorf("SP. Expected 0x%04X, got 0x%04X", addr+2, cpu.registers.getSP())
				}
			} else {
				if addr != cpu.registers.getSP() {
					t.Errorf("SP. Expected 0x%04X, got 0x%04X", addr, cpu.registers.getSP())
				}
			}
			if value1 != value2 {
				t.Errorf("PC. Expected 0x%04X, got 0x%04X", value1, value2)
			}
		})
	}
}

// RETI
func TestRETInstructions4(t *testing.T) {
	memory := memory.NewDMGMemory()
	interrupts := controllers.NewInterrupts(memory)
	cpu := NewCPU(memory, interrupts, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
	}{
		{cpu.instructions[0xD9]},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			addr := cpu.registers.getSP()
			msb := uint16(cpu.memory.Read(addr + 1))
			lsb := uint16(cpu.memory.Read(addr))
			value1 := msb<<8 | lsb
			tc.instruction.exec(cpu)
			value2 := cpu.registers.getPC()
			if uint8(msb) != cpu.memory.Read(addr+1) {
				t.Errorf("SP-1. Expected 0x%02X, got 0x%02X", uint8(msb), cpu.memory.Read(addr+1))
			}
			if uint8(lsb) != cpu.memory.Read(addr) {
				t.Errorf("SP-2. Expected 0x%02X, got 0x%02X", uint8(lsb), cpu.memory.Read(addr))
			}
			if addr+2 != cpu.registers.getSP() {
				t.Errorf("SP. Expected 0x%04X, got 0x%04X", addr+2, cpu.registers.getSP())
			}
			if value1 != value2 {
				t.Errorf("PC. Expected 0x%04X, got 0x%04X", value1, value2)
			}
			if true != cpu.interrupts.IsMasterEnabled() {
				t.Errorf("Master disabled after RETI. Expected %v, got %v", true, cpu.interrupts.IsMasterEnabled())
			}
		})
	}
}

// RST h
func TestRSTInstructions1(t *testing.T) {
	memory := memory.NewDMGMemory()
	cpu := NewCPU(memory, nil, nil, nil, nil)

	randomizeMemory(memory)

	instructions := []struct {
		instruction instruction
		vector      uint16
	}{
		{cpu.instructions[0xC7], 0x00},
		{cpu.instructions[0xCF], 0x08},
		{cpu.instructions[0xD7], 0x10},
		{cpu.instructions[0xDF], 0x18},
		{cpu.instructions[0xE7], 0x20},
		{cpu.instructions[0xEF], 0x28},
		{cpu.instructions[0xF7], 0x30},
		{cpu.instructions[0xFF], 0x38},
	}
	for _, tc := range instructions {
		testName := fmt.Sprintf("Executes %s", tc.instruction.mnemonic)
		t.Run(testName, func(t *testing.T) {
			randomizeRegisters(cpu.registers)
			value1 := tc.vector
			addr := cpu.registers.getSP()
			if addr == controllers.ADDR_DIV_COUNTER+1 || addr == controllers.ADDR_DIV_COUNTER+2 { // writing to 0xFF04 does nothing; skip
				return
			}
			msb := uint8(cpu.registers.getPC() >> 8)
			lsb := uint8(cpu.registers.getPC() & 0xFF)
			tc.instruction.exec(cpu)
			value2 := cpu.registers.getPC()
			if msb != cpu.memory.Read(addr-1) {
				t.Errorf("SP-1. Expected 0x%02X, got 0x%02X", msb, cpu.memory.Read(addr-1))
			}
			if lsb != cpu.memory.Read(addr-2) {
				t.Errorf("SP-2. Expected 0x%02X, got 0x%02X", lsb, cpu.memory.Read(addr-2))
			}
			if addr-2 != cpu.registers.getSP() {
				t.Errorf("SP. Expected 0x%04X, got 0x%04X", addr-2, cpu.registers.getSP())
			}
			if value1 != value2 {
				t.Errorf("PC. Expected 0x%04X, got 0x%04X", value1, value2)
			}
		})
	}
}

// TODO(somerussianlad): Finish these tests
// DI
func TestDI(t *testing.T) {
	memory := memory.NewDMGMemory()
	interrupts := controllers.NewInterrupts(memory)
	cpu := NewCPU(memory, interrupts, nil, nil, nil)

	interrupts.EnableMaster()
	cpu.instructions[0xF3].exec(cpu)

	expectedMaster := false
	expectedDelay := false
	gotMaster := interrupts.IsMasterEnabled()
	gotDelay := interrupts.IsDelayed()

	if expectedMaster != gotMaster {
		t.Errorf("Wrong master value. Expected %v, got %v", expectedMaster, gotMaster)
	}
	if expectedDelay != gotDelay {
		t.Errorf("Wrong master value. Expected %v, got %v", expectedMaster, gotMaster)
	}
}

// EI
func TestEI(t *testing.T) {
	memory := memory.NewDMGMemory()
	interrupts := controllers.NewInterrupts(memory)
	cpu := NewCPU(memory, interrupts, nil, nil, nil)

	interrupts.DisableMaster()
	cpu.instructions[0xFB].exec(cpu)

	expectedMaster := true
	expectedDelay := true
	gotMaster := interrupts.IsMasterEnabled()
	gotDelay := interrupts.IsDelayed()

	if expectedMaster != gotMaster {
		t.Errorf("Wrong master value. Expected %v, got %v", expectedMaster, gotMaster)
	}
	if expectedDelay != gotDelay {
		t.Errorf("Wrong master value. Expected %v, got %v", expectedMaster, gotMaster)
	}
}
