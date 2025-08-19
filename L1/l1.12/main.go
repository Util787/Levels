package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(set(arr))
}

func set(arr []string) []string {
	res := make([]string, 0, 10) // я не знаю размеры массивов в тесткейсах поэтому предвыделил 10
	m := make(map[string]struct{})

	for _, val := range arr {
		m[val] = struct{}{} // использую мапу как сет
	}

	for k := range m {
		res = append(res, k)
	}

	return res
}
