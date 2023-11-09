package factorymethod

import (
	"go-design-patterns/helpers"
	"testing"
)

// TestNewCar checks if NewCar returns a Car with the correct initialized name.
func TestNewCar(t *testing.T) {
	name := "Speedy"
	car := NewToy(car, name)
	if car.Name() != name {
		t.Errorf("Expected Car name to be %s, got %s", name, car.Name())
	}
}

func TestNewPlane(t *testing.T) {
	name := "Glider"
	plane := NewToy(plane, name)
	if plane.Name() != name {
		t.Errorf("Expected Plane name to be %s, got %s", name, plane.Name())
	}
}

func TestNewTrain(t *testing.T) {
	name := "ChooChoo"
	train := NewToy(train, name)
	if train.Name() != name {
		t.Errorf("Expected Train name to be %s, got %s", name, train.Name())
	}
}

func TestPlanePlay(t *testing.T) {
	plane := NewToy(plane, "Glider")
	output := helpers.CaptureOutput(func() {
		plane.Play()
	})

	expected := "Flying high!\n" // fmt.Println adds a newline.
	if output != expected {
		t.Errorf("Expected Plane Play output to be %q, got %q", expected, output)
	}
}
