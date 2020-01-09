package main

import (
	"encoding/json"
	"fmt"
)

// T 空のstruct
type T struct{}

// Person jsonの値を入れるstruct　marshal時に小文字でパラメータ化する指定付
type Person struct {
	// 空の場合には表示しないomitempty
	Name      string   `json:"name,omitempty"`
	Age       int      `json:"age,omitempty"`
	Nicknames []string `json:"nicknames"`
	T         *T       `json:"T,omitempty"`
}

/*
type Person struct {
	// nameを除いてjson化する場合は'-'を指定
	Name      string   `json:"-"`
	// Ageの値をstring値としてjson化
	Age       int      `json:"age, string"`
	// nicknamesの値をhandlenamesとして出力
	Nicknames []string `json:"handlenames"`
}
*/

// UnmarshalJSON メソッド名はこれじゃないとダメ
func (p *Person) UnmarshalJSON(b []byte) error {
	// このfunc内だけの独自のstruct
	type Person2 struct {
		Name string
	}
	var p2 Person2
	err := json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Println(err)
	}
	p.Name = p2.Name + "!"
	return err
}

// // MarshalJSON メソッド名はこれじゃないとダメ
// func (p Person) MarshalJSON() ([]byte, error) {
// 	// &で直接にstructを書いている
// 	v, err := json.Marshal(&struct {
// 		Name string
// 	}{
// 		Name: "Mr " + p.Name,
// 	})
// 	return v, err
// }

func main() {
	b := []byte(`{"name":"Mike","age":20,"nicknames":["a","b","c"]}`)
	var p Person
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)

	v, _ := json.Marshal(p)
	fmt.Println(string(v))
}
