package main

import "fmt"

// Pi big P means Public
const Pi = 3.14

// ユーザ名とパスワード
const (
	Username = "test_user"
	Password = "test_pass"
)

//int64の最大値を超えると宣言でエラー
//var big int64 = 9223372036854775807 + 1
// ところが、const は untype なので宣言時ならint64最大を超えられる
const big = 9223372036854775807 + 1

func main() {
	fmt.Println(Pi, Username, Password)
	fmt.Println(big - 1)
}
