package main

import "fmt"

func one(x int) {
	x = 1
}

func two(x *int) {
	// ポイントが指す実体を「デリファレンス」という
	*x = 2
}

// ポインタ2
func main() {

	var n = 100
	one(n)
	fmt.Println(n)

	two(&n)
	fmt.Println(n)
	fmt.Println("-----------")
	fmt.Println(n)
	fmt.Println(&n)
	fmt.Println(*&n)

}
