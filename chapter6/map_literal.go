package main

import "fmt"

func main() {

	// マップを作成
	captials := map[string]string{
		"ja": "Tokyo",
		"us": "washinton",
		"ch": "pekin"}

	// インデックスを使用して値を取り出す
	fmt.Println(captials)
	fmt.Println(captials["ja"])
	fmt.Println(captials["us"])
	fmt.Println(captials["ch"])

	fmt.Println()

	// rangeパターンで key, valueの順番で取り出せる
	for country, captial := range captials {
		fmt.Println(country, captial)
	}
}
