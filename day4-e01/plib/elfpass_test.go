package plib

import (
	"testing"
)

func TestPasswordToDigitSequence(t *testing.T) {
	seq := PasswordToDigitSequence(947)
	if seq[0] != 9 || seq[1] != 4 || seq[2] != 7 {
		t.Fail()
	}
}

func TestHasTwoAdjacentDigits(t *testing.T) {
	if !HasTwoAdjacentDigits(PasswordToDigitSequence(1234456)) {
		t.Fail()
	}

	if HasTwoAdjacentDigits(PasswordToDigitSequence(12345)) {
		t.Fail()
	}
}

func TestDigitsOnlyIncrease(t *testing.T) {
	if !DigitsOnlyIncrease(PasswordToDigitSequence(123)) {
		t.Fail()
	}

	if DigitsOnlyIncrease(PasswordToDigitSequence(123434)) {
		t.Fail()
	}
}

func TestCheckPassword(t *testing.T) {
	if !CheckPassword(PasswordToDigitSequence(111111)) {
		t.Fail()
	}

	if CheckPassword(PasswordToDigitSequence(223450)) {
		t.Fail()
	}

	if CheckPassword(PasswordToDigitSequence(123789)) {
		t.Fail()
	}

}
