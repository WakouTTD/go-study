package main

import "fmt"

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		c <- sum
	}
	close(c)
}

// rangeとclose
func main() {

	s := []int{1, 2, 3, 4, 5}

	//数がわからないなら、数の引数無し(アンバッファー)でも良いが、len(s)で数がわかるなら指定してやる
	c := make(chan int, len(s))
	go goroutine1(s, c)
	for i := range c {
		fmt.Println(i)
	}

}
