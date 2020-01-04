package main

import "fmt"

// Vertex 図形
type Vertex struct {
	// x座標、y座標
	x, y int
}

// Area 面積
func (v Vertex) Area() int {
	return v.x * v.y
}

// Scale 書き換え アスタリスクが無いと書き換えにならない　ポインターレシーバー
func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

//-----------------

// Vertex3D 図形
type Vertex3D struct {
	// x座標、y座標
	Vertex
	z int
}

// Area3D 面積
func (v Vertex3D) Area3D() int {
	return v.x * v.y * v.z
}

// Scale3D 書き換え アスタリスクが無いと書き換えにならない　ポインターレシーバー
func (v *Vertex3D) Scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}

// New 小文字変数のVertexを生成
func New(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}

// embedded(組み込み) goには継承が無い代わりにembeddedがある
func main() {

	v := New(3, 4, 5)
	// 書き換え
	v.Scale3D(10)
	//fmt.Println(v.Area())
	fmt.Println(v.Area3D())

}
