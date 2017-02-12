package main

import "fmt"

func main() {

    channel := make(chan int)
    ok := make(chan bool)

    go func() {
	for i := 0; i < 10; i++ {
	    channel <- i
	}

	ok <- true //inform semaphore

    }()

    go func() {
	for i := 0; i < 10; i++ {
	    channel <- i
	}

	ok <- true //inform semaphore

    }()

    go func() {
	<- ok //release semaphore
	<- ok //release semaphore
	close(channel) // close channel
    }()

    for number := range channel  {
	fmt.Println(number)
    }

}
