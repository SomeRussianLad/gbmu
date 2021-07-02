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

		0xCB00: {"RLC B", 2, 8, (*CPU).InstructionCB00},
		0xCB01: {"RLC C", 2, 8, (*CPU).InstructionCB01},
		0xCB02: {"RLC D", 2, 8, (*CPU).InstructionCB02},
		0xCB03: {"RLC E", 2, 8, (*CPU).InstructionCB03},
		0xCB04: {"RLC H", 2, 8, (*CPU).InstructionCB04},
		0xCB05: {"RLC L", 2, 8, (*CPU).InstructionCB05},
		0xCB06: {"RLC (HL)", 2, 16, (*CPU).InstructionCB06},
		0xCB07: {"RLC A", 2, 8, (*CPU).InstructionCB07},
		0xCB08: {"RRC B", 2, 8, (*CPU).InstructionCB08},
		0xCB09: {"RRC C", 2, 8, (*CPU).InstructionCB09},
		0xCB0A: {"RRC D", 2, 8, (*CPU).InstructionCB0A},
		0xCB0B: {"RRC E", 2, 8, (*CPU).InstructionCB0B},
		0xCB0C: {"RRC H", 2, 8, (*CPU).InstructionCB0C},
		0xCB0D: {"RRC L", 2, 8, (*CPU).InstructionCB0D},
		0xCB0E: {"RRC (HL)", 2, 16, (*CPU).InstructionCB0E},
		0xCB0F: {"RRC A", 2, 8, (*CPU).InstructionCB0F},

		0xCB10: {"RL B", 2, 8, (*CPU).InstructionCB10},
		0xCB11: {"RL C", 2, 8, (*CPU).InstructionCB11},
		0xCB12: {"RL D", 2, 8, (*CPU).InstructionCB12},
		0xCB13: {"RL E", 2, 8, (*CPU).InstructionCB13},
		0xCB14: {"RL H", 2, 8, (*CPU).InstructionCB14},
		0xCB15: {"RL L", 2, 8, (*CPU).InstructionCB15},
		0xCB16: {"RL (HL)", 2, 16, (*CPU).InstructionCB16},
		0xCB17: {"RL A", 2, 8, (*CPU).InstructionCB17},
		0xCB18: {"RR B", 2, 8, (*CPU).InstructionCB18},
		0xCB19: {"RR C", 2, 8, (*CPU).InstructionCB19},
		0xCB1A: {"RR D", 2, 8, (*CPU).InstructionCB1A},
		0xCB1B: {"RR E", 2, 8, (*CPU).InstructionCB1B},
		0xCB1C: {"RR H", 2, 8, (*CPU).InstructionCB1C},
		0xCB1D: {"RR L", 2, 8, (*CPU).InstructionCB1D},
		0xCB1E: {"RR (HL)", 2, 16, (*CPU).InstructionCB1E},
		0xCB1F: {"RR A", 2, 8, (*CPU).InstructionCB1F},

		0xCB20: {"SLA B", 2, 8, (*CPU).InstructionCB20},
		0xCB21: {"SLA C", 2, 8, (*CPU).InstructionCB21},
		0xCB22: {"SLA D", 2, 8, (*CPU).InstructionCB22},
		0xCB23: {"SLA E", 2, 8, (*CPU).InstructionCB23},
		0xCB24: {"SLA H", 2, 8, (*CPU).InstructionCB24},
		0xCB25: {"SLA L", 2, 8, (*CPU).InstructionCB25},
		0xCB26: {"SLA (HL)", 2, 16, (*CPU).InstructionCB26},
		0xCB27: {"SLA A", 2, 8, (*CPU).InstructionCB27},
		0xCB28: {"SRA B", 2, 8, (*CPU).InstructionCB28},
		0xCB29: {"SRA C", 2, 8, (*CPU).InstructionCB29},
		0xCB2A: {"SRA D", 2, 8, (*CPU).InstructionCB2A},
		0xCB2B: {"SRA E", 2, 8, (*CPU).InstructionCB2B},
		0xCB2C: {"SRA H", 2, 8, (*CPU).InstructionCB2C},
		0xCB2D: {"SRA L", 2, 8, (*CPU).InstructionCB2D},
		0xCB2E: {"SRA (HL)", 2, 16, (*CPU).InstructionCB2E},
		0xCB2F: {"SRA A", 2, 8, (*CPU).InstructionCB2F},

		0xCB30: {"SWAP B", 2, 8, (*CPU).InstructionCB30},
		0xCB31: {"SWAP C", 2, 8, (*CPU).InstructionCB31},
		0xCB32: {"SWAP D", 2, 8, (*CPU).InstructionCB32},
		0xCB33: {"SWAP E", 2, 8, (*CPU).InstructionCB33},
		0xCB34: {"SWAP H", 2, 8, (*CPU).InstructionCB34},
		0xCB35: {"SWAP L", 2, 8, (*CPU).InstructionCB35},
		0xCB36: {"SWAP (HL)", 2, 16, (*CPU).InstructionCB36},
		0xCB37: {"SWAP A", 2, 8, (*CPU).InstructionCB37},
		0xCB38: {"SRL B", 2, 8, (*CPU).InstructionCB38},
		0xCB39: {"SRL C", 2, 8, (*CPU).InstructionCB39},
		0xCB3A: {"SRL D", 2, 8, (*CPU).InstructionCB3A},
		0xCB3B: {"SRL E", 2, 8, (*CPU).InstructionCB3B},
		0xCB3C: {"SRL H", 2, 8, (*CPU).InstructionCB3C},
		0xCB3D: {"SRL L", 2, 8, (*CPU).InstructionCB3D},
		0xCB3E: {"SRL (HL)", 2, 16, (*CPU).InstructionCB3E},
		0xCB3F: {"SRL A", 2, 8, (*CPU).InstructionCB3F},

		0xCB40: {"BIT 0,B", 2, 8, (*CPU).InstructionCB40},
		0xCB41: {"BIT 0,C", 2, 8, (*CPU).InstructionCB41},
		0xCB42: {"BIT 0,D", 2, 8, (*CPU).InstructionCB42},
		0xCB43: {"BIT 0,E", 2, 8, (*CPU).InstructionCB43},
		0xCB44: {"BIT 0,H", 2, 8, (*CPU).InstructionCB44},
		0xCB45: {"BIT 0,L", 2, 8, (*CPU).InstructionCB45},
		0xCB46: {"BIT 0,(HL)", 2, 16, (*CPU).InstructionCB46},
		0xCB47: {"BIT 0,A", 2, 8, (*CPU).InstructionCB47},
		0xCB48: {"BIT 1,B", 2, 8, (*CPU).InstructionCB48},
		0xCB49: {"BIT 1,C", 2, 8, (*CPU).InstructionCB49},
		0xCB4A: {"BIT 1,D", 2, 8, (*CPU).InstructionCB4A},
		0xCB4B: {"BIT 1,E", 2, 8, (*CPU).InstructionCB4B},
		0xCB4C: {"BIT 1,H", 2, 8, (*CPU).InstructionCB4C},
		0xCB4D: {"BIT 1,L", 2, 8, (*CPU).InstructionCB4D},
		0xCB4E: {"BIT 1,(HL)", 2, 16, (*CPU).InstructionCB4E},
		0xCB4F: {"BIT 1,A", 2, 8, (*CPU).InstructionCB4F},
		0xCB50: {"BIT 2,B", 2, 8, (*CPU).InstructionCB50},

		0xCB51: {"BIT 2,C", 2, 8, (*CPU).InstructionCB51},
		0xCB52: {"BIT 2,D", 2, 8, (*CPU).InstructionCB52},
		0xCB53: {"BIT 2,E", 2, 8, (*CPU).InstructionCB53},
		0xCB54: {"BIT 2,H", 2, 8, (*CPU).InstructionCB54},
		0xCB55: {"BIT 2,L", 2, 8, (*CPU).InstructionCB55},
		0xCB56: {"BIT 2,(HL)", 2, 16, (*CPU).InstructionCB56},
		0xCB57: {"BIT 2,A", 2, 8, (*CPU).InstructionCB57},
		0xCB58: {"BIT 3,B", 2, 8, (*CPU).InstructionCB58},
		0xCB59: {"BIT 3,C", 2, 8, (*CPU).InstructionCB59},
		0xCB5A: {"BIT 3,D", 2, 8, (*CPU).InstructionCB5A},
		0xCB5B: {"BIT 3,E", 2, 8, (*CPU).InstructionCB5B},
		0xCB5C: {"BIT 3,H", 2, 8, (*CPU).InstructionCB5C},
		0xCB5D: {"BIT 3,L", 2, 8, (*CPU).InstructionCB5D},
		0xCB5E: {"BIT 3,(HL)", 2, 16, (*CPU).InstructionCB5E},
		0xCB5F: {"BIT 3,A", 2, 8, (*CPU).InstructionCB5F},

		0xCB60: {"BIT 4,B", 2, 8, (*CPU).InstructionCB60},
		0xCB61: {"BIT 4,C", 2, 8, (*CPU).InstructionCB61},
		0xCB62: {"BIT 4,D", 2, 8, (*CPU).InstructionCB62},
		0xCB63: {"BIT 4,E", 2, 8, (*CPU).InstructionCB63},
		0xCB64: {"BIT 4,H", 2, 8, (*CPU).InstructionCB64},
		0xCB65: {"BIT 4,L", 2, 8, (*CPU).InstructionCB65},
		0xCB66: {"BIT 4,(HL)", 2, 16, (*CPU).InstructionCB66},
		0xCB67: {"BIT 4,A", 2, 8, (*CPU).InstructionCB67},
		0xCB68: {"BIT 5,B", 2, 8, (*CPU).InstructionCB68},
		0xCB69: {"BIT 5,C", 2, 8, (*CPU).InstructionCB69},
		0xCB6A: {"BIT 5,D", 2, 8, (*CPU).InstructionCB6A},
		0xCB6B: {"BIT 5,E", 2, 8, (*CPU).InstructionCB6B},
		0xCB6C: {"BIT 5,H", 2, 8, (*CPU).InstructionCB6C},
		0xCB6D: {"BIT 5,L", 2, 8, (*CPU).InstructionCB6D},
		0xCB6E: {"BIT 5,(HL)", 2, 16, (*CPU).InstructionCB6E},
		0xCB6F: {"BIT 5,A", 2, 8, (*CPU).InstructionCB6F},

		0xCB70: {"BIT 6,B", 2, 8, (*CPU).InstructionCB70},
		0xCB71: {"BIT 6,C", 2, 8, (*CPU).InstructionCB71},
		0xCB72: {"BIT 6,D", 2, 8, (*CPU).InstructionCB72},
		0xCB73: {"BIT 6,E", 2, 8, (*CPU).InstructionCB73},
		0xCB74: {"BIT 6,H", 2, 8, (*CPU).InstructionCB74},
		0xCB75: {"BIT 6,L", 2, 8, (*CPU).InstructionCB75},
		0xCB76: {"BIT 6,(HL)", 2, 16, (*CPU).InstructionCB76},
		0xCB77: {"BIT 6,A", 2, 8, (*CPU).InstructionCB77},
		0xCB78: {"BIT 7,B", 2, 8, (*CPU).InstructionCB78},
		0xCB79: {"BIT 7,C", 2, 8, (*CPU).InstructionCB79},
		0xCB7A: {"BIT 7,D", 2, 8, (*CPU).InstructionCB7A},
		0xCB7B: {"BIT 7,E", 2, 8, (*CPU).InstructionCB7B},
		0xCB7C: {"BIT 7,H", 2, 8, (*CPU).InstructionCB7C},
		0xCB7D: {"BIT 7,L", 2, 8, (*CPU).InstructionCB7D},
		0xCB7E: {"BIT 7,(HL)", 2, 16, (*CPU).InstructionCB7E},
		0xCB7F: {"BIT 7,A", 2, 8, (*CPU).InstructionCB7F},

		0xCB80: {"RES 0,B", 2, 8, (*CPU).InstructionCB80},
		0xCB81: {"RES 0,C", 2, 8, (*CPU).InstructionCB81},
		0xCB82: {"RES 0,D", 2, 8, (*CPU).InstructionCB82},
		0xCB83: {"RES 0,E", 2, 8, (*CPU).InstructionCB83},
		0xCB84: {"RES 0,H", 2, 8, (*CPU).InstructionCB84},
		0xCB85: {"RES 0,L", 2, 8, (*CPU).InstructionCB85},
		0xCB86: {"RES 0,(HL)", 2, 16, (*CPU).InstructionCB86},
		0xCB87: {"RES 0,A", 2, 8, (*CPU).InstructionCB87},
		0xCB88: {"RES 1,B", 2, 8, (*CPU).InstructionCB88},
		0xCB89: {"RES 1,C", 2, 8, (*CPU).InstructionCB89},
		0xCB8A: {"RES 1,D", 2, 8, (*CPU).InstructionCB8A},
		0xCB8B: {"RES 1,E", 2, 8, (*CPU).InstructionCB8B},
		0xCB8C: {"RES 1,H", 2, 8, (*CPU).InstructionCB8C},
		0xCB8D: {"RES 1,L", 2, 8, (*CPU).InstructionCB8D},
		0xCB8E: {"RES 1,(HL)", 2, 16, (*CPU).InstructionCB8E},
		0xCB8F: {"RES 1,A", 2, 8, (*CPU).InstructionCB8F},

		0xCB90: {"RES 2,B", 2, 8, (*CPU).InstructionCB90},
		0xCB91: {"RES 2,C", 2, 8, (*CPU).InstructionCB91},
		0xCB92: {"RES 2,D", 2, 8, (*CPU).InstructionCB92},
		0xCB93: {"RES 2,E", 2, 8, (*CPU).InstructionCB93},
		0xCB94: {"RES 2,H", 2, 8, (*CPU).InstructionCB94},
		0xCB95: {"RES 2,L", 2, 8, (*CPU).InstructionCB95},
		0xCB96: {"RES 2,(HL)", 2, 16, (*CPU).InstructionCB96},
		0xCB97: {"RES 2,A", 2, 8, (*CPU).InstructionCB97},
		0xCB98: {"RES 3,B", 2, 8, (*CPU).InstructionCB98},
		0xCB99: {"RES 3,C", 2, 8, (*CPU).InstructionCB99},
		0xCB9A: {"RES 3,D", 2, 8, (*CPU).InstructionCB9A},
		0xCB9B: {"RES 3,E", 2, 8, (*CPU).InstructionCB9B},
		0xCB9C: {"RES 3,H", 2, 8, (*CPU).InstructionCB9C},
		0xCB9D: {"RES 3,L", 2, 8, (*CPU).InstructionCB9D},
		0xCB9E: {"RES 3,(HL)", 2, 16, (*CPU).InstructionCB9E},
		0xCB9F: {"RES 3,A", 2, 8, (*CPU).InstructionCB9F},

		0xCBA0: {"RES 4,B", 2, 8, (*CPU).InstructionCBA0},
		0xCBA1: {"RES 4,C", 2, 8, (*CPU).InstructionCBA1},
		0xCBA2: {"RES 4,D", 2, 8, (*CPU).InstructionCBA2},
		0xCBA3: {"RES 4,E", 2, 8, (*CPU).InstructionCBA3},
		0xCBA4: {"RES 4,H", 2, 8, (*CPU).InstructionCBA4},
		0xCBA5: {"RES 4,L", 2, 8, (*CPU).InstructionCBA5},
		0xCBA6: {"RES 4,(HL)", 2, 16, (*CPU).InstructionCBA6},
		0xCBA7: {"RES 4,A", 2, 8, (*CPU).InstructionCBA7},
		0xCBA8: {"RES 5,B", 2, 8, (*CPU).InstructionCBA8},
		0xCBA9: {"RES 5,C", 2, 8, (*CPU).InstructionCBA9},
		0xCBAA: {"RES 5,D", 2, 8, (*CPU).InstructionCBAA},
		0xCBAB: {"RES 5,E", 2, 8, (*CPU).InstructionCBAB},
		0xCBAC: {"RES 5,H", 2, 8, (*CPU).InstructionCBAC},
		0xCBAD: {"RES 5,L", 2, 8, (*CPU).InstructionCBAD},
		0xCBAE: {"RES 5,(HL)", 2, 16, (*CPU).InstructionCBAE},
		0xCBAF: {"RES 5,A", 2, 8, (*CPU).InstructionCBAF},

		0xCBB0: {"RES 6,B", 2, 8, (*CPU).InstructionCBB0},
		0xCBB1: {"RES 6,C", 2, 8, (*CPU).InstructionCBB1},
		0xCBB2: {"RES 6,D", 2, 8, (*CPU).InstructionCBB2},
		0xCBB3: {"RES 6,E", 2, 8, (*CPU).InstructionCBB3},
		0xCBB4: {"RES 6,H", 2, 8, (*CPU).InstructionCBB4},
		0xCBB5: {"RES 6,L", 2, 8, (*CPU).InstructionCBB5},
		0xCBB6: {"RES 6,(HL)", 2, 16, (*CPU).InstructionCBB6},
		0xCBB7: {"RES 6,A", 2, 8, (*CPU).InstructionCBB7},
		0xCBB8: {"RES 7,B", 2, 8, (*CPU).InstructionCBB8},
		0xCBB9: {"RES 7,C", 2, 8, (*CPU).InstructionCBB9},
		0xCBBA: {"RES 7,D", 2, 8, (*CPU).InstructionCBBA},
		0xCBBB: {"RES 7,E", 2, 8, (*CPU).InstructionCBBB},
		0xCBBC: {"RES 7,H", 2, 8, (*CPU).InstructionCBBC},
		0xCBBD: {"RES 7,L", 2, 8, (*CPU).InstructionCBBD},
		0xCBBE: {"RES 7,(HL)", 2, 16, (*CPU).InstructionCBBE},
		0xCBBF: {"RES 7,A", 2, 8, (*CPU).InstructionCBBF},

		0xCBC0: {"SET 0,B", 2, 8, (*CPU).InstructionCBC0},
		0xCBC1: {"SET 0,C", 2, 8, (*CPU).InstructionCBC1},
		0xCBC2: {"SET 0,D", 2, 8, (*CPU).InstructionCBC2},
		0xCBC3: {"SET 0,E", 2, 8, (*CPU).InstructionCBC3},
		0xCBC4: {"SET 0,H", 2, 8, (*CPU).InstructionCBC4},
		0xCBC5: {"SET 0,L", 2, 8, (*CPU).InstructionCBC5},
		0xCBC6: {"SET 0,(HL)", 2, 16, (*CPU).InstructionCBC6},
		0xCBC7: {"SET 0,A", 2, 8, (*CPU).InstructionCBC7},
		0xCBC8: {"SET 1,B", 2, 8, (*CPU).InstructionCBC8},
		0xCBC9: {"SET 1,C", 2, 8, (*CPU).InstructionCBC9},
		0xCBCA: {"SET 1,D", 2, 8, (*CPU).InstructionCBCA},
		0xCBCB: {"SET 1,E", 2, 8, (*CPU).InstructionCBCB},
		0xCBCC: {"SET 1,H", 2, 8, (*CPU).InstructionCBCC},
		0xCBCD: {"SET 1,L", 2, 8, (*CPU).InstructionCBCD},
		0xCBCE: {"SET 1,(HL)", 2, 16, (*CPU).InstructionCBCE},
		0xCBCF: {"SET 1,A", 2, 8, (*CPU).InstructionCBCF},

		0xCBD0: {"SET 2,B", 2, 8, (*CPU).InstructionCBD0},
		0xCBD1: {"SET 2,C", 2, 8, (*CPU).InstructionCBD1},
		0xCBD2: {"SET 2,D", 2, 8, (*CPU).InstructionCBD2},
		0xCBD3: {"SET 2,E", 2, 8, (*CPU).InstructionCBD3},
		0xCBD4: {"SET 2,H", 2, 8, (*CPU).InstructionCBD4},
		0xCBD5: {"SET 2,L", 2, 8, (*CPU).InstructionCBD5},
		0xCBD6: {"SET 2,(HL)", 2, 16, (*CPU).InstructionCBD6},
		0xCBD7: {"SET 2,A", 2, 8, (*CPU).InstructionCBD7},
		0xCBD8: {"SET 3,B", 2, 8, (*CPU).InstructionCBD8},
		0xCBD9: {"SET 3,C", 2, 8, (*CPU).InstructionCBD9},
		0xCBDA: {"SET 3,D", 2, 8, (*CPU).InstructionCBDA},
		0xCBDB: {"SET 3,E", 2, 8, (*CPU).InstructionCBDB},
		0xCBDC: {"SET 3,H", 2, 8, (*CPU).InstructionCBDC},
		0xCBDD: {"SET 3,L", 2, 8, (*CPU).InstructionCBDD},
		0xCBDE: {"SET 3,(HL)", 2, 16, (*CPU).InstructionCBDE},
		0xCBDF: {"SET 3,A", 2, 8, (*CPU).InstructionCBDF},

		0xCBE0: {"SET 4,B", 2, 8, (*CPU).InstructionCBE0},
		0xCBE1: {"SET 4,C", 2, 8, (*CPU).InstructionCBE1},
		0xCBE2: {"SET 4,D", 2, 8, (*CPU).InstructionCBE2},
		0xCBE3: {"SET 4,E", 2, 8, (*CPU).InstructionCBE3},
		0xCBE4: {"SET 4,H", 2, 8, (*CPU).InstructionCBE4},
		0xCBE5: {"SET 4,L", 2, 8, (*CPU).InstructionCBE5},
		0xCBE6: {"SET 4,(HL)", 2, 16, (*CPU).InstructionCBE6},
		0xCBE7: {"SET 4,A", 2, 8, (*CPU).InstructionCBE7},
		0xCBE8: {"SET 5,B", 2, 8, (*CPU).InstructionCBE8},
		0xCBE9: {"SET 5,C", 2, 8, (*CPU).InstructionCBE9},
		0xCBEA: {"SET 5,D", 2, 8, (*CPU).InstructionCBEA},
		0xCBEB: {"SET 5,E", 2, 8, (*CPU).InstructionCBEB},
		0xCBEC: {"SET 5,H", 2, 8, (*CPU).InstructionCBEC},
		0xCBED: {"SET 5,L", 2, 8, (*CPU).InstructionCBED},
		0xCBEE: {"SET 5,(HL)", 2, 16, (*CPU).InstructionCBEE},
		0xCBEF: {"SET 5,A", 2, 8, (*CPU).InstructionCBEF},

		0xCBF0: {"SET 6,B", 2, 8, (*CPU).InstructionCBF0},
		0xCBF1: {"SET 6,C", 2, 8, (*CPU).InstructionCBF1},
		0xCBF2: {"SET 6,D", 2, 8, (*CPU).InstructionCBF2},
		0xCBF3: {"SET 6,E", 2, 8, (*CPU).InstructionCBF3},
		0xCBF4: {"SET 6,H", 2, 8, (*CPU).InstructionCBF4},
		0xCBF5: {"SET 6,L", 2, 8, (*CPU).InstructionCBF5},
		0xCBF6: {"SET 6,(HL)", 2, 16, (*CPU).InstructionCBF6},
		0xCBF7: {"SET 6,A", 2, 8, (*CPU).InstructionCBF7},
		0xCBF8: {"SET 7,B", 2, 8, (*CPU).InstructionCBF8},
		0xCBF9: {"SET 7,C", 2, 8, (*CPU).InstructionCBF9},
		0xCBFA: {"SET 7,D", 2, 8, (*CPU).InstructionCBFA},
		0xCBFB: {"SET 7,E", 2, 8, (*CPU).InstructionCBFB},
		0xCBFC: {"SET 7,H", 2, 8, (*CPU).InstructionCBFC},
		0xCBFD: {"SET 7,L", 2, 8, (*CPU).InstructionCBFD},
		0xCBFE: {"SET 7,(HL)", 2, 16, (*CPU).InstructionCBFE},
		0xCBFF: {"SET 7,A", 2, 8, (*CPU).InstructionCBFF},
	}
}

func (c *CPU) Instruction00() {}

func (c *CPU) Instruction01() {
	value := c.read16BitOperand()
	c.Registers.SetBC(value)
}

func (c *CPU) Instruction02() {
	addr := c.Registers.GetBC()
	value := c.Registers.GetA()
	c.Memory.Write(addr, value)
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
	value := c.read8BitOperand()
	c.Registers.SetB(value)
}

func (c *CPU) Instruction07() {
	value := c.Registers.GetA()
	bit7 := value >> 7
	value = (value << 1) + bit7
	c.Flags.SetZ(false)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit7 == 1)
	c.Registers.SetA(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction08() {
	addr := c.read16BitOperand()
	value1 := c.Registers.GetP()
	value2 := c.Registers.GetS()
	c.Memory.Write(addr, value1)
	c.Memory.Write(addr+1, value2)
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
	value := c.Memory.Read(addr)
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
	value := c.read8BitOperand()
	c.Registers.SetC(value)
}

func (c *CPU) Instruction0F() {
	value := c.Registers.GetA()
	bit0 := value << 7
	value = (value >> 1) + bit0
	c.Flags.SetZ(false)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit0 == 128)
	c.Registers.SetA(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction10() {
}

func (c *CPU) Instruction11() {
	value := c.read16BitOperand()
	c.Registers.SetDE(value)
}

func (c *CPU) Instruction12() {
	addr := c.Registers.GetDE()
	value := c.Registers.GetA()
	c.Memory.Write(addr, value)
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
	value := c.read8BitOperand()
	c.Registers.SetD(value)
}

func (c *CPU) Instruction17() {
	carry := c.Flags.GetCarryAsValue()
	value := c.Registers.GetA()
	bit7 := value >> 7
	value = (value << 1) + carry
	c.Flags.SetZ(false)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit7 == 1)
	c.Registers.SetA(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction18() {
	value1 := int(c.Registers.GetPC())
	value2 := int(int8(c.read8BitOperand()))
	c.Registers.SetPC(uint16(value1 + value2))
}

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
	value := c.Memory.Read(addr)
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
	value := c.read8BitOperand()
	c.Registers.SetE(value)
}

func (c *CPU) Instruction1F() {
	carry := c.Flags.GetCarryAsValue() << 7
	value := c.Registers.GetA()
	bit0 := value << 7
	value = (value >> 1) + carry
	c.Flags.SetZ(false)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit0 == 128)
	c.Registers.SetA(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction20() {
	value1 := int(c.Registers.GetPC())
	value2 := int(int8(c.read8BitOperand()))
	if !c.Flags.GetZ() {
		c.Registers.SetPC(uint16(value1 + value2))
	}
}

func (c *CPU) Instruction21() {
	value := c.read16BitOperand()
	c.Registers.SetHL(value)
}

func (c *CPU) Instruction22() {
	addr := c.Registers.GetHL()
	value := c.Registers.GetA()
	c.Memory.Write(addr, value)
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
	value := c.read8BitOperand()
	c.Registers.SetH(value)
}

func (c *CPU) Instruction27() {
	value := uint16(c.Registers.GetA())
	if !c.Flags.GetN() {
		if c.Flags.GetH() || (value&0xF) > 0x9 {
			value += 0x6
		}
		if c.Flags.GetC() || value > 0x9F {
			value += 0x60
			c.Flags.SetC(true)
		}
	} else {
		if c.Flags.GetH() {
			value -= 0x6
		}
		if c.Flags.GetC() {
			value -= 0x60
		}
	}
	c.Flags.SetZ((value & 0xFF) == 0)
	c.Flags.SetH(false)
	c.Registers.SetA(uint8(value))
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction28() {
	value1 := int(c.Registers.GetPC())
	value2 := int(int8(c.read8BitOperand()))
	if c.Flags.GetZ() {
		c.Registers.SetPC(uint16(value1 + value2))
	}
}

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
	value := c.Memory.Read(addr)
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
	value := c.read8BitOperand()
	c.Registers.SetL(value)
}

func (c *CPU) Instruction2F() {
	value := c.Registers.GetA() ^ 0xFF
	c.Flags.SetN(true)
	c.Flags.SetH(true)
	c.Registers.SetA(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction30() {
	value1 := int(c.Registers.GetPC())
	value2 := int(int8(c.read8BitOperand()))
	if !c.Flags.GetC() {
		c.Registers.SetPC(uint16(value1 + value2))
	}
}

func (c *CPU) Instruction31() {
	value := c.read16BitOperand()
	c.Registers.SetSP(value)
}

func (c *CPU) Instruction32() {
	addr := c.Registers.GetHL()
	value := c.Registers.GetA()
	c.Memory.Write(addr, value)
	c.Registers.DecHL()
}

func (c *CPU) Instruction33() {
	c.Registers.IncSP()
}

func (c *CPU) Instruction34() {
	addr := c.Registers.GetHL()
	value1 := c.Memory.Read(addr)
	value2 := value1 + 1
	c.Flags.SetZ(value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(value2&0xF == 0)
	c.Memory.Write(addr, value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction35() {
	addr := c.Registers.GetHL()
	value1 := c.Memory.Read(addr)
	value2 := value1 - 1
	c.Flags.SetZ(value2 == 0)
	c.Flags.SetN(true)
	c.Flags.SetH(value1&0xF == 0)
	c.Memory.Write(addr, value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction36() {
	addr := c.Registers.GetHL()
	value := c.read8BitOperand()
	c.Memory.Write(addr, value)
}

func (c *CPU) Instruction37() {
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) Instruction38() {
	value1 := int(c.Registers.GetPC())
	value2 := int(int8(c.read8BitOperand()))
	if c.Flags.GetC() {
		c.Registers.SetPC(uint16(value1 + value2))
	}
}

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
	value := c.Memory.Read(addr)
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
	value := c.read8BitOperand()
	c.Registers.SetA(value)
}

func (c *CPU) Instruction3F() {
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(!c.Flags.GetC())
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

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
	value := c.Memory.Read(addr)
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
	value := c.Memory.Read(addr)
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
	value := c.Memory.Read(addr)
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
	value := c.Memory.Read(addr)
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
	value := c.Memory.Read(addr)
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
	value := c.Memory.Read(addr)
	c.Registers.SetL(value)
}

func (c *CPU) Instruction6F() {
	value := c.Registers.GetA()
	c.Registers.SetL(value)
}

func (c *CPU) Instruction70() {
	addr := c.Registers.GetHL()
	value := c.Registers.GetB()
	c.Memory.Write(addr, value)
}

func (c *CPU) Instruction71() {
	addr := c.Registers.GetHL()
	value := c.Registers.GetC()
	c.Memory.Write(addr, value)
}

func (c *CPU) Instruction72() {
	addr := c.Registers.GetHL()
	value := c.Registers.GetD()
	c.Memory.Write(addr, value)
}

func (c *CPU) Instruction73() {
	addr := c.Registers.GetHL()
	value := c.Registers.GetE()
	c.Memory.Write(addr, value)
}

func (c *CPU) Instruction74() {
	addr := c.Registers.GetHL()
	value := c.Registers.GetH()
	c.Memory.Write(addr, value)
}

func (c *CPU) Instruction75() {
	addr := c.Registers.GetHL()
	value := c.Registers.GetL()
	c.Memory.Write(addr, value)
}

func (c *CPU) Instruction76() {
}

func (c *CPU) Instruction77() {
	addr := c.Registers.GetHL()
	value := c.Registers.GetA()
	c.Memory.Write(addr, value)
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
	value := c.Memory.Read(addr)
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
	value2 := c.Memory.Read(addr)
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
	value2 := c.Memory.Read(addr)
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
	value2 := c.Memory.Read(addr)
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
	value2 := c.Memory.Read(addr)
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
	value2 := c.Memory.Read(addr)
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
	value2 := c.Memory.Read(addr)
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
	value2 := c.Memory.Read(addr)
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
	value2 := c.Memory.Read(addr)
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

func (c *CPU) InstructionC0() {
	if !c.Flags.GetZ() {
		addr := c.Registers.GetSP()
		msb := uint16(c.Memory.Read(addr + 1))
		lsb := uint16(c.Memory.Read(addr))
		value := msb<<8 | lsb
		c.Registers.SetPC(value)
		c.Registers.SetSP(addr + 2)
	}
}

func (c *CPU) InstructionC1() {
	addr := c.Registers.GetSP()
	msb := uint16(c.Memory.Read(addr + 1))
	lsb := uint16(c.Memory.Read(addr))
	value := msb<<8 | lsb
	c.Registers.SetBC(value)
	c.Registers.SetSP(addr + 2)
}

func (c *CPU) InstructionC2() {
	lsb := uint16(c.read8BitOperand())
	msb := uint16(c.read8BitOperand())
	value := msb<<8 | lsb
	if !c.Flags.GetZ() {
		c.Registers.SetPC(value)
	}
}

func (c *CPU) InstructionC3() {
	lsb := uint16(c.read8BitOperand())
	msb := uint16(c.read8BitOperand())
	value := msb<<8 | lsb
	c.Registers.SetPC(value)
}

func (c *CPU) InstructionC4() {
	value := c.read16BitOperand()
	if !c.Flags.GetZ() {
		addr := c.Registers.GetSP()
		msb := uint8(c.Registers.GetPC() >> 8)
		lsb := uint8(c.Registers.GetPC() & 0xFF)
		c.Memory.Write(addr-1, msb)
		c.Memory.Write(addr-2, lsb)
		c.Registers.SetSP(addr - 2)
		c.Registers.SetPC(value)
	}
}

func (c *CPU) InstructionC5() {
	addr := c.Registers.GetSP()
	msb := c.Registers.GetB()
	lsb := c.Registers.GetC()
	c.Memory.Write(addr-1, msb)
	c.Memory.Write(addr-2, lsb)
	c.Registers.SetSP(addr - 2)
}

func (c *CPU) InstructionC6() {
	value1 := c.Registers.GetA()
	value2 := c.read8BitOperand()
	c.Flags.SetZ(value1+value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH((value1&0xF)+(value2&0xF) > 0xF)
	c.Flags.SetC(uint16(value1)+uint16(value2) > 0xFF)
	c.Registers.SetA(value1 + value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionC7() {
	value := uint16(0x00)
	addr := c.Registers.GetSP()
	msb := uint8(c.Registers.GetPC() >> 8)
	lsb := uint8(c.Registers.GetPC() & 0xFF)
	c.Memory.Write(addr-1, msb)
	c.Memory.Write(addr-2, lsb)
	c.Registers.SetSP(addr - 2)
	c.Registers.SetPC(value)
}

func (c *CPU) InstructionC8() {
	if c.Flags.GetZ() {
		addr := c.Registers.GetSP()
		msb := uint16(c.Memory.Read(addr + 1))
		lsb := uint16(c.Memory.Read(addr))
		value := msb<<8 | lsb
		c.Registers.SetPC(value)
		c.Registers.SetSP(addr + 2)
	}
}

func (c *CPU) InstructionC9() {
	addr := c.Registers.GetSP()
	msb := uint16(c.Memory.Read(addr + 1))
	lsb := uint16(c.Memory.Read(addr))
	value := msb<<8 | lsb
	c.Registers.SetPC(value)
	c.Registers.SetSP(addr + 2)
}

func (c *CPU) InstructionCA() {
	lsb := uint16(c.read8BitOperand())
	msb := uint16(c.read8BitOperand())
	value := msb<<8 | lsb
	if c.Flags.GetZ() {
		c.Registers.SetPC(value)
	}
}

func (c *CPU) InstructionCB() {
}

func (c *CPU) InstructionCC() {
	value := c.read16BitOperand()
	if c.Flags.GetZ() {
		addr := c.Registers.GetSP()
		msb := uint8(c.Registers.GetPC() >> 8)
		lsb := uint8(c.Registers.GetPC() & 0xFF)
		c.Memory.Write(addr-1, msb)
		c.Memory.Write(addr-2, lsb)
		c.Registers.SetSP(addr - 2)
		c.Registers.SetPC(value)
	}
}

func (c *CPU) InstructionCD() {
	value := c.read16BitOperand()
	addr := c.Registers.GetSP()
	msb := uint8(c.Registers.GetPC() >> 8)
	lsb := uint8(c.Registers.GetPC() & 0xFF)
	c.Memory.Write(addr-1, msb)
	c.Memory.Write(addr-2, lsb)
	c.Registers.SetSP(addr - 2)
	c.Registers.SetPC(value)
}

func (c *CPU) InstructionCE() {
	carry := c.Flags.GetCarryAsValue()
	value1 := c.Registers.GetA()
	value2 := c.read8BitOperand()
	c.Flags.SetZ(value1+value2+carry == 0)
	c.Flags.SetN(false)
	c.Flags.SetH((value1&0xF)+(value2&0xF)+carry > 0xF)
	c.Flags.SetC(uint16(value1)+uint16(value2)+uint16(carry) > 0xFF)
	c.Registers.SetA(value1 + value2 + carry)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCF() {
	value := uint16(0x08)
	addr := c.Registers.GetSP()
	msb := uint8(c.Registers.GetPC() >> 8)
	lsb := uint8(c.Registers.GetPC() & 0xFF)
	c.Memory.Write(addr-1, msb)
	c.Memory.Write(addr-2, lsb)
	c.Registers.SetSP(addr - 2)
	c.Registers.SetPC(value)
}

func (c *CPU) InstructionD0() {
	if !c.Flags.GetC() {
		addr := c.Registers.GetSP()
		msb := uint16(c.Memory.Read(addr + 1))
		lsb := uint16(c.Memory.Read(addr))
		value := msb<<8 | lsb
		c.Registers.SetPC(value)
		c.Registers.SetSP(addr + 2)
	}
}

func (c *CPU) InstructionD1() {
	addr := c.Registers.GetSP()
	msb := uint16(c.Memory.Read(addr + 1))
	lsb := uint16(c.Memory.Read(addr))
	value := msb<<8 | lsb
	c.Registers.SetDE(value)
	c.Registers.SetSP(addr + 2)
}

func (c *CPU) InstructionD2() {
	lsb := uint16(c.read8BitOperand())
	msb := uint16(c.read8BitOperand())
	value := msb<<8 | lsb
	if !c.Flags.GetC() {
		c.Registers.SetPC(value)
	}
}

func (c *CPU) InstructionD3() {
}

func (c *CPU) InstructionD4() {
	value := c.read16BitOperand()
	if !c.Flags.GetC() {
		addr := c.Registers.GetSP()
		msb := uint8(c.Registers.GetPC() >> 8)
		lsb := uint8(c.Registers.GetPC() & 0xFF)
		c.Memory.Write(addr-1, msb)
		c.Memory.Write(addr-2, lsb)
		c.Registers.SetSP(addr - 2)
		c.Registers.SetPC(value)
	}
}

func (c *CPU) InstructionD5() {
	addr := c.Registers.GetSP()
	msb := c.Registers.GetD()
	lsb := c.Registers.GetE()
	c.Memory.Write(addr-1, msb)
	c.Memory.Write(addr-2, lsb)
	c.Registers.SetSP(addr - 2)
}

func (c *CPU) InstructionD6() {
	value1 := c.Registers.GetA()
	value2 := c.read8BitOperand()
	c.Flags.SetZ(value1-value2 == 0)
	c.Flags.SetN(true)
	c.Flags.SetH((value1 & 0xF) < (value2 & 0xF))
	c.Flags.SetC(uint16(value1) < uint16(value2))
	c.Registers.SetA(value1 - value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionD7() {
	value := uint16(0x10)
	addr := c.Registers.GetSP()
	msb := uint8(c.Registers.GetPC() >> 8)
	lsb := uint8(c.Registers.GetPC() & 0xFF)
	c.Memory.Write(addr-1, msb)
	c.Memory.Write(addr-2, lsb)
	c.Registers.SetSP(addr - 2)
	c.Registers.SetPC(value)
}

func (c *CPU) InstructionD8() {
	if c.Flags.GetC() {
		addr := c.Registers.GetSP()
		msb := uint16(c.Memory.Read(addr + 1))
		lsb := uint16(c.Memory.Read(addr))
		value := msb<<8 | lsb
		c.Registers.SetPC(value)
		c.Registers.SetSP(addr + 2)
	}
}

func (c *CPU) InstructionD9() {
	addr := c.Registers.GetSP()
	msb := uint16(c.Memory.Read(addr + 1))
	lsb := uint16(c.Memory.Read(addr))
	value := msb<<8 | lsb
	c.Registers.SetPC(value)
	c.Registers.SetSP(addr + 2)
	//	Exec EI
}

func (c *CPU) InstructionDA() {
	lsb := uint16(c.read8BitOperand())
	msb := uint16(c.read8BitOperand())
	value := msb<<8 | lsb
	if c.Flags.GetC() {
		c.Registers.SetPC(value)
	}
}

func (c *CPU) InstructionDB() {
}

func (c *CPU) InstructionDC() {
	value := c.read16BitOperand()
	if c.Flags.GetC() {
		addr := c.Registers.GetSP()
		msb := uint8(c.Registers.GetPC() >> 8)
		lsb := uint8(c.Registers.GetPC() & 0xFF)
		c.Memory.Write(addr-1, msb)
		c.Memory.Write(addr-2, lsb)
		c.Registers.SetSP(addr - 2)
		c.Registers.SetPC(value)
	}
}

func (c *CPU) InstructionDD() {
}

func (c *CPU) InstructionDE() {
	carry := c.Flags.GetCarryAsValue()
	value1 := c.Registers.GetA()
	value2 := c.read8BitOperand()
	c.Flags.SetZ(value1-value2-carry == 0)
	c.Flags.SetN(true)
	c.Flags.SetH((value1 & 0xF) < ((value2 & 0xF) + carry))
	c.Flags.SetC(uint16(value1) < (uint16(value2) + uint16(carry)))
	c.Registers.SetA(value1 - value2 - carry)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionDF() {
	value := uint16(0x18)
	addr := c.Registers.GetSP()
	msb := uint8(c.Registers.GetPC() >> 8)
	lsb := uint8(c.Registers.GetPC() & 0xFF)
	c.Memory.Write(addr-1, msb)
	c.Memory.Write(addr-2, lsb)
	c.Registers.SetSP(addr - 2)
	c.Registers.SetPC(value)
}

func (c *CPU) InstructionE0() {
	value := c.Registers.GetA()
	msb := uint16(0xFF)
	lsb := uint16(c.read8BitOperand())
	addr := msb<<8 | lsb
	c.Memory.Write(addr, value)
}

func (c *CPU) InstructionE1() {
	addr := c.Registers.GetSP()
	msb := uint16(c.Memory.Read(addr + 1))
	lsb := uint16(c.Memory.Read(addr))
	value := msb<<8 | lsb
	c.Registers.SetHL(value)
	c.Registers.SetSP(addr + 2)
}

func (c *CPU) InstructionE2() {
	value := c.Registers.GetA()
	msb := uint16(0xFF)
	lsb := uint16(c.Registers.GetC())
	addr := msb<<8 | lsb
	c.Memory.Write(addr, value)
}

func (c *CPU) InstructionE3() {
}

func (c *CPU) InstructionE4() {
}

func (c *CPU) InstructionE5() {
	addr := c.Registers.GetSP()
	msb := c.Registers.GetH()
	lsb := c.Registers.GetL()
	c.Memory.Write(addr-1, msb)
	c.Memory.Write(addr-2, lsb)
	c.Registers.SetSP(addr - 2)
}

func (c *CPU) InstructionE6() {
	value1 := c.Registers.GetA()
	value2 := c.read8BitOperand()
	c.Flags.SetZ(value1&value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Flags.SetC(false)
	c.Registers.SetA(value1 & value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionE7() {
	value := uint16(0x20)
	addr := c.Registers.GetSP()
	msb := uint8(c.Registers.GetPC() >> 8)
	lsb := uint8(c.Registers.GetPC() & 0xFF)
	c.Memory.Write(addr-1, msb)
	c.Memory.Write(addr-2, lsb)
	c.Registers.SetSP(addr - 2)
	c.Registers.SetPC(value)
}

func (c *CPU) InstructionE8() {
	value1 := int(c.Registers.GetSP())
	value2 := int(c.read8BitOperand())
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

func (c *CPU) InstructionE9() {
	value := c.Registers.GetHL()
	c.Registers.SetPC(value)
}

func (c *CPU) InstructionEA() {
	addr := c.read16BitOperand()
	value := c.Registers.GetA()
	c.Memory.Write(addr, value)
}

func (c *CPU) InstructionEB() {
}

func (c *CPU) InstructionEC() {
}

func (c *CPU) InstructionED() {
}

func (c *CPU) InstructionEE() {
	value1 := c.Registers.GetA()
	value2 := c.read8BitOperand()
	c.Flags.SetZ(value1^value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(false)
	c.Registers.SetA(value1 ^ value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionEF() {
	value := uint16(0x28)
	addr := c.Registers.GetSP()
	msb := uint8(c.Registers.GetPC() >> 8)
	lsb := uint8(c.Registers.GetPC() & 0xFF)
	c.Memory.Write(addr-1, msb)
	c.Memory.Write(addr-2, lsb)
	c.Registers.SetSP(addr - 2)
	c.Registers.SetPC(value)
}

func (c *CPU) InstructionF0() {
	msb := uint16(0xFF)
	lsb := uint16(c.read8BitOperand())
	addr := msb<<8 | lsb
	value := c.Memory.Read(addr)
	c.Registers.SetA(value)
}

func (c *CPU) InstructionF1() {
	addr := c.Registers.GetSP()
	msb := uint16(c.Memory.Read(addr + 1))
	lsb := uint16(c.Memory.Read(addr))
	value := msb<<8 | lsb
	c.Registers.SetAF(value)
	c.Registers.SetSP(addr + 2)
	c.Flags.SetFlagsFromValue(c.Registers.GetF())
}

func (c *CPU) InstructionF2() {
	msb := uint16(0xFF)
	lsb := uint16(c.Registers.GetC())
	addr := msb<<8 | lsb
	value := c.Memory.Read(addr)
	c.Registers.SetA(value)
}

func (c *CPU) InstructionF3() {
}

func (c *CPU) InstructionF4() {
}

func (c *CPU) InstructionF5() {
	addr := c.Registers.GetSP()
	msb := c.Registers.GetA()
	lsb := c.Registers.GetF()
	c.Memory.Write(addr-1, msb)
	c.Memory.Write(addr-2, lsb)
	c.Registers.SetSP(addr - 2)
}

func (c *CPU) InstructionF6() {
	value1 := c.Registers.GetA()
	value2 := c.read8BitOperand()
	c.Flags.SetZ(value1|value2 == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(false)
	c.Registers.SetA(value1 | value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionF7() {
	value := uint16(0x30)
	addr := c.Registers.GetSP()
	msb := uint8(c.Registers.GetPC() >> 8)
	lsb := uint8(c.Registers.GetPC() & 0xFF)
	c.Memory.Write(addr-1, msb)
	c.Memory.Write(addr-2, lsb)
	c.Registers.SetSP(addr - 2)
	c.Registers.SetPC(value)
}

func (c *CPU) InstructionF8() {
	value1 := int(c.Registers.GetSP())
	value2 := int(int8(c.read8BitOperand()))
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
	addr := c.read16BitOperand()
	value := c.Memory.Read(addr)
	c.Registers.SetA(value)
}

func (c *CPU) InstructionFB() {
}

func (c *CPU) InstructionFC() {
}

func (c *CPU) InstructionFD() {
}

func (c *CPU) InstructionFE() {
	value1 := c.Registers.GetA()
	value2 := c.read8BitOperand()
	c.Flags.SetZ(value1 == value2)
	c.Flags.SetN(true)
	c.Flags.SetH((value1 & 0xF) < (value2 & 0xF))
	c.Flags.SetC(value1 < value2)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionFF() {
	value := uint16(0x38)
	addr := c.Registers.GetSP()
	msb := uint8(c.Registers.GetPC() >> 8)
	lsb := uint8(c.Registers.GetPC() & 0xFF)
	c.Memory.Write(addr-1, msb)
	c.Memory.Write(addr-2, lsb)
	c.Registers.SetSP(addr - 2)
	c.Registers.SetPC(value)
}

func (c *CPU) InstructionCB00() {
	value := c.Registers.GetB()
	bit7 := value >> 7
	value = (value << 1) + bit7
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit7 == 1)
	c.Registers.SetB(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB01() {
	value := c.Registers.GetC()
	bit7 := value >> 7
	value = (value << 1) + bit7
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit7 == 1)
	c.Registers.SetC(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB02() {
	value := c.Registers.GetD()
	bit7 := value >> 7
	value = (value << 1) + bit7
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit7 == 1)
	c.Registers.SetD(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB03() {
	value := c.Registers.GetE()
	bit7 := value >> 7
	value = (value << 1) + bit7
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit7 == 1)
	c.Registers.SetE(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB04() {
	value := c.Registers.GetH()
	bit7 := value >> 7
	value = (value << 1) + bit7
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit7 == 1)
	c.Registers.SetH(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB05() {
	value := c.Registers.GetL()
	bit7 := value >> 7
	value = (value << 1) + bit7
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit7 == 1)
	c.Registers.SetL(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB06() {
	addr := c.Registers.GetHL()
	value := c.Memory.Read(addr)
	bit7 := value >> 7
	value = (value << 1) + bit7
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit7 == 1)
	c.Memory.Write(addr, value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB07() {
	value := c.Registers.GetA()
	bit7 := value >> 7
	value = (value << 1) + bit7
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit7 == 1)
	c.Registers.SetA(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB08() {
	value := c.Registers.GetB()
	bit0 := value << 7
	value = (value >> 1) + bit0
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit0 == 128)
	c.Registers.SetB(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB09() {
	value := c.Registers.GetC()
	bit0 := value << 7
	value = (value >> 1) + bit0
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit0 == 128)
	c.Registers.SetC(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB0A() {
	value := c.Registers.GetD()
	bit0 := value << 7
	value = (value >> 1) + bit0
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit0 == 128)
	c.Registers.SetD(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB0B() {
	value := c.Registers.GetE()
	bit0 := value << 7
	value = (value >> 1) + bit0
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit0 == 128)
	c.Registers.SetE(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB0C() {
	value := c.Registers.GetH()
	bit0 := value << 7
	value = (value >> 1) + bit0
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit0 == 128)
	c.Registers.SetH(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB0D() {
	value := c.Registers.GetL()
	bit0 := value << 7
	value = (value >> 1) + bit0
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit0 == 128)
	c.Registers.SetL(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB0E() {
	addr := c.Registers.GetHL()
	value := c.Memory.Read(addr)
	bit0 := value << 7
	value = (value >> 1) + bit0
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit0 == 128)
	c.Memory.Write(addr, value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB0F() {
	value := c.Registers.GetA()
	bit0 := value << 7
	value = (value >> 1) + bit0
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit0 == 128)
	c.Registers.SetA(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB10() {
	carry := c.Flags.GetCarryAsValue()
	value := c.Registers.GetB()
	bit7 := value >> 7
	value = (value << 1) + carry
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit7 == 1)
	c.Registers.SetB(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB11() {
	carry := c.Flags.GetCarryAsValue()
	value := c.Registers.GetC()
	bit7 := value >> 7
	value = (value << 1) + carry
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit7 == 1)
	c.Registers.SetC(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB12() {
	carry := c.Flags.GetCarryAsValue()
	value := c.Registers.GetD()
	bit7 := value >> 7
	value = (value << 1) + carry
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit7 == 1)
	c.Registers.SetD(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB13() {
	carry := c.Flags.GetCarryAsValue()
	value := c.Registers.GetE()
	bit7 := value >> 7
	value = (value << 1) + carry
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit7 == 1)
	c.Registers.SetE(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB14() {
	carry := c.Flags.GetCarryAsValue()
	value := c.Registers.GetH()
	bit7 := value >> 7
	value = (value << 1) + carry
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit7 == 1)
	c.Registers.SetH(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB15() {
	carry := c.Flags.GetCarryAsValue()
	value := c.Registers.GetL()
	bit7 := value >> 7
	value = (value << 1) + carry
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit7 == 1)
	c.Registers.SetL(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB16() {
	addr := c.Registers.GetHL()
	carry := c.Flags.GetCarryAsValue()
	value := c.Memory.Read(addr)
	bit7 := value >> 7
	value = (value << 1) + carry
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit7 == 1)
	c.Memory.Write(addr, value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB17() {
	carry := c.Flags.GetCarryAsValue()
	value := c.Registers.GetA()
	bit7 := value >> 7
	value = (value << 1) + carry
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit7 == 1)
	c.Registers.SetA(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB18() {
	carry := c.Flags.GetCarryAsValue() << 7
	value := c.Registers.GetB()
	bit0 := value << 7
	value = (value >> 1) + carry
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit0 == 128)
	c.Registers.SetB(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB19() {
	carry := c.Flags.GetCarryAsValue() << 7
	value := c.Registers.GetC()
	bit0 := value << 7
	value = (value >> 1) + carry
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit0 == 128)
	c.Registers.SetC(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB1A() {
	carry := c.Flags.GetCarryAsValue() << 7
	value := c.Registers.GetD()
	bit0 := value << 7
	value = (value >> 1) + carry
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit0 == 128)
	c.Registers.SetD(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB1B() {
	carry := c.Flags.GetCarryAsValue() << 7
	value := c.Registers.GetE()
	bit0 := value << 7
	value = (value >> 1) + carry
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit0 == 128)
	c.Registers.SetE(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB1C() {
	carry := c.Flags.GetCarryAsValue() << 7
	value := c.Registers.GetH()
	bit0 := value << 7
	value = (value >> 1) + carry
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit0 == 128)
	c.Registers.SetH(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB1D() {
	carry := c.Flags.GetCarryAsValue() << 7
	value := c.Registers.GetL()
	bit0 := value << 7
	value = (value >> 1) + carry
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit0 == 128)
	c.Registers.SetL(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB1E() {
	addr := c.Registers.GetHL()
	carry := c.Flags.GetCarryAsValue() << 7
	value := c.Memory.Read(addr)
	bit7 := value << 7
	value = (value >> 1) + carry
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit7 == 128)
	c.Memory.Write(addr, value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB1F() {
	carry := c.Flags.GetCarryAsValue() << 7
	value := c.Registers.GetA()
	bit0 := value << 7
	value = (value >> 1) + carry
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit0 == 128)
	c.Registers.SetA(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB20() {
	value := c.Registers.GetB()
	bit7 := value >> 7
	value = value << 1
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit7 == 1)
	c.Registers.SetB(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB21() {
	value := c.Registers.GetC()
	bit7 := value >> 7
	value = value << 1
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit7 == 1)
	c.Registers.SetC(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB22() {
	value := c.Registers.GetD()
	bit7 := value >> 7
	value = value << 1
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit7 == 1)
	c.Registers.SetD(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB23() {
	value := c.Registers.GetE()
	bit7 := value >> 7
	value = value << 1
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit7 == 1)
	c.Registers.SetE(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB24() {
	value := c.Registers.GetH()
	bit7 := value >> 7
	value = value << 1
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit7 == 1)
	c.Registers.SetH(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB25() {
	value := c.Registers.GetL()
	bit7 := value >> 7
	value = value << 1
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit7 == 1)
	c.Registers.SetL(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB26() {
	addr := c.Registers.GetHL()
	value := c.Memory.Read(addr)
	bit7 := value >> 7
	value = value << 1
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit7 == 1)
	c.Memory.Write(addr, value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB27() {
	value := c.Registers.GetA()
	bit7 := value >> 7
	value = value << 1
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit7 == 1)
	c.Registers.SetA(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB28() {
	value := c.Registers.GetB()
	bit0 := value << 7
	bit7 := value & (1 << 7)
	value = (value >> 1) + bit7
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit0 == 128)
	c.Registers.SetB(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB29() {
	value := c.Registers.GetC()
	bit0 := value << 7
	bit7 := value & (1 << 7)
	value = (value >> 1) + bit7
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit0 == 128)
	c.Registers.SetC(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB2A() {
	value := c.Registers.GetD()
	bit0 := value << 7
	bit7 := value & (1 << 7)
	value = (value >> 1) + bit7
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit0 == 128)
	c.Registers.SetD(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB2B() {
	value := c.Registers.GetE()
	bit0 := value << 7
	bit7 := value & (1 << 7)
	value = (value >> 1) + bit7
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit0 == 128)
	c.Registers.SetE(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB2C() {
	value := c.Registers.GetH()
	bit0 := value << 7
	bit7 := value & (1 << 7)
	value = (value >> 1) + bit7
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit0 == 128)
	c.Registers.SetH(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB2D() {
	value := c.Registers.GetL()
	bit0 := value << 7
	bit7 := value & (1 << 7)
	value = (value >> 1) + bit7
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit0 == 128)
	c.Registers.SetL(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB2E() {
	addr := c.Registers.GetHL()
	value := c.Memory.Read(addr)
	bit0 := value << 7
	bit7 := value & (1 << 7)
	value = (value >> 1) + bit7
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit0 == 128)
	c.Memory.Write(addr, value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB2F() {
	value := c.Registers.GetA()
	bit0 := value << 7
	bit7 := value & (1 << 7)
	value = (value >> 1) + bit7
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit0 == 128)
	c.Registers.SetA(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB30() {
	value := c.Registers.GetB()
	lo := value & 0xF
	hi := value >> 4
	value = lo<<4 | hi
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(false)
	c.Registers.SetB(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB31() {
	value := c.Registers.GetC()
	lo := value & 0xF
	hi := value >> 4
	value = lo<<4 | hi
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(false)
	c.Registers.SetC(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB32() {
	value := c.Registers.GetD()
	lo := value & 0xF
	hi := value >> 4
	value = lo<<4 | hi
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(false)
	c.Registers.SetD(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB33() {
	value := c.Registers.GetE()
	lo := value & 0xF
	hi := value >> 4
	value = lo<<4 | hi
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(false)
	c.Registers.SetE(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB34() {
	value := c.Registers.GetH()
	lo := value & 0xF
	hi := value >> 4
	value = lo<<4 | hi
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(false)
	c.Registers.SetH(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB35() {
	value := c.Registers.GetL()
	lo := value & 0xF
	hi := value >> 4
	value = lo<<4 | hi
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(false)
	c.Registers.SetL(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB36() {
	addr := c.Registers.GetHL()
	value := c.Memory.Read(addr)
	lo := value & 0xF
	hi := value >> 4
	value = lo<<4 | hi
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(false)
	c.Memory.Write(addr, value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB37() {
	value := c.Registers.GetA()
	lo := value & 0xF
	hi := value >> 4
	value = lo<<4 | hi
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(false)
	c.Registers.SetA(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB38() {
	value := c.Registers.GetB()
	bit0 := value << 7
	value = value >> 1
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit0 == 128)
	c.Registers.SetB(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB39() {
	value := c.Registers.GetC()
	bit0 := value << 7
	value = value >> 1
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit0 == 128)
	c.Registers.SetC(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB3A() {
	value := c.Registers.GetD()
	bit0 := value << 7
	value = value >> 1
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit0 == 128)
	c.Registers.SetD(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB3B() {
	value := c.Registers.GetE()
	bit0 := value << 7
	value = value >> 1
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit0 == 128)
	c.Registers.SetE(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB3C() {
	value := c.Registers.GetH()
	bit0 := value << 7
	value = value >> 1
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit0 == 128)
	c.Registers.SetH(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB3D() {
	value := c.Registers.GetL()
	bit0 := value << 7
	value = value >> 1
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit0 == 128)
	c.Registers.SetL(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB3E() {
	addr := c.Registers.GetHL()
	value := c.Memory.Read(addr)
	bit0 := value << 7
	value = value >> 1
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit0 == 128)
	c.Memory.Write(addr, value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB3F() {
	value := c.Registers.GetA()
	bit0 := value << 7
	value = value >> 1
	c.Flags.SetZ(value == 0)
	c.Flags.SetN(false)
	c.Flags.SetH(false)
	c.Flags.SetC(bit0 == 128)
	c.Registers.SetA(value)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB40() {
	value := c.Registers.GetB()
	bit := (value & (1 << 0)) == (1 << 0)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB41() {
	value := c.Registers.GetC()
	bit := (value & (1 << 0)) == (1 << 0)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB42() {
	value := c.Registers.GetD()
	bit := (value & (1 << 0)) == (1 << 0)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB43() {
	value := c.Registers.GetE()
	bit := (value & (1 << 0)) == (1 << 0)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB44() {
	value := c.Registers.GetH()
	bit := (value & (1 << 0)) == (1 << 0)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB45() {
	value := c.Registers.GetL()
	bit := (value & (1 << 0)) == (1 << 0)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB46() {
	addr := c.Registers.GetHL()
	value := c.Memory.Read(addr)
	bit := (value & (1 << 0)) == (1 << 0)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB47() {
	value := c.Registers.GetA()
	bit := (value & (1 << 0)) == (1 << 0)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB48() {
	value := c.Registers.GetB()
	bit := (value & (1 << 1)) == (1 << 1)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB49() {
	value := c.Registers.GetC()
	bit := (value & (1 << 1)) == (1 << 1)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB4A() {
	value := c.Registers.GetD()
	bit := (value & (1 << 1)) == (1 << 1)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB4B() {
	value := c.Registers.GetE()
	bit := (value & (1 << 1)) == (1 << 1)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB4C() {
	value := c.Registers.GetH()
	bit := (value & (1 << 1)) == (1 << 1)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB4D() {
	value := c.Registers.GetL()
	bit := (value & (1 << 1)) == (1 << 1)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB4E() {
	addr := c.Registers.GetHL()
	value := c.Memory.Read(addr)
	bit := (value & (1 << 1)) == (1 << 1)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB4F() {
	value := c.Registers.GetA()
	bit := (value & (1 << 1)) == (1 << 1)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB50() {
	value := c.Registers.GetB()
	bit := (value & (1 << 2)) == (1 << 2)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB51() {
	value := c.Registers.GetC()
	bit := (value & (1 << 2)) == (1 << 2)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB52() {
	value := c.Registers.GetD()
	bit := (value & (1 << 2)) == (1 << 2)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB53() {
	value := c.Registers.GetE()
	bit := (value & (1 << 2)) == (1 << 2)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB54() {
	value := c.Registers.GetH()
	bit := (value & (1 << 2)) == (1 << 2)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB55() {
	value := c.Registers.GetL()
	bit := (value & (1 << 2)) == (1 << 2)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB56() {
	addr := c.Registers.GetHL()
	value := c.Memory.Read(addr)
	bit := (value & (1 << 2)) == (1 << 2)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB57() {
	value := c.Registers.GetA()
	bit := (value & (1 << 2)) == (1 << 2)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB58() {
	value := c.Registers.GetB()
	bit := (value & (1 << 3)) == (1 << 3)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB59() {
	value := c.Registers.GetC()
	bit := (value & (1 << 3)) == (1 << 3)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB5A() {
	value := c.Registers.GetD()
	bit := (value & (1 << 3)) == (1 << 3)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB5B() {
	value := c.Registers.GetE()
	bit := (value & (1 << 3)) == (1 << 3)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB5C() {
	value := c.Registers.GetH()
	bit := (value & (1 << 3)) == (1 << 3)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB5D() {
	value := c.Registers.GetL()
	bit := (value & (1 << 3)) == (1 << 3)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB5E() {
	addr := c.Registers.GetHL()
	value := c.Memory.Read(addr)
	bit := (value & (1 << 3)) == (1 << 3)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB5F() {
	value := c.Registers.GetA()
	bit := (value & (1 << 3)) == (1 << 3)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB60() {
	value := c.Registers.GetB()
	bit := (value & (1 << 4)) == (1 << 4)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB61() {
	value := c.Registers.GetC()
	bit := (value & (1 << 4)) == (1 << 4)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB62() {
	value := c.Registers.GetD()
	bit := (value & (1 << 4)) == (1 << 4)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB63() {
	value := c.Registers.GetE()
	bit := (value & (1 << 4)) == (1 << 4)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB64() {
	value := c.Registers.GetH()
	bit := (value & (1 << 4)) == (1 << 4)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB65() {
	value := c.Registers.GetL()
	bit := (value & (1 << 4)) == (1 << 4)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB66() {
	addr := c.Registers.GetHL()
	value := c.Memory.Read(addr)
	bit := (value & (1 << 4)) == (1 << 4)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB67() {
	value := c.Registers.GetA()
	bit := (value & (1 << 4)) == (1 << 4)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB68() {
	value := c.Registers.GetB()
	bit := (value & (1 << 5)) == (1 << 5)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB69() {
	value := c.Registers.GetC()
	bit := (value & (1 << 5)) == (1 << 5)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB6A() {
	value := c.Registers.GetD()
	bit := (value & (1 << 5)) == (1 << 5)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB6B() {
	value := c.Registers.GetE()
	bit := (value & (1 << 5)) == (1 << 5)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB6C() {
	value := c.Registers.GetH()
	bit := (value & (1 << 5)) == (1 << 5)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB6D() {
	value := c.Registers.GetL()
	bit := (value & (1 << 5)) == (1 << 5)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB6E() {
	addr := c.Registers.GetHL()
	value := c.Memory.Read(addr)
	bit := (value & (1 << 5)) == (1 << 5)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB6F() {
	value := c.Registers.GetA()
	bit := (value & (1 << 5)) == (1 << 5)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB70() {
	value := c.Registers.GetB()
	bit := (value & (1 << 6)) == (1 << 6)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB71() {
	value := c.Registers.GetC()
	bit := (value & (1 << 6)) == (1 << 6)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB72() {
	value := c.Registers.GetD()
	bit := (value & (1 << 6)) == (1 << 6)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB73() {
	value := c.Registers.GetE()
	bit := (value & (1 << 6)) == (1 << 6)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB74() {
	value := c.Registers.GetH()
	bit := (value & (1 << 6)) == (1 << 6)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB75() {
	value := c.Registers.GetL()
	bit := (value & (1 << 6)) == (1 << 6)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB76() {
	addr := c.Registers.GetHL()
	value := c.Memory.Read(addr)
	bit := (value & (1 << 6)) == (1 << 6)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB77() {
	value := c.Registers.GetA()
	bit := (value & (1 << 6)) == (1 << 6)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB78() {
	value := c.Registers.GetB()
	bit := (value & (1 << 7)) == (1 << 7)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB79() {
	value := c.Registers.GetC()
	bit := (value & (1 << 7)) == (1 << 7)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB7A() {
	value := c.Registers.GetD()
	bit := (value & (1 << 7)) == (1 << 7)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB7B() {
	value := c.Registers.GetE()
	bit := (value & (1 << 7)) == (1 << 7)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB7C() {
	value := c.Registers.GetH()
	bit := (value & (1 << 7)) == (1 << 7)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB7D() {
	value := c.Registers.GetL()
	bit := (value & (1 << 7)) == (1 << 7)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB7E() {
	addr := c.Registers.GetHL()
	value := c.Memory.Read(addr)
	bit := (value & (1 << 7)) == (1 << 7)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB7F() {
	value := c.Registers.GetA()
	bit := (value & (1 << 7)) == (1 << 7)
	c.Flags.SetZ(!bit)
	c.Flags.SetN(false)
	c.Flags.SetH(true)
	c.Registers.SetF(c.Flags.GetFlagsAsValue() | (c.Registers.GetF() & 0xF))
}

func (c *CPU) InstructionCB80() {
	value := c.Registers.GetB() & (0xFF - (1 << 0))
	c.Registers.SetB(value)
}

func (c *CPU) InstructionCB81() {
	value := c.Registers.GetC() & (0xFF - (1 << 0))
	c.Registers.SetC(value)
}

func (c *CPU) InstructionCB82() {
	value := c.Registers.GetD() & (0xFF - (1 << 0))
	c.Registers.SetD(value)
}

func (c *CPU) InstructionCB83() {
	value := c.Registers.GetE() & (0xFF - (1 << 0))
	c.Registers.SetE(value)
}

func (c *CPU) InstructionCB84() {
	value := c.Registers.GetH() & (0xFF - (1 << 0))
	c.Registers.SetH(value)
}

func (c *CPU) InstructionCB85() {
	value := c.Registers.GetL() & (0xFF - (1 << 0))
	c.Registers.SetL(value)
}

func (c *CPU) InstructionCB86() {
	addr := c.Registers.GetHL()
	value := c.Memory.Read(addr) & (0xFF - (1 << 0))
	c.Memory.Write(addr, value)
}

func (c *CPU) InstructionCB87() {
	value := c.Registers.GetA() & (0xFF - (1 << 0))
	c.Registers.SetA(value)
}

func (c *CPU) InstructionCB88() {
	value := c.Registers.GetB() & (0xFF - (1 << 1))
	c.Registers.SetB(value)
}

func (c *CPU) InstructionCB89() {
	value := c.Registers.GetC() & (0xFF - (1 << 1))
	c.Registers.SetC(value)
}

func (c *CPU) InstructionCB8A() {
	value := c.Registers.GetD() & (0xFF - (1 << 1))
	c.Registers.SetD(value)
}

func (c *CPU) InstructionCB8B() {
	value := c.Registers.GetE() & (0xFF - (1 << 1))
	c.Registers.SetE(value)
}

func (c *CPU) InstructionCB8C() {
	value := c.Registers.GetH() & (0xFF - (1 << 1))
	c.Registers.SetH(value)
}

func (c *CPU) InstructionCB8D() {
	value := c.Registers.GetL() & (0xFF - (1 << 1))
	c.Registers.SetL(value)
}

func (c *CPU) InstructionCB8E() {
	addr := c.Registers.GetHL()
	value := c.Memory.Read(addr) & (0xFF - (1 << 1))
	c.Memory.Write(addr, value)
}

func (c *CPU) InstructionCB8F() {
	value := c.Registers.GetA() & (0xFF - (1 << 1))
	c.Registers.SetA(value)
}

func (c *CPU) InstructionCB90() {
	value := c.Registers.GetB() & (0xFF - (1 << 2))
	c.Registers.SetB(value)
}

func (c *CPU) InstructionCB91() {
	value := c.Registers.GetC() & (0xFF - (1 << 2))
	c.Registers.SetC(value)
}

func (c *CPU) InstructionCB92() {
	value := c.Registers.GetD() & (0xFF - (1 << 2))
	c.Registers.SetD(value)
}

func (c *CPU) InstructionCB93() {
	value := c.Registers.GetE() & (0xFF - (1 << 2))
	c.Registers.SetE(value)
}

func (c *CPU) InstructionCB94() {
	value := c.Registers.GetH() & (0xFF - (1 << 2))
	c.Registers.SetH(value)
}

func (c *CPU) InstructionCB95() {
	value := c.Registers.GetL() & (0xFF - (1 << 2))
	c.Registers.SetL(value)
}

func (c *CPU) InstructionCB96() {
	addr := c.Registers.GetHL()
	value := c.Memory.Read(addr) & (0xFF - (1 << 2))
	c.Memory.Write(addr, value)
}

func (c *CPU) InstructionCB97() {
	value := c.Registers.GetA() & (0xFF - (1 << 2))
	c.Registers.SetA(value)
}

func (c *CPU) InstructionCB98() {
	value := c.Registers.GetB() & (0xFF - (1 << 3))
	c.Registers.SetB(value)
}

func (c *CPU) InstructionCB99() {
	value := c.Registers.GetC() & (0xFF - (1 << 3))
	c.Registers.SetC(value)
}

func (c *CPU) InstructionCB9A() {
	value := c.Registers.GetD() & (0xFF - (1 << 3))
	c.Registers.SetD(value)
}

func (c *CPU) InstructionCB9B() {
	value := c.Registers.GetE() & (0xFF - (1 << 3))
	c.Registers.SetE(value)
}

func (c *CPU) InstructionCB9C() {
	value := c.Registers.GetH() & (0xFF - (1 << 3))
	c.Registers.SetH(value)
}

func (c *CPU) InstructionCB9D() {
	value := c.Registers.GetL() & (0xFF - (1 << 3))
	c.Registers.SetL(value)
}

func (c *CPU) InstructionCB9E() {
	addr := c.Registers.GetHL()
	value := c.Memory.Read(addr) & (0xFF - (1 << 3))
	c.Memory.Write(addr, value)
}

func (c *CPU) InstructionCB9F() {
	value := c.Registers.GetA() & (0xFF - (1 << 3))
	c.Registers.SetA(value)
}

func (c *CPU) InstructionCBA0() {
	value := c.Registers.GetB() & (0xFF - (1 << 4))
	c.Registers.SetB(value)
}

func (c *CPU) InstructionCBA1() {
	value := c.Registers.GetC() & (0xFF - (1 << 4))
	c.Registers.SetC(value)
}

func (c *CPU) InstructionCBA2() {
	value := c.Registers.GetD() & (0xFF - (1 << 4))
	c.Registers.SetD(value)
}

func (c *CPU) InstructionCBA3() {
	value := c.Registers.GetE() & (0xFF - (1 << 4))
	c.Registers.SetE(value)
}

func (c *CPU) InstructionCBA4() {
	value := c.Registers.GetH() & (0xFF - (1 << 4))
	c.Registers.SetH(value)
}

func (c *CPU) InstructionCBA5() {
	value := c.Registers.GetL() & (0xFF - (1 << 4))
	c.Registers.SetL(value)
}

func (c *CPU) InstructionCBA6() {
	addr := c.Registers.GetHL()
	value := c.Memory.Read(addr) & (0xFF - (1 << 4))
	c.Memory.Write(addr, value)
}

func (c *CPU) InstructionCBA7() {
	value := c.Registers.GetA() & (0xFF - (1 << 4))
	c.Registers.SetA(value)
}

func (c *CPU) InstructionCBA8() {
	value := c.Registers.GetB() & (0xFF - (1 << 5))
	c.Registers.SetB(value)
}

func (c *CPU) InstructionCBA9() {
	value := c.Registers.GetC() & (0xFF - (1 << 5))
	c.Registers.SetC(value)
}

func (c *CPU) InstructionCBAA() {
	value := c.Registers.GetD() & (0xFF - (1 << 5))
	c.Registers.SetD(value)
}

func (c *CPU) InstructionCBAB() {
	value := c.Registers.GetE() & (0xFF - (1 << 5))
	c.Registers.SetE(value)
}

func (c *CPU) InstructionCBAC() {
	value := c.Registers.GetH() & (0xFF - (1 << 5))
	c.Registers.SetH(value)
}

func (c *CPU) InstructionCBAD() {
	value := c.Registers.GetL() & (0xFF - (1 << 5))
	c.Registers.SetL(value)
}

func (c *CPU) InstructionCBAE() {
	addr := c.Registers.GetHL()
	value := c.Memory.Read(addr) & (0xFF - (1 << 5))
	c.Memory.Write(addr, value)
}

func (c *CPU) InstructionCBAF() {
	value := c.Registers.GetA() & (0xFF - (1 << 5))
	c.Registers.SetA(value)
}

func (c *CPU) InstructionCBB0() {
	value := c.Registers.GetB() & (0xFF - (1 << 6))
	c.Registers.SetB(value)
}

func (c *CPU) InstructionCBB1() {
	value := c.Registers.GetC() & (0xFF - (1 << 6))
	c.Registers.SetC(value)
}

func (c *CPU) InstructionCBB2() {
	value := c.Registers.GetD() & (0xFF - (1 << 6))
	c.Registers.SetD(value)
}

func (c *CPU) InstructionCBB3() {
	value := c.Registers.GetE() & (0xFF - (1 << 6))
	c.Registers.SetE(value)
}

func (c *CPU) InstructionCBB4() {
	value := c.Registers.GetH() & (0xFF - (1 << 6))
	c.Registers.SetH(value)
}

func (c *CPU) InstructionCBB5() {
	value := c.Registers.GetL() & (0xFF - (1 << 6))
	c.Registers.SetL(value)
}

func (c *CPU) InstructionCBB6() {
	addr := c.Registers.GetHL()
	value := c.Memory.Read(addr) & (0xFF - (1 << 6))
	c.Memory.Write(addr, value)
}

func (c *CPU) InstructionCBB7() {
	value := c.Registers.GetA() & (0xFF - (1 << 6))
	c.Registers.SetA(value)
}

func (c *CPU) InstructionCBB8() {
	value := c.Registers.GetB() & (0xFF - (1 << 7))
	c.Registers.SetB(value)
}

func (c *CPU) InstructionCBB9() {
	value := c.Registers.GetC() & (0xFF - (1 << 7))
	c.Registers.SetC(value)
}

func (c *CPU) InstructionCBBA() {
	value := c.Registers.GetD() & (0xFF - (1 << 7))
	c.Registers.SetD(value)
}

func (c *CPU) InstructionCBBB() {
	value := c.Registers.GetE() & (0xFF - (1 << 7))
	c.Registers.SetE(value)
}

func (c *CPU) InstructionCBBC() {
	value := c.Registers.GetH() & (0xFF - (1 << 7))
	c.Registers.SetH(value)
}

func (c *CPU) InstructionCBBD() {
	value := c.Registers.GetL() & (0xFF - (1 << 7))
	c.Registers.SetL(value)
}

func (c *CPU) InstructionCBBE() {
	addr := c.Registers.GetHL()
	value := c.Memory.Read(addr) & (0xFF - (1 << 7))
	c.Memory.Write(addr, value)
}

func (c *CPU) InstructionCBBF() {
	value := c.Registers.GetA() & (0xFF - (1 << 7))
	c.Registers.SetA(value)
}

func (c *CPU) InstructionCBC0() {
	value := c.Registers.GetB() | (1 << 0)
	c.Registers.SetB(value)
}

func (c *CPU) InstructionCBC1() {
	value := c.Registers.GetC() | (1 << 0)
	c.Registers.SetC(value)
}

func (c *CPU) InstructionCBC2() {
	value := c.Registers.GetD() | (1 << 0)
	c.Registers.SetD(value)
}

func (c *CPU) InstructionCBC3() {
	value := c.Registers.GetE() | (1 << 0)
	c.Registers.SetE(value)
}

func (c *CPU) InstructionCBC4() {
	value := c.Registers.GetH() | (1 << 0)
	c.Registers.SetH(value)
}

func (c *CPU) InstructionCBC5() {
	value := c.Registers.GetL() | (1 << 0)
	c.Registers.SetL(value)
}

func (c *CPU) InstructionCBC6() {
	addr := c.Registers.GetHL()
	value := c.Memory.Read(addr) | (1 << 0)
	c.Memory.Write(addr, value)
}

func (c *CPU) InstructionCBC7() {
	value := c.Registers.GetA() | (1 << 0)
	c.Registers.SetA(value)
}

func (c *CPU) InstructionCBC8() {
	value := c.Registers.GetB() | (1 << 1)
	c.Registers.SetB(value)
}

func (c *CPU) InstructionCBC9() {
	value := c.Registers.GetC() | (1 << 1)
	c.Registers.SetC(value)
}

func (c *CPU) InstructionCBCA() {
	value := c.Registers.GetD() | (1 << 1)
	c.Registers.SetD(value)
}

func (c *CPU) InstructionCBCB() {
	value := c.Registers.GetE() | (1 << 1)
	c.Registers.SetE(value)
}

func (c *CPU) InstructionCBCC() {
	value := c.Registers.GetH() | (1 << 1)
	c.Registers.SetH(value)
}

func (c *CPU) InstructionCBCD() {
	value := c.Registers.GetL() | (1 << 1)
	c.Registers.SetL(value)
}

func (c *CPU) InstructionCBCE() {
	addr := c.Registers.GetHL()
	value := c.Memory.Read(addr) | (1 << 1)
	c.Memory.Write(addr, value)
}

func (c *CPU) InstructionCBCF() {
	value := c.Registers.GetA() | (1 << 1)
	c.Registers.SetA(value)
}

func (c *CPU) InstructionCBD0() {
	value := c.Registers.GetB() | (1 << 2)
	c.Registers.SetB(value)
}

func (c *CPU) InstructionCBD1() {
	value := c.Registers.GetC() | (1 << 2)
	c.Registers.SetC(value)
}

func (c *CPU) InstructionCBD2() {
	value := c.Registers.GetD() | (1 << 2)
	c.Registers.SetD(value)
}

func (c *CPU) InstructionCBD3() {
	value := c.Registers.GetE() | (1 << 2)
	c.Registers.SetE(value)
}

func (c *CPU) InstructionCBD4() {
	value := c.Registers.GetH() | (1 << 2)
	c.Registers.SetH(value)
}

func (c *CPU) InstructionCBD5() {
	value := c.Registers.GetL() | (1 << 2)
	c.Registers.SetL(value)
}

func (c *CPU) InstructionCBD6() {
	addr := c.Registers.GetHL()
	value := c.Memory.Read(addr) | (1 << 2)
	c.Memory.Write(addr, value)
}

func (c *CPU) InstructionCBD7() {
	value := c.Registers.GetA() | (1 << 2)
	c.Registers.SetA(value)
}

func (c *CPU) InstructionCBD8() {
	value := c.Registers.GetB() | (1 << 3)
	c.Registers.SetB(value)
}

func (c *CPU) InstructionCBD9() {
	value := c.Registers.GetC() | (1 << 3)
	c.Registers.SetC(value)
}

func (c *CPU) InstructionCBDA() {
	value := c.Registers.GetD() | (1 << 3)
	c.Registers.SetD(value)
}

func (c *CPU) InstructionCBDB() {
	value := c.Registers.GetE() | (1 << 3)
	c.Registers.SetE(value)
}

func (c *CPU) InstructionCBDC() {
	value := c.Registers.GetH() | (1 << 3)
	c.Registers.SetH(value)
}

func (c *CPU) InstructionCBDD() {
	value := c.Registers.GetL() | (1 << 3)
	c.Registers.SetL(value)
}

func (c *CPU) InstructionCBDE() {
	addr := c.Registers.GetHL()
	value := c.Memory.Read(addr) | (1 << 3)
	c.Memory.Write(addr, value)
}

func (c *CPU) InstructionCBDF() {
	value := c.Registers.GetA() | (1 << 3)
	c.Registers.SetA(value)
}

func (c *CPU) InstructionCBE0() {
	value := c.Registers.GetB() | (1 << 4)
	c.Registers.SetB(value)
}

func (c *CPU) InstructionCBE1() {
	value := c.Registers.GetC() | (1 << 4)
	c.Registers.SetC(value)
}

func (c *CPU) InstructionCBE2() {
	value := c.Registers.GetD() | (1 << 4)
	c.Registers.SetD(value)
}

func (c *CPU) InstructionCBE3() {
	value := c.Registers.GetE() | (1 << 4)
	c.Registers.SetE(value)
}

func (c *CPU) InstructionCBE4() {
	value := c.Registers.GetH() | (1 << 4)
	c.Registers.SetH(value)
}

func (c *CPU) InstructionCBE5() {
	value := c.Registers.GetL() | (1 << 4)
	c.Registers.SetL(value)
}

func (c *CPU) InstructionCBE6() {
	addr := c.Registers.GetHL()
	value := c.Memory.Read(addr) | (1 << 4)
	c.Memory.Write(addr, value)
}

func (c *CPU) InstructionCBE7() {
	value := c.Registers.GetA() | (1 << 4)
	c.Registers.SetA(value)
}

func (c *CPU) InstructionCBE8() {
	value := c.Registers.GetB() | (1 << 5)
	c.Registers.SetB(value)
}

func (c *CPU) InstructionCBE9() {
	value := c.Registers.GetC() | (1 << 5)
	c.Registers.SetC(value)
}

func (c *CPU) InstructionCBEA() {
	value := c.Registers.GetD() | (1 << 5)
	c.Registers.SetD(value)
}

func (c *CPU) InstructionCBEB() {
	value := c.Registers.GetE() | (1 << 5)
	c.Registers.SetE(value)
}

func (c *CPU) InstructionCBEC() {
	value := c.Registers.GetH() | (1 << 5)
	c.Registers.SetH(value)
}

func (c *CPU) InstructionCBED() {
	value := c.Registers.GetL() | (1 << 5)
	c.Registers.SetL(value)
}

func (c *CPU) InstructionCBEE() {
	addr := c.Registers.GetHL()
	value := c.Memory.Read(addr) | (1 << 5)
	c.Memory.Write(addr, value)
}

func (c *CPU) InstructionCBEF() {
	value := c.Registers.GetA() | (1 << 5)
	c.Registers.SetA(value)
}

func (c *CPU) InstructionCBF0() {
	value := c.Registers.GetB() | (1 << 6)
	c.Registers.SetB(value)
}

func (c *CPU) InstructionCBF1() {
	value := c.Registers.GetC() | (1 << 6)
	c.Registers.SetC(value)
}

func (c *CPU) InstructionCBF2() {
	value := c.Registers.GetD() | (1 << 6)
	c.Registers.SetD(value)
}

func (c *CPU) InstructionCBF3() {
	value := c.Registers.GetE() | (1 << 6)
	c.Registers.SetE(value)
}

func (c *CPU) InstructionCBF4() {
	value := c.Registers.GetH() | (1 << 6)
	c.Registers.SetH(value)
}

func (c *CPU) InstructionCBF5() {
	value := c.Registers.GetL() | (1 << 6)
	c.Registers.SetL(value)
}

func (c *CPU) InstructionCBF6() {
	addr := c.Registers.GetHL()
	value := c.Memory.Read(addr) | (1 << 6)
	c.Memory.Write(addr, value)
}

func (c *CPU) InstructionCBF7() {
	value := c.Registers.GetA() | (1 << 6)
	c.Registers.SetA(value)
}

func (c *CPU) InstructionCBF8() {
	value := c.Registers.GetB() | (1 << 7)
	c.Registers.SetB(value)
}

func (c *CPU) InstructionCBF9() {
	value := c.Registers.GetC() | (1 << 7)
	c.Registers.SetC(value)
}

func (c *CPU) InstructionCBFA() {
	value := c.Registers.GetD() | (1 << 7)
	c.Registers.SetD(value)
}

func (c *CPU) InstructionCBFB() {
	value := c.Registers.GetE() | (1 << 7)
	c.Registers.SetE(value)
}

func (c *CPU) InstructionCBFC() {
	value := c.Registers.GetH() | (1 << 7)
	c.Registers.SetH(value)
}

func (c *CPU) InstructionCBFD() {
	value := c.Registers.GetL() | (1 << 7)
	c.Registers.SetL(value)
}

func (c *CPU) InstructionCBFE() {
	addr := c.Registers.GetHL()
	value := c.Memory.Read(addr) | (1 << 7)
	c.Memory.Write(addr, value)
}

func (c *CPU) InstructionCBFF() {
	value := c.Registers.GetA() | (1 << 7)
	c.Registers.SetA(value)
}
