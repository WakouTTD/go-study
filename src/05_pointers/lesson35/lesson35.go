package main

import "fmt"

// Vertex struct
type Vertex struct {
	X, Y int
	S    string
}

func changeVertex(v Vertex) {
	v.X = 1000
}

func changeVertex2(v *Vertex) {
	v.X = 2000
}

func changeVertex3(v *Vertex) {
	(*v).X = 3000
}

// struct　構造体(structure) の続き
func main() {

	v := Vertex{1, 2, "test"}
	changeVertex(v)
	fmt.Println(v)
	fmt.Printf("%T %v\n", v, v)

	fmt.Println("------------")

	changeVertex2(&v)
	fmt.Println(v)
	fmt.Printf("%T %v\n", v, v)
	fmt.Println("------------")

	changeVertex3(&v)
	fmt.Println(v)
	fmt.Printf("%T %v\n", v, v)

	fmt.Println("------------")
	vp := &Vertex{1, 2, "test"}
	changeVertex2(vp)
	fmt.Println(vp)
	fmt.Println(*vp)
	fmt.Printf("%T %v\n", vp, vp)

}
