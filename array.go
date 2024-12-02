package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Ahmad"
	names[1] = "Zidan"
	names[2] = "Nur Hakim"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names)

	var values = [3]int{
		90, 80, 95,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	var values2 = [...]int{
		90, 80, 95, 40, 90,
	}
	fmt.Println(values2)
	fmt.Println(len(values2))
	values[0] = 100
	fmt.Println(values)
}
