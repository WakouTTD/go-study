package main

import "fmt"

// make
func main() {

	// makeはmapとsliceで使う
	// 空のスライス
	s := make([]int, 0)
	// []int
	fmt.Printf("%T\n", s)

	// 空のmap
	m := make(map[string]int)
	// map[string]int
	fmt.Printf("%T\n", m)

	// チャネルにはmake
	ch := make(chan int)
	fmt.Printf("%T\n", ch)

	var p *int = new(int)
	// ポインタを返すものにはnewを使う、そうでないものはmakeを使う
	// *int
	fmt.Printf("%T\n", p)

	// ストラクトにはnew
	var st = new(struct{})
	// *struct{}
	fmt.Printf("%T\n", st)

}
