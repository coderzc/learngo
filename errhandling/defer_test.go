package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("err occurred")
	defer fmt.Println(4)
	return
}

func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	err = errors.New("this is a ")
	if err != nil {
		fmt.Println("Error:", err)
		// Type assertion
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close() //writeFile("./fib.txt")

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fibonacci()
	for i := 0; i < 20; i++ {
		writer.WriteString(strconv.Itoa(f()) + "\n")
	}
}

func TestDefer(t *testing.T) {
	//tryDefer()
	writeFile("./fib.txt")
}
