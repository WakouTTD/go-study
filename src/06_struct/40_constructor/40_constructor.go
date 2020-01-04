package main

import "fmt"

// Vertex 変数を小文字にしたため、同じパッケージ内からしか生成できない
type Vertex struct {
	// x座標、y座標
	x, y int
}

// Area 面積　範囲
// func Area(v Vertex) int {
// 	// 縦かける横
// 	return v.x * v.y
// }

// Area2 goにはクラスが無い代わりに別のやり方がある　値レシーバー
func (v Vertex) Area2() int {
	return v.x * v.y
}

// New 小文字変数のVertexを生成
func New(x, y int) *Vertex {
	return &Vertex{x, y}
}

// Scale 書き換え アスタリスクが無いと書き換えにならない　ポインターレシーバー
func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i

}

// メソッド
func main() {

	// v := Vertex{3, 4}
	// fmt.Println(Area(v))
	// //
	// fmt.Println(v.Area2())

	/*
	 こうやってNew というfuncを使って他のパッケージからstructを呼び出すことが多い
	 デザインパターンというよりGoのよくあるパターン
	 https://golang.org/search?q=New
	 errors.New,log.New,rand.Newなどなど
	 パッケージ名ドットNewで、structを返す
	*/
	v := New(3, 4)
	// 書き換え
	v.Scale(10)
	fmt.Println(v.Area2())

}
