package main

import (
	"fmt"
	"time"
)

type slyce []string

func (s slyce) output() {
	for index, value := range s {
		fmt.Println(index, value)
	}
}

type car struct {
	brand string
	price int
}

func changeCar(c car, newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

func (c car) changeCar1(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

func (c *car) realChangeCar(newBrand string, newPrice int) {
	(*c).brand = newBrand
	(*c).price = newPrice

}

func main() {
	const day = 24 * time.Hour
	seconds := day.Seconds()
	fmt.Printf("%T,\n %v\n", seconds, seconds)

	people := slyce{"Abass", "Tunji"}
	people.output()
	slyce.output(people)

	newCar := car{brand: "Innoson", price: 150000}
	fmt.Println(newCar)
	changeCar(newCar, "Toyota", 35000)
	fmt.Println(newCar) //doesnt change

	newCar.changeCar1("Mazda", 50000)
	fmt.Println(newCar) //doesnt change

	newCar.realChangeCar("Audi", 156000)
	fmt.Println(newCar) //changed

}
