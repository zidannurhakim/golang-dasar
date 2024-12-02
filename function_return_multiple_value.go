package main

import "fmt"

func getFullName() (string, string) {
	return "Ahmad", "Zidan"
}

func main() {
	//Semua
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	//Pilih salah satu pakai "_"
	//firstName, _ := getFullName()
	//fmt.Println(firstName)
}
