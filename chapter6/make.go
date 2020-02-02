package main

import "fmt"

func main() {

	// 長さ10, キャパシティ20のスライスを作成
	s1 := make([]int, 10, 20)
	fmt.Println(s1)
	fmt.Println("len :", len(s1))
	fmt.Println("cap :", cap(s1))

	// 見やすくするため
	fmt.Println()

	// 長さ5, キャパシティ5のスライスを作成
	s2 := make([]float32, 5)
	fmt.Println(s2)
	fmt.Println("len :", len(s2))
	fmt.Println("cap :", cap(s2))
}
