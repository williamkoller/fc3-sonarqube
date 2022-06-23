package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(20, 20)

	if total != 40 {
		t.Error("The total must be 40")
	}
}