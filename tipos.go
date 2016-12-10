package main

import (
	"fmt"
	"strconv"
)

func main()  {
    edad:="22"
    edad_int, _ := strconv.Atoi(edad)
    fmt.Print(10 + edad_int)
}

