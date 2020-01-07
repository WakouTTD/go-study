package main

import (
	"fmt"
	"sync"
	"time"
)

/*
イメージとしてはmainファンクションから
①GoroutineのProducerを並列で走らせて
Producerの中の実行結果をchannelに入れて

②GoroutineのConsumerを並列で走らせて
channel経由でProducerから来た値を集めて処理をする
---------------
たとえば、
Producerでいろんなサーバーに行ってログを解析して取ってきて
Consumerでそれらの処理をする
*/

// producer　このプログラムではforループ内から10回呼ばれる
func producer(ch chan int, i int) {
	//Something
	ch <- i * 2
}

// consumer
func consumer(ch chan int, wg *sync.WaitGroup) {
	/*
		この書き方だとwg.Done()に到るまでの処理--ここだとfmt.Printlnだけだから発生しないが--
		で何かエラーが発生した場合、wg.Doneが呼ばれないので処理の異常が波及する
		for i := range ch {
			fmt.Println("process", i*1000)
			wg.Done()
		}
	*/

	for i := range ch {

		func() {
			// このインナーファンクションが終わったらwg.Done()するという書き方
			defer wg.Done()
			fmt.Println("process", i*1000)
		}()
	}

	//　この##は出力されない
	// consumerが呼ばれた直後にcloseされるので、rangeのループが終わった頃には
	// プログラム自体が終了
	fmt.Println("#############")
}

// producer consumer
func main() {

	var wg sync.WaitGroup
	ch := make(chan int)

	// 10回producerを作る、前処理としてDoneしてくださいね、というwg.Addする
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i)
	}

	// Consumerの中でDone chanのrange分forループするので呼び出しは1回
	go consumer(ch, &wg)
	// 全てのconsumerの処理の終了を待つ
	wg.Wait()

	// chanクローズしないといつまでもconsumerが終わらない。chのrange forループだから
	// まだchがあるのでは？って待ち続ける
	close(ch)

	// もしconsumer内の##を出力させてやりたければ、プログラムの終了を待つ必要がある
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}
