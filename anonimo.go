package main

import (
	"fmt"
)

type Human struct {
    name string
    lastname string
}

func (this Human) hablar() string {
    return "Bla bla bla"
}

type Tutor struct {
    Human
    profesion string
}

func main()  {
    tutor := Tutor{Human{"Jhon", "Salguero"}, "programador"}
    fmt.Println(tutor.hablar())
}