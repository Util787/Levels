package main

import "fmt"

func binSearch(arr []int, num int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == num {
			return mid
		} else if arr[mid] < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	fmt.Println(binSearch(arr, 0))

	fmt.Println(binSearch(arr, 17))
}
