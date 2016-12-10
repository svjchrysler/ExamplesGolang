package main

import "gopkg.in/zabawaba99/firego.v1"
import "log"


func main()  {
    v := map[string]string{"food": "message"}
    fb := firego.New("https://android-chat-f55e4.firebaseio.com/", nil)
    //fbresult, _ := fb.Push(v)
    
    if err := fb.Set(v); err != nil {
        log.Fatal(err)
    }

    log.Fatal("ddd")

}