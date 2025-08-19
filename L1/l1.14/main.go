package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(findType(42))
	fmt.Println(findType("hello"))
	fmt.Println(findType(true))
	fmt.Println(findType(make(chan int)))
	fmt.Println(findType(make(chan string)))
	fmt.Println(findType(make(<-chan string)))
	fmt.Println(findType(make(chan<- string)))
	fmt.Println(findType([]int{1, 2, 3}))
}

func findType(val any) string {
	switch val.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	default:
		v := reflect.ValueOf(val) // заместо того чтобы прописывать все виды каналов сделал так
		if v.Kind() == reflect.Chan {
			return "chan"
		}
		return "unknown"
	}
}
