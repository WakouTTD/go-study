package main

import (
	"fmt"
	"strconv"
)

//      (x, y int)と書くこともできる
func plus(x int, y int) int {
	return x + y
}

func plusminus(x, y int) (int, int) {
	return x + y, x - y
}

func cal(price, item int) (result int) {
	result = price * item
	//ネイキッドリターンなので、return resultと書かなくてもresultが返る
	return
}

func convert(price int) float64 {
	return float64(price)
}

// 関数
func main() {
	// ショートデクレアレーションで返り値を受け取る
	r := plus(1, 1)
	// Itaoは、Integer to Ascii
	fmt.Println("plus function:" + strconv.Itoa(r))

	r1, r2 := plusminus(1, 1)
	// Itaoは、Integer to Ascii
	fmt.Println("plusminus function:" + strconv.Itoa(r1) + ", " + strconv.Itoa(r2))

	r3 := cal(100, 2)
	// Itaoは、Integer to Ascii
	fmt.Println("cal function:" + strconv.Itoa(r3))

	r4 := convert(10)
	// Itaoは、Integer to Ascii
	fmt.Println("convert function:" + strconv.FormatFloat(r4, 'f', 4, 64))

	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(1)

	// 変数に入れないで実行することもできる
	func(x int) {
		fmt.Println("inner func", x)
	}(2)

	// 並列化
	// go func(x int) {
	// 	fmt.Println("go inner func", x)
	// }(3)

}
