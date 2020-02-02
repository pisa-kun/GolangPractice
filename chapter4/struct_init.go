package main

import "fmt"

// 構造体型の宣言
type Person struct {
	name string
	age  int
}

func main() {

	// 構造体型リテラルを使用せずフィールドに個別に初期化
	var p1 Person
	p1.name = "hoge"
	p1.age = 11

	// 構造体リテラルで初期化
	p2 := Person{age: 31, name: "tom"}

	// 構造体リテラルで初期化2
	p3 := Person{"jane", 42}

	p4 := &Person{"mike", 22}

	fmt.Println(p1, p2, p3, p4)
}
