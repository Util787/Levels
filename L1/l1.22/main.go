package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1 << 30) // 2^30
	b := big.NewInt(1 << 30) // 2^30
	fmt.Println(Add(a, b))
	fmt.Println(Subtract(a, b))
	fmt.Println(Multiply(a, b))
	fmt.Println(Divide(a, b))
}

// сделал везде поинтеры заместо передачи по значению чтобы избежать копирования структуры

func Add(a, b *big.Int) *big.Int {
	res := new(big.Int)
	return res.Add(a, b)
}

func Subtract(a, b *big.Int) *big.Int {
	res := new(big.Int)
	return res.Sub(a, b)
}

func Multiply(a, b *big.Int) *big.Int {
	res := new(big.Int)
	return res.Mul(a, b)
}

func Divide(a, b *big.Int) *big.Int {
	res := new(big.Int)
	return res.Div(a, b)
}
