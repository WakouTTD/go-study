package main

import "fmt"

// フォーマットしたい時は下記のようにターミナルからgofmtを-wオプション付きで打つ
// gofmt -w /Users/tateda/work2019/go-study/src/3_definition/lesson06.go

func main() {

	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("10 - 1 =", 10-1)
	fmt.Println("10 / 2 =", 10/2)
	fmt.Println("10 / 3 =", 10/3)
	fmt.Println("10.0 / 3 =", 10.0/3)
	fmt.Println("10 / 3.0 =", 10/3.0)
	fmt.Println("10 % 2 =", 10%2)
	fmt.Println("10 % 3 =", 10%3)

}
