// 型の変換
// 変換先の型(変換する型)

package main

import "fmt"

func main(){
	
	// int型の変数を宣言
	var i int = 1234

	// int型からuint32
	var u uint32 = uint32(i)

	var f float32 = float32(u)

	var s string = string(i)

	// string型から配列(正確にはスライス)への変換
	var b []byte = []byte("abc")

	fmt.Println(i)
	fmt.Println(u)
	fmt.Println(f)
	fmt.Println(s)
	fmt.Println(b)
}