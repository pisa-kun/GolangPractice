package main

import "fmt"

// 演算インターフェース型
type Calculator interface {
	// 関数の定義
	Calculate(a int, b int) int
}

// 足し算型
type Add struct {
	// フィールドはもたない
}

// Add型にCaluculatorインターフェースのCalculate関数を定義
func (x Add) Calculate(a int, b int) int {
	return a + b
}

type Sub struct {
}

func (x Sub) Calculate(a int, b int) int {
	return a - b
}

func main() {
	// Calculatorインターフェースを実装した型の変数
	var add Add
	var sub Sub

	// Calculatorインターフェース型の変数を宣言
	var cal Calculator

	cal = add

	fmt.Println("和:", cal.Calculate(1, 2))

	cal = sub

	fmt.Println("差:", cal.Calculate(1, 2))
}
