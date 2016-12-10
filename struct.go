package main

import (
	"fmt"
	"strconv"
)

type User struct {
    edad int
    nombre string
    apellido string   
}

func (u *User) getFullInfo() string {
    return u.nombre + " " + u.apellido + " " + strconv.Itoa(u.edad)
}

func main()  {
    //user := User{edad:10,nombre:"Jhon",apellido:"salguero"}

    user := User{edad: 22, nombre: "Jhon", apellido: "Salguero"}

    fmt.Print(user.getFullInfo())
}