package main

import "fmt"

func main() {
	var i interface{}

	i = 1
	fmt.Printf("type: %T, value: %v\n", i, i)

	i = "hello interface{}"
	fmt.Printf("type: %T, value: %v\n", i, i)

	// 声明a变量, 类型int, 初始值为1
	var a = 1
	// 声明i变量, 类型为interface{}, 初始值为a, 此时i的值变为1
	var i2 interface{} = a
	// 声明b变量, 尝试赋值i 需要强转
	var b = i2.(int)
	fmt.Println(b)

	fmt.Println(i2.(int))
}
