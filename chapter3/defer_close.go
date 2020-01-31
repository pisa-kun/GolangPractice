package main

import "os"

func main() {
	// ファイルのオープン
	file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE, 066)

	// オープンに失敗したときは終了
	if err != nil {
		os.Exit(1)
	}

	// ファイルのクローズを遅延実行
	defer file.Close()

	// ファイルへテキスト出力
	file.WriteString("aiueo")
}
