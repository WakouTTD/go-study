package main

import "fmt"

// 雛形
func main() {
	x := 0
	fmt.Println(x)
	x++
	fmt.Println(x)
	x--
	fmt.Println(x)

	fmt.Println("-- 次にシフトしてみる ---------")
	fmt.Println(1 << 0) // 0001 0001
	fmt.Println(1 << 1) // 0001 0010
	fmt.Println(1 << 2) // 0001 0100
	fmt.Println(1 << 3) // 0001 0100
	fmt.Println(1 << 4) // 0001 1000

}
