package main

import "testing"

func TestComputeJoltProduct(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: `16
10
15
5
1
11
7
19
6
12
4`,
			want: 35,
		},
	}
	for _, tt := range tests {
		t.Run("TestComputeJoltProduct", func(t *testing.T) {
			if got := ComputeJoltProduct(tt.input); got != tt.want {
				t.Errorf("ComputeJoltProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
