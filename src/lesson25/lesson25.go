package main

import (
	"fmt"
	"os"
)

// defer 遅延実行 続き
func main() {

	fmt.Println("run")
	// deferはスタック形式なのでLIFO、つまり3,2,1と出力される
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("success")

	// deferの使い方をファイルオープンで説明
	file, _ := os.Open("/Users/tateda/work2019/go-study/src/lesson25/lesson25.go")
	defer file.Close()

	fmt.Println("-----------")
	// byteスライス
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))
	fmt.Println("-----------")

}
