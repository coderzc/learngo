package base

import (
	"fmt"
	"testing"
)

func TestSliceops(t *testing.T) {
	fmt.Println("create slice")
	var s []int //Zero value for slice is nil

	for i := 0; i < 100; i++ {
		//fmt.Println(s==nil)
		printSlice(s)
		s = append(s, 2*i+1)
	}

	s1 := []int{2, 3}
	fmt.Print("s1:")
	printSlice(s1)

	s2 := []int{2, 4, 6, 8}
	fmt.Print("s2:")
	printSlice(s2)

	s3 := make([]int, 10, 32)
	printSlice(s3)

	fmt.Println("copy slice")
	copy(s2, s1)
	fmt.Print("s2:")
	printSlice(s2)
	fmt.Println(s2)

	fmt.Println("delete slice")
	s2 = append(s2[:3], s[4:]...) // 展开列表
	fmt.Println(s2)
}

func printSlice(s []int) {
	fmt.Printf("len:%d,cap:%d\n", len(s), cap(s))
}
