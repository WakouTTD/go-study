package main

import "fmt"

// Vertex 大文字のことをキャピタルと言う
type Vertex struct {
	// X座標、Y座標
	X, Y int
}

// Area 面積　範囲　ただの関数でVertexとの結びつきが無い
func Area(v Vertex) int {
	// 縦かける横
	return v.X * v.Y
}

// Area2 goにはクラスが無い代わりに別のやり方がある　値レシーバー
func (v Vertex) Area2() int {
	return v.X * v.Y
}

// Scale 書き換え アスタリスクが無いと書き換えにならない　ポインターレシーバー
func (v *Vertex) Scale(i int) {
	v.X = v.X * i
	v.Y = v.Y * i

}

// メソッドとポイントレシーバーと値レシーバー
func main() {

	v := Vertex{3, 4}
	fmt.Println(Area(v))
	// 上記の書き方だとAreaというfunctionが存在することを知ってなければならない
	// 下記の書き方だとv.と打った後にArea2のコード補完が出る
	fmt.Println(v.Area2())

	// 書き換え
	v.Scale(10)
	fmt.Println(v.Area2())

}
