package cpu

type Instruction struct {
	Mnemonic string
	Length   int
	Cycles   int
	Exec     func(*CPU)
}

type Instructions map[int]Instruction

func NewInstructions() Instructions {
	return Instructions{
		0x00: {"NOP", 1, 4, (*CPU).Instruction00},
		0x01: {"LD BC,nn", 3, 12, (*CPU).Instruction01},
		0x02: {"LD (BC),A", 1, 8, (*CPU).Instruction02},
		0x03: {"INC BC", 1, 8, (*CPU).Instruction03},
		0x04: {"INC B", 1, 4, (*CPU).Instruction04},
		0x05: {"DEC B", 1, 4, (*CPU).Instruction05},
		0x06: {"LD B,n", 2, 8, (*CPU).Instruction06},
		0x07: {"RLCA", 1, 4, (*CPU).Instruction07},
		0x08: {"LD (nn),SP", 3, 20, (*CPU).Instruction08},
		0x09: {"ADD HL,BC", 1, 8, (*CPU).Instruction09},
		0x0A: {"LD A,(BC)", 1, 8, (*CPU).Instruction0A},
		0x0B: {"DEC BC", 1, 8, (*CPU).Instruction0B},
		0x0C: {"INC C", 1, 4, (*CPU).Instruction0C},
		0x0D: {"DEC C", 1, 4, (*CPU).Instruction0D},
		0x0E: {"LD C,n", 2, 8, (*CPU).Instruction0E},
		0x0F: {"RRCA", 1, 4, (*CPU).Instruction0F},

		0x10: {"STOP", 2, 4, (*CPU).Instruction10},
		0x11: {"LD DE,nn", 3, 12, (*CPU).Instruction11},
		0x12: {"LD (DE),A", 1, 8, (*CPU).Instruction12},
		0x13: {"INC DE", 1, 8, (*CPU).Instruction13},
		0x14: {"INC D", 1, 4, (*CPU).Instruction14},
		0x15: {"DEC D", 1, 4, (*CPU).Instruction15},
		0x16: {"LD D,n", 2, 8, (*CPU).Instruction16},
		0x17: {"RLA", 1, 4, (*CPU).Instruction17},
		0x18: {"JR n", 2, 12, (*CPU).Instruction18},
		0x19: {"ADD HL,DE", 1, 8, (*CPU).Instruction19},
		0x1A: {"LD A,(DE)", 1, 8, (*CPU).Instruction1A},
		0x1B: {"DEC DE", 1, 8, (*CPU).Instruction1B},
		0x1C: {"INC E", 1, 4, (*CPU).Instruction1C},
		0x1D: {"DEC E", 1, 4, (*CPU).Instruction1D},
		0x1E: {"LD E,n", 2, 8, (*CPU).Instruction1E},
		0x1F: {"RRA", 1, 4, (*CPU).Instruction1F},

		0x20: {"JR NZ,n", 2, 8, (*CPU).Instruction20},
		0x21: {"LD HL,nn", 3, 12, (*CPU).Instruction21},
		0x22: {"LDI (HL),A", 1, 8, (*CPU).Instruction22},
		0x23: {"INC HL", 1, 8, (*CPU).Instruction23},
		0x24: {"INC H", 1, 4, (*CPU).Instruction24},
		0x25: {"DEC H", 1, 4, (*CPU).Instruction25},
		0x26: {"LD H,n", 2, 8, (*CPU).Instruction26},
		0x27: {"DAA", 1, 4, (*CPU).Instruction27},
		0x28: {"JR Z,n", 2, 8, (*CPU).Instruction28},
		0x29: {"ADD HL,HL", 1, 8, (*CPU).Instruction29},
		0x2A: {"LDI A,(HL)", 1, 8, (*CPU).Instruction2A},
		0x2B: {"DEC HL", 1, 8, (*CPU).Instruction2B},
		0x2C: {"INC L", 1, 4, (*CPU).Instruction2C},
		0x2D: {"DEC L", 1, 4, (*CPU).Instruction2D},
		0x2E: {"LD L,n", 2, 8, (*CPU).Instruction2E},
		0x2F: {"CPL", 1, 4, (*CPU).Instruction2F},

		0x30: {"JR NC,n", 2, 8, (*CPU).Instruction30},
		0x31: {"LD SP,nn", 3, 12, (*CPU).Instruction31},
		0x32: {"LDD (HL),A", 1, 8, (*CPU).Instruction32},
		0x33: {"INC SP", 1, 8, (*CPU).Instruction33},
		0x34: {"INC (HL)", 1, 12, (*CPU).Instruction34},
		0x35: {"DEC (HL)", 1, 12, (*CPU).Instruction35},
		0x36: {"LD (HL),n", 2, 12, (*CPU).Instruction36},
		0x37: {"SCF", 1, 4, (*CPU).Instruction37},
		0x38: {"JR C,n", 2, 8, (*CPU).Instruction38},
		0x39: {"ADD HL,SP", 1, 8, (*CPU).Instruction39},
		0x3A: {"LDD A,(HL)", 1, 8, (*CPU).Instruction3A},
		0x3B: {"DEC SP", 1, 8, (*CPU).Instruction3B},
		0x3C: {"INC A", 1, 4, (*CPU).Instruction3C},
		0x3D: {"DEC A", 1, 4, (*CPU).Instruction3D},
		0x3E: {"LD A,n", 2, 8, (*CPU).Instruction3E},
		0x3F: {"CCF", 1, 4, (*CPU).Instruction3F},

		0x40: {"LD B,B", 1, 4, (*CPU).Instruction40},
		0x41: {"LD B,C", 1, 4, (*CPU).Instruction41},
		0x42: {"LD B,D", 1, 4, (*CPU).Instruction42},
		0x43: {"LD B,E", 1, 4, (*CPU).Instruction43},
		0x44: {"LD B,H", 1, 4, (*CPU).Instruction44},
		0x45: {"LD B,L", 1, 4, (*CPU).Instruction45},
		0x46: {"LD B,(HL)", 1, 8, (*CPU).Instruction46},
		0x47: {"LD B,A", 1, 4, (*CPU).Instruction47},
		0x48: {"LD C,B", 1, 4, (*CPU).Instruction48},
		0x49: {"LD C,C", 1, 4, (*CPU).Instruction49},
		0x4A: {"LD C,D", 1, 4, (*CPU).Instruction4A},
		0x4B: {"LD C,E", 1, 4, (*CPU).Instruction4B},
		0x4C: {"LD C,H", 1, 4, (*CPU).Instruction4C},
		0x4D: {"LD C,L", 1, 4, (*CPU).Instruction4D},
		0x4E: {"LD C,(HL)", 1, 8, (*CPU).Instruction4E},
		0x4F: {"LD C,A", 1, 4, (*CPU).Instruction4F},

		0x50: {"LD D,B", 1, 4, (*CPU).Instruction50},
		0x51: {"LD D,C", 1, 4, (*CPU).Instruction51},
		0x52: {"LD D,D", 1, 4, (*CPU).Instruction52},
		0x53: {"LD D,E", 1, 4, (*CPU).Instruction53},
		0x54: {"LD D,H", 1, 4, (*CPU).Instruction54},
		0x55: {"LD D,L", 1, 4, (*CPU).Instruction55},
		0x56: {"LD D,(HL)", 1, 8, (*CPU).Instruction56},
		0x57: {"LD D,A", 1, 4, (*CPU).Instruction57},
		0x58: {"LD E,B", 1, 4, (*CPU).Instruction58},
		0x59: {"LD E,C", 1, 4, (*CPU).Instruction59},
		0x5A: {"LD E,D", 1, 4, (*CPU).Instruction5A},
		0x5B: {"LD E,E", 1, 4, (*CPU).Instruction5B},
		0x5C: {"LD E,H", 1, 4, (*CPU).Instruction5C},
		0x5D: {"LD E,L", 1, 4, (*CPU).Instruction5D},
		0x5E: {"LD E,(HL)", 1, 8, (*CPU).Instruction5E},
		0x5F: {"LD E,A", 1, 4, (*CPU).Instruction5F},

		0x60: {"LD H,B", 1, 4, (*CPU).Instruction60},
		0x61: {"LD H,C", 1, 4, (*CPU).Instruction61},
		0x62: {"LD H,D", 1, 4, (*CPU).Instruction62},
		0x63: {"LD H,E", 1, 4, (*CPU).Instruction63},
		0x64: {"LD H,H", 1, 4, (*CPU).Instruction64},
		0x65: {"LD H,L", 1, 4, (*CPU).Instruction65},
		0x66: {"LD H,(HL)", 1, 8, (*CPU).Instruction66},
		0x67: {"LD H,A", 1, 4, (*CPU).Instruction67},
		0x68: {"LD L,B", 1, 4, (*CPU).Instruction68},
		0x69: {"LD L,C", 1, 4, (*CPU).Instruction69},
		0x6A: {"LD L,D", 1, 4, (*CPU).Instruction6A},
		0x6B: {"LD L,E", 1, 4, (*CPU).Instruction6B},
		0x6C: {"LD L,H", 1, 4, (*CPU).Instruction6C},
		0x6D: {"LD L,L", 1, 4, (*CPU).Instruction6D},
		0x6E: {"LD L,(HL)", 1, 8, (*CPU).Instruction6E},
		0x6F: {"LD L,A", 1, 4, (*CPU).Instruction6F},

		0x70: {"LD (HL),B", 1, 8, (*CPU).Instruction70},
		0x71: {"LD (HL),C", 1, 8, (*CPU).Instruction71},
		0x72: {"LD (HL),D", 1, 8, (*CPU).Instruction72},
		0x73: {"LD (HL),E", 1, 8, (*CPU).Instruction73},
		0x74: {"LD (HL),H", 1, 8, (*CPU).Instruction74},
		0x75: {"LD (HL),L", 1, 8, (*CPU).Instruction75},
		0x76: {"HALT", 1, 4, (*CPU).Instruction76},
		0x77: {"LD (HL),A", 1, 8, (*CPU).Instruction77},
		0x78: {"LD A,B", 1, 4, (*CPU).Instruction78},
		0x79: {"LD A,C", 1, 4, (*CPU).Instruction79},
		0x7A: {"LD A,D", 1, 4, (*CPU).Instruction7A},
		0x7B: {"LD A,E", 1, 4, (*CPU).Instruction7B},
		0x7C: {"LD A,H", 1, 4, (*CPU).Instruction7C},
		0x7D: {"LD A,L", 1, 4, (*CPU).Instruction7D},
		0x7E: {"LD A,(HL)", 1, 8, (*CPU).Instruction7E},
		0x7F: {"LD A,A", 1, 4, (*CPU).Instruction7F},

		0x80: {"ADD A,B", 1, 4, (*CPU).Instruction80},
		0x81: {"ADD A,C", 1, 4, (*CPU).Instruction81},
		0x82: {"ADD A,D", 1, 4, (*CPU).Instruction82},
		0x83: {"ADD A,E", 1, 4, (*CPU).Instruction83},
		0x84: {"ADD A,H", 1, 4, (*CPU).Instruction84},
		0x85: {"ADD A,L", 1, 4, (*CPU).Instruction85},
		0x86: {"ADD A,(HL)", 1, 8, (*CPU).Instruction86},
		0x87: {"ADD A,A", 1, 4, (*CPU).Instruction87},
		0x88: {"ADC A,B", 1, 4, (*CPU).Instruction88},
		0x89: {"ADC A,C", 1, 4, (*CPU).Instruction89},
		0x8A: {"ADC A,D", 1, 4, (*CPU).Instruction8A},
		0x8B: {"ADC A,E", 1, 4, (*CPU).Instruction8B},
		0x8C: {"ADC A,H", 1, 4, (*CPU).Instruction8C},
		0x8D: {"ADC A,L", 1, 4, (*CPU).Instruction8D},
		0x8E: {"ADC A,(HL)", 1, 8, (*CPU).Instruction8E},
		0x8F: {"ADC A,A", 1, 4, (*CPU).Instruction8F},

		0x90: {"SUB B", 1, 4, (*CPU).Instruction90},
		0x91: {"SUB C", 1, 4, (*CPU).Instruction91},
		0x92: {"SUB D", 1, 4, (*CPU).Instruction92},
		0x93: {"SUB E", 1, 4, (*CPU).Instruction93},
		0x94: {"SUB H", 1, 4, (*CPU).Instruction94},
		0x95: {"SUB L", 1, 4, (*CPU).Instruction95},
		0x96: {"SUB (HL)", 1, 8, (*CPU).Instruction96},
		0x97: {"SUB A", 1, 4, (*CPU).Instruction97},
		0x98: {"SBC A,B", 1, 4, (*CPU).Instruction98},
		0x99: {"SBC A,C", 1, 4, (*CPU).Instruction99},
		0x9A: {"SBC A,D", 1, 4, (*CPU).Instruction9A},
		0x9B: {"SBC A,E", 1, 4, (*CPU).Instruction9B},
		0x9C: {"SBC A,H", 1, 4, (*CPU).Instruction9C},
		0x9D: {"SBC A,L", 1, 4, (*CPU).Instruction9D},
		0x9E: {"SBC A,(HL)", 1, 8, (*CPU).Instruction9E},
		0x9F: {"SBC A,A", 1, 4, (*CPU).Instruction9F},

		0xA0: {"AND B", 1, 4, (*CPU).InstructionA0},
		0xA1: {"AND C", 1, 4, (*CPU).InstructionA1},
		0xA2: {"AND D", 1, 4, (*CPU).InstructionA2},
		0xA3: {"AND E", 1, 4, (*CPU).InstructionA3},
		0xA4: {"AND H", 1, 4, (*CPU).InstructionA4},
		0xA5: {"AND L", 1, 4, (*CPU).InstructionA5},
		0xA6: {"AND (HL)", 1, 8, (*CPU).InstructionA6},
		0xA7: {"AND A", 1, 4, (*CPU).InstructionA7},
		0xA8: {"XOR B", 1, 4, (*CPU).InstructionA8},
		0xA9: {"XOR C", 1, 4, (*CPU).InstructionA9},
		0xAA: {"XOR D", 1, 4, (*CPU).InstructionAA},
		0xAB: {"XOR E", 1, 4, (*CPU).InstructionAB},
		0xAC: {"XOR H", 1, 4, (*CPU).InstructionAC},
		0xAD: {"XOR L", 1, 4, (*CPU).InstructionAD},
		0xAE: {"XOR (HL)", 1, 8, (*CPU).InstructionAE},
		0xAF: {"XOR A", 1, 4, (*CPU).InstructionAF},

		0xB0: {"OR B", 1, 4, (*CPU).InstructionB0},
		0xB1: {"OR C", 1, 4, (*CPU).InstructionB1},
		0xB2: {"OR D", 1, 4, (*CPU).InstructionB2},
		0xB3: {"OR E", 1, 4, (*CPU).InstructionB3},
		0xB4: {"OR H", 1, 4, (*CPU).InstructionB4},
		0xB5: {"OR L", 1, 4, (*CPU).InstructionB5},
		0xB6: {"OR (HL)", 1, 8, (*CPU).InstructionB6},
		0xB7: {"OR A", 1, 4, (*CPU).InstructionB7},
		0xB8: {"CP B", 1, 4, (*CPU).InstructionB8},
		0xB9: {"CP C", 1, 4, (*CPU).InstructionB9},
		0xBA: {"CP D", 1, 4, (*CPU).InstructionBA},
		0xBB: {"CP E", 1, 4, (*CPU).InstructionBB},
		0xBC: {"CP H", 1, 4, (*CPU).InstructionBC},
		0xBD: {"CP L", 1, 4, (*CPU).InstructionBD},
		0xBE: {"CP (HL)", 1, 8, (*CPU).InstructionBE},
		0xBF: {"CP A", 1, 4, (*CPU).InstructionBF},

		0xC0: {"RET NZ", 1, 8, (*CPU).InstructionC0},
		0xC1: {"POP BC", 1, 12, (*CPU).InstructionC1},
		0xC2: {"JP NZ,nn", 3, 12, (*CPU).InstructionC2},
		0xC3: {"JP nn", 3, 16, (*CPU).InstructionC3},
		0xC4: {"CALL NZ,nn", 3, 12, (*CPU).InstructionC4},
		0xC5: {"PUSH BC", 1, 16, (*CPU).InstructionC5},
		0xC6: {"ADD A,n", 2, 8, (*CPU).InstructionC6},
		0xC7: {"RST 00H", 1, 16, (*CPU).InstructionC7},
		0xC8: {"RET Z", 1, 8, (*CPU).InstructionC8},
		0xC9: {"RET", 1, 16, (*CPU).InstructionC9},
		0xCA: {"JP Z,nn", 3, 12, (*CPU).InstructionCA},
		0xCB: {"PREFIX CB", 1, 4, (*CPU).InstructionCB},
		0xCC: {"CALL Z,nn", 3, 12, (*CPU).InstructionCC},
		0xCD: {"CALL nn", 3, 24, (*CPU).InstructionCD},
		0xCE: {"ADC A,n", 2, 8, (*CPU).InstructionCE},
		0xCF: {"RST 08H", 1, 16, (*CPU).InstructionCF},

		0xD0: {"RET NC", 1, 8, (*CPU).InstructionD0},
		0xD1: {"POP DE", 1, 12, (*CPU).InstructionD1},
		0xD2: {"JP NC,nn", 3, 12, (*CPU).InstructionD2},
		0xD4: {"CALL NC,nn", 3, 12, (*CPU).InstructionD4},
		0xD5: {"PUSH DE", 1, 16, (*CPU).InstructionD5},
		0xD6: {"SUB n", 2, 8, (*CPU).InstructionD6},
		0xD7: {"RST 10H", 1, 16, (*CPU).InstructionD7},
		0xD8: {"RET C", 1, 8, (*CPU).InstructionD8},
		0xD9: {"RETI", 1, 16, (*CPU).InstructionD9},
		0xDA: {"JP C,nn", 3, 12, (*CPU).InstructionDA},
		0xDC: {"CALL C,nn", 3, 12, (*CPU).InstructionDC},
		0xDE: {"SBC A,n", 2, 8, (*CPU).InstructionDE},
		0xDF: {"RST 18H", 1, 16, (*CPU).InstructionDF},

		0xE0: {"LDH (n),A", 2, 12, (*CPU).InstructionE0},
		0xE1: {"POP HL", 1, 12, (*CPU).InstructionE1},
		0xE2: {"LD (C),A", 1, 8, (*CPU).InstructionE2},
		0xE5: {"PUSH HL", 1, 16, (*CPU).InstructionE5},
		0xE6: {"AND n", 2, 8, (*CPU).InstructionE6},
		0xE7: {"RST 20H", 1, 16, (*CPU).InstructionE7},
		0xE8: {"ADD SP,e", 2, 16, (*CPU).InstructionE8},
		0xE9: {"JP HL", 1, 4, (*CPU).InstructionE9},
		0xEA: {"LD (nn),A", 3, 16, (*CPU).InstructionEA},
		0xEE: {"XOR n", 2, 8, (*CPU).InstructionEE},
		0xEF: {"RST 28H", 1, 16, (*CPU).InstructionEF},

		0xF0: {"LDH A,(n)", 2, 12, (*CPU).InstructionF0},
		0xF1: {"POP AF", 1, 12, (*CPU).InstructionF1},
		0xF2: {"LD A,(C)", 1, 8, (*CPU).InstructionF2},
		0xF3: {"DI", 1, 4, (*CPU).InstructionF3},
		0xF5: {"PUSH AF", 1, 16, (*CPU).InstructionF5},
		0xF6: {"OR n", 2, 8, (*CPU).InstructionF6},
		0xF7: {"RST 30H", 1, 16, (*CPU).InstructionF7},
		0xF8: {"LDHL SP+e", 2, 12, (*CPU).InstructionF8},
		0xF9: {"LD SP,HL", 1, 8, (*CPU).InstructionF9},
		0xFA: {"LD A,(nn)", 3, 16, (*CPU).InstructionFA},
		0xFB: {"EI", 1, 4, (*CPU).InstructionFB},
		0xFE: {"CP n", 2, 8, (*CPU).InstructionFE},
		0xFF: {"RST 38H", 1, 16, (*CPU).InstructionFF},
	}
}

func (c *CPU) Instruction00() {}

func (c *CPU) Instruction01() {
	msb := uint16(c.ReadPC())
	lsb := uint16(c.ReadPC())
	value := msb<<8 | lsb
	c.Registers.SetBC(value)
}

func (c *CPU) Instruction02() {
	value := c.Registers.GetA()
	addr := c.Registers.GetBC()
	c.WriteMemory(addr, value)
}

func (c *CPU) Instruction03() {
	c.Registers.IncBC()
}

func (c *CPU) Instruction04() {
	value1 := c.Registers.GetB()
	value2 := value1 + 1
	c.Flags.SetZ(value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(value2&0xF == 0)
	c.Registers.IncB()
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction05() {
	value1 := c.Registers.GetB()
	value2 := value1 - 1
	c.Flags.SetZ(value2 == 0)
	c.Flags.SetN(true)
	c.Flags.SetH(value1&0xF == 0)
	c.Registers.DecB()
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction06() {
	value := c.ReadPC()
	c.Registers.SetB(value)
}

func (c *CPU) Instruction07() {}

func (c *CPU) Instruction08() {
	msb := uint16(c.ReadPC())
	lsb := uint16(c.ReadPC())
	addr := msb<<8 | lsb
	c.WriteMemory(addr, c.Registers.GetP())
	c.WriteMemory(addr+1, c.Registers.GetS())
}

func (c *CPU) Instruction09() {
	value1 := c.Registers.GetHL()
	value2 := c.Registers.GetBC()
	c.Flags.SetN(false)
	c.Flags.SetH((value1&0xFFF)+(value2&0xFFF) > 0xFFF)
	c.Flags.SetC(uint32(value1)+uint32(value2) > 0xFFFF)
	c.Registers.SetHL(value1 + value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction0A() {
	addr := c.Registers.GetBC()
	value := c.ReadMemory(addr)
	c.Registers.SetA(value)
}

func (c *CPU) Instruction0B() {
	c.Registers.DecBC()
}

func (c *CPU) Instruction0C() {
	value1 := c.Registers.GetC()
	value2 := value1 + 1
	c.Flags.SetZ(value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(value2&0xF == 0)
	c.Registers.IncC()
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction0D() {
	value1 := c.Registers.GetC()
	value2 := value1 - 1
	c.Flags.SetZ(value2 == 0)
	c.Flags.SetN(true)
	c.Flags.SetH(value1&0xF == 0)
	c.Registers.DecC()
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction0E() {
	value := c.ReadPC()
	c.Registers.SetC(value)
}

func (c *CPU) Instruction0F() {}

func (c *CPU) Instruction10() {}

func (c *CPU) Instruction11() {
	msb := uint16(c.ReadPC())
	lsb := uint16(c.ReadPC())
	value := msb<<8 | lsb
	c.Registers.SetDE(value)
}

func (c *CPU) Instruction12() {
	value := c.Registers.GetA()
	addr := c.Registers.GetDE()
	c.WriteMemory(addr, value)
}

func (c *CPU) Instruction13() {
	c.Registers.IncDE()
}

func (c *CPU) Instruction14() {
	value1 := c.Registers.GetD()
	value2 := value1 + 1
	c.Flags.SetZ(value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(value2&0xF == 0)
	c.Registers.IncD()
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction15() {
	value1 := c.Registers.GetD()
	value2 := value1 - 1
	c.Flags.SetZ(value2 == 0)
	c.Flags.SetN(true)
	c.Flags.SetH(value1&0xF == 0)
	c.Registers.DecD()
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction16() {
	value := c.ReadPC()
	c.Registers.SetD(value)
}

func (c *CPU) Instruction17() {}

func (c *CPU) Instruction18() {}

func (c *CPU) Instruction19() {
	value1 := c.Registers.GetHL()
	value2 := c.Registers.GetDE()
	c.Flags.SetN(false)
	c.Flags.SetH((value1&0xFFF)+(value2&0xFFF) > 0xFFF)
	c.Flags.SetC(uint32(value1)+uint32(value2) > 0xFFFF)
	c.Registers.SetHL(value1 + value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction1A() {
	addr := c.Registers.GetDE()
	value := c.ReadMemory(addr)
	c.Registers.SetA(value)
}

func (c *CPU) Instruction1B() {
	c.Registers.DecDE()
}

func (c *CPU) Instruction1C() {
	value1 := c.Registers.GetE()
	value2 := value1 + 1
	c.Flags.SetZ(value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(value2&0xF == 0)
	c.Registers.IncE()
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction1D() {
	value1 := c.Registers.GetE()
	value2 := value1 - 1
	c.Flags.SetZ(value2 == 0)
	c.Flags.SetN(true)
	c.Flags.SetH(value1&0xF == 0)
	c.Registers.DecE()
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction1E() {
	value := c.ReadPC()
	c.Registers.SetE(value)
}

func (c *CPU) Instruction1F() {}

func (c *CPU) Instruction20() {}

func (c *CPU) Instruction21() {
	msb := uint16(c.ReadPC())
	lsb := uint16(c.ReadPC())
	value := msb<<8 | lsb
	c.Registers.SetHL(value)
}

func (c *CPU) Instruction22() {
	value := c.Registers.GetA()
	addr := c.Registers.GetHL()
	c.WriteMemory(addr, value)
	c.Registers.IncHL()
}

func (c *CPU) Instruction23() {
	c.Registers.IncHL()
}

func (c *CPU) Instruction24() {
	value1 := c.Registers.GetH()
	value2 := value1 + 1
	c.Flags.SetZ(value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(value2&0xF == 0)
	c.Registers.IncH()
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction25() {
	value1 := c.Registers.GetH()
	value2 := value1 - 1
	c.Flags.SetZ(value2 == 0)
	c.Flags.SetN(true)
	c.Flags.SetH(value1&0xF == 0)
	c.Registers.DecH()
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction26() {
	value := c.ReadPC()
	c.Registers.SetH(value)
}

func (c *CPU) Instruction27() {

}

func (c *CPU) Instruction28() {}

func (c *CPU) Instruction29() {
	value1 := c.Registers.GetHL()
	value2 := c.Registers.GetHL()
	c.Flags.SetN(false)
	c.Flags.SetH((value1&0xFFF)+(value2&0xFFF) > 0xFFF)
	c.Flags.SetC(uint32(value1)+uint32(value2) > 0xFFFF)
	c.Registers.SetHL(value1 + value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction2A() {
	addr := c.Registers.GetHL()
	value := c.ReadMemory(addr)
	c.Registers.SetA(value)
	c.Registers.IncHL()
}

func (c *CPU) Instruction2B() {
	c.Registers.DecHL()
}

func (c *CPU) Instruction2C() {
	value1 := c.Registers.GetL()
	value2 := value1 + 1
	c.Flags.SetZ(value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(value2&0xF == 0)
	c.Registers.IncL()
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction2D() {
	value1 := c.Registers.GetL()
	value2 := value1 - 1
	c.Flags.SetZ(value2 == 0)
	c.Flags.SetN(true)
	c.Flags.SetH(value1&0xF == 0)
	c.Registers.DecL()
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction2E() {
	value := c.ReadPC()
	c.Registers.SetL(value)
}

func (c *CPU) Instruction2F() {}

func (c *CPU) Instruction30() {}

func (c *CPU) Instruction31() {
	msb := uint16(c.ReadPC())
	lsb := uint16(c.ReadPC())
	value := msb<<8 | lsb
	c.Registers.SetSP(value)
}

func (c *CPU) Instruction32() {
	value := c.Registers.GetA()
	addr := c.Registers.GetHL()
	c.WriteMemory(addr, value)
	c.Registers.DecHL()
}

func (c *CPU) Instruction33() {
	c.Registers.IncSP()
}

func (c *CPU) Instruction34() {
	addr := c.Registers.GetHL()
	value1 := c.ReadMemory(addr)
	value2 := value1 + 1
	c.Flags.SetZ(value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(value2&0xF == 0)
	c.WriteMemory(addr, value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction35() {
	addr := c.Registers.GetHL()
	value1 := c.ReadMemory(addr)
	value2 := value1 - 1
	c.Flags.SetZ(value2 == 0)
	c.Flags.SetN(true)
	c.Flags.SetH(value1&0xF == 0)
	c.WriteMemory(addr, value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction36() {
	value := c.ReadPC()
	addr := c.Registers.GetHL()
	c.WriteMemory(addr, value)
}

func (c *CPU) Instruction37() {}

func (c *CPU) Instruction38() {}

func (c *CPU) Instruction39() {
	value1 := c.Registers.GetHL()
	value2 := c.Registers.GetSP()
	c.Flags.SetN(false)
	c.Flags.SetH((value1&0xFFF)+(value2&0xFFF) > 0xFFF)
	c.Flags.SetC(uint32(value1)+uint32(value2) > 0xFFFF)
	c.Registers.SetHL(value1 + value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction3A() {
	addr := c.Registers.GetHL()
	value := c.ReadMemory(addr)
	c.Registers.SetA(value)
	c.Registers.DecHL()
}

func (c *CPU) Instruction3B() {
	c.Registers.DecSP()
}

func (c *CPU) Instruction3C() {
	value1 := c.Registers.GetA()
	value2 := value1 + 1
	c.Flags.SetZ(value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(value2&0xF == 0)
	c.Registers.IncA()
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction3D() {
	value1 := c.Registers.GetA()
	value2 := value1 - 1
	c.Flags.SetZ(value2 == 0)
	c.Flags.SetN(true)
	c.Flags.SetH(value1&0xF == 0)
	c.Registers.DecA()
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction3E() {
	value := c.ReadPC()
	c.Registers.SetA(value)
}

func (c *CPU) Instruction3F() {}

func (c *CPU) Instruction40() {
	value := c.Registers.GetB()
	c.Registers.SetB(value)
}

func (c *CPU) Instruction41() {
	value := c.Registers.GetC()
	c.Registers.SetB(value)
}

func (c *CPU) Instruction42() {
	value := c.Registers.GetD()
	c.Registers.SetB(value)
}

func (c *CPU) Instruction43() {
	value := c.Registers.GetE()
	c.Registers.SetB(value)
}

func (c *CPU) Instruction44() {
	value := c.Registers.GetH()
	c.Registers.SetB(value)
}

func (c *CPU) Instruction45() {
	value := c.Registers.GetL()
	c.Registers.SetB(value)
}

func (c *CPU) Instruction46() {
	addr := c.Registers.GetHL()
	value := c.ReadMemory(addr)
	c.Registers.SetB(value)
}

func (c *CPU) Instruction47() {
	value := c.Registers.GetA()
	c.Registers.SetB(value)
}

func (c *CPU) Instruction48() {
	value := c.Registers.GetB()
	c.Registers.SetC(value)
}

func (c *CPU) Instruction49() {
	value := c.Registers.GetC()
	c.Registers.SetC(value)
}

func (c *CPU) Instruction4A() {
	value := c.Registers.GetD()
	c.Registers.SetC(value)
}

func (c *CPU) Instruction4B() {
	value := c.Registers.GetE()
	c.Registers.SetC(value)
}

func (c *CPU) Instruction4C() {
	value := c.Registers.GetH()
	c.Registers.SetC(value)
}

func (c *CPU) Instruction4D() {
	value := c.Registers.GetL()
	c.Registers.SetC(value)
}

func (c *CPU) Instruction4E() {
	addr := c.Registers.GetHL()
	value := c.ReadMemory(addr)
	c.Registers.SetC(value)
}

func (c *CPU) Instruction4F() {
	value := c.Registers.GetA()
	c.Registers.SetC(value)
}

func (c *CPU) Instruction50() {
	value := c.Registers.GetB()
	c.Registers.SetD(value)
}

func (c *CPU) Instruction51() {
	value := c.Registers.GetC()
	c.Registers.SetD(value)
}

func (c *CPU) Instruction52() {
	value := c.Registers.GetD()
	c.Registers.SetD(value)
}

func (c *CPU) Instruction53() {
	value := c.Registers.GetE()
	c.Registers.SetD(value)
}

func (c *CPU) Instruction54() {
	value := c.Registers.GetH()
	c.Registers.SetD(value)
}

func (c *CPU) Instruction55() {
	value := c.Registers.GetL()
	c.Registers.SetD(value)
}

func (c *CPU) Instruction56() {
	addr := c.Registers.GetHL()
	value := c.ReadMemory(addr)
	c.Registers.SetD(value)
}

func (c *CPU) Instruction57() {
	value := c.Registers.GetA()
	c.Registers.SetD(value)
}

func (c *CPU) Instruction58() {
	value := c.Registers.GetB()
	c.Registers.SetE(value)
}

func (c *CPU) Instruction59() {
	value := c.Registers.GetC()
	c.Registers.SetE(value)
}

func (c *CPU) Instruction5A() {
	value := c.Registers.GetD()
	c.Registers.SetE(value)
}

func (c *CPU) Instruction5B() {
	value := c.Registers.GetE()
	c.Registers.SetE(value)
}

func (c *CPU) Instruction5C() {
	value := c.Registers.GetH()
	c.Registers.SetE(value)
}

func (c *CPU) Instruction5D() {
	value := c.Registers.GetL()
	c.Registers.SetE(value)
}

func (c *CPU) Instruction5E() {
	addr := c.Registers.GetHL()
	value := c.ReadMemory(addr)
	c.Registers.SetE(value)
}

func (c *CPU) Instruction5F() {
	value := c.Registers.GetA()
	c.Registers.SetE(value)
}

func (c *CPU) Instruction60() {
	value := c.Registers.GetB()
	c.Registers.SetH(value)
}

func (c *CPU) Instruction61() {
	value := c.Registers.GetC()
	c.Registers.SetH(value)
}

func (c *CPU) Instruction62() {
	value := c.Registers.GetD()
	c.Registers.SetH(value)
}

func (c *CPU) Instruction63() {
	value := c.Registers.GetE()
	c.Registers.SetH(value)
}

func (c *CPU) Instruction64() {
	value := c.Registers.GetH()
	c.Registers.SetH(value)
}

func (c *CPU) Instruction65() {
	value := c.Registers.GetL()
	c.Registers.SetH(value)
}

func (c *CPU) Instruction66() {
	addr := c.Registers.GetHL()
	value := c.ReadMemory(addr)
	c.Registers.SetH(value)
}

func (c *CPU) Instruction67() {
	value := c.Registers.GetA()
	c.Registers.SetH(value)
}

func (c *CPU) Instruction68() {
	value := c.Registers.GetB()
	c.Registers.SetL(value)
}

func (c *CPU) Instruction69() {
	value := c.Registers.GetC()
	c.Registers.SetL(value)
}

func (c *CPU) Instruction6A() {
	value := c.Registers.GetD()
	c.Registers.SetL(value)
}

func (c *CPU) Instruction6B() {
	value := c.Registers.GetE()
	c.Registers.SetL(value)
}

func (c *CPU) Instruction6C() {
	value := c.Registers.GetH()
	c.Registers.SetL(value)
}

func (c *CPU) Instruction6D() {
	value := c.Registers.GetL()
	c.Registers.SetL(value)
}

func (c *CPU) Instruction6E() {
	addr := c.Registers.GetHL()
	value := c.ReadMemory(addr)
	c.Registers.SetL(value)
}

func (c *CPU) Instruction6F() {
	value := c.Registers.GetA()
	c.Registers.SetL(value)
}

func (c *CPU) Instruction70() {
	value := c.Registers.GetB()
	addr := c.Registers.GetHL()
	c.WriteMemory(addr, value)
}

func (c *CPU) Instruction71() {
	value := c.Registers.GetC()
	addr := c.Registers.GetHL()
	c.WriteMemory(addr, value)
}

func (c *CPU) Instruction72() {
	value := c.Registers.GetD()
	addr := c.Registers.GetHL()
	c.WriteMemory(addr, value)
}

func (c *CPU) Instruction73() {
	value := c.Registers.GetE()
	addr := c.Registers.GetHL()
	c.WriteMemory(addr, value)
}

func (c *CPU) Instruction74() {
	value := c.Registers.GetH()
	addr := c.Registers.GetHL()
	c.WriteMemory(addr, value)
}

func (c *CPU) Instruction75() {
	value := c.Registers.GetL()
	addr := c.Registers.GetHL()
	c.WriteMemory(addr, value)
}

func (c *CPU) Instruction76() {}

func (c *CPU) Instruction77() {
	value := c.Registers.GetA()
	addr := c.Registers.GetHL()
	c.WriteMemory(addr, value)
}

func (c *CPU) Instruction78() {
	value := c.Registers.GetB()
	c.Registers.SetA(value)
}

func (c *CPU) Instruction79() {
	value := c.Registers.GetC()
	c.Registers.SetA(value)
}

func (c *CPU) Instruction7A() {
	value := c.Registers.GetD()
	c.Registers.SetA(value)
}

func (c *CPU) Instruction7B() {
	value := c.Registers.GetE()
	c.Registers.SetA(value)
}

func (c *CPU) Instruction7C() {
	value := c.Registers.GetH()
	c.Registers.SetA(value)
}

func (c *CPU) Instruction7D() {
	value := c.Registers.GetL()
	c.Registers.SetA(value)
}

func (c *CPU) Instruction7E() {
	addr := c.Registers.GetHL()
	value := c.ReadMemory(addr)
	c.Registers.SetA(value)
}

func (c *CPU) Instruction7F() {
	value := c.Registers.GetA()
	c.Registers.SetA(value)
}

func (c *CPU) Instruction80() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetB()
	c.Flags.SetZ(value1+value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH((value1&0xF)+(value2&0xF) > 0xF)
	c.Flags.SetC(uint16(value1)+uint16(value2) > 0xFF)
	c.Registers.SetA(value1 + value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction81() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetC()
	c.Flags.SetZ(value1+value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH((value1&0xF)+(value2&0xF) > 0xF)
	c.Flags.SetC(uint16(value1)+uint16(value2) > 0xFF)
	c.Registers.SetA(value1 + value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction82() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetD()
	c.Registers.SetA(value1 + value2)
	c.Flags.SetZ(value1+value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH((value1&0xF)+(value2&0xF) > 0xF)
	c.Flags.SetC(uint16(value1)+uint16(value2) > 0xFF)
	c.Registers.SetA(value1 + value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction83() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetE()
	c.Registers.SetA(value1 + value2)
	c.Flags.SetZ(value1+value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH((value1&0xF)+(value2&0xF) > 0xF)
	c.Flags.SetC(uint16(value1)+uint16(value2) > 0xFF)
	c.Registers.SetA(value1 + value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction84() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetH()
	c.Registers.SetA(value1 + value2)
	c.Flags.SetZ(value1+value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH((value1&0xF)+(value2&0xF) > 0xF)
	c.Flags.SetC(uint16(value1)+uint16(value2) > 0xFF)
	c.Registers.SetA(value1 + value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction85() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetL()
	c.Registers.SetA(value1 + value2)
	c.Flags.SetZ(value1+value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH((value1&0xF)+(value2&0xF) > 0xF)
	c.Flags.SetC(uint16(value1)+uint16(value2) > 0xFF)
	c.Registers.SetA(value1 + value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction86() {
	addr := c.Registers.GetHL()
	value1 := c.Registers.GetA()
	value2 := c.ReadMemory(addr)
	c.Flags.SetZ(value1+value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH((value1&0xF)+(value2&0xF) > 0xF)
	c.Flags.SetC(uint16(value1)+uint16(value2) > 0xFF)
	c.Registers.SetA(value1 + value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction87() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetA()
	c.Registers.SetA(value1 + value2)
	c.Flags.SetZ(value1+value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH((value1&0xF)+(value2&0xF) > 0xF)
	c.Flags.SetC(uint16(value1)+uint16(value2) > 0xFF)
	c.Registers.SetA(value1 + value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction88() {
	carry := c.Flags.GetCarryAsValue()
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetB()
	c.Flags.SetZ(value1+value2+carry == 0)
	c.Flags.SetN(false)
	c.Flags.SetH((value1&0xF)+(value2&0xF)+carry > 0xF)
	c.Flags.SetC(uint16(value1)+uint16(value2)+uint16(carry) > 0xFF)
	c.Registers.SetA(value1 + value2 + carry)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction89() {
	carry := c.Flags.GetCarryAsValue()
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetC()
	c.Flags.SetZ(value1+value2+carry == 0)
	c.Flags.SetN(false)
	c.Flags.SetH((value1&0xF)+(value2&0xF)+carry > 0xF)
	c.Flags.SetC(uint16(value1)+uint16(value2)+uint16(carry) > 0xFF)
	c.Registers.SetA(value1 + value2 + carry)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction8A() {
	carry := c.Flags.GetCarryAsValue()
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetD()
	c.Flags.SetZ(value1+value2+carry == 0)
	c.Flags.SetN(false)
	c.Flags.SetH((value1&0xF)+(value2&0xF)+carry > 0xF)
	c.Flags.SetC(uint16(value1)+uint16(value2)+uint16(carry) > 0xFF)
	c.Registers.SetA(value1 + value2 + carry)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction8B() {
	carry := c.Flags.GetCarryAsValue()
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetE()
	c.Flags.SetZ(value1+value2+carry == 0)
	c.Flags.SetN(false)
	c.Flags.SetH((value1&0xF)+(value2&0xF)+carry > 0xF)
	c.Flags.SetC(uint16(value1)+uint16(value2)+uint16(carry) > 0xFF)
	c.Registers.SetA(value1 + value2 + carry)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction8C() {
	carry := c.Flags.GetCarryAsValue()
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetH()
	c.Flags.SetZ(value1+value2+carry == 0)
	c.Flags.SetN(false)
	c.Flags.SetH((value1&0xF)+(value2&0xF)+carry > 0xF)
	c.Flags.SetC(uint16(value1)+uint16(value2)+uint16(carry) > 0xFF)
	c.Registers.SetA(value1 + value2 + carry)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction8D() {
	carry := c.Flags.GetCarryAsValue()
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetL()
	c.Flags.SetZ(value1+value2+carry == 0)
	c.Flags.SetN(false)
	c.Flags.SetH((value1&0xF)+(value2&0xF)+carry > 0xF)
	c.Flags.SetC(uint16(value1)+uint16(value2)+uint16(carry) > 0xFF)
	c.Registers.SetA(value1 + value2 + carry)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction8E() {
	carry := c.Flags.GetCarryAsValue()
	addr := c.Registers.GetHL()
	value1 := c.Registers.GetA()
	value2 := c.ReadMemory(addr)
	c.Flags.SetZ(value1+value2+carry == 0)
	c.Flags.SetN(false)
	c.Flags.SetH((value1&0xF)+(value2&0xF)+carry > 0xF)
	c.Flags.SetC(uint16(value1)+uint16(value2)+uint16(carry) > 0xFF)
	c.Registers.SetA(value1 + value2 + carry)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction8F() {
	carry := c.Flags.GetCarryAsValue()
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetA()
	c.Flags.SetZ(value1+value2+carry == 0)
	c.Flags.SetN(false)
	c.Flags.SetH((value1&0xF)+(value2&0xF)+carry > 0xF)
	c.Flags.SetC(uint16(value1)+uint16(value2)+uint16(carry) > 0xFF)
	c.Registers.SetA(value1 + value2 + carry)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction90() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetB()
	c.Flags.SetZ(value1-value2 == 0)
	c.Flags.SetN(true)
	c.Flags.SetH((value1 & 0xF) < (value2 & 0xF))
	c.Flags.SetC(uint16(value1) < uint16(value2))
	c.Registers.SetA(value1 - value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction91() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetC()
	c.Flags.SetZ(value1-value2 == 0)
	c.Flags.SetN(true)
	c.Flags.SetH((value1 & 0xF) < (value2 & 0xF))
	c.Flags.SetC(uint16(value1) < uint16(value2))
	c.Registers.SetA(value1 - value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction92() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetD()
	c.Flags.SetZ(value1-value2 == 0)
	c.Flags.SetN(true)
	c.Flags.SetH((value1 & 0xF) < (value2 & 0xF))
	c.Flags.SetC(uint16(value1) < uint16(value2))
	c.Registers.SetA(value1 - value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction93() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetE()
	c.Flags.SetZ(value1-value2 == 0)
	c.Flags.SetN(true)
	c.Flags.SetH((value1 & 0xF) < (value2 & 0xF))
	c.Flags.SetC(uint16(value1) < uint16(value2))
	c.Registers.SetA(value1 - value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction94() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetH()
	c.Flags.SetZ(value1-value2 == 0)
	c.Flags.SetN(true)
	c.Flags.SetH((value1 & 0xF) < (value2 & 0xF))
	c.Flags.SetC(uint16(value1) < uint16(value2))
	c.Registers.SetA(value1 - value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction95() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetL()
	c.Flags.SetZ(value1-value2 == 0)
	c.Flags.SetN(true)
	c.Flags.SetH((value1 & 0xF) < (value2 & 0xF))
	c.Flags.SetC(uint16(value1) < uint16(value2))
	c.Registers.SetA(value1 - value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction96() {
	addr := c.Registers.GetHL()
	value1 := c.Registers.GetA()
	value2 := c.ReadMemory(addr)
	c.Flags.SetZ(value1-value2 == 0)
	c.Flags.SetN(true)
	c.Flags.SetH((value1 & 0xF) < (value2 & 0xF))
	c.Flags.SetC(uint16(value1) < uint16(value2))
	c.Registers.SetA(value1 - value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction97() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetA()
	c.Flags.SetZ(value1-value2 == 0)
	c.Flags.SetN(true)
	c.Flags.SetH((value1 & 0xF) < (value2 & 0xF))
	c.Flags.SetC(uint16(value1) < uint16(value2))
	c.Registers.SetA(value1 - value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction98() {
	carry := c.Flags.GetCarryAsValue()
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetB()
	c.Flags.SetZ(value1-value2-carry == 0)
	c.Flags.SetN(true)
	c.Flags.SetH((value1 & 0xF) < ((value2 & 0xF) + carry))
	c.Flags.SetC(uint16(value1) < (uint16(value2) + uint16(carry)))
	c.Registers.SetA(value1 - value2 - carry)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction99() {
	carry := c.Flags.GetCarryAsValue()
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetC()
	c.Flags.SetZ(value1-value2-carry == 0)
	c.Flags.SetN(true)
	c.Flags.SetH((value1 & 0xF) < ((value2 & 0xF) + carry))
	c.Flags.SetC(uint16(value1) < (uint16(value2) + uint16(carry)))
	c.Registers.SetA(value1 - value2 - carry)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction9A() {
	carry := c.Flags.GetCarryAsValue()
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetD()
	c.Flags.SetZ(value1-value2-carry == 0)
	c.Flags.SetN(true)
	c.Flags.SetH((value1 & 0xF) < ((value2 & 0xF) + carry))
	c.Flags.SetC(uint16(value1) < (uint16(value2) + uint16(carry)))
	c.Registers.SetA(value1 - value2 - carry)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction9B() {
	carry := c.Flags.GetCarryAsValue()
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetE()
	c.Flags.SetZ(value1-value2-carry == 0)
	c.Flags.SetN(true)
	c.Flags.SetH((value1 & 0xF) < ((value2 & 0xF) + carry))
	c.Flags.SetC(uint16(value1) < (uint16(value2) + uint16(carry)))
	c.Registers.SetA(value1 - value2 - carry)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction9C() {
	carry := c.Flags.GetCarryAsValue()
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetH()
	c.Flags.SetZ(value1-value2-carry == 0)
	c.Flags.SetN(true)
	c.Flags.SetH((value1 & 0xF) < ((value2 & 0xF) + carry))
	c.Flags.SetC(uint16(value1) < (uint16(value2) + uint16(carry)))
	c.Registers.SetA(value1 - value2 - carry)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction9D() {
	carry := c.Flags.GetCarryAsValue()
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetL()
	c.Flags.SetZ(value1-value2-carry == 0)
	c.Flags.SetN(true)
	c.Flags.SetH((value1 & 0xF) < ((value2 & 0xF) + carry))
	c.Flags.SetC(uint16(value1) < (uint16(value2) + uint16(carry)))
	c.Registers.SetA(value1 - value2 - carry)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction9E() {
	carry := c.Flags.GetCarryAsValue()
	addr := c.Registers.GetHL()
	value1 := c.Registers.GetA()
	value2 := c.ReadMemory(addr)
	c.Flags.SetZ(value1-value2-carry == 0)
	c.Flags.SetN(true)
	c.Flags.SetH((value1 & 0xF) < ((value2 & 0xF) + carry))
	c.Flags.SetC(uint16(value1) < (uint16(value2) + uint16(carry)))
	c.Registers.SetA(value1 - value2 - carry)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction9F() {
	carry := c.Flags.GetCarryAsValue()
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetA()
	c.Flags.SetZ(value1-value2-carry == 0)
	c.Flags.SetN(true)
	c.Flags.SetH((value1 & 0xF) < ((value2 & 0xF) + carry))
	c.Flags.SetC(uint16(value1) < (uint16(value2) + uint16(carry)))
	c.Registers.SetA(value1 - value2 - carry)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionA0() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetB()
	c.Flags.SetZ(value1&value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Flags.SetC(false)
	c.Registers.SetA(value1 & value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionA1() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetC()
	c.Flags.SetZ(value1&value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Flags.SetC(false)
	c.Registers.SetA(value1 & value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionA2() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetD()
	c.Flags.SetZ(value1&value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Flags.SetC(false)
	c.Registers.SetA(value1 & value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionA3() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetE()
	c.Flags.SetZ(value1&value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Flags.SetC(false)
	c.Registers.SetA(value1 & value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionA4() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetH()
	c.Flags.SetZ(value1&value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Flags.SetC(false)
	c.Registers.SetA(value1 & value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionA5() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetL()
	c.Flags.SetZ(value1&value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Flags.SetC(false)
	c.Registers.SetA(value1 & value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionA6() {
	addr := c.Registers.GetHL()
	value1 := c.Registers.GetA()
	value2 := c.ReadMemory(addr)
	c.Flags.SetZ(value1&value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Flags.SetC(false)
	c.Registers.SetA(value1 & value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionA7() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetA()
	c.Flags.SetZ(value1&value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Flags.SetC(false)
	c.Registers.SetA(value1 & value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionA8() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetB()
	c.Flags.SetZ(value1^value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(false)
	c.Registers.SetA(value1 ^ value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionA9() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetC()
	c.Flags.SetZ(value1^value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(false)
	c.Registers.SetA(value1 ^ value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionAA() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetD()
	c.Flags.SetZ(value1^value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(false)
	c.Registers.SetA(value1 ^ value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionAB() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetE()
	c.Flags.SetZ(value1^value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(false)
	c.Registers.SetA(value1 ^ value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionAC() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetH()
	c.Flags.SetZ(value1^value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(false)
	c.Registers.SetA(value1 ^ value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionAD() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetL()
	c.Flags.SetZ(value1^value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(false)
	c.Registers.SetA(value1 ^ value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionAE() {
	addr := c.Registers.GetHL()
	value1 := c.Registers.GetA()
	value2 := c.ReadMemory(addr)
	c.Flags.SetZ(value1^value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(false)
	c.Registers.SetA(value1 ^ value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionAF() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetA()
	c.Flags.SetZ(value1^value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(false)
	c.Registers.SetA(value1 ^ value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionB0() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetB()
	c.Flags.SetZ(value1|value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(false)
	c.Registers.SetA(value1 | value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionB1() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetC()
	c.Flags.SetZ(value1|value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(false)
	c.Registers.SetA(value1 | value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionB2() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetD()
	c.Flags.SetZ(value1|value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(false)
	c.Registers.SetA(value1 | value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionB3() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetE()
	c.Flags.SetZ(value1|value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(false)
	c.Registers.SetA(value1 | value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionB4() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetH()
	c.Flags.SetZ(value1|value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(false)
	c.Registers.SetA(value1 | value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionB5() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetL()
	c.Flags.SetZ(value1|value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(false)
	c.Registers.SetA(value1 | value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionB6() {
	addr := c.Registers.GetHL()
	value1 := c.Registers.GetA()
	value2 := c.ReadMemory(addr)
	c.Flags.SetZ(value1|value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(false)
	c.Registers.SetA(value1 | value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionB7() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetA()
	c.Flags.SetZ(value1|value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(false)
	c.Registers.SetA(value1 | value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionB8() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetB()
	c.Flags.SetZ(value1 == value2)
	c.Flags.SetN(true)
	c.Flags.SetH((value1 & 0xF) < (value2 & 0xF))
	c.Flags.SetC(value1 < value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionB9() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetC()
	c.Flags.SetZ(value1 == value2)
	c.Flags.SetN(true)
	c.Flags.SetH((value1 & 0xF) < (value2 & 0xF))
	c.Flags.SetC(value1 < value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionBA() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetD()
	c.Flags.SetZ(value1 == value2)
	c.Flags.SetN(true)
	c.Flags.SetH((value1 & 0xF) < (value2 & 0xF))
	c.Flags.SetC(value1 < value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionBB() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetE()
	c.Flags.SetZ(value1 == value2)
	c.Flags.SetN(true)
	c.Flags.SetH((value1 & 0xF) < (value2 & 0xF))
	c.Flags.SetC(value1 < value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionBC() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetH()
	c.Flags.SetZ(value1 == value2)
	c.Flags.SetN(true)
	c.Flags.SetH((value1 & 0xF) < (value2 & 0xF))
	c.Flags.SetC(value1 < value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionBD() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetL()
	c.Flags.SetZ(value1 == value2)
	c.Flags.SetN(true)
	c.Flags.SetH((value1 & 0xF) < (value2 & 0xF))
	c.Flags.SetC(value1 < value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionBE() {
	addr := c.Registers.GetHL()
	value1 := c.Registers.GetA()
	value2 := c.ReadMemory(addr)
	c.Flags.SetZ(value1 == value2)
	c.Flags.SetN(true)
	c.Flags.SetH((value1 & 0xF) < (value2 & 0xF))
	c.Flags.SetC(value1 < value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionBF() {
	value1 := c.Registers.GetA()
	value2 := c.Registers.GetA()
	c.Flags.SetZ(value1 == value2)
	c.Flags.SetN(true)
	c.Flags.SetH((value1 & 0xF) < (value2 & 0xF))
	c.Flags.SetC(value1 < value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionC0() {}

func (c *CPU) InstructionC1() {
	addr := c.Registers.GetSP()
	msb := uint16(c.ReadMemory(addr + 1))
	lsb := uint16(c.ReadMemory(addr))
	value := msb<<8 | lsb
	c.Registers.SetBC(value)
	c.Registers.SetSP(addr + 2)
}

func (c *CPU) InstructionC2() {}

func (c *CPU) InstructionC3() {}

func (c *CPU) InstructionC4() {}

func (c *CPU) InstructionC5() {
	msb := c.Registers.GetB()
	lsb := c.Registers.GetC()
	addr := c.Registers.GetSP()
	c.WriteMemory(addr-1, msb)
	c.WriteMemory(addr-2, lsb)
	c.Registers.SetSP(addr - 2)
}

func (c *CPU) InstructionC6() {
	value1 := c.Registers.GetA()
	value2 := c.ReadPC()
	c.Flags.SetZ(value1+value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH((value1&0xF)+(value2&0xF) > 0xF)
	c.Flags.SetC(uint16(value1)+uint16(value2) > 0xFF)
	c.Registers.SetA(value1 + value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionC7() {}

func (c *CPU) InstructionC8() {}

func (c *CPU) InstructionC9() {}

func (c *CPU) InstructionCA() {}

func (c *CPU) InstructionCB() {}

func (c *CPU) InstructionCC() {}

func (c *CPU) InstructionCD() {}

func (c *CPU) InstructionCE() {
	carry := c.Flags.GetCarryAsValue()
	value1 := c.Registers.GetA()
	value2 := c.ReadPC()
	c.Flags.SetZ(value1+value2+carry == 0)
	c.Flags.SetN(false)
	c.Flags.SetH((value1&0xF)+(value2&0xF)+carry > 0xF)
	c.Flags.SetC(uint16(value1)+uint16(value2)+uint16(carry) > 0xFF)
	c.Registers.SetA(value1 + value2 + carry)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCF() {}

func (c *CPU) InstructionD0() {}

func (c *CPU) InstructionD1() {
	addr := c.Registers.GetSP()
	msb := uint16(c.ReadMemory(addr + 1))
	lsb := uint16(c.ReadMemory(addr))
	value := msb<<8 | lsb
	c.Registers.SetDE(value)
	c.Registers.SetSP(addr + 2)
}

func (c *CPU) InstructionD2() {}

func (c *CPU) InstructionD3() {}

func (c *CPU) InstructionD4() {}

func (c *CPU) InstructionD5() {
	msb := c.Registers.GetD()
	lsb := c.Registers.GetE()
	addr := c.Registers.GetSP()
	c.WriteMemory(addr-1, msb)
	c.WriteMemory(addr-2, lsb)
	c.Registers.SetSP(addr - 2)
}

func (c *CPU) InstructionD6() {
	value1 := c.Registers.GetA()
	value2 := c.ReadPC()
	c.Flags.SetZ(value1-value2 == 0)
	c.Flags.SetN(true)
	c.Flags.SetH((value1 & 0xF) < (value2 & 0xF))
	c.Flags.SetC(uint16(value1) < uint16(value2))
	c.Registers.SetA(value1 - value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionD7() {}

func (c *CPU) InstructionD8() {}

func (c *CPU) InstructionD9() {}

func (c *CPU) InstructionDA() {}

func (c *CPU) InstructionDB() {}

func (c *CPU) InstructionDC() {}

func (c *CPU) InstructionDD() {}

func (c *CPU) InstructionDE() {
	carry := c.Flags.GetCarryAsValue()
	value1 := c.Registers.GetA()
	value2 := c.ReadPC()
	c.Flags.SetZ(value1-value2-carry == 0)
	c.Flags.SetN(true)
	c.Flags.SetH((value1 & 0xF) < ((value2 & 0xF) + carry))
	c.Flags.SetC(uint16(value1) < (uint16(value2) + uint16(carry)))
	c.Registers.SetA(value1 - value2 - carry)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionDF() {}

func (c *CPU) InstructionE0() {
	value := c.Registers.GetA()
	msb := uint16(0xFF)
	lsb := uint16(c.ReadPC())
	addr := msb<<8 | lsb
	c.WriteMemory(addr, value)
}

func (c *CPU) InstructionE1() {
	addr := c.Registers.GetSP()
	msb := uint16(c.ReadMemory(addr + 1))
	lsb := uint16(c.ReadMemory(addr))
	value := msb<<8 | lsb
	c.Registers.SetHL(value)
	c.Registers.SetSP(addr + 2)
}

func (c *CPU) InstructionE2() {
	value := c.Registers.GetA()
	msb := uint16(0xFF)
	lsb := uint16(c.Registers.GetC())
	addr := msb<<8 | lsb
	c.WriteMemory(addr, value)
}

func (c *CPU) InstructionE3() {}

func (c *CPU) InstructionE4() {}

func (c *CPU) InstructionE5() {
	msb := c.Registers.GetH()
	lsb := c.Registers.GetL()
	addr := c.Registers.GetSP()
	c.WriteMemory(addr-1, msb)
	c.WriteMemory(addr-2, lsb)
	c.Registers.SetSP(addr - 2)
}

func (c *CPU) InstructionE6() {
	value1 := c.Registers.GetA()
	value2 := c.ReadPC()
	c.Flags.SetZ(value1&value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Flags.SetC(false)
	c.Registers.SetA(value1 & value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionE7() {}

//	WARNING: POTENTIALLY BROKEN HALF-CARRY EVALUATION!
func (c *CPU) InstructionE8() {
	value1 := int(c.Registers.GetSP())
	value2 := int(c.ReadPC())
	value2 = (value2 & 127) - (value2 & 128)
	c.Flags.SetZ(false)
	c.Flags.SetN(false)
	if value2 < 0 {
		c.Flags.SetH((value1 & 0xF) < (value2 & 0xF))
		c.Flags.SetC((value1 & 0xFF) < (value2 & 0xFF))
	} else {
		c.Flags.SetH((value1&0xF)+(value2&0xF) > 0xF)
		c.Flags.SetC((value1&0xFF)+(value2&0xFF) > 0xFF)
	}
	c.Registers.SetSP(uint16(value1 + value2))
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))

}

func (c *CPU) InstructionE9() {}

func (c *CPU) InstructionEA() {
	value := c.Registers.GetA()
	msb := uint16(c.ReadPC())
	lsb := uint16(c.ReadPC())
	addr := msb<<8 | lsb
	c.WriteMemory(addr, value)
}

func (c *CPU) InstructionEB() {}

func (c *CPU) InstructionEC() {}

func (c *CPU) InstructionED() {}

func (c *CPU) InstructionEE() {
	value1 := c.Registers.GetA()
	value2 := c.ReadPC()
	c.Flags.SetZ(value1^value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(false)
	c.Registers.SetA(value1 ^ value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionEF() {}

func (c *CPU) InstructionF0() {
	msb := uint16(0xFF)
	lsb := uint16(c.ReadPC())
	addr := msb<<8 | lsb
	value := c.ReadMemory(addr)
	c.Registers.SetA(value)
}

func (c *CPU) InstructionF1() {
	addr := c.Registers.GetSP()
	msb := uint16(c.ReadMemory(addr + 1))
	lsb := uint16(c.ReadMemory(addr))
	value := msb<<8 | lsb
	c.Registers.SetAF(value)
	c.Registers.SetSP(addr + 2)
	c.Flags.SetFlagsFromValue(c.Registers.GetF())
}

func (c *CPU) InstructionF2() {
	msb := uint16(0xFF)
	lsb := uint16(c.Registers.GetC())
	addr := msb<<8 | lsb
	value := c.ReadMemory(addr)
	c.Registers.SetA(value)
}

func (c *CPU) InstructionF3() {}

func (c *CPU) InstructionF4() {}

func (c *CPU) InstructionF5() {
	msb := c.Registers.GetA()
	lsb := c.Registers.GetF()
	addr := c.Registers.GetSP()
	c.WriteMemory(addr-1, msb)
	c.WriteMemory(addr-2, lsb)
	c.Registers.SetSP(addr - 2)
}

func (c *CPU) InstructionF6() {
	value1 := c.Registers.GetA()
	value2 := c.ReadPC()
	c.Flags.SetZ(value1|value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(false)
	c.Registers.SetA(value1 | value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionF7() {}

//	WARNING: POTENTIALLY BROKEN HALF-CARRY EVALUATION!
func (c *CPU) InstructionF8() {
	value1 := int(c.Registers.GetSP())
	value2 := int(c.ReadPC())
	value2 = (value2 & 127) - (value2 & 128)
	c.Flags.SetZ(false)
	c.Flags.SetN(false)
	if value2 < 0 {
		c.Flags.SetH((value1 & 0xF) < (value2 & 0xF))
		c.Flags.SetC((value1 & 0xFF) < (value2 & 0xFF))
	} else {
		c.Flags.SetH((value1&0xF)+(value2&0xF) > 0xF)
		c.Flags.SetC((value1&0xFF)+(value2&0xFF) > 0xFF)
	}
	c.Registers.SetHL(uint16(value1 + value2))
	c.Registers.SetF(c.Flags.GetFlagsAsValue())
}

func (c *CPU) InstructionF9() {
	value := c.Registers.GetHL()
	c.Registers.SetSP(value)
}

func (c *CPU) InstructionFA() {
	msb := uint16(c.ReadPC())
	lsb := uint16(c.ReadPC())
	addr := msb<<8 | lsb
	value := c.ReadMemory(addr)
	c.Registers.SetA(value)
}

func (c *CPU) InstructionFB() {}

func (c *CPU) InstructionFC() {}

func (c *CPU) InstructionFD() {}

func (c *CPU) InstructionFE() {
	value1 := c.Registers.GetA()
	value2 := c.ReadPC()
	c.Flags.SetZ(value1 == value2)
	c.Flags.SetN(true)
	c.Flags.SetH((value1 & 0xF) < (value2 & 0xF))
	c.Flags.SetC(value1 < value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionFF() {}
