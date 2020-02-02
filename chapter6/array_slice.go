package main

import "fmt"

func main() {
	// sliceは参照型なので呼び出し先でスライスの中身を変えることが可能

	// 配列を宣言
	values := []int{0, 1, 2, 3, 4}

	// 配列を関数に受け渡す
	double(values)

	// valuesの値をコピーして関数に渡しているため、valuesがそのまま出力される
	fmt.Println(values)

	// スライスを関数に受け渡す
	double(values[:])

	// valuesの値が二倍になっている
	fmt.Println(values)
}

func double(values []int) {

	// 要素を倍にする
	for i := 0; i < len(values); i++ {
		values[i] *= 2
	}
}
