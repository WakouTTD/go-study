package main

import (
	"fmt"
	"sync"
	"time"
)

// Counter インクリメント用struct
type Counter struct {
	v map[string]int

	// Counterにsync.Mutexを持たせているとgo-vetによる下記の警告が出たので修正
	// "call of fmt.Println copies lock value: 07_goroutine/57_sync_mutex.Counter contains sync.Mutex"
	//mux sync.Mutex
}

// Inc インクリメント
func (c *Counter) Inc(key string, mutex *sync.Mutex) {
	mutex.Lock()
	defer mutex.Unlock()
	c.v[key]++
}

// Value 取り出し
func (c *Counter) Value(key string, mutex *sync.Mutex) int {
	mutex.Lock()
	defer mutex.Unlock()
	return c.v[key]
}

// sync mutax
func main() {

	//c := make(map[string]int)
	//fatal error: concurrent map writes発生の危険性

	var mutex sync.Mutex
	c := Counter{v: make(map[string]int)}
	go func() {
		for i := 0; i < 10; i++ {
			//c["key"]++
			c.Inc("Key", &mutex)
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			//c["key"]++
			c.Inc("Key", &mutex)
		}
	}()
	time.Sleep(1 * time.Second)
	//fmt.Println(c, c["key"])

	fmt.Println(c, c.Value("Key", &mutex))

}
