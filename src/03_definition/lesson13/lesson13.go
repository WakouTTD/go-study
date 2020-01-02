package main

import "fmt"

// スライスのmakeとcapacity
func main() {
	// intgerのsliceで、長さは3つ　キャパシティは5
	n := make([]int, 3, 5)
	fmt.Printf("n type=%T len=%d cap=%d value=%v\n", n, len(n), cap(n), n)

	n = append(n, 0, 0)
	fmt.Printf("n type=%T len=%d cap=%d value=%v\n", n, len(n), cap(n), n)

	n = append(n, 1, 2, 3, 4, 5)
	fmt.Printf("n type=%T len=%d cap=%d value=%v\n", n, len(n), cap(n), n)

	n = append(n, 11)
	fmt.Printf("n type=%T len=%d cap=%d value=%v\n", n, len(n), cap(n), n)

	fmt.Println("----------")
	// 長さ3でキャパシティ入れないと、長さもキャパシティも3になる
	a := make([]int, 3)
	fmt.Printf("a type=%T len=%d cap=%d value=%v\n", a, len(a), cap(a), a)

	// ゼロのスライス
	b := make([]int, 0)
	// nilになる
	var c []int
	fmt.Printf("b type=%T len=%d cap=%d value=%v\n", b, len(b), cap(b), b)
	fmt.Printf("c type=%T len=%d cap=%d value=%v\n", c, len(c), cap(c), c)

	fmt.Println("----------")

	d := make([]int, 3)
	for i := 0; i < 5; i++ {
		d = append(d, i)
		fmt.Println(d)
	}
	fmt.Println(d)

	fmt.Println("----------")

	e := make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		e = append(e, i)
		fmt.Println(e)
	}
	fmt.Println(e)

	/* anser
	----------
	[0 0 0 0]
	[0 0 0 0 1]
	[0 0 0 0 1 2]
	[0 0 0 0 1 2 3]
	[0 0 0 0 1 2 3 4]
	[0 0 0 0 1 2 3 4]
	----------
	[0]
	[0 1]
	[0 1 2]
	[0 1 2 3]
	[0 1 2 3 4]
	[0 1 2 3 4]
	*/
}
