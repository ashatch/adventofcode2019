package plib

import (
	"fmt"
	"strconv"
)

func PasswordToDigitSequence(pass int) []int {
	str := fmt.Sprintf("%d", pass)
	sequence := []int{}

	for _, c := range str {
		cstr, _ := strconv.Atoi(string(c))
		sequence = append(sequence, cstr)
	}

	return sequence
}

func HasTwoAdjacentDigits(sequence []int) bool {
	sequenceLength := len(sequence)
	if sequenceLength < 2 {
		return false
	}

	for i := 0; i < sequenceLength-1; i++ {
		if sequence[i] == sequence[i+1] {
			return true
		}
	}
	return false
}

func HasTwoAdjacentDigitsInLargerGroup(sequence []int) bool {
	sequenceLength := len(sequence)

	adjFreqCount := [10]int{}

	for n := 0; n < 10; n++ {
		count := 0
		latch := true
		for i := 0; i < sequenceLength; i++ {
			if sequence[i] == n {
				if latch {
					count++
				}
			} else {
				if count > 0 {
					latch = false
				}
			}
		}

		adjFreqCount[n] = count
	}

	var hasTwo bool
	var hasLarger bool

	for _, i := range adjFreqCount {
		if i == 2 {
			hasTwo = true
		}
		if i > 2 {
			hasLarger = true
		}
	}

	return hasLarger && !hasTwo
}

func DigitsOnlyIncrease(sequence []int) bool {
	sequenceLength := len(sequence)
	if sequenceLength < 2 {
		return true
	}

	for i := 0; i < sequenceLength-1; i++ {
		if sequence[i] > sequence[i+1] {
			return false
		}
	}
	return true
}

func CheckPassword(sequence []int) bool {
	return len(sequence) == 6 && HasTwoAdjacentDigits(sequence) && DigitsOnlyIncrease(sequence) && (!HasTwoAdjacentDigitsInLargerGroup(sequence))
}

func CheckPasswordRange(start int, end int) int {
	count := 0
	for i := start; i <= end; i++ {
		if CheckPassword(PasswordToDigitSequence(i)) {
			count++
		}
	}
	return count
}
