package main

import (
	"fmt"
	"log"
)

// https://golang.org/pkg/log/
// log
func main() {

	// 標準pkgのみだと機能が少ないため、サードパーティのloggingを使用すべきかもしれない
	log.Println("基本pkgのlog.Println")
	log.Printf("%T %v\n", "test", "test")

	log.Fatalln("基本pkgのlog.Fatalln")

	// このokは表示されない。なぜなら上記のfatalで終了してるから
	fmt.Println("ok")

}
