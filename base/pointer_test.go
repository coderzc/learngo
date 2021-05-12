package base

import (
	"fmt"
	"testing"
)

func swap(a, b int) {
	a, b = b, a
}

func swap2(a, b *int) {
	*a, *b = *b, *a
}

func TestPointer(t *testing.T) {
	var v int = 2
	var pv *int = &v
	*pv = 3
	fmt.Println(v)

	// pv = pv + 4 这是错的，go不支持指针运算

	a := 3
	b := 4

	swap(a, b)
	fmt.Println(a, b)

	swap2(&a, &b)
	fmt.Println(a, b)
}
