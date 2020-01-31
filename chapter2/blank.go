// 変数代入
package main

import "fmt"

func main(){
	
	var a, b, c, d, e int

	// 同時に複数の変数に代入可能
	a, b = 1,2
	fmt.Println(a, b)

	// 2つ以上の値を返す関数fnの戻り値を代入する
	c, d = fn(10, 20)

	// 返り値を使用しないときはコンパイルエラーがでるらしい
	// goでは "_" は ブランク識別子として用いる
	_, e = fn(1,2)

	// 宣言した変数は必ず使用すること
	fmt.Println(c, d, e)
}

// fn(引数)(戻り値)
func fn(a int, b int)(add int, sub int){
	add = a + b
	sub = a - b

	return add, sub
}