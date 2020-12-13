package main

import "testing"

func TestComputeEncryptionWeakness(t *testing.T) {
	tests := []struct {
		numbers         []int
		violatingNumber int
		want            int
	}{
		{
			numbers:         []int{35, 20, 15, 25, 47, 40, 62, 55, 65},
			violatingNumber: 127,
			want:            62,
		},
		{
			numbers:         []int{150, 182, 127, 219, 15, 25, 47, 40},
			violatingNumber: 127,
			want:            62,
		},
		{
			numbers:         []int{35, 20},
			violatingNumber: 55,
			want:            55,
		},
	}
	for _, tt := range tests {
		t.Run("TestComputeEncryptionWeakness", func(t *testing.T) {
			if got := ComputeEncryptionWeakness(tt.numbers, tt.violatingNumber); got != tt.want {
				t.Errorf("ComputeEncryptionWeakness() = %v, want %v", got, tt.want)
			}
		})
	}
}
