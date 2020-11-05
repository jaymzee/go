package main

import "testing"

var slice64 []float64

func FuncVar(f32 []float32) []float64 {
    f64 := make([]float64, len(f32))
    var f float32
    var i int
    for i, f = range f32 {
        f64[i] = float64(f)
    }
    return f64
}

func BenchmarkFuncVar(b *testing.B) {
    f32 := make([]float32, 1024)
    b.ReportAllocs()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        slice64 = FuncVar(f32)
    }
}

func RangeVar(f32 []float32) []float64 {
    f64 := make([]float64, len(f32))
    for i, f := range f32 {
        f64[i] = float64(f)
    }
    return f64
}

func BenchmarkRangeVar(b *testing.B) {
    f32 := make([]float32, 1024)
    b.ReportAllocs()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        slice64 = RangeVar(f32)
    }
}

