package main
import "fmt"

func main(){
	const G float32 = 9.8065
	const pi float64 = 3.14159

	// 複数の定数を一度に宣言
	const a, b, c int = 1, 2, 3

	// 〇カッコを使ったグルーピング
	const(
		ja string = "日本語"
		en string = "English"
	)

	// グルーピング時、値を省略すると前の行と同じ値になる
	const(
		X int = 1
		Y
		Z
	)

	fmt.Println(X,Y,Z)

	// IOTA列挙子
	const(
		ONE = iota
		TWO = iota
		THREE // 二番目以降は省略可能
		FOUR
	)
	fmt.Println(ONE, TWO, THREE, FOUR)
}