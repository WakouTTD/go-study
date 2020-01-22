package main

import "fmt"

// new
func main() {

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

	fmt.Println("--------------")

	// new を使わないと
	var p2 *int
	fmt.Println(p2) // nilが表示される
	// nilに大してインクリメントしようとするとpanicが発生する
	// panic: runtime error: invalid memory address or nil pointer dereference
	// *p2++
	// fmt.Println(*p2)
}
