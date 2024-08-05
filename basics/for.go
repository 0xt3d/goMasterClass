package main

import "fmt"

func main() {
	i := 0 //this is intialtion and declaration in one go
	var a = 10 int;
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j:=0; j<3; j++{
		fmt.Println(j)
	}

	for{
		fmt.Println("Hello")
		break
	}

	for n := range a {
		if n%2 == 0{
			continue
		}
		fmt.Println(n)
	}
}
