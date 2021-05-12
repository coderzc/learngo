package base

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
	"testing"
)

/*
函数
*/

// 返回多个值  带余数除法
func div(a, b int) (int, int) {
	return a / b, a % b
}

// 返回多个值 并给返回值 指定变量名
func div2(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return q, r
}

//  a,b int ---> a,b 都定义成 int类型
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		//panic(fmt.Sprintf("unsupported op :%s", op))
		return 0, fmt.Errorf("unsupported op ---> %s\n", op)
	}
}

/*
类似函数指针的功能,回调函数
*/
type cb func(int, int) int

func apply(op cb, a, b int) int {
	pointer := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(pointer).Name()
	fmt.Printf("Callback func ----> %s ; with args ----> (%d,%d)\n", opName, a, b)

	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

/**
变长参数
*/
func sum(numbers ...int) int {
	s := 0
	for i, number := range numbers {
		s += number
		fmt.Println(i)
	}
	return s
}

func TestFunc(t *testing.T) {

	i, i2 := div(13, 3)
	fmt.Printf("q:%d，r:%d\n", i, i2)

	q, r := div2(13, 3)
	fmt.Printf("q:%d，r:%d\n", q, r)

	// 只要一个值怎么用 ----> 用_占位符
	q2, _ := div2(13, 3)
	fmt.Println(q2)

	/*
		go 的多返回值函数与python不同，python实质上是返回会的一个list，而go不是
		第二个返回值一般用于error，go官法建议使用error返回值手动处理每一个异常，go没有try catch
	*/
	if i3, err := eval(3, 4, "*"); err != nil {
		fmt.Printf("error:%s", err)
	} else {
		fmt.Println(i3)
	}

	if i3, err := eval(3, 4, "="); err != nil {
		fmt.Printf("error:%s", err)
	} else {
		fmt.Println(i3)
	}

	fmt.Println(apply(pow, 3, 4))

	// 匿名回调函数写法 (类似于java匿名类)
	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5))
}
