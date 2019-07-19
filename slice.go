package main

import "fmt"

/*
切片
*/

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr[2:6])
	fmt.Println(arr[:6])
	fmt.Println(arr[:])

	s1 := arr[2:]
	s2 := arr[:]

	fmt.Println("s1 = ", s1)
	updateSlice(s1)
	fmt.Println("s1 update= ", s1)
	fmt.Println("arr = ", arr)

	fmt.Println("s2 = ", s2)
	updateSlice(s2)
	fmt.Println("s2 update= ", s2)
	fmt.Println("arr = ", arr)

	s3 := s1[:5]
	fmt.Println("s3 = ", s3)

	fmt.Println("Extending Slice-------")
	arr2 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	k1 := arr2[2:6]
	fmt.Println("k1：", k1)

	//println("k1[4]：",k1[4]) //这样会报错

	k2 := k1[3:5] // 这里不会报错
	fmt.Println("k2：", k2)

	arr3 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	m1 := arr3[2:6]
	m2 := m1[3:5]

	fmt.Println("m1：", m1)
	fmt.Println("m2：", m2)

	m3 := append(m1, 10)
	m4 := append(s3, 11)
	m5 := append(m4, 12)

}
