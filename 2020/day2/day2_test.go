package main

import (
	"testing"
)

func TestCountValidPasswords(t *testing.T) {
	passwordPolicies := readInputIntoPolicySlice("1-3 a: abcde\n1-3 b: cdefg\n2-9 c: ccccccccc")
	validPasswordCount := CountValidPasswords(passwordPolicies)
	if validPasswordCount != 2 {
		t.Errorf("expected CountValidPasswords(%v) = %d; got %d", passwordPolicies, 2, validPasswordCount)
	}
}

func TestCountValidPasswordsGivenNewInterpretation(t *testing.T) {
	passwordPolicies := readInputIntoPolicySlice("1-3 a: abcde\n1-3 b: cdefg\n2-9 c: ccccccccc")
	validPasswordCount := CountValidPasswordsGivenNewInterpretation(passwordPolicies)
	if validPasswordCount != 1 {
		t.Errorf("expected CountValidPasswords(%v) = %d; got %d", passwordPolicies, 1, validPasswordCount)
	}
}
