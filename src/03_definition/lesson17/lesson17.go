package main

import "fmt"

func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

// どんな時にクロージャを使うのか
func circleArea(pi float64) func(radius float64) float64 {
	// pi 円周率　radius 半径
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

// クロージャ
func main() {

	counter := incrementGenerator()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	// 円周率を3.14にしてたまま、半径だけ変えて2回計算
	c1 := circleArea(3.14)
	fmt.Println(c1(2))
	fmt.Println(c1(3))

	// 円周率を3にしてたまま、半径だけ変えて2回計算
	c2 := circleArea(3)
	fmt.Println(c2(2))
	fmt.Println(c2(3))

}
