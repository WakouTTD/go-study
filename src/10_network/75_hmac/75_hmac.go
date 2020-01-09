package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// DB サーバーサイドが持っているDBの情報
var DB = map[string]string{
	"User1Key": "User1Secret",
	"User2Key": "User2Secret",
}

// Server 認証処理
func Server(apiKey, sign string, data []byte) {
	apiSecret := DB[apiKey]
	h := hmac.New(sha256.New, []byte(apiSecret))
	//
	h.Write(data)
	expectedHMAC := hex.EncodeToString(h.Sum(nil))
	// クライアントから送られてきた署名と期待値は一致するかを出力
	fmt.Println(sign == expectedHMAC)
}

// hmacはよくAPIでauthenticationのヘッダーに含める
func main() {
	// クライアントがサーバーに送り際にサービスが提供するAPIキーを持っている
	const apiKey = "User2Key"
	const apiSecret = "User2Secret"

	data := []byte("data")
	// sha256を使ってハッシュを作る
	h := hmac.New(sha256.New, []byte(apiSecret))
	fmt.Println(h)

	//そのハッシュにサーバーサイドに送りたいデータを書き込む
	h.Write(data)

	// 16進数でエンコード
	sign := hex.EncodeToString(h.Sum(nil))

	fmt.Println(sign)

	// サーバーに情報を投げる
	Server(apiKey, sign, data)
}
