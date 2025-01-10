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

func basicDataType(){
    // Boolean
    var isActive bool = true
    fmt.Println("Boolean:", isActive)

    // Integer Types
    var age int = 25
    var height uint = 180 // unsigned integer
    fmt.Println("Integer:", age)
    fmt.Println("Unsigned Integer:", height)

    // Floating Point Numbers
    var temperature float32 = 36.6
    var pi float64 = 3.141592653589793
    fmt.Println("Float32:", temperature)
    fmt.Println("Float64:", pi)

    // Complex Numbers
    var complexNum complex64 = 1 + 2i
    fmt.Println("Complex Number:", complexNum)

    // String
    var greeting string = "Hello, Golang!"
    fmt.Println("String:", greeting)

    // Array
    var days [7]string = [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
    fmt.Println("Array of Days:", days)

    // Slice
    fruits := []string{"Apple", "Banana", "Cherry"}
    fmt.Println("Slice of Fruits:", fruits)

    // Map
    capitals := map[string]string{
        "India":   "New Delhi",
        "USA":     "Washington D.C.",
        "France":  "Paris",
        "Japan":   "Tokyo",
    }
    
    fmt.Println("Map of Capitals:")
    for country, capital := range capitals {
        fmt.Printf("%s: %s\n", country, capital)
    }
}


