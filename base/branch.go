package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

/*
程序结构
*/

const filename = "abc.txt"

func main() {

	testIf()

	fmt.Println(
		grad(90),
		grad(0),
		grad(100),
		grad(70),
		grad(100),
	)

	fmt.Println(
		covertToBit(5),
		covertToBit(13),
		covertToBit(72387885),
		covertToBit(0),
	)

	printFile(filename)
}

// if
func testIf() {

	bytes, e := ioutil.ReadFile(filename)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Printf("%s\n", bytes)
	}

	if bytes, e := ioutil.ReadFile(filename); e != nil {
		fmt.Println(e)
	} else {
		fmt.Printf("%s\n", bytes)
	}
}

// switch
func grad(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Worong score %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}

	return g
}

// 整数转二进制表示
func covertToBit(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

// for 相当于 while
func printFile(filename string) {
	file, e := os.Open(filename)
	if e != nil {
		panic(e)
	}
	printFileContent(file)
}
func printFileContent(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	// 相当于while
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 死循环
func forever() {
	for {
		fmt.Println("abc")
	}
}
