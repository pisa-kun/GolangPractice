package main

import "fmt"

func main() {

	// 長さ5の配列
	array1 := [5]float32{}

	// 残り要素には0
	array2 := [6]int{1, 2, 3, 4}

	// 要素数が長さとして使用
	array3 := [...]string{"one", "two", "three"}

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
}
