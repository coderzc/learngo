package main

import "fmt"

func fib(n int) int {
	memo := []int{1, 2}
	for i := 2; i <= n; i++ {
		memo = append(memo, memo[i-1]+memo[i-2])
	}
	return memo[len(memo)-1]
}

func main() {
	fmt.Println(fib(10))
}
