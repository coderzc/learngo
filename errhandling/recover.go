package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(err)
		}
	}()
	//panic(errors.New("this is a error"))
	panic(123)
}

func main() {
	tryRecover()
}
