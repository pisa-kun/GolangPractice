package main

import "fmt"

func main(){
	// for
	for i := 0; i < 10; i++{
		fmt.Println(i)
	}

	// goにはwhileが存在しない

	var count int = 0

	for{
		fmt.Println(count)
		count++

		// 終了条件
		if count == 5{
			break;
		}
	}

	// 拡張for はrange
	arr := [...]int{0,1,2,3,4}

	for i := range arr{
		fmt.Println(i)
	}
}