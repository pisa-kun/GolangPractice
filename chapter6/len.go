package main

import "fmt"

func main() {
	// 配列型要素の宣言
	var array1 [1]int
	var array2 [2]*int
	var array3 [3][2]float64
	var array4 [2]struct{ x, y int }

	fmt.Println(len(array1))
	fmt.Println(len(array2))
	fmt.Println(len(array3), len(array3[0]))
	fmt.Println(len(array4))
}
