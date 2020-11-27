package main

import (
	"testing"
)

type Fixture struct {
	RocketMasses []int
	Expected     int
}

func TestGetTotalFuelRequirements(t *testing.T) {
	fixtures := []Fixture{
		{[]int{12}, 2},
		{[]int{14, 0, 0}, 2},
		{[]int{1969}, 654},
	}

	for _, fixture := range fixtures {
		got := GetTotalFuelRequirements(fixture.RocketMasses)
		if got != fixture.Expected {
			t.Errorf("GetTotalFuelRequirements(%d) = %d; want %d", fixture.RocketMasses, got, fixture.Expected)
		}
	}
}
