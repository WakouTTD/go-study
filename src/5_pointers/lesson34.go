package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

// struct　構造体(structure)
func main() {

	v := Vertex{X: 1, Y: 2}
	fmt.Printf("%T %v", v, v)

}
