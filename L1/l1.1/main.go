package main

import "fmt"

type Human struct {
	Name string
}

func (h *Human) SayMyName() {
	fmt.Println(h.Name)
}

func (h *Human) SayHi() {
	fmt.Println("hi")
}

type Action struct {
	Human
}

func main() {
	a := Action{
		Human: Human{
			Name: "Walter",
		},
	}

	a.SayMyName() // т.к. поле не именовано можно вызывать напрямую, будь у поля Human название пришлось бы вызывать через имя поля (уже не было бы встраиванием) ну и еще я не мог бы написать fmt.Println(a.Name) т.к. поле не встроилось бы и пришлось тоже через название поля брать
	a.SayHi()
}
