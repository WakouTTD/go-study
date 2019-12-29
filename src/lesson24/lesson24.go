package main

import "fmt"

func foo() {
	// foo関数の処理後にdeferを実行する
	defer fmt.Println("world foo")
	fmt.Println("hello foo")

}

// defer 遅延実行
func main() {

	foo()
	// ここでのdeferでは、main関数を処理後にdeferを実行しろ、という意味
	// hello が先に表示される
	defer fmt.Println("world")
	fmt.Println("hello")

}
