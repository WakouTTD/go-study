package main

import "fmt"

func one(x int) {
	x = 1
}

func two(x *int) {
	// ポイントが指す実体を「デリファレンス」という
	*x = 2
}

// ポインタ2
func main() {

	var n = 100
	one(n)
	// 参照渡しなので、nに変更は無い
	fmt.Println(n)

	// ポインタ型を引数にしている場合は、アドレスを渡さなければならないのでアンバサンド
	two(&n)
	// ２が表示される
	fmt.Println(n)
	fmt.Println("-----------")

	fmt.Println(n)
	fmt.Println(&n)
	fmt.Println(*&n)
	fmt.Println(&*&n)

}
