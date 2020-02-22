
package main

import (
	"fmt"
)

func main(){
	f := sumCalc()
	for i := 1; i< 5; i++{
		srcVal, sumVal := f(i)
		fmt.Printf("add val : %d, sum val : %d\n", srcVal, sumVal)
	}
}

// int 引数を一つ取り、int変数を二つ返す関数
// を返り値にもつ関数
func sumCalc() func(int) (int,int){
	// 変数sumは関数が呼び出されるたびに初期化(0代入)されているようにみえるが
	// sumCalcに属する変数として、sumCaclが破棄されるまでメモリ上に値を残す
	sum := 0
	return func(x int) (int, int){
		sum += x
		return x, sum
	}
}