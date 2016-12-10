package main

import (
	"fmt"
)

type User interface {
    Permisos() int // 1 - 5
    Nombre() string
}

type Admin struct {
    name string
}

func (this Admin) Permisos() int {
    return 5
}

func (this Admin) Nombre() string {
    return this.name
}

type Editor struct {
    name string
}

func (this Editor) Permisos() int {
    return 3
}

func (this Editor) Nombre() string {
    return this.name
}

func auth(user User) string {
    if user.Permisos() == 4 {
        return user.Nombre() + " Tiene permisos de Administrador"
    } else if user.Permisos() == 3 {
        return user.Nombre() + " Tiene permisos de Editor"
    }
    return user.Nombre() + " Otro Permiso"
}

func main()  {
    //admin := Admin{"jhon salguero"}
    editor := Editor{"luis mercado"}
       
    fmt.Println(auth(editor))   
}