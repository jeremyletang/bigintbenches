package main

import "testing"

func benchmarkFibGoff(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibGoff(i)
	}
}

func BenchmarkFibGoff1(b *testing.B)  { benchmarkFibGoff(1, b) }
func BenchmarkFibGoff2(b *testing.B)  { benchmarkFibGoff(2, b) }
func BenchmarkFibGoff3(b *testing.B)  { benchmarkFibGoff(3, b) }
func BenchmarkFibGoff10(b *testing.B) { benchmarkFibGoff(10, b) }
func BenchmarkFibGoff20(b *testing.B) { benchmarkFibGoff(20, b) }
func BenchmarkFibGoff40(b *testing.B) { benchmarkFibGoff(40, b) }

func benchmarkFibBigInt(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibBigInt(i)
	}
}

func BenchmarkFibBigInt1(b *testing.B)  { benchmarkFibBigInt(1, b) }
func BenchmarkFibBigInt2(b *testing.B)  { benchmarkFibBigInt(2, b) }
func BenchmarkFibBigInt3(b *testing.B)  { benchmarkFibBigInt(3, b) }
func BenchmarkFibBigInt10(b *testing.B) { benchmarkFibBigInt(10, b) }
func BenchmarkFibBigInt20(b *testing.B) { benchmarkFibBigInt(20, b) }
func BenchmarkFibBigInt40(b *testing.B) { benchmarkFibBigInt(40, b) }

// func benchmarkFibGmpInt(i int, b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		FibGmpInt(i)
// 	}
// }

// func BenchmarkFibGmpInt1(b *testing.B)  { benchmarkFibGmpInt(1, b) }
// func BenchmarkFibGmpInt2(b *testing.B)  { benchmarkFibGmpInt(2, b) }
// func BenchmarkFibGmpInt3(b *testing.B)  { benchmarkFibGmpInt(3, b) }
// func BenchmarkFibGmpInt10(b *testing.B) { benchmarkFibGmpInt(10, b) }
// func BenchmarkFibGmpInt20(b *testing.B) { benchmarkFibGmpInt(20, b) }
// func BenchmarkFibGmpInt40(b *testing.B) { benchmarkFibGmpInt(40, b) }

func benchmarkFibInt(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibInt(i)
	}
}

func BenchmarkFibInt1(b *testing.B)  { benchmarkFibInt(1, b) }
func BenchmarkFibInt2(b *testing.B)  { benchmarkFibInt(2, b) }
func BenchmarkFibInt3(b *testing.B)  { benchmarkFibInt(3, b) }
func BenchmarkFibInt10(b *testing.B) { benchmarkFibInt(10, b) }
func BenchmarkFibInt20(b *testing.B) { benchmarkFibInt(20, b) }
func BenchmarkFibInt40(b *testing.B) { benchmarkFibInt(40, b) }
