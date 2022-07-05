package main

import (
	"fmt"
	"unicode/utf8"
)

//rune相当于go的char
//使用range遍历pos，rune对
//使用len获得的是字节长度
//使用[]byte获得字节
//使用utf8.RuneCountInString获得字符数量

func main() {
	s := "lkidszzzz易大山"
	fmt.Println("s:", s)
	fmt.Println("len: ", len(s))
	//UTF-8
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	//unicode
	for i, ch := range s {
		//ch是rune
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}
	fmt.Println()
}

//其他字符串操作：
//Fields,Split,Join
//Contains,Index
//ToLower,ToUpper
//Trim,TrimRight,TrimLeft
