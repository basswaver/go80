package disassembler

import (
	"fmt"
	"os"
	"strings"
)

var instructions map[byte]string = map[byte]string{
    0x00: "NOP",		0x02: "STAX B",  	0x12: "STAX D",
    0x10: "NOP",		0x03: "INX B",   	0x06: "MVI B, %",
    0x20: "NOP",		0x13: "INX D",   	0x16: "MVI D, %",
    0x30: "NOP",		0x23: "INX H",   	0x26: "MVI H, %",
    0x08: "NOP",		0x33: "INX SP",  	0x36: "MVI M, %",
    0x18: "NOP",
    0x28: "NOP",		0x04: "INR B",   	0x05: "DCR B",
    0x38: "NOP",		0x14: "INR D",   	0x15: "DCR D",
    0x01: "LXI B, %%",	0x24: "INR H",   	0x25: "DCR H",
    0x11: "LXI D, %%",	0x34: "INR M",   	0x35: "DCR M",
    0x21: "LXI H, %%",	0x2a: "LHLD %%",
    0x31: "LXI SP, %%",	0x22: "SHLD $%%",
	0x07: "RLC",		0x17: "RAL",
	0x0f: "RCC",		0x1f: "RAR",		0x2f: "CMA",
	0x27: "DAA",		0x37: "STC",		0x3f: "CMC"

	0x3a: "LDA %%",		0x32: "STA $%%",
	0x0a: "LDAX B",		0x1a: "LDAX D",
	0x09: "DAD B",		0x0b: "DCX B",
	0x19: "DAD D",		0x1b: "DCX D",
	0x29: "DAD H",		0x2b: "DCX H",
	0x39: "DAD SP",		0x3b: "DCX SP",

	0x0c: "INR C",		0x0d: "DCR C",		0x0e: "MVI C, %%",
	0x1c: "INR E",		0x1d: "DCR E",		0x1e: "MVI E, %%",
	0x2c: "INR L",		0x2d: "DCR L",		0x2e: "MVI L, %%",
	0x3c: "INR A",		0x3d: "DCR A",		0x3e: "MVI A, %%",

	0x76: "HLT"

	0x40: "MOV B, B", 	0x41: "MOV B, C",	0x42: "MOV B, D",	// MOV B
	0x43: "MOV B, E", 	0x44: "MOV B, H", 	0x45: "MOV B, L",
	0x46: "MOV B, M",	0x47: "MOV B, A"

	0x48: "MOV C, B", 	0x49: "MOV C, C",	0x4a: "MOV C, D",	// MOV C
	0x4b: "MOV C, E", 	0x4c: "MOV C, H", 	0x4d: "MOV C, L",
	0x4e: "MOV C, M",	0x4f: "MOV C, A"

	0x50: "MOV D, B", 	0x51: "MOV D, C",	0x52: "MOV D, D",	// MOV D
	0x53: "MOV D, E", 	0x54: "MOV D, H", 	0x55: "MOV D, L",
	0x56: "MOV D, M",	0x57: "MOV D, A"

	0x58: "MOV E, B", 	0x59: "MOV E, C",	0x5a: "MOV E, D",	// MOV E
	0x5b: "MOV E, E", 	0x5c: "MOV E, H", 	0x5d: "MOV E, L",
	0x5e: "MOV E, M",	0x5f: "MOV E, A"

	0x60: "MOV H, B", 	0x61: "MOV H, C",	0x62: "MOV H, D",	// MOV H
	0x63: "MOV H, E", 	0x64: "MOV H, H", 	0x65: "MOV H, L",
	0x66: "MOV H, M",	0x67: "MOV H, A"

	0x68: "MOV L, B", 	0x69: "MOV L, C",	0x6a: "MOV L, D",	// MOV L
	0x6b: "MOV L, E", 	0x6c: "MOV L, H", 	0x6d: "MOV L, L",
	0x6e: "MOV L, M",	0x6f: "MOV L, A"

	0x70: "MOV M, B", 	0x71: "MOV M, C",	0x72: "MOV M, D",	// MOV M
	0x73: "MOV M, E", 	0x74: "MOV M, H", 	0x75: "MOV M, L",
						0x77: "MOV M, A"

	0x78: "MOV A, B", 	0x79: "MOV A, C",	0x7a: "MOV A, D",	// MOV A
	0x7b: "MOV A, E", 	0x7c: "MOV A, H", 	0x7d: "MOV A, L",
	0x7e: "MOV A, M",	0x7f: "MOV A, A"
}

func bytes_of(path string) ([]byte, int64, error) {
	var stat os.FileInfo
	var err error
	stat, err = os.Stat(path)

	if err != nil {
		return make([]byte, 0), 0, err
	}

	var size int64
	size = stat.Size()

	var bytes []byte = make([]byte, size)

	var file *os.File
	file, err = os.Open(path)

	if err != nil {
		return bytes, size, err
	}

	defer file.Close()
	file.Read(bytes)

	return bytes, size, nil
}

func disassemble_bytes(bytes []byte, size int64) ([]string, error) {
	var index int64 = 0
	var argc int64
	var instruction string
	var args []byte

	for index < size {
		instruction = instructions[bytes[index]]
		argc = int64(strings.Count(instruction, "%"))
		args = bytes[index+1 : index+1+argc]
		instruction = strings.ReplaceAll(instruction, "%", "")
		index += argc + 1

		for argc > 0 {
			argc--
			instruction += fmt.Sprintf("%02X", args[argc])
		}

		fmt.Println(instruction)
	}

	return []string{}, nil
}

func push(array *[]string, item string) {
	var new []string = append(*array, item)
	array = &new
}

func T() {
	var bytes []byte
	var size int64
	bytes, size, _ = bytes_of("./source/invaders.h")
	disassemble_bytes(bytes, size)
}
