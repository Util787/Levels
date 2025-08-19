package main

import "fmt"

func main() {
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}
	fmt.Println(res(A, B))
}

func res(arrA []int, arrB []int) []int {
	res := make([]int, 0, 10) // я не знаю размеры массивов в тесткейсах поэтому предвыделил 10
	set := make(map[int]struct{})

	for _, val := range arrA {
		set[val] = struct{}{}
	}

	for _, val := range arrB {
		if _, ok := set[val]; ok {
			res = append(res, val)
		}
	}
	return res
}
