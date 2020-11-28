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

func TestGetTotalFuelRequirementsPartB(t *testing.T) {
	fixtures := []Fixture{
		{[]int{14}, 2},
		{[]int{1969}, 966},
		{[]int{100756}, 50346},
	}

	for _, fixture := range fixtures {
		got := GetTotalFuelRequirementsPartB(fixture.RocketMasses)
		if got != fixture.Expected {
			t.Errorf("GetTotalFuelRequirementsPartB(%d) = %d; want %d", fixture.RocketMasses, got, fixture.Expected)
		}
	}
}
