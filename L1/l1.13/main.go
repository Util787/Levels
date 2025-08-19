package main

import "fmt"

func main() {
	a := 5
	b := 30
	fmt.Println(a)
	fmt.Println(b)

	a = b + a
	b = a - b
	a = a - b

	fmt.Println(a)
	fmt.Println(b)
}
