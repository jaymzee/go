package main

import "testing"

func TestAddInt(t *testing.T) {
	got := AddInt(2, 3)
	if got != 5 {
		t.Errorf("AddInt(2, 3) = %d; want 6", got)
	}
}

func TestMultInt_small(t *testing.T) {
	lo, _ := MultInt(2, 3)
	if lo != 6 {
		t.Errorf("MultInt(2, 3) = %d; want 6", lo)
	}
}

func TestMultInt_big(t *testing.T) {
	a, b := 0x20003, 0x3002
	lo, hi := MultInt(a, b)
	if hi != 0 || lo != 0x60049006 {
		t.Errorf("MultInt(%#x, %#x) = %#x-%x; want 0x6-0", a, b, hi, lo)
	}
}

func TestMultFloat32(t *testing.T) {
	var a, b float32 = 2.0, 3.0
	got := MultFloat32(a, b)
	if got != 6.0 {
		t.Errorf("MultFloat32(%v, %v) = %v; want 6", a, b, got)
	}
}
