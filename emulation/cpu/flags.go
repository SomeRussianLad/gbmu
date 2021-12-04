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
func (f *flags) setValue(v uint8) {
	f.value = v & 0xF0
}

// getCarry returns 1 if the Carry flag is set, 0 otherwise
func (f *flags) getCarry() uint8 {
	return (f.value >> 4) % 2
}

// getZ returns true if the Zero flag is set, false otherwise
func (f *flags) getZ() bool {
	return (f.value & 128) > 0
}

// getZ returns true if the Subtract flag is set, false otherwise
func (f *flags) getN() bool {
	return (f.value & 64) > 0
}

// getZ returns true if the Half Carry flag is set, false otherwise
func (f *flags) getH() bool {
	return (f.value & 32) > 0
}

// getZ returns true if the Carry flag is set, false otherwise
func (f *flags) getC() bool {
	return (f.value & 16) > 0
}

// setZ sets the value of the Zero flag
func (f *flags) setZ(v bool) {
	if v {
		f.value |= 128
	} else {
		f.value &= 127
	}
}

// setN sets the value of the Subtract flag
func (f *flags) setN(v bool) {
	if v {
		f.value |= 64
	} else {
		f.value &= 191
	}
}

// setH sets the value of the Half Carry flag
func (f *flags) setH(v bool) {
	if v {
		f.value |= 32
	} else {
		f.value &= 223
	}
}

// setC sets the value of the Carry flag
func (f *flags) setC(v bool) {
	if v {
		f.value |= 16
	} else {
		f.value &= 239
	}
}
