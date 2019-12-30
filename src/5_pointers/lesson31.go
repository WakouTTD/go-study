package main

import "fmt"

// ポインタ
func main() {

	var n int = 100
	fmt.Println(n)
	// アンバサンドを変数につけると　0xc000016098　といったようなメモリアドレスが16進数で表示される
	fmt.Println(&n)
	// integerのポイント型という意味　アンバサンドではなくアスタリスク
	var p *int = &n
	//0xc000016098
	fmt.Println(p)
	// 100
	fmt.Println(*p)

}
