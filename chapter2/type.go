// 型の宣言
// func 型名 元になる型

package main

import "fmt"

func main(){
	// int型からintger型を宣言
	type myInteger int

	var i myInteger = 123

	i += 1

	fmt.Println(i)

	// 新しい構造体を作成
	type myStruct struct{
		a int
		b int
	}
}