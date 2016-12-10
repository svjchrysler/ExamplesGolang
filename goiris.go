package main

import (
	"github.com/kataras/iris"
)

func main()  {
    
    iris.Get("/", func (ctx *iris.Context) {
        err := ctx.Text(200, "Hola")
        if err != nil {
            
        }
    })
    
    iris.Listen("localhost:8080")

}