package main

import (
	"fmt"

	"08_package/60_package/mylib"
	"08_package/60_package/mylib/under"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s))

	mylib.Say()
	under.Hello()
}
