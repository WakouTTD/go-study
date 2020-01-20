package main

import "fmt"

// Buffered Channels
// Goroutineとmainなど、並列で走っているスレッド間でデータのやりとりをする
func main() {

	// channelの数を2にしてみる　bufferd channel
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))

	// 3つ目のchannelを追加するとどうなるか？
	// ch <- 300　// ここでruntime errorになる
	// fmt.Println(len(ch))

	// しかし、channelを1つ取り出すとまた2つまでなら、channelを追加できる
	x := <-ch
	fmt.Println(x)
	ch <- 300
	fmt.Println(len(ch))

	// closeで終了しなければchの範囲がわからないため、
	// 下記のrangeでerrorになる
	close(ch)

	fmt.Println("-------------")
	for c := range ch {
		fmt.Println(c)
	}

}
