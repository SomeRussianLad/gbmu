package cpu

type instruction struct {
	mnemonic string     // conventional name; primarily for tests and logging
	length   int        // instruction size in bytes
	cycles   int        // amount of machine cycles taken when executed; changes when instruction has variable execution duration
	exec     func(*CPU) // execution call
}

type instructions map[int]instruction

func newInstructions() instructions {
	return instructions{
		0x00: {"NOP", 1, 4, (*CPU).instruction00},
		0x01: {"LD BC,nn", 3, 12, (*CPU).instruction01},
		0x02: {"LD (BC),A", 1, 8, (*CPU).instruction02},
		0x03: {"INC BC", 1, 8, (*CPU).instruction03},
		0x04: {"INC B", 1, 4, (*CPU).instruction04},
		0x05: {"DEC B", 1, 4, (*CPU).instruction05},
		0x06: {"LD B,n", 2, 8, (*CPU).instruction06},
		0x07: {"RLCA", 1, 4, (*CPU).instruction07},
		0x08: {"LD (nn),SP", 3, 20, (*CPU).instruction08},
		0x09: {"ADD HL,BC", 1, 8, (*CPU).instruction09},
		0x0A: {"LD A,(BC)", 1, 8, (*CPU).instruction0A},
		0x0B: {"DEC BC", 1, 8, (*CPU).instruction0B},
		0x0C: {"INC C", 1, 4, (*CPU).instruction0C},
		0x0D: {"DEC C", 1, 4, (*CPU).instruction0D},
		0x0E: {"LD C,n", 2, 8, (*CPU).instruction0E},
		0x0F: {"RRCA", 1, 4, (*CPU).instruction0F},

		0x10: {"STOP", 2, 4, (*CPU).instruction10},
		0x11: {"LD DE,nn", 3, 12, (*CPU).instruction11},
		0x12: {"LD (DE),A", 1, 8, (*CPU).instruction12},
		0x13: {"INC DE", 1, 8, (*CPU).instruction13},
		0x14: {"INC D", 1, 4, (*CPU).instruction14},
		0x15: {"DEC D", 1, 4, (*CPU).instruction15},
		0x16: {"LD D,n", 2, 8, (*CPU).instruction16},
		0x17: {"RLA", 1, 4, (*CPU).instruction17},
		0x18: {"JR n", 2, 12, (*CPU).instruction18},
		0x19: {"ADD HL,DE", 1, 8, (*CPU).instruction19},
		0x1A: {"LD A,(DE)", 1, 8, (*CPU).instruction1A},
		0x1B: {"DEC DE", 1, 8, (*CPU).instruction1B},
		0x1C: {"INC E", 1, 4, (*CPU).instruction1C},
		0x1D: {"DEC E", 1, 4, (*CPU).instruction1D},
		0x1E: {"LD E,n", 2, 8, (*CPU).instruction1E},
		0x1F: {"RRA", 1, 4, (*CPU).instruction1F},

		0x20: {"JR NZ,n", 2, 8, (*CPU).instruction20},
		0x21: {"LD HL,nn", 3, 12, (*CPU).instruction21},
		0x22: {"LDI (HL),A", 1, 8, (*CPU).instruction22},
		0x23: {"INC HL", 1, 8, (*CPU).instruction23},
		0x24: {"INC H", 1, 4, (*CPU).instruction24},
		0x25: {"DEC H", 1, 4, (*CPU).instruction25},
		0x26: {"LD H,n", 2, 8, (*CPU).instruction26},
		0x27: {"DAA", 1, 4, (*CPU).instruction27},
		0x28: {"JR Z,n", 2, 8, (*CPU).instruction28},
		0x29: {"ADD HL,HL", 1, 8, (*CPU).instruction29},
		0x2A: {"LDI A,(HL)", 1, 8, (*CPU).instruction2A},
		0x2B: {"DEC HL", 1, 8, (*CPU).instruction2B},
		0x2C: {"INC L", 1, 4, (*CPU).instruction2C},
		0x2D: {"DEC L", 1, 4, (*CPU).instruction2D},
		0x2E: {"LD L,n", 2, 8, (*CPU).instruction2E},
		0x2F: {"CPL", 1, 4, (*CPU).instruction2F},

		0x30: {"JR NC,n", 2, 8, (*CPU).instruction30},
		0x31: {"LD SP,nn", 3, 12, (*CPU).instruction31},
		0x32: {"LDD (HL),A", 1, 8, (*CPU).instruction32},
		0x33: {"INC SP", 1, 8, (*CPU).instruction33},
		0x34: {"INC (HL)", 1, 12, (*CPU).instruction34},
		0x35: {"DEC (HL)", 1, 12, (*CPU).instruction35},
		0x36: {"LD (HL),n", 2, 12, (*CPU).instruction36},
		0x37: {"SCF", 1, 4, (*CPU).instruction37},
		0x38: {"JR C,n", 2, 8, (*CPU).instruction38},
		0x39: {"ADD HL,SP", 1, 8, (*CPU).instruction39},
		0x3A: {"LDD A,(HL)", 1, 8, (*CPU).instruction3A},
		0x3B: {"DEC SP", 1, 8, (*CPU).instruction3B},
		0x3C: {"INC A", 1, 4, (*CPU).instruction3C},
		0x3D: {"DEC A", 1, 4, (*CPU).instruction3D},
		0x3E: {"LD A,n", 2, 8, (*CPU).instruction3E},
		0x3F: {"CCF", 1, 4, (*CPU).instruction3F},

		0x40: {"LD B,B", 1, 4, (*CPU).instruction40},
		0x41: {"LD B,C", 1, 4, (*CPU).instruction41},
		0x42: {"LD B,D", 1, 4, (*CPU).instruction42},
		0x43: {"LD B,E", 1, 4, (*CPU).instruction43},
		0x44: {"LD B,H", 1, 4, (*CPU).instruction44},
		0x45: {"LD B,L", 1, 4, (*CPU).instruction45},
		0x46: {"LD B,(HL)", 1, 8, (*CPU).instruction46},
		0x47: {"LD B,A", 1, 4, (*CPU).instruction47},
		0x48: {"LD C,B", 1, 4, (*CPU).instruction48},
		0x49: {"LD C,C", 1, 4, (*CPU).instruction49},
		0x4A: {"LD C,D", 1, 4, (*CPU).instruction4A},
		0x4B: {"LD C,E", 1, 4, (*CPU).instruction4B},
		0x4C: {"LD C,H", 1, 4, (*CPU).instruction4C},
		0x4D: {"LD C,L", 1, 4, (*CPU).instruction4D},
		0x4E: {"LD C,(HL)", 1, 8, (*CPU).instruction4E},
		0x4F: {"LD C,A", 1, 4, (*CPU).instruction4F},

		0x50: {"LD D,B", 1, 4, (*CPU).instruction50},
		0x51: {"LD D,C", 1, 4, (*CPU).instruction51},
		0x52: {"LD D,D", 1, 4, (*CPU).instruction52},
		0x53: {"LD D,E", 1, 4, (*CPU).instruction53},
		0x54: {"LD D,H", 1, 4, (*CPU).instruction54},
		0x55: {"LD D,L", 1, 4, (*CPU).instruction55},
		0x56: {"LD D,(HL)", 1, 8, (*CPU).instruction56},
		0x57: {"LD D,A", 1, 4, (*CPU).instruction57},
		0x58: {"LD E,B", 1, 4, (*CPU).instruction58},
		0x59: {"LD E,C", 1, 4, (*CPU).instruction59},
		0x5A: {"LD E,D", 1, 4, (*CPU).instruction5A},
		0x5B: {"LD E,E", 1, 4, (*CPU).instruction5B},
		0x5C: {"LD E,H", 1, 4, (*CPU).instruction5C},
		0x5D: {"LD E,L", 1, 4, (*CPU).instruction5D},
		0x5E: {"LD E,(HL)", 1, 8, (*CPU).instruction5E},
		0x5F: {"LD E,A", 1, 4, (*CPU).instruction5F},

		0x60: {"LD H,B", 1, 4, (*CPU).instruction60},
		0x61: {"LD H,C", 1, 4, (*CPU).instruction61},
		0x62: {"LD H,D", 1, 4, (*CPU).instruction62},
		0x63: {"LD H,E", 1, 4, (*CPU).instruction63},
		0x64: {"LD H,H", 1, 4, (*CPU).instruction64},
		0x65: {"LD H,L", 1, 4, (*CPU).instruction65},
		0x66: {"LD H,(HL)", 1, 8, (*CPU).instruction66},
		0x67: {"LD H,A", 1, 4, (*CPU).instruction67},
		0x68: {"LD L,B", 1, 4, (*CPU).instruction68},
		0x69: {"LD L,C", 1, 4, (*CPU).instruction69},
		0x6A: {"LD L,D", 1, 4, (*CPU).instruction6A},
		0x6B: {"LD L,E", 1, 4, (*CPU).instruction6B},
		0x6C: {"LD L,H", 1, 4, (*CPU).instruction6C},
		0x6D: {"LD L,L", 1, 4, (*CPU).instruction6D},
		0x6E: {"LD L,(HL)", 1, 8, (*CPU).instruction6E},
		0x6F: {"LD L,A", 1, 4, (*CPU).instruction6F},

		0x70: {"LD (HL),B", 1, 8, (*CPU).instruction70},
		0x71: {"LD (HL),C", 1, 8, (*CPU).instruction71},
		0x72: {"LD (HL),D", 1, 8, (*CPU).instruction72},
		0x73: {"LD (HL),E", 1, 8, (*CPU).instruction73},
		0x74: {"LD (HL),H", 1, 8, (*CPU).instruction74},
		0x75: {"LD (HL),L", 1, 8, (*CPU).instruction75},
		0x76: {"HALT", 1, 4, (*CPU).instruction76},
		0x77: {"LD (HL),A", 1, 8, (*CPU).instruction77},
		0x78: {"LD A,B", 1, 4, (*CPU).instruction78},
		0x79: {"LD A,C", 1, 4, (*CPU).instruction79},
		0x7A: {"LD A,D", 1, 4, (*CPU).instruction7A},
		0x7B: {"LD A,E", 1, 4, (*CPU).instruction7B},
		0x7C: {"LD A,H", 1, 4, (*CPU).instruction7C},
		0x7D: {"LD A,L", 1, 4, (*CPU).instruction7D},
		0x7E: {"LD A,(HL)", 1, 8, (*CPU).instruction7E},
		0x7F: {"LD A,A", 1, 4, (*CPU).instruction7F},

		0x80: {"ADD A,B", 1, 4, (*CPU).instruction80},
		0x81: {"ADD A,C", 1, 4, (*CPU).instruction81},
		0x82: {"ADD A,D", 1, 4, (*CPU).instruction82},
		0x83: {"ADD A,E", 1, 4, (*CPU).instruction83},
		0x84: {"ADD A,H", 1, 4, (*CPU).instruction84},
		0x85: {"ADD A,L", 1, 4, (*CPU).instruction85},
		0x86: {"ADD A,(HL)", 1, 8, (*CPU).instruction86},
		0x87: {"ADD A,A", 1, 4, (*CPU).instruction87},
		0x88: {"ADC A,B", 1, 4, (*CPU).instruction88},
		0x89: {"ADC A,C", 1, 4, (*CPU).instruction89},
		0x8A: {"ADC A,D", 1, 4, (*CPU).instruction8A},
		0x8B: {"ADC A,E", 1, 4, (*CPU).instruction8B},
		0x8C: {"ADC A,H", 1, 4, (*CPU).instruction8C},
		0x8D: {"ADC A,L", 1, 4, (*CPU).instruction8D},
		0x8E: {"ADC A,(HL)", 1, 8, (*CPU).instruction8E},
		0x8F: {"ADC A,A", 1, 4, (*CPU).instruction8F},

		0x90: {"SUB B", 1, 4, (*CPU).instruction90},
		0x91: {"SUB C", 1, 4, (*CPU).instruction91},
		0x92: {"SUB D", 1, 4, (*CPU).instruction92},
		0x93: {"SUB E", 1, 4, (*CPU).instruction93},
		0x94: {"SUB H", 1, 4, (*CPU).instruction94},
		0x95: {"SUB L", 1, 4, (*CPU).instruction95},
		0x96: {"SUB (HL)", 1, 8, (*CPU).instruction96},
		0x97: {"SUB A", 1, 4, (*CPU).instruction97},
		0x98: {"SBC A,B", 1, 4, (*CPU).instruction98},
		0x99: {"SBC A,C", 1, 4, (*CPU).instruction99},
		0x9A: {"SBC A,D", 1, 4, (*CPU).instruction9A},
		0x9B: {"SBC A,E", 1, 4, (*CPU).instruction9B},
		0x9C: {"SBC A,H", 1, 4, (*CPU).instruction9C},
		0x9D: {"SBC A,L", 1, 4, (*CPU).instruction9D},
		0x9E: {"SBC A,(HL)", 1, 8, (*CPU).instruction9E},
		0x9F: {"SBC A,A", 1, 4, (*CPU).instruction9F},

		0xA0: {"AND B", 1, 4, (*CPU).instructionA0},
		0xA1: {"AND C", 1, 4, (*CPU).instructionA1},
		0xA2: {"AND D", 1, 4, (*CPU).instructionA2},
		0xA3: {"AND E", 1, 4, (*CPU).instructionA3},
		0xA4: {"AND H", 1, 4, (*CPU).instructionA4},
		0xA5: {"AND L", 1, 4, (*CPU).instructionA5},
		0xA6: {"AND (HL)", 1, 8, (*CPU).instructionA6},
		0xA7: {"AND A", 1, 4, (*CPU).instructionA7},
		0xA8: {"XOR B", 1, 4, (*CPU).instructionA8},
		0xA9: {"XOR C", 1, 4, (*CPU).instructionA9},
		0xAA: {"XOR D", 1, 4, (*CPU).instructionAA},
		0xAB: {"XOR E", 1, 4, (*CPU).instructionAB},
		0xAC: {"XOR H", 1, 4, (*CPU).instructionAC},
		0xAD: {"XOR L", 1, 4, (*CPU).instructionAD},
		0xAE: {"XOR (HL)", 1, 8, (*CPU).instructionAE},
		0xAF: {"XOR A", 1, 4, (*CPU).instructionAF},

		0xB0: {"OR B", 1, 4, (*CPU).instructionB0},
		0xB1: {"OR C", 1, 4, (*CPU).instructionB1},
		0xB2: {"OR D", 1, 4, (*CPU).instructionB2},
		0xB3: {"OR E", 1, 4, (*CPU).instructionB3},
		0xB4: {"OR H", 1, 4, (*CPU).instructionB4},
		0xB5: {"OR L", 1, 4, (*CPU).instructionB5},
		0xB6: {"OR (HL)", 1, 8, (*CPU).instructionB6},
		0xB7: {"OR A", 1, 4, (*CPU).instructionB7},
		0xB8: {"CP B", 1, 4, (*CPU).instructionB8},
		0xB9: {"CP C", 1, 4, (*CPU).instructionB9},
		0xBA: {"CP D", 1, 4, (*CPU).instructionBA},
		0xBB: {"CP E", 1, 4, (*CPU).instructionBB},
		0xBC: {"CP H", 1, 4, (*CPU).instructionBC},
		0xBD: {"CP L", 1, 4, (*CPU).instructionBD},
		0xBE: {"CP (HL)", 1, 8, (*CPU).instructionBE},
		0xBF: {"CP A", 1, 4, (*CPU).instructionBF},

		0xC0: {"RET NZ", 1, 8, (*CPU).instructionC0},
		0xC1: {"POP BC", 1, 12, (*CPU).instructionC1},
		0xC2: {"JP NZ,nn", 3, 12, (*CPU).instructionC2},
		0xC3: {"JP nn", 3, 16, (*CPU).instructionC3},
		0xC4: {"CALL NZ,nn", 3, 12, (*CPU).instructionC4},
		0xC5: {"PUSH BC", 1, 16, (*CPU).instructionC5},
		0xC6: {"ADD A,n", 2, 8, (*CPU).instructionC6},
		0xC7: {"RST 00H", 1, 16, (*CPU).instructionC7},
		0xC8: {"RET Z", 1, 8, (*CPU).instructionC8},
		0xC9: {"RET", 1, 16, (*CPU).instructionC9},
		0xCA: {"JP Z,nn", 3, 12, (*CPU).instructionCA},
		0xCB: {"PREFIX CB", 1, 4, (*CPU).instructionCB},
		0xCC: {"CALL Z,nn", 3, 12, (*CPU).instructionCC},
		0xCD: {"CALL nn", 3, 24, (*CPU).instructionCD},
		0xCE: {"ADC A,n", 2, 8, (*CPU).instructionCE},
		0xCF: {"RST 08H", 1, 16, (*CPU).instructionCF},

		0xD0: {"RET NC", 1, 8, (*CPU).instructionD0},
		0xD1: {"POP DE", 1, 12, (*CPU).instructionD1},
		0xD2: {"JP NC,nn", 3, 12, (*CPU).instructionD2},
		0xD4: {"CALL NC,nn", 3, 12, (*CPU).instructionD4},
		0xD5: {"PUSH DE", 1, 16, (*CPU).instructionD5},
		0xD6: {"SUB n", 2, 8, (*CPU).instructionD6},
		0xD7: {"RST 10H", 1, 16, (*CPU).instructionD7},
		0xD8: {"RET C", 1, 8, (*CPU).instructionD8},
		0xD9: {"RETI", 1, 16, (*CPU).instructionD9},
		0xDA: {"JP C,nn", 3, 12, (*CPU).instructionDA},
		0xDC: {"CALL C,nn", 3, 12, (*CPU).instructionDC},
		0xDE: {"SBC A,n", 2, 8, (*CPU).instructionDE},
		0xDF: {"RST 18H", 1, 16, (*CPU).instructionDF},

		0xE0: {"LDH (n),A", 2, 12, (*CPU).instructionE0},
		0xE1: {"POP HL", 1, 12, (*CPU).instructionE1},
		0xE2: {"LD (C),A", 1, 8, (*CPU).instructionE2},
		0xE5: {"PUSH HL", 1, 16, (*CPU).instructionE5},
		0xE6: {"AND n", 2, 8, (*CPU).instructionE6},
		0xE7: {"RST 20H", 1, 16, (*CPU).instructionE7},
		0xE8: {"ADD SP,e", 2, 16, (*CPU).instructionE8},
		0xE9: {"JP HL", 1, 4, (*CPU).instructionE9},
		0xEA: {"LD (nn),A", 3, 16, (*CPU).instructionEA},
		0xEE: {"XOR n", 2, 8, (*CPU).instructionEE},
		0xEF: {"RST 28H", 1, 16, (*CPU).instructionEF},

		0xF0: {"LDH A,(n)", 2, 12, (*CPU).instructionF0},
		0xF1: {"POP AF", 1, 12, (*CPU).instructionF1},
		0xF2: {"LD A,(C)", 1, 8, (*CPU).instructionF2},
		0xF3: {"DI", 1, 4, (*CPU).instructionF3},
		0xF5: {"PUSH AF", 1, 16, (*CPU).instructionF5},
		0xF6: {"OR n", 2, 8, (*CPU).instructionF6},
		0xF7: {"RST 30H", 1, 16, (*CPU).instructionF7},
		0xF8: {"LDHL SP+e", 2, 12, (*CPU).instructionF8},
		0xF9: {"LD SP,HL", 1, 8, (*CPU).instructionF9},
		0xFA: {"LD A,(nn)", 3, 16, (*CPU).instructionFA},
		0xFB: {"EI", 1, 4, (*CPU).instructionFB},
		0xFE: {"CP n", 2, 8, (*CPU).instructionFE},
		0xFF: {"RST 38H", 1, 16, (*CPU).instructionFF},

		0xCB00: {"RLC B", 2, 8, (*CPU).instructionCB00},
		0xCB01: {"RLC C", 2, 8, (*CPU).instructionCB01},
		0xCB02: {"RLC D", 2, 8, (*CPU).instructionCB02},
		0xCB03: {"RLC E", 2, 8, (*CPU).instructionCB03},
		0xCB04: {"RLC H", 2, 8, (*CPU).instructionCB04},
		0xCB05: {"RLC L", 2, 8, (*CPU).instructionCB05},
		0xCB06: {"RLC (HL)", 2, 16, (*CPU).instructionCB06},
		0xCB07: {"RLC A", 2, 8, (*CPU).instructionCB07},
		0xCB08: {"RRC B", 2, 8, (*CPU).instructionCB08},
		0xCB09: {"RRC C", 2, 8, (*CPU).instructionCB09},
		0xCB0A: {"RRC D", 2, 8, (*CPU).instructionCB0A},
		0xCB0B: {"RRC E", 2, 8, (*CPU).instructionCB0B},
		0xCB0C: {"RRC H", 2, 8, (*CPU).instructionCB0C},
		0xCB0D: {"RRC L", 2, 8, (*CPU).instructionCB0D},
		0xCB0E: {"RRC (HL)", 2, 16, (*CPU).instructionCB0E},
		0xCB0F: {"RRC A", 2, 8, (*CPU).instructionCB0F},

		0xCB10: {"RL B", 2, 8, (*CPU).instructionCB10},
		0xCB11: {"RL C", 2, 8, (*CPU).instructionCB11},
		0xCB12: {"RL D", 2, 8, (*CPU).instructionCB12},
		0xCB13: {"RL E", 2, 8, (*CPU).instructionCB13},
		0xCB14: {"RL H", 2, 8, (*CPU).instructionCB14},
		0xCB15: {"RL L", 2, 8, (*CPU).instructionCB15},
		0xCB16: {"RL (HL)", 2, 16, (*CPU).instructionCB16},
		0xCB17: {"RL A", 2, 8, (*CPU).instructionCB17},
		0xCB18: {"RR B", 2, 8, (*CPU).instructionCB18},
		0xCB19: {"RR C", 2, 8, (*CPU).instructionCB19},
		0xCB1A: {"RR D", 2, 8, (*CPU).instructionCB1A},
		0xCB1B: {"RR E", 2, 8, (*CPU).instructionCB1B},
		0xCB1C: {"RR H", 2, 8, (*CPU).instructionCB1C},
		0xCB1D: {"RR L", 2, 8, (*CPU).instructionCB1D},
		0xCB1E: {"RR (HL)", 2, 16, (*CPU).instructionCB1E},
		0xCB1F: {"RR A", 2, 8, (*CPU).instructionCB1F},

		0xCB20: {"SLA B", 2, 8, (*CPU).instructionCB20},
		0xCB21: {"SLA C", 2, 8, (*CPU).instructionCB21},
		0xCB22: {"SLA D", 2, 8, (*CPU).instructionCB22},
		0xCB23: {"SLA E", 2, 8, (*CPU).instructionCB23},
		0xCB24: {"SLA H", 2, 8, (*CPU).instructionCB24},
		0xCB25: {"SLA L", 2, 8, (*CPU).instructionCB25},
		0xCB26: {"SLA (HL)", 2, 16, (*CPU).instructionCB26},
		0xCB27: {"SLA A", 2, 8, (*CPU).instructionCB27},
		0xCB28: {"SRA B", 2, 8, (*CPU).instructionCB28},
		0xCB29: {"SRA C", 2, 8, (*CPU).instructionCB29},
		0xCB2A: {"SRA D", 2, 8, (*CPU).instructionCB2A},
		0xCB2B: {"SRA E", 2, 8, (*CPU).instructionCB2B},
		0xCB2C: {"SRA H", 2, 8, (*CPU).instructionCB2C},
		0xCB2D: {"SRA L", 2, 8, (*CPU).instructionCB2D},
		0xCB2E: {"SRA (HL)", 2, 16, (*CPU).instructionCB2E},
		0xCB2F: {"SRA A", 2, 8, (*CPU).instructionCB2F},

		0xCB30: {"SWAP B", 2, 8, (*CPU).instructionCB30},
		0xCB31: {"SWAP C", 2, 8, (*CPU).instructionCB31},
		0xCB32: {"SWAP D", 2, 8, (*CPU).instructionCB32},
		0xCB33: {"SWAP E", 2, 8, (*CPU).instructionCB33},
		0xCB34: {"SWAP H", 2, 8, (*CPU).instructionCB34},
		0xCB35: {"SWAP L", 2, 8, (*CPU).instructionCB35},
		0xCB36: {"SWAP (HL)", 2, 16, (*CPU).instructionCB36},
		0xCB37: {"SWAP A", 2, 8, (*CPU).instructionCB37},
		0xCB38: {"SRL B", 2, 8, (*CPU).instructionCB38},
		0xCB39: {"SRL C", 2, 8, (*CPU).instructionCB39},
		0xCB3A: {"SRL D", 2, 8, (*CPU).instructionCB3A},
		0xCB3B: {"SRL E", 2, 8, (*CPU).instructionCB3B},
		0xCB3C: {"SRL H", 2, 8, (*CPU).instructionCB3C},
		0xCB3D: {"SRL L", 2, 8, (*CPU).instructionCB3D},
		0xCB3E: {"SRL (HL)", 2, 16, (*CPU).instructionCB3E},
		0xCB3F: {"SRL A", 2, 8, (*CPU).instructionCB3F},

		0xCB40: {"BIT 0,B", 2, 8, (*CPU).instructionCB40},
		0xCB41: {"BIT 0,C", 2, 8, (*CPU).instructionCB41},
		0xCB42: {"BIT 0,D", 2, 8, (*CPU).instructionCB42},
		0xCB43: {"BIT 0,E", 2, 8, (*CPU).instructionCB43},
		0xCB44: {"BIT 0,H", 2, 8, (*CPU).instructionCB44},
		0xCB45: {"BIT 0,L", 2, 8, (*CPU).instructionCB45},
		0xCB46: {"BIT 0,(HL)", 2, 16, (*CPU).instructionCB46},
		0xCB47: {"BIT 0,A", 2, 8, (*CPU).instructionCB47},
		0xCB48: {"BIT 1,B", 2, 8, (*CPU).instructionCB48},
		0xCB49: {"BIT 1,C", 2, 8, (*CPU).instructionCB49},
		0xCB4A: {"BIT 1,D", 2, 8, (*CPU).instructionCB4A},
		0xCB4B: {"BIT 1,E", 2, 8, (*CPU).instructionCB4B},
		0xCB4C: {"BIT 1,H", 2, 8, (*CPU).instructionCB4C},
		0xCB4D: {"BIT 1,L", 2, 8, (*CPU).instructionCB4D},
		0xCB4E: {"BIT 1,(HL)", 2, 16, (*CPU).instructionCB4E},
		0xCB4F: {"BIT 1,A", 2, 8, (*CPU).instructionCB4F},
		0xCB50: {"BIT 2,B", 2, 8, (*CPU).instructionCB50},

		0xCB51: {"BIT 2,C", 2, 8, (*CPU).instructionCB51},
		0xCB52: {"BIT 2,D", 2, 8, (*CPU).instructionCB52},
		0xCB53: {"BIT 2,E", 2, 8, (*CPU).instructionCB53},
		0xCB54: {"BIT 2,H", 2, 8, (*CPU).instructionCB54},
		0xCB55: {"BIT 2,L", 2, 8, (*CPU).instructionCB55},
		0xCB56: {"BIT 2,(HL)", 2, 16, (*CPU).instructionCB56},
		0xCB57: {"BIT 2,A", 2, 8, (*CPU).instructionCB57},
		0xCB58: {"BIT 3,B", 2, 8, (*CPU).instructionCB58},
		0xCB59: {"BIT 3,C", 2, 8, (*CPU).instructionCB59},
		0xCB5A: {"BIT 3,D", 2, 8, (*CPU).instructionCB5A},
		0xCB5B: {"BIT 3,E", 2, 8, (*CPU).instructionCB5B},
		0xCB5C: {"BIT 3,H", 2, 8, (*CPU).instructionCB5C},
		0xCB5D: {"BIT 3,L", 2, 8, (*CPU).instructionCB5D},
		0xCB5E: {"BIT 3,(HL)", 2, 16, (*CPU).instructionCB5E},
		0xCB5F: {"BIT 3,A", 2, 8, (*CPU).instructionCB5F},

		0xCB60: {"BIT 4,B", 2, 8, (*CPU).instructionCB60},
		0xCB61: {"BIT 4,C", 2, 8, (*CPU).instructionCB61},
		0xCB62: {"BIT 4,D", 2, 8, (*CPU).instructionCB62},
		0xCB63: {"BIT 4,E", 2, 8, (*CPU).instructionCB63},
		0xCB64: {"BIT 4,H", 2, 8, (*CPU).instructionCB64},
		0xCB65: {"BIT 4,L", 2, 8, (*CPU).instructionCB65},
		0xCB66: {"BIT 4,(HL)", 2, 16, (*CPU).instructionCB66},
		0xCB67: {"BIT 4,A", 2, 8, (*CPU).instructionCB67},
		0xCB68: {"BIT 5,B", 2, 8, (*CPU).instructionCB68},
		0xCB69: {"BIT 5,C", 2, 8, (*CPU).instructionCB69},
		0xCB6A: {"BIT 5,D", 2, 8, (*CPU).instructionCB6A},
		0xCB6B: {"BIT 5,E", 2, 8, (*CPU).instructionCB6B},
		0xCB6C: {"BIT 5,H", 2, 8, (*CPU).instructionCB6C},
		0xCB6D: {"BIT 5,L", 2, 8, (*CPU).instructionCB6D},
		0xCB6E: {"BIT 5,(HL)", 2, 16, (*CPU).instructionCB6E},
		0xCB6F: {"BIT 5,A", 2, 8, (*CPU).instructionCB6F},

		0xCB70: {"BIT 6,B", 2, 8, (*CPU).instructionCB70},
		0xCB71: {"BIT 6,C", 2, 8, (*CPU).instructionCB71},
		0xCB72: {"BIT 6,D", 2, 8, (*CPU).instructionCB72},
		0xCB73: {"BIT 6,E", 2, 8, (*CPU).instructionCB73},
		0xCB74: {"BIT 6,H", 2, 8, (*CPU).instructionCB74},
		0xCB75: {"BIT 6,L", 2, 8, (*CPU).instructionCB75},
		0xCB76: {"BIT 6,(HL)", 2, 16, (*CPU).instructionCB76},
		0xCB77: {"BIT 6,A", 2, 8, (*CPU).instructionCB77},
		0xCB78: {"BIT 7,B", 2, 8, (*CPU).instructionCB78},
		0xCB79: {"BIT 7,C", 2, 8, (*CPU).instructionCB79},
		0xCB7A: {"BIT 7,D", 2, 8, (*CPU).instructionCB7A},
		0xCB7B: {"BIT 7,E", 2, 8, (*CPU).instructionCB7B},
		0xCB7C: {"BIT 7,H", 2, 8, (*CPU).instructionCB7C},
		0xCB7D: {"BIT 7,L", 2, 8, (*CPU).instructionCB7D},
		0xCB7E: {"BIT 7,(HL)", 2, 16, (*CPU).instructionCB7E},
		0xCB7F: {"BIT 7,A", 2, 8, (*CPU).instructionCB7F},

		0xCB80: {"RES 0,B", 2, 8, (*CPU).instructionCB80},
		0xCB81: {"RES 0,C", 2, 8, (*CPU).instructionCB81},
		0xCB82: {"RES 0,D", 2, 8, (*CPU).instructionCB82},
		0xCB83: {"RES 0,E", 2, 8, (*CPU).instructionCB83},
		0xCB84: {"RES 0,H", 2, 8, (*CPU).instructionCB84},
		0xCB85: {"RES 0,L", 2, 8, (*CPU).instructionCB85},
		0xCB86: {"RES 0,(HL)", 2, 16, (*CPU).instructionCB86},
		0xCB87: {"RES 0,A", 2, 8, (*CPU).instructionCB87},
		0xCB88: {"RES 1,B", 2, 8, (*CPU).instructionCB88},
		0xCB89: {"RES 1,C", 2, 8, (*CPU).instructionCB89},
		0xCB8A: {"RES 1,D", 2, 8, (*CPU).instructionCB8A},
		0xCB8B: {"RES 1,E", 2, 8, (*CPU).instructionCB8B},
		0xCB8C: {"RES 1,H", 2, 8, (*CPU).instructionCB8C},
		0xCB8D: {"RES 1,L", 2, 8, (*CPU).instructionCB8D},
		0xCB8E: {"RES 1,(HL)", 2, 16, (*CPU).instructionCB8E},
		0xCB8F: {"RES 1,A", 2, 8, (*CPU).instructionCB8F},

		0xCB90: {"RES 2,B", 2, 8, (*CPU).instructionCB90},
		0xCB91: {"RES 2,C", 2, 8, (*CPU).instructionCB91},
		0xCB92: {"RES 2,D", 2, 8, (*CPU).instructionCB92},
		0xCB93: {"RES 2,E", 2, 8, (*CPU).instructionCB93},
		0xCB94: {"RES 2,H", 2, 8, (*CPU).instructionCB94},
		0xCB95: {"RES 2,L", 2, 8, (*CPU).instructionCB95},
		0xCB96: {"RES 2,(HL)", 2, 16, (*CPU).instructionCB96},
		0xCB97: {"RES 2,A", 2, 8, (*CPU).instructionCB97},
		0xCB98: {"RES 3,B", 2, 8, (*CPU).instructionCB98},
		0xCB99: {"RES 3,C", 2, 8, (*CPU).instructionCB99},
		0xCB9A: {"RES 3,D", 2, 8, (*CPU).instructionCB9A},
		0xCB9B: {"RES 3,E", 2, 8, (*CPU).instructionCB9B},
		0xCB9C: {"RES 3,H", 2, 8, (*CPU).instructionCB9C},
		0xCB9D: {"RES 3,L", 2, 8, (*CPU).instructionCB9D},
		0xCB9E: {"RES 3,(HL)", 2, 16, (*CPU).instructionCB9E},
		0xCB9F: {"RES 3,A", 2, 8, (*CPU).instructionCB9F},

		0xCBA0: {"RES 4,B", 2, 8, (*CPU).instructionCBA0},
		0xCBA1: {"RES 4,C", 2, 8, (*CPU).instructionCBA1},
		0xCBA2: {"RES 4,D", 2, 8, (*CPU).instructionCBA2},
		0xCBA3: {"RES 4,E", 2, 8, (*CPU).instructionCBA3},
		0xCBA4: {"RES 4,H", 2, 8, (*CPU).instructionCBA4},
		0xCBA5: {"RES 4,L", 2, 8, (*CPU).instructionCBA5},
		0xCBA6: {"RES 4,(HL)", 2, 16, (*CPU).instructionCBA6},
		0xCBA7: {"RES 4,A", 2, 8, (*CPU).instructionCBA7},
		0xCBA8: {"RES 5,B", 2, 8, (*CPU).instructionCBA8},
		0xCBA9: {"RES 5,C", 2, 8, (*CPU).instructionCBA9},
		0xCBAA: {"RES 5,D", 2, 8, (*CPU).instructionCBAA},
		0xCBAB: {"RES 5,E", 2, 8, (*CPU).instructionCBAB},
		0xCBAC: {"RES 5,H", 2, 8, (*CPU).instructionCBAC},
		0xCBAD: {"RES 5,L", 2, 8, (*CPU).instructionCBAD},
		0xCBAE: {"RES 5,(HL)", 2, 16, (*CPU).instructionCBAE},
		0xCBAF: {"RES 5,A", 2, 8, (*CPU).instructionCBAF},

		0xCBB0: {"RES 6,B", 2, 8, (*CPU).instructionCBB0},
		0xCBB1: {"RES 6,C", 2, 8, (*CPU).instructionCBB1},
		0xCBB2: {"RES 6,D", 2, 8, (*CPU).instructionCBB2},
		0xCBB3: {"RES 6,E", 2, 8, (*CPU).instructionCBB3},
		0xCBB4: {"RES 6,H", 2, 8, (*CPU).instructionCBB4},
		0xCBB5: {"RES 6,L", 2, 8, (*CPU).instructionCBB5},
		0xCBB6: {"RES 6,(HL)", 2, 16, (*CPU).instructionCBB6},
		0xCBB7: {"RES 6,A", 2, 8, (*CPU).instructionCBB7},
		0xCBB8: {"RES 7,B", 2, 8, (*CPU).instructionCBB8},
		0xCBB9: {"RES 7,C", 2, 8, (*CPU).instructionCBB9},
		0xCBBA: {"RES 7,D", 2, 8, (*CPU).instructionCBBA},
		0xCBBB: {"RES 7,E", 2, 8, (*CPU).instructionCBBB},
		0xCBBC: {"RES 7,H", 2, 8, (*CPU).instructionCBBC},
		0xCBBD: {"RES 7,L", 2, 8, (*CPU).instructionCBBD},
		0xCBBE: {"RES 7,(HL)", 2, 16, (*CPU).instructionCBBE},
		0xCBBF: {"RES 7,A", 2, 8, (*CPU).instructionCBBF},

		0xCBC0: {"SET 0,B", 2, 8, (*CPU).instructionCBC0},
		0xCBC1: {"SET 0,C", 2, 8, (*CPU).instructionCBC1},
		0xCBC2: {"SET 0,D", 2, 8, (*CPU).instructionCBC2},
		0xCBC3: {"SET 0,E", 2, 8, (*CPU).instructionCBC3},
		0xCBC4: {"SET 0,H", 2, 8, (*CPU).instructionCBC4},
		0xCBC5: {"SET 0,L", 2, 8, (*CPU).instructionCBC5},
		0xCBC6: {"SET 0,(HL)", 2, 16, (*CPU).instructionCBC6},
		0xCBC7: {"SET 0,A", 2, 8, (*CPU).instructionCBC7},
		0xCBC8: {"SET 1,B", 2, 8, (*CPU).instructionCBC8},
		0xCBC9: {"SET 1,C", 2, 8, (*CPU).instructionCBC9},
		0xCBCA: {"SET 1,D", 2, 8, (*CPU).instructionCBCA},
		0xCBCB: {"SET 1,E", 2, 8, (*CPU).instructionCBCB},
		0xCBCC: {"SET 1,H", 2, 8, (*CPU).instructionCBCC},
		0xCBCD: {"SET 1,L", 2, 8, (*CPU).instructionCBCD},
		0xCBCE: {"SET 1,(HL)", 2, 16, (*CPU).instructionCBCE},
		0xCBCF: {"SET 1,A", 2, 8, (*CPU).instructionCBCF},

		0xCBD0: {"SET 2,B", 2, 8, (*CPU).instructionCBD0},
		0xCBD1: {"SET 2,C", 2, 8, (*CPU).instructionCBD1},
		0xCBD2: {"SET 2,D", 2, 8, (*CPU).instructionCBD2},
		0xCBD3: {"SET 2,E", 2, 8, (*CPU).instructionCBD3},
		0xCBD4: {"SET 2,H", 2, 8, (*CPU).instructionCBD4},
		0xCBD5: {"SET 2,L", 2, 8, (*CPU).instructionCBD5},
		0xCBD6: {"SET 2,(HL)", 2, 16, (*CPU).instructionCBD6},
		0xCBD7: {"SET 2,A", 2, 8, (*CPU).instructionCBD7},
		0xCBD8: {"SET 3,B", 2, 8, (*CPU).instructionCBD8},
		0xCBD9: {"SET 3,C", 2, 8, (*CPU).instructionCBD9},
		0xCBDA: {"SET 3,D", 2, 8, (*CPU).instructionCBDA},
		0xCBDB: {"SET 3,E", 2, 8, (*CPU).instructionCBDB},
		0xCBDC: {"SET 3,H", 2, 8, (*CPU).instructionCBDC},
		0xCBDD: {"SET 3,L", 2, 8, (*CPU).instructionCBDD},
		0xCBDE: {"SET 3,(HL)", 2, 16, (*CPU).instructionCBDE},
		0xCBDF: {"SET 3,A", 2, 8, (*CPU).instructionCBDF},

		0xCBE0: {"SET 4,B", 2, 8, (*CPU).instructionCBE0},
		0xCBE1: {"SET 4,C", 2, 8, (*CPU).instructionCBE1},
		0xCBE2: {"SET 4,D", 2, 8, (*CPU).instructionCBE2},
		0xCBE3: {"SET 4,E", 2, 8, (*CPU).instructionCBE3},
		0xCBE4: {"SET 4,H", 2, 8, (*CPU).instructionCBE4},
		0xCBE5: {"SET 4,L", 2, 8, (*CPU).instructionCBE5},
		0xCBE6: {"SET 4,(HL)", 2, 16, (*CPU).instructionCBE6},
		0xCBE7: {"SET 4,A", 2, 8, (*CPU).instructionCBE7},
		0xCBE8: {"SET 5,B", 2, 8, (*CPU).instructionCBE8},
		0xCBE9: {"SET 5,C", 2, 8, (*CPU).instructionCBE9},
		0xCBEA: {"SET 5,D", 2, 8, (*CPU).instructionCBEA},
		0xCBEB: {"SET 5,E", 2, 8, (*CPU).instructionCBEB},
		0xCBEC: {"SET 5,H", 2, 8, (*CPU).instructionCBEC},
		0xCBED: {"SET 5,L", 2, 8, (*CPU).instructionCBED},
		0xCBEE: {"SET 5,(HL)", 2, 16, (*CPU).instructionCBEE},
		0xCBEF: {"SET 5,A", 2, 8, (*CPU).instructionCBEF},

		0xCBF0: {"SET 6,B", 2, 8, (*CPU).instructionCBF0},
		0xCBF1: {"SET 6,C", 2, 8, (*CPU).instructionCBF1},
		0xCBF2: {"SET 6,D", 2, 8, (*CPU).instructionCBF2},
		0xCBF3: {"SET 6,E", 2, 8, (*CPU).instructionCBF3},
		0xCBF4: {"SET 6,H", 2, 8, (*CPU).instructionCBF4},
		0xCBF5: {"SET 6,L", 2, 8, (*CPU).instructionCBF5},
		0xCBF6: {"SET 6,(HL)", 2, 16, (*CPU).instructionCBF6},
		0xCBF7: {"SET 6,A", 2, 8, (*CPU).instructionCBF7},
		0xCBF8: {"SET 7,B", 2, 8, (*CPU).instructionCBF8},
		0xCBF9: {"SET 7,C", 2, 8, (*CPU).instructionCBF9},
		0xCBFA: {"SET 7,D", 2, 8, (*CPU).instructionCBFA},
		0xCBFB: {"SET 7,E", 2, 8, (*CPU).instructionCBFB},
		0xCBFC: {"SET 7,H", 2, 8, (*CPU).instructionCBFC},
		0xCBFD: {"SET 7,L", 2, 8, (*CPU).instructionCBFD},
		0xCBFE: {"SET 7,(HL)", 2, 16, (*CPU).instructionCBFE},
		0xCBFF: {"SET 7,A", 2, 8, (*CPU).instructionCBFF},
	}
}

// TODO(somerussianlad): Fix every instruction that reads 16 bytes and use swapBytes(), if needed

//     01,                     08,
//     11,
//     21,
//     31,
// C0, C1, C2, C3, C4, C5, C7, C8, C9, CA, CC, CD, CF,
// D0, D1, D2,     D4, D5, D7, D8, D9, DA, DC,     DF,
//     E1,             E5, E7,     E9, EA,         EF,
//     F1,             F5, F7,         FA,         FF,

func (c *CPU) read8BitOperand() uint8 {
	value := c.readPC()
	return value
}

func (c *CPU) read16BitOperand() uint16 {
	lsb := uint16(c.readPC())
	msb := uint16(c.readPC())
	value := msb<<8 | lsb
	return value
}

func (c *CPU) instruction00() {
	// NOP
}

func (c *CPU) instruction01() {
	value := c.read16BitOperand()
	c.registers.setBC(value)
}

func (c *CPU) instruction02() {
	addr := c.registers.getBC()
	value := c.registers.getA()
	c.memory.Write(addr, value)
}

func (c *CPU) instruction03() {
	c.registers.incBC()
}

func (c *CPU) instruction04() {
	value1 := c.registers.getB()
	value2 := value1 + 1
	c.registers.f.setZ(value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(value2&0xF == 0)
	c.registers.incB()
}

func (c *CPU) instruction05() {
	value1 := c.registers.getB()
	value2 := value1 - 1
	c.registers.f.setZ(value2 == 0)
	c.registers.f.setN(true)
	c.registers.f.setH(value1&0xF == 0)
	c.registers.decB()
}

func (c *CPU) instruction06() {
	value := c.read8BitOperand()
	c.registers.setB(value)
}

func (c *CPU) instruction07() {
	value := c.registers.getA()
	bit7 := value >> 7
	value = (value << 1) + bit7
	c.registers.f.setZ(false)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit7 == 1)
	c.registers.setA(value)
}

func (c *CPU) instruction08() {
	addr := c.read16BitOperand()
	lsb := uint8(c.registers.getSP() & 0xFF)
	msb := uint8(c.registers.getSP() >> 8)
	c.memory.Write(addr, lsb)
	c.memory.Write(addr+1, msb)
}

func (c *CPU) instruction09() {
	value1 := c.registers.getHL()
	value2 := c.registers.getBC()
	c.registers.f.setN(false)
	c.registers.f.setH((value1&0xFFF)+(value2&0xFFF) > 0xFFF)
	c.registers.f.setC(uint32(value1)+uint32(value2) > 0xFFFF)
	c.registers.setHL(value1 + value2)
}

func (c *CPU) instruction0A() {
	addr := c.registers.getBC()
	value := c.memory.Read(addr)
	c.registers.setA(value)
}

func (c *CPU) instruction0B() {
	c.registers.decBC()
}

func (c *CPU) instruction0C() {
	value1 := c.registers.getC()
	value2 := value1 + 1
	c.registers.f.setZ(value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(value2&0xF == 0)
	c.registers.incC()
}

func (c *CPU) instruction0D() {
	value1 := c.registers.getC()
	value2 := value1 - 1
	c.registers.f.setZ(value2 == 0)
	c.registers.f.setN(true)
	c.registers.f.setH(value1&0xF == 0)
	c.registers.decC()
}

func (c *CPU) instruction0E() {
	value := c.read8BitOperand()
	c.registers.setC(value)
}

func (c *CPU) instruction0F() {
	value := c.registers.getA()
	bit0 := value << 7
	value = (value >> 1) + bit0
	c.registers.f.setZ(false)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit0 == 128)
	c.registers.setA(value)
}

func (c *CPU) instruction10() {
	// STOP
	// c.divider.stop()
	// c.divider.reset()
	c.registers.incPC() // skipping the next byte 0x00 as it is part of STOP
}

func (c *CPU) instruction11() {
	value := c.read16BitOperand()
	c.registers.setDE(value)
}

func (c *CPU) instruction12() {
	addr := c.registers.getDE()
	value := c.registers.getA()
	c.memory.Write(addr, value)
}

func (c *CPU) instruction13() {
	c.registers.incDE()
}

func (c *CPU) instruction14() {
	value1 := c.registers.getD()
	value2 := value1 + 1
	c.registers.f.setZ(value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(value2&0xF == 0)
	c.registers.incD()
}

func (c *CPU) instruction15() {
	value1 := c.registers.getD()
	value2 := value1 - 1
	c.registers.f.setZ(value2 == 0)
	c.registers.f.setN(true)
	c.registers.f.setH(value1&0xF == 0)
	c.registers.decD()
}

func (c *CPU) instruction16() {
	value := c.read8BitOperand()
	c.registers.setD(value)
}

func (c *CPU) instruction17() {
	carry := c.registers.f.getCarry()
	value := c.registers.getA()
	bit7 := value >> 7
	value = (value << 1) + carry
	c.registers.f.setZ(false)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit7 == 1)
	c.registers.setA(value)
}

func (c *CPU) instruction18() {
	value1 := int(c.registers.getPC())
	value2 := int(int8(c.read8BitOperand()))
	c.registers.setPC(uint16(value1 + value2))
}

func (c *CPU) instruction19() {
	value1 := c.registers.getHL()
	value2 := c.registers.getDE()
	c.registers.f.setN(false)
	c.registers.f.setH((value1&0xFFF)+(value2&0xFFF) > 0xFFF)
	c.registers.f.setC(uint32(value1)+uint32(value2) > 0xFFFF)
	c.registers.setHL(value1 + value2)
}

func (c *CPU) instruction1A() {
	addr := c.registers.getDE()
	value := c.memory.Read(addr)
	c.registers.setA(value)
}

func (c *CPU) instruction1B() {
	c.registers.decDE()
}

func (c *CPU) instruction1C() {
	value1 := c.registers.getE()
	value2 := value1 + 1
	c.registers.f.setZ(value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(value2&0xF == 0)
	c.registers.incE()
}

func (c *CPU) instruction1D() {
	value1 := c.registers.getE()
	value2 := value1 - 1
	c.registers.f.setZ(value2 == 0)
	c.registers.f.setN(true)
	c.registers.f.setH(value1&0xF == 0)
	c.registers.decE()
}

func (c *CPU) instruction1E() {
	value := c.read8BitOperand()
	c.registers.setE(value)
}

func (c *CPU) instruction1F() {
	carry := c.registers.f.getCarry() << 7
	value := c.registers.getA()
	bit0 := value << 7
	value = (value >> 1) + carry
	c.registers.f.setZ(false)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit0 == 128)
	c.registers.setA(value)
}

func (c *CPU) instruction20() {
	value1 := int(c.registers.getPC())
	value2 := int(int8(c.read8BitOperand()))
	instruction := c.instructions[0x20]
	if !c.registers.f.getZ() {
		c.registers.setPC(uint16(value1 + value2))
		instruction.cycles = 12
	} else {
		instruction.cycles = 8
	}
	c.instructions[0x20] = instruction
}

func (c *CPU) instruction21() {
	value := c.read16BitOperand()
	c.registers.setHL(value)
}

func (c *CPU) instruction22() {
	addr := c.registers.getHL()
	value := c.registers.getA()
	c.memory.Write(addr, value)
	c.registers.incHL()
}

func (c *CPU) instruction23() {
	c.registers.incHL()
}

func (c *CPU) instruction24() {
	value1 := c.registers.getH()
	value2 := value1 + 1
	c.registers.f.setZ(value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(value2&0xF == 0)
	c.registers.incH()
}

func (c *CPU) instruction25() {
	value1 := c.registers.getH()
	value2 := value1 - 1
	c.registers.f.setZ(value2 == 0)
	c.registers.f.setN(true)
	c.registers.f.setH(value1&0xF == 0)
	c.registers.decH()
}

func (c *CPU) instruction26() {
	value := c.read8BitOperand()
	c.registers.setH(value)
}

func (c *CPU) instruction27() {
	value := uint16(c.registers.getA())
	if !c.registers.f.getN() {
		if c.registers.f.getH() || (value&0xF) > 0x9 {
			value += 0x6
		}
		if c.registers.f.getC() || value > 0x9F {
			value += 0x60
			c.registers.f.setC(true)
		}
	} else {
		if c.registers.f.getH() {
			value -= 0x6
		}
		if c.registers.f.getC() {
			value -= 0x60
		}
	}
	c.registers.f.setZ((value & 0xFF) == 0)
	c.registers.f.setH(false)
	c.registers.setA(uint8(value))
}

func (c *CPU) instruction28() {
	value1 := int(c.registers.getPC())
	value2 := int(int8(c.read8BitOperand()))
	instruction := c.instructions[0x28]
	if c.registers.f.getZ() {
		c.registers.setPC(uint16(value1 + value2))
		instruction.cycles = 12
	} else {
		instruction.cycles = 8
	}
	c.instructions[0x28] = instruction
}

func (c *CPU) instruction29() {
	value1 := c.registers.getHL()
	value2 := c.registers.getHL()
	c.registers.f.setN(false)
	c.registers.f.setH((value1&0xFFF)+(value2&0xFFF) > 0xFFF)
	c.registers.f.setC(uint32(value1)+uint32(value2) > 0xFFFF)
	c.registers.setHL(value1 + value2)
}

func (c *CPU) instruction2A() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr)
	c.registers.setA(value)
	c.registers.incHL()
}

func (c *CPU) instruction2B() {
	c.registers.decHL()
}

func (c *CPU) instruction2C() {
	value1 := c.registers.getL()
	value2 := value1 + 1
	c.registers.f.setZ(value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(value2&0xF == 0)
	c.registers.incL()
}

func (c *CPU) instruction2D() {
	value1 := c.registers.getL()
	value2 := value1 - 1
	c.registers.f.setZ(value2 == 0)
	c.registers.f.setN(true)
	c.registers.f.setH(value1&0xF == 0)
	c.registers.decL()
}

func (c *CPU) instruction2E() {
	value := c.read8BitOperand()
	c.registers.setL(value)
}

func (c *CPU) instruction2F() {
	value := c.registers.getA() ^ 0xFF
	c.registers.f.setN(true)
	c.registers.f.setH(true)
	c.registers.setA(value)
}

func (c *CPU) instruction30() {
	value1 := int(c.registers.getPC())
	value2 := int(int8(c.read8BitOperand()))
	instruction := c.instructions[0x30]
	if !c.registers.f.getC() {
		c.registers.setPC(uint16(value1 + value2))
		instruction.cycles = 12
	} else {
		instruction.cycles = 8
	}
	c.instructions[0x30] = instruction
}

func (c *CPU) instruction31() {
	value := c.read16BitOperand()
	c.registers.setSP(value)
}

func (c *CPU) instruction32() {
	addr := c.registers.getHL()
	value := c.registers.getA()
	c.memory.Write(addr, value)
	c.registers.decHL()
}

func (c *CPU) instruction33() {
	c.registers.incSP()
}

func (c *CPU) instruction34() {
	addr := c.registers.getHL()
	value1 := c.memory.Read(addr)
	value2 := value1 + 1
	c.registers.f.setZ(value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(value2&0xF == 0)
	c.memory.Write(addr, value2)
}

func (c *CPU) instruction35() {
	addr := c.registers.getHL()
	value1 := c.memory.Read(addr)
	value2 := value1 - 1
	c.registers.f.setZ(value2 == 0)
	c.registers.f.setN(true)
	c.registers.f.setH(value1&0xF == 0)
	c.memory.Write(addr, value2)
}

func (c *CPU) instruction36() {
	addr := c.registers.getHL()
	value := c.read8BitOperand()
	c.memory.Write(addr, value)
}

func (c *CPU) instruction37() {
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(true)
}

func (c *CPU) instruction38() {
	value1 := int(c.registers.getPC())
	value2 := int(int8(c.read8BitOperand()))
	instruction := c.instructions[0x38]
	if c.registers.f.getC() {
		c.registers.setPC(uint16(value1 + value2))
		instruction.cycles = 12
	} else {
		instruction.cycles = 8
	}
	c.instructions[0x38] = instruction
}

func (c *CPU) instruction39() {
	value1 := c.registers.getHL()
	value2 := c.registers.getSP()
	c.registers.f.setN(false)
	c.registers.f.setH((value1&0xFFF)+(value2&0xFFF) > 0xFFF)
	c.registers.f.setC(uint32(value1)+uint32(value2) > 0xFFFF)
	c.registers.setHL(value1 + value2)
}

func (c *CPU) instruction3A() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr)
	c.registers.setA(value)
	c.registers.decHL()
}

func (c *CPU) instruction3B() {
	c.registers.decSP()
}

func (c *CPU) instruction3C() {
	value1 := c.registers.getA()
	value2 := value1 + 1
	c.registers.f.setZ(value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(value2&0xF == 0)
	c.registers.incA()
}

func (c *CPU) instruction3D() {
	value1 := c.registers.getA()
	value2 := value1 - 1
	c.registers.f.setZ(value2 == 0)
	c.registers.f.setN(true)
	c.registers.f.setH(value1&0xF == 0)
	c.registers.decA()
}

func (c *CPU) instruction3E() {
	value := c.read8BitOperand()
	c.registers.setA(value)
}

func (c *CPU) instruction3F() {
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(!c.registers.f.getC())
}

func (c *CPU) instruction40() {
	value := c.registers.getB()
	c.registers.setB(value)
}

func (c *CPU) instruction41() {
	value := c.registers.getC()
	c.registers.setB(value)
}

func (c *CPU) instruction42() {
	value := c.registers.getD()
	c.registers.setB(value)
}

func (c *CPU) instruction43() {
	value := c.registers.getE()
	c.registers.setB(value)
}

func (c *CPU) instruction44() {
	value := c.registers.getH()
	c.registers.setB(value)
}

func (c *CPU) instruction45() {
	value := c.registers.getL()
	c.registers.setB(value)
}

func (c *CPU) instruction46() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr)
	c.registers.setB(value)
}

func (c *CPU) instruction47() {
	value := c.registers.getA()
	c.registers.setB(value)
}

func (c *CPU) instruction48() {
	value := c.registers.getB()
	c.registers.setC(value)
}

func (c *CPU) instruction49() {
	value := c.registers.getC()
	c.registers.setC(value)
}

func (c *CPU) instruction4A() {
	value := c.registers.getD()
	c.registers.setC(value)
}

func (c *CPU) instruction4B() {
	value := c.registers.getE()
	c.registers.setC(value)
}

func (c *CPU) instruction4C() {
	value := c.registers.getH()
	c.registers.setC(value)
}

func (c *CPU) instruction4D() {
	value := c.registers.getL()
	c.registers.setC(value)
}

func (c *CPU) instruction4E() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr)
	c.registers.setC(value)
}

func (c *CPU) instruction4F() {
	value := c.registers.getA()
	c.registers.setC(value)
}

func (c *CPU) instruction50() {
	value := c.registers.getB()
	c.registers.setD(value)
}

func (c *CPU) instruction51() {
	value := c.registers.getC()
	c.registers.setD(value)
}

func (c *CPU) instruction52() {
	value := c.registers.getD()
	c.registers.setD(value)
}

func (c *CPU) instruction53() {
	value := c.registers.getE()
	c.registers.setD(value)
}

func (c *CPU) instruction54() {
	value := c.registers.getH()
	c.registers.setD(value)
}

func (c *CPU) instruction55() {
	value := c.registers.getL()
	c.registers.setD(value)
}

func (c *CPU) instruction56() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr)
	c.registers.setD(value)
}

func (c *CPU) instruction57() {
	value := c.registers.getA()
	c.registers.setD(value)
}

func (c *CPU) instruction58() {
	value := c.registers.getB()
	c.registers.setE(value)
}

func (c *CPU) instruction59() {
	value := c.registers.getC()
	c.registers.setE(value)
}

func (c *CPU) instruction5A() {
	value := c.registers.getD()
	c.registers.setE(value)
}

func (c *CPU) instruction5B() {
	value := c.registers.getE()
	c.registers.setE(value)
}

func (c *CPU) instruction5C() {
	value := c.registers.getH()
	c.registers.setE(value)
}

func (c *CPU) instruction5D() {
	value := c.registers.getL()
	c.registers.setE(value)
}

func (c *CPU) instruction5E() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr)
	c.registers.setE(value)
}

func (c *CPU) instruction5F() {
	value := c.registers.getA()
	c.registers.setE(value)
}

func (c *CPU) instruction60() {
	value := c.registers.getB()
	c.registers.setH(value)
}

func (c *CPU) instruction61() {
	value := c.registers.getC()
	c.registers.setH(value)
}

func (c *CPU) instruction62() {
	value := c.registers.getD()
	c.registers.setH(value)
}

func (c *CPU) instruction63() {
	value := c.registers.getE()
	c.registers.setH(value)
}

func (c *CPU) instruction64() {
	value := c.registers.getH()
	c.registers.setH(value)
}

func (c *CPU) instruction65() {
	value := c.registers.getL()
	c.registers.setH(value)
}

func (c *CPU) instruction66() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr)
	c.registers.setH(value)
}

func (c *CPU) instruction67() {
	value := c.registers.getA()
	c.registers.setH(value)
}

func (c *CPU) instruction68() {
	value := c.registers.getB()
	c.registers.setL(value)
}

func (c *CPU) instruction69() {
	value := c.registers.getC()
	c.registers.setL(value)
}

func (c *CPU) instruction6A() {
	value := c.registers.getD()
	c.registers.setL(value)
}

func (c *CPU) instruction6B() {
	value := c.registers.getE()
	c.registers.setL(value)
}

func (c *CPU) instruction6C() {
	value := c.registers.getH()
	c.registers.setL(value)
}

func (c *CPU) instruction6D() {
	value := c.registers.getL()
	c.registers.setL(value)
}

func (c *CPU) instruction6E() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr)
	c.registers.setL(value)
}

func (c *CPU) instruction6F() {
	value := c.registers.getA()
	c.registers.setL(value)
}

func (c *CPU) instruction70() {
	addr := c.registers.getHL()
	value := c.registers.getB()
	c.memory.Write(addr, value)
}

func (c *CPU) instruction71() {
	addr := c.registers.getHL()
	value := c.registers.getC()
	c.memory.Write(addr, value)
}

func (c *CPU) instruction72() {
	addr := c.registers.getHL()
	value := c.registers.getD()
	c.memory.Write(addr, value)
}

func (c *CPU) instruction73() {
	addr := c.registers.getHL()
	value := c.registers.getE()
	c.memory.Write(addr, value)
}

func (c *CPU) instruction74() {
	addr := c.registers.getHL()
	value := c.registers.getH()
	c.memory.Write(addr, value)
}

func (c *CPU) instruction75() {
	addr := c.registers.getHL()
	value := c.registers.getL()
	c.memory.Write(addr, value)
}

func (c *CPU) instruction76() {
	c.enableHalt()
}

func (c *CPU) instruction77() {
	addr := c.registers.getHL()
	value := c.registers.getA()
	c.memory.Write(addr, value)
}

func (c *CPU) instruction78() {
	value := c.registers.getB()
	c.registers.setA(value)
}

func (c *CPU) instruction79() {
	value := c.registers.getC()
	c.registers.setA(value)
}

func (c *CPU) instruction7A() {
	value := c.registers.getD()
	c.registers.setA(value)
}

func (c *CPU) instruction7B() {
	value := c.registers.getE()
	c.registers.setA(value)
}

func (c *CPU) instruction7C() {
	value := c.registers.getH()
	c.registers.setA(value)
}

func (c *CPU) instruction7D() {
	value := c.registers.getL()
	c.registers.setA(value)
}

func (c *CPU) instruction7E() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr)
	c.registers.setA(value)
}

func (c *CPU) instruction7F() {
	value := c.registers.getA()
	c.registers.setA(value)
}

func (c *CPU) instruction80() {
	value1 := c.registers.getA()
	value2 := c.registers.getB()
	c.registers.f.setZ(value1+value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH((value1&0xF)+(value2&0xF) > 0xF)
	c.registers.f.setC(uint16(value1)+uint16(value2) > 0xFF)
	c.registers.setA(value1 + value2)
}

func (c *CPU) instruction81() {
	value1 := c.registers.getA()
	value2 := c.registers.getC()
	c.registers.f.setZ(value1+value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH((value1&0xF)+(value2&0xF) > 0xF)
	c.registers.f.setC(uint16(value1)+uint16(value2) > 0xFF)
	c.registers.setA(value1 + value2)
}

func (c *CPU) instruction82() {
	value1 := c.registers.getA()
	value2 := c.registers.getD()
	c.registers.setA(value1 + value2)
	c.registers.f.setZ(value1+value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH((value1&0xF)+(value2&0xF) > 0xF)
	c.registers.f.setC(uint16(value1)+uint16(value2) > 0xFF)
	c.registers.setA(value1 + value2)
}

func (c *CPU) instruction83() {
	value1 := c.registers.getA()
	value2 := c.registers.getE()
	c.registers.setA(value1 + value2)
	c.registers.f.setZ(value1+value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH((value1&0xF)+(value2&0xF) > 0xF)
	c.registers.f.setC(uint16(value1)+uint16(value2) > 0xFF)
	c.registers.setA(value1 + value2)
}

func (c *CPU) instruction84() {
	value1 := c.registers.getA()
	value2 := c.registers.getH()
	c.registers.setA(value1 + value2)
	c.registers.f.setZ(value1+value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH((value1&0xF)+(value2&0xF) > 0xF)
	c.registers.f.setC(uint16(value1)+uint16(value2) > 0xFF)
	c.registers.setA(value1 + value2)
}

func (c *CPU) instruction85() {
	value1 := c.registers.getA()
	value2 := c.registers.getL()
	c.registers.setA(value1 + value2)
	c.registers.f.setZ(value1+value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH((value1&0xF)+(value2&0xF) > 0xF)
	c.registers.f.setC(uint16(value1)+uint16(value2) > 0xFF)
	c.registers.setA(value1 + value2)
}

func (c *CPU) instruction86() {
	addr := c.registers.getHL()
	value1 := c.registers.getA()
	value2 := c.memory.Read(addr)
	c.registers.f.setZ(value1+value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH((value1&0xF)+(value2&0xF) > 0xF)
	c.registers.f.setC(uint16(value1)+uint16(value2) > 0xFF)
	c.registers.setA(value1 + value2)
}

func (c *CPU) instruction87() {
	value1 := c.registers.getA()
	value2 := c.registers.getA()
	c.registers.setA(value1 + value2)
	c.registers.f.setZ(value1+value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH((value1&0xF)+(value2&0xF) > 0xF)
	c.registers.f.setC(uint16(value1)+uint16(value2) > 0xFF)
	c.registers.setA(value1 + value2)
}

func (c *CPU) instruction88() {
	carry := c.registers.f.getCarry()
	value1 := c.registers.getA()
	value2 := c.registers.getB()
	c.registers.f.setZ(value1+value2+carry == 0)
	c.registers.f.setN(false)
	c.registers.f.setH((value1&0xF)+(value2&0xF)+carry > 0xF)
	c.registers.f.setC(uint16(value1)+uint16(value2)+uint16(carry) > 0xFF)
	c.registers.setA(value1 + value2 + carry)
}

func (c *CPU) instruction89() {
	carry := c.registers.f.getCarry()
	value1 := c.registers.getA()
	value2 := c.registers.getC()
	c.registers.f.setZ(value1+value2+carry == 0)
	c.registers.f.setN(false)
	c.registers.f.setH((value1&0xF)+(value2&0xF)+carry > 0xF)
	c.registers.f.setC(uint16(value1)+uint16(value2)+uint16(carry) > 0xFF)
	c.registers.setA(value1 + value2 + carry)
}

func (c *CPU) instruction8A() {
	carry := c.registers.f.getCarry()
	value1 := c.registers.getA()
	value2 := c.registers.getD()
	c.registers.f.setZ(value1+value2+carry == 0)
	c.registers.f.setN(false)
	c.registers.f.setH((value1&0xF)+(value2&0xF)+carry > 0xF)
	c.registers.f.setC(uint16(value1)+uint16(value2)+uint16(carry) > 0xFF)
	c.registers.setA(value1 + value2 + carry)
}

func (c *CPU) instruction8B() {
	carry := c.registers.f.getCarry()
	value1 := c.registers.getA()
	value2 := c.registers.getE()
	c.registers.f.setZ(value1+value2+carry == 0)
	c.registers.f.setN(false)
	c.registers.f.setH((value1&0xF)+(value2&0xF)+carry > 0xF)
	c.registers.f.setC(uint16(value1)+uint16(value2)+uint16(carry) > 0xFF)
	c.registers.setA(value1 + value2 + carry)
}

func (c *CPU) instruction8C() {
	carry := c.registers.f.getCarry()
	value1 := c.registers.getA()
	value2 := c.registers.getH()
	c.registers.f.setZ(value1+value2+carry == 0)
	c.registers.f.setN(false)
	c.registers.f.setH((value1&0xF)+(value2&0xF)+carry > 0xF)
	c.registers.f.setC(uint16(value1)+uint16(value2)+uint16(carry) > 0xFF)
	c.registers.setA(value1 + value2 + carry)
}

func (c *CPU) instruction8D() {
	carry := c.registers.f.getCarry()
	value1 := c.registers.getA()
	value2 := c.registers.getL()
	c.registers.f.setZ(value1+value2+carry == 0)
	c.registers.f.setN(false)
	c.registers.f.setH((value1&0xF)+(value2&0xF)+carry > 0xF)
	c.registers.f.setC(uint16(value1)+uint16(value2)+uint16(carry) > 0xFF)
	c.registers.setA(value1 + value2 + carry)
}

func (c *CPU) instruction8E() {
	carry := c.registers.f.getCarry()
	addr := c.registers.getHL()
	value1 := c.registers.getA()
	value2 := c.memory.Read(addr)
	c.registers.f.setZ(value1+value2+carry == 0)
	c.registers.f.setN(false)
	c.registers.f.setH((value1&0xF)+(value2&0xF)+carry > 0xF)
	c.registers.f.setC(uint16(value1)+uint16(value2)+uint16(carry) > 0xFF)
	c.registers.setA(value1 + value2 + carry)
}

func (c *CPU) instruction8F() {
	carry := c.registers.f.getCarry()
	value1 := c.registers.getA()
	value2 := c.registers.getA()
	c.registers.f.setZ(value1+value2+carry == 0)
	c.registers.f.setN(false)
	c.registers.f.setH((value1&0xF)+(value2&0xF)+carry > 0xF)
	c.registers.f.setC(uint16(value1)+uint16(value2)+uint16(carry) > 0xFF)
	c.registers.setA(value1 + value2 + carry)
}

func (c *CPU) instruction90() {
	value1 := c.registers.getA()
	value2 := c.registers.getB()
	c.registers.f.setZ(value1-value2 == 0)
	c.registers.f.setN(true)
	c.registers.f.setH((value1 & 0xF) < (value2 & 0xF))
	c.registers.f.setC(uint16(value1) < uint16(value2))
	c.registers.setA(value1 - value2)
}

func (c *CPU) instruction91() {
	value1 := c.registers.getA()
	value2 := c.registers.getC()
	c.registers.f.setZ(value1-value2 == 0)
	c.registers.f.setN(true)
	c.registers.f.setH((value1 & 0xF) < (value2 & 0xF))
	c.registers.f.setC(uint16(value1) < uint16(value2))
	c.registers.setA(value1 - value2)
}

func (c *CPU) instruction92() {
	value1 := c.registers.getA()
	value2 := c.registers.getD()
	c.registers.f.setZ(value1-value2 == 0)
	c.registers.f.setN(true)
	c.registers.f.setH((value1 & 0xF) < (value2 & 0xF))
	c.registers.f.setC(uint16(value1) < uint16(value2))
	c.registers.setA(value1 - value2)
}

func (c *CPU) instruction93() {
	value1 := c.registers.getA()
	value2 := c.registers.getE()
	c.registers.f.setZ(value1-value2 == 0)
	c.registers.f.setN(true)
	c.registers.f.setH((value1 & 0xF) < (value2 & 0xF))
	c.registers.f.setC(uint16(value1) < uint16(value2))
	c.registers.setA(value1 - value2)
}

func (c *CPU) instruction94() {
	value1 := c.registers.getA()
	value2 := c.registers.getH()
	c.registers.f.setZ(value1-value2 == 0)
	c.registers.f.setN(true)
	c.registers.f.setH((value1 & 0xF) < (value2 & 0xF))
	c.registers.f.setC(uint16(value1) < uint16(value2))
	c.registers.setA(value1 - value2)
}

func (c *CPU) instruction95() {
	value1 := c.registers.getA()
	value2 := c.registers.getL()
	c.registers.f.setZ(value1-value2 == 0)
	c.registers.f.setN(true)
	c.registers.f.setH((value1 & 0xF) < (value2 & 0xF))
	c.registers.f.setC(uint16(value1) < uint16(value2))
	c.registers.setA(value1 - value2)
}

func (c *CPU) instruction96() {
	addr := c.registers.getHL()
	value1 := c.registers.getA()
	value2 := c.memory.Read(addr)
	c.registers.f.setZ(value1-value2 == 0)
	c.registers.f.setN(true)
	c.registers.f.setH((value1 & 0xF) < (value2 & 0xF))
	c.registers.f.setC(uint16(value1) < uint16(value2))
	c.registers.setA(value1 - value2)
}

func (c *CPU) instruction97() {
	value1 := c.registers.getA()
	value2 := c.registers.getA()
	c.registers.f.setZ(value1-value2 == 0)
	c.registers.f.setN(true)
	c.registers.f.setH((value1 & 0xF) < (value2 & 0xF))
	c.registers.f.setC(uint16(value1) < uint16(value2))
	c.registers.setA(value1 - value2)
}

func (c *CPU) instruction98() {
	carry := c.registers.f.getCarry()
	value1 := c.registers.getA()
	value2 := c.registers.getB()
	c.registers.f.setZ(value1-value2-carry == 0)
	c.registers.f.setN(true)
	c.registers.f.setH((value1 & 0xF) < ((value2 & 0xF) + carry))
	c.registers.f.setC(uint16(value1) < (uint16(value2) + uint16(carry)))
	c.registers.setA(value1 - value2 - carry)
}

func (c *CPU) instruction99() {
	carry := c.registers.f.getCarry()
	value1 := c.registers.getA()
	value2 := c.registers.getC()
	c.registers.f.setZ(value1-value2-carry == 0)
	c.registers.f.setN(true)
	c.registers.f.setH((value1 & 0xF) < ((value2 & 0xF) + carry))
	c.registers.f.setC(uint16(value1) < (uint16(value2) + uint16(carry)))
	c.registers.setA(value1 - value2 - carry)
}

func (c *CPU) instruction9A() {
	carry := c.registers.f.getCarry()
	value1 := c.registers.getA()
	value2 := c.registers.getD()
	c.registers.f.setZ(value1-value2-carry == 0)
	c.registers.f.setN(true)
	c.registers.f.setH((value1 & 0xF) < ((value2 & 0xF) + carry))
	c.registers.f.setC(uint16(value1) < (uint16(value2) + uint16(carry)))
	c.registers.setA(value1 - value2 - carry)
}

func (c *CPU) instruction9B() {
	carry := c.registers.f.getCarry()
	value1 := c.registers.getA()
	value2 := c.registers.getE()
	c.registers.f.setZ(value1-value2-carry == 0)
	c.registers.f.setN(true)
	c.registers.f.setH((value1 & 0xF) < ((value2 & 0xF) + carry))
	c.registers.f.setC(uint16(value1) < (uint16(value2) + uint16(carry)))
	c.registers.setA(value1 - value2 - carry)
}

func (c *CPU) instruction9C() {
	carry := c.registers.f.getCarry()
	value1 := c.registers.getA()
	value2 := c.registers.getH()
	c.registers.f.setZ(value1-value2-carry == 0)
	c.registers.f.setN(true)
	c.registers.f.setH((value1 & 0xF) < ((value2 & 0xF) + carry))
	c.registers.f.setC(uint16(value1) < (uint16(value2) + uint16(carry)))
	c.registers.setA(value1 - value2 - carry)
}

func (c *CPU) instruction9D() {
	carry := c.registers.f.getCarry()
	value1 := c.registers.getA()
	value2 := c.registers.getL()
	c.registers.f.setZ(value1-value2-carry == 0)
	c.registers.f.setN(true)
	c.registers.f.setH((value1 & 0xF) < ((value2 & 0xF) + carry))
	c.registers.f.setC(uint16(value1) < (uint16(value2) + uint16(carry)))
	c.registers.setA(value1 - value2 - carry)
}

func (c *CPU) instruction9E() {
	carry := c.registers.f.getCarry()
	addr := c.registers.getHL()
	value1 := c.registers.getA()
	value2 := c.memory.Read(addr)
	c.registers.f.setZ(value1-value2-carry == 0)
	c.registers.f.setN(true)
	c.registers.f.setH((value1 & 0xF) < ((value2 & 0xF) + carry))
	c.registers.f.setC(uint16(value1) < (uint16(value2) + uint16(carry)))
	c.registers.setA(value1 - value2 - carry)
}

func (c *CPU) instruction9F() {
	carry := c.registers.f.getCarry()
	value1 := c.registers.getA()
	value2 := c.registers.getA()
	c.registers.f.setZ(value1-value2-carry == 0)
	c.registers.f.setN(true)
	c.registers.f.setH((value1 & 0xF) < ((value2 & 0xF) + carry))
	c.registers.f.setC(uint16(value1) < (uint16(value2) + uint16(carry)))
	c.registers.setA(value1 - value2 - carry)
}

func (c *CPU) instructionA0() {
	value1 := c.registers.getA()
	value2 := c.registers.getB()
	c.registers.f.setZ(value1&value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
	c.registers.f.setC(false)
	c.registers.setA(value1 & value2)
}

func (c *CPU) instructionA1() {
	value1 := c.registers.getA()
	value2 := c.registers.getC()
	c.registers.f.setZ(value1&value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
	c.registers.f.setC(false)
	c.registers.setA(value1 & value2)
}

func (c *CPU) instructionA2() {
	value1 := c.registers.getA()
	value2 := c.registers.getD()
	c.registers.f.setZ(value1&value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
	c.registers.f.setC(false)
	c.registers.setA(value1 & value2)
}

func (c *CPU) instructionA3() {
	value1 := c.registers.getA()
	value2 := c.registers.getE()
	c.registers.f.setZ(value1&value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
	c.registers.f.setC(false)
	c.registers.setA(value1 & value2)
}

func (c *CPU) instructionA4() {
	value1 := c.registers.getA()
	value2 := c.registers.getH()
	c.registers.f.setZ(value1&value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
	c.registers.f.setC(false)
	c.registers.setA(value1 & value2)
}

func (c *CPU) instructionA5() {
	value1 := c.registers.getA()
	value2 := c.registers.getL()
	c.registers.f.setZ(value1&value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
	c.registers.f.setC(false)
	c.registers.setA(value1 & value2)
}

func (c *CPU) instructionA6() {
	addr := c.registers.getHL()
	value1 := c.registers.getA()
	value2 := c.memory.Read(addr)
	c.registers.f.setZ(value1&value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
	c.registers.f.setC(false)
	c.registers.setA(value1 & value2)
}

func (c *CPU) instructionA7() {
	value1 := c.registers.getA()
	value2 := c.registers.getA()
	c.registers.f.setZ(value1&value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
	c.registers.f.setC(false)
	c.registers.setA(value1 & value2)
}

func (c *CPU) instructionA8() {
	value1 := c.registers.getA()
	value2 := c.registers.getB()
	c.registers.f.setZ(value1^value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(false)
	c.registers.setA(value1 ^ value2)
}

func (c *CPU) instructionA9() {
	value1 := c.registers.getA()
	value2 := c.registers.getC()
	c.registers.f.setZ(value1^value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(false)
	c.registers.setA(value1 ^ value2)
}

func (c *CPU) instructionAA() {
	value1 := c.registers.getA()
	value2 := c.registers.getD()
	c.registers.f.setZ(value1^value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(false)
	c.registers.setA(value1 ^ value2)
}

func (c *CPU) instructionAB() {
	value1 := c.registers.getA()
	value2 := c.registers.getE()
	c.registers.f.setZ(value1^value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(false)
	c.registers.setA(value1 ^ value2)
}

func (c *CPU) instructionAC() {
	value1 := c.registers.getA()
	value2 := c.registers.getH()
	c.registers.f.setZ(value1^value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(false)
	c.registers.setA(value1 ^ value2)
}

func (c *CPU) instructionAD() {
	value1 := c.registers.getA()
	value2 := c.registers.getL()
	c.registers.f.setZ(value1^value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(false)
	c.registers.setA(value1 ^ value2)
}

func (c *CPU) instructionAE() {
	addr := c.registers.getHL()
	value1 := c.registers.getA()
	value2 := c.memory.Read(addr)
	c.registers.f.setZ(value1^value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(false)
	c.registers.setA(value1 ^ value2)
}

func (c *CPU) instructionAF() {
	value1 := c.registers.getA()
	value2 := c.registers.getA()
	c.registers.f.setZ(value1^value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(false)
	c.registers.setA(value1 ^ value2)
}

func (c *CPU) instructionB0() {
	value1 := c.registers.getA()
	value2 := c.registers.getB()
	c.registers.f.setZ(value1|value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(false)
	c.registers.setA(value1 | value2)
}

func (c *CPU) instructionB1() {
	value1 := c.registers.getA()
	value2 := c.registers.getC()
	c.registers.f.setZ(value1|value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(false)
	c.registers.setA(value1 | value2)
}

func (c *CPU) instructionB2() {
	value1 := c.registers.getA()
	value2 := c.registers.getD()
	c.registers.f.setZ(value1|value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(false)
	c.registers.setA(value1 | value2)
}

func (c *CPU) instructionB3() {
	value1 := c.registers.getA()
	value2 := c.registers.getE()
	c.registers.f.setZ(value1|value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(false)
	c.registers.setA(value1 | value2)
}

func (c *CPU) instructionB4() {
	value1 := c.registers.getA()
	value2 := c.registers.getH()
	c.registers.f.setZ(value1|value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(false)
	c.registers.setA(value1 | value2)
}

func (c *CPU) instructionB5() {
	value1 := c.registers.getA()
	value2 := c.registers.getL()
	c.registers.f.setZ(value1|value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(false)
	c.registers.setA(value1 | value2)
}

func (c *CPU) instructionB6() {
	addr := c.registers.getHL()
	value1 := c.registers.getA()
	value2 := c.memory.Read(addr)
	c.registers.f.setZ(value1|value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(false)
	c.registers.setA(value1 | value2)
}

func (c *CPU) instructionB7() {
	value1 := c.registers.getA()
	value2 := c.registers.getA()
	c.registers.f.setZ(value1|value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(false)
	c.registers.setA(value1 | value2)
}

func (c *CPU) instructionB8() {
	value1 := c.registers.getA()
	value2 := c.registers.getB()
	c.registers.f.setZ(value1 == value2)
	c.registers.f.setN(true)
	c.registers.f.setH((value1 & 0xF) < (value2 & 0xF))
	c.registers.f.setC(value1 < value2)
}

func (c *CPU) instructionB9() {
	value1 := c.registers.getA()
	value2 := c.registers.getC()
	c.registers.f.setZ(value1 == value2)
	c.registers.f.setN(true)
	c.registers.f.setH((value1 & 0xF) < (value2 & 0xF))
	c.registers.f.setC(value1 < value2)
}

func (c *CPU) instructionBA() {
	value1 := c.registers.getA()
	value2 := c.registers.getD()
	c.registers.f.setZ(value1 == value2)
	c.registers.f.setN(true)
	c.registers.f.setH((value1 & 0xF) < (value2 & 0xF))
	c.registers.f.setC(value1 < value2)
}

func (c *CPU) instructionBB() {
	value1 := c.registers.getA()
	value2 := c.registers.getE()
	c.registers.f.setZ(value1 == value2)
	c.registers.f.setN(true)
	c.registers.f.setH((value1 & 0xF) < (value2 & 0xF))
	c.registers.f.setC(value1 < value2)
}

func (c *CPU) instructionBC() {
	value1 := c.registers.getA()
	value2 := c.registers.getH()
	c.registers.f.setZ(value1 == value2)
	c.registers.f.setN(true)
	c.registers.f.setH((value1 & 0xF) < (value2 & 0xF))
	c.registers.f.setC(value1 < value2)
}

func (c *CPU) instructionBD() {
	value1 := c.registers.getA()
	value2 := c.registers.getL()
	c.registers.f.setZ(value1 == value2)
	c.registers.f.setN(true)
	c.registers.f.setH((value1 & 0xF) < (value2 & 0xF))
	c.registers.f.setC(value1 < value2)
}

func (c *CPU) instructionBE() {
	addr := c.registers.getHL()
	value1 := c.registers.getA()
	value2 := c.memory.Read(addr)
	c.registers.f.setZ(value1 == value2)
	c.registers.f.setN(true)
	c.registers.f.setH((value1 & 0xF) < (value2 & 0xF))
	c.registers.f.setC(value1 < value2)
}

func (c *CPU) instructionBF() {
	value1 := c.registers.getA()
	value2 := c.registers.getA()
	c.registers.f.setZ(value1 == value2)
	c.registers.f.setN(true)
	c.registers.f.setH((value1 & 0xF) < (value2 & 0xF))
	c.registers.f.setC(value1 < value2)
}

func (c *CPU) instructionC0() {
	instruction := c.instructions[0xC0]
	if !c.registers.f.getZ() {
		addr := c.registers.getSP()
		msb := uint16(c.memory.Read(addr + 1))
		lsb := uint16(c.memory.Read(addr))
		value := msb<<8 | lsb
		c.registers.setPC(value)
		c.registers.setSP(addr + 2)
		instruction.cycles = 20
	} else {
		instruction.cycles = 8
	}
	c.instructions[0xC0] = instruction
}

func (c *CPU) instructionC1() {
	addr := c.registers.getSP()
	msb := uint16(c.memory.Read(addr + 1))
	lsb := uint16(c.memory.Read(addr))
	value := msb<<8 | lsb
	c.registers.setBC(value)
	c.registers.setSP(addr + 2)
}

func (c *CPU) instructionC2() {
	value := c.read16BitOperand()
	instruction := c.instructions[0xC2]
	if !c.registers.f.getZ() {
		c.registers.setPC(value)
		instruction.cycles = 16
	} else {
		instruction.cycles = 12
	}
	c.instructions[0xC2] = instruction
}

func (c *CPU) instructionC3() {
	value := c.read16BitOperand()
	c.registers.setPC(value)
}

func (c *CPU) instructionC4() {
	value := c.read16BitOperand()
	instruction := c.instructions[0xC4]
	if !c.registers.f.getZ() {
		addr := c.registers.getSP()
		msb := uint8(c.registers.getPC() >> 8)
		lsb := uint8(c.registers.getPC() & 0xFF)
		c.memory.Write(addr-1, msb)
		c.memory.Write(addr-2, lsb)
		c.registers.setSP(addr - 2)
		c.registers.setPC(value)
		instruction.cycles = 24
	} else {
		instruction.cycles = 12
	}
	c.instructions[0xC4] = instruction
}

func (c *CPU) instructionC5() {
	addr := c.registers.getSP()
	msb := c.registers.getB()
	lsb := c.registers.getC()
	c.memory.Write(addr-1, msb)
	c.memory.Write(addr-2, lsb)
	c.registers.setSP(addr - 2)
}

func (c *CPU) instructionC6() {
	value1 := c.registers.getA()
	value2 := c.read8BitOperand()
	c.registers.f.setZ(value1+value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH((value1&0xF)+(value2&0xF) > 0xF)
	c.registers.f.setC(uint16(value1)+uint16(value2) > 0xFF)
	c.registers.setA(value1 + value2)
}

func (c *CPU) instructionC7() {
	value := uint16(0x00)
	addr := c.registers.getSP()
	msb := uint8(c.registers.getPC() >> 8)
	lsb := uint8(c.registers.getPC() & 0xFF)
	c.memory.Write(addr-1, msb)
	c.memory.Write(addr-2, lsb)
	c.registers.setSP(addr - 2)
	c.registers.setPC(value)
}

func (c *CPU) instructionC8() {
	instruction := c.instructions[0xC8]
	if c.registers.f.getZ() {
		addr := c.registers.getSP()
		msb := uint16(c.memory.Read(addr + 1))
		lsb := uint16(c.memory.Read(addr))
		value := msb<<8 | lsb
		c.registers.setPC(value)
		c.registers.setSP(addr + 2)
		instruction.cycles = 20
	} else {
		instruction.cycles = 8
	}
	c.instructions[0xC8] = instruction
}

func (c *CPU) instructionC9() {
	addr := c.registers.getSP()
	msb := uint16(c.memory.Read(addr + 1))
	lsb := uint16(c.memory.Read(addr))
	value := msb<<8 | lsb
	c.registers.setPC(value)
	c.registers.setSP(addr + 2)
}

func (c *CPU) instructionCA() {
	value := c.read16BitOperand()
	instruction := c.instructions[0xCA]
	if c.registers.f.getZ() {
		c.registers.setPC(value)
		instruction.cycles = 16
	} else {
		instruction.cycles = 12
	}
	c.instructions[0xCA] = instruction
}

func (c *CPU) instructionCB() {
	// PREFIX CB
}

func (c *CPU) instructionCC() {
	value := c.read16BitOperand()
	instruction := c.instructions[0xCC]
	if c.registers.f.getZ() {
		addr := c.registers.getSP()
		msb := uint8(c.registers.getPC() >> 8)
		lsb := uint8(c.registers.getPC() & 0xFF)
		c.memory.Write(addr-1, msb)
		c.memory.Write(addr-2, lsb)
		c.registers.setSP(addr - 2)
		c.registers.setPC(value)
		instruction.cycles = 24
	} else {
		instruction.cycles = 12
	}
	c.instructions[0xCC] = instruction
}

func (c *CPU) instructionCD() {
	value := c.read16BitOperand()
	addr := c.registers.getSP()
	msb := uint8(c.registers.getPC() >> 8)
	lsb := uint8(c.registers.getPC() & 0xFF)
	c.memory.Write(addr-1, msb)
	c.memory.Write(addr-2, lsb)
	c.registers.setSP(addr - 2)
	c.registers.setPC(value)
}

func (c *CPU) instructionCE() {
	carry := c.registers.f.getCarry()
	value1 := c.registers.getA()
	value2 := c.read8BitOperand()
	c.registers.f.setZ(value1+value2+carry == 0)
	c.registers.f.setN(false)
	c.registers.f.setH((value1&0xF)+(value2&0xF)+carry > 0xF)
	c.registers.f.setC(uint16(value1)+uint16(value2)+uint16(carry) > 0xFF)
	c.registers.setA(value1 + value2 + carry)
}

func (c *CPU) instructionCF() {
	value := uint16(0x08)
	addr := c.registers.getSP()
	msb := uint8(c.registers.getPC() >> 8)
	lsb := uint8(c.registers.getPC() & 0xFF)
	c.memory.Write(addr-1, msb)
	c.memory.Write(addr-2, lsb)
	c.registers.setSP(addr - 2)
	c.registers.setPC(value)
}

func (c *CPU) instructionD0() {
	instruction := c.instructions[0xD0]
	if !c.registers.f.getC() {
		addr := c.registers.getSP()
		msb := uint16(c.memory.Read(addr + 1))
		lsb := uint16(c.memory.Read(addr))
		value := msb<<8 | lsb
		c.registers.setPC(value)
		c.registers.setSP(addr + 2)
		instruction.cycles = 20
	} else {
		instruction.cycles = 8
	}
	c.instructions[0xD0] = instruction
}

func (c *CPU) instructionD1() {
	addr := c.registers.getSP()
	msb := uint16(c.memory.Read(addr + 1))
	lsb := uint16(c.memory.Read(addr))
	value := msb<<8 | lsb
	c.registers.setDE(value)
	c.registers.setSP(addr + 2)
}

func (c *CPU) instructionD2() {
	value := c.read16BitOperand()
	instruction := c.instructions[0xD2]
	if !c.registers.f.getC() {
		c.registers.setPC(value)
		instruction.cycles = 16
	} else {
		instruction.cycles = 12
	}
	c.instructions[0xD2] = instruction
}

func (c *CPU) instructionD4() {
	value := c.read16BitOperand()
	instruction := c.instructions[0xD4]
	if !c.registers.f.getC() {
		addr := c.registers.getSP()
		msb := uint8(c.registers.getPC() >> 8)
		lsb := uint8(c.registers.getPC() & 0xFF)
		c.memory.Write(addr-1, msb)
		c.memory.Write(addr-2, lsb)
		c.registers.setSP(addr - 2)
		c.registers.setPC(value)
		instruction.cycles = 24
	} else {
		instruction.cycles = 12
	}
	c.instructions[0xD4] = instruction
}

func (c *CPU) instructionD5() {
	addr := c.registers.getSP()
	msb := c.registers.getD()
	lsb := c.registers.getE()
	c.memory.Write(addr-1, msb)
	c.memory.Write(addr-2, lsb)
	c.registers.setSP(addr - 2)
}

func (c *CPU) instructionD6() {
	value1 := c.registers.getA()
	value2 := c.read8BitOperand()
	c.registers.f.setZ(value1-value2 == 0)
	c.registers.f.setN(true)
	c.registers.f.setH((value1 & 0xF) < (value2 & 0xF))
	c.registers.f.setC(uint16(value1) < uint16(value2))
	c.registers.setA(value1 - value2)
}

func (c *CPU) instructionD7() {
	value := uint16(0x10)
	addr := c.registers.getSP()
	msb := uint8(c.registers.getPC() >> 8)
	lsb := uint8(c.registers.getPC() & 0xFF)
	c.memory.Write(addr-1, msb)
	c.memory.Write(addr-2, lsb)
	c.registers.setSP(addr - 2)
	c.registers.setPC(value)
}

func (c *CPU) instructionD8() {
	instruction := c.instructions[0xD8]
	if c.registers.f.getC() {
		addr := c.registers.getSP()
		msb := uint16(c.memory.Read(addr + 1))
		lsb := uint16(c.memory.Read(addr))
		value := msb<<8 | lsb
		c.registers.setPC(value)
		c.registers.setSP(addr + 2)
		instruction.cycles = 20
	} else {
		instruction.cycles = 8
	}
	c.instructions[0xD8] = instruction
}

func (c *CPU) instructionD9() {
	addr := c.registers.getSP()
	msb := uint16(c.memory.Read(addr + 1))
	lsb := uint16(c.memory.Read(addr))
	value := msb<<8 | lsb
	c.registers.setPC(value)
	c.registers.setSP(addr + 2)
	c.interrupts.EnableMaster()
}

func (c *CPU) instructionDA() {
	value := c.read16BitOperand()
	instruction := c.instructions[0xDA]
	if c.registers.f.getC() {
		c.registers.setPC(value)
		instruction.cycles = 16
	} else {
		instruction.cycles = 12
	}
	c.instructions[0xDA] = instruction
}

func (c *CPU) instructionDC() {
	value := c.read16BitOperand()
	instruction := c.instructions[0xDC]
	if c.registers.f.getC() {
		addr := c.registers.getSP()
		msb := uint8(c.registers.getPC() >> 8)
		lsb := uint8(c.registers.getPC() & 0xFF)
		c.memory.Write(addr-1, msb)
		c.memory.Write(addr-2, lsb)
		c.registers.setSP(addr - 2)
		c.registers.setPC(value)
		instruction.cycles = 24
	} else {
		instruction.cycles = 12
	}
	c.instructions[0xDC] = instruction
}

func (c *CPU) instructionDE() {
	carry := c.registers.f.getCarry()
	value1 := c.registers.getA()
	value2 := c.read8BitOperand()
	c.registers.f.setZ(value1-value2-carry == 0)
	c.registers.f.setN(true)
	c.registers.f.setH((value1 & 0xF) < ((value2 & 0xF) + carry))
	c.registers.f.setC(uint16(value1) < (uint16(value2) + uint16(carry)))
	c.registers.setA(value1 - value2 - carry)
}

func (c *CPU) instructionDF() {
	value := uint16(0x18)
	addr := c.registers.getSP()
	msb := uint8(c.registers.getPC() >> 8)
	lsb := uint8(c.registers.getPC() & 0xFF)
	c.memory.Write(addr-1, msb)
	c.memory.Write(addr-2, lsb)
	c.registers.setSP(addr - 2)
	c.registers.setPC(value)
}

func (c *CPU) instructionE0() {
	value := c.registers.getA()
	msb := uint16(0xFF)
	lsb := uint16(c.read8BitOperand())
	addr := msb<<8 | lsb
	c.memory.Write(addr, value)
}

func (c *CPU) instructionE1() {
	addr := c.registers.getSP()
	msb := uint16(c.memory.Read(addr + 1))
	lsb := uint16(c.memory.Read(addr))
	value := msb<<8 | lsb
	c.registers.setHL(value)
	c.registers.setSP(addr + 2)
}

func (c *CPU) instructionE2() {
	value := c.registers.getA()
	msb := uint16(0xFF)
	lsb := uint16(c.registers.getC())
	addr := msb<<8 | lsb
	c.memory.Write(addr, value)
}

func (c *CPU) instructionE5() {
	addr := c.registers.getSP()
	msb := c.registers.getH()
	lsb := c.registers.getL()
	c.memory.Write(addr-1, msb)
	c.memory.Write(addr-2, lsb)
	c.registers.setSP(addr - 2)
}

func (c *CPU) instructionE6() {
	value1 := c.registers.getA()
	value2 := c.read8BitOperand()
	c.registers.f.setZ(value1&value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
	c.registers.f.setC(false)
	c.registers.setA(value1 & value2)
}

func (c *CPU) instructionE7() {
	value := uint16(0x20)
	addr := c.registers.getSP()
	msb := uint8(c.registers.getPC() >> 8)
	lsb := uint8(c.registers.getPC() & 0xFF)
	c.memory.Write(addr-1, msb)
	c.memory.Write(addr-2, lsb)
	c.registers.setSP(addr - 2)
	c.registers.setPC(value)
}

func (c *CPU) instructionE8() {
	value1 := int(c.registers.getSP())
	value2 := int(c.read8BitOperand())
	value2 = (value2 & 127) - (value2 & 128)
	c.registers.f.setZ(false)
	c.registers.f.setN(false)
	if value2 < 0 {
		c.registers.f.setH((value1 & 0xF) < (value2 & 0xF))
		c.registers.f.setC((value1 & 0xFF) < (value2 & 0xFF))
	} else {
		c.registers.f.setH((value1&0xF)+(value2&0xF) > 0xF)
		c.registers.f.setC((value1&0xFF)+(value2&0xFF) > 0xFF)
	}
	c.registers.setSP(uint16(value1 + value2))
}

func (c *CPU) instructionE9() {
	value := c.registers.getHL()
	c.registers.setPC(value)
}

func (c *CPU) instructionEA() {
	addr := c.read16BitOperand()
	value := c.registers.getA()
	c.memory.Write(addr, value)
}

func (c *CPU) instructionEE() {
	value1 := c.registers.getA()
	value2 := c.read8BitOperand()
	c.registers.f.setZ(value1^value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(false)
	c.registers.setA(value1 ^ value2)
}

func (c *CPU) instructionEF() {
	value := uint16(0x28)
	addr := c.registers.getSP()
	msb := uint8(c.registers.getPC() >> 8)
	lsb := uint8(c.registers.getPC() & 0xFF)
	c.memory.Write(addr-1, msb)
	c.memory.Write(addr-2, lsb)
	c.registers.setSP(addr - 2)
	c.registers.setPC(value)
}

func (c *CPU) instructionF0() {
	msb := uint16(0xFF)
	lsb := uint16(c.read8BitOperand())
	addr := msb<<8 | lsb
	value := c.memory.Read(addr)
	c.registers.setA(value)
}

func (c *CPU) instructionF1() {
	addr := c.registers.getSP()
	msb := uint16(c.memory.Read(addr + 1))
	lsb := uint16(c.memory.Read(addr))
	value := msb<<8 | lsb
	c.registers.setAF(value)
	c.registers.setSP(addr + 2)
}

func (c *CPU) instructionF2() {
	msb := uint16(0xFF)
	lsb := uint16(c.registers.getC())
	addr := msb<<8 | lsb
	value := c.memory.Read(addr)
	c.registers.setA(value)
}

func (c *CPU) instructionF3() {
	c.interrupts.DisableMaster()
}

func (c *CPU) instructionF5() {
	addr := c.registers.getSP()
	msb := c.registers.getA()
	lsb := c.registers.getF()
	c.memory.Write(addr-1, msb)
	c.memory.Write(addr-2, lsb)
	c.registers.setSP(addr - 2)
}

func (c *CPU) instructionF6() {
	value1 := c.registers.getA()
	value2 := c.read8BitOperand()
	c.registers.f.setZ(value1|value2 == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(false)
	c.registers.setA(value1 | value2)
}

func (c *CPU) instructionF7() {
	value := uint16(0x30)
	addr := c.registers.getSP()
	msb := uint8(c.registers.getPC() >> 8)
	lsb := uint8(c.registers.getPC() & 0xFF)
	c.memory.Write(addr-1, msb)
	c.memory.Write(addr-2, lsb)
	c.registers.setSP(addr - 2)
	c.registers.setPC(value)
}

func (c *CPU) instructionF8() {
	value1 := int(c.registers.getSP())
	value2 := int(int8(c.read8BitOperand()))
	c.registers.f.setZ(false)
	c.registers.f.setN(false)
	if value2 < 0 {
		c.registers.f.setH((value1 & 0xF) < (value2 & 0xF))
		c.registers.f.setC((value1 & 0xFF) < (value2 & 0xFF))
	} else {
		c.registers.f.setH((value1&0xF)+(value2&0xF) > 0xF)
		c.registers.f.setC((value1&0xFF)+(value2&0xFF) > 0xFF)
	}
	c.registers.setHL(uint16(value1 + value2))
}

func (c *CPU) instructionF9() {
	value := c.registers.getHL()
	c.registers.setSP(value)
}

func (c *CPU) instructionFA() {
	addr := c.read16BitOperand()
	value := c.memory.Read(addr)
	c.registers.setA(value)
}

func (c *CPU) instructionFB() {
	c.interrupts.EnableMaster()
	c.interrupts.EnableDelay()
}

func (c *CPU) instructionFE() {
	value1 := c.registers.getA()
	value2 := c.read8BitOperand()
	c.registers.f.setZ(value1 == value2)
	c.registers.f.setN(true)
	c.registers.f.setH((value1 & 0xF) < (value2 & 0xF))
	c.registers.f.setC(value1 < value2)
}

func (c *CPU) instructionFF() {
	value := uint16(0x38)
	addr := c.registers.getSP()
	msb := uint8(c.registers.getPC() >> 8)
	lsb := uint8(c.registers.getPC() & 0xFF)
	c.memory.Write(addr-1, msb)
	c.memory.Write(addr-2, lsb)
	c.registers.setSP(addr - 2)
	c.registers.setPC(value)
}

func (c *CPU) instructionCB00() {
	value := c.registers.getB()
	bit7 := value >> 7
	value = (value << 1) + bit7
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit7 == 1)
	c.registers.setB(value)
}

func (c *CPU) instructionCB01() {
	value := c.registers.getC()
	bit7 := value >> 7
	value = (value << 1) + bit7
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit7 == 1)
	c.registers.setC(value)
}

func (c *CPU) instructionCB02() {
	value := c.registers.getD()
	bit7 := value >> 7
	value = (value << 1) + bit7
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit7 == 1)
	c.registers.setD(value)
}

func (c *CPU) instructionCB03() {
	value := c.registers.getE()
	bit7 := value >> 7
	value = (value << 1) + bit7
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit7 == 1)
	c.registers.setE(value)
}

func (c *CPU) instructionCB04() {
	value := c.registers.getH()
	bit7 := value >> 7
	value = (value << 1) + bit7
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit7 == 1)
	c.registers.setH(value)
}

func (c *CPU) instructionCB05() {
	value := c.registers.getL()
	bit7 := value >> 7
	value = (value << 1) + bit7
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit7 == 1)
	c.registers.setL(value)
}

func (c *CPU) instructionCB06() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr)
	bit7 := value >> 7
	value = (value << 1) + bit7
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit7 == 1)
	c.memory.Write(addr, value)
}

func (c *CPU) instructionCB07() {
	value := c.registers.getA()
	bit7 := value >> 7
	value = (value << 1) + bit7
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit7 == 1)
	c.registers.setA(value)
}

func (c *CPU) instructionCB08() {
	value := c.registers.getB()
	bit0 := value << 7
	value = (value >> 1) + bit0
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit0 == 128)
	c.registers.setB(value)
}

func (c *CPU) instructionCB09() {
	value := c.registers.getC()
	bit0 := value << 7
	value = (value >> 1) + bit0
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit0 == 128)
	c.registers.setC(value)
}

func (c *CPU) instructionCB0A() {
	value := c.registers.getD()
	bit0 := value << 7
	value = (value >> 1) + bit0
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit0 == 128)
	c.registers.setD(value)
}

func (c *CPU) instructionCB0B() {
	value := c.registers.getE()
	bit0 := value << 7
	value = (value >> 1) + bit0
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit0 == 128)
	c.registers.setE(value)
}

func (c *CPU) instructionCB0C() {
	value := c.registers.getH()
	bit0 := value << 7
	value = (value >> 1) + bit0
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit0 == 128)
	c.registers.setH(value)
}

func (c *CPU) instructionCB0D() {
	value := c.registers.getL()
	bit0 := value << 7
	value = (value >> 1) + bit0
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit0 == 128)
	c.registers.setL(value)
}

func (c *CPU) instructionCB0E() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr)
	bit0 := value << 7
	value = (value >> 1) + bit0
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit0 == 128)
	c.memory.Write(addr, value)
}

func (c *CPU) instructionCB0F() {
	value := c.registers.getA()
	bit0 := value << 7
	value = (value >> 1) + bit0
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit0 == 128)
	c.registers.setA(value)
}

func (c *CPU) instructionCB10() {
	carry := c.registers.f.getCarry()
	value := c.registers.getB()
	bit7 := value >> 7
	value = (value << 1) + carry
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit7 == 1)
	c.registers.setB(value)
}

func (c *CPU) instructionCB11() {
	carry := c.registers.f.getCarry()
	value := c.registers.getC()
	bit7 := value >> 7
	value = (value << 1) + carry
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit7 == 1)
	c.registers.setC(value)
}

func (c *CPU) instructionCB12() {
	carry := c.registers.f.getCarry()
	value := c.registers.getD()
	bit7 := value >> 7
	value = (value << 1) + carry
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit7 == 1)
	c.registers.setD(value)
}

func (c *CPU) instructionCB13() {
	carry := c.registers.f.getCarry()
	value := c.registers.getE()
	bit7 := value >> 7
	value = (value << 1) + carry
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit7 == 1)
	c.registers.setE(value)
}

func (c *CPU) instructionCB14() {
	carry := c.registers.f.getCarry()
	value := c.registers.getH()
	bit7 := value >> 7
	value = (value << 1) + carry
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit7 == 1)
	c.registers.setH(value)
}

func (c *CPU) instructionCB15() {
	carry := c.registers.f.getCarry()
	value := c.registers.getL()
	bit7 := value >> 7
	value = (value << 1) + carry
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit7 == 1)
	c.registers.setL(value)
}

func (c *CPU) instructionCB16() {
	addr := c.registers.getHL()
	carry := c.registers.f.getCarry()
	value := c.memory.Read(addr)
	bit7 := value >> 7
	value = (value << 1) + carry
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit7 == 1)
	c.memory.Write(addr, value)
}

func (c *CPU) instructionCB17() {
	carry := c.registers.f.getCarry()
	value := c.registers.getA()
	bit7 := value >> 7
	value = (value << 1) + carry
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit7 == 1)
	c.registers.setA(value)
}

func (c *CPU) instructionCB18() {
	carry := c.registers.f.getCarry() << 7
	value := c.registers.getB()
	bit0 := value << 7
	value = (value >> 1) + carry
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit0 == 128)
	c.registers.setB(value)
}

func (c *CPU) instructionCB19() {
	carry := c.registers.f.getCarry() << 7
	value := c.registers.getC()
	bit0 := value << 7
	value = (value >> 1) + carry
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit0 == 128)
	c.registers.setC(value)
}

func (c *CPU) instructionCB1A() {
	carry := c.registers.f.getCarry() << 7
	value := c.registers.getD()
	bit0 := value << 7
	value = (value >> 1) + carry
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit0 == 128)
	c.registers.setD(value)
}

func (c *CPU) instructionCB1B() {
	carry := c.registers.f.getCarry() << 7
	value := c.registers.getE()
	bit0 := value << 7
	value = (value >> 1) + carry
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit0 == 128)
	c.registers.setE(value)
}

func (c *CPU) instructionCB1C() {
	carry := c.registers.f.getCarry() << 7
	value := c.registers.getH()
	bit0 := value << 7
	value = (value >> 1) + carry
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit0 == 128)
	c.registers.setH(value)
}

func (c *CPU) instructionCB1D() {
	carry := c.registers.f.getCarry() << 7
	value := c.registers.getL()
	bit0 := value << 7
	value = (value >> 1) + carry
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit0 == 128)
	c.registers.setL(value)
}

func (c *CPU) instructionCB1E() {
	addr := c.registers.getHL()
	carry := c.registers.f.getCarry() << 7
	value := c.memory.Read(addr)
	bit7 := value << 7
	value = (value >> 1) + carry
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit7 == 128)
	c.memory.Write(addr, value)
}

func (c *CPU) instructionCB1F() {
	carry := c.registers.f.getCarry() << 7
	value := c.registers.getA()
	bit0 := value << 7
	value = (value >> 1) + carry
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit0 == 128)
	c.registers.setA(value)
}

func (c *CPU) instructionCB20() {
	value := c.registers.getB()
	bit7 := value >> 7
	value = value << 1
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit7 == 1)
	c.registers.setB(value)
}

func (c *CPU) instructionCB21() {
	value := c.registers.getC()
	bit7 := value >> 7
	value = value << 1
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit7 == 1)
	c.registers.setC(value)
}

func (c *CPU) instructionCB22() {
	value := c.registers.getD()
	bit7 := value >> 7
	value = value << 1
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit7 == 1)
	c.registers.setD(value)
}

func (c *CPU) instructionCB23() {
	value := c.registers.getE()
	bit7 := value >> 7
	value = value << 1
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit7 == 1)
	c.registers.setE(value)
}

func (c *CPU) instructionCB24() {
	value := c.registers.getH()
	bit7 := value >> 7
	value = value << 1
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit7 == 1)
	c.registers.setH(value)
}

func (c *CPU) instructionCB25() {
	value := c.registers.getL()
	bit7 := value >> 7
	value = value << 1
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit7 == 1)
	c.registers.setL(value)
}

func (c *CPU) instructionCB26() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr)
	bit7 := value >> 7
	value = value << 1
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit7 == 1)
	c.memory.Write(addr, value)
}

func (c *CPU) instructionCB27() {
	value := c.registers.getA()
	bit7 := value >> 7
	value = value << 1
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit7 == 1)
	c.registers.setA(value)
}

func (c *CPU) instructionCB28() {
	value := c.registers.getB()
	bit0 := value << 7
	bit7 := value & (1 << 7)
	value = (value >> 1) + bit7
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit0 == 128)
	c.registers.setB(value)
}

func (c *CPU) instructionCB29() {
	value := c.registers.getC()
	bit0 := value << 7
	bit7 := value & (1 << 7)
	value = (value >> 1) + bit7
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit0 == 128)
	c.registers.setC(value)
}

func (c *CPU) instructionCB2A() {
	value := c.registers.getD()
	bit0 := value << 7
	bit7 := value & (1 << 7)
	value = (value >> 1) + bit7
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit0 == 128)
	c.registers.setD(value)
}

func (c *CPU) instructionCB2B() {
	value := c.registers.getE()
	bit0 := value << 7
	bit7 := value & (1 << 7)
	value = (value >> 1) + bit7
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit0 == 128)
	c.registers.setE(value)
}

func (c *CPU) instructionCB2C() {
	value := c.registers.getH()
	bit0 := value << 7
	bit7 := value & (1 << 7)
	value = (value >> 1) + bit7
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit0 == 128)
	c.registers.setH(value)
}

func (c *CPU) instructionCB2D() {
	value := c.registers.getL()
	bit0 := value << 7
	bit7 := value & (1 << 7)
	value = (value >> 1) + bit7
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit0 == 128)
	c.registers.setL(value)
}

func (c *CPU) instructionCB2E() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr)
	bit0 := value << 7
	bit7 := value & (1 << 7)
	value = (value >> 1) + bit7
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit0 == 128)
	c.memory.Write(addr, value)
}

func (c *CPU) instructionCB2F() {
	value := c.registers.getA()
	bit0 := value << 7
	bit7 := value & (1 << 7)
	value = (value >> 1) + bit7
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit0 == 128)
	c.registers.setA(value)
}

func (c *CPU) instructionCB30() {
	value := c.registers.getB()
	lo := value & 0xF
	hi := value >> 4
	value = lo<<4 | hi
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(false)
	c.registers.setB(value)
}

func (c *CPU) instructionCB31() {
	value := c.registers.getC()
	lo := value & 0xF
	hi := value >> 4
	value = lo<<4 | hi
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(false)
	c.registers.setC(value)
}

func (c *CPU) instructionCB32() {
	value := c.registers.getD()
	lo := value & 0xF
	hi := value >> 4
	value = lo<<4 | hi
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(false)
	c.registers.setD(value)
}

func (c *CPU) instructionCB33() {
	value := c.registers.getE()
	lo := value & 0xF
	hi := value >> 4
	value = lo<<4 | hi
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(false)
	c.registers.setE(value)
}

func (c *CPU) instructionCB34() {
	value := c.registers.getH()
	lo := value & 0xF
	hi := value >> 4
	value = lo<<4 | hi
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(false)
	c.registers.setH(value)
}

func (c *CPU) instructionCB35() {
	value := c.registers.getL()
	lo := value & 0xF
	hi := value >> 4
	value = lo<<4 | hi
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(false)
	c.registers.setL(value)
}

func (c *CPU) instructionCB36() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr)
	lo := value & 0xF
	hi := value >> 4
	value = lo<<4 | hi
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(false)
	c.memory.Write(addr, value)
}

func (c *CPU) instructionCB37() {
	value := c.registers.getA()
	lo := value & 0xF
	hi := value >> 4
	value = lo<<4 | hi
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(false)
	c.registers.setA(value)
}

func (c *CPU) instructionCB38() {
	value := c.registers.getB()
	bit0 := value << 7
	value = value >> 1
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit0 == 128)
	c.registers.setB(value)
}

func (c *CPU) instructionCB39() {
	value := c.registers.getC()
	bit0 := value << 7
	value = value >> 1
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit0 == 128)
	c.registers.setC(value)
}

func (c *CPU) instructionCB3A() {
	value := c.registers.getD()
	bit0 := value << 7
	value = value >> 1
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit0 == 128)
	c.registers.setD(value)
}

func (c *CPU) instructionCB3B() {
	value := c.registers.getE()
	bit0 := value << 7
	value = value >> 1
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit0 == 128)
	c.registers.setE(value)
}

func (c *CPU) instructionCB3C() {
	value := c.registers.getH()
	bit0 := value << 7
	value = value >> 1
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit0 == 128)
	c.registers.setH(value)
}

func (c *CPU) instructionCB3D() {
	value := c.registers.getL()
	bit0 := value << 7
	value = value >> 1
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit0 == 128)
	c.registers.setL(value)
}

func (c *CPU) instructionCB3E() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr)
	bit0 := value << 7
	value = value >> 1
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit0 == 128)
	c.memory.Write(addr, value)
}

func (c *CPU) instructionCB3F() {
	value := c.registers.getA()
	bit0 := value << 7
	value = value >> 1
	c.registers.f.setZ(value == 0)
	c.registers.f.setN(false)
	c.registers.f.setH(false)
	c.registers.f.setC(bit0 == 128)
	c.registers.setA(value)
}

func (c *CPU) instructionCB40() {
	value := c.registers.getB()
	bit := (value & (1 << 0)) == (1 << 0)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB41() {
	value := c.registers.getC()
	bit := (value & (1 << 0)) == (1 << 0)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB42() {
	value := c.registers.getD()
	bit := (value & (1 << 0)) == (1 << 0)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB43() {
	value := c.registers.getE()
	bit := (value & (1 << 0)) == (1 << 0)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB44() {
	value := c.registers.getH()
	bit := (value & (1 << 0)) == (1 << 0)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB45() {
	value := c.registers.getL()
	bit := (value & (1 << 0)) == (1 << 0)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB46() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr)
	bit := (value & (1 << 0)) == (1 << 0)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB47() {
	value := c.registers.getA()
	bit := (value & (1 << 0)) == (1 << 0)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB48() {
	value := c.registers.getB()
	bit := (value & (1 << 1)) == (1 << 1)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB49() {
	value := c.registers.getC()
	bit := (value & (1 << 1)) == (1 << 1)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB4A() {
	value := c.registers.getD()
	bit := (value & (1 << 1)) == (1 << 1)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB4B() {
	value := c.registers.getE()
	bit := (value & (1 << 1)) == (1 << 1)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB4C() {
	value := c.registers.getH()
	bit := (value & (1 << 1)) == (1 << 1)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB4D() {
	value := c.registers.getL()
	bit := (value & (1 << 1)) == (1 << 1)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB4E() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr)
	bit := (value & (1 << 1)) == (1 << 1)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB4F() {
	value := c.registers.getA()
	bit := (value & (1 << 1)) == (1 << 1)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB50() {
	value := c.registers.getB()
	bit := (value & (1 << 2)) == (1 << 2)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB51() {
	value := c.registers.getC()
	bit := (value & (1 << 2)) == (1 << 2)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB52() {
	value := c.registers.getD()
	bit := (value & (1 << 2)) == (1 << 2)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB53() {
	value := c.registers.getE()
	bit := (value & (1 << 2)) == (1 << 2)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB54() {
	value := c.registers.getH()
	bit := (value & (1 << 2)) == (1 << 2)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB55() {
	value := c.registers.getL()
	bit := (value & (1 << 2)) == (1 << 2)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB56() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr)
	bit := (value & (1 << 2)) == (1 << 2)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB57() {
	value := c.registers.getA()
	bit := (value & (1 << 2)) == (1 << 2)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB58() {
	value := c.registers.getB()
	bit := (value & (1 << 3)) == (1 << 3)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB59() {
	value := c.registers.getC()
	bit := (value & (1 << 3)) == (1 << 3)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB5A() {
	value := c.registers.getD()
	bit := (value & (1 << 3)) == (1 << 3)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB5B() {
	value := c.registers.getE()
	bit := (value & (1 << 3)) == (1 << 3)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB5C() {
	value := c.registers.getH()
	bit := (value & (1 << 3)) == (1 << 3)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB5D() {
	value := c.registers.getL()
	bit := (value & (1 << 3)) == (1 << 3)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB5E() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr)
	bit := (value & (1 << 3)) == (1 << 3)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB5F() {
	value := c.registers.getA()
	bit := (value & (1 << 3)) == (1 << 3)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB60() {
	value := c.registers.getB()
	bit := (value & (1 << 4)) == (1 << 4)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB61() {
	value := c.registers.getC()
	bit := (value & (1 << 4)) == (1 << 4)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB62() {
	value := c.registers.getD()
	bit := (value & (1 << 4)) == (1 << 4)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB63() {
	value := c.registers.getE()
	bit := (value & (1 << 4)) == (1 << 4)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB64() {
	value := c.registers.getH()
	bit := (value & (1 << 4)) == (1 << 4)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB65() {
	value := c.registers.getL()
	bit := (value & (1 << 4)) == (1 << 4)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB66() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr)
	bit := (value & (1 << 4)) == (1 << 4)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB67() {
	value := c.registers.getA()
	bit := (value & (1 << 4)) == (1 << 4)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB68() {
	value := c.registers.getB()
	bit := (value & (1 << 5)) == (1 << 5)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB69() {
	value := c.registers.getC()
	bit := (value & (1 << 5)) == (1 << 5)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB6A() {
	value := c.registers.getD()
	bit := (value & (1 << 5)) == (1 << 5)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB6B() {
	value := c.registers.getE()
	bit := (value & (1 << 5)) == (1 << 5)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB6C() {
	value := c.registers.getH()
	bit := (value & (1 << 5)) == (1 << 5)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB6D() {
	value := c.registers.getL()
	bit := (value & (1 << 5)) == (1 << 5)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB6E() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr)
	bit := (value & (1 << 5)) == (1 << 5)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB6F() {
	value := c.registers.getA()
	bit := (value & (1 << 5)) == (1 << 5)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB70() {
	value := c.registers.getB()
	bit := (value & (1 << 6)) == (1 << 6)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB71() {
	value := c.registers.getC()
	bit := (value & (1 << 6)) == (1 << 6)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB72() {
	value := c.registers.getD()
	bit := (value & (1 << 6)) == (1 << 6)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB73() {
	value := c.registers.getE()
	bit := (value & (1 << 6)) == (1 << 6)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB74() {
	value := c.registers.getH()
	bit := (value & (1 << 6)) == (1 << 6)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB75() {
	value := c.registers.getL()
	bit := (value & (1 << 6)) == (1 << 6)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB76() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr)
	bit := (value & (1 << 6)) == (1 << 6)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB77() {
	value := c.registers.getA()
	bit := (value & (1 << 6)) == (1 << 6)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB78() {
	value := c.registers.getB()
	bit := (value & (1 << 7)) == (1 << 7)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB79() {
	value := c.registers.getC()
	bit := (value & (1 << 7)) == (1 << 7)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB7A() {
	value := c.registers.getD()
	bit := (value & (1 << 7)) == (1 << 7)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB7B() {
	value := c.registers.getE()
	bit := (value & (1 << 7)) == (1 << 7)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB7C() {
	value := c.registers.getH()
	bit := (value & (1 << 7)) == (1 << 7)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB7D() {
	value := c.registers.getL()
	bit := (value & (1 << 7)) == (1 << 7)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB7E() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr)
	bit := (value & (1 << 7)) == (1 << 7)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB7F() {
	value := c.registers.getA()
	bit := (value & (1 << 7)) == (1 << 7)
	c.registers.f.setZ(!bit)
	c.registers.f.setN(false)
	c.registers.f.setH(true)
}

func (c *CPU) instructionCB80() {
	value := c.registers.getB() & (0xFF - (1 << 0))
	c.registers.setB(value)
}

func (c *CPU) instructionCB81() {
	value := c.registers.getC() & (0xFF - (1 << 0))
	c.registers.setC(value)
}

func (c *CPU) instructionCB82() {
	value := c.registers.getD() & (0xFF - (1 << 0))
	c.registers.setD(value)
}

func (c *CPU) instructionCB83() {
	value := c.registers.getE() & (0xFF - (1 << 0))
	c.registers.setE(value)
}

func (c *CPU) instructionCB84() {
	value := c.registers.getH() & (0xFF - (1 << 0))
	c.registers.setH(value)
}

func (c *CPU) instructionCB85() {
	value := c.registers.getL() & (0xFF - (1 << 0))
	c.registers.setL(value)
}

func (c *CPU) instructionCB86() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr) & (0xFF - (1 << 0))
	c.memory.Write(addr, value)
}

func (c *CPU) instructionCB87() {
	value := c.registers.getA() & (0xFF - (1 << 0))
	c.registers.setA(value)
}

func (c *CPU) instructionCB88() {
	value := c.registers.getB() & (0xFF - (1 << 1))
	c.registers.setB(value)
}

func (c *CPU) instructionCB89() {
	value := c.registers.getC() & (0xFF - (1 << 1))
	c.registers.setC(value)
}

func (c *CPU) instructionCB8A() {
	value := c.registers.getD() & (0xFF - (1 << 1))
	c.registers.setD(value)
}

func (c *CPU) instructionCB8B() {
	value := c.registers.getE() & (0xFF - (1 << 1))
	c.registers.setE(value)
}

func (c *CPU) instructionCB8C() {
	value := c.registers.getH() & (0xFF - (1 << 1))
	c.registers.setH(value)
}

func (c *CPU) instructionCB8D() {
	value := c.registers.getL() & (0xFF - (1 << 1))
	c.registers.setL(value)
}

func (c *CPU) instructionCB8E() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr) & (0xFF - (1 << 1))
	c.memory.Write(addr, value)
}

func (c *CPU) instructionCB8F() {
	value := c.registers.getA() & (0xFF - (1 << 1))
	c.registers.setA(value)
}

func (c *CPU) instructionCB90() {
	value := c.registers.getB() & (0xFF - (1 << 2))
	c.registers.setB(value)
}

func (c *CPU) instructionCB91() {
	value := c.registers.getC() & (0xFF - (1 << 2))
	c.registers.setC(value)
}

func (c *CPU) instructionCB92() {
	value := c.registers.getD() & (0xFF - (1 << 2))
	c.registers.setD(value)
}

func (c *CPU) instructionCB93() {
	value := c.registers.getE() & (0xFF - (1 << 2))
	c.registers.setE(value)
}

func (c *CPU) instructionCB94() {
	value := c.registers.getH() & (0xFF - (1 << 2))
	c.registers.setH(value)
}

func (c *CPU) instructionCB95() {
	value := c.registers.getL() & (0xFF - (1 << 2))
	c.registers.setL(value)
}

func (c *CPU) instructionCB96() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr) & (0xFF - (1 << 2))
	c.memory.Write(addr, value)
}

func (c *CPU) instructionCB97() {
	value := c.registers.getA() & (0xFF - (1 << 2))
	c.registers.setA(value)
}

func (c *CPU) instructionCB98() {
	value := c.registers.getB() & (0xFF - (1 << 3))
	c.registers.setB(value)
}

func (c *CPU) instructionCB99() {
	value := c.registers.getC() & (0xFF - (1 << 3))
	c.registers.setC(value)
}

func (c *CPU) instructionCB9A() {
	value := c.registers.getD() & (0xFF - (1 << 3))
	c.registers.setD(value)
}

func (c *CPU) instructionCB9B() {
	value := c.registers.getE() & (0xFF - (1 << 3))
	c.registers.setE(value)
}

func (c *CPU) instructionCB9C() {
	value := c.registers.getH() & (0xFF - (1 << 3))
	c.registers.setH(value)
}

func (c *CPU) instructionCB9D() {
	value := c.registers.getL() & (0xFF - (1 << 3))
	c.registers.setL(value)
}

func (c *CPU) instructionCB9E() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr) & (0xFF - (1 << 3))
	c.memory.Write(addr, value)
}

func (c *CPU) instructionCB9F() {
	value := c.registers.getA() & (0xFF - (1 << 3))
	c.registers.setA(value)
}

func (c *CPU) instructionCBA0() {
	value := c.registers.getB() & (0xFF - (1 << 4))
	c.registers.setB(value)
}

func (c *CPU) instructionCBA1() {
	value := c.registers.getC() & (0xFF - (1 << 4))
	c.registers.setC(value)
}

func (c *CPU) instructionCBA2() {
	value := c.registers.getD() & (0xFF - (1 << 4))
	c.registers.setD(value)
}

func (c *CPU) instructionCBA3() {
	value := c.registers.getE() & (0xFF - (1 << 4))
	c.registers.setE(value)
}

func (c *CPU) instructionCBA4() {
	value := c.registers.getH() & (0xFF - (1 << 4))
	c.registers.setH(value)
}

func (c *CPU) instructionCBA5() {
	value := c.registers.getL() & (0xFF - (1 << 4))
	c.registers.setL(value)
}

func (c *CPU) instructionCBA6() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr) & (0xFF - (1 << 4))
	c.memory.Write(addr, value)
}

func (c *CPU) instructionCBA7() {
	value := c.registers.getA() & (0xFF - (1 << 4))
	c.registers.setA(value)
}

func (c *CPU) instructionCBA8() {
	value := c.registers.getB() & (0xFF - (1 << 5))
	c.registers.setB(value)
}

func (c *CPU) instructionCBA9() {
	value := c.registers.getC() & (0xFF - (1 << 5))
	c.registers.setC(value)
}

func (c *CPU) instructionCBAA() {
	value := c.registers.getD() & (0xFF - (1 << 5))
	c.registers.setD(value)
}

func (c *CPU) instructionCBAB() {
	value := c.registers.getE() & (0xFF - (1 << 5))
	c.registers.setE(value)
}

func (c *CPU) instructionCBAC() {
	value := c.registers.getH() & (0xFF - (1 << 5))
	c.registers.setH(value)
}

func (c *CPU) instructionCBAD() {
	value := c.registers.getL() & (0xFF - (1 << 5))
	c.registers.setL(value)
}

func (c *CPU) instructionCBAE() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr) & (0xFF - (1 << 5))
	c.memory.Write(addr, value)
}

func (c *CPU) instructionCBAF() {
	value := c.registers.getA() & (0xFF - (1 << 5))
	c.registers.setA(value)
}

func (c *CPU) instructionCBB0() {
	value := c.registers.getB() & (0xFF - (1 << 6))
	c.registers.setB(value)
}

func (c *CPU) instructionCBB1() {
	value := c.registers.getC() & (0xFF - (1 << 6))
	c.registers.setC(value)
}

func (c *CPU) instructionCBB2() {
	value := c.registers.getD() & (0xFF - (1 << 6))
	c.registers.setD(value)
}

func (c *CPU) instructionCBB3() {
	value := c.registers.getE() & (0xFF - (1 << 6))
	c.registers.setE(value)
}

func (c *CPU) instructionCBB4() {
	value := c.registers.getH() & (0xFF - (1 << 6))
	c.registers.setH(value)
}

func (c *CPU) instructionCBB5() {
	value := c.registers.getL() & (0xFF - (1 << 6))
	c.registers.setL(value)
}

func (c *CPU) instructionCBB6() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr) & (0xFF - (1 << 6))
	c.memory.Write(addr, value)
}

func (c *CPU) instructionCBB7() {
	value := c.registers.getA() & (0xFF - (1 << 6))
	c.registers.setA(value)
}

func (c *CPU) instructionCBB8() {
	value := c.registers.getB() & (0xFF - (1 << 7))
	c.registers.setB(value)
}

func (c *CPU) instructionCBB9() {
	value := c.registers.getC() & (0xFF - (1 << 7))
	c.registers.setC(value)
}

func (c *CPU) instructionCBBA() {
	value := c.registers.getD() & (0xFF - (1 << 7))
	c.registers.setD(value)
}

func (c *CPU) instructionCBBB() {
	value := c.registers.getE() & (0xFF - (1 << 7))
	c.registers.setE(value)
}

func (c *CPU) instructionCBBC() {
	value := c.registers.getH() & (0xFF - (1 << 7))
	c.registers.setH(value)
}

func (c *CPU) instructionCBBD() {
	value := c.registers.getL() & (0xFF - (1 << 7))
	c.registers.setL(value)
}

func (c *CPU) instructionCBBE() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr) & (0xFF - (1 << 7))
	c.memory.Write(addr, value)
}

func (c *CPU) instructionCBBF() {
	value := c.registers.getA() & (0xFF - (1 << 7))
	c.registers.setA(value)
}

func (c *CPU) instructionCBC0() {
	value := c.registers.getB() | (1 << 0)
	c.registers.setB(value)
}

func (c *CPU) instructionCBC1() {
	value := c.registers.getC() | (1 << 0)
	c.registers.setC(value)
}

func (c *CPU) instructionCBC2() {
	value := c.registers.getD() | (1 << 0)
	c.registers.setD(value)
}

func (c *CPU) instructionCBC3() {
	value := c.registers.getE() | (1 << 0)
	c.registers.setE(value)
}

func (c *CPU) instructionCBC4() {
	value := c.registers.getH() | (1 << 0)
	c.registers.setH(value)
}

func (c *CPU) instructionCBC5() {
	value := c.registers.getL() | (1 << 0)
	c.registers.setL(value)
}

func (c *CPU) instructionCBC6() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr) | (1 << 0)
	c.memory.Write(addr, value)
}

func (c *CPU) instructionCBC7() {
	value := c.registers.getA() | (1 << 0)
	c.registers.setA(value)
}

func (c *CPU) instructionCBC8() {
	value := c.registers.getB() | (1 << 1)
	c.registers.setB(value)
}

func (c *CPU) instructionCBC9() {
	value := c.registers.getC() | (1 << 1)
	c.registers.setC(value)
}

func (c *CPU) instructionCBCA() {
	value := c.registers.getD() | (1 << 1)
	c.registers.setD(value)
}

func (c *CPU) instructionCBCB() {
	value := c.registers.getE() | (1 << 1)
	c.registers.setE(value)
}

func (c *CPU) instructionCBCC() {
	value := c.registers.getH() | (1 << 1)
	c.registers.setH(value)
}

func (c *CPU) instructionCBCD() {
	value := c.registers.getL() | (1 << 1)
	c.registers.setL(value)
}

func (c *CPU) instructionCBCE() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr) | (1 << 1)
	c.memory.Write(addr, value)
}

func (c *CPU) instructionCBCF() {
	value := c.registers.getA() | (1 << 1)
	c.registers.setA(value)
}

func (c *CPU) instructionCBD0() {
	value := c.registers.getB() | (1 << 2)
	c.registers.setB(value)
}

func (c *CPU) instructionCBD1() {
	value := c.registers.getC() | (1 << 2)
	c.registers.setC(value)
}

func (c *CPU) instructionCBD2() {
	value := c.registers.getD() | (1 << 2)
	c.registers.setD(value)
}

func (c *CPU) instructionCBD3() {
	value := c.registers.getE() | (1 << 2)
	c.registers.setE(value)
}

func (c *CPU) instructionCBD4() {
	value := c.registers.getH() | (1 << 2)
	c.registers.setH(value)
}

func (c *CPU) instructionCBD5() {
	value := c.registers.getL() | (1 << 2)
	c.registers.setL(value)
}

func (c *CPU) instructionCBD6() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr) | (1 << 2)
	c.memory.Write(addr, value)
}

func (c *CPU) instructionCBD7() {
	value := c.registers.getA() | (1 << 2)
	c.registers.setA(value)
}

func (c *CPU) instructionCBD8() {
	value := c.registers.getB() | (1 << 3)
	c.registers.setB(value)
}

func (c *CPU) instructionCBD9() {
	value := c.registers.getC() | (1 << 3)
	c.registers.setC(value)
}

func (c *CPU) instructionCBDA() {
	value := c.registers.getD() | (1 << 3)
	c.registers.setD(value)
}

func (c *CPU) instructionCBDB() {
	value := c.registers.getE() | (1 << 3)
	c.registers.setE(value)
}

func (c *CPU) instructionCBDC() {
	value := c.registers.getH() | (1 << 3)
	c.registers.setH(value)
}

func (c *CPU) instructionCBDD() {
	value := c.registers.getL() | (1 << 3)
	c.registers.setL(value)
}

func (c *CPU) instructionCBDE() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr) | (1 << 3)
	c.memory.Write(addr, value)
}

func (c *CPU) instructionCBDF() {
	value := c.registers.getA() | (1 << 3)
	c.registers.setA(value)
}

func (c *CPU) instructionCBE0() {
	value := c.registers.getB() | (1 << 4)
	c.registers.setB(value)
}

func (c *CPU) instructionCBE1() {
	value := c.registers.getC() | (1 << 4)
	c.registers.setC(value)
}

func (c *CPU) instructionCBE2() {
	value := c.registers.getD() | (1 << 4)
	c.registers.setD(value)
}

func (c *CPU) instructionCBE3() {
	value := c.registers.getE() | (1 << 4)
	c.registers.setE(value)
}

func (c *CPU) instructionCBE4() {
	value := c.registers.getH() | (1 << 4)
	c.registers.setH(value)
}

func (c *CPU) instructionCBE5() {
	value := c.registers.getL() | (1 << 4)
	c.registers.setL(value)
}

func (c *CPU) instructionCBE6() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr) | (1 << 4)
	c.memory.Write(addr, value)
}

func (c *CPU) instructionCBE7() {
	value := c.registers.getA() | (1 << 4)
	c.registers.setA(value)
}

func (c *CPU) instructionCBE8() {
	value := c.registers.getB() | (1 << 5)
	c.registers.setB(value)
}

func (c *CPU) instructionCBE9() {
	value := c.registers.getC() | (1 << 5)
	c.registers.setC(value)
}

func (c *CPU) instructionCBEA() {
	value := c.registers.getD() | (1 << 5)
	c.registers.setD(value)
}

func (c *CPU) instructionCBEB() {
	value := c.registers.getE() | (1 << 5)
	c.registers.setE(value)
}

func (c *CPU) instructionCBEC() {
	value := c.registers.getH() | (1 << 5)
	c.registers.setH(value)
}

func (c *CPU) instructionCBED() {
	value := c.registers.getL() | (1 << 5)
	c.registers.setL(value)
}

func (c *CPU) instructionCBEE() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr) | (1 << 5)
	c.memory.Write(addr, value)
}

func (c *CPU) instructionCBEF() {
	value := c.registers.getA() | (1 << 5)
	c.registers.setA(value)
}

func (c *CPU) instructionCBF0() {
	value := c.registers.getB() | (1 << 6)
	c.registers.setB(value)
}

func (c *CPU) instructionCBF1() {
	value := c.registers.getC() | (1 << 6)
	c.registers.setC(value)
}

func (c *CPU) instructionCBF2() {
	value := c.registers.getD() | (1 << 6)
	c.registers.setD(value)
}

func (c *CPU) instructionCBF3() {
	value := c.registers.getE() | (1 << 6)
	c.registers.setE(value)
}

func (c *CPU) instructionCBF4() {
	value := c.registers.getH() | (1 << 6)
	c.registers.setH(value)
}

func (c *CPU) instructionCBF5() {
	value := c.registers.getL() | (1 << 6)
	c.registers.setL(value)
}

func (c *CPU) instructionCBF6() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr) | (1 << 6)
	c.memory.Write(addr, value)
}

func (c *CPU) instructionCBF7() {
	value := c.registers.getA() | (1 << 6)
	c.registers.setA(value)
}

func (c *CPU) instructionCBF8() {
	value := c.registers.getB() | (1 << 7)
	c.registers.setB(value)
}

func (c *CPU) instructionCBF9() {
	value := c.registers.getC() | (1 << 7)
	c.registers.setC(value)
}

func (c *CPU) instructionCBFA() {
	value := c.registers.getD() | (1 << 7)
	c.registers.setD(value)
}

func (c *CPU) instructionCBFB() {
	value := c.registers.getE() | (1 << 7)
	c.registers.setE(value)
}

func (c *CPU) instructionCBFC() {
	value := c.registers.getH() | (1 << 7)
	c.registers.setH(value)
}

func (c *CPU) instructionCBFD() {
	value := c.registers.getL() | (1 << 7)
	c.registers.setL(value)
}

func (c *CPU) instructionCBFE() {
	addr := c.registers.getHL()
	value := c.memory.Read(addr) | (1 << 7)
	c.memory.Write(addr, value)
}

func (c *CPU) instructionCBFF() {
	value := c.registers.getA() | (1 << 7)
	c.registers.setA(value)
}
