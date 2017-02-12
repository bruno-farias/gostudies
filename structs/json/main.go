package main

import (
    "encoding/json"
    "fmt"
)

type Car struct {
    Name string //`json:"Model"` //altera o nome da propriedade exportada
    Year int //`json:"-"` // nao exporta a propriedade no json
    Color string
}

func main()  {

    /*car := Car{"Gol", 2017, "Amarelo"}

    result, _ := json.Marshal(car)

    fmt.Println(result)
    fmt.Println(string(result))*/

    var car Car
    data := []byte(`{"Name":"Gol", "Year": 2017, "Color": "Black"}`)
    json.Unmarshal(data, &car)

    fmt.Println(car.Name, car.Year, car.Color)

}
