package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	sayHelloTo("Ahmad Zidan", "Nur Hakim")
	sayHelloTo("Safira Machruza", "Amilia")
}
