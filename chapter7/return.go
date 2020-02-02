package main

import (
	"fmt"
	"os"

)

func main() {
	// ファイルのオープン
	file, err := os.Open("test.txt")

	// オープンに成功したかの判定
	if err != nil {
		fmt.Println(err.Error())

		os.Exit(1)
	}

	file.Close()

	fmt.Println("OK")
}
