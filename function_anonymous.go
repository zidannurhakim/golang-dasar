package main

import "fmt"

type setBlacklist func(string) bool

func registerUser(name string, blacklist setBlacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	//Model 1
	blacklist := func(name string) bool {
		return name == "Anjing"
	}
	registerUser("Zidan", blacklist)
	//Model 2
	registerUser("Anjing", func(name string) bool {
		return name == "Anjing"
	})
}
