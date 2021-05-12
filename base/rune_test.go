package base

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestRune(t *testing.T) {
	s := "YES 我爱你中国"
	fmt.Println("lens is :", len(s))

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}

	fmt.Println("\n---")

	for i, ch := range s { // ch is rune
		fmt.Printf("(%d,%X)  ", i, ch)
	}

	fmt.Println("\n---")

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

	var zi rune = '国'
	fmt.Printf("zi type:%T\n", zi)
	fmt.Printf("aas:%c\n", zi)

	emjom := "😂"
	var emjomRune = []rune(emjom)[0]
	fmt.Printf("emjomRune:%c\n", emjomRune)

}
