package main

import (
	"fmt"
	"bufio"
	"os"
)

func main()  {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Ingresa tu nombre")

    text, err := reader.ReadString('\n')

    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("Hola %s", text)
    }
}