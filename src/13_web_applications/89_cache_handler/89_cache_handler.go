package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

// Page ページを管理するstruct
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := "/Users/tateda/work2020/go-study/src/13_web_applications/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {

	fmt.Println("loadPage")

	filename := "/Users/tateda/work2020/go-study/src/13_web_applications/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// キャッシュする
var templates = template.Must(template.ParseFiles("/Users/tateda/work2020/go-study/src/13_web_applications/edit.html", "/Users/tateda/work2020/go-study/src/13_web_applications/view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	fmt.Println("renderTemplate")

	// このrenderTemplateは、viewHandlerとeditHandlerが呼ばれる度にParseFilesするのはパフォーマンス的に良くないのでキャッシュすれば良い
	//t, _ := template.ParseFiles("/Users/tateda/work2020/go-study/src/13_web_applications/" + tmpl + ".html")

	// ExecuteTemplateでキャッシュから取得する際には、フルパスで書かない
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {

	fmt.Println("viewHandler")

	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {

	fmt.Println("editHandler")

	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {

	fmt.Println("makeHandler")

	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		// m[2]には(edit|save|view)以降の文字列が入る
		fn(w, r, m[2])
	}
}

func main() {

	fmt.Println("main")

	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
