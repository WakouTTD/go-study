package main

import (
	"fmt"
	"time"
)

func getOsName() string {
	return "alkdjnc"
}

// switch文
func main() {

	os := "mac"
	switch os {
	case "mac":
		fmt.Println("Mac!!")
	case "windows":
		fmt.Println("Windows!!")
	default:
		fmt.Println("Default!!")
	}
	fmt.Println(os)

	// switch文の中でしか使わないのであれば、1行で書ける
	switch os2 := getOsName(); os2 {
	case "mac":
		fmt.Println("Mac!!")
	case "windows":
		fmt.Println("Windows!!")
	default:
		fmt.Println("Default!!", os2)
	}
	// ただし、外でos2を使えない
	//fmt.Println(os2)

	// time.Now()は、goだと厳密に言うとオブジェクトではなくStruct
	t := time.Now()
	fmt.Println(t.Hour())
	fmt.Println(t.Minute())
	fmt.Println(t.Second())
	fmt.Println(t.Location())
	fmt.Println(t.Local())

	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 17:
		fmt.Println("Afternoon")
	case t.Hour() >= 18:
		fmt.Println("Night")
	}

}
