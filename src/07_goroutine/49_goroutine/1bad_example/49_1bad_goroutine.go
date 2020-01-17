package main

import (
	"fmt"
)

func goroutine(s string) {
	for i := 0; i < 5; i++ {
		//time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		//time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// Goroutine
// 軽量なスレッド、並列処理、他の言語だとマルチプロセス、マルチスレッド、イベント駆動とか言う
// Goだと暗黙的にやってくれる
func main() {

	// goroutineがスレッドを作成している間にnormalが処理を終わらせてしまう
	// helloのみが全出力した後、プログラムも終了するためworldは一件も出力されずに終わる
	// goroutineの処理が終わらなくてもプログラム自体は終わることを覚えておく必要がある
	go goroutine("world")
	normal("hello")

	//normalの処理終わっても、goroutineの処理を続行させるにはプログラム終了を待つ必要がある
	//time.Sleep(2000 * time.Millisecond)
}
