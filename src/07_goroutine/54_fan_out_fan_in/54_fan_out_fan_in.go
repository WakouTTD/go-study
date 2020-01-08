package main

import "fmt"

/*
Goroutine1
  Producer

Goroutine2
  Second Stage

Goroutine3
  Third Stage
*/

func producer(first chan int) {
	defer close(first)
	for i := 0; i < 10; i++ {
		first <- i
	}
}

// func multi2(first chan int, second chan int) {
// 下記のように<-でchannelを明示的にした方が良い
func multi2(first <-chan int, second chan<- int) {
	defer close(second)
	// 最初のgoroutine(producer)で作成されたchanから取得するのでrange使う
	for i := range first {
		second <- i * 2
	}
}

func multi4(second <-chan int, third chan<- int) {
	defer close(third)
	// 2番目ののgoroutine(multi2)で作成されたchanから取得するのでrange使う
	for i := range second {
		third <- i * 4
	}
}

// fan out fan in
func main() {

	first := make(chan int)
	second := make(chan int)
	third := make(chan int)

	go producer(first)
	// × 2　の意味
	go multi2(first, second)
	// × 4　の意味
	go multi4(second, third)

	for result := range third {
		fmt.Println(result)
	}
}
