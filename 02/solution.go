package main

import (
	"fmt"
	"strconv"
	"strings"
)

const Incr = 4 // Step to get the following instruction
const Sep = "," // Char to split the program

// OPCODES
const OpAdd = 1
const OpMul = 2
const OpRet = 99

func assert(got string, expected string) {
	if got != expected {
		panic(fmt.Sprintf("ERROR: %s != %s", got, expected))
	}
}

func add(l, r int) int {
	return l + r
}

func applyOp(mem []int, pos int, op func(int, int) int) {
	lPos, rPos, retPos := mem[pos + 1], mem[pos + 2], mem[pos + 3]
	mem[retPos] = op(mem[lPos], mem[rPos])
}

func mul(l, r int) int {
	return l * r
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

// TODO: Check for panic when there are still iterations to do!
func findInputPair(code string, target int) string {
	baseMemSlice := strToIntSlice(code, Sep)
	for l := 0; l < 100; l++ {
		for r := 0; r < 100; r++ {
			memToUse := append([]int(nil), baseMemSlice...)
			memToUse[1], memToUse[2] = l, r
			runProgram(memToUse)
			if memToUse[0] == target {
				return fmt.Sprintf("noun: %d\nverb: %d", memToUse[1], memToUse[2])
			}
		}
	}
	panic("There is no combination that matches the target!")
}

// It runs the code and returns the changed memory
func runProgram(memSlice []int) {
	for instPointer := 0; memSlice[instPointer] != 99; instPointer += Incr {
		op := memSlice[instPointer]
		if op == OpAdd {
			applyOp(memSlice, instPointer, add)
		} else if op == OpMul {
			applyOp(memSlice, instPointer, mul)
		} else {
			fmt.Println(memSlice)
			panic(fmt.Sprintf("Unexpected opcode %d!\n", op))
		}
	}
}

func main() {
	// Test cases
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

	// Programs
	correctProg := "1,12,2,3,1,1,2,3,1,3,4,3,1,5,0,3,2,13,1,19,1,19,6,23,1,23,6,27,1,13,27,31,2,13,31,35,1,5,35,39,2,39,13,43,1,10,43,47,2,13,47,51,1,6,51,55,2,55,13,59,1,59,10,63,1,63,10,67,2,10,67,71,1,6,71,75,1,10,75,79,1,79,9,83,2,83,6,87,2,87,9,91,1,5,91,95,1,6,95,99,1,99,9,103,2,10,103,107,1,107,6,111,2,9,111,115,1,5,115,119,1,10,119,123,1,2,123,127,1,127,6,0,99,2,14,0,0"
	
	// Pt 1
	fmt.Println(prepRunCode(correctProg))
	// Pt 2
	fmt.Println(findInputPair(correctProg, 19690720))
}
