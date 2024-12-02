package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	goodBye := getGoodBye
	misalKan := getGoodBye
	fmt.Println(goodBye("Zidan"))
	fmt.Println(misalKan("Zidan"))
}
