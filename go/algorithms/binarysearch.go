package main

import "fmt"

func binarySearch(item int, arr []int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (high + low) / 2

		if item < arr[mid] {
			fmt.Println("Muito baixo")
			high = mid - 1
		} else if item > arr[mid] {
			fmt.Println("Muito alto")
			low = mid + 1
		} else {
			fmt.Println("Achou")
			return mid
		}
	}
	return -1
}

func main() {
	arr := make([]int, 100)

	for i := range arr {
		arr[i] = i + 1
	}

	index := binarySearch(42, arr)

	if index == -1 {
		fmt.Println("não achou")
		return
	}

	fmt.Printf("index = %v", index)
}
