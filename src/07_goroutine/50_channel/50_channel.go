package main

import "fmt"

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func goroutine2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

// Channel
// Goroutineとmainなど、並列で走っているスレッド間でデータのやりとりをする
func main() {

	s := []int{1, 2, 3, 4, 5}

	// integerを受け付けるchannel　キューのような役割
	c := make(chan int)
	go goroutine1(s, c)
	go goroutine2(s, c)

	// ブロッキングでsumからデータが入ってくるまでずっと待ってる
	// なのでsync waitで待たなくても良い
	x := <-c
	fmt.Println(x)
	y := <-c
	fmt.Println(y)
}
