package main

import(
	"fmt"
)

func main()  {
	var  a = 1
	var  b = 2

	if a > b{
		fmt.Println("%a First Variable is greater", a)
	}else{
		fmt.Println("%b is greater", b)
	}
}