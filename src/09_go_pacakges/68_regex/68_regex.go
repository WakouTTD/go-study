package main

import (
	"fmt"
	"regexp"
)

// 正規表現の学習
func main() {
	//最初がaで、間がa-zで1個以上あるもので、末尾がeなので、appleはtrue
	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match)

	match2, _ := regexp.MatchString("a([a-z0-9]+)e", "app0e")
	fmt.Println(match2)

	//高速で何度も使うマッチングであれば、下記のようにMustCompileを使う
	r := regexp.MustCompile("a([a-z]+)e")
	ms := r.MatchString("apple")
	fmt.Println(ms)

	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	// ひっかかった文字列全部を取ってくるFindString
	fs := r2.FindString("/view/test")
	fmt.Println(fs)

	fss := r2.FindStringSubmatch("/view/test")
	// 0番目が全部 1番目 view 2番目 test
	fmt.Println(fss, fss[0], fss[1], fss[2])

	fss = r2.FindStringSubmatch("/edit/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])

	fss = r2.FindStringSubmatch("/save/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])
}
