package main

// 先頭が大文字のためエクスポートされる
const Export = true

// 先頭が小文字だとエクスポートされない
const export = false

func main(){
	// 先頭が大文字でもローカル変数のためエクスポートされない
	const Z = 123
}