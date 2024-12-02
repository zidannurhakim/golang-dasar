package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

func main() {
	result := getHello("Zidan")
	fmt.Println(result)
	fmt.Println(getHello("Safira"))
}
