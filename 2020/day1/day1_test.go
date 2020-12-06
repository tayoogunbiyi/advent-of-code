package main

import (
	"reflect"
	"sort"
	"testing"
)

type Fixture struct {
	Numbers   []int
	TargetSum int
	Expected  []int
}

func TestTwoSum(t *testing.T) {
	fixtures := []Fixture{
		{[]int{1721, 979, 366, 299, 675, 14567}, 2020, []int{299, 1721}},
		{[]int{2, 7, 11, 9}, 9, []int{2, 7}},
		{[]int{-1, 0}, -1, []int{-1, 0}},
		{[]int{-1, 0}, -1, []int{-1, 0}},
	}

	for _, fixture := range fixtures {
		i, j := TwoSum(fixture.Numbers, fixture.TargetSum)
		a := fixture.Numbers[i]
		b := fixture.Numbers[j]
		if a > b {
			a, b = b, a
		}
		if !reflect.DeepEqual([]int{a, b}, fixture.Expected) {
			t.Errorf("TwoSum(%v) = %v,%v; want %v", fixture.Numbers, a, b, fixture.Expected)
		}
	}
}

func TestTwoSumWhenNoNumbersSatisfy(t *testing.T) {
	i, j := TwoSum([]int{1, 2, 3}, 100)

	if !(i == -1 && j == -1) {
		t.Errorf("expected %d,%d; got %d,%d", -1, -1, i, j)
	}

}

func TestThreeSum(t *testing.T) {
	fixtures := []Fixture{
		{[]int{1721, 979, 366, 299, 675, 14567}, 2020, []int{366, 675, 979}},
		{[]int{2, 7, 11, 9, 0}, 9, []int{0, 2, 7}},
	}

	for _, fixture := range fixtures {
		i, j, k := ThreeSum(fixture.Numbers, fixture.TargetSum)
		a := fixture.Numbers[i]
		b := fixture.Numbers[j]
		c := fixture.Numbers[k]
		got := []int{a, b, c}
		sort.Ints(got)
		if !reflect.DeepEqual(got, fixture.Expected) {
			t.Errorf("ThreeSum(%v,%d) = %v want %v", fixture.Numbers, fixture.TargetSum, got, fixture.Expected)
		}
	}
}

func TestThreeSumWhenNoNumbersSatisfy(t *testing.T) {
	i, j, k := ThreeSum([]int{1, 2, 3}, 100)

	if !(i == -1 && j == -1 && k == -1) {
		t.Errorf("expected %d,%d,%d; got %d,%d,%d", -1, -1, -1, i, j, k)
	}

}
