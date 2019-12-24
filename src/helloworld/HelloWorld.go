package main

import "fmt"

// fmt.Printlnか
// fmt.printlnか
// Pが大文字か小文字かで、Publicかprivateか
//

func init() {
	fmt.Println("init")
}

func bazz() {
	fmt.Println("bazz")
}

func main() {
	bazz()
	fmt.Println("hello, world\n", "test test")
}
