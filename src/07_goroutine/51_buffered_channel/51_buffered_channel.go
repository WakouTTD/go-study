package main

import "fmt"

// Buffered Channels
// Goroutineとmainなど、並列で走っているスレッド間でデータのやりとりをする
func main() {

	// channelの数を2にしてみる
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))

	// x := <-ch
	// fmt.Println(x)
	// ch <- 300
	// fmt.Println(len(ch))

	// closeで終了しなければchの範囲がわからない
	close(ch)

	fmt.Println("-------------")
	for c := range ch {
		fmt.Println(c)
	}

}
