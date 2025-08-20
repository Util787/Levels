package main

import "fmt"

func reverse(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-1-i] = runes[len(runes)-1-i], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println(reverse("ABC"))
	fmt.Println(reverse("ABCD"))
	fmt.Println(reverse("главрыба"))
}
