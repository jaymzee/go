package main

import (
	"fmt"
	"math"
)

func clamp(x float64, mn float64, mx float64) float64 {
	return math.Min(math.Max(x, mn), mx)
}

func main() {
	nums := []float64 {1.1, 2.1,3.4,4.5,5.6, 6.3, 7.9, 8.3, 9.5}
	for _, x := range nums {
		fmt.Println(clamp(x, 3.0, 7.0));
	}
}
