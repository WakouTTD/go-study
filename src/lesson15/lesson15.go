package main

import "fmt"

// byte
func main() {

	// https://www.ascii-code.com/
	b := []byte{72, 73}
	fmt.Println(b)
	fmt.Println(string(b))

	// キャストできる
	c := []byte("HIJ")
	fmt.Println(c)
	fmt.Println(string(c))
}
