package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 3)
	if total != 8 {
		t.Error("Fucking Error")
	}
}
