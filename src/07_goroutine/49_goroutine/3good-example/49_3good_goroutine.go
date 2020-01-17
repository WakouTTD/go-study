package main

import (
	"fmt"
	"sync"
)

func goroutine(s string, wg *sync.WaitGroup) {
	// Done()はdeferにした方が見やすい
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

// Goroutine
// 軽量なスレッド、並列処理、他の言語だとマルチプロセス、マルチスレッド、イベント駆動とか言う
// Goだと暗黙的にやってくれる
func main() {

	var wg sync.WaitGroup
	// 1つのsyncグループがありますよという宣言
	wg.Add(1)
	// waitGroupのポインタを渡す
	go goroutine("world", &wg)
	normal("hello")

	// 終了するのを待ってやらないと、worldが出力し終わる前に、helloの出力が終わった時点で、
	// プログラムが終わる。
	// goroutineの処理が終わらなくても、プログラムの処理自体が終わる
	//　そうならないようにするためにsyncのwaitグループがある
	// waitGroupに1つの並列処理があることを告げているため、
	// time.Sleepを使わなくても
	// goroutineメソッド内のDone()が実行されるまで待ってくれる
	wg.Wait()
}
