package main

import "fmt"

// 構造体型の宣言
type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
	Person
}

func main() {
	// 構造体リテラルで初期化
	e := Employee{1, Person{"Nao", 28}}
	fmt.Println(e)
}
