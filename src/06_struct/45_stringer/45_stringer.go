/*
Stringers

type Stringer interface {
	String() string
}

fmtパッケージのprintに書かれていて
Stringのメソッドをstructに定義していれば、
その通りにprint時に出力されるよ

*/

package main

import "fmt"

// Person struct　{ Name string Age int}
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	// return "My name is " + p.Name と同じ意味
	return fmt.Sprintf("My name is %v", p.Name)
}

// Stringer
func main() {

	mike := Person{"Mike", 22}
	// Stringerを定義していなければ　"{Mike 22}"　と表示される
	fmt.Println(mike)

}
