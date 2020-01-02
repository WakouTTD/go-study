package main

import "fmt"

// Vertex 頂上　頂点
type Vertex struct {
	X int
	Y int
}

// struct　構造体(structure)
func main() {

	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)
	fmt.Printf("%T %v", v, v)

}
