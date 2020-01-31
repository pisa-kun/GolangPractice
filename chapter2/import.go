package main

// インポートはパッケージのパスを記述
// import 識別子 "パッケージ"
import form "fmt"

// インポートするパッケージは必ず使用すること
// コンパイルエラーが発生
// import "archive/zip"

// import(
// 	"fmt"
// 	"time"
// )

func main(){
	var a int = 1
	form.Println(a)
}