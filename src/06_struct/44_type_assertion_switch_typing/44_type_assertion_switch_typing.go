package main

import "fmt"

// do iのinterfaceを引数に取る関数。空のインターフェースはどんな型でも引数として受けつける
func do(i interface{}) {
	/*
		// iがinterfaceなのでタイプアサーションが必要
		ii := i.(int)
		//i = ii * 2
		ii *= 2
		fmt.Println(ii)
			stringならstringのタイプアサーション
			ss := i.(string)
			fmt.Println(ss + "!")
	*/

	// switch type文
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("I don't know %T\n", v)

	}
}

// タイプアサーション
func main() {

	/*
		do(10)は、これと同じ
		var i interface{} = 10
		do(i)
	*/

	// 空のinterfaceなら何でも渡せる
	do(10)
	do("Mike")
	do(true)

	var i int = 10
	//キャストと言っていたが、goでは type conversionと言う
	ii := float64(i)
	fmt.Println(i, ii)
}
