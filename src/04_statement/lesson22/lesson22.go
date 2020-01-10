package main

import "fmt"

// range
func main() {

	l := []string{"python", "go", "java"}
	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}

	// 同様の処理がgoのrangeで短く書ける
	for i, v := range l {
		fmt.Println(i, v)
	}
	// index番号使わない場合は、_にする
	for _, v := range l {
		fmt.Println(v)
	}

	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}

}



