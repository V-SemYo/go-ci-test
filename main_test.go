package main

import "testing"

func TestMaxInt(t *testing.T) {
	a, b := 2, 7

	result := MaxInt(a, b)

	if result != b {
		t.Errorf("expected: %d, got: %d", b, result)
	}
}
