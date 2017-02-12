package main

import (
    "fmt"
    "time"
    "math/rand"
    "sync"
)

var result int
var m sync.Mutex

func main()  {
    go runProcess("p1", 20)
    go runProcess("p2", 20)

    var s string
    fmt.Scanln(&s)
    fmt.Println("Final result:", result)
}

func runProcess(name string, total int)  {
    for i := 0; i < total; i++ {
	t := time.Duration(rand.Intn(255))
	time.Sleep(time.Millisecond * t)
	m.Lock()
	result++
	fmt.Println(name, " -> ", i, " partial result: ", result)
	m.Unlock()
    }
}