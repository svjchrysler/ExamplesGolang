package main

import (
	"gopkg.in/mgo.v2"
	"log"
    "gopkg.in/mgo.v2/bson"
	"fmt"
)

type Person struct {
    Name string
    Phone string
}

func main()  {
    session, err := mgo.Dial("localhost:27017")
    if err != nil {
        panic(err)
    }

    defer session.Close()

    session.SetMode(mgo.Monotonic, true)

    c:=session.DB("gomongo").C("people")
    err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
	               &Person{"Cla", "+55 53 8402 8510"})

    if err != nil {
        log.Fatal(err)
    }

    result := Person{}
    err = c.Find(bson.M{"name": "Ale"}).One(&result)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("phone:", result.Phone)
    
}