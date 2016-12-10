package main

import (
	"fmt"
)

func main()  {
    /*array:= [10]int{10,9,8,7,6,5,4,3,2,1}

    for i := 0; i < len(array); i++ {     
        if i == 2 {
            array[i] = 5
        }   
        fmt.Println(array[i])
    }*/


    var matriz [2][2]int
    matriz[0][1] = 10
    fmt.Print(matriz)
}
