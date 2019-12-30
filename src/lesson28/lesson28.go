package main

import (
	"fmt"
	"log"
	"os"
)

// エラーハンドリング
func main() {

	file, err := os.Open("/Users/tateda/work2019/go-study/src/lesson28/lesson28.go")
	//エラーがnilでなければ、出力する
	if err != nil {
		log.Fatal("Error!")
	}
	defer file.Close()
	data := make([]byte, 500)
	// ショートデクレアレーション':='は、countの方を初期化しているが、errは、上記のerrを上書きしているだけ
	// ショートデクレアレーションは、1つ初期化すればエラーにならない
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("Error")
	}

	// ChdirのようにError1つの返り値の場合、ショートデクレアレーションを使えないため
	// 同じerr変数を使う場合は、イコールでオーバーライド
	err = os.Chdir("test")
	if err != nil {
		log.Fatalln("Error")
	}
	// 返り値が1つであれば次のようにも書ける
	if err = os.Chdir("test"); err != nil {
		log.Fatalln("Error")
	}

	fmt.Println(count, string(data))

}
