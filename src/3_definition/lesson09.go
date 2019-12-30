package main

import (
	"fmt"
)

// 論理値
func main() {

	var t, f bool = true, false
	fmt.Printf("t type=%T value=%v\n", t, t)
	fmt.Printf("f type=%T value=%v\n", f, f)

	a, b := true, false
	fmt.Printf("a type=%T value=%v\n", a, a)
	fmt.Printf("b type=%T value=%v\n", b, b)

	// printfの小文字 %t　the word true or false
	fmt.Printf("t bool word=%t\n", t)
	fmt.Printf("f bool word=%t\n", f)
	fmt.Printf("a bool word=%t\n", a)
	fmt.Printf("b bool word=%t\n", b)

	fmt.Println("-------------------------------------------")

	fmt.Println("true && true=", true && true)
	fmt.Println("true && true=", true && false)
	fmt.Println("false && false=", false && false)
	fmt.Println("false && true=", false && true)
	fmt.Println("-------------------------------------------")
	fmt.Println("true || true=", true || true)
	fmt.Println("true || true=", true || false)
	fmt.Println("false || false=", false || false)
	fmt.Println("false || true=", false || true)
	fmt.Println("-------------------------------------------")
	fmt.Println(!true)
	fmt.Println(!false)

}
