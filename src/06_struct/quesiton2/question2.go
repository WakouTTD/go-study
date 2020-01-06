package main

import (
	"fmt"
)

// Vertex 構造体
type Vertex struct {
	X, Y int
}

// X is 3! Y is 4! と表示されるStringerを作成してください、という課題の回答
func (v Vertex) String() string {
	return fmt.Sprintf("X is %v! Y is %v!", v.X, v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v)
}
