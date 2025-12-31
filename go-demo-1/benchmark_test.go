package main

import "testing"

// Бенчмарк для варианта с прямой записью в слайс
func BenchmarkCalculateSumInParallel(b *testing.B) {
	arr := make([]int, 10000)
	for i := range arr {
		arr[i] = i + 1
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = calculateSumInParallel(arr, 3)
	}
}

// Бенчмарк для варианта с каналами
func BenchmarkCalculateSumInParallelWithChannels(b *testing.B) {
	arr := make([]int, 10000)
	for i := range arr {
		arr[i] = i + 1
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = calculateSumInParallelWithChannels(arr, 3)
	}
}

