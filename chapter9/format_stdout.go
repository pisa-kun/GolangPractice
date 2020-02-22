package main

import "fmt"

func main() {
	// 文字列にフォーマット
	fmt.Print("東京", "の降水確率は", 90, "%です\n")

	// 書式を使用して標準出力
	fmt.Printf("%sの降水確率は%d%%です", "福岡", 50)
}
