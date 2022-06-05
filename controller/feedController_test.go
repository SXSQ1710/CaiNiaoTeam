package controller

import "testing"

func BenchmarkSetUrlTestA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetUrlTestA(AllVideoList)
	}
}

func BenchmarkSetUrlTestB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetUrlTestB(AllVideoList)
	}
}
