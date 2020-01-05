package main

import "fmt"

/*
通常のstructはブレイシズで囲って変数を宣言
type MyInt struct{
	myInt int
}
*/

// MyInt だよ
type MyInt int

// Double だがint返す
func (i MyInt) Double() int {
	fmt.Printf("%T %v \n", i, i)
	fmt.Printf("%T %v \n", 1, 1)
	// intにキャストして返してる
	return int(i * 2)
}

// non struct type にもメソッドが使える
// レアだからあまり使わないかもしれない
func main() {

	myInt := MyInt(10)
	fmt.Println(myInt.Double())

}
