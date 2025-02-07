package Action_Demo

import (
	"testing"
)

func BenchmarkFib10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var _ = Fib(20)
	}
}

func BenchmarkFib20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var _ = Fib(10)
	}
}

func BenchmarkFib30(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var _ = Fib(30)
	}
}

func BenchmarkFib30(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var _ = Fib(20)
	}
}
