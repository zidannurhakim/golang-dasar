package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Ahmad Zidan Nur Hakim",
		"address": "Banyuwangi",
	}
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Zidan"
	book["ups"] = "Salah"
	fmt.Println(book)
	delete(book, "ups")
	fmt.Println(book)
}
