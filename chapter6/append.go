package main

import "fmt"

func main() {

	// スライスを宣言
	s1 := []int{1, 2, 3, 4}

	// スライスに要素を追加
	// append(追加先のスライス、 追加する要素1, 追加する要素2, ...)
	s2 := append(s1, 5, 6)

	// スライスにスライスを追加
	s3 := append(s2, s1...)

	fmt.Println(s3)
}
