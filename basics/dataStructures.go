package main

import "fmt"

func main()  {
	var array = [5]int{1,2,3,4,5}

	fmt.Println(array[0])//Prints 1
	fmt.Println(array[4])//Prints 5
	for i :=range array {
		fmt.Println(array[i]) //Print index values on new line
		//1
		//2
		//3
		//4
		//5
	}

	array[4] = 100
	fmt.Println("Updating Index of 4:", array)
	fmt.Println("getting lenth of array using len(array) function", len(array))
	

	
	





}