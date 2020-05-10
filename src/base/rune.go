package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "YES 我爱你中国"
	fmt.Println("lens is :", len(s))

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}

	fmt.Println()

	for i, ch := range s {
		fmt.Printf("(%d,%d)  ", i, ch)
	}

	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	bytes := []byte(s)

	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		fmt.Printf("%c ", ch)
		bytes = bytes[size:]
	}

	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("%d, %c\n", i, ch)
	}

	subStr := s[5:]

	fmt.Println(subStr)

}
