package main

import "testing"

func TestMultInt_bigger(t *testing.T) {
	a, b := 0x2_00000005, 0x3_00000007
	lo, hi := MultInt(a, b)
	if hi != 6 || lo != 0x1d_00000023 {
		t.Errorf("MultInt(%#x, %#x) = %#x-%x; want 0x6-0", a, b, hi, lo)
	}
}
