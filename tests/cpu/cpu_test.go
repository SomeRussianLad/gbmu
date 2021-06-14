package cpu_test

import (
	"fmt"
	"goboy/tests"
	"testing"
)

func TestReadPC(t *testing.T) {
	c := tests.InitCPU()
	for i := 0; i < 10; i++ {
		tests.RandRegisters(&c)
		pcBeforeExec := c.Registers.GetPC()
		testName := fmt.Sprintf("Reads value 0x%02X in memory address 0x%04X and increments PC by 1", c.Memory.ReadMemory(pcBeforeExec), pcBeforeExec)
		t.Run(testName, func(t *testing.T) {
			value := c.ReadPC()
			pcAfterExec := c.Registers.GetPC()
			tests.Equals(t, value, c.Memory.ReadMemory(pcBeforeExec), "Wrong value read by PC. Expected 0x%02X, got 0x%02X")
			tests.Equals(t, pcAfterExec, c.Registers.GetPC(), "Wrong value in register PC. Expected 0x%04X, got 0x%04X")
		})
	}
}
