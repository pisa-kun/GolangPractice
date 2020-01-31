package main

import "fmt"

// 四則演算を行う関数の宣下、4つの戻り値を持つ関数
// 戻り値に変数名をつけることで、return で返せる
func calc(a int, b int) (add int, sub int, mul int, div float32) {
	// 戻り値を返す
	add = a + b
	sub = a - b
	mul = a * b
	div = float32(a) / float32(b)

	// ret
	return
}

// main関数はパラメータ、戻り値をもたない
func main() {
	add, sub, mul, div := calc(1, 2)
	fmt.Println("add", add)
	fmt.Println("sub", sub)
	fmt.Println("mul", mul)
	fmt.Println("div", div)
}
