package main

import "fmt"

func thirdPartyConnectDB() {
	// panicは自分で例外を投げられる　stdError
	panic("Unable to connect database!")
}

func save() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	// defer recover が先に無ければpanicをキャッチできない
	thirdPartyConnectDB()
}

// panic とrecover
// まぁできるだけプロダクトコードでpanicは書かないことが推奨されている
func main() {

	save()
	fmt.Println("OK?")

}
