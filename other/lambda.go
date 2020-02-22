// go言語における無名関数の例

package main

import (
	"fmt"

)

// 関数として定義
func decorate(str string) string {
	return "<<" + str + ">>"
}

func main() {
	// goでは関数を変数に格納可能
	f := decorate
	ret := f("Naobo")
	fmt.Println("decorete result ->", ret)

	// 変数に "string型を引数にとり srcを加工したstringを返す関数"をセット
	f2 := func(src string) string {
		return "<<" + src + ">>"
	}

	ret2 := f2("Toyama")
	fmt.Println("decorete result2 ->", ret2)
	fmt.Printf("function type -> %T\n", f2)
}
