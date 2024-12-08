package main

import "fmt"

func main() {
    rows := 5
    for i := 1; i <= rows; i++ {
        for j := i; j < rows; j++ {
            fmt.Print(" ")
        }
        for k := 1; k <= (2*i - 1); k++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
    for i := rows - 1; i >= 1; i-- {
        for j := rows; j > i; j-- {
            fmt.Print(" ")
        }
        for k := 1; k <= (2*i - 1); k++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
}
