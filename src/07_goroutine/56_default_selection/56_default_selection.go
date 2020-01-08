package main

import (
	"fmt"
	"time"
)

// default selection
func main() {

	// timeのTickはchannelを返す
	tick := time.Tick(100 * time.Millisecond)
	// timeのTickはchannelを返す
	boom := time.After(500 * time.Millisecond)

OuterLoop:
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		// こう書いたら時間も出力される
		// case t := <-tick:
		// 	fmt.Println("tick.", t)
		case <-boom:
			fmt.Println("BOOM!")
			// このreturnで、mainファンクションから抜ける
			//return
			// もし BOOM!出力後に#####を出力させたいなら
			break OuterLoop
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
	fmt.Println("#########")
}
