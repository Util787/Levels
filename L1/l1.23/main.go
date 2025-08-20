package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(DeleteFromSlice(arr, 2))
}

// честно говоря я всегда просто свапал последний элемент с удаляемым но способ ниже позволяет сохранить порядок элементов
func DeleteFromSlice(sl []int, ind int) []int {
	copy(sl[ind:], sl[ind+1:]) // 1 2 4 5 5 то есть на место dest(3,4,5) скопировалось src(4,5) и из-за длины src последний элемент dest не затрагивается, он до сих пор 5
	return sl[:len(sl)-1]
}
