package main

import (
	"fmt"
	"os"

)

func main() {
	read()
}

// 読み込み
func read() {
	// ファイルのオープン
	file, err := os.OpenFile("test.txt", os.O_RDONLY, 0)

	// エラーチェック
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// ファイルクローズ
	defer func() {
		file.Close()
	}()

	// ファイルから文字列を読み込む
	var str string
	fmt.Fscanf(file, "%s", &str) // 改行またはスペースまで読み込む

	// 読み込んだ文字列の出力
	fmt.Println(str)

	// 以降のデータをbyteスライスに読み込む
	buf := make([]byte, 16)

	size, _ := file.Read(buf)

	// 読み込んだバイナリデータの出力
	fmt.Println(size, buf)
}
