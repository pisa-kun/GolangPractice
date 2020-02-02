package main

import "fmt"

func main() {

	// スライスを宣言
	dst := []int{1, 2, 3, 4}
	src := []int{5, 6, 7}

	// 要素のコピー
	// 戻り値はコピーされた要素の数
	// コピー先スライスとコピー元スライスのどちらか長さが短いほう
	count := copy(dst[2:], src)

	// 出力
	fmt.Println(dst)
	fmt.Println("count :", count)
}
