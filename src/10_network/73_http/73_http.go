package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// // URLアクセス　エラーハンドリングは後回し
	// resp, _ := http.Get("http://example.com")
	// // クローズ必要
	// defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	// fmt.Println("------------------------")

	// urlとして正しいか　正しくない場合はerr(ここでは'_'にしているが)にinvalidが入る
	base, _ := url.Parse("https://example.com/")
	reference, _ := url.Parse("/test?a=1&b=2")
	endpoint := base.ResolveReference(reference).String()
	//https://example.com/test?a=1&b=2　と出力される
	fmt.Println(endpoint)

	// req, _ := http.NewRequest("GET", endpoint, nil)
	req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte("password")))
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	q := req.URL.Query()
	q.Add("c", "3&%")
	fmt.Println(q)
	fmt.Println(q.Encode())
	req.URL.RawQuery = q.Encode()

	var client *http.Client = &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
