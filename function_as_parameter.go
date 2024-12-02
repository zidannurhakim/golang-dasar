package main

import "fmt"

type setFilter func(string) string

//type declaration sebagai alias untuk format data sehingga lebih efisien

func sayHelloWithFilter(name string, filter setFilter) {
	filteredName := filter(name)
	fmt.Println("Hello " + filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Zidan", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
