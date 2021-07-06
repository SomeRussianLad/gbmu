package cpu

type Registers struct {
	A  uint8
	F  uint8
	B  uint8
	C  uint8
	D  uint8
	E  uint8
	H  uint8
	L  uint8
	SP uint16
	PC uint16
}

func NewRegisters() Registers {
	return Registers{}
}

func (r *Registers) GetA() uint8 {
	return r.A
}

func (r *Registers) GetF() uint8 {
	// return r.F
	return r.F & 0xF0
}

func (r *Registers) GetB() uint8 {
	return r.B
}

func (r *Registers) GetC() uint8 {
	return r.C
}

func (r *Registers) GetD() uint8 {
	return r.D
}

func (r *Registers) GetE() uint8 {
	return r.E
}

func (r *Registers) GetH() uint8 {
	return r.H
}

func (r *Registers) GetL() uint8 {
	return r.L
}

func (r *Registers) GetS() uint8 {
	return uint8(r.GetSP() >> 8)
}

func (r *Registers) GetP() uint8 {
	return uint8(r.GetSP() & 0xFF)
}

func (r *Registers) GetAF() uint16 {
	msb := uint16(r.GetA())
	// lsb := uint16(r.GetF())
	lsb := uint16(r.GetF() & 0xF0)
	value := msb<<8 | lsb
	return value
}

func (r *Registers) GetBC() uint16 {
	msb := uint16(r.GetB())
	lsb := uint16(r.GetC())
	value := msb<<8 | lsb
	return value
}

func (r *Registers) GetDE() uint16 {
	msb := uint16(r.GetD())
	lsb := uint16(r.GetE())
	value := msb<<8 | lsb
	return value
}

func (r *Registers) GetHL() uint16 {
	msb := uint16(r.GetH())
	lsb := uint16(r.GetL())
	value := msb<<8 | lsb
	return value
}

func (r *Registers) GetSP() uint16 {
	return r.SP
}

func (r *Registers) GetPC() uint16 {
	return r.PC
}

func (r *Registers) SetA(v uint8) {
	r.A = v
}

func (r *Registers) SetF(v uint8) {
	// r.F = v
	r.F = v & 0xF0
}

func (r *Registers) SetB(v uint8) {
	r.B = v
}

func (r *Registers) SetC(v uint8) {
	r.C = v
}

func (r *Registers) SetD(v uint8) {
	r.D = v
}

func (r *Registers) SetE(v uint8) {
	r.E = v
}

func (r *Registers) SetH(v uint8) {
	r.H = v
}

func (r *Registers) SetL(v uint8) {
	r.L = v
}

func (r *Registers) SetAF(v uint16) {
	msb := uint8(v >> 8)
	lsb := uint8(v & 0xFF)
	r.SetA(msb)
	// r.SetF(lsb)
	r.SetF(lsb & 0xF0)
}

func (r *Registers) SetBC(v uint16) {
	msb := uint8(v >> 8)
	lsb := uint8(v & 0xFF)
	r.SetB(msb)
	r.SetC(lsb)
}

func (r *Registers) SetDE(v uint16) {
	msb := uint8(v >> 8)
	lsb := uint8(v & 0xFF)
	r.SetD(msb)
	r.SetE(lsb)
}

func (r *Registers) SetHL(v uint16) {
	msb := uint8(v >> 8)
	lsb := uint8(v & 0xFF)
	r.SetH(msb)
	r.SetL(lsb)
}

func (r *Registers) SetSP(v uint16) {
	r.SP = v
}

func (r *Registers) SetPC(v uint16) {
	r.PC = v
}

func (r *Registers) IncA() {
	r.A++
}

func (r *Registers) IncB() {
	r.B++
}

func (r *Registers) IncC() {
	r.C++

}

func (r *Registers) IncD() {
	r.D++

}

func (r *Registers) IncE() {
	r.E++

}

func (r *Registers) IncH() {
	r.H++

}

func (r *Registers) IncL() {
	r.L++
}

func (r *Registers) IncBC() {
	value := r.GetBC()
	r.SetBC(value + 1)
}

func (r *Registers) IncDE() {
	value := r.GetDE()
	r.SetDE(value + 1)
}

func (r *Registers) IncHL() {
	value := r.GetHL()
	r.SetHL(value + 1)
}

func (r *Registers) IncSP() {
	r.SP++
}

func (r *Registers) IncPC() {
	r.PC++
}

func (r *Registers) DecA() {
	r.A--
}

func (r *Registers) DecB() {
	r.B--
}

func (r *Registers) DecC() {
	r.C--

}

func (r *Registers) DecD() {
	r.D--

}

func (r *Registers) DecE() {
	r.E--

}

func (r *Registers) DecH() {
	r.H--

}

func (r *Registers) DecL() {
	r.L--
}

func (r *Registers) DecBC() {
	value := r.GetBC()
	r.SetBC(value - 1)
}

func (r *Registers) DecDE() {
	value := r.GetDE()
	r.SetDE(value - 1)
}

func (r *Registers) DecHL() {
	value := r.GetHL()
	r.SetHL(value - 1)
}

func (r *Registers) DecSP() {
	r.SP--
}
