package main

import (
	"fmt"
	"strconv"
)

// 型 キャスト
func main() {

	var x int64 = 1
	xx := float64(x)

	fmt.Printf("xx type=%T value=%v float=%f\n", xx, xx, xx)

	fmt.Println("------------------")

	var y float64 = 1.2
	yy := int64(y)
	fmt.Printf("yy type=%T value=%v integer=%d\n", yy, yy, yy)

	fmt.Println("------------------")

	var s string = "14"

	// var i = int64(s) // can not conversion string to int

	// 文字列からintへのキャストはできないため、strconvを使う
	//              Ascii to Integer
	i, _ := strconv.Atoi(s)
	//上記のアンダースコアは、本来ならerrだが、今回は使わないため、変数宣言してない
	// i, err := strconv.Atoi(s)
	// if err != nil {
	// 	fmt.Println("ERROR")
	// }
	fmt.Printf("i type=%T value=%v integer=%d\n", i, i, i)

	// byteから文字列へのキャストはできる
	h := "Hello World"
	fmt.Println(h[0])
	fmt.Println(string(h[0]))
}
