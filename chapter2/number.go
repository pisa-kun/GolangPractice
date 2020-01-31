package main

import "fmt"

func main(){
	// int型の変数を代入
	var i int = 12345

	// int64変換
	var i64 int64 = int64(i)

	fmt.Println(i64)
}