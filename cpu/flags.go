package cpu

type Flags struct {
	Z bool
	N bool
	H bool
	C bool
}

func NewFlags() Flags {
	return Flags{}
}

func (f *Flags) GetFlagsAsValue() uint8 {
	flags := uint8(0)
	if f.Z {
		flags += 1 << 7
	}
	if f.N {
		flags += 1 << 6
	}
	if f.H {
		flags += 1 << 5
	}
	if f.C {
		flags += 1 << 4
	}
	return flags
}

func (f *Flags) GetCarryAsValue() uint8 {
	if f.C {
		return 1
	}
	return 0
}

// func (c *CPU) GetZ() bool {
// 	return c.Registers.GetF()%128 == 128
// }

// func (c *CPU) GetN() bool {
// 	return c.Registers.GetF()%64 == 64
// }

// func (c *CPU) GetH() bool {
// 	return c.Registers.GetF()%32 == 32
// }

// func (c *CPU) GetC() bool {
// 	return c.Registers.GetF()%16 == 16
// }

func (f *Flags) SetZ(v bool) {
	f.Z = v
}

func (f *Flags) SetN(v bool) {
	f.N = v
}

func (f *Flags) SetH(v bool) {
	f.H = v
}

func (f *Flags) SetC(v bool) {
	f.C = v
}

func (f *Flags) SetFlagsFromValue(v uint8) {
	f.SetZ(v&128 == 128)
	f.SetN(v&64 == 64)
	f.SetH(v&32 == 32)
	f.SetC(v&16 == 16)
}
