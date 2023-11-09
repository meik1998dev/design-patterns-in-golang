package factorymethod

import "fmt"

type Toy interface {
	Name() string
	Play()
}

// Plan
type Plane struct {
	name string
}

func NewPlane(name string) Toy {
	return &Plane{name: name}
}

func (p *Plane) Name() string {
	return p.name
}

func (p *Plane) Play() {
	fmt.Println("Flying high!")
}

// Car
type Car struct {
	name string
}

func NewCar(name string) Toy {
	return &Car{name: name}
}

func (p *Car) Name() string {
	return p.name
}

func (p *Car) Play() {
	fmt.Println("Zooming around!")
}

// Train
type Train struct {
	name string
}

func NewTrain(name string) Toy {
	return &Train{name: name}
}

func (p *Train) Name() string {
	return p.name
}

func (p *Train) Play() {
	fmt.Println("Chugging along the tracks!")
}
