package main

import "testing"

func TestSum(t *testing.T) {
	total := sum(10, 10)
	if total != 20 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 20)
	}
}
