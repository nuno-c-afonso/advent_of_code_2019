/*---------|
|- PART 1 -|
|---------*/
// Fuel required to launch a given module is based on its mass. Specifically, to
// find the fuel required for a module, take its mass, divide by three, round
// down, and subtract 2.
// For example:
// - For a mass of 12, divide by 3 and round down to get 4, then subtract 2 to
// 	 get 2.
// - For a mass of 14, dividing by 3 and rounding down still yields 4, so the
// 	 fuel required is also 2.
// - For a mass of 1969, the fuel required is 654.
// - For a mass of 100756, the fuel required is 33583.
// The Fuel Counter-Upper needs to know the total fuel requirement. To find it,
// individually calculate the fuel needed for the mass of each module (your
// puzzle input), then add together all the fuel values.
// What is the sum of the fuel requirements for all of the modules on your
// spacecraft?

/*---------|
|- PART 2 -|
|---------*/
// During the second Go / No Go poll, the Elf in charge of the Rocket Equation
// Double-Checker stops the launch sequence. Apparently, you forgot to include
// additional fuel for the fuel you just added.
// Fuel itself requires fuel just like a module - take its mass, divide by
// three, round down, and subtract 2. However, that fuel also requires fuel,
// and that fuel requires fuel, and so on. Any mass that would require negative
// fuel should instead be treated as if it requires zero fuel; the remaining
// mass, if any, is instead handled by wishing really hard, which has no mass
// and is outside the scope of this calculation.
// So, for each module mass, calculate its fuel and add it to the total. Then,
// treat the fuel amount you just calculated as the input mass and repeat the
// process, continuing until a fuel requirement is zero or negative.
// For example:
// - A module of mass 14 requires 2 fuel. This fuel requires no further fuel (2
// 	divided by 3 and rounded down is 0, which would call for a negative fuel),
// 	so the total fuel required is still just 2.
// - At first, a module of mass 1969 requires 654 fuel. Then, this fuel requires
// 	216 more fuel (654 / 3 - 2). 216 then requires 70 more fuel, which requires
// 	21 fuel, which requires 5 fuel, which requires no further fuel. So, the
// 	total fuel required for a module of mass 1969 is
// 	654 + 216 + 70 + 21 + 5 = 966.
// - The fuel required by a module of mass 100756 and its fuel is:
// 	33583 + 11192 + 3728 + 1240 + 411 + 135 + 43 + 12 + 2 = 50346.
// What is the sum of the fuel requirements for all of the modules on your
// spacecraft when also taking into account the mass of the added fuel?
// (Calculate the fuel requirements for each module separately, then add them
// all up at the end.)

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func assert(got int, expected int) {
	if got != expected {
		panic(fmt.Sprintf("ERROR: %d != %d", got, expected))
	}
}

func massToFuel(mass int) int {
	return int(mass / 3) - 2
}

func rollingFuel(fuel int) int {
	fuel = massToFuel(fuel)
	if fuel <= 0 {
		return 0
	}
	return fuel + rollingFuel(fuel)
}

func fuel(filename string, conversion func(mass int) int) int {
	finalFuel := 0
	f, _ := os.Open(filename)
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		mass, _ := strconv.Atoi(line)
		finalFuel += conversion(mass)
	}
	return finalFuel
}

func main() {
	// Test cases - pt1
	assert(massToFuel(12), 2)
	assert(massToFuel(14), 2)
	assert(massToFuel(1969), 654)
	assert(massToFuel(100756), 33583)

	// Test cases - pt2
	assert(rollingFuel(2), 0)
	assert(rollingFuel(14), 2)
	assert(rollingFuel(1969), 966)
	assert(rollingFuel(100756), 50346)

	// Get the amounts of fuel
	filename := "input.txt"
	initFuel := fuel(filename, massToFuel)
	fmt.Printf("Part 1: %d\n", initFuel)
	fmt.Printf("Part 2 (wrong): %d\n", initFuel + rollingFuel(initFuel))
	fmt.Printf("Part 2 (right): %d\n", fuel(filename, rollingFuel))
}
