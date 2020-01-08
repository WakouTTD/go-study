package main

import (
	"fmt"

	"08_package/61_public_and_private/mylib"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s))

	mylib.Say()
	person := mylib.Person{Name: "Mike", Age: 20}
	fmt.Println(person)

	fmt.Println(mylib.Public)

	// 小文字の変数は外パッケージから読めない
	// fmt.Println(mylib.private)
}
