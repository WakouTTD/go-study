package main

import "fmt"

func main() {

	i := 1
	f64 := 1.2
	s := "test"
	t, f := true, false

	fmt.Println(i, f64, s, t, f)
	fmt.Printf("%T \n", f64)

	var f32 float32 = 1.2
	fmt.Printf("%T", f32)

}
