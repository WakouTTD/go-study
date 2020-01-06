package main

import "fmt"

// Buffered Channels
// Goroutineとmainなど、並列で走っているスレッド間でデータのやりとりをする
func main() {

	// channelの数を2にしてみる
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
}
