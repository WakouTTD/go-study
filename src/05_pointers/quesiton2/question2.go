package main

import "fmt"

func main() {

	// Q2 何が出力されるか
	/*
	   var i int = 100
	   var j int = 200
	   var p1 *int
	   var p2 *int
	   p1 = &i
	   p2 = &j
	   i = *p1 + *p2
	   p2 = p1
	   j = *p2 + i
	   fmt.Println(j)
	*/

	var i int = 100
	var j int = 200
	var p1 *int
	var p2 *int
	p1 = &i
	p2 = &j
	fmt.Println(p1, p2)
	// i=100+200で300になるが、p1 = &iなので、*p1も300になる
	i = *p1 + *p2
	fmt.Println(i, *p1, *p2)
	// p2=100
	p2 = p1
	fmt.Println(p2, *p2)
	// j=400
	j = *p2 + i
	fmt.Println(j, *p2, i)
	// 400が出力じゃない？
}
