package main

import "testing"

func TestAddInt(t *testing.T) {
	got := AddInt(2, 3)
	if got != 5 {
		t.Errorf("AddInt(2, 3) = %d; want 6", got)
	}
}
