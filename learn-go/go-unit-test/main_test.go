package main

import "testing"

func TestSum(t *testing.T) {
	total := sum(5, 5)

	if total != 11 {
		t.Errorf("suma was incorrect got: %d but need %d", total, 11)
	}
}
