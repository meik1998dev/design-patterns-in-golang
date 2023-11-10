package abstractfactory

import (
	"testing"
)

// TestHotWheelsCarCreation tests the creation of a HotWheels car.
func TestHotWheelsCarCreation(t *testing.T) {
	factory := HotWheelsFactory{}
	car := factory.MakeCar()
	if car == nil {
		t.Error("HotWheelsFactory was unable to make a Car")
	}
}

// TestHotWheelsPlaneCreation tests that HotWheelsFactory cannot create a Plane.
func TestHotWheelsPlaneCreation(t *testing.T) {
	factory := HotWheelsFactory{}
	plane := factory.MakePlane()
	if plane != nil {
		t.Error("HotWheelsFactory should not be able to make a Plane")
	}
}

// TestJetSetPlaneCreation tests the creation of a JetSet plane.
func TestJetSetPlaneCreation(t *testing.T) {
	factory := JetSetFactory{}
	plane := factory.MakePlane()
	if plane == nil {
		t.Error("JetSetFactory was unable to make a Plane")
	}
}

// TestJetSetCarCreation tests that JetSetFactory cannot create a Car.
func TestJetSetCarCreation(t *testing.T) {
	factory := JetSetFactory{}
	car := factory.MakeCar()
	if car != nil {
		t.Error("JetSetFactory should not be able to make a Car")
	}
}
