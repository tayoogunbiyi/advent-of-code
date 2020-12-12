// https://adventofcode.com/2020/day/7
package main

import "testing"

func TestCountBagsWithin(t *testing.T) {
	tests := []struct {
		input          string
		targetBagColor string
		want           int
	}{
		{
			input: `shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.
			`,
			targetBagColor: "shiny gold",
			want:           126,
		},
		{
			input: `shiny gold bags contain 2 dark red bags.
dark red bags contain 1 purple black bag.
purple black bags contains no other bags.
			`,
			targetBagColor: "shiny gold",
			want:           4,
		},
	}
	for _, tt := range tests {
		t.Run("TestCountBagsWithin", func(t *testing.T) {
			if got := CountBagsWithin(tt.input, tt.targetBagColor); got != tt.want {
				t.Errorf("CountBagsWithin() = %v, want %v", got, tt.want)
			}
		})
	}
}
