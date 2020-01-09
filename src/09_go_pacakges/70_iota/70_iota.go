package main

import "fmt"

const (
	c1 = iota
	c2
	c3
)

const (
	_ = iota
	// KB 1024 シフト
	KB int = 1 << (10 * iota)
	// MB 1024 * 1024
	MB
	// GB 1024 * 1024 * 1024
	GB
)

// constの連番を振るiota untyped int
func main() {
	fmt.Println(c1, c2, c3)
	fmt.Println(KB, MB, GB)
}
