package cpu

// Flags emulate the behavior of CPU flags and are required for arithmetic
// operations to work properly.
type flags struct {
	value uint8
}

func newFlags() *flags {
	return &flags{}
}

// getValue returns the value of the F register.
func (f *flags) getValue() uint8 {
	return f.value & 0xF0
}

// setValue sets the value of the F register.
// Four least significant (rightmost) bits are always cleared
func (f *flags) setValue(value uint8) {
	f.value = value & 0xF0
}

// getCarry returns 1 if the Carry flag is set, 0 otherwise
func (f *flags) getCarry() uint8 {
	return (f.value >> 4) & 1
}

// getZ returns true if the Zero flag is set, false otherwise
func (f *flags) getZ() bool {
	return (f.value>>7)&1 == 1
}

// getZ returns true if the Subtract flag is set, false otherwise
func (f *flags) getN() bool {
	return (f.value>>6)&1 == 1
}

// getZ returns true if the Half Carry flag is set, false otherwise
func (f *flags) getH() bool {
	return (f.value>>5)&1 == 1
}

// getZ returns true if the Carry flag is set, false otherwise
func (f *flags) getC() bool {
	return (f.value>>4)&1 == 1
}

// setZ sets the value of the Zero flag
func (f *flags) setZ(value bool) {
	if value {
		f.value |= 128
	} else {
		f.value &= 127
	}
}

// setN sets the value of the Subtract flag
func (f *flags) setN(value bool) {
	if value {
		f.value |= 64
	} else {
		f.value &= 191
	}
}

// setH sets the value of the Half Carry flag
func (f *flags) setH(value bool) {
	if value {
		f.value |= 32
	} else {
		f.value &= 223
	}
}

// setC sets the value of the Carry flag
func (f *flags) setC(value bool) {
	if value {
		f.value |= 16
	} else {
		f.value &= 239
	}
}
