package main

import (
	"fmt"
)

// Vertex 構造体
type Vertex struct {
	X, Y int
}

// Plus 7と表示されるメソッドを作成してください。という課題の回答
func (v Vertex) Plus() int {
	return v.X + v.Y
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Plus())
}
