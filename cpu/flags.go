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

func (f *Flags) GetZ() bool {
	return f.Z
}

func (f *Flags) GetN() bool {
	return f.N
}

func (f *Flags) GetH() bool {
	return f.H
}

func (f *Flags) GetC() bool {
	return f.C
}

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
