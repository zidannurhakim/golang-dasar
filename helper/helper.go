package helper

import "fmt"

var version = "1.0.0"
var Application = "Golang"

func sayGoodBye(name string) string {
	return "Good Bye " + name
}

func SayHello(name string) string {
	return "Hello " + name

}

func Bye() {
	sayGoodBye("Zidan")
	fmt.Println(version)
}
