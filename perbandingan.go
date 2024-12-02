package main

import "fmt"

func main() {
	var name1 = "Zidan"
	var name2 = "Zidan"

	var result bool = name1 == name2
	fmt.Println(result)

	var alpabet1 = "b"
	var alpabet2 = "a"
	//dibaca b>a maka true
	var result2 bool = alpabet1 > alpabet2
	fmt.Println(result2)
}
