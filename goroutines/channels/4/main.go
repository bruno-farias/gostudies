package main

import (
    "fmt"
)

func main() {

    channel := make(chan int)

    go func() { //need this to avoid deadlocks
	channel <- 10
    }()

    fmt.Println(<-channel)

}