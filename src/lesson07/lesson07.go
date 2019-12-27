package main

import "fmt"

// シフト
func main() {
	x := 0
	fmt.Println(x)
	x++
	fmt.Println(x)
	x--
	fmt.Println(x)

	fmt.Println("-- 次にシフトしてみる ---------")
	fmt.Println(1 << 0) // 0001 00001
	fmt.Println(1 << 1) // 0001 00010
	fmt.Println(1 << 2) // 0001 00100
	fmt.Println(1 << 3) // 0001 01000
	fmt.Println(1 << 4) // 0001 10000

}
