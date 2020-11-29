package main

import (
	"testing"
)

type Fixture struct {
	Instructions []int
	Expected     []int
}

func TestProcessInstructions(t *testing.T) {
	fixtures := []Fixture{
		{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}

	for _, fixture := range fixtures {
		t.Log(fixture)
		output := ProcessInstructions(fixture.Instructions, fixture.Instructions[1], fixture.Instructions[2])

		if output != fixture.Expected[0] {
			t.Errorf("Got %d as ouput; want %d", output, fixture.Expected[0])

		}
	}

}
