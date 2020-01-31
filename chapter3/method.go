package main

import "fmt"

// int型からmyType型を宣言
type myType int

// mytypeをレシーバに持つ関数、すなわちmytype型のメソッドを宣言
func (value myType) println() {
	fmt.Println(value)
}

func main() {
	var z myType = 1234

	// myType型の変数zがレシーバとなって、
	// OOPっぽいレシーバにできる
	z.println()
}
