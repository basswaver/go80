package disassembler

import (
    "os"
    "fmt"
    "strings"
)

var instructions map[byte]string = map[byte]string{
    0x00: "NOP",                0x02: "STAX B",     0x12: "STAX D",
    0x10: "NOP",                0x03: "INX B",      0x06: "MVI B, %",
    0x20: "NOP",                0x13: "INX D",      0x16: "MVI D, %",
    0x30: "NOP",                0x23: "INX H",      0x26: "MVI H, %",
    0x08: "NOP",                0x33: "INX SP",     0x36: "MVI M, %",
    0x18: "NOP",
    0x28: "NOP",                0x04: "INR B",      0x05: "DCR B",
    0x38: "NOP",                0x14: "INR D",      0x15: "DCR D",
    0x01: "LXI B, %%",       0x24: "INR H",      0x25: "DCR H",
    0x11: "LXI D, %%",       0x34: "INR M",      0x35: "DCR M",
    0x21: "LXI H, %%",
    0x31: "LXI SP, %%",      0x22: "SHLD %%", 0x32: "STA %%",
}

func bytes_of(path string) ([]byte, int64, error) {
    var stat os.FileInfo
    var err error
    stat, err = os.Stat(path)

    if (err != nil) {
        return make([]byte, 0), 0, err
    }

    var size int64
    size = stat.Size()

    var bytes []byte = make([]byte, size)

    var file *os.File
    file, err = os.Open(path)

    if (err != nil) {
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

    for (index < size) {
        instruction = instructions[bytes[index]]
        argc = int64(strings.Count(instruction, "%"))
        args = bytes[index + 1: index + 1 + argc]
        instruction = strings.ReplaceAll(instruction, "%", "")
        index += argc + 1

        for (argc > 0) {
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
