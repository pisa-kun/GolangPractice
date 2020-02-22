package main

import (
	"fmt"
)

// stringを引数にとってstringを返す関数
// を返す関数
func returnDecoreteFunc() func(string) string{
	return func(str string) string{return "||" + str + "||"}
}

// 関数を引数にとる関数を定義することも可能
// 返り値がないのでvoid型
func decoreteAndPrint(str string, decoreteStrategy func(string) string){
	decoretedStr := decoreteStrategy(str)
	fmt.Printf("original string: %s -> decoreted string: %s", str, decoretedStr)
}

func main(){
	// 一度変数に受ける場合
	f := returnDecoreteFunc()
	ret1 := f("ruby")
	fmt.Println("decorete result ->", ret1)

	// 変数を経由しない場合
	ret2 := returnDecoreteFunc()("c#")
	fmt.Println("decorete result ->", ret2)
	
	// 第一引数のstringと第二引数の関数を元に結果を出力
	decoreteAndPrint("toyama nao", func(str string)string{
		return "ito " + str + " emoshi"
	})
}