package main

import "fmt"

type vehicle interface {
    start() string
}

type Car struct {
    name string
}

type Motorcycle struct {
    name string
}

func (c Car) start() string  {
    return  "The car " + c.name + " has been started";
}

func (m Motorcycle) start() string  {
    return  "The Motorcycle " + m.name + " has been started";
}

func startVehicle(v vehicle) string  {
    return v.start();
}

func main()  {

    c := Car{"Gol"}
    m := Motorcycle{"Harley"}
    fmt.Println(startVehicle(c))
    fmt.Println(startVehicle(m))
    //fmt.Println(c.startCar())
    //fmt.Println(m.startMotorcycle())

}
