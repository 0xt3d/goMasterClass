package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}

func variables(){
	var integerVar int = 42
    var floatVar float64 = 3.14
    var stringVar string = "Hello, Golang!"
    var boolVar bool = true

    fmt.Println(integerVar, floatVar, stringVar, boolVar)
}