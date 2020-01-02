package main

import "fmt"

// パーレンティスでvarの連続を省略できる
var (
	j      int     = 1
	flt64  float64 = 1.2
	str    string  = "test"
	tt, ff bool    = false, true
)

func main() {

	// もっとも基本的な変数の宣言方法
	// var 変数名 型 = 初期値
	var i int = 1
	var f64 float64 = 1.2
	var s string = "test"
	var t, f bool = true, false

	fmt.Println(i, f64, s, t, f)

	// 関数の外で宣言している
	fmt.Println(j, flt64, str, tt, ff)

	//初期化しないでいるとデフォルト値が入る
	var noiniti int
	var noinitf64 float64
	var noinits string
	var noinitt, noinitf bool
	fmt.Println("noinit", noiniti, noinitf64, noinits, noinitt, noinitf)

	// 初期値を与えた時は型を省略して記述できる
	var firstnumber = 1234
	// 関数内であれば、varキーワードも:=を使用することで省略可能 ショートヴァリアブルデクレアレーション
	secondnumber := 5678
	fmt.Println(firstnumber, secondnumber)

	//定数の宣言には
}
