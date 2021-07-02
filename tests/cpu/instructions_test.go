package cpu_test

import (
	"fmt"
	"goboy/cpu"
	"goboy/tests"
	"testing"
)

//	MAJOR TODO: Найти тесты формата черного ящика. Одного алгоритма и рандомных значений не хватит

//	Checks for proper increments of PC after execution
//	TODO: Добавить IncPC() для всех префиксных инструкций
// func TestPC(t *testing.T) {
// 	c := tests.InitCPU()
// 	instructions := cpu.NewInstructions()
// 	reSkip, _ := regexp.Compile(`JP|JR|CALL|RET`)
// 	// rePrefix, _ := regexp.Compile(`RLC|RRC|RL}`)
// 	for _, instruction := range instructions {
// 		if reSkip.Match([]byte(instruction.Mnemonic)) {
// 			continue
// 		}
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
			instruction.Exec(&c)
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

//	DAA
func TestDAAInstructions(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0x27]},
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
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			for _, operation := range operations {
				c.Registers.SetA(operation.PrevA)
				c.Flags.SetZ(operation.PrevZ)
				c.Flags.SetN(operation.PrevN)
				c.Flags.SetH(operation.PrevH)
				c.Flags.SetC(operation.PrevC)
				testCase.Instruction.Exec(&c)
				tests.Equals(t, operation.CurrA, c.Registers.GetA(), "Expected 0x%02X, got 0x%02X")
				tests.Equals(t, operation.CurrZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
				tests.Equals(t, operation.CurrN, c.Flags.N, "Flag N. Expected %t, got %t")
				tests.Equals(t, operation.CurrH, c.Flags.H, "Flag H. Expected %t, got %t")
				tests.Equals(t, operation.CurrC, c.Flags.C, "Flag C. Expected %t, got %t")
			}
		})
	}
}

//	CPL
func TestCPLInstructions(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0x2F]},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			value1 := c.Registers.GetA() ^ 0xFF
			testCase.Instruction.Exec(&c)
			value2 := c.Registers.GetA()
			flagN := true
			flagH := true
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
		})
	}
}

//	RLCA
func TestRLCA(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0x07]},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			bit7 := c.Registers.GetA() >> 7
			value1 := (c.Registers.GetA() << 1) + bit7
			testCase.Instruction.Exec(&c)
			value2 := c.Registers.GetA()
			flagC := bit7 == 1
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, false, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	RLA
func TestRLA(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0x17]},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			carry := c.Flags.GetCarryAsValue()
			bit7 := c.Registers.GetA() >> 7
			value1 := (c.Registers.GetA() << 1) + carry
			testCase.Instruction.Exec(&c)
			value2 := c.Registers.GetA()
			flagC := bit7 == 1
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, false, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	RRCA
func TestRRCA(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0x0F]},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			bit0 := c.Registers.GetA() << 7
			value1 := (c.Registers.GetA() >> 1) + bit0
			testCase.Instruction.Exec(&c)
			value2 := c.Registers.GetA()
			flagC := bit0 == 128
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, false, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	RRA
func TestRRA(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0x1F]},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			carry := c.Flags.GetCarryAsValue() << 7
			bit0 := c.Registers.GetA() << 7
			value1 := (c.Registers.GetA() >> 1) + carry
			testCase.Instruction.Exec(&c)
			value2 := c.Registers.GetA()
			flagC := bit0 == 128
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, false, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	RLC r
func TestRLCInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint8
	}{
		{c.Instructions[0xCB00], c.Registers.GetB},
		{c.Instructions[0xCB01], c.Registers.GetC},
		{c.Instructions[0xCB02], c.Registers.GetD},
		{c.Instructions[0xCB03], c.Registers.GetE},
		{c.Instructions[0xCB04], c.Registers.GetH},
		{c.Instructions[0xCB05], c.Registers.GetL},
		{c.Instructions[0xCB07], c.Registers.GetA},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			bit7 := testCase.To() >> 7
			value1 := (testCase.To() << 1) + bit7
			testCase.Instruction.Exec(&c)
			value2 := testCase.To()
			flagZ := value2 == 0
			flagC := bit7 == 1
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	RLC (HL)
func TestRLCInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xCB06]},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := c.Registers.GetHL()
			bit7 := c.Memory.Read(addr) >> 7
			value1 := (c.Memory.Read(addr) << 1) + bit7
			testCase.Instruction.Exec(&c)
			value2 := c.Memory.Read(addr)
			flagZ := value2 == 0
			flagC := bit7 == 1
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	RRC r
func TestRRCInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint8
	}{
		{c.Instructions[0xCB08], c.Registers.GetB},
		{c.Instructions[0xCB09], c.Registers.GetC},
		{c.Instructions[0xCB0A], c.Registers.GetD},
		{c.Instructions[0xCB0B], c.Registers.GetE},
		{c.Instructions[0xCB0C], c.Registers.GetH},
		{c.Instructions[0xCB0D], c.Registers.GetL},
		{c.Instructions[0xCB0F], c.Registers.GetA},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			bit0 := testCase.To() << 7
			value1 := (testCase.To() >> 1) + bit0
			testCase.Instruction.Exec(&c)
			value2 := testCase.To()
			flagZ := value2 == 0
			flagC := bit0 == 128
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	RRC (HL)
func TestRRCInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xCB0E]},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := c.Registers.GetHL()
			bit0 := c.Memory.Read(addr) << 7
			value1 := (c.Memory.Read(addr) >> 1) + bit0
			testCase.Instruction.Exec(&c)
			value2 := c.Memory.Read(addr)
			flagZ := value2 == 0
			flagC := bit0 == 128
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	RL r
func TestRLInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint8
	}{
		{c.Instructions[0xCB10], c.Registers.GetB},
		{c.Instructions[0xCB11], c.Registers.GetC},
		{c.Instructions[0xCB12], c.Registers.GetD},
		{c.Instructions[0xCB13], c.Registers.GetE},
		{c.Instructions[0xCB14], c.Registers.GetH},
		{c.Instructions[0xCB15], c.Registers.GetL},
		{c.Instructions[0xCB17], c.Registers.GetA},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			carry := c.Flags.GetCarryAsValue()
			bit7 := testCase.To() >> 7
			value1 := (testCase.To() << 1) + carry
			testCase.Instruction.Exec(&c)
			value2 := testCase.To()
			flagZ := value2 == 0
			flagC := bit7 == 1
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	RL (HL)
func TestRLInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xCB16]},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := c.Registers.GetHL()
			carry := c.Flags.GetCarryAsValue()
			bit7 := c.Memory.Read(addr) >> 7
			value1 := (c.Memory.Read(addr) << 1) + carry
			testCase.Instruction.Exec(&c)
			value2 := c.Memory.Read(addr)
			flagZ := value2 == 0
			flagC := bit7 == 1
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	RR r
func TestRRInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint8
	}{
		{c.Instructions[0xCB18], c.Registers.GetB},
		{c.Instructions[0xCB19], c.Registers.GetC},
		{c.Instructions[0xCB1A], c.Registers.GetD},
		{c.Instructions[0xCB1B], c.Registers.GetE},
		{c.Instructions[0xCB1C], c.Registers.GetH},
		{c.Instructions[0xCB1D], c.Registers.GetL},
		{c.Instructions[0xCB1F], c.Registers.GetA},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			carry := c.Flags.GetCarryAsValue() << 7
			bit0 := testCase.To() << 7
			value1 := (testCase.To() >> 1) + carry
			testCase.Instruction.Exec(&c)
			value2 := testCase.To()
			flagZ := value2 == 0
			flagC := bit0 == 128
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	RR (HL)
func TestRRInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xCB1E]},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := c.Registers.GetHL()
			carry := c.Flags.GetCarryAsValue() << 7
			bit0 := c.Memory.Read(addr) << 7
			value1 := (c.Memory.Read(addr) >> 1) + carry
			testCase.Instruction.Exec(&c)
			value2 := c.Memory.Read(addr)
			flagZ := value2 == 0
			flagC := bit0 == 128
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	SLA r
func TestSLAInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint8
	}{
		{c.Instructions[0xCB20], c.Registers.GetB},
		{c.Instructions[0xCB21], c.Registers.GetC},
		{c.Instructions[0xCB22], c.Registers.GetD},
		{c.Instructions[0xCB23], c.Registers.GetE},
		{c.Instructions[0xCB24], c.Registers.GetH},
		{c.Instructions[0xCB25], c.Registers.GetL},
		{c.Instructions[0xCB27], c.Registers.GetA},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			bit7 := testCase.To() >> 7
			value1 := testCase.To() << 1
			testCase.Instruction.Exec(&c)
			value2 := testCase.To()
			flagZ := value2 == 0
			flagC := bit7 == 1
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	SLA (HL)
func TestSLAInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xCB26]},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := c.Registers.GetHL()
			bit7 := c.Memory.Read(addr) >> 7
			value1 := c.Memory.Read(addr) << 1
			testCase.Instruction.Exec(&c)
			value2 := c.Memory.Read(addr)
			flagZ := value2 == 0
			flagC := bit7 == 1
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	SRA r
func TestSRAInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint8
	}{
		{c.Instructions[0xCB28], c.Registers.GetB},
		{c.Instructions[0xCB29], c.Registers.GetC},
		{c.Instructions[0xCB2A], c.Registers.GetD},
		{c.Instructions[0xCB2B], c.Registers.GetE},
		{c.Instructions[0xCB2C], c.Registers.GetH},
		{c.Instructions[0xCB2D], c.Registers.GetL},
		{c.Instructions[0xCB2F], c.Registers.GetA},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			bit0 := testCase.To() << 7
			bit7 := testCase.To() & 128
			value1 := (testCase.To() >> 1) + bit7
			testCase.Instruction.Exec(&c)
			value2 := testCase.To()
			flagZ := value2 == 0
			flagC := bit0 == 128
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	SRA (HL)
func TestSRAInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xCB2E]},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := c.Registers.GetHL()
			bit0 := c.Memory.Read(addr) << 7
			bit7 := c.Memory.Read(addr) & 128
			value1 := (c.Memory.Read(addr) >> 1) + bit7
			testCase.Instruction.Exec(&c)
			value2 := c.Memory.Read(addr)
			flagZ := value2 == 0
			flagC := bit0 == 128
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	SWAP r
func TestSWAPInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint8
	}{
		{c.Instructions[0xCB30], c.Registers.GetB},
		{c.Instructions[0xCB31], c.Registers.GetC},
		{c.Instructions[0xCB32], c.Registers.GetD},
		{c.Instructions[0xCB33], c.Registers.GetE},
		{c.Instructions[0xCB34], c.Registers.GetH},
		{c.Instructions[0xCB35], c.Registers.GetL},
		{c.Instructions[0xCB37], c.Registers.GetA},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			lo := testCase.To() % 16
			hi := testCase.To() >> 4
			value1 := lo<<4 | hi
			testCase.Instruction.Exec(&c)
			value2 := testCase.To()
			flagZ := value2 == 0
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	SWAP (HL)
func TestSWAPInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xCB36]},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := c.Registers.GetHL()
			lo := c.Memory.Read(addr) % 16
			hi := c.Memory.Read(addr) >> 4
			value1 := lo<<4 | hi
			testCase.Instruction.Exec(&c)
			value2 := c.Memory.Read(addr)
			flagZ := value2 == 0
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	SRL r
func TestSRLInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          func() uint8
	}{
		{c.Instructions[0xCB38], c.Registers.GetB},
		{c.Instructions[0xCB39], c.Registers.GetC},
		{c.Instructions[0xCB3A], c.Registers.GetD},
		{c.Instructions[0xCB3B], c.Registers.GetE},
		{c.Instructions[0xCB3C], c.Registers.GetH},
		{c.Instructions[0xCB3D], c.Registers.GetL},
		{c.Instructions[0xCB3F], c.Registers.GetA},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			bit0 := testCase.To() << 7
			value1 := testCase.To() >> 1
			testCase.Instruction.Exec(&c)
			value2 := testCase.To()
			flagZ := value2 == 0
			flagC := bit0 == 128
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	SRL (HL)
func TestSRLInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xCB3E]},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := c.Registers.GetHL()
			bit0 := c.Memory.Read(addr) << 7
			value1 := c.Memory.Read(addr) >> 1
			testCase.Instruction.Exec(&c)
			value2 := c.Memory.Read(addr)
			flagZ := value2 == 0
			flagC := bit0 == 128
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	BIT n,r
func TestBITInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		BitIndex    int
		To          func() uint8
	}{
		{c.Instructions[0xCB40], 0, c.Registers.GetB},
		{c.Instructions[0xCB41], 0, c.Registers.GetC},
		{c.Instructions[0xCB42], 0, c.Registers.GetD},
		{c.Instructions[0xCB43], 0, c.Registers.GetE},
		{c.Instructions[0xCB44], 0, c.Registers.GetH},
		{c.Instructions[0xCB45], 0, c.Registers.GetL},
		{c.Instructions[0xCB47], 0, c.Registers.GetA},
		{c.Instructions[0xCB48], 1, c.Registers.GetB},
		{c.Instructions[0xCB49], 1, c.Registers.GetC},
		{c.Instructions[0xCB4A], 1, c.Registers.GetD},
		{c.Instructions[0xCB4B], 1, c.Registers.GetE},
		{c.Instructions[0xCB4C], 1, c.Registers.GetH},
		{c.Instructions[0xCB4D], 1, c.Registers.GetL},
		{c.Instructions[0xCB4F], 1, c.Registers.GetA},
		{c.Instructions[0xCB50], 2, c.Registers.GetB},
		{c.Instructions[0xCB51], 2, c.Registers.GetC},
		{c.Instructions[0xCB52], 2, c.Registers.GetD},
		{c.Instructions[0xCB53], 2, c.Registers.GetE},
		{c.Instructions[0xCB54], 2, c.Registers.GetH},
		{c.Instructions[0xCB55], 2, c.Registers.GetL},
		{c.Instructions[0xCB57], 2, c.Registers.GetA},
		{c.Instructions[0xCB58], 3, c.Registers.GetB},
		{c.Instructions[0xCB59], 3, c.Registers.GetC},
		{c.Instructions[0xCB5A], 3, c.Registers.GetD},
		{c.Instructions[0xCB5B], 3, c.Registers.GetE},
		{c.Instructions[0xCB5C], 3, c.Registers.GetH},
		{c.Instructions[0xCB5D], 3, c.Registers.GetL},
		{c.Instructions[0xCB5F], 3, c.Registers.GetA},
		{c.Instructions[0xCB60], 4, c.Registers.GetB},
		{c.Instructions[0xCB61], 4, c.Registers.GetC},
		{c.Instructions[0xCB62], 4, c.Registers.GetD},
		{c.Instructions[0xCB63], 4, c.Registers.GetE},
		{c.Instructions[0xCB64], 4, c.Registers.GetH},
		{c.Instructions[0xCB65], 4, c.Registers.GetL},
		{c.Instructions[0xCB67], 4, c.Registers.GetA},
		{c.Instructions[0xCB68], 5, c.Registers.GetB},
		{c.Instructions[0xCB69], 5, c.Registers.GetC},
		{c.Instructions[0xCB6A], 5, c.Registers.GetD},
		{c.Instructions[0xCB6B], 5, c.Registers.GetE},
		{c.Instructions[0xCB6C], 5, c.Registers.GetH},
		{c.Instructions[0xCB6D], 5, c.Registers.GetL},
		{c.Instructions[0xCB6F], 5, c.Registers.GetA},
		{c.Instructions[0xCB70], 6, c.Registers.GetB},
		{c.Instructions[0xCB71], 6, c.Registers.GetC},
		{c.Instructions[0xCB72], 6, c.Registers.GetD},
		{c.Instructions[0xCB73], 6, c.Registers.GetE},
		{c.Instructions[0xCB74], 6, c.Registers.GetH},
		{c.Instructions[0xCB75], 6, c.Registers.GetL},
		{c.Instructions[0xCB77], 6, c.Registers.GetA},
		{c.Instructions[0xCB78], 7, c.Registers.GetB},
		{c.Instructions[0xCB79], 7, c.Registers.GetC},
		{c.Instructions[0xCB7A], 7, c.Registers.GetD},
		{c.Instructions[0xCB7B], 7, c.Registers.GetE},
		{c.Instructions[0xCB7C], 7, c.Registers.GetH},
		{c.Instructions[0xCB7D], 7, c.Registers.GetL},
		{c.Instructions[0xCB7F], 7, c.Registers.GetA},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			value1 := (testCase.To() & (1 << testCase.BitIndex)) == (1 << testCase.BitIndex)
			testCase.Instruction.Exec(&c)
			value2 := !c.Flags.Z
			flagZ := !value1
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, true, c.Flags.H, "Flag H. Expected %t, got %t")
		})
	}
}

//	BIT n,(HL)
func TestBITInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		BitIndex    int
	}{
		{c.Instructions[0xCB46], 0},
		{c.Instructions[0xCB4E], 1},
		{c.Instructions[0xCB56], 2},
		{c.Instructions[0xCB5E], 3},
		{c.Instructions[0xCB66], 4},
		{c.Instructions[0xCB6E], 5},
		{c.Instructions[0xCB76], 6},
		{c.Instructions[0xCB7E], 7},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := c.Registers.GetHL()
			value1 := (c.Memory.Read(addr) & (1 << testCase.BitIndex)) == (1 << testCase.BitIndex)
			testCase.Instruction.Exec(&c)
			value2 := !c.Flags.Z
			flagZ := !value1
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, false, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, true, c.Flags.H, "Flag H. Expected %t, got %t")
		})
	}
}

//	RES n,r
func TestRESInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		BitIndex    int
		To          func() uint8
	}{
		{c.Instructions[0xCB80], 0, c.Registers.GetB},
		{c.Instructions[0xCB81], 0, c.Registers.GetC},
		{c.Instructions[0xCB82], 0, c.Registers.GetD},
		{c.Instructions[0xCB83], 0, c.Registers.GetE},
		{c.Instructions[0xCB84], 0, c.Registers.GetH},
		{c.Instructions[0xCB85], 0, c.Registers.GetL},
		{c.Instructions[0xCB87], 0, c.Registers.GetA},
		{c.Instructions[0xCB88], 1, c.Registers.GetB},
		{c.Instructions[0xCB89], 1, c.Registers.GetC},
		{c.Instructions[0xCB8A], 1, c.Registers.GetD},
		{c.Instructions[0xCB8B], 1, c.Registers.GetE},
		{c.Instructions[0xCB8C], 1, c.Registers.GetH},
		{c.Instructions[0xCB8D], 1, c.Registers.GetL},
		{c.Instructions[0xCB8F], 1, c.Registers.GetA},
		{c.Instructions[0xCB90], 2, c.Registers.GetB},
		{c.Instructions[0xCB91], 2, c.Registers.GetC},
		{c.Instructions[0xCB92], 2, c.Registers.GetD},
		{c.Instructions[0xCB93], 2, c.Registers.GetE},
		{c.Instructions[0xCB94], 2, c.Registers.GetH},
		{c.Instructions[0xCB95], 2, c.Registers.GetL},
		{c.Instructions[0xCB97], 2, c.Registers.GetA},
		{c.Instructions[0xCB98], 3, c.Registers.GetB},
		{c.Instructions[0xCB99], 3, c.Registers.GetC},
		{c.Instructions[0xCB9A], 3, c.Registers.GetD},
		{c.Instructions[0xCB9B], 3, c.Registers.GetE},
		{c.Instructions[0xCB9C], 3, c.Registers.GetH},
		{c.Instructions[0xCB9D], 3, c.Registers.GetL},
		{c.Instructions[0xCB9F], 3, c.Registers.GetA},
		{c.Instructions[0xCBA0], 4, c.Registers.GetB},
		{c.Instructions[0xCBA1], 4, c.Registers.GetC},
		{c.Instructions[0xCBA2], 4, c.Registers.GetD},
		{c.Instructions[0xCBA3], 4, c.Registers.GetE},
		{c.Instructions[0xCBA4], 4, c.Registers.GetH},
		{c.Instructions[0xCBA5], 4, c.Registers.GetL},
		{c.Instructions[0xCBA7], 4, c.Registers.GetA},
		{c.Instructions[0xCBA8], 5, c.Registers.GetB},
		{c.Instructions[0xCBA9], 5, c.Registers.GetC},
		{c.Instructions[0xCBAA], 5, c.Registers.GetD},
		{c.Instructions[0xCBAB], 5, c.Registers.GetE},
		{c.Instructions[0xCBAC], 5, c.Registers.GetH},
		{c.Instructions[0xCBAD], 5, c.Registers.GetL},
		{c.Instructions[0xCBAF], 5, c.Registers.GetA},
		{c.Instructions[0xCBB0], 6, c.Registers.GetB},
		{c.Instructions[0xCBB1], 6, c.Registers.GetC},
		{c.Instructions[0xCBB2], 6, c.Registers.GetD},
		{c.Instructions[0xCBB3], 6, c.Registers.GetE},
		{c.Instructions[0xCBB4], 6, c.Registers.GetH},
		{c.Instructions[0xCBB5], 6, c.Registers.GetL},
		{c.Instructions[0xCBB7], 6, c.Registers.GetA},
		{c.Instructions[0xCBB8], 7, c.Registers.GetB},
		{c.Instructions[0xCBB9], 7, c.Registers.GetC},
		{c.Instructions[0xCBBA], 7, c.Registers.GetD},
		{c.Instructions[0xCBBB], 7, c.Registers.GetE},
		{c.Instructions[0xCBBC], 7, c.Registers.GetH},
		{c.Instructions[0xCBBD], 7, c.Registers.GetL},
		{c.Instructions[0xCBBF], 7, c.Registers.GetA},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			value1 := testCase.To() & (0xFF - (1 << testCase.BitIndex))
			testCase.Instruction.Exec(&c)
			value2 := testCase.To()
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
		})
	}
}

//	RES n,(HL)
func TestRESInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		BitIndex    int
	}{
		{c.Instructions[0xCB86], 0},
		{c.Instructions[0xCB8E], 1},
		{c.Instructions[0xCB96], 2},
		{c.Instructions[0xCB9E], 3},
		{c.Instructions[0xCBA6], 4},
		{c.Instructions[0xCBAE], 5},
		{c.Instructions[0xCBB6], 6},
		{c.Instructions[0xCBBE], 7},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := c.Registers.GetHL()
			value1 := c.Memory.Read(addr) & (0xFF - (1 << testCase.BitIndex))
			testCase.Instruction.Exec(&c)
			value2 := c.Memory.Read(addr)
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
		})
	}
}

//	SET n,r
func TestSETInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		BitIndex    int
		To          func() uint8
	}{
		{c.Instructions[0xCBC0], 0, c.Registers.GetB},
		{c.Instructions[0xCBC1], 0, c.Registers.GetC},
		{c.Instructions[0xCBC2], 0, c.Registers.GetD},
		{c.Instructions[0xCBC3], 0, c.Registers.GetE},
		{c.Instructions[0xCBC4], 0, c.Registers.GetH},
		{c.Instructions[0xCBC5], 0, c.Registers.GetL},
		{c.Instructions[0xCBC7], 0, c.Registers.GetA},
		{c.Instructions[0xCBC8], 1, c.Registers.GetB},
		{c.Instructions[0xCBC9], 1, c.Registers.GetC},
		{c.Instructions[0xCBCA], 1, c.Registers.GetD},
		{c.Instructions[0xCBCB], 1, c.Registers.GetE},
		{c.Instructions[0xCBCC], 1, c.Registers.GetH},
		{c.Instructions[0xCBCD], 1, c.Registers.GetL},
		{c.Instructions[0xCBCF], 1, c.Registers.GetA},
		{c.Instructions[0xCBD0], 2, c.Registers.GetB},
		{c.Instructions[0xCBD1], 2, c.Registers.GetC},
		{c.Instructions[0xCBD2], 2, c.Registers.GetD},
		{c.Instructions[0xCBD3], 2, c.Registers.GetE},
		{c.Instructions[0xCBD4], 2, c.Registers.GetH},
		{c.Instructions[0xCBD5], 2, c.Registers.GetL},
		{c.Instructions[0xCBD7], 2, c.Registers.GetA},
		{c.Instructions[0xCBD8], 3, c.Registers.GetB},
		{c.Instructions[0xCBD9], 3, c.Registers.GetC},
		{c.Instructions[0xCBDA], 3, c.Registers.GetD},
		{c.Instructions[0xCBDB], 3, c.Registers.GetE},
		{c.Instructions[0xCBDC], 3, c.Registers.GetH},
		{c.Instructions[0xCBDD], 3, c.Registers.GetL},
		{c.Instructions[0xCBDF], 3, c.Registers.GetA},
		{c.Instructions[0xCBE0], 4, c.Registers.GetB},
		{c.Instructions[0xCBE1], 4, c.Registers.GetC},
		{c.Instructions[0xCBE2], 4, c.Registers.GetD},
		{c.Instructions[0xCBE3], 4, c.Registers.GetE},
		{c.Instructions[0xCBE4], 4, c.Registers.GetH},
		{c.Instructions[0xCBE5], 4, c.Registers.GetL},
		{c.Instructions[0xCBE7], 4, c.Registers.GetA},
		{c.Instructions[0xCBE8], 5, c.Registers.GetB},
		{c.Instructions[0xCBE9], 5, c.Registers.GetC},
		{c.Instructions[0xCBEA], 5, c.Registers.GetD},
		{c.Instructions[0xCBEB], 5, c.Registers.GetE},
		{c.Instructions[0xCBEC], 5, c.Registers.GetH},
		{c.Instructions[0xCBED], 5, c.Registers.GetL},
		{c.Instructions[0xCBEF], 5, c.Registers.GetA},
		{c.Instructions[0xCBF0], 6, c.Registers.GetB},
		{c.Instructions[0xCBF1], 6, c.Registers.GetC},
		{c.Instructions[0xCBF2], 6, c.Registers.GetD},
		{c.Instructions[0xCBF3], 6, c.Registers.GetE},
		{c.Instructions[0xCBF4], 6, c.Registers.GetH},
		{c.Instructions[0xCBF5], 6, c.Registers.GetL},
		{c.Instructions[0xCBF7], 6, c.Registers.GetA},
		{c.Instructions[0xCBF8], 7, c.Registers.GetB},
		{c.Instructions[0xCBF9], 7, c.Registers.GetC},
		{c.Instructions[0xCBFA], 7, c.Registers.GetD},
		{c.Instructions[0xCBFB], 7, c.Registers.GetE},
		{c.Instructions[0xCBFC], 7, c.Registers.GetH},
		{c.Instructions[0xCBFD], 7, c.Registers.GetL},
		{c.Instructions[0xCBFF], 7, c.Registers.GetA},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			value1 := testCase.To() | (1 << testCase.BitIndex)
			testCase.Instruction.Exec(&c)
			value2 := testCase.To()
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
		})
	}
}

//	SET n,(HL)
func TestSETInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		BitIndex    int
	}{
		{c.Instructions[0xCBC6], 0},
		{c.Instructions[0xCBCE], 1},
		{c.Instructions[0xCBD6], 2},
		{c.Instructions[0xCBDE], 3},
		{c.Instructions[0xCBE6], 4},
		{c.Instructions[0xCBEE], 5},
		{c.Instructions[0xCBF6], 6},
		{c.Instructions[0xCBFE], 7},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := c.Registers.GetHL()
			value1 := c.Memory.Read(addr) | (1 << testCase.BitIndex)
			testCase.Instruction.Exec(&c)
			value2 := c.Memory.Read(addr)
			tests.Equals(t, value1, value2, "Expected 0x%02X, got 0x%02X")
		})
	}
}

//	SCF
func TestSCFInstructions(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0x37]},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			flagN := false
			flagH := false
			flagC := true
			testCase.Instruction.Exec(&c)
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	CCF
func TestCCFInstructions(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0x3F]},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			flagN := false
			flagH := false
			flagC := !c.Flags.C
			testCase.Instruction.Exec(&c)
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	NOP
func TestNOPInstructions(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0x00]},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			regA := c.Registers.GetA()
			regF := c.Registers.GetF()
			regB := c.Registers.GetB()
			regC := c.Registers.GetC()
			regD := c.Registers.GetD()
			regE := c.Registers.GetE()
			regH := c.Registers.GetH()
			regL := c.Registers.GetL()
			regSP := c.Registers.GetSP()
			regPC := c.Registers.GetPC()
			flagZ := c.Flags.Z
			flagN := c.Flags.N
			flagH := c.Flags.H
			flagC := c.Flags.C
			testCase.Instruction.Exec(&c)
			tests.Equals(t, regA, c.Registers.GetA(), "Reg A. Expected 0x%02X, got 0x%02X")
			tests.Equals(t, regF, c.Registers.GetF(), "Reg F. Expected 0x%02X, got 0x%02X")
			tests.Equals(t, regB, c.Registers.GetB(), "Reg B. Expected 0x%02X, got 0x%02X")
			tests.Equals(t, regC, c.Registers.GetC(), "Reg C. Expected 0x%02X, got 0x%02X")
			tests.Equals(t, regD, c.Registers.GetD(), "Reg D. Expected 0x%02X, got 0x%02X")
			tests.Equals(t, regE, c.Registers.GetE(), "Reg E. Expected 0x%02X, got 0x%02X")
			tests.Equals(t, regH, c.Registers.GetH(), "Reg H. Expected 0x%02X, got 0x%02X")
			tests.Equals(t, regL, c.Registers.GetL(), "Reg L. Expected 0x%02X, got 0x%02X")
			tests.Equals(t, regSP, c.Registers.GetSP(), "Reg SP. Expected 0x%04X, got 0x%04X")
			tests.Equals(t, regPC, c.Registers.GetPC(), "Reg PC. Expected 0x%04X, got 0x%04X")
			tests.Equals(t, flagZ, c.Flags.Z, "Flag Z. Expected %t, got %t")
			tests.Equals(t, flagN, c.Flags.N, "Flag N. Expected %t, got %t")
			tests.Equals(t, flagH, c.Flags.H, "Flag H. Expected %t, got %t")
			tests.Equals(t, flagC, c.Flags.C, "Flag C. Expected %t, got %t")
		})
	}
}

//	JP nn
func TestJPInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xC3]},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := c.Registers.GetPC()
			msb := uint16(c.Memory.Read(addr + 1))
			lsb := uint16(c.Memory.Read(addr))
			value1 := msb<<8 | lsb
			testCase.Instruction.Exec(&c)
			value2 := c.Registers.GetPC()
			tests.Equals(t, value1, value2, "Expected 0x%04X, got 0x%04X")
		})
	}
}

//	JP HL
func TestJPInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xE9]},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			value1 := c.Registers.GetHL()
			testCase.Instruction.Exec(&c)
			value2 := c.Registers.GetPC()
			tests.Equals(t, value1, value2, "Expected 0x%04X, got 0x%04X")
		})
	}
}

//	JP !c,nn
func TestJPInstructions3(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		Flag        func() bool
	}{
		{c.Instructions[0xC2], c.Flags.GetZ},
		{c.Instructions[0xD2], c.Flags.GetC},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			var value1 uint16
			addr := c.Registers.GetPC()
			msb := uint16(c.Memory.Read(addr + 1))
			lsb := uint16(c.Memory.Read(addr))
			if !testCase.Flag() {
				value1 = msb<<8 | lsb
			} else {
				value1 = c.Registers.GetPC() + 2
			}
			testCase.Instruction.Exec(&c)
			value2 := c.Registers.GetPC()
			tests.Equals(t, value1, value2, "Expected 0x%04X, got 0x%04X")
		})
	}
}

//	JP f,nn
func TestJPInstructions4(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		Flag        func() bool
	}{
		{c.Instructions[0xCA], c.Flags.GetZ},
		{c.Instructions[0xDA], c.Flags.GetC},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			var value1 uint16
			addr := c.Registers.GetPC()
			msb := uint16(c.Memory.Read(addr + 1))
			lsb := uint16(c.Memory.Read(addr))
			if testCase.Flag() {
				value1 = msb<<8 | lsb
			} else {
				value1 = c.Registers.GetPC() + 2
			}
			testCase.Instruction.Exec(&c)
			value2 := c.Registers.GetPC()
			tests.Equals(t, value1, value2, "Expected 0x%04X, got 0x%04X")
		})
	}
}

//	JR e
func TestJRInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0x18]},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pc := int(c.Registers.GetPC())
			e := int(int8(tests.Read8BitOperand(&c)))
			value1 := uint16(pc + e)
			testCase.Instruction.Exec(&c)
			value2 := c.Registers.GetPC()
			tests.Equals(t, value1, value2, "Expected 0x%04X, got 0x%04X")
		})
	}
}

//	JR !c,nn
func TestJRInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		Flag        func() bool
	}{
		{c.Instructions[0x20], c.Flags.GetZ},
		{c.Instructions[0x30], c.Flags.GetC},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			var value1 uint16
			tests.RandRegisters(&c)
			pc := int(c.Registers.GetPC())
			e := int(int8(tests.Read8BitOperand(&c)))
			if !testCase.Flag() {
				value1 = uint16(pc + e)
			} else {
				value1 = c.Registers.GetPC() + 1
			}
			testCase.Instruction.Exec(&c)
			value2 := c.Registers.GetPC()
			tests.Equals(t, value1, value2, "Expected 0x%04X, got 0x%04X")
		})
	}
}

//	JR f,e
func TestJRInstructions3(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		Flag        func() bool
	}{
		{c.Instructions[0x28], c.Flags.GetZ},
		{c.Instructions[0x38], c.Flags.GetC},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			var value1 uint16
			tests.RandRegisters(&c)
			pc := int(c.Registers.GetPC())
			e := int(int8(tests.Read8BitOperand(&c)))
			if testCase.Flag() {
				value1 = uint16(pc + e)
			} else {
				value1 = c.Registers.GetPC() + 1
			}
			testCase.Instruction.Exec(&c)
			value2 := c.Registers.GetPC()
			tests.Equals(t, value1, value2, "Expected 0x%04X, got 0x%04X")
		})
	}
}

//	CALL nn
func TestCALLInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xCD]},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			value1 := tests.Read16BitOperand(&c)
			addr := c.Registers.GetSP()
			msb := uint8((c.Registers.GetPC() + 2) >> 8)
			lsb := uint8((c.Registers.GetPC() + 2) & 0xFF)
			testCase.Instruction.Exec(&c)
			value2 := c.Registers.GetPC()
			tests.Equals(t, msb, c.Memory.Read(addr-1), "SP-1. Expected 0x%02X, got 0x%02X")
			tests.Equals(t, lsb, c.Memory.Read(addr-2), "SP-2. Expected 0x%02X, got 0x%02X")
			tests.Equals(t, addr-2, c.Registers.GetSP(), "SP. Expected 0x%04X, got 0x%04X")
			tests.Equals(t, value1, value2, "PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

//	CALL !f, nn
func TestCALLInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		Flag        func() bool
	}{
		{c.Instructions[0xC4], c.Flags.GetZ},
		{c.Instructions[0xD4], c.Flags.GetC},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			var value1 uint16
			var msb, lsb uint8
			addr := c.Registers.GetSP()
			if !testCase.Flag() {
				value1 = tests.Read16BitOperand(&c)
				msb = uint8((c.Registers.GetPC() + 2) >> 8)
				lsb = uint8((c.Registers.GetPC() + 2) & 0xFF)
			} else {
				value1 = c.Registers.GetPC() + 2
				msb = c.Memory.Read(addr - 1)
				lsb = c.Memory.Read(addr - 2)
			}
			testCase.Instruction.Exec(&c)
			value2 := c.Registers.GetPC()
			if !testCase.Flag() {
				tests.Equals(t, addr-2, c.Registers.GetSP(), "SP. Expected 0x%04X, got 0x%04X")
			} else {
				tests.Equals(t, addr, c.Registers.GetSP(), "SP. Expected 0x%04X, got 0x%04X")
			}
			tests.Equals(t, msb, c.Memory.Read(addr-1), "SP-1. Expected 0x%02X, got 0x%02X")
			tests.Equals(t, lsb, c.Memory.Read(addr-2), "SP-2. Expected 0x%02X, got 0x%02X")
			tests.Equals(t, value1, value2, "PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

//	CALL f, nn
func TestCALLInstructions3(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		Flag        func() bool
	}{
		{c.Instructions[0xCC], c.Flags.GetZ},
		{c.Instructions[0xDC], c.Flags.GetC},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			var value1 uint16
			var msb, lsb uint8
			addr := c.Registers.GetSP()
			if testCase.Flag() {
				value1 = tests.Read16BitOperand(&c)
				msb = uint8((c.Registers.GetPC() + 2) >> 8)
				lsb = uint8((c.Registers.GetPC() + 2) & 0xFF)
			} else {
				value1 = c.Registers.GetPC() + 2
				msb = c.Memory.Read(addr - 1)
				lsb = c.Memory.Read(addr - 2)
			}
			testCase.Instruction.Exec(&c)
			value2 := c.Registers.GetPC()
			if testCase.Flag() {
				tests.Equals(t, addr-2, c.Registers.GetSP(), "SP. Expected 0x%04X, got 0x%04X")
			} else {
				tests.Equals(t, addr, c.Registers.GetSP(), "SP. Expected 0x%04X, got 0x%04X")
			}
			tests.Equals(t, msb, c.Memory.Read(addr-1), "SP-1. Expected 0x%02X, got 0x%02X")
			tests.Equals(t, lsb, c.Memory.Read(addr-2), "SP-2. Expected 0x%02X, got 0x%02X")
			tests.Equals(t, value1, value2, "PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

//	RET
func TestRETInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xC9]},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := c.Registers.GetSP()
			msb := uint16(c.Memory.Read(addr + 1))
			lsb := uint16(c.Memory.Read(addr))
			value1 := msb<<8 | lsb
			testCase.Instruction.Exec(&c)
			value2 := c.Registers.GetPC()
			tests.Equals(t, uint8(msb), c.Memory.Read(addr+1), "SP-1. Expected 0x%02X, got 0x%02X")
			tests.Equals(t, uint8(lsb), c.Memory.Read(addr), "SP-2. Expected 0x%02X, got 0x%02X")
			tests.Equals(t, addr+2, c.Registers.GetSP(), "SP. Expected 0x%04X, got 0x%04X")
			tests.Equals(t, value1, value2, "PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

//	RET !f
func TestRETInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		Flag        func() bool
	}{
		{c.Instructions[0xC0], c.Flags.GetZ},
		{c.Instructions[0xD0], c.Flags.GetC},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			var value1 uint16
			addr := c.Registers.GetSP()
			if !testCase.Flag() {
				msb := uint16(c.Memory.Read(addr + 1))
				lsb := uint16(c.Memory.Read(addr))
				value1 = msb<<8 | lsb
			} else {
				value1 = c.Registers.GetPC()
			}
			testCase.Instruction.Exec(&c)
			value2 := c.Registers.GetPC()
			if !testCase.Flag() {
				tests.Equals(t, addr+2, c.Registers.GetSP(), "SP. Expected 0x%04X, got 0x%04X")
			} else {
				tests.Equals(t, addr, c.Registers.GetSP(), "SP. Expected 0x%04X, got 0x%04X")
			}
			tests.Equals(t, value1, value2, "PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

//	RET f
func TestRETInstructions3(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		Flag        func() bool
	}{
		{c.Instructions[0xC8], c.Flags.GetZ},
		{c.Instructions[0xD8], c.Flags.GetC},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			var value1 uint16
			addr := c.Registers.GetSP()
			if testCase.Flag() {
				msb := uint16(c.Memory.Read(addr + 1))
				lsb := uint16(c.Memory.Read(addr))
				value1 = msb<<8 | lsb
			} else {
				value1 = c.Registers.GetPC()
			}
			testCase.Instruction.Exec(&c)
			value2 := c.Registers.GetPC()
			if testCase.Flag() {
				tests.Equals(t, addr+2, c.Registers.GetSP(), "SP. Expected 0x%04X, got 0x%04X")
			} else {
				tests.Equals(t, addr, c.Registers.GetSP(), "SP. Expected 0x%04X, got 0x%04X")
			}
			tests.Equals(t, value1, value2, "PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

//	RETI
func TestRETInstructions4(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xD9]},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			addr := c.Registers.GetSP()
			msb := uint16(c.Memory.Read(addr + 1))
			lsb := uint16(c.Memory.Read(addr))
			value1 := msb<<8 | lsb
			testCase.Instruction.Exec(&c)
			value2 := c.Registers.GetPC()
			tests.Equals(t, uint8(msb), c.Memory.Read(addr+1), "SP-1. Expected 0x%02X, got 0x%02X")
			tests.Equals(t, uint8(lsb), c.Memory.Read(addr), "SP-2. Expected 0x%02X, got 0x%02X")
			tests.Equals(t, addr+2, c.Registers.GetSP(), "SP. Expected 0x%04X, got 0x%04X")
			tests.Equals(t, value1, value2, "PC. Expected 0x%04X, got 0x%04X")
			// Test if Interrupt flag enabled
		})
	}
}

//	RST h
func TestRSTInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		Vector      uint16
	}{
		{c.Instructions[0xC7], 0x00},
		{c.Instructions[0xCF], 0x08},
		{c.Instructions[0xD7], 0x10},
		{c.Instructions[0xDF], 0x18},
		{c.Instructions[0xE7], 0x20},
		{c.Instructions[0xEF], 0x28},
		{c.Instructions[0xF7], 0x30},
		{c.Instructions[0xFF], 0x38},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			value1 := testCase.Vector
			addr := c.Registers.GetSP()
			msb := uint8(c.Registers.GetPC() >> 8)
			lsb := uint8(c.Registers.GetPC() & 0xFF)
			testCase.Instruction.Exec(&c)
			value2 := c.Registers.GetPC()
			tests.Equals(t, msb, c.Memory.Read(addr-1), "SP-1. Expected 0x%02X, got 0x%02X")
			tests.Equals(t, lsb, c.Memory.Read(addr-2), "SP-2. Expected 0x%02X, got 0x%02X")
			tests.Equals(t, addr-2, c.Registers.GetSP(), "SP. Expected 0x%04X, got 0x%04X")
			tests.Equals(t, value1, value2, "PC. Expected 0x%04X, got 0x%04X")
		})
	}
}
