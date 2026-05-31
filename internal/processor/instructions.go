package processor

type InstructionType int

const (
	I_0NNN InstructionType = iota
	I_00E0
	I_00EE
	I_1NNN
	I_2NNN
	I_3XNN
	I_4XNN
	I_5XY0
	I_6XNN
	I_7XNN
	I_8XY0
	I_8XY1
	I_8XY2
	I_8XY3
	I_8XY4
	I_8XY5
	I_8XY6
	I_8XY7
	I_8XYE
	I_9XY0
	I_ANNN
	I_BNNN
	I_CXNN
	I_DXYN
	I_EX9E
	I_EXA1
	I_FX07
	I_FX0A
	I_FX15
	I_FX18
	I_FX1E
	I_FX29
	I_FX33
	I_FX55
	I_FX65
)

var proc8Array = []InstructionProc{
	0x0: LDRR,
	0x1: ORRR,
	0x2: ANDRR,
	0x3: XORRR,
	0x4: ADDRRC,
	0x5: SUBXYB,
	0x6: SHIFTR,
	0x7: SUBYXB,
	0xE: SHIFTL,
}

var procFArray = []InstructionProc{
	0x07: StoreDelay,
	0x0A: WAIT,
	0x15: SetDelay,
	0x18: SetSound,
	0x1E: ADDRI,
	0x29: LDRI,
	0x33: BCD,
	0x55: WriteMRX,
	0x65: WriteRXM,
}

var instructionPointers = []InstructionProc{
	0x0: Proc0Collector,
	0x1: JP,
	0x2: CALL,
	0x3: SER,
	0x4: SNER,
	0x5: SERR,
	0x6: LDR,
	0x7: ADDRN,
	0x8: Proc8Collector,
	0x9: SNERR,
	0xA: STOREMI,
	0xB: JUMP,
	0xC: RANDR,
	0xD: DRAW,
	0xE: ProcECollector,
	0xF: ProcFCollector,
}
