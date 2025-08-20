package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	arrLength := len(arr)

	if arrLength < 2 {
		return arr
	}

	pivotIndex := arrLength / 2
	pivot := arr[pivotIndex]

	var left []int
	var right []int

	for i, value := range arr {
		if i == pivotIndex { // надо скипнуть pivot иначе он войдет в рекурсии и будут дубликаты
			continue
		}
		if value <= pivot {
			left = append(left, value)
		} else {
			right = append(right, value)
		}
	}

	sortedLeft := quickSort(left)
	sortedRight := quickSort(right)

	result := append(sortedLeft, pivot)
	result = append(result, sortedRight...)

	return result
}

func main() {
	arr := []int{111, 43, 67, 43, 10, 5, 2, 3, 8, 6, 7, 4, 9, 1, 5, 7, 2, 5, 87, 3, 5}
	fmt.Println("orig:", arr)

	res := quickSort(arr)
	fmt.Println("res:  ", res)
}
