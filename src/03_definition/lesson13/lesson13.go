package main

import "fmt"

// スライスのmakeとcapacity
func main() {
	// intgerのsliceで、長さは3つ　キャパシティは5
	n := make([]int, 3, 5)
	// キャパシティはcap()で取得できる
	fmt.Printf("n type=%T len=%d cap=%d value=%v\n", n, len(n), cap(n), n)

	n = append(n, 0, 0)
	fmt.Printf("n type=%T len=%d cap=%d value=%v\n", n, len(n), cap(n), n)

	fmt.Println("そもそも5個のキャパシティのsliceに5個追加すると10個のキャパシティになる")

	n = append(n, 1, 2, 3, 4, 5)
	fmt.Printf("n type=%T len=%d cap=%d value=%v\n", n, len(n), cap(n), n)

	// 次に10個のキャパシティに1個appendすると、倍の20になる
	n = append(n, 6)
	fmt.Printf("n type=%T len=%d cap=%d value=%v\n", n, len(n), cap(n), n)

	fmt.Println("\n\n----------")
	// 長さ3でキャパシティ入れないと、長さもキャパシティも3になる
	a := make([]int, 3)
	fmt.Printf("a type=%T len=%d cap=%d value=%v\n", a, len(a), cap(a), a)

	// ゼロのスライス
	b := make([]int, 0)
	fmt.Printf("b type=%T len=%d cap=%d value=%v\n", b, len(b), cap(b), b)
	//b type=[]int len=0 cap=0 value=[]

	// nilになるが、表示上はゼロのスライスと同じ
	var c []int
	fmt.Printf("c type=%T len=%d cap=%d value=%v\n", c, len(c), cap(c), c)
	//c type=[]int len=0 cap=0 value=[]

	fmt.Println("\n\n----------")

	d := make([]int, 3)
	for i := 0; i < 5; i++ {
		d = append(d, i)
		fmt.Println(d)
	}
	fmt.Println(d)

	fmt.Println("\n\n----------")

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
