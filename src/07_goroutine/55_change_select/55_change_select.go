package main

import (
	"fmt"
	"time"
)

func goroutine1(ch chan string) {
	for {
		ch <- "packet from 1"
		time.Sleep(3 * time.Second)
	}
}

func goroutine2(ch chan string) {
	for {
		ch <- "packet from 2"
		time.Sleep(1 * time.Second)
	}
}

/*
1方でネットワークパケット受診
別方でネットワークパケットを受信

お互いのパケットをブロックすることないように作る
*/

// change select
func main() {

	c1 := make(chan string)
	c2 := make(chan string)
	go goroutine1(c1)
	go goroutine2(c2)

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
