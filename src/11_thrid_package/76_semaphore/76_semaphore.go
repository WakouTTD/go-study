package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

var s *semaphore.Weighted = semaphore.NewWeighted(1)

func longProcess(ctx context.Context) {
	isAcquire := s.TryAcquire(1)
	if !isAcquire {
		fmt.Println("Could not get lock")
		return
	}
	/*
		if err := s.Acquire(ctx, 1); err != nil{
			fmt.Println(err)
			return
		}
	*/
	defer s.Release(1)
	fmt.Println("Wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}

// サードパーティの便利はパッケージ紹介
// セマフォ (英: semaphore) とは、計算機科学において、
// 並列プログラミング環境での複数のプロセスが共有する資源にアクセスするのを制御する際の単純だが便利な抽象化を提供する変数または抽象データ
// semaphore
func main() {
	ctx := context.TODO()
	go longProcess(ctx)
	go longProcess(ctx)
	go longProcess(ctx)
	time.Sleep(5 * time.Second)
}
