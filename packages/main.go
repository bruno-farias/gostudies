package main

import (
    "packages/car"
    "fmt"
)

func main()  {
    car := car.Car{
	"Gol",
	"Black",
    }
    fmt.Println(car.Start())
}
