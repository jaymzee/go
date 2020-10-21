package main

import "testing"

func TestMultInt_bigger(t *testing.T) {
	a, b := 0x200000005, 0x300000007
	lo, hi := MultInt(a, b)
	if hi != 6 || lo != 0x1d00000023 {
		t.Errorf("MultInt(%#x, %#x) = %#x-%x; want 0x6-0", a, b, hi, lo)
	}
}
