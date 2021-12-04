package controllers

import (
	"math/rand"
	"testing"
	"time"

	"gbmu/emulation/memory"
)

func randomizeDMASourceMemory(memory memory.Memory) {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0x0000; i < 0xF1A0; i++ {
		memory.Write(uint16(i), uint8(rand.Int()))
	}
}

func TestDMAUpdate(t *testing.T) {
	memory := memory.NewDMGMemory()
	dma := NewDMA(memory)

	testCases := []struct {
		cycles int
	}{}

	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < (1 << 8); i++ {
		testCases = append(testCases, struct{ cycles int }{(rand.Intn(5) + 1) * 4})
	}

	t.Run("Checks whether update works properly when DMA is NOT launched", func(t *testing.T) {
		var cyclesSum int

		expectedOAM := func() []uint8 {
			oam := make([]uint8, 160)
			for src, i := uint16(0xFE00), uint16(0); i < 0xA0; i++ {
				oam[i] = memory.Read(src + i)
			}
			return oam
		}()

		randomizeDMASourceMemory(memory)

		for _, tc := range testCases {
			dma.Update(tc.cycles)
			cyclesSum += tc.cycles

			expectedIsLaunched := false
			expectedPtr := uint16(0)

			gotIsLaunched := dma.isLaunched
			gotPtr := dma.ptr

			if expectedIsLaunched != gotIsLaunched {
				t.Errorf("DMA is launched. Expected %v, got %v", expectedIsLaunched, gotIsLaunched)
				break
			}
			if expectedPtr != gotPtr {
				t.Errorf("DMA pointer is offset. Expected 0x%04X, got 0x%04X", expectedPtr, gotPtr)
				break
			}

			gotOAM := func() []uint8 {
				oam := make([]uint8, 160)
				for src, i := uint16(0xFE00), uint16(0); i < 0xA0; i++ {
					oam[i] = memory.Read(src + i)
				}
				return oam
			}()

			for i := range expectedOAM {
				if expectedOAM[i] != gotOAM[i] {
					t.Errorf("DMA transfer not authorized. SRC and OAM do not contain the same data")
					t.Logf("Machine cycles: %v", cyclesSum)
					t.Logf("EXP_OAM: %v", expectedOAM)
					t.Logf("GOT_OAM: %v", gotOAM)
					break
				}
			}
		}
	})

	t.Run("Checks whether update works properly when DMA is launched", func(t *testing.T) {
		var cyclesSum int

		transfer := uint8(rand.Int() % 0xF2)
		memory.Write(ADDR_DMA_TRANSFER, transfer)
		randomizeDMASourceMemory(memory)

		for _, tc := range testCases {
			dma.Update(tc.cycles)
			cyclesSum += tc.cycles

			expectedIsLaunched := func() bool {
				return cyclesSum < 160*4
			}()
			expectedPtr := func() uint16 {
				if cyclesSum >= 160*4 {
					return 0
				}
				return uint16(cyclesSum / 4)
			}()

			gotIsLaunched := dma.isLaunched
			gotPtr := dma.ptr

			if expectedIsLaunched != gotIsLaunched {
				t.Errorf("DMA is launched. Expected %v, got %v", expectedIsLaunched, gotIsLaunched)
				break
			}
			if expectedPtr != gotPtr {
				t.Errorf("DMA pointer is offset. Expected 0x%04X, got 0x%04X", expectedPtr, gotPtr)
				break
			}

			expectedOAM := func() []uint8 {
				oam := make([]uint8, 160)
				if cyclesSum < 160*4 {
					for src, i := uint16(transfer)<<8, uint16(0); i < expectedPtr; i++ {
						oam[i] = memory.Read(src + i)
					}
				} else {
					for src, i := uint16(transfer)<<8, uint16(0); i < 160; i++ {
						oam[i] = memory.Read(src + i)
					}
				}
				return oam
			}()
			gotOAM := func() []uint8 {
				oam := make([]uint8, 160)
				for src, i := uint16(0xFE00), uint16(0); i < 160; i++ {
					oam[i] = memory.Read(src + i)
				}
				return oam
			}()

			for i := range expectedOAM {
				if expectedOAM[i] != gotOAM[i] {
					t.Errorf("SRC and OAM do not contain the same data")
					t.Logf("Machine cycles: %v", cyclesSum)
					t.Logf("EXP_OAM: %v", expectedOAM)
					t.Logf("GOT_OAM: %v", gotOAM)
					break
				}
			}
		}
	})
}

func TestDMAHandlers(t *testing.T) {
	memory := memory.NewDMGMemory()
	dma := NewDMA(memory)

	testCases := []struct {
		value uint8
	}{}

	for i := 0; i < 10; i++ {
		testCases = append(testCases, struct{ value uint8 }{uint8(rand.Int())})
	}

	testName := "Checks whether exported getter works properly"
	t.Run(testName, func(t *testing.T) {
		for _, tc := range testCases {
			dma.transfer = tc.value
			dma.isLaunched = true

			expectedDma := tc.value
			expectedDmaLaunched := true

			gotDma := memory.Read(ADDR_DMA_TRANSFER)
			gotDmaLaunched := dma.isLaunched

			if expectedDma != gotDma {
				t.Errorf("Transfer getter reads incorrect value. Expected 0x%02X, got 0x%02X", expectedDma, gotDma)
				break
			}
			if expectedDmaLaunched != gotDmaLaunched {
				t.Errorf("Wrong isLaunched value. DMA transfer not started. Expected %v, got %v", expectedDmaLaunched, gotDmaLaunched)
				break
			}
		}
	})

	testName = "Checks whether exported setter works properly"
	t.Run(testName, func(t *testing.T) {
		for _, tc := range testCases {
			memory.Write(ADDR_DMA_TRANSFER, tc.value)

			expectedDma := tc.value
			expectedDmaLaunched := true

			gotDma := dma.transfer
			gotDmaLaunched := dma.isLaunched

			if expectedDma != gotDma {
				t.Errorf("Transfer setter writes incorrect value. Expected 0x%02X, got 0x%02X", expectedDma, gotDma)
				break
			}
			if expectedDmaLaunched != gotDmaLaunched {
				t.Errorf("Wrong isLaunched value. DMA transfer not started. Expected %v, got %v", expectedDmaLaunched, gotDmaLaunched)
				break
			}
		}
	})

	testName = "Checks whether exported getter/setter works properly"
	t.Run(testName, func(t *testing.T) {
		for _, tc := range testCases {
			memory.Write(ADDR_DMA_TRANSFER, tc.value)

			expectedDma := tc.value
			expectedDmaLaunched := true

			gotDma := memory.Read(ADDR_DMA_TRANSFER)
			gotDmaLaunched := dma.isLaunched

			if expectedDma != gotDma {
				t.Errorf("Transfer getter/setter reads/writes incorrect value. Expected 0x%02X, got 0x%02X", expectedDma, gotDma)
				break
			}
			if expectedDmaLaunched != gotDmaLaunched {
				t.Errorf("Wrong isLaunched value. DMA transfer not started. Expected %v, got %v", expectedDmaLaunched, gotDmaLaunched)
				break
			}
		}
	})
}
