package matrix

import (
	"fmt"
	"testing"
)

var rows = [4][4]float32{
	{1, 2, 3, 4},
	{5, 6, 7, 8},
	{9, 3, 2, 1},
	{5, 4, 7, 6},
}

func ExampleMatrix4x4v() {

	x := [4]float32{1, 2, 3, 4}
	var y [4]float32
	Matrix4x4v(&rows, &x, &y)
	fmt.Println(y)
	// Output:
	// [30 70 25 58]
}

func BenchmarkMatrix4x4v(b *testing.B) {
	x := [4]float32{1, 2, 3, 4}
	var y [4]float32

	for n := 0; n < b.N; n++ {
		Matrix4x4v(&rows, &x, &y)
	}
}
