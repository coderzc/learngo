package base

import (
	"fmt"
	"testing"
)

func printArray(arr [5]int) {
	arr[0] = 100
	for k, v := range arr {
		//fmt.Println(arr3[k])
		fmt.Println(k, v)
	}
}

func printArrayPointer(arr *[5]int) {
	(*arr)[0] = 100 //它还可以这么写：arr[0] = 100这点go有点混淆了
	for k, v := range arr {
		//fmt.Println(arr3[k])
		fmt.Println(k, v)
	}
}

func TestArray(t *testing.T) {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}

	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	printArray(arr1)
	printArray(arr3)

	// go的数组是按值传递的，这点和其他不太一样
	fmt.Println(arr1[0])

	fmt.Println(&arr1[0])
}
