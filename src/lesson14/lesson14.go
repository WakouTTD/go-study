package main

import "fmt"

// map
func main() {
	//pythonの辞書型と同じようなもの
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["banana"] = 300
	fmt.Println(m)
	m["orange"] = 400
	fmt.Println(m)
	// 存在しないキーと設定するとデフォルト値のゼロになる
	fmt.Println(m["nothing"])

	fmt.Println("--------------")
	v, ok := m["apple"]
	fmt.Println(v, ok)

	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2)

	// 2つ目の返り値を無視することもできる
	v3 := m["banana"]
	fmt.Println(v3)

	fmt.Println("--------------")
	m2 := make(map[string]int)
	fmt.Println(m2)
	m2["pc"] = 5000
	fmt.Println(m2)

	var m3 = map[string]int{"aaa": 2}
	fmt.Println(m3)

	/*
		var m4 map[string]int
		panic: assignment to entry in nil map
		m4["pc"] = 5000
		fmt.Println(m4)
	*/
	fmt.Println("------------")
	var s []int
	if s == nil {
		fmt.Println("Nil")
	}
	// 出力は空配列のように`[]`と表示される
	fmt.Println(s)

}
