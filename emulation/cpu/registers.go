package cpu

type registers struct {
	a  uint8
	f  *flags
	b  uint8
	c  uint8
	d  uint8
	e  uint8
	h  uint8
	l  uint8
	sp uint16
	pc uint16
}

func newRegisters() *registers {
	flags := newFlags()
	return &registers{f: flags}
}

func (r *registers) getA() uint8 {
	return r.a
}

func (r *registers) getF() uint8 {
	return r.f.getValue()
}

func (r *registers) getB() uint8 {
	return r.b
}

func (r *registers) getC() uint8 {
	return r.c
}

func (r *registers) getD() uint8 {
	return r.d
}

func (r *registers) getE() uint8 {
	return r.e
}

func (r *registers) getH() uint8 {
	return r.h
}

func (r *registers) getL() uint8 {
	return r.l
}

func (r *registers) getBC() uint16 {
	msb := uint16(r.getB())
	lsb := uint16(r.getC())
	value := msb<<8 | lsb
	return value
}

func (r *registers) getDE() uint16 {
	msb := uint16(r.getD())
	lsb := uint16(r.getE())
	value := msb<<8 | lsb
	return value
}

func (r *registers) getHL() uint16 {
	msb := uint16(r.getH())
	lsb := uint16(r.getL())
	value := msb<<8 | lsb
	return value
}

func (r *registers) getSP() uint16 {
	return r.sp
}

func (r *registers) getPC() uint16 {
	return r.pc
}

func (r *registers) setA(value uint8) {
	r.a = value
}

func (r *registers) setF(value uint8) {
	r.f.setValue(value)
}

func (r *registers) setB(value uint8) {
	r.b = value
}

func (r *registers) setC(value uint8) {
	r.c = value
}

func (r *registers) setD(value uint8) {
	r.d = value
}

func (r *registers) setE(value uint8) {
	r.e = value
}

func (r *registers) setH(value uint8) {
	r.h = value
}

func (r *registers) setL(value uint8) {
	r.l = value
}

func (r *registers) setAF(value uint16) {
	msb := uint8(value >> 8)
	lsb := uint8(value & 0xFF)
	r.setA(msb)
	r.setF(lsb)
}

func (r *registers) setBC(value uint16) {
	msb := uint8(value >> 8)
	lsb := uint8(value & 0xFF)
	r.setB(msb)
	r.setC(lsb)
}

func (r *registers) setDE(value uint16) {
	msb := uint8(value >> 8)
	lsb := uint8(value & 0xFF)
	r.setD(msb)
	r.setE(lsb)
}

func (r *registers) setHL(value uint16) {
	msb := uint8(value >> 8)
	lsb := uint8(value & 0xFF)
	r.setH(msb)
	r.setL(lsb)
}

func (r *registers) setSP(value uint16) {
	r.sp = value
}

func (r *registers) setPC(value uint16) {
	r.pc = value
}

func (r *registers) incA() {
	r.a++
}

func (r *registers) incB() {
	r.b++
}

func (r *registers) incC() {
	r.c++
}

func (r *registers) incD() {
	r.d++
}

func (r *registers) incE() {
	r.e++
}

func (r *registers) incH() {
	r.h++
}

func (r *registers) incL() {
	r.l++
}

func (r *registers) incBC() {
	value := r.getBC()
	r.setBC(value + 1)
}

func (r *registers) incDE() {
	value := r.getDE()
	r.setDE(value + 1)
}

func (r *registers) incHL() {
	value := r.getHL()
	r.setHL(value + 1)
}

func (r *registers) incSP() {
	r.sp++
}

func (r *registers) incPC() {
	r.pc++
}

func (r *registers) decA() {
	r.a--
}

func (r *registers) decB() {
	r.b--
}

func (r *registers) decC() {
	r.c--
}

func (r *registers) decD() {
	r.d--
}

func (r *registers) decE() {
	r.e--
}

func (r *registers) decH() {
	r.h--
}

func (r *registers) decL() {
	r.l--
}

func (r *registers) decBC() {
	value := r.getBC()
	r.setBC(value - 1)
}

func (r *registers) decDE() {
	value := r.getDE()
	r.setDE(value - 1)
}

func (r *registers) decHL() {
	value := r.getHL()
	r.setHL(value - 1)
}

func (r *registers) decSP() {
	r.sp--
}
