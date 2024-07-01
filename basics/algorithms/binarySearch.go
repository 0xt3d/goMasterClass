package main

import (
	"fmt"
	"sort"
)

func binary(array []int, key int)int  {
	var high = len(array) - 1
	var low = 0

	for low <= high{
		mid := low + (high-low) /2
		if array[mid] == key{
			return mid
		}else if array[mid] > key {
			 high = mid - 1
		}else{
			low = mid +1
		}
	}
	return -1
}

func main()  {
	var array= []int{4,5,6,7,8,1,2,3,9}
	var key = 9
	sort.Ints(array)
	var result= binary(array, key)
	fmt.Println(result)
}
