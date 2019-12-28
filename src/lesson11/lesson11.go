package main

import "fmt"

// 配列
func main() {

	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	fmt.Println("-------------")

	var b [3]int
	b[0] = 100
	b[1] = 200
	//b[2]、つまり３個目はデフォルトのゼロが表示される
	fmt.Println(b)

	fmt.Println("-------------")

	// 初期設定時にも3つ目を設定しないと、ゼロが入る
	var c [3]int = [3]int{100, 200}
	// 配列はappendできない
	//c.append(c, 300) //コンパイルできません
	fmt.Println(c)

	fmt.Println("-------------")

	// スライス
	var d = []int{100, 200}
	d = append(d, 300)
	// d = d.append(d, 200)
	// d = d.append(d, 300)
	fmt.Printf("%v\n", d)
}
