package mylib

import "fmt"

// Public という文字列
var Public string = "Public"
var private string = "private"

// Person struct
type Person struct {
	Name string
	Age  int
}

// Say と出力
func Say() {
	fmt.Println("Human!")
}
