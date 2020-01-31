package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

// main関数はパラメータ、戻り値をもたない
func main() {
	answer := plus(114, 514)
	fmt.Println(answer)
}
