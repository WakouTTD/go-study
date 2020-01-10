package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Page ページ管理
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := "/Users/tateda/work2019/go-study/src/13_web_applications/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := "/Users/tateda/work2019/go-study/src/13_web_applications/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {

	//直接にResponseWriterを書いても良い
	//w.Header().Set("何か値")
	//w.Write([]byte{"<HTML></THMT>"})
	// /view/test　のtest部分がtitleに入る
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	// wはResponseWriter titleとbodyを埋め込んだhtmlを返す
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {

	//"/view/"にアクセスされたら、viewHandlerを呼び出す
	http.HandleFunc("/view/", viewHandler)

	// 8080はポート、コロンの前に何も書かなければ、localhost、nilを指定すると、デフォルトになりスラッシュでアクセスした時にpage not foundを返す
	// エラーを返した場合は、そのエラーを出力してサーバーを閉じるため、log.Fatalを使用
	log.Fatal(http.ListenAndServe(":8080", nil))
}
