package main

import (
	"strings"
	"fmt"
	"time"
)

func main()  {
    go miNombreLento("jhon")
    go miNombreLento("chrysler")
    go miNombreLento("salguero")
    go miNombreLento("veizaga")
    fmt.Println("Que aburrido esperar este metodo")
    var wait string
    fmt.Scanln(&wait)
}


func miNombreLento(name string)  {
    letras := strings.Split(name, "")

    for _, value:= range(letras) {
        time.Sleep(1000 * time.Millisecond)
        fmt.Println(value)
    }
}