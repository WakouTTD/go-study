package main

import "fmt"

// スライス続き
func main() {

	// スライス
	n := []int{1, 2, 3, 4, 5, 6}
	n = append(n, 700)
	// d = d.append(d, 200)
	// d = d.append(d, 300)
	fmt.Printf("n type=%T value=%v\n", n, n)

	fmt.Println(n[2])
	fmt.Println(n[2:4])

	fmt.Println("-------------")
	// 何も書かないと最初から
	fmt.Println(n[:4])
	// 何も書かないと最後まで
	fmt.Println(n[2:])
	// 両方とも書かないと最初から最後まで全て
	fmt.Println(n[:])

	fmt.Println("-------------")
	n[2] = 300
	fmt.Println(n)

	n = append(n, 11, 12, 13, 14)

	fmt.Println(n)

	fmt.Println("-------------")

	var board = [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}
	fmt.Println(board)

	fmt.Println("-------------")
	var board2 = [][]int{
		{0, 1, 2},
		{},
		{6, 7, 8},
		{},
	}
	fmt.Println(board2)

	fmt.Println("-------------")

}
