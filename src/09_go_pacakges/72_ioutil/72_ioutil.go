package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

/*
lesson27でosパッケージのos.OpenFileでファイルを読み込んだが、
osパッケージは文字どおりosに関連するものを読み込む際に使われる

ioutilは、ネットワーク関連で用いられるパッケージ
たとえばパケットを読んだり書き出したりする際に使用
*/

// ioutil アイオーユーティル
func main() {
	// content, err := ioutil.ReadFile("/Users/tateda/work2019/go-study/src/09_go_pacakges/72_ioutil/72_ioutil.go")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(string(content))

	// if err := ioutil.WriteFile("/Users/tateda/work2019/go-study/src/09_go_pacakges/72_ioutil/ioutil_temp.go", content, 0666); err != nil {
	// 	log.Fatalln(err)
	// }

	r := bytes.NewBuffer([]byte("abc"))
	content, _ := ioutil.ReadAll(r)
	//byteなのでstringにコンバージョンして出力
	fmt.Println(string(content))
}
