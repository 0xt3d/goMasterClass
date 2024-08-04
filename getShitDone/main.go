package main

import "fmt"

func display()  {
	fmt.Println("Display")
	
}

func add()  {
	fmt.Println("Add")
	
}

func remove()  {
	fmt.Println("remove")
	
}

func main()  {
	var action int
	fmt.Println("Get-Shit-Done List")
	fmt.Println("1. Display List")
	fmt.Println("2. Add Task")
	fmt.Println("3. Delete Task")
	fmt.Scan(&action)
	fmt.Println(action)

	if action == 1{
		display()
	}else if action == 2{
		add()
	}else if action == 3{
		remove()
	}
}