package main

import "fmt"

// do iのinterfaceを引数に取る関数。空のインターフェースはどんな型でも引数として受けつける
func do(i interface{}) {
	// iがinterfaceなのでタイプアサーションが必要
	ii := i.(int)
	//i = ii * 2
	ii *= 2
	fmt.Println(ii)

}

// タイプアサーション
func main() {

	do(10)
	// do("Mike")
	// do(true)

}
