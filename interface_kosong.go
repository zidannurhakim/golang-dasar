package main

import "fmt"

/*
interface{} = any
*/
func Ups() interface{} {
	return "Ups"
}

func cobaAny() any {
	//return true
	//return 1
	return "Coba Any"
}

func main() {
	kosong := Ups()
	fmt.Println(kosong)

	var coba any = cobaAny()
	fmt.Println(coba)
}
