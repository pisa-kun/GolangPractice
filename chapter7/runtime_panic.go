package main

func main() {
	// 要素を3つもつ配列
	array := [...]int{1, 2, 3}

	index := 10

	// インデックス外への代入(ランタイムパニック)
	array[index] = index
}
