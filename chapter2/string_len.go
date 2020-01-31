package main

import "fmt"

func main(){
	var str string

	str = "a"

	str = str + "i"

	str += "u"

	var ja string = "Go言語"

	fmt.Println(str)
	fmt.Println(len(str))

	// utf08エンコードされたバイト数である「8」が出力される
	fmt.Println(ja, "len: ", len(ja))
}