package main

import "fmt"

func main(){
	// bool型の変数を宣言
	var b bool

	// bool型の変数にリテラル変数truとfalseを代入
	b = false
	b = true

	// bool型の変数に論理演算した結果を代入
	b = true || false

	// output
	fmt.Println(b)
}