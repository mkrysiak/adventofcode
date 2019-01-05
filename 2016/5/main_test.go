package main

import "testing"

func BenchmarkPassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		password("wtnhxymk")
	}
}
