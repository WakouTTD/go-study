package main

import "fmt"

// for文
func main() {

	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("continue")
			continue
		}

		if i > 5 {
			fmt.Println("break")
			break
		}
		fmt.Println(i)
	}

	fmt.Println("---------------")

	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}

	// ただの無限ループ
	for {
		fmt.Println("Hello")
	}

}
