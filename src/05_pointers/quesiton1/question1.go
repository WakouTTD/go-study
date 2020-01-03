package main

import "fmt"

func main() {

	// Q1 何が出力されるか
	var i int = 10
	var p *int
	p = &i
	var j int
	j = *p
	fmt.Println(j)
}
