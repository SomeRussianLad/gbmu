package controllers

import (
	"math/rand"
	"testing"

	"gbmu/emulation/memory"
)

// TODO(somerussianlad) Complete this test with remaining registers
func TestLCDHandlers(t *testing.T) {
	memory := memory.NewDMGMemory()
	lcd := NewLCD(memory)

	testCases := []struct {
		value uint8
	}{}

	for i := 0; i < 10; i++ {
		testCases = append(testCases, struct{ value uint8 }{uint8(rand.Int())})
	}

	testName := "Checks whether exported getters work properly"
	t.Run(testName, func(t *testing.T) {
		for _, tc := range testCases {
			lcd.control = tc.value
			lcd.status = tc.value
			lcd.scy = tc.value
			lcd.scx = tc.value
			lcd.ly = tc.value
			lcd.lyc = tc.value
			lcd.wy = tc.value
			lcd.wx = tc.value
			lcd.bgp = tc.value
			lcd.obp0 = tc.value
			lcd.obp1 = tc.value

			expectedControl := tc.value
			expectedStatus := tc.value
			expectedSCY := tc.value
			expectedSCX := tc.value
			expectedLY := lcd.ly
			expectedLYC := tc.value
			expectedWY := tc.value
			expectedWX := tc.value
			expectedBGP := tc.value
			expectedOBP0 := tc.value
			expectedOBP1 := tc.value

			gotControl := memory.Read(ADDR_LCD_CONTROL)
			gotStatus := memory.Read(ADDR_LCD_STATUS)
			gotSCY := memory.Read(ADDR_LCD_SCY)
			gotSCX := memory.Read(ADDR_LCD_SCX)
			gotLY := memory.Read(ADDR_LCD_LY)
			gotLYC := memory.Read(ADDR_LCD_LYC)
			gotWY := memory.Read(ADDR_LCD_WY)
			gotWX := memory.Read(ADDR_LCD_WX)
			gotBGP := memory.Read(ADDR_LCD_BGP)
			gotOBP0 := memory.Read(ADDR_LCD_OBP0)
			gotOBP1 := memory.Read(ADDR_LCD_OBP1)

			if expectedControl != gotControl {
				t.Errorf("Control getter reads incorrect value. Expected 0x%02X, got 0x%02X", expectedControl, gotControl)
				break
			}
			if expectedStatus != gotStatus {
				t.Errorf("Status getter reads incorrect value. Expected 0x%02X, got 0x%02X", expectedStatus, gotStatus)
				break
			}
			if expectedSCY != gotSCY {
				t.Errorf("SCY getter reads incorrect value. Expected 0x%02X, got 0x%02X", expectedSCY, gotSCY)
				break
			}
			if expectedSCX != gotSCX {
				t.Errorf("SCX getter reads incorrect value. Expected 0x%02X, got 0x%02X", expectedSCX, gotSCX)
				break
			}
			if expectedLY != gotLY {
				t.Errorf("LY getter reads incorrect value. Expected 0x%02X, got 0x%02X", expectedLY, gotLY)
				break
			}
			if expectedLYC != gotLYC {
				t.Errorf("LYC getter reads incorrect value. Expected 0x%02X, got 0x%02X", expectedLYC, gotLYC)
				break
			}
			if expectedWY != gotWY {
				t.Errorf("WY getter reads incorrect value. Expected 0x%02X, got 0x%02X", expectedWY, gotWY)
				break
			}
			if expectedWX != gotWX {
				t.Errorf("WX getter reads incorrect value. Expected 0x%02X, got 0x%02X", expectedWX, gotWX)
				break
			}
			if expectedBGP != gotBGP {
				t.Errorf("BGP getter reads incorrect value. Expected 0x%02X, got 0x%02X", expectedBGP, gotBGP)
				break
			}
			if expectedOBP0 != gotOBP0 {
				t.Errorf("OBP0 getter reads incorrect value. Expected 0x%02X, got 0x%02X", expectedOBP0, gotOBP0)
				break
			}
			if expectedOBP1 != gotOBP1 {
				t.Errorf("OBP1 getter reads incorrect value. Expected 0x%02X, got 0x%02X", expectedOBP1, gotOBP1)
				break
			}
		}
	})

	testName = "Checks whether exported setters work properly"
	t.Run(testName, func(t *testing.T) {
		for _, tc := range testCases {
			memory.Write(ADDR_LCD_CONTROL, tc.value)
			memory.Write(ADDR_LCD_STATUS, tc.value)

			memory.Write(ADDR_LCD_SCY, tc.value)
			memory.Write(ADDR_LCD_SCX, tc.value)
			memory.Write(ADDR_LCD_LY, tc.value)
			memory.Write(ADDR_LCD_LYC, tc.value)
			memory.Write(ADDR_LCD_WY, tc.value)
			memory.Write(ADDR_LCD_WX, tc.value)

			memory.Write(ADDR_LCD_BGP, tc.value)
			memory.Write(ADDR_LCD_OBP0, tc.value)
			memory.Write(ADDR_LCD_OBP1, tc.value)

			expectedControl := tc.value
			expectedStatus := tc.value
			expectedSCY := tc.value
			expectedSCX := tc.value
			expectedLY := uint8(0)
			expectedLYC := tc.value
			expectedWY := tc.value
			expectedWX := tc.value
			expectedBGP := tc.value
			expectedOBP0 := tc.value
			expectedOBP1 := tc.value

			gotControl := lcd.control
			gotStatus := lcd.status
			gotSCY := lcd.scy
			gotSCX := lcd.scx
			gotLY := lcd.ly
			gotLYC := lcd.lyc
			gotWY := lcd.wy
			gotWX := lcd.wx
			gotBGP := lcd.bgp
			gotOBP0 := lcd.obp0
			gotOBP1 := lcd.obp1

			if expectedControl != gotControl {
				t.Errorf("Control setter writes incorrect value. Expected 0x%02X, got 0x%02X", expectedControl, gotControl)
				break
			}
			if expectedStatus != gotStatus {
				t.Errorf("Status setter writes incorrect value. Expected 0x%02X, got 0x%02X", expectedStatus, gotStatus)
				break
			}
			if expectedSCY != gotSCY {
				t.Errorf("SCY setter writes incorrect value. Expected 0x%02X, got 0x%02X", expectedSCY, gotSCY)
				break
			}
			if expectedSCX != gotSCX {
				t.Errorf("SCX setter writes incorrect value. Expected 0x%02X, got 0x%02X", expectedSCX, gotSCX)
				break
			}
			if expectedLY != gotLY {
				t.Errorf("LY setter writes incorrect value. Expected 0x%02X, got 0x%02X", expectedLY, gotLY)
				break
			}
			if expectedLYC != gotLYC {
				t.Errorf("LYC setter writes incorrect value. Expected 0x%02X, got 0x%02X", expectedLYC, gotLYC)
				break
			}
			if expectedWY != gotWY {
				t.Errorf("WY setter writes incorrect value. Expected 0x%02X, got 0x%02X", expectedWY, gotWY)
				break
			}
			if expectedWX != gotWX {
				t.Errorf("WX setter writes incorrect value. Expected 0x%02X, got 0x%02X", expectedWX, gotWX)
				break
			}
			if expectedBGP != gotBGP {
				t.Errorf("BGP setter writes incorrect value. Expected 0x%02X, got 0x%02X", expectedBGP, gotBGP)
				break
			}
			if expectedOBP0 != gotOBP0 {
				t.Errorf("OBP0 setter writes incorrect value. Expected 0x%02X, got 0x%02X", expectedOBP0, gotOBP0)
				break
			}
			if expectedOBP1 != gotOBP1 {
				t.Errorf("OBP1 setter writes incorrect value. Expected 0x%02X, got 0x%02X", expectedOBP1, gotOBP1)
				break
			}
		}
	})

	testName = "Checks whether exported getters/setters work properly"
	t.Run(testName, func(t *testing.T) {
		for _, tc := range testCases {
			memory.Write(ADDR_LCD_CONTROL, tc.value)
			memory.Write(ADDR_LCD_STATUS, tc.value)

			memory.Write(ADDR_LCD_SCY, tc.value)
			memory.Write(ADDR_LCD_SCX, tc.value)
			memory.Write(ADDR_LCD_LY, tc.value)
			memory.Write(ADDR_LCD_LYC, tc.value)
			memory.Write(ADDR_LCD_WY, tc.value)
			memory.Write(ADDR_LCD_WX, tc.value)

			memory.Write(ADDR_LCD_BGP, tc.value)
			memory.Write(ADDR_LCD_OBP0, tc.value)
			memory.Write(ADDR_LCD_OBP1, tc.value)

			expectedControl := tc.value
			expectedStatus := tc.value
			expectedSCY := tc.value
			expectedSCX := tc.value
			expectedLY := uint8(0)
			expectedLYC := tc.value
			expectedWY := tc.value
			expectedWX := tc.value
			expectedBGP := tc.value
			expectedOBP0 := tc.value
			expectedOBP1 := tc.value

			gotControl := memory.Read(ADDR_LCD_CONTROL)
			gotStatus := memory.Read(ADDR_LCD_STATUS)
			gotSCY := memory.Read(ADDR_LCD_SCY)
			gotSCX := memory.Read(ADDR_LCD_SCX)
			gotLY := memory.Read(ADDR_LCD_LY)
			gotLYC := memory.Read(ADDR_LCD_LYC)
			gotWY := memory.Read(ADDR_LCD_WY)
			gotWX := memory.Read(ADDR_LCD_WX)
			gotBGP := memory.Read(ADDR_LCD_BGP)
			gotOBP0 := memory.Read(ADDR_LCD_OBP0)
			gotOBP1 := memory.Read(ADDR_LCD_OBP1)

			if expectedControl != gotControl {
				t.Errorf("Control getter/setter reads/writes incorrect value. Expected 0x%02X, got 0x%02X", expectedControl, gotControl)
				break
			}
			if expectedStatus != gotStatus {
				t.Errorf("Status getter/setter reads/writes incorrect value. Expected 0x%02X, got 0x%02X", expectedStatus, gotStatus)
				break
			}
			if expectedSCY != gotSCY {
				t.Errorf("SCY getter/setter reads/writes incorrect value. Expected 0x%02X, got 0x%02X", expectedSCY, gotSCY)
				break
			}
			if expectedSCX != gotSCX {
				t.Errorf("SCX getter/setter reads/writes incorrect value. Expected 0x%02X, got 0x%02X", expectedSCX, gotSCX)
				break
			}
			if expectedLY != gotLY {
				t.Errorf("LY getter/setter reads/writes incorrect value. Expected 0x%02X, got 0x%02X", expectedLY, gotLY)
				break
			}
			if expectedLYC != gotLYC {
				t.Errorf("LYC getter/setter reads/writes incorrect value. Expected 0x%02X, got 0x%02X", expectedLYC, gotLYC)
				break
			}
			if expectedWY != gotWY {
				t.Errorf("WY getter/setter reads/writes incorrect value. Expected 0x%02X, got 0x%02X", expectedWY, gotWY)
				break
			}
			if expectedWX != gotWX {
				t.Errorf("WX getter/setter reads/writes incorrect value. Expected 0x%02X, got 0x%02X", expectedWX, gotWX)
				break
			}
			if expectedBGP != gotBGP {
				t.Errorf("BGP getter/setter reads/writes incorrect value. Expected 0x%02X, got 0x%02X", expectedBGP, gotBGP)
				break
			}
			if expectedOBP0 != gotOBP0 {
				t.Errorf("OBP0 getter/setter reads/writes incorrect value. Expected 0x%02X, got 0x%02X", expectedOBP0, gotOBP0)
				break
			}
			if expectedOBP1 != gotOBP1 {
				t.Errorf("OBP1 getter/setter reads/writes incorrect value. Expected 0x%02X, got 0x%02X", expectedOBP1, gotOBP1)
				break
			}
		}
	})
}
