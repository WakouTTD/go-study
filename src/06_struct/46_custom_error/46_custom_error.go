/*
stringerと同じようにカスタムエラーも
interfaceが実装されている

type error interface {
	Error() string
}
Stringerと同じように
Errorファンクションを実装すればカスタムエラーが作成できる
*/

package main

import (
	"fmt"
)

// UserNotFound エラー
type UserNotFound struct {
	Username string
}

// Error やろ
func (e *UserNotFound) Error() string {
	//	return e.Username
	return fmt.Sprintf("User not found: %v", e.Username)
}

func myFunc() error {
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{Username: "mike"}
}

// カスタムエラー
func main() {

	if err := myFunc(); err != nil {
		fmt.Println(err)
	}

	/*
		なぜ、	myFuncで &UserNotFound{"mike"}で、
		(e *UserNotFound) Error()なの？

		アンバサンド無しのアスタリスク無しでも良いのでは？
		ってのの理由としてアプリケーション開発だと
		同じエラーメッセージでも同じエラーではない場合があるため、
		アンバサンドとアスタリスクで別のものだと認識できるようにしておく

	*/
	e1 := UserNotFound{"mike"}
	e2 := UserNotFound{"mike"}
	fmt.Println(e1 == e2)

	e3 := &UserNotFound{"mike"}
	e4 := &UserNotFound{"mike"}
	fmt.Println(e3 == e4)

}
