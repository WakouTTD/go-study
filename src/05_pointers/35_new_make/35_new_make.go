package main

import "fmt"

// newとmake
func main() {

	// var n int = 100
	// var p *int = &n

	// 値を入れない状態でメモリにポインタが入るメモリ領域(アドレス)を確保したい場合、newを使う
	// mapとスライスについてはmake使った
	var p *int = new(int)
	// 0xc000016098というようなアドレスが表示される
	fmt.Println(p)
	// デリファレンス(値)はデフォルトのゼロ
	fmt.Println(*p)
	*p++
	fmt.Println(*p)
	*p++
	fmt.Println(*p)

	// new を使わないと
	var p2 *int
	fmt.Println(p2) // nilが表示される
	// nilに大してインクリメントしようとするとpanicが発生する
	// panic: runtime error: invalid memory address or nil pointer dereference
	// *p2++
	// fmt.Println(*p2)

	fmt.Println("--------------")

	// make例
	// 空のスライス
	s := make([]int, 0)
	fmt.Printf("%T\n", s)

	// 空のmap
	m := make(map[string]int)
	fmt.Printf("%T\n", m)

	// チャネルにはmake
	ch := make(chan int)
	fmt.Printf("%T\n", ch)

	// 上で宣言した var p *int = new(int)
	// ポインタを返すものにはnewを使う
	fmt.Printf("%T\n", p)

	// ストラクトにはnew
	var st = new(struct{})
	fmt.Printf("%T\n", st)

}
