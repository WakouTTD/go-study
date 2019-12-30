package main

import "fmt"

// 演習
func main() {

	// Q1 一番小さい値を出力
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	var min int
	for i, v := range l {

		if i == 0 {
			min = v
			continue
		}

		if v < min {
			min = v
		}
	}
	fmt.Println(min)

	// 価格の合計
	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}
	var sum int = 0
	for _, v := range m {
		sum += v
	}
	fmt.Println(sum)

}
