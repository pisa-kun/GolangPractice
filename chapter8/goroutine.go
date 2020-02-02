package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main:開始")

	fmt.Println("testを通常の関数として実行")
	test()

	fmt.Println("testをゴルーチンとして実行")
	go test()

	// 3秒wait
	time.Sleep(3 * time.Second)

	fmt.Println("main:終了")
}

// ゴルーチンとして起動する関数
func test() {
	for i := 0; i < 5; i++ {
		// 連番を出力
		fmt.Println(i)

		// 1秒wait
		time.Sleep(1 * time.Second)
	}
}
