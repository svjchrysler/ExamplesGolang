package main

import (
	"fmt"
	"os"
	"bufio"
)

func main()  {
    channel := make(chan string)

    go func (channel chan string)  {
        for{
            reader := bufio.NewReader(os.Stdin)
            fmt.Println("Ingresa tu nombre porfavor..!")
            name, _ := reader.ReadString('\n')
            channel <- name
        }
    }(channel)

    msg := <- channel

    fmt.Println(msg)

    msg = <- channel

    fmt.Println(msg)
}