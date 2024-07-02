package main

import "fmt"

func linear(array []int, key int)int  {
	for i := range array{
		if array[i]==key{
			return i
		}
	}
	return -1
}


func main()  {
	
	var array = []int{1,2,3,4,5,6,7,8,9}
	var key = 5
	var result = linear(array, key)
	fmt.Println(result) 

}