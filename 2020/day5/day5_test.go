package main

import (
	"fmt"
	"testing"
)

func TestGetHighestSeatID(t *testing.T) {
	result := GetHighestSeatID("BFFFBBFRRR\nFFFBBBFRRR\nBBFFBBFRLL")
	fmt.Println(result)
	if result != 820 {
		t.Fatalf("Expected GetHighestSeatID(...) = %d; got %d", 820, result)
	}
}
