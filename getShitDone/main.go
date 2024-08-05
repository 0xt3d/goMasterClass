package main

import (
	"fmt"
// "golang.org/x/exp/slices"
)

var task []string

func display()  {
	fmt.Println("Display")
	for i := range task{
		fmt.Println("-----------------------------")
		fmt.Println(task[i])
		fmt.Println("-----------------------------")
	}
	
}

func add(item string)  {
	fmt.Println("Add")
	task = append(task, item);
	
}

func remove()  {
	fmt.Println("remove")
	
}

func main()  {

	var item string;
	var action int;
	fmt.Println("Get-Shit-Done List")
	fmt.Println("1. Display List")
	fmt.Println("2. Add Task")
	fmt.Println("3. Delete Task")
	fmt.Println("4. Exit")
	fmt.Scan(&action)
	fmt.Println(action)
	for{
	if action == 1{
		display()
	}else if action == 2{
		fmt.Scan(&item);
		add(item)
		display()
	}else if action == 3{
		remove()
	}else if action ==4{
		break
	}

	}
}
