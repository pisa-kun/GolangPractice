package main

import "fmt"

func main() {
	// 文字列にフォーマット
	s := fmt.Sprintf("%sの降水確率は%d%%です", "沖縄", 30)

	// 降水確率を出力
	fmt.Println(s)
}
