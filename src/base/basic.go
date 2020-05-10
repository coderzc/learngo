package main

/*
go 变量
*/

import (
	"fmt"
	"math"
)

// 定义变量(会自动初始化有默认值)
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

// 初始化变量
func variableInitiaValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

// 自动类型推断
func variableTypeDeduction() {
	var a, b, s = 3, 4, "def"
	fmt.Println(a, b, s)
}

// 使用:= 定义并初始化变量
func variableSimple() {
	a, b, s := 3, 4, "def"
	fmt.Println(a, b, s)
}

// 包变量
// 在函数外面定义变量，不能使用:=的形式，并且可以不使用不会报错
// 错误的：name := "tom"
var name = "tom"

// 集中定义变量，一般包变量这么用
var (
	aaa = "aaa"
	bbb = 1
)

// 强制类型转换
func triangle() {
	a, b := 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

// 常量
// 可以定义不使用
func consts() {
	const (
		filename string = "abc.txt"
		a, b            = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

// 枚举 _ 是占位符
func enums() {
	const (
		cpp = iota //iota 自增值
		_
		python
		golang
		scala
	)

	fmt.Println(cpp, scala, python, golang)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitiaValue()
	variableTypeDeduction()
	variableSimple()
	triangle()
	consts()
	enums()
}
