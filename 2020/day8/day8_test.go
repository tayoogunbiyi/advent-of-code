package main

import "testing"

func TestFindFinalAccumulatorValueWithCycle(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{

		{input: `nop +0
acc +1 
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6

`, want: 8,
		},
	}
	for _, tt := range tests {
		t.Run("TestFindFinalAccumulatorValueWithCycle", func(t *testing.T) {
			if got := FindFinalAccumulatorValueWithCycle(tt.input); got != tt.want {
				t.Errorf("FindFinalAccumulatorValueWithCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
