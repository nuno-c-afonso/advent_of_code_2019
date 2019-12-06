package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const Sep = "," // Char to split the program

// OPCODES
const OpAdd = 1
const OpMul = 2
const OpIn = 3
const OpOut = 4
const OpRet = 99

// MODES
const MPos = 0
const MImm = 1

func assert(got string, expected string) {
	if got != expected {
		panic(fmt.Sprintf("ERROR: %s != %s", got, expected))
	}
}

func assertInt(got int, expected int) {
	if got != expected {
		panic(fmt.Sprintf("ERROR: %d != %d", got, expected))
	}
}

// Operations

func add(l, r int) int {
	return l + r
}

func mul(l, r int) int {
	return l * r
}

func readStdin(mem []int, pos int, modes int) int {
	if modes == MImm {
		panic("ERROR: Cannot input value into immediate!")
	}
	operand := modeToOperand(modes, mem, pos + 1, 1)
	fmt.Print("Input number: ")
	fmt.Scanf("%d", operand[0])
	return 2
}

func writeStdout(mem []int, pos int, modes int) int {
	operand := modeToOperand(modes, mem, pos + 1, 1)
	fmt.Printf("Output: %d\n", *operand[0])
	return 2
}

// Helpers

func digAt(n int, pos int) int {
	ignore := int(math.Pow10(pos))
	return int(n / ignore) % 10
}

func modeToOperand(addressModes int, mem []int, pos int, nArgs int) []*int {
	var operands []*int
	for i := 0; i < nArgs; i++ {
		addressMode := digAt(addressModes, i)
		memContent := &mem[pos + i]
		if addressMode == MImm {
			operands = append(operands, memContent)
		} else if addressMode == MPos {
			operands = append(operands, &mem[*memContent])
		} else {
			panic(fmt.Sprintf("Unknown addressing mode: %d!", addressMode))
		}
	}
	return operands
}

func applyOp(mem []int, pos int, modes int, op func(int, int) int) int {
	addressModes := int(mem[pos] / 100)
	operands := modeToOperand(addressModes, mem, pos + 1, 3)
	*operands[2] = op(*operands[0], *operands[1]) // The ret pos is always MPos
	return 4
}

func strToIntSlice(str string, sep string) []int {
	var finalSlice []int
	slice := strings.Split(str, sep)
	sliceSize := len(slice)
	for i := 0; i < sliceSize; i++ {
		n, _ := strconv.Atoi(slice[i])
		finalSlice = append(finalSlice, n)
	}
	return finalSlice
}

func prepRunCode(code string) string {
	memSlice := strToIntSlice(code, Sep)
	runProgram(memSlice)
	return fmt.Sprint(memSlice)
}

// It runs the code and returns the changed memory
func runProgram(memSlice []int) {
	incr := 0
	for instPointer := 0; memSlice[instPointer] != 99; instPointer += incr {
		op := memSlice[instPointer]
		opCode := op % 100
		modes := int(op / 100)
		if opCode == OpAdd {
			incr = applyOp(memSlice, instPointer, modes, add)
		} else if opCode == OpMul {
			incr = applyOp(memSlice, instPointer, modes, mul)
		} else if opCode == OpIn {
			incr = readStdin(memSlice, instPointer, modes)
		}	else if opCode == OpOut {
			incr = writeStdout(memSlice, instPointer, modes)
		}	else {
			fmt.Println(memSlice)
			panic(fmt.Sprintf("Unexpected opcode %d!\n", op))
		}
	}
}

func main() {
	// Test cases - Day 02
	assert(
		prepRunCode("1,9,10,3,2,3,11,0,99,30,40,50"),
		"[3500 9 10 70 2 3 11 0 99 30 40 50]")
	assert(
		prepRunCode("1,0,0,0,99"),
		"[2 0 0 0 99]")
	assert(
		prepRunCode("2,3,0,3,99"),
		"[2 3 0 6 99]")
	assert(
		prepRunCode("2,4,4,5,99,0"),
		"[2 4 4 5 99 9801]")
	assert(
		prepRunCode("1,1,1,4,99,5,6,0,99"),
		"[30 1 1 4 2 5 6 0 99]")
	
		// Test cases - Day 05
	assertInt(digAt(123, 0), 3)
	assertInt(digAt(123, 1), 2)
	assertInt(digAt(123, 2), 1)

	// Programs
	prog := "3,225,1,225,6,6,1100,1,238,225,104,0,1102,46,47,225,2,122,130,224,101,-1998,224,224,4,224,1002,223,8,223,1001,224,6,224,1,224,223,223,1102,61,51,225,102,32,92,224,101,-800,224,224,4,224,1002,223,8,223,1001,224,1,224,1,223,224,223,1101,61,64,225,1001,118,25,224,101,-106,224,224,4,224,1002,223,8,223,101,1,224,224,1,224,223,223,1102,33,25,225,1102,73,67,224,101,-4891,224,224,4,224,1002,223,8,223,1001,224,4,224,1,224,223,223,1101,14,81,225,1102,17,74,225,1102,52,67,225,1101,94,27,225,101,71,39,224,101,-132,224,224,4,224,1002,223,8,223,101,5,224,224,1,224,223,223,1002,14,38,224,101,-1786,224,224,4,224,102,8,223,223,1001,224,2,224,1,223,224,223,1,65,126,224,1001,224,-128,224,4,224,1002,223,8,223,101,6,224,224,1,224,223,223,1101,81,40,224,1001,224,-121,224,4,224,102,8,223,223,101,4,224,224,1,223,224,223,4,223,99,0,0,0,677,0,0,0,0,0,0,0,0,0,0,0,1105,0,99999,1105,227,247,1105,1,99999,1005,227,99999,1005,0,256,1105,1,99999,1106,227,99999,1106,0,265,1105,1,99999,1006,0,99999,1006,227,274,1105,1,99999,1105,1,280,1105,1,99999,1,225,225,225,1101,294,0,0,105,1,0,1105,1,99999,1106,0,300,1105,1,99999,1,225,225,225,1101,314,0,0,106,0,0,1105,1,99999,1008,677,226,224,1002,223,2,223,1005,224,329,1001,223,1,223,107,677,677,224,102,2,223,223,1005,224,344,101,1,223,223,1107,677,677,224,102,2,223,223,1005,224,359,1001,223,1,223,1108,226,226,224,1002,223,2,223,1006,224,374,101,1,223,223,107,226,226,224,1002,223,2,223,1005,224,389,1001,223,1,223,108,226,226,224,1002,223,2,223,1005,224,404,1001,223,1,223,1008,677,677,224,1002,223,2,223,1006,224,419,1001,223,1,223,1107,677,226,224,102,2,223,223,1005,224,434,1001,223,1,223,108,226,677,224,102,2,223,223,1006,224,449,1001,223,1,223,8,677,226,224,102,2,223,223,1006,224,464,1001,223,1,223,1007,677,226,224,1002,223,2,223,1006,224,479,1001,223,1,223,1007,677,677,224,1002,223,2,223,1005,224,494,1001,223,1,223,1107,226,677,224,1002,223,2,223,1006,224,509,101,1,223,223,1108,226,677,224,102,2,223,223,1005,224,524,1001,223,1,223,7,226,226,224,102,2,223,223,1005,224,539,1001,223,1,223,8,677,677,224,1002,223,2,223,1005,224,554,101,1,223,223,107,677,226,224,102,2,223,223,1006,224,569,1001,223,1,223,7,226,677,224,1002,223,2,223,1005,224,584,1001,223,1,223,1008,226,226,224,1002,223,2,223,1006,224,599,101,1,223,223,1108,677,226,224,102,2,223,223,1006,224,614,101,1,223,223,7,677,226,224,102,2,223,223,1005,224,629,1001,223,1,223,8,226,677,224,1002,223,2,223,1006,224,644,101,1,223,223,1007,226,226,224,102,2,223,223,1005,224,659,101,1,223,223,108,677,677,224,1002,223,2,223,1006,224,674,1001,223,1,223,4,223,99,226"
	
	// Pt 1
	prepRunCode(prog)
}
