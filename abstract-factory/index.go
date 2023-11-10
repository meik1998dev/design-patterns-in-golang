package abstractfactory

import "fmt"

type Toy interface {
	Play()
}

type Car interface {
	Toy
	Drive()
}

type Plane interface {
	Toy
	Fly()
}

type ToyFactory interface {
	MakeCar() Car
	MakePlane() Plane
}

type HotWheelsFactory struct{}

// HotWheelsCar factory
type HotWheelsCar struct{}

func (f HotWheelsFactory) MakeCar() Car {
	return &HotWheelsCar{}
}

func (f HotWheelsFactory) MakePlane() Plane {
	return nil
}

func (f HotWheelsCar) Play() {
	fmt.Printf("Zooming around!\n")
}

func (f HotWheelsCar) Drive() {
	fmt.Printf("Driving!\n")
}

// JetSetFactory factory
type JetSetFactory struct{}

func (f JetSetFactory) MakePlane() Plane {
	return &JetSetFactory{}
}

func (f JetSetFactory) MakeCar() Car {
	return nil
}

func (f JetSetFactory) Play() {
	fmt.Printf("plying high!\n")
}

func (f JetSetFactory) Fly() {
	fmt.Printf("Flying high!\n")
}
