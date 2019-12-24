package main

import (
	"fmt"
	"os/user"
	"time"
)

// packageはこのページを参照　https://golang.org/pkg/

func main() {
	fmt.Println("hello, world\n", time.Now())
	fmt.Println(user.Current())
}
