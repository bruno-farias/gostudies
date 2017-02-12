package main

import (
    "fmt"
    "time"
    "math/rand"
    "sync"
    "runtime"
)

var waitingGroup sync.WaitGroup

func main() {
    fmt.Println(runtime.NumCPU())
    waitingGroup.Add(2)
    go runProcess("p1", 20)
    go runProcess("p2", 20)
    waitingGroup.Wait()
}

func runProcess(name string, total int) {
    for i := 0; i < total; i++ {
	fmt.Println(name, " -> ", i)
	t := time.Duration(rand.Intn(255))
	time.Sleep(time.Millisecond * t)
    }
    waitingGroup.Done()
}