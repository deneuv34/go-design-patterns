package main

import "fmt"

// PaintStruct struct type
type PaintStruct struct {
	color string
}

// CarStruct struct type
type CarStruct struct {
	wheelBrand string
	body       string
}

type Car struct{}

func (carType Car) create(c CarStruct) string {
	return "This car has " + c.wheelBrand + " and " + c.body
}

// PaintCar struct type
type PaintCar struct{}

// paint function
func (s PaintCar) paint(p PaintStruct) string {
	return ", has paint with " + p.color + " color"
}

// Accessories struct type
type Accessories struct{}

func (a Accessories) addLogo() string {
	return " and has logo"
}

// CreateCarFacade struct type
type CreateCarFacade struct {
}

// CreateSomeCar function
func (s CreateCarFacade) CreateBrandedCar() string {
	car := Car{}
	paintCar := PaintCar{}
	accessories := Accessories{}

	createdCar := car.create(CarStruct{body: "metal", wheelBrand: "swallow"})
	paintedCar := createdCar + paintCar.paint(PaintStruct{color: "blue"})
	finishAdded := paintedCar + accessories.addLogo()

	return finishAdded
}

// implementation
func main() {
	var facade CreateCarFacade
	fmt.Println(facade.CreateBrandedCar())
	return
}
