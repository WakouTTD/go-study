package main

import "fmt"

// Human interfaceは実装を持たない、型のようなイメージ
type Human interface {
	Say() string
}

// Person struct
type Person struct {
	Name string
}

// Say Human interfaceの実装
func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}

// Say2 書き換え
func (p Person) Say2() {
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
}

// インターフェースのダックタイピング
// DriveCar Humanインターフェースを受け取る
func DriveCar(human Human) {
	if human.Say() == "Mr.Mike" {
		fmt.Println("Run")
	} else {
		fmt.Println("Get out")
	}
}

// インターフェースとダックタイピング
func main() {

	var mike Human = &Person{"Mike"}
	//mike.Say()
	// Say2は、Human　interfaceに無いため、Human型だとエラーになる
	//mike2.Say2()

	//fmt.Printf("%T %v\n", mike, mike)
	DriveCar(mike)
	var x Human = &Person{"X"}
	DriveCar(x)

	// Humanインターフェースに入れなくてもPerson型らしい
	var mike2 = Person{"Mike2"}
	fmt.Printf("%T %v\n", mike2, mike2)
	mike2.Say2()

}
