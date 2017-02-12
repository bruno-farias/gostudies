package main

import (
    "fmt"
    "time"
    "math/rand"
)

func main()  {
    go runProcess("p1", 20)
    go runProcess("p2", 20)

    var s string
    fmt.Scanln(&s)
}

func runProcess(name string, total int)  {
    for i := 0; i < total; i++ {
	fmt.Println(name, " -> ", i)
	t := time.Duration(rand.Intn(255))
	time.Sleep(time.Millisecond * t)
    }
}