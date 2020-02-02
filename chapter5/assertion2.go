package main

import "fmt"

func main() {
	// 空インターフェースにstringの値を格納
	var i interface{} = "test"

	// 型アサーション
	s1, ok := i.(string)
	fmt.Println(s1, ok)

	// 失敗
	// string型はdummyメソッドをもたないので変換できない
	s2, ok := i.(interface {
		dummy()
	})
	fmt.Println(s2, ok)
}
