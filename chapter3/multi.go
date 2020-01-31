package main

import "fmt"

// 四則演算を行う関数の宣下、4つの戻り値を持つ関数
func calc(a int, b int) (int, int, int, float32) {
	// 戻り値を返す
	return a + b, a - b, a * b, float32(a) / float32(b)
}

// main関数はパラメータ、戻り値をもたない
func main() {
	add, sub, mul, div := calc(1, 2)
	fmt.Println("add", add)
	fmt.Println("sub", sub)
	fmt.Println("mul", mul)
	fmt.Println("div", div)
}
