package controllers

import "gbmu/emulation/memory"

const ADDR_DMA_TRANSFER = uint16(0xFF46)

type DMA struct {
	transfer   uint8
	isLaunched bool   // isLaunched is set to true if DMA transfer is in progress, false otherwise
	ptr        uint16 // ptr stores the least significant byte of src and dst addresses of the current transfer iteration

	memory memory.Memory
}

func NewDMA(memory memory.Memory) *DMA {
	dma := &DMA{memory: memory}

	handlers := []struct {
		addr   uint16
		getter func() uint8
		setter func(uint8)
	}{
		{ADDR_DMA_TRANSFER, dma.dmaGetter(), dma.dmaSetter()},
	}

	for _, h := range handlers {
		memory.RegisterGetter(h.addr, h.getter)
		memory.RegisterSetter(h.addr, h.setter)
	}

	return dma
}

func (d *DMA) Update(cycles int) {
	for ; d.isLaunched && cycles >= 4; cycles -= 4 {
		srcAddr := uint16(d.transfer)<<8 + d.ptr
		dstAddr := uint16(0xFE00) + d.ptr

		value := d.memory.Read(srcAddr)
		d.memory.Write(dstAddr, value) // TODO(somerussianlad) Come up with a fix! Write will be ignored in LCD Mode 2-3

		d.ptr++
		if d.ptr == 160 {
			d.isLaunched = false
			d.ptr = 0
		}
	}
}

func (d *DMA) dmaGetter() func() uint8 {
	return func() uint8 {
		return d.transfer
	}
}

func (d *DMA) dmaSetter() func(uint8) {
	return func(value uint8) {
		// TODO(somerussianlad) Is it us who must enforce the 0x00-0xF1 range? Do gamedevs respect it?
		// d.transfer = value % 0xF2
		d.transfer = value
		d.isLaunched = true
	}
}
