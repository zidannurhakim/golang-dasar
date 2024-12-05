package main

import "fmt"

/*
Interface adalah tipe data Abstract.
Interface tidak memiliki implementasi secara langsung.
Biasanya interface digunakan sebagai kontrak.
*/

type HasName interface {
	//Method "GetName" dengan parameter "string"
	GetName() string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

/*
	Setiap tipe data yang sesuai kontrak interface, secara otomatis dianggap sebagai interface tersebut.
	Hal ini agak berbeda dengan bahasa pemrograman lain yang ketika membuat interface,
	kita harus menyebutkan secara eksplisit akan menggunakan interface mana.
*/

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "Zidan"}
	SayHello(person)

	animal := Animal{Name: "Dino"}
	SayHello(animal)
}
