package main

import "fmt"

type Car struct {
    Name string
    Year int
    Color string
}

//herda de Car
type SuperCar struct {
    Car
    Name string //override
    CanFly bool
}

//metodo de car

func (c Car) info() string {
    return fmt.Sprintf("Car: %s\n Year: %d\n Color: %s\n", c.Name, c.Year, c.Color)
}

func main()  {

    car1 := Car{"Corolla", 2017, "White"}
    sCar := SuperCar{
	Car: Car{
	    "Fusca",
	    2030,
	    "Blue",
	},
	CanFly: true,
	Name: "Delorean", //override de propriedade
    }
    //car2 := Car{"BMW 320i", 2017, "Black"}
    //fmt.Println(car1.Name, car1.Year, car1.Color)
    //fmt.Println(car2.Name, car2.Year, car2.Color)

    fmt.Println(sCar.Name)
    fmt.Printf(car1.info())

}
