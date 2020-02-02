package main

import "fmt"

func main() {
	var str [7]string

	str[0] = "月曜日"
	str[1] = "火曜日"
	str[2] = "水曜日"
	str[3] = "木曜日"
	str[4] = "金曜日"
	str[5] = "土曜日"
	str[6] = "日曜日"

	// インデックスを使用して要素から値を取り出す
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}

	fmt.Println()

	// range
	// rangeから返ってくる第一変数はindexのためブランク使用
	for _, value := range str {
		fmt.Println(value, " ")
	}
}
