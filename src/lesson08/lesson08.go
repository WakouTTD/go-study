package main

import (
	"fmt"
	"strings"
)

// 文字列
func main() {

	fmt.Println("Hello World")
	fmt.Println("Hello " + "World")
	// スライスのゼロ番のアスキーコード
	fmt.Println("Hello World"[0])
	// スライスのゼロ番の文字
	fmt.Println(string("Hello World"[0]))

	var s string = "Hello World"
	// こういうのはできません
	//s[0] = "X"

	fmt.Println(strings.Replace(s, "H", "X", 1))
	fmt.Println(strings.Replace("HHHHH", "H", "X", 4))

	// sは置き換わっていないため、Hello Worldのまま
	fmt.Println(s)
	// 置き換えたいならsに代入
	s = string(strings.Replace(s, "H", "X", 1))
	fmt.Println(s)

	fmt.Println(strings.Contains(s, "World"))

	fmt.Println("test\n" +
		"test")

	fmt.Println(`test
	　　　　　　　　　　　　　　　test
test`)

	fmt.Println("\"")
	fmt.Println(`"`)

}
