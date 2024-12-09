package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "Validation Error"}
	}

	if id != "Zidan" {
		return &notFoundError{Message: "Data Not Found"}
	}
	//OK
	return nil
}

func main() {
	err := SaveData("Eko", nil)
	if err != nil {
		//Terjadi Error
		//Contoh 1
		//if validationErr, ok := err.(*validationError); ok {
		//	fmt.Println("Validation Error:", validationErr.Error())
		//} else if notFoundErr, ok := err.(*notFoundError); ok {
		//	fmt.Println("Not Found Error:", notFoundErr.Error())
		//} else {
		//	fmt.Println("Unknown Error:", err.Error())
		//}
		//Contoh 2 (Switch Case)
		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("Validation Error:", finalError.Message)
		case *notFoundError:
			fmt.Println("Not Found Error:", finalError.Message)
		default:
			fmt.Println("Unknown error:", finalError.Error())
		}
	} else {
		//Sukses
		fmt.Println("Sukses")
	}
}
