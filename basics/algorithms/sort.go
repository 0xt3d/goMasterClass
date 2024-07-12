package main

import (
	"fmt"
)

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(result, right...)
		}
		if len(right) == 0 {
			return append(result, left...)
		}
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	return result
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[len(arr)/2]
	left := make([]int, 0)
	right := make([]int, 0)
	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		}
	}
	sortedLeft := quickSort(left)
	sortedRight := quickSort(right)
	return append(append(sortedLeft, pivot), sortedRight...)
}

func main() {
	arr := []int{34, 21, 65, 12, 98, 32, 76, 45, 89, 23}
	fmt.Println("Original array:", arr)

	bubbleSort(arr)
	fmt.Println("Bubble sort:", arr)

	arr = []int{34, 21, 65, 12, 98, 32, 76, 45, 89, 23}
	selectionSort(arr)
	fmt.Println("Selection sort:", arr)

	arr = []int{34, 21, 65, 12, 98, 32, 76, 45, 89, 23}
	insertionSort(arr)
	fmt.Println("Insertion sort:", arr)

	arr = []int{34, 21, 65, 12, 98, 32, 76, 45, 89, 23}
	arr = mergeSort(arr)
	fmt.Println("Merge sort:", arr)

	arr = []int{34, 21, 65, 12, 98, 32, 76, 45, 89, 23}

}