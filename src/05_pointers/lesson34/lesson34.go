package main

import "fmt"

// Vertex 頂上　頂点
type Vertex struct {
	X int
	Y int
}

// VertexS Vertexにstring変数を加えたもの
type VertexS struct {
	X, Y int
	S    string
}

/*
構造体の変数は、
大文字でパブリック
小文字でプライベート

大文字：capital letters」「小文字はsmall letters
以前はcapital/smallが主流だったが，コンピュータの影響があり，
最近ではupper/lowerを使うようになった。
*/

// struct　構造体(structure)
func main() {

	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)
	fmt.Printf("%T %v", v, v)

	fmt.Println("-----------")
	v.X = 100
	fmt.Println(v.X, v.Y)

	fmt.Println("-----------")
	v２ := Vertex{X: 1}
	fmt.Println(v２)

	fmt.Println("-----------")
	// X以外は、デフォルト0,空文字
	vs := VertexS{X: 1}
	fmt.Println(vs)

	fmt.Println("-----------")
	vs2 := VertexS{10, 20, "test"}
	fmt.Println(vs2)

	fmt.Println("-----------")
	// デフォルトの0,0,空文字
	vsEmpty := VertexS{}
	fmt.Println(vsEmpty)
	fmt.Printf("vsEmpty: %T %v\n", vsEmpty, vsEmpty)

	fmt.Println("-----------")
	// varで宣言のみの場合、スライスやmapではnilだったが、struct構造体ではnilではない
	// vsEmpty := VertexS{}と同じ結果になる
	var varV VertexS
	fmt.Println(varV)
	fmt.Printf("var varV VertexS: %T %v\n", varV, varV)

	fmt.Println("----------- この後 v6 := new(VertexS)")
	// 値を入れない状態でメモリにポインタが入るメモリ領域(アドレス)を確保したい場合、newを使う
	// mapとスライスについてはmake使った

	v6 := new(VertexS)
	fmt.Println(v6)
	fmt.Printf("new(VertexS): %T %v\n", v6, v6)

	fmt.Println("----------- &VertexS{}")
	v7 := &VertexS{}
	fmt.Println(v7)
	fmt.Printf("&VertexS{}:   %T %v\n", v7, v7)
	/*
	   new(VertexS)と&VertexS{}は、同じ結果になるが、
	   実務では&VertexS{}の方が視認しやすいとして好まれる傾向が多い
	*/

	fmt.Println("----------- sliceの場合")
	/*
		値を入れない状態でメモリにポインタが入るメモリ領域(アドレス)を確保したい場合、newを使う
		mapとスライスについてはmake使った
		下記のような場合は、makeとエンプティ宣言でmakeが推奨されている
	*/
	// sliceの場合
	s := make([]int, 0)
	fmt.Println(s)
	fmt.Printf("slice   %T %v\n", s, s)

	s2 := []int{}
	fmt.Println(s2)
	fmt.Printf("slice   %T %v\n", s2, s2)

}
