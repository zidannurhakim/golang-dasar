package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	addressKe1 := Address{"Banyuwangi", "Jawa Timur", "Indonesia"}
	addressKe2 := &addressKe1 //Pointer

	addressKe2.City = "Jember"
	fmt.Println(addressKe1) //Berubah karena fitur pointer &
	fmt.Println(addressKe2) //Berubah

	//addressKe2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"} //Tidak akan berubah
	*addressKe2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(addressKe1)
	fmt.Println(addressKe2)
}
