package main

import (
	"fmt"
	"io/ioutil"
)

// Page ページを管理するstruct
type Page struct {
	Title string
	Body  []byte
}

// save テキストに情報を書き込むメソッド
func (p *Page) save() error {
	filename := "/Users/tateda/work2019/go-study/src/13_web_applications/85_web_applications/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := "/Users/tateda/work2019/go-study/src/13_web_applications/85_web_applications/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

/*
簡単なwebページを作るクイックスタート
https://golang.org/doc/articles/wiki/
*/
func main() {
	p1 := &Page{Title: "test", Body: []byte("This is a sample Page.")}
	p1.save()

	p2, _ := loadPage(p1.Title)
	fmt.Println(string(p2.Body))
}
