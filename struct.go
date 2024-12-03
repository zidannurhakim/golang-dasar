package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// Program Struct Method
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {
	//Model 1
	var cust Customer
	cust.Name = "Ahmad Zidan N.H."
	cust.Address = "Banyuwangi, East Java, Indonesia"
	cust.Age = 25
	fmt.Println(cust)
	fmt.Println(cust.Name)
	fmt.Println(cust.Address)
	fmt.Println(cust.Age)

	//Model 2
	zidan := Customer{
		Name:    "Zidan",
		Address: "Banyuwangi, East Java, Indonesia",
		Age:     25,
	}
	fmt.Println(zidan)

	//Model 3
	budi := Customer{"Budi", "Banyuwangi, East Java, Indonesia", 25}
	fmt.Println(budi)

	//Jalankan Struct Method
	//kevin := Customer{Name: "Kevin"}
	//kevin.sayHello()

	budi.sayHello("Zidan")
	zidan.sayHello("Zidan")
}
