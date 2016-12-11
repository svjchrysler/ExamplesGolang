package main

import (
	"os"
	"fmt"
	"bufio"
)

func main()  {
    ejecucion := readFile()
    fmt.Println(ejecucion)
}

func readFile() bool {
    archivo, err := os.Open("./txt.txt")   
    defer func ()  {
      archivo.Close()
      fmt.Println("Defer")  
      recover()
    }()
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(archivo)

    for scanner.Scan() {
        linea := scanner.Text()

        fmt.Println(linea)
    }

    if true {
        return true
    }

    fmt.Println("Nunca me ejecuto")

    return true
    
}