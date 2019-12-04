package main

import (
	"fmt"
)

func assert(got, expected bool) {
	if got != expected {
		panic (fmt.Sprintf("Got %t, expected %t!", got, expected))
	}
}

func mapContainsValue(m map[int]int, val int) bool {
	for _, freq := range m {
		if freq == val {
			return true
		}
	}
	return false
}

// It goes from right to left
func isValidPassword(attempt int) bool {
	// Rules:
	// - There are two adjacent digits that are the same;
	// - From left to right, the digits never decrease.

	foundSameDig := false
	prevDig := 10
	for attempt > 0 {
		currentDig := attempt % 10
		if currentDig == prevDig {
			foundSameDig = true
		} else if prevDig < currentDig {
			return false
		}
		prevDig = currentDig
		attempt = int(attempt / 10)
	}
	return foundSameDig
}

// It goes from right to left
func isValidPasswordExtended(attempt int) bool {
	// Rules:
	// - Same as in `isValidPassword`;
	// - There is a repetition of same digit with length 2.

	sameDigFreq := make(map[int] int)
	prevDig := 10
	for attempt > 0 {
		currentDig := attempt % 10
		if currentDig == prevDig {
			sameDigFreq[currentDig]++
		} else if prevDig > currentDig {
			sameDigFreq[currentDig] = 1
		} else {
			return false
		}
		prevDig = currentDig
		attempt = int(attempt / 10)
	}
	return mapContainsValue(sameDigFreq, 2)
}

func nPasswords(from, to int, followRules func(int) bool) int {
	nValidPasswords := 0
	for password := from; password <= to; password++ {
		if followRules(password) {
			nValidPasswords ++
		}
	}
	return nValidPasswords
}

func main() {
	// Test cases
	assert(isValidPasswordExtended(112233), true)
	assert(isValidPasswordExtended(123444), false)
	assert(isValidPasswordExtended(111122), true)

	// Problem
	fmt.Println(nPasswords(271973, 785961, isValidPassword)) // Pt1
	fmt.Println(nPasswords(271973, 785961, isValidPasswordExtended)) // Pt2
}
