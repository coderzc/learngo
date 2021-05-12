package algor

import (
	"testing"
)

func fib(n int) int {
	memo := []int{0, 1}
	for i := 2; i <= n; i++ {
		memo = append(memo, memo[i-1]+memo[i-2])
	}
	return memo[len(memo)-1]
}

func TestFib(t *testing.T) {
	tests := []struct{ input, output int }{
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13}}

	for _, tt := range tests {
		if res := fib(tt.input); res != tt.output {
			t.Errorf(`fib(%d); got %d; expr %d`, tt.input, res, tt.output)
		}
	}
}

func BenchmarkFib(b *testing.B) {
	input := 7
	output := 13
	for i := 0; i < b.N; i++ {
		if res := fib(input); res != output {
			b.Errorf(`fib(%d); got %d; expr %d`, input, res, output)
		}
	}
}
