package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan bool,1)
    ch1 <- true
    fmt.Println("tdata")

    //ch2 := make(chan bool)
    var ch2 chan bool

    go func() {
       ch2 <- true
       fmt.Println("send true in go function")
    }()

    select {
    case <- time.After(time.Second):
        fmt.Println("after time 1 second")
    case <- ch2:
        fmt.Println("from ch2...")
    }
}
