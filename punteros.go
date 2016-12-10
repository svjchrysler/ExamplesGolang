package main

import (
	"fmt"
)

func main()  {
    var x,y *int

    entero := 5
    
    x = &entero

    y = x

    fmt.Println(*x,*y)

    *x = 10

    fmt.Println(*x,*y)
}