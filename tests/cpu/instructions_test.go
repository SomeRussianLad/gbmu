package cpu_test

import (
	"fmt"
	"goboy/cpu"
	"goboy/tests"
	"testing"
)

//	Checks for proper increments of PC after execution
// func TestPC(t *testing.T) {
// 	c := tests.InitCPU()
// 	instructions := cpu.NewInstructions()
// 	for _, instruction := range instructions {
// 		tests.RandRegisters(&c)
// 		testName := fmt.Sprintf("Executes %s", instruction.Mnemonic)
// 		t.Run(testName, func(t *testing.T) {
// 			pc0 := c.Registers.GetPC()

// 			//	Fake instruction read which is supposed to auto-increment PC
// 			c.Registers.IncPC()

// 			instruction.Exec(&c)
// 			pc1 := c.Registers.GetPC()
// 			tests.Equals(t, int(pc1-pc0), instruction.Length, "PC incremented incorrectly. Expected %v, got %v")
// 		})
// 	}
// }

//	Checks for conformity between Flags and Registers structs (Flags == Registers.F)
func TestFlagsConformity(t *testing.T) {
	c := tests.InitCPU()
	instructions := cpu.NewInstructions()
	for _, instruction := range instructions {
		tests.RandRegisters(&c)
		testName := fmt.Sprintf("Executes %s", instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			value := c.Registers.GetF()
			tests.Equals(t, c.Flags.Z, (value&128) == 128, "Conformity broken. Flag Z: %t, RegF[7]: %t")
			tests.Equals(t, c.Flags.N, (value&64) == 64, "Conformity broken. Flag N: %t, RegF[6]: %t")
			tests.Equals(t, c.Flags.H, (value&32) == 32, "Conformity broken. Flag H: %t, RegF[5]: %t")
			tests.Equals(t, c.Flags.C, (value&16) == 16, "Conformity broken. Flag C: %t, RegF[4]: %t")
		})
	}
}

//	LD r,r
func TestLDInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To, From    func() uint8
	}{
		{c.Instructions[0x40], c.Registers.GetB, c.Registers.GetB},
		{c.Instructions[0x41], c.Registers.GetB, c.Registers.GetC},
		{c.Instructions[0x42], c.Registers.GetB, c.Registers.GetD},
		{c.Instructions[0x43], c.Registers.GetB, c.Registers.GetE},
		{c.Instructions[0x44], c.Registers.GetB, c.Registers.GetH},
		{c.Instructions[0x45], c.Registers.GetB, c.Registers.GetL},
		{c.Instructions[0x47], c.Registers.GetB, c.Registers.GetA},
		{c.Instructions[0x48], c.Registers.GetC, c.Registers.GetB},
		{c.Instructions[0x49], c.Registers.GetC, c.Registers.GetC},
		{c.Instructions[0x4A], c.Registers.GetC, c.Registers.GetD},
		{c.Instructions[0x4B], c.Registers.GetC, c.Registers.GetE},
		{c.Instructions[0x4C], c.Registers.GetC, c.Registers.GetH},
		{c.Instructions[0x4D], c.Registers.GetC, c.Registers.GetL},
		{c.Instructions[0x4F], c.Registers.GetC, c.Registers.GetA},
		{c.Instructions[0x50], c.Registers.GetD, c.Registers.GetB},
		{c.Instructions[0x51], c.Registers.GetD, c.Registers.GetC},
		{c.Instructions[0x52], c.Registers.GetD, c.Registers.GetD},
		{c.Instructions[0x53], c.Registers.GetD, c.Registers.GetE},
		{c.Instructions[0x54], c.Registers.GetD, c.Registers.GetH},
		{c.Instructions[0x55], c.Registers.GetD, c.Registers.GetL},
		{c.Instructions[0x57], c.Registers.GetD, c.Registers.GetA},
		{c.Instructions[0x58], c.Registers.GetE, c.Registers.GetB},
		{c.Instructions[0x59], c.Registers.GetE, c.Registers.GetC},
		{c.Instructions[0x5A], c.Registers.GetE, c.Registers.GetD},
		{c.Instructions[0x5B], c.Registers.GetE, c.Registers.GetE},
		{c.Instructions[0x5C], c.Registers.GetE, c.Registers.GetH},
		{c.Instructions[0x5D], c.Registers.GetE, c.Registers.GetL},
		{c.Instructions[0x5F], c.Registers.GetE, c.Registers.GetA},
		{c.Instructions[0x60], c.Registers.GetH, c.Registers.GetB},
		{c.Instructions[0x61], c.Registers.GetH, c.Registers.GetC},
		{c.Instructions[0x62], c.Registers.GetH, c.Registers.GetD},
		{c.Instructions[0x63], c.Registers.GetH, c.Registers.GetE},
		{c.Instructions[0x64], c.Registers.GetH, c.Registers.GetH},
		{c.Instructions[0x65], c.Registers.GetH, c.Registers.GetL},
		{c.Instructions[0x67], c.Registers.GetH, c.Registers.GetA},
		{c.Instructions[0x68], c.Registers.GetL, c.Registers.GetB},
		{c.Instructions[0x69], c.Registers.GetL, c.Registers.GetC},
		{c.Instructions[0x6A], c.Registers.GetL, c.Registers.GetD},
		{c.Instructions[0x6B], c.Registers.GetL, c.Registers.GetE},
		{c.Instructions[0x6C], c.Registers.GetL, c.Registers.GetH},
		{c.Instructions[0x6D], c.Registers.GetL, c.Registers.GetL},
		{c.Instructions[0x6F], c.Registers.GetL, c.Registers.GetA},
		{c.Instructions[0x78], c.Registers.GetA, c.Registers.GetB},
		{c.Instructions[0x79], c.Registers.GetA, c.Registers.GetC},
		{c.Instructions[0x7A], c.Registers.GetA, c.Registers.GetD},
		{c.Instructions[0x7B], c.Registers.GetA, c.Registers.GetE},
		{c.Instructions[0x7C], c.Registers.GetA, c.Registers.GetH},
		{c.Instructions[0x7D], c.Registers.GetA, c.Registers.GetL},
		{c.Instructions[0x7F], c.Registers.GetA, c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			testCase.Instruction.Exec(&c)
			value1 := testCase.From()
			value2 := testCase.To()
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
		})
	}
}

//	LD r,n
func TestLDInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint8
	}{
		{c.Instructions[0x06], c.Registers.GetB},
		{c.Instructions[0x0E], c.Registers.GetC},
		{c.Instructions[0x16], c.Registers.GetD},
		{c.Instructions[0x1E], c.Registers.GetE},
		{c.Instructions[0x26], c.Registers.GetH},
		{c.Instructions[0x2E], c.Registers.GetL},
		{c.Instructions[0x3E], c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			value1 := tests.Read8BitOperand(&c)
			testCase.Instruction.Exec(&c)
			value2 := testCase.To()
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
		})
	}
}

//	LD r,(HL)
func TestLDInstructions3(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint8
	}{
		{c.Instructions[0x46], c.Registers.GetB},
		{c.Instructions[0x4E], c.Registers.GetC},
		{c.Instructions[0x56], c.Registers.GetD},
		{c.Instructions[0x5E], c.Registers.GetE},
		{c.Instructions[0x66], c.Registers.GetH},
		{c.Instructions[0x6E], c.Registers.GetL},
		{c.Instructions[0x7E], c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := c.Registers.GetHL()
			value1 := c.Memory.Read(addr)
			testCase.Instruction.Exec(&c)
			value2 := testCase.To()
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
		})
	}
}

//	LD (HL),r
func TestLDInstructions4(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		From        func() uint8
	}{
		{c.Instructions[0x70], c.Registers.GetB},
		{c.Instructions[0x71], c.Registers.GetC},
		{c.Instructions[0x72], c.Registers.GetD},
		{c.Instructions[0x73], c.Registers.GetE},
		{c.Instructions[0x74], c.Registers.GetH},
		{c.Instructions[0x75], c.Registers.GetL},
		{c.Instructions[0x77], c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			testCase.Instruction.Exec(&c)
			addr := c.Registers.GetHL()
			value1 := testCase.From()
			value2 := c.Memory.Read(addr)
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
		})
	}
}

//	LD (HL),n
func TestLDInstructions5(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0x36]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := c.Registers.GetHL()
			value1 := tests.Read8BitOperand(&c)
			testCase.Instruction.Exec(&c)
			value2 := c.Memory.Read(addr)
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
		})
	}
}

//	LD A,(BC)
func TestLDInstructions6(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint8
	}{
		{c.Instructions[0x0A], c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			testCase.Instruction.Exec(&c)
			addr := c.Registers.GetBC()
			value1 := c.Memory.Read(addr)
			value2 := testCase.To()
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
		})
	}
}

//	LD A,(DE)
func TestLDInstructions7(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint8
	}{
		{c.Instructions[0x1A], c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			testCase.Instruction.Exec(&c)
			addr := c.Registers.GetDE()
			value1 := c.Memory.Read(addr)
			value2 := testCase.To()
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
		})
	}
}

//	LD A,(nn)
func TestLDInstructions8(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint8
	}{
		{c.Instructions[0xFA], c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := tests.Read16BitOperand(&c)
			testCase.Instruction.Exec(&c)
			value1 := c.Memory.Read(addr)
			value2 := testCase.To()
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
		})
	}
}

//	LD (BC),A
func TestLDInstructions9(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		From        func() uint8
	}{
		{c.Instructions[0x02], c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			testCase.Instruction.Exec(&c)
			addr := c.Registers.GetBC()
			value1 := testCase.From()
			value2 := c.Memory.Read(addr)
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
		})
	}
}

//	LD (DE),A
func TestLDInstructions10(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		From        func() uint8
	}{
		{c.Instructions[0x12], c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			testCase.Instruction.Exec(&c)
			addr := c.Registers.GetDE()
			value1 := testCase.From()
			value2 := c.Memory.Read(addr)
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
		})
	}
}

//	LD (nn),A
func TestLDInstructions11(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		From        func() uint8
	}{
		{c.Instructions[0xEA], c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := tests.Read16BitOperand(&c)
			testCase.Instruction.Exec(&c)
			value1 := testCase.From()
			value2 := c.Memory.Read(addr)
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
		})
	}
}

//	LDH A,(n)
func TestLDInstructions12(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint8
	}{
		{c.Instructions[0xF0], c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			msb := uint16(0xFF)
			lsb := uint16(tests.Read8BitOperand(&c))
			addr := msb<<8 | lsb
			testCase.Instruction.Exec(&c)
			value1 := c.Memory.Read(addr)
			value2 := testCase.To()
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
		})
	}
}

//	LDH (n),A
func TestLDInstructions13(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		From        func() uint8
	}{
		{c.Instructions[0xE0], c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			msb := uint16(0xFF)
			lsb := uint16(tests.Read8BitOperand(&c))
			addr := msb<<8 | lsb
			testCase.Instruction.Exec(&c)
			value1 := testCase.From()
			value2 := c.Memory.Read(addr)
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
		})
	}
}

//	LD A,(C)
func TestLDInstructions14(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To, From    func() uint8
	}{
		{c.Instructions[0xF2], c.Registers.GetA, c.Registers.GetC},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			testCase.Instruction.Exec(&c)
			msb := uint16(0xFF)
			lsb := uint16(testCase.From())
			addr := msb<<8 | lsb
			value1 := c.Memory.Read(addr)
			value2 := testCase.To()
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
		})
	}
}

//	LD (C),A
func TestLDInstructions15(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To, From    func() uint8
	}{
		{c.Instructions[0xE2], c.Registers.GetC, c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			testCase.Instruction.Exec(&c)
			msb := uint16(0xFF)
			lsb := uint16(testCase.To())
			addr := msb<<8 | lsb
			value1 := testCase.From()
			value2 := c.Memory.Read(addr)
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
		})
	}
}

//	LDI (HL),A
func TestLDInstructions16(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		From        func() uint8
	}{
		{c.Instructions[0x22], c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := c.Registers.GetHL()
			testCase.Instruction.Exec(&c)
			value1 := testCase.From()
			value2 := c.Memory.Read(addr)
			prevAddr := addr
			currAddr := c.Registers.GetHL()
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, prevAddr+1, currAddr, "Expected 0x%04X, got 0x%04X")
		})
	}
}

//	LDI A,(HL)
func TestLDInstructions17(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint8
	}{
		{c.Instructions[0x2A], c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := c.Registers.GetHL()
			testCase.Instruction.Exec(&c)
			value1 := c.Memory.Read(addr)
			value2 := testCase.To()
			prevAddr := addr
			currAddr := c.Registers.GetHL()
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, prevAddr+1, currAddr, "Expected 0x%04X, got 0x%04X")
		})
	}
}

//	LDD (HL),A
func TestLDInstructions18(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		From        func() uint8
	}{
		{c.Instructions[0x32], c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := c.Registers.GetHL()
			testCase.Instruction.Exec(&c)
			value1 := testCase.From()
			value2 := c.Memory.Read(addr)
			prevAddr := addr
			currAddr := c.Registers.GetHL()
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, prevAddr-1, currAddr, "Expected 0x%04X, got 0x%04X")
		})
	}
}

//	LDD A,(HL)
func TestLDInstructions19(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint8
	}{
		{c.Instructions[0x3A], c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := c.Registers.GetHL()
			testCase.Instruction.Exec(&c)
			value1 := c.Memory.Read(addr)
			value2 := testCase.To()
			prevAddr := addr
			currAddr := c.Registers.GetHL()
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, prevAddr-1, currAddr, "Expected 0x%04X, got 0x%04X")
		})
	}
}

//	LD rr,nn
func TestLDInstructions20(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint16
	}{
		{c.Instructions[0x01], c.Registers.GetBC},
		{c.Instructions[0x11], c.Registers.GetDE},
		{c.Instructions[0x21], c.Registers.GetHL},
		{c.Instructions[0x31], c.Registers.GetSP},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			value1 := tests.Read16BitOperand(&c)
			testCase.Instruction.Exec(&c)
			value2 := testCase.To()
			tests.Equals(t, value1, value2, "Expected 0x%04X, got 0x%04X")
		})
	}
}

//	LD (nn),SP
func TestLDInstructions21(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction  cpu.Instruction
		From1, From2 func() uint8
	}{
		{c.Instructions[0x08], c.Registers.GetP, c.Registers.GetS},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := tests.Read16BitOperand(&c)
			testCase.Instruction.Exec(&c)
			value1 := c.Memory.Read(addr)
			value2 := c.Memory.Read(addr + 1)
			tests.Equals(t, value1, testCase.From1(), "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, value2, testCase.From2(), "Expected 0x%02X, got 0x%02X")
		})
	}
}

//	LDHL SP+e
func TestLDInstructions22(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint16
	}{
		{c.Instructions[0xF8], c.Registers.GetHL},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			sp := int(c.Registers.GetSP())
			e := int(int8(tests.Read8BitOperand(&c)))
			testCase.Instruction.Exec(&c)
			value1 := uint16(sp + e)
			value2 := testCase.To()
			flagH := false
			flagC := false
			if e < 0 {
				flagH = (sp & 0xF) < (e & 0xF)
				flagC = (sp & 0xFF) < (e & 0xFF)
			} else {
				flagH = (sp&0xF)+(e&0xF) > 0xF
				flagC = (sp&0xFF)+(e&0xFF) > 0xFF
			}
			tests.Equals(t, value1, value2, "Expected 0x%04X, got 0x%04X")
			tests.Equals(t, false, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	LD SP,HL
func TestLDInstructions23(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To, From    func() uint16
	}{
		{c.Instructions[0xF9], c.Registers.GetSP, c.Registers.GetHL},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			testCase.Instruction.Exec(&c)
			tests.Equals(t, testCase.From(), testCase.To(), "Expected 0x%04X, got 0x%04X")
		})
	}
}

//	PUSH rr
func TestPUSHInstructions(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction  cpu.Instruction
		From1, From2 func() uint8
	}{
		{c.Instructions[0xC5], c.Registers.GetB, c.Registers.GetC},
		{c.Instructions[0xD5], c.Registers.GetD, c.Registers.GetE},
		{c.Instructions[0xE5], c.Registers.GetH, c.Registers.GetL},
		{c.Instructions[0xF5], c.Registers.GetA, c.Registers.GetF},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := c.Registers.GetSP()
			testCase.Instruction.Exec(&c)
			value1 := testCase.From1()
			value2 := testCase.From2()
			value3 := c.Memory.Read(addr - 1)
			value4 := c.Memory.Read(addr - 2)
			prevAddr := addr
			currAddr := c.Registers.GetSP()
			tests.Equals(t, value1, value3, "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, value2, value4, "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, prevAddr-2, currAddr, "Expected 0x%04X, got 0x%04X")
		})
	}
}

//	POP rr
func TestPOPInstructions(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To1, To2    func() uint8
	}{
		{c.Instructions[0xC1], c.Registers.GetC, c.Registers.GetB},
		{c.Instructions[0xD1], c.Registers.GetE, c.Registers.GetD},
		{c.Instructions[0xE1], c.Registers.GetL, c.Registers.GetH},
		{c.Instructions[0xF1], c.Registers.GetF, c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := c.Registers.GetSP()
			testCase.Instruction.Exec(&c)
			value1 := c.Memory.Read(addr)
			value2 := c.Memory.Read(addr + 1)
			value3 := testCase.To1()
			value4 := testCase.To2()
			prevAddr := addr
			currAddr := c.Registers.GetSP()
			tests.Equals(t, value1, value3, "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, value2, value4, "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, prevAddr+2, currAddr, "Expected 0x%04X, got 0x%04X")
		})
	}
}

//	ADD A,r
func TestADDInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To, From    func() uint8
	}{
		{c.Instructions[0x80], c.Registers.GetA, c.Registers.GetB},
		{c.Instructions[0x81], c.Registers.GetA, c.Registers.GetC},
		{c.Instructions[0x82], c.Registers.GetA, c.Registers.GetD},
		{c.Instructions[0x83], c.Registers.GetA, c.Registers.GetE},
		{c.Instructions[0x84], c.Registers.GetA, c.Registers.GetH},
		{c.Instructions[0x85], c.Registers.GetA, c.Registers.GetL},
		{c.Instructions[0x87], c.Registers.GetA, c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			value1 := testCase.To()
			value2 := testCase.From()
			flagZ := value1+value2 == 0
			flagN := false
			flagH := (value1&0xF)+(value2&0xF) > 0xF
			flagC := uint16(value1)+uint16(value2) > 0xFF
			testCase.Instruction.Exec(&c)
			tests.Equals(t, value1+value2, testCase.To(), "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	ADD A,n
func TestADDInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint8
	}{
		{c.Instructions[0xC6], c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			value1 := testCase.To()
			value2 := tests.Read8BitOperand(&c)
			flagZ := value1+value2 == 0
			flagN := false
			flagH := (value1&0xF)+(value2&0xF) > 0xF
			flagC := uint16(value1)+uint16(value2) > 0xFF
			testCase.Instruction.Exec(&c)
			tests.Equals(t, value1+value2, testCase.To(), "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	ADD A,(HL)
func TestADDInstructions3(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint8
	}{
		{c.Instructions[0x86], c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := c.Registers.GetHL()
			value1 := testCase.To()
			value2 := c.Memory.Read(addr)
			flagZ := value1+value2 == 0
			flagN := false
			flagH := (value1&0xF)+(value2&0xF) > 0xF
			flagC := uint16(value1)+uint16(value2) > 0xFF
			testCase.Instruction.Exec(&c)
			tests.Equals(t, value1+value2, testCase.To(), "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	ADD HL,rr
func TestADDInstructions4(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To, From    func() uint16
	}{
		{c.Instructions[0x09], c.Registers.GetHL, c.Registers.GetBC},
		{c.Instructions[0x19], c.Registers.GetHL, c.Registers.GetDE},
		{c.Instructions[0x29], c.Registers.GetHL, c.Registers.GetHL},
		{c.Instructions[0x39], c.Registers.GetHL, c.Registers.GetSP},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			value1 := testCase.To()
			value2 := testCase.From()
			flagZ := c.Flags.Z
			flagN := false
			flagH := (value1&0xFFF)+(value2&0xFFF) > 0xFFF
			flagC := uint32(value1)+uint32(value2) > 0xFFFF
			testCase.Instruction.Exec(&c)
			tests.Equals(t, value1+value2, testCase.To(), "Expected 0x%04X, got 0x%04X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	ADD SP,e
func TestADDInstructions5(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint16
	}{
		{c.Instructions[0xE8], c.Registers.GetSP},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			value1 := int(c.Registers.GetSP())
			value2 := int(int8(tests.Read8BitOperand(&c)))
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
			testCase.Instruction.Exec(&c)
			tests.Equals(t, uint16(value1+value2), testCase.To(), "Expected 0x%04X, got 0x%04X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	ADC A,r
func TestADCInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To, From    func() uint8
	}{
		{c.Instructions[0x88], c.Registers.GetA, c.Registers.GetB},
		{c.Instructions[0x89], c.Registers.GetA, c.Registers.GetC},
		{c.Instructions[0x8A], c.Registers.GetA, c.Registers.GetD},
		{c.Instructions[0x8B], c.Registers.GetA, c.Registers.GetE},
		{c.Instructions[0x8C], c.Registers.GetA, c.Registers.GetH},
		{c.Instructions[0x8D], c.Registers.GetA, c.Registers.GetL},
		{c.Instructions[0x8F], c.Registers.GetA, c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			carry := c.Flags.GetCarryAsValue()
			value1 := testCase.To()
			value2 := testCase.From()
			flagZ := value1+value2+carry == 0
			flagN := false
			flagH := (value1&0xF)+(value2&0xF)+carry > 0xF
			flagC := uint16(value1)+uint16(value2)+uint16(carry) > 0xFF
			testCase.Instruction.Exec(&c)
			tests.Equals(t, value1+value2+carry, testCase.To(), "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	ADC A,n
func TestADCInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint8
	}{
		{c.Instructions[0xCE], c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			carry := c.Flags.GetCarryAsValue()
			value1 := testCase.To()
			value2 := tests.Read8BitOperand(&c)
			flagZ := value1+value2+carry == 0
			flagN := false
			flagH := (value1&0xF)+(value2&0xF)+carry > 0xF
			flagC := uint16(value1)+uint16(value2)+uint16(carry) > 0xFF
			testCase.Instruction.Exec(&c)
			tests.Equals(t, value1+value2+carry, testCase.To(), "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	ADC A,(HL)
func TestADCInstructions3(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint8
	}{
		{c.Instructions[0x8E], c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := c.Registers.GetHL()
			carry := c.Flags.GetCarryAsValue()
			value1 := testCase.To()
			value2 := c.Memory.Read(addr)
			flagZ := value1+value2+carry == 0
			flagN := false
			flagH := (value1&0xF)+(value2&0xF)+carry > 0xF
			flagC := uint16(value1)+uint16(value2)+uint16(carry) > 0xFF
			testCase.Instruction.Exec(&c)
			tests.Equals(t, value1+value2+carry, testCase.To(), "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	SUB A,r
func TestSUBInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To, From    func() uint8
	}{
		{c.Instructions[0x90], c.Registers.GetA, c.Registers.GetB},
		{c.Instructions[0x91], c.Registers.GetA, c.Registers.GetC},
		{c.Instructions[0x92], c.Registers.GetA, c.Registers.GetD},
		{c.Instructions[0x93], c.Registers.GetA, c.Registers.GetE},
		{c.Instructions[0x94], c.Registers.GetA, c.Registers.GetH},
		{c.Instructions[0x95], c.Registers.GetA, c.Registers.GetL},
		{c.Instructions[0x97], c.Registers.GetA, c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			value1 := testCase.To()
			value2 := testCase.From()
			flagZ := value1-value2 == 0
			flagN := true
			flagH := (value1 & 0xF) < (value2 & 0xF)
			flagC := uint16(value1) < uint16(value2)
			testCase.Instruction.Exec(&c)
			tests.Equals(t, value1-value2, testCase.To(), "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	SUB A,n
func TestSUBInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint8
	}{
		{c.Instructions[0xD6], c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			value1 := testCase.To()
			value2 := tests.Read8BitOperand(&c)
			flagZ := value1-value2 == 0
			flagN := true
			flagH := (value1 & 0xF) < (value2 & 0xF)
			flagC := uint16(value1) < uint16(value2)
			testCase.Instruction.Exec(&c)
			tests.Equals(t, value1-value2, testCase.To(), "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	SUB A,(HL)
func TestSUBInstructions3(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint8
	}{
		{c.Instructions[0x96], c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := c.Registers.GetHL()
			value1 := testCase.To()
			value2 := c.Memory.Read(addr)
			flagZ := value1-value2 == 0
			flagN := true
			flagH := (value1 & 0xF) < (value2 & 0xF)
			flagC := uint16(value1) < uint16(value2)
			testCase.Instruction.Exec(&c)
			tests.Equals(t, value1-value2, testCase.To(), "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	SBC A,r
func TestSBCInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To, From    func() uint8
	}{
		{c.Instructions[0x98], c.Registers.GetA, c.Registers.GetB},
		{c.Instructions[0x99], c.Registers.GetA, c.Registers.GetC},
		{c.Instructions[0x9A], c.Registers.GetA, c.Registers.GetD},
		{c.Instructions[0x9B], c.Registers.GetA, c.Registers.GetE},
		{c.Instructions[0x9C], c.Registers.GetA, c.Registers.GetH},
		{c.Instructions[0x9D], c.Registers.GetA, c.Registers.GetL},
		{c.Instructions[0x9F], c.Registers.GetA, c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			carry := c.Flags.GetCarryAsValue()
			value1 := testCase.To()
			value2 := testCase.From()
			flagZ := value1-value2-carry == 0
			flagN := true
			flagH := (value1 & 0xF) < ((value2 & 0xF) + carry)
			flagC := uint16(value1) < (uint16(value2) + uint16(carry))
			testCase.Instruction.Exec(&c)
			tests.Equals(t, value1-value2-carry, testCase.To(), "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	SBC A,n
func TestSBCInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint8
	}{
		{c.Instructions[0xDE], c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			carry := c.Flags.GetCarryAsValue()
			value1 := testCase.To()
			value2 := tests.Read8BitOperand(&c)
			flagZ := value1-value2-carry == 0
			flagN := true
			flagH := (value1 & 0xF) < ((value2 & 0xF) + carry)
			flagC := uint16(value1) < (uint16(value2) + uint16(carry))
			testCase.Instruction.Exec(&c)
			tests.Equals(t, value1-value2-carry, testCase.To(), "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	SBC A,(HL)
func TestSBCInstructions3(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint8
	}{
		{c.Instructions[0x9E], c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := c.Registers.GetHL()
			carry := c.Flags.GetCarryAsValue()
			value1 := testCase.To()
			value2 := c.Memory.Read(addr)
			flagZ := value1-value2-carry == 0
			flagN := true
			flagH := (value1 & 0xF) < ((value2 & 0xF) + carry)
			flagC := uint16(value1) < (uint16(value2) + uint16(carry))
			testCase.Instruction.Exec(&c)
			tests.Equals(t, value1-value2-carry, testCase.To(), "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	AND A,r
func TestANDInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To, From    func() uint8
	}{
		{c.Instructions[0xA0], c.Registers.GetA, c.Registers.GetB},
		{c.Instructions[0xA1], c.Registers.GetA, c.Registers.GetC},
		{c.Instructions[0xA2], c.Registers.GetA, c.Registers.GetD},
		{c.Instructions[0xA3], c.Registers.GetA, c.Registers.GetE},
		{c.Instructions[0xA4], c.Registers.GetA, c.Registers.GetH},
		{c.Instructions[0xA5], c.Registers.GetA, c.Registers.GetL},
		{c.Instructions[0xA7], c.Registers.GetA, c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			value1 := testCase.To()
			value2 := testCase.From()
			flagZ := value1&value2 == 0
			flagN := false
			flagH := true
			flagC := false
			testCase.Instruction.Exec(&c)
			tests.Equals(t, value1&value2, testCase.To(), "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	AND A,n
func TestANDInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint8
	}{
		{c.Instructions[0xE6], c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			value1 := testCase.To()
			value2 := tests.Read8BitOperand(&c)
			flagZ := value1&value2 == 0
			flagN := false
			flagH := true
			flagC := false
			testCase.Instruction.Exec(&c)
			tests.Equals(t, value1&value2, testCase.To(), "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	AND A,(HL)
func TestANDInstructions3(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint8
	}{
		{c.Instructions[0xA6], c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := c.Registers.GetHL()
			value1 := testCase.To()
			value2 := c.Memory.Read(addr)
			flagZ := value1&value2 == 0
			flagN := false
			flagH := true
			flagC := false
			testCase.Instruction.Exec(&c)
			tests.Equals(t, value1&value2, testCase.To(), "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	XOR A,r
func TestXORInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To, From    func() uint8
	}{
		{c.Instructions[0xA8], c.Registers.GetA, c.Registers.GetB},
		{c.Instructions[0xA9], c.Registers.GetA, c.Registers.GetC},
		{c.Instructions[0xAA], c.Registers.GetA, c.Registers.GetD},
		{c.Instructions[0xAB], c.Registers.GetA, c.Registers.GetE},
		{c.Instructions[0xAC], c.Registers.GetA, c.Registers.GetH},
		{c.Instructions[0xAD], c.Registers.GetA, c.Registers.GetL},
		{c.Instructions[0xAF], c.Registers.GetA, c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			value1 := testCase.To()
			value2 := testCase.From()
			flagZ := value1^value2 == 0
			flagN := false
			flagH := false
			flagC := false
			testCase.Instruction.Exec(&c)
			tests.Equals(t, value1^value2, testCase.To(), "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	XOR A,n
func TestXORInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint8
	}{
		{c.Instructions[0xEE], c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			value1 := testCase.To()
			value2 := tests.Read8BitOperand(&c)
			flagZ := value1^value2 == 0
			flagN := false
			flagH := false
			flagC := false
			testCase.Instruction.Exec(&c)
			tests.Equals(t, value1^value2, testCase.To(), "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	XOR A,(HL)
func TestXORInstructions3(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint8
	}{
		{c.Instructions[0xAE], c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := c.Registers.GetHL()
			value1 := c.Registers.GetA()
			value2 := c.Memory.Read(addr)
			flagZ := value1^value2 == 0
			flagN := false
			flagH := false
			flagC := false
			testCase.Instruction.Exec(&c)
			tests.Equals(t, value1^value2, testCase.To(), "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	OR A,r
func TestORInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To, From    func() uint8
	}{
		{c.Instructions[0xB0], c.Registers.GetA, c.Registers.GetB},
		{c.Instructions[0xB1], c.Registers.GetA, c.Registers.GetC},
		{c.Instructions[0xB2], c.Registers.GetA, c.Registers.GetD},
		{c.Instructions[0xB3], c.Registers.GetA, c.Registers.GetE},
		{c.Instructions[0xB4], c.Registers.GetA, c.Registers.GetH},
		{c.Instructions[0xB5], c.Registers.GetA, c.Registers.GetL},
		{c.Instructions[0xB7], c.Registers.GetA, c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			value1 := testCase.To()
			value2 := testCase.From()
			flagZ := value1|value2 == 0
			flagN := false
			flagH := false
			flagC := false
			testCase.Instruction.Exec(&c)
			tests.Equals(t, value1|value2, testCase.To(), "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	OR A,n
func TestORInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint8
	}{
		{c.Instructions[0xF6], c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			value1 := testCase.To()
			value2 := tests.Read8BitOperand(&c)
			flagZ := value1|value2 == 0
			flagN := false
			flagH := false
			flagC := false
			testCase.Instruction.Exec(&c)
			tests.Equals(t, value1|value2, testCase.To(), "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	OR A,(HL)
func TestORInstructions3(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint8
	}{
		{c.Instructions[0xB6], c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := c.Registers.GetHL()
			value1 := testCase.To()
			value2 := c.Memory.Read(addr)
			flagZ := value1|value2 == 0
			flagN := false
			flagH := false
			flagC := false
			testCase.Instruction.Exec(&c)
			tests.Equals(t, value1|value2, testCase.To(), "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	CP A,r
func TestCPInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To, From    func() uint8
	}{
		{c.Instructions[0xB8], c.Registers.GetA, c.Registers.GetB},
		{c.Instructions[0xB9], c.Registers.GetA, c.Registers.GetC},
		{c.Instructions[0xBA], c.Registers.GetA, c.Registers.GetD},
		{c.Instructions[0xBB], c.Registers.GetA, c.Registers.GetE},
		{c.Instructions[0xBC], c.Registers.GetA, c.Registers.GetH},
		{c.Instructions[0xBD], c.Registers.GetA, c.Registers.GetL},
		{c.Instructions[0xBF], c.Registers.GetA, c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			value1 := testCase.To()
			value2 := testCase.From()
			flagZ := value1 == value2
			flagN := true
			flagH := (value1 & 0xF) < (value2 & 0xF)
			flagC := value1 < value2
			testCase.Instruction.Exec(&c)
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	CP A,n
func TestCPInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint8
	}{
		{c.Instructions[0xD6], c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			value1 := testCase.To()
			value2 := tests.Read8BitOperand(&c)
			flagZ := value1 == value2
			flagN := true
			flagH := (value1 & 0xF) < (value2 & 0xF)
			flagC := value1 < value2
			testCase.Instruction.Exec(&c)
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	CP A,(HL)
func TestCPInstructions3(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint8
	}{
		{c.Instructions[0xBE], c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := c.Registers.GetHL()
			value1 := testCase.To()
			value2 := c.Memory.Read(addr)
			flagZ := value1 == value2
			flagN := true
			flagH := (value1 & 0xF) < (value2 & 0xF)
			flagC := value1 < value2
			testCase.Instruction.Exec(&c)
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	INC r
func TestINCInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		From        func() uint8
	}{
		{c.Instructions[0x04], c.Registers.GetB},
		{c.Instructions[0x0C], c.Registers.GetC},
		{c.Instructions[0x14], c.Registers.GetD},
		{c.Instructions[0x1C], c.Registers.GetE},
		{c.Instructions[0x24], c.Registers.GetH},
		{c.Instructions[0x2C], c.Registers.GetL},
		{c.Instructions[0x3C], c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			value1 := testCase.From()
			value2 := value1 + 1
			flagZ := value2 == 0
			flagN := false
			flagH := value2&0xF == 0
			testCase.Instruction.Exec(&c)
			tests.Equals(t, value2, testCase.From(), "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
		})
	}
}

//	INC (HL)
func TestINCInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0x34]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := c.Registers.GetHL()
			value1 := c.Memory.Read(addr)
			value2 := value1 + 1
			flagZ := value2 == 0
			flagN := false
			flagH := value2&0xF == 0
			testCase.Instruction.Exec(&c)
			tests.Equals(t, value2, c.Memory.Read(addr), "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
		})
	}
}

//	INC rr
func TestINCInstructions3(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		From        func() uint16
	}{
		{c.Instructions[0x03], c.Registers.GetBC},
		{c.Instructions[0x13], c.Registers.GetDE},
		{c.Instructions[0x23], c.Registers.GetHL},
		{c.Instructions[0x33], c.Registers.GetSP},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			value1 := testCase.From()
			value2 := value1 + 1
			testCase.Instruction.Exec(&c)
			tests.Equals(t, value2, testCase.From(), "Expected 0x%04X, got 0x%04X")
		})
	}
}

//	DEC r
func TestDECInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		From        func() uint8
	}{
		{c.Instructions[0x05], c.Registers.GetB},
		{c.Instructions[0x0D], c.Registers.GetC},
		{c.Instructions[0x15], c.Registers.GetD},
		{c.Instructions[0x1D], c.Registers.GetE},
		{c.Instructions[0x25], c.Registers.GetH},
		{c.Instructions[0x2D], c.Registers.GetL},
		{c.Instructions[0x3D], c.Registers.GetA},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			value1 := testCase.From()
			value2 := value1 - 1
			flagZ := value2 == 0
			flagN := true
			flagH := value1&0xF == 0
			testCase.Instruction.Exec(&c)
			tests.Equals(t, value2, testCase.From(), "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
		})
	}
}

//	DEC (HL)
func TestDECInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0x35]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := c.Registers.GetHL()
			value1 := c.Memory.Read(addr)
			value2 := value1 - 1
			flagZ := value2 == 0
			flagN := true
			flagH := value1&0xF == 0
			testCase.Instruction.Exec(&c)
			tests.Equals(t, value2, c.Memory.Read(addr), "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
		})
	}
}

//	DEC rr
func TestDECInstructions3(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		From        func() uint16
	}{
		{c.Instructions[0x0B], c.Registers.GetBC},
		{c.Instructions[0x1B], c.Registers.GetDE},
		{c.Instructions[0x2B], c.Registers.GetHL},
		{c.Instructions[0x3B], c.Registers.GetSP},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			value1 := testCase.From()
			value2 := value1 - 1
			testCase.Instruction.Exec(&c)
			tests.Equals(t, value2, testCase.From(), "Expected 0x%04X, got 0x%04X")
		})
	}
}

// func TestRLCAInstruction(t *testing.T) {
// 	c := tests.InitCPU()
// 	instructions := []struct {
// 		Instruction cpu.Instruction
// 	}{
// 		{c.Instructions[0x07]},
// 	}

// 	for _, testCase := range instructions {
// 		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
// 		t.Run(testName, func(t *testing.T) {
// 			tests.RandRegisters(&c)
// 			value := c.Registers.GetA()
// 			bit7 := value & (1 << 7)
// 			value = (value << 1) + bit7
// 			flagZ := false
// 			flagN := false
// 			flagH := false
// 			flagC := bit7 == 1
// 			testCase.Instruction.Exec(&c)
// 			tests.Equals(
// 				t,
// 				value,
// 				c.Registers.GetA(),
// 				"Wrong value in register A. Expected 0x%02X. Got 0x%02X",
// 			)
// 			tests.Equals(
// 				t,
// 				flagZ,
// 				c.Flags.Z,
// 				"Wrong value in Flag Z. Expected %v, got %v",
// 			)
// 			tests.Equals(
// 				t,
// 				flagN,
// 				c.Flags.N,
// 				"Wrong value in Flag N. Expected %v, got %v",
// 			)
// 			tests.Equals(
// 				t,
// 				flagH,
// 				c.Flags.H,
// 				"Wrong value in Flag H. Expected %v, got %v",
// 			)
// 			tests.Equals(
// 				t,
// 				flagC,
// 				c.Flags.C,
// 				"Wrong value in Flag C. Expected %v, got %v",
// 			)
// 			flags := c.Registers.GetF()
// 			tests.Equals(
// 				t,
// 				(flags&(1<<7)) == 1<<7,
// 				c.Flags.Z,
// 				"Wrong value in flag Z. Expected %v, got %v",
// 			)
// 			tests.Equals(
// 				t,
// 				(flags&(1<<6)) == 1<<6,
// 				c.Flags.N,
// 				"Wrong value in flag N. Expected %v, got %v",
// 			)
// 			tests.Equals(
// 				t,
// 				(flags&(1<<5)) == 1<<5,
// 				c.Flags.H,
// 				"Wrong value in flag H. Expected %v, got %v",
// 			)
// 			tests.Equals(
// 				t,
// 				(flags&(1<<4)) == 1<<4,
// 				c.Flags.C,
// 				"Wrong value in flag C. Expected %v, got %v",
// 			)
// 		})
// 	}
// }

// func TestRLAInstruction(t *testing.T) {
// 	c := tests.InitCPU()
// 	instructions := []struct {
// 		Instruction cpu.Instruction
// 	}{
// 		{c.Instructions[0x17]},
// 	}

// 	for _, testCase := range instructions {
// 		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
// 		t.Run(testName, func(t *testing.T) {
// 			tests.RandRegisters(&c)
// 			carry := c.Flags.GetCarryAsValue()
// 			value := c.Registers.GetA()
// 			bit7 := value & (1 << 7)
// 			value = (value << 1) + carry
// 			flagZ := false
// 			flagN := false
// 			flagH := false
// 			flagC := bit7 == 1
// 			testCase.Instruction.Exec(&c)
// 			tests.Equals(
// 				t,
// 				value,
// 				c.Registers.GetA(),
// 				"Wrong value in register A. Expected 0x%02X. Got 0x%02X",
// 			)
// 			tests.Equals(
// 				t,
// 				flagZ,
// 				c.Flags.Z,
// 				"Wrong value in Flag Z. Expected %v, got %v",
// 			)
// 			tests.Equals(
// 				t,
// 				flagN,
// 				c.Flags.N,
// 				"Wrong value in Flag N. Expected %v, got %v",
// 			)
// 			tests.Equals(
// 				t,
// 				flagH,
// 				c.Flags.H,
// 				"Wrong value in Flag H. Expected %v, got %v",
// 			)
// 			tests.Equals(
// 				t,
// 				flagC,
// 				c.Flags.C,
// 				"Wrong value in Flag C. Expected %v, got %v",
// 			)
// 			flags := c.Registers.GetF()
// 			tests.Equals(
// 				t,
// 				(flags&(1<<7)) == 1<<7,
// 				c.Flags.Z,
// 				"Wrong value in flag Z. Expected %v, got %v",
// 			)
// 			tests.Equals(
// 				t,
// 				(flags&(1<<6)) == 1<<6,
// 				c.Flags.N,
// 				"Wrong value in flag N. Expected %v, got %v",
// 			)
// 			tests.Equals(
// 				t,
// 				(flags&(1<<5)) == 1<<5,
// 				c.Flags.H,
// 				"Wrong value in flag H. Expected %v, got %v",
// 			)
// 			tests.Equals(
// 				t,
// 				(flags&(1<<4)) == 1<<4,
// 				c.Flags.C,
// 				"Wrong value in flag C. Expected %v, got %v",
// 			)
// 		})
// 	}
// }
