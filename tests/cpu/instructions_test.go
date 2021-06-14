package cpu_test

import (
	"fmt"
	"goboy/cpu"
	"goboy/tests"
	"testing"
)

func TestLDInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To, From    string
	}{
		{c.Instructions[0x40], "B", "B"},
		{c.Instructions[0x41], "B", "C"},
		{c.Instructions[0x42], "B", "D"},
		{c.Instructions[0x43], "B", "E"},
		{c.Instructions[0x44], "B", "H"},
		{c.Instructions[0x45], "B", "L"},
		{c.Instructions[0x47], "B", "A"},
		{c.Instructions[0x48], "C", "B"},
		{c.Instructions[0x49], "C", "C"},
		{c.Instructions[0x4A], "C", "D"},
		{c.Instructions[0x4B], "C", "E"},
		{c.Instructions[0x4C], "C", "H"},
		{c.Instructions[0x4D], "C", "L"},
		{c.Instructions[0x4F], "C", "A"},
		{c.Instructions[0x50], "D", "B"},
		{c.Instructions[0x51], "D", "C"},
		{c.Instructions[0x52], "D", "D"},
		{c.Instructions[0x53], "D", "E"},
		{c.Instructions[0x54], "D", "H"},
		{c.Instructions[0x55], "D", "L"},
		{c.Instructions[0x57], "D", "A"},
		{c.Instructions[0x58], "E", "B"},
		{c.Instructions[0x59], "E", "C"},
		{c.Instructions[0x5A], "E", "D"},
		{c.Instructions[0x5B], "E", "E"},
		{c.Instructions[0x5C], "E", "H"},
		{c.Instructions[0x5D], "E", "L"},
		{c.Instructions[0x5F], "E", "A"},
		{c.Instructions[0x60], "H", "B"},
		{c.Instructions[0x61], "H", "C"},
		{c.Instructions[0x62], "H", "D"},
		{c.Instructions[0x63], "H", "E"},
		{c.Instructions[0x64], "H", "H"},
		{c.Instructions[0x65], "H", "L"},
		{c.Instructions[0x67], "H", "A"},
		{c.Instructions[0x68], "L", "B"},
		{c.Instructions[0x69], "L", "C"},
		{c.Instructions[0x6A], "L", "D"},
		{c.Instructions[0x6B], "L", "E"},
		{c.Instructions[0x6C], "L", "H"},
		{c.Instructions[0x6D], "L", "L"},
		{c.Instructions[0x6F], "L", "A"},
		{c.Instructions[0x78], "A", "B"},
		{c.Instructions[0x79], "A", "C"},
		{c.Instructions[0x7A], "A", "D"},
		{c.Instructions[0x7B], "A", "E"},
		{c.Instructions[0x7C], "A", "H"},
		{c.Instructions[0x7D], "A", "L"},
		{c.Instructions[0x7F], "A", "A"},
	}
	registers := map[string]*uint8{
		"A": &c.Registers.A,
		"B": &c.Registers.B,
		"C": &c.Registers.C,
		"D": &c.Registers.D,
		"E": &c.Registers.E,
		"H": &c.Registers.H,
		"L": &c.Registers.L,
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			c.ReadPC()
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				*registers[testCase.From],
				*registers[testCase.To],
				fmt.Sprintf("Wrong value in register %s. Expected 0x%%02X, got 0x%%02X", testCase.To),
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestLDInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          string
	}{
		{c.Instructions[0x06], "B"},
		{c.Instructions[0x0E], "C"},
		{c.Instructions[0x16], "D"},
		{c.Instructions[0x1E], "E"},
		{c.Instructions[0x26], "H"},
		{c.Instructions[0x2E], "L"},
		{c.Instructions[0x3E], "A"},
	}
	registers := map[string]*uint8{
		"A": &c.Registers.A,
		"B": &c.Registers.B,
		"C": &c.Registers.C,
		"D": &c.Registers.D,
		"E": &c.Registers.E,
		"H": &c.Registers.H,
		"L": &c.Registers.L,
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			c.ReadPC()
			immediateData := c.Memory.ReadMemory(pcBeforeExec + 1)
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				immediateData,
				*registers[testCase.To],
				fmt.Sprintf("Wrong value in register %s. Expected 0x%%02X, got 0x%%02X", testCase.To),
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestLDInstructions3(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          string
	}{
		{c.Instructions[0x46], "B"},
		{c.Instructions[0x4E], "C"},
		{c.Instructions[0x56], "D"},
		{c.Instructions[0x5E], "E"},
		{c.Instructions[0x66], "H"},
		{c.Instructions[0x6E], "L"},
		{c.Instructions[0x7E], "A"},
	}
	registers := map[string]*uint8{
		"A": &c.Registers.A,
		"B": &c.Registers.B,
		"C": &c.Registers.C,
		"D": &c.Registers.D,
		"E": &c.Registers.E,
		"H": &c.Registers.H,
		"L": &c.Registers.L,
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			dataFromAddress := c.ReadMemory(c.Registers.GetHL())
			c.ReadPC()
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				dataFromAddress,
				*registers[testCase.To],
				fmt.Sprintf("Wrong value in register %s. Expected 0x%%02X, got 0x%%02X", testCase.To),
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestLDInstructions4(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		From        string
	}{
		{c.Instructions[0x70], "B"},
		{c.Instructions[0x71], "C"},
		{c.Instructions[0x72], "D"},
		{c.Instructions[0x73], "E"},
		{c.Instructions[0x74], "H"},
		{c.Instructions[0x75], "L"},
		{c.Instructions[0x77], "A"},
	}
	registers := map[string]*uint8{
		"A": &c.Registers.A,
		"B": &c.Registers.B,
		"C": &c.Registers.C,
		"D": &c.Registers.D,
		"E": &c.Registers.E,
		"H": &c.Registers.H,
		"L": &c.Registers.L,
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			address := c.Registers.GetHL()
			c.ReadPC()
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				*registers[testCase.From],
				c.ReadMemory(c.Registers.GetHL()),
				fmt.Sprintf("Wrong value in address 0x%04X. Expected 0x%%02X, got 0x%%02X", address),
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

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
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			address := c.Registers.GetHL()
			c.ReadPC()
			immediateData := c.Memory.ReadMemory(pcBeforeExec + 1)
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				immediateData,
				c.ReadMemory(c.Registers.GetHL()),
				fmt.Sprintf("Wrong value in address 0x%04X. Expected 0x%%02X, got 0x%%02X", address),
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestLDInstructions6(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0x0A]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			address := c.Registers.GetBC()
			value := c.ReadMemory(address)
			c.ReadPC()
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				value,
				c.Registers.GetA(),
				"Wrong value in register A. Expected 0x%02X, got 0x%02X",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestLDInstructions7(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0x1A]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			address := c.Registers.GetDE()
			value := c.ReadMemory(address)
			c.ReadPC()
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				value,
				c.Registers.GetA(),
				"Wrong value in register A. Expected 0x%02X, got 0x%02X",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestLDInstructions8(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0x02]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			c.ReadPC()
			testCase.Instruction.Exec(&c)
			address := c.Registers.GetBC()
			value := c.ReadMemory(address)
			tests.Equals(
				t,
				c.Registers.GetA(),
				value,
				fmt.Sprintf("Wrong value in address 0x%04X. Expected 0x%%02X, got 0x%%02X", address),
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestLDInstructions9(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0x12]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			c.ReadPC()
			testCase.Instruction.Exec(&c)
			address := c.Registers.GetDE()
			value := c.ReadMemory(address)
			tests.Equals(
				t,
				c.Registers.GetA(),
				value,
				fmt.Sprintf("Wrong value in address 0x%04X. Expected 0x%%02X, got 0x%%02X", address),
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestLDInstructions10(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xFA]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			msb := uint16(c.ReadMemory(pcBeforeExec + 1))
			lsb := uint16(c.ReadMemory(pcBeforeExec + 2))
			value := c.ReadMemory(msb<<8 | lsb)
			c.ReadPC()
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				value,
				c.Registers.GetA(),
				"Wrong value in register A. Expected 0x%02X, got 0x%02X",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestLDInstructions11(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xEA]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			msb := uint16(c.ReadMemory(pcBeforeExec + 1))
			lsb := uint16(c.ReadMemory(pcBeforeExec + 2))
			c.ReadPC()
			testCase.Instruction.Exec(&c)
			value := c.ReadMemory(msb<<8 | lsb)
			tests.Equals(
				t,
				c.Registers.GetA(),
				value,
				fmt.Sprintf("Wrong value in address 0x%04X. Expected 0x%%02X, got 0x%%02X", msb<<8|lsb),
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestLDInstructions12(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xF2]},
	}
	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			msb := uint16(0xFF)
			lsb := uint16(c.Registers.GetC())
			value := c.ReadMemory(msb<<8 | lsb)
			c.ReadPC()
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				value,
				c.Registers.GetA(),
				"Wrong value in register A. Expected 0x%02X, got 0x%02X",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestLDInstructions13(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xE2]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			msb := uint16(0xFF)
			lsb := uint16(c.Registers.GetC())
			c.ReadPC()
			testCase.Instruction.Exec(&c)
			value := c.ReadMemory(msb<<8 | lsb)
			tests.Equals(
				t,
				c.Registers.GetA(),
				value,
				fmt.Sprintf("Wrong value in address 0x%04X. Expected 0x%%02X, got 0x%%02X", msb<<8|lsb),
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestLDInstructions14(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xF0]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			msb := uint16(0xFF)
			lsb := uint16(c.ReadMemory(pcBeforeExec + 1))
			value := c.ReadMemory(msb<<8 | lsb)
			c.ReadPC()
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				value,
				c.Registers.GetA(),
				"Wrong value in register A. Expected 0x%02X, got 0x%02X",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestLDInstructions15(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xE0]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			msb := uint16(0xFF)
			lsb := uint16(c.ReadMemory(pcBeforeExec + 1))
			c.ReadPC()
			testCase.Instruction.Exec(&c)
			value := c.ReadMemory(msb<<8 | lsb)
			tests.Equals(
				t,
				c.Registers.GetA(),
				value,
				fmt.Sprintf("Wrong value in address 0x%04X. Expected 0x%%02X, got 0x%%02X", msb<<8|lsb),
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestLDInstructions16(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0x3A]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			address := c.Registers.GetHL()
			dataFromAddress := c.ReadMemory(address)
			c.ReadPC()
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				dataFromAddress,
				c.Registers.GetA(),
				"Wrong value in register A. Expected 0x%02X, got 0x%02X",
			)
			tests.Equals(t, c.Registers.GetHL()+1, address, "Wrong value in register HL. Expected 0x%04X, got 0x%04X")
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestLDInstructions17(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0x32]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			address := c.Registers.GetHL()
			c.ReadPC()
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				c.Registers.GetA(),
				c.ReadMemory(address),
				fmt.Sprintf("Wrong value in address 0x%04X. Expected 0x%%02X, got 0x%%02X", address),
			)
			tests.Equals(t, c.Registers.GetHL()+1, address, "Wrong value in register HL. Expected 0x%04X, got 0x%04X")
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestLDInstructions18(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0x2A]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			address := c.Registers.GetHL()
			dataFromAddress := c.ReadMemory(address)
			c.ReadPC()
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				dataFromAddress,
				c.Registers.GetA(),
				"Wrong value in register A. Expected 0x%02X, got 0x%02X",
			)
			tests.Equals(t, c.Registers.GetHL()-1, address, "Wrong value in register HL. Expected 0x%04X, got 0x%04X")
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestLDInstructions19(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0x22]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			address := c.Registers.GetHL()
			c.ReadPC()
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				c.Registers.GetA(),
				c.ReadMemory(address),
				fmt.Sprintf("Wrong value in address 0x%04X. Expected 0x%%02X, got 0x%%02X", address),
			)
			tests.Equals(t, c.Registers.GetHL()-1, address, "Wrong value in register HL. Expected 0x%04X, got 0x%04X")
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestLDInstructions20(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          string
	}{
		{c.Instructions[0x01], "BC"},
		{c.Instructions[0x11], "DE"},
		{c.Instructions[0x21], "HL"},
		{c.Instructions[0x31], "SP"},
	}
	registers := map[string]func() uint16{
		"BC": c.Registers.GetBC,
		"DE": c.Registers.GetDE,
		"HL": c.Registers.GetHL,
		"SP": c.Registers.GetSP,
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			c.ReadPC()
			msb := uint16(c.ReadMemory(pcBeforeExec + 1))
			lsb := uint16(c.ReadMemory(pcBeforeExec + 2))
			immediateData := msb<<8 | lsb
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				immediateData,
				registers[testCase.To](),
				fmt.Sprintf("Wrong value in register %s. Expected 0x%%04X, got 0x%%04X", testCase.To),
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestLDInstructions21(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0x08]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			msb := uint16(c.ReadMemory(pcBeforeExec + 1))
			lsb := uint16(c.ReadMemory(pcBeforeExec + 2))
			c.ReadPC()
			testCase.Instruction.Exec(&c)
			addr := msb<<8 | lsb
			value1 := c.ReadMemory(addr)
			value2 := c.ReadMemory(addr + 1)
			tests.Equals(
				t,
				c.Registers.GetP(),
				value1,
				fmt.Sprintf("Wrong value in address 0x%04X. Expected 0x%%02X, got 0x%%02X", addr),
			)
			tests.Equals(
				t,
				c.Registers.GetS(),
				value2,
				fmt.Sprintf("Wrong value in address 0x%04X. Expected 0x%%02X, got 0x%%02X", addr+1),
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestLDInstructions22(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xF8]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			c.ReadPC()
			immediateData := int(c.ReadMemory(pcBeforeExec + 1))
			immediateData = (immediateData & 127) - (immediateData & 128)
			value := uint16(int(c.Registers.GetSP()) + immediateData)
			testCase.Instruction.Exec(&c)
			// if immediateData > int(c.Registers.SP) && !c.Flags.H {
			// 	t.Logf("S8: 0x%04b (%v) (positive: %v)\n", immediateData&0xF, int(immediateData&0xF), immediateData > 0)
			// 	t.Logf("SP: 0x%04b (%v)\n", c.Registers.SP&0xF, int(c.Registers.SP&0xF))
			// 	t.Logf("HC: %v\n", c.Flags.H)
			// }
			flagZ := false
			flagN := false
			// if immediateData < 0 {
			// 	c.Flags.SetH((sp & 0xF) < (immediateData & 0xF))
			// 	c.Flags.SetC((sp & 0xFF) < (immediateData & 0xFF))
			// } else {
			// 	c.Flags.SetH((sp&0xF)+(s8&0xF) > 0xF)
			// 	c.Flags.SetC((sp&0xFF)+(s8&0xFF) > 0xFF)
			// }
			// flagH := (c.Registers.GetA()&0xF)+(*registers[testCase.From]&0xF) > 0xF
			// flagC := uint16(c.Registers.GetA())+uint16(*registers[testCase.From]) > 0xFF
			tests.Equals(
				t,
				value,
				c.Registers.GetHL(),
				"Wrong value in register HL. Expected 0x%02X, got 0x%02X",
			)
			tests.Equals(
				t,
				flagZ,
				c.Flags.Z,
				"Wrong value in Flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagN,
				c.Flags.N,
				"Wrong value in Flag N. Expected %v, got %v",
			)
			// tests.Equals(
			// 	t,
			// 	flagH,
			// 	c.Flags.H,
			// 	"Wrong value in Flag H. Expected %v, got %v",
			// )
			// tests.Equals(
			// 	t,
			// 	flagC,
			// 	c.Flags.C,
			// 	"Wrong value in Flag C. Expected %v, got %v",
			// )
			flags := c.Registers.GetF()
			// t.Logf("REG_F: 0x%08b\n", c.Registers.GetF())
			// t.Logf("FLAGS: 0x%08b\n", c.Flags.GetFlagsAsValue())
			tests.Equals(
				t,
				flags&(1<<7) == 1<<7,
				c.Flags.Z,
				"Wrong value in flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<6) == 1<<6,
				c.Flags.N,
				"Wrong value in flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<5) == 1<<5,
				c.Flags.H,
				"Wrong value in flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<4) == 1<<4,
				c.Flags.C,
				"Wrong value in flag C. Expected %v, got %v",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestLDInstructions23(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xF9]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			c.ReadPC()
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				c.Registers.GetHL(),
				c.Registers.GetSP(),
				"Wrong value in register SP. Expected 0x%02X, got 0x%02X",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestPUSHInstructions(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          string
	}{
		{c.Instructions[0xC5], "BC"},
		{c.Instructions[0xD5], "DE"},
		{c.Instructions[0xE5], "HL"},
		{c.Instructions[0xF5], "AF"},
	}
	registers := map[string]func() uint16{
		"BC": c.Registers.GetBC,
		"DE": c.Registers.GetDE,
		"HL": c.Registers.GetHL,
		"AF": c.Registers.GetAF,
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			c.ReadPC()
			addr := c.Registers.GetSP()
			testCase.Instruction.Exec(&c)
			msb := uint16(c.ReadMemory(addr - 1))
			lsb := uint16(c.ReadMemory(addr - 2))
			tests.Equals(
				t,
				registers[testCase.To](),
				msb<<8|lsb,
				fmt.Sprintf("Wrong value in address 0x%04X. Expected 0x%%04X, got 0x%%04X", addr-1),
			)
			tests.Equals(
				t,
				addr,
				c.Registers.GetSP()+2,
				"Wrong value in register SP. Expected 0x%02X, got 0x%02X",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestPOPInstructions(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		To          string
	}{
		{c.Instructions[0xC1], "BC"},
		{c.Instructions[0xD1], "DE"},
		{c.Instructions[0xE1], "HL"},
		{c.Instructions[0xF1], "AF"},
	}
	registers := map[string]func() uint16{
		"BC": c.Registers.GetBC,
		"DE": c.Registers.GetDE,
		"HL": c.Registers.GetHL,
		"AF": c.Registers.GetAF,
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			c.ReadPC()
			addr := c.Registers.GetSP()
			testCase.Instruction.Exec(&c)
			msb := uint16(c.ReadMemory(addr + 1))
			lsb := uint16(c.ReadMemory(addr))
			tests.Equals(
				t,
				msb<<8|lsb,
				registers[testCase.To](),
				fmt.Sprintf("Wrong value in address 0x%04X. Expected 0x%%04X, got 0x%%04X", addr-1),
			)
			tests.Equals(
				t,
				addr,
				c.Registers.GetSP()-2,
				"Wrong value in register SP. Expected 0x%02X, got 0x%02X",
			)
			if testCase.To == "AF" {
				flags := c.Registers.GetF()
				// t.Logf("REG_F: 0x%08b\n", c.Registers.GetF())
				// t.Logf("FLAGS: 0x%08b\n", c.Flags.GetFlagsAsValue())
				tests.Equals(
					t,
					flags&(1<<7) == 1<<7,
					c.Flags.Z,
					"Wrong value in flag Z. Expected %v, got %v",
				)
				tests.Equals(
					t,
					flags&(1<<6) == 1<<6,
					c.Flags.N,
					"Wrong value in flag N. Expected %v, got %v",
				)
				tests.Equals(
					t,
					flags&(1<<5) == 1<<5,
					c.Flags.H,
					"Wrong value in flag H. Expected %v, got %v",
				)
				tests.Equals(
					t,
					flags&(1<<4) == 1<<4,
					c.Flags.C,
					"Wrong value in flag C. Expected %v, got %v",
				)
			}
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestADDInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		From        string
	}{
		{c.Instructions[0x80], "B"},
		{c.Instructions[0x81], "C"},
		{c.Instructions[0x82], "D"},
		{c.Instructions[0x83], "E"},
		{c.Instructions[0x84], "H"},
		{c.Instructions[0x85], "L"},
		{c.Instructions[0x87], "A"},
	}
	registers := map[string]*uint8{
		"B": &c.Registers.B,
		"C": &c.Registers.C,
		"D": &c.Registers.D,
		"E": &c.Registers.E,
		"H": &c.Registers.H,
		"L": &c.Registers.L,
		"A": &c.Registers.A,
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			c.ReadPC()
			value := uint8(c.Registers.GetA() + *registers[testCase.From])
			flagZ := value == 0
			flagN := false
			flagH := (c.Registers.GetA()&0xF)+(*registers[testCase.From]&0xF) > 0xF
			flagC := uint16(c.Registers.GetA())+uint16(*registers[testCase.From]) > 0xFF
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				flagZ,
				c.Flags.Z,
				"Wrong value in Flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagN,
				c.Flags.N,
				"Wrong value in Flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagH,
				c.Flags.H,
				"Wrong value in Flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagC,
				c.Flags.C,
				"Wrong value in Flag C. Expected %v, got %v",
			)
			tests.Equals(
				t,
				value,
				c.Registers.GetA(),
				"Wrong value in register A. Expected 0x%02X, got 0x%02X",
			)
			flags := c.Registers.GetF()
			// t.Logf("REG_F: 0x%08b\n", c.Registers.GetF())
			// t.Logf("FLAGS: 0x%08b\n", c.Flags.GetFlagsAsValue())
			tests.Equals(
				t,
				flags&(1<<7) == 1<<7,
				c.Flags.Z,
				"Wrong value in flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<6) == 1<<6,
				c.Flags.N,
				"Wrong value in flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<5) == 1<<5,
				c.Flags.H,
				"Wrong value in flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<4) == 1<<4,
				c.Flags.C,
				"Wrong value in flag C. Expected %v, got %v",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestADDInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xC6]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			immediateData := c.ReadMemory(pcBeforeExec + 1)
			c.ReadPC()
			value := uint8(c.Registers.GetA() + immediateData)
			flagZ := value == 0
			flagN := false
			flagH := (c.Registers.GetA()&0xF)+(immediateData&0xF) > 0xF
			flagC := uint16(c.Registers.GetA())+uint16(immediateData) > 0xFF
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				flagZ,
				c.Flags.Z,
				"Wrong value in Flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagN,
				c.Flags.N,
				"Wrong value in Flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagH,
				c.Flags.H,
				"Wrong value in Flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagC,
				c.Flags.C,
				"Wrong value in Flag C. Expected %v, got %v",
			)
			tests.Equals(
				t,
				value,
				c.Registers.GetA(),
				"Wrong value in register A. Expected 0x%02X, got 0x%02X",
			)
			flags := c.Registers.GetF()
			// t.Logf("REG_F: 0x%08b\n", c.Registers.GetF())
			// t.Logf("FLAGS: 0x%08b\n", c.Flags.GetFlagsAsValue())
			tests.Equals(
				t,
				flags&(1<<7) == 1<<7,
				c.Flags.Z,
				"Wrong value in flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<6) == 1<<6,
				c.Flags.N,
				"Wrong value in flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<5) == 1<<5,
				c.Flags.H,
				"Wrong value in flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<4) == 1<<4,
				c.Flags.C,
				"Wrong value in flag C. Expected %v, got %v",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestADDInstructions3(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0x86]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			immediateData := c.ReadMemory(c.Registers.GetHL())
			c.ReadPC()
			value := uint8(c.Registers.GetA() + immediateData)
			flagZ := value == 0
			flagN := false
			flagH := (c.Registers.GetA()&0xF)+(immediateData&0xF) > 0xF
			flagC := uint16(c.Registers.GetA())+uint16(immediateData) > 0xFF
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				flagZ,
				c.Flags.Z,
				"Wrong value in Flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagN,
				c.Flags.N,
				"Wrong value in Flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagH,
				c.Flags.H,
				"Wrong value in Flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagC,
				c.Flags.C,
				"Wrong value in Flag C. Expected %v, got %v",
			)
			tests.Equals(
				t,
				value,
				c.Registers.GetA(),
				"Wrong value in register A. Expected 0x%02X, got 0x%02X",
			)
			flags := c.Registers.GetF()
			// t.Logf("REG_F: 0x%08b\n", c.Registers.GetF())
			// t.Logf("FLAGS: 0x%08b\n", c.Flags.GetFlagsAsValue())
			tests.Equals(
				t,
				flags&(1<<7) == 1<<7,
				c.Flags.Z,
				"Wrong value in flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<6) == 1<<6,
				c.Flags.N,
				"Wrong value in flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<5) == 1<<5,
				c.Flags.H,
				"Wrong value in flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<4) == 1<<4,
				c.Flags.C,
				"Wrong value in flag C. Expected %v, got %v",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestADDInstructions4(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		From        string
	}{
		{c.Instructions[0x09], "BC"},
		{c.Instructions[0x19], "DE"},
		{c.Instructions[0x29], "HL"},
		{c.Instructions[0x39], "SP"},
	}
	registers := map[string]func() uint16{
		"BC": c.Registers.GetBC,
		"DE": c.Registers.GetDE,
		"HL": c.Registers.GetHL,
		"SP": c.Registers.GetSP,
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			c.ReadPC()
			value := c.Registers.GetHL() + registers[testCase.From]()
			flagN := false
			flagH := (c.Registers.GetHL()&0xFFF)+(registers[testCase.From]()&0xFFF) > 0xFFF
			flagC := uint32(c.Registers.GetHL())+uint32(registers[testCase.From]()) > 0xFFFF
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				flagN,
				c.Flags.N,
				"Wrong value in Flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagH,
				c.Flags.H,
				"Wrong value in Flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagC,
				c.Flags.C,
				"Wrong value in Flag C. Expected %v, got %v",
			)
			tests.Equals(
				t,
				value,
				c.Registers.GetHL(),
				"Wrong value in register HL. Expected 0x%04X, got 0x%04X",
			)
			flags := c.Registers.GetF()
			// t.Logf("REG_F: 0x%08b\n", c.Registers.GetF())
			// t.Logf("FLAGS: 0x%08b\n", c.Flags.GetFlagsAsValue())
			tests.Equals(
				t,
				flags&(1<<7) == 1<<7,
				c.Flags.Z,
				"Wrong value in flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<6) == 1<<6,
				c.Flags.N,
				"Wrong value in flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<5) == 1<<5,
				c.Flags.H,
				"Wrong value in flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<4) == 1<<4,
				c.Flags.C,
				"Wrong value in flag C. Expected %v, got %v",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestADDInstructions5(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xE8]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			c.ReadPC()
			immediateData := int(c.ReadMemory(pcBeforeExec + 1))
			immediateData = (immediateData & 127) - (immediateData & 128)
			value := uint16(int(c.Registers.GetSP()) + immediateData)
			testCase.Instruction.Exec(&c)
			// if immediateData > int(c.Registers.SP) && !c.Flags.H {
			// 	t.Logf("S8: 0x%04b (%v) (positive: %v)\n", immediateData&0xF, int(immediateData&0xF), immediateData > 0)
			// 	t.Logf("SP: 0x%04b (%v)\n", c.Registers.SP&0xF, int(c.Registers.SP&0xF))
			// 	t.Logf("HC: %v\n", c.Flags.H)
			// }
			flagZ := false
			flagN := false
			// if immediateData < 0 {
			// 	c.Flags.SetH((sp & 0xF) < (immediateData & 0xF))
			// 	c.Flags.SetC((sp & 0xFF) < (immediateData & 0xFF))
			// } else {
			// 	c.Flags.SetH((sp&0xF)+(s8&0xF) > 0xF)
			// 	c.Flags.SetC((sp&0xFF)+(s8&0xFF) > 0xFF)
			// }
			// flagH := (c.Registers.GetA()&0xF)+(*registers[testCase.From]&0xF) > 0xF
			// flagC := uint16(c.Registers.GetA())+uint16(*registers[testCase.From]) > 0xFF
			tests.Equals(
				t,
				value,
				c.Registers.GetSP(),
				"Wrong value in register SP. Expected 0x%02X, got 0x%02X",
			)
			tests.Equals(
				t,
				flagZ,
				c.Flags.Z,
				"Wrong value in Flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagN,
				c.Flags.N,
				"Wrong value in Flag N. Expected %v, got %v",
			)
			// tests.Equals(
			// 	t,
			// 	flagH,
			// 	c.Flags.H,
			// 	"Wrong value in Flag H. Expected %v, got %v",
			// )
			// tests.Equals(
			// 	t,
			// 	flagC,
			// 	c.Flags.C,
			// 	"Wrong value in Flag C. Expected %v, got %v",
			// )
			flags := c.Registers.GetF()
			// t.Logf("REG_F: 0x%08b\n", c.Registers.GetF())
			// t.Logf("FLAGS: 0x%08b\n", c.Flags.GetFlagsAsValue())
			tests.Equals(
				t,
				flags&(1<<7) == 1<<7,
				c.Flags.Z,
				"Wrong value in flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<6) == 1<<6,
				c.Flags.N,
				"Wrong value in flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<5) == 1<<5,
				c.Flags.H,
				"Wrong value in flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<4) == 1<<4,
				c.Flags.C,
				"Wrong value in flag C. Expected %v, got %v",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestADCInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		From        string
	}{
		{c.Instructions[0x88], "B"},
		{c.Instructions[0x89], "C"},
		{c.Instructions[0x8A], "D"},
		{c.Instructions[0x8B], "E"},
		{c.Instructions[0x8C], "H"},
		{c.Instructions[0x8D], "L"},
		{c.Instructions[0x8F], "A"},
	}
	registers := map[string]*uint8{
		"B": &c.Registers.B,
		"C": &c.Registers.C,
		"D": &c.Registers.D,
		"E": &c.Registers.E,
		"H": &c.Registers.H,
		"L": &c.Registers.L,
		"A": &c.Registers.A,
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			c.ReadPC()
			carry := c.Flags.GetCarryAsValue()
			value := uint8(c.Registers.GetA() + *registers[testCase.From] + carry)
			flagZ := value == 0
			flagN := false
			flagH := (c.Registers.GetA()&0xF)+(*registers[testCase.From]&0xF)+carry > 0xF
			flagC := uint16(c.Registers.GetA())+uint16(*registers[testCase.From])+uint16(carry) > 0xFF
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				flagZ,
				c.Flags.Z,
				"Wrong value in Flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagN,
				c.Flags.N,
				"Wrong value in Flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagH,
				c.Flags.H,
				"Wrong value in Flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagC,
				c.Flags.C,
				"Wrong value in Flag C. Expected %v, got %v",
			)
			tests.Equals(
				t,
				value,
				c.Registers.GetA(),
				"Wrong value in register A. Expected 0x%02X, got 0x%02X",
			)
			flags := c.Registers.GetF()
			// t.Logf("REG_F: 0x%08b\n", c.Registers.GetF())
			// t.Logf("FLAGS: 0x%08b\n", c.Flags.GetFlagsAsValue())
			tests.Equals(
				t,
				flags&(1<<7) == 1<<7,
				c.Flags.Z,
				"Wrong value in flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<6) == 1<<6,
				c.Flags.N,
				"Wrong value in flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<5) == 1<<5,
				c.Flags.H,
				"Wrong value in flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<4) == 1<<4,
				c.Flags.C,
				"Wrong value in flag C. Expected %v, got %v",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestADCInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xCE]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			immediateData := c.ReadMemory(pcBeforeExec + 1)
			c.ReadPC()
			carry := c.Flags.GetCarryAsValue()
			value := uint8(c.Registers.GetA() + immediateData + carry)
			flagZ := value == 0
			flagN := false
			flagH := (c.Registers.GetA()&0xF)+(immediateData&0xF)+carry > 0xF
			flagC := uint16(c.Registers.GetA())+uint16(immediateData)+uint16(carry) > 0xFF
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				flagZ,
				c.Flags.Z,
				"Wrong value in Flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagN,
				c.Flags.N,
				"Wrong value in Flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagH,
				c.Flags.H,
				"Wrong value in Flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagC,
				c.Flags.C,
				"Wrong value in Flag C. Expected %v, got %v",
			)
			tests.Equals(
				t,
				value,
				c.Registers.GetA(),
				"Wrong value in register A. Expected 0x%02X, got 0x%02X",
			)
			flags := c.Registers.GetF()
			// t.Logf("REG_F: 0x%08b\n", c.Registers.GetF())
			// t.Logf("FLAGS: 0x%08b\n", c.Flags.GetFlagsAsValue())
			tests.Equals(
				t,
				flags&(1<<7) == 1<<7,
				c.Flags.Z,
				"Wrong value in flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<6) == 1<<6,
				c.Flags.N,
				"Wrong value in flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<5) == 1<<5,
				c.Flags.H,
				"Wrong value in flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<4) == 1<<4,
				c.Flags.C,
				"Wrong value in flag C. Expected %v, got %v",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestADCInstructions3(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0x8E]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			immediateData := c.ReadMemory(c.Registers.GetHL())
			c.ReadPC()
			carry := c.Flags.GetCarryAsValue()
			value := uint8(c.Registers.GetA() + immediateData + carry)
			flagZ := value == 0
			flagN := false
			flagH := (c.Registers.GetA()&0xF)+(immediateData&0xF)+carry > 0xF
			flagC := uint16(c.Registers.GetA())+uint16(immediateData)+uint16(carry) > 0xFF
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				flagZ,
				c.Flags.Z,
				"Wrong value in Flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagN,
				c.Flags.N,
				"Wrong value in Flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagH,
				c.Flags.H,
				"Wrong value in Flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagC,
				c.Flags.C,
				"Wrong value in Flag C. Expected %v, got %v",
			)
			tests.Equals(
				t,
				value,
				c.Registers.GetA(),
				"Wrong value in register A. Expected 0x%02X, got 0x%02X",
			)
			flags := c.Registers.GetF()
			// t.Logf("REG_F: 0x%08b\n", c.Registers.GetF())
			// t.Logf("FLAGS: 0x%08b\n", c.Flags.GetFlagsAsValue())
			tests.Equals(
				t,
				flags&(1<<7) == 1<<7,
				c.Flags.Z,
				"Wrong value in flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<6) == 1<<6,
				c.Flags.N,
				"Wrong value in flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<5) == 1<<5,
				c.Flags.H,
				"Wrong value in flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<4) == 1<<4,
				c.Flags.C,
				"Wrong value in flag C. Expected %v, got %v",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestSUBInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		From        string
	}{
		{c.Instructions[0x90], "B"},
		{c.Instructions[0x91], "C"},
		{c.Instructions[0x92], "D"},
		{c.Instructions[0x93], "E"},
		{c.Instructions[0x94], "H"},
		{c.Instructions[0x95], "L"},
		{c.Instructions[0x97], "A"},
	}
	registers := map[string]*uint8{
		"B": &c.Registers.B,
		"C": &c.Registers.C,
		"D": &c.Registers.D,
		"E": &c.Registers.E,
		"H": &c.Registers.H,
		"L": &c.Registers.L,
		"A": &c.Registers.A,
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			c.ReadPC()
			value := uint8(c.Registers.GetA() - *registers[testCase.From])
			flagZ := value == 0
			flagN := true
			flagH := (c.Registers.GetA() & 0xF) < (*registers[testCase.From] & 0xF)
			flagC := uint16(c.Registers.GetA()) < uint16(*registers[testCase.From])
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				flagZ,
				c.Flags.Z,
				"Wrong value in Flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagN,
				c.Flags.N,
				"Wrong value in Flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagH,
				c.Flags.H,
				"Wrong value in Flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagC,
				c.Flags.C,
				"Wrong value in Flag C. Expected %v, got %v",
			)
			tests.Equals(
				t,
				value,
				c.Registers.GetA(),
				"Wrong value in register A. Expected 0x%02X, got 0x%02X",
			)
			flags := c.Registers.GetF()
			// t.Logf("REG_F: 0x%08b\n", c.Registers.GetF())
			// t.Logf("FLAGS: 0x%08b\n", c.Flags.GetFlagsAsValue())
			tests.Equals(
				t,
				flags&(1<<7) == 1<<7,
				c.Flags.Z,
				"Wrong value in flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<6) == 1<<6,
				c.Flags.N,
				"Wrong value in flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<5) == 1<<5,
				c.Flags.H,
				"Wrong value in flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<4) == 1<<4,
				c.Flags.C,
				"Wrong value in flag C. Expected %v, got %v",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestSUBInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xD6]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			immediateData := c.ReadMemory(pcBeforeExec + 1)
			c.ReadPC()
			value := uint8(c.Registers.GetA() - immediateData)
			flagZ := value == 0
			flagN := true
			flagH := (c.Registers.GetA() & 0xF) < (immediateData & 0xF)
			flagC := uint16(c.Registers.GetA()) < uint16(immediateData)
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				flagZ,
				c.Flags.Z,
				"Wrong value in Flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagN,
				c.Flags.N,
				"Wrong value in Flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagH,
				c.Flags.H,
				"Wrong value in Flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagC,
				c.Flags.C,
				"Wrong value in Flag C. Expected %v, got %v",
			)
			tests.Equals(
				t,
				value,
				c.Registers.GetA(),
				"Wrong value in register A. Expected 0x%02X, got 0x%02X",
			)
			flags := c.Registers.GetF()
			// t.Logf("REG_F: 0x%08b\n", c.Registers.GetF())
			// t.Logf("FLAGS: 0x%08b\n", c.Flags.GetFlagsAsValue())
			tests.Equals(
				t,
				flags&(1<<7) == 1<<7,
				c.Flags.Z,
				"Wrong value in flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<6) == 1<<6,
				c.Flags.N,
				"Wrong value in flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<5) == 1<<5,
				c.Flags.H,
				"Wrong value in flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<4) == 1<<4,
				c.Flags.C,
				"Wrong value in flag C. Expected %v, got %v",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestSUBInstructions3(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0x96]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			immediateData := c.ReadMemory(c.Registers.GetHL())
			c.ReadPC()
			value := uint8(c.Registers.GetA() - immediateData)
			flagZ := value == 0
			flagN := true
			flagH := (c.Registers.GetA() & 0xF) < (immediateData & 0xF)
			flagC := uint16(c.Registers.GetA()) < uint16(immediateData)
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				flagZ,
				c.Flags.Z,
				"Wrong value in Flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagN,
				c.Flags.N,
				"Wrong value in Flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagH,
				c.Flags.H,
				"Wrong value in Flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagC,
				c.Flags.C,
				"Wrong value in Flag C. Expected %v, got %v",
			)
			tests.Equals(
				t,
				value,
				c.Registers.GetA(),
				"Wrong value in register A. Expected 0x%02X, got 0x%02X",
			)
			flags := c.Registers.GetF()
			// t.Logf("REG_F: 0x%08b\n", c.Registers.GetF())
			// t.Logf("FLAGS: 0x%08b\n", c.Flags.GetFlagsAsValue())
			tests.Equals(
				t,
				flags&(1<<7) == 1<<7,
				c.Flags.Z,
				"Wrong value in flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<6) == 1<<6,
				c.Flags.N,
				"Wrong value in flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<5) == 1<<5,
				c.Flags.H,
				"Wrong value in flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<4) == 1<<4,
				c.Flags.C,
				"Wrong value in flag C. Expected %v, got %v",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestSBCInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		From        string
	}{
		{c.Instructions[0x98], "B"},
		{c.Instructions[0x99], "C"},
		{c.Instructions[0x9A], "D"},
		{c.Instructions[0x9B], "E"},
		{c.Instructions[0x9C], "H"},
		{c.Instructions[0x9D], "L"},
		{c.Instructions[0x9F], "A"},
	}
	registers := map[string]*uint8{
		"B": &c.Registers.B,
		"C": &c.Registers.C,
		"D": &c.Registers.D,
		"E": &c.Registers.E,
		"H": &c.Registers.H,
		"L": &c.Registers.L,
		"A": &c.Registers.A,
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			c.ReadPC()
			carry := c.Flags.GetCarryAsValue()
			value := uint8(c.Registers.GetA() - *registers[testCase.From] - carry)
			flagZ := value == 0
			flagN := true
			flagH := (c.Registers.GetA() & 0xF) < ((*registers[testCase.From] & 0xF) + carry)
			flagC := uint16(c.Registers.GetA()) < (uint16(*registers[testCase.From]) + uint16(carry))
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				flagZ,
				c.Flags.Z,
				"Wrong value in Flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagN,
				c.Flags.N,
				"Wrong value in Flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagH,
				c.Flags.H,
				"Wrong value in Flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagC,
				c.Flags.C,
				"Wrong value in Flag C. Expected %v, got %v",
			)
			tests.Equals(
				t,
				value,
				c.Registers.GetA(),
				"Wrong value in register A. Expected 0x%02X, got 0x%02X",
			)
			flags := c.Registers.GetF()
			// t.Logf("REG_F: 0x%08b\n", c.Registers.GetF())
			// t.Logf("FLAGS: 0x%08b\n", c.Flags.GetFlagsAsValue())
			tests.Equals(
				t,
				flags&(1<<7) == 1<<7,
				c.Flags.Z,
				"Wrong value in flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<6) == 1<<6,
				c.Flags.N,
				"Wrong value in flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<5) == 1<<5,
				c.Flags.H,
				"Wrong value in flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<4) == 1<<4,
				c.Flags.C,
				"Wrong value in flag C. Expected %v, got %v",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestSBCInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xDE]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			immediateData := c.ReadMemory(pcBeforeExec + 1)
			c.ReadPC()
			carry := c.Flags.GetCarryAsValue()
			value := uint8(c.Registers.GetA() - immediateData - carry)
			flagZ := value == 0
			flagN := true
			flagH := (c.Registers.GetA() & 0xF) < ((immediateData & 0xF) + carry)
			flagC := uint16(c.Registers.GetA()) < (uint16(immediateData) + uint16(carry))
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				flagZ,
				c.Flags.Z,
				"Wrong value in Flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagN,
				c.Flags.N,
				"Wrong value in Flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagH,
				c.Flags.H,
				"Wrong value in Flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagC,
				c.Flags.C,
				"Wrong value in Flag C. Expected %v, got %v",
			)
			tests.Equals(
				t,
				value,
				c.Registers.GetA(),
				"Wrong value in register A. Expected 0x%02X, got 0x%02X",
			)
			flags := c.Registers.GetF()
			// t.Logf("REG_F: 0x%08b\n", c.Registers.GetF())
			// t.Logf("FLAGS: 0x%08b\n", c.Flags.GetFlagsAsValue())
			tests.Equals(
				t,
				flags&(1<<7) == 1<<7,
				c.Flags.Z,
				"Wrong value in flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<6) == 1<<6,
				c.Flags.N,
				"Wrong value in flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<5) == 1<<5,
				c.Flags.H,
				"Wrong value in flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<4) == 1<<4,
				c.Flags.C,
				"Wrong value in flag C. Expected %v, got %v",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestSBCInstructions3(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0x9E]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			immediateData := c.ReadMemory(c.Registers.GetHL())
			c.ReadPC()
			carry := c.Flags.GetCarryAsValue()
			value := uint8(c.Registers.GetA() - immediateData - carry)
			flagZ := value == 0
			flagN := true
			flagH := (c.Registers.GetA() & 0xF) < ((immediateData & 0xF) + carry)
			flagC := uint16(c.Registers.GetA()) < (uint16(immediateData) + uint16(carry))
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				flagZ,
				c.Flags.Z,
				"Wrong value in Flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagN,
				c.Flags.N,
				"Wrong value in Flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagH,
				c.Flags.H,
				"Wrong value in Flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagC,
				c.Flags.C,
				"Wrong value in Flag C. Expected %v, got %v",
			)
			tests.Equals(
				t,
				value,
				c.Registers.GetA(),
				"Wrong value in register A. Expected 0x%02X, got 0x%02X",
			)
			flags := c.Registers.GetF()
			// t.Logf("REG_F: 0x%08b\n", c.Registers.GetF())
			// t.Logf("FLAGS: 0x%08b\n", c.Flags.GetFlagsAsValue())
			tests.Equals(
				t,
				flags&(1<<7) == 1<<7,
				c.Flags.Z,
				"Wrong value in flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<6) == 1<<6,
				c.Flags.N,
				"Wrong value in flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<5) == 1<<5,
				c.Flags.H,
				"Wrong value in flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<4) == 1<<4,
				c.Flags.C,
				"Wrong value in flag C. Expected %v, got %v",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestANDInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		From        string
	}{
		{c.Instructions[0xA0], "B"},
		{c.Instructions[0xA1], "C"},
		{c.Instructions[0xA2], "D"},
		{c.Instructions[0xA3], "E"},
		{c.Instructions[0xA4], "H"},
		{c.Instructions[0xA5], "L"},
		{c.Instructions[0xA7], "A"},
	}
	registers := map[string]*uint8{
		"B": &c.Registers.B,
		"C": &c.Registers.C,
		"D": &c.Registers.D,
		"E": &c.Registers.E,
		"H": &c.Registers.H,
		"L": &c.Registers.L,
		"A": &c.Registers.A,
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			c.ReadPC()
			value := uint8(c.Registers.GetA() & *registers[testCase.From])
			flagZ := value == 0
			flagN := false
			flagH := true
			flagC := false
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				flagZ,
				c.Flags.Z,
				"Wrong value in Flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagN,
				c.Flags.N,
				"Wrong value in Flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagH,
				c.Flags.H,
				"Wrong value in Flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagC,
				c.Flags.C,
				"Wrong value in Flag C. Expected %v, got %v",
			)
			tests.Equals(
				t,
				value,
				c.Registers.GetA(),
				"Wrong value in register A. Expected 0x%02X, got 0x%02X",
			)
			flags := c.Registers.GetF()
			// t.Logf("REG_F: 0x%08b\n", c.Registers.GetF())
			// t.Logf("FLAGS: 0x%08b\n", c.Flags.GetFlagsAsValue())
			tests.Equals(
				t,
				flags&(1<<7) == 1<<7,
				c.Flags.Z,
				"Wrong value in flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<6) == 1<<6,
				c.Flags.N,
				"Wrong value in flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<5) == 1<<5,
				c.Flags.H,
				"Wrong value in flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<4) == 1<<4,
				c.Flags.C,
				"Wrong value in flag C. Expected %v, got %v",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestANDInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xE6]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			immediateData := c.ReadMemory(pcBeforeExec + 1)
			c.ReadPC()
			value := uint8(c.Registers.GetA() & immediateData)
			flagZ := value == 0
			flagN := false
			flagH := true
			flagC := false
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				flagZ,
				c.Flags.Z,
				"Wrong value in Flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagN,
				c.Flags.N,
				"Wrong value in Flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagH,
				c.Flags.H,
				"Wrong value in Flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagC,
				c.Flags.C,
				"Wrong value in Flag C. Expected %v, got %v",
			)
			tests.Equals(
				t,
				value,
				c.Registers.GetA(),
				"Wrong value in register A. Expected 0x%02X, got 0x%02X",
			)
			flags := c.Registers.GetF()
			// t.Logf("REG_F: 0x%08b\n", c.Registers.GetF())
			// t.Logf("FLAGS: 0x%08b\n", c.Flags.GetFlagsAsValue())
			tests.Equals(
				t,
				flags&(1<<7) == 1<<7,
				c.Flags.Z,
				"Wrong value in flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<6) == 1<<6,
				c.Flags.N,
				"Wrong value in flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<5) == 1<<5,
				c.Flags.H,
				"Wrong value in flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<4) == 1<<4,
				c.Flags.C,
				"Wrong value in flag C. Expected %v, got %v",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestANDInstructions3(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xA6]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			immediateData := c.ReadMemory(c.Registers.GetHL())
			c.ReadPC()
			value := uint8(c.Registers.GetA() & immediateData)
			flagZ := value == 0
			flagN := false
			flagH := true
			flagC := false
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				flagZ,
				c.Flags.Z,
				"Wrong value in Flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagN,
				c.Flags.N,
				"Wrong value in Flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagH,
				c.Flags.H,
				"Wrong value in Flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagC,
				c.Flags.C,
				"Wrong value in Flag C. Expected %v, got %v",
			)
			tests.Equals(
				t,
				value,
				c.Registers.GetA(),
				"Wrong value in register A. Expected 0x%02X, got 0x%02X",
			)
			flags := c.Registers.GetF()
			// t.Logf("REG_F: 0x%08b\n", c.Registers.GetF())
			// t.Logf("FLAGS: 0x%08b\n", c.Flags.GetFlagsAsValue())
			tests.Equals(
				t,
				flags&(1<<7) == 1<<7,
				c.Flags.Z,
				"Wrong value in flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<6) == 1<<6,
				c.Flags.N,
				"Wrong value in flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<5) == 1<<5,
				c.Flags.H,
				"Wrong value in flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<4) == 1<<4,
				c.Flags.C,
				"Wrong value in flag C. Expected %v, got %v",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestXORInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		From        string
	}{
		{c.Instructions[0xA8], "B"},
		{c.Instructions[0xA9], "C"},
		{c.Instructions[0xAA], "D"},
		{c.Instructions[0xAB], "E"},
		{c.Instructions[0xAC], "H"},
		{c.Instructions[0xAD], "L"},
		{c.Instructions[0xAF], "A"},
	}
	registers := map[string]*uint8{
		"B": &c.Registers.B,
		"C": &c.Registers.C,
		"D": &c.Registers.D,
		"E": &c.Registers.E,
		"H": &c.Registers.H,
		"L": &c.Registers.L,
		"A": &c.Registers.A,
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			c.ReadPC()
			value := uint8(c.Registers.GetA() ^ *registers[testCase.From])
			flagZ := value == 0
			flagN := false
			flagH := false
			flagC := false
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				flagZ,
				c.Flags.Z,
				"Wrong value in Flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagN,
				c.Flags.N,
				"Wrong value in Flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagH,
				c.Flags.H,
				"Wrong value in Flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagC,
				c.Flags.C,
				"Wrong value in Flag C. Expected %v, got %v",
			)
			tests.Equals(
				t,
				value,
				c.Registers.GetA(),
				"Wrong value in register A. Expected 0x%02X, got 0x%02X",
			)
			flags := c.Registers.GetF()
			// t.Logf("REG_F: 0x%08b\n", c.Registers.GetF())
			// t.Logf("FLAGS: 0x%08b\n", c.Flags.GetFlagsAsValue())
			tests.Equals(
				t,
				flags&(1<<7) == 1<<7,
				c.Flags.Z,
				"Wrong value in flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<6) == 1<<6,
				c.Flags.N,
				"Wrong value in flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<5) == 1<<5,
				c.Flags.H,
				"Wrong value in flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<4) == 1<<4,
				c.Flags.C,
				"Wrong value in flag C. Expected %v, got %v",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestXORInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xEE]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			immediateData := c.ReadMemory(pcBeforeExec + 1)
			c.ReadPC()
			value := uint8(c.Registers.GetA() ^ immediateData)
			flagZ := value == 0
			flagN := false
			flagH := false
			flagC := false
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				flagZ,
				c.Flags.Z,
				"Wrong value in Flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagN,
				c.Flags.N,
				"Wrong value in Flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagH,
				c.Flags.H,
				"Wrong value in Flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagC,
				c.Flags.C,
				"Wrong value in Flag C. Expected %v, got %v",
			)
			tests.Equals(
				t,
				value,
				c.Registers.GetA(),
				"Wrong value in register A. Expected 0x%02X, got 0x%02X",
			)
			flags := c.Registers.GetF()
			// t.Logf("REG_F: 0x%08b\n", c.Registers.GetF())
			// t.Logf("FLAGS: 0x%08b\n", c.Flags.GetFlagsAsValue())
			tests.Equals(
				t,
				flags&(1<<7) == 1<<7,
				c.Flags.Z,
				"Wrong value in flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<6) == 1<<6,
				c.Flags.N,
				"Wrong value in flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<5) == 1<<5,
				c.Flags.H,
				"Wrong value in flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<4) == 1<<4,
				c.Flags.C,
				"Wrong value in flag C. Expected %v, got %v",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestXORInstructions3(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xAE]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			immediateData := c.ReadMemory(c.Registers.GetHL())
			c.ReadPC()
			value := uint8(c.Registers.GetA() ^ immediateData)
			flagZ := value == 0
			flagN := false
			flagH := false
			flagC := false
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				flagZ,
				c.Flags.Z,
				"Wrong value in Flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagN,
				c.Flags.N,
				"Wrong value in Flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagH,
				c.Flags.H,
				"Wrong value in Flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagC,
				c.Flags.C,
				"Wrong value in Flag C. Expected %v, got %v",
			)
			tests.Equals(
				t,
				value,
				c.Registers.GetA(),
				"Wrong value in register A. Expected 0x%02X, got 0x%02X",
			)
			flags := c.Registers.GetF()
			// t.Logf("REG_F: 0x%08b\n", c.Registers.GetF())
			// t.Logf("FLAGS: 0x%08b\n", c.Flags.GetFlagsAsValue())
			tests.Equals(
				t,
				flags&(1<<7) == 1<<7,
				c.Flags.Z,
				"Wrong value in flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<6) == 1<<6,
				c.Flags.N,
				"Wrong value in flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<5) == 1<<5,
				c.Flags.H,
				"Wrong value in flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<4) == 1<<4,
				c.Flags.C,
				"Wrong value in flag C. Expected %v, got %v",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestORInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		From        string
	}{
		{c.Instructions[0xB0], "B"},
		{c.Instructions[0xB1], "C"},
		{c.Instructions[0xB2], "D"},
		{c.Instructions[0xB3], "E"},
		{c.Instructions[0xB4], "H"},
		{c.Instructions[0xB5], "L"},
		{c.Instructions[0xB7], "A"},
	}
	registers := map[string]*uint8{
		"B": &c.Registers.B,
		"C": &c.Registers.C,
		"D": &c.Registers.D,
		"E": &c.Registers.E,
		"H": &c.Registers.H,
		"L": &c.Registers.L,
		"A": &c.Registers.A,
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			c.ReadPC()
			value := uint8(c.Registers.GetA() & *registers[testCase.From])
			flagZ := value == 0
			flagN := false
			flagH := false
			flagC := false
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				flagZ,
				c.Flags.Z,
				"Wrong value in Flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagN,
				c.Flags.N,
				"Wrong value in Flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagH,
				c.Flags.H,
				"Wrong value in Flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagC,
				c.Flags.C,
				"Wrong value in Flag C. Expected %v, got %v",
			)
			tests.Equals(
				t,
				value,
				c.Registers.GetA(),
				"Wrong value in register A. Expected 0x%02X, got 0x%02X",
			)
			flags := c.Registers.GetF()
			// t.Logf("REG_F: 0x%08b\n", c.Registers.GetF())
			// t.Logf("FLAGS: 0x%08b\n", c.Flags.GetFlagsAsValue())
			tests.Equals(
				t,
				flags&(1<<7) == 1<<7,
				c.Flags.Z,
				"Wrong value in flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<6) == 1<<6,
				c.Flags.N,
				"Wrong value in flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<5) == 1<<5,
				c.Flags.H,
				"Wrong value in flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<4) == 1<<4,
				c.Flags.C,
				"Wrong value in flag C. Expected %v, got %v",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestORInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xF6]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			immediateData := c.ReadMemory(pcBeforeExec + 1)
			c.ReadPC()
			value := uint8(c.Registers.GetA() & immediateData)
			flagZ := value == 0
			flagN := false
			flagH := false
			flagC := false
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				flagZ,
				c.Flags.Z,
				"Wrong value in Flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagN,
				c.Flags.N,
				"Wrong value in Flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagH,
				c.Flags.H,
				"Wrong value in Flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagC,
				c.Flags.C,
				"Wrong value in Flag C. Expected %v, got %v",
			)
			tests.Equals(
				t,
				value,
				c.Registers.GetA(),
				"Wrong value in register A. Expected 0x%02X, got 0x%02X",
			)
			flags := c.Registers.GetF()
			// t.Logf("REG_F: 0x%08b\n", c.Registers.GetF())
			// t.Logf("FLAGS: 0x%08b\n", c.Flags.GetFlagsAsValue())
			tests.Equals(
				t,
				flags&(1<<7) == 1<<7,
				c.Flags.Z,
				"Wrong value in flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<6) == 1<<6,
				c.Flags.N,
				"Wrong value in flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<5) == 1<<5,
				c.Flags.H,
				"Wrong value in flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<4) == 1<<4,
				c.Flags.C,
				"Wrong value in flag C. Expected %v, got %v",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestORInstructions3(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xB6]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			immediateData := c.ReadMemory(c.Registers.GetHL())
			c.ReadPC()
			value := uint8(c.Registers.GetA() & immediateData)
			flagZ := value == 0
			flagN := false
			flagH := false
			flagC := false
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				flagZ,
				c.Flags.Z,
				"Wrong value in Flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagN,
				c.Flags.N,
				"Wrong value in Flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagH,
				c.Flags.H,
				"Wrong value in Flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagC,
				c.Flags.C,
				"Wrong value in Flag C. Expected %v, got %v",
			)
			tests.Equals(
				t,
				value,
				c.Registers.GetA(),
				"Wrong value in register A. Expected 0x%02X, got 0x%02X",
			)
			flags := c.Registers.GetF()
			// t.Logf("REG_F: 0x%08b\n", c.Registers.GetF())
			// t.Logf("FLAGS: 0x%08b\n", c.Flags.GetFlagsAsValue())
			tests.Equals(
				t,
				flags&(1<<7) == 1<<7,
				c.Flags.Z,
				"Wrong value in flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<6) == 1<<6,
				c.Flags.N,
				"Wrong value in flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<5) == 1<<5,
				c.Flags.H,
				"Wrong value in flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<4) == 1<<4,
				c.Flags.C,
				"Wrong value in flag C. Expected %v, got %v",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestCPInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		From        string
	}{
		{c.Instructions[0xB8], "B"},
		{c.Instructions[0xB9], "C"},
		{c.Instructions[0xBA], "D"},
		{c.Instructions[0xBB], "E"},
		{c.Instructions[0xBC], "H"},
		{c.Instructions[0xBD], "L"},
		{c.Instructions[0xBF], "A"},
	}
	registers := map[string]*uint8{
		"B": &c.Registers.B,
		"C": &c.Registers.C,
		"D": &c.Registers.D,
		"E": &c.Registers.E,
		"H": &c.Registers.H,
		"L": &c.Registers.L,
		"A": &c.Registers.A,
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			c.ReadPC()
			value1 := c.Registers.GetA()
			value2 := *registers[testCase.From]
			flagZ := value1 == value2
			flagN := true
			flagH := (value1 & 0xF) < (value2 & 0xF)
			flagC := value1 < value2
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				flagZ,
				c.Flags.Z,
				"Wrong value in Flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagN,
				c.Flags.N,
				"Wrong value in Flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagH,
				c.Flags.H,
				"Wrong value in Flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagC,
				c.Flags.C,
				"Wrong value in Flag C. Expected %v, got %v",
			)
			flags := c.Registers.GetF()
			// t.Logf("REG_F: 0x%08b\n", c.Registers.GetF())
			// t.Logf("FLAGS: 0x%08b\n", c.Flags.GetFlagsAsValue())
			tests.Equals(
				t,
				flags&(1<<7) == 1<<7,
				c.Flags.Z,
				"Wrong value in flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<6) == 1<<6,
				c.Flags.N,
				"Wrong value in flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<5) == 1<<5,
				c.Flags.H,
				"Wrong value in flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<4) == 1<<4,
				c.Flags.C,
				"Wrong value in flag C. Expected %v, got %v",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestCPInstructions2(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xD6]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			immediateData := c.ReadMemory(pcBeforeExec + 1)
			c.ReadPC()
			value := uint8(c.Registers.GetA() - immediateData)
			flagZ := value == 0
			flagN := true
			flagH := (c.Registers.GetA() & 0xF) < (immediateData & 0xF)
			flagC := uint16(c.Registers.GetA()) < uint16(immediateData)
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				flagZ,
				c.Flags.Z,
				"Wrong value in Flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagN,
				c.Flags.N,
				"Wrong value in Flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagH,
				c.Flags.H,
				"Wrong value in Flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagC,
				c.Flags.C,
				"Wrong value in Flag C. Expected %v, got %v",
			)
			tests.Equals(
				t,
				value,
				c.Registers.GetA(),
				"Wrong value in register A. Expected 0x%02X, got 0x%02X",
			)
			flags := c.Registers.GetF()
			// t.Logf("REG_F: 0x%08b\n", c.Registers.GetF())
			// t.Logf("FLAGS: 0x%08b\n", c.Flags.GetFlagsAsValue())
			tests.Equals(
				t,
				flags&(1<<7) == 1<<7,
				c.Flags.Z,
				"Wrong value in flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<6) == 1<<6,
				c.Flags.N,
				"Wrong value in flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<5) == 1<<5,
				c.Flags.H,
				"Wrong value in flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<4) == 1<<4,
				c.Flags.C,
				"Wrong value in flag C. Expected %v, got %v",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestCPInstructions3(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
	}{
		{c.Instructions[0xBE]},
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			c.ReadPC()
			value1 := c.Registers.GetA()
			value2 := c.ReadMemory(c.Registers.GetHL())
			flagZ := value1 == value2
			flagN := true
			flagH := (value1 & 0xF) < (value2 & 0xF)
			flagC := value1 < value2
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				flagZ,
				c.Flags.Z,
				"Wrong value in Flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagN,
				c.Flags.N,
				"Wrong value in Flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagH,
				c.Flags.H,
				"Wrong value in Flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagC,
				c.Flags.C,
				"Wrong value in Flag C. Expected %v, got %v",
			)
			flags := c.Registers.GetF()
			// t.Logf("REG_F: 0x%08b\n", c.Registers.GetF())
			// t.Logf("FLAGS: 0x%08b\n", c.Flags.GetFlagsAsValue())
			tests.Equals(
				t,
				flags&(1<<7) == 1<<7,
				c.Flags.Z,
				"Wrong value in flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<6) == 1<<6,
				c.Flags.N,
				"Wrong value in flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<5) == 1<<5,
				c.Flags.H,
				"Wrong value in flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<4) == 1<<4,
				c.Flags.C,
				"Wrong value in flag C. Expected %v, got %v",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestINCInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		From        string
	}{
		{c.Instructions[0x04], "B"},
		{c.Instructions[0x0C], "C"},
		{c.Instructions[0x14], "D"},
		{c.Instructions[0x1C], "E"},
		{c.Instructions[0x24], "H"},
		{c.Instructions[0x2C], "L"},
		{c.Instructions[0x3C], "A"},
	}
	registers := map[string]*uint8{
		"B": &c.Registers.B,
		"C": &c.Registers.C,
		"D": &c.Registers.D,
		"E": &c.Registers.E,
		"H": &c.Registers.H,
		"L": &c.Registers.L,
		"A": &c.Registers.A,
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			c.ReadPC()
			value1 := *registers[testCase.From]
			value2 := value1 + 1
			flagZ := value2 == 0
			flagN := false
			flagH := value2&0xF == 0
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				flagZ,
				c.Flags.Z,
				"Wrong value in Flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagN,
				c.Flags.N,
				"Wrong value in Flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagH,
				c.Flags.H,
				"Wrong value in Flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				value2,
				*registers[testCase.From],
				fmt.Sprintf("Wrong value in register %v. Expected 0x%%02X. Got 0x%%02X", testCase.From),
			)
			flags := c.Registers.GetF()
			// t.Logf("REG_F: 0x%08b\n", c.Registers.GetF())
			// t.Logf("FLAGS: 0x%08b\n", c.Flags.GetFlagsAsValue())
			tests.Equals(
				t,
				flags&(1<<7) == 1<<7,
				c.Flags.Z,
				"Wrong value in flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<6) == 1<<6,
				c.Flags.N,
				"Wrong value in flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<5) == 1<<5,
				c.Flags.H,
				"Wrong value in flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<4) == 1<<4,
				c.Flags.C,
				"Wrong value in flag C. Expected %v, got %v",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

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
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			c.ReadPC()
			addr := c.Registers.GetHL()
			value1 := c.ReadMemory(addr)
			value2 := value1 + 1
			flagZ := value2 == 0
			flagN := false
			flagH := value2&0xF == 0
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				flagZ,
				c.Flags.Z,
				"Wrong value in Flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagN,
				c.Flags.N,
				"Wrong value in Flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagH,
				c.Flags.H,
				"Wrong value in Flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				value2,
				c.ReadMemory(addr),
				fmt.Sprintf("Wrong value in address 0x%04X. Expected 0x%%02X. Got 0x%%02X", addr),
			)
			flags := c.Registers.GetF()
			// t.Logf("REG_F: 0x%08b\n", c.Registers.GetF())
			// t.Logf("FLAGS: 0x%08b\n", c.Flags.GetFlagsAsValue())
			tests.Equals(
				t,
				flags&(1<<7) == 1<<7,
				c.Flags.Z,
				"Wrong value in flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<6) == 1<<6,
				c.Flags.N,
				"Wrong value in flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<5) == 1<<5,
				c.Flags.H,
				"Wrong value in flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<4) == 1<<4,
				c.Flags.C,
				"Wrong value in flag C. Expected %v, got %v",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestINCInstructions3(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		From        string
	}{
		{c.Instructions[0x03], "BC"},
		{c.Instructions[0x13], "DE"},
		{c.Instructions[0x23], "HL"},
		{c.Instructions[0x33], "SP"},
	}
	registers := map[string]func() uint16{
		"BC": c.Registers.GetBC,
		"DE": c.Registers.GetDE,
		"HL": c.Registers.GetHL,
		"SP": c.Registers.GetSP,
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			c.ReadPC()
			value1 := registers[testCase.From]()
			value2 := value1 + 1
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				value2,
				registers[testCase.From](),
				fmt.Sprintf("Wrong value in register %v. Expected 0x%%04X. Got 0x%%04X", testCase.From),
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestDECInstructions1(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		From        string
	}{
		{c.Instructions[0x05], "B"},
		{c.Instructions[0x0D], "C"},
		{c.Instructions[0x15], "D"},
		{c.Instructions[0x1D], "E"},
		{c.Instructions[0x25], "H"},
		{c.Instructions[0x2D], "L"},
		{c.Instructions[0x3D], "A"},
	}
	registers := map[string]*uint8{
		"B": &c.Registers.B,
		"C": &c.Registers.C,
		"D": &c.Registers.D,
		"E": &c.Registers.E,
		"H": &c.Registers.H,
		"L": &c.Registers.L,
		"A": &c.Registers.A,
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			c.ReadPC()
			value1 := *registers[testCase.From]
			value2 := value1 - 1
			flagZ := value2 == 0
			flagN := true
			flagH := value1&0xF == 0
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				flagZ,
				c.Flags.Z,
				"Wrong value in Flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagN,
				c.Flags.N,
				"Wrong value in Flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagH,
				c.Flags.H,
				"Wrong value in Flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				value2,
				*registers[testCase.From],
				fmt.Sprintf("Wrong value in register %v. Expected 0x%%02X. Got 0x%%02X", testCase.From),
			)
			flags := c.Registers.GetF()
			// t.Logf("REG_F: 0x%08b\n", c.Registers.GetF())
			// t.Logf("FLAGS: 0x%08b\n", c.Flags.GetFlagsAsValue())
			tests.Equals(
				t,
				flags&(1<<7) == 1<<7,
				c.Flags.Z,
				"Wrong value in flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<6) == 1<<6,
				c.Flags.N,
				"Wrong value in flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<5) == 1<<5,
				c.Flags.H,
				"Wrong value in flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<4) == 1<<4,
				c.Flags.C,
				"Wrong value in flag C. Expected %v, got %v",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

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
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			c.ReadPC()
			addr := c.Registers.GetHL()
			value1 := c.ReadMemory(addr)
			value2 := value1 - 1
			flagZ := value2 == 0
			flagN := true
			flagH := value1&0xF == 0
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				flagZ,
				c.Flags.Z,
				"Wrong value in Flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagN,
				c.Flags.N,
				"Wrong value in Flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flagH,
				c.Flags.H,
				"Wrong value in Flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				value2,
				c.ReadMemory(addr),
				fmt.Sprintf("Wrong value in address 0x%04X. Expected 0x%%02X. Got 0x%%02X", addr),
			)
			flags := c.Registers.GetF()
			// t.Logf("REG_F: 0x%08b\n", c.Registers.GetF())
			// t.Logf("FLAGS: 0x%08b\n", c.Flags.GetFlagsAsValue())
			tests.Equals(
				t,
				flags&(1<<7) == 1<<7,
				c.Flags.Z,
				"Wrong value in flag Z. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<6) == 1<<6,
				c.Flags.N,
				"Wrong value in flag N. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<5) == 1<<5,
				c.Flags.H,
				"Wrong value in flag H. Expected %v, got %v",
			)
			tests.Equals(
				t,
				flags&(1<<4) == 1<<4,
				c.Flags.C,
				"Wrong value in flag C. Expected %v, got %v",
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}

func TestDECInstructions3(t *testing.T) {
	c := tests.InitCPU()
	instructions := []struct {
		Instruction cpu.Instruction
		From        string
	}{
		{c.Instructions[0x0B], "BC"},
		{c.Instructions[0x1B], "DE"},
		{c.Instructions[0x2B], "HL"},
		{c.Instructions[0x3B], "SP"},
	}
	registers := map[string]func() uint16{
		"BC": c.Registers.GetBC,
		"DE": c.Registers.GetDE,
		"HL": c.Registers.GetHL,
		"SP": c.Registers.GetSP,
	}

	for _, testCase := range instructions {
		testName := fmt.Sprintf("Executes %s", testCase.Instruction.Mnemonic)
		t.Run(testName, func(t *testing.T) {
			tests.RandRegisters(&c)
			pcBeforeExec := c.Registers.GetPC()
			pcAfterExec := pcBeforeExec + uint16(testCase.Instruction.Length)
			c.ReadPC()
			value1 := registers[testCase.From]()
			value2 := value1 - 1
			testCase.Instruction.Exec(&c)
			tests.Equals(
				t,
				value2,
				registers[testCase.From](),
				fmt.Sprintf("Wrong value in register %v. Expected 0x%%04X. Got 0x%%04X", testCase.From),
			)
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}
