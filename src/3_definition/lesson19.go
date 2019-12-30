package main

import "fmt"

// 演習
func main() {

	// Q1
	f := 1.11
	i := int(f)
	fmt.Printf("%T %v\n", i, i)

	// Q2
	s := []int{1, 2, 5, 6, 2, 3, 1}
	// 5 6が出力される
	fmt.Println(s[2:4])

	//Q3
	// map[string]int map[Mike:20 Nancy:24 Messi:30]
	m := map[string]int{
		"Mike":  20,
		"Nancy": 24,
		"Messi": 30,
	}
	fmt.Printf("%T %v", m, m)
}
