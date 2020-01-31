package main

import "fmt"

func main(){
	var ptr *int

	var i int = 123

	ptr = &i

	fmt.Println("i : ", i)
	fmt.Println("ptr : ", *ptr)
	fmt.Println("i p: ", &i)
	fmt.Println("ptr p: ", ptr)

	// 値変更
	i = 321
	fmt.Println("i : ", i)
	fmt.Println("ptr : ", *ptr)
	fmt.Println("i p: ", &i)
	fmt.Println("ptr p: ", ptr)

	// ptrの値変更
	*ptr = 444
	fmt.Println("i : ", i)
	fmt.Println("ptr : ", *ptr)
	fmt.Println("i p: ", &i)
	fmt.Println("ptr p: ", ptr)
}