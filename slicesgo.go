package main

import (
	"fmt"
)

func main()  {
    var  matriz []int
    //matriz := []int{1,2,3}
    fmt.Println(matriz)
    
    fmt.Printf("capacidad: %d longitud: %d \n", cap(matriz), len(matriz) )

    matriz = append(matriz, 4)
    matriz = append(matriz, 5)
    matriz = append(matriz, 6)
    matriz = append(matriz, 7)
    matriz = append(matriz, 8)
    matriz = append(matriz, 9)
    matriz = append(matriz, 10)
    matriz = append(matriz, 11)
    matriz = append(matriz, 12)
    matriz = append(matriz, 13)
    
    fmt.Printf("capacidad: %d longitud: %d \n", cap(matriz), len(matriz) )

    fmt.Println(matriz)
}