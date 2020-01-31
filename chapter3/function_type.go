package main

import "fmt"

func main() {

	// 関数型の変数宣言
	var f func(int, int) int

	// 関数型の変数に、関数リテラルを代入
	f = func(a int, b int) int {
		return a + b
	}

	fmt.Println(f(22, 33))

	// 関数型の変数に関数の値を代入
	f = multiply

	fmt.Println(f(22, 33))
}

func multiply(x int, y int) int {
	return x * y
}
