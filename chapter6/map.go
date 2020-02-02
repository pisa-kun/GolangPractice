package main

import "fmt"

func main() {

	// マップを作成
	captials := make(map[string]string)

	// インデックスを使用してマップに値を格納
	captials["ja"] = "Tokyo"
	captials["us"] = "washinton"
	captials["ch"] = "pekin"

	// インデックスを使用して値を取り出す
	fmt.Println(captials)
	fmt.Println(captials["ja"])
	fmt.Println(captials["us"])
	fmt.Println(captials["ch"])

	fmt.Println()

	// rangeパターンで key, valueの順番で取り出せる
	for country, captial := range captials{
		fmt.Println(country, captial)
	}
}
